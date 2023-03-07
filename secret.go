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

type Secret struct {
	cc *Client

	ptr *libvirt.Secret

	spec.Secret
}

func (c *Client) GetAllSecrets() ([]*Secret, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	secrets, err := cc.ListAllSecrets(libvirt.CONNECT_LIST_SECRETS_EPHEMERAL | libvirt.CONNECT_LIST_SECRETS_NO_EPHEMERAL)
	if err != nil {
		return nil, err
	}
	outs := make([]*Secret, 0, len(secrets))
	for i, sec := range secrets {
		doc, err := sec.GetXMLDesc(0)
		if err != nil {
			continue
		}
		p := &Secret{cc: c, ptr: &secrets[i]}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetSecretByName(typ libvirt.SecretUsageType, usage string) (*Secret, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupSecretByUsage(typ, usage)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &Secret{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetSecretByUUID(uuid string) (*Secret, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.LookupSecretByUUIDString(uuid)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}
	p := &Secret{cc: c, ptr: net}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) SecretDefineXML(xml string) (*Secret, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	net, err := cc.SecretDefineXML(xml, 0)
	if err != nil {
		return nil, err
	}
	doc, err := net.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &Secret{cc: c, ptr: net}
	err = out.UnmarshalX(doc)
	return out, err
}

func (n *Secret) UnDefine() error {
	return n.ptr.Undefine()
}

func (n *Secret) Deref() *libvirt.Secret {
	return n.ptr
}
