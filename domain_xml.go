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

type DomainOSSmbiosMode string

const (
	DomainOSSmbiosModeSysInfo DomainOSSmbiosMode = "sysinfo"
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

	// The content of this element defines the maximum number of virtual CPUs allocated for the guest OS,
	// which must be between 1 and the maximum supported by the hypervisor.
	Vcpu Vcpu `xml:"vcpu,omitempty" json:"vcpu,omitempty"`

	// The vcpus element allows to control state of individual vCPUs. The id attribute specifies the vCPU
	// id as used by libvirt in other places such as vCPU pinning, scheduler information and NUMA assignment.
	// Note that the vCPU ID as seen in the guest may differ from libvirt ID in certain cases. Valid IDs are
	// from 0 to the maximum vCPU count as set by the vcpu element minus 1. The enabled attribute allows to
	// control the state of the vCPU. Valid values are yes and no. hotpluggable controls whether given vCPU
	// can be hotplugged and hotunplugged in cases when the CPU is enabled at boot. Note that all disabled vCPUs
	// must be hotpluggable. Valid values are yes and no. order allows to specify the order to add the online
	// vCPUs. For hypervisors/platforms that require to insert multiple vCPUs at once the order may be duplicated
	// across all vCPUs that need to be enabled at once. Specifying order is not necessary, vCPUs are then added
	// in an arbitrary order. If order info is used, it must be used for all online vCPUs. Hypervisors may clear
	// or update ordering information during certain operations to assure valid configuration. Note that hypervisors
	// may create hotpluggable vCPUs differently from boot vCPUs thus special initialization may be necessary.
	// Hypervisors may require that vCPUs enabled on boot which are not hotpluggable are clustered at the beginning
	// starting with ID 0. It may be also required that vCPU 0 is always present and non-hotpluggable.
	// Note that providing state for individual CPUs may be necessary to enable support of addressable vCPU hotplug
	// and this feature may not be supported by all hypervisors. For QEMU the following conditions are required.
	// vCPU 0 needs to be enabled and non-hotpluggable. On PPC64 along with it vCPUs that are in the same core need
	// to be enabled as well. All non-hotpluggable CPUs present at boot need to be grouped after vCPU 0.
	// Since 2.2.0 (QEMU only)
	Vcpus *Vcpus `xml:"vcpus,omitempty" json:"vcpus,omitempty"`
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
	Firmware *DomainOSFirmware `json:"firmware,omitempty" json:"firmware,omitempty"`

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
	Boot []DomainBoot `xml:"boot,omitempty" json:"boot,omitempty"`

	// Whether or not to enable an interactive boot menu prompt on guest startup. The enable attribute can
	// be either "yes" or "no". If not specified, the hypervisor default is used. Since 0.8.3 Additional
	// attribute timeout takes the number of milliseconds the boot menu should wait until it times out.
	// Allowed values are numbers in range [0, 65535] inclusive and it is ignored unless enable is set to "yes".
	// Since 1.2.8
	BootMenu *DomainBootMenu `xml:"bootmenu" json:"bootmenu"`

	// How to populate SMBIOS information visible in the guest. The mode attribute must be specified, and
	// is either "emulate" (let the hypervisor generate all values), "host" (copy all of Block 0 and Block 1,
	// except for the UUID, from the host's SMBIOS values; the virConnectGetSysinfo call can be used to see
	// what values are copied), or "sysinfo" (use the values in the sysinfo element). If not specified,
	// the hypervisor default is used.
	// Since 0.8.7
	Smbios *DomainSmbios `xml:"smbios,omitempty" json:"smbios,omitempty"`

	// This element has attribute useserial with possible values yes or no. It enables or disables Serial Graphics
	// Adapter which allows users to see BIOS messages on a serial port. Therefore, one needs to have serial port
	// defined. Since 0.9.4 . Since 0.10.2 (QEMU only) there is another attribute, rebootTimeout that controls
	// whether and after how long the guest should start booting again in case the boot fails (according to BIOS).
	// The value is in milliseconds with maximum of 65535 and special value -1 disables the reboot.
	Bios *DomainBios `xml:"bios,omitempty" json:"bios,omitempty"`

	// ========================================= Direct kernel boot =========================================
	// When installing a new guest OS it is often useful to boot directly from a kernel and initrd stored in
	// the host OS, allowing command line arguments to be passed directly to the installer. This capability
	// is usually available for both para and full virtualized guests.

	// Type this element has the same semantics as described earlier in the BIOS boot section
	// Loader this element has the same semantics as described earlier in the BIOS boot section

	// The contents of this element specify the fully-qualified path to the kernel image in the host OS.
	Kernel string `xml:"kernel,omitempty" json:"kernel,omitempty"`

	// The contents of this element specify the fully-qualified path to the (optional) ramdisk image in the host OS.
	Initrd string `xml:"initrd,omitempty" json:"initrd,omitempty"`

	// The contents of this element specify arguments to be passed to the kernel (or installer) at boot time.
	// This is often used to specify an alternate primary console (eg serial port), or the installation media
	// source / kickstart file
	Cmdline string `xml:"cmdline,omitempty" json:"cmdline,omitempty"`

	// The contents of this element specify the fully-qualified path to the (optional) device tree binary (dtb)
	// image in the host OS.
	// Since 1.0.4
	Dtb string `xml:"dtb,omitempty" json:"dtb,omitempty"`

	// The table element contains a fully-qualified path to the ACPI table. The type attribute contains the
	// ACPI table type (currently only slic is supported) Since 1.3.5 (QEMU) Since 5.9.0 (Xen)
	Acpi *DomainAcpi `xml:"acpi,omitempty" json:"acpi,omitempty"`

	// =========================================== Container boot ===========================================
	// When booting a domain using container based virtualization, instead of a kernel / boot image, a path
	// to the init binary is required, using the init element. By default this will be launched with no arguments.
	// To specify the initial argv, use the initarg element, repeated as many time as is required. The cmdline
	// element, if set will be used to provide an equivalent to /proc/cmdline but will not affect init argv.
	//
	// To set environment variables, use the initenv element, one for each variable.
	//
	// To set a custom work directory for the init, use the initdir element.
	//
	// To run the init command as a given user or group, use the inituser or initgroup elements respectively.
	// Both elements can be provided either a user (resp. group) id or a name. Prefixing the user or group id
	// with a + will force it to be considered like a numeric value. Without this, it will be first tried
	// as a user or group name.
	//
	// Type
	Init      string   `xml:"init,omitempty" json:"init,omitempty"`
	Initarg   []string `xml:"initarg,omitempty" json:"initarg,omitempty"`
	Initenv   string   `xml:"initenv,omitempty" json:"initenv,omitempty"`
	Initdir   string   `xml:"initdir,omitempty" json:"initdir,omitempty"`
	Inituser  string   `xml:"inituser,omitempty" json:"inituser,omitempty"`
	Initgroup string   `xml:"initgroup,omitempty" json:"initgroup,omitempty"`

	// Some hypervisors allow control over what system information is presented to the guest
	// (for example, SMBIOS fields can be populated by a hypervisor and inspected via the dmidecode
	// command in the guest). The optional sysinfo element covers all such categories of information.
	// Since 0.8.7
	SysInfo []*DomainSysInfo `xml:"sysinfo,omitempty" json:"sysinfo,omitempty"`
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

	Timeout int64 `xml:"timeout,attr" json:"timeout"`
}

