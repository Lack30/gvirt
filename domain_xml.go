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

	// The content of this optional element defines the number of IOThreads to be assigned to the domain for use
	// by supported target storage devices. There should be only 1 or 2 IOThreads per host CPU. There may be more
	// than one supported device assigned to each IOThread.
	// Since 1.2.8
	IOThreads *DomainIOThreads `xml:"iothreads,omitempty" json:"iothreads,omitempty"`

	// The optional iothreadids element provides the capability to specifically define the IOThread ID's for the
	// domain. By default, IOThread ID's are sequentially numbered starting from 1 through the number of iothreads
	// defined for the domain. The id attribute is used to define the IOThread ID. The id attribute must be a positive
	// integer greater than 0. If there are less iothreadids defined than iothreads defined for the domain, then
	// libvirt will sequentially fill iothreadids starting at 1 avoiding any predefined id. If there are more
	// iothreadids defined than iothreads defined for the domain, then the iothreads value will be adjusted accordingly.
	// Since 1.2.15
	IOThreadIds *DomainIOThreadIds `xml:"iothreadids,omitempty" json:"iothreadids,omitempty"`

	// The optional cputune element provides details regarding the CPU tunable parameters for the domain.
	// Note: for the qemu driver, the optional vcpupin and emulatorpin pinning settings are honored after
	// the emulator is launched and NUMA constraints considered. This means that it is expected that other
	// physical CPUs of the host will be used during this time by the domain, which will be reflected by
	// the output of virsh cpu-stats.
	// Since 0.9.0
	CpuTune *DomainCpuTune `xml:"cputune,omitempty" json:"cputune,omitempty"`

	// The run time maximum memory allocation of the guest. The initial memory specified by either the <memory>
	// element or the NUMA cell size configuration can be increased by hot-plugging of memory to the limit
	// specified by this element. The unit attribute behaves the same as for <memory>. The slots attribute
	// specifies the number of slots available for adding memory to the guest. The bounds are hypervisor
	// specific. Note that due to alignment of the memory chunks added via memory hotplug the full size
	// allocation specified by this element may be impossible to achieve.
	// Since 1.2.14 supported by the QEMU driver.
	MaxMemory *DomainMaxMemory `xml:"maxMemory,omitempty" json:"maxMemory,omitempty"`

	// The maximum allocation of memory for the guest at boot time. The memory allocation includes possible
	// additional memory devices specified at start or hotplugged later. The units for this value are determined
	// by the optional attribute unit, which defaults to "KiB" (kibibytes, 210 or blocks of 1024 bytes). Valid units are
	// "b" or "bytes" for bytes, "KB" for kilobytes (103 or 1,000 bytes), "k" or "KiB" for kibibytes (1024 bytes),
	// "MB" for megabytes (106 or 1,000,000 bytes), "M" or "MiB" for mebibytes (220 or 1,048,576 bytes),
	// "GB" for gigabytes (109 or 1,000,000,000 bytes), "G" or "GiB" for gibibytes (230 or 1,073,741,824 bytes),
	// "TB" for terabytes (1012 or 1,000,000,000,000 bytes), or "T" or "TiB" for tebibytes (240 or 1,099,511,627,776 bytes).
	// However, the value will be rounded up to the nearest kibibyte by libvirt, and may be further rounded to
	// the granularity supported by the hypervisor. Some hypervisors also enforce a minimum, such as 4000KiB.
	// In case NUMA is configured for the guest the memory element can be omitted. In the case of crash,
	// optional attribute dumpCore can be used to control whether the guest memory should be included in the generated
	// coredump or not (values "on", "off").
	// unit since 0.9.11 , dumpCore since 0.10.2 (QEMU only)
	Memory DomainMemory `xml:"memory,omitempty" json:"memory,omitempty"`

	// The actual allocation of memory for the guest. This value can be less than the maximum allocation, to allow
	// for ballooning up the guests memory on the fly. If this is omitted, it defaults to the same value as the
	// memory element. The unit attribute behaves the same as for memory.
	CurrentMemory *DomainCurrentMemory `xml:"currentMemory,omitempty" json:"currentMemory,omitempty"`

	// The optional memoryBacking element may contain several elements that influence how virtual memory
	// pages are backed by host pages.
	MemoryBacking *DomainMemoryBacking `xml:"memoryBacking,omitempty" json:"memoryBacking,omitempty"`

	// The optional memtune element provides details regarding the memory tunable parameters for the domain.
	// If this is omitted, it defaults to the OS provided defaults. For QEMU/KVM, the parameters are applied
	// to the QEMU process as a whole. Thus, when counting them, one needs to add up guest RAM, guest video
	// RAM, and some memory overhead of QEMU itself. The last piece is hard to determine so one needs guess
	// and try. For each tunable, it is possible to designate which unit the number is in on input, using the
	// same values as for <memory>. For backwards compatibility, output is always in KiB. unit since 0.9.11
	// Possible values for all *_limit parameters are in range from 0 to VIR_DOMAIN_MEMORY_PARAM_UNLIMITED.
	MemTune *DomainMemTune `xml:"memtune,omitempty" json:"memtune,omitempty"`

	// The optional numatune element provides details of how to tune the performance of a NUMA host via controlling
	// NUMA policy for domain process. NB, only supported by QEMU driver.
	// Since 0.9.3
	NumaTune *DomainNumaTune `xml:"numatune,omitempty" json:"numatune,omitempty"`

	// The optional blkiotune element provides the ability to tune Blkio cgroup tunable parameters for the domain.
	// If this is omitted, it defaults to the OS provided defaults.
	// Since 0.8.8
	BlkioTune *DomainBlkioTune `xml:"blkiotune,omitempty" json:"blkiotune,omitempty"`

	// Hypervisors may allow for virtual machines to be placed into resource partitions, potentially with nesting
	// of said partitions. The resource element groups together configuration related to resource partitioning.
	// It currently supports a child element partition whose content defines the absolute path of the resource
	// partition in which to place the domain. If no partition is listed, then the domain will be placed in a
	// default partition. It is the responsibility of the app/admin to ensure that the partition exists prior to
	// starting the guest. Only the (hypervisor specific) default partition can be assumed to exist by default.
	//
	// Resource partitions are currently supported by the QEMU and LXC drivers, which map partition paths to
	// cgroups directories, in all mounted controllers.
	// Since 1.0.5
	Resource *DomainResource `xml:"resource,omitempty" json:"resource,omitempty"`
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

