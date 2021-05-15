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

	// The cpu element is the main container for describing guest CPU requirements. Its match attribute specifies
	// how strictly the virtual CPU provided to the guest matches these requirements. Since 0.7.6 the match attribute
	// can be omitted if topology is the only element within cpu.
	CPU *DomainCPU `xml:"cpu,omitempty" json:"cpu,omitempty"`

	// It is sometimes necessary to override the default actions taken on various events. Not all hypervisors
	// support all events and actions. The actions may be taken as a result of calls to libvirt APIs
	// virDomainReboot , virDomainShutdown , or virDomainShutdownFlags .
	// Using virsh reboot or virsh shutdown would also trigger the event.
	//
	// The content of this element specifies the action to take when the guest requests a poweroff.
	OnPowerOff EventState `xml:"on_poweroff,omitempty" json:"onPoweroff,omitempty"`

	// The content of this element specifies the action to take when the guest requests a reboot.
	OnRestart EventState `xml:"on_restart,omitempty" json:"onRestart,omitempty"`

	// The content of this element specifies the action to take when the guest crashes.
	OnCrash EventState `xml:"on_crash,omitempty" json:"onCrash,omitempty"`

	// The on_lockfailure element ( since 1.0.0 ) may be used to configure what action should be taken
	// when a lock manager loses resource locks. The following actions are recognized by libvirt, although
	// not all of them need to be supported by individual lock managers. When no action is specified,
	// each lock manager will take its default action.
	//  poweroff  : The domain will be forcefully powered off.
	//  restart   : The domain will be powered off and started up again to reacquire its locks.
	//  pause     : The domain will be paused so that it can be manually resumed when lock issues are solved.
	//  ignore    : Keep the domain running as if nothing happened.
	OnLockFailure EventState `xml:"on_lockfailure,omitempty" json:"onLockfailure,omitempty"`

	// These elements enable ('yes') or disable ('no') BIOS support for S3 (suspend-to-mem) and S4
	// (suspend-to-disk) ACPI sleep states. If nothing is specified, then the hypervisor will be left
	// with its default value. Note: This setting cannot prevent the guest OS from performing a suspend
	// as the guest OS itself can choose to circumvent the unavailability of the sleep states
	// (e.g. S4 by turning off completely).
	Pm *DomainPm `xml:"pm,omitempty" json:"pm,omitempty"`

	// Hypervisors may allow certain CPU / machine features to be toggled on/off.
	Features *DomainFeatures `xml:"features,omitempty" json:"features,omitempty"`

	// The guest clock is typically initialized from the host clock. Most operating systems expect the hardware
	// clock to be kept in UTC, and this is the default. Windows, however, expects it to be in so called 'localtime'.
	Clock *DomainClock `xml:"clock,omitempty" json:"clock,omitempty"`

	// Some platforms allow monitoring of performance of the virtual machine and the code executed inside.
	// To enable the performance monitoring events you can either specify them in the perf element or enable
	// them via virDomainSetPerfEvents API. The performance values are then retrieved using the
	// virConnectGetAllDomainStats API. Since 2.0.0
	Perf *DomainPerformance `xml:"perf,omitempty" json:"perf,omitempty"`

	// The final set of XML elements are all used to describe devices provided to the guest domain.
	// All devices occur as children of the main devices element. Since 0.1.3
	Devices []*DomainDevices `xml:"devices,omitempty" json:"devices,omitempty"`
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

	Value DomainOsType `xml:",chardata" json:"value"`
}

type DomainBoot struct {
	Dev string `xml:"dev,attr" json:"dev"`
}

type DomainOSLoader struct {
	ReadOnly ButtonState `xml:"readonly,attr,omitempty" json:"readonly,omitempty"`

	Secure ButtonState `xml:"secure,attr,omitempty" json:"secure,omitempty"`

	Type DomainLoaderType `xml:"type,attr,omitempty" json:"type,omitempty"`

	Value string `xml:",chardata" json:"value"`
}

type DomainNvram struct {
	Template string `xml:"template,attr,omitempty" json:"template,omitempty"`

	Value string `xml:",chardata" json:"value"`
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
	Type  DomainAcpiType `xml:"type,attr" json:"type"`
	Value string         `xml:",chardata" json:"value"`
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
	Id int32 `xml:"id,attr,omitempty" json:"id,omitempty"`
}

type DomainIOThreads struct {
	Value int32 `xml:",chardata" json:"value"`
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
	GlobalPeriod int64 `xml:"global_period,omitempty" json:"globalPeriod,omitempty"`

	// The optional global_quota element specifies the maximum allowed bandwidth (unit: microseconds) within
	// a period for the whole domain. A domain with global_quota as any negative value indicates that the domain
	// has infinite bandwidth, which means that it is not bandwidth controlled. The value should be in range
	// [1000, 17592186044415] or less than 0. A global_quota with value 0 means no value.
	// Only QEMU driver support since 1.3.3
	GlobalQuota int32 `xml:"global_quota,omitempty" json:"globalQuota,omitempty"`

	// The optional emulator_period element specifies the enforcement interval (unit: microseconds).
	// Within emulator_period, emulator threads (those excluding vCPUs) of the domain will not be allowed
	// to consume more than emulator_quota worth of runtime. The value should be in range [1000, 1000000].
	// A period with value 0 means no value.
	// Only QEMU driver support since 0.10.0
	EmulatorPeriod int64 `xml:"emulator_period,omitempty" json:"emulatorPeriod,omitempty"`

	// The optional emulator_quota element specifies the maximum allowed bandwidth (unit: microseconds)
	// for domain's emulator threads (those excluding vCPUs). A domain with emulator_quota as any negative
	// value indicates that the domain has infinite bandwidth for emulator threads (those excluding vCPUs),
	// which means that it is not bandwidth controlled. The value should be in range [1000, 17592186044415]
	// or less than 0. A quota with value 0 means no value.
	// Only QEMU driver support since 0.10.0
	EmulatorQuota int32 `xml:"emulator_quota,omitempty" json:"emulatorQuota,omitempty"`

	// The optional iothread_period element specifies the enforcement interval (unit: microseconds) for IOThreads.
	// Within iothread_period, each IOThread of the domain will not be allowed to consume more than iothread_quota
	// worth of runtime. The value should be in range [1000, 1000000]. An iothread_period with value 0 means no value.
	// Only QEMU driver support since 2.1.0
	IOThreadPeriod int64 `xml:"iothread_period,omitempty" json:"iothreadPeriod,omitempty"`

	// The optional iothread_quota element specifies the maximum allowed bandwidth (unit: microseconds) for IOThreads.
	// A domain with iothread_quota as any negative value indicates that the domain IOThreads have infinite bandwidth,
	// which means that it is not bandwidth controlled. The value should be in range [1000, 17592186044415] or less
	// than 0. An iothread_quota with value 0 means no value. You can use this feature to ensure that all IOThreads
	// run at the same speed.
	// Only QEMU driver support since 2.1.0
	IOThreadQuota int32 `xml:"iothread_quota,omitempty" json:"iothreadQuota,omitempty"`

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
	Value int64 `xml:",chardata" json:"value"`
}

type DomainMemory struct {
	Unit  Unit  `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value int64 `xml:",chardata" json:"value"`
}

type DomainCurrentMemory struct {
	Unit  Unit  `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value int64 `xml:",chardata" json:"value"`
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
	HardLimit Size `xml:"hard_limit,omitempty" json:"hardLimit,omitempty"`

	// The optional soft_limit element is the memory limit to enforce during memory contention.
	// The units for this value are kibibytes (i.e. blocks of 1024 bytes)
	SoftLimit Size `xml:"soft_limit,omitempty" json:"softLimit,omitempty"`

	// The optional swap_hard_limit element is the maximum memory plus swap the guest can use. The units for
	// this value are kibibytes (i.e. blocks of 1024 bytes). This has to be more than hard_limit value provided
	SwapHardLimit Size `xml:"swap_hard_limit,omitempty" json:"swapHardLimit,omitempty"`

	// The optional min_guarantee element is the guaranteed minimum memory allocation for the guest.
	// The units for this value are kibibytes (i.e. blocks of 1024 bytes). This element is only supported
	// by VMware ESX and OpenVZ drivers.
	MinGuarantee Size `xml:"min_guarantee,omitempty" json:"minGuarantee,omitempty"`
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
	Path          string `xml:"path,omitempty" json:"path,omitempty"`
	Weight        int64  `xml:"weight,omitempty" json:"weight,omitempty"`
	ReadBytesSec  int64  `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty"`
	WriteBytesSec int64  `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty"`
	ReadIOPSSec   int64  `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty"`
	WriteIOPSSec  int64  `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty" json:"partition,omitempty"`
}

type DomainCPUMatch string

const (
	DomainCPUMatchMinimum DomainCPUMatch = "minimum"
	DomainCPUMatchExact   DomainCPUMatch = "exact"
	DomainCPUMatchStrict  DomainCPUMatch = "strict"
	DomainCPUMatchNone    DomainCPUMatch = "none"
	DomainCPUMatchPartial DomainCPUMatch = "partial"
	DomainCPUMatchFull    DomainCPUMatch = "full"
)

type DomainCPUMode string

const (
	DomainCPUModeCustom          DomainCPUMode = "custom"
	DomainCPUModeHostMode        DomainCPUMode = "host-mode"
	DomainCPUModeHostPassthrough DomainCPUMode = "host-passthrough"
	DomainCPUModeMaximum         DomainCPUMode = "maximum"
)

type DomainCPU struct {
	// Its match attribute specifies how strictly the virtual CPU provided to the guest matches these requirements.
	// Since 0.7.6 the match attribute can be omitted if topology is the only element within cpu. Possible values
	// for the match attribute are:
	//  minimum : The specified CPU model and features describes the minimum requested CPU. A better CPU will be
	//            provided to the guest if it is possible with the requested hypervisor on the current host. This
	//            is a constrained host-model mode; the domain will not be created if the provided virtual CPU does
	//            not meet the requirements.
	//  exact   : The virtual CPU provided to the guest should exactly match the specification. If such CPU is not
	//            supported, libvirt will refuse to start the domain.
	//  strict  : The domain will not be created unless the host CPU exactly matches the specification. This is
	//            not very useful in practice and should only be used if there is a real reason.
	//
	// Since 0.8.5 the match attribute can be omitted and will default to exact. Sometimes the hypervisor is not
	// able to create a virtual CPU exactly matching the specification passed by libvirt. Since 3.2.0 , an optional
	// check attribute can be used to request a specific way of checking whether the virtual CPU matches the
	// specification. It is usually safe to omit this attribute when starting a domain and stick with the default
	// value. Once the domain starts, libvirt will automatically change the check attribute to the best supported
	// value to ensure the virtual CPU does not change when the domain is migrated to another host. The following
	// values can be used:
	//  none    : Libvirt does no checking and it is up to the hypervisor to refuse to start the domain if it
	//            cannot provide the requested CPU. With QEMU this means no checking is done at all since the
	//            default behavior of QEMU is to emit warnings, but start the domain anyway.
	//  partial : Libvirt will check the guest CPU specification before starting a domain, but the rest is left
	//            on the hypervisor. It can still provide a different virtual CPU.
	//  full    : The virtual CPU created by the hypervisor will be checked against the CPU specification and
	//            the domain will not be started unless the two CPUs match.
	Match DomainCPUMatch `xml:"match,attr,omitempty" json:"match,omitempty"`

	// Since 0.9.10 , an optional mode attribute may be used to make it easier to configure a guest CPU to be as
	// close to host CPU as possible. Possible values for the mode attribute are:
	//  custom          : In this mode, the cpu element describes the CPU that should be presented to the guest.
	//                    This is the default when no mode attribute is specified. This mode makes it so that a persistent
	//                    guest will see the same hardware no matter what host the guest is booted on.
	//  host-model      : The host-model mode is essentially a shortcut to copying host CPU definition from capabilities
	//                    XML into domain XML. Since the CPU definition is copied just before starting a domain, exactly
	//                    the same XML can be used on different hosts while still providing the best guest CPU each host
	//                    supports. The match attribute can't be used in this mode. Specifying CPU model is not supported
	//                    either, but model's fallback attribute may still be used. Using the feature element, specific
	//                    flags may be enabled or disabled specifically in addition to the host model. This may be used
	//                    to fine tune features that can be emulated. (Since 1.1.1) . Libvirt does not model every aspect
	//                    of each CPU so the guest CPU will not match the host CPU exactly. On the other hand, the ABI provided
	//                    to the guest is reproducible. During migration, complete CPU model definition is transferred to the
	//                    destination host so the migrated guest will see exactly the same CPU model for the running instance
	//                    of the guest, even if the destination host contains more capable CPUs or newer kernel; but shutting
	//                    down and restarting the guest may present different hardware to the guest according to the capabilities
	//                    of the new host. Prior to libvirt 3.2.0 and QEMU 2.9.0 detection of the host CPU model via QEMU is
	//                    not supported. Thus the CPU configuration created using host-model may not work as expected.
	//                    Since 3.2.0 and QEMU 2.9.0 this mode works the way it was designed and it is indicated by
	//                    the fallback attribute set to forbid in the host-model CPU definition advertised in domain
	//                    capabilities XML. When fallback attribute is set to allow in the domain capabilities XML,
	//                    it is recommended to use custom mode with just the CPU model from the host capabilities XML.
	//                    Since 1.2.11 PowerISA allows processors to run VMs in binary compatibility mode supporting
	//                    an older version of ISA. Libvirt on PowerPC architecture uses the host-model to signify
	//                    a guest mode CPU running in binary compatibility mode. Example: When a user needs a power7 VM
	//                    to run in compatibility mode on a Power8 host, this can be described in XML as follows :
	//                    ```
	//                     <cpu mode='host-model'>
	//                     <model>power7</model>
	//                     </cpu>
	//                    ```
	//  host-passthrough : With this mode, the CPU visible to the guest should be exactly the same as the host CPU
	//                     even in the aspects that libvirt does not understand. Though the downside of this mode
	//                     is that the guest environment cannot be reproduced on different hardware. Thus, if you
	//                     hit any bugs, you are on your own. Further details of that CPU can be changed using feature
	//                     elements. Migration of a guest using host-passthrough is dangerous if the source and
	//                     destination hosts are not identical in both hardware, QEMU version, microcode version
	//                     and configuration. If such a migration is attempted then the guest may hang or crash
	//                     upon resuming execution on the destination host. Depending on hypervisor version the
	//                     virtual CPU may or may not contain features which may block migration even to an identical
	//                     host. Since 6.5.0 optional migratable attribute may be used to explicitly request such
	//                     features to be removed from (on) or kept in (off) the virtual CPU. This attribute does
	//                     not make migration to another host safer: even with migratable='on' migration will be
	//                     dangerous unless both hosts are identical as described above.
	//  maximum          : When running a guest with hardware virtualization this CPU model is functionally identical
	//                     to host-passthrough, so refer to the docs above.
	//                     When running a guest with CPU emulation, this CPU model will enable the maximum set of
	//                     features that the emulation engine is able to support. Note that even with migratable='on'
	//                     migration will be dangerous unless both hosts are running identical versions of the emulation
	//                     code. Since 7.1.0 with the QEMU driver.
	//
	// Both host-model and host-passthrough modes make sense when a domain can run directly on the host CPUs
	// (for example, domains with type kvm). The actual host CPU is irrelevant for domains with emulated virtual
	// CPUs (such as domains with type qemu). However, for backward compatibility host-model may be implemented
	// even for domains running on emulated CPUs in which case the best CPU the hypervisor is able to emulate may
	// be used rather then trying to mimic the host CPU model.
	// If an application does not care about a specific CPU, just wants the best featureset without a need for
	// migration compatibility, the maximum model is a good choice on hypervisors where it is available.
	Mode DomainCPUMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`

	//
	Migratable TurnState `xml:"migratable,attr,omitempty" json:"migratable,omitempty"`

	// The content of the model element specifies CPU model requested by the guest. The list of available CPU
	// models and their definition can be found in directory cpu_map, installed in libvirt's data directory.
	// If a hypervisor is not able to use the exact CPU model, libvirt automatically falls back to a closest
	// model supported by the hypervisor while maintaining the list of CPU features. Since 0.9.10 , an optional
	// fallback attribute can be used to forbid this behavior, in which case an attempt to start a domain requesting
	// an unsupported CPU model will fail. Supported values for fallback attribute are: allow (this is the default),
	// and forbid. The optional vendor_id attribute ( Since 0.10.0 ) can be used to set the vendor id seen by the guest.
	// It must be exactly 12 characters long. If not set the vendor id of the host is used. Typical possible values are
	// "AuthenticAMD" and "GenuineIntel".
	Model *DomainCPUModel `xml:"model,omitempty" json:"model,omitempty"`

	// Since 0.8.3 the content of the vendor element specifies CPU vendor requested by the guest. If this element is
	// missing, the guest can be run on a CPU matching given features regardless on its vendor. The list of supported
	// vendors can be found in cpu_map/*_vendors.xml.
	Vendor string `xml:"vendor,omitempty" json:"vendor,omitempty"`

	// The topology element specifies requested topology of virtual CPU provided to the guest. Four attributes,
	// sockets, dies, cores, and threads, accept non-zero positive integer values. They refer to the number of
	// CPU sockets per NUMA node, number of dies per socket, number of cores per die, and number of threads per
	// core, respectively. The dies attribute is optional and will default to 1 if omitted, while the other
	// attributes are all mandatory. Hypervisors may require that the maximum number of vCPUs specified by
	// the cpus element equals to the number of vcpus resulting from the topology.
	Topology *DomainCPUTopology `xml:"topology,omitempty" json:"topology,omitempty"`

	// Since 3.3.0 the cache element describes the virtual CPU cache. If the element is missing, the hypervisor
	// will use a sensible default.
	Cache *DomainCPUCache `xml:"cache,omitempty" json:"cache,omitempty"`

	// The cpu element can contain zero or more feature elements used to fine-tune features provided by
	// the selected CPU model. The list of known feature names can be found in the same file as CPU models.
	// The meaning of each feature element depends on its policy attribute, which has to be set to one of
	// the following values:
	//  force    : The virtual CPU will claim the feature is supported regardless of it being supported by host CPU.
	//  require  : Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it.
	//  optional : The feature will be supported by virtual CPU if and only if it is supported by host CPU.
	//  disable  : The feature will not be supported by virtual CPU.
	//  forbid   : Guest creation will fail if the feature is supported by host CPU.
	//
	// Since 0.8.5 the policy attribute can be omitted and will default to require. Individual CPU feature names
	// are specified as part of the name attribute. For example, to explicitly specify the 'pcid' feature with
	// Intel IvyBridge CPU model:
	// ```
	//	...
	//	<cpu match='exact'>
	//  	<model fallback='forbid'>IvyBridge</model>
	//  	<vendor>Intel</vendor>
	//  	<feature policy='require' name='pcid'/>
	//	</cpu>
	//	...
	// ```
	Feature []*DomainCPUFeature `xml:"feature,omitempty" json:"feature,omitempty"`

	// Guest NUMA topology can be specified using the numa element.
	// Since 0.9.8
	Numa *DomainCPUNuma `xml:"numa,omitempty" json:"numa,omitempty"`
}

type DomainCPUModelFallback string

const (
	DomainCPUModelFallbackAllow  DomainCPUModelFallback = "allow"
	DomainCPUModelFallbackForbid DomainCPUModelFallback = "forbid"
)

type DomainCPUModel struct {
	Fallback DomainCPUModelFallback `xml:"fallback,attr,omitempty" json:"fallback,omitempty"`
	Value    string                 `xml:",chardata" json:"value"`
}

type DomainCPUTopology struct {
	Sockets int32 `xml:"sockets,attr,omitempty" json:"sockets,omitempty"`
	Dies    int32 `xml:"dies,attr,omitempty" json:"dies,omitempty"`
	Cores   int32 `xml:"cores,attr,omitempty" json:"cores,omitempty"`
	Threads int32 `xml:"threads,attr,omitempty" json:"threads,omitempty"`
}

type DomainCPUPolicy string

const (
	DomainCPUPolicyForce    DomainCPUPolicy = "force"
	DomainCPUPolicyRequire  DomainCPUPolicy = "require"
	DomainCPUPolicyOptional DomainCPUPolicy = "optional"
	DomainCPUPolicyDisable  DomainCPUPolicy = "disable"
	DomainCPUPolicyForbid   DomainCPUPolicy = "forbid"
)