type DomainSmbios struct {
	Mode DomainOSSmbiosMode `xml:"mode,attr" json:"mode"`
}

type DomainBios struct {
	UseSerial     ButtonState `xml:"useserial,attr" json:"useserial"`
	RebootTimeout int32       `xml:"rebootTimeout,attr" json:"rebootTimeout"`
}

type DomainAcpi struct {
	Table DomainAcpiTable `xml:"table" json:"table"`
}

type DomainAcpiType string

const (
	DomainAcpiTypeSlic DomainAcpiType = "slic"
)

type DomainAcpiTable struct {
	Type DomainAcpiType `xml:"type,attr" json:"type"`
	Data string         `xml:",chardata" json:"data"`
}

type DomainSysInfoType string

const (
	DomainSysInfoTypeSmbios DomainSysInfoType = "smbios"
	DomainSysInfoTypeFwCfg  DomainSysInfoType = "fwcfg"
)

type DomainSysInfo struct {
	Type DomainSysInfoType `xml:"type,attr" json:"type,omitempty"`

	// This is block 0 of SMBIOS, with entry names drawn from:
	//  vendor  : BIOS Vendor's Name
	//  version : BIOS Version
	//  date    : BIOS release date. If supplied, is in either mm/dd/yy or mm/dd/yyyy format.
	//            If the year portion of the string is two digits, the year is assumed to be 19yy.
	//  release : System BIOS Major and Minor release number values concatenated together as one
	//            string separated by a period, for example, 10.22.
	Bios *Entries `xml:"bios,omitempty" json:"bios,omitempty"`

	// This is block 1 of SMBIOS, with entry names drawn from:
	//  manufacturer  : Manufacturer of BIOS
	//  product       : Product Name
	//  version       : Version of the product
	//  serial        : Serial number
	//  uuid          : Universal Unique ID number. If this entry is provided alongside a top-level uuid
	//                  element, then the two values must match.
	//  sku           : SKU number to identify a particular configuration.
	//  family        : Identify the family a particular computer belongs to.
	System *Entries `xml:"system,omitempty" json:"system,omitempty"`

	// This is block 2 of SMBIOS. This element can be repeated multiple times to describe all the base boards;
	// however, not all hypervisors necessarily support the repetition. The element can have the following children:
	//  manufacturer  : Manufacturer of BIOS
	//  product       : Product Name
	//  version       : Version of the product
	//  serial        : Serial number
	//  asset         : Asset tag
	//  location      : Location in chassis
	//
	// NB: Incorrectly supplied entries for the bios, system or baseBoard blocks will be ignored without error.
	// Other than uuid validation and date format checking, all values are passed as strings to the hypervisor driver.
	BaseBoard *Entries `xml:"baseBoard,omitempty" json:"baseBoard,omitempty"`

	// Since 4.1.0, this is block 3 of SMBIOS, with entry names drawn from:
	//  manufacturer : Manufacturer of Chassis
	//  version      : Version of the Chassis
	//  serial       : Serial number
	//  asset        : Asset tag
	//  sku          : SKU number
	Chassis *Entries `xml:"chassis,omitempty" json:"chassis,omitempty"`

	// This is block 11 of SMBIOS. This element should appear once and can have multiple entry child elements,
	// each providing arbitrary string data. There are no restrictions on what data can be provided in the entries,
	// however, if the data is intended to be consumed by an application in the guest, it is recommended to use
	// the application name as a prefix in the string.
	// ( Since 4.1.0 )
	OemStrings *Entries `xml:"oemStrings,omitempty" json:"oemStrings,omitempty"`
}

