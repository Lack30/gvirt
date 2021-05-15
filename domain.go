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
	"github.com/lack-io/gvirt/spec"
	"libvirt.org/libvirt-go"
)

func (c *Client) GetDomains() ([]*spec.Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	domains, err := cc.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := make([]*spec.Domain, 0, len(domains))
	for _, dom := range domains {
		doc, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE | libvirt.DOMAIN_XML_INACTIVE)
		if err != nil {
			continue
		}
		d := &spec.Domain{}
		if err := d.UnmarshalX(doc); err == nil {
			out = append(out, d)
		}
	}
	return out, nil
}
