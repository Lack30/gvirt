package spec

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type DomainControllerPCIHole64 struct {
	Size int64  `xml:",chardata" json:"size"`
	Unit string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainControllerPCIModel struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainControllerPCITarget struct {
	ChassisNr *int32 `json:"chassisNr,omitempty"`
	Chassis   *int32 `json:"chassis,omitempty"`
	Port      *int32 `json:"port,omitempty"`
	BusNr     *int32 `json:"busNr,omitempty"`
	Index     *int32 `json:"index,omitempty"`
	NUMANode  *int32 `json:"numaNode,omitempty"`
	Hotplug   string `json:"hotplug,omitempty"`
}

type DomainControllerPCI struct {
	Model  *DomainControllerPCIModel  `xml:"model" json:"model,omitempty"`
	Target *DomainControllerPCITarget `xml:"target" json:"target,omitempty"`
	Hole64 *DomainControllerPCIHole64 `xml:"pcihole64" json:"pcihole64,omitempty"`
}

type DomainControllerUSBMaster struct {
	StartPort int32 `xml:"startport,attr" json:"startPort"`
}

type DomainControllerUSB struct {
	Port   *int32                     `xml:"ports,attr" json:"port,omitempty"`
	Master *DomainControllerUSBMaster `xml:"master,omitempty" json:"master,omitempty"`
}

type DomainControllerVirtIOSerial struct {
	Ports   *int32 `xml:"ports,attr" json:"ports,omitempty"`
	Vectors *int32 `xml:"vectors,attr" json:"vectors,omitempty"`
}

type DomainControllerXenBus struct {
	MaxGrantFrames   int32 `xml:"maxGrantFrames,attr,omitempty" json:"maxGrantFrames,omitempty"`
	MaxEventChannels int32 `xml:"maxEventChannels,attr,omitempty" json:"maxEventChannels,omitempty"`
}

type DomainControllerDriver struct {
	Queues     *int32 `xml:"queues,attr" json:"queues,omitempty"`
	CmdPerLUN  *int32 `xml:"cmd_per_lun,attr" json:"cmdPerLun,omitempty"`
	MaxSectors *int32 `xml:"max_sectors,attr" json:"maxSectors,omitempty"`
	IOEventFD  string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`
	IOThread   int32  `xml:"iothread,attr,omitempty" json:"iothread,omitempty"`
	IOMMU      string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS        string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed     string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainController struct {
	XMLName      xml.Name                      `xml:"controller" json:"-"`
	Type         string                        `xml:"type,attr" json:"type,omitempty"`
	Index        *int32                        `xml:"index,attr" json:"index,omitempty"`
	Model        string                        `xml:"model,attr,omitempty" json:"model,omitempty"`
	Driver       *DomainControllerDriver       `xml:"driver" json:"driver,omitempty"`
	PCI          *DomainControllerPCI          `xml:"-" json:"pci,omitempty"`
	USB          *DomainControllerUSB          `xml:"-" json:"usb,omitempty"`
	VirtIOSerial *DomainControllerVirtIOSerial `xml:"-" json:"virtioSerial,omitempty"`
	XenBus       *DomainControllerXenBus       `xml:"-" json:"xenBus,omitempty"`
	ACPI         *DomainDeviceACPI             `xml:"acpi" json:"acpi,omitempty"`
	Alias        *DomainAlias                  `xml:"alias" json:"alias,omitempty"`
	Address      *DomainAddress                `xml:"address" json:"address,omitempty"`
}

type DomainDiskSecret struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
}

type DomainDiskAuth struct {
	Username string            `xml:"username,attr,omitempty" json:"username,omitempty"`
	Secret   *DomainDiskSecret `xml:"secret" json:"secret,omitempty"`
}

type DomainDiskSourceHost struct {
	Transport string `xml:"transport,attr,omitempty" json:"transport,omitempty"`
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Port      string `xml:"port,attr,omitempty" json:"port,omitempty"`
	Socket    string `xml:"socket,attr,omitempty" json:"socket,omitempty"`
}

type DomainDiskSourceSSL struct {
	Verify string `xml:"verify,attr" json:"verify"`
}

type DomainDiskCookie struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:",chardata" json:"value"`
}

type DomainDiskCookies struct {
	Cookies []DomainDiskCookie `xml:"cookie" json:"cookies"`
}

type DomainDiskSourceReadahead struct {
	Size string `xml:"size,attr" json:"size"`
}

type DomainDiskSourceTimeout struct {
	Seconds string `xml:"seconds,attr" json:"seconds"`
}

type DomainDiskReservationsSource DomainChardevSource

type DomainDiskReservations struct {
	Enabled string                        `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`
	Managed string                        `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	Source  *DomainDiskReservationsSource `xml:"source" json:"source"`
}

type DomainDiskSource struct {
	File          *DomainDiskSourceFile      `xml:"-" json:"file,omitempty"`
	Block         *DomainDiskSourceBlock     `xml:"-" json:"block,omitempty"`
	Dir           *DomainDiskSourceDir       `xml:"-" json:"dir,omitempty"`
	Network       *DomainDiskSourceNetwork   `xml:"-" json:"network,omitempty"`
	Volume        *DomainDiskSourceVolume    `xml:"-" json:"volume,omitempty"`
	NVME          *DomainDiskSourceNVME      `xml:"-" json:"nvme,omitempty"`
	VHostUser     *DomainDiskSourceVHostUser `xml:"-" json:"vHostUser,omitempty"`
	StartupPolicy string                     `xml:"startupPolicy,attr,omitempty" json:"startupPolicy,omitempty"`
	Index         int32                      `xml:"index,attr,omitempty" json:"index,omitempty"`
	Encryption    *DomainDiskEncryption      `xml:"encryption" json:"encryption,omitempty"`
	Reservations  *DomainDiskReservations    `xml:"reservations" json:"reservations,omitempty"`
	Slices        *DomainDiskSlices          `xml:"slices" json:"slices,omitempty"`
	SSL           *DomainDiskSourceSSL       `xml:"ssl" json:"ssl,omitempty"`
	Cookies       *DomainDiskCookies         `xml:"cookies" json:"cookies,omitempty"`
	Readahead     *DomainDiskSourceReadahead `xml:"readahead" json:"readahead,omitempty"`
	Timeout       *DomainDiskSourceTimeout   `xml:"timeout" json:"timeout,omitempty"`
}

type DomainDiskSlices struct {
	Slices []DomainDiskSlice `xml:"slice" json:"slices"`
}

type DomainDiskSlice struct {
	Type   string `xml:"type,attr" json:"type"`
	Offset int32  `xml:"offset,attr" json:"offset"`
	Size   int32  `xml:"size,attr" json:"size"`
}

type DomainDiskSourceFile struct {
	File     string                 `xml:"file,attr,omitempty" json:"file,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel"`
}

type DomainDiskSourceNVME struct {
	PCI *DomainDiskSourceNVMEPCI `json:"pci,omitempty"`
}

type DomainDiskSourceNVMEPCI struct {
	Managed   string            `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	Namespace int64             `xml:"namespace,attr,omitempty" json:"namespace,omitempty"`
	Address   *DomainAddressPCI `xml:"address" json:"address,omitempty"`
}

type DomainDiskSourceBlock struct {
	Dev      string                 `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel"`
}

type DomainDiskSourceDir struct {
	Dir string `xml:"dir,attr,omitempty" json:"dir,omitempty"`
}

type DomainDiskSourceNetwork struct {
	Protocol  string                            `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	Name      string                            `xml:"name,attr,omitempty" json:"name,omitempty"`
	Query     string                            `xml:"query,attr,omitempty" json:"query,omitempty"`
	TLS       string                            `xml:"tls,attr,omitempty" json:"tls,omitempty"`
	Hosts     []DomainDiskSourceHost            `xml:"host" json:"hosts"`
	Identity  *DomainDiskSourceNetworkIdentity  `xml:"identity" json:"identity,omitempty"`
	Initiator *DomainDiskSourceNetworkInitiator `xml:"initiator" json:"initiator,omitempty"`
	Snapshot  *DomainDiskSourceNetworkSnapshot  `xml:"snapshot" json:"snapshot,omitempty"`
	Config    *DomainDiskSourceNetworkConfig    `xml:"config" json:"config,omitempty"`
	Auth      *DomainDiskAuth                   `xml:"auth" json:"auth,omitempty"`
}

type DomainDiskSourceNetworkIdentity struct {
	User  string `xml:"user,attr" json:"user"`
	Group string `xml:"group,attr" json:"group"`
}

type DomainDiskSourceNetworkInitiator struct {
	IQN *DomainDiskSourceNetworkIQN `xml:"iqn" json:"iqn,omitempty"`
}

type DomainDiskSourceNetworkIQN struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainDiskSourceNetworkSnapshot struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainDiskSourceNetworkConfig struct {
	File string `xml:"file,attr" json:"file"`
}

type DomainDiskSourceVolume struct {
	Pool     string                 `xml:"pool,attr,omitempty" json:"pool,omitempty"`
	Volume   string                 `xml:"volume,attr,omitempty" json:"volume,omitempty"`
	Mode     string                 `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"sec_label"`
}

type DomainDiskSourceVHostUser DomainChardevSource

type DomainDiskMetadataCache struct {
	MaxSize *DomainDiskMetadataCacheSize `xml:"max_size" json:"maxSize"`
}

type DomainDiskMetadataCacheSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value int    `xml:",cdata" json:"value"`
}

type DomainDiskDriver struct {
	Name          string                   `xml:"name,attr,omitempty" json:"name,omitempty"`
	Type          string                   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Cache         string                   `xml:"cache,attr,omitempty" json:"cache,omitempty"`
	ErrorPolicy   string                   `xml:"error_policy,attr,omitempty" json:"errorPolicy,omitempty"`
	RErrorPolicy  string                   `xml:"rerror_policy,attr,omitempty" json:"rerrorPolicy,omitempty"`
	IO            string                   `xml:"io,attr,omitempty" json:"io,omitempty"`
	IOEventFD     string                   `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`
	EventIDX      string                   `xml:"event_idx,attr,omitempty" json:"eventIdx,omitempty"`
	CopyOnRead    string                   `xml:"copy_on_read,attr,omitempty" json:"copyOnRead,omitempty"`
	Discard       string                   `xml:"discard,attr,omitempty" json:"discard,omitempty"`
	IOThread      *int32                   `xml:"iothread,attr" json:"iothread,omitempty"`
	DetectZeros   string                   `xml:"detect_zeroes,attr,omitempty" json:"detectZeros,omitempty"`
	Queues        *int32                   `xml:"queues,attr" json:"queues,omitempty"`
	IOMMU         string                   `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS           string                   `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed        string                   `xml:"packed,attr,omitempty" json:"packed,omitempty"`
	MetadataCache *DomainDiskMetadataCache `xml:"metadata_cache" json:"metadata_cache,omitempty"`
}

type DomainDiskTarget struct {
	Dev          string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Bus          string `xml:"bus,attr,omitempty" json:"bus,omitempty"`
	Tray         string `xml:"tray,attr,omitempty" json:"tray,omitempty"`
	Removable    string `xml:"removable,attr,omitempty" json:"removable,omitempty"`
	RotationRate int32  `xml:"rotation_rate,attr,omitempty" json:"rotationRate,omitempty"`
}

type DomainDiskEncryption struct {
	Format string            `xml:"format,attr,omitempty" json:"format,omitempty"`
	Secret *DomainDiskSecret `xml:"secret" json:"secret,omitempty"`
}

type DomainDiskReadOnly struct {
}

type DomainDiskShareable struct {
}

type DomainDiskTransient struct {
}

type DomainDiskIOTune struct {
	TotalBytesSec          int64  `xml:"total_bytes_sec,omitempty" json:"totalBytesSec,omitempty"`
	ReadBytesSec           int64  `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty"`
	WriteBytesSec          int64  `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty"`
	TotalIopsSec           int64  `xml:"total_iops_sec,omitempty" json:"totalIopsSec,omitempty"`
	ReadIopsSec            int64  `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty"`
	WriteIopsSec           int64  `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty"`
	TotalBytesSecMax       int64  `xml:"total_bytes_sec_max,omitempty" json:"totalBytesSecMax,omitempty"`
	ReadBytesSecMax        int64  `xml:"read_bytes_sec_max,omitempty" json:"readBytesSecMax,omitempty"`
	WriteBytesSecMax       int64  `xml:"write_bytes_sec_max,omitempty" json:"writeBytesSecMax,omitempty"`
	TotalIopsSecMax        int64  `xml:"total_iops_sec_max,omitempty" json:"totalIopsSecMax,omitempty"`
	ReadIopsSecMax         int64  `xml:"read_iops_sec_max,omitempty" json:"readIopsSecMax,omitempty"`
	WriteIopsSecMax        int64  `xml:"write_iops_sec_max,omitempty" json:"writeIopsSecMax,omitempty"`
	TotalBytesSecMaxLength int64  `xml:"total_bytes_sec_max_length,omitempty" json:"totalBytesSecMaxLength,omitempty"`
	ReadBytesSecMaxLength  int64  `xml:"read_bytes_sec_max_length,omitempty" json:"readBytesSecMaxLength,omitempty"`
	WriteBytesSecMaxLength int64  `xml:"write_bytes_sec_max_length,omitempty" json:"writeBytesSecMaxLength,omitempty"`
	TotalIopsSecMaxLength  int64  `xml:"total_iops_sec_max_length,omitempty" json:"totalIopsSecMaxLength,omitempty"`
	ReadIopsSecMaxLength   int64  `xml:"read_iops_sec_max_length,omitempty" json:"readIopsSecMaxLength,omitempty"`
	WriteIopsSecMaxLength  int64  `xml:"write_iops_sec_max_length,omitempty" json:"writeIopsSecMaxLength,omitempty"`
	SizeIopsSec            int64  `xml:"size_iops_sec,omitempty" json:"sizeIopsSec,omitempty"`
	GroupName              string `xml:"group_name,omitempty" json:"groupName,omitempty"`
}

type DomainDiskGeometry struct {
	Cylinders int32  `xml:"cyls,attr" json:"cyli"`
	Headers   int32  `xml:"heads,attr" json:"heads"`
	Sectors   int32  `xml:"secs,attr" json:"secs"`
	Trans     string `xml:"trans,attr,omitempty" json:"trans,omitempty"`
}

type DomainDiskBlockIO struct {
	LogicalBlockSize  int32 `xml:"logical_block_size,attr,omitempty" json:"logicalBlockSize,omitempty"`
	PhysicalBlockSize int32 `xml:"physical_block_size,attr,omitempty" json:"physicalBlockSize,omitempty"`
}

type DomainDiskFormat struct {
	Type          string                   `xml:"type,attr" json:"type"`
	MetadataCache *DomainDiskMetadataCache `xml:"metadata_cache" json:"metadataCache,omitempty"`
}

type DomainDiskBackingStore struct {
	Index        int32                   `xml:"index,attr,omitempty" json:"index,omitempty"`
	Format       *DomainDiskFormat       `xml:"format" json:"format,omitempty"`
	Source       *DomainDiskSource       `xml:"source" json:"source,omitempty"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty"`
}

type DomainDiskMirror struct {
	Job          string                  `xml:"job,attr,omitempty" json:"job,omitempty"`
	Ready        string                  `xml:"ready,attr,omitempty" json:"ready,omitempty"`
	Format       *DomainDiskFormat       `xml:"format" json:"format,omitempty"`
	Source       *DomainDiskSource       `xml:"source" json:"source,omitempty"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty"`
}

type DomainBackendDomain struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainDisk struct {
	XMLName       xml.Name                `xml:"disk" json:"-"`
	Device        string                  `xml:"device,attr,omitempty" json:"device,omitempty"`
	RawIO         string                  `xml:"rawio,attr,omitempty" json:"rawio,omitempty"`
	SGIO          string                  `xml:"sgio,attr,omitempty" json:"sgio,omitempty"`
	Snapshot      string                  `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty"`
	Model         string                  `xml:"model,attr,omitempty" json:"model,omitempty"`
	Driver        *DomainDiskDriver       `xml:"driver" json:"driver,omitempty"`
	Auth          *DomainDiskAuth         `xml:"auth" json:"auth,omitempty"`
	Source        *DomainDiskSource       `xml:"source" json:"source,omitempty"`
	BackingStore  *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty"`
	BackendDomain *DomainBackendDomain    `xml:"backenddomain" json:"backendDomain,omitempty"`
	Geometry      *DomainDiskGeometry     `xml:"geometry" json:"geometry,omitempty"`
	BlockIO       *DomainDiskBlockIO      `xml:"blockio" json:"blockio,omitempty"`
	Mirror        *DomainDiskMirror       `xml:"mirror" json:"mirror,omitempty"`
	Target        *DomainDiskTarget       `xml:"target" json:"target,omitempty"`
	IOTune        *DomainDiskIOTune       `xml:"iotune" json:"iotune,omitempty"`
	ReadOnly      *DomainDiskReadOnly     `xml:"readonly" json:"readonly,omitempty"`
	Shareable     *DomainDiskShareable    `xml:"shareable" json:"shareable,omitempty"`
	Transient     *DomainDiskTransient    `xml:"transient" json:"transient,omitempty"`
	Serial        string                  `xml:"serial,omitempty" json:"serial,omitempty"`
	WWN           string                  `xml:"wwn,omitempty" json:"wwn,omitempty"`
	Vendor        string                  `xml:"vendor,omitempty" json:"vendor,omitempty"`
	Product       string                  `xml:"product,omitempty" json:"product,omitempty"`
	Encryption    *DomainDiskEncryption   `xml:"encryption" json:"encryption,omitempty"`
	Boot          *DomainDeviceBoot       `xml:"boot" json:"boot,omitempty"`
	ACPI          *DomainDeviceACPI       `xml:"acpi" json:"acpi,omitempty"`
	Alias         *DomainAlias            `xml:"alias" json:"alias,omitempty"`
	Address       *DomainAddress          `xml:"address" json:"address,omitempty"`
}

type DomainFilesystemDriver struct {
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Format   string `xml:"format,attr,omitempty" json:"format,omitempty"`
	Name     string `xml:"name,attr,omitempty" json:"name,omitempty"`
	WRPolicy string `xml:"wrpolicy,attr,omitempty" json:"wrpolicy,omitempty"`
	IOMMU    string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS      string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed   string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
	Queue    int32  `xml:"queue,attr,omitempty" json:"queue,omitempty"`
}

type DomainFilesystemSource struct {
	Mount    *DomainFilesystemSourceMount    `xml:"-" json:"mount,omitempty"`
	Block    *DomainFilesystemSourceBlock    `xml:"-" json:"block,omitempty"`
	File     *DomainFilesystemSourceFile     `xml:"-" json:"file,omitempty"`
	Template *DomainFilesystemSourceTemplate `xml:"-" json:"template,omitempty"`
	RAM      *DomainFilesystemSourceRAM      `xml:"-" json:"ram,omitempty"`
	Bind     *DomainFilesystemSourceBind     `xml:"-" json:"bind,omitempty"`
	Volume   *DomainFilesystemSourceVolume   `xml:"-" json:"volume,omitempty"`
}

type DomainFilesystemSourceMount struct {
	Dir    string `xml:"dir,attr,omitempty" json:"dir,omitempty"`
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty"`
}

type DomainFilesystemSourceBlock struct {
	Dev string `xml:"dev,attr" json:"dev"`
}

type DomainFilesystemSourceFile struct {
	File string `xml:"file,attr" json:"file"`
}

type DomainFilesystemSourceTemplate struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainFilesystemSourceRAM struct {
	Usage int32  `xml:"usage,attr" json:"usage"`
	Units string `xml:"units,attr,omitempty" json:"units,omitempty"`
}

type DomainFilesystemSourceBind struct {
	Dir string `xml:"dir,attr" json:"dir"`
}

type DomainFilesystemSourceVolume struct {
	Pool   string `xml:"pool,attr" json:"pool"`
	Volume string `xml:"volume,attr" json:"volume"`
}

type DomainFilesystemTarget struct {
	Dir string `xml:"dir,attr" json:"dir"`
}

type DomainFilesystemReadOnly struct {
}

type DomainFilesystemSpaceHardLimit struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainFilesystemSpaceSoftLimit struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainFilesystemBinaryCache struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainFilesystemBinarySandbox struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainFilesystemBinaryLock struct {
	POSIX string `xml:"posix,attr,omitempty" json:"posix,omitempty"`
	Flock string `xml:"flock,attr,omitempty" json:"flock,omitempty"`
}

type DomainFilesystemBinary struct {
	Path    string                         `xml:"path,attr,omitempty" json:"path,omitempty"`
	XAttr   string                         `xml:"xattr,attr,omitempty" json:"xattr,omitempty"`
	Cache   *DomainFilesystemBinaryCache   `xml:"cache" json:"cache,omitempty"`
	Sandbox *DomainFilesystemBinarySandbox `xml:"sandbox" json:"sandbox,omitempty"`
	Lock    *DomainFilesystemBinaryLock    `xml:"lock" json:"lock,omitempty"`
}

type DomainFilesystem struct {
	XMLName        xml.Name                        `xml:"filesystem" json:"-"`
	AccessMode     string                          `xml:"accessmode,attr,omitempty" json:"accessMode,omitempty"`
	Model          string                          `xml:"model,attr,omitempty" json:"model,omitempty"`
	MultiDevs      string                          `xml:"multidevs,attr,omitempty" json:"multidevs,omitempty"`
	FMode          string                          `xml:"fmode,attr,omitempty" json:"fmode,omitempty"`
	DMode          string                          `xml:"dmode,attr,omitempty" json:"dmode,omitempty"`
	Driver         *DomainFilesystemDriver         `xml:"driver" json:"driver,omitempty"`
	Binary         *DomainFilesystemBinary         `xml:"binary" json:"binary,omitempty"`
	Source         *DomainFilesystemSource         `xml:"source" json:"source,omitempty"`
	Target         *DomainFilesystemTarget         `xml:"target" json:"target,omitempty"`
	ReadOnly       *DomainFilesystemReadOnly       `xml:"readonly" json:"readonly,omitempty"`
	SpaceHardLimit *DomainFilesystemSpaceHardLimit `xml:"space_hard_limit" json:"spaceHardLimit,omitempty"`
	SpaceSoftLimit *DomainFilesystemSpaceSoftLimit `xml:"space_soft_limit" json:"spaceSoftLimit,omitempty"`
	Boot           *DomainDeviceBoot               `xml:"boot" json:"boot,omitempty"`
	ACPI           *DomainDeviceACPI               `xml:"acpi" json:"acpi,omitempty"`
	Alias          *DomainAlias                    `xml:"alias" json:"alias,omitempty"`
	Address        *DomainAddress                  `xml:"address" json:"address,omitempty"`
}

type DomainInterfaceMAC struct {
	Address string `xml:"address,attr" json:"address"`
	Type    string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Check   string `xml:"check,attr,omitempty" json:"check,omitempty"`
}

type DomainInterfaceModel struct {
	Type string `xml:"type,attr" json:"type"`
}

type DomainInterfaceSource struct {
	User      *DomainInterfaceSourceUser     `xml:"-" json:"user,omitempty"`
	Ethernet  *DomainInterfaceSourceEthernet `xml:"-" json:"ethernet,omitempty"`
	VHostUser *DomainChardevSource           `xml:"-" json:"vHostUser,omitempty"`
	Server    *DomainInterfaceSourceServer   `xml:"-" json:"server,omitempty"`
	Client    *DomainInterfaceSourceClient   `xml:"-" json:"client,omitempty"`
	MCast     *DomainInterfaceSourceMCast    `xml:"-" json:"mcast,omitempty"`
	Network   *DomainInterfaceSourceNetwork  `xml:"-" json:"network,omitempty"`
	Bridge    *DomainInterfaceSourceBridge   `xml:"-" json:"bridge,omitempty"`
	Internal  *DomainInterfaceSourceInternal `xml:"-" json:"internal,omitempty"`
	Direct    *DomainInterfaceSourceDirect   `xml:"-" json:"direct,omitempty"`
	Hostdev   *DomainInterfaceSourceHostdev  `xml:"-" json:"hostdev,omitempty"`
	UDP       *DomainInterfaceSourceUDP      `xml:"-" json:"udp,omitempty"`
	VDPA      *DomainInterfaceSourceVDPA     `xml:"-" json:"vdpa,omitempty"`
}

type DomainInterfaceSourceUser struct {
}

type DomainInterfaceSourceEthernet struct {
	IP    []DomainInterfaceIP    `xml:"ip" json:"ip"`
	Route []DomainInterfaceRoute `xml:"route" json:"route"`
}

type DomainInterfaceSourceServer struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty"`
}

