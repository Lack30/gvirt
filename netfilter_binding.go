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

type NWFilterBinding struct {
	cc *Client

	ptr *libvirt.NWFilterBinding

	spec.NWFilterBinding
}

func (c *Client) GetAllNWFilterBindings() ([]*NWFilterBinding, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	bindings, err := cc.ListAllNWFilterBindings(0)
	if err != nil {
		return nil, err
	}
	outs := make([]*NWFilterBinding, 0, len(bindings))
	for i, dev := range bindings {
		doc, err := dev.GetXMLDesc(0)
		if err != nil {
			continue
		}
		p := &NWFilterBinding{cc: c, ptr: &bindings[i]}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetNWFilterBindingByName(name string) (*NWFilterBinding, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	binding, err := cc.LookupNWFilterBindingByPortDev(name)
	if err != nil {
		return nil, err
	}
	doc, err := binding.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &NWFilterBinding{cc: c, ptr: binding}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) NWFilterBindingCreateXML(xml string) (*NWFilterBinding, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	binding, err := cc.NWFilterBindingCreateXML(xml, 0)
	if err != nil {
		return nil, err
	}
	doc, err := binding.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &NWFilterBinding{cc: c, ptr: binding}
	err = out.UnmarshalX(doc)
	return out, err
}

func (binding *NWFilterBinding) Delete() error {
	return binding.ptr.Delete()
}

func (binding *NWFilterBinding) Deref() *libvirt.NWFilterBinding {
	return binding.ptr
}