type DomainOSSmbiosMode string

const (
	DomainOSSmbiosModeSysInfo DomainOSSmbiosMode = "sysinfo"
)

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

type DomainIOThreadIds struct {
	DomainIOThread []DomainIOThread `xml:"iothread,omitempty" json:"iothread,omitempty"`
}

type DomainIOThread struct {
	Id int32 `xml:"id,attr" json:"id,omitempty"`
}

type DomainIOThreads struct {
	Data int32 `xml:",chardata" json:"data"`
}

type DomainCpuTune struct {
	// The optional vcpupin element specifies which of host's physical CPUs the domain vCPU will
	// be pinned to. If this is omitted, and attribute cpuset of element vcpu is not specified,
	// the vCPU is pinned to all the physical CPUs by default. It contains two required attributes,
	// the attribute vcpu specifies vCPU id, and the attribute cpuset is same as attribute cpuset of
	// element vcpu. QEMU driver support since 0.9.0, Xen driver support since 0.9.1
	VcpuPin []DomainVcpuPin `xml:"vcpupin,omitempty" json:"vcpupin,omitempty"`

	// The optional emulatorpin element specifies which of host physical CPUs the "emulator",
	// a subset of a domain not including vCPU or iothreads will be pinned to. If this is omitted,
	// and attribute cpuset of element vcpu is not specified, "emulator" is pinned to all the physical
	// CPUs by default. It contains one required attribute cpuset specifying which physical CPUs to pin to.
	Emulatorpin DomainEmulatorpin `xml:"emulatorpin,omitempty" json:"emulatorpin,omitempty"`

	// The optional iothreadpin element specifies which of host physical CPUs the IOThreads will be pinned to.
	// If this is omitted and attribute cpuset of element vcpu is not specified, the IOThreads are pinned to
	// all the physical CPUs by default. There are two required attributes, the attribute iothread specifies
	// the IOThread ID and the attribute cpuset specifying which physical CPUs to pin to. See the iothreadids
	// description for valid iothread values.
	// Since 1.2.9
	IOThreadPin []DomainIOThreadPin `xml:"iothreadpin,omitempty" json:"iothreadpin,omitempty"`

	// The optional shares element specifies the proportional weighted share for the domain. If this is omitted,
	// it defaults to the OS provided defaults. NB, There is no unit for the value, it's a relative measure based
	// on the setting of other VM, e.g. A VM configured with value 2048 will get twice as much CPU time as a
	// VM configured with value 1024. The value should be in range [2, 262144].
	// Since 0.9.0
	Shares int64 `xml:"shares,omitempty" json:"shares,omitempty"`

	// The optional period element specifies the enforcement interval (unit: microseconds). Within period, each
	// vCPU of the domain will not be allowed to consume more than quota worth of runtime. The value should be
	// in range [1000, 1000000]. A period with value 0 means no value.
	// Only QEMU driver support since 0.9.4, LXC since 0.9.10
	Period int64 `xml:"period,omitempty" json:"period,omitempty"`

	// The optional quota element specifies the maximum allowed bandwidth (unit: microseconds). A domain with
	// quota as any negative value indicates that the domain has infinite bandwidth for vCPU threads, which
	// means that it is not bandwidth controlled. The value should be in range [1000, 17592186044415] or
	// less than 0. A quota with value 0 means no value. You can use this feature to ensure that all vCPUs
	// run at the same speed.
	// Only QEMU driver support since 0.9.4, LXC since 0.9.10
	Quota int32 `xml:"quota,omitempty" json:"quota,omitempty"`

	// The optional global_period element specifies the enforcement CFS scheduler interval (unit: microseconds)
	// for the whole domain in contrast with period which enforces the interval per vCPU. The value should
	// be in range 1000, 1000000]. A global_period with value 0 means no value.
	// Only QEMU driver support since 1.3.3
	GlobalPeriod int64 `xml:"global_period,omitempty" json:"global_period,omitempty"`

	// The optional global_quota element specifies the maximum allowed bandwidth (unit: microseconds) within
	// a period for the whole domain. A domain with global_quota as any negative value indicates that the domain
	// has infinite bandwidth, which means that it is not bandwidth controlled. The value should be in range
	// [1000, 17592186044415] or less than 0. A global_quota with value 0 means no value.
	// Only QEMU driver support since 1.3.3
	GlobalQuota int32 `xml:"global_quota,omitempty" json:"global_quota,omitempty"`

	// The optional emulator_period element specifies the enforcement interval (unit: microseconds).
	// Within emulator_period, emulator threads (those excluding vCPUs) of the domain will not be allowed
	// to consume more than emulator_quota worth of runtime. The value should be in range [1000, 1000000].
	// A period with value 0 means no value.
	// Only QEMU driver support since 0.10.0
	EmulatorPeriod int64 `xml:"emulator_period,omitempty" json:"emulator_period,omitempty"`

	// The optional emulator_quota element specifies the maximum allowed bandwidth (unit: microseconds)
	// for domain's emulator threads (those excluding vCPUs). A domain with emulator_quota as any negative
	// value indicates that the domain has infinite bandwidth for emulator threads (those excluding vCPUs),
	// which means that it is not bandwidth controlled. The value should be in range [1000, 17592186044415]
	// or less than 0. A quota with value 0 means no value.
	// Only QEMU driver support since 0.10.0
	EmulatorQuota int32 `xml:"emulator_quota,omitempty" json:"emulator_quota,omitempty"`

	// The optional iothread_period element specifies the enforcement interval (unit: microseconds) for IOThreads.
	// Within iothread_period, each IOThread of the domain will not be allowed to consume more than iothread_quota
	// worth of runtime. The value should be in range [1000, 1000000]. An iothread_period with value 0 means no value.
	// Only QEMU driver support since 2.1.0
	IOThreadPeriod int64 `xml:"iothread_period,omitempty" json:"iothread_period,omitempty"`

	// The optional iothread_quota element specifies the maximum allowed bandwidth (unit: microseconds) for IOThreads.
	// A domain with iothread_quota as any negative value indicates that the domain IOThreads have infinite bandwidth,
	// which means that it is not bandwidth controlled. The value should be in range [1000, 17592186044415] or less
	// than 0. An iothread_quota with value 0 means no value. You can use this feature to ensure that all IOThreads
	// run at the same speed.
	// Only QEMU driver support since 2.1.0
	IOThreadQuota int32 `xml:"iothread_quota,omitempty" json:"iothread_quota,omitempty"`

	// The optional vcpusched, iothreadsched and emulatorsched elements specify the scheduler type
	// (values batch, idle, fifo, rr) for particular vCPU, IOThread and emulator threads respectively.
	// For vcpusched and iothreadsched the attributes vcpus and iothreads select which vCPUs/IOThreads
	// this setting applies to, leaving them out sets the default. The element emulatorsched does not have
	// that attribute. Valid vcpus values start at 0 through one less than the number of vCPU's defined
	// for the domain. Valid iothreads values are described in the iothreadids description. If no
	// iothreadids are defined, then libvirt numbers IOThreads from 1 to the number of iothreads available
	// for the domain. For real-time schedulers (fifo, rr), priority must be specified as well (and is
	// ignored for non-real-time ones). The value range for the priority depends on the host kernel
	// (usually 1-99).
	// Since 1.2.13 emulatorsched since 5.3.0
	Vcpusched *DomainVcpuSched `xml:"vcpusched,omitempty" json:"vcpusched,omitempty"`

	IOThreadSched *DomainIOThreadSched `xml:"iothreadsched,omitempty" json:"iothreadsched,omitempty"`

	// Optional cachetune element can control allocations for CPU caches using the resctrl on the host.
	// Whether or not is this supported can be gathered from capabilities where some limitations like
	// minimum size and required granularity are reported as well. The required attribute vcpus specifies
	// to which vCPUs this allocation applies. A vCPU can only be member of one cachetune element allocation.
	// The vCPUs specified by cachetune can be identical with those in memorytune, however they are not
	// allowed to overlap.
	CacheTune []DomainCacheTune `xml:"cachetune,omitempty" json:"cachetune,omitempty"`

	// Optional memorytune element can control allocations for memory bandwidth using the resctrl on the host.
	// Whether or not is this supported can be gathered from capabilities where some limitations like minimum
	// bandwidth and required granularity are reported as well. The required attribute vcpus specifies to which
	// vCPUs this allocation applies. A vCPU can only be member of one memorytune element allocation.
	// The vcpus specified by memorytune can be identical to those specified by cachetune.
	// However they are not allowed to overlap each other.
	MemoryTune *DomainMemoryTune `xml:"memorytune,omitempty" json:"memorytune,omitempty"`
}