type DomainInterfaceSourceClient struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty"`
}

type DomainInterfaceSourceMCast struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty"`
}

type DomainInterfaceSourceNetwork struct {
	Network   string `xml:"network,attr,omitempty" json:"network,omitempty"`
	PortGroup string `xml:"portgroup,attr,omitempty" json:"portGroup,omitempty"`
	Bridge    string `xml:"bridge,attr,omitempty" json:"bridge,omitempty"`
	PortID    string `xml:"portid,attr,omitempty" json:"portId,omitempty"`
}

type DomainInterfaceSourceBridge struct {
	Bridge string `xml:"bridge,attr" json:"bridge"`
}

type DomainInterfaceSourceInternal struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainInterfaceSourceDirect struct {
	Dev  string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainInterfaceSourceHostdev struct {
	PCI *DomainHostdevSubsysPCISource `xml:"-" json:"pci,omitempty"`
	USB *DomainHostdevSubsysUSBSource `xml:"-" json:"usb,omitempty"`
}

type DomainInterfaceSourceUDP struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty"`
}

type DomainInterfaceSourceVDPA struct {
	Device string `xml:"dev,attr,omitempty" json:"device,omitempty"`
}

type DomainInterfaceSourceLocal struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    int32  `xml:"port,attr,omitempty" json:"port,omitempty"`
}

type DomainInterfaceTarget struct {
	Dev     string `xml:"dev,attr" json:"dev"`
	Managed string `xml:"managed,attr,omitempty" json:"managed,omitempty"`
}

type DomainInterfaceLink struct {
	State string `xml:"state,attr" json:"state"`
}

type DomainDeviceBoot struct {
	Order    int32  `xml:"order,attr" json:"order"`
	LoadParm string `xml:"loadparm,attr,omitempty" json:"loadParm,omitempty"`
}

type DomainInterfaceScript struct {
	Path string `xml:"path,attr" json:"path"`
}

type DomainInterfaceDriver struct {
	Name        string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	TXMode      string                      `xml:"txmode,attr,omitempty" json:"txMode,omitempty"`
	IOEventFD   string                      `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`
	EventIDX    string                      `xml:"event_idx,attr,omitempty" json:"eventIdx,omitempty"`
	Queues      int32                       `xml:"queues,attr,omitempty" json:"queues,omitempty"`
	RXQueueSize int32                       `xml:"rx_queue_size,attr,omitempty" json:"rxQueueSize,omitempty"`
	TXQueueSize int32                       `xml:"tx_queue_size,attr,omitempty" json:"txQueueSize,omitempty"`
	IOMMU       string                      `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS         string                      `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed      string                      `xml:"packed,attr,omitempty" json:"packed,omitempty"`
	Host        *DomainInterfaceDriverHost  `xml:"host" json:"host,omitempty"`
	Guest       *DomainInterfaceDriverGuest `xml:"guest" json:"guest,omitempty"`
}

type DomainInterfaceDriverHost struct {
	CSum     string `xml:"csum,attr,omitempty" json:"csum,omitempty"`
	GSO      string `xml:"gso,attr,omitempty" json:"gso,omitempty"`
	TSO4     string `xml:"tso4,attr,omitempty" json:"tso4,omitempty"`
	TSO6     string `xml:"tso6,attr,omitempty" json:"tso6,omitempty"`
	ECN      string `xml:"ecn,attr,omitempty" json:"ecn,omitempty"`
	UFO      string `xml:"ufo,attr,omitempty" json:"ufo,omitempty"`
	MrgRXBuf string `xml:"mrg_rxbuf,attr,omitempty" json:"mrgRxBuf,omitempty"`
}

type DomainInterfaceDriverGuest struct {
	CSum string `xml:"csum,attr,omitempty" json:"csum,omitempty"`
	TSO4 string `xml:"tso4,attr,omitempty" json:"tso4,omitempty"`
	TSO6 string `xml:"tso6,attr,omitempty" json:"tso6,omitempty"`
	ECN  string `xml:"ecn,attr,omitempty" json:"ecn,omitempty"`
	UFO  string `xml:"ufo,attr,omitempty" json:"ufo,omitempty"`
}

type DomainInterfaceVirtualPort struct {
	Params *DomainInterfaceVirtualPortParams `xml:"parameters" json:"parameters,omitempty"`
}

type DomainInterfaceVirtualPortParams struct {
	Any          *DomainInterfaceVirtualPortParamsAny          `xml:"-" json:"any,omitempty"`
	VEPA8021QBG  *DomainInterfaceVirtualPortParamsVEPA8021QBG  `xml:"-" json:"vepa8021_qbg,omitempty"`
	VNTag8011QBH *DomainInterfaceVirtualPortParamsVNTag8021QBH `xml:"-" json:"vntag8011_qbh,omitempty"`
	OpenVSwitch  *DomainInterfaceVirtualPortParamsOpenVSwitch  `xml:"-" json:"openvswitch,omitempty"`
	MidoNet      *DomainInterfaceVirtualPortParamsMidoNet      `xml:"-" json:"midonet,omitempty"`
}

type DomainInterfaceVirtualPortParamsAny struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
}

type DomainInterfaceVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty"`
}

type DomainInterfaceVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
}

type DomainInterfaceVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
}

type DomainInterfaceVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
}

type DomainInterfaceBandwidthParams struct {
	Average *int32 `xml:"average,attr" json:"average,omitempty"`
	Peak    *int32 `xml:"peak,attr" json:"peak,omitempty"`
	Burst   *int32 `xml:"burst,attr" json:"burst,omitempty"`
	Floor   *int32 `xml:"floor,attr" json:"floor,omitempty"`
}

type DomainInterfaceBandwidth struct {
	Inbound  *DomainInterfaceBandwidthParams `xml:"inbound" json:"inbound,omitempty"`
	Outbound *DomainInterfaceBandwidthParams `xml:"outbound" json:"outbound,omitempty"`
}

type DomainInterfaceVLan struct {
	Trunk string                   `xml:"trunk,attr,omitempty" json:"trunk,omitempty"`
	Tags  []DomainInterfaceVLanTag `xml:"tag" json:"tags"`
}

type DomainInterfaceVLanTag struct {
	ID         int32  `xml:"id,attr" json:"id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode"`
}

type DomainInterfaceGuest struct {
	Dev    string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Actual string `xml:"actual,attr,omitempty" json:"actual,omitempty"`
}

type DomainInterfaceFilterRef struct {
	Filter     string                       `xml:"filter,attr" json:"filter"`
	Parameters []DomainInterfaceFilterParam `xml:"parameter" json:"parameter"`
}

type DomainInterfaceFilterParam struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

type DomainInterfaceBackend struct {
	Tap   string `xml:"tap,attr,omitempty" json:"tap,omitempty"`
	VHost string `xml:"vhost,attr,omitempty" json:"vhost,omitempty"`
}

type DomainInterfaceTune struct {
	SndBuf int32 `xml:"sndbuf" json:"sndbuf"`
}

type DomainInterfaceMTU struct {
	Size int32 `xml:"size,attr" json:"size"`
}

type DomainInterfaceCoalesce struct {
	RX *DomainInterfaceCoalesceRX `xml:"rx" json:"rx,omitempty"`
}

type DomainInterfaceCoalesceRX struct {
	Frames *DomainInterfaceCoalesceRXFrames `xml:"frames" json:"frames,omitempty"`
}

type DomainInterfaceCoalesceRXFrames struct {
	Max *int32 `xml:"max,attr" json:"max,omitempty"`
}

type DomainROM struct {
	Bar     string `xml:"bar,attr,omitempty" json:"bar,omitempty"`
	File    string `xml:"file,attr,omitempty" json:"file,omitempty"`
	Enabled string `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`
}

type DomainInterfaceIP struct {
	Address string `xml:"address,attr" json:"address"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	Peer    string `xml:"peer,attr,omitempty" json:"peer,omitempty"`
}

type DomainInterfaceRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty"`
	Address string `xml:"address,attr" json:"address"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	Gateway string `xml:"gateway,attr" json:"gateway"`
	Metric  int32  `xml:"metric,attr,omitempty" json:"metric,omitempty"`
}

type DomainInterfaceTeaming struct {
	Type       string `xml:"type,attr" json:"type"`
	Persistent string `xml:"persistent,attr,omitempty" json:"persistent,omitempty"`
}

type DomainInterfacePortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty"`
}

type DomainInterface struct {
	XMLName             xml.Name                    `xml:"interface" json:"-"`
	Managed             string                      `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	TrustGuestRXFilters string                      `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty"`
	MAC                 *DomainInterfaceMAC         `xml:"mac" json:"mac,omitempty"`
	Source              *DomainInterfaceSource      `xml:"source" json:"source,omitempty"`
	Boot                *DomainDeviceBoot           `xml:"boot" json:"boot,omitempty"`
	VLan                *DomainInterfaceVLan        `xml:"vlan" json:"vlan,omitempty"`
	VirtualPort         *DomainInterfaceVirtualPort `xml:"virtualport" json:"virtualPort,omitempty"`
	IP                  []DomainInterfaceIP         `xml:"ip" json:"ip"`
	Route               []DomainInterfaceRoute      `xml:"route" json:"route"`
	Script              *DomainInterfaceScript      `xml:"script" json:"script,omitempty"`
	DownScript          *DomainInterfaceScript      `xml:"downscript" json:"downscript,omitempty"`
	BackendDomain       *DomainBackendDomain        `xml:"backenddomain" json:"backenddomain,omitempty"`
	Target              *DomainInterfaceTarget      `xml:"target" json:"target,omitempty"`
	Guest               *DomainInterfaceGuest       `xml:"guest" json:"guest,omitempty"`
	Model               *DomainInterfaceModel       `xml:"model" json:"model,omitempty"`
	Driver              *DomainInterfaceDriver      `xml:"driver" json:"driver,omitempty"`
	Backend             *DomainInterfaceBackend     `xml:"backend" json:"backend,omitempty"`
	FilterRef           *DomainInterfaceFilterRef   `xml:"filterref" json:"filterref,omitempty"`
	Tune                *DomainInterfaceTune        `xml:"tune" json:"tune,omitempty"`
	Teaming             *DomainInterfaceTeaming     `xml:"teaming" json:"teaming,omitempty"`
	Link                *DomainInterfaceLink        `xml:"link" json:"link,omitempty"`
	MTU                 *DomainInterfaceMTU         `xml:"mtu" json:"mtu,omitempty"`
	Bandwidth           *DomainInterfaceBandwidth   `xml:"bandwidth" json:"bandwidth,omitempty"`
	PortOptions         *DomainInterfacePortOptions `xml:"port" json:"portOptions,omitempty"`
	Coalesce            *DomainInterfaceCoalesce    `xml:"coalesce" json:"coalesce,omitempty"`
	ROM                 *DomainROM                  `xml:"rom" json:"rom,omitempty"`
	ACPI                *DomainDeviceACPI           `xml:"acpi" json:"acpi,omitempty"`
	Alias               *DomainAlias                `xml:"alias" json:"alias,omitempty"`
	Address             *DomainAddress              `xml:"address" json:"address,omitempty"`
}

type DomainChardevSource struct {
	Null      *DomainChardevSourceNull      `xml:"-" json:"null,omitempty"`
	VC        *DomainChardevSourceVC        `xml:"-" json:"vc,omitempty"`
	Pty       *DomainChardevSourcePty       `xml:"-" json:"pty,omitempty"`
	Dev       *DomainChardevSourceDev       `xml:"-" json:"dev,omitempty"`
	File      *DomainChardevSourceFile      `xml:"-" json:"file,omitempty"`
	Pipe      *DomainChardevSourcePipe      `xml:"-" json:"pipe,omitempty"`
	StdIO     *DomainChardevSourceStdIO     `xml:"-" json:"stdio,omitempty"`
	UDP       *DomainChardevSourceUDP       `xml:"-" json:"udp,omitempty"`
	TCP       *DomainChardevSourceTCP       `xml:"-" json:"tcp,omitempty"`
	UNIX      *DomainChardevSourceUNIX      `xml:"-" json:"unix,omitempty"`
	SpiceVMC  *DomainChardevSourceSpiceVMC  `xml:"-" json:"spicevmc,omitempty"`
	SpicePort *DomainChardevSourceSpicePort `xml:"-" json:"spiceport,omitempty"`
	NMDM      *DomainChardevSourceNMDM      `xml:"-" json:"nmdm,omitempty"`
}

type DomainChardevSourceNull struct {
}

type DomainChardevSourceVC struct {
}

type DomainChardevSourcePty struct {
	Path     string                 `xml:"path,attr" json:"path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel"`
}

type DomainChardevSourceDev struct {
	Path     string                 `xml:"path,attr" json:"path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"sec_label"`
}

type DomainChardevSourceFile struct {
	Path     string                 `xml:"path,attr" json:"path"`
	Append   string                 `xml:"append,attr,omitempty" json:"append,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel"`
}

type DomainChardevSourcePipe struct {
	Path     string                 `xml:"path,attr" json:"path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel"`
}

type DomainChardevSourceStdIO struct {
}

type DomainChardevSourceUDP struct {
	BindHost       string `xml:"-" json:"bindHost"`
	BindService    string `xml:"-" json:"bindService"`
	ConnectHost    string `xml:"-" json:"connectHost"`
	ConnectService string `xml:"-" json:"connectService"`
}

type DomainChardevSourceReconnect struct {
	Enabled string `xml:"enabled,attr" json:"enabled"`
	Timeout *int32 `xml:"timeout,attr" json:"timeout,omitempty"`
}

type DomainChardevSourceTCP struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Host      string                        `xml:"host,attr,omitempty" json:"host,omitempty"`
	Service   string                        `xml:"service,attr,omitempty" json:"service,omitempty"`
	TLS       string                        `xml:"tls,attr,omitempty" json:"tls,omitempty"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect,omitempty"`
}

type DomainChardevSourceUNIX struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Path      string                        `xml:"path,attr,omitempty" json:"path,omitempty"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect,omitempty"`
	SecLabel  []DomainDeviceSecLabel        `xml:"seclabel" json:"seclabel"`
}

type DomainChardevSourceSpiceVMC struct {
}

type DomainChardevSourceSpicePort struct {
	Channel string `xml:"channel,attr" json:"channel"`
}

type DomainChardevSourceNMDM struct {
	Master string `xml:"master,attr" json:"master"`
	Slave  string `xml:"slave,attr" json:"slave"`
}

type DomainChardevTarget struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty"` // is guest agent connected?
	Port  *int32 `xml:"port,attr" json:"port,omitempty"`
}

type DomainConsoleTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Port *int32 `xml:"port,attr" json:"port,omitempty"`
}

type DomainSerialTarget struct {
	Type  string                   `xml:"type,attr,omitempty" json:"type,omitempty"`
	Port  *int32                   `xml:"port,attr" json:"port,omitempty"`
	Model *DomainSerialTargetModel `xml:"model" json:"model,omitempty"`
}

type DomainSerialTargetModel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainParallelTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Port *int32 `xml:"port,attr" json:"port,omitempty"`
}

type DomainChannelTarget struct {
	VirtIO   *DomainChannelTargetVirtIO   `xml:"-" json:"virtio,omitempty"`
	Xen      *DomainChannelTargetXen      `xml:"-" json:"xen,omitempty"`
	GuestFWD *DomainChannelTargetGuestFWD `xml:"-" json:"guestfwd,omitempty"`
}

type DomainChannelTargetVirtIO struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty"` // is guest agent connected?
}

type DomainChannelTargetXen struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty"` // is guest agent connected?
}

type DomainChannelTargetGuestFWD struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Port    string `xml:"port,attr,omitempty" json:"port,omitempty"`
}

type DomainAlias struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainDeviceACPI struct {
	Index int32 `xml:"index,attr,omitempty" json:"index,omitempty"`
}

type DomainAddressPCI struct {
	Domain        *int32             `xml:"domain,attr" json:"domain,omitempty"`
	Bus           *int32             `xml:"bus,attr" json:"bus,omitempty"`
	Slot          *int32             `xml:"slot,attr" json:"slot,omitempty"`
	Function      *int32             `xml:"function,attr" json:"function,omitempty"`
	MultiFunction string             `xml:"multifunction,attr,omitempty" json:"multifunction,omitempty"`
	ZPCI          *DomainAddressZPCI `xml:"zpci" json:"zpci,omitempty"`
}

type DomainAddressZPCI struct {
	UID *int32 `xml:"uid,attr,omitempty" json:"uid,omitempty"`
	FID *int32 `xml:"fid,attr,omitempty" json:"fid,omitempty"`
}

type DomainAddressUSB struct {
	Bus    *int32 `xml:"bus,attr" json:"bus,omitempty"`
	Port   string `xml:"port,attr,omitempty" json:"port,omitempty"`
	Device *int32 `xml:"device,attr" json:"device,omitempty"`
}

type DomainAddressDrive struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty"`
	Bus        *int32 `xml:"bus,attr" json:"bus,omitempty"`
	Target     *int32 `xml:"target,attr" json:"target,omitempty"`
	Unit       *int32 `xml:"unit,attr" json:"unit,omitempty"`
}

type DomainAddressDIMM struct {
	Slot *int32 `xml:"slot,attr" json:"slot,omitempty"`
	Base *int64 `xml:"base,attr" json:"base,omitempty"`
}

type DomainAddressISA struct {
	IOBase *int32 `xml:"iobase,attr" json:"iobase,omitempty"`
	IRQ    *int32 `xml:"irq,attr" json:"irq,omitempty"`
}

type DomainAddressVirtioMMIO struct {
}

type DomainAddressCCW struct {
	CSSID *int32 `xml:"cssid,attr" json:"cssid,omitempty"`
	SSID  *int32 `xml:"ssid,attr" json:"ssid,omitempty"`
	DevNo *int32 `xml:"devno,attr" json:"devno,omitempty"`
}

type DomainAddressVirtioSerial struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty"`
	Bus        *int32 `xml:"bus,attr" json:"bus,omitempty"`
	Port       *int32 `xml:"port,attr" json:"port,omitempty"`
}

type DomainAddressSpaprVIO struct {
	Reg *int64 `xml:"reg,attr" json:"reg,omitempty"`
}

type DomainAddressCCID struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty"`
	Slot       *int32 `xml:"slot,attr" json:"slot,omitempty"`
}

type DomainAddressVirtioS390 struct {
}

type DomainAddressUnassigned struct {
}

type DomainAddress struct {
	PCI          *DomainAddressPCI          `json:"pci,omitempty"`
	Drive        *DomainAddressDrive        `json:"drive,omitempty"`
	VirtioSerial *DomainAddressVirtioSerial `json:"virtioSerial,omitempty"`
	CCID         *DomainAddressCCID         `json:"ccid,omitempty"`
	USB          *DomainAddressUSB          `json:"usb,omitempty"`
	SpaprVIO     *DomainAddressSpaprVIO     `json:"spaprvio,omitempty"`
	VirtioS390   *DomainAddressVirtioS390   `json:"virtioS390,omitempty"`
	CCW          *DomainAddressCCW          `json:"ccw,omitempty"`
	VirtioMMIO   *DomainAddressVirtioMMIO   `json:"virtiommio,omitempty"`
	ISA          *DomainAddressISA          `json:"isa,omitempty"`
	DIMM         *DomainAddressDIMM         `json:"dimm,omitempty"`
	Unassigned   *DomainAddressUnassigned   `json:"unassigned,omitempty"`
}

type DomainChardevLog struct {
	File   string `xml:"file,attr" json:"file"`
	Append string `xml:"append,attr,omitempty" json:"append,omitempty"`
}

