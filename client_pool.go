// MIT License
//
// Copyright (c) 2021 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package gkvm

import (
	"net/url"
	"sync"
	"time"

	"libvirt.org/libvirt-go"
)

type pool struct {
	size int
	ttl  int64

	// max streams on a *poolConn
	maxStreams int
	// max idle conns
	maxIdle int

	sync.Mutex
	conns map[string]*streamsPool
}

type streamsPool struct {
	// head of list
	head *poolConn
	// busy conns list
	busy *poolConn
	// the size of list
	count int
	// idle conn
	idle int
}

type poolConn struct {
	*libvirt.Connect
	u   url.URL
	err error

	// pool and streams pool
	pool    *pool
	sp      *streamsPool
	streams int
	created int64

	// list
	pre  *poolConn
	next *poolConn
	in   bool
}

func newPool(size int, ttl time.Duration, idle int, ms int) *pool {
	if ms <= 0 {
		ms = 1
	}
	if idle < 0 {
		idle = 0
	}
	return &pool{
		size:       size,
		ttl:        int64(ttl.Seconds()),
		maxStreams: ms,
		maxIdle:    idle,
		conns:      make(map[string]*streamsPool),
	}
}

func (p *pool) getConn(u url.URL) (*poolConn, error) {
	uu := u.String()
	now := time.Now().Unix()
	p.Lock()
	sp, ok := p.conns[uu]
	if !ok {
		sp = &streamsPool{head: &poolConn{}, busy: &poolConn{}, count: 0, idle: 0}
		p.conns[uu] = sp
	}
	// while we have conn check streams and then return one
	// otherwise we'll create a new conn
	conn := sp.head.next
	for conn != nil {
		alived, err := conn.IsAlive()
		if err != nil || !alived {
			next := conn.next
			if conn.streams == 0 {
				removeConn(conn)
				sp.idle--
			}
			conn = next
			continue
		}

		// a old conn
		if now-conn.created > p.ttl {
			next := conn.next
			if conn.streams == 0 {
				removeConn(conn)
				conn.Connect.Close()
				sp.idle--
			}
			conn = next
			continue
		}
		// a busy conn
		if conn.streams >= p.maxStreams {
			next := conn.next
			removeConn(conn)
			addConnAfter(conn, sp.busy)
			conn = next
			continue
		}
		// a idle conn
		if conn.streams == 0 {
			sp.idle--
		}
		// a good conn
		conn.streams++
		p.Unlock()
		return conn, nil
	}
	p.Unlock()

	// create new conn
	cc, err := libvirt.NewConnect(uu)
	if err != nil {
		return nil, err
	}
	conn = &poolConn{cc, u, nil, p, sp, 1, time.Now().Unix(), nil, nil, false}

	// add conn to streams pool
	p.Lock()
	if sp.count < p.size {
		addConnAfter(conn, sp.head)
	}
	p.Unlock()

	return conn, nil
}

func (p *pool) release(u url.URL, conn *poolConn, err error) {
	p.Lock()
	p, sp, created := conn.pool, conn.sp, conn.created
	// try to add conn
	if !conn.in && sp.count < p.size {
		addConnAfter(conn, sp.head)
	}
	if !conn.in {
		p.Unlock()
		conn.Connect.Close()
		return
	}
	// a busy conn
	if conn.streams >= p.maxStreams {
		removeConn(conn)
		addConnAfter(conn, sp.head)
	}
	conn.streams--
	// if streams == 0, we can do something
	if conn.streams == 0 {
		// 1. it has errored
		// 2. too many idle conn or
		// 3. conn is too old
		now := time.Now().Unix()
		if err != nil || sp.idle >= p.maxIdle || now-created > p.ttl {
			removeConn(conn)
			p.Unlock()
			conn.Connect.Close()
			return
		}
		sp.idle++
	}
	p.Unlock()
	return
}

func (p *pool) destroy() {
	p.Lock()
	defer p.Unlock()
	for _, sp := range p.conns {
		conn := sp.head.next
		for conn != nil {
			next := conn.next
			removeConn(conn)
			conn = next
		}
	}
	p.size = 0
}

func (c *poolConn) Close() {
	c.pool.release(c.u, c, c.err)
}

func removeConn(conn *poolConn) {
	if conn.pre != nil {
		conn.pre.next = conn.next
	}
	if conn.next != nil {
		conn.next.pre = conn.pre
	}
	conn.pre = nil
	conn.next = nil
	conn.in = false
	conn.sp.count--
	return
}

func addConnAfter(conn *poolConn, after *poolConn) {
	conn.next = after.next
	conn.pre = after
	if after.next != nil {
		after.next.pre = conn
	}
	after.next = conn
	conn.in = true
	conn.sp.count++
	return
}
