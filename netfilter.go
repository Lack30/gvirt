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
	"github.com/lack-io/gvirt/spec"
	"libvirt.org/libvirt-go"
)

type NWFilter struct {
	cc *Client

	ptr *libvirt.NWFilter

	spec.NWFilter
}

func (c *Client) GetAllNWFilters() ([]*NWFilter, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	filters, err := cc.ListAllNWFilters(0)
	if err != nil {
		return nil, err
	}
	outs := make([]*NWFilter, 0, len(filters))
	for i, dev := range filters {
		doc, err := dev.GetXMLDesc(0)
		if err != nil {
			continue
		}
		p := &NWFilter{cc: c, ptr: &filters[i]}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetNWFilterByName(name string) (*NWFilter, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	filter, err := cc.LookupNWFilterByName(name)
	if err != nil {
		return nil, err
	}
	doc, err := filter.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &NWFilter{cc: c, ptr: filter}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetNWFilterByUUID(uuid string) (*NWFilter, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	filter, err := cc.LookupNWFilterByUUIDString(uuid)
	if err != nil {
		return nil, err
	}
	doc, err := filter.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &NWFilter{cc: c, ptr: filter}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) NWFilterDefineXML(xml string) (*NWFilter, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	filter, err := cc.NWFilterDefineXML(xml)
	if err != nil {
		return nil, err
	}
	doc, err := filter.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &NWFilter{cc: c, ptr: filter}
	err = out.UnmarshalX(doc)
	return out, err
}

func (f *NWFilter) UnDefine() error {
	return f.ptr.Undefine()
}

func (f *NWFilter) Deref() *libvirt.NWFilter {
	return f.ptr
}