type DomainConsole struct {
	XMLName  xml.Name               `xml:"console" json:"-"`
	TTY      string                 `xml:"tty,attr,omitempty" json:"tty,omitempty"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
	Target   *DomainConsoleTarget   `xml:"target" json:"target,omitempty"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainSerial struct {
	XMLName  xml.Name               `xml:"serial" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
	Target   *DomainSerialTarget    `xml:"target" json:"target,omitempty"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainParallel struct {
	XMLName  xml.Name               `xml:"parallel" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
	Target   *DomainParallelTarget  `xml:"target" json:"target,omitempty"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainChardevProtocol struct {
	Type string `xml:"type,attr" json:"type"`
}

type DomainChannel struct {
	XMLName  xml.Name               `xml:"channel" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
	Target   *DomainChannelTarget   `xml:"target" json:"target,omitempty"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainRedirDev struct {
	XMLName  xml.Name               `xml:"redirdev" json:"-"`
	Bus      string                 `xml:"bus,attr,omitempty" json:"bus,omitempty"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
	Boot     *DomainDeviceBoot      `xml:"boot" json:"boot,omitempty"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainRedirFilter struct {
	USB []DomainRedirFilterUSB `xml:"usbdev" json:"usb"`
}

type DomainRedirFilterUSB struct {
	Class   *int32 `xml:"class,attr" json:"class,omitempty"`
	Vendor  *int32 `xml:"vendor,attr" json:"vendor,omitempty"`
	Product *int32 `xml:"product,attr" json:"product,omitempty"`
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
	Allow   string `xml:"allow,attr" json:"allow"`
}

type DomainInput struct {
	XMLName xml.Name           `xml:"input" json:"-"`
	Type    string             `xml:"type,attr" json:"type"`
	Bus     string             `xml:"bus,attr,omitempty" json:"bus,omitempty"`
	Model   string             `xml:"model,attr,omitempty" json:"model,omitempty"`
	Driver  *DomainInputDriver `xml:"driver" json:"driver,omitempty"`
	Source  *DomainInputSource `xml:"source" json:"source,omitempty"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty"`
}

type DomainInputDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainInputSource struct {
	EVDev string `xml:"evdev,attr" json:"evdev"`
}

type DomainGraphicListenerAddress struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
}

type DomainGraphicListenerNetwork struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Network string `xml:"network,attr,omitempty" json:"network,omitempty"`
}

type DomainGraphicListenerSocket struct {
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty"`
}

type DomainGraphicListener struct {
	Address *DomainGraphicListenerAddress `xml:"-" json:"address,omitempty"`
	Network *DomainGraphicListenerNetwork `xml:"-" json:"network,omitempty"`
	Socket  *DomainGraphicListenerSocket  `xml:"-" json:"socket,omitempty"`
}

type DomainGraphicChannel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainGraphicFileTransfer struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty"`
}

type DomainGraphicsSDLGL struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty"`
}

type DomainGraphicSDL struct {
	Display    string               `xml:"display,attr,omitempty" json:"display,omitempty"`
	XAuth      string               `xml:"xauth,attr,omitempty" json:"xauth,omitempty"`
	FullScreen string               `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty"`
	GL         *DomainGraphicsSDLGL `xml:"gl" json:"gl,omitempty"`
}

type DomainGraphicVNC struct {
	Socket        string                  `xml:"socket,attr,omitempty" json:"socket,omitempty"`
	Port          int                     `xml:"port,attr,omitempty" json:"port,omitempty"`
	AutoPort      string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty"`
	WebSocket     int                     `xml:"websocket,attr,omitempty" json:"websocket,omitempty"`
	Keymap        string                  `xml:"keymap,attr,omitempty" json:"keymap,omitempty"`
	SharePolicy   string                  `xml:"sharePolicy,attr,omitempty" json:"sharePolicy,omitempty"`
	Passwd        string                  `xml:"passwd,attr,omitempty" json:"passwd,omitempty"`
	PasswdValidTo string                  `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty"`
	Connected     string                  `xml:"connected,attr,omitempty" json:"connected,omitempty"`
	PowerControl  string                  `xml:"powerControl,attr,omitempty" json:"powerControl,omitempty"`
	Listen        string                  `xml:"listen,attr,omitempty" json:"listen,omitempty"`
	Listeners     []DomainGraphicListener `xml:"listen" json:"listeners"`
}

type DomainGraphicRDP struct {
	Port        int                     `xml:"port,attr,omitempty" json:"port,omitempty"`
	AutoPort    string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty"`
	ReplaceUser string                  `xml:"replaceUser,attr,omitempty" json:"replaceuser,omitempty"`
	MultiUser   string                  `xml:"multiUser,attr,omitempty" json:"multiuser,omitempty"`
	Listen      string                  `xml:"listen,attr,omitempty" json:"listen,omitempty"`
	Listeners   []DomainGraphicListener `xml:"listen" json:"listeners"`
}

type DomainGraphicDesktop struct {
	Display    string `xml:"display,attr,omitempty" json:"display,omitempty"`
	FullScreen string `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty"`
}

type DomainGraphicSpiceChannel struct {
	Name string `xml:"name,attr" json:"name"`
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainGraphicSpiceImage struct {
	Compression string `xml:"compression,attr" json:"compression"`
}

type DomainGraphicSpiceJPEG struct {
	Compression string `xml:"compression,attr" json:"compression"`
}

type DomainGraphicSpiceZLib struct {
	Compression string `xml:"compression,attr" json:"compression"`
}

type DomainGraphicSpicePlayback struct {
	Compression string `xml:"compression,attr" json:"compression"`
}

type DomainGraphicSpiceStreaming struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainGraphicSpiceMouse struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainGraphicSpiceClipBoard struct {
	CopyPaste string `xml:"copypaste,attr" json:"copypaste"`
}

type DomainGraphicSpiceFileTransfer struct {
	Enable string `xml:"enable,attr" json:"enable"`
}

type DomainGraphicSpiceGL struct {
	Enable     string `xml:"enable,attr,omitempty" json:"enable,omitempty"`
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty"`
}

type DomainGraphicSpice struct {
	Port          int                             `xml:"port,attr,omitempty" json:"port,omitempty"`
	TLSPort       int                             `xml:"tlsPort,attr,omitempty" json:"tlsPort,omitempty"`
	AutoPort      string                          `xml:"autoport,attr,omitempty" json:"autoPort,omitempty"`
	Listen        string                          `xml:"listen,attr,omitempty" json:"listen,omitempty"`
	Keymap        string                          `xml:"keymap,attr,omitempty" json:"keymap,omitempty"`
	DefaultMode   string                          `xml:"defaultMode,attr,omitempty" json:"defaultMode,omitempty"`
	Passwd        string                          `xml:"passwd,attr,omitempty" json:"passwd,omitempty"`
	PasswdValidTo string                          `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty"`
	Connected     string                          `xml:"connected,attr,omitempty" json:"connected,omitempty"`
	Listeners     []DomainGraphicListener         `xml:"listen" json:"listeners"`
	Channel       []DomainGraphicSpiceChannel     `xml:"channel" json:"channel"`
	Image         *DomainGraphicSpiceImage        `xml:"image" json:"image,omitempty"`
	JPEG          *DomainGraphicSpiceJPEG         `xml:"jpeg" json:"jpeg,omitempty"`
	ZLib          *DomainGraphicSpiceZLib         `xml:"zlib" json:"zlib,omitempty"`
	Playback      *DomainGraphicSpicePlayback     `xml:"playback" json:"playback,omitempty"`
	Streaming     *DomainGraphicSpiceStreaming    `xml:"streaming" json:"streaming,omitempty"`
	Mouse         *DomainGraphicSpiceMouse        `xml:"mouse" json:"mouse,omitempty"`
	ClipBoard     *DomainGraphicSpiceClipBoard    `xml:"clipboard" json:"clipboard,omitempty"`
	FileTransfer  *DomainGraphicSpiceFileTransfer `xml:"filetransfer" json:"filetransfer,omitempty"`
	GL            *DomainGraphicSpiceGL           `xml:"gl" json:"gl,omitempty"`
}

type DomainGraphicEGLHeadlessGL struct {
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty"`
}

type DomainGraphicEGLHeadless struct {
	GL *DomainGraphicEGLHeadlessGL `xml:"gl" json:"gl,omitempty"`
}

type DomainGraphicAudio struct {
	ID int32 `xml:"id,attr,omitempty" json:"id,omitempty"`
}

type DomainGraphic struct {
	XMLName     xml.Name                  `xml:"graphics" json:"-"`
	SDL         *DomainGraphicSDL         `xml:"-" json:"sdl,omitempty"`
	VNC         *DomainGraphicVNC         `xml:"-" json:"vnc,omitempty"`
	RDP         *DomainGraphicRDP         `xml:"-" json:"rdp,omitempty"`
	Desktop     *DomainGraphicDesktop     `xml:"-" json:"desktop,omitempty"`
	Spice       *DomainGraphicSpice       `xml:"-" json:"spice,omitempty"`
	EGLHeadless *DomainGraphicEGLHeadless `xml:"-" json:"eglHeadless,omitempty"`
	Audio       *DomainGraphicAudio       `xml:"audio" json:"audio,omitempty"`
}

type DomainVideoAccel struct {
	Accel3D    string `xml:"accel3d,attr,omitempty" json:"accel3d,omitempty"`
	Accel2D    string `xml:"accel2d,attr,omitempty" json:"accel2d,omitempty"`
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty"`
}

type DomainVideoResolution struct {
	X int32 `xml:"x,attr" json:"x"`
	Y int32 `xml:"y,attr" json:"y"`
}

type DomainVideoModel struct {
	Type       string                 `xml:"type,attr" json:"type"`
	Heads      int32                  `xml:"heads,attr,omitempty" json:"heads,omitempty"`
	Ram        int32                  `xml:"ram,attr,omitempty" json:"ram,omitempty"`
	VRam       int32                  `xml:"vram,attr,omitempty" json:"vram,omitempty"`
	VRam64     int32                  `xml:"vram64,attr,omitempty" json:"vram64,omitempty"`
	VGAMem     int32                  `xml:"vgamem,attr,omitempty" json:"vgamem,omitempty"`
	Primary    string                 `xml:"primary,attr,omitempty" json:"primary,omitempty"`
	Accel      *DomainVideoAccel      `xml:"acceleration" json:"accel,omitempty"`
	Resolution *DomainVideoResolution `xml:"resolution" json:"resolution,omitempty"`
}

type DomainVideo struct {
	XMLName xml.Name           `xml:"video" json:"-"`
	Model   DomainVideoModel   `xml:"model" json:"model"`
	Driver  *DomainVideoDriver `xml:"driver" json:"driver,omitempty"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty"`
}

type DomainVideoDriver struct {
	Name    string `xml:"name,attr,omitempty" json:"name,omitempty"`
	VGAConf string `xml:"vgaconf,attr,omitempty" json:"vgaconf,omitempty"`
	IOMMU   string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS     string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed  string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainMemBalloonStats struct {
	Period int32 `xml:"period,attr" json:"period,omitempty"`
}

type DomainMemBalloon struct {
	XMLName           xml.Name                `xml:"memballoon" json:"-"`
	Model             string                  `xml:"model,attr" json:"model"`
	AutoDeflate       string                  `xml:"autodeflate,attr,omitempty" json:"autodeflate,omitempty"`
	FreePageReporting string                  `xml:"freePageReporting,attr,omitempty" json:"freePageReporting,omitempty"`
	Driver            *DomainMemBalloonDriver `xml:"driver" json:"driver,omitempty"`
	Stats             *DomainMemBalloonStats  `xml:"stats" json:"stats,omitempty"`
	ACPI              *DomainDeviceACPI       `xml:"acpi" json:"acpi,omitempty"`
	Alias             *DomainAlias            `xml:"alias" json:"alias,omitempty"`
	Address           *DomainAddress          `xml:"address" json:"address,omitempty"`
}

type DomainVSockCID struct {
	Auto    string `xml:"auto,attr,omitempty" json:"auto,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
}

type DomainVSockDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainVSock struct {
	XMLName xml.Name           `xml:"vsock" json:"-"`
	Model   string             `xml:"model,attr,omitempty" json:"model,omitempty"`
	CID     *DomainVSockCID    `xml:"cid" json:"cid,omitempty"`
	Driver  *DomainVSockDriver `xml:"driver" json:"driver,omitempty"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty"`
}

type DomainMemBalloonDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainPanic struct {
	XMLName xml.Name          `xml:"panic" json:"-"`
	Model   string            `xml:"model,attr,omitempty" json:"model,omitempty"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainSoundCodec struct {
	Type string `xml:"type,attr" json:"type"`
}

type DomainSound struct {
	XMLName xml.Name           `xml:"sound" json:"-"`
	Model   string             `xml:"model,attr" json:"model"`
	Codec   []DomainSoundCodec `xml:"codec" json:"codec"`
	Audio   *DomainSoundAudio  `xml:"audio" json:"audio,omitempty"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty"`
}

type DomainSoundAudio struct {
	ID int32 `xml:"id,attr" json:"id"`
}

type DomainAudio struct {
	XMLName    xml.Name               `xml:"audio" json:"-"`
	ID         int                    `xml:"id,attr" json:"id"`
	None       *DomainAudioNone       `xml:"-" json:"none,omitempty"`
	ALSA       *DomainAudioALSA       `xml:"-" json:"alsa,omitempty"`
	CoreAudio  *DomainAudioCoreAudio  `xml:"-" json:"coreAudio,omitempty"`
	Jack       *DomainAudioJack       `xml:"-" json:"jack,omitempty"`
	OSS        *DomainAudioOSS        `xml:"-" json:"oss,omitempty"`
	PulseAudio *DomainAudioPulseAudio `xml:"-" json:"pulseAudio,omitempty"`
	SDL        *DomainAudioSDL        `xml:"-" json:"sdl,omitempty"`
	SPICE      *DomainAudioSPICE      `xml:"-" json:"spice,omitempty"`
	File       *DomainAudioFile       `xml:"-" json:"file,omitempty"`
}

type DomainAudioChannel struct {
	MixingEngine  string                      `xml:"mixingEngine,attr,omitempty" json:"mixingEngine,omitempty"`
	FixedSettings string                      `xml:"fixedSettings,attr,omitempty" json:"fixedSettings,omitempty"`
	Voices        int32                       `xml:"voices,attr,omitempty" json:"voices,omitempty"`
	Settings      *DomainAudioChannelSettings `xml:"settings" json:"settings,omitempty"`
	BufferLength  int32                       `xml:"bufferLength,attr,omitempty" json:"bufferLength,omitempty"`
}

type DomainAudioChannelSettings struct {
	Frequency int32  `xml:"frequency,attr,omitempty" json:"frequency,omitempty"`
	Channels  int32  `xml:"channels,attr,omitempty" json:"channels,omitempty"`
	Format    string `xml:"format,attr,omitempty" json:"format,omitempty"`
}

type DomainAudioNone struct {
	Input  *DomainAudioNoneChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioNoneChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioNoneChannel struct {
	DomainAudioChannel `protobuf:"bytes,1,opt,name=domainAudioChannel"`
}

type DomainAudioALSA struct {
	Input  *DomainAudioALSAChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioALSAChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioALSAChannel struct {
	DomainAudioChannel `json:",inline"`
	Dev                string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
}

type DomainAudioCoreAudio struct {
	Input  *DomainAudioCoreAudioChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioCoreAudioChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioCoreAudioChannel struct {
	DomainAudioChannel `json:",inline"`
	BufferCount        int32 `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty"`
}

type DomainAudioJack struct {
	Input  *DomainAudioJackChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioJackChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioJackChannel struct {
	DomainAudioChannel `json:",inline"`
	ServerName         string `xml:"serverName,attr,omitempty" json:"serverName,omitempty"`
	ClientName         string `xml:"clientName,attr,omitempty" json:"clientName,omitempty"`
	ConnectPorts       string `xml:"connectPorts,attr,omitempty" json:"connectPorts,omitempty"`
	ExactName          string `xml:"exactName,attr,omitempty" json:"exactName,omitempty"`
}

type DomainAudioOSS struct {
	TryMMap   string `xml:"tryMMap,attr,omitempty" json:"tryMMap,omitempty"`
	Exclusive string `xml:"exclusive,attr,omitempty" json:"exclusive,omitempty"`
	DSPPolicy *int32 `xml:"dspPolicy,attr" json:"dspPolicy,omitempty"`

	Input  *DomainAudioOSSChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioOSSChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioOSSChannel struct {
	DomainAudioChannel `json:",inline"`
	Dev                string `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	BufferCount        int32  `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty"`
	TryPoll            string `xml:"tryPoll,attr,omitempty" json:"tryPoll,omitempty"`
}

type DomainAudioPulseAudio struct {
	ServerName string                        `xml:"serverName,attr,omitempty" json:"serverName,omitempty"`
	Input      *DomainAudioPulseAudioChannel `xml:"input" json:"input,omitempty"`
	Output     *DomainAudioPulseAudioChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioPulseAudioChannel struct {
	DomainAudioChannel `json:",inline"`
	Name               string `xml:"name,attr,omitempty" json:"name,omitempty"`
	StreamName         string `xml:"streamName,attr,omitempty" json:"streamName,omitempty"`
	Latency            int32  `xml:"latency,attr,omitempty" json:"latency,omitempty"`
}

type DomainAudioSDL struct {
	Driver string                 `xml:"driver,attr,omitempty" json:"driver,omitempty"`
	Input  *DomainAudioSDLChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioSDLChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioSDLChannel struct {
	DomainAudioChannel `json:",inline"`
	BufferCount        int32 `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty"`
}

type DomainAudioSPICE struct {
	Input  *DomainAudioSPICEChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioSPICEChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioSPICEChannel struct {
	DomainAudioChannel `json:",inline"`
}

type DomainAudioFile struct {
	Path   string                  `xml:"path,attr,omitempty" json:"path,omitempty"`
	Input  *DomainAudioFileChannel `xml:"input" json:"input,omitempty"`
	Output *DomainAudioFileChannel `xml:"output" json:"output,omitempty"`
}

type DomainAudioFileChannel struct {
	DomainAudioChannel `json:",inline"`
}

type DomainRNGRate struct {
	Bytes  int32 `xml:"bytes,attr" json:"bytes"`
	Period int32 `xml:"period,attr,omitempty" json:"period,omitempty"`
}

type DomainRNGBackend struct {
	Random  *DomainRNGBackendRandom  `xml:"-" json:"random,omitempty"`
	EGD     *DomainRNGBackendEGD     `xml:"-" json:"egd,omitempty"`
	BuiltIn *DomainRNGBackendBuiltIn `xml:"-" json:"builtin,omitempty"`
}

type DomainRNGBackendEGD struct {
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty"`
}

type DomainRNGBackendRandom struct {
	Device string `xml:",chardata" json:"device"`
}

type DomainRNGBackendBuiltIn struct {
}

type DomainRNG struct {
	XMLName xml.Name          `xml:"rng" json:"-"`
	Model   string            `xml:"model,attr" json:"model"`
	Driver  *DomainRNGDriver  `xml:"driver" json:"driver,omitempty"`
	Rate    *DomainRNGRate    `xml:"rate" json:"rate,omitempty"`
	Backend *DomainRNGBackend `xml:"backend" json:"backend,omitempty"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainRNGDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty"`
}

type DomainHostdevSubsysUSB struct {
	Source *DomainHostdevSubsysUSBSource `xml:"source" json:"source,omitempty"`
}

type DomainHostdevSubsysUSBSource struct {
	Address *DomainAddressUSB `xml:"address" json:"address,omitempty"`
}

type DomainHostdevSubsysSCSI struct {
	SGIO      string                         `xml:"sgio,attr,omitempty" json:"sgio,omitempty"`
	RawIO     string                         `xml:"rawio,attr,omitempty" json:"rawio,omitempty"`
	Source    *DomainHostdevSubsysSCSISource `xml:"source" json:"source,omitempty"`
	ReadOnly  *DomainDiskReadOnly            `xml:"readonly" json:"readonly,omitempty"`
	Shareable *DomainDiskShareable           `xml:"shareable" json:"shareable,omitempty"`
}

type DomainHostdevSubsysSCSISource struct {
	Host  *DomainHostdevSubsysSCSISourceHost  `xml:"-" json:"host,omitempty"`
	ISCSI *DomainHostdevSubsysSCSISourceISCSI `xml:"-" json:"iscsi,omitempty"`
}

type DomainHostdevSubsysSCSIAdapter struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainHostdevSubsysSCSISourceHost struct {
	Adapter *DomainHostdevSubsysSCSIAdapter `xml:"adapter" json:"adapter,omitempty"`
	Address *DomainAddressDrive             `xml:"address" json:"address,omitempty"`
}

type DomainHostdevSubsysSCSISourceISCSI struct {
	Name      string                                  `xml:"name,attr" json:"name"`
	Host      []DomainDiskSourceHost                  `xml:"host" json:"host"`
	Auth      *DomainDiskAuth                         `xml:"auth" json:"auth,omitempty"`
	Initiator *DomainHostdevSubsysSCSISourceInitiator `xml:"initiator" json:"initiator,omitempty"`
}

type DomainHostdevSubsysSCSISourceInitiator struct {
	IQN DomainHostdevSubsysSCSISourceIQN `xml:"iqn" json:"iqn"`
}

type DomainHostdevSubsysSCSISourceIQN struct {
	Name string `xml:"name,attr" json:"name"`
}

type DomainHostdevSubsysSCSIHost struct {
	Model  string                             `xml:"model,attr,omitempty" json:"model,omitempty"`
	Source *DomainHostdevSubsysSCSIHostSource `xml:"source" json:"source,omitempty"`
}

type DomainHostdevSubsysSCSIHostSource struct {
	Protocol string `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	WWPN     string `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty"`
}

type DomainHostdevSubsysPCISource struct {
	WriteFiltering string            `xml:"writeFiltering,attr,omitempty" json:"writeFiltering,omitempty"`
	Address        *DomainAddressPCI `xml:"address" json:"address,omitempty"`
}

type DomainHostdevSubsysPCIDriver struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainHostdevSubsysPCI struct {
	Driver  *DomainHostdevSubsysPCIDriver `xml:"driver" json:"driver,omitempty"`
	Source  *DomainHostdevSubsysPCISource `xml:"source" json:"source,omitempty"`
	Teaming *DomainInterfaceTeaming       `xml:"teaming" json:"teaming,omitempty"`
}

type DomainAddressMDev struct {
	UUID string `xml:"uuid,attr" json:"uuid"`
}

type DomainHostdevSubsysMDevSource struct {
	Address *DomainAddressMDev `xml:"address" json:"address,omitempty"`
}

type DomainHostdevSubsysMDev struct {
	Model   string                         `xml:"model,attr,omitempty" json:"model,omitempty"`
	Display string                         `xml:"display,attr,omitempty" json:"display,omitempty"`
	RamFB   string                         `xml:"ramfb,attr,omitempty" json:"ramfb,omitempty"`
	Source  *DomainHostdevSubsysMDevSource `xml:"source" json:"source,omitempty"`
}

type DomainHostdevCapsStorage struct {
	Source *DomainHostdevCapsStorageSource `xml:"source" json:"source,omitempty"`
}

type DomainHostdevCapsStorageSource struct {
	Block string `xml:"block" json:"block"`
}

type DomainHostdevCapsMisc struct {
	Source *DomainHostdevCapsMiscSource `xml:"source" json:"source,omitempty"`
}

type DomainHostdevCapsMiscSource struct {
	Char string `xml:"char" json:"char"`
}

type DomainIP struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty"`
	Prefix  *int32 `xml:"prefix,attr" json:"prefix,omitempty"`
}

type DomainRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty"`
}

type DomainHostdevCapsNet struct {
	Source *DomainHostdevCapsNetSource `xml:"source" json:"source,omitempty"`
	IP     []DomainIP                  `xml:"ip" json:"ip"`
	Route  []DomainRoute               `xml:"route" json:"route"`
}

type DomainHostdevCapsNetSource struct {
	Interface string `xml:"interface" json:"interface"`
}

type DomainHostdev struct {
	Managed        string                       `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	SubsysUSB      *DomainHostdevSubsysUSB      `xml:"-" json:"usb,omitempty"`
	SubsysSCSI     *DomainHostdevSubsysSCSI     `xml:"-" json:"scsi,omitempty"`
	SubsysSCSIHost *DomainHostdevSubsysSCSIHost `xml:"-" json:"scsiHost,omitempty"`
	SubsysPCI      *DomainHostdevSubsysPCI      `xml:"-" json:"pci,omitempty"`
	SubsysMDev     *DomainHostdevSubsysMDev     `xml:"-" json:"mdev,omitempty"`
	CapsStorage    *DomainHostdevCapsStorage    `xml:"-" json:"storage,omitempty"`
	CapsMisc       *DomainHostdevCapsMisc       `xml:"-" json:"misc,omitempty"`
	CapsNet        *DomainHostdevCapsNet        `xml:"-" json:"net,omitempty"`
	Boot           *DomainDeviceBoot            `xml:"boot" json:"boot,omitempty"`
	ROM            *DomainROM                   `xml:"rom" json:"rom,omitempty"`
	ACPI           *DomainDeviceACPI            `xml:"acpi" json:"acpi,omitempty"`
	Alias          *DomainAlias                 `xml:"alias" json:"alias,omitempty"`
	Address        *DomainAddress               `xml:"address" json:"address,omitempty"`
}

type DomainMemorydevSource struct {
	NodeMask  string                          `xml:"nodemask,omitempty"`
	PageSize  *DomainMemorydevSourcePagesize  `xml:"pagesize" json:"pageSize,omitempty"`
	Path      string                          `xml:"path,omitempty" json:"path,omitempty"`
	AlignSize *DomainMemorydevSourceAlignsize `xml:"alignsize" json:"alignSize,omitempty"`
	PMem      *DomainMemorydevSourcePMem      `xml:"pmem" json:"pmem,omitempty"`
}

type DomainMemorydevSourcePMem struct {
}

type DomainMemorydevSourcePagesize struct {
	Value int64  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainMemorydevSourceAlignsize struct {
	Value int64  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainMemorydevTargetNode struct {
	Value int32 `xml:",chardata" json:"value"`
}

type DomainMemorydevTargetReadOnly struct {
}

type DomainMemorydevTargetSize struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainMemorydevTargetLabel struct {
	Size *DomainMemorydevTargetSize `xml:"size" json:"size,omitempty"`
}

type DomainMemorydevTarget struct {
	Size     *DomainMemorydevTargetSize     `xml:"size" json:"size,omitempty"`
	Node     *DomainMemorydevTargetNode     `xml:"node" json:"node,omitempty"`
	Label    *DomainMemorydevTargetLabel    `xml:"label" json:"label,omitempty"`
	ReadOnly *DomainMemorydevTargetReadOnly `xml:"readonly" json:"readonly,omitempty"`
}

type DomainMemorydev struct {
	XMLName xml.Name               `xml:"memory" json:"-"`
	Model   string                 `xml:"model,attr" json:"model"`
	Access  string                 `xml:"access,attr,omitempty" json:"access,omitempty"`
	Discard string                 `xml:"discard,attr,omitempty" json:"discard,omitempty"`
	UUID    string                 `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Source  *DomainMemorydevSource `xml:"source" json:"source,omitempty"`
	Target  *DomainMemorydevTarget `xml:"target" json:"target,omitempty"`
	ACPI    *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias           `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress         `xml:"address" json:"address,omitempty"`
}

type DomainWatchdog struct {
	XMLName xml.Name          `xml:"watchdog" json:"-"`
	Model   string            `xml:"model,attr" json:"model"`
	Action  string            `xml:"action,attr,omitempty" json:"action,omitempty"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainHub struct {
	Type    string            `xml:"type,attr" json:"type"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainIOMMU struct {
	Model  string             `xml:"model,attr" json:"model"`
	Driver *DomainIOMMUDriver `xml:"driver" json:"driver,omitempty"`
}

type DomainIOMMUDriver struct {
	IntRemap    string `xml:"intremap,attr,omitempty" json:"intremap,omitempty"`
	CachingMode string `xml:"caching_mode,attr,omitempty" json:"cachingMode,omitempty"`
	EIM         string `xml:"eim,attr,omitempty" json:"eim,omitempty"`
	IOTLB       string `xml:"iotlb,attr,omitempty" json:"iotlb,omitempty"`
	AWBits      int32  `xml:"aw_bits,attr,omitempty" json:"awBits,omitempty"`
}

type DomainNVRAM struct {
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainLease struct {
	Lockspace string             `xml:"lockspace" json:"lockspace"`
	Key       string             `xml:"key" json:"key"`
	Target    *DomainLeaseTarget `xml:"target" json:"target,omitempty"`
}

type DomainLeaseTarget struct {
	Path   string `xml:"path,attr" json:"path"`
	Offset int64  `xml:"offset,attr,omitempty" json:"offset,omitempty"`
}

type DomainSmartcard struct {
	XMLName     xml.Name                  `xml:"smartcard" json:"-"`
	Passthrough *DomainChardevSource      `xml:"source" json:"passthrough,omitempty"`
	Protocol    *DomainChardevProtocol    `xml:"protocol" json:"protocol,omitempty"`
	Host        *DomainSmartcardHost      `xml:"-" json:"host,omitempty"`
	HostCerts   []DomainSmartcardHostCert `xml:"certificate" json:"certificate,omitempty"`
	Database    string                    `xml:"database,omitempty" json:"database,omitempty"`
	ACPI        *DomainDeviceACPI         `xml:"acpi" json:"acpi,omitempty"`
	Alias       *DomainAlias              `xml:"alias" json:"alias,omitempty"`
	Address     *DomainAddress            `xml:"address" json:"address,omitempty"`
}

type DomainSmartcardHost struct {
}

type DomainSmartcardHostCert struct {
	File string `xml:",chardata" json:"file"`
}

type DomainTPM struct {
	XMLName xml.Name          `xml:"tpm" json:"-"`
	Model   string            `xml:"model,attr,omitempty" json:"model,omitempty"`
	Backend *DomainTPMBackend `xml:"backend" json:"backend,omitempty"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty"`
}

type DomainTPMBackend struct {
	Passthrough *DomainTPMBackendPassthrough `xml:"-" json:"passthrough,omitempty"`
	Emulator    *DomainTPMBackendEmulator    `xml:"-" json:"emulator,omitempty"`
}

type DomainTPMBackendPassthrough struct {
	Device *DomainTPMBackendDevice `xml:"device" json:"device,omitempty"`
}

type DomainTPMBackendEmulator struct {
	Version         string                      `xml:"version,attr,omitempty" json:"version,omitempty"`
	Encryption      *DomainTPMBackendEncryption `xml:"encryption" json:"encryption,omitempty"`
	PersistentState string                      `xml:"persistent_state,attr,omitempty" json:"persistentState,omitempty"`
}

type DomainTPMBackendEncryption struct {
	Secret string `xml:"secret,attr" json:"secret"`
}

type DomainTPMBackendDevice struct {
	Path string `xml:"path,attr" json:"path"`
}

type DomainShmem struct {
	XMLName xml.Name           `xml:"shmem" json:"-"`
	Name    string             `xml:"name,attr" json:"name"`
	Role    string             `xml:"role,attr,omitempty" json:"role,omitempty"`
	Size    *DomainShmemSize   `xml:"size" json:"size,omitempty"`
	Model   *DomainShmemModel  `xml:"model" json:"model,omitempty"`
	Server  *DomainShmemServer `xml:"server" json:"server,omitempty"`
	MSI     *DomainShmemMSI    `xml:"msi" json:"msi,omitempty"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty"`
}

type DomainShmemSize struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainShmemModel struct {
	Type string `xml:"type,attr" json:"type"`
}

type DomainShmemServer struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`
}

type DomainShmemMSI struct {
	Enabled   string `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`
	Vectors   int32  `xml:"vectors,attr,omitempty" json:"vectors,omitempty"`
	IOEventFD string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty"`
}

type DomainDeviceList struct {
	Emulator     string              `xml:"emulator,omitempty" json:"emulator,omitempty"`
	Disks        []DomainDisk        `xml:"disk" json:"disks"`
	Controllers  []DomainController  `xml:"controller" json:"controllers"`
	Leases       []DomainLease       `xml:"lease" json:"leases"`
	Filesystems  []DomainFilesystem  `xml:"filesystem" json:"filesystems"`
	Interfaces   []DomainInterface   `xml:"interface" json:"interfaces"`
	Smartcards   []DomainSmartcard   `xml:"smartcard" json:"smartcards"`
	Serials      []DomainSerial      `xml:"serial" json:"serials"`
	Parallels    []DomainParallel    `xml:"parallel" json:"parallels"`
	Consoles     []DomainConsole     `xml:"console" json:"consoles"`
	Channels     []DomainChannel     `xml:"channel" json:"channels"`
	Inputs       []DomainInput       `xml:"input" json:"inputs"`
	TPMs         []DomainTPM         `xml:"tpm" json:"tpms"`
	Graphics     []DomainGraphic     `xml:"graphics" json:"graphics"`
	Sounds       []DomainSound       `xml:"sound" json:"sounds"`
	Audios       []DomainAudio       `xml:"audio" json:"audios"`
	Videos       []DomainVideo       `xml:"video" json:"videos"`
	Hostdevs     []DomainHostdev     `xml:"hostdev" json:"hostdevs"`
	RedirDevs    []DomainRedirDev    `xml:"redirdev" json:"redirDevs"`
	RedirFilters []DomainRedirFilter `xml:"redirfilter" json:"redirfilters"`
	Hubs         []DomainHub         `xml:"hub" json:"hubs"`
	Watchdog     *DomainWatchdog     `xml:"watchdog" json:"watchdog,omitempty"`
	MemBalloon   *DomainMemBalloon   `xml:"memballoon" json:"memballoon,omitempty"`
	RNGs         []DomainRNG         `xml:"rng" json:"rngs"`
	NVRAM        *DomainNVRAM        `xml:"nvram" json:"nvram"`
	Panics       []DomainPanic       `xml:"panic" json:"panics"`
	Shmems       []DomainShmem       `xml:"shmem" json:"shmems"`
	Memorydevs   []DomainMemorydev   `xml:"memory" json:"memorydevs"`
	IOMMU        *DomainIOMMU        `xml:"iommu" json:"iommu,omitempty"`
	VSock        *DomainVSock        `xml:"vsock" json:"vsock,omitempty"`
}

type DomainMemory struct {
	Value    int32  `xml:",chardata" json:"value"`
	Unit     string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	DumpCore string `xml:"dumpCore,attr,omitempty" json:"dumpCore,omitempty"`
}

type DomainCurrentMemory struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainMaxMemory struct {
	Value int32  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Slots int32  `xml:"slots,attr,omitempty" json:"slots,omitempty"`
}

type DomainMemoryHugepage struct {
	Size    int32  `xml:"size,attr" json:"size"`
	Unit    string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Nodeset string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty"`
}

type DomainMemoryHugepages struct {
	Hugepages []DomainMemoryHugepage `xml:"page" json:"hugepages"`
}

type DomainMemoryNosharepages struct {
}

type DomainMemoryLocked struct {
}

type DomainMemorySource struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainMemoryAccess struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainMemoryAllocation struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainMemoryDiscard struct {
}

type DomainMemoryBacking struct {
	MemoryHugePages    *DomainMemoryHugepages    `xml:"hugepages" json:"memoryHugePages,omitempty"`
	MemoryNosharepages *DomainMemoryNosharepages `xml:"nosharepages" json:"memoryNosharepages,omitempty"`
	MemoryLocked       *DomainMemoryLocked       `xml:"locked" json:"memoryLocked"`
	MemorySource       *DomainMemorySource       `xml:"source" json:"memorySource"`
	MemoryAccess       *DomainMemoryAccess       `xml:"access" json:"memoryAccess"`
	MemoryAllocation   *DomainMemoryAllocation   `xml:"allocation" json:"memoryAllocation"`
	MemoryDiscard      *DomainMemoryDiscard      `xml:"discard" json:"memoryDiscard"`
}

type DomainOSType struct {
	Arch    string `xml:"arch,attr,omitempty" json:"arch,omitempty"`
	Machine string `xml:"machine,attr,omitempty" json:"machine,omitempty"`
	Type    string `xml:",chardata" json:"type"`
}

type DomainSMBios struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type DomainNVRam struct {
	NVRam    string `xml:",chardata" json:"nvram"`
	Template string `xml:"template,attr,omitempty" json:"template,omitempty"`
}

type DomainBootDevice struct {
	Dev string `xml:"dev,attr" json:"dev"`
}

type DomainBootMenu struct {
	Enable  string `xml:"enable,attr,omitempty" json:"enable,omitempty"`
	Timeout string `xml:"timeout,attr,omitempty" json:"timeout,omitempty"`
}

type DomainSysInfoBIOS struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoSystem struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoBaseBoard struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoProcessor struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoMemory struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoChassis struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfoOEMStrings struct {
	Entry []string `xml:"entry" json:"entry"`
}

type DomainSysInfoSMBIOS struct {
	BIOS       *DomainSysInfoBIOS       `xml:"bios" json:"bios,omitempty"`
	System     *DomainSysInfoSystem     `xml:"system" json:"system,omitempty"`
	BaseBoard  []DomainSysInfoBaseBoard `xml:"baseBoard" json:"baseBoard"`
	Chassis    *DomainSysInfoChassis    `xml:"chassis" json:"chassis,omitempty"`
	Processor  []DomainSysInfoProcessor `xml:"processor" json:"processor"`
	Memory     []DomainSysInfoMemory    `xml:"memory" json:"memory"`
	OEMStrings *DomainSysInfoOEMStrings `xml:"oemStrings" json:"oemStrings,omitempty"`
}

type DomainSysInfoFWCfg struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry"`
}

type DomainSysInfo struct {
	SMBIOS *DomainSysInfoSMBIOS `xml:"-" json:"smbios,omitempty"`
	FWCfg  *DomainSysInfoFWCfg  `xml:"-" json:"fwcfg,omitempty"`
}

type DomainSysInfoEntry struct {
	Name  string `xml:"name,attr" json:"name"`
	File  string `xml:"file,attr,omitempty" json:"file,omitempty"`
	Value string `xml:",chardata" json:"value"`
}

type DomainBIOS struct {
	UseSerial     string `xml:"useserial,attr,omitempty" json:"useSerial,omitempty"`
	RebootTimeout *int32 `xml:"rebootTimeout,attr" json:"rebootTimeout"`
}

type DomainLoader struct {
	Path     string `xml:",chardata" json:"path"`
	Readonly string `xml:"readonly,attr,omitempty" json:"readonly,omitempty"`
	Secure   string `xml:"secure,attr,omitempty" json:"secure,omitempty"`
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty"`
}

type DomainACPI struct {
	Tables []DomainACPITable `xml:"table" json:"tables"`
}

type DomainACPITable struct {
	Type string `xml:"type,attr" json:"type"`
	Path string `xml:",chardata" json:"path"`
}

type DomainOSInitEnv struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:",chardata" json:"value"`
}

type DomainOSFirmwareInfo struct {
	Features []DomainOSFirmwareFeature `xml:"feature" json:"features"`
}

type DomainOSFirmwareFeature struct {
	Enabled string `xml:"enabled,attr,omitempty" json:"enabled"`
	Name    string `xml:"name,attr,omitempty" json:"name"`
}

type DomainOS struct {
	Type         *DomainOSType         `xml:"type" json:"type,omitempty"`
	Firmware     string                `xml:"firmware,attr,omitempty" json:"firmware,omitempty"`
	FirmwareInfo *DomainOSFirmwareInfo `xml:"firmware" json:"firmwareInfo,omitempty"`
	Init         string                `xml:"init,omitempty" json:"init,omitempty"`
	InitArgs     []string              `xml:"initarg" json:"initargs"`
	InitEnv      []DomainOSInitEnv     `xml:"initenv" json:"initenv"`
	InitDir      string                `xml:"initdir,omitempty" json:"initdir,omitempty"`
	InitUser     string                `xml:"inituser,omitempty" json:"inituser,omitempty"`
	InitGroup    string                `xml:"initgroup,omitempty" json:"initgroup,omitempty"`
	Loader       *DomainLoader         `xml:"loader" json:"loader,omitempty"`
	NVRam        *DomainNVRam          `xml:"nvram" json:"nvram,omitempty"`
	Kernel       string                `xml:"kernel,omitempty" json:"kernel,omitempty"`
	Initrd       string                `xml:"initrd,omitempty" json:"initrd,omitempty"`
	Cmdline      string                `xml:"cmdline,omitempty" json:"cmdline,omitempty"`
	DTB          string                `xml:"dtb,omitempty" json:"dtb,omitempty"`
	ACPI         *DomainACPI           `xml:"acpi" json:"acpi,omitempty"`
	BootDevices  []DomainBootDevice    `xml:"boot" json:"bootDevices"`
	BootMenu     *DomainBootMenu       `xml:"bootmenu" json:"bootMenu,omitempty"`
	BIOS         *DomainBIOS           `xml:"bios" json:"bios,omitempty"`
	SMBios       *DomainSMBios         `xml:"smbios" json:"smbios,omitempty"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty" json:"partition,omitempty"`
}

type DomainVCPU struct {
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty"`
	CPUSet    string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty"`
	Current   int32  `xml:"current,attr,omitempty" json:"current,omitempty"`
	Value     int32  `xml:",chardata" json:"value"`
}

type DomainVCPUsVCPU struct {
	Id           *int32 `xml:"id,attr" json:"id,omitempty"`
	Enabled      string `xml:"enabled,attr,omitempty" json:"enabled,omitempty"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty" json:"hotpluggable,omitempty"`
	Order        *int32 `xml:"order,attr" json:"order,omitempty"`
}

type DomainVCPUs struct {
	VCPU []DomainVCPUsVCPU `xml:"vcpu" json:"vcpu"`
}

type DomainCPUModel struct {
	Fallback string `xml:"fallback,attr,omitempty" json:"fallback,omitempty"`
	Value    string `xml:",chardata" json:"value"`
	VendorID string `xml:"vendor_id,attr,omitempty" json:"vendorId,omitempty"`
}

type DomainCPUTopology struct {
	Sockets int `xml:"sockets,attr,omitempty" json:"sockets,omitempty"`
	Dies    int `xml:"dies,attr,omitempty" json:"dies,omitempty"`
	Cores   int `xml:"cores,attr,omitempty" json:"cores,omitempty"`
	Threads int `xml:"threads,attr,omitempty" json:"threads,omitempty"`
}

type DomainCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty" json:"policy,omitempty"`
	Name   string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type DomainCPUCache struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty"`
	Mode  string `xml:"mode,attr" json:"mode,omitempty"`
}

type DomainCPU struct {
	XMLName    xml.Name           `xml:"cpu" json:"-"`
	Match      string             `xml:"match,attr,omitempty" json:"match,omitempty"`
	Mode       string             `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Check      string             `xml:"check,attr,omitempty" json:"check,omitempty"`
	Migratable string             `xml:"migratable,attr,omitempty" json:"migratable,omitempty"`
	Model      *DomainCPUModel    `xml:"model" json:"model,omitempty"`
	Vendor     string             `xml:"vendor,omitempty" json:"vendor"`
	Topology   *DomainCPUTopology `xml:"topology" json:"topology,omitempty"`
	Cache      *DomainCPUCache    `xml:"cache" json:"cache,omitempty"`
	Features   []DomainCPUFeature `xml:"feature" json:"features"`
	Numa       *DomainNuma        `xml:"numa" json:"numa,omitempty"`
}

type DomainNuma struct {
	Cell          []DomainCell             `xml:"cell" json:"cell"`
	Interconnects *DomainNUMAInterconnects `xml:"interconnects" json:"interconnects,omitempty"`
}

type DomainCell struct {
	ID        *int32               `xml:"id,attr" json:"id,omitempty"`
	CPUs      string               `xml:"cpus,attr,omitempty" json:"cpus"`
	Memory    int32                `xml:"memory,attr" json:"memory"`
	Unit      string               `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	MemAccess string               `xml:"memAccess,attr,omitempty" json:"memAccess,omitempty"`
	Discard   string               `xml:"discard,attr,omitempty" json:"discard,omitempty"`
	Distances *DomainCellDistances `xml:"distances" json:"distances,omitempty"`
	Caches    []DomainCellCache    `xml:"cache" json:"caches"`
}

type DomainCellDistances struct {
	Siblings []DomainCellSibling `xml:"sibling" json:"siblings"`
}

type DomainCellSibling struct {
	ID    int32 `xml:"id,attr" json:"id"`
	Value int32 `xml:"value,attr" json:"value"`
}

type DomainCellCache struct {
	Level         int32               `xml:"level,attr" json:"level"`
	Associativity string              `xml:"associativity,attr" json:"associativity"`
	Policy        string              `xml:"policy,attr" json:"policy"`
	Size          DomainCellCacheSize `xml:"size" json:"size"`
	Line          DomainCellCacheLine `xml:"line" json:"line"`
}

type DomainCellCacheSize struct {
	Value string `xml:"value,attr" json:"value"`
	Unit  string `xml:"unit,attr" json:"unit"`
}

type DomainCellCacheLine struct {
	Value string `xml:"value,attr" json:"value"`
	Unit  string `xml:"unit,attr" json:"unit"`
}

type DomainNUMAInterconnects struct {
	Latencies  []DomainNUMAInterconnectLatency   `xml:"latency" json:"latencies"`
	Bandwidths []DomainNUMAInterconnectBandwidth `xml:"bandwidth" json:"bandwidths"`
}

type DomainNUMAInterconnectLatency struct {
	Initiator int32  `xml:"initiator,attr" json:"initiator"`
	Target    int32  `xml:"target,attr" json:"target"`
	Cache     int32  `xml:"cache,attr,omitempty" json:"cache,omitempty"`
	Type      string `xml:"type,attr" json:"type"`
	Value     int32  `xml:"value,attr" json:"value"`
}

type DomainNUMAInterconnectBandwidth struct {
	Initiator int32  `xml:"initiator,attr" json:"initiator"`
	Target    int32  `xml:"target,attr" json:"target"`
	Type      string `xml:"type,attr" json:"type"`
	Value     int32  `xml:"value,attr" json:"value"`
	Unit      string `xml:"unit,attr" json:"unit"`
}

type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty" json:"offset,omitempty"`
	Basis      string        `xml:"basis,attr,omitempty" json:"basis,omitempty"`
	Adjustment string        `xml:"adjustment,attr,omitempty" json:"adjustment,omitempty"`
	TimeZone   string        `xml:"timezone,attr,omitempty" json:"time_zone,omitempty"`
	Timer      []DomainTimer `xml:"timer" json:"timer"`
}

type DomainTimer struct {
	Name       string              `xml:"name,attr" json:"name"`
	Track      string              `xml:"track,attr,omitempty" json:"track,omitempty"`
	TickPolicy string              `xml:"tickpolicy,attr,omitempty" json:"tickpolicy,omitempty"`
	CatchUp    *DomainTimerCatchUp `xml:"catchup" json:"catchup,omitempty"`
	Frequency  int64               `xml:"frequency,attr,omitempty" json:"frequency,omitempty"`
	Mode       string              `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Present    string              `xml:"present,attr,omitempty" json:"present,omitempty"`
}

type DomainTimerCatchUp struct {
	Threshold int32 `xml:"threshold,attr,omitempty" json:"threshold,omitempty"`
	Slew      int32 `xml:"slew,attr,omitempty" json:"slew,omitempty"`
	Limit     int32 `xml:"limit,attr,omitempty" json:"limit,omitempty"`
}

type DomainFeature struct {
}

type DomainFeatureState struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty"`
}

type DomainFeatureAPIC struct {
	EOI string `xml:"eoi,attr,omitempty" json:"eoi,omitempty"`
}

type DomainFeatureHyperVVendorId struct {
	DomainFeatureState `json:",inline"`
	Value              string `xml:"value,attr,omitempty" json:"value"`
}

type DomainFeatureHyperVSpinlocks struct {
	DomainFeatureState `json:",inline"`
	Retries            int32 `xml:"retries,attr,omitempty" json:"retries,omitempty"`
}

type DomainFeatureHyperVSTimer struct {
	DomainFeatureState `json:",inline"`
	Direct             *DomainFeatureState `xml:"direct" json:"direct,omitempty"`
}

type DomainFeatureHyperV struct {
	DomainFeature   `json:",inline"`
	Relaxed         *DomainFeatureState           `xml:"relaxed" json:"relaxed,omitempty"`
	VAPIC           *DomainFeatureState           `xml:"vapic" json:"vapic,omitempty"`
	Spinlocks       *DomainFeatureHyperVSpinlocks `xml:"spinlocks" json:"spinlocks,omitempty"`
	VPIndex         *DomainFeatureState           `xml:"vpindex" json:"vpindex,omitempty"`
	Runtime         *DomainFeatureState           `xml:"runtime" json:"runtime,omitempty"`
	Synic           *DomainFeatureState           `xml:"synic" json:"synic,omitempty"`
	STimer          *DomainFeatureHyperVSTimer    `xml:"stimer" json:"stimer,omitempty"`
	Reset_          *DomainFeatureState           `xml:"reset" json:"reset,omitempty"`
	VendorId        *DomainFeatureHyperVVendorId  `xml:"vendor_id" json:"vendorId,omitempty"`
	Frequencies     *DomainFeatureState           `xml:"frequencies" json:"frequencies,omitempty"`
	ReEnlightenment *DomainFeatureState           `xml:"reenlightenment" json:"reenlightenment,omitempty"`
	TLBFlush        *DomainFeatureState           `xml:"tlbflush" json:"tlb_flush,omitempty"`
	IPI             *DomainFeatureState           `xml:"ipi" json:"ipi,omitempty"`
	EVMCS           *DomainFeatureState           `xml:"evmcs" json:"evmcs,omitempty"`
}

type DomainFeatureKVM struct {
	Hidden        *DomainFeatureState `xml:"hidden" json:"hidden,omitempty"`
	HintDedicated *DomainFeatureState `xml:"hint-dedicated" json:"hintDedicated,omitempty"`
	PollControl   *DomainFeatureState `xml:"poll-control" json:"pollControl,omitempty"`
}

type DomainFeatureXenPassthrough struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty"`
	Mode  string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
}

type DomainFeatureXenE820Host struct {
	State string `xml:"state,attr" json:"state"`
}

type DomainFeatureXen struct {
	E820Host    *DomainFeatureXenE820Host    `xml:"e820_host" json:"e820Host,omitempty"`
	Passthrough *DomainFeatureXenPassthrough `xml:"passthrough" json:"passthrough,omitempty"`
}

type DomainFeatureGIC struct {
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
}

type DomainFeatureIOAPIC struct {
	Driver string `xml:"driver,attr,omitempty" json:"driver,omitempty"`
}

type DomainFeatureHPT struct {
	Resizing    string                    `xml:"resizing,attr,omitempty" json:"resizing,omitempty"`
	MaxPageSize *DomainFeatureHPTPageSize `xml:"maxpagesize" json:"maxpagesize,omitempty"`
}

type DomainFeatureHPTPageSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value string `xml:",chardata" json:"value"`
}

type DomainFeatureSMM struct {
	State string                `xml:"state,attr,omitempty" json:"state,omitempty"`
	TSeg  *DomainFeatureSMMTSeg `xml:"tseg" json:"tseg,omitempty"`
}

type DomainFeatureSMMTSeg struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value int32  `xml:",chardata" json:"value"`
}

type DomainFeatureCapability struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty"`
}

type DomainLaunchSecurity struct {
	SEV *DomainLaunchSecuritySEV `xml:"-" json:"sev,omitempty"`
}

type DomainLaunchSecuritySEV struct {
	CBitPos         *int32 `xml:"cbitpos" json:"cBitPos,omitempty"`
	ReducedPhysBits *int32 `xml:"reducedPhysBits" json:"reducedPhysBits,omitempty"`
	Policy          *int32 `xml:"policy" json:"policy,omitempty"`
	DHCert          string `xml:"dhCert" json:"dh_cert"`
	Session         string `xml:"session" json:"session"`
}

type DomainFeatureCapabilities struct {
	Policy         string                   `xml:"policy,attr,omitempty" json:"policy,omitempty"`
	AuditControl   *DomainFeatureCapability `xml:"audit_control" json:"auditControl,omitempty"`
	AuditWrite     *DomainFeatureCapability `xml:"audit_write" json:"auditWrite,omitempty"`
	BlockSuspend   *DomainFeatureCapability `xml:"block_suspend" json:"blockSuspend,omitempty"`
	Chown          *DomainFeatureCapability `xml:"chown" json:"chown,omitempty"`
	DACOverride    *DomainFeatureCapability `xml:"dac_override" json:"dacOverride,omitempty"`
	DACReadSearch  *DomainFeatureCapability `xml:"dac_read_Search" json:"dacReadSearch,omitempty"`
	FOwner         *DomainFeatureCapability `xml:"fowner" json:"fowner,omitempty"`
	FSetID         *DomainFeatureCapability `xml:"fsetid" json:"fsetid,omitempty"`
	IPCLock        *DomainFeatureCapability `xml:"ipc_lock" json:"ipcLock,omitempty"`
	IPCOwner       *DomainFeatureCapability `xml:"ipc_owner" json:"ipcOwner,omitempty"`
	Kill           *DomainFeatureCapability `xml:"kill" json:"kill,omitempty"`
	Lease          *DomainFeatureCapability `xml:"lease" json:"lease,omitempty"`
	LinuxImmutable *DomainFeatureCapability `xml:"linux_immutable" json:"linuxImmutable,omitempty"`
	MACAdmin       *DomainFeatureCapability `xml:"mac_admin" json:"macAdmin,omitempty"`
	MACOverride    *DomainFeatureCapability `xml:"mac_override" json:"macOverride,omitempty"`
	MkNod          *DomainFeatureCapability `xml:"mknod" json:"mknod,omitempty"`
	NetAdmin       *DomainFeatureCapability `xml:"net_admin" json:"netAdmin,omitempty"`
	NetBindService *DomainFeatureCapability `xml:"net_bind_service" json:"netBindService,omitempty"`
	NetBroadcast   *DomainFeatureCapability `xml:"net_broadcast" json:"netBroadcast,omitempty"`
	NetRaw         *DomainFeatureCapability `xml:"net_raw" json:"netraw,omitempty"`
	SetGID         *DomainFeatureCapability `xml:"setgid" json:"setgid,omitempty"`
	SetFCap        *DomainFeatureCapability `xml:"setfcap" json:"setfcap,omitempty"`
	SetPCap        *DomainFeatureCapability `xml:"setpcap" json:"setpcap,omitempty"`
	SetUID         *DomainFeatureCapability `xml:"setuid" json:"setuid,omitempty"`
	SysAdmin       *DomainFeatureCapability `xml:"sys_admin" json:"admin,omitempty"`
	SysBoot        *DomainFeatureCapability `xml:"sys_boot" json:"boot,omitempty"`
	SysChRoot      *DomainFeatureCapability `xml:"sys_chroot" json:"chroot,omitempty"`
	SysModule      *DomainFeatureCapability `xml:"sys_module" json:"module,omitempty"`
	SysNice        *DomainFeatureCapability `xml:"sys_nice" json:"nice,omitempty"`
	SysPAcct       *DomainFeatureCapability `xml:"sys_pacct" json:"pacct,omitempty"`
	SysPTrace      *DomainFeatureCapability `xml:"sys_ptrace" json:"ptrace,omitempty"`
	SysRawIO       *DomainFeatureCapability `xml:"sys_rawio" json:"rawio,omitempty"`
	SysResource    *DomainFeatureCapability `xml:"sys_resource" json:"resource,omitempty"`
	SysTime        *DomainFeatureCapability `xml:"sys_time" json:"time,omitempty"`
	SysTTYConfig   *DomainFeatureCapability `xml:"sys_tty_config" json:"ttyConfig,omitempty"`
	SysLog         *DomainFeatureCapability `xml:"syslog" json:"sysLog,omitempty"`
	WakeAlarm      *DomainFeatureCapability `xml:"wake_alarm" json:"wakeAlarm,omitempty"`
}

type DomainFeatureMSRS struct {
	Unknown string `xml:"unknown,attr" json:"unknown"`
}

type DomainFeatureCFPC struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainFeatureSBBC struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainFeatureIBS struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainFeatureList struct {
	PAE          *DomainFeature             `xml:"pae" json:"pae,omitempty"`
	ACPI         *DomainFeature             `xml:"acpi" json:"acpi,omitempty"`
	APIC         *DomainFeatureAPIC         `xml:"apic" json:"apic,omitempty"`
	HAP          *DomainFeatureState        `xml:"hap" json:"hap,omitempty"`
	Viridian     *DomainFeature             `xml:"viridian" json:"viridian,omitempty"`
	PrivNet      *DomainFeature             `xml:"privnet" json:"privnet,omitempty"`
	HyperV       *DomainFeatureHyperV       `xml:"hyperv" json:"hyperv,omitempty"`
	KVM          *DomainFeatureKVM          `xml:"kvm" json:"kvm,omitempty"`
	Xen          *DomainFeatureXen          `xml:"xen" json:"xen,omitempty"`
	PVSpinlock   *DomainFeatureState        `xml:"pvspinlock" json:"pvspinlock,omitempty"`
	PMU          *DomainFeatureState        `xml:"pmu" json:"pmu,omitempty"`
	VMPort       *DomainFeatureState        `xml:"vmport" json:"vmport,omitempty"`
	GIC          *DomainFeatureGIC          `xml:"gic" json:"gic,omitempty"`
	SMM          *DomainFeatureSMM          `xml:"smm" json:"smm,omitempty"`
	IOAPIC       *DomainFeatureIOAPIC       `xml:"ioapic" json:"ioapic,omitempty"`
	HPT          *DomainFeatureHPT          `xml:"hpt" json:"hpt,omitempty"`
	HTM          *DomainFeatureState        `xml:"htm" json:"htm,omitempty"`
	NestedHV     *DomainFeatureState        `xml:"nested-hv" json:"nestedHv,omitempty"`
	Capabilities *DomainFeatureCapabilities `xml:"capabilities" json:"capabilities,omitempty"`
	VMCoreInfo   *DomainFeatureState        `xml:"vmcoreinfo" json:"vmCoreInfo,omitempty"`
	MSRS         *DomainFeatureMSRS         `xml:"msrs" json:"msrs,omitempty"`
	CCFAssist    *DomainFeatureState        `xml:"ccf-assist" json:"ccfAssist,omitempty"`
	CFPC         *DomainFeatureCFPC         `xml:"cfpc" json:"cfpc,omitempty"`
	SBBC         *DomainFeatureSBBC         `xml:"sbbc" json:"sbbc,omitempty"`
	IBS          *DomainFeatureIBS          `xml:"ibs" json:"ibs,omitempty"`
}

type DomainCPUTuneShares struct {
	Value int32 `xml:",chardata" json:"value"`
}

type DomainCPUTunePeriod struct {
	Value int64 `xml:",chardata" json:"value"`
}

type DomainCPUTuneQuota struct {
	Value int64 `xml:",chardata" json:"value"`
}

type DomainCPUTuneVCPUPin struct {
	VCPU   int32  `xml:"vcpu,attr" json:"vcpu"`
	CPUSet string `xml:"cpuset,attr" json:"cpu_set"`
}

type DomainCPUTuneEmulatorPin struct {
	CPUSet string `xml:"cpuset,attr" json:"cpuset"`
}

type DomainCPUTuneIOThreadPin struct {
	IOThread int32  `xml:"iothread,attr" json:"iothread"`
	CPUSet   string `xml:"cpuset,attr" json:"cpuset"`
}

type DomainCPUTuneVCPUSched struct {
	VCPUs     string `xml:"vcpus,attr" json:"vcpus"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty"`
	Priority  *int32 `xml:"priority,attr" json:"priority,omitempty"`
}

type DomainCPUTuneIOThreadSched struct {
	IOThreads string `xml:"iothreads,attr" json:"iothreads"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty"`
	Priority  *int32 `xml:"priority,attr" json:"priority"`
}

type DomainCPUTuneEmulatorSched struct {
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty"`
	Priority  *int32 `xml:"priority,attr" json:"priority,omitempty"`
}

type DomainCPUCacheTune struct {
	VCPUs   string                      `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`
	Cache   []DomainCPUCacheTuneCache   `xml:"cache" json:"cache"`
	Monitor []DomainCPUCacheTuneMonitor `xml:"monitor" json:"monitor"`
}

type DomainCPUCacheTuneCache struct {
	ID    int32  `xml:"id,attr" json:"id"`
	Level int32  `xml:"level,attr" json:"level"`
	Type  string `xml:"type,attr" json:"type"`
	Size  int32  `xml:"size,attr" json:"size"`
	Unit  string `xml:"unit,attr" json:"unit"`
}

type DomainCPUCacheTuneMonitor struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty"`
	VCPUs string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`
}

type DomainCPUMemoryTune struct {
	VCPUs   string                       `xml:"vcpus,attr" json:"vcp_us"`
	Nodes   []DomainCPUMemoryTuneNode    `xml:"node" json:"nodes"`
	Monitor []DomainCPUMemoryTuneMonitor `xml:"monitor" json:"monitor"`
}

type DomainCPUMemoryTuneNode struct {
	ID        int32 `xml:"id,attr" json:"id"`
	Bandwidth int32 `xml:"bandwidth,attr" json:"bandwidth"`
}

type DomainCPUMemoryTuneMonitor struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty"`
	VCPUs string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty"`
}