type DomainVcpuPin struct {
	Vcpu   int32  `xml:"vcpu,attr,omitempty" json:"vcpu,omitempty"`
	CpuSet string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty"`
}

type DomainEmulatorpin struct {
	CpuSet string `xml:"cpuset,omitempty" json:"cpuset,omitempty"`
}

type DomainIOThreadPin struct {
	IOThread int32  `xml:"iothread,attr,omitempty" json:"iothread,omitempty"`
	CpuSet   string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty"`
}

type CPUSchedType string

const (
	CPUSchedTypeBatch CPUSchedType = "batch"
	CPUSchedTypeIdle  CPUSchedType = "idle"
	CPUSchedTypeFifo  CPUSchedType = "fifo"
	CPUSchedTypeRR    CPUSchedType = "rr"
)

type DomainVcpuSched struct {
	Vcpus     string       `xml:"vcpus,omitempty" json:"vcpus,omitempty"`
	Scheduler CPUSchedType `xml:"scheduler,omitempty" json:"scheduler,omitempty"`
	Priority  int32        `xml:"priority,omitempty" json:"priority,omitempty"`
}

type DomainIOThreadSched struct {
	IOThreads int32        `xml:"iothreads,omitempty" json:"iothreads,omitempty"`
	Scheduler CPUSchedType `xml:"scheduler,omitempty" json:"scheduler,omitempty"`
}

