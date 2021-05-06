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

type DomainType string

const (
	DomainTypeKvm DomainType = "kvm"
)

type DomainOsType string

const (
	DomainOsTypeHVM DomainOsType = "hvm"
)

type DomainOSFirmware string

const (
	DomainOSFirmwareBIOS DomainOSFirmware = "bios"
	DomainOSFirmwareEFI  DomainOSFirmware = "efi"
)

type DomainLoaderType string

const (
	DomainLoaderTypeRom DomainLoaderType = "rom"
)

type DomainXml struct {
	// Type specifies the hypervisor used for running the domain. The allowed values
	// are driver specific, but include "xen", "kvm", "qemu" and "lxc".
	Type DomainType `xml:"type,attr" json:"type"`

	// Id which is a unique integer identifier for the running guest machine.
	// Inactive machines have no id value.
	Id string `xml:"id,attr" json:"id"`

	// The content of the name element provides a short name for the virtual machine.
	// This name should consist only of alphanumeric characters and is required to be
	// unique within the scope of a single host. It is often used to form the filename
	// for storing the persistent configuration file.
	// Since 0.0.1
	Name string `xml:"name" json:"name"`

	// The content of the uuid element provides a globally unique identifier for the
	// virtual machine. The format must be RFC 4122 compliant, eg 3e3fce45-4f53-4fa7-bb32-11f34168b82b.
	// If omitted when defining/creating a new machine, a random UUID is generated. It is also possible
	// to provide the UUID via a sysinfo specification.
	// Since 0.0.1, sysinfo since 0.8.7
	UUID string `xml:"uuid" json:"uuid"`

	// The genid element can be used to add a Virtual Machine Generation ID which exposes a 128-bit,
	// cryptographically random, integer value identifier, referred to as a Globally Unique Identifier
	// (GUID) using the same format as the uuid. The value is used to help notify the guest operating
	// system when the virtual machine is re-executing something that has already executed before, such as:
	// 	1.VM starts executing a snapshot
	//	2.VM is recovered from backup
	//  3.VM is failover in a disaster recovery environment
	//  4.VM is imported, copied, or cloned
	//
	// The guest operating system notices the change and is then able to react as appropriate by marking
	// its copies of distributed databases as dirty, re-initializing its random number generator, etc.
	// The libvirt XML parser will accept both a provided GUID value or just <genid/> in which case a GUID
	// will be generated and saved in the XML. For the transitions such as above, libvirt will change the
	// GUID before re-executing.
	// Since 4.4.0
	GenId string `xml:"genid" json:"genid"`

	// The optional element title provides space for a short description of the domain.
	// The title should not contain any newlines.
	// Since 0.9.10.
	Title string `xml:"title" json:"title"`

	// The content of the description element provides a human readable description of the virtual machine.
	// This data is not used by libvirt in any way, it can contain any information the user wants.
	// Since 0.7.2
	Description string `xml:"description" json:"description"`

	// The metadata node can be used by applications to store custom metadata in the form of XML nodes/trees.
	// Applications must use custom namespaces on their XML nodes/trees, with only one top-level element per
	// namespace (if the application needs structure, they should have sub-elements to their namespace element).
	// Since 0.9.10
	Metadata *DomainMetadata `xml:"metadata,omitempty" json:"metadata,omitempty"`

	OS DomainOS `xml:"os" json:"os"`
}

type DomainMetadata struct {
}

