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
	"fmt"

	"github.com/lack-io/gvirt/spec"
	"libvirt.org/libvirt-go"
)

func (c *Client) GetAllDomains() ([]*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	domains, err := cc.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := make([]*Domain, 0, len(domains))
	for _, dom := range domains {
		doc, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE | libvirt.DOMAIN_XML_INACTIVE)
		if err != nil {
			continue
		}
		d := &Domain{cc: c, ptr: &dom}
		if err := d.UnmarshalX(doc); err == nil {
			out = append(out, d)
		}
	}
	return out, nil
}

func (c *Client) GetDomainById(id uint32) (*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dom, err := cc.LookupDomainById(id)
	if err != nil {
		return nil, err
	}
	xml, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &Domain{cc: c, ptr: dom}
	err = out.UnmarshalX(xml)
	return out, err
}

func (c *Client) GetDomainByUUID(uuid string) (*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dom, err := cc.LookupDomainByUUIDString(uuid)
	if err != nil {
		return nil, err
	}
	xml, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &Domain{cc: c, ptr: dom}
	err = out.UnmarshalX(xml)
	return out, err
}

func (c *Client) GetDomainByName(name string) (*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dom, err := cc.LookupDomainByName(name)
	if err != nil {
		return nil, err
	}
	xml, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &Domain{cc: c, ptr: dom}
	err = out.UnmarshalX(xml)
	return out, err
}

func (c *Client) CreateDomain(domain *spec.Domain) (*Domain, error) {
	doc, err := domain.MarshalX()
	if err != nil {
		return nil, err
	}
	return c.DomainCreateXML(doc)
}

func (c *Client) DomainCreateXML(xml string) (*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dom, err := cc.DomainCreateXML(xml, libvirt.DOMAIN_START_VALIDATE)
	if err != nil {
		return nil, err
	}
	doc, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &Domain{cc: c, ptr: dom}
	err = out.UnmarshalX(doc)
	return out, err
}

func (c *Client) DefineDomain(domain *spec.Domain) (*Domain, error) {
	doc, err := domain.MarshalX()
	if err != nil {
		return nil, err
	}
	return c.DomainDefineXML(doc)
}

func (c *Client) DomainDefineXML(xml string) (*Domain, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	dom, err := cc.DomainDefineXMLFlags(xml, libvirt.DOMAIN_DEFINE_VALIDATE)
	if err != nil {
		return nil, err
	}
	doc, err := dom.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &Domain{cc: c, ptr: dom}
	err = out.UnmarshalX(doc)
	return out, err
}

type Domain struct {
	cc *Client

	ptr *libvirt.Domain

	spec.Domain
}

func (d *Domain) Deref() *libvirt.Domain {
	return d.ptr
}

func (d *Domain) Define() error {
	cc, err := d.cc.NewSession()
	if err != nil {
		return err
	}
	defer cc.Close()

	doc, err := d.MarshalX()
	if err != nil {
		return err
	}
	_, err = cc.DomainDefineXML(doc)
	return err
}

// SetVCPUs set up domain cpu, placement has value static
func (d *Domain) SetVCPUs(vcpu *spec.DomainVCPU, cpus ...spec.DomainVCPUsVCPU) error {
	d.VCPU = vcpu
	d.VCPUs = &spec.DomainVCPUs{VCPU: cpus}
	return d.Define()
}

func (d *Domain) SetAutoStart(b bool) error {
	return d.ptr.SetAutostart(b)
}

func (d *Domain) AttachDevice(xml string) (*Domain, error) {
	err := d.ptr.AttachDeviceFlags(xml, libvirt.DOMAIN_DEVICE_MODIFY_LIVE)
	if err != nil {
		return nil, err
	}
	return d.cc.GetDomainByUUID(d.UUID)
}

func (d *Domain) AttachInterface(di *spec.DomainInterface) (*Domain, error) {
	x, err := di.MarshalX()
	if err != nil {
		return nil, fmt.Errorf("invalid domain interface: %v", err)
	}
	return d.AttachDevice(x)
}

func (d *Domain) AttachDisk(disk *spec.DomainInterface) (*Domain, error) {
	x, err := disk.MarshalX()
	if err != nil {
		return nil, fmt.Errorf("invalid domain disk: %v", err)
	}
	return d.AttachDevice(x)
}

func (d *Domain) DetachDevice(xml string) (*Domain, error) {
	err := d.ptr.DetachDeviceFlags(xml, libvirt.DOMAIN_DEVICE_MODIFY_LIVE)
	if err != nil {
		return nil, err
	}
	return d.cc.GetDomainByUUID(d.UUID)
}

func (d *Domain) DetachDeviceAlias(alias string) (*Domain, error) {
	err := d.ptr.DetachDeviceAlias(alias, libvirt.DOMAIN_DEVICE_MODIFY_FORCE)
	if err != nil {
		return nil, err
	}

	return d.cc.GetDomainByUUID(d.UUID)
}