type DomainCacheTune struct {
	Vcpus string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`

	// This optional element controls the allocation of CPU cache
	Cache []DomainTuneCache `xml:"cache,omitempty" json:"cache,omitempty"`

	// The optional element monitor creates the cache monitor(s) for current cache allocation
	Monitor []DomainTuneMonitor `xml:"monitor,omitempty" json:"monitor,omitempty"`
}

type CacheType string

const (
	CacheTypeCode CacheType = "code"
	CacheTypeData CacheType = "data"
	CacheTypeBoth CacheType = "both"
)

type DomainTuneCache struct {
	// Host cache id from which to allocate.
	Id int32 `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Host cache level from which to allocate.
	Level int32 `xml:"level,attr,omitempty" json:"level,omitempty"`

	// Type of allocation. Can be code for code (instructions), data for data or both for both code and data (unified).
	// Currently the allocation can be done only with the same type as the host supports, meaning you cannot request
	// both for host with CDP (code/data prioritization) enabled.
	Type CacheType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// The size of the region to allocate. The value by default is in bytes, but the unit attribute can
	// be used to scale the value.
	Size int64 `xml:"size,attr,omitempty" json:"size,omitempty"`

	// If specified it is the unit such as KiB, MiB, GiB, or TiB (described in the memory element for Memory
	// Allocation) in which size is specified, defaults to bytes.
	Unit Unit `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainTuneMonitor struct {
	// Host cache level the monitor belongs to.
	Level int32 `xml:"level,attr,omitempty" json:"level,omitempty"`

	// vCPU list the monitor applies to. A monitor's vCPU list can only be the member(s) of the vCPU list
	// of the associated allocation. The default monitor has the same vCPU list as the associated allocation.
	// For non-default monitors, overlapping vCPUs are not permitted.
	Vcpus string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`
}