type CpuPlacement string

const (
	CpuPlacementAuto   CpuPlacement = "auto"
	CpuPlacementStatic CpuPlacement = "static"
)

type Vcpu struct {
	// The optional attribute placement can be used to indicate the CPU placement mode for domain process.
	// The value can be either "static" or "auto", but defaults to placement of numatune or "static" if
	// cpuset is specified. Using "auto" indicates the domain process will be pinned to the advisory
	// nodeset from querying numad and the value of attribute cpuset will be ignored if it's specified.
	// If both cpuset and placement are not specified or if placement is "static", but no cpuset is specified,
	// the domain process will be pinned to all the available physical CPUs.
	// Since 0.9.11 (QEMU and KVM only)
	Placement CpuPlacement `xml:"placement,attr,omitempty" json:"placement,omitempty"`

	// The optional attribute cpuset is a comma-separated list of physical CPU numbers that domain process
	// and virtual CPUs can be pinned to by default. (NB: The pinning policy of domain process and virtual
	// CPUs can be specified separately by cputune. If the attribute emulatorpin of cputune is specified,
	// the cpuset specified by vcpu here will be ignored. Similarly, for virtual CPUs which have the vcpupin
	// specified, the cpuset specified by cpuset here will be ignored. For virtual CPUs which don't have vcpupin
	// specified, each will be pinned to the physical CPUs specified by cpuset here). Each element in that list
	// is either a single CPU number, a range of CPU numbers, or a caret followed by a CPU number to be excluded
	// from a previous range.
	// Since 0.4.4
	Cpuset string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty"`

	// The optional attribute current can be used to specify whether fewer than the maximum number of virtual
	// CPUs should be enabled.
	// Since 0.8.5
	Current int32 `xml:"current,attr,omitempty" json:"current,omitempty"`

	Id int64 `xml:"id,attr,omitempty" json:"id,omitempty"`

	Enabled ButtonState `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`

	Hotpluggable ButtonState `xml:"hotpluggable,attr,omitempty" json:"hotpluggable,omitempty"`

	Order int32 `xml:"order,attr,omitempty" json:"order,omitempty"`
}

type Vcpus struct {
	Vcpu []Vcpu `xml:"vcpu,omitempty" json:"vcpu,omitempty"`
}
