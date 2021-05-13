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

package gkvm

type StorageType string

const (
	StorageTypeDir         StorageType = "dir"          // since libvirt 0.4.1
	StorageTypeFs          StorageType = "fs"           // since libvirt 0.4.1
	StorageTypeNetFs       StorageType = "netfs"        // since libvirt 0.4.1
	StorageTypeDisk        StorageType = "disk"         // since libvirt 0.4.1
	StorageTypeIscsi       StorageType = "iscsi"        // since libvirt 0.4.1
	StorageTypeLogical     StorageType = "logical"      // since libvirt 0.4.1
	StorageTypeScsi        StorageType = "scsi"         // since libvirt 0.4.1
	StorageTypeMPath       StorageType = "mpath"        // since libvirt 0.7.1
	StorageTypeRdb         StorageType = "rdb"          // since libvirt 0.9.13
	StorageTypeSheepDog    StorageType = "sheepdog"     // since libvirt 0.10.0
	StorageTypeGluster     StorageType = "gluster"      // since libvirt 1.2.0
	StorageTypeZfs         StorageType = "zfs"          // since libvirt 1.2.8
	StorageTypeVStorage    StorageType = "vstorage"     // since libvirt 3.1.0
	StorageTypeIscsiDirect StorageType = "iscsi-direct" // since libvirt 4.7.0
)

type StoragePoolSourceAuthType string

const (
	AuthTypeChap StoragePoolSourceAuthType = "chap"
	AuthTypeCeph StoragePoolSourceAuthType = "ceph"
)

type StoragePoolXml struct {
	Type StorageType `xml:"type,attr"`

	// Storage Pool Namespaces
	// Since 5.1.0
	XmlFS string `xml:"xmlns:fs,omitempty" json:"ns:fs,omitempty"`

	// Since 5.1.0
	XmlRdb string `xml:"xmlns:rdb,omitempty" json:"ns:rdb,omitempty"`

	Name string `xml:"name,attr" json:"name"`

	UUID string `xml:"uuid" json:"uuid"`

	Allocation Size `xml:"allocation" json:"allocation"`

	Capacity Size `xml:"capacity" json:"capacity"`

	Available Size `xml:"available" json:"available"`

	// some pools support optional features
	Features *StoragePoolFeatures `xml:"features,omitempty" json:"features,omitempty"`

	// A single source element is contained within the top level pool element.
	// This tag is used to describe the source of the storage pool.
	Source StoragePoolSource `xml:"source,omitempty" json:"source,omitempty"`

	// A single target element is contained within the top level pool element for
	// some types of pools (pool types dir, fs, netfs, logical, disk, iscsi, scsi, mpath, zfs).
	// This tag is used to describe the mapping of the storage pool into the host filesystem.
	Target StoragePoolTarget `xml:"target" json:"target"`

	// since 5.1.0
	FSMountOpts *FSMountOpts `xml:"fs:mount_opts,omitempty" json:"fsMountOpts,omitempty"`

	// since 5.1.0
	RdbConfigOpts *RdbConfigOpts `xml:"rdb:config_opts,omitempty" json:"rdbConfigOpts,omitempty"`
}

type StoragePoolFeatures struct {
	// Controls whether the filesystem performs copy-on-write (COW) for images in the pool.
	// This may only be set for directory / filesystem pools on the btrfs filesystem.
	// If not set then libvirt will attempt to disable COW on any btrfs filesystems.
	// Since 6.6.0.
	Cow StoragePoolCow `xml:"cow" json:"cow"`
}

type StoragePoolCow struct {
	State ButtonState `xml:"state,attr" json:"state"`
}

