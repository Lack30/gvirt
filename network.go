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

func (c *Client) GetAllNetworks() ([]*Network, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	nets, err := cc.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE | libvirt.CONNECT_LIST_NETWORKS_INACTIVE)
	if err != nil {
		return nil, err
	}
	outs := make([]*Network, 0, len(nets))
	for i, net := range nets {
		doc, err := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
		if err != nil {
			continue
		}
		p := &Network{cc: c, ptr: &nets[i]}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetNetworkByName(name string) (*Network, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupNetworkByName(name)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &Network{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetNetworkByUUID(uuid string) (*Network, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupNetworkByUUIDString(uuid)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &Network{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) NetworkCreateXML(xml string) (*Network, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.NetworkCreateXML(xml)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := &Network{cc: c, ptr: net}
	err = out.UnmarshalX(doc)
	return out, err
}

func (c *Client) NetworkDefineXML(xml string) (*Network, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.NetworkDefineXML(xml)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(libvirt.NETWORK_XML_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := &Network{cc: c, ptr: net}
	err = out.UnmarshalX(doc)
	return out, err
}

type Network struct {
	cc *Client

	ptr *libvirt.Network

	spec.Network
}

func (n *Network) Deref() *libvirt.Network {
	return n.ptr
}