type DomainCPUTune struct {
	Shares         *DomainCPUTuneShares         `xml:"shares" json:"shares,omitempty"`
	Period         *DomainCPUTunePeriod         `xml:"period" json:"period,omitempty"`
	Quota          *DomainCPUTuneQuota          `xml:"quota" json:"quota,omitempty"`
	GlobalPeriod   *DomainCPUTunePeriod         `xml:"global_period" json:"globalPeriod,omitempty"`
	GlobalQuota    *DomainCPUTuneQuota          `xml:"global_quota" json:"globalQuota,omitempty"`
	EmulatorPeriod *DomainCPUTunePeriod         `xml:"emulator_period" json:"emulatorPeriod,omitempty"`
	EmulatorQuota  *DomainCPUTuneQuota          `xml:"emulator_quota" json:"emulatorQuota,omitempty"`
	IOThreadPeriod *DomainCPUTunePeriod         `xml:"iothread_period" json:"iothreadPeriod,omitempty"`
	IOThreadQuota  *DomainCPUTuneQuota          `xml:"iothread_quota" json:"iothreadQuota,omitempty"`
	VCPUPin        []DomainCPUTuneVCPUPin       `xml:"vcpupin" json:"vcpupin,omitempty"`
	EmulatorPin    *DomainCPUTuneEmulatorPin    `xml:"emulatorpin" json:"emulatorpin,omitempty"`
	IOThreadPin    []DomainCPUTuneIOThreadPin   `xml:"iothreadpin" json:"iothreadpin,omitempty"`
	VCPUSched      []DomainCPUTuneVCPUSched     `xml:"vcpusched" json:"vcpusched,omitempty"`
	EmulatorSched  *DomainCPUTuneEmulatorSched  `xml:"emulatorsched" json:"emulatorsched,omitempty"`
	IOThreadSched  []DomainCPUTuneIOThreadSched `xml:"iothreadsched" json:"iothreadsched,omitempty"`
	CacheTune      []DomainCPUCacheTune         `xml:"cachetune" json:"cachetune,omitempty"`
	MemoryTune     []DomainCPUMemoryTune        `xml:"memorytune" json:"memorytune,omitempty"`
}

type DomainQEMUCommandlineArg struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainQEMUCommandlineEnv struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type DomainQEMUCommandline struct {
	XMLName xml.Name                   `xml:"http://libvirt.org/schemas/domain/qemu/1.0 commandline" json:"-"`
	Args    []DomainQEMUCommandlineArg `xml:"arg" json:"args"`
	Envs    []DomainQEMUCommandlineEnv `xml:"env" json:"envs"`
}

type DomainQEMUCapabilitiesEntry struct {
	Name string `xml:"capability,attr" json:"name"`
}

type DomainQEMUCapabilities struct {
	XMLName xml.Name                      `xml:"http://libvirt.org/schemas/domain/qemu/1.0 capabilities" json:"-"`
	Add     []DomainQEMUCapabilitiesEntry `xml:"add" json:"add"`
	Del     []DomainQEMUCapabilitiesEntry `xml:"del" json:"del"`
}

type DomainQEMUDeprecation struct {
	XMLName  xml.Name `xml:"http://libvirt.org/schemas/domain/qemu/1.0 deprecation" json:"-"`
	Behavior string   `xml:"behavior,attr,omitempty" json:"behavior,omitempty"`
}

type DomainLXCNamespace struct {
	XMLName  xml.Name               `xml:"http://libvirt.org/schemas/domain/lxc/1.0 namespace" json:"-"`
	ShareNet *DomainLXCNamespaceMap `xml:"sharenet" json:"sharenet,omitempty"`
	ShareIPC *DomainLXCNamespaceMap `xml:"shareipc" json:"shareipc,omitempty"`
	ShareUTS *DomainLXCNamespaceMap `xml:"shareuts" json:"shareuts,omitempty"`
}

type DomainLXCNamespaceMap struct {
	Type  string `xml:"type,attr" json:"type"`
	Value string `xml:"value,attr" json:"value"`
}

type DomainBHyveCommandlineArg struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainBHyveCommandlineEnv struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

type DomainBHyveCommandline struct {
	XMLName xml.Name                    `xml:"http://libvirt.org/schemas/domain/bhyve/1.0 commandline" json:"-"`
	Args    []DomainBHyveCommandlineArg `xml:"arg" json:"args"`
	Envs    []DomainBHyveCommandlineEnv `xml:"env" json:"envs"`
}

type DomainXenCommandlineArg struct {
	Value string `xml:"value,attr" json:"value"`
}

type DomainXenCommandline struct {
	XMLName xml.Name                  `xml:"http://libvirt.org/schemas/domain/xen/1.0 commandline" json:"-"`
	Args    []DomainXenCommandlineArg `xml:"arg" json:"args"`
}

type DomainBlockIOTune struct {
	Weight int32                     `xml:"weight,omitempty" json:"weight"`
	Device []DomainBlockIOTuneDevice `xml:"device" json:"device"`
}

type DomainBlockIOTuneDevice struct {
	Path          string `xml:"path" json:"path"`
	Weight        int32  `xml:"weight,omitempty" json:"weight,omitempty"`
	ReadIopsSec   int32  `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty"`
	WriteIopsSec  int32  `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty"`
	ReadBytesSec  int32  `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty"`
	WriteBytesSec int32  `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty"`
}

type DomainPM struct {
	SuspendToMem  *DomainPMPolicy `xml:"suspend-to-mem" json:"suspendToMem,omitempty"`
	SuspendToDisk *DomainPMPolicy `xml:"suspend-to-disk" json:"suspendToDisk,omitempty"`
}

type DomainPMPolicy struct {
	Enabled string `xml:"enabled,attr" json:"enabled"`
}

type DomainSecLabel struct {
	Type       string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Model      string `xml:"model,attr,omitempty" json:"model,omitempty"`
	Relabel    string `xml:"relabel,attr,omitempty" json:"relabel,omitempty"`
	Label      string `xml:"label,omitempty" json:"label,omitempty"`
	ImageLabel string `xml:"imagelabel,omitempty" json:"imageLabel,omitempty"`
	BaseLabel  string `xml:"baselabel,omitempty" json:"baseLabel,omitempty"`
}

type DomainDeviceSecLabel struct {
	Model     string `xml:"model,attr,omitempty" json:"model,omitempty"`
	LabelSkip string `xml:"labelskip,attr,omitempty" json:"labelSkip,omitempty"`
	Relabel   string `xml:"relabel,attr,omitempty" json:"relabel,omitempty"`
	Label     string `xml:"label,omitempty" json:"label,omitempty"`
}

type DomainNUMATune struct {
	Memory   *DomainNUMATuneMemory   `xml:"memory" json:"memory"`
	MemNodes []DomainNUMATuneMemNode `xml:"memnode" json:"memNodes"`
}