type DomainOS struct {
	// The firmware attribute allows management applications to automatically fill <loader/> and <nvram/>
	// elements and possibly enable some features required by selected firmware. Accepted values are bios
	// and efi. The selection process scans for files describing installed firmware images in specified
	// location and uses the most specific one which fulfils domain requirements. The locations in order
	// of preference (from generic to most specific one) are:
	//  1.usr/share/qemu/firmware
	//  2.etc/qemu/firmware
	//  3.$XDG_CONFIG_HOME/qemu/firmware
	//
	// For more information refer to firmware metadata specification as described in docs/interop/firmware.json
	// in QEMU repository. Regular users do not need to bother. Since 5.2.0 (QEMU and KVM only) For VMware guests,
	// this is set to efi when the guest uses UEFI, and it is not set when using BIOS.
	// Since 5.3.0 (VMware ESX and Workstation/Player)
	Firmware DomainOSFirmware `json:"firmware,omitempty" json:"firmware,omitempty"`

	// The content of the type element specifies the type of operating system to be booted in the virtual machine.
	// hvm indicates that the OS is one designed to run on bare metal, so requires full virtualization.
	// linux (badly named!) refers to an OS that supports the Xen 3 hypervisor guest ABI. There are also two
	// optional attributes, arch specifying the CPU architecture to virtualization, and machine referring to
	// the machine type. The Capabilities XML provides details on allowed values for these. If arch is omitted
	// then for most hypervisor drivers, the host native arch will be chosen. For the test, ESX and VMWare
	// hypervisor drivers, however, the i686 arch will always be chosen even on an x86_64 host.
	// Since 0.0.1
	Type DomainOsType `xml:"type" json:"type"`

	// The optional loader tag refers to a firmware blob, which is specified by absolute path, used to assist
	// the domain creation process. It is used by Xen fully virtualized domains as well as setting the QEMU BIOS
	// file path for QEMU/KVM domains. Xen since 0.1.0, QEMU/KVM since 0.9.12 Then, since 1.2.8 it's possible for
	// the element to have two optional attributes: readonly (accepted values are yes and no) to reflect the fact
	// that the image should be writable or read-only. The second attribute type accepts values rom and pflash.
	// It tells the hypervisor where in the guest memory the file should be mapped. For instance, if the loader
	// path points to an UEFI image, type should be pflash. Moreover, some firmwares may implement the Secure
	// boot feature. Attribute secure can be used to tell the hypervisor that the firmware is capable of Secure
	// Boot feature. It cannot be used to enable or disable the feature itself in the firmware.
	// Since 2.1.0
	Loader *DomainOSLoader `xml:"loader,omitempty" json:"loader,omitempty"`

	// Some UEFI firmwares may want to use a non-volatile memory to store some variables. In the host, this
	// is represented as a file and the absolute path to the file is stored in this element. Moreover, when
	// the domain is started up libvirt copies so called master NVRAM store file defined in qemu.conf. If needed,
	// the template attribute can be used to per domain override map of master NVRAM stores from the config file.
	// Note, that for transient domains if the NVRAM file has been created by libvirt it is left behind and it
	// is management application's responsibility to save and remove file (if needed to be persistent).
	// Since 1.2.8
	Nvram *DomainNvram `xml:"nvram,omitempty" json:"nvram,omitempty"`

	// The dev attribute takes one of the values "fd", "hd", "cdrom" or "network" and is used to specify the next
	// boot device to consider. The boot element can be repeated multiple times to setup a priority list of boot
	// devices to try in turn. Multiple devices of the same type are sorted according to their targets while preserving
	// the order of buses. After defining the domain, its XML configuration returned by libvirt
	// (through virDomainGetXMLDesc) lists devices in the sorted order. Once sorted, the first device is marked
	// as bootable. Thus, e.g., a domain configured to boot from "hd" with vdb, hda, vda, and hdc disks assigned
	// to it will boot from vda (the sorted list is vda, vdb, hda, hdc). Similar domain with hdc, vda, vdb, and
	// hda disks will boot from hda (sorted disks are: hda, hdc, vda, vdb). It can be tricky to configure in
	// the desired way, which is why per-device boot elements (see disks, network interfaces, and USB and
	// PCI devices sections below) were introduced and they are the preferred way providing full control over
	// booting order. The boot element and per-device boot elements are mutually exclusive.
	// Since 0.1.3, per-device boot since 0.8.8
	Boot []DomainBoot `xml:"boot" json:"boot"`

	// Whether or not to enable an interactive boot menu prompt on guest startup. The enable attribute can
	// be either "yes" or "no". If not specified, the hypervisor default is used. Since 0.8.3 Additional
	// attribute timeout takes the number of milliseconds the boot menu should wait until it times out.
	// Allowed values are numbers in range [0, 65535] inclusive and it is ignored unless enable is set to "yes".
	// Since 1.2.8
	BootMenu *DomainBootMenu `xml:"bootmenu" json:"bootmenu"`

	Smbios DomainSmbios `xml:"smbios,omitempty" json:"smbios,omitempty"`

	Bios DomainBios `xml:"bios,omitempty" json:"bios,omitempty"`
}

type DomainOSType struct {
	Arch string `xml:"arch,attr,omitempty" json:"arch,omitempty"`

	Machine string `xml:"machine,attr,omitempty" json:"machine,omitempty"`

	Data DomainOsType `xml:",chardata" json:"data"`
}

type DomainBoot struct {
	Dev string `xml:"dev,attr" json:"dev"`
}

type DomainOSLoader struct {
	ReadOnly ButtonState `xml:"readonly,attr" json:"readonly"`

	Secure ButtonState `xml:"secure,attr" json:"secure"`

	Type DomainLoaderType `xml:"type,attr" json:"type"`

	Data string `xml:",chardata" json:"data"`
}

type DomainNvram struct {
	Template string `xml:"template,attr" json:"template"`

	Data string `xml:",chardata" json:"data"`
}

type DomainBootMenu struct {
	Enable ButtonState `xml:"enable,attr" json:"enable"`

	Timeout string `xml:"timeout,attr" json:"timeout"`
}

type DomainSmbios struct {

}

type DomainBios struct {

}