type DomainCPUFeature struct {
	Policy DomainCPUPolicy `xml:"policy,attr,omitempty" json:"policy,omitempty"`
	Name   string          `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainCPUCacheMode string

const (
	DomainCPUCacheModeEmulate     DomainCPUCacheMode = "emulate"
	DomainCPUCacheModePassthrough DomainCPUCacheMode = "passthrough"
	DomainCPUCacheModeDisable     DomainCPUCacheMode = "disable"
)

type DomainCPUCache struct {
	// This optional attribute specifies which cache level is described by the element. Missing attribute means
	// the element describes all CPU cache levels at once. Mixing cache elements with the level attribute set
	// and those without the attribute is forbidden.
	Level int32 `xml:"level,attr,omitempty" json:"level,omitempty"`

	// The following values are supported:
	//  emulate     : The hypervisor will provide a fake CPU cache data.
	//  passthrough : The real CPU cache data reported by the host CPU will be passed through to the virtual CPU.
	//  disable     : The virtual CPU will report no CPU cache of the specified level (or no cache at all if the
	//                level attribute is missing).
	Mode DomainCPUCacheMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainCPUNuma struct {
	// Each cell element specifies a NUMA cell or a NUMA node. cpus specifies the CPU or range of CPUs that are
	// part of the node. Since 6.5.0 For the qemu driver, if the emulator binary supports disjointed cpus ranges
	// in each cell, the sum of all CPUs declared in each cell will be matched with the maximum number of virtual
	// CPUs declared in the vcpu element. This is done by filling any remaining CPUs into the first NUMA cell.
	// Users are encouraged to supply a complete NUMA topology, where the sum of the NUMA CPUs matches the maximum
	// virtual CPUs number declared in vcpus, to make the domain consistent across qemu and libvirt versions.
	// memory specifies the node memory in kibibytes (i.e. blocks of 1024 bytes). Since 6.6.0 the cpus attribute
	// is optional and if omitted a CPU-less NUMA node is created. Since 1.2.11 one can use an additional unit
	// attribute to define units in which memory is specified. Since 1.2.7 all cells should have id attribute in
	// case referring to some cell is necessary in the code, otherwise the cells are assigned ids in the increasing
	// order starting from 0. Mixing cells with and without the id attribute is not recommended as it may result in
	// unwanted behaviour. Since 1.2.9 the optional attribute memAccess can control whether the memory is to be
	// mapped as "shared" or "private". This is valid only for hugepages-backed memory and nvdimm modules.
	// Each cell element can have an optional discard attribute which fine tunes the discard feature for given
	// numa node as described under Memory Backing. Accepted values are yes and no. Since 4.4.0
	//
	// This guest NUMA specification is currently available only for QEMU/KVM and Xen.
	//
	// A NUMA hardware architecture supports the notion of distances between NUMA cells. Since 3.10.0 it is
	// possible to define the distance between NUMA cells using the distances element within a NUMA cell description.
	// The sibling sub-element is used to specify the distance value between sibling NUMA cells. For more details,
	// ee the chapter explaining the system's SLIT (System Locality Information Table) within the ACPI
	// (Advanced Configuration and Power Interface) specification.
	Cell []*DomainCPUNumaCell `xml:"cell,omitempty" json:"cell,omitempty"`

	// The NUMA description has an optional interconnects element that describes the normalized memory read/write
	// latency, read/write bandwidth between Initiator Proximity Domains (Processor or I/O) and Target Proximity
	// Domains (Memory).
	//
	// The interconnects element can have zero or more latency child elements to describe latency between two
	// memory nodes and zero or more bandwidth child elements to describe bandwidth between two memory nodes.
	// Both these have the following mandatory attributes:
	//
	// To describe latency from one NUMA node to a cache of another NUMA node the latency element has optional
	// cache attribute which in combination with target attribute creates full reference to distant NUMA node's
	// cache level. For instance, target='0' cache='1' refers to the first level cache of NUMA node 0.
	Interconnects *DomainCPUNumaInterconnects `xml:"interconnects,omitempty" json:"interconnects,omitempty"`
}

type DomainCPUNumaMemAccess string

const (
	DomainCPUNumaMemAccessShared  DomainCPUNumaMemAccess = "shared"
	DomainCPUNumaMemAccessPrivate DomainCPUNumaMemAccess = "private"
)

type DomainCPUNumaCell struct {
	Id        int64                  `xml:"id,attr,omitempty" json:"id,omitempty"`
	Cpus      string                 `xml:"cpus,omitempty" json:"cpus,omitempty"`
	Memory    int64                  `xml:"memory,omitempty" json:"memory,omitempty"`
	Unit      Unit                   `xml:"unit,omitempty" json:"unit,omitempty"`
	Discard   ButtonState            `xml:"discard,omitempty" json:"discard,omitempty"`
	MemAccess DomainCPUNumaMemAccess `xml:"memaccess,omitempty" json:"memaccess,omitempty"`

	// Describing distances between NUMA cells is currently only supported by Xen and QEMU.
	// If no distances are given to describe the SLIT data between different cells,
	// it will default to a scheme using 10 for local and 20 for remote distances.
	Distances *DomainCPUNumaDistances `xml:"distances,omitempty" json:"distances,omitempty"`

	// Since 6.6.0 the cell element can have a cache child element which describes memory side cache for memory
	// proximity domains. The cache element has a level attribute describing the cache level and thus the element
	// can be repeated multiple times to describe different levels of the cache.
	// The cache element has two mandatory child elements then: size and line which describe cache size and cache
	// line size. Both elements accept two attributes: value and unit which set the value of corresponding cache
	// attribute.
	Cache *DomainCPUCache `xml:"cache,omitempty" json:"cache,omitempty"`
}

type DomainCPUNumaDistances struct {
	Sibling []*DomainCPUNumaDistanceSibling `xml:"sibling,omitempty" json:"sibling,omitempty"`
}

type DomainCPUNumaDistanceSibling struct {
	Id    int64 `xml:"id,attr,omitempty" json:"id,omitempty"`
	Value int64 `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type DomainCPUNumaCacheAssociativity string

const (
	DomainCPUNumaCacheAssociativityNone   DomainCPUNumaCacheAssociativity = "none"
	DomainCPUNumaCacheAssociativityDirect DomainCPUNumaCacheAssociativity = "direct"
	DomainCPUNumaCacheAssociativityFull   DomainCPUNumaCacheAssociativity = "full"
)

type DomainCPUNumaCachePolicy string

const (
	DomainCPUNumaCachePolicyNone         DomainCPUNumaCachePolicy = "none"
	DomainCPUNumaCachePolicyWriteBack    DomainCPUNumaCachePolicy = "writeback"
	DomainCPUNumaCachePolicyWritethrough DomainCPUNumaCachePolicy = "writethrough"
)

type DomainCPUNumaCache struct {
	// Level of the cache this description refers to.
	Level int32 `xml:"level,attr,omitempty" json:"level,omitempty"`

	// Describes cache associativity (accepted values are none, direct and full).
	Associativity DomainCPUNumaCacheAssociativity `xml:"associativity,attr,omitempty" json:"associativity,omitempty"`

	// Describes cache write associativity (accepted values are none, writeback and writethrough).
	Policy DomainCPUNumaCachePolicy `xml:"policy,attr,omitempty" json:"policy,omitempty"`

	Size *DomainCPUNumaCacheSize `xml:"size,omitempty" json:"size,omitempty"`
	Line *DomainCPUNumaCacheSize `xml:"line,omitempty" json:"line,omitempty"`
}

type DomainCPUNumaCacheSize struct {
	Value int64 `xml:"value,attr,omitempty" json:"value,omitempty"`
	Unit  Unit  `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainCPUNumaInterconnects struct {
	Latency   []*DomainCPUNumaInterconnect `xml:"latency,omitempty" json:"latency,omitempty"`
	Bandwidth *DomainCPUNumaInterconnect   `xml:"bandwidth,omitempty" json:"bandwidth,omitempty"`
}

type DomainCPUNumaInterconnectType string

const (
	DomainCPUNumaInterconnectTypeAccess DomainCPUNumaInterconnectType = "access"
	DomainCPUNumaInterconnectTypeRead   DomainCPUNumaInterconnectType = "read"
	DomainCPUNumaInterconnectTypeWrite  DomainCPUNumaInterconnectType = "write"
)

type DomainCPUNumaInterconnect struct {
	// Refers to the source NUMA node
	Initiator int32 `xml:"initiator,attr,omitempty" json:"initiator,omitempty"`

	// Refers to the target NUMA node
	Target int32 `xml:"target,attr,omitempty" json:"target,omitempty"`

	Cache int32 `xml:"cache,attr,omitempty" json:"cache,omitempty"`

	// The type of the access. Accepted values: access, read, write
	Type DomainCPUNumaInterconnectType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// The actual value. For latency this is delay in nanoseconds, for bandwidth this value is in kibibytes per
	// second. Use additional unit attribute to change the units.
	Value int64 `xml:"value,attr,omitempty" json:"value,omitempty"`

	Unit Unit `xml:"unit,omitempty" json:"unit,omitempty"`
}

type EventState string

const (
	EventStateDestroy         EventState = "destroy"
	EventStateRestart         EventState = "restart"
	EventStatePreserve        EventState = "preserve"
	EventStateRenameRestart   EventState = "rename-restart"
	EventStateCoredumpDestroy EventState = "coredump-destroy"
	EventStateCoredumpRestart EventState = "coredump-restart"
	EventStatePoweroff        EventState = "poweroff"
	EventStatePause           EventState = "pause"
	EventStateIgnore          EventState = "ignore"
)

type DomainPm struct {
	SuspendToDisk *DomainPmCase `xml:"suspend-to-disk,omitempty" json:"suspendToDisk,omitempty"`
	SuspendToMem  *DomainPmCase `xml:"suspend-to-mem,omitempty" json:"suspendToMem,omitempty"`
}

type DomainPmCase struct {
	Enabled ButtonState `xml:"enabled,omitempty" json:"enabled,omitempty"`
}

type DomainFeatures struct {
	// Physical address extension mode allows 32-bit guests to address more than 4 GB of memory.
	Pae *Empty `xml:"pae,omitempty" json:"pae,omitempty"`

	// ACPI is useful for power management, for example, with KVM guests it is required for graceful shutdown to work.
	Acpi *Empty `xml:"acpi,omitempty" json:"acpi,omitempty"`

	// APIC allows the use of programmable IRQ management. Since 0.10.2 (QEMU only) there is an optional
	// attribute eoi with values on and off which toggles the availability of EOI (End of Interrupt) for the guest.
	Apic *DomainFeatureApic `xml:"apic,omitempty" json:"apic,omitempty"`

	// Depending on the state attribute (values on, off) enable or disable use of Hardware Assisted Paging.
	// The default is on if the hypervisor detects availability of Hardware Assisted Paging.
	Hap *DomainFeatureCase `xml:"hap,omitempty" json:"hap,omitempty"`

	// Enable Viridian hypervisor extensions for paravirtualizing guest operating systems
	Viridian *Empty `xml:"viridian,omitempty" json:"viridian,omitempty"`

	// Always create a private network namespace. This is automatically set if any interface devices are defined.
	// This feature is only relevant for container based virtualization drivers, such as LXC.
	Privnet *Empty `xml:"privnet,omitempty" json:"privnet,omitempty"`

	// Enable various features improving behavior of guests running Microsoft Windows.
	Hyperv *DomainFeatureHyperv `xml:"hyperv,omitempty" json:"hyperv,omitempty"`

	// Various features to change the behavior of the KVM hypervisor.
	Kvm *DomainFeatureKVM `xml:"kvm,omitempty" json:"kvm,omitempty"`

	// Various features to change the behavior of the Xen hypervisor.
	Xen *DomainFeatureXen `xml:"xen,omitempty" json:"xen,omitempty"`

	// Notify the guest that the host supports paravirtual spinlocks for example by exposing the pvticketlocks
	// mechanism. This feature can be explicitly disabled by using state='off' attribute.
	Pvspinlock *DomainFeatureCase `xml:"pvspinlock,omitempty" json:"pvspinlock,omitempty"`

	// Depending on the state attribute (values on, off, default on) enable or disable the performance
	// monitoring unit for the guest. Since 1.2.12
	Pmu *DomainFeatureCase `xml:"pmu,omitempty" json:"pmu,omitempty"`

	// Depending on the state attribute (values on, off, default on) enable or disable the emulation of
	// VMware IO port, for vmmouse etc. Since 1.2.16
	Vmport *DomainFeatureCase `xml:"vmport,omitempty" json:"vmport,omitempty"`

	// Enable for architectures using a General Interrupt Controller instead of APIC in order to handle interrupts.
	// For example, the 'aarch64' architecture uses gic instead of apic. The optional attribute version specifies
	// the GIC version; however, it may not be supported by all hypervisors. Accepted values are 2, 3 and host.
	// Since 1.2.16
	Gic *DomainFeatureCase `xml:"gic,omitempty" json:"gic,omitempty"`

	// Tune the I/O APIC. Possible values for the driver attribute are: kvm (default for KVM domains) and qemu
	// which puts I/O APIC in userspace which is also known as a split I/O APIC mode.
	// Since 3.4.0 (QEMU/KVM only)
	IOApic *DomainFeatureCase `xml:"ioapic,omitempty" json:"ioapic,omitempty"`

	// Configure the HPT (Hash Page Table) of a pSeries guest. Possible values for the resizing attribute are enabled,
	// which causes HPT resizing to be enabled if both the guest and the host support it; disabled, which causes HPT
	// resizing to be disabled regardless of guest and host support; and required, which prevents the guest from
	// starting unless both the guest and the host support HPT resizing. If the attribute is not defined, the
	// hypervisor default will be used. Since 3.10.0 (QEMU/KVM only).
	//
	// The optional maxpagesize subelement can be used to limit the usable page size for HPT guests. Common values
	// are 64 KiB, 16 MiB and 16 GiB; when not specified, the hypervisor default will be used.
	// Since 4.5.0 (QEMU/KVM only).
	HPT *DomainFeatureHPT `xml:"hpt,omitempty" json:"hpt,omitempty"`

	// Enable QEMU vmcoreinfo device to let the guest kernel save debug details. Since 4.4.0 (QEMU only)
	VmCoreInfo *DomainFeatureCase `xml:"vmcoreinfo,omitempty" json:"vmcoreinfo,omitempty"`

	// Depending on the state attribute (values on, off, default on) enable or disable System Management Mode.
	// Since 2.1.0
	//
	// Optional sub-element tseg can be used to specify the amount of memory dedicated to SMM's extended TSEG.
	// That offers a fourth option size apart from the existing ones (1 MiB, 2 MiB and 8 MiB) that the guest
	// OS (or rather loader) can choose from. The size can be specified as a value of that element, optional
	// attribute unit can be used to specify the unit of the aforementioned value (defaults to 'MiB'). If set
	// to 0 the extended size is not advertised and only the default ones (see above) are available.
	//
	// If the VM is booting you should leave this option alone, unless you are very certain you know what you are doing.
	//
	// This value is configurable due to the fact that the calculation cannot be done right with the guarantee
	// that it will work correctly. In QEMU, the user-configurable extended TSEG feature was unavailable up to
	// and including pc-q35-2.9. Starting with pc-q35-2.10 the feature is available, with default size 16 MiB.
	// That should suffice for up to roughly 272 vCPUs, 5 GiB guest RAM in total, no hotplug memory range, and
	// 32 GiB of 64-bit PCI MMIO aperture. Or for 48 vCPUs, with 1TB of guest RAM, no hotplug DIMM range, and
	// 32GB of 64-bit PCI MMIO aperture. The values may also vary based on the loader the VM is using.
	//
	// Additional size might be needed for significantly higher vCPU counts or increased address space (that can
	// be memory, maxMemory, 64-bit PCI MMIO aperture size; roughly 8 MiB of TSEG per 1 TiB of address space) which
	// can also be rounded up.
	//
	// Due to the nature of this setting being similar to "how much RAM should the guest have" users are advised to
	// either consult the documentation of the guest OS or loader (if there is any), or test this by trial-and-error
	// changing the value until the VM boots successfully. Yet another guiding value for users might be the fact
	// that 48 MiB should be enough for pretty large guests (240 vCPUs and 4TB guest RAM), but it is on purpose
	// not set as default as 48 MiB of unavailable RAM might be too much for small guests (e.g. with 512 MiB of RAM).
	//
	// See Memory Allocation for more details about the unit attribute.
	// Since 4.5.0 (QEMU only)
	Smm *DomainFeatureSmm `xml:"smm,omitempty" json:"smm,omitempty"`

	// Configure HTM (Hardware Transational Memory) availability for pSeries guests. Possible values for the
	// state attribute are on and off. If the attribute is not defined, the hypervisor default will be used.
	// Since 4.6.0 (QEMU/KVM only)
	HTM *DomainFeatureCase `xml:"htm,omitempty" json:"htm,omitempty"`

	// Configure nested HV availability for pSeries guests. This needs to be enabled from the host (L0) in
	// order to be effective; having HV support in the (L1) guest is very desiderable if it's planned to run
	// nested (L2) guests inside it, because it will result in those nested guests having much better performance
	// than they would when using KVM PR or TCG. Possible values for the state attribute are on and off. If the
	// attribute is not defined, the hypervisor default will be used. Since 4.10.0 (QEMU/KVM only)
	NestedHv *DomainFeatureCase `xml:"nested-hv,omitempty" json:"nestedHv,omitempty"`

	// Configure ccf-assist (Count Cache Flush Assist) availability for pSeries guests. Possible values for the
	// state attribute are on and off. If the attribute is not defined, the hypervisor default will be used.
	// Since 5.9.0 (QEMU/KVM only)
	CcfAssist *DomainFeatureCase `xml:"ccf-assist,omitempty" json:"ccfAssist,omitempty"`

	// Some guests might require ignoring unknown Model Specific Registers (MSRs) reads and writes.
	// It's possible to switch this by setting unknown attribute of msrs to ignore. If the attribute
	// is not defined, or set to fault, unknown reads and writes will not be ignored.
	// Since 5.1.0 (bhyve only)
	Msrs *DomainFeatureMsrs `xml:"msrs,omitempty" json:"msrs,omitempty"`

	// Configure cfpc (Cache Flush on Privilege Change) availability for pSeries guests. Possible values for the
	// value attribute are broken (no protection), workaround (software workaround available) and fixed (fixed in
	// hardware). If the attribute is not defined, the hypervisor default will be used.
	// Since 6.3.0 (QEMU/KVM only)
	Cfpc *DomainFeatureCfgCase `xml:"cfpc,omitempty" json:"cfpc,omitempty"`

	// Configure sbbc (Speculation Barrier Bounds Checking) availability for pSeries guests. Possible values for
	// the value attribute are broken (no protection), workaround (software workaround available) and fixed
	// (fixed in hardware). If the attribute is not defined, the hypervisor default will be used.
	// Since 6.3.0 (QEMU/KVM only)
	Sbbc *DomainFeatureCfgCase `xml:"sbbc,omitempty" json:"sbbc,omitempty"`

	// Configure ibs (Indirect Branch Speculation) availability for pSeries guests. Possible values for the
	// value attribute are broken (no protection), workaround (count cache flush), fixed-ibs (fixed by serializing indirect branches),
	// fixed-ccd (fixed by disabling the cache count) and fixed-na (fixed in hardware - no longer applicable).
	// If the attribute is not defined, the hypervisor default will be used.
	// Since 6.3.0 (QEMU/KVM only)
	Ibs *DomainFeatureCfgCase `xml:"ibs,omitempty" json:"ibs,omitempty"`
}

type DomainFeatureApic struct {
	EOI TurnState `xml:"eoi,omitempty" json:"eoi,omitempty"`
}

type DomainFeatureHyperv struct {
	// Relax constraints on timers
	Relaxed *DomainFeatureCase `xml:"relaxed,omitempty" json:"relaxed,omitempty"`

	// Enable virtual APIC
	Vapic *DomainFeatureCase `xml:"vapic,omitempty" json:"vapic,omitempty"`

	// Enable spinlock support , retries - at least 4095
	Spinlocks *DomainFeatureCase `xml:"spinlocks,omitempty" json:"spinlocks,omitempty"`

	// Virtual processor index
	Vpindex *DomainFeatureCase `xml:"vpindex,omitempty" json:"vpindex,omitempty"`

	// Processor time spent on running guest code and on behalf of guest code
	Runtime *DomainFeatureCase `xml:"runtime,omitempty" json:"runtime,omitempty"`

	// Enable Synthetic Interrupt Controller (SynIC)
	Synic *DomainFeatureCase `xml:"synic,omitempty" json:"synic,omitempty"`

	// Enable SynIC timers, optionally with Direct Mode support
	Stimer *DomainFeatureHypervStimer `xml:"stimer,omitempty" json:"stimer,omitempty"`

	// Enable hypervisor reset
	Reset *DomainFeatureCase `xml:"reset,omitempty" json:"reset,omitempty"`

	// Set hypervisor vendor id
	VendorId *DomainFeatureCase `xml:"vendor_id,omitempty" json:"vendorId,omitempty"`

	// Expose frequency MSRs
	Frequencies *DomainFeatureCase `xml:"frequencies,omitempty" json:"frequencies,omitempty"`

	// Enable re-enlightenment notification on migration
	Reenlightenment *DomainFeatureCase `xml:"reenlightenment,omitempty" json:"reenlightenment,omitempty"`

	// Enable PV TLB flush support
	Tlbflush *DomainFeatureCase `xml:"tlbflush,omitempty" json:"tlbflush,omitempty"`

	// Enable PV IPI support
	Ipi *DomainFeatureCase `xml:"ipi,omitempty" json:"ipi,omitempty"`

	// Enable Enlightened VMCS
	Evmcs *DomainFeatures `xml:"evmcs,omitempty" json:"evmcs,omitempty"`
}

type DomainFeatureHypervStimer struct {
	State TurnState `xml:"state,attr,omitempty" json:"state,omitempty"`

	//
	Direct *DomainFeatureCase `xml:"direct,omitempty" json:"direct,omitempty"`
}

type DomainFeatureKVM struct {
	// Hide the KVM hypervisor from standard MSR based discovery
	Hidden *DomainFeatureCase `xml:"hidden,omitempty" json:"hidden,omitempty"`

	// Allows a guest to enable optimizations when running on dedicated vCPUs
	HintDedicated *DomainFeatureCase `xml:"hint-dedicated,omitempty" json:"hintDedicated,omitempty"`

	// Decrease IO completion latency by introducing a grace period of busy waiting
	PollControl *DomainFeatureCase `xml:"poll-control,omitempty" json:"pollControl,omitempty"`
}

type DomainFeatureCaseMode string

const (
	DomainFeatureCaseModeShare DomainFeatureCaseMode = "share_pt"
	DomainFeatureCaseModeSync  DomainFeatureCaseMode = "sync_pt"
)

type DomainFeatureXen struct {
	// Expose the host e820 to the guest (PV only)
	E820Host *DomainFeatureCase `xml:"e820_host,omitempty" json:"e820Host,omitempty"`

	// Enable IOMMU mappings allowing PCI passthrough
	Passthrough *DomainFeatureCase `xml:"passthrough,omitempty" json:"passthrough,omitempty"`
}

type DomainFeatureHPT struct {
	Maxpagesize *Size `xml:"maxpagesize,omitempty" json:"maxpagesize,omitempty"`
}

type DomainFeatureSmm struct {
	Tseg Size `xml:"tseg,omitempty" json:"tseg,omitempty"`
}

type DomainFeatureMsrs struct {
	Unknown string `xml:"unknown,attr,omitempty" json:"unknown,omitempty"`
}

type DomainFeatureCfg string

const (
	DomainFeatureCfgBroken     DomainFeatureCfg = "broken"     // no protection
	DomainFeatureCfgWorkaround DomainFeatureCfg = "workaround" // count cache flush
	DomainFeatureCfgFixed      DomainFeatureCfg = "fixed"      // fixed in hardware
	DomainFeatureCfgFixedIbs   DomainFeatureCfg = "fixed-ibs"  // fixed by serializing indirect branches
	DomainFeatureCfgFixedCcd   DomainFeatureCfg = "fixed-ccd"  // fixed by disabling the cache count
	DomainFeatureCfgFixedNa    DomainFeatureCfg = "fixed-na"   // fixed in hardware - no longer applicable
)

type DomainFeatureCfgCase struct {
	Value DomainFeatureCfg `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type DomainFeatureCase struct {
	State   TurnState             `xml:"state,attr,omitempty" json:"state,omitempty"`
	Retries int64                 `xml:"retries,attr,omitempty" json:"retries,omitempty"`
	Value   string                `xml:"value,attr,omitempty" json:"value,omitempty"`
	Mode    DomainFeatureCaseMode `xml:"mode,omitempty" json:"mode,omitempty"`
}

type DomainClockOffset string

const (
	// DomainClockOffsetUTC the guest clock will always be synchronized to UTC when booted.
	// Since 0.9.11 'utc' mode can be converted to 'variable' mode, which can be controlled
	// by using the adjustment attribute. If the value is 'reset', the conversion is never
	// done (not all hypervisors can synchronize to UTC on each boot; use of 'reset' will cause
	// an error on those hypervisors). A numeric value forces the conversion to 'variable' mode
	// using the value as the initial adjustment. The default adjustment is hypervisor specific.
	DomainClockOffsetUTC DomainClockOffset = "utc"

	// DomainClockOffsetLocaltime the guest clock will be synchronized to the host's configured
	// timezone when booted, if any. Since 0.9.11, the adjustment attribute behaves the same as
	// in 'utc' mode.
	DomainClockOffsetLocaltime DomainClockOffset = "localtime"

	// DomainClockOffsetTimezone the guest clock will be synchronized to the requested timezone
	// using the timezone attribute. Since 0.7.7
	DomainClockOffsetTimezone DomainClockOffset = "timezone"

	// DomainClockOffsetVariable the guest clock will have an arbitrary offset applied relative
	// to UTC or localtime, depending on the basis attribute. The delta relative to UTC (or localtime)
	// is specified in seconds, using the adjustment attribute. The guest is free to adjust the RTC over
	// time and expect that it will be honored at next reboot. This is in contrast to 'utc' and 'localtime'
	// mode (with the optional attribute adjustment='reset'), where the RTC adjustments are lost at each reboot.
	// Since 0.7.7 Since 0.9.11 the basis attribute can be either 'utc' (default) or 'localtime'.
	DomainClockOffsetVariable DomainClockOffset = "variable"
)

type DomainClock struct {
	// The offset attribute takes four possible values, allowing fine grained control over how the guest clock
	// is synchronized to the host. NB, not all hypervisors support all modes.
	Offset DomainClockOffset `xml:"offset,attr,omitempty" json:"offset,omitempty"`

	// Each timer element requires a name attribute, and has other optional attributes that depend on the name
	// specified. Various hypervisors support different combinations of attributes.
	Timer []*DomainClockTimer `xml:"timer,omitempty" json:"timer,omitempty"`
}

type DomainClockTimerName string

const (
	DomainClockTimerNamePlatform    DomainClockTimerName = "platform"
	DomainClockTimerNameHpet        DomainClockTimerName = "hpet"
	DomainClockTimerNameKvmClock    DomainClockTimerName = "kvmclock"
	DomainClockTimerNamePit         DomainClockTimerName = "pic"
	DomainClockTimerNameRtc         DomainClockTimerName = "rtc"
	DomainClockTimerNameTsc         DomainClockTimerName = "tsc"
	DomainClockTimerNameHypervClock DomainClockTimerName = "hypervclock"
	DomainClockTimerNameArmvtimer   DomainClockTimerName = "armvtimer"
)

type DomainClockTimerTrack string

const (
	DomainClockTimerTrackBoot     DomainClockTimerTrack = "boot"
	DomainClockTimerTrackGuest    DomainClockTimerTrack = "guest"
	DomainClockTimerTrackWall     DomainClockTimerTrack = "wall"
	DomainClockTimerTrackRealtime DomainClockTimerTrack = "realtime"
)

type DomainClockTimerTickPolicy string

const (
	// DomainClockTimerTickPolicyDelay continue to deliver ticks at the normal rate. The guest OS will not notice
	// anything is amiss, as from its point of view time will have continued to flow normally. The time in the
	// guest should now be behind the time in the host by exactly the amount of time during which ticks have been missed.
	DomainClockTimerTickPolicyDelay DomainClockTimerTickPolicy = "delay"

	// DomainClockTimerTickPolicyCatchup deliver ticks at a higher rate to catch up with the missed ticks.
	// The guest OS will not notice anything is amiss, as from its point of view time will have continued
	// to flow normally. Once the timer has managed to catch up with all the missing ticks, the time in
	// the guest and in the host should match.
	DomainClockTimerTickPolicyCatchup DomainClockTimerTickPolicy = "catchup"

	// DomainClockTimerTickPolicyMerge merge the missed tick(s) into one tick and inject. The guest time
	// may be delayed, depending on how the OS reacts to the merging of ticks
	DomainClockTimerTickPolicyMerge DomainClockTimerTickPolicy = "merge"

	// DomainClockTimerTickPolicyDiscard Throw away the missed ticks and continue with future injection normally.
	// The guest OS will see the timer jump ahead by a potentially quite significant amount all at once, as if the
	// intervening chunk of time had simply not existed; needless to say, such a sudden jump can easily confuse
	// a guest OS which is not specifically prepared to deal with it. Assuming the guest OS can deal correctly
	// with the time jump, the time in the guest and in the host should now match.
	DomainClockTimerTickPolicyDiscard DomainClockTimerTickPolicy = "discard"
)

type DomainClockTimerMode string

const (
	DomainClockTimerModeAuto     DomainClockTimerMode = "auto"
	DomainClockTimerModeNative   DomainClockTimerMode = "native"
	DomainClockTimerModeEmulate  DomainClockTimerMode = "emulate"
	DomainClockTimerModeParaVirt DomainClockTimerMode = "paravirt"
	DomainClockTimerModeSmpSafe  DomainClockTimerMode = "smpsafe"
)

type DomainClockTimer struct {
	// The name attribute selects which timer is being modified, and can be one of "platform" (currently unsupported),
	// "hpet" (xen, qemu, lxc), "kvmclock" (qemu), "pit" (qemu), "rtc" (qemu, lxc), "tsc" (xen, qemu - since 3.2.0 ),
	// "hypervclock" (qemu - since 1.2.2 ) or "armvtimer" (qemu - since 6.1.0 ). The hypervclock timer adds support
	// for the reference time counter and the reference page for iTSC feature for guests running the Microsoft
	// Windows operating system.
	Name DomainClockTimerName `xml:"name,attr,omitempty" json:"name,omitempty"`

	// The track attribute specifies what the timer tracks, and can be "boot", "guest", or "wall", or "realtime".
	// Only valid for name="rtc" or name="platform".
	Track DomainClockTimerTrack `xml:"track,attr,omitempty" json:"track,omitempty"`

	// The tickpolicy attribute determines what happens when QEMU misses a deadline for injecting a tick to the guest.
	// This can happen, for example, because the guest was paused.
	TickPolicy DomainClockTimerTickPolicy `xml:"tickpolicy,attr,omitempty" json:"tickpolicy,omitempty"`

	// If the policy is "catchup", there can be further details in the catchup sub-element.
	// The catchup element has three optional attributes, each a positive integer. The attributes are
	// threshold, slew, and limit.
	Catchup *DomainClockTimerCatchup `xml:"catchup,attr,omitempty" json:"catchup,omitempty"`

	// The frequency attribute is an unsigned integer specifying the frequency at which name="tsc" runs.
	Frequency uint64 `xml:"frequency,attr,omitempty" json:"frequency,omitempty"`

	// The mode attribute controls how the name="tsc" timer is managed, and can be "auto", "native", "emulate",
	// "paravirt", or "smpsafe". Other timers are always emulated.
	Mode DomainClockTimerMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`

	// The present attribute can be "yes" or "no" to specify whether a particular timer is available to the guest.
	Present ButtonState `xml:"present,attr,omitempty" json:"present,omitempty"`
}

type DomainClockTimerCatchup struct {
	Threshold int32 `xml:"threshold,attr,omitempty" json:"threshold,omitempty"`
	Slew      int32 `xml:"slew,attr,omitempty" json:"slew,omitempty"`
	Limit     int32 `xml:"limit,attr,omitempty" json:"limit,omitempty"`
}

type DomainPerformance struct {
	Event []*DomainPerformanceEvent `xml:"event,omitempty" json:"event,omitempty"`
}

type DomainPerformanceEventName string

const (
	// DomainPerformanceEventNameCmt usage of l3 cache in bytes by applications running on the platform (perf.cmt)
	DomainPerformanceEventNameCmt DomainPerformanceEventName = "cmt"

	// DomainPerformanceEventNameMbmt total system bandwidth from one level of cache (perf.mbmt)
	DomainPerformanceEventNameMbmt DomainPerformanceEventName = "mbmt"

	// DomainPerformanceEventNameMbml bandwidth of memory traffic for a memory controller (perf.mbml)
	DomainPerformanceEventNameMbml DomainPerformanceEventName = "mbml"

	// DomainPerformanceEventNameCpuCycles the count of CPU cycles (total/elapsed) (perf.cpu_cycles)
	DomainPerformanceEventNameCpuCycles DomainPerformanceEventName = "cpu_cycles"

	// DomainPerformanceEventNameInstructions the count of instructions by applications running on the platform
	// (perf.instructions)
	DomainPerformanceEventNameInstructions DomainPerformanceEventName = "instructions"

	// DomainPerformanceEventNameCacheReferences the count of cache hits by applications running on the platform
	// (perf.cache_references)
	DomainPerformanceEventNameCacheReferences DomainPerformanceEventName = "cache_references"

	// DomainPerformanceEventNameCacheMisses the count of cache misses by applications running on the platform
	// (perf.cache_misses)
	DomainPerformanceEventNameCacheMisses DomainPerformanceEventName = "cache_misses"

	// DomainPerformanceEventNameBranchInstructions the count of branch instructions by applications running on
	// the platform (perf.branch_instructions)
	DomainPerformanceEventNameBranchInstructions DomainPerformanceEventName = "branch_instructions"

	// DomainPerformanceEventNameBranchMisses the count of branch misses by applications running on the platform
	// (perf.branch_misses)
	DomainPerformanceEventNameBranchMisses DomainPerformanceEventName = "branch_misses"

	// DomainPerformanceEventNameBusCycles the count of bus cycles by applications running on the platform
	// (perf.bus_cycles)
	DomainPerformanceEventNameBusCycles DomainPerformanceEventName = "bus_cycles"

	// DomainPerformanceEventNameStalledCyclesFrontend the count of stalled CPU cycles in the frontend of the
	// instruction processor pipeline by applications running on the platform (perf.stalled_cycles_frontend)
	DomainPerformanceEventNameStalledCyclesFrontend DomainPerformanceEventName = "stalled_cycles_frontend"

	// DomainPerformanceEventNameStalledCyclesBackend the count of stalled CPU cycles in the backend of the
	// instruction processor pipeline by applications running on the platform (perf.stalled_cycles_backend)
	DomainPerformanceEventNameStalledCyclesBackend DomainPerformanceEventName = "stalled_cycles_backend"

	// DomainPerformanceEventNameRefCpuCycles the count of total CPU cycles not affected by CPU frequency scaling
	// by applications running on the platform (perf.ref_cpu_cycles)
	DomainPerformanceEventNameRefCpuCycles DomainPerformanceEventName = "ref_cpu_cycles"

	// DomainPerformanceEventNameCpuClock the count of CPU clock time, as measured by a monotonic high-resolution
	// per-CPU timer, by applications running on the platform (perf.cpu_clock)
	DomainPerformanceEventNameCpuClock DomainPerformanceEventName = "cpu_clock"

	// DomainPerformanceEventNameTaskClock the count of task clock time, as measured by a monotonic high-resolution
	// CPU timer, specific to the task that is run by applications running on the platform (perf.task_clock)
	DomainPerformanceEventNameTaskClock DomainPerformanceEventName = "task_clock"

	// DomainPerformanceEventNamePageFaults the count of page faults by applications running on the platform.
	// This includes minor, major, invalid and other types of page faults (perf.page_faults)
	DomainPerformanceEventNamePageFaults DomainPerformanceEventName = "page_faults"

	// DomainPerformanceEventNameContextSwitches the count of context switches by applications running on the platform
	// (perf.context_switches)
	DomainPerformanceEventNameContextSwitches DomainPerformanceEventName = "context_switches"

	// DomainPerformanceEventNameCpuMigrations the count of CPU migrations, that is, where the process moved from one
	// logical processor to another, by applications running on the platform (perf.cpu_migrations)
	DomainPerformanceEventNameCpuMigrations DomainPerformanceEventName = "cpu_migrations"

	// DomainPerformanceEventNamePageFaultsMin the count of minor page faults, that is, where the page was present in
	// the page cache, and therefore the fault avoided loading it from storage, by applications running on the platform
	// (perf.page_faults_min)
	DomainPerformanceEventNamePageFaultsMin DomainPerformanceEventName = "page_faults_min"

	// DomainPerformanceEventNamePageFaultsMaj the count of major page faults, that is, where the page was not
	// present in the page cache, and therefore had to be fetched from storage, by applications running on the platform
	// (perf.page_faults_maj)
	DomainPerformanceEventNamePageFaultsMaj DomainPerformanceEventName = "page_faults_maj"

	// DomainPerformanceEventNameAlignmentFaults the count of alignment faults, that is when the load or store is
	// not aligned properly, by applications running on the platform (perf.alignment_faults)
	DomainPerformanceEventNameAlignmentFaults DomainPerformanceEventName = "alignment_faults"

	// DomainPerformanceEventNameEmulationFaults the count of emulation faults, that is when the kernel traps on
	// unimplemented instructions and emulates them for user space, by applications running on the platform
	// (perf.emulation_faults)
	DomainPerformanceEventNameEmulationFaults DomainPerformanceEventName = "emulation_faults"
)

type DomainPerformanceEvent struct {
	Name   DomainPerformanceEventName `xml:"name,attr,omitempty" json:"name,omitempty"`
	Enable TurnState                  `xml:"enable,attr,omitempty" json:"enablem,omitempty"`
}

type DomainDevices struct {
	// The contents of the emulator element specify the fully qualified path to the device model emulator binary.
	// The capabilities XML specifies the recommended default emulator to use for each particular domain
	// type / architecture combination.
	Emulator string `xml:"emulator,omitempty" json:"emulator,omitempty"`

	// Any device that looks like a disk, be it a floppy, harddisk, cdrom, or paravirtualized driver is
	// specified via the disk element.
	Disk []*DomainDisk `xml:"disk,omitempty" json:"disk,omitempty"`

	// A directory on the host that can be accessed directly from the guest. since 0.3.3, since 0.8.5 for QEMU/KVM
	Filesystem []*DomainFilesystem `xml:"filesystem,omitempty" json:"filesystem,omitempty"`

	// Depending on the guest architecture, some device buses can appear more than once, with a group of virtual
	// devices tied to a virtual controller. Normally, libvirt can automatically infer such controllers without
	// requiring explicit XML markup, but sometimes it is necessary to provide an explicit controller element,
	// notably when planning the PCI topology for guests where device hotplug is expected.
	//
	//```
	//<devices>
	//  <controller type='ide' index='0'/>
	//  <controller type='virtio-serial' index='0' ports='16' vectors='4'/>
	//  <controller type='virtio-serial' index='1'>
	//    <address type='pci' domain='0x0000' bus='0x00' slot='0x0a' function='0x0'/>
	//  </controller>
	//  <controller type='scsi' index='0' model='virtio-scsi'>
	//    <driver iothread='4'/>
	//    <address type='pci' domain='0x0000' bus='0x00' slot='0x0b' function='0x0'/>
	//  </controller>
	//  <controller type='xenbus' maxGrantFrames='64' maxEventChannels='2047'/>
	//  ...
	//</devices>
	//```
	// Each controller has a mandatory attribute type, which must be one of 'ide', 'fdc', 'scsi', 'sata', 'usb',
	// 'ccid', 'virtio-serial' or 'pci', and a mandatory attribute index which is the decimal integer describing
	// in which order the bus controller is encountered (for use in controller attributes of <address> elements).
	// Since 1.3.5 the index is optional; if not specified, it will be auto-assigned to be the lowest unused index
	// for the given controller type.
	Controller []*DomainController `xml:"controller,omitempty" json:"controller,omitempty"`

	// When using a lock manager, it may be desirable to record device leases against a VM. The lock manager
	// will ensure the VM won't start unless the leases can be acquired.
	Lease *DomainLease `xml:"lease,omitempty" json:"lease,omitempty"`

	// USB, PCI and SCSI devices attached to the host can be passed through to the guest using the hostdev element.
	// since after 0.4.4 for USB, 0.6.0 for PCI (KVM only) and 1.0.6 for SCSI (KVM only)
	HostDev []*DomainHostDev `xml:"hostdev,omitempty" json:"hostdev,omitempty"`

	// USB device redirection through a character device is supported since after 0.9.5 (KVM only) :
	ReDirDev []*DomainReDirDev `xml:"redirdev,omitempty" json:"redirdev,omitempty"`

	// Theredirfilterelement is used for creating the filter rule to filter out certain devices from redirection.
	// It uses sub-element <usbdev> to define each filter rule. class attribute is the USB Class code, for example,
	// 0x08 represents mass storage devices. The USB device can be addressed by vendor / product id using the vendor
	// and product attributes. version is the device revision from the bcdDevice field (not the version of the
	// USB protocol). These four attributes are optional and -1 can be used to allow any value for them. allow
	// attribute is mandatory, 'yes' means allow, 'no' for deny.
	ReDirFilter []*DomainReDirFilter `xml:"redirfilter,omitempty" json:"redirfilter,omitempty"`

	// A virtual smartcard device can be supplied to the guest via the smartcard element. A USB smartcard reader
	// device on the host cannot be used on a guest with simple device passthrough, since it will then not be
	// available on the host, possibly locking the host computer when it is "removed". Therefore, some hypervisors
	// provide a specialized virtual device that can present a smartcard interface to the guest, with several modes
	// for describing how credentials are obtained from the host or even a from a channel created to a third-party
	// smartcard provider. Since 0.8.8
	SmartCard []*DomainSmartCard `xml:"smartcard,omitempty" json:"smartcard,omitempty"`

	Interface []*DomainInterface `xml:"interface,omitempty" json:"interface,omitempty"`

	// Input devices allow interaction with the graphical framebuffer in the guest virtual machine. When enabling
	// the framebuffer, an input device is automatically provided.
	Input []*DomainInput `xml:"input,omitempty" json:"input,omitempty"`

	// A hub is a device that expands a single port into several so that there are more ports available to
	// connect devices to a host system.
	Hub []*DomainHub `xml:"hub,omitempty" json:"hub,omitempty"`

	// A graphics device allows for graphical interaction with the guest OS. A guest will typically have either
	// a framebuffer or a text console configured to allow interaction with the admin.
	Graphics []*DomainGraphics `xml:"graphics,omitempty" json:"graphics,omitempty"`

	// The video element is the container for describing video devices. For backwards compatibility, if no video
	// is set but there is a graphics in domain xml, then libvirt will add a default video according to the guest
	// type.
	//
	// For a guest of type "kvm", the default video is: type with value "cirrus", vram with value "16384" and heads
	// with value "1". By default, the first video device in domain xml is the primary one, but the optional attribute
	// primary ( since 1.0.2 ) with value 'yes' can be used to mark the primary in cases of multiple video device.
	// The non-primary must be type of "qxl" or ( since 2.4.0 ) "virtio".
	Video []*DomainVideo `xml:"vidio,omitempty" json:"vidio,omitempty"`

	//
	Parallel []*DomainParallel `xml:"parallel,omitempty" json:"parallel,omitempty"`
	Serial   []*DomainSerial   `xml:"serial,omitempty" json:"serial,omitempty"`
	Console  []*DomainConsole  `xml:"console,omitempty" json:"console,omitempty"`
	Channel  []*DomainChannel  `xml:"channel,omitempty" json:"channel,omitempty"`

	Sound []*DomainSound `xml:"sound,omitempty" json:"sound,omitempty"`
	Audio []*DomainAudio `xml:"audio,omitempty" json:"audio,omitempty"`

	WatchDog   *DomainWatchDog   `xml:"watchdog,omitempty" json:"watchdog,omitempty"`
	MemBalloon *DomainMemBalloon `xml:"memballoon,omitempty" json:"memballoon,omitempty"`
}

type DomainDiskType string

const (
	DomainDiskTypeFile      DomainDiskType = "file"
	DomainDiskTypeBlock     DomainDiskType = "block"
	DomainDiskTypeDir       DomainDiskType = "dir"
	DomainDiskTypeNetwork   DomainDiskType = "network"
	DomainDiskTypeVolume    DomainDiskType = "volume"
	DomainDiskTypeNvme      DomainDiskType = "nvme"
	DomainDiskTypeVHostUser DomainDiskType = "vhostuser"
)

type DomainDiskDevice string

const (
	DomainDiskDeviceFloppy DomainDiskDevice = "floppy"
	DomainDiskDeviceDisk   DomainDiskDevice = "disk"
	DomainDiskDeviceCdrom  DomainDiskDevice = "cdrom"
	DomainDiskDeviceLun    DomainDiskDevice = "lun"
)

type DomainDiskSgio string

const (
	DomainDiskSgioFiltered   DomainDiskSgio = "filtered"
	DomainDiskSgioUnfiltered DomainDiskSgio = "unfiltered"
)

type DomainDiskSnapshot string

const (
	DomainDiskSnapshotInternal DomainDiskSnapshot = "internal"
	DomainDiskSnapshotExternal DomainDiskSnapshot = "external"
)

type DomainDisk struct {
	// Valid values are "file", "block", "dir" ( since 0.7.5 ), "network" ( since 0.8.7 ),
	// or "volume" ( since 1.0.5 ), or "nvme" ( since 6.0.0 ), or "vhostuser" ( since 7.1.0 )
	// and refer to the underlying source for the disk. Since 0.0.3
	Type DomainDiskType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// Indicates how the disk is to be exposed to the guest OS. Possible values for this attribute are "floppy",
	// "disk", "cdrom", and "lun", defaulting to "disk". Using "lun" ( since 0.9.10 ) is only valid when the type
	// is "block" or "network" for protocol='iscsi' or when the type is "volume" when using an iSCSI source pool
	// for mode "host" or as an NPIV virtual Host Bus Adapter (vHBA) using a Fibre Channel storage pool.
	// Configured in this manner, the LUN behaves identically to "disk", except that generic SCSI commands from
	// the guest are accepted and passed through to the physical device. Also note that device='lun' will only
	// be recognized for actual raw devices, but never for individual partitions or LVM partitions (in those cases,
	// the kernel will reject the generic SCSI commands, making it identical to device='disk').
	// Since 0.1.4
	Device DomainDiskDevice `xml:"device,attr,omitempty" json:"device,omitempty"`

	// Indicates the emulated device model of the disk. Typically this is indicated solely by the bus property
	// but for bus "virtio" the model can be specified further with "virtio-transitional", "virtio-non-transitional",
	// or "virtio". See Virtio transitional devices for more details. Since 5.2.0
	Model string `xml:"model,attr,omitempty" json:"model,omitempty"`

	// Indicates whether the disk needs rawio capability. Valid settings are "yes" or "no" (default is "no").
	// If any one disk in a domain has rawio='yes', rawio capability will be enabled for all disks in the domain
	// (because, in the case of QEMU, this capability can only be set on a per-process basis). This attribute is
	// only valid when device is "lun". NB, rawio intends to confine the capability per-device, however, current
	// QEMU implementation gives the domain process broader capability than that (per-process basis, affects all
	// the domain disks). To confine the capability as much as possible for QEMU driver as this stage, sgio is
	// recommended, it's more secure than rawio. Since 0.9.10
	RawIO ButtonState `xml:"rawio,attr,omitempty" json:"rawio,omitempty"`

	// If supported by the hypervisor and OS, indicates whether unprivileged SG_IO commands are filtered for the
	// disk. Valid settings are "filtered" or "unfiltered" where the default is "filtered". Only available when
	// the device is 'lun'. Since 1.0.2
	SgIO DomainDiskSgio `xml:"sgio,attr,omitempty" json:"sgio,omitempty"`

	// Indicates the default behavior of the disk during disk snapshots: "internal" requires a file format such
	// as qcow2 that can store both the snapshot and the data changes since the snapshot; "external" will separate
	// the snapshot from the live data; and "no" means the disk will not participate in snapshots. Read-only disks
	// default to "no", while the default for other disks depends on the hypervisor's capabilities. Some hypervisors
	// allow a per-snapshot choice as well, during domain snapshot creation. Not all snapshot modes are supported;
	// for example, enabling snapshots with a transient disk generally does not make sense. Since 0.9.5
	Snapshot DomainDiskSnapshot `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty"`

	// The optional driver element allows specifying further details related to the hypervisor driver
	// used to provide the disk
	Driver *DomainDriver `xml:"driver,attr,omitempty" json:"driver,omitempty"`

	// Representation of the disk source depends on the disk type
	Source *DomainDiskSource `xml:"source,omitempty" json:"source,omitempty"`

	// This element describes the backing store used by the disk specified by sibling source element. Since 1.2.4.
	// If the hypervisor driver does not support the backingStoreInput ( Since 5.10.0 ) domain feature the backingStore
	// is ignored on input and only used for output to describe the detected backing chains of running domains.
	// If backingStoreInput is supported the backingStore is used as the backing image of source or other
	// backingStore overriding any backing image information recorded in the image metadata. An empty backingStore
	// element means the sibling source is self-contained and is not based on any backing store. For the detected
	// backing chain information to be accurate, the backing format must be correctly specified in the metadata of
	// each file of the chain (files created by libvirt satisfy this property, but using existing external files
	// for snapshot or block copy operations requires the end user to pre-create the file correctly).
	BackingStore *DomainDiskBackingStore `xml:"backingStore,omitempty" json:"backingStore,omitempty"`

	// The optional backenddomain element allows specifying a backend domain (aka driver domain) hosting the disk.
	// Use the name attribute to specify the backend domain name. Since 1.2.13 (Xen only)
	BackendDomain *DomainBackendDomain `xml:"backenddomain,omitempty" json:"backenddomain,omitempty"`

	Geometry *DomainDiskGeometry `xml:"geometry,omitempty" json:"geometry,omitempty"`

	BlockIO *DomainDiskBlockIO `xml:"blockio,omitempty" json:"blockio,omitempty"`

	// This element is present if the hypervisor has started a long-running block job operation, where the mirror
	// location in the source sub-element will eventually have the same contents as the source, and with the file
	// format in the sub-element format (which might differ from the format of the source). The details of the
	// source sub-element are determined by the type attribute of the mirror, similar to what is done for the
	// overall disk device element. The job attribute mentions which API started the operation ("copy" for the
	// virDomainBlockRebase API, or "active-commit" for the virDomainBlockCommit API), since 1.2.7 . The attribute
	// ready, if present, tracks progress of the job: yes if the disk is known to be ready to pivot, or, since 1.2.7 ,
	// abort or pivot if the job is in the process of completing. If ready is not present, the disk is probably still
	// copying. For now, this element only valid in output; it is ignored on input. The source sub-element exists for
	// all two-phase jobs since 1.2.6 . Older libvirt supported only block copy to a file, since 0.9.12 ; for
	// compatibility with older clients, such jobs include redundant information in the attributes file and format
	// in the mirror element.
	Mirror *DomainDiskMirror `xml:"mirror,omitempty" json:"mirror,omitempty"`

	// The target element controls the bus / device under which the disk is exposed to the guest OS. The dev
	// attribute indicates the "logical" device name. The actual device name specified is not guaranteed to
	// map to the device name in the guest OS. Treat it as a device ordering hint. The optional bus attribute
	// specifies the type of disk device to emulate; possible values are driver specific, with typical values
	// being "ide", "scsi", "virtio", "xen", "usb", "sata", or "sd" "sd" since 1.1.2 . If omitted, the bus
	// type is inferred from the style of the device name (e.g. a device named 'sda' will typically be exported
	// using a SCSI bus). The optional attribute tray indicates the tray status of the removable disks
	// (i.e. CDROM or Floppy disk), the value can be either "open" or "closed", defaults to "closed". NB,
	// the value of tray could be updated while the domain is running. The optional attribute removable sets
	// the removable flag for USB disks, and its value can be either "on" or "off", defaulting to "off". The
	// optional attribute rotation_rate sets the rotation rate of the storage for disks on a SCSI, IDE, or
	// SATA bus. Values in the range 1025 to 65534 are used to indicate rotational media speed in revolutions
	// per minute. A value of one is used to indicate solid state, or otherwise non-rotational, storage. These
	// values are not required to match the values of the underlying host storage. Since 0.0.3; bus attribute
	// since 0.4.3; tray attribute since 0.9.11; "usb" attribute value since after 0.4.4; "sata" attribute value
	// since 0.9.7; "removable" attribute value since 1.1.3; "rotation_rate" attribute value since 7.3.0
	Target *DomainDiskTarget `xml:"target,omitempty" json:"target,omitempty"`

	// The optional iotune element provides the ability to provide additional per-device I/O tuning, with values
	// that can vary for each device (contrast this to the <blkiotune> element, which applies globally to the
	// domain). Currently, the only tuning available is Block I/O throttling for qemu. This element has optional
	// sub-elements; any sub-element not specified or given with a value of 0 implies no limit. Since 0.9.8
	IOTune *DomainDiskIOTune `xml:"iotune,omitempty" json:"iotune,omitempty"`

	// If present, this indicates the device cannot be modified by the guest. For now, this is the default for
	// disks with attribute device='cdrom'.
	ReadOnly *Empty `xml:"readonly,omitempty" json:"readonly,omitempty"`

	// If present, this indicates the device is expected to be shared between domains (assuming the hypervisor
	// and OS support this), which means that caching should be deactivated for that device.
	Shareable *Empty `xml:"shareable,omitempty" json:"shareable,omitempty"`

	// If present, this indicates that changes to the device contents should be reverted automatically when the
	// guest exits. With some hypervisors, marking a disk transient prevents the domain from participating in migration,
	// snapshots, or blockjobs. Only supported in vmx hypervisor (Since 0.9.5) and qemu hypervisor (Since 6.9.0).
	Transient *Empty `xml:"transient,omitempty" json:"transient,omitempty"`

	// If present, this specify serial number of virtual hard drive. For example, it may look like
	// <serial>WD-WMAP9A966149</serial>. Not supported for scsi-block devices, that is those using
	// disk type 'block' using device 'lun' on bus 'scsi'. Since 0.7.1
	Serial string `xml:"serial,omitempty" json:"serial,omitempty"`

	// If present, this element specifies the WWN (World Wide Name) of a virtual hard disk or CD-ROM drive.
	// It must be composed of 16 hexadecimal digits. Since 0.10.1
	WWN string `xml:"wwn,omitempty" json:"wwn,omitempty"`

	// If present, this element specifies the vendor of a virtual hard disk or CD-ROM device. It must not be longer
	// than 8 printable characters. Since 1.0.1
	Vendor string `xml:"vendor,omitempty" json:"vendor,omitempty"`

	// If present, this element specifies the product of a virtual hard disk or CD-ROM device. It must not be longer
	// than 16 printable characters. Since 1.0.1
	Product string `xml:"product,omitempty" json:"product,omitempty"`

	// Starting with libvirt 3.9.0 the encryption element is preferred to be a sub-element of the source element.
	// If present, specifies how the volume is encrypted using "qcow". See the Storage Encryption page for more
	// information.
	Encryption *StorageEncryption `xml:"encryption,omitempty" json:"encryption,omitempty"`

	// Specifies that the disk is bootable. The order attribute determines the order in which devices will be
	// tried during boot sequence. On the S390 architecture only the first boot device is used. The optional
	// loadparm attribute is an 8 character string which can be queried by guests on S390 via sclp or diag
	// 308. Linux guests on S390 can use loadparm to select a boot entry. Since 3.5.0 The per-device boot
	// elements cannot be used together with general boot elements in BIOS bootloader section. Since 0.8.8
	Boot *DomainDiskBoot `xml:"boot,omitempty" json:"boot,omitempty"`

	Alias *DomainAlias `xml:"alias,omitempty" json:"alias,omitempty"`

	// If present, the address element ties the disk to a given slot of a controller (the actual <controller>
	// device can often be inferred by libvirt, although it can be explicitly specified). The type attribute is
	// mandatory, and is typically "pci" or "drive". For a "pci" controller, additional attributes for bus,
	// slot, and function must be present, as well as optional domain and multifunction. Multifunction defaults
	// to 'off'; any other value requires QEMU 0.1.3 and libvirt 0.9.7 . For a "drive" controller, additional
	// attributes controller, bus, target ( libvirt 0.9.11 ), and unit are available, each defaulting to 0.
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainAlias struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainDeviceACPI struct {
	Index uint `xml:"index,attr,omitempty"`
}

type DomainDriverName string

const (
	DomainDriverNameTap       DomainDriverName = "tap"
	DomainDriverNameTap2      DomainDriverName = "tap2"
	DomainDriverNamePhy       DomainDriverName = "phy"
	DomainDriverNameFile      DomainDriverName = "file"
	DomainDriverNameQemu      DomainDriverName = "qemu"
	DomainDriverNameVhost     DomainDriverName = "vhost"
	DomainDriverNameVhostUser DomainDriverName = "vhostuser"
)

type DomainDriverType string

const (
	DomainDiskDriverTypeAio   DomainDriverType = "aio"
	DomainDiskDriverTypeRaw   DomainDriverType = "raw"
	DomainDiskDriverTypeBochs DomainDriverType = "bochs"
	DomainDiskDriverTypeQcow2 DomainDriverType = "qcow2"
	DomainDiskDriverTypeQed   DomainDriverType = "qed"
)

type DomainDriverCache string

const (
	DomainDriverTypeNone         DomainDriverType = "none"
	DomainDriverTypeDefault      DomainDriverType = "default"
	DomainDriverTypeWriteThrough DomainDriverType = "writethroug"
	DomainDriverTypeWriteBack    DomainDriverType = "writeback"
	DomainDriverTypeDirectSync   DomainDriverType = "directsync"
)

type DomainDriverErrorPolicy string

const (
	DomainDriverErrorPolicyStop     DomainDriverErrorPolicy = "stop"
	DomainDriverErrorPolicyReport   DomainDriverErrorPolicy = "report"
	DomainDriverErrorPolicyIgnore   DomainDriverErrorPolicy = "ignore"
	DomainDriverErrorPolicyEnospace DomainDriverErrorPolicy = "enospace"
)

type DomainDriverIO string

const (
	DomainDriverIOThreads DomainDriverIO = "threads"
	DomainDriverIONative  DomainDriverIO = "native"
	DomainDriverIOIOUring DomainDriverIO = "io_uring"
)

type DomainDriverDiscard string

const (
	DomainDriverDiscardUnmap  DomainDriverDiscard = "unmap"
	DomainDriverDiscardIgnore DomainDriverDiscard = "ignore"
)

type DomainDriverDetectZeroes string

const (
	DomainDriverDetectZeroesOn    DomainDriverDetectZeroes = "on"
	DomainDriverDetectZeroesOff   DomainDriverDetectZeroes = "off"
	DomainDriverDetectZeroesUnmap DomainDriverDetectZeroes = "unmap"
)

type DomainDriver struct {
	// If the hypervisor supports multiple backend drivers, then the name attribute selects the primary backend
	// driver name, while the optional type attribute provides the sub-type. For example, xen supports a name of
	// "tap", "tap2", "phy", or "file", with a type of "aio", while qemu only supports a name of "qemu", but
	// multiple types including "raw", "bochs", "qcow2", and "qed".
	Name DomainDriverName `xml:"name,attr,omitempty" json:"name,omitempty"`
	Type DomainDriverType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// The optional cache attribute controls the cache mechanism, possible values are "default", "none",
	// "writethrough", "writeback", "directsync" (like "writethrough", but it bypasses the host page cache)
	// and "unsafe" (host may cache all disk io, and sync requests from guest are ignored). Since 0.6.0,
	// "directsync" since 0.9.5, "unsafe" since 0.9.7
	Cache DomainDriverCache `xml:"cache,attr,omitempty" json:"cache,omitempty"`

	// The optional error_policy attribute controls how the hypervisor will behave on a disk read or write error,
	// possible values are "stop", "report", "ignore", and "enospace". Since 0.8.0, "report" since 0.9.7 The default
	// is left to the discretion of the hypervisor. There is also an optional rerror_policy that controls behavior for
	// read errors only. Since 0.9.7 . If no rerror_policy is given, error_policy is used for both read and write errors.
	// rerror_policy is given, it overrides the error_policy for read errors. Also note that "enospace" is not a valid
	// policy for read errors, so if error_policy is set to "enospace" and no rerror_policy is given, the read error
	// policy will be left at its default.
	ErrorPolicy  DomainDriverErrorPolicy `xml:"error_policy,attr,omitempty" json:"errorPolicy,omitempty"`
	RErrorPolicy DomainDriverErrorPolicy `xml:"rerror_policy,attr,omitempty" json:"rerrorPolicy,omitempty"`

	// The optional io attribute controls specific policies on I/O; qemu guests support "threads" and "native"
	// Since 0.8.8 , io_uring Since 6.3.0 (QEMU 5.0) .
	IO DomainDriverIO `xml:"io,attr,omitempty" json:"io,omitempty"`

	// The optional ioeventfd attribute allows users to set domain I/O asynchronous handling for disk device.
	// The default is left to the discretion of the hypervisor. Accepted values are "on" and "off". Enabling
	// this allows qemu to execute VM while a separate thread handles I/O. Typically guests experiencing high
	// system CPU utilization during I/O will benefit from this. On the other hand, on overloaded host it
	// could increase guest I/O latency. Since 0.9.3 (QEMU and KVM only) In general you should leave this
	// option alone, unless you are very certain you know what you are doing.
	IOEventFd TurnState `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`

	// The optional event_idx attribute controls some aspects of device event processing. The value can be
	// either 'on' or 'off' - if it is on, it will reduce the number of interrupts and exits for the guest.
	// The default is determined by QEMU; usually if the feature is supported, default is on. In case there
	// is a situation where this behavior is suboptimal, this attribute provides a way to force the feature
	// off. Since 0.9.5 (QEMU and KVM only) In general you should leave this option alone, unless you are
	// very certain you know what you are doing.
	EventIdx TurnState `xml:"event_idx,attr,omitempty" json:"eventIdx,omitempty"`

	// The optional copy_on_read attribute controls whether to copy read backing file into the image file.
	// The value can be either "on" or "off". Copy-on-read avoids accessing the same backing file sectors
	// repeatedly and is useful when the backing file is over a slow network. By default copy-on-read is off.
	// Since 0.9.10 (QEMU and KVM only)
	CopyOnRead TurnState `xml:"copy_on_read,attr,omitempty" json:"copyOnRead,omitempty"`

	// The optional discard attribute controls whether discard requests (also known as "trim" or "unmap")
	// are ignored or passed to the filesystem. The value can be either "unmap" (allow the discard request
	// to be passed) or "ignore" (ignore the discard request). Since 1.0.6 (QEMU and KVM only)
	Discard DomainDriverDiscard `xml:"discard,attr,omitempty" json:"discard,omitempty"`

	// The optional detect_zeroes attribute controls whether to detect zero write requests. The value
	// can be "off", "on" or "unmap". First two values turn the detection off and on, respectively.
	// The third value ("unmap") turns the detection on and additionally tries to discard such areas
	// from the image based on the value of discard above (it will act as "on" if discard is set to
	// "ignore"). NB enabling the detection is a compute intensive operation, but can save file space
	// and/or time on slow media. Since 2.0.0
	DetectZeroes DomainDriverDetectZeroes `xml:"detect_zeroes,attr,omitempty" json:"detectZeroes,omitempty"`

	// The optional iothread attribute assigns the disk to an IOThread as defined by the range for the domain iothreads
	// value. Multiple disks may be assigned to the same IOThread and are numbered from 1 to the domain iothreads value.
	// Available for a disk device target configured to use "virtio" bus and "pci" or "ccw" address types.
	// Since 1.2.8 (QEMU 2.1)
	IOThread int32 `xml:"iothread,attr,omitempty" json:"iothread,omitempty"`

	// The optional queues attribute specifies the number of virt queues for virtio-blk. ( Since 3.9.0 )
	Queues int32 `xml:"queues,attr,omitempty" json:"queues,omitempty"`

	// The optional metadata_cache subelement controls aspects related to the format specific caching of
	// storage image metadata. Note that this setting applies only on the top level image; the identically
	// named subelement of backingStore's format element can be used to specify cache settings for the backing image.
	//
	// Since 7.0.0 the maximum size of the metadata cache of qcow2 format driver of the qemu hypervisor can be
	// controlled via the max_size subelement (see example below).
	//
	//In the majority of cases the default configuration used by the hypervisor is sufficient so modifying this
	// setting should not be necessary. For specifics on how the metadata cache of qcow2 in qemu behaves refer
	// to the qemu qcow2 cache docs
	MetadataCache *DomainDriverMetadataCache `xml:"metadata_cache,omitempty" json:"metadataCache,omitempty"`

	// The optional rx_queue_size attribute controls the size of virtio ring for each queue as described above.
	// The default value is hypervisor dependent and may change across its releases. Moreover, some hypervisors
	// may pose some restrictions on actual value. For instance, latest QEMU (as of 2016-09-01) requires value
	// to be a power of two from [256, 1024] range. Since 2.3.0 (QEMU and KVM only) In general you should leave
	// this option alone, unless you are very certain you know what you are doing.
	RxQueueSize int64 `xml:"rx_queue_size,attr,omitempty" json:"rxQueueSize,omitempty"`

	// The optional tx_queue_size attribute controls the size of virtio ring for each queue as described above.
	// The default value is hypervisor dependent and may change across its releases. Moreover, some hypervisors
	// may pose some restrictions on actual value. For instance, QEMU v2.9 requires value to be a power of two
	// from [256, 1024] range. In addition to that, this may work only for a subset of interface types, e.g.
	// aforementioned QEMU enables this option only for vhostuser type. Since 3.7.0 (QEMU and KVM only) In general
	// you should leave this option alone, unless you are very certain you know what you are doing.
	TxQueueSize int64 `xml:"txQueueSize,omitempty" json:"txQueueSize,omitempty"`

	// The csum, gso, tso4, tso6, ecn and ufo attributes with possible values on and off can be used to turn
	// off host offloading options. By default, the supported offloads are enabled by QEMU. Since 1.2.9 (QEMU
	// only) The mrg_rxbuf attribute can be used to control mergeable rx buffers on the host side. Possible
	// values are on (default) and off. Since 1.2.13 (QEMU only)
	Host *DomainDriverHost `xml:"host,omitempty" json:"host,omitempty"`

	// The csum, tso4, tso6, ecn and ufo attributes with possible values on and off can be used to turn off guest
	// offloading options. By default, the supported offloads are enabled by QEMU. Since 1.2.9 (QEMU only)
	Guest *DomainDriverGuest `xml:"guest,omitempty" json:"guest,omitempty"`
}

type DomainDriverMetadataCache struct {
	MaxSize Size `xml:"max_size,omitempty" json:"max_size,omitempty"`
}

type DomainDriverHost struct {
	Csum TurnState `xml:"csum,attr,omitempty" json:"csum,omitempty"`
	Gso  TurnState `xml:"gso,attr,omitempty" json:"gso,omitempty"`
	Tso4 TurnState `xml:"tso4,attr,omitempty" json:"tso4,omitempty"`
	Tso6 TurnState `xml:"tso6,attr,omitempty" json:"tso6,omitempty"`
	Ecn  TurnState `xml:"ecn,attr,omitempty" json:"ecn,omitempty"`
	Ufo  TurnState `xml:"ufo,attr,omitempty" json:"ufo,omitempty"`
}

type DomainDriverGuest struct {
	Csum TurnState `xml:"csum,attr,omitempty" json:"csum,omitempty"`
	Gso  TurnState `xml:"gso,attr,omitempty" json:"gso,omitempty"`
	Tso4 TurnState `xml:"tso4,attr,omitempty" json:"tso4,omitempty"`
	Tso6 TurnState `xml:"tso6,attr,omitempty" json:"tso6,omitempty"`
	Ecn  TurnState `xml:"ecn,attr,omitempty" json:"ecn,omitempty"`
	Ufo  TurnState `xml:"ufo,attr,omitempty" json:"ufo,omitempty"`
}

type DomainDiskSourceProtocol string

const (
	DomainDiskSourceProtocolNbd      DomainDiskSourceProtocol = "nbd"
	DomainDiskSourceProtocolIscsi    DomainDiskSourceProtocol = "iscsi"
	DomainDiskSourceProtocolRbd      DomainDiskSourceProtocol = "rbd"
	DomainDiskSourceProtocolSheepDog DomainDiskSourceProtocol = "sheepdog"
	DomainDiskSourceProtocolGluster  DomainDiskSourceProtocol = "gluster"
	DomainDiskSourceProtocolVxhs     DomainDiskSourceProtocol = "vxhs"
	DomainDiskSourceProtocolNfs      DomainDiskSourceProtocol = "nfs"
	DomainDiskSourceProtocolHttp     DomainDiskSourceProtocol = "http"
	DomainDiskSourceProtocolHttps    DomainDiskSourceProtocol = "https"
	DomainDiskSourceProtocolFtp      DomainDiskSourceProtocol = "ftp"
	DomainDiskSourceProtocolFtps     DomainDiskSourceProtocol = "ftps"
	DomainDiskSourceProtocolTftp     DomainDiskSourceProtocol = "tftp"
)

type DomainDiskSourceMode string

const (
	DomainDiskSourceModeHost   DomainDiskSourceMode = "host"
	DomainDiskSourceModeDirect DomainDiskSourceMode = "direct"
)

type DomainDiskSourceType string

const (
	DomainDiskSourceTypePCI  DomainDiskSourceType = "pci"
	DomainDiskSourceTypeUnit DomainDiskSourceType = "unix"
)

type DomainDiskSourceStartupPolicy string

const (
	DomainDiskSourceStartupPolicyMandatory DomainDiskSourceStartupPolicy = "mandatory"
	DomainDiskSourceStartupPolicyRequisite DomainDiskSourceStartupPolicy = "requisite"
	DomainDiskSourceStartupPolicyOptional  DomainDiskSourceStartupPolicy = "optional"
)

type DomainDiskSource struct {
	// (type=file) The file attribute specifies the fully-qualified path to the file holding the disk. Since 0.0.3
	File string `xml:"file,attr,omitempty" json:"file,omitempty"`

	// (type=block)The dev attribute specifies the fully-qualified path to the host device to serve as the disk. Since 0.0.3
	Block string `xml:"block,attr,omitempty" json:"block,omitempty"`

	// (type=dir) The dir attribute specifies the fully-qualified path to the directory to use as the disk. Since 0.7.5
	Dir string `xml:"dir,attr,omitempty" json:"dir,omitempty"`

	// (type=network) The protocol attribute specifies the protocol to access to the requested image.
	// Possible values are "nbd", "iscsi", "rbd", "sheepdog", "gluster", "vxhs", "nfs", "http", "https", "ftp",
	// ftps", or "tftp".
	//
	// For any protocol other than nbd an additional attribute name is mandatory to specify which volume/image
	// will be used.
	//
	// For "nbd", the name attribute is optional. TLS transport for NBD can be enabled by setting the tls attribute
	// to yes. For the QEMU hypervisor, usage of a TLS environment can also be globally controlled on the host by
	// the nbd_tls and nbd_tls_x509_cert_dir in /etc/libvirt/qemu.conf. ('tls' Since 4.5.0 )
	//
	// For protocols http and https an optional attribute query specifies the query string. ( Since 6.2.0 )
	//
	// For "iscsi" ( since 1.0.4 ), the name attribute may include a logical unit number, separated from the target's
	// name by a slash (e.g., iqn.2013-07.com.example:iscsi-pool/1). If not specified, the default LUN is zero.
	//
	// For "vxhs" ( since 3.8.0 ), the name is the UUID of the volume, assigned by the HyperScale server. Additionally,
	// an optional attribute tls (QEMU only) can be used to control whether a VxHS block device would utilize a
	// hypervisor configured TLS X.509 certificate environment in order to encrypt the data channel. For the QEMU
	// hypervisor, usage of a TLS environment can also be globally controlled on the host by the vxhs_tls and
	// vxhs_tls_x509_cert_dir or default_tls_x509_cert_dir settings in the file /etc/libvirt/qemu.conf. If vxhs_tls
	// is enabled, then unless the domain tls attribute is set to "no", libvirt will use the host configured TLS
	// environment. If the tls attribute is set to "yes", then regardless of the qemu.conf setting, TLS authentication
	// will be attempted.
	Protocol DomainDiskSourceProtocol `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	Name     string                   `xml:"name,attr,omitempty" json:"name,omitempty"`

	Query string `xml:"query,attr,omitempty" json:"query,omitempty"`

	// (type=volume) The underlying disk source is represented by attributes pool and volume. Attribute pool specifies the name
	// of the storage pool (managed by libvirt) where the disk source resides. Attribute volume specifies the name
	// of storage volume (managed by libvirt) used as the disk source. The value for the volume attribute will be
	// the output from the "Name" column of a virsh vol-list [pool-name] command.
	//
	// Use the attribute mode ( since 1.1.1 ) to indicate how to represent the LUN as the disk source. Valid values
	// are "direct" and "host". If mode is not specified, the default is to use "host". Using "direct" as the mode
	// value indicates to use the storage pool's source element host attribute as the disk source to generate the
	// libiscsi URI (e.g. 'file=iscsi://example.com:3260/iqn.2013-07.com.example:iscsi-pool/1'). Using "host" as
	// the mode value indicates to use the LUN's path as it shows up on host
	// (e.g. 'file=/dev/disk/by-path/ip-example.com:3260-iscsi-iqn.2013-07.com.example:iscsi-pool-lun-1').
	// Using a LUN from an iSCSI source pool provides the same features as a disk configured using type 'block'
	// or 'network' and device of 'lun' with respect to how the LUN is presented to and may be used by the guest.
	// Since 1.0.5
	Pool   string               `xml:"pool,attr,omitempty" json:"pool,omitempty"`
	Volume string               `xml:"volume,attr,omitempty" json:"volume,omitempty"`
	Mode   DomainDiskSourceMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`

	// pci or unix
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`

	// (type=nvme) To specify disk source for NVMe disk the source element has the following attributes:
	// The type of address specified in address sub-element. Currently, only pci value is accepted.
	// The difference between <disk type='nvme'> and <hostdev/> is that the latter is plain host device
	// assignment with all its limitations (e.g. no live migration), while the former makes hypervisor to
	// run the NVMe disk through hypervisor's block layer thus enabling all features provided by the layer
	// (e.g. snapshots, domain migration, etc.). Moreover, since the NVMe disk is unbinded from its PCI driver,
	// the host kernel storage stack is not involved (compared to passing say /dev/nvme0n1 via <disk type='block'>
	// and therefore lower latencies can be achieved.
	//
	// This attribute instructs libvirt to detach NVMe controller automatically on domain startup (yes) or
	// expect the controller to be detached by system administrator (no).
	Managed ButtonState `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	// The namespace ID which should be assigned to the domain. According to NVMe standard, namespace numbers
	// start from 1, including.
	Namespace int32                `xml:"namespace,attr,omitempty" json:"namespace,omitempty"`
	Address   *DomainDeviceAddress `xml:"address,attr,omitempty" json:"address,omitempty"`

	// (type=vhostuser) Enables the hypervisor to connect to another process using vhost-user protocol.
	// Requires shared memory configured for the VM, for more details see access mode for memoryBacking element.
	// The source element has following mandatory attributes:
	// Type: The type of char device. Currently only unix type is supported.
	// Path to the unix socket to be used as disk source.
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`

	// For a "file" or "volume" disk type which represents a cdrom or floppy (the device attribute), it is possible
	// to define policy what to do with the disk if the source file is not accessible. (NB, startupPolicy is not
	// valid for "volume" disk unless the specified storage volume is of "file" type). This is done by the
	// startupPolicy attribute ( since 0.9.7 ), accepting these values:
	//  mandatory  : fail if missing for any reason (the default)
	//  requisite  : fail if missing on boot up, drop if missing on migrate/restore/revert
	//  optional   : drop if missing at any start attempt
	//
	// Since 1.1.2 the startupPolicy is extended to support hard disks besides cdrom and floppy. On guest cold
	// bootup, if a certain disk is not accessible or its disk chain is broken, with startupPolicy 'optional'
	// the guest will drop this disk. This feature doesn't support migration currently.
	StartupPolicy DomainDiskSourceStartupPolicy `xml:"startupPolicy,attr,omitempty" json:"startupPolicy,omitempty"`

	// When the disk type is "network", the source may have zero or more host sub-elements used to specify the
	// hosts to connect. The host element supports 4 attributes, viz. "name", "port", "transport" and "socket",
	// which specify the hostname, the port number, transport type and path to socket, respectively.
	Host *DomainDiskSourceHost `xml:"host,omitempty" json:"host,omitempty"`

	// The name attribute of snapshot element can optionally specify an internal snapshot name to be used
	// as the source for storage protocols. Supported for 'rbd' since 1.2.11 (QEMU only).
	Snapshot *DomainDiskSourceSnapshot `xml:"snapshot,omitempty" json:"snapshot,omitempty"`

	// The file attribute for the config element provides a fully qualified path to a configuration file to
	// be provided as a parameter to the client of a networked storage protocol.
	// Supported for 'rbd' since 1.2.11 (QEMU only).
	Config *DomainDiskSourceConfig `xml:"config,omitempty" json:"config,omitempty"`

	// Since libvirt 3.9.0 , the auth element is supported for a disk type "network" that is using a source
	// element with the protocol attributes "rbd" or "iscsi". If present, the auth element provides the
	// authentication credentials needed to access the source. It includes a mandatory attribute username,
	// which identifies the username to use during authentication, as well as a sub-element secret with mandatory
	// attribute type, to tie back to a libvirt secret object that holds the actual password or other credentials
	// (the domain XML intentionally does not expose the password, only the reference to the object that does manage
	// the password). Known secret types are "ceph" for Ceph RBD network sources and "iscsi" for CHAP authentication
	// of iSCSI targets. Both will require either a uuid attribute with the UUID of the secret object or a usage
	// attribute matching the key that was specified in the secret object.
	Auth *DomainDiskSourceAuth `xml:"auth,attr,omitempty" json:"auth,omitempty"`

	// Since libvirt 3.9.0 , the encryption can be a sub-element of the source element for encrypted storage sources.
	// If present, specifies how the storage source is encrypted See the Storage Encryption page for more information.
	// Note that the 'qcow' format of encryption is broken and thus is no longer supported for use with disk images.
	// ( Since libvirt 4.5.0 )
	Encryption *StorageEncryption `xml:"encryption,omitempty" json:"encryption,omitempty"`

	// Since libvirt 4.4.0 , the reservations can be a sub-element of the source element for storage sources
	// (QEMU driver only). If present it enables persistent reservations for SCSI based disks. The element has
	// one mandatory attribute managed with accepted values yes and no. If managed is enabled libvirt prepares
	// and manages any resources needed. When the persistent reservations are unmanaged, then the hypervisor
	// acts as a client and the path to the server socket must be provided in the child element source, which
	// currently accepts only the following attributes: type with one value unix, path path to the socket, and
	// finally mode which accepts one value client specifying the role of hypervisor. It's recommended to allow
	// libvirt manage the persistent reservations.
	Reservations *DomainDiskSourceReservations `xml:"reservations,omitempty" json:"reservations,omitempty"`

	// Since libvirt 4.7.0 , the initiator element is supported for a disk type "network" that is using a source
	// element with the protocol attribute "iscsi". If present, the initiator element provides the initiator IQN
	// needed to access the source via mandatory attribute name.
	Initiator *DomainDiskSourceInitiator `xml:"initiator,omitempty" json:"initiator,omitempty"`

	// The slices element using its slice sub-elements allows configuring offset and size of either the location
	// of the image format (slice type='storage') inside the storage source or the guest data inside the image
	// format container (future expansion). The offset and size values are in bytes. Since 6.1.0
	Slice []*DomainDiskSourceSlice `xml:"slice,omitempty" json:"slice,omitempty"`

	// For https and ftps accessed storage it's possible to tweak the SSL transport parameters with this element.
	// The verify attribute allows to turn on or off SSL certificate validation. Supported values are yes and no.
	// Since 6.2.0
	SSL *DomainDiskSourceSSL `xml:"ssl,omitempty" json:"ssl,omitempty"`

	// For http and https accessed storage it's possible to pass one or more cookies. The cookie name and value
	// must conform to the HTTP specification. Since 6.2.0
	Cookies []*DomainDiskSourceCookie `xml:"cookies,omitempty" json:"cookies,omitempty"`

	// Specifies the size of the readahead buffer for protocols which support it. (all 'curl' based drivers in qemu).
	// The size is in bytes. Note that '0' is considered as if the value is not provided. Since 6.2.0
	Readahead *DomainDiskSourceReadahead `xml:"readahead,omitempty" json:"readahead,omitempty"`

	// Specifies the connection timeout for protocols which support it. Note that '0' is considered as if the
	// value is not provided. Since 6.2.0
	Timeout *DomainDiskSourceTimeout `xml:"timeout,omitempty" json:"timeout,omitempty"`

	// When using an nfs protocol, this is used to provide information on the configuration of the user and group.
	// The element has two attributes, user and group. The user can provide these elements as user or group strings,
	// or as user and group ID numbers directly if the string is formatted using a "+" at the beginning of the ID
	// number. If either of these attributes is omitted, then that field is assumed to be the default value for the
	// current system. If both user and group are intended to be default, then the entire element may be omitted.
	Identity *DomainDiskSourceIdentity `xml:"identity,omitempty" json:"identity,omitempty"`

	// For disk type vhostuser configures reconnect timeout if the connection is lost.
	Reconnect *DomainDiskSourceReconnect `xml:"reconnect,omitempty" json:"reconnect,omitempty"`
}

type DomainDiskSourceHostTransport string

const (
	DomainDiskSourceHostTransportTcp  DomainDiskSourceHostTransport = "tcp"
	DomainDiskSourceHostTransportRdma DomainDiskSourceHostTransport = "rdma"
	DomainDiskSourceHostTransportUnix DomainDiskSourceHostTransport = "unix"
)

type DomainDiskSourceHostSocket string

const (
	DomainDiskSourceHostTransportAFUNIX DomainDiskSourceHostTransport = "AF_UNIX"
)

type DomainDiskSourceHost struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Port int32  `xml:"port,attr,omitempty" json:"port,omitempty"`
	// gluster supports "tcp", "rdma", "unix" as valid values for the transport attribute. nbd supports
	// "tcp" and "unix". Others only support "tcp". If nothing is specified, "tcp" is assumed. If the
	// transport is "unix"
	Transport DomainDiskSourceHostTransport `xml:"transport,attr,omitempty" json:"transport,omitempty"`
	// the socket attribute specifies the path to an AF_UNIX socket. nfs only supports the use of a "tcp" transport,
	// and does not support using a port at all so it must be omitted.
	Socket DomainDiskSourceHostSocket `xml:"socket,attr,omitempty" json:"socket,omitempty"`
}

type DomainDiskSourceSnapshot struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainDiskSourceConfig struct {
	File string `xml:"file,attr,omitempty" json:"file,omitempty"`
}

type DomainDiskSourceAuth struct {
	Username string `xml:"username,attr,omitempty" json:"usernamem,omitempty"`

	Secret *DomainDiskSourceAuthSecret `xml:"secret,attr,omitempty" json:"secret,omitempty"`
}

type DomainDiskSourceAuthSecretType string

const (
	DomainDiskSourceAuthSecretTypeCeph  DomainDiskSourceAuthSecretType = "ceph"
	DomainDiskSourceAuthSecretTypeIscsi DomainDiskSourceAuthSecretType = "iscsi"
)

type DomainDiskSourceAuthSecret struct {
	Type  DomainDiskSourceAuthSecretType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Usage string                         `xml:"usage,attr,omitempty" json:"usage,omitempty"`
}

type DomainDiskSourceReservations struct {
	Managed ButtonState                         `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	Source  *DomainDiskSourceReservationsSource `xml:"source,omitempty" json:"source,omitempty"`
}

type DomainDiskSourceReservationsSource struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainDiskSourceInitiator struct {
	Iqn *DomainDiskSourceIqn `xml:"iqn,omitempty" json:"iqn,omitempty"`
}

type DomainDiskSourceIqn struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainDiskSourceSliceType string

const (
	DomainDiskSourceSliceTypeStorage DomainDiskSourceSliceType = "storage"
)

type DomainDiskSourceSlice struct {
	Type   DomainDiskSourceSliceType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Offset int32                     `xml:"offset,attr,omitempty" json:"offset,omitempty"`
	Size   int64                     `xml:"size,attr,omitempty" json:"size,omitempty"`
}

type DomainDiskSourceSSL struct {
	Verify ButtonState `xml:"verify,attr,omitempty" json:"verify,omitempty"`
}

type DomainDiskSourceCookie struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Value string `xml:",chardata" json:"value"`
}

type DomainDiskSourceReadahead struct {
	Size int64 `xml:"size,attr" json:"size"`
}

type DomainDiskSourceTimeout struct {
	Seconds int64 `xml:"seconds,attr" json:"seconds"`
}

type DomainDiskSourceIdentity struct {
	User  string `xml:"user,attr,omitempty" json:"user,omitempty"`
	Group string `xml:"group,attr,omitempty" json:"group,omitempty"`
}

type DomainDiskSourceReconnect struct {
	// If the reconnect feature is enabled, accepts yes and no
	Enable ButtonState `xml:"enable,attr,omitempty" json:"enable,omitempty"`
	// The amount of seconds after which hypervisor tries to reconnect.
	Timeout int32 `xml:"timeout,attr,omitempty" json:"timeout,omitempty"`
}

type DomainDiskBackingStoreType string

const (
	DomainDiskBackingStoreTypeFile  DomainDiskBackingStoreType = "file"
	DomainDiskBackingStoreTypeBlock DomainDiskBackingStoreType = "block"
)

type DomainDiskBackingStore struct {
	// The type attribute represents the type of disk used by the backing store, see disk type attribute above
	// for more details and possible values.
	Type DomainDiskBackingStoreType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// This attribute is only valid in output (and ignored on input) and it can be used to refer to a specific
	// part of the disk chain when doing block operations (such as via the virDomainBlockRebase API). For example,
	// vda[2] refers to the backing store with index='2' of the disk with vda target.
	Index int32 `xml:"index,attr,omitempty" json:"index,omitempty"`

	// The format element contains type attribute which specifies the internal format of the backing store,
	// such as raw or qcow2.
	//
	// The format element can contain metadata_cache subelement, which has identical semantics to the identically
	// named subelement of driver of a disk.
	Format *DomainDiskFormat `xml:"format,omitempty" json:"format,omitempty"`

	// This element has the same structure as the source element in disk. It specifies which file, device, or
	// network location contains the data of the described backing store.
	Source *DomainDiskBackingStoreSource `xml:"source,omitempty" json:"source,omitempty"`

	// If the backing store is not self-contained, the next element in the chain is described by nested
	// backingStore element.
	BackingStore *DomainDiskBackingStore `xml:"backingStore,omitempty" json:"backingStore,omitempty"`
}

type DomainDiskBackingStoreFormatType string

const (
	DomainDiskBackingStoreFormatTypeBlock DomainDiskBackingStoreFormatType = "block"
)

type DomainDiskFormat struct {
	Type          DomainDiskBackingStoreFormatType `xml:"type,attr,omitempty" json:"type,omitempty"`
	MetadataCache *DomainDriverMetadataCache       `xml:"metadata_cache,omitempty" json:"metadataCache,omitempty"`
}

type DomainDiskBackingStoreSource struct {
	Dev string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
}

type DomainDiskMirror struct {
	Job          string                  `xml:"job,attr,omitempty"`
	Ready        string                  `xml:"ready,attr,omitempty"`
	Format       *DomainDiskFormat       `xml:"format"`
	Source       *DomainDiskSource       `xml:"source"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore"`
}

type DomainDiskTargetBus string

const (
	DomainDiskTargetBusIde    DomainDiskTargetBus = "ide"
	DomainDiskTargetBusScsi   DomainDiskTargetBus = "scsi"
	DomainDiskTargetBusVirtio DomainDiskTargetBus = "virtio"
	DomainDiskTargetBusXen    DomainDiskTargetBus = "xen"
	DomainDiskTargetBusUSB    DomainDiskTargetBus = "usb"
	DomainDiskTargetBusSata   DomainDiskTargetBus = "sata"
	DomainDiskTargetBusSd     DomainDiskTargetBus = "sd"
)

type DomainDiskTargetTray string

const (
	DomainDiskTargetTrayOpen   DomainDiskTargetTray = "open"
	DomainDiskTargetTrayClosed DomainDiskTargetTray = "closed"
)

type DomainDiskTarget struct {
	Dev          string               `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Bus          DomainDiskTargetBus  `xml:"bus,attr,omitempty" json:"bus,omitempty"`
	Tray         DomainDiskTargetTray `xml:"tray,attr,omitempty" json:"tray,omitempty"`
	Removable    TurnState            `xml:"removable,attr,omitempty" json:"removable,omitempty"`
	RotationRate int32                `xml:"rotation_rate,attr,omitempty" json:"rotationRate,omitempty"`
}

type DomainDiskIOTune struct {
	// The optional total_bytes_sec element is the total throughput limit in bytes per second. This cannot
	// appear with read_bytes_sec or write_bytes_sec.
	TotalBytesSec int64 `xml:"total_bytes_sec,omitempty" json:"totalBytesSec,omitempty"`
	// The optional read_bytes_sec element is the read throughput limit in bytes per second.
	ReadBytesSec int64 `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty"`
	// The optional write_bytes_sec element is the write throughput limit in bytes per second.
	WriteBytesSec int64 `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty"`
	// The optional total_iops_sec element is the total I/O operations per second. This cannot appear with
	// read_iops_sec or write_iops_sec.
	TotalIOPSSec int64 `xml:"total_iops_sec,omitempty" json:"totalIopsSec,omitempty"`
	// The optional read_iops_sec element is the read I/O operations per second.
	ReadIOPSSec int64 `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty"`
	// The optional write_iops_sec element is the write I/O operations per second.
	WriteIOPSSec int64 `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty"`
	// The optional total_bytes_sec_max element is the maximum total throughput limit in bytes per second.
	// This cannot appear with read_bytes_sec_max or write_bytes_sec_max.
	TotalBytesSecMax int64 `xml:"total_bytes_sec_max,omitempty" json:"totalBytesSecMax,omitempty"`
	// The optional read_bytes_sec_max element is the maximum read throughput limit in bytes per second.
	ReadBytesSecMax int64 `xml:"read_bytes_sec_max,omitempty" json:"readBytesSecMax,omitempty"`
	// The optional write_bytes_sec_max element is the maximum write throughput limit in bytes per second.
	WriteBytesSecMax int64 `xml:"write_bytes_sec_max,omitempty" json:"writeBytesSecMax,omitempty"`
	// The optional total_iops_sec_max element is the maximum total I/O operations per second. This cannot
	// appear with read_iops_sec_max or write_iops_sec_max.
	TotalIOPSSecMax int64 `xml:"total_iops_sec_max,omitempty" json:"totalIopsSecMax,omitempty"`
	// The optional read_iops_sec_max element is the maximum read I/O operations per second.
	ReadIOPSSecMax int64 `xml:"read_iops_sec_max,omitempty" json:"readIopsSecMax,omitempty"`
	// The optional write_iops_sec_max element is the maximum write I/O operations per second.
	WriteIOPSSecMax int64 `xml:"write_iops_sec_max,omitempty" json:"writeIopsSecMax,omitempty"`
	// The optional size_iops_sec element is the size of I/O operations per second.
	SizeIOPSSec int64 `xml:"size_iops_sec,omitempty" json:"sizeIopsSec,omitempty"`
	// The optional group_name provides the cability to share I/O throttling quota between multiple drives.
	// This prevents end-users from circumventing a hosting provider's throttling policy by splitting 1 large
	// drive in N small drives and getting N times the normal throttling quota. Any name may be used.
	GroupName string `xml:"group_name,omitempty" json:"groupName,omitempty"`
	// The optional total_bytes_sec_max_length element is the maximum duration in seconds for the total_bytes_sec_max
	// burst period. Only valid when the total_bytes_sec_max is set.
	TotalBytesSecMaxLength int64 `xml:"total_bytes_sec_max_length,omitempty" json:"totalBytesSecMaxLength,omitempty"`
	// The optional read_bytes_sec_max_length element is the maximum duration in seconds for the read_bytes_sec_max
	// burst period. Only valid when the read_bytes_sec_max is set.
	ReadBytesSecMaxLength int64 `xml:"read_bytes_sec_max_length,omitempty" json:"readBytesSecMaxLength,omitempty"`
	// The optional write_bytes_sec_max_length element is the maximum duration in seconds for the write_bytes_sec_max
	// burst period. Only valid when the write_bytes_sec_max is set.
	WriteBytesSecMaxLength int64 `xml:"write_bytes_sec_max_length,omitempty" json:"writeBytesSecMaxLength,omitempty"`
	// The optional total_iops_sec_max_length element is the maximum duration in seconds for the total_iops_sec_max
	// burst period. Only valid when the total_iops_sec_max is set.
	TotalIOPSSecMaxLength int64 `xml:"total_iops_sec_max_length,omitempty" json:"totalIopsSecMaxLength,omitempty"`
	// The optional read_iops_sec_max_length element is the maximum duration in seconds for the read_iops_sec_max
	// burst period. Only valid when the read_iops_sec_max is set.
	ReadIOPSSecMaxLength int64 `xml:"read_iops_sec_max_length,omitempty" json:"readIopsSecMaxLength,omitempty"`
	// The optional write_iops_sec_max_length element is the maximum duration in seconds for the write_iops_sec_max
	// burst period. Only valid when the write_iops_sec_max is set.
	WriteIOPSSecMaxLength int64 `xml:"write_iops_sec_max_length,omitempty" json:"writeIopsSecMaxLength,omitempty"`
}

type DomainBackendDomain struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainDiskBoot struct {
	Order    uint32 `xml:"order,attr,omitempty" json:"order,omitempty"`
	LoadParm string `xml:"loadparm,attr,omitempty" json:"loadparm,omitempty"`
}

type DomainDiskGeometry struct {
	Cylinders uint   `xml:"cyls,attr,omitempty" json:"cylinders,omitempty"`
	Headers   uint   `xml:"heads,attr,omitempty" json:"headers,omitempty"`
	Sectors   uint   `xml:"secs,attr,omitempty" json:"sectors,omitempty"`
	Trans     string `xml:"trans,attr,omitempty" json:"trans,omitempty"`
}

type DomainDiskBlockIO struct {
	LogicalBlockSize  int64 `xml:"logical_block_size,attr,omitempty" json:"logicalBlockSize,omitempty"`
	PhysicalBlockSize int64 `xml:"physical_block_size,attr,omitempty" json:"physicalBlockSize,omitempty"`
}

type DomainFilesystemType string

const (
	// DomainFilesystemTypeTemplate A host directory to mount in the guest. Used by LXC, OpenVZ
	// (since 0.6.2) and QEMU/KVM (since 0.8.5) . This is the default type if one is not specified.
	// This mode also has an optional sub-element driver, with an attribute type='path' or type='handle'
	// (since 0.9.7) . The driver block has an optional attribute wrpolicy that further controls interaction
	// with the host page cache; omitting the attribute gives default behavior, while the value immediate
	// means that a host writeback is immediately triggered for all pages touched during a guest file write
	// operation (since 0.9.10) . Since 6.2.0 , type='virtiofs' is also supported. Using virtiofs requires
	// setting up shared memory, see the guide: Virtio-FS
	DomainFilesystemTypeTemplate DomainFilesystemType = "template"

	// DomainFilesystemTypeMount OpenVZ filesystem template. Only used by OpenVZ driver.
	DomainFilesystemTypeMount DomainFilesystemType = "mount"

	// DomainFilesystemTypeFile a host file will be treated as an image and mounted in the guest. The
	// filesystem format will be autodetected. Only used by LXC driver.
	DomainFilesystemTypeFile DomainFilesystemType = "file"

	// DomainFilesystemTypeBlock a host block device to mount in the guest. The filesystem format will be
	// autodetected. Only used by LXC driver (since 0.9.5) .
	DomainFilesystemTypeBlock DomainFilesystemType = "block"

	// DomainFilesystemTypeRam an in-memory filesystem, using memory from the host OS. The source element
	// has a single attribute usage which gives the memory usage limit in KiB, unless units are specified
	// by the units attribute. Only used by LXC driver. (since 0.9.13)
	DomainFilesystemTypeRam DomainFilesystemType = "ram"

	// DomainFilesystemTypeBind a directory inside the guest will be bound to another directory inside the guest.
	// Only used by LXC driver (since 0.9.13)
	DomainFilesystemTypeBind DomainFilesystemType = "bind"
)

type DomainFilesystemAccessMode string

const (
	// DomainFilesystemAccessModePassthrough the source is accessed with the permissions of the user inside the guest.
	// This is the default accessmode if one is not specified.
	DomainFilesystemAccessModePassthrough DomainFilesystemAccessMode = "passthrough"

	// DomainFilesystemAccessModeMapped the source is accessed with the permissions of the hypervisor (QEMU process).
	DomainFilesystemAccessModeMapped DomainFilesystemAccessMode = "mapped"

	// DomainFilesystemAccessModeSquash similar to 'passthrough', the exception is that failure of privileged
	// operations like 'chown' are ignored. This makes a passthrough-like mode usable for people who run the
	// hypervisor as non-root.
	DomainFilesystemAccessModeSquash DomainFilesystemAccessMode = "squash"
)

type DomainFilesystemMultidevs string

const (
	// DomainFilesystemMultidevsDefault use QEMU's default setting (which currently is warn).
	DomainFilesystemMultidevsDefault DomainFilesystemMultidevs = "default"

	// DomainFilesystemMultidevsRemap this setting allows guest to access multiple devices per
	// export without encountering misbehaviours. Inode numbers from host are automatically remapped
	// on guest to actively prevent file ID collisions if guest accesses one export containing multiple devices.
	DomainFilesystemMultidevsRemap DomainFilesystemMultidevs = "remap"

	// DomainFilesystemMultidevsForbid only allow to access one device per export by guest. Attempts to access
	// additional devices on the same export will cause the individual filesystem access by guest to fail with
	// an error and being logged (once) as error on host side.
	DomainFilesystemMultidevsForbid DomainFilesystemMultidevs = "forbid"

	// DomainFilesystemMultidevsWarn this setting resembles the behaviour of 9pfs prior to QEMU 4.2, that is
	// no action is performed to prevent any potential file ID collisions if an export contains multiple devices,
	// with the only exception: a warning is logged (once) on host side now. This setting may lead to misbehaviours
	// on guest side if more than one device is exported per export, due to the potential file ID collisions this may
	// cause on guest side in that case.
	DomainFilesystemMultidevsWarn DomainFilesystemMultidevs = "warn"
)

type DomainFilesystem struct {
	// The filesystem attribute type specifies the type of the source.
	Type DomainFilesystemType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// The filesystem element has an optional attribute accessmode which specifies the security mode for accessing
	// the source (since 0.8.5) . Currently this only works with type='mount' for the QEMU/KVM driver. For driver
	// type virtiofs, only passthrough is supported.
	AccessMode DomainFilesystemAccessMode `xml:"accessmode,attr,omitempty" json:"accessmode,omitempty"`

	// Since 5.2.0 , the filesystem element has an optional attribute model with supported values "virtio-transitional",
	// "virtio-non-transitional", or "virtio". See Virtio transitional devices for more details.
	Model DomainVirtioModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	// The filesystem element has optional attributes fmode and dmode. These two attributes control the creation mode
	// for files and directories when used with the mapped value for accessmode (since 6.10.0, requires QEMU 2.10 ).
	// If not specified, QEMU creates files with mode 600 and directories with mode 700. The setuid, setgid, and sticky
	// bit are unsupported.
	FMode int32 `xml:"fmode,attr,omitempty" json:"fmode,omitempty"`
	DMode int32 `xml:"dmode,attr,omitempty" json:"dmode,omitempty"`

	// The filesystem element has an optional attribute multidevs which specifies how to deal with a filesystem
	// export containing more than one device, in order to avoid file ID collisions on guest when using
	// 9pfs ( since 6.3.0, requires QEMU 4.2 ).
	Multidevs DomainFilesystemMultidevs `xml:"multidevs,attr,omitempty" json:"multidevs,omitempty"`

	// The optional driver element allows specifying further details related to the hypervisor driver used to provide
	// the filesystem. Since 1.0.6
	//
	// If the hypervisor supports multiple backend drivers, then the type attribute selects the primary backend driver
	// name, while the format attribute provides the format type. For example, LXC supports a type of "loop", with a
	// format of "raw" or "nbd" with any format. QEMU supports a type of "path" or "handle", but no formats. Virtuozzo
	// driver supports a type of "ploop" with a format of "ploop".
	//
	// For virtio-backed devices, Virtio-specific options can also be set. ( Since 3.5.0 )
	//
	// For virtiofs, the queue attribute can be used to specify the queue size (i.e. how many requests can the
	// queue fit). ( Since 6.2.0 )
	Driver *DomainFilesystemDriver `xml:"driver,omitempty" json:"driver,omitempty"`

	// The optional binary element can tune the options for virtiofsd. All of the following attributes and elements
	// are optional. The attribute path can be used to override the path to the daemon. Attribute xattr enables the
	// use of filesystem extended attributes. Caching can be tuned via the cache element, possible mode values being
	// none and always. Locking can be controlled via the lock element - attributes posix and flock both accepting
	// values on or off. ( Since 6.2.0 ) The sandboxing method used by virtiofsd can be configured with the sandbox
	// element, possible mode values being namespace and chroot, see the virtiofsd documentation for more details.
	// ( Since 7.2.0 )
	Binary *DomainFilesystemBinary `xml:"binary,omitempty" json:"binary,omitempty"`

	// The resource on the host that is being accessed in the guest. The name attribute must be used with
	// type='template', and the dir attribute must be used with type='mount'. For virtiofs, the socket attribute
	// can be used to connect to a virtiofsd daemon launched outside of libvirt. In that case, the target element
	// does not apply and neither do most virtiofs-related options, since they are controlled by virtiofsd, not
	// libvirtd. The usage attribute is used with type='ram' to set the memory limit in KiB, unless units are
	// specified by the units attribute.
	Source *DomainFilesystemSource `xml:"source,omitempty" json:"source,omitempty"`

	// Where the source can be accessed in the guest. For most drivers this is an automatic mount point, but for
	// QEMU/KVM this is merely an arbitrary string tag that is exported to the guest as a hint for where to mount.
	Target *DomainFilesystemTarget `xml:"target,omitempty" json:"target,omitempty"`

	// Enables exporting filesystem as a readonly mount for guest, by default read-write access is given (currently
	// only works for QEMU/KVM driver).
	Readonly *Empty `xml:"readonly,omitempty" json:"readonly,omitempty"`

	// Maximum space available to this guest's filesystem. Since 0.9.13
	SpaceHardLimit int64 `xml:"space_hard_limit,omitempty" json:"spaceHardLimit,omitempty"`

	// Maximum space available to this guest's filesystem. The container is permitted to exceed its soft limits
	// for a grace period of time. Afterwards the hard limit is enforced. Since 0.9.13
	SpaceSoftLimit int64 `xml:"space_soft_limit,omitempty" json:"spaceSoftLimit,omitempty"`
}

type DomainFilesystemDriverType string

const (
	DomainFilesystemDriverTypePath     DomainFilesystemDriverType = "path"
	DomainFilesystemDriverTypeHandle   DomainFilesystemDriverType = "handle"
	DomainFilesystemDriverTypeLoop     DomainFilesystemDriverType = "loop"
	DomainFilesystemDriverTypePloop    DomainFilesystemDriverType = "ploop"
	DomainFilesystemDriverTypeVirtioFS DomainFilesystemDriverType = "virtiofs"
)

type DomainFilesystemDriverFormat string

const (
	DomainFilesystemDriverFormatRaw   DomainFilesystemDriverFormat = "raw"
	DomainFilesystemDriverFormatNbd   DomainFilesystemDriverFormat = "nbd"
	DomainFilesystemDriverFormatPloop DomainFilesystemDriverFormat = "ploop"
)

type DomainFilesystemDriverWrpolicy string

const (
	DomainFilesystemDriverWrpolicyImmediate DomainFilesystemDriverWrpolicy = "immediate"
)

type DomainFilesystemDriver struct {
	Type     DomainFilesystemDriverType     `xml:"type,attr,omitempty" json:"type,omitempty"`
	Format   DomainFilesystemDriverWrpolicy `xml:"format,attr,omitempty" json:"format,omitempty"`
	Wrpolicy string                         `xml:"wrpolicy,attr,omitempty" json:"wrpolicy,omitempty"`
	Queue    int32                          `xml:"queue,attr,omitempty" json:"queue,omitempty"`
}

type DomainFilesystemBinary struct {
	Path    string                         `xml:"path,attr,omitempty" json:"path,omitempty"`
	Xattr   TurnState                      `xml:"xattr,attr,omitempty" json:"xattr,omitempty"`
	Cache   *DomainFilesystemBinaryCache   `xml:"cache,omitempty" json:"cache,omitempty"`
	Sandbox *DomainFilesystemBinarySandbox `xml:"sandbox,omitempty" json:"sandbox,omitempty"`
	Lock    *DomainFilesystemBinaryLock    `xml:"lock,omitempty" json:"lock,omitempty"`
}

type DomainFilesystemBinaryCacheMode string

const (
	DomainFilesystemBinaryCacheModeNone   DomainFilesystemBinaryCacheMode = "none"
	DomainFilesystemBinaryCacheModeAlways DomainFilesystemBinaryCacheMode = "always"
)

type DomainFilesystemBinaryCache struct {
	Mode DomainFilesystemBinaryCacheMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainFilesystemBinarySandboxMode string

const (
	DomainFilesystemBinarySandboxModeNamespace DomainFilesystemBinarySandboxMode = "namespace"
	DomainFilesystemBinarySandboxModeChroot    DomainFilesystemBinarySandboxMode = "chroot"
)

type DomainFilesystemBinarySandbox struct {
	Mode DomainFilesystemBinarySandboxMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainFilesystemBinaryLock struct {
	Posix TurnState `xml:"posix,attr,omitempty" json:"posix,omitempty"`
	Flock TurnState `xml:"flock,attr,omitempty" json:"flock,omitempty"`
}

type DomainFilesystemSource struct {
	Name   string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Dir    string `xml:"dir,attr,omitempty" json:"dir,omitempty"`
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty"`
	Target string `xml:"target,attr,omitempty" json:"target,omitempty"`
}

type DomainFilesystemTarget struct {
	Dir string `xml:"dir,attr,omitempty" json:"dir,omitempty"`
}

type DomainControllerType string

const (
	DomainControllerTypeIde          DomainControllerType = "ide"
	DomainControllerTypeFdc          DomainControllerType = "fdc"
	DomainControllerTypeScsi         DomainControllerType = "scsi"
	DomainControllerTypeSata         DomainControllerType = "sata"
	DomainControllerTypeUSB          DomainControllerType = "usb"
	DomainControllerTypeCcid         DomainControllerType = "ccid"
	DomainControllerTypeVirtioSerial DomainControllerType = "virtio-serial"
	DomainControllerTypeXenbus       DomainControllerType = "xenbus"
	DomainControllerTypePci          DomainControllerType = "pci"
)

type DomainControllerModel string

const (
	DomainControllerModelAuto                  DomainControllerModel = "auto"                    // scsi model
	DomainControllerModelBuslogic              DomainControllerModel = "buslogic"                // scsi model
	DomainControllerModelIbmvscsi              DomainControllerModel = "ibmvscsi"                // scsi model
	DomainControllerModelIsilogic              DomainControllerModel = "lsilogic"                // scsi model
	DomainControllerModelIsisas1068            DomainControllerModel = "lsisas1068"              // scsi model
	DomainControllerModelIsisas1078            DomainControllerModel = "lsisas1078"              // scsi model
	DomainControllerModelVirtioScsi            DomainControllerModel = "virtio-scsi"             // scsi model
	DomainControllerModelVmpvscsi              DomainControllerModel = "vmpvscsi"                // scsi model
	DomainControllerModelVirtioTransitional    DomainControllerModel = "virtio-transitional"     // scsi model
	DomainControllerModelVirtioNonTransitional DomainControllerModel = "virtio-non-transitional" // scsi model
	DomainControllerModelNcr53c90              DomainControllerModel = "ncr53c90"                // scsi model

	DomainControllerModelPiix3Uhci     DomainControllerModel = "piix3-uhci"     // usb model
	DomainControllerModelPii4Uhci      DomainControllerModel = "piix4-uhci"     // usb model
	DomainControllerModelEhci          DomainControllerModel = "ehci"           // usb model
	DomainControllerModelIch9Echci1    DomainControllerModel = "ich9-echci1"    // usb model
	DomainControllerModelIch9Echci2    DomainControllerModel = "ich9-echci2"    // usb model
	DomainControllerModelIch9Echci3    DomainControllerModel = "ich9-echci3"    // usb model
	DomainControllerModelVt82c686bUhci DomainControllerModel = "vt82c686b-uhci" // usb model
	DomainControllerModelPciOhci       DomainControllerModel = "pci-ohci"       // usb model
	DomainControllerModelNecXhci       DomainControllerModel = "nec-xhci"       // usb model
	DomainControllerModelQusb1         DomainControllerModel = "qusb1"          // usb model
	DomainControllerModelQusb2         DomainControllerModel = "qusb2"          // usb model
	DomainControllerModelQemuXhci      DomainControllerModel = "qemu-xhci"      // usb model
	DomainControllerModelNone          DomainControllerModel = "none"           // usb model

	DomainControllerModelPiix3 DomainControllerModel = "piix3" // ide model
	DomainControllerModelPiix4 DomainControllerModel = "piix4" // ide model
	DomainControllerModelIch6  DomainControllerModel = "ich6"  // ide model

	DomainControllerModelPciRoot                  DomainControllerModel = "pci-root"                    // pci model
	DomainControllerModelPciBridge                DomainControllerModel = "pci-bridge"                  // pci model
	DomainControllerModelPcieRoot                 DomainControllerModel = "pcie-root"                   // pci model
	DomainControllerModelDmiToPciBridge           DomainControllerModel = "dmi-to-pci-bridge"           // pci model
	DomainControllerModelPcieRootPort             DomainControllerModel = "pcie-root-port"              // pci model
	DomainControllerModelPcieSwitchUpstreamPort   DomainControllerModel = "pcie-switch-upstream-port"   // pci model
	DomainControllerModelPcieSwitchDownstreamPort DomainControllerModel = "pcie-switch-downstream-port" // pci model
	DomainControllerModelPciExpanderBus           DomainControllerModel = "pci-expander-bus"            // pci model
	DomainControllerModelPcieExpanderBus          DomainControllerModel = "pcie-expander-bus"           // pci model
	DomainControllerModelPcieToPciBridge          DomainControllerModel = "pcie-to-pci-bridge"          // pci model
)

type DomainController struct {
	Type             DomainControllerType  `xml:"type,attr,omitempty" json:"type,omitempty"`
	Index            int32                 `xml:"index,attr,omitempty" json:"index,omitempty"`
	Model            DomainControllerModel `xml:"model,attr,omitempty" json:"model,omitempty"`
	Ports            int32                 `xml:"ports,attr,omitempty" json:"ports,omitempty"`
	Vectors          int32                 `xml:"vectors,attr,omitempty" json:"vectors,omitempty"`
	MaxGrantFrames   int64                 `xml:"maxGrantFrames,attr,omitempty" json:"maxGrantFrames,omitempty"`
	MaxEventChannels int64                 `xml:"maxEventChannels,attr,omitempty" json:"maxEventChannels,omitempty"`

	Alias   *DomainControllerAlias  `xml:"alias,omitempty" json:"alias,omitempty"`
	Driver  *DomainControllerDriver `xml:"driver,omitempty" json:"driver,omitempty"`
	Master  *DomainControllerMaster `xml:"master,omitempty" json:"master,omitempty"`
	Address *DomainDeviceAddress    `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainControllerAlias struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainControllerDriver struct {
	// The optional queues attribute specifies the number of queues for the controller. For best performance,
	// it's recommended to specify a value matching the number of vCPUs. Since 1.0.5 (QEMU and KVM only)
	Queues int32 `xml:"queues,attr,omitempty" json:"queues,omitempty"`

	// The optional cmd_per_lun attribute specifies the maximum number of commands that can be queued on devices
	// controlled by the host. Since 1.2.7 (QEMU and KVM only)
	CmdPerLun int32 `xml:"cmd_per_lun,attr,omitempty" json:"cmdPerLun,omitempty"`

	// The optional max_sectors attribute specifies the maximum amount of data in bytes that will be transferred
	// to or from the device in a single command. The transfer length is measured in sectors, where a sector is
	// 512 bytes. Since 1.2.7 (QEMU and KVM only)
	MaxSectors int32 `xml:"max_sectors,omitempty" json:"maxSectors,omitempty"`

	// The optional ioeventfd attribute specifies whether the controller should use I/O asynchronous handling or not.
	// Accepted values are "on" and "off". Since 1.2.18
	IOEventFD TurnState `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`

	// Supported for controller type scsi using model virtio-scsi for address types pci and ccw since 1.3.5 (QEMU 2.4) .
	// The optional iothread attribute assigns the controller to an IOThread as defined by the range for the domain
	// iothreads value. Each SCSI disk assigned to use the specified controller will utilize the same IOThread.
	// If a specific IOThread is desired for a specific SCSI disk, then multiple controllers must be defined each
	// having a specific iothread value. The iothread value must be within the range 1 to the domain iothreads value.
	IOThread int32 `xml:"iothread,attr,omitempty" json:"iothread,omitempty"`
}

type DomainControllerMaster struct {
	StartPort int32 `xml:"startPort,attr,omitempty" json:"startPort,omitempty"`
}

type DomainLease struct {
	// This is an arbitrary string, identifying the lockspace within which the key is held. Lock managers
	// may impose extra restrictions on the format, or length of the lockspace name.
	LockSpace string `xml:"lockspace,omitempty" json:"lockspace,omitempty"`

	// This is an arbitrary string, uniquely identifying the lease to be acquired. Lock managers may impose
	// extra restrictions on the format, or length of the key.
	Key string `xml:"key,omitempty" json:"key,omitempty"`

	// This is the fully qualified path of the file associated with the lockspace. The offset specifies where
	// the lease is stored within the file. If the lock manager does not require an offset, just pass 0.
	Target *DomainLeaseTarget `xml:"target,omitempty" json:"target,omitempty"`
}

type DomainLeaseTarget struct {
	Path   string `xml:"path,attr,omitempty" json:"path,omitempty"`
	Offset int32  `xml:"offset,attr,omitempty" json:"offset,omitempty"`
}

type DomainHostDevMode string

const (
	DomainHostDevModeSubsystem    DomainHostDevMode = "subsystem"
	DomainHostDevModeCapabilities DomainHostDevMode = "capabilities"
)

type DomainHostDevType string

const (
	DomainHostDevTypeUSB      DomainHostDevType = "usb"
	DomainHostDevTypePci      DomainHostDevType = "pci"
	DomainHostDevTypeScsi     DomainHostDevType = "scsi"
	DomainHostDevTypeScsiHost DomainHostDevType = "scsi_host"
	DomainHostDevTypeMdev     DomainHostDevType = "mdev"

	DomainHostDevTypeStorage DomainHostDevType = "storage"
	DomainHostDevTypeMisc    DomainHostDevType = "misc"
	DomainHostDevTypeNet     DomainHostDevType = "net"
)

type DomainHostDevModel string

const (
	DomainHostDevModelVirtioTransitional    DomainHostDevModel = "virtio-transitional"
	DomainHostDevModelVirtioNonTransitional DomainHostDevModel = "virtio-non-transitional"
	DomainHostDevModelVirtio                DomainHostDevModel = "virtio"
	DomainHostDevModelCapabilities          DomainHostDevModel = "capabilities"
)

type DomainHostDev struct {
	// the mode is always "subsystem"
	Mode DomainHostDevMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`

	// the type is one of the following values with additional attributes noted.
	//  usb 	: USB devices are detached from the host on guest startup and reattached after the guest exits
	//            or the device is hot-unplugged.

	//  pci 	: For PCI devices, when managed is "yes" it is detached from the host before being passed on to
	//            the guest and reattached to the host after the guest exits. If managed is omitted or "no", the
	//            user is responsible to call virNodeDeviceDetachFlags (or virsh nodedev-detach before starting
	//            the guest or hot-plugging the device and virNodeDeviceReAttach (or virsh nodedev-reattach) after
	//            hot-unplug or stopping the guest.

	//  scsi 	: For SCSI devices, user is responsible to make sure the device is not used by host. If supported
	//            by the hypervisor and OS, the optional sgio ( since 1.0.6 ) attribute indicates whether unprivileged
	//            SG_IO commands are filtered for the disk. Valid settings are "filtered" or "unfiltered", where the
	//            default is "filtered". The optional rawio ( since 1.2.9 ) attribute indicates whether the lun needs
	//            the rawio capability. Valid settings are "yes" or "no". See the rawio description within the disk
	//            section. If a disk lun in the domain already has the rawio capability, then this setting not
	//            required.
	//
	// scsi_host: since 2.5.0 For SCSI devices, user is responsible to make sure the device is not used by host.
	//            This type passes all LUNs presented by a single HBA to the guest. Since 5.2.0, the model attribute
	//            can be specified further with "virtio-transitional", "virtio-non-transitional", or "virtio". See
	//            Virtio transitional devices for more details.

	// mdev 	: For mediated devices ( Since 3.2.0 ) the model attribute specifies the device API which determines
	//            how the host's vfio driver will expose the device to the guest. Currently, model='vfio-pci',
	//            model='vfio-ccw' ( Since 4.4.0 ) and model='vfio-ap' ( Since 4.9.0 ) is supported. MDEV section
	//            provides more information about mediated devices as well as how to create mediated devices on the
	//            host. Since 4.6.0 (QEMU 2.12) an optional display attribute may be used to enable or disable support
	//            for an accelerated remote desktop backed by a mediated device (such as NVIDIA vGPU or Intel GVT-g)
	//            as an alternative to emulated video devices. This attribute is limited to model='vfio-pci' only.
	//            Supported values are either on or off (default is 'off'). It is required to use a graphical
	//            framebuffer in order to use this attribute, currently only supported with VNC, Spice and egl-headless
	//            graphics devices. Since version 5.10.0 , there is an optional ramfb attribute for devices with
	//            model='vfio-pci'. Supported values are either on or off (default is 'off'). When enabled, this
	//            attribute provides a memory framebuffer device to the guest. This framebuffer will be used as a boot
	//            display when a vgpu device is the primary display.
	//
	// Note: There are also some implications on the usage of guest's address type depending on the model attribute,
	// see the address element below.
	Type DomainHostDevType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// Note: The managed attribute is only used with type='pci' and is ignored by all the other device types,
	// thus setting managed explicitly with other than a PCI device has the same effect as omitting it. Similarly,
	// model attribute is only supported by mediated devices and ignored by all other device types.
	Manged ButtonState `xml:"manged,attr,omitempty" json:"manged,omitempty"`

	SgIO  DomainDiskSgio `xml:"sgio,attr,omitempty" json:"sgio,omitempty"`
	RawIO ButtonState    `xml:"rawio,attr,omitempty" json:"rawio,omitempty"`

	Model *DomainHostDevModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	Display TurnState `xml:"display,attr,omitempty" json:"display,omitempty"`
	RamFb   TurnState `xml:"ramfb,attr,omitempty" json:"ramfb,omitempty"`

	Source *DomainHostDevSource `xml:"source,omitempty" json:"source,omitempty"`

	// Indicates that the device is readonly, only supported by SCSI host device now. Since 1.0.6 (QEMU and KVM only)
	Readonly *Empty `xml:"readonly,omitempty" json:"readonly,omitempty"`

	// If present, this indicates the device is expected to be shared between domains (assuming the hypervisor
	// and OS support this). Only supported by SCSI host device. Since 1.0.6
	Shareable *Empty `xml:"shareable,omitempty" json:"shareable,omitempty"`

	// Specifies that the device is bootable. The order attribute determines the order in which devices will
	// be tried during boot sequence. The per-device boot elements cannot be used together with general boot
	// elements in BIOS bootloader section. Since 0.8.8 for PCI devices, Since 1.0.1 for USB devices.
	Boot *DomainDeviceBoot `xml:"boot,omitempty" json:"boot,omitempty"`

	// The rom element is used to change how a PCI device's ROM is presented to the guest. The optional bar
	// attribute can be set to "on" or "off", and determines whether or not the device's ROM will be visible
	// in the guest's memory map. (In PCI documentation, the "rombar" setting controls the presence of the
	// Base Address Register for the ROM). If no rom bar is specified, the qemu default will be used (older
	// versions of qemu used a default of "off", while newer qemus have a default of "on"). Since 0.9.7 (QEMU
	// and KVM only) . The optional file attribute contains an absolute path to a binary file to be presented
	// to the guest as the device's ROM BIOS. This can be useful, for example, to provide a PXE boot ROM for
	// a virtual function of an sr-iov capable ethernet device (which has no boot ROMs for the VFs). Since
	// 0.9.10 (QEMU and KVM only) . The optional enabled attribute can be set to no to disable PCI ROM loading
	// completely for the device; if PCI ROM loading is disabled through this attribute, attempts to tweak the
	// loading process further using the bar or file attributes will be rejected. Since 4.3.0 (QEMU and KVM only).
	Rom *DomainDeviceRom `xml:"rom,omitempty" json:"rom,omitempty"`

	// PCI devices can have an optional driver subelement that specifies which backend driver to use for PCI
	// device assignment. Use the name attribute to select either "vfio" (for the new VFIO device assignment
	// backend, which is compatible with UEFI SecureBoot) or "kvm" (the legacy device assignment handled directly
	// by the KVM kernel module) Since 1.0.5 (QEMU and KVM only, requires kernel 3.6 or newer) . When specified,
	// device assignment will fail if the requested method of device assignment isn't available on the host. When
	// not specified, the default is "vfio" on systems where the VFIO driver is available and loaded, and "kvm"
	// on older systems, or those where the VFIO driver hasn't been loaded Since 1.1.3 (prior to that the default
	// was always "kvm").
	Driver *DomainHostDevDriver `xml:"driver,omitempty" json:"driver,omitempty"`

	// The address element for USB devices has a bus and device attribute to specify the USB bus and device number
	// the device appears at on the host. The values of these attributes can be given in decimal, hexadecimal
	// (starting with 0x) or octal (starting with 0) form. For PCI devices the element carries 4 attributes allowing
	// to designate the device as can be found with the lspci or with virsh nodedev-list. For SCSI devices a 'drive'
	// address type must be used. For mediated devices, which are software-only devices defining an allocation of
	// resources on the physical parent device, the address type used must conform to the model attribute of element
	// hostdev, e.g. any address type other than PCI for vfio-pci device API or any address type other than CCW for
	// vfio-ccw device API will result in an error. See above for more details on the address element.
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`

	IP    []*DomainInterfaceIP    `xml:"ip,omitempty" json:"ip,omitempty"`
	Route []*DomainInterfaceRoute `xml:"route,omitempty" json:"route,omitempty"`
}

type DomainHostDevSourceStartupPolicy string

const (
	DomainHostDevSourceStartupPolicyMandatory DomainHostDevSourceStartupPolicy = "mandatory"
	DomainHostDevSourceStartupPolicyRequisite DomainHostDevSourceStartupPolicy = "requisite"
	DomainHostDevSourceStartupPolicyOptional  DomainHostDevSourceStartupPolicy = "optional"
)

type DomainHostDevSourceProtocol string

const (
	DomainHostDevSourceProtocolIscsi DomainHostDevSourceProtocol = "iscsi"
	DomainHostDevSourceProtocolVhost DomainHostDevSourceProtocol = "vhost"
)

type DomainHostDevSource struct {
	StartupPolicy DomainHostDevSourceStartupPolicy `xml:"startupPolicy,attr,omitempty" json:"startupPolicy,omitempty"`
	Vendor        *DomainHostDevSourceVendor       `xml:"vendor,omitempty" json:"vendor,omitempty"`
	Product       *DomainHostDevSourceProduct      `xml:"product,omitempty" json:"product,omitempty"`

	WriteFiltering ButtonState          `xml:"writeFiltering,attr,omitempty" json:"writeFiltering,omitempty"`
	Address        *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`

	Protocol DomainHostDevSourceProtocol `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	Name     string                      `xml:"name,attr,omitempty" json:"name,omitempty"`

	Host      *DomainDiskSourceHost      `xml:"host,omitempty" json:"host,omitempty"`
	Auth      *DomainDiskSourceAuth      `xml:"auth,omitempty" json:"auth,omitempty"`
	Initiator *DomainDiskSourceInitiator `xml:"initiator,omitempty" json:"initiator,omitempty"`

	Wwpn string `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty"`

	// The hostdev element is the main container for describing host devices. For block/character device passthrough
	// mode is always "capabilities" and type is "storage" for a block device, "misc" for a character device and "net"
	// for a host network interface.
	//
	// The source element describes the device as seen from the host. For block devices, the path to the block device
	// in the host OS is provided in the nested "block" element, while for character devices the "char" element is used.
	// For network interfaces, the name of the interface is provided in the "interface" element.
	Block     string `xml:"block,omitempty" json:"block,omitempty"`
	Char      string `xml:"char,omitempty" json:"char,omitempty"`
	Interface string `xml:"interface,omitempty" json:"interface,omitempty"`
}

type DomainHostDevSourceVendor struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

type DomainHostDevSourceProduct struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

type DomainDeviceBoot struct {
	Order int32 `xml:"order,attr,omitempty" json:"order,omitempty"`
}

type DomainDeviceRom struct {
	Bar  TurnState `xml:"bar,attr,omitempty" json:"bar,omitempty"`
	File string    `xml:"file,attr,omitempty" json:"file,omitempty"`
}

type DomainHostDevDriverName string

const (
	DomainHostDevDriverNameVfio DomainHostDevDriverName = "vfio"
	DomainHostDevDriverNameKVM  DomainHostDevDriverName = "kvm"
)

type DomainHostDevDriver struct {
	Name DomainHostDevDriverName `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainReDirDevBus string

const (
	DomainReDirDevBusTCP      DomainReDirDevBus = "tcp"
	DomainReDirDevBusSpiceVmc DomainReDirDevBus = "spicevmc"
)

// DomainReDirDev the redirdev element is the main container for describing redirected devices. bus must be "usb"
// for a USB device. An additional attribute type is required, matching one of the supported serial device types,
// to describe the host side of the tunnel; type='tcp' or type='spicevmc' (which uses the usbredir channel of a
// SPICE graphics device) are typical. The redirdev element has an optional sub-element <address> which can tie
// the device to a particular controller. Further sub-elements, such as <source>, may be required according to the
// given type, although a <target> sub-element is not required (since the consumer of the character device is the
// hypervisor itself, rather than a device visible in the guest).
type DomainReDirDev struct {
	Bus DomainReDirDevBus `xml:"bus,attr,omitempty" json:"bus,omitempty"`

	Source *DomainReDirDevSource `xml:"source,omitempty" json:"source,omitempty"`

	// Specifies that the device is bootable. The order attribute determines the order in which devices will
	// be tried during boot sequence. The per-device boot elements cannot be used together with general boot
	// elements in BIOS bootloader section. ( Since 1.0.1 )
	Boot *DomainReDirDevBoot `xml:"boot,omitempty" json:"boot,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainReDirDevSourceMode string

const (
	DomainReDirDevSourceModeConnect DomainReDirDevSourceMode = "connect"
)

type DomainReDirDevSource struct {
	Mode    DomainReDirDevSourceMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Host    string                   `xml:"host,attr,omitempty" json:"host,omitempty"`
	Service int32                    `xml:"service,attr,omitempty" json:"service,omitempty"`
}

type DomainReDirDevBoot struct {
	Order int32 `xml:"order,attr,omitempty" json:"order,omitempty"`
}

type DomainReDirFilter struct {
	UsbDev []*DomainReDirFilterUsbDev `xml:"usbdev,omitempty" json:"usbdev,omitempty"`
}

type DomainReDirFilterUsbDev struct {
	Class   string      `xml:"class,attr,omitempty" json:"class,omitempty"`
	Vendor  string      `xml:"vendor,attr,omitempty" json:"vendor,omitempty"`
	Product string      `xml:"product,attr,omitempty" json:"product,omitempty"`
	Version string      `xml:"version,attr,omitempty" json:"version,omitempty"`
	Allow   ButtonState `xml:"allow,attr,omitempty" json:"allow,omitempty"`
}

type DomainSmartCardMode string

const (
	DomainSmartCardModeHost             DomainSmartCardMode = "host"
	DomainSmartCardModeHostCertificates DomainSmartCardMode = "host-certificates"
	DomainSmartCardModePassthrough      DomainSmartCardMode = "passthrough"
)

type DomainSmartCardType string

const (
	DomainSmartCardTypeTCP      DomainSmartCardType = "tcp"
	DomainSmartCardTypeSpiceVmc DomainSmartCardType = "spicevmc"
)

// DomainSmartCard the <smartcard> element has a mandatory attribute mode. The following modes are supported;
// in each mode, the guest sees a device on its USB bus that behaves like a physical USB CCID (Chip/Smart Card
// Interface Device) card.
// 	host : The simplest operation, where the hypervisor relays all requests from the guest into direct access to
//	       the host's smartcard via NSS. No other attributes or sub-elements are required. See below about the use
//	       of an optional <address> sub-element.
//
//  host-certificates : Rather than requiring a smartcard to be plugged into the host, it is possible to provide
//                      three NSS certificate names residing in a database on the host. These certificates can be
//                      generated via the command certutil -d /etc/pki/nssdb -x -t CT,CT,CT -S -s CN=cert1 -n cert1,
//                      and the resulting three certificate names must be supplied as the content of each of three
//                      <certificate> sub-elements. An additional sub-element <database> can specify the absolute
//                      path to an alternate directory (matching the -d option of the certutil command when creating
//                      the certificates); if not present, it defaults to /etc/pki/nssdb.
//
//  passthrough : Rather than having the hypervisor directly communicate with the host, it is possible to tunnel
//                all requests through a secondary character device to a third-party provider (which may in turn be
//                talking to a smartcard or using three certificate files). In this mode of operation, an additional
//                attribute type is required, matching one of the supported serial device types, to describe the host
//                side of the tunnel; type='tcp' or type='spicevmc' (which uses the smartcard channel of a SPICE
//                graphics device) are typical. Further sub-elements, such as <source>, may be required according to
//                the given type, although a <target> sub-element is not required (since the consumer of the character
//                device is the hypervisor itself, rather than a device visible in the guest).
//
// Each mode supports an optional sub-element <address>, which fine-tunes the correlation between the smartcard and
// a ccid bus controller, documented above. For now, qemu only supports at most one smartcard, with an address of
// bus=0 slot=0.
type DomainSmartCard struct {
	Mode DomainSmartCardMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Type DomainSmartCardType `xml:"type,attr,omitempty" json:"type,omitempty"`

	Certificate []string `xml:"certificate,omitempty" json:"certificate,omitempty"`

	Database string `xml:"database,omitempty" json:"database,omitempty"`

	Source *DomainSmartCardSource `xml:"source,omitempty" json:"source,omitempty"`

	Protocol *DomainSmartCardProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainSmartCardSourceMode string

const (
	DomainSmartCardSourceBind DomainSmartCardSourceMode = "bind"
)

type DomainSmartCardSource struct {
	Mode    DomainSmartCardSourceMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Host    string                    `xml:"host,attr,omitempty" json:"host,omitempty"`
	Service int32                     `xml:"service,attr,omitempty" json:"service,omitempty"`
}

type DomainSmartCardProtocolType string

const (
	DomainSmartCardProtocolTypeRaw DomainSmartCardProtocolType = "raw"
)

type DomainSmartCardProtocol struct {
	Type DomainSmartCardProtocolType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainInterfaceType string

const (
	DomainInterfaceTypeDirect    DomainInterfaceType = "direct"
	DomainInterfaceTypeNetwork   DomainInterfaceType = "network"
	DomainInterfaceTypeBridge    DomainInterfaceType = "bridge"
	DomainInterfaceTypeUser      DomainInterfaceType = "user"
	DomainInterfaceTypeEthernet  DomainInterfaceType = "ethernet"
	DomainInterfaceTypeHostDev   DomainInterfaceType = "hostdev"
	DomainInterfaceTypeVdpa      DomainInterfaceType = "vdpa"
	DomainInterfaceTypeMcast     DomainInterfaceType = "mcast"
	DomainInterfaceTypeServer    DomainInterfaceType = "server"
	DomainInterfaceTypeClient    DomainInterfaceType = "client"
	DomainInterfaceTypeUDP       DomainInterfaceType = "udp"
	DomainInterfaceTypeVHostUser DomainInterfaceType = "vhostuser"
)

type DomainInterface struct {
	Type DomainInterfaceType `xml:"type,attr,omitempty" json:"type,omitempty"`

	TrustGuestRxFilters ButtonState `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty"`

	Driver *DomainDriver `xml:"driver,omitempty" json:"driver,omitempty"`

	Source  *DomainInterfaceSource  `xml:"source,omitempty" json:"source,omitempty"`
	Mac     *DomainInterfaceMac     `xml:"mac,omitempty" json:"mac,omitempty"`
	Teaming *DomainInterfaceTeaming `xml:"teaming,omitempty" json:"teaming,omitempty"`

	VirtualPort *DomainInterfaceMacVirtualPort `xml:"virtualport,omitempty" json:"virtualport,omitempty"`

	IP    []*DomainInterfaceIP    `xml:"ip,omitempty" json:"ip,omitempty"`
	Route []*DomainInterfaceRoute `xml:"route,omitempty" json:"route,omitempty"`

	Scripts     *DomainInterfaceScripts     `xml:"scripts,omitempty" json:"scripts,omitempty"`
	DownScripts *DomainInterfaceDownScripts `xml:"downscripts,omitempty" json:"downscripts,omitempty"`

	Target *DomainInterfaceTarget `xml:"target,omitempty" json:"target,omitempty"`
	Model  *DomainInterfaceModel  `xml:"model,omitempty" json:"model,omitempty"`

	Backend *DomainInterfaceBackend `xml:"backend,omitempty" json:"backend,omitempty"`
	Tune    *DomainInterfaceTune    `xml:"tune,omitempty" json:"tune,omitempty"`

	Guest *DomainInterfaceGuest `xml:"guest,omitempty" json:"guest,omitempty"`

	Boot *DomainBoot       `xml:"boot,omitempty" json:"boot,omitempty"`
	Rom  *DomainDeviceRom  `xml:"rom,omitempty" json:"rom,omitempty"`
	ACPI *DomainDeviceACPI `xml:"acpi,omitempty" json:"acpi,omitempty"`

	BackendDomain *DomainBackendDomain      `xml:"backenddomain,omitempty" json:"backenddomain,omitempty"`
	Bandwidth     *DomainInterfaceBandwidth `xml:"bandwidth,omitempty" json:"bandwidth,omitempty"`
	Vlan          *DomainInterfaceVlan      `xml:"vlan,omitempty" json:"vlan,omitempty"`

	Port *DomainInterfacePort `xml:"port,omitempty" json:"port,omitempty"`

	Link *DomainInterfaceLink `xml:"link,omitempty" json:"link,omitempty"`

	Mtu *DomainInterfaceMtu `xml:"mtu,omitempty" json:"mtu,omitempty"`

	Coalesce *DomainInterfaceCoalesce `xml:"coalesce,omitempty" json:"coalesce,omitempty"`

	FilterRef []*DomainFilterRef `xml:"filterref,omitempty" json:"filterRef,omitempty"`
}

type DomainInterfaceDriverName string

const (
	DomainInterfaceDriverNameVfio DomainInterfaceDriverName = "vfio"
	DomainInterfaceDriverNameKVM  DomainInterfaceDriverName = "kvm"
)

type DomainInterfaceDriver struct {
	Name DomainInterfaceDriverName `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainInterfaceSourceNetwork string

const (
	DomainInterfaceSourceNetworkDefault DomainInterfaceSourceNetwork = "default"
)

type DomainInterfaceSourceMode string

const (
	// DomainInterfaceSourceModeVepa all VMs' packets are sent to the external bridge. Packets whose destination
	// is a VM on the same host as where the packet originates from are sent back to the host by the VEPA capable
	// bridge (today's bridges are typically not VEPA capable).
	DomainInterfaceSourceModeVepa DomainInterfaceSourceMode = "vepa"

	// DomainInterfaceSourceModeBridge packets whose destination is on the same host as where they originate
	// from are directly delivered to the target macvtap device. Both origin and destination devices need to
	// be in bridge mode for direct delivery. If either one of them is in vepa mode, a VEPA capable bridge
	// is required.
	DomainInterfaceSourceModeBridge DomainInterfaceSourceMode = "bridge"

	// DomainInterfaceSourceModePrivate all packets are sent to the external bridge and will only be delivered
	// to a target VM on the same host if they are sent through an external router or gateway and that device
	// sends them back to the host. This procedure is followed if either the source or destination device is
	// in private mode.
	DomainInterfaceSourceModePrivate DomainInterfaceSourceMode = "private"

	// DomainInterfaceSourceModePassthrough this feature attaches a virtual function of a SRIOV capable NIC
	// directly to a VM without losing the migration capability. All packets are sent to the VF/IF of the
	// configured network device. Depending on the capabilities of the device additional prerequisites or
	// limitations may apply; for example, on Linux this requires kernel 2.6.38 or newer. Since 0.9.2
	DomainInterfaceSourceModePassthrough DomainInterfaceSourceMode = "passthrough"

	DomainInterfaceSourceModeNe2kPCI DomainInterfaceSourceMode = "ne2k_pci"
)

type DomainInterfaceSource struct {
	Dev       string                       `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Network   DomainInterfaceSourceNetwork `xml:"network,attr,omitempty" json:"network,omitempty"`
	PostGroup string                       `xml:"postgroup,attr,omitempty" json:"postgroup,omitempty"`
	Path      string                       `xml:"path,attr,omitempty" json:"path,omitempty"`
	Mode      DomainInterfaceSourceMode    `xml:"mode,attr,omitempty" json:"mode,omitempty"`

	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`

	AddressAttr string `xml:"address,attr,omitempty" json:"addressAttr,omitempty"`
	Port        int32  `xml:"port,attr,omitempty" json:"port,omitempty"`

	Local *DomainInterfaceSourceLocal `xml:"local,omitempty" json:"local,omitempty"`

	Reconnect *DomainInterfaceSourceReconnect `xml:"reconnect,omitempty" json:"reconnect,omitempty"`
}

type DomainInterfaceSourceLocal struct {
	AddressAttr string `xml:"address,attr,omitempty" json:"addressAttr,omitempty"`
	Port        int32  `xml:"port,attr,omitempty" json:"port,omitempty"`
}

type DomainInterfaceSourceReconnect struct {
	Enabled ButtonState `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`
	Timeout int32       `xml:"timeout,attr,omitempty" json:"timeout,omitempty"`
}

type DomainInterfaceMac struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
}

type DomainInterfaceTeamingType string

const (
	DomainInterfaceTeamingTypeTransient  DomainInterfaceTeamingType = "transient"
	DomainInterfaceTeamingTypePersistent DomainInterfaceTeamingType = "persistent"
)

type DomainInterfaceTeaming struct {
	Type       DomainInterfaceTeamingType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Persistent string                     `xml:"persistent,attr,omitempty" json:"persistent,omitempty"`
}

type DomainInterfaceMacVirtualPortType string

const (
	DomainInterfaceMacVirtualPortTypeOpenVSwitch DomainInterfaceMacVirtualPortType = "openvswitch"
	DomainInterfaceMacVirtualPortTypeMidonet     DomainInterfaceMacVirtualPortType = "midonet"
)

type DomainInterfaceMacVirtualPort struct {
	Type       DomainInterfaceMacVirtualPortType   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Parameters *DomainInterfaceMacVirtualPortParam `xml:"parameters,omitempty" json:"parameters,omitempty"`
}

type DomainInterfaceMacVirtualPortParamProfileId string

const (
	DomainInterfaceMacVirtualPortParamProfileIdMenial  DomainInterfaceMacVirtualPortParamProfileId = "menial"
	DomainInterfaceMacVirtualPortParamProfileIdFinance DomainInterfaceMacVirtualPortParamProfileId = "finance"
)

type DomainInterfaceMacVirtualPortParam struct {
	ProfileId DomainInterfaceMacVirtualPortParamProfileId `xml:"profileid,attr,omitempty" json:"profileid,omitempty"`

	InstanceId string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty"`

	// Please note that IEEE 802.1Qbg requires a non-zero value for the VLAN ID.
	//
	// The VSI Manager ID identifies the database containing the VSI type and instance definitions.
	// This is an integer value and the value 0 is reserved.
	ManagerId int32 `xml:"managerid,attr,omitempty" json:"managerid,omitempty"`

	// The VSI Type ID identifies a VSI type characterizing the network access. VSI types are typically
	// managed by network administrator. This is an integer value.
	TypeId int64 `xml:"typeid,attr,omitempty" json:"typeid,omitempty"`

	// The VSI Type Version allows multiple versions of a VSI Type. This is an integer value.
	TypeIdVersion int32 `xml:"typeidversion,attr,omitempty" json:"typeidversion,omitempty"`

	// The VSI Instance ID Identifier is generated when a VSI instance (i.e. a virtual interface of
	// a virtual machine) is created. This is a globally unique identifier.
	// InstanceId string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty"`
}

type DomainInterfaceIPFamily string

const (
	DomainInterfaceIPFamilyIPV4 DomainInterfaceIPFamily = "ipv4"
	DomainInterfaceIPFamilyIPV6 DomainInterfaceIPFamily = "ipv6"
)

type DomainInterfaceIP struct {
	Family  DomainInterfaceIPFamily `xml:"family,attr,omitempty" json:"family,omitempty"`
	Address string                  `xml:"address,attr,omitempty" json:"address,omitempty"`
	Prefix  int32                   `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	Peer    string                  `xml:"peer,attr,omitempty" json:"peer,omitempty"`
}

type DomainInterfaceRoute struct {
	Family  DomainInterfaceIPFamily `xml:"family,attr,omitempty" json:"family,omitempty"`
	Address string                  `xml:"address,attr,omitempty" json:"address,omitempty"`
	Prefix  int32                   `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	Peer    string                  `xml:"peer,attr,omitempty" json:"peer,omitempty"`
	Gateway string                  `xml:"gateway,attr,omitempty" json:"gateway,omitempty"`
}

type DomainInterfaceScripts struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`
}

type DomainInterfaceDownScripts struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`
}

type DomainInterfaceTarget struct {
	Dev     string      `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Managed ButtonState `xml:"managed,attr,omitempty" json:"managed,omitempty"`
}

type DomainInterfaceModelType string

const (
	DomainInterfaceModelTypeVirtio  DomainInterfaceModelType = "virtio"
	DomainInterfaceModelType8021Qbg DomainInterfaceModelType = "802.1Qbg"
	DomainInterfaceModelType8021Qbh DomainInterfaceModelType = "802.1Qbh"
)

type DomainInterfaceModel struct {
	Type DomainInterfaceModelType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainInterfaceBackend struct {
	Tap   string `xml:"tap,attr,omitempty" json:"tap,omitempty"`
	Vhost string `xml:"vhost,attr,omitempty" json:"vhost,omitempty"`
}

type DomainInterfaceTune struct {
	Sndbuf int64 `xml:"sndbuf,omitempty" json:"sndbuf,omitempty"`
}

type DomainInterfaceGuest struct {
	Dev string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
}

type DomainInterfaceBandwidth struct {
	Inbound  *DomainQos `xml:"inbound,omitempty" json:"inbound,omitempty"`
	Outbound *DomainQos `xml:"outbound,omitempty" json:"outbound,omitempty"`
}

type DomainInterfaceVlan struct {
	Trunk *ButtonState `xml:"trunk,attr,omitempty" json:"trunk,omitempty"`

	Tag []*DomainInterfaceVlanTag `xml:"tag,omitempty" json:"tag,omitempty"`
}

type DomainInterfacePort struct {
	Isolated ButtonState `xml:"isolated,attr,omitempty" json:"isolated,omitempty"`
}

type DomainInterfaceLinkState string

const (
	DomainInterfaceLinkStateUP   DomainInterfaceLinkState = "up"
	DomainInterfaceLinkStateDown DomainInterfaceLinkState = "down"
)

type DomainInterfaceLink struct {
	State DomainInterfaceLinkState `xml:"state,attr,omitempty" json:"state,omitempty"`
}

type DomainInterfaceMtu struct {
	Size int32 `xml:"size,attr,omitempty" json:"size,omitempty"`
}

type DomainInterfaceCoalesce struct {
	Tx *DomainInterfaceCoalesceTx `xml:"tx,omitempty" json:"tx,omitempty"`
}

type DomainInterfaceCoalesceTx struct {
	Frames *DomainInterfaceCoalesceTxFrames `xml:"frames,omitempty" json:"frames,omitempty"`
}

type DomainInterfaceCoalesceTxFrames struct {
	Max int32 `xml:"max,attr,omitempty" json:"max,omitempty"`
}

type DomainInterfaceVlanTagMode string

const (
	DomainInterfaceVlanTagModeTagged   DomainInterfaceVlanTagMode = "tagged"
	DomainInterfaceVlanTagModeUnTagged DomainInterfaceVlanTagMode = "untagged"
)

type DomainInterfaceVlanTag struct {
	Id int32 `xml:"id,attr,omitempty" json:"id,omitempty"`

	NativeMode DomainInterfaceVlanTagMode `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty"`
}

type DomainQos struct {
	Average int64 `xml:"average,attr,omitempty" json:"average,omitempty"`
	Peak    int64 `xml:"peak,attr,omitempty" json:"peak,omitempty"`
	Floor   int64 `xml:"floor,attr,omitempty" json:"floor,omitempty"`
	Burst   int64 `xml:"burst,attr,omitempty" json:"burst,omitempty"`
}

type DomainFilterRefFilter string

const (
	DomainFilterRefFilterBridge   DomainFilterRefFilter = "bridge"
	DomainFilterRefFilterNetwork  DomainFilterRefFilter = "network"
	DomainFilterRefFilterEthernet DomainFilterRefFilter = "ethernet"
)

type DomainFilterRef struct {
	Filter    DomainFilterRefFilter   `xml:"filter,attr,omitempty" json:"filter,omitempty"`
	Parameter []*DomainFilterRefParam `xml:"parameter,omitempty" json:"parameter,omitempty"`
}

type DomainFilterRefParamName string

const (
	DomainFilterRefParamNameIP      DomainFilterRefParamName = "IP"
	DomainFilterRefParamNameIP6Addr DomainFilterRefParamName = "IP6_ADDR"
	DomainFilterRefParamNameIP6Mask DomainFilterRefParamName = "IP6_MASK"
)

type DomainFilterRefParam struct {
	Name  DomainFilterRefParamName `xml:"name,attr,omitempty" json:"name,omitempty"`
	Value string                   `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type DomainInputType string

const (
	DomainInputTypeMouse       DomainInputType = "mouse"
	DomainInputTypeKeyBoard    DomainInputType = "keyboard"
	DomainInputTypeTablet      DomainInputType = "tablet"
	DomainInputTypePassthrough DomainInputType = "passthrough"
)

type DomainInputBus string

const (
	DomainInputBusVirtio DomainInputBus = "virtio"
	DomainInputBusUSB    DomainInputBus = "usb"
)

type DomainInput struct {
	Type DomainInputType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Bus  DomainInputBus  `xml:"bus,attr,omitempty" json:"bus,omitempty"`

	Source *DomainInputSource `xml:"source,omitempty" json:"source,omitempty"`
	Driver *DomainDriver      `xml:"driver,omitempty" json:"driver,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainInputSource struct {
	EvDev string `xml:"evdev,attr,omitempty" json:"evdev,omitempty"`
}

type DomainHubType string

const (
	DomainHubTypeUSB DomainHubType = "usb"
)

type DomainHub struct {
	Type DomainHubType `xml:"type,attr,omitempty" json:"type,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainGraphicsType string

const (
	DomainGraphicsTypeSDL         DomainGraphicsType = "sdl"
	DomainGraphicsTypeVNC         DomainGraphicsType = "vnc"
	DomainGraphicsTypeSpice       DomainGraphicsType = "spice"
	DomainGraphicsTypeRdp         DomainGraphicsType = "rdp"
	DomainGraphicsTypeDesktop     DomainGraphicsType = "desktop"
	DomainGraphicsTypeEglHeadless DomainGraphicsType = "egl-headless"
)

type DomainGraphicsSharePolicy string

const (
	DomainGraphicsSharePolicyAllowExclusive DomainGraphicsSharePolicy = "allow-exclusive"
	DomainGraphicsSharePolicyForceShare     DomainGraphicsSharePolicy = "force-share"
	DomainGraphicsSharePolicyIgnore         DomainGraphicsSharePolicy = "ignore"
)

type DomainGraphics struct {
	Type    DomainGraphicsType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Display string             `xml:"display,attr,omitempty" json:"display,omitempty"`

	AutoPort    ButtonState               `xml:"autoport,attr,omitempty" json:"autoport,omitempty"`
	Port        int32                     `xml:"port,attr,omitempty" json:"port,omitempty"`
	TlsPort     int32                     `xml:"tlsPort,attr,omitempty" json:"tlsPort,omitempty"`
	SharePolicy DomainGraphicsSharePolicy `xml:"sharePolicy,attr,omitempty" json:"sharePolicy,omitempty"`
	Listen      *DomainGraphicsListen     `xml:"listen,omitempty" json:"listen,omitempty"`

	Passwd string `xml:"passwd,attr,omitempty" json:"passwd,omitempty"`

	MultiUser ButtonState `xml:"multiUser,attr,omitempty" json:"multiUser,omitempty"`

	FullScreen ButtonState `xml:"fullScreen,attr,omitempty" json:"fullScreen,omitempty"`

	Channel      []*DomainGraphicsChannel    `xml:"channel,omitempty" json:"channel,omitempty"`
	Image        *DomainGraphicsImage        `xml:"image,omitempty" json:"image,omitempty"`
	Streaming    *DomainGraphicsStreaming    `xml:"streaming,omitempty" json:"streaming,omitempty"`
	Clipboard    *DomainGraphicsClipboard    `xml:"clipboard,omitempty" json:"clipboard,omitempty"`
	Mouse        *DomainGraphicsMouse        `xml:"mouse,omitempty" json:"mouse,omitempty"`
	FileTransfer *DomainGraphicsFileTransfer `xml:"filetransfer,omitempty" json:"filetransfer,omitempty"`

	GL *DomainGraphicsGL `xml:"gl,omitempty" json:"gl,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainGraphicsListenType string

const (
	DomainGraphicsListenTypeAddress DomainGraphicsListenType = "address"
	DomainGraphicsListenTypeNetwork DomainGraphicsListenType = "network"
	DomainGraphicsListenTypeSocket  DomainGraphicsListenType = "socket"
	DomainGraphicsListenTypeNone    DomainGraphicsListenType = "none"
)

type DomainGraphicsListen struct {
	Type    DomainGraphicsListenType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Address string                   `xml:"address,attr,omitempty" json:"address,omitempty"`
	Network string                   `xml:"network,attr,omitempty" json:"network,omitempty"`
	Socket  string                   `xml:"socket,attr,omitempty" json:"socket,omitempty"`
}

type DomainGraphicsChannelName string

const (
	DomainGraphicsChannelNameMain      DomainGraphicsChannelName = "main"
	DomainGraphicsChannelNameDisplay   DomainGraphicsChannelName = "display"
	DomainGraphicsChannelNameInputs    DomainGraphicsChannelName = "inputs"
	DomainGraphicsChannelNameCursor    DomainGraphicsChannelName = "cursor"
	DomainGraphicsChannelNamePlayBack  DomainGraphicsChannelName = "playback"
	DomainGraphicsChannelNameRecord    DomainGraphicsChannelName = "record"
	DomainGraphicsChannelNameSmartCard DomainGraphicsChannelName = "smartcard"
	DomainGraphicsChannelNameUsbReDir  DomainGraphicsChannelName = "usbredir"
)

type DomainGraphicsChannelMode string

const (
	DomainGraphicsChannelModeSecure   DomainGraphicsChannelMode = "secure"
	DomainGraphicsChannelModeInsecure DomainGraphicsChannelMode = "insecure"
)

type DomainGraphicsChannel struct {
	Name DomainGraphicsChannelName `xml:"name,attr,omitempty" json:"name,omitempty"`
	Mode DomainGraphicsChannelMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainGraphicsImageCompression string

const (
	DomainGraphicsImageCompressionAutoGLZ DomainGraphicsImageCompression = "auto_glz"
	DomainGraphicsImageCompressionAutoLZ  DomainGraphicsImageCompression = "auto_lz"
	DomainGraphicsImageCompressionQuic    DomainGraphicsImageCompression = "quic"
	DomainGraphicsImageCompressionGLZ     DomainGraphicsImageCompression = "glz"
	DomainGraphicsImageCompressionLZ      DomainGraphicsImageCompression = "lz"
)

type DomainGraphicsImage struct {
	Compression DomainGraphicsImageCompression `xml:"compression,attr,omitempty" json:"compression,omitempty"`
}

type DomainGraphicsStreamingMode string

const (
	DomainGraphicsStreamingModeFilter DomainGraphicsStreamingMode = "filter"
	DomainGraphicsStreamingModeAll    DomainGraphicsStreamingMode = "all"
	DomainGraphicsStreamingModeOff    DomainGraphicsStreamingMode = "off"
)

type DomainGraphicsStreaming struct {
	// Streaming mode is set by the streaming element, settings its mode attribute to one of filter, all or off.
	// Since 0.9.2
	Mode DomainGraphicsStreamingMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainGraphicsClipboard struct {
	CopyPaste ButtonState `xml:"copypaste,attr,omitempty" json:"copypaste,omitempty"`
}

type DomainGraphicsMouseMode string

const (
	DomainGraphicsMouseModeClient DomainGraphicsMouseMode = "client"
	DomainGraphicsMouseModeServer DomainGraphicsMouseMode = "server"
)

type DomainGraphicsMouse struct {
	// Mouse mode is set by the mouse element, setting its mode attribute to one of server or client.
	// If no mode is specified, the qemu default will be used (client mode). Since 0.9.11
	Mode DomainGraphicsMouseMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainGraphicsFileTransfer struct {
	// File transfer functionality (via Spice agent) is set using the filetransfer element. It is enabled
	// by default, and can be disabled by setting the enable property to no. Since 1.2.2
	Enable ButtonState `xml:"enable,attr,omitempty" json:"enable,omitempty"`
}

type DomainGraphicsGL struct {
	Enable     ButtonState `xml:"enable,attr,omitempty" json:"enable,omitempty"`
	RenderNode string      `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty"`
}

type DomainVideo struct {
	// The model element has a mandatory type attribute which takes the value "vga", "cirrus", "vmvga", "xen", "vbox",
	// "qxl" ( since 0.8.6 ), "virtio" ( since 1.3.0 ), "gop" ( since 3.2.0 ), "bochs" ( since 5.6.0 ), "ramfb"
	// ( since 5.9.0 ), or "none" ( since 4.6.0 ), depending on the hypervisor features available.
	//
	// Note: The purpose of the type none is to instruct libvirt not to add a default video device in the guest
	// (see the video element description above), since such behaviour is inconvenient in cases where GPU mediated
	// devices are meant to be the only rendering device within a guest. If this is your use case specify a none type
	// video device in the XML to stop the default behaviour. Refer to Host device assignment to see how to add a
	// mediated device into a guest.
	//
	// You can provide the amount of video memory in kibibytes (blocks of 1024 bytes) using vram. This is supported
	// only for guest type of "vz", "qemu", "vbox", "vmx" and "xen". If no value is provided the default is used.
	// If the size is not a power of two it will be rounded to closest one.
	//
	// The number of screen can be set using heads. This is supported only for guests type of "vz", "kvm",
	// "vbox" and "vmx".
	//
	// For guest type of "kvm" or "qemu" and model type "qxl" there are optional attributes. Attribute ram ( since 1.0.2 )
	// specifies the size of the primary bar, while the attribute vram specifies the secondary bar size. If ram or
	// vram are not supplied a default value is used. The ram should also be rounded to power of two as vram. There
	// is also optional attribute vgamem ( since 1.2.11 ) to set the size of VGA framebuffer for fallback mode of QXL
	// device. Attribute vram64 ( since 1.3.3 ) extends secondary bar and makes it addressable as 64bit memory.
	//
	// Since 5.9.0 , the model element may also have an optional resolution sub-element. The resolution element has
	// attributes x and y to set the minimum resolution for the video device. This sub-element is valid for model types
	// "vga", "qxl", "bochs", "gop", and "virtio".
	Model *DomainVideoModel `xml:"model,omitempty" json:"model,omitempty"`

	Acceleration *DomainVideoAcceleration `xml:"acceleration,omitempty" json:"acceleration,omitempty"`

	Alias   *DomainAlias         `xml:"alias,attr,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
	Driver  *DomainVideoDriver   `xml:"driver,omitempty" json:"driver,omitempty"`
}

type DomainVideoModelType string

const (
	DomainVideoModelTypeVga    DomainVideoModelType = "vga"
	DomainVideoModelTypeCirrus DomainVideoModelType = "cirrus"
	DomainVideoModelTypeVmvga  DomainVideoModelType = "vmvga"
	DomainVideoModelTypeXen    DomainVideoModelType = "xen"
	DomainVideoModelTypeVbox   DomainVideoModelType = "vbox"
	DomainVideoModelTypeQxl    DomainVideoModelType = "qxl"
	DomainVideoModelTypeVirtio DomainVideoModelType = "virtio"
	DomainVideoModelTypeGop    DomainVideoModelType = "gop"
	DomainVideoModelTypeBochs  DomainVideoModelType = "bochs"
	DomainVideoModelTypeRamfb  DomainVideoModelType = "ramfb"
)

type DomainVideoModel struct {
	Type  DomainVideoModelType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Vram  int64                `xml:"vram,attr,omitempty" json:"vram,omitempty"`
	Heads int32                `xml:"heads,attr,omitempty" json:"heads,omitempty"`
}

type DomainVideoAccelerationRenderNode string

const (
	DomainVideoAccelerationRenderNodeVHostUser DomainVideoAccelerationRenderNode = "vhostuser"
)

type DomainVideoAcceleration struct {
	Accel2d    ButtonState                       `xml:"accel2d,attr,omitempty" json:"accel2d,omitempty"`
	Accel3d    ButtonState                       `xml:"accel3d,attr,omitempty" json:"accel3d,omitempty"`
	RenderNode DomainVideoAccelerationRenderNode `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty"`
}

type DomainVideoDriverVgaConf string

const (
	DomainVideoDriverVgaConfIO  DomainVideoDriverVgaConf = "io"
	DomainVideoDriverVgaConfOn  DomainVideoDriverVgaConf = "on"
	DomainVideoDriverVgaConfOff DomainVideoDriverVgaConf = "off"
)

type DomainVideoDriver struct {
	Name    DomainDriverName         `xml:"name,attr,omitempty" json:"name,omitempty"`
	VgaConf DomainVideoDriverVgaConf `xml:"vgaconf,attr,omitempty" json:"vgaconf,omitempty"`
}

type DomainDeviceType string

const (
	DomainDeviceTypePty       DomainDeviceType = "pty"
	DomainDeviceTypeFile      DomainDeviceType = "file"
	DomainDeviceTypeUnix      DomainDeviceType = "unix"
	DomainDeviceTypeStdio     DomainDeviceType = "stdio"
	DomainDeviceTypeVc        DomainDeviceType = "vc"
	DomainDeviceTypeNull      DomainDeviceType = "null"
	DomainDeviceTypeDev       DomainDeviceType = "dev"
	DomainDeviceTypePipe      DomainDeviceType = "pipe"
	DomainDeviceTypeTCP       DomainDeviceType = "tcp"
	DomainDeviceTypeUDP       DomainDeviceType = "udp"
	DomainDeviceTypeSpicePort DomainDeviceType = "spiceport"
	DomainDeviceTypeNmdm      DomainDeviceType = "nmdm"
)

type DomainParallel struct {
	Type     DomainDeviceType      `xml:"type,attr,omitempty" json:"type,omitempty"`
	Source   *DomainDeviceSource   `xml:"source,omitempty" json:"source,omitempty"`
	Protocol *DomainDeviceProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`
	Target   *DomainDeviceTarget   `xml:"target,omitempty" json:"target,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainSerial struct {
	Type     DomainDeviceType      `xml:"type,attr,omitempty" json:"type,omitempty"`
	Source   *DomainDeviceSource   `xml:"source,omitempty" json:"source,omitempty"`
	Protocol *DomainDeviceProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`
	Target   *DomainDeviceTarget   `xml:"target,omitempty" json:"target,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainConsole struct {
	Type     DomainDeviceType      `xml:"type,attr,omitempty" json:"type,omitempty"`
	Source   *DomainDeviceSource   `xml:"source,omitempty" json:"source,omitempty"`
	Protocol *DomainDeviceProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`
	Target   *DomainDeviceTarget   `xml:"target,omitempty" json:"target,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainChannel struct {
	Type     DomainDeviceType      `xml:"type,attr,omitempty" json:"type,omitempty"`
	Source   *DomainDeviceSource   `xml:"source,omitempty" json:"source,omitempty"`
	Protocol *DomainDeviceProtocol `xml:"protocol,omitempty" json:"protocol,omitempty"`
	Target   *DomainDeviceTarget   `xml:"target,omitempty" json:"target,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainDeviceSourceMode string

const (
	DomainDeviceSourceModeBind    DomainDeviceSourceMode = "bind"
	DomainDeviceSourceModeConnect DomainDeviceSourceMode = "connect"
)

type DomainDeviceSource struct {
	Mode     DomainDeviceSourceMode `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Path     string                 `xml:"path,attr,omitempty" json:"path,omitempty"`
	Append   TurnState              `xml:"append,attr,omitempty" json:"append,omitempty"`
	SecLabel *DomainSecLabel        `xml:"seclabel,omitempty" json:"seclabel,omitempty"`

	Host    string      `xml:"host,attr,omitempty" json:"host,omitempty"`
	Service int32       `xml:"service,attr,omitempty" json:"service,omitempty"`
	TLS     ButtonState `xml:"tls,attr,omitempty" json:"tls,omitempty"`

	Channel string `xml:"channel,attr,omitempty" json:"channel,omitempty"`

	// Master device of the pair, that is passed to the hypervisor. Device is specified by a fully qualified path.
	Master string `xml:"master,attr,omitempty" json:"master,omitempty"`
	// Slave device of the pair, that is passed to the clients for connection to the guest console. Device is
	// specified by a fully qualified path.
	Slave string `xml:"slave,attr,omitempty" json:"slave,omitempty"`
}

type DomainSecLabelType string

const (
	DomainSecLabelTypeStatic  DomainSecLabelType = "static"
	DomainSecLabelTypeDynamic DomainSecLabelType = "dynamic"
	DomainSecLabelTypeNone    DomainSecLabelType = "none"
)

type DomainSecLabelModel string

const (
	DomainSecLabelModelDac DomainSecLabelModel = "dac"
)

type DomainDeviceProtocolType string

const (
	DomainDeviceProtocolTypeRaw    DomainDeviceProtocolType = "raw"
	DomainDeviceProtocolTypeTelnet DomainDeviceProtocolType = "telnet"
)

type DomainDeviceProtocol struct {
	Type DomainDeviceProtocolType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainSecLabel struct {
	Type    DomainSecLabelType  `xml:"type,attr,omitempty" json:"type,omitempty"`
	Model   DomainSecLabelModel `xml:"model,attr,omitempty" json:"model,omitempty"`
	Relabel ButtonState         `xml:"relabel,attr,omitempty" json:"relabel,omitempty"`
}

type DomainDeviceTargetType string

const (
	DomainDeviceTargetTypeGuestFwd DomainDeviceTargetType = "guestfwd" // channel device
	DomainDeviceTargetTypeXen      DomainDeviceTargetType = "xen"      // channel device
	DomainDeviceTargetTypeSpiceVmc DomainDeviceTargetType = "spicevmc" // channel device

	DomainDeviceTargetTypeIsaSerial       DomainDeviceTargetType = "isa-serial"       // serial device
	DomainDeviceTargetTypeUsbSerial       DomainDeviceTargetType = "usb-serial"       // serial device
	DomainDeviceTargetTypePciSerial       DomainDeviceTargetType = "pci-serial"       // serial device
	DomainDeviceTargetTypeSpaprVioSerinal DomainDeviceTargetType = "spapr-vio-serial" // serial device
	DomainDeviceTargetTypeSystemSerial    DomainDeviceTargetType = "system-serial"    // serial device
	DomainDeviceTargetTypeSclpSerial      DomainDeviceTargetType = "sclp-serial"      // serial device

	DomainDeviceTargetTypeSerial DomainDeviceTargetType = "serial" // console device
	DomainDeviceTargetTypeVirtio DomainDeviceTargetType = "virtio" // console device | channel device
)

type DomainDeviceTargetState string

const (
	DomainDeviceTargetStateConnect    DomainDeviceTargetState = "connect"
	DomainDeviceTargetStateDisconnect DomainDeviceTargetState = "disconnect"
)

// DomainDeviceTarget the target element can have an optional port attribute, which specifies the port number
// (starting from 0), and an optional type attribute: valid values are, since 1.0.2 , isa-serial (usable with
// x86 guests), usb-serial (usable whenever USB support is available) and pci-serial (usable whenever PCI support
// is available); since 3.10.0 , spapr-vio-serial (usable with ppc64/pseries guests), system-serial (usable with
// aarch64/virt and, since 4.7.0 , riscv/virt guests) and sclp-serial (usable with s390 and s390x guests) are
// available as well.
type DomainDeviceTarget struct {
	Type    DomainDeviceTargetType `xml:"type,attr,omitempty" json:"type,omitempty"`
	Address string                 `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32                  `xml:"port,attr,omitempty" json:"port,omitempty"`

	Name  string                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	State DomainDeviceTargetState `xml:"state,attr,omitempty" json:"state,omitempty"`

	Model *DomainDeviceTargetModel `xml:"model,omitempty" json:"model,omitempty"`
}

type DomainDeviceTargetModelName string

const (
	DomainDeviceTargetModelNameIsaSerial DomainDeviceTargetModelName = "isa-serial"
	DomainDeviceTargetModelNameUsbSerial DomainDeviceTargetModelName = "usb-serial"
	DomainDeviceTargetModelNamePciSerial DomainDeviceTargetModelName = "pci-serial"
	DomainDeviceTargetModelNameSpaprVty  DomainDeviceTargetModelName = "spapr-vty"
)

// DomainDeviceTargetModel since 3.10.0 , the target element can have an optional model subelement; valid values
// for its name attribute are: isa-serial (usable with the isa-serial target type); usb-serial (usable with the
// usb-serial target type); pci-serial (usable with the pci-serial target type); spapr-vty (usable with the
// spapr-vio-serial target type); pl011 and, since 4.7.0 , 16550a (usable with the system-serial target type);
// sclpconsole and sclplmconsole (usable with the sclp-serial target type). Providing a target model is usually
// unnecessary: libvirt will automatically pick one that's suitable for the chosen target type, and overriding
// that value is generally not recommended.
type DomainDeviceTargetModel struct {
	Name DomainDeviceTargetModelName `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainSoundModel string

const (
	DomainSoundModelES1370 DomainSoundModel = "es1370"
	DomainSoundModelSB16   DomainSoundModel = "sb16"
	DomainSoundModelAC97   DomainSoundModel = "ac97"
	DomainSoundModelICH6   DomainSoundModel = "ich6"
	DomainSoundModelUSB    DomainSoundModel = "usb"
)

// DomainSound the sound element has one mandatory attribute, model, which specifies what real sound device is
// emulated. Valid values are specific to the underlying hypervisor, though typical choices are 'es1370', 'sb16',
// 'ac97', 'ich6' and 'usb'. ( 'ac97' only since 0.6.0, 'ich6' only since 0.8.8, 'usb' only since 1.2.7 )
type DomainSound struct {
	Model DomainSoundModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	// Since 0.9.13 , a sound element with ich6 model can have optional sub-elements <codec> to attach various
	// audio codecs to the audio device. If not specified, a default codec will be attached to allow playback
	// and recording.
	//
	// Valid values are:
	//  'duplex'	: advertise a line-in and a line-out
	//  'micro'		: advertise a speaker and a microphone
	//  'output'	: advertise a line-out Since 4.4.0
	Codec *DomainSoundCodec `xml:"codec,omitempty" json:"codec,omitempty"`

	Audio *DomainSoundAudio `xml:"audio,omitempty" json:"audio,omitempty"`
}

type DomainSoundCodecType string

const (
	DomainSoundCodecTypeDuplex DomainSoundCodecType = "duplex"
	DomainSoundCodecTypeMicro  DomainSoundCodecType = "micro"
	DomainSoundCodecTypeOutput DomainSoundCodecType = "output"
)

type DomainSoundCodec struct {
	Type DomainSoundCodecType `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainSoundAudio struct {
	Id int32 `xml:"id,attr,omitempty" json:"id,omitempty"`
}

type DomainAudioType string

const (
	DomainAudioTypeNone       DomainAudioType = "none"
	DomainAudioTypeAlsa       DomainAudioType = "alsa"
	DomainAudioTypeCoreAudio  DomainAudioType = "coreaudio"
	DomainAudioTypeJack       DomainAudioType = "jack"
	DomainAudioTypeOSS        DomainAudioType = "oss"
	DomainAudioTypePulseAudio DomainAudioType = "pluseaudio"
	DomainAudioTypeSDL        DomainAudioType = "sdl"
	DomainAudioTypeSpice      DomainAudioType = "spice"
	DomainAudioTypeFile       DomainAudioType = "file"
)

type DomainAudioDriver string

const (
	DomainAudioDriverESD        DomainAudioDriver = "esd"
	DomainAudioDriverAlsa       DomainAudioDriver = "alsa"
	DomainAudioDriverArts       DomainAudioDriver = "arts"
	DomainAudioDriverPulseaudio DomainAudioDriver = "pulseaudio"
)

type DomainAudio struct {
	Id   int32           `xml:"id,attr,omitempty" json:"id,omitempty"`
	Type DomainAudioType `xml:"type,attr,omitempty" json:"type,omitempty"`

	// OSS audio backend
	// Attempt to use mmap for data transfer
	TryMMap ButtonState `xml:"tryMMmap,attr,omitempty" json:"tryMMap,omitempty"`
	// Enforce exclusive access to the host device
	Exclusive ButtonState `xml:"exclusive,attr,omitempty" json:"exclusive,omitempty"`
	// Set the timing policy of the device, values between -1 and 10. Smaller numbers result in
	// lower latency but higher CPU usage. A negative value requests use of fragment mode.
	DspPolicy int32 `xml:"dspPolicy,attr,omitempty" json:"dspPolicy,omitempty"`

	// PulseAudio audio backend
	ServerName string `xml:"serverName,attr,omitempty" json:"serverName,omitempty"`

	// SDL audio backend
	// SDL audio driver. The name attribute specifies SDL driver name, one of 'esd', 'alsa', 'arts', 'pulseaudio'.
	Driver DomainAudioDriver `xml:"driver,attr,omitempty" json:"driver,omitempty"`

	// The 'file' audio backend is an output only driver which records audio to a file. The file format is
	// implementation defined, and defaults to 'WAV' with QEMU.
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`

	Input  *DomainAudioIO `xml:"input,omitempty" json:"input,omitempty"`
	Output *DomainAudioIO `xml:"output,omitempty" json:"output,omitempty"`
}

type DomainAudioIO struct {
	// Control whether the host mixing engine is used to convert between different audio formats and
	// sampling rates. When the mixing engine is disabled it is possible to make use of improved audio
	// formats such as 5.1/7.1. If not specified, a hypervisor default applies.
	MixingEngine ButtonState `xml:"mixingEngine,attr,omitempty" json:"mixingEngine,omitempty"`
	// Control whether the mixing engine can dynamically choose settings to minimize format conversion.
	// This is only valid when the mixing engine is explicitly enabled.
	FixedSettings ButtonState `xml:"fixedSettings,attr,omitempty" json:"fixedSettings,omitempty"`
	// The number of voices voices to use, usually defaults to 1
	Voices int32 `xml:"voices,attr,omitempty" json:"voices,omitempty"`
	// The length of the audio buffer in microseconds. Default is backend specific.
	BufferLength int32 `xml:"bufferLength,attr,omitempty" json:"bufferLength,omitempty"`

	// Path to the host device node to connect the backend to. A hypervisor specific default applies if not specified.
	Dev string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	// The number of buffers. It is recommended to set the bufferLength attribute at the same time.
	BufferCount int32 `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty"`

	TryPoll ButtonState `xml:"tryPoll,attr,omitempty" json:"tryPoll,omitempty"`

	// Select the Jack server instance to connect to.
	ServerName string `xml:"serverName,attr,omitempty" json:"serverName,omitempty"`
	// The client name to identify as. The server may modify this to ensure uniqueness unless exactName is enabled
	ClientName string `xml:"clientName,attr,omitempty" json:"clientName,omitempty"`
	// A regular expression of Jack client port names to monitor and connect to.
	ConnectPorts string `xml:"connectPorts,attr,omitempty" json:"connectPorts,omitempty"`
	// Use the exact clientName requested
	ExactName ButtonState `xml:"exactName,attr,omitempty" json:"exactName,omitempty"`

	StreamName string `xml:"streamName,attr,omitempty" json:"streamName,omitempty"`
	Latency    int32  `xml:"latency,attr,omitempty" json:"latency,omitempty"`
}

type DomainAudioIOFormat string

const (
	DomainAudioIOFormatS8  DomainAudioIOFormat = "s8"
	DomainAudioIOFormatU8  DomainAudioIOFormat = "u8"
	DomainAudioIOFormatS16 DomainAudioIOFormat = "s16"
	DomainAudioIOFormatU16 DomainAudioIOFormat = "u16"
	DomainAudioIOFormatS32 DomainAudioIOFormat = "s32"
	DomainAudioIOFormatU32 DomainAudioIOFormat = "u32"
	DomainAudioIOFormatF32 DomainAudioIOFormat = "f32"
)

type DomainAudioIOSetting struct {
	// The frequency in HZ, usually defaulting to 44100
	Frequency int64 `xml:"frequency,attr,omitempty" json:"frequency,omitempty"`
	// The number of channels, usually defaulting to 2. The permitted max number of channels is hypervisor specific.
	Channels int32 `xml:"channels,attr,omitempty" json:"channels,omitempty"`
	// The audio format, one of s8, u8, s16, u16, s32, u32, f32. The default is hypervisor specific.
	Format DomainAudioIOFormat `xml:"format,attr,omitempty" json:"format,omitempty"`
}

type DomainWatchDogModel string

const (
	DomainWatchDogModelI6300ESB DomainWatchDogModel = "i6300esb"
	DomainWatchDogModelIB700    DomainWatchDogModel = "ib700"
	DomainWatchDogModelDiag288  DomainWatchDogModel = "diag288"
)

type DomainWatchDogAction string

const (
	DomainWatchDogActionReset     DomainWatchDogAction = "reset"
	DomainWatchDogActionShutdown  DomainWatchDogAction = "shutdown"
	DomainWatchDogActionPowerOff  DomainWatchDogAction = "poweroff"
	DomainWatchDogActionPause     DomainWatchDogAction = "pause"
	DomainWatchDogActionNone      DomainWatchDogAction = "none"
	DomainWatchDogActionDump      DomainWatchDogAction = "dump"
	DomainWatchDogActionInjectNmi DomainWatchDogAction = "injectNmi"
)

// DomainWatchDog a virtual hardware watchdog device can be added to the guest via the watchdog element.
// Since 0.7.3, QEMU and KVM only
// The watchdog device requires an additional driver and management daemon in the guest. Just enabling the
// watchdog in the libvirt configuration does not do anything useful on its own.
type DomainWatchDog struct {
	// The required model attribute specifies what real watchdog device is emulated. Valid values are
	// specific to the underlying hypervisor.
	// QEMU and KVM support:
	//  'i6300esb'	: the recommended device, emulating a PCI Intel 6300ESB
	//  'ib700'		: emulating an ISA iBase IB700
	//  'diag288'	: emulating an S390 DIAG288 device Since 1.2.17
	Model DomainWatchDogModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	// The optional action attribute describes what action to take when the watchdog expires. Valid values
	// are specific to the underlying hypervisor.
	// QEMU and KVM support:
	//  'reset'		: default, forcefully reset the guest
	//  'shutdown'	: gracefully shutdown the guest (not recommended)
	//  'poweroff'  : forcefully power off the guest
	//  'pause'		: pause the guest
	//  'none'		: do nothing
	//  'dump'		: automatically dump the guest Since 0.8.7
	// 	'inject-nmi': inject a non-maskable interrupt into the guest Since 1.2.17
	//
	// Note 1: the 'shutdown' action requires that the guest is responsive to ACPI signals. In the sort of situations
	//         where the watchdog has expired, guests are usually unable to respond to ACPI signals. Therefore using
	//         'shutdown' is not recommended.
	// Note 2: the directory to save dump files can be configured by auto_dump_path in file /etc/libvirt/qemu.conf.
	Action DomainWatchDogAction `xml:"action,attr,omitempty" json:"action,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainMemBalloonModel string

const (
	DomainMemBalloonModelVirtio                DomainMemBalloonModel = "virtio"
	DomainMemBalloonModelVirtioTransitional    DomainMemBalloonModel = "virtio-transitional"
	DomainMemBalloonModelVirtioNonTransitional DomainMemBalloonModel = "virtio-non-transitional"
	DomainMemBalloonModelXen                   DomainMemBalloonModel = "xen"
)

// DomainMemBalloon a virtual memory balloon device is added to all Xen and KVM/QEMU guests. It will be seen
// as memballoon element. It will be automatically added when appropriate, so there is no need to explicitly
// add this element in the guest XML unless a specific PCI slot needs to be assigned. Since 0.8.3, Xen, QEMU
// and KVM only Additionally, since 0.8.4 , if the memballoon device needs to be explicitly disabled,
// model='none' may be used.
type DomainMemBalloon struct {
	Model DomainMemBalloonModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	// The optional autodeflate attribute allows to enable/disable (values "on"/"off", respectively) the
	// ability of the QEMU virtio memory balloon to release some memory at the last moment before a guest's
	// process get killed by Out of Memory killer. Since 1.3.1, QEMU and KVM only
	AutoDeflate TurnState `xml:"autodeflate,attr,omitempty" json:"autodeflate,omitempty"`

	// The optional freePageReporting attribute allows to enable/disable ("on"/"off", respectively) the ability
	// of the QEMU virtio memory balloon to return unused pages back to the hypervisor to be used by other guests
	// or processes. Please note that despite its name it has no effect on free memory as reported by
	// virDomainMemoryStats() and/or virsh dommemstat. Since 6.9.0, QEMU and KVM only
	FreePageReporting TurnState `xml:"freePageReporting,attr,omitempty" json:"freePageReporting,omitempty"`

	// The optional period allows the QEMU virtio memory balloon driver to provide statistics through the
	// virsh dommemstat [domain] command. By default, collection is not enabled. In order to enable, use the
	// virsh dommemstat [domain] --period [number] command or virsh edit command to add the option to the XML
	// definition. The virsh dommemstat will accept the options --live, --current, or --config. If an option
	// is not provided, the change for a running domain will only be made to the active guest. If the QEMU
	// driver is not at the right revision, the attempt to set the period will fail. Large values
	// (e.g. many years) might be ignored. Since 1.1.1, requires QEMU 1.5
	Period int32 `xml:"period,attr,omitempty" json:"period,omitempty"`

	Stat DomainMemBalloonStats `xml:"stat,omitempty" json:"stat,omitempty"`

	// For model virtio memballoon, Virtio-specific options can also be set. ( Since 3.5.0 )
	Driver *DomainDeviceDriver `xml:"driver,omitempty" json:"driver,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainMemBalloonStats struct {
	Period int32 `xml:"period,attr,omitempty" json:"period,omitempty"`
}

type DomainDeviceDriver struct {
	IOMMU ButtonState `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	Ats   ButtonState `xml:"ats,attr,omitempty" json:"ats,omitempty"`
}

// DomainRng the virtual random number generator device allows the host to pass through entropy to guest
// operating systems. Since 1.0.3
type DomainRng struct {
	// The required model attribute specifies what type of RNG device is provided.
	Model DomainVirtioModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	// The optional rate element allows limiting the rate at which entropy can be consumed from the source.
	// The mandatory attribute bytes specifies how many bytes are permitted to be consumed per period. An
	// optional period attribute specifies the duration of a period in milliseconds; if omitted, the period
	// is taken as 1000 milliseconds (1 second). Since 1.0.4
	Rate *DomainRngRate `xml:"rate,omitempty" json:"rate,omitempty"`

	// The backend element specifies the source of entropy to be used for the domain. The source model is
	// configured using the model attribute. Supported source models are:
	//  random	: This backend type expects a non-blocking character device as input. The file name is specified
	//            as contents of the backend element. Since 1.3.4 any path is accepted. Before that /dev/random
	//            and /dev/hwrng were the only accepted paths. When no file name is specified, the hypervisor
	//            default is used. For QEMU, the default is /dev/random. However, the recommended source of
	//            entropy is /dev/urandom (as it doesn't have the limitations of /dev/random).
	//
	//  egd		: This backend connects to a source using the EGD protocol. The source is specified as a character
	//            device. Refer to character device host interface for more information.
	//
	//  builtin	: This backend uses qemu builtin random generator, which uses getrandom() syscall as the source of
	//            entropy. ( Since 6.1.0 and QEMU 4.2 )
	Backend *DomainRngBackend `xml:"backend,omitempty" json:"backend,omitempty"`

	Driver *DomainDeviceDriver `xml:"driver,omitempty" json:"driver,omitempty"`

	ACPI    *DomainDeviceACPI    `xml:"acpi,omitempty" json:"acpi,omitempty"`
	Alias   *DomainAlias         `xml:"alias,omitempty" json:"alias,omitempty"`
	Address *DomainDeviceAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type DomainRngRate struct {
	Period int32 `xml:"period,attr,omitempty" json:"period,omitempty"`
	Bytes  int64 `xml:"bytes,attr,omitempty" json:"bytes,omitempty"`
}

type DomainRngBackendModel string

const (
	DomainRngBackendModelRandom  DomainRngBackendModel = "random"
	DomainRngBackendModelEGD     DomainRngBackendModel = "egd"
	DomainRngBackendModelBuiltin DomainRngBackendModel = "builtin"
)

type DomainRngBackend struct {
	Model DomainRngBackendModel `xml:"model,attr,omitempty" json:"model,omitempty"`

	Value string `xml:",chardata" json:"value"`
}