type DomainNUMATuneMemory struct {
	Mode      string `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Nodeset   string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty"`
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty"`
}

type DomainNUMATuneMemNode struct {
	CellID  int32  `xml:"cellid,attr" json:"cellId,omitempty"`
	Mode    string `xml:"mode,attr" json:"mode,omitempty"`
	Nodeset string `xml:"nodeset,attr" json:"nodeset,omitempty"`
}

type DomainIOThreadIDs struct {
	IOThreads []DomainIOThread `xml:"iothread" json:"iothreads"`
}

type DomainIOThread struct {
	ID int32 `xml:"id,attr" json:"id"`
}

type DomainKeyWrap struct {
	Ciphers []DomainKeyWrapCipher `xml:"cipher" json:"ciphers"`
}

type DomainKeyWrapCipher struct {
	Name  string `xml:"name,attr" json:"name"`
	State string `xml:"state,attr" json:"state"`
}

type DomainIDMap struct {
	UIDs []DomainIDMapRange `xml:"uid" json:"uid"`
	GIDs []DomainIDMapRange `xml:"gid" json:"gid"`
}

type DomainIDMapRange struct {
	Start  int32 `xml:"start,attr" json:"start"`
	Target int32 `xml:"target,attr" json:"target"`
	Count  int32 `xml:"count,attr" json:"count"`
}

