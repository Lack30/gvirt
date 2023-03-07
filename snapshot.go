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

type DomainSnapshot struct {
	cc *Client

	ptr *libvirt.DomainSnapshot

	spec.DomainSnapshot
}

func (d *Domain) GetAllDomainSnapshots() ([]*DomainSnapshot, error) {
	snapshots, err := d.ptr.ListAllSnapshots(libvirt.DOMAIN_SNAPSHOT_LIST_ACTIVE | libvirt.DOMAIN_SNAPSHOT_LIST_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := make([]*DomainSnapshot, 0, len(snapshots))
	for i, snapshot := range snapshots {
		doc, err := snapshot.GetXMLDesc(libvirt.DOMAIN_SNAPSHOT_XML_SECURE)
		if err != nil {
			continue
		}
		s := &DomainSnapshot{cc: d.cc, ptr: &snapshots[i]}
		if err := s.UnmarshalX(doc); err == nil {
			out = append(out, s)
		}
	}
	return out, nil
}

func (d *Domain) GetDomainSnapshotByName(name string) (*DomainSnapshot, error) {
	snapshot, err := d.ptr.SnapshotLookupByName(name, 0)
	if err != nil {
		return nil, err
	}

	xml, err := snapshot.GetXMLDesc(libvirt.DOMAIN_SNAPSHOT_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &DomainSnapshot{cc: d.cc, ptr: snapshot}
	err = out.UnmarshalX(xml)
	return out, err
}

func (d *Domain) DomainSnapshotCreateXML(xml string) (*DomainSnapshot, error) {
	snapshot, err := d.ptr.CreateSnapshotXML(xml, libvirt.DOMAIN_SNAPSHOT_CREATE_ATOMIC|
		libvirt.DOMAIN_SNAPSHOT_CREATE_VALIDATE)
	if err != nil {
		return nil, err
	}

	doc, err := snapshot.GetXMLDesc(libvirt.DOMAIN_SNAPSHOT_XML_SECURE)
	if err != nil {
		return nil, err
	}

	out := &DomainSnapshot{cc: d.cc, ptr: snapshot}
	err = out.UnmarshalX(doc)
	return out, err
}

func (s *DomainSnapshot) Delete() error {
	return s.ptr.Delete(libvirt.DOMAIN_SNAPSHOT_DELETE_CHILDREN)
}

func (s *DomainSnapshot) Deref() *libvirt.DomainSnapshot {
	return s.ptr
}
