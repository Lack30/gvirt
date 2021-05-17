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

type Domain struct {
	cc *Client

	ptr *libvirt.Domain

	spec.Domain
}

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

func (d *Domain) Create() error {
	return d.ptr.Create()
}

func (d *Domain) Shutdown() error {
	return d.ptr.ShutdownFlags(libvirt.DOMAIN_SHUTDOWN_ACPI_POWER_BTN)
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

func (d *Domain) Suspend() error {
	return d.ptr.Suspend()
}

func (d *Domain) Resume() error {
	return d.ptr.Resume()
}

func (d *Domain) Reboot() error {
	return d.ptr.Reboot(libvirt.DOMAIN_REBOOT_ACPI_POWER_BTN)
}

func (d *Domain) Destroy() error {
	return d.ptr.DestroyFlags(libvirt.DOMAIN_DESTROY_GRACEFUL)
}

func (d *Domain) UnDefine() error {
	return d.ptr.UndefineFlags(libvirt.DOMAIN_UNDEFINE_MANAGED_SAVE |
		libvirt.DOMAIN_UNDEFINE_SNAPSHOTS_METADATA |
		libvirt.DOMAIN_UNDEFINE_NVRAM)
}

// SetVCPUs set up domain cpu, placement has value static
func (d *Domain) SetVCPUs(vcpu *spec.DomainVCPU, cpus ...spec.DomainVCPUsVCPU) error {
	d.VCPU = vcpu
	d.VCPUs = &spec.DomainVCPUs{VCPU: cpus}
	return d.Define()
}


