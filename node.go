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

type NodeDevice struct {
	cc *Client

	ptr *libvirt.NodeDevice

	spec.NodeDevice
}

func (c *Client) GetAllNodeDevices() ([]*NodeDevice, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	devs, err := cc.ListAllNodeDevices(libvirt.CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_NET |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_SCSI |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_STORAGE |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_VPORTS |
		libvirt.CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC)
	if err != nil {
		return nil, err
	}
	outs := make([]*NodeDevice, 0, len(devs))
	for i, dev := range devs {
		doc, err := dev.GetXMLDesc(0)
		if err != nil {
			continue
		}
		p := &NodeDevice{cc: c, ptr: &devs[i]}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetNodeDeviceByName(name string) (*NodeDevice, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dev, err := cc.LookupDeviceByName(name)
	if err != nil {
		return nil, err
	}
	doc, err := dev.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &NodeDevice{cc: c, ptr: dev}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetNodeDeviceByWWN(wwnn, wwpn string) (*NodeDevice, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dev, err := cc.LookupDeviceSCSIHostByWWN(wwnn, wwpn, 0)
	if err != nil {
		return nil, err
	}
	doc, err := dev.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &NodeDevice{cc: c, ptr: dev}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) NodeDeviceCreateXML(xml string) (*NodeDevice, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.DeviceCreateXML(xml, 0)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &NodeDevice{cc: c, ptr: net}
	err = out.UnmarshalX(doc)
	return out, err
}

func (d *NodeDevice) Destroy() error {
	return d.ptr.Destroy()
}

func (d *NodeDevice) Detach() error {
	return d.ptr.Detach()
}

func (d *NodeDevice) Deref() *libvirt.NodeDevice {
	return d.ptr
}
