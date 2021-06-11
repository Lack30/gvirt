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

package gvirt

import (
	"net/url"
	"strings"
	"time"
)

var (
	// DefaultPoolSize sets the connection pool size
	DefaultPoolSize = 100
	// DefaultPoolTTL sets the connection pool ttl
	DefaultPoolTTL = time.Minute
	// DefaultPoolMaxStreams maximum streams on a connections (20)
	DefaultPoolMaxStreams = 20
	// DefaultPoolMaxIdle maximum idle conns of a pool (50)
	DefaultPoolMaxIdle = 50
)

type ClientKind string

const (
	KindQemu ClientKind = "qemu"
)

type Options struct {
	// the url of libvirt api
	u url.URL

	Kind ClientKind
	// Connection pool
	PoolSize       int
	PoolTTL        time.Duration
	PoolMaxStreams int
	PoolMaxIdle    int
}

func newOptions(opts ...Option) Options {
	options := Options{
		u: url.URL{
			Path: "/system",
		},
		Kind:           KindQemu,
		PoolSize:       DefaultPoolSize,
		PoolTTL:        DefaultPoolTTL,
		PoolMaxStreams: DefaultPoolMaxStreams,
		PoolMaxIdle:    DefaultPoolMaxIdle,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

// Kind specifies the Kind for Options
func Kind(kind ClientKind) Option {
	return func(o *Options) {
		o.Kind = kind
	}
}

// URL specifies the value for url.URL
func URL(u url.URL) Option {
	return func(o *Options) {
		if strings.HasPrefix(u.String(), string(KindQemu)) {
			o.Kind = KindQemu
		}
		o.u = u
	}
}

// Addr specifies the Schema and Host for url.URL
func Addr(schema, host string) Option {
	return func(o *Options) {
		switch o.Kind {
		case KindQemu:
			if !strings.HasPrefix(schema, "qemu") {
				schema = "qemu+" + schema
			}
		}
		o.u.Scheme = schema
		o.u.Host = host
	}
}

// Auth specifies the authorization information for url.URL
func Auth(username, password string) Option {
	return func(o *Options) {
		var info *url.Userinfo
		if len(password) != 0 {
			info = url.UserPassword(username, password)
		} else {
			info = url.User(username)
		}
		o.u.User = info
	}
}

// PoolSize sets the connection pool size
func PoolSize(d int) Option {
	return func(o *Options) {
		o.PoolSize = d
	}
}

// PoolMaxStreams sets maximum streams on a connections
func PoolMaxStreams(d int) Option {
	return func(o *Options) {
		o.PoolMaxStreams = d
	}
}

// PoolMaxIdle sets maximum idle conns of a pool
func PoolMaxIdle(d int) Option {
	return func(o *Options) {
		o.PoolMaxIdle = d
	}
}

// PoolTTL sets the connection pool ttl
func PoolTTL(d time.Duration) Option {
	return func(o *Options) {
		o.PoolTTL = d
	}
}
