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
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
)

type Client struct {
	options Options

	pool *pool
	once atomic.Value

	// the version of libvirt
	libVersion string
	// the version of qemu
	qemuVersion string
}

func NewClient(opts ...Option) (*Client, error) {
	options := newOptions(opts...)

	p := newPool(options.PoolSize, options.PoolTTL, options.PoolMaxStreams, options.PoolMaxIdle)

	cc, err := p.getConn(options.u)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	v1, _ := cc.GetLibVersion()
	v2, _ := cc.GetVersion()

	rc := &Client{
		options:     options,
		pool:        p,
		libVersion:  parseVersion(v1),
		qemuVersion: parseVersion(v2),
	}
	rc.once.Store(true)

	return rc, nil
}

func (c *Client) NewSession() (*poolConn, error) {
	return c.pool.getConn(c.options.u)
}

func (c *Client) Version() (string, string) {
	return c.libVersion, c.qemuVersion
}

func (c *Client) Close() error {
	if c.once.Load().(bool) {
		c.pool.destroy()
		return nil
	}
	return fmt.Errorf("client no initialized")
}

func parseVersion(n uint32) string {
	sp := make([]string, 0)
	for n > 1 {
		sp = append([]string{strconv.FormatUint(uint64(n%1000), 10)}, sp...)
		n = n / 1000
	}
	return strings.Join(sp, ".")
}
