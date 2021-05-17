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

type Interface struct {
	cc *Client

	ptr *libvirt.Interface

	spec.Interface
}

func (c *Client) GetAllInterfaces() ([]*Interface, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	interfaces, err := cc.ListAllInterfaces(libvirt.CONNECT_LIST_INTERFACES_ACTIVE | libvirt.CONNECT_LIST_INTERFACES_INACTIVE)
	if err != nil {
		return nil, err
	}
	outs := make([]*Interface, 0, len(interfaces))
	for _, in := range interfaces {
		doc, err := in.GetXMLDesc(libvirt.INTERFACE_XML_INACTIVE)
		if err != nil {
			continue
		}
		p := &Interface{cc: c, ptr: &in}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetInterfaceByName(name string) (*Interface, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupInterfaceByName(name)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.INTERFACE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &Interface{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetInterfaceByMAC(mac string) (*Interface, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupInterfaceByMACString(mac)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.INTERFACE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &Interface{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) InterfaceDefineXML(xml string) (*Interface, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.InterfaceDefineXML(xml, 0)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.INTERFACE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := &Interface{cc: c, ptr: net}
	err = out.UnmarshalX(doc)
	return out, err
}

func (i *Interface) Create() error {
	return i.ptr.Create(0)
}

func (i *Interface) UnDefine() error {
	return i.ptr.Undefine()
}
