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

type StoragePool struct {
	cc *Client

	ptr *libvirt.StoragePool

	spec.StoragePool
}

func (c *Client) GetAllStoragePools() ([]*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pools, err := cc.ListAllStoragePools(libvirt.CONNECT_LIST_STORAGE_POOLS_ACTIVE | libvirt.CONNECT_LIST_STORAGE_POOLS_INACTIVE)
	if err != nil {
		return nil, err
	}
	outs := make([]*StoragePool, 0, len(pools))
	for _, pool := range pools {
		doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
		if err != nil {
			continue
		}
		p := &StoragePool{cc: c, ptr: &pool}
		if err := p.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, p)
	}
	return outs, nil
}

func (c *Client) GetStoragePoolByName(name string) (*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pool, err := cc.LookupStoragePoolByName(name)
	if err != nil {
		return nil, err
	}
	doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &StoragePool{cc: c, ptr: pool}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetStoragePoolByUUID(uuid string) (*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pool, err := cc.LookupStoragePoolByUUIDString(uuid)
	if err != nil {
		return nil, err
	}
	doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &StoragePool{cc: c, ptr: pool}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) GetStoragePoolByTargetPath(target string) (*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pool, err := cc.LookupStoragePoolByTargetPath(target)
	if err != nil {
		return nil, err
	}
	doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}
	p := &StoragePool{cc: c, ptr: pool}
	err = p.UnmarshalX(doc)
	return p, err
}

func (c *Client) StoragePoolCreateXML(xml string) (*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pool, err := cc.StoragePoolCreateXML(xml, libvirt.STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE)
	if err != nil {
		return nil, err
	}
	doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := &StoragePool{cc: c, ptr: pool}
	err = out.UnmarshalX(doc)
	return out, err
}

func (c *Client) StoragePoolDefineXML(xml string) (*StoragePool, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	pool, err := cc.StoragePoolDefineXML(xml, 0)
	if err != nil {
		return nil, err
	}
	doc, err := pool.GetXMLDesc(libvirt.STORAGE_XML_INACTIVE)
	if err != nil {
		return nil, err
	}

	out := &StoragePool{cc: c, ptr: pool}
	err = out.UnmarshalX(doc)
	return out, err
}

func (sp *StoragePool) Refresh() error {
	return sp.ptr.Refresh(0)
}

func (sp *StoragePool) Create() error {
	return sp.ptr.Create(libvirt.STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE)
}

func (sp *StoragePool) Define() error {
	return sp.ptr.Create(libvirt.STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE)
}

func (sp *StoragePool) UnDefine() error {
	return sp.ptr.Undefine()
}

func (sp *StoragePool) Destroy() error {
	return sp.ptr.Destroy()
}

func (sp *StoragePool) Deref() *libvirt.StoragePool {
	return sp.ptr
}

type StorageVolume struct {
	cc *Client

	ptr *libvirt.StorageVol

	spec.StorageVolume
}

func (sp *StoragePool) GetAllStorageVolumes() ([]*StorageVolume, error) {
	volumes, err := sp.ptr.ListAllStorageVolumes(0)
	if err != nil {
		return nil, err
	}

	outs := make([]*StorageVolume, 0, len(volumes))
	for _, volume := range volumes {
		doc, err := volume.GetXMLDesc(0)
		if err != nil {
			continue
		}
		vol := &StorageVolume{cc: sp.cc, ptr: &volume}
		if err := vol.UnmarshalX(doc); err != nil {
			continue
		}
		outs = append(outs, vol)
	}
	return outs, nil
}

func (sp *StoragePool) StorageVolumeCreateXML(xml string) (*StorageVolume, error) {
	vol, err := sp.ptr.StorageVolCreateXML(xml, libvirt.STORAGE_VOL_CREATE_PREALLOC_METADATA)
	if err != nil {
		return nil, err
	}
	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &StorageVolume{cc: sp.cc, ptr: vol}
	err = out.UnmarshalX(doc)
	return out, err
}

func (sp *StoragePool) StorageVolumeCloneXML(xml string, v *StorageVolume) (*StorageVolume, error) {
	vol, err := sp.ptr.StorageVolCreateXMLFrom(xml, v.ptr, libvirt.STORAGE_VOL_CREATE_PREALLOC_METADATA)
	if err != nil {
		return nil, err
	}
	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &StorageVolume{cc: sp.cc, ptr: vol}
	err = out.UnmarshalX(doc)
	return out, err
}

func (c *Client) GetStorageVolumeByKey(key string) (*StorageVolume, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	vol, err := cc.LookupStorageVolByKey(key)
	if err != nil {
		return nil, err
	}
	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &StorageVolume{cc: c, ptr: vol}
	err = out.UnmarshalX(doc)
	return out, err
}

func (c *Client) GetStorageVolumeByPath(path string) (*StorageVolume, error) {
	cc, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	vol, err := cc.LookupStorageVolByPath(path)
	if err != nil {
		return nil, err
	}
	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	out := &StorageVolume{cc: c, ptr: vol}
	err = out.UnmarshalX(doc)
	return out, err
}

func (v *StorageVolume) Delete() error {
	return v.ptr.Delete(libvirt.STORAGE_VOL_DELETE_WITH_SNAPSHOTS)
}

func (v *StorageVolume) Resize(cap uint64) error {
	return v.ptr.Resize(cap, libvirt.STORAGE_VOL_RESIZE_ALLOCATE)
}

func (v *StorageVolume) Deref() *libvirt.StorageVol {
	return v.ptr
}