type DomainMemoryTuneLimit struct {
	Value int64  `xml:",chardata" json:"value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type DomainMemoryTune struct {
	HardLimit     *DomainMemoryTuneLimit `xml:"hard_limit" json:"hardLimit,omitempty"`
	SoftLimit     *DomainMemoryTuneLimit `xml:"soft_limit" json:"softLimit,omitempty"`
	MinGuarantee  *DomainMemoryTuneLimit `xml:"min_guarantee" json:"minGuarantee,omitempty"`
	SwapHardLimit *DomainMemoryTuneLimit `xml:"swap_hard_limit" json:"swapHardLimit,omitempty"`
}

type DomainMetadata struct {
	XML string `xml:",innerxml" json:",inline"`
}

type DomainVMWareDataCenterPath struct {
	XMLName xml.Name `xml:"http://libvirt.org/schemas/domain/vmware/1.0 datacenterpath" json:"-"`
	Value   string   `xml:",chardata" json:"value"`
}

type DomainPerf struct {
	Events []DomainPerfEvent `xml:"event" json:"events"`
}

type DomainPerfEvent struct {
	Name    string `xml:"name,attr" json:"name"`
	Enabled string `xml:"enabled,attr" json:"enabled"`
}

type DomainGenID struct {
	Value string `xml:",chardata" json:"value"`
}

// Domain NB, try to keep the order of fields in this struct
// matching the order of XML elements that libvirt
// will generate when dumping XML.
// +gogo:deepcopy-gen=true

type Domain struct {
	XMLName        xml.Name              `xml:"domain" json:"-"`
	Type           string                `xml:"type,attr,omitempty" json:"type,omitempty"`
	ID             *int32                `xml:"id,attr" json:"id"`
	Name           string                `xml:"name,omitempty" json:"name,omitempty"`
	UUID           string                `xml:"uuid,omitempty" json:"uuid,omitempty"`
	GenID          *DomainGenID          `xml:"genid" json:"genId,omitempty"`
	Title          string                `xml:"title,omitempty" json:"title,omitempty"`
	Description    string                `xml:"description,omitempty" json:"description,omitempty"`
	Metadata       *DomainMetadata       `xml:"metadata" json:"metadata,omitempty"`
	MaximumMemory  *DomainMaxMemory      `xml:"maxMemory" json:"maxMemory,omitempty"`
	Memory         *DomainMemory         `xml:"memory" json:"memory,omitempty"`
	CurrentMemory  *DomainCurrentMemory  `xml:"currentMemory" json:"currentMemory,omitempty"`
	BlockIOTune    *DomainBlockIOTune    `xml:"blkiotune" json:"blockIoTune,omitempty"`
	MemoryTune     *DomainMemoryTune     `xml:"memtune" json:"memTune,omitempty"`
	MemoryBacking  *DomainMemoryBacking  `xml:"memoryBacking" json:"memoryBacking,omitempty"`
	VCPU           *DomainVCPU           `xml:"vcpu" json:"vcpu,omitempty"`
	VCPUs          *DomainVCPUs          `xml:"vcpus" json:"vcpus,omitempty"`
	IOThreads      int32                 `xml:"iothreads,omitempty" json:"iothreads,omitempty"`
	IOThreadIDs    *DomainIOThreadIDs    `xml:"iothreadids" json:"iothreadids,omitempty"`
	CPUTune        *DomainCPUTune        `xml:"cputune" json:"cputune,omitempty"`
	NUMATune       *DomainNUMATune       `xml:"numatune" json:"numatune,omitempty"`
	Resource       *DomainResource       `xml:"resource" json:"resource,omitempty"`
	SysInfo        []DomainSysInfo       `xml:"sysinfo" json:"sysinfo"`
	Bootloader     string                `xml:"bootloader,omitempty" json:"bootloader,omitempty"`
	BootloaderArgs string                `xml:"bootloader_args,omitempty" json:"bootloaderArgs,omitempty"`
	OS             *DomainOS             `xml:"os" json:"os,omitempty"`
	IDMap          *DomainIDMap          `xml:"idmap" json:"idMap,omitempty"`
	Features       *DomainFeatureList    `xml:"features" json:"features,omitempty"`
	CPU            *DomainCPU            `xml:"cpu" json:"cpu,omitempty"`
	Clock          *DomainClock          `xml:"clock" json:"clock,omitempty"`
	OnPoweroff     string                `xml:"on_poweroff,omitempty" json:"onPoweroff,omitempty"`
	OnReboot       string                `xml:"on_reboot,omitempty" json:"onReboot,omitempty"`
	OnCrash        string                `xml:"on_crash,omitempty" json:"onCrash,omitempty"`
	PM             *DomainPM             `xml:"pm" json:"pm,omitempty"`
	Perf           *DomainPerf           `xml:"perf" json:"perf,omitempty"`
	Devices        *DomainDeviceList     `xml:"devices" json:"devices,omitempty"`
	SecLabel       []DomainSecLabel      `xml:"seclabel" json:"seclabel"`
	KeyWrap        *DomainKeyWrap        `xml:"keywrap" json:"keywrap,omitempty"`
	LaunchSecurity *DomainLaunchSecurity `xml:"launchSecurity" json:"launchSecurity,omitempty"`

	/* Hypervisor namespaces must all be last */
	QEMUCommandline      *DomainQEMUCommandline      `json:"qemuCommandline,omitempty"`
	QEMUCapabilities     *DomainQEMUCapabilities     `json:"qemuCapabilities,omitempty"`
	QEMUDeprecation      *DomainQEMUDeprecation      `json:"qemuDeprecation,omitempty"`
	LXCNamespace         *DomainLXCNamespace         `json:"lxcNamespace,omitempty"`
	BHyveCommandline     *DomainBHyveCommandline     `json:"bHyveCommandline,omitempty"`
	VMWareDataCenterPath *DomainVMWareDataCenterPath `json:"vmWareDataCenterPath,omitempty"`
	XenCommandline       *DomainXenCommandline       `json:"xenCommandline,omitempty"`
}

func (d *Domain) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *Domain) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainController DomainController

type domainControllerPCI struct {
	DomainControllerPCI
	domainController
}

type domainControllerUSB struct {
	DomainControllerUSB
	domainController
}

type domainControllerVirtIOSerial struct {
	DomainControllerVirtIOSerial
	domainController
}

type domainControllerXenBus struct {
	DomainControllerXenBus
	domainController
}

func (a *DomainControllerPCITarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "chassisNr", a.ChassisNr, "%d")
	marshalUintAttr(&start, "chassis", a.Chassis, "%d")
	marshalUintAttr(&start, "port", a.Port, "%d")
	marshalUintAttr(&start, "busNr", a.BusNr, "%d")
	marshalUintAttr(&start, "index", a.Index, "%d")
	if a.Hotplug != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "hotplug"}, a.Hotplug,
		})
	}
	e.EncodeToken(start)
	if a.NUMANode != nil {
		node := xml.StartElement{
			Name: xml.Name{Local: "node"},
		}
		e.EncodeToken(node)
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.NUMANode)))
		e.EncodeToken(node.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainControllerPCITarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "chassisNr" {
			if err := unmarshalIntAttr(attr.Value, &a.ChassisNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "chassis" {
			if err := unmarshalIntAttr(attr.Value, &a.Chassis, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshalIntAttr(attr.Value, &a.Port, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "busNr" {
			if err := unmarshalIntAttr(attr.Value, &a.BusNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "index" {
			if err := unmarshalIntAttr(attr.Value, &a.Index, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "hotplug" {
			a.Hotplug = attr.Value
		}
	}
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "node" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					val, err := strconv.ParseUint(string(data), 10, 64)
					if err != nil {
						return err
					}
					vali := int32(val)
					a.NUMANode = &vali
				}
			}
		}
	}
	return nil
}

func (a *DomainController) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "controller"
	if a.Type == "pci" {
		pci := domainControllerPCI{}
		pci.domainController = domainController(*a)
		if a.PCI != nil {
			pci.DomainControllerPCI = *a.PCI
		}
		return e.EncodeElement(pci, start)
	} else if a.Type == "usb" {
		usb := domainControllerUSB{}
		usb.domainController = domainController(*a)
		if a.USB != nil {
			usb.DomainControllerUSB = *a.USB
		}
		return e.EncodeElement(usb, start)
	} else if a.Type == "virtio-serial" {
		vioserial := domainControllerVirtIOSerial{}
		vioserial.domainController = domainController(*a)
		if a.VirtIOSerial != nil {
			vioserial.DomainControllerVirtIOSerial = *a.VirtIOSerial
		}
		return e.EncodeElement(vioserial, start)
	} else if a.Type == "xenbus" {
		xenbus := domainControllerXenBus{}
		xenbus.domainController = domainController(*a)
		if a.XenBus != nil {
			xenbus.DomainControllerXenBus = *a.XenBus
		}
		return e.EncodeElement(xenbus, start)
	} else {
		gen := domainController(*a)
		return e.EncodeElement(gen, start)
	}
}

func getAttr(attrs []xml.Attr, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func (a *DomainController) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if typ == "pci" {
		var pci domainControllerPCI
		err := d.DecodeElement(&pci, &start)
		if err != nil {
			return err
		}
		*a = DomainController(pci.domainController)
		a.PCI = &pci.DomainControllerPCI
		return nil
	} else if typ == "usb" {
		var usb domainControllerUSB
		err := d.DecodeElement(&usb, &start)
		if err != nil {
			return err
		}
		*a = DomainController(usb.domainController)
		a.USB = &usb.DomainControllerUSB
		return nil
	} else if typ == "virtio-serial" {
		var vioserial domainControllerVirtIOSerial
		err := d.DecodeElement(&vioserial, &start)
		if err != nil {
			return err
		}
		*a = DomainController(vioserial.domainController)
		a.VirtIOSerial = &vioserial.DomainControllerVirtIOSerial
		return nil
	} else if typ == "xenbus" {
		var xenbus domainControllerXenBus
		err := d.DecodeElement(&xenbus, &start)
		if err != nil {
			return err
		}
		*a = DomainController(xenbus.domainController)
		a.XenBus = &xenbus.DomainControllerXenBus
		return nil
	} else {
		var gen domainController
		err := d.DecodeElement(&gen, &start)
		if err != nil {
			return err
		}
		*a = DomainController(gen)
		return nil
	}
}

func (d *DomainGraphic) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainGraphic) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainController) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainController) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainDiskReservationsSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "source"
	src := DomainChardevSource(*a)
	typ := getChardevSourceType(&src)
	if typ != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, typ,
		})
	}
	return e.EncodeElement(&src, start)
}

func (a *DomainDiskReservationsSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "unix"
	}
	src := createChardevSource(typ)
	err := d.DecodeElement(&src, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskReservationsSource(*src)
	return nil
}

func (a *DomainDiskSourceVHostUser) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "source"
	src := DomainChardevSource(*a)
	typ := getChardevSourceType(&src)
	if typ != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, typ,
		})
	}
	return e.EncodeElement(&src, start)
}

func (a *DomainDiskSourceVHostUser) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "unix"
	}
	src := createChardevSource(typ)
	err := d.DecodeElement(&src, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskSourceVHostUser(*src)
	return nil
}

type domainDiskSource DomainDiskSource

type domainDiskSourceFile struct {
	DomainDiskSourceFile
	domainDiskSource
}

type domainDiskSourceBlock struct {
	DomainDiskSourceBlock
	domainDiskSource
}

type domainDiskSourceDir struct {
	DomainDiskSourceDir
	domainDiskSource
}

type domainDiskSourceNetwork struct {
	DomainDiskSourceNetwork
	domainDiskSource
}

type domainDiskSourceVolume struct {
	DomainDiskSourceVolume
	domainDiskSource
}

type domainDiskSourceNVMEPCI struct {
	DomainDiskSourceNVMEPCI
	domainDiskSource
}

type domainDiskSourceVHostUser struct {
	DomainDiskSourceVHostUser
	domainDiskSource
}

func (a *DomainDiskSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.File != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.File.File == "" {
			return nil
		}
		file := domainDiskSourceFile{
			*a.File, domainDiskSource(*a),
		}
		return e.EncodeElement(&file, start)
	} else if a.Block != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.Block.Dev == "" {
			return nil
		}
		block := domainDiskSourceBlock{
			*a.Block, domainDiskSource(*a),
		}
		return e.EncodeElement(&block, start)
	} else if a.Dir != nil {
		dir := domainDiskSourceDir{
			*a.Dir, domainDiskSource(*a),
		}
		return e.EncodeElement(&dir, start)
	} else if a.Network != nil {
		network := domainDiskSourceNetwork{
			*a.Network, domainDiskSource(*a),
		}
		return e.EncodeElement(&network, start)
	} else if a.Volume != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.Volume.Pool == "" && a.Volume.Volume == "" {
			return nil
		}
		volume := domainDiskSourceVolume{
			*a.Volume, domainDiskSource(*a),
		}
		return e.EncodeElement(&volume, start)
	} else if a.NVME != nil {
		if a.NVME.PCI != nil {
			nvme := domainDiskSourceNVMEPCI{
				*a.NVME.PCI, domainDiskSource(*a),
			}
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "pci",
			})
			return e.EncodeElement(&nvme, start)
		}
	} else if a.VHostUser != nil {
		vhost := domainDiskSourceVHostUser{
			*a.VHostUser, domainDiskSource(*a),
		}
		return e.EncodeElement(&vhost, start)
	}
	return nil
}

func (a *DomainDiskSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.File != nil {
		file := domainDiskSourceFile{
			*a.File, domainDiskSource(*a),
		}
		err := d.DecodeElement(&file, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(file.domainDiskSource)
		a.File = &file.DomainDiskSourceFile
	} else if a.Block != nil {
		block := domainDiskSourceBlock{
			*a.Block, domainDiskSource(*a),
		}
		err := d.DecodeElement(&block, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(block.domainDiskSource)
		a.Block = &block.DomainDiskSourceBlock
	} else if a.Dir != nil {
		dir := domainDiskSourceDir{
			*a.Dir, domainDiskSource(*a),
		}
		err := d.DecodeElement(&dir, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(dir.domainDiskSource)
		a.Dir = &dir.DomainDiskSourceDir
	} else if a.Network != nil {
		network := domainDiskSourceNetwork{
			*a.Network, domainDiskSource(*a),
		}
		err := d.DecodeElement(&network, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(network.domainDiskSource)
		a.Network = &network.DomainDiskSourceNetwork
	} else if a.Volume != nil {
		volume := domainDiskSourceVolume{
			*a.Volume, domainDiskSource(*a),
		}
		err := d.DecodeElement(&volume, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(volume.domainDiskSource)
		a.Volume = &volume.DomainDiskSourceVolume
	} else if a.NVME != nil {
		typ, ok := getAttr(start.Attr, "type")
		if !ok {
			return fmt.Errorf("Missing nvme source type")
		}
		if typ == "pci" {
			a.NVME.PCI = &DomainDiskSourceNVMEPCI{}
			nvme := domainDiskSourceNVMEPCI{
				*a.NVME.PCI, domainDiskSource(*a),
			}
			err := d.DecodeElement(&nvme, &start)
			if err != nil {
				return err
			}
			*a = DomainDiskSource(nvme.domainDiskSource)
			a.NVME.PCI = &nvme.DomainDiskSourceNVMEPCI
		}
	} else if a.VHostUser != nil {
		vhost := domainDiskSourceVHostUser{
			*a.VHostUser, domainDiskSource(*a),
		}
		err := d.DecodeElement(&vhost, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(vhost.domainDiskSource)
		a.VHostUser = &vhost.DomainDiskSourceVHostUser
	}
	return nil
}

type domainDiskBackingStore DomainDiskBackingStore

func (a *DomainDiskBackingStore) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backingStore"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		} else if a.Source.VHostUser != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vhostuser",
			})
		}
	}
	disk := domainDiskBackingStore(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDiskBackingStore) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	} else if typ == "vhostuser" {
		a.Source.VHostUser = &DomainDiskSourceVHostUser{}
	}
	disk := domainDiskBackingStore(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskBackingStore(disk)
	if !ok && a.Source.File.File == "" {
		a.Source.File = nil
	}
	return nil
}

type domainDiskMirror DomainDiskMirror

func (a *DomainDiskMirror) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "mirror"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
			if a.Source.File.File != "" {
				start.Attr = append(start.Attr, xml.Attr{
					xml.Name{Local: "file"}, a.Source.File.File,
				})
			}
			if a.Format != nil && a.Format.Type != "" {
				start.Attr = append(start.Attr, xml.Attr{
					xml.Name{Local: "format"}, a.Format.Type,
				})
			}
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		} else if a.Source.VHostUser != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vhostuser",
			})
		}
	}
	disk := domainDiskMirror(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDiskMirror) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	} else if typ == "vhostuser" {
		a.Source.VHostUser = &DomainDiskSourceVHostUser{}
	}
	disk := domainDiskMirror(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskMirror(disk)
	if !ok {
		if a.Source.File.File == "" {
			file, ok := getAttr(start.Attr, "file")
			if ok {
				a.Source.File.File = file
			} else {
				a.Source.File = nil
			}
		}
		if a.Format == nil {
			format, ok := getAttr(start.Attr, "format")
			if ok {
				a.Format = &DomainDiskFormat{
					Type: format,
				}
			}
		}
	}
	return nil
}

type domainDisk DomainDisk

func (a *DomainDisk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "disk"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		} else if a.Source.NVME != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "nvme",
			})
		} else if a.Source.VHostUser != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vhostuser",
			})
		}
	}
	disk := domainDisk(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDisk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	} else if typ == "nvme" {
		a.Source.NVME = &DomainDiskSourceNVME{}
	} else if typ == "vhostuser" {
		a.Source.VHostUser = &DomainDiskSourceVHostUser{}
	}
	disk := domainDisk(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDisk(disk)
	return nil
}

func (d *DomainDisk) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainDisk) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainFilesystemSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Mount != nil {
		return e.EncodeElement(a.Mount, start)
	} else if a.Block != nil {
		return e.EncodeElement(a.Block, start)
	} else if a.File != nil {
		return e.EncodeElement(a.File, start)
	} else if a.Template != nil {
		return e.EncodeElement(a.Template, start)
	} else if a.RAM != nil {
		return e.EncodeElement(a.RAM, start)
	} else if a.Bind != nil {
		return e.EncodeElement(a.Bind, start)
	} else if a.Volume != nil {
		return e.EncodeElement(a.Volume, start)
	}
	return nil
}

func (a *DomainFilesystemSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Mount != nil {
		return d.DecodeElement(a.Mount, &start)
	} else if a.Block != nil {
		return d.DecodeElement(a.Block, &start)
	} else if a.File != nil {
		return d.DecodeElement(a.File, &start)
	} else if a.Template != nil {
		return d.DecodeElement(a.Template, &start)
	} else if a.RAM != nil {
		return d.DecodeElement(a.RAM, &start)
	} else if a.Bind != nil {
		return d.DecodeElement(a.Bind, &start)
	} else if a.Volume != nil {
		return d.DecodeElement(a.Volume, &start)
	}
	return nil
}

type domainFilesystem DomainFilesystem

func (a *DomainFilesystem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "filesystem"
	if a.Source != nil {
		if a.Source.Mount != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "mount",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Template != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "template",
			})
		} else if a.Source.RAM != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "ram",
			})
		} else if a.Source.Bind != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "bind",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	fs := domainFilesystem(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainFilesystem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "mount"
	}
	a.Source = &DomainFilesystemSource{}
	if typ == "mount" {
		a.Source.Mount = &DomainFilesystemSourceMount{}
	} else if typ == "block" {
		a.Source.Block = &DomainFilesystemSourceBlock{}
	} else if typ == "file" {
		a.Source.File = &DomainFilesystemSourceFile{}
	} else if typ == "template" {
		a.Source.Template = &DomainFilesystemSourceTemplate{}
	} else if typ == "ram" {
		a.Source.RAM = &DomainFilesystemSourceRAM{}
	} else if typ == "bind" {
		a.Source.Bind = &DomainFilesystemSourceBind{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainFilesystemSourceVolume{}
	}
	fs := domainFilesystem(*a)
	err := d.DecodeElement(&fs, &start)
	if err != nil {
		return err
	}
	*a = DomainFilesystem(fs)
	return nil
}

func (d *DomainFilesystem) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainFilesystem) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainInterfaceVirtualPortParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "parameters"
	if a.Any != nil {
		return e.EncodeElement(a.Any, start)
	} else if a.VEPA8021QBG != nil {
		return e.EncodeElement(a.VEPA8021QBG, start)
	} else if a.VNTag8011QBH != nil {
		return e.EncodeElement(a.VNTag8011QBH, start)
	} else if a.OpenVSwitch != nil {
		return e.EncodeElement(a.OpenVSwitch, start)
	} else if a.MidoNet != nil {
		return e.EncodeElement(a.MidoNet, start)
	}
	return nil
}

func (a *DomainInterfaceVirtualPortParams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Any != nil {
		return d.DecodeElement(a.Any, &start)
	} else if a.VEPA8021QBG != nil {
		return d.DecodeElement(a.VEPA8021QBG, &start)
	} else if a.VNTag8011QBH != nil {
		return d.DecodeElement(a.VNTag8011QBH, &start)
	} else if a.OpenVSwitch != nil {
		return d.DecodeElement(a.OpenVSwitch, &start)
	} else if a.MidoNet != nil {
		return d.DecodeElement(a.MidoNet, &start)
	}
	return nil
}

type domainInterfaceVirtualPort DomainInterfaceVirtualPort

func (a *DomainInterfaceVirtualPort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "virtualport"
	if a.Params != nil {
		if a.Params.Any != nil {
			/* no type attr wanted */
		} else if a.Params.VEPA8021QBG != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbg",
			})
		} else if a.Params.VNTag8011QBH != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbh",
			})
		} else if a.Params.OpenVSwitch != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "openvswitch",
			})
		} else if a.Params.MidoNet != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "midonet",
			})
		}
	}
	vp := domainInterfaceVirtualPort(*a)
	return e.EncodeElement(&vp, start)
}

func (a *DomainInterfaceVirtualPort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	a.Params = &DomainInterfaceVirtualPortParams{}
	if !ok {
		var any DomainInterfaceVirtualPortParamsAny
		a.Params.Any = &any
	} else if typ == "802.1Qbg" {
		var vepa DomainInterfaceVirtualPortParamsVEPA8021QBG
		a.Params.VEPA8021QBG = &vepa
	} else if typ == "802.1Qbh" {
		var vntag DomainInterfaceVirtualPortParamsVNTag8021QBH
		a.Params.VNTag8011QBH = &vntag
	} else if typ == "openvswitch" {
		var ovs DomainInterfaceVirtualPortParamsOpenVSwitch
		a.Params.OpenVSwitch = &ovs
	} else if typ == "midonet" {
		var mido DomainInterfaceVirtualPortParamsMidoNet
		a.Params.MidoNet = &mido
	}

	vp := domainInterfaceVirtualPort(*a)
	err := d.DecodeElement(&vp, &start)
	if err != nil {
		return err
	}
	*a = DomainInterfaceVirtualPort(vp)
	return nil
}

func (a *DomainInterfaceSourceHostdev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if a.PCI != nil {
		addr := xml.StartElement{
			Name: xml.Name{Local: "address"},
		}
		addr.Attr = append(addr.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		e.EncodeElement(a.PCI.Address, addr)
	} else if a.USB != nil {
		addr := xml.StartElement{
			Name: xml.Name{Local: "address"},
		}
		addr.Attr = append(addr.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		e.EncodeElement(a.USB.Address, addr)
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainInterfaceSourceHostdev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "address" {
				typ, ok := getAttr(tok.Attr, "type")
				if !ok {
					return fmt.Errorf("Missing hostdev address type attribute")
				}

				if typ == "pci" {
					a.PCI = &DomainHostdevSubsysPCISource{
						"",
						&DomainAddressPCI{},
					}
					err := d.DecodeElement(a.PCI.Address, &tok)
					if err != nil {
						return err
					}
				} else if typ == "usb" {
					a.USB = &DomainHostdevSubsysUSBSource{
						&DomainAddressUSB{},
					}
					err := d.DecodeElement(a.USB, &tok)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainInterfaceSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.User != nil {
		/* We don't want an empty <source></source> for User mode */
		//return e.EncodeElement(a.User, start)
		return nil
	} else if a.Ethernet != nil {
		if len(a.Ethernet.IP) > 0 && len(a.Ethernet.Route) > 0 {
			return e.EncodeElement(a.Ethernet, start)
		}
		return nil
	} else if a.VHostUser != nil {
		typ := getChardevSourceType(a.VHostUser)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
		return e.EncodeElement(a.VHostUser, start)
	} else if a.Server != nil {
		return e.EncodeElement(a.Server, start)
	} else if a.Client != nil {
		return e.EncodeElement(a.Client, start)
	} else if a.MCast != nil {
		return e.EncodeElement(a.MCast, start)
	} else if a.Network != nil {
		return e.EncodeElement(a.Network, start)
	} else if a.Bridge != nil {
		return e.EncodeElement(a.Bridge, start)
	} else if a.Internal != nil {
		return e.EncodeElement(a.Internal, start)
	} else if a.Direct != nil {
		return e.EncodeElement(a.Direct, start)
	} else if a.Hostdev != nil {
		return e.EncodeElement(a.Hostdev, start)
	} else if a.UDP != nil {
		return e.EncodeElement(a.UDP, start)
	} else if a.VDPA != nil {
		return e.EncodeElement(a.VDPA, start)
	}
	return nil
}

func (a *DomainInterfaceSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.User != nil {
		return d.DecodeElement(a.User, &start)
	} else if a.Ethernet != nil {
		return d.DecodeElement(a.Ethernet, &start)
	} else if a.VHostUser != nil {
		typ, ok := getAttr(start.Attr, "type")
		if !ok {
			typ = "pty"
		}
		a.VHostUser = createChardevSource(typ)
		return d.DecodeElement(a.VHostUser, &start)
	} else if a.Server != nil {
		return d.DecodeElement(a.Server, &start)
	} else if a.Client != nil {
		return d.DecodeElement(a.Client, &start)
	} else if a.MCast != nil {
		return d.DecodeElement(a.MCast, &start)
	} else if a.Network != nil {
		return d.DecodeElement(a.Network, &start)
	} else if a.Bridge != nil {
		return d.DecodeElement(a.Bridge, &start)
	} else if a.Internal != nil {
		return d.DecodeElement(a.Internal, &start)
	} else if a.Direct != nil {
		return d.DecodeElement(a.Direct, &start)
	} else if a.Hostdev != nil {
		return d.DecodeElement(a.Hostdev, &start)
	} else if a.UDP != nil {
		return d.DecodeElement(a.UDP, &start)
	} else if a.VDPA != nil {
		return d.DecodeElement(a.VDPA, &start)
	}
	return nil
}

type domainInterface DomainInterface

func (a *DomainInterface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "interface"
	if a.Source != nil {
		if a.Source.User != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "user",
			})
		} else if a.Source.Ethernet != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "ethernet",
			})
		} else if a.Source.VHostUser != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vhostuser",
			})
		} else if a.Source.Server != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "server",
			})
		} else if a.Source.Client != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "client",
			})
		} else if a.Source.MCast != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "mcast",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Bridge != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "bridge",
			})
		} else if a.Source.Internal != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "internal",
			})
		} else if a.Source.Direct != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "direct",
			})
		} else if a.Source.Hostdev != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "hostdev",
			})
		} else if a.Source.UDP != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "udp",
			})
		} else if a.Source.VDPA != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vdpa",
			})
		}
	}
	fs := domainInterface(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainInterface) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing interface type attribute")
	}
	a.Source = &DomainInterfaceSource{}
	if typ == "user" {
		a.Source.User = &DomainInterfaceSourceUser{}
	} else if typ == "ethernet" {
		a.Source.Ethernet = &DomainInterfaceSourceEthernet{}
	} else if typ == "vhostuser" {
		a.Source.VHostUser = &DomainChardevSource{}
	} else if typ == "server" {
		a.Source.Server = &DomainInterfaceSourceServer{}
	} else if typ == "client" {
		a.Source.Client = &DomainInterfaceSourceClient{}
	} else if typ == "mcast" {
		a.Source.MCast = &DomainInterfaceSourceMCast{}
	} else if typ == "network" {
		a.Source.Network = &DomainInterfaceSourceNetwork{}
	} else if typ == "bridge" {
		a.Source.Bridge = &DomainInterfaceSourceBridge{}
	} else if typ == "internal" {
		a.Source.Internal = &DomainInterfaceSourceInternal{}
	} else if typ == "direct" {
		a.Source.Direct = &DomainInterfaceSourceDirect{}
	} else if typ == "hostdev" {
		a.Source.Hostdev = &DomainInterfaceSourceHostdev{}
	} else if typ == "udp" {
		a.Source.UDP = &DomainInterfaceSourceUDP{}
	} else if typ == "vdpa" {
		a.Source.VDPA = &DomainInterfaceSourceVDPA{}
	}
	fs := domainInterface(*a)
	err := d.DecodeElement(&fs, &start)
	if err != nil {
		return err
	}
	*a = DomainInterface(fs)
	return nil
}

func (d *DomainInterface) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInterface) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainSmartcard DomainSmartcard

func (a *DomainSmartcard) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "smartcard"
	if a.Passthrough != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "passthrough",
		})
		typ := getChardevSourceType(a.Passthrough)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	} else if a.Host != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "host",
		})
	} else if len(a.HostCerts) != 0 {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "host-certificates",
		})
	}
	smartcard := domainSmartcard(*a)
	return e.EncodeElement(smartcard, start)
}

func (a *DomainSmartcard) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mode, ok := getAttr(start.Attr, "mode")
	if !ok {
		return fmt.Errorf("Missing mode on smartcard device")
	}
	if mode == "host" {
		a.Host = &DomainSmartcardHost{}
	} else if mode == "passthrough" {
		typ, ok := getAttr(start.Attr, "type")
		if !ok {
			typ = "pty"
		}
		a.Passthrough = createChardevSource(typ)
	}
	smartcard := domainSmartcard(*a)
	err := d.DecodeElement(&smartcard, &start)
	if err != nil {
		return err
	}
	*a = DomainSmartcard(smartcard)
	return nil
}

func (d *DomainSmartcard) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSmartcard) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainTPMBackend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backend"
	if a.Passthrough != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "passthrough",
		})
		err := e.EncodeElement(a.Passthrough, start)
		if err != nil {
			return err
		}
	} else if a.Emulator != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "emulator",
		})
		err := e.EncodeElement(a.Emulator, start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *DomainTPMBackend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing TPM backend type")
	}
	if typ == "passthrough" {
		a.Passthrough = &DomainTPMBackendPassthrough{}
		err := d.DecodeElement(a.Passthrough, &start)
		if err != nil {
			return err
		}
	} else if typ == "emulator" {
		a.Emulator = &DomainTPMBackendEmulator{}
		err := d.DecodeElement(a.Emulator, &start)
		if err != nil {
			return err
		}
	} else {
		d.Skip()
	}
	return nil
}

func (d *DomainTPM) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainTPM) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainShmem) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainShmem) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func getChardevSourceType(s *DomainChardevSource) string {
	if s.Null != nil {
		return "null"
	} else if s.VC != nil {
		return "vc"
	} else if s.Pty != nil {
		return "pty"
	} else if s.Dev != nil {
		return "dev"
	} else if s.File != nil {
		return "file"
	} else if s.Pipe != nil {
		return "pipe"
	} else if s.StdIO != nil {
		return "stdio"
	} else if s.UDP != nil {
		return "udp"
	} else if s.TCP != nil {
		return "tcp"
	} else if s.UNIX != nil {
		return "unix"
	} else if s.SpiceVMC != nil {
		return "spicevmc"
	} else if s.SpicePort != nil {
		return "spiceport"
	} else if s.NMDM != nil {
		return "nmdm"
	}
	return ""
}

func createChardevSource(typ string) *DomainChardevSource {
	switch typ {
	case "null":
		return &DomainChardevSource{
			Null: &DomainChardevSourceNull{},
		}
	case "vc":
		return &DomainChardevSource{
			VC: &DomainChardevSourceVC{},
		}
	case "pty":
		return &DomainChardevSource{
			Pty: &DomainChardevSourcePty{},
		}
	case "dev":
		return &DomainChardevSource{
			Dev: &DomainChardevSourceDev{},
		}
	case "file":
		return &DomainChardevSource{
			File: &DomainChardevSourceFile{},
		}
	case "pipe":
		return &DomainChardevSource{
			Pipe: &DomainChardevSourcePipe{},
		}
	case "stdio":
		return &DomainChardevSource{
			StdIO: &DomainChardevSourceStdIO{},
		}
	case "udp":
		return &DomainChardevSource{
			UDP: &DomainChardevSourceUDP{},
		}
	case "tcp":
		return &DomainChardevSource{
			TCP: &DomainChardevSourceTCP{},
		}
	case "unix":
		return &DomainChardevSource{
			UNIX: &DomainChardevSourceUNIX{},
		}
	case "spicevmc":
		return &DomainChardevSource{
			SpiceVMC: &DomainChardevSourceSpiceVMC{},
		}
	case "spiceport":
		return &DomainChardevSource{
			SpicePort: &DomainChardevSourceSpicePort{},
		}
	case "nmdm":
		return &DomainChardevSource{
			NMDM: &DomainChardevSourceNMDM{},
		}
	}

	return nil
}

type domainChardevSourceUDPFlat struct {
	Mode    string `xml:"mode,attr"`
	Host    string `xml:"host,attr,omitempty"`
	Service string `xml:"service,attr,omitempty"`
}

func (a *DomainChardevSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Null != nil {
		return nil
	} else if a.VC != nil {
		return nil
	} else if a.Pty != nil {
		if a.Pty.Path != "" {
			return e.EncodeElement(a.Pty, start)
		}
		return nil
	} else if a.Dev != nil {
		return e.EncodeElement(a.Dev, start)
	} else if a.File != nil {
		return e.EncodeElement(a.File, start)
	} else if a.Pipe != nil {
		return e.EncodeElement(a.Pipe, start)
	} else if a.StdIO != nil {
		return nil
	} else if a.UDP != nil {
		srcs := []domainChardevSourceUDPFlat{
			domainChardevSourceUDPFlat{
				Mode:    "bind",
				Host:    a.UDP.BindHost,
				Service: a.UDP.BindService,
			},
			domainChardevSourceUDPFlat{
				Mode:    "connect",
				Host:    a.UDP.ConnectHost,
				Service: a.UDP.ConnectService,
			},
		}
		if srcs[0].Host != "" || srcs[0].Service != "" {
			err := e.EncodeElement(&srcs[0], start)
			if err != nil {
				return err
			}
		}
		if srcs[1].Host != "" || srcs[1].Service != "" {
			err := e.EncodeElement(&srcs[1], start)
			if err != nil {
				return err
			}
		}
	} else if a.TCP != nil {
		return e.EncodeElement(a.TCP, start)
	} else if a.UNIX != nil {
		if a.UNIX.Path == "" && a.UNIX.Mode == "" {
			return nil
		}
		return e.EncodeElement(a.UNIX, start)
	} else if a.SpiceVMC != nil {
		return nil
	} else if a.SpicePort != nil {
		return e.EncodeElement(a.SpicePort, start)
	} else if a.NMDM != nil {
		return e.EncodeElement(a.NMDM, start)
	}
	return nil
}

func (a *DomainChardevSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Null != nil {
		d.Skip()
		return nil
	} else if a.VC != nil {
		d.Skip()
		return nil
	} else if a.Pty != nil {
		return d.DecodeElement(a.Pty, &start)
	} else if a.Dev != nil {
		return d.DecodeElement(a.Dev, &start)
	} else if a.File != nil {
		return d.DecodeElement(a.File, &start)
	} else if a.Pipe != nil {
		return d.DecodeElement(a.Pipe, &start)
	} else if a.StdIO != nil {
		d.Skip()
		return nil
	} else if a.UDP != nil {
		src := domainChardevSourceUDPFlat{}
		err := d.DecodeElement(&src, &start)
		if src.Mode == "connect" {
			a.UDP.ConnectHost = src.Host
			a.UDP.ConnectService = src.Service
		} else {
			a.UDP.BindHost = src.Host
			a.UDP.BindService = src.Service
		}
		return err
	} else if a.TCP != nil {
		return d.DecodeElement(a.TCP, &start)
	} else if a.UNIX != nil {
		return d.DecodeElement(a.UNIX, &start)
	} else if a.SpiceVMC != nil {
		d.Skip()
		return nil
	} else if a.SpicePort != nil {
		return d.DecodeElement(a.SpicePort, &start)
	} else if a.NMDM != nil {
		return d.DecodeElement(a.NMDM, &start)
	}
	return nil
}

type domainConsole DomainConsole

func (a *DomainConsole) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "console"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainConsole(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainConsole) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainConsole(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainConsole(con)
	return nil
}

func (d *DomainConsole) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainConsole) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainSerial DomainSerial

func (a *DomainSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "serial"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	s := domainSerial(*a)
	return e.EncodeElement(s, start)
}

func (a *DomainSerial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainSerial(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainSerial(con)
	return nil
}

func (d *DomainSerial) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSerial) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainParallel DomainParallel

func (a *DomainParallel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "parallel"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	s := domainParallel(*a)
	return e.EncodeElement(s, start)
}

func (a *DomainParallel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainParallel(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainParallel(con)
	return nil
}

func (d *DomainParallel) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainParallel) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainInput) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInput) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainVideo) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainVideo) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainChannelTarget DomainChannelTarget

func (a *DomainChannelTarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.VirtIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio",
		})
		return e.EncodeElement(a.VirtIO, start)
	} else if a.Xen != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "xen",
		})
		return e.EncodeElement(a.Xen, start)
	} else if a.GuestFWD != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "guestfwd",
		})
		return e.EncodeElement(a.GuestFWD, start)
	}
	return nil
}

func (a *DomainChannelTarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing channel target type")
	}
	if typ == "virtio" {
		a.VirtIO = &DomainChannelTargetVirtIO{}
		return d.DecodeElement(a.VirtIO, &start)
	} else if typ == "xen" {
		a.Xen = &DomainChannelTargetXen{}
		return d.DecodeElement(a.Xen, &start)
	} else if typ == "guestfwd" {
		a.GuestFWD = &DomainChannelTargetGuestFWD{}
		return d.DecodeElement(a.GuestFWD, &start)
	}
	d.Skip()
	return nil
}

type domainChannel DomainChannel

func (a *DomainChannel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "channel"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainChannel(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainChannel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainChannel(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainChannel(con)
	return nil
}

func (d *DomainChannel) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainChannel) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainRedirFilterUSB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "class", a.Class, "0x%02x")
	marshalUintAttr(&start, "vendor", a.Vendor, "0x%04x")
	marshalUintAttr(&start, "product", a.Product, "0x%04x")
	if a.Version != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "version"}, a.Version,
		})
	}
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "allow"}, a.Allow,
	})
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainRedirFilterUSB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "class" && attr.Value != "-1" {
			if err := unmarshalIntAttr(attr.Value, &a.Class, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "product" && attr.Value != "-1" {
			if err := unmarshalIntAttr(attr.Value, &a.Product, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "vendor" && attr.Value != "-1" {
			if err := unmarshalIntAttr(attr.Value, &a.Vendor, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "version" && attr.Value != "-1" {
			a.Version = attr.Value
		} else if attr.Name.Local == "allow" {
			a.Allow = attr.Value
		}
	}
	d.Skip()
	return nil
}

type domainRedirDev DomainRedirDev

func (a *DomainRedirDev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "redirdev"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainRedirDev(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainRedirDev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainRedirDev(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainRedirDev(con)
	return nil
}

func (d *DomainRedirDev) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainRedirDev) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainMemBalloon) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemBalloon) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainVSock) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainVSock) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainSound) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSound) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainRNGBackendEGD DomainRNGBackendEGD

func (a *DomainRNGBackendEGD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backend"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	egd := domainRNGBackendEGD(*a)
	return e.EncodeElement(egd, start)
}

func (a *DomainRNGBackendEGD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainRNGBackendEGD(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainRNGBackendEGD(con)
	return nil
}

func (a *DomainRNGBackend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Random != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "random",
		})
		return e.EncodeElement(a.Random, start)
	} else if a.EGD != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "egd",
		})
		return e.EncodeElement(a.EGD, start)
	} else if a.BuiltIn != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "builtin",
		})
		return e.EncodeElement(a.BuiltIn, start)
	}
	return nil
}

func (a *DomainRNGBackend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	model, ok := getAttr(start.Attr, "model")
	if !ok {
		return nil
	}
	if model == "random" {
		a.Random = &DomainRNGBackendRandom{}
		err := d.DecodeElement(a.Random, &start)
		if err != nil {
			return err
		}
	} else if model == "egd" {
		a.EGD = &DomainRNGBackendEGD{}
		err := d.DecodeElement(a.EGD, &start)
		if err != nil {
			return err
		}
	} else if model == "builtin" {
		a.BuiltIn = &DomainRNGBackendBuiltIn{}
		err := d.DecodeElement(a.BuiltIn, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
}

func (d *DomainRNG) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainRNG) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainHostdevSubsysSCSISource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Host != nil {
		return e.EncodeElement(a.Host, start)
	} else if a.ISCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "protocol"}, "iscsi",
		})
		return e.EncodeElement(a.ISCSI, start)
	}
	return nil
}

func (a *DomainHostdevSubsysSCSISource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	proto, ok := getAttr(start.Attr, "protocol")
	if !ok {
		a.Host = &DomainHostdevSubsysSCSISourceHost{}
		err := d.DecodeElement(a.Host, &start)
		if err != nil {
			return err
		}
	}
	if proto == "iscsi" {
		a.ISCSI = &DomainHostdevSubsysSCSISourceISCSI{}
		err := d.DecodeElement(a.ISCSI, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
}

type domainHostdev DomainHostdev

type domainHostdevSubsysSCSI struct {
	DomainHostdevSubsysSCSI
	domainHostdev
}

type domainHostdevSubsysSCSIHost struct {
	DomainHostdevSubsysSCSIHost
	domainHostdev
}

type domainHostdevSubsysUSB struct {
	DomainHostdevSubsysUSB
	domainHostdev
}

type domainHostdevSubsysPCI struct {
	DomainHostdevSubsysPCI
	domainHostdev
}

type domainHostdevSubsysMDev struct {
	DomainHostdevSubsysMDev
	domainHostdev
}

type domainHostdevCapsStorage struct {
	DomainHostdevCapsStorage
	domainHostdev
}

type domainHostdevCapsMisc struct {
	DomainHostdevCapsMisc
	domainHostdev
}

type domainHostdevCapsNet struct {
	DomainHostdevCapsNet
	domainHostdev
}

func (a *DomainHostdev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "hostdev"
	if a.SubsysSCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		scsi := domainHostdevSubsysSCSI{}
		scsi.domainHostdev = domainHostdev(*a)
		scsi.DomainHostdevSubsysSCSI = *a.SubsysSCSI
		return e.EncodeElement(scsi, start)
	} else if a.SubsysSCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		scsi_host := domainHostdevSubsysSCSIHost{}
		scsi_host.domainHostdev = domainHostdev(*a)
		scsi_host.DomainHostdevSubsysSCSIHost = *a.SubsysSCSIHost
		return e.EncodeElement(scsi_host, start)
	} else if a.SubsysUSB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		usb := domainHostdevSubsysUSB{}
		usb.domainHostdev = domainHostdev(*a)
		usb.DomainHostdevSubsysUSB = *a.SubsysUSB
		return e.EncodeElement(usb, start)
	} else if a.SubsysPCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		pci := domainHostdevSubsysPCI{}
		pci.domainHostdev = domainHostdev(*a)
		pci.DomainHostdevSubsysPCI = *a.SubsysPCI
		return e.EncodeElement(pci, start)
	} else if a.SubsysMDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		mdev := domainHostdevSubsysMDev{}
		mdev.domainHostdev = domainHostdev(*a)
		mdev.DomainHostdevSubsysMDev = *a.SubsysMDev
		return e.EncodeElement(mdev, start)
	} else if a.CapsStorage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		storage := domainHostdevCapsStorage{}
		storage.domainHostdev = domainHostdev(*a)
		storage.DomainHostdevCapsStorage = *a.CapsStorage
		return e.EncodeElement(storage, start)
	} else if a.CapsMisc != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "misc",
		})
		misc := domainHostdevCapsMisc{}
		misc.domainHostdev = domainHostdev(*a)
		misc.DomainHostdevCapsMisc = *a.CapsMisc
		return e.EncodeElement(misc, start)
	} else if a.CapsNet != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		net := domainHostdevCapsNet{}
		net.domainHostdev = domainHostdev(*a)
		net.DomainHostdevCapsNet = *a.CapsNet
		return e.EncodeElement(net, start)
	} else {
		gen := domainHostdev(*a)
		return e.EncodeElement(gen, start)
	}
}

func (a *DomainHostdev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mode, ok := getAttr(start.Attr, "mode")
	if !ok {
		return fmt.Errorf("Missing 'mode' attribute on domain hostdev")
	}
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if mode == "subsystem" {
		if typ == "scsi" {
			var scsi domainHostdevSubsysSCSI
			err := d.DecodeElement(&scsi, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi.domainHostdev)
			a.SubsysSCSI = &scsi.DomainHostdevSubsysSCSI
			return nil
		} else if typ == "scsi_host" {
			var scsi_host domainHostdevSubsysSCSIHost
			err := d.DecodeElement(&scsi_host, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi_host.domainHostdev)
			a.SubsysSCSIHost = &scsi_host.DomainHostdevSubsysSCSIHost
			return nil
		} else if typ == "usb" {
			var usb domainHostdevSubsysUSB
			err := d.DecodeElement(&usb, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(usb.domainHostdev)
			a.SubsysUSB = &usb.DomainHostdevSubsysUSB
			return nil
		} else if typ == "pci" {
			var pci domainHostdevSubsysPCI
			err := d.DecodeElement(&pci, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(pci.domainHostdev)
			a.SubsysPCI = &pci.DomainHostdevSubsysPCI
			return nil
		} else if typ == "mdev" {
			var mdev domainHostdevSubsysMDev
			err := d.DecodeElement(&mdev, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(mdev.domainHostdev)
			a.SubsysMDev = &mdev.DomainHostdevSubsysMDev
			return nil
		}
	} else if mode == "capabilities" {
		if typ == "storage" {
			var storage domainHostdevCapsStorage
			err := d.DecodeElement(&storage, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(storage.domainHostdev)
			a.CapsStorage = &storage.DomainHostdevCapsStorage
			return nil
		} else if typ == "misc" {
			var misc domainHostdevCapsMisc
			err := d.DecodeElement(&misc, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(misc.domainHostdev)
			a.CapsMisc = &misc.DomainHostdevCapsMisc
			return nil
		} else if typ == "net" {
			var net domainHostdevCapsNet
			err := d.DecodeElement(&net, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(net.domainHostdev)
			a.CapsNet = &net.DomainHostdevCapsNet
			return nil
		}
	}
	var gen domainHostdev
	err := d.DecodeElement(&gen, &start)
	if err != nil {
		return err
	}
	*a = DomainHostdev(gen)
	return nil
}

func (d *DomainHostdev) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainHostdev) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainGraphicListener) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "listen"
	if a.Address != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "address",
		})
		return e.EncodeElement(a.Address, start)
	} else if a.Network != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "network",
		})
		return e.EncodeElement(a.Network, start)
	} else if a.Socket != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "socket",
		})
		return e.EncodeElement(a.Socket, start)
	} else {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "none",
		})
		e.EncodeToken(start)
		e.EncodeToken(start.End())
	}
	return nil
}

func (a *DomainGraphicListener) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain graphics listen")
	}
	if typ == "address" {
		var addr DomainGraphicListenerAddress
		err := d.DecodeElement(&addr, &start)
		if err != nil {
			return err
		}
		a.Address = &addr
		return nil
	} else if typ == "network" {
		var net DomainGraphicListenerNetwork
		err := d.DecodeElement(&net, &start)
		if err != nil {
			return err
		}
		a.Network = &net
		return nil
	} else if typ == "socket" {
		var sock DomainGraphicListenerSocket
		err := d.DecodeElement(&sock, &start)
		if err != nil {
			return err
		}
		a.Socket = &sock
		return nil
	} else if typ == "none" {
		d.Skip()
	}
	return nil
}

type domainGraphicSDL struct {
	DomainGraphicSDL
	Audio *DomainGraphicAudio `xml:"audio"`
}

type domainGraphicVNC struct {
	DomainGraphicVNC
	Audio *DomainGraphicAudio `xml:"audio"`
}

type domainGraphicRDP struct {
	DomainGraphicRDP
	Audio *DomainGraphicAudio `xml:"audio"`
}

type domainGraphicDesktop struct {
	DomainGraphicDesktop
	Audio *DomainGraphicAudio `xml:"audio"`
}

type domainGraphicSpice struct {
	DomainGraphicSpice
	Audio *DomainGraphicAudio `xml:"audio"`
}

type domainGraphicEGLHeadless struct {
	DomainGraphicEGLHeadless
	Audio *DomainGraphicAudio `xml:"audio"`
}

func (a *DomainGraphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "graphics"
	if a.SDL != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "sdl",
		})
		sdl := domainGraphicSDL{*a.SDL, a.Audio}
		return e.EncodeElement(sdl, start)
	} else if a.VNC != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "vnc",
		})
		vnc := domainGraphicVNC{*a.VNC, a.Audio}
		return e.EncodeElement(vnc, start)
	} else if a.RDP != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "rdp",
		})
		rdp := domainGraphicRDP{*a.RDP, a.Audio}
		return e.EncodeElement(rdp, start)
	} else if a.Desktop != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "desktop",
		})
		desktop := domainGraphicDesktop{*a.Desktop, a.Audio}
		return e.EncodeElement(desktop, start)
	} else if a.Spice != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spice",
		})
		spice := domainGraphicSpice{*a.Spice, a.Audio}
		return e.EncodeElement(spice, start)
	} else if a.EGLHeadless != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "egl-headless",
		})
		egl := domainGraphicEGLHeadless{*a.EGLHeadless, a.Audio}
		return e.EncodeElement(egl, start)
	}
	return nil
}

func (a *DomainGraphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain graphics")
	}
	if typ == "sdl" {
		var sdl domainGraphicSDL
		err := d.DecodeElement(&sdl, &start)
		if err != nil {
			return err
		}
		a.SDL = &sdl.DomainGraphicSDL
		a.Audio = sdl.Audio
		return nil
	} else if typ == "vnc" {
		var vnc domainGraphicVNC
		err := d.DecodeElement(&vnc, &start)
		if err != nil {
			return err
		}
		a.VNC = &vnc.DomainGraphicVNC
		a.Audio = vnc.Audio
		return nil
	} else if typ == "rdp" {
		var rdp domainGraphicRDP
		err := d.DecodeElement(&rdp, &start)
		if err != nil {
			return err
		}
		a.RDP = &rdp.DomainGraphicRDP
		a.Audio = rdp.Audio
		return nil
	} else if typ == "desktop" {
		var desktop domainGraphicDesktop
		err := d.DecodeElement(&desktop, &start)
		if err != nil {
			return err
		}
		a.Desktop = &desktop.DomainGraphicDesktop
		a.Audio = desktop.Audio
		return nil
	} else if typ == "spice" {
		var spice domainGraphicSpice
		err := d.DecodeElement(&spice, &start)
		if err != nil {
			return err
		}
		a.Spice = &spice.DomainGraphicSpice
		a.Audio = spice.Audio
		return nil
	} else if typ == "egl-headless" {
		var egl domainGraphicEGLHeadless
		err := d.DecodeElement(&egl, &start)
		if err != nil {
			return err
		}
		a.EGLHeadless = &egl.DomainGraphicEGLHeadless
		a.Audio = egl.Audio
		return nil
	}
	return nil
}

func (a *DomainAudio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "audio"
	if a.ID != 0 {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "id"}, fmt.Sprintf("%d", a.ID),
		})
	}
	if a.None != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "none",
		})
		return e.EncodeElement(a.None, start)
	} else if a.ALSA != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "alsa",
		})
		return e.EncodeElement(a.ALSA, start)
	} else if a.CoreAudio != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "coreaudio",
		})
		return e.EncodeElement(a.CoreAudio, start)
	} else if a.Jack != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "jack",
		})
		return e.EncodeElement(a.Jack, start)
	} else if a.OSS != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "oss",
		})
		return e.EncodeElement(a.OSS, start)
	} else if a.PulseAudio != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pulseaudio",
		})
		return e.EncodeElement(a.PulseAudio, start)
	} else if a.SDL != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "sdl",
		})
		return e.EncodeElement(a.SDL, start)
	} else if a.SPICE != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spice",
		})
		return e.EncodeElement(a.SPICE, start)
	} else if a.File != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "file",
		})
		return e.EncodeElement(a.File, start)
	}
	return nil
}

func (a *DomainAudio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain audio")
	}
	id, ok := getAttr(start.Attr, "id")
	if ok {
		idval, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			return err
		}
		a.ID = int(idval)
	}

	if typ == "none" {
		var none DomainAudioNone
		err := d.DecodeElement(&none, &start)
		if err != nil {
			return err
		}
		a.None = &none
		return nil
	} else if typ == "alsa" {
		var alsa DomainAudioALSA
		err := d.DecodeElement(&alsa, &start)
		if err != nil {
			return err
		}
		a.ALSA = &alsa
		return nil
	} else if typ == "coreaudio" {
		var coreaudio DomainAudioCoreAudio
		err := d.DecodeElement(&coreaudio, &start)
		if err != nil {
			return err
		}
		a.CoreAudio = &coreaudio
		return nil
	} else if typ == "jack" {
		var jack DomainAudioJack
		err := d.DecodeElement(&jack, &start)
		if err != nil {
			return err
		}
		a.Jack = &jack
		return nil
	} else if typ == "oss" {
		var oss DomainAudioOSS
		err := d.DecodeElement(&oss, &start)
		if err != nil {
			return err
		}
		a.OSS = &oss
		return nil
	} else if typ == "pulseaudio" {
		var pulseaudio DomainAudioPulseAudio
		err := d.DecodeElement(&pulseaudio, &start)
		if err != nil {
			return err
		}
		a.PulseAudio = &pulseaudio
		return nil
	} else if typ == "sdl" {
		var sdl DomainAudioSDL
		err := d.DecodeElement(&sdl, &start)
		if err != nil {
			return err
		}
		a.SDL = &sdl
		return nil
	} else if typ == "spice" {
		var spice DomainAudioSPICE
		err := d.DecodeElement(&spice, &start)
		if err != nil {
			return err
		}
		a.SPICE = &spice
		return nil
	} else if typ == "file" {
		var file DomainAudioFile
		err := d.DecodeElement(&file, &start)
		if err != nil {
			return err
		}
		a.File = &file
		return nil
	}
	return nil
}

func (d *DomainMemorydev) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemorydev) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainWatchdog) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainWatchdog) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func marshalUintAttr(start *xml.StartElement, name string, val *int32, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func marshalUint64Attr(start *xml.StartElement, name string, val *int64, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func (a *DomainAddressPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	if a.MultiFunction != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "multifunction"}, a.MultiFunction,
		})
	}
	e.EncodeToken(start)
	if a.ZPCI != nil {
		zpci := xml.StartElement{}
		zpci.Name.Local = "zpci"
		err := e.EncodeElement(a.ZPCI, zpci)
		if err != nil {
			return err
		}
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressZPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "uid", a.UID, "0x%04x")
	marshalUintAttr(&start, "fid", a.FID, "0x%04x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressUSB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	if a.Port != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "port"}, a.Port,
		})
	}
	marshalUintAttr(&start, "device", a.Device, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDrive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	marshalUintAttr(&start, "target", a.Target, "%d")
	marshalUintAttr(&start, "unit", a.Unit, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDIMM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "slot", a.Slot, "%d")
	marshalUint64Attr(&start, "base", a.Base, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressISA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "iobase", a.IOBase, "0x%x")
	marshalUintAttr(&start, "irq", a.IRQ, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioMMIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCW) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "cssid", a.CSSID, "0x%x")
	marshalUintAttr(&start, "ssid", a.SSID, "0x%x")
	marshalUintAttr(&start, "devno", a.DevNo, "0x%04x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	marshalUintAttr(&start, "port", a.Port, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressSpaprVIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUint64Attr(&start, "reg", a.Reg, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "slot", a.Slot, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioS390) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressUnassigned) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return e.EncodeElement(a.USB, start)
	} else if a.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(a.PCI, start)
	} else if a.Drive != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drive",
		})
		return e.EncodeElement(a.Drive, start)
	} else if a.DIMM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "dimm",
		})
		return e.EncodeElement(a.DIMM, start)
	} else if a.ISA != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "isa",
		})
		return e.EncodeElement(a.ISA, start)
	} else if a.VirtioMMIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-mmio",
		})
		return e.EncodeElement(a.VirtioMMIO, start)
	} else if a.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return e.EncodeElement(a.CCW, start)
	} else if a.VirtioSerial != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-serial",
		})
		return e.EncodeElement(a.VirtioSerial, start)
	} else if a.SpaprVIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spapr-vio",
		})
		return e.EncodeElement(a.SpaprVIO, start)
	} else if a.CCID != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccid",
		})
		return e.EncodeElement(a.CCID, start)
	} else if a.VirtioS390 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-s390",
		})
		return e.EncodeElement(a.VirtioS390, start)
	} else if a.Unassigned != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "unassigned",
		})
		return e.EncodeElement(a.Unassigned, start)
	} else {
		return nil
	}
}

func unmarshalInt64Attr(valstr string, valptr **int64, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseInt(valstr, base, 64)
	if err != nil {
		return err
	}
	*valptr = &val
	return nil
}

func unmarshalIntAttr(valstr string, valptr **int32, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseInt(valstr, base, 64)
	if err != nil {
		return err
	}
	vali := int32(val)
	*valptr = &vali
	return nil
}

func (a *DomainAddressUSB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "bus" {
			if err := unmarshalIntAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			a.Port = attr.Value
		} else if attr.Name.Local == "device" {
			if err := unmarshalIntAttr(attr.Value, &a.Device, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "domain" {
			if err := unmarshalIntAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalIntAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalIntAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshalIntAttr(attr.Value, &a.Function, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "multifunction" {
			a.MultiFunction = attr.Value
		}
	}

	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "zpci" {
				a.ZPCI = &DomainAddressZPCI{}
				err = d.DecodeElement(a.ZPCI, &tok)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *DomainAddressZPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "fid" {
			if err := unmarshalIntAttr(attr.Value, &a.FID, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "uid" {
			if err := unmarshalIntAttr(attr.Value, &a.UID, 0); err != nil {
				return err
			}
		}
	}

	d.Skip()
	return nil
}

func (a *DomainAddressDrive) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalIntAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalIntAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "target" {
			if err := unmarshalIntAttr(attr.Value, &a.Target, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "unit" {
			if err := unmarshalIntAttr(attr.Value, &a.Unit, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressDIMM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "slot" {
			if err := unmarshalIntAttr(attr.Value, &a.Slot, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "base" {
			if err := unmarshalInt64Attr(attr.Value, &a.Base, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressISA) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "iobase" {
			if err := unmarshalIntAttr(attr.Value, &a.IOBase, 16); err != nil {
				return err
			}
		} else if attr.Name.Local == "irq" {
			if err := unmarshalIntAttr(attr.Value, &a.IRQ, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioMMIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddressCCW) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "cssid" {
			if err := unmarshalIntAttr(attr.Value, &a.CSSID, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "ssid" {
			if err := unmarshalIntAttr(attr.Value, &a.SSID, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "devno" {
			if err := unmarshalIntAttr(attr.Value, &a.DevNo, 0); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioSerial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalIntAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalIntAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshalIntAttr(attr.Value, &a.Port, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressSpaprVIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "reg" {
			if err := unmarshalInt64Attr(attr.Value, &a.Reg, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressCCID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalIntAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalIntAttr(attr.Value, &a.Slot, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioS390) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddressUnassigned) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
			break
		}
	}
	if typ == "" {
		d.Skip()
		return nil
	}

	if typ == "usb" {
		a.USB = &DomainAddressUSB{}
		return d.DecodeElement(a.USB, &start)
	} else if typ == "pci" {
		a.PCI = &DomainAddressPCI{}
		return d.DecodeElement(a.PCI, &start)
	} else if typ == "drive" {
		a.Drive = &DomainAddressDrive{}
		return d.DecodeElement(a.Drive, &start)
	} else if typ == "dimm" {
		a.DIMM = &DomainAddressDIMM{}
		return d.DecodeElement(a.DIMM, &start)
	} else if typ == "isa" {
		a.ISA = &DomainAddressISA{}
		return d.DecodeElement(a.ISA, &start)
	} else if typ == "virtio-mmio" {
		a.VirtioMMIO = &DomainAddressVirtioMMIO{}
		return d.DecodeElement(a.VirtioMMIO, &start)
	} else if typ == "ccw" {
		a.CCW = &DomainAddressCCW{}
		return d.DecodeElement(a.CCW, &start)
	} else if typ == "virtio-serial" {
		a.VirtioSerial = &DomainAddressVirtioSerial{}
		return d.DecodeElement(a.VirtioSerial, &start)
	} else if typ == "spapr-vio" {
		a.SpaprVIO = &DomainAddressSpaprVIO{}
		return d.DecodeElement(a.SpaprVIO, &start)
	} else if typ == "ccid" {
		a.CCID = &DomainAddressCCID{}
		return d.DecodeElement(a.CCID, &start)
	} else if typ == "virtio-s390" {
		a.VirtioS390 = &DomainAddressVirtioS390{}
		return d.DecodeElement(a.VirtioS390, &start)
	} else if typ == "unassigned" {
		a.Unassigned = &DomainAddressUnassigned{}
		return d.DecodeElement(a.Unassigned, &start)
	}

	return nil
}

func (d *DomainCPU) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainCPU) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainLaunchSecuritySEV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)

	if a.CBitPos != nil {
		cbitpos := xml.StartElement{
			Name: xml.Name{Local: "cbitpos"},
		}
		e.EncodeToken(cbitpos)
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.CBitPos)))
		e.EncodeToken(cbitpos.End())
	}

	if a.ReducedPhysBits != nil {
		reducedPhysBits := xml.StartElement{
			Name: xml.Name{Local: "reducedPhysBits"},
		}
		e.EncodeToken(reducedPhysBits)
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.ReducedPhysBits)))
		e.EncodeToken(reducedPhysBits.End())
	}

	if a.Policy != nil {
		policy := xml.StartElement{
			Name: xml.Name{Local: "policy"},
		}
		e.EncodeToken(policy)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *a.Policy)))
		e.EncodeToken(policy.End())
	}

	dhcert := xml.StartElement{
		Name: xml.Name{Local: "dhCert"},
	}
	e.EncodeToken(dhcert)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%s", a.DHCert)))
	e.EncodeToken(dhcert.End())

	session := xml.StartElement{
		Name: xml.Name{Local: "session"},
	}
	e.EncodeToken(session)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%s", a.Session)))
	e.EncodeToken(session.End())

	e.EncodeToken(start.End())

	return nil
}

func (a *DomainLaunchSecuritySEV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "policy" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalIntAttr(string(data), &a.Policy, 16); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "cbitpos" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalIntAttr(string(data), &a.CBitPos, 10); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "reducedPhysBits" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalIntAttr(string(data), &a.ReducedPhysBits, 10); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "dhCert" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					a.DHCert = string(data)
				}
			} else if tok.Name.Local == "session" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					a.Session = string(data)
				}
			}
		}
	}
	return nil
}

func (a *DomainLaunchSecurity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if a.SEV != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "sev",
		})
		return e.EncodeElement(a.SEV, start)
	} else {
		return nil
	}

}

func (a *DomainLaunchSecurity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
		}
	}

	if typ == "" {
		d.Skip()
		return nil
	}

	if typ == "sev" {
		a.SEV = &DomainLaunchSecuritySEV{}
		return d.DecodeElement(a.SEV, &start)
	}

	return nil
}

type domainSysInfo DomainSysInfo

type domainSysInfoSMBIOS struct {
	DomainSysInfoSMBIOS
	domainSysInfo
}

type domainSysInfoFWCfg struct {
	DomainSysInfoFWCfg
	domainSysInfo
}

func (a *DomainSysInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "sysinfo"
	if a.SMBIOS != nil {
		smbios := domainSysInfoSMBIOS{}
		smbios.domainSysInfo = domainSysInfo(*a)
		smbios.DomainSysInfoSMBIOS = *a.SMBIOS
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "smbios",
		})
		return e.EncodeElement(smbios, start)
	} else if a.FWCfg != nil {
		fwcfg := domainSysInfoFWCfg{}
		fwcfg.domainSysInfo = domainSysInfo(*a)
		fwcfg.DomainSysInfoFWCfg = *a.FWCfg
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "fwcfg",
		})
		return e.EncodeElement(fwcfg, start)
	} else {
		gen := domainSysInfo(*a)
		return e.EncodeElement(gen, start)
	}
}

func (a *DomainSysInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if typ == "smbios" {
		var smbios domainSysInfoSMBIOS
		err := d.DecodeElement(&smbios, &start)
		if err != nil {
			return err
		}
		*a = DomainSysInfo(smbios.domainSysInfo)
		a.SMBIOS = &smbios.DomainSysInfoSMBIOS
		return nil
	} else if typ == "fwcfg" {
		var fwcfg domainSysInfoFWCfg
		err := d.DecodeElement(&fwcfg, &start)
		if err != nil {
			return err
		}
		*a = DomainSysInfo(fwcfg.domainSysInfo)
		a.FWCfg = &fwcfg.DomainSysInfoFWCfg
		return nil
	} else {
		var gen domainSysInfo
		err := d.DecodeElement(&gen, &start)
		if err != nil {
			return err
		}
		*a = DomainSysInfo(gen)
		return nil
	}
}