type DomainMemoryTune struct {
	Vcpus string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`

	// This element controls the allocation of CPU memory bandwidth
	Node []DomainTuneNode `xml:"node,attr" json:"node,omitempty"`
}

type DomainTuneNode struct {
	// Host node id from which to allocate memory bandwidth.
	Id int32 `xml:"id,attr,omitempty" json:"id,omitempty"`

	// The memory bandwidth to allocate from this node. The value by default is in percentage.
	BandWidth int32 `xml:"bandwidth,attr,omitempty" json:"bandwidth,omitempty"`
}

type DomainMaxMemory struct {
	Slots int32 `xml:"slots,attr,omitempty" json:"slots,omitempty"`
	Unit  Unit  `xml:"unit,attr,attr" json:"unit,omitempty,omitempty"`
	Data  int64 `xml:",chardata" json:"data"`
}

type DomainMemory struct {
	Unit Unit  `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Data int64 `xml:",chardata" json:"data"`
}

type DomainCurrentMemory struct {
	Unit Unit  `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Data int64 `xml:",chardata" json:"data"`
}

type DomainMemoryBacking struct {
	// This tells the hypervisor that the guest should have its memory allocated using hugepages instead of the
	// normal native page size. Since 1.2.5 it's possible to set hugepages more specifically per numa node.
	// The page element is introduced. It has one compulsory attribute size which specifies which hugepages
	// should be used (especially useful on systems supporting hugepages of different sizes). The default
	// unit for the size attribute is kilobytes (multiplier of 1024). If you want to use different unit,
	// use optional unit attribute. For systems with NUMA, the optional nodeset attribute may come handy
	// as it ties given guest's NUMA nodes to certain hugepage sizes. From the example snippet, one gigabyte
	// hugepages are used for every NUMA node except node number four. For the correct syntax see this.
	HugePages *DomainHugePages `xml:"hugepages,omitempty" json:"hugepages,omitempty"`

	// Instructs hypervisor to disable shared pages (memory merge, KSM) for this domain. Since 1.0.6
	NoSharePages *Empty `xml:"nosharepages,omitempty" json:"nosharepages,omitempty"`

	// When set and supported by the hypervisor, memory pages belonging to the domain will be locked in
	// host's memory and the host will not be allowed to swap them out, which might be required for some
	// workloads such as real-time. For QEMU/KVM guests, the memory used by the QEMU process itself will
	// be locked too: unlike guest memory, this is an amount libvirt has no way of figuring out in advance,
	// so it has to remove the limit on locked memory altogether. Thus, enabling this option opens up to a
	// potential security risk: the host will be unable to reclaim the locked memory back from the guest when
	// it's running out of memory, which means a malicious guest allocating large amounts of locked memory
	// could cause a denial-of-service attack on the host. Because of this, using this option is discouraged
	// unless your workload demands it; even then, it's highly recommended to set a hard_limit (see memory tuning)
	// on memory allocation suitable for the specific environment at the same time to mitigate the risks described above.
	// Since 1.0.6
	Locked *Empty `xml:"locked,omitempty" json:"locked,omitempty"`

	// Using the type attribute, it's possible to provide "file" to utilize file memorybacking or keep the default "anonymous".
	// Since 4.10.0 , you may choose "memfd" backing. (QEMU/KVM only)
	Source *DomainMemorySource `xml:"source,attr,omitempty" json:"source,omitempty"`

	// Using the mode attribute, specify if the memory is to be "shared" or "private". This can be overridden per numa node by memAccess.
	Access *DomainMemoryAccess `xml:"access,attr,omitempty" json:"access,omitempty"`

	// Using the mode attribute, specify when to allocate the memory by supplying either "immediate" or "ondemand".
	Allocation *DomainMemoryAllocation `xml:"allocation,attr,omitempty" json:"allocation,omitempty"`

	// When set and supported by hypervisor the memory content is discarded just before guest shuts down
	// (or when DIMM module is unplugged). Please note that this is just an optimization and is not guaranteed
	// to work in all cases (e.g. when hypervisor crashes).
	// Since 4.4.0 (QEMU/KVM only)
	Discard *Empty `xml:"discard,omitempty" json:"discard,omitempty"`
}

type DomainHugePages struct {
	Page []DomainPage `xml:"page,omitempty" json:"page,omitempty"`
}

type DomainPage struct {
	Size    int64  `xml:"size,attr" json:"size,omitempty"`
	Unit    Unit   `xml:"unit,attr" json:"unit,omitempty"`
	NodeSet string `xml:"nodeset,omitempty" json:"nodeset,omitempty"`
}

type MemorySourceType string

const (
	MemorySourceTypeFile      MemorySourceType = "file"
	MemorySourceTypeAnonymous MemorySourceType = "anonymous"
	MemorySourceTypeMemFd     MemorySourceType = "memfd"
)

type MemoryAccessMode string

const (
	MemoryAccessModeShared  MemoryAccessMode = "shared"
	MemoryAccessModePrivate MemoryAccessMode = "private"
)

type DomainMemorySource struct {
	Type MemorySourceType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainMemoryAccess struct {
	Mode MemoryAccessMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type MemoryAllocationMode string

const (
	MemoryAllocationModeShared  MemoryAllocationMode = "immediate"
	MemoryAllocationModePrivate MemoryAllocationMode = "ondemand"
)

type DomainMemoryAllocation struct {
	Mode MemoryAllocationMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainMemTune struct {
	// The optional hard_limit element is the maximum memory the guest can use. The units for this value are
	// kibibytes (i.e. blocks of 1024 bytes). Users of QEMU and KVM are strongly advised not to set this limit
	// as domain may get killed by the kernel if the guess is too low, and determining the memory needed for
	// a process to run is an undecidable problem; that said, if you already set locked in memory backing
	// because your workload demands it, you'll have to take into account the specifics of your deployment
	// and figure out a value for hard_limit that is large enough to support the memory requirements of your
	// guest, but small enough to protect your host against a malicious guest locking all memory.
	HardLimit Size `xml:"hard_limit,omitempty" json:"hard_limit,omitempty"`

	// The optional soft_limit element is the memory limit to enforce during memory contention.
	// The units for this value are kibibytes (i.e. blocks of 1024 bytes)
	SoftLimit Size `xml:"soft_limit,omitempty" json:"soft_limit,omitempty"`

	// The optional swap_hard_limit element is the maximum memory plus swap the guest can use. The units for
	// this value are kibibytes (i.e. blocks of 1024 bytes). This has to be more than hard_limit value provided
	SwapHardLimit Size `xml:"swap_hard_limit,omitempty" json:"swap_hard_limit,omitempty"`

	// The optional min_guarantee element is the guaranteed minimum memory allocation for the guest.
	// The units for this value are kibibytes (i.e. blocks of 1024 bytes). This element is only supported
	// by VMware ESX and OpenVZ drivers.
	MinGuarantee Size `xml:"min_guarantee,omitempty" json:"min_guarantee,omitempty"`
}

type DomainNumaTune struct {
	// The optional memory element specifies how to allocate memory for the domain process on a NUMA host.
	// It contains several optional attributes. Attribute mode is either 'interleave', 'strict', 'preferred', or 'restrictive',
	// defaults to 'strict'. The value 'restrictive' specifies using system default policy and only cgroups is used to
	// restrict the memory nodes, and it requires setting mode to 'restrictive' in memnode elements. Attribute
	// nodeset specifies the NUMA nodes, using the same syntax as attribute cpuset of element vcpu. Attribute
	// placement ( since 0.9.12 ) can be used to indicate the memory placement mode for domain process, its
	// value can be either "static" or "auto", defaults to placement of vcpu, or "static" if nodeset is specified.
	// "auto" indicates the domain process will only allocate memory from the advisory nodeset returned from querying
	// numad, and the value of attribute nodeset will be ignored if it's specified. If placement of vcpu is 'auto',
	// and numatune is not specified, a default numatune with placement 'auto' and mode 'strict' will be added implicitly.
	// Since 0.9.3
	Memory *DomainNumaMemory `xml:"memory,omitempty" json:"memory,omitempty"`

	// Optional memnode elements can specify memory allocation policies per each guest NUMA node.
	// For those nodes having no corresponding memnode element, the default from element memory will be used.
	// Attribute cellid addresses guest NUMA node for which the settings are applied. Attributes mode and
	// nodeset have the same meaning and syntax as in memory element. This setting is not compatible with
	// automatic placement.
	// QEMU Since 1.2.7
	MemNode []DomainNumaMemNode `xml:"memnode,omitempty" json:"memnode,omitempty"`
}

type DomainNumaMemoryMode string

const (
	DomainNumaMemoryModeInterleave  DomainNumaMemoryMode = "interleave"
	DomainNumaMemoryModeStrict      DomainNumaMemoryMode = "strict"
	DomainNumaMemoryModePreferred   DomainNumaMemoryMode = "preferred"
	DomainNumaMemoryModeRestrictive DomainNumaMemoryMode = "restrictive"
)

type DomainNumaMemory struct {
	Mode    DomainNumaMemoryMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	NodeSet string               `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty"`
}