type StoragePoolSource struct {
	// Provides the source for pools backed by storage from a named element
	// (pool types logical, rbd, sheepdog, gluster). Contains a string identifier.
	// Since 0.4.5
	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Host *StoragePoolSourceHost `xml:"host,omitempty" json:"host,omitempty"`

	// Required by the iscsi-direct pool in order to provide the iSCSI Qualified Name (IQN)
	// to communicate with the pool's device target IQN. There is one sub-element iqn with
	// the name attribute to describe the IQN for the initiator.
	// Since 4.7.0
	Initiator *StoragePoolSourceInitiator `xml:"initiator,omitempty" json:"initiator,omitempty"`

	Device *StoragePoolSourceDevice `xml:"device,omitempty" json:"device,omitempty"`

	// If present, the auth element provides the authentication credentials needed to access
	// the source by the setting of the type attribute (pool types iscsi, iscsi-direct, rbd).
	// The type must be either "chap" or "ceph". Use "ceph" for Ceph RBD (Rados Block Device)
	// network sources and use "iscsi" for CHAP (Challenge-Handshake Authentication Protocol)
	// iSCSI targets. Additionally a mandatory attribute username identifies the username to
	// use during authentication as well as a sub-element secret with a mandatory attribute type,
	// to tie back to a libvirt secret object that holds the actual password or other credentials.
	// The domain XML intentionally does not expose the password, only the reference to the object
	// that manages the password. The secret element requires either a uuid attribute with the UUID
	// of the secret object or a usage attribute matching the key that was specified in the secret object.
	// Since 0.9.7 for "ceph" and 1.1.1 for "chap"
	Auth *StoragePoolSourceAuth `xml:"auth,omitempty" json:"auth,omitempty"`

	// Provides optional information about the vendor of the storage device.
	// This contains a single attribute name whose value is backend specific.
	// Since 0.8.4
	Vendor *StoragePoolSourceVendor `xml:"vendor,omitempty" json:"vendor,omitempty"`

	// Provides an optional product name of the storage device.
	// This contains a single attribute name whose value is backend specific.
	// Since 0.8.4
	Product *StoragePoolSourceProduct `xml:"product,omitempty" json:"product,omitempty"`

	// Provides information about the format of the pool (pool types fs, netfs, disk, logical).
	// This contains a single attribute type whose value is backend specific.
	// This is typically used to indicate filesystem type, or network filesystem type, or
	// partition table type, or LVM metadata type.
	// All drivers are required to have a default value for this, so it is optional.
	// Since 0.4.1
	Format *StoragePoolSourceFormat `xml:"format,omitempty" json:"format,omitempty"`

	// For a netfs Storage Pool provide a mechanism to define which NFS protocol
	// version number will be used to contact the server's NFS service.
	// The attribute ver accepts an unsigned integer as the version number to use.
	// Since 5.1.0
	Protocol *StoragePoolSourceProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`

	Adapter *StoragePoolSourceAdapter `xml:"adapter,omitempty" json:"adapter,omitempty"`
}

type StoragePoolSourceHost struct {
	Name string `xml:"name,attr" json:"name"`
}

type StoragePoolSourceInitiator struct {
	Iqn  string `xml:"iqn,attr" json:"iqn"`
	Name string `xml:"name,attr" json:"name"`
}

type StoragePoolSourceDevice struct {
	Path          string      `xml:"path,attr,omitempty" json:"path,omitempty"`
	PartSeparator ButtonState `xml:"part_separator,attr" json:"partSeparator,omitempty"`
}

type StoragePoolSourceAuth struct {
	Type StoragePoolSourceAuthType `xml:"type,attr,omitempty" json:"type,omitempty"`

	Username string `xml:"username,attr,omitempty" json:"username,omitempty"`

	Secret StoragePoolSourceAuthSecret `xml:"secret" json:"secret"`
}

type StoragePoolSourceAuthSecret struct {
	Usage string `xml:"usage" json:"usage"`
}

type StoragePoolSourceFormat struct {
	Type string `xml:"type,attr" json:"type,omitempty"`
}

type StoragePoolSourceProtocol struct {
	Ver string `xml:"ver,attr" json:"ver,omitempty"`
}

type StoragePoolSourceAdapter struct {
	Type   string `xml:"type,attr" json:"type"`
	Parent string `xml:"parent,attr,omitempty" json:"parent,omitempty"`
	Wwnn   string `xml:"wwnn,attr,omitempty" json:"wwnn,omitempty"`
	Wwpn   string `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty"`

	ParentAddr *StoragePoolSourceAdapterParentAddr `xml:"parentaddr,omitempty" json:"parentaddr,omitempty"`
}

type StoragePoolSourceAdapterParentAddr struct {
	UniqueId string `xml:"unique_id,attr,omitempty" json:"uniqueId,omitempty"`

	Address StoragePoolSourceAdapterParentAddress `xml:"address" json:"address"`
}

type StoragePoolSourceAdapterParentAddress struct {
	Domain string `xml:"domain,attr" json:"domain"`
	Bus    string `xml:"bus,attr" json:"bus"`
	Slot   string `xml:"slot,attr" json:"slot"`
	Addr   string `xml:"addr,attr" json:"addr"`
}

type StoragePoolSourceVendor struct {
	Name string `xml:"name,attr" json:"name"`
}

type StoragePoolSourceProduct struct {
	Name string `xml:"name,attr" json:"name"`
}

type StoragePoolTarget struct {
	Path        string      `xml:"path" json:"path"`
	Permissions Permissions `xml:"permissions" json:"permissions"`
}

type FSMountOpts struct {
	Opts []FSMountOpt `xml:"fs:option" json:"option"`
}

type FSMountOpt struct {
	Name string `xml:"name,attr" json:"name"`
}

type RdbConfigOpts struct {
	Opts []RdbConfigOpt `xml:"rdb:option" json:"option"`
}

type RdbConfigOpt struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

type StorageVolumeFormatType string

const (
	StorageVolumeTypeRaw   StorageVolumeFormatType = "raw"
	StorageVolumeTypeQcow2 StorageVolumeFormatType = "qcow2"
)

type StorageVolumeXml struct {
	Name string `xml:"name" json:"name"`

	Allocation Size `xml:"allocation" json:"allocation"`

	Capacity Size `xml:"capacity" json:"capacity"`

	Target StoragePoolTarget `xml:"target" json:"target"`
}

type StorageVolumeTarget struct {
	Path string `xml:"path" json:"path"`

	Format StorageVolumeFormatType `xml:"format" json:"format"`

	Permissions *Permissions `xml:"permissions,omitempty" json:"permissions,omitempty"`

	Encryption *StorageEncryption `xml:"encryption,omitempty" json:"encryption,omitempty"`
}