type DomainNumaMemNode struct {
	CellId  int32                `xml:"cellid,attr,omitempty" json:"cellid,omitempty"`
	Mode    DomainNumaMemoryMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	NodeSet string               `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty"`
}

type DomainBlkioTune struct {
	// The optional weight element is the overall I/O weight of the guest. The value should be in the
	// range [100, 1000]. After kernel 2.6.39, the value could be in the range [10, 1000].
	Weight int64 `xml:"weight,omitempty" json:"weight,omitempty"`

	// The domain may have multiple device elements that further tune the weights for each host block device
	// in use by the domain. Note that multiple guest disks can share a single host block device, if they are
	// backed by files within the same host file system, which is why this tuning parameter is at the global
	// domain level rather than associated with each guest disk device (contrast this to the <iotune> element
	// which can apply to an individual <disk>). Each device element has two mandatory sub-elements, path describing
	// the absolute path of the device, and weight giving the relative weight of that device, in the range [100, 1000].
	// After kernel 2.6.39, the value could be in the range [10, 1000]. Since 0.9.8 Additionally, the following
	// optional sub-elements can be used:
	// 	read_bytes_sec  : Read throughput limit in bytes per second. Since 1.2.2
	//  write_bytes_sec : Write throughput limit in bytes per second. Since 1.2.2
	//  read_iops_sec   : Read I/O operations per second limit. Since 1.2.2
	//  write_iops_sec  : Write I/O operations per second limit. Since 1.2.2
	Device []DomainBlkioDevice `xml:"device,omitempty" json:"device,omitempty"`
}

type DomainBlkioDevice struct {
	Path         string `xml:"path,omitempty" json:"path,omitempty"`
	Weight       int64  `xml:"weight,omitempty" json:"weight,omitempty"`
	ReadBytesSec int64  `xml:"read_bytes_sec,omitempty" json:"read_bytes_sec,omitempty"`
	WriteByteSec int64  `xml:"write_byte_sec,omitempty" json:"write_byte_sec,omitempty"`
	ReadIOPSSec  int64  `xml:"read_iops_sec,omitempty" json:"read_iops_sec,omitempty"`
	WriteIOPSSec int64  `xml:"write_iops_sec,omitempty" json:"write_iops_sec,omitempty"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty" json:"partition,omitempty"`
}