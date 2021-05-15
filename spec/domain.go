package spec

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// +gogo:genproto=true
type DomainControllerPCIHole64 struct {
	Size_ int64  `xml:",chardata" json:"size" protobuf:"varint,1,opt,name=size"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainControllerPCIModel struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainControllerPCITarget struct {
	ChassisNr *int32 `json:"chassisNr,omitempty" protobuf:"varint,1,opt,name=chassisNr"`
	Chassis   *int32 `json:"chassis,omitempty" protobuf:"varint,2,opt,name=chassis"`
	Port      *int32 `json:"port,omitempty" protobuf:"varint,3,opt,name=port"`
	BusNr     *int32 `json:"busNr,omitempty" protobuf:"varint,4,opt,name=busNr"`
	Index     *int32 `json:"index,omitempty" protobuf:"varint,5,opt,name=index"`
	NUMANode  *int32 `json:"numaNode,omitempty" protobuf:"varint,6,opt,name=numaNode"`
	Hotplug   string `json:"hotplug,omitempty" protobuf:"bytes,7,opt,name=hotplug"`
}

// +gogo:genproto=true
type DomainControllerPCI struct {
	Model  *DomainControllerPCIModel  `xml:"model" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	Target *DomainControllerPCITarget `xml:"target" json:"target,omitempty" protobuf:"bytes,2,opt,name=target"`
	Hole64 *DomainControllerPCIHole64 `xml:"pcihole64" json:"pcihole64,omitempty" protobuf:"bytes,3,opt,name=pcihole64"`
}

// +gogo:genproto=true
type DomainControllerUSBMaster struct {
	StartPort int32 `xml:"startport,attr" json:"startPort" protobuf:"varint,1,opt,name=startPort"`
}

// +gogo:genproto=true
type DomainControllerUSB struct {
	Port   *int32                     `xml:"ports,attr" json:"port,omitempty" protobuf:"varint,1,opt,name=port"`
	Master *DomainControllerUSBMaster `xml:"master,omitempty" json:"master,omitempty" protobuf:"bytes,2,opt,name=master"`
}

// +gogo:genproto=true
type DomainControllerVirtIOSerial struct {
	Ports   *int32 `xml:"ports,attr" json:"ports,omitempty" protobuf:"varint,1,opt,name=ports"`
	Vectors *int32 `xml:"vectors,attr" json:"vectors,omitempty" protobuf:"varint,2,opt,name=vectors"`
}

// +gogo:genproto=true
type DomainControllerXenBus struct {
	MaxGrantFrames   int32 `xml:"maxGrantFrames,attr,omitempty" json:"maxGrantFrames,omitempty" protobuf:"varint,1,opt,name=maxGrantFrames"`
	MaxEventChannels int32 `xml:"maxEventChannels,attr,omitempty" json:"maxEventChannels,omitempty" protobuf:"varint,2,opt,name=maxEventChannels"`
}

// +gogo:genproto=true
type DomainControllerDriver struct {
	Queues     *int32 `xml:"queues,attr" json:"queues,omitempty" protobuf:"varint,1,opt,name=queues"`
	CmdPerLUN  *int32 `xml:"cmd_per_lun,attr" json:"cmdPerLun,omitempty" protobuf:"varint,2,opt,name=cmdPerLun"`
	MaxSectors *int32 `xml:"max_sectors,attr" json:"maxSectors,omitempty" protobuf:"varint,3,opt,name=maxSectors"`
	IOEventFD  string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" protobuf:"bytes,4,opt,name=ioeventfd"`
	IOThread   int32  `xml:"iothread,attr,omitempty" json:"iothread,omitempty" protobuf:"varint,5,opt,name=iothread"`
	IOMMU      string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,6,opt,name=iommu"`
	ATS        string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,7,opt,name=ats"`
	Packed     string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,8,opt,name=packed"`
}

// +gogo:genproto=true
type DomainController struct {
	XMLName      xml.Name                      `xml:"controller" json:"-"`
	Type         string                        `xml:"type,attr" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Index        *int32                        `xml:"index,attr" json:"index,omitempty" protobuf:"varint,2,opt,name=index"`
	Model        string                        `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,3,opt,name=model"`
	Driver       *DomainControllerDriver       `xml:"driver" json:"driver,omitempty" protobuf:"bytes,4,opt,name=driver"`
	PCI          *DomainControllerPCI          `xml:"-" json:"pci,omitempty" protobuf:"bytes,5,opt,name=pci"`
	USB          *DomainControllerUSB          `xml:"-" json:"usb,omitempty" protobuf:"bytes,6,opt,name=usb"`
	VirtIOSerial *DomainControllerVirtIOSerial `xml:"-" json:"virtioSerial,omitempty" protobuf:"bytes,7,opt,name=virtioSerial"`
	XenBus       *DomainControllerXenBus       `xml:"-" json:"xenBus,omitempty" protobuf:"bytes,8,opt,name=xenBus"`
	ACPI         *DomainDeviceACPI             `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,9,opt,name=acpi"`
	Alias        *DomainAlias                  `xml:"alias" json:"alias,omitempty" protobuf:"bytes,10,opt,name=alias"`
	Address      *DomainAddress                `xml:"address" json:"address,omitempty" protobuf:"bytes,11,opt,name=address"`
}

// +gogo:genproto=true
type DomainDiskSecret struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty" protobuf:"bytes,2,opt,name=usage"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty" protobuf:"bytes,3,opt,name=uuid"`
}

// +gogo:genproto=true
type DomainDiskAuth struct {
	Username string            `xml:"username,attr,omitempty" json:"username,omitempty" protobuf:"bytes,1,opt,name=username"`
	Secret   *DomainDiskSecret `xml:"secret" json:"secret,omitempty" protobuf:"bytes,2,opt,name=secret"`
}

// +gogo:genproto=true
type DomainDiskSourceHost struct {
	Transport string `xml:"transport,attr,omitempty" json:"transport,omitempty" protobuf:"bytes,1,opt,name=transport"`
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	Port      string `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"bytes,3,opt,name=port"`
	Socket    string `xml:"socket,attr,omitempty" json:"socket,omitempty" protobuf:"bytes,4,opt,name=socket"`
}

// +gogo:genproto=true
type DomainDiskSourceSSL struct {
	Verify string `xml:"verify,attr" json:"verify" protobuf:"bytes,1,opt,name=verify"`
}

// +gogo:genproto=true
type DomainDiskCookie struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:",chardata" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainDiskCookies struct {
	Cookies []DomainDiskCookie `xml:"cookie" json:"cookies" protobuf:"bytes,1,rep,name=cookies"`
}

// +gogo:genproto=true
type DomainDiskSourceReadahead struct {
	Size_ string `xml:"size,attr" json:"size" protobuf:"bytes,1,opt,name=size"`
}

// +gogo:genproto=true
type DomainDiskSourceTimeout struct {
	Seconds string `xml:"seconds,attr" json:"seconds" protobuf:"bytes,1,opt,name=seconds"`
}

// +gogo:genproto=true
type DomainDiskReservationsSource DomainChardevSource

// +gogo:genproto=true
type DomainDiskReservations struct {
	Enabled string                        `xml:"enabled,attr,omitempty" json:"enabled,omitempty" protobuf:"bytes,1,opt,name=enabled"`
	Managed string                        `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,2,opt,name=managed"`
	Source  *DomainDiskReservationsSource `xml:"source" json:"source" protobuf:"bytes,3,opt,name=source"`
}

// +gogo:genproto=true
type DomainDiskSource struct {
	File          *DomainDiskSourceFile      `xml:"-" json:"file,omitempty" protobuf:"bytes,1,opt,name=file"`
	Block         *DomainDiskSourceBlock     `xml:"-" json:"block,omitempty" protobuf:"bytes,2,opt,name=block"`
	Dir           *DomainDiskSourceDir       `xml:"-" json:"dir,omitempty" protobuf:"bytes,3,opt,name=dir"`
	Network       *DomainDiskSourceNetwork   `xml:"-" json:"network,omitempty" protobuf:"bytes,4,opt,name=network"`
	Volume        *DomainDiskSourceVolume    `xml:"-" json:"volume,omitempty" protobuf:"bytes,5,opt,name=volume"`
	NVME          *DomainDiskSourceNVME      `xml:"-" json:"nvme,omitempty" protobuf:"bytes,6,opt,name=nvme"`
	VHostUser     *DomainDiskSourceVHostUser `xml:"-" json:"vHostUser,omitempty" protobuf:"bytes,7,opt,name=vHostUser"`
	StartupPolicy string                     `xml:"startupPolicy,attr,omitempty" json:"startupPolicy,omitempty" protobuf:"bytes,8,opt,name=startupPolicy"`
	Index         int32                      `xml:"index,attr,omitempty" json:"index,omitempty" protobuf:"varint,9,opt,name=index"`
	Encryption    *DomainDiskEncryption      `xml:"encryption" json:"encryption,omitempty" protobuf:"bytes,10,opt,name=encryption"`
	Reservations  *DomainDiskReservations    `xml:"reservations" json:"reservations,omitempty" protobuf:"bytes,11,opt,name=reservations"`
	Slices        *DomainDiskSlices          `xml:"slices" json:"slices,omitempty" protobuf:"bytes,12,opt,name=slices"`
	SSL           *DomainDiskSourceSSL       `xml:"ssl" json:"ssl,omitempty" protobuf:"bytes,13,opt,name=ssl"`
	Cookies       *DomainDiskCookies         `xml:"cookies" json:"cookies,omitempty" protobuf:"bytes,14,opt,name=cookies"`
	Readahead     *DomainDiskSourceReadahead `xml:"readahead" json:"readahead,omitempty" protobuf:"bytes,15,opt,name=readahead"`
	Timeout       *DomainDiskSourceTimeout   `xml:"timeout" json:"timeout,omitempty" protobuf:"bytes,16,opt,name=timeout"`
}

// +gogo:genproto=true
type DomainDiskSlices struct {
	Slices []DomainDiskSlice `xml:"slice" json:"slices" protobuf:"bytes,1,rep,name=slices"`
}

// +gogo:genproto=true
type DomainDiskSlice struct {
	Type   string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Offset int32  `xml:"offset,attr" json:"offset" protobuf:"varint,2,opt,name=offset"`
	Size_  int32  `xml:"size,attr" json:"size" protobuf:"varint,3,opt,name=size"`
}

// +gogo:genproto=true
type DomainDiskSourceFile struct {
	File     string                 `xml:"file,attr,omitempty" json:"file,omitempty" protobuf:"bytes,1,opt,name=file"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" protobuf:"bytes,2,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainDiskSourceNVME struct {
	PCI *DomainDiskSourceNVMEPCI `json:"pci,omitempty" protobuf:"bytes,1,opt,name=pci"`
}

// +gogo:genproto=true
type DomainDiskSourceNVMEPCI struct {
	Managed   string            `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,1,opt,name=managed"`
	Namespace int64             `xml:"namespace,attr,omitempty" json:"namespace,omitempty" protobuf:"varint,2,opt,name=namespace"`
	Address   *DomainAddressPCI `xml:"address" json:"address,omitempty" protobuf:"bytes,3,opt,name=address"`
}

// +gogo:genproto=true
type DomainDiskSourceBlock struct {
	Dev      string                 `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,1,opt,name=dev"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" protobuf:"bytes,2,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainDiskSourceDir struct {
	Dir string `xml:"dir,attr,omitempty" json:"dir,omitempty" protobuf:"bytes,1,opt,name=dir"`
}

// +gogo:genproto=true
type DomainDiskSourceNetwork struct {
	Protocol  string                            `xml:"protocol,attr,omitempty" json:"protocol,omitempty" protobuf:"bytes,1,opt,name=protocol"`
	Name      string                            `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	Query     string                            `xml:"query,attr,omitempty" json:"query,omitempty" protobuf:"bytes,3,opt,name=query"`
	TLS       string                            `xml:"tls,attr,omitempty" json:"tls,omitempty" protobuf:"bytes,4,opt,name=tls"`
	Hosts     []DomainDiskSourceHost            `xml:"host" json:"hosts" protobuf:"bytes,5,rep,name=hosts"`
	Identity  *DomainDiskSourceNetworkIdentity  `xml:"identity" json:"identity,omitempty" protobuf:"bytes,6,opt,name=identity"`
	Initiator *DomainDiskSourceNetworkInitiator `xml:"initiator" json:"initiator,omitempty" protobuf:"bytes,7,opt,name=initiator"`
	Snapshot  *DomainDiskSourceNetworkSnapshot  `xml:"snapshot" json:"snapshot,omitempty" protobuf:"bytes,8,opt,name=snapshot"`
	Config    *DomainDiskSourceNetworkConfig    `xml:"config" json:"config,omitempty" protobuf:"bytes,9,opt,name=config"`
	Auth      *DomainDiskAuth                   `xml:"auth" json:"auth,omitempty" protobuf:"bytes,10,opt,name=auth"`
}

// +gogo:genproto=true
type DomainDiskSourceNetworkIdentity struct {
	User  string `xml:"user,attr" json:"user" protobuf:"bytes,1,opt,name=user"`
	Group string `xml:"group,attr" json:"group" protobuf:"bytes,2,opt,name=group"`
}

// +gogo:genproto=true
type DomainDiskSourceNetworkInitiator struct {
	IQN *DomainDiskSourceNetworkIQN `xml:"iqn" json:"iqn,omitempty" protobuf:"bytes,1,opt,name=iqn"`
}

// +gogo:genproto=true
type DomainDiskSourceNetworkIQN struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainDiskSourceNetworkSnapshot struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainDiskSourceNetworkConfig struct {
	File string `xml:"file,attr" json:"file" protobuf:"bytes,1,opt,name=file"`
}

// +gogo:genproto=true
type DomainDiskSourceVolume struct {
	Pool     string                 `xml:"pool,attr,omitempty" json:"pool,omitempty" protobuf:"bytes,1,opt,name=pool"`
	Volume   string                 `xml:"volume,attr,omitempty" json:"volume,omitempty" protobuf:"bytes,2,opt,name=volume"`
	Mode     string                 `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,3,opt,name=mode"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"sec_label" protobuf:"bytes,4,rep,name=sec_label,json=secLabel"`
}

// +gogo:genproto=true
type DomainDiskSourceVHostUser DomainChardevSource

// +gogo:genproto=true
type DomainDiskMetadataCache struct {
	MaxSize *DomainDiskMetadataCacheSize `xml:"max_size" json:"maxSize" protobuf:"bytes,1,opt,name=maxSize"`
}

// +gogo:genproto=true
type DomainDiskMetadataCacheSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,1,opt,name=unit"`
	Value int    `xml:",cdata" json:"value" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainDiskDriver struct {
	Name          string                   `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Type          string                   `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
	Cache         string                   `xml:"cache,attr,omitempty" json:"cache,omitempty" protobuf:"bytes,3,opt,name=cache"`
	ErrorPolicy   string                   `xml:"error_policy,attr,omitempty" json:"errorPolicy,omitempty" protobuf:"bytes,4,opt,name=errorPolicy"`
	RErrorPolicy  string                   `xml:"rerror_policy,attr,omitempty" json:"rerrorPolicy,omitempty" protobuf:"bytes,5,opt,name=rerrorPolicy"`
	IO            string                   `xml:"io,attr,omitempty" json:"io,omitempty" protobuf:"bytes,6,opt,name=io"`
	IOEventFD     string                   `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" protobuf:"bytes,7,opt,name=ioeventfd"`
	EventIDX      string                   `xml:"event_idx,attr,omitempty" json:"eventIdx,omitempty" protobuf:"bytes,8,opt,name=eventIdx"`
	CopyOnRead    string                   `xml:"copy_on_read,attr,omitempty" json:"copyOnRead,omitempty" protobuf:"bytes,9,opt,name=copyOnRead"`
	Discard       string                   `xml:"discard,attr,omitempty" json:"discard,omitempty" protobuf:"bytes,10,opt,name=discard"`
	IOThread      *int32                   `xml:"iothread,attr" json:"iothread,omitempty" protobuf:"varint,11,opt,name=iothread"`
	DetectZeros   string                   `xml:"detect_zeroes,attr,omitempty" json:"detectZeros,omitempty" protobuf:"bytes,12,opt,name=detectZeros"`
	Queues        *int32                   `xml:"queues,attr" json:"queues,omitempty" protobuf:"varint,13,opt,name=queues"`
	IOMMU         string                   `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,14,opt,name=iommu"`
	ATS           string                   `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,15,opt,name=ats"`
	Packed        string                   `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,16,opt,name=packed"`
	MetadataCache *DomainDiskMetadataCache `xml:"metadata_cache" json:"metadata_cache,omitempty" protobuf:"bytes,17,opt,name=metadata_cache,json=metadataCache"`
}

// +gogo:genproto=true
type DomainDiskTarget struct {
	Dev          string `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,1,opt,name=dev"`
	Bus          string `xml:"bus,attr,omitempty" json:"bus,omitempty" protobuf:"bytes,2,opt,name=bus"`
	Tray         string `xml:"tray,attr,omitempty" json:"tray,omitempty" protobuf:"bytes,3,opt,name=tray"`
	Removable    string `xml:"removable,attr,omitempty" json:"removable,omitempty" protobuf:"bytes,4,opt,name=removable"`
	RotationRate int32  `xml:"rotation_rate,attr,omitempty" json:"rotationRate,omitempty" protobuf:"varint,5,opt,name=rotationRate"`
}

// +gogo:genproto=true
type DomainDiskEncryption struct {
	Format string            `xml:"format,attr,omitempty" json:"format,omitempty" protobuf:"bytes,1,opt,name=format"`
	Secret *DomainDiskSecret `xml:"secret" json:"secret,omitempty" protobuf:"bytes,2,opt,name=secret"`
}

// +gogo:genproto=true
type DomainDiskReadOnly struct {
}

// +gogo:genproto=true
type DomainDiskShareable struct {
}

// +gogo:genproto=true
type DomainDiskTransient struct {
}

// +gogo:genproto=true
type DomainDiskIOTune struct {
	TotalBytesSec          int64  `xml:"total_bytes_sec,omitempty" json:"totalBytesSec,omitempty" protobuf:"varint,1,opt,name=totalBytesSec"`
	ReadBytesSec           int64  `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty" protobuf:"varint,2,opt,name=readBytesSec"`
	WriteBytesSec          int64  `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty" protobuf:"varint,3,opt,name=writeBytesSec"`
	TotalIopsSec           int64  `xml:"total_iops_sec,omitempty" json:"totalIopsSec,omitempty" protobuf:"varint,4,opt,name=totalIopsSec"`
	ReadIopsSec            int64  `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty" protobuf:"varint,5,opt,name=readIopsSec"`
	WriteIopsSec           int64  `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty" protobuf:"varint,6,opt,name=writeIopsSec"`
	TotalBytesSecMax       int64  `xml:"total_bytes_sec_max,omitempty" json:"totalBytesSecMax,omitempty" protobuf:"varint,7,opt,name=totalBytesSecMax"`
	ReadBytesSecMax        int64  `xml:"read_bytes_sec_max,omitempty" json:"readBytesSecMax,omitempty" protobuf:"varint,8,opt,name=readBytesSecMax"`
	WriteBytesSecMax       int64  `xml:"write_bytes_sec_max,omitempty" json:"writeBytesSecMax,omitempty" protobuf:"varint,9,opt,name=writeBytesSecMax"`
	TotalIopsSecMax        int64  `xml:"total_iops_sec_max,omitempty" json:"totalIopsSecMax,omitempty" protobuf:"varint,10,opt,name=totalIopsSecMax"`
	ReadIopsSecMax         int64  `xml:"read_iops_sec_max,omitempty" json:"readIopsSecMax,omitempty" protobuf:"varint,11,opt,name=readIopsSecMax"`
	WriteIopsSecMax        int64  `xml:"write_iops_sec_max,omitempty" json:"writeIopsSecMax,omitempty" protobuf:"varint,12,opt,name=writeIopsSecMax"`
	TotalBytesSecMaxLength int64  `xml:"total_bytes_sec_max_length,omitempty" json:"totalBytesSecMaxLength,omitempty" protobuf:"varint,13,opt,name=totalBytesSecMaxLength"`
	ReadBytesSecMaxLength  int64  `xml:"read_bytes_sec_max_length,omitempty" json:"readBytesSecMaxLength,omitempty" protobuf:"varint,14,opt,name=readBytesSecMaxLength"`
	WriteBytesSecMaxLength int64  `xml:"write_bytes_sec_max_length,omitempty" json:"writeBytesSecMaxLength,omitempty" protobuf:"varint,15,opt,name=writeBytesSecMaxLength"`
	TotalIopsSecMaxLength  int64  `xml:"total_iops_sec_max_length,omitempty" json:"totalIopsSecMaxLength,omitempty" protobuf:"varint,16,opt,name=totalIopsSecMaxLength"`
	ReadIopsSecMaxLength   int64  `xml:"read_iops_sec_max_length,omitempty" json:"readIopsSecMaxLength,omitempty" protobuf:"varint,17,opt,name=readIopsSecMaxLength"`
	WriteIopsSecMaxLength  int64  `xml:"write_iops_sec_max_length,omitempty" json:"writeIopsSecMaxLength,omitempty" protobuf:"varint,18,opt,name=writeIopsSecMaxLength"`
	SizeIopsSec            int64  `xml:"size_iops_sec,omitempty" json:"sizeIopsSec,omitempty" protobuf:"varint,19,opt,name=sizeIopsSec"`
	GroupName              string `xml:"group_name,omitempty" json:"groupName,omitempty" protobuf:"bytes,20,opt,name=groupName"`
}

// +gogo:genproto=true
type DomainDiskGeometry struct {
	Cylinders int32  `xml:"cyls,attr" json:"cyli" protobuf:"varint,1,opt,name=cyli"`
	Headers   int32  `xml:"heads,attr" json:"heads" protobuf:"varint,2,opt,name=heads"`
	Sectors   int32  `xml:"secs,attr" json:"secs" protobuf:"varint,3,opt,name=secs"`
	Trans     string `xml:"trans,attr,omitempty" json:"trans,omitempty" protobuf:"bytes,4,opt,name=trans"`
}

// +gogo:genproto=true
type DomainDiskBlockIO struct {
	LogicalBlockSize  int32 `xml:"logical_block_size,attr,omitempty" json:"logicalBlockSize,omitempty" protobuf:"varint,1,opt,name=logicalBlockSize"`
	PhysicalBlockSize int32 `xml:"physical_block_size,attr,omitempty" json:"physicalBlockSize,omitempty" protobuf:"varint,2,opt,name=physicalBlockSize"`
}

// +gogo:genproto=true
type DomainDiskFormat struct {
	Type          string                   `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	MetadataCache *DomainDiskMetadataCache `xml:"metadata_cache" json:"metadataCache,omitempty" protobuf:"bytes,2,opt,name=metadataCache"`
}

// +gogo:genproto=true
type DomainDiskBackingStore struct {
	Index        int32                   `xml:"index,attr,omitempty" json:"index,omitempty" protobuf:"varint,1,opt,name=index"`
	Format       *DomainDiskFormat       `xml:"format" json:"format,omitempty" protobuf:"bytes,2,opt,name=format"`
	Source       *DomainDiskSource       `xml:"source" json:"source,omitempty" protobuf:"bytes,3,opt,name=source"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty" protobuf:"bytes,4,opt,name=backingStore"`
}

// +gogo:genproto=true
type DomainDiskMirror struct {
	Job          string                  `xml:"job,attr,omitempty" json:"job,omitempty" protobuf:"bytes,1,opt,name=job"`
	Ready        string                  `xml:"ready,attr,omitempty" json:"ready,omitempty" protobuf:"bytes,2,opt,name=ready"`
	Format       *DomainDiskFormat       `xml:"format" json:"format,omitempty" protobuf:"bytes,3,opt,name=format"`
	Source       *DomainDiskSource       `xml:"source" json:"source,omitempty" protobuf:"bytes,4,opt,name=source"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty" protobuf:"bytes,5,opt,name=backingStore"`
}

// +gogo:genproto=true
type DomainBackendDomain struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainDisk struct {
	XMLName       xml.Name                `xml:"disk" json:"-"`
	Device        string                  `xml:"device,attr,omitempty" json:"device,omitempty" protobuf:"bytes,1,opt,name=device"`
	RawIO         string                  `xml:"rawio,attr,omitempty" json:"rawio,omitempty" protobuf:"bytes,2,opt,name=rawio"`
	SGIO          string                  `xml:"sgio,attr,omitempty" json:"sgio,omitempty" protobuf:"bytes,3,opt,name=sgio"`
	Snapshot      string                  `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty" protobuf:"bytes,4,opt,name=snapshot"`
	Model         string                  `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,5,opt,name=model"`
	Driver        *DomainDiskDriver       `xml:"driver" json:"driver,omitempty" protobuf:"bytes,6,opt,name=driver"`
	Auth          *DomainDiskAuth         `xml:"auth" json:"auth,omitempty" protobuf:"bytes,7,opt,name=auth"`
	Source        *DomainDiskSource       `xml:"source" json:"source,omitempty" protobuf:"bytes,8,opt,name=source"`
	BackingStore  *DomainDiskBackingStore `xml:"backingStore" json:"backingStore,omitempty" protobuf:"bytes,9,opt,name=backingStore"`
	BackendDomain *DomainBackendDomain    `xml:"backenddomain" json:"backendDomain,omitempty" protobuf:"bytes,10,opt,name=backendDomain"`
	Geometry      *DomainDiskGeometry     `xml:"geometry" json:"geometry,omitempty" protobuf:"bytes,11,opt,name=geometry"`
	BlockIO       *DomainDiskBlockIO      `xml:"blockio" json:"blockio,omitempty" protobuf:"bytes,12,opt,name=blockio"`
	Mirror        *DomainDiskMirror       `xml:"mirror" json:"mirror,omitempty" protobuf:"bytes,13,opt,name=mirror"`
	Target        *DomainDiskTarget       `xml:"target" json:"target,omitempty" protobuf:"bytes,14,opt,name=target"`
	IOTune        *DomainDiskIOTune       `xml:"iotune" json:"iotune,omitempty" protobuf:"bytes,15,opt,name=iotune"`
	ReadOnly      *DomainDiskReadOnly     `xml:"readonly" json:"readonly,omitempty" protobuf:"bytes,16,opt,name=readonly"`
	Shareable     *DomainDiskShareable    `xml:"shareable" json:"shareable,omitempty" protobuf:"bytes,17,opt,name=shareable"`
	Transient     *DomainDiskTransient    `xml:"transient" json:"transient,omitempty" protobuf:"bytes,18,opt,name=transient"`
	Serial        string                  `xml:"serial,omitempty" json:"serial,omitempty" protobuf:"bytes,19,opt,name=serial"`
	WWN           string                  `xml:"wwn,omitempty" json:"wwn,omitempty" protobuf:"bytes,20,opt,name=wwn"`
	Vendor        string                  `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,21,opt,name=vendor"`
	Product       string                  `xml:"product,omitempty" json:"product,omitempty" protobuf:"bytes,22,opt,name=product"`
	Encryption    *DomainDiskEncryption   `xml:"encryption" json:"encryption,omitempty" protobuf:"bytes,23,opt,name=encryption"`
	Boot          *DomainDeviceBoot       `xml:"boot" json:"boot,omitempty" protobuf:"bytes,24,opt,name=boot"`
	ACPI          *DomainDeviceACPI       `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,25,opt,name=acpi"`
	Alias         *DomainAlias            `xml:"alias" json:"alias,omitempty" protobuf:"bytes,26,opt,name=alias"`
	Address       *DomainAddress          `xml:"address" json:"address,omitempty" protobuf:"bytes,27,opt,name=address"`
}

// +gogo:genproto=true
type DomainFilesystemDriver struct {
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Format   string `xml:"format,attr,omitempty" json:"format,omitempty" protobuf:"bytes,2,opt,name=format"`
	Name     string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	WRPolicy string `xml:"wrpolicy,attr,omitempty" json:"wrpolicy,omitempty" protobuf:"bytes,4,opt,name=wrpolicy"`
	IOMMU    string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,5,opt,name=iommu"`
	ATS      string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,6,opt,name=ats"`
	Packed   string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,7,opt,name=packed"`
	Queue    int32  `xml:"queue,attr,omitempty" json:"queue,omitempty" protobuf:"varint,8,opt,name=queue"`
}

// +gogo:genproto=true
type DomainFilesystemSource struct {
	Mount    *DomainFilesystemSourceMount    `xml:"-" json:"mount,omitempty" protobuf:"bytes,1,opt,name=mount"`
	Block    *DomainFilesystemSourceBlock    `xml:"-" json:"block,omitempty" protobuf:"bytes,2,opt,name=block"`
	File     *DomainFilesystemSourceFile     `xml:"-" json:"file,omitempty" protobuf:"bytes,3,opt,name=file"`
	Template *DomainFilesystemSourceTemplate `xml:"-" json:"template,omitempty" protobuf:"bytes,4,opt,name=template"`
	RAM      *DomainFilesystemSourceRAM      `xml:"-" json:"ram,omitempty" protobuf:"bytes,5,opt,name=ram"`
	Bind     *DomainFilesystemSourceBind     `xml:"-" json:"bind,omitempty" protobuf:"bytes,6,opt,name=bind"`
	Volume   *DomainFilesystemSourceVolume   `xml:"-" json:"volume,omitempty" protobuf:"bytes,7,opt,name=volume"`
}

// +gogo:genproto=true
type DomainFilesystemSourceMount struct {
	Dir    string `xml:"dir,attr,omitempty" json:"dir,omitempty" protobuf:"bytes,1,opt,name=dir"`
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty" protobuf:"bytes,2,opt,name=socket"`
}

// +gogo:genproto=true
type DomainFilesystemSourceBlock struct {
	Dev string `xml:"dev,attr" json:"dev" protobuf:"bytes,1,opt,name=dev"`
}

// +gogo:genproto=true
type DomainFilesystemSourceFile struct {
	File string `xml:"file,attr" json:"file" protobuf:"bytes,1,opt,name=file"`
}

// +gogo:genproto=true
type DomainFilesystemSourceTemplate struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainFilesystemSourceRAM struct {
	Usage int32  `xml:"usage,attr" json:"usage" protobuf:"varint,1,opt,name=usage"`
	Units string `xml:"units,attr,omitempty" json:"units,omitempty" protobuf:"bytes,2,opt,name=units"`
}

// +gogo:genproto=true
type DomainFilesystemSourceBind struct {
	Dir string `xml:"dir,attr" json:"dir" protobuf:"bytes,1,opt,name=dir"`
}

// +gogo:genproto=true
type DomainFilesystemSourceVolume struct {
	Pool   string `xml:"pool,attr" json:"pool" protobuf:"bytes,1,opt,name=pool"`
	Volume string `xml:"volume,attr" json:"volume" protobuf:"bytes,2,opt,name=volume"`
}

// +gogo:genproto=true
type DomainFilesystemTarget struct {
	Dir string `xml:"dir,attr" json:"dir" protobuf:"bytes,1,opt,name=dir"`
}

// +gogo:genproto=true
type DomainFilesystemReadOnly struct {
}

// +gogo:genproto=true
type DomainFilesystemSpaceHardLimit struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainFilesystemSpaceSoftLimit struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainFilesystemBinaryCache struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainFilesystemBinarySandbox struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainFilesystemBinaryLock struct {
	POSIX string `xml:"posix,attr,omitempty" json:"posix,omitempty" protobuf:"bytes,1,opt,name=posix"`
	Flock string `xml:"flock,attr,omitempty" json:"flock,omitempty" protobuf:"bytes,2,opt,name=flock"`
}

// +gogo:genproto=true
type DomainFilesystemBinary struct {
	Path    string                         `xml:"path,attr,omitempty" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	XAttr   string                         `xml:"xattr,attr,omitempty" json:"xattr,omitempty" protobuf:"bytes,2,opt,name=xattr"`
	Cache   *DomainFilesystemBinaryCache   `xml:"cache" json:"cache,omitempty" protobuf:"bytes,3,opt,name=cache"`
	Sandbox *DomainFilesystemBinarySandbox `xml:"sandbox" json:"sandbox,omitempty" protobuf:"bytes,4,opt,name=sandbox"`
	Lock    *DomainFilesystemBinaryLock    `xml:"lock" json:"lock,omitempty" protobuf:"bytes,5,opt,name=lock"`
}

// +gogo:genproto=true
type DomainFilesystem struct {
	XMLName        xml.Name                        `xml:"filesystem" json:"-"`
	AccessMode     string                          `xml:"accessmode,attr,omitempty" json:"accessMode,omitempty" protobuf:"bytes,1,opt,name=accessMode"`
	Model          string                          `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,2,opt,name=model"`
	MultiDevs      string                          `xml:"multidevs,attr,omitempty" json:"multidevs,omitempty" protobuf:"bytes,3,opt,name=multidevs"`
	FMode          string                          `xml:"fmode,attr,omitempty" json:"fmode,omitempty" protobuf:"bytes,4,opt,name=fmode"`
	DMode          string                          `xml:"dmode,attr,omitempty" json:"dmode,omitempty" protobuf:"bytes,5,opt,name=dmode"`
	Driver         *DomainFilesystemDriver         `xml:"driver" json:"driver,omitempty" protobuf:"bytes,6,opt,name=driver"`
	Binary         *DomainFilesystemBinary         `xml:"binary" json:"binary,omitempty" protobuf:"bytes,7,opt,name=binary"`
	Source         *DomainFilesystemSource         `xml:"source" json:"source,omitempty" protobuf:"bytes,8,opt,name=source"`
	Target         *DomainFilesystemTarget         `xml:"target" json:"target,omitempty" protobuf:"bytes,9,opt,name=target"`
	ReadOnly       *DomainFilesystemReadOnly       `xml:"readonly" json:"readonly,omitempty" protobuf:"bytes,10,opt,name=readonly"`
	SpaceHardLimit *DomainFilesystemSpaceHardLimit `xml:"space_hard_limit" json:"spaceHardLimit,omitempty" protobuf:"bytes,11,opt,name=spaceHardLimit"`
	SpaceSoftLimit *DomainFilesystemSpaceSoftLimit `xml:"space_soft_limit" json:"spaceSoftLimit,omitempty" protobuf:"bytes,12,opt,name=spaceSoftLimit"`
	Boot           *DomainDeviceBoot               `xml:"boot" json:"boot,omitempty" protobuf:"bytes,13,opt,name=boot"`
	ACPI           *DomainDeviceACPI               `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,14,opt,name=acpi"`
	Alias          *DomainAlias                    `xml:"alias" json:"alias,omitempty" protobuf:"bytes,15,opt,name=alias"`
	Address        *DomainAddress                  `xml:"address" json:"address,omitempty" protobuf:"bytes,16,opt,name=address"`
}

// +gogo:genproto=true
type DomainInterfaceMAC struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
	Type    string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
	Check   string `xml:"check,attr,omitempty" json:"check,omitempty" protobuf:"bytes,3,opt,name=check"`
}

// +gogo:genproto=true
type DomainInterfaceModel struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type DomainInterfaceSource struct {
	User      *DomainInterfaceSourceUser     `xml:"-" json:"user,omitempty" protobuf:"bytes,1,opt,name=user"`
	Ethernet  *DomainInterfaceSourceEthernet `xml:"-" json:"ethernet,omitempty" protobuf:"bytes,2,opt,name=ethernet"`
	VHostUser *DomainChardevSource           `xml:"-" json:"vHostUser,omitempty" protobuf:"bytes,3,opt,name=vHostUser"`
	Server    *DomainInterfaceSourceServer   `xml:"-" json:"server,omitempty" protobuf:"bytes,4,opt,name=server"`
	Client    *DomainInterfaceSourceClient   `xml:"-" json:"client,omitempty" protobuf:"bytes,5,opt,name=client"`
	MCast     *DomainInterfaceSourceMCast    `xml:"-" json:"mcast,omitempty" protobuf:"bytes,6,opt,name=mcast"`
	Network   *DomainInterfaceSourceNetwork  `xml:"-" json:"network,omitempty" protobuf:"bytes,7,opt,name=network"`
	Bridge    *DomainInterfaceSourceBridge   `xml:"-" json:"bridge,omitempty" protobuf:"bytes,8,opt,name=bridge"`
	Internal  *DomainInterfaceSourceInternal `xml:"-" json:"internal,omitempty" protobuf:"bytes,9,opt,name=internal"`
	Direct    *DomainInterfaceSourceDirect   `xml:"-" json:"direct,omitempty" protobuf:"bytes,10,opt,name=direct"`
	Hostdev   *DomainInterfaceSourceHostdev  `xml:"-" json:"hostdev,omitempty" protobuf:"bytes,11,opt,name=hostdev"`
	UDP       *DomainInterfaceSourceUDP      `xml:"-" json:"udp,omitempty" protobuf:"bytes,12,opt,name=udp"`
	VDPA      *DomainInterfaceSourceVDPA     `xml:"-" json:"vdpa,omitempty" protobuf:"bytes,13,opt,name=vdpa"`
}

// +gogo:genproto=true
type DomainInterfaceSourceUser struct {
}

// +gogo:genproto=true
type DomainInterfaceSourceEthernet struct {
	IP    []DomainInterfaceIP    `xml:"ip" json:"ip" protobuf:"bytes,1,rep,name=ip"`
	Route []DomainInterfaceRoute `xml:"route" json:"route" protobuf:"bytes,2,rep,name=route"`
}

// +gogo:genproto=true
type DomainInterfaceSourceServer struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty" protobuf:"bytes,3,opt,name=local"`
}

// +gogo:genproto=true
type DomainInterfaceSourceClient struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty" protobuf:"bytes,3,opt,name=local"`
}

// +gogo:genproto=true
type DomainInterfaceSourceMCast struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty" protobuf:"bytes,3,opt,name=local"`
}

// +gogo:genproto=true
type DomainInterfaceSourceNetwork struct {
	Network   string `xml:"network,attr,omitempty" json:"network,omitempty" protobuf:"bytes,1,opt,name=network"`
	PortGroup string `xml:"portgroup,attr,omitempty" json:"portGroup,omitempty" protobuf:"bytes,2,opt,name=portGroup"`
	Bridge    string `xml:"bridge,attr,omitempty" json:"bridge,omitempty" protobuf:"bytes,3,opt,name=bridge"`
	PortID    string `xml:"portid,attr,omitempty" json:"portId,omitempty" protobuf:"bytes,4,opt,name=portId"`
}

// +gogo:genproto=true
type DomainInterfaceSourceBridge struct {
	Bridge string `xml:"bridge,attr" json:"bridge" protobuf:"bytes,1,opt,name=bridge"`
}

// +gogo:genproto=true
type DomainInterfaceSourceInternal struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainInterfaceSourceDirect struct {
	Dev  string `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,1,opt,name=dev"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type DomainInterfaceSourceHostdev struct {
	PCI *DomainHostdevSubsysPCISource `xml:"-" json:"pci,omitempty" protobuf:"bytes,1,opt,name=pci"`
	USB *DomainHostdevSubsysUSBSource `xml:"-" json:"usb,omitempty" protobuf:"bytes,2,opt,name=usb"`
}

// +gogo:genproto=true
type DomainInterfaceSourceUDP struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    int32                       `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local,omitempty" protobuf:"bytes,3,opt,name=local"`
}

// +gogo:genproto=true
type DomainInterfaceSourceVDPA struct {
	Device string `xml:"dev,attr,omitempty" json:"device,omitempty" protobuf:"bytes,1,opt,name=device"`
}

// +gogo:genproto=true
type DomainInterfaceSourceLocal struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    int32  `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
}

// +gogo:genproto=true
type DomainInterfaceTarget struct {
	Dev     string `xml:"dev,attr" json:"dev" protobuf:"bytes,1,opt,name=dev"`
	Managed string `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,2,opt,name=managed"`
}

// +gogo:genproto=true
type DomainInterfaceLink struct {
	State string `xml:"state,attr" json:"state" protobuf:"bytes,1,opt,name=state"`
}

// +gogo:genproto=true
type DomainDeviceBoot struct {
	Order    int32  `xml:"order,attr" json:"order" protobuf:"varint,1,opt,name=order"`
	LoadParm string `xml:"loadparm,attr,omitempty" json:"loadParm,omitempty" protobuf:"bytes,2,opt,name=loadParm"`
}

// +gogo:genproto=true
type DomainInterfaceScript struct {
	Path string `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
}

// +gogo:genproto=true
type DomainInterfaceDriver struct {
	Name        string                      `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	TXMode      string                      `xml:"txmode,attr,omitempty" json:"txMode,omitempty" protobuf:"bytes,2,opt,name=txMode"`
	IOEventFD   string                      `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" protobuf:"bytes,3,opt,name=ioeventfd"`
	EventIDX    string                      `xml:"event_idx,attr,omitempty" json:"eventIdx,omitempty" protobuf:"bytes,4,opt,name=eventIdx"`
	Queues      int32                       `xml:"queues,attr,omitempty" json:"queues,omitempty" protobuf:"varint,5,opt,name=queues"`
	RXQueueSize int32                       `xml:"rx_queue_size,attr,omitempty" json:"rxQueueSize,omitempty" protobuf:"varint,6,opt,name=rxQueueSize"`
	TXQueueSize int32                       `xml:"tx_queue_size,attr,omitempty" json:"txQueueSize,omitempty" protobuf:"varint,7,opt,name=txQueueSize"`
	IOMMU       string                      `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,8,opt,name=iommu"`
	ATS         string                      `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,9,opt,name=ats"`
	Packed      string                      `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,10,opt,name=packed"`
	Host        *DomainInterfaceDriverHost  `xml:"host" json:"host,omitempty" protobuf:"bytes,11,opt,name=host"`
	Guest       *DomainInterfaceDriverGuest `xml:"guest" json:"guest,omitempty" protobuf:"bytes,12,opt,name=guest"`
}

// +gogo:genproto=true
type DomainInterfaceDriverHost struct {
	CSum     string `xml:"csum,attr,omitempty" json:"csum,omitempty" protobuf:"bytes,1,opt,name=csum"`
	GSO      string `xml:"gso,attr,omitempty" json:"gso,omitempty" protobuf:"bytes,2,opt,name=gso"`
	TSO4     string `xml:"tso4,attr,omitempty" json:"tso4,omitempty" protobuf:"bytes,3,opt,name=tso4"`
	TSO6     string `xml:"tso6,attr,omitempty" json:"tso6,omitempty" protobuf:"bytes,4,opt,name=tso6"`
	ECN      string `xml:"ecn,attr,omitempty" json:"ecn,omitempty" protobuf:"bytes,5,opt,name=ecn"`
	UFO      string `xml:"ufo,attr,omitempty" json:"ufo,omitempty" protobuf:"bytes,6,opt,name=ufo"`
	MrgRXBuf string `xml:"mrg_rxbuf,attr,omitempty" json:"mrgRxBuf,omitempty" protobuf:"bytes,7,opt,name=mrgRxBuf"`
}

// +gogo:genproto=true
type DomainInterfaceDriverGuest struct {
	CSum string `xml:"csum,attr,omitempty" json:"csum,omitempty" protobuf:"bytes,1,opt,name=csum"`
	TSO4 string `xml:"tso4,attr,omitempty" json:"tso4,omitempty" protobuf:"bytes,2,opt,name=tso4"`
	TSO6 string `xml:"tso6,attr,omitempty" json:"tso6,omitempty" protobuf:"bytes,3,opt,name=tso6"`
	ECN  string `xml:"ecn,attr,omitempty" json:"ecn,omitempty" protobuf:"bytes,4,opt,name=ecn"`
	UFO  string `xml:"ufo,attr,omitempty" json:"ufo,omitempty" protobuf:"bytes,5,opt,name=ufo"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPort struct {
	Params *DomainInterfaceVirtualPortParams `xml:"parameters" json:"parameters,omitempty" protobuf:"bytes,1,opt,name=parameters"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParams struct {
	Any          *DomainInterfaceVirtualPortParamsAny          `xml:"-" json:"any,omitempty" protobuf:"bytes,1,opt,name=any"`
	VEPA8021QBG  *DomainInterfaceVirtualPortParamsVEPA8021QBG  `xml:"-" json:"vepa8021_qbg,omitempty" protobuf:"bytes,2,opt,name=vepa8021_qbg,json=vepa8021Qbg"`
	VNTag8011QBH *DomainInterfaceVirtualPortParamsVNTag8021QBH `xml:"-" json:"vntag8011_qbh,omitempty" protobuf:"bytes,3,opt,name=vntag8011_qbh,json=vntag8011Qbh"`
	OpenVSwitch  *DomainInterfaceVirtualPortParamsOpenVSwitch  `xml:"-" json:"openvswitch,omitempty" protobuf:"bytes,4,opt,name=openvswitch"`
	MidoNet      *DomainInterfaceVirtualPortParamsMidoNet      `xml:"-" json:"midonet,omitempty" protobuf:"bytes,5,opt,name=midonet"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParamsAny struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty" protobuf:"varint,1,opt,name=managerId"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty" protobuf:"varint,2,opt,name=typeId"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty" protobuf:"varint,3,opt,name=typeIdVersion"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty" protobuf:"bytes,4,opt,name=instanceId"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,5,opt,name=profileId"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,6,opt,name=interfaceId"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty" protobuf:"varint,1,opt,name=managerId"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty" protobuf:"varint,2,opt,name=typeId"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty" protobuf:"varint,3,opt,name=typeIdVersion"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty" protobuf:"bytes,4,opt,name=instanceId"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,1,opt,name=profileId"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,1,opt,name=interfaceId"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,2,opt,name=profileId"`
}

// +gogo:genproto=true
type DomainInterfaceVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,1,opt,name=interfaceId"`
}

// +gogo:genproto=true
type DomainInterfaceBandwidthParams struct {
	Average *int32 `xml:"average,attr" json:"average,omitempty" protobuf:"varint,1,opt,name=average"`
	Peak    *int32 `xml:"peak,attr" json:"peak,omitempty" protobuf:"varint,2,opt,name=peak"`
	Burst   *int32 `xml:"burst,attr" json:"burst,omitempty" protobuf:"varint,3,opt,name=burst"`
	Floor   *int32 `xml:"floor,attr" json:"floor,omitempty" protobuf:"varint,4,opt,name=floor"`
}

// +gogo:genproto=true
type DomainInterfaceBandwidth struct {
	Inbound  *DomainInterfaceBandwidthParams `xml:"inbound" json:"inbound,omitempty" protobuf:"bytes,1,opt,name=inbound"`
	Outbound *DomainInterfaceBandwidthParams `xml:"outbound" json:"outbound,omitempty" protobuf:"bytes,2,opt,name=outbound"`
}

// +gogo:genproto=true
type DomainInterfaceVLan struct {
	Trunk string                   `xml:"trunk,attr,omitempty" json:"trunk,omitempty" protobuf:"bytes,1,opt,name=trunk"`
	Tags  []DomainInterfaceVLanTag `xml:"tag" json:"tags" protobuf:"bytes,2,rep,name=tags"`
}

// +gogo:genproto=true
type DomainInterfaceVLanTag struct {
	ID         int32  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode" protobuf:"bytes,2,opt,name=nativeMode"`
}

// +gogo:genproto=true
type DomainInterfaceGuest struct {
	Dev    string `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,1,opt,name=dev"`
	Actual string `xml:"actual,attr,omitempty" json:"actual,omitempty" protobuf:"bytes,2,opt,name=actual"`
}

// +gogo:genproto=true
type DomainInterfaceFilterRef struct {
	Filter     string                       `xml:"filter,attr" json:"filter" protobuf:"bytes,1,opt,name=filter"`
	Parameters []DomainInterfaceFilterParam `xml:"parameter" json:"parameter" protobuf:"bytes,2,rep,name=parameter"`
}

// +gogo:genproto=true
type DomainInterfaceFilterParam struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainInterfaceBackend struct {
	Tap   string `xml:"tap,attr,omitempty" json:"tap,omitempty" protobuf:"bytes,1,opt,name=tap"`
	VHost string `xml:"vhost,attr,omitempty" json:"vhost,omitempty" protobuf:"bytes,2,opt,name=vhost"`
}

// +gogo:genproto=true
type DomainInterfaceTune struct {
	SndBuf int32 `xml:"sndbuf" json:"sndbuf" protobuf:"varint,1,opt,name=sndbuf"`
}

// +gogo:genproto=true
type DomainInterfaceMTU struct {
	Size_ int32 `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
}

// +gogo:genproto=true
type DomainInterfaceCoalesce struct {
	RX *DomainInterfaceCoalesceRX `xml:"rx" json:"rx,omitempty" protobuf:"bytes,1,opt,name=rx"`
}

// +gogo:genproto=true
type DomainInterfaceCoalesceRX struct {
	Frames *DomainInterfaceCoalesceRXFrames `xml:"frames" json:"frames,omitempty" protobuf:"bytes,1,opt,name=frames"`
}

// +gogo:genproto=true
type DomainInterfaceCoalesceRXFrames struct {
	Max *int32 `xml:"max,attr" json:"max,omitempty" protobuf:"varint,1,opt,name=max"`
}

// +gogo:genproto=true
type DomainROM struct {
	Bar     string `xml:"bar,attr,omitempty" json:"bar,omitempty" protobuf:"bytes,1,opt,name=bar"`
	File    string `xml:"file,attr,omitempty" json:"file,omitempty" protobuf:"bytes,2,opt,name=file"`
	Enabled string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" protobuf:"bytes,3,opt,name=enabled"`
}

// +gogo:genproto=true
type DomainInterfaceIP struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,2,opt,name=family"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty" protobuf:"varint,3,opt,name=prefix"`
	Peer    string `xml:"peer,attr,omitempty" json:"peer,omitempty" protobuf:"bytes,4,opt,name=peer"`
}

// +gogo:genproto=true
type DomainInterfaceRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,1,opt,name=family"`
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,2,opt,name=address"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty" protobuf:"bytes,3,opt,name=netmask"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty" protobuf:"varint,4,opt,name=prefix"`
	Gateway string `xml:"gateway,attr" json:"gateway" protobuf:"bytes,5,opt,name=gateway"`
	Metric  int32  `xml:"metric,attr,omitempty" json:"metric,omitempty" protobuf:"varint,6,opt,name=metric"`
}

// +gogo:genproto=true
type DomainInterfaceTeaming struct {
	Type       string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Persistent string `xml:"persistent,attr,omitempty" json:"persistent,omitempty" protobuf:"bytes,2,opt,name=persistent"`
}

// +gogo:genproto=true
type DomainInterfacePortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty" protobuf:"bytes,1,opt,name=isolated"`
}

// +gogo:genproto=true
type DomainInterface struct {
	XMLName             xml.Name                    `xml:"interface" json:"-"`
	Managed             string                      `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,1,opt,name=managed"`
	TrustGuestRXFilters string                      `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty" protobuf:"bytes,2,opt,name=trustGuestRxFilters"`
	MAC                 *DomainInterfaceMAC         `xml:"mac" json:"mac,omitempty" protobuf:"bytes,3,opt,name=mac"`
	Source              *DomainInterfaceSource      `xml:"source" json:"source,omitempty" protobuf:"bytes,4,opt,name=source"`
	Boot                *DomainDeviceBoot           `xml:"boot" json:"boot,omitempty" protobuf:"bytes,5,opt,name=boot"`
	VLan                *DomainInterfaceVLan        `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,6,opt,name=vlan"`
	VirtualPort         *DomainInterfaceVirtualPort `xml:"virtualport" json:"virtualPort,omitempty" protobuf:"bytes,7,opt,name=virtualPort"`
	IP                  []DomainInterfaceIP         `xml:"ip" json:"ip" protobuf:"bytes,8,rep,name=ip"`
	Route               []DomainInterfaceRoute      `xml:"route" json:"route" protobuf:"bytes,9,rep,name=route"`
	Script              *DomainInterfaceScript      `xml:"script" json:"script,omitempty" protobuf:"bytes,10,opt,name=script"`
	DownScript          *DomainInterfaceScript      `xml:"downscript" json:"downscript,omitempty" protobuf:"bytes,11,opt,name=downscript"`
	BackendDomain       *DomainBackendDomain        `xml:"backenddomain" json:"backenddomain,omitempty" protobuf:"bytes,12,opt,name=backenddomain"`
	Target              *DomainInterfaceTarget      `xml:"target" json:"target,omitempty" protobuf:"bytes,13,opt,name=target"`
	Guest               *DomainInterfaceGuest       `xml:"guest" json:"guest,omitempty" protobuf:"bytes,14,opt,name=guest"`
	Model               *DomainInterfaceModel       `xml:"model" json:"model,omitempty" protobuf:"bytes,15,opt,name=model"`
	Driver              *DomainInterfaceDriver      `xml:"driver" json:"driver,omitempty" protobuf:"bytes,16,opt,name=driver"`
	Backend             *DomainInterfaceBackend     `xml:"backend" json:"backend,omitempty" protobuf:"bytes,17,opt,name=backend"`
	FilterRef           *DomainInterfaceFilterRef   `xml:"filterref" json:"filterref,omitempty" protobuf:"bytes,18,opt,name=filterref"`
	Tune                *DomainInterfaceTune        `xml:"tune" json:"tune,omitempty" protobuf:"bytes,19,opt,name=tune"`
	Teaming             *DomainInterfaceTeaming     `xml:"teaming" json:"teaming,omitempty" protobuf:"bytes,20,opt,name=teaming"`
	Link                *DomainInterfaceLink        `xml:"link" json:"link,omitempty" protobuf:"bytes,21,opt,name=link"`
	MTU                 *DomainInterfaceMTU         `xml:"mtu" json:"mtu,omitempty" protobuf:"bytes,22,opt,name=mtu"`
	Bandwidth           *DomainInterfaceBandwidth   `xml:"bandwidth" json:"bandwidth,omitempty" protobuf:"bytes,23,opt,name=bandwidth"`
	PortOptions         *DomainInterfacePortOptions `xml:"port" json:"portOptions,omitempty" protobuf:"bytes,24,opt,name=portOptions"`
	Coalesce            *DomainInterfaceCoalesce    `xml:"coalesce" json:"coalesce,omitempty" protobuf:"bytes,25,opt,name=coalesce"`
	ROM                 *DomainROM                  `xml:"rom" json:"rom,omitempty" protobuf:"bytes,26,opt,name=rom"`
	ACPI                *DomainDeviceACPI           `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,27,opt,name=acpi"`
	Alias               *DomainAlias                `xml:"alias" json:"alias,omitempty" protobuf:"bytes,28,opt,name=alias"`
	Address             *DomainAddress              `xml:"address" json:"address,omitempty" protobuf:"bytes,29,opt,name=address"`
}

// +gogo:genproto=true
type DomainChardevSource struct {
	Null      *DomainChardevSourceNull      `xml:"-" json:"null,omitempty" protobuf:"bytes,1,opt,name=null"`
	VC        *DomainChardevSourceVC        `xml:"-" json:"vc,omitempty" protobuf:"bytes,2,opt,name=vc"`
	Pty       *DomainChardevSourcePty       `xml:"-" json:"pty,omitempty" protobuf:"bytes,3,opt,name=pty"`
	Dev       *DomainChardevSourceDev       `xml:"-" json:"dev,omitempty" protobuf:"bytes,4,opt,name=dev"`
	File      *DomainChardevSourceFile      `xml:"-" json:"file,omitempty" protobuf:"bytes,5,opt,name=file"`
	Pipe      *DomainChardevSourcePipe      `xml:"-" json:"pipe,omitempty" protobuf:"bytes,6,opt,name=pipe"`
	StdIO     *DomainChardevSourceStdIO     `xml:"-" json:"stdio,omitempty" protobuf:"bytes,7,opt,name=stdio"`
	UDP       *DomainChardevSourceUDP       `xml:"-" json:"udp,omitempty" protobuf:"bytes,8,opt,name=udp"`
	TCP       *DomainChardevSourceTCP       `xml:"-" json:"tcp,omitempty" protobuf:"bytes,9,opt,name=tcp"`
	UNIX      *DomainChardevSourceUNIX      `xml:"-" json:"unix,omitempty" protobuf:"bytes,10,opt,name=unix"`
	SpiceVMC  *DomainChardevSourceSpiceVMC  `xml:"-" json:"spicevmc,omitempty" protobuf:"bytes,11,opt,name=spicevmc"`
	SpicePort *DomainChardevSourceSpicePort `xml:"-" json:"spiceport,omitempty" protobuf:"bytes,12,opt,name=spiceport"`
	NMDM      *DomainChardevSourceNMDM      `xml:"-" json:"nmdm,omitempty" protobuf:"bytes,13,opt,name=nmdm"`
}

// +gogo:genproto=true
type DomainChardevSourceNull struct {
}

// +gogo:genproto=true
type DomainChardevSourceVC struct {
}

// +gogo:genproto=true
type DomainChardevSourcePty struct {
	Path     string                 `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" protobuf:"bytes,2,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainChardevSourceDev struct {
	Path     string                 `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"sec_label" protobuf:"bytes,2,rep,name=sec_label,json=secLabel"`
}

// +gogo:genproto=true
type DomainChardevSourceFile struct {
	Path     string                 `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
	Append   string                 `xml:"append,attr,omitempty" json:"append,omitempty" protobuf:"bytes,2,opt,name=append"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" protobuf:"bytes,3,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainChardevSourcePipe struct {
	Path     string                 `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" protobuf:"bytes,2,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainChardevSourceStdIO struct {
}

// +gogo:genproto=true
type DomainChardevSourceUDP struct {
	BindHost       string `xml:"-" json:"bindHost" protobuf:"bytes,1,opt,name=bindHost"`
	BindService    string `xml:"-" json:"bindService" protobuf:"bytes,2,opt,name=bindService"`
	ConnectHost    string `xml:"-" json:"connectHost" protobuf:"bytes,3,opt,name=connectHost"`
	ConnectService string `xml:"-" json:"connectService" protobuf:"bytes,4,opt,name=connectService"`
}

// +gogo:genproto=true
type DomainChardevSourceReconnect struct {
	Enabled string `xml:"enabled,attr" json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
	Timeout *int32 `xml:"timeout,attr" json:"timeout,omitempty" protobuf:"varint,2,opt,name=timeout"`
}

// +gogo:genproto=true
type DomainChardevSourceTCP struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
	Host      string                        `xml:"host,attr,omitempty" json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
	Service   string                        `xml:"service,attr,omitempty" json:"service,omitempty" protobuf:"bytes,3,opt,name=service"`
	TLS       string                        `xml:"tls,attr,omitempty" json:"tls,omitempty" protobuf:"bytes,4,opt,name=tls"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect,omitempty" protobuf:"bytes,5,opt,name=reconnect"`
}

// +gogo:genproto=true
type DomainChardevSourceUNIX struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
	Path      string                        `xml:"path,attr,omitempty" json:"path,omitempty" protobuf:"bytes,2,opt,name=path"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect,omitempty" protobuf:"bytes,3,opt,name=reconnect"`
	SecLabel  []DomainDeviceSecLabel        `xml:"seclabel" json:"seclabel" protobuf:"bytes,4,rep,name=seclabel"`
}

// +gogo:genproto=true
type DomainChardevSourceSpiceVMC struct {
}

// +gogo:genproto=true
type DomainChardevSourceSpicePort struct {
	Channel string `xml:"channel,attr" json:"channel" protobuf:"bytes,1,opt,name=channel"`
}

// +gogo:genproto=true
type DomainChardevSourceNMDM struct {
	Master string `xml:"master,attr" json:"master" protobuf:"bytes,1,opt,name=master"`
	Slave  string `xml:"slave,attr" json:"slave" protobuf:"bytes,2,opt,name=slave"`
}

// +gogo:genproto=true
type DomainChardevTarget struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,3,opt,name=state"` // is guest agent connected?
	Port  *int32 `xml:"port,attr" json:"port,omitempty" protobuf:"varint,4,opt,name=port"`
}

// +gogo:genproto=true
type DomainConsoleTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Port *int32 `xml:"port,attr" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
}

// +gogo:genproto=true
type DomainSerialTarget struct {
	Type  string                   `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Port  *int32                   `xml:"port,attr" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	Model *DomainSerialTargetModel `xml:"model" json:"model,omitempty" protobuf:"bytes,3,opt,name=model"`
}

// +gogo:genproto=true
type DomainSerialTargetModel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainParallelTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Port *int32 `xml:"port,attr" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
}

// +gogo:genproto=true
type DomainChannelTarget struct {
	VirtIO   *DomainChannelTargetVirtIO   `xml:"-" json:"virtio,omitempty" protobuf:"bytes,1,opt,name=virtio"`
	Xen      *DomainChannelTargetXen      `xml:"-" json:"xen,omitempty" protobuf:"bytes,2,opt,name=xen"`
	GuestFWD *DomainChannelTargetGuestFWD `xml:"-" json:"guestfwd,omitempty" protobuf:"bytes,3,opt,name=guestfwd"`
}

// +gogo:genproto=true
type DomainChannelTargetVirtIO struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,2,opt,name=state"` // is guest agent connected?
}

// +gogo:genproto=true
type DomainChannelTargetXen struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,2,opt,name=state"` // is guest agent connected?
}

// +gogo:genproto=true
type DomainChannelTargetGuestFWD struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Port    string `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`
}

// +gogo:genproto=true
type DomainAlias struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainDeviceACPI struct {
	Index int32 `xml:"index,attr,omitempty" json:"index,omitempty" protobuf:"varint,1,opt,name=index"`
}

// +gogo:genproto=true
type DomainAddressPCI struct {
	Domain        *int32             `xml:"domain,attr" json:"domain,omitempty" protobuf:"varint,1,opt,name=domain"`
	Bus           *int32             `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Slot          *int32             `xml:"slot,attr" json:"slot,omitempty" protobuf:"varint,3,opt,name=slot"`
	Function      *int32             `xml:"function,attr" json:"function,omitempty" protobuf:"varint,4,opt,name=function"`
	MultiFunction string             `xml:"multifunction,attr,omitempty" json:"multifunction,omitempty" protobuf:"bytes,5,opt,name=multifunction"`
	ZPCI          *DomainAddressZPCI `xml:"zpci" json:"zpci,omitempty" protobuf:"bytes,6,opt,name=zpci"`
}

// +gogo:genproto=true
type DomainAddressZPCI struct {
	UID *int32 `xml:"uid,attr,omitempty" json:"uid,omitempty" protobuf:"varint,1,opt,name=uid"`
	FID *int32 `xml:"fid,attr,omitempty" json:"fid,omitempty" protobuf:"varint,2,opt,name=fid"`
}

// +gogo:genproto=true
type DomainAddressUSB struct {
	Bus    *int32 `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,1,opt,name=bus"`
	Port   string `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`
	Device *int32 `xml:"device,attr" json:"device,omitempty" protobuf:"varint,3,opt,name=device"`
}

// +gogo:genproto=true
type DomainAddressDrive struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty" protobuf:"varint,1,opt,name=controller"`
	Bus        *int32 `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Target     *int32 `xml:"target,attr" json:"target,omitempty" protobuf:"varint,3,opt,name=target"`
	Unit       *int32 `xml:"unit,attr" json:"unit,omitempty" protobuf:"varint,4,opt,name=unit"`
}

// +gogo:genproto=true
type DomainAddressDIMM struct {
	Slot *int32 `xml:"slot,attr" json:"slot,omitempty" protobuf:"varint,1,opt,name=slot"`
	Base *int64 `xml:"base,attr" json:"base,omitempty" protobuf:"varint,2,opt,name=base"`
}

// +gogo:genproto=true
type DomainAddressISA struct {
	IOBase *int32 `xml:"iobase,attr" json:"iobase,omitempty" protobuf:"varint,1,opt,name=iobase"`
	IRQ    *int32 `xml:"irq,attr" json:"irq,omitempty" protobuf:"varint,2,opt,name=irq"`
}

// +gogo:genproto=true
type DomainAddressVirtioMMIO struct {
}

// +gogo:genproto=true
type DomainAddressCCW struct {
	CSSID *int32 `xml:"cssid,attr" json:"cssid,omitempty" protobuf:"varint,1,opt,name=cssid"`
	SSID  *int32 `xml:"ssid,attr" json:"ssid,omitempty" protobuf:"varint,2,opt,name=ssid"`
	DevNo *int32 `xml:"devno,attr" json:"devno,omitempty" protobuf:"varint,3,opt,name=devno"`
}

// +gogo:genproto=true
type DomainAddressVirtioSerial struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty" protobuf:"varint,1,opt,name=controller"`
	Bus        *int32 `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Port       *int32 `xml:"port,attr" json:"port,omitempty" protobuf:"varint,3,opt,name=port"`
}

// +gogo:genproto=true
type DomainAddressSpaprVIO struct {
	Reg *int64 `xml:"reg,attr" json:"reg,omitempty" protobuf:"varint,1,opt,name=reg"`
}

// +gogo:genproto=true
type DomainAddressCCID struct {
	Controller *int32 `xml:"controller,attr" json:"controller,omitempty" protobuf:"varint,1,opt,name=controller"`
	Slot       *int32 `xml:"slot,attr" json:"slot,omitempty" protobuf:"varint,2,opt,name=slot"`
}

// +gogo:genproto=true
type DomainAddressVirtioS390 struct {
}

// +gogo:genproto=true
type DomainAddressUnassigned struct {
}

// +gogo:genproto=true
type DomainAddress struct {
	PCI          *DomainAddressPCI          `json:"pci,omitempty" protobuf:"bytes,1,opt,name=pci"`
	Drive        *DomainAddressDrive        `json:"drive,omitempty" protobuf:"bytes,2,opt,name=drive"`
	VirtioSerial *DomainAddressVirtioSerial `json:"virtioSerial,omitempty" protobuf:"bytes,3,opt,name=virtioSerial"`
	CCID         *DomainAddressCCID         `json:"ccid,omitempty" protobuf:"bytes,4,opt,name=ccid"`
	USB          *DomainAddressUSB          `json:"usb,omitempty" protobuf:"bytes,5,opt,name=usb"`
	SpaprVIO     *DomainAddressSpaprVIO     `json:"spaprvio,omitempty" protobuf:"bytes,6,opt,name=spaprvio"`
	VirtioS390   *DomainAddressVirtioS390   `json:"virtioS390,omitempty" protobuf:"bytes,7,opt,name=virtioS390"`
	CCW          *DomainAddressCCW          `json:"ccw,omitempty" protobuf:"bytes,8,opt,name=ccw"`
	VirtioMMIO   *DomainAddressVirtioMMIO   `json:"virtiommio,omitempty" protobuf:"bytes,9,opt,name=virtiommio"`
	ISA          *DomainAddressISA          `json:"isa,omitempty" protobuf:"bytes,10,opt,name=isa"`
	DIMM         *DomainAddressDIMM         `json:"dimm,omitempty" protobuf:"bytes,11,opt,name=dimm"`
	Unassigned   *DomainAddressUnassigned   `json:"unassigned,omitempty" protobuf:"bytes,12,opt,name=unassigned"`
}

// +gogo:genproto=true
type DomainChardevLog struct {
	File   string `xml:"file,attr" json:"file" protobuf:"bytes,1,opt,name=file"`
	Append string `xml:"append,attr,omitempty" json:"append,omitempty" protobuf:"bytes,2,opt,name=append"`
}

// +gogo:genproto=true
type DomainConsole struct {
	XMLName  xml.Name               `xml:"console" json:"-"`
	TTY      string                 `xml:"tty,attr,omitempty" json:"tty,omitempty" protobuf:"bytes,1,opt,name=tty"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,2,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,3,opt,name=protocol"`
	Target   *DomainConsoleTarget   `xml:"target" json:"target,omitempty" protobuf:"bytes,4,opt,name=target"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty" protobuf:"bytes,5,opt,name=log"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,6,opt,name=acpi"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,7,opt,name=alias"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,8,opt,name=address"`
}

// +gogo:genproto=true
type DomainSerial struct {
	XMLName  xml.Name               `xml:"serial" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
	Target   *DomainSerialTarget    `xml:"target" json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty" protobuf:"bytes,4,opt,name=log"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,5,opt,name=acpi"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,6,opt,name=alias"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,7,opt,name=address"`
}

// +gogo:genproto=true
type DomainParallel struct {
	XMLName  xml.Name               `xml:"parallel" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
	Target   *DomainParallelTarget  `xml:"target" json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty" protobuf:"bytes,4,opt,name=log"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,5,opt,name=acpi"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,6,opt,name=alias"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,7,opt,name=address"`
}

// +gogo:genproto=true
type DomainChardevProtocol struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type DomainChannel struct {
	XMLName  xml.Name               `xml:"channel" json:"-"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
	Target   *DomainChannelTarget   `xml:"target" json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
	Log      *DomainChardevLog      `xml:"log" json:"log,omitempty" protobuf:"bytes,4,opt,name=log"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,5,opt,name=acpi"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,6,opt,name=alias"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,7,opt,name=address"`
}

// +gogo:genproto=true
type DomainRedirDev struct {
	XMLName  xml.Name               `xml:"redirdev" json:"-"`
	Bus      string                 `xml:"bus,attr,omitempty" json:"bus,omitempty" protobuf:"bytes,1,opt,name=bus"`
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,2,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,3,opt,name=protocol"`
	Boot     *DomainDeviceBoot      `xml:"boot" json:"boot,omitempty" protobuf:"bytes,4,opt,name=boot"`
	ACPI     *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,5,opt,name=acpi"`
	Alias    *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,6,opt,name=alias"`
	Address  *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,7,opt,name=address"`
}

// +gogo:genproto=true
type DomainRedirFilter struct {
	USB []DomainRedirFilterUSB `xml:"usbdev" json:"usb" protobuf:"bytes,1,rep,name=usb"`
}

// +gogo:genproto=true
type DomainRedirFilterUSB struct {
	Class   *int32 `xml:"class,attr" json:"class,omitempty" protobuf:"varint,1,opt,name=class"`
	Vendor  *int32 `xml:"vendor,attr" json:"vendor,omitempty" protobuf:"varint,2,opt,name=vendor"`
	Product *int32 `xml:"product,attr" json:"product,omitempty" protobuf:"varint,3,opt,name=product"`
	Version string `xml:"version,attr,omitempty" json:"version,omitempty" protobuf:"bytes,4,opt,name=version"`
	Allow   string `xml:"allow,attr" json:"allow" protobuf:"bytes,5,opt,name=allow"`
}

// +gogo:genproto=true
type DomainInput struct {
	XMLName xml.Name           `xml:"input" json:"-"`
	Type    string             `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Bus     string             `xml:"bus,attr,omitempty" json:"bus,omitempty" protobuf:"bytes,2,opt,name=bus"`
	Model   string             `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,3,opt,name=model"`
	Driver  *DomainInputDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,4,opt,name=driver"`
	Source  *DomainInputSource `xml:"source" json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,6,opt,name=acpi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty" protobuf:"bytes,7,opt,name=alias"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty" protobuf:"bytes,8,opt,name=address"`
}

// +gogo:genproto=true
type DomainInputDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,1,opt,name=iommu"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,2,opt,name=ats"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,3,opt,name=packed"`
}

// +gogo:genproto=true
type DomainInputSource struct {
	EVDev string `xml:"evdev,attr" json:"evdev" protobuf:"bytes,1,opt,name=evdev"`
}

// +gogo:genproto=true
type DomainGraphicListenerAddress struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type DomainGraphicListenerNetwork struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Network string `xml:"network,attr,omitempty" json:"network,omitempty" protobuf:"bytes,2,opt,name=network"`
}

// +gogo:genproto=true
type DomainGraphicListenerSocket struct {
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty" protobuf:"bytes,1,opt,name=socket"`
}

// +gogo:genproto=true
type DomainGraphicListener struct {
	Address *DomainGraphicListenerAddress `xml:"-" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Network *DomainGraphicListenerNetwork `xml:"-" json:"network,omitempty" protobuf:"bytes,2,opt,name=network"`
	Socket  *DomainGraphicListenerSocket  `xml:"-" json:"socket,omitempty" protobuf:"bytes,3,opt,name=socket"`
}

// +gogo:genproto=true
type DomainGraphicChannel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type DomainGraphicFileTransfer struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`
}

// +gogo:genproto=true
type DomainGraphicsSDLGL struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`
}

// +gogo:genproto=true
type DomainGraphicSDL struct {
	Display    string               `xml:"display,attr,omitempty" json:"display,omitempty" protobuf:"bytes,1,opt,name=display"`
	XAuth      string               `xml:"xauth,attr,omitempty" json:"xauth,omitempty" protobuf:"bytes,2,opt,name=xauth"`
	FullScreen string               `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty" protobuf:"bytes,3,opt,name=fullscreen"`
	GL         *DomainGraphicsSDLGL `xml:"gl" json:"gl,omitempty" protobuf:"bytes,4,opt,name=gl"`
}

// +gogo:genproto=true
type DomainGraphicVNC struct {
	Socket        string                  `xml:"socket,attr,omitempty" json:"socket,omitempty" protobuf:"bytes,1,opt,name=socket"`
	Port          int                     `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
	AutoPort      string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty" protobuf:"bytes,3,opt,name=autoport"`
	WebSocket     int                     `xml:"websocket,attr,omitempty" json:"websocket,omitempty" protobuf:"varint,4,opt,name=websocket"`
	Keymap        string                  `xml:"keymap,attr,omitempty" json:"keymap,omitempty" protobuf:"bytes,5,opt,name=keymap"`
	SharePolicy   string                  `xml:"sharePolicy,attr,omitempty" json:"sharePolicy,omitempty" protobuf:"bytes,6,opt,name=sharePolicy"`
	Passwd        string                  `xml:"passwd,attr,omitempty" json:"passwd,omitempty" protobuf:"bytes,7,opt,name=passwd"`
	PasswdValidTo string                  `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty" protobuf:"bytes,8,opt,name=passwdValidTo"`
	Connected     string                  `xml:"connected,attr,omitempty" json:"connected,omitempty" protobuf:"bytes,9,opt,name=connected"`
	PowerControl  string                  `xml:"powerControl,attr,omitempty" json:"powerControl,omitempty" protobuf:"bytes,10,opt,name=powerControl"`
	Listen        string                  `xml:"listen,attr,omitempty" json:"listen,omitempty" protobuf:"bytes,11,opt,name=listen"`
	Listeners     []DomainGraphicListener `xml:"listen" json:"listeners" protobuf:"bytes,12,rep,name=listeners"`
}

// +gogo:genproto=true
type DomainGraphicRDP struct {
	Port        int                     `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,1,opt,name=port"`
	AutoPort    string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty" protobuf:"bytes,2,opt,name=autoport"`
	ReplaceUser string                  `xml:"replaceUser,attr,omitempty" json:"replaceuser,omitempty" protobuf:"bytes,3,opt,name=replaceuser"`
	MultiUser   string                  `xml:"multiUser,attr,omitempty" json:"multiuser,omitempty" protobuf:"bytes,4,opt,name=multiuser"`
	Listen      string                  `xml:"listen,attr,omitempty" json:"listen,omitempty" protobuf:"bytes,5,opt,name=listen"`
	Listeners   []DomainGraphicListener `xml:"listen" json:"listeners" protobuf:"bytes,6,rep,name=listeners"`
}

// +gogo:genproto=true
type DomainGraphicDesktop struct {
	Display    string `xml:"display,attr,omitempty" json:"display,omitempty" protobuf:"bytes,1,opt,name=display"`
	FullScreen string `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty" protobuf:"bytes,2,opt,name=fullscreen"`
}

// +gogo:genproto=true
type DomainGraphicSpiceChannel struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type DomainGraphicSpiceImage struct {
	Compression string `xml:"compression,attr" json:"compression" protobuf:"bytes,1,opt,name=compression"`
}

// +gogo:genproto=true
type DomainGraphicSpiceJPEG struct {
	Compression string `xml:"compression,attr" json:"compression" protobuf:"bytes,1,opt,name=compression"`
}

// +gogo:genproto=true
type DomainGraphicSpiceZLib struct {
	Compression string `xml:"compression,attr" json:"compression" protobuf:"bytes,1,opt,name=compression"`
}

// +gogo:genproto=true
type DomainGraphicSpicePlayback struct {
	Compression string `xml:"compression,attr" json:"compression" protobuf:"bytes,1,opt,name=compression"`
}

// +gogo:genproto=true
type DomainGraphicSpiceStreaming struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainGraphicSpiceMouse struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainGraphicSpiceClipBoard struct {
	CopyPaste string `xml:"copypaste,attr" json:"copypaste" protobuf:"bytes,1,opt,name=copypaste"`
}

// +gogo:genproto=true
type DomainGraphicSpiceFileTransfer struct {
	Enable string `xml:"enable,attr" json:"enable" protobuf:"bytes,1,opt,name=enable"`
}

// +gogo:genproto=true
type DomainGraphicSpiceGL struct {
	Enable     string `xml:"enable,attr,omitempty" json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty" protobuf:"bytes,2,opt,name=rendernode"`
}

// +gogo:genproto=true
type DomainGraphicSpice struct {
	Port          int                             `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,1,opt,name=port"`
	TLSPort       int                             `xml:"tlsPort,attr,omitempty" json:"tlsPort,omitempty" protobuf:"varint,2,opt,name=tlsPort"`
	AutoPort      string                          `xml:"autoport,attr,omitempty" json:"autoPort,omitempty" protobuf:"bytes,3,opt,name=autoPort"`
	Listen        string                          `xml:"listen,attr,omitempty" json:"listen,omitempty" protobuf:"bytes,4,opt,name=listen"`
	Keymap        string                          `xml:"keymap,attr,omitempty" json:"keymap,omitempty" protobuf:"bytes,5,opt,name=keymap"`
	DefaultMode   string                          `xml:"defaultMode,attr,omitempty" json:"defaultMode,omitempty" protobuf:"bytes,6,opt,name=defaultMode"`
	Passwd        string                          `xml:"passwd,attr,omitempty" json:"passwd,omitempty" protobuf:"bytes,7,opt,name=passwd"`
	PasswdValidTo string                          `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty" protobuf:"bytes,8,opt,name=passwdValidTo"`
	Connected     string                          `xml:"connected,attr,omitempty" json:"connected,omitempty" protobuf:"bytes,9,opt,name=connected"`
	Listeners     []DomainGraphicListener         `xml:"listen" json:"listeners" protobuf:"bytes,10,rep,name=listeners"`
	Channel       []DomainGraphicSpiceChannel     `xml:"channel" json:"channel" protobuf:"bytes,11,rep,name=channel"`
	Image         *DomainGraphicSpiceImage        `xml:"image" json:"image,omitempty" protobuf:"bytes,12,opt,name=image"`
	JPEG          *DomainGraphicSpiceJPEG         `xml:"jpeg" json:"jpeg,omitempty" protobuf:"bytes,13,opt,name=jpeg"`
	ZLib          *DomainGraphicSpiceZLib         `xml:"zlib" json:"zlib,omitempty" protobuf:"bytes,14,opt,name=zlib"`
	Playback      *DomainGraphicSpicePlayback     `xml:"playback" json:"playback,omitempty" protobuf:"bytes,15,opt,name=playback"`
	Streaming     *DomainGraphicSpiceStreaming    `xml:"streaming" json:"streaming,omitempty" protobuf:"bytes,16,opt,name=streaming"`
	Mouse         *DomainGraphicSpiceMouse        `xml:"mouse" json:"mouse,omitempty" protobuf:"bytes,17,opt,name=mouse"`
	ClipBoard     *DomainGraphicSpiceClipBoard    `xml:"clipboard" json:"clipboard,omitempty" protobuf:"bytes,18,opt,name=clipboard"`
	FileTransfer  *DomainGraphicSpiceFileTransfer `xml:"filetransfer" json:"filetransfer,omitempty" protobuf:"bytes,19,opt,name=filetransfer"`
	GL            *DomainGraphicSpiceGL           `xml:"gl" json:"gl,omitempty" protobuf:"bytes,20,opt,name=gl"`
}

// +gogo:genproto=true
type DomainGraphicEGLHeadlessGL struct {
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty" protobuf:"bytes,1,opt,name=rendernode"`
}

// +gogo:genproto=true
type DomainGraphicEGLHeadless struct {
	GL *DomainGraphicEGLHeadlessGL `xml:"gl" json:"gl,omitempty" protobuf:"bytes,1,opt,name=gl"`
}

// +gogo:genproto=true
type DomainGraphicAudio struct {
	ID int32 `xml:"id,attr,omitempty" json:"id,omitempty" protobuf:"varint,1,opt,name=id"`
}

// +gogo:genproto=true
type DomainGraphic struct {
	XMLName     xml.Name                  `xml:"graphics" json:"-"`
	SDL         *DomainGraphicSDL         `xml:"-" json:"sdl,omitempty" protobuf:"bytes,1,opt,name=sdl"`
	VNC         *DomainGraphicVNC         `xml:"-" json:"vnc,omitempty" protobuf:"bytes,2,opt,name=vnc"`
	RDP         *DomainGraphicRDP         `xml:"-" json:"rdp,omitempty" protobuf:"bytes,3,opt,name=rdp"`
	Desktop     *DomainGraphicDesktop     `xml:"-" json:"desktop,omitempty" protobuf:"bytes,4,opt,name=desktop"`
	Spice       *DomainGraphicSpice       `xml:"-" json:"spice,omitempty" protobuf:"bytes,5,opt,name=spice"`
	EGLHeadless *DomainGraphicEGLHeadless `xml:"-" json:"eglHeadless,omitempty" protobuf:"bytes,6,opt,name=eglHeadless"`
	Audio       *DomainGraphicAudio       `xml:"audio" json:"audio,omitempty" protobuf:"bytes,7,opt,name=audio"`
}

// +gogo:genproto=true
type DomainVideoAccel struct {
	Accel3D    string `xml:"accel3d,attr,omitempty" json:"accel3d,omitempty" protobuf:"bytes,1,opt,name=accel3d"`
	Accel2D    string `xml:"accel2d,attr,omitempty" json:"accel2d,omitempty" protobuf:"bytes,2,opt,name=accel2d"`
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty" protobuf:"bytes,3,opt,name=rendernode"`
}

// +gogo:genproto=true
type DomainVideoResolution struct {
	X int32 `xml:"x,attr" json:"x" protobuf:"varint,1,opt,name=x"`
	Y int32 `xml:"y,attr" json:"y" protobuf:"varint,2,opt,name=y"`
}

// +gogo:genproto=true
type DomainVideoModel struct {
	Type       string                 `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Heads      int32                  `xml:"heads,attr,omitempty" json:"heads,omitempty" protobuf:"varint,2,opt,name=heads"`
	Ram        int32                  `xml:"ram,attr,omitempty" json:"ram,omitempty" protobuf:"varint,3,opt,name=ram"`
	VRam       int32                  `xml:"vram,attr,omitempty" json:"vram,omitempty" protobuf:"varint,4,opt,name=vram"`
	VRam64     int32                  `xml:"vram64,attr,omitempty" json:"vram64,omitempty" protobuf:"varint,5,opt,name=vram64"`
	VGAMem     int32                  `xml:"vgamem,attr,omitempty" json:"vgamem,omitempty" protobuf:"varint,6,opt,name=vgamem"`
	Primary    string                 `xml:"primary,attr,omitempty" json:"primary,omitempty" protobuf:"bytes,7,opt,name=primary"`
	Accel      *DomainVideoAccel      `xml:"acceleration" json:"accel,omitempty" protobuf:"bytes,8,opt,name=accel"`
	Resolution *DomainVideoResolution `xml:"resolution" json:"resolution,omitempty" protobuf:"bytes,9,opt,name=resolution"`
}

// +gogo:genproto=true
type DomainVideo struct {
	XMLName xml.Name           `xml:"video" json:"-"`
	Model   DomainVideoModel   `xml:"model" json:"model" protobuf:"bytes,1,opt,name=model"`
	Driver  *DomainVideoDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,2,opt,name=driver"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,3,opt,name=acpi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty" protobuf:"bytes,4,opt,name=alias"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
}

// +gogo:genproto=true
type DomainVideoDriver struct {
	Name    string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	VGAConf string `xml:"vgaconf,attr,omitempty" json:"vgaconf,omitempty" protobuf:"bytes,2,opt,name=vgaconf"`
	IOMMU   string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,3,opt,name=iommu"`
	ATS     string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,4,opt,name=ats"`
	Packed  string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,5,opt,name=packed"`
}

// +gogo:genproto=true
type DomainMemBalloonStats struct {
	Period int32 `xml:"period,attr" json:"period,omitempty" protobuf:"varint,1,opt,name=period"`
}

// +gogo:genproto=true
type DomainMemBalloon struct {
	XMLName           xml.Name                `xml:"memballoon" json:"-"`
	Model             string                  `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	AutoDeflate       string                  `xml:"autodeflate,attr,omitempty" json:"autodeflate,omitempty" protobuf:"bytes,2,opt,name=autodeflate"`
	FreePageReporting string                  `xml:"freePageReporting,attr,omitempty" json:"freePageReporting,omitempty" protobuf:"bytes,3,opt,name=freePageReporting"`
	Driver            *DomainMemBalloonDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,4,opt,name=driver"`
	Stats             *DomainMemBalloonStats  `xml:"stats" json:"stats,omitempty" protobuf:"bytes,5,opt,name=stats"`
	ACPI              *DomainDeviceACPI       `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,6,opt,name=acpi"`
	Alias             *DomainAlias            `xml:"alias" json:"alias,omitempty" protobuf:"bytes,7,opt,name=alias"`
	Address           *DomainAddress          `xml:"address" json:"address,omitempty" protobuf:"bytes,8,opt,name=address"`
}

// +gogo:genproto=true
type DomainVSockCID struct {
	Auto    string `xml:"auto,attr,omitempty" json:"auto,omitempty" protobuf:"bytes,1,opt,name=auto"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
}

// +gogo:genproto=true
type DomainVSockDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,1,opt,name=iommu"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,2,opt,name=ats"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,3,opt,name=packed"`
}

// +gogo:genproto=true
type DomainVSock struct {
	XMLName xml.Name           `xml:"vsock" json:"-"`
	Model   string             `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	CID     *DomainVSockCID    `xml:"cid" json:"cid,omitempty" protobuf:"bytes,2,opt,name=cid"`
	Driver  *DomainVSockDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,3,opt,name=driver"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,4,opt,name=acpi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty" protobuf:"bytes,5,opt,name=alias"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty" protobuf:"bytes,6,opt,name=address"`
}

// +gogo:genproto=true
type DomainMemBalloonDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,1,opt,name=iommu"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,2,opt,name=ats"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,3,opt,name=packed"`
}

// +gogo:genproto=true
type DomainPanic struct {
	XMLName xml.Name          `xml:"panic" json:"-"`
	Model   string            `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,2,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,3,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,4,opt,name=address"`
}

// +gogo:genproto=true
type DomainSoundCodec struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type DomainSound struct {
	XMLName xml.Name           `xml:"sound" json:"-"`
	Model   string             `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	Codec   []DomainSoundCodec `xml:"codec" json:"codec" protobuf:"bytes,2,rep,name=codec"`
	Audio   *DomainSoundAudio  `xml:"audio" json:"audio,omitempty" protobuf:"bytes,3,opt,name=audio"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,4,opt,name=acpi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty" protobuf:"bytes,5,opt,name=alias"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty" protobuf:"bytes,6,opt,name=address"`
}

// +gogo:genproto=true
type DomainSoundAudio struct {
	ID int32 `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
}

// +gogo:genproto=true
type DomainAudio struct {
	XMLName    xml.Name               `xml:"audio" json:"-"`
	ID         int                    `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	None       *DomainAudioNone       `xml:"-" json:"none,omitempty" protobuf:"bytes,2,opt,name=none"`
	ALSA       *DomainAudioALSA       `xml:"-" json:"alsa,omitempty" protobuf:"bytes,3,opt,name=alsa"`
	CoreAudio  *DomainAudioCoreAudio  `xml:"-" json:"coreAudio,omitempty" protobuf:"bytes,4,opt,name=coreAudio"`
	Jack       *DomainAudioJack       `xml:"-" json:"jack,omitempty" protobuf:"bytes,5,opt,name=jack"`
	OSS        *DomainAudioOSS        `xml:"-" json:"oss,omitempty" protobuf:"bytes,6,opt,name=oss"`
	PulseAudio *DomainAudioPulseAudio `xml:"-" json:"pulseAudio,omitempty" protobuf:"bytes,7,opt,name=pulseAudio"`
	SDL        *DomainAudioSDL        `xml:"-" json:"sdl,omitempty" protobuf:"bytes,8,opt,name=sdl"`
	SPICE      *DomainAudioSPICE      `xml:"-" json:"spice,omitempty" protobuf:"bytes,9,opt,name=spice"`
	File       *DomainAudioFile       `xml:"-" json:"file,omitempty" protobuf:"bytes,10,opt,name=file"`
}

// +gogo:genproto=true
type DomainAudioChannel struct {
	MixingEngine  string                      `xml:"mixingEngine,attr,omitempty" json:"mixingEngine,omitempty" protobuf:"bytes,1,opt,name=mixingEngine"`
	FixedSettings string                      `xml:"fixedSettings,attr,omitempty" json:"fixedSettings,omitempty" protobuf:"bytes,2,opt,name=fixedSettings"`
	Voices        int32                       `xml:"voices,attr,omitempty" json:"voices,omitempty" protobuf:"varint,3,opt,name=voices"`
	Settings      *DomainAudioChannelSettings `xml:"settings" json:"settings,omitempty" protobuf:"bytes,4,opt,name=settings"`
	BufferLength  int32                       `xml:"bufferLength,attr,omitempty" json:"bufferLength,omitempty" protobuf:"varint,5,opt,name=bufferLength"`
}

// +gogo:genproto=true
type DomainAudioChannelSettings struct {
	Frequency int32  `xml:"frequency,attr,omitempty" json:"frequency,omitempty" protobuf:"varint,1,opt,name=frequency"`
	Channels  int32  `xml:"channels,attr,omitempty" json:"channels,omitempty" protobuf:"varint,2,opt,name=channels"`
	Format    string `xml:"format,attr,omitempty" json:"format,omitempty" protobuf:"bytes,3,opt,name=format"`
}

// +gogo:genproto=true
type DomainAudioNone struct {
	Input  *DomainAudioNoneChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,1,opt,name=input"`
	Output *DomainAudioNoneChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioNoneChannel struct {
	DomainAudioChannel `protobuf:"bytes,1,opt,name=domainAudioChannel"`
}

// +gogo:genproto=true
type DomainAudioALSA struct {
	Input  *DomainAudioALSAChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,1,opt,name=input"`
	Output *DomainAudioALSAChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioALSAChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	Dev                string `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,2,opt,name=dev"`
}

// +gogo:genproto=true
type DomainAudioCoreAudio struct {
	Input  *DomainAudioCoreAudioChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,1,opt,name=input"`
	Output *DomainAudioCoreAudioChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioCoreAudioChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	BufferCount        int32 `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty" protobuf:"varint,2,opt,name=bufferCount"`
}

// +gogo:genproto=true
type DomainAudioJack struct {
	Input  *DomainAudioJackChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,1,opt,name=input"`
	Output *DomainAudioJackChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioJackChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	ServerName         string `xml:"serverName,attr,omitempty" json:"serverName,omitempty" protobuf:"bytes,2,opt,name=serverName"`
	ClientName         string `xml:"clientName,attr,omitempty" json:"clientName,omitempty" protobuf:"bytes,3,opt,name=clientName"`
	ConnectPorts       string `xml:"connectPorts,attr,omitempty" json:"connectPorts,omitempty" protobuf:"bytes,4,opt,name=connectPorts"`
	ExactName          string `xml:"exactName,attr,omitempty" json:"exactName,omitempty" protobuf:"bytes,5,opt,name=exactName"`
}

// +gogo:genproto=true
type DomainAudioOSS struct {
	TryMMap   string `xml:"tryMMap,attr,omitempty" json:"tryMMap,omitempty" protobuf:"bytes,1,opt,name=tryMMap"`
	Exclusive string `xml:"exclusive,attr,omitempty" json:"exclusive,omitempty" protobuf:"bytes,2,opt,name=exclusive"`
	DSPPolicy *int32 `xml:"dspPolicy,attr" json:"dspPolicy,omitempty" protobuf:"varint,3,opt,name=dspPolicy"`

	Input  *DomainAudioOSSChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,4,opt,name=input"`
	Output *DomainAudioOSSChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,5,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioOSSChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	Dev                string `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,2,opt,name=dev"`
	BufferCount        int32  `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty" protobuf:"varint,3,opt,name=bufferCount"`
	TryPoll            string `xml:"tryPoll,attr,omitempty" json:"tryPoll,omitempty" protobuf:"bytes,4,opt,name=tryPoll"`
}

// +gogo:genproto=true
type DomainAudioPulseAudio struct {
	ServerName string                        `xml:"serverName,attr,omitempty" json:"serverName,omitempty" protobuf:"bytes,1,opt,name=serverName"`
	Input      *DomainAudioPulseAudioChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,2,opt,name=input"`
	Output     *DomainAudioPulseAudioChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,3,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioPulseAudioChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	Name               string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	StreamName         string `xml:"streamName,attr,omitempty" json:"streamName,omitempty" protobuf:"bytes,3,opt,name=streamName"`
	Latency            int32  `xml:"latency,attr,omitempty" json:"latency,omitempty" protobuf:"varint,4,opt,name=latency"`
}

// +gogo:genproto=true
type DomainAudioSDL struct {
	Driver string                 `xml:"driver,attr,omitempty" json:"driver,omitempty" protobuf:"bytes,1,opt,name=driver"`
	Input  *DomainAudioSDLChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,2,opt,name=input"`
	Output *DomainAudioSDLChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,3,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioSDLChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
	BufferCount        int32 `xml:"bufferCount,attr,omitempty" json:"bufferCount,omitempty" protobuf:"varint,2,opt,name=bufferCount"`
}

// +gogo:genproto=true
type DomainAudioSPICE struct {
	Input  *DomainAudioSPICEChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,1,opt,name=input"`
	Output *DomainAudioSPICEChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,2,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioSPICEChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
}

// +gogo:genproto=true
type DomainAudioFile struct {
	Path   string                  `xml:"path,attr,omitempty" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	Input  *DomainAudioFileChannel `xml:"input" json:"input,omitempty" protobuf:"bytes,2,opt,name=input"`
	Output *DomainAudioFileChannel `xml:"output" json:"output,omitempty" protobuf:"bytes,3,opt,name=output"`
}

// +gogo:genproto=true
type DomainAudioFileChannel struct {
	DomainAudioChannel `json:",inline" protobuf:"bytes,1,opt,name=domainAudioChannel"`
}

// +gogo:genproto=true
type DomainRNGRate struct {
	Bytes  int32 `xml:"bytes,attr" json:"bytes" protobuf:"varint,1,opt,name=bytes"`
	Period int32 `xml:"period,attr,omitempty" json:"period,omitempty" protobuf:"varint,2,opt,name=period"`
}

// +gogo:genproto=true
type DomainRNGBackend struct {
	Random  *DomainRNGBackendRandom  `xml:"-" json:"random,omitempty" protobuf:"bytes,1,opt,name=random"`
	EGD     *DomainRNGBackendEGD     `xml:"-" json:"egd,omitempty" protobuf:"bytes,2,opt,name=egd"`
	BuiltIn *DomainRNGBackendBuiltIn `xml:"-" json:"builtin,omitempty" protobuf:"bytes,3,opt,name=builtin"`
}

// +gogo:genproto=true
type DomainRNGBackendEGD struct {
	Source   *DomainChardevSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
}

// +gogo:genproto=true
type DomainRNGBackendRandom struct {
	Device string `xml:",chardata" json:"device" protobuf:"bytes,1,opt,name=device"`
}

// +gogo:genproto=true
type DomainRNGBackendBuiltIn struct {
}

// +gogo:genproto=true
type DomainRNG struct {
	XMLName xml.Name          `xml:"rng" json:"-"`
	Model   string            `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	Driver  *DomainRNGDriver  `xml:"driver" json:"driver,omitempty" protobuf:"bytes,2,opt,name=driver"`
	Rate    *DomainRNGRate    `xml:"rate" json:"rate,omitempty" protobuf:"bytes,3,opt,name=rate"`
	Backend *DomainRNGBackend `xml:"backend" json:"backend,omitempty" protobuf:"bytes,4,opt,name=backend"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,5,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,6,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,7,opt,name=address"`
}

// +gogo:genproto=true
type DomainRNGDriver struct {
	IOMMU  string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" protobuf:"bytes,1,opt,name=iommu"`
	ATS    string `xml:"ats,attr,omitempty" json:"ats,omitempty" protobuf:"bytes,2,opt,name=ats"`
	Packed string `xml:"packed,attr,omitempty" json:"packed,omitempty" protobuf:"bytes,3,opt,name=packed"`
}

// +gogo:genproto=true
type DomainHostdevSubsysUSB struct {
	Source *DomainHostdevSubsysUSBSource `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
}

// +gogo:genproto=true
type DomainHostdevSubsysUSBSource struct {
	Address *DomainAddressUSB `xml:"address" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSI struct {
	SGIO      string                         `xml:"sgio,attr,omitempty" json:"sgio,omitempty" protobuf:"bytes,1,opt,name=sgio"`
	RawIO     string                         `xml:"rawio,attr,omitempty" json:"rawio,omitempty" protobuf:"bytes,2,opt,name=rawio"`
	Source    *DomainHostdevSubsysSCSISource `xml:"source" json:"source,omitempty" protobuf:"bytes,3,opt,name=source"`
	ReadOnly  *DomainDiskReadOnly            `xml:"readonly" json:"readonly,omitempty" protobuf:"bytes,4,opt,name=readonly"`
	Shareable *DomainDiskShareable           `xml:"shareable" json:"shareable,omitempty" protobuf:"bytes,5,opt,name=shareable"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSISource struct {
	Host  *DomainHostdevSubsysSCSISourceHost  `xml:"-" json:"host,omitempty" protobuf:"bytes,1,opt,name=host"`
	ISCSI *DomainHostdevSubsysSCSISourceISCSI `xml:"-" json:"iscsi,omitempty" protobuf:"bytes,2,opt,name=iscsi"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSIAdapter struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSISourceHost struct {
	Adapter *DomainHostdevSubsysSCSIAdapter `xml:"adapter" json:"adapter,omitempty" protobuf:"bytes,1,opt,name=adapter"`
	Address *DomainAddressDrive             `xml:"address" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSISourceISCSI struct {
	Name      string                                  `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Host      []DomainDiskSourceHost                  `xml:"host" json:"host" protobuf:"bytes,2,rep,name=host"`
	Auth      *DomainDiskAuth                         `xml:"auth" json:"auth,omitempty" protobuf:"bytes,3,opt,name=auth"`
	Initiator *DomainHostdevSubsysSCSISourceInitiator `xml:"initiator" json:"initiator,omitempty" protobuf:"bytes,4,opt,name=initiator"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSISourceInitiator struct {
	IQN DomainHostdevSubsysSCSISourceIQN `xml:"iqn" json:"iqn" protobuf:"bytes,1,opt,name=iqn"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSISourceIQN struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSIHost struct {
	Model  string                             `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	Source *DomainHostdevSubsysSCSIHostSource `xml:"source" json:"source,omitempty" protobuf:"bytes,2,opt,name=source"`
}

// +gogo:genproto=true
type DomainHostdevSubsysSCSIHostSource struct {
	Protocol string `xml:"protocol,attr,omitempty" json:"protocol,omitempty" protobuf:"bytes,1,opt,name=protocol"`
	WWPN     string `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty" protobuf:"bytes,2,opt,name=wwpn"`
}

// +gogo:genproto=true
type DomainHostdevSubsysPCISource struct {
	WriteFiltering string            `xml:"writeFiltering,attr,omitempty" json:"writeFiltering,omitempty" protobuf:"bytes,1,opt,name=writeFiltering"`
	Address        *DomainAddressPCI `xml:"address" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
}

// +gogo:genproto=true
type DomainHostdevSubsysPCIDriver struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainHostdevSubsysPCI struct {
	Driver  *DomainHostdevSubsysPCIDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,1,opt,name=driver"`
	Source  *DomainHostdevSubsysPCISource `xml:"source" json:"source,omitempty" protobuf:"bytes,2,opt,name=source"`
	Teaming *DomainInterfaceTeaming       `xml:"teaming" json:"teaming,omitempty" protobuf:"bytes,3,opt,name=teaming"`
}

// +gogo:genproto=true
type DomainAddressMDev struct {
	UUID string `xml:"uuid,attr" json:"uuid" protobuf:"bytes,1,opt,name=uuid"`
}

// +gogo:genproto=true
type DomainHostdevSubsysMDevSource struct {
	Address *DomainAddressMDev `xml:"address" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type DomainHostdevSubsysMDev struct {
	Model   string                         `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	Display string                         `xml:"display,attr,omitempty" json:"display,omitempty" protobuf:"bytes,2,opt,name=display"`
	RamFB   string                         `xml:"ramfb,attr,omitempty" json:"ramfb,omitempty" protobuf:"bytes,3,opt,name=ramfb"`
	Source  *DomainHostdevSubsysMDevSource `xml:"source" json:"source,omitempty" protobuf:"bytes,4,opt,name=source"`
}

// +gogo:genproto=true
type DomainHostdevCapsStorage struct {
	Source *DomainHostdevCapsStorageSource `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
}

// +gogo:genproto=true
type DomainHostdevCapsStorageSource struct {
	Block string `xml:"block" json:"block" protobuf:"bytes,1,opt,name=block"`
}

// +gogo:genproto=true
type DomainHostdevCapsMisc struct {
	Source *DomainHostdevCapsMiscSource `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
}

// +gogo:genproto=true
type DomainHostdevCapsMiscSource struct {
	Char string `xml:"char" json:"char" protobuf:"bytes,1,opt,name=char"`
}

// +gogo:genproto=true
type DomainIP struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,2,opt,name=family"`
	Prefix  *int32 `xml:"prefix,attr" json:"prefix,omitempty" protobuf:"varint,3,opt,name=prefix"`
}

// +gogo:genproto=true
type DomainRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,1,opt,name=family"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty" protobuf:"bytes,3,opt,name=gateway"`
}

// +gogo:genproto=true
type DomainHostdevCapsNet struct {
	Source *DomainHostdevCapsNetSource `xml:"source" json:"source,omitempty" protobuf:"bytes,1,opt,name=source"`
	IP     []DomainIP                  `xml:"ip" json:"ip" protobuf:"bytes,2,rep,name=ip"`
	Route  []DomainRoute               `xml:"route" json:"route" protobuf:"bytes,3,rep,name=route"`
}

// +gogo:genproto=true
type DomainHostdevCapsNetSource struct {
	Interface string `xml:"interface" json:"interface" protobuf:"bytes,1,opt,name=interface"`
}

// +gogo:genproto=true
type DomainHostdev struct {
	Managed        string                       `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,1,opt,name=managed"`
	SubsysUSB      *DomainHostdevSubsysUSB      `xml:"-" json:"usb,omitempty" protobuf:"bytes,2,opt,name=usb"`
	SubsysSCSI     *DomainHostdevSubsysSCSI     `xml:"-" json:"scsi,omitempty" protobuf:"bytes,3,opt,name=scsi"`
	SubsysSCSIHost *DomainHostdevSubsysSCSIHost `xml:"-" json:"scsiHost,omitempty" protobuf:"bytes,4,opt,name=scsiHost"`
	SubsysPCI      *DomainHostdevSubsysPCI      `xml:"-" json:"pci,omitempty" protobuf:"bytes,5,opt,name=pci"`
	SubsysMDev     *DomainHostdevSubsysMDev     `xml:"-" json:"mdev,omitempty" protobuf:"bytes,6,opt,name=mdev"`
	CapsStorage    *DomainHostdevCapsStorage    `xml:"-" json:"storage,omitempty" protobuf:"bytes,7,opt,name=storage"`
	CapsMisc       *DomainHostdevCapsMisc       `xml:"-" json:"misc,omitempty" protobuf:"bytes,8,opt,name=misc"`
	CapsNet        *DomainHostdevCapsNet        `xml:"-" json:"net,omitempty" protobuf:"bytes,9,opt,name=net"`
	Boot           *DomainDeviceBoot            `xml:"boot" json:"boot,omitempty" protobuf:"bytes,10,opt,name=boot"`
	ROM            *DomainROM                   `xml:"rom" json:"rom,omitempty" protobuf:"bytes,11,opt,name=rom"`
	ACPI           *DomainDeviceACPI            `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,12,opt,name=acpi"`
	Alias          *DomainAlias                 `xml:"alias" json:"alias,omitempty" protobuf:"bytes,13,opt,name=alias"`
	Address        *DomainAddress               `xml:"address" json:"address,omitempty" protobuf:"bytes,14,opt,name=address"`
}

// +gogo:genproto=true
type DomainMemorydevSource struct {
	NodeMask  string                          `xml:"nodemask,omitempty" protobuf:"bytes,1,opt,name=nodeMask"`
	PageSize  *DomainMemorydevSourcePagesize  `xml:"pagesize" json:"pageSize,omitempty" protobuf:"bytes,2,opt,name=pageSize"`
	Path      string                          `xml:"path,omitempty" json:"path,omitempty" protobuf:"bytes,3,opt,name=path"`
	AlignSize *DomainMemorydevSourceAlignsize `xml:"alignsize" json:"alignSize,omitempty" protobuf:"bytes,4,opt,name=alignSize"`
	PMem      *DomainMemorydevSourcePMem      `xml:"pmem" json:"pmem,omitempty" protobuf:"bytes,5,opt,name=pmem"`
}

// +gogo:genproto=true
type DomainMemorydevSourcePMem struct {
}

// +gogo:genproto=true
type DomainMemorydevSourcePagesize struct {
	Value int64  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainMemorydevSourceAlignsize struct {
	Value int64  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainMemorydevTargetNode struct {
	Value int32 `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainMemorydevTargetReadOnly struct {
}

// +gogo:genproto=true
type DomainMemorydevTargetSize struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainMemorydevTargetLabel struct {
	Size_ *DomainMemorydevTargetSize `xml:"size" json:"size,omitempty" protobuf:"bytes,1,opt,name=size"`
}

// +gogo:genproto=true
type DomainMemorydevTarget struct {
	Size_    *DomainMemorydevTargetSize     `xml:"size" json:"size,omitempty" protobuf:"bytes,1,opt,name=size"`
	Node     *DomainMemorydevTargetNode     `xml:"node" json:"node,omitempty" protobuf:"bytes,2,opt,name=node"`
	Label    *DomainMemorydevTargetLabel    `xml:"label" json:"label,omitempty" protobuf:"bytes,3,opt,name=label"`
	ReadOnly *DomainMemorydevTargetReadOnly `xml:"readonly" json:"readonly,omitempty" protobuf:"bytes,4,opt,name=readonly"`
}

// +gogo:genproto=true
type DomainMemorydev struct {
	XMLName xml.Name               `xml:"memory" json:"-"`
	Model   string                 `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	Access  string                 `xml:"access,attr,omitempty" json:"access,omitempty" protobuf:"bytes,2,opt,name=access"`
	Discard string                 `xml:"discard,attr,omitempty" json:"discard,omitempty" protobuf:"bytes,3,opt,name=discard"`
	UUID    string                 `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,4,opt,name=uuid"`
	Source  *DomainMemorydevSource `xml:"source" json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
	Target  *DomainMemorydevTarget `xml:"target" json:"target,omitempty" protobuf:"bytes,6,opt,name=target"`
	ACPI    *DomainDeviceACPI      `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,7,opt,name=acpi"`
	Alias   *DomainAlias           `xml:"alias" json:"alias,omitempty" protobuf:"bytes,8,opt,name=alias"`
	Address *DomainAddress         `xml:"address" json:"address,omitempty" protobuf:"bytes,9,opt,name=address"`
}

// +gogo:genproto=true
type DomainWatchdog struct {
	XMLName xml.Name          `xml:"watchdog" json:"-"`
	Model   string            `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	Action  string            `xml:"action,attr,omitempty" json:"action,omitempty" protobuf:"bytes,2,opt,name=action"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,3,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,4,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
}

// +gogo:genproto=true
type DomainHub struct {
	Type    string            `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,2,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,3,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,4,opt,name=address"`
}

// +gogo:genproto=true
type DomainIOMMU struct {
	Model  string             `xml:"model,attr" json:"model" protobuf:"bytes,1,opt,name=model"`
	Driver *DomainIOMMUDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,2,opt,name=driver"`
}

// +gogo:genproto=true
type DomainIOMMUDriver struct {
	IntRemap    string `xml:"intremap,attr,omitempty" json:"intremap,omitempty" protobuf:"bytes,1,opt,name=intremap"`
	CachingMode string `xml:"caching_mode,attr,omitempty" json:"cachingMode,omitempty" protobuf:"bytes,2,opt,name=cachingMode"`
	EIM         string `xml:"eim,attr,omitempty" json:"eim,omitempty" protobuf:"bytes,3,opt,name=eim"`
	IOTLB       string `xml:"iotlb,attr,omitempty" json:"iotlb,omitempty" protobuf:"bytes,4,opt,name=iotlb"`
	AWBits      int32  `xml:"aw_bits,attr,omitempty" json:"awBits,omitempty" protobuf:"varint,5,opt,name=awBits"`
}

// +gogo:genproto=true
type DomainNVRAM struct {
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,1,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,2,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,3,opt,name=address"`
}

// +gogo:genproto=true
type DomainLease struct {
	Lockspace string             `xml:"lockspace" json:"lockspace" protobuf:"bytes,1,opt,name=lockspace"`
	Key       string             `xml:"key" json:"key" protobuf:"bytes,2,opt,name=key"`
	Target    *DomainLeaseTarget `xml:"target" json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
}

// +gogo:genproto=true
type DomainLeaseTarget struct {
	Path   string `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
	Offset int64  `xml:"offset,attr,omitempty" json:"offset,omitempty" protobuf:"varint,2,opt,name=offset"`
}

// +gogo:genproto=true
type DomainSmartcard struct {
	XMLName     xml.Name                  `xml:"smartcard" json:"-"`
	Passthrough *DomainChardevSource      `xml:"source" json:"passthrough,omitempty" protobuf:"bytes,1,opt,name=passthrough"`
	Protocol    *DomainChardevProtocol    `xml:"protocol" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
	Host        *DomainSmartcardHost      `xml:"-" json:"host,omitempty" protobuf:"bytes,3,opt,name=host"`
	HostCerts   []DomainSmartcardHostCert `xml:"certificate" json:"certificate,omitempty" protobuf:"bytes,4,rep,name=certificate"`
	Database    string                    `xml:"database,omitempty" json:"database,omitempty" protobuf:"bytes,5,opt,name=database"`
	ACPI        *DomainDeviceACPI         `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,6,opt,name=acpi"`
	Alias       *DomainAlias              `xml:"alias" json:"alias,omitempty" protobuf:"bytes,7,opt,name=alias"`
	Address     *DomainAddress            `xml:"address" json:"address,omitempty" protobuf:"bytes,8,opt,name=address"`
}

// +gogo:genproto=true
type DomainSmartcardHost struct {
}

// +gogo:genproto=true
type DomainSmartcardHostCert struct {
	File string `xml:",chardata" json:"file" protobuf:"bytes,1,opt,name=file"`
}

// +gogo:genproto=true
type DomainTPM struct {
	XMLName xml.Name          `xml:"tpm" json:"-"`
	Model   string            `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	Backend *DomainTPMBackend `xml:"backend" json:"backend,omitempty" protobuf:"bytes,2,opt,name=backend"`
	ACPI    *DomainDeviceACPI `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,3,opt,name=acpi"`
	Alias   *DomainAlias      `xml:"alias" json:"alias,omitempty" protobuf:"bytes,4,opt,name=alias"`
	Address *DomainAddress    `xml:"address" json:"address,omitempty" protobuf:"bytes,5,opt,name=address"`
}

// +gogo:genproto=true
type DomainTPMBackend struct {
	Passthrough *DomainTPMBackendPassthrough `xml:"-" json:"passthrough,omitempty" protobuf:"bytes,1,opt,name=passthrough"`
	Emulator    *DomainTPMBackendEmulator    `xml:"-" json:"emulator,omitempty" protobuf:"bytes,2,opt,name=emulator"`
}

// +gogo:genproto=true
type DomainTPMBackendPassthrough struct {
	Device *DomainTPMBackendDevice `xml:"device" json:"device,omitempty" protobuf:"bytes,1,opt,name=device"`
}

// +gogo:genproto=true
type DomainTPMBackendEmulator struct {
	Version         string                      `xml:"version,attr,omitempty" json:"version,omitempty" protobuf:"bytes,1,opt,name=version"`
	Encryption      *DomainTPMBackendEncryption `xml:"encryption" json:"encryption,omitempty" protobuf:"bytes,2,opt,name=encryption"`
	PersistentState string                      `xml:"persistent_state,attr,omitempty" json:"persistentState,omitempty" protobuf:"bytes,3,opt,name=persistentState"`
}

// +gogo:genproto=true
type DomainTPMBackendEncryption struct {
	Secret string `xml:"secret,attr" json:"secret" protobuf:"bytes,1,opt,name=secret"`
}

// +gogo:genproto=true
type DomainTPMBackendDevice struct {
	Path string `xml:"path,attr" json:"path" protobuf:"bytes,1,opt,name=path"`
}

// +gogo:genproto=true
type DomainShmem struct {
	XMLName xml.Name           `xml:"shmem" json:"-"`
	Name    string             `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Role    string             `xml:"role,attr,omitempty" json:"role,omitempty" protobuf:"bytes,2,opt,name=role"`
	Size_   *DomainShmemSize   `xml:"size" json:"size,omitempty" protobuf:"bytes,3,opt,name=size"`
	Model   *DomainShmemModel  `xml:"model" json:"model,omitempty" protobuf:"bytes,4,opt,name=model"`
	Server  *DomainShmemServer `xml:"server" json:"server,omitempty" protobuf:"bytes,5,opt,name=server"`
	MSI     *DomainShmemMSI    `xml:"msi" json:"msi,omitempty" protobuf:"bytes,6,opt,name=msi"`
	ACPI    *DomainDeviceACPI  `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,7,opt,name=acpi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias,omitempty" protobuf:"bytes,8,opt,name=alias"`
	Address *DomainAddress     `xml:"address" json:"address,omitempty" protobuf:"bytes,9,opt,name=address"`
}

// +gogo:genproto=true
type DomainShmemSize struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainShmemModel struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type DomainShmemServer struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

// +gogo:genproto=true
type DomainShmemMSI struct {
	Enabled   string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" protobuf:"bytes,1,opt,name=enabled"`
	Vectors   int32  `xml:"vectors,attr,omitempty" json:"vectors,omitempty" protobuf:"varint,2,opt,name=vectors"`
	IOEventFD string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" protobuf:"bytes,3,opt,name=ioeventfd"`
}

// +gogo:genproto=true
type DomainDeviceList struct {
	Emulator     string              `xml:"emulator,omitempty" json:"emulator,omitempty" protobuf:"bytes,1,opt,name=emulator"`
	Disks        []DomainDisk        `xml:"disk" json:"disks" protobuf:"bytes,2,rep,name=disks"`
	Controllers  []DomainController  `xml:"controller" json:"controllers" protobuf:"bytes,3,rep,name=controllers"`
	Leases       []DomainLease       `xml:"lease" json:"leases" protobuf:"bytes,4,rep,name=leases"`
	Filesystems  []DomainFilesystem  `xml:"filesystem" json:"filesystems" protobuf:"bytes,5,rep,name=filesystems"`
	Interfaces   []DomainInterface   `xml:"interface" json:"interfaces" protobuf:"bytes,6,rep,name=interfaces"`
	Smartcards   []DomainSmartcard   `xml:"smartcard" json:"smartcards" protobuf:"bytes,7,rep,name=smartcards"`
	Serials      []DomainSerial      `xml:"serial" json:"serials" protobuf:"bytes,8,rep,name=serials"`
	Parallels    []DomainParallel    `xml:"parallel" json:"parallels" protobuf:"bytes,9,rep,name=parallels"`
	Consoles     []DomainConsole     `xml:"console" json:"consoles" protobuf:"bytes,10,rep,name=consoles"`
	Channels     []DomainChannel     `xml:"channel" json:"channels" protobuf:"bytes,11,rep,name=channels"`
	Inputs       []DomainInput       `xml:"input" json:"inputs" protobuf:"bytes,12,rep,name=inputs"`
	TPMs         []DomainTPM         `xml:"tpm" json:"tpms" protobuf:"bytes,13,rep,name=tpms"`
	Graphics     []DomainGraphic     `xml:"graphics" json:"graphics" protobuf:"bytes,14,rep,name=graphics"`
	Sounds       []DomainSound       `xml:"sound" json:"sounds" protobuf:"bytes,15,rep,name=sounds"`
	Audios       []DomainAudio       `xml:"audio" json:"audios" protobuf:"bytes,16,rep,name=audios"`
	Videos       []DomainVideo       `xml:"video" json:"videos" protobuf:"bytes,17,rep,name=videos"`
	Hostdevs     []DomainHostdev     `xml:"hostdev" json:"hostdevs" protobuf:"bytes,18,rep,name=hostdevs"`
	RedirDevs    []DomainRedirDev    `xml:"redirdev" json:"redirDevs" protobuf:"bytes,19,rep,name=redirDevs"`
	RedirFilters []DomainRedirFilter `xml:"redirfilter" json:"redirfilters" protobuf:"bytes,20,rep,name=redirfilters"`
	Hubs         []DomainHub         `xml:"hub" json:"hubs" protobuf:"bytes,21,rep,name=hubs"`
	Watchdog     *DomainWatchdog     `xml:"watchdog" json:"watchdog,omitempty" protobuf:"bytes,22,opt,name=watchdog"`
	MemBalloon   *DomainMemBalloon   `xml:"memballoon" json:"memballoon,omitempty" protobuf:"bytes,23,opt,name=memballoon"`
	RNGs         []DomainRNG         `xml:"rng" json:"rngs" protobuf:"bytes,24,rep,name=rngs"`
	NVRAM        *DomainNVRAM        `xml:"nvram" json:"nvram" protobuf:"bytes,25,opt,name=nvram"`
	Panics       []DomainPanic       `xml:"panic" json:"panics" protobuf:"bytes,26,rep,name=panics"`
	Shmems       []DomainShmem       `xml:"shmem" json:"shmems" protobuf:"bytes,27,rep,name=shmems"`
	Memorydevs   []DomainMemorydev   `xml:"memory" json:"memorydevs" protobuf:"bytes,28,rep,name=memorydevs"`
	IOMMU        *DomainIOMMU        `xml:"iommu" json:"iommu,omitempty" protobuf:"bytes,29,opt,name=iommu"`
	VSock        *DomainVSock        `xml:"vsock" json:"vsock,omitempty" protobuf:"bytes,30,opt,name=vsock"`
}

// +gogo:genproto=true
type DomainMemory struct {
	Value    int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit     string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
	DumpCore string `xml:"dumpCore,attr,omitempty" json:"dumpCore,omitempty" protobuf:"bytes,3,opt,name=dumpCore"`
}

// +gogo:genproto=true
type DomainCurrentMemory struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainMaxMemory struct {
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
	Slots int32  `xml:"slots,attr,omitempty" json:"slots,omitempty" protobuf:"varint,3,opt,name=slots"`
}

// +gogo:genproto=true
type DomainMemoryHugepage struct {
	Size_   int32  `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
	Unit    string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
	Nodeset string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty" protobuf:"bytes,3,opt,name=nodeset"`
}

// +gogo:genproto=true
type DomainMemoryHugepages struct {
	Hugepages []DomainMemoryHugepage `xml:"page" json:"hugepages" protobuf:"bytes,1,rep,name=hugepages"`
}

// +gogo:genproto=true
type DomainMemoryNosharepages struct {
}

// +gogo:genproto=true
type DomainMemoryLocked struct {
}

// +gogo:genproto=true
type DomainMemorySource struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type DomainMemoryAccess struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainMemoryAllocation struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainMemoryDiscard struct {
}

// +gogo:genproto=true
type DomainMemoryBacking struct {
	MemoryHugePages    *DomainMemoryHugepages    `xml:"hugepages" json:"memoryHugePages,omitempty" protobuf:"bytes,1,opt,name=memoryHugePages"`
	MemoryNosharepages *DomainMemoryNosharepages `xml:"nosharepages" json:"memoryNosharepages,omitempty" protobuf:"bytes,2,opt,name=memoryNosharepages"`
	MemoryLocked       *DomainMemoryLocked       `xml:"locked" json:"memoryLocked" protobuf:"bytes,3,opt,name=memoryLocked"`
	MemorySource       *DomainMemorySource       `xml:"source" json:"memorySource" protobuf:"bytes,4,opt,name=memorySource"`
	MemoryAccess       *DomainMemoryAccess       `xml:"access" json:"memoryAccess" protobuf:"bytes,5,opt,name=memoryAccess"`
	MemoryAllocation   *DomainMemoryAllocation   `xml:"allocation" json:"memoryAllocation" protobuf:"bytes,6,opt,name=memoryAllocation"`
	MemoryDiscard      *DomainMemoryDiscard      `xml:"discard" json:"memoryDiscard" protobuf:"bytes,7,opt,name=memoryDiscard"`
}

// +gogo:genproto=true
type DomainOSType struct {
	Arch    string `xml:"arch,attr,omitempty" json:"arch,omitempty" protobuf:"bytes,1,opt,name=arch"`
	Machine string `xml:"machine,attr,omitempty" json:"machine,omitempty" protobuf:"bytes,2,opt,name=machine"`
	Type    string `xml:",chardata" json:"type" protobuf:"bytes,3,opt,name=type"`
}

// +gogo:genproto=true
type DomainSMBios struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type DomainNVRam struct {
	NVRam    string `xml:",chardata" json:"nvram" protobuf:"bytes,1,opt,name=nvram"`
	Template string `xml:"template,attr,omitempty" json:"template,omitempty" protobuf:"bytes,2,opt,name=template"`
}

// +gogo:genproto=true
type DomainBootDevice struct {
	Dev string `xml:"dev,attr" json:"dev" protobuf:"bytes,1,opt,name=dev"`
}

// +gogo:genproto=true
type DomainBootMenu struct {
	Enable  string `xml:"enable,attr,omitempty" json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`
	Timeout string `xml:"timeout,attr,omitempty" json:"timeout,omitempty" protobuf:"bytes,2,opt,name=timeout"`
}

// +gogo:genproto=true
type DomainSysInfoBIOS struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoSystem struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoBaseBoard struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoProcessor struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoMemory struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoChassis struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoOEMStrings struct {
	Entry []string `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfoSMBIOS struct {
	BIOS       *DomainSysInfoBIOS       `xml:"bios" json:"bios,omitempty" protobuf:"bytes,1,opt,name=bios"`
	System     *DomainSysInfoSystem     `xml:"system" json:"system,omitempty" protobuf:"bytes,2,opt,name=system"`
	BaseBoard  []DomainSysInfoBaseBoard `xml:"baseBoard" json:"baseBoard" protobuf:"bytes,3,rep,name=baseBoard"`
	Chassis    *DomainSysInfoChassis    `xml:"chassis" json:"chassis,omitempty" protobuf:"bytes,4,opt,name=chassis"`
	Processor  []DomainSysInfoProcessor `xml:"processor" json:"processor" protobuf:"bytes,5,rep,name=processor"`
	Memory     []DomainSysInfoMemory    `xml:"memory" json:"memory" protobuf:"bytes,6,rep,name=memory"`
	OEMStrings *DomainSysInfoOEMStrings `xml:"oemStrings" json:"oemStrings,omitempty" protobuf:"bytes,7,opt,name=oemStrings"`
}

// +gogo:genproto=true
type DomainSysInfoFWCfg struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" protobuf:"bytes,1,rep,name=entry"`
}

// +gogo:genproto=true
type DomainSysInfo struct {
	SMBIOS *DomainSysInfoSMBIOS `xml:"-" json:"smbios,omitempty" protobuf:"bytes,1,opt,name=smbios"`
	FWCfg  *DomainSysInfoFWCfg  `xml:"-" json:"fwcfg,omitempty" protobuf:"bytes,2,opt,name=fwcfg"`
}

// +gogo:genproto=true
type DomainSysInfoEntry struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	File  string `xml:"file,attr,omitempty" json:"file,omitempty" protobuf:"bytes,2,opt,name=file"`
	Value string `xml:",chardata" json:"value" protobuf:"bytes,3,opt,name=value"`
}

// +gogo:genproto=true
type DomainBIOS struct {
	UseSerial     string `xml:"useserial,attr,omitempty" json:"useSerial,omitempty" protobuf:"bytes,1,opt,name=useSerial"`
	RebootTimeout *int32 `xml:"rebootTimeout,attr" json:"rebootTimeout" protobuf:"varint,2,opt,name=rebootTimeout"`
}

// +gogo:genproto=true
type DomainLoader struct {
	Path     string `xml:",chardata" json:"path" protobuf:"bytes,1,opt,name=path"`
	Readonly string `xml:"readonly,attr,omitempty" json:"readonly,omitempty" protobuf:"bytes,2,opt,name=readonly"`
	Secure   string `xml:"secure,attr,omitempty" json:"secure,omitempty" protobuf:"bytes,3,opt,name=secure"`
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,4,opt,name=type"`
}

// +gogo:genproto=true
type DomainACPI struct {
	Tables []DomainACPITable `xml:"table" json:"tables" protobuf:"bytes,1,rep,name=tables"`
}

// +gogo:genproto=true
type DomainACPITable struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Path string `xml:",chardata" json:"path" protobuf:"bytes,2,opt,name=path"`
}

// +gogo:genproto=true
type DomainOSInitEnv struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:",chardata" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainOSFirmwareInfo struct {
	Features []DomainOSFirmwareFeature `xml:"feature" json:"features" protobuf:"bytes,1,rep,name=features"`
}

// +gogo:genproto=true
type DomainOSFirmwareFeature struct {
	Enabled string `xml:"enabled,attr,omitempty" json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
	Name    string `xml:"name,attr,omitempty" json:"name" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type DomainOS struct {
	Type         *DomainOSType         `xml:"type" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Firmware     string                `xml:"firmware,attr,omitempty" json:"firmware,omitempty" protobuf:"bytes,2,opt,name=firmware"`
	FirmwareInfo *DomainOSFirmwareInfo `xml:"firmware" json:"firmwareInfo,omitempty" protobuf:"bytes,3,opt,name=firmwareInfo"`
	Init         string                `xml:"init,omitempty" json:"init,omitempty" protobuf:"bytes,4,opt,name=init"`
	InitArgs     []string              `xml:"initarg" json:"initargs" protobuf:"bytes,5,rep,name=initargs"`
	InitEnv      []DomainOSInitEnv     `xml:"initenv" json:"initenv" protobuf:"bytes,6,rep,name=initenv"`
	InitDir      string                `xml:"initdir,omitempty" json:"initdir,omitempty" protobuf:"bytes,7,opt,name=initdir"`
	InitUser     string                `xml:"inituser,omitempty" json:"inituser,omitempty" protobuf:"bytes,8,opt,name=inituser"`
	InitGroup    string                `xml:"initgroup,omitempty" json:"initgroup,omitempty" protobuf:"bytes,9,opt,name=initgroup"`
	Loader       *DomainLoader         `xml:"loader" json:"loader,omitempty" protobuf:"bytes,10,opt,name=loader"`
	NVRam        *DomainNVRam          `xml:"nvram" json:"nvram,omitempty" protobuf:"bytes,11,opt,name=nvram"`
	Kernel       string                `xml:"kernel,omitempty" json:"kernel,omitempty" protobuf:"bytes,12,opt,name=kernel"`
	Initrd       string                `xml:"initrd,omitempty" json:"initrd,omitempty" protobuf:"bytes,13,opt,name=initrd"`
	Cmdline      string                `xml:"cmdline,omitempty" json:"cmdline,omitempty" protobuf:"bytes,14,opt,name=cmdline"`
	DTB          string                `xml:"dtb,omitempty" json:"dtb,omitempty" protobuf:"bytes,15,opt,name=dtb"`
	ACPI         *DomainACPI           `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,16,opt,name=acpi"`
	BootDevices  []DomainBootDevice    `xml:"boot" json:"bootDevices" protobuf:"bytes,17,rep,name=bootDevices"`
	BootMenu     *DomainBootMenu       `xml:"bootmenu" json:"bootMenu,omitempty" protobuf:"bytes,18,opt,name=bootMenu"`
	BIOS         *DomainBIOS           `xml:"bios" json:"bios,omitempty" protobuf:"bytes,19,opt,name=bios"`
	SMBios       *DomainSMBios         `xml:"smbios" json:"smbios,omitempty" protobuf:"bytes,20,opt,name=smbios"`
}

// +gogo:genproto=true
type DomainResource struct {
	Partition string `xml:"partition,omitempty" json:"partition,omitempty" protobuf:"bytes,1,opt,name=partition"`
}

// +gogo:genproto=true
type DomainVCPU struct {
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty" protobuf:"bytes,1,opt,name=placement"`
	CPUSet    string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty" protobuf:"bytes,2,opt,name=cpuset"`
	Current   int32  `xml:"current,attr,omitempty" json:"current,omitempty" protobuf:"varint,3,opt,name=current"`
	Value     int32  `xml:",chardata" json:"value" protobuf:"varint,4,opt,name=value"`
}

// +gogo:genproto=true
type DomainVCPUsVCPU struct {
	Id           *int32 `xml:"id,attr" json:"id,omitempty" protobuf:"varint,1,opt,name=id"`
	Enabled      string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" protobuf:"bytes,2,opt,name=enabled"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty" json:"hotpluggable,omitempty" protobuf:"bytes,3,opt,name=hotpluggable"`
	Order        *int32 `xml:"order,attr" json:"order,omitempty" protobuf:"varint,4,opt,name=order"`
}

// +gogo:genproto=true
type DomainVCPUs struct {
	VCPU []DomainVCPUsVCPU `xml:"vcpu" json:"vcpu" protobuf:"bytes,1,rep,name=vcpu"`
}

// +gogo:genproto=true
type DomainCPUModel struct {
	Fallback string `xml:"fallback,attr,omitempty" json:"fallback,omitempty" protobuf:"bytes,1,opt,name=fallback"`
	Value    string `xml:",chardata" json:"value" protobuf:"bytes,2,opt,name=value"`
	VendorID string `xml:"vendor_id,attr,omitempty" json:"vendorId,omitempty" protobuf:"bytes,3,opt,name=vendorId"`
}

// +gogo:genproto=true
type DomainCPUTopology struct {
	Sockets int `xml:"sockets,attr,omitempty" json:"sockets,omitempty" protobuf:"varint,1,opt,name=sockets"`
	Dies    int `xml:"dies,attr,omitempty" json:"dies,omitempty" protobuf:"varint,2,opt,name=dies"`
	Cores   int `xml:"cores,attr,omitempty" json:"cores,omitempty" protobuf:"varint,3,opt,name=cores"`
	Threads int `xml:"threads,attr,omitempty" json:"threads,omitempty" protobuf:"varint,4,opt,name=threads"`
}

// +gogo:genproto=true
type DomainCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty" json:"policy,omitempty" protobuf:"bytes,1,opt,name=policy"`
	Name   string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type DomainCPUCache struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty" protobuf:"varint,1,opt,name=level"`
	Mode  string `xml:"mode,attr" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type DomainCPU struct {
	XMLName    xml.Name           `xml:"cpu" json:"-"`
	Match      string             `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	Mode       string             `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
	Check      string             `xml:"check,attr,omitempty" json:"check,omitempty" protobuf:"bytes,3,opt,name=check"`
	Migratable string             `xml:"migratable,attr,omitempty" json:"migratable,omitempty" protobuf:"bytes,4,opt,name=migratable"`
	Model      *DomainCPUModel    `xml:"model" json:"model,omitempty" protobuf:"bytes,5,opt,name=model"`
	Vendor     string             `xml:"vendor,omitempty" json:"vendor" protobuf:"bytes,6,opt,name=vendor"`
	Topology   *DomainCPUTopology `xml:"topology" json:"topology,omitempty" protobuf:"bytes,7,opt,name=topology"`
	Cache      *DomainCPUCache    `xml:"cache" json:"cache,omitempty" protobuf:"bytes,8,opt,name=cache"`
	Features   []DomainCPUFeature `xml:"feature" json:"features" protobuf:"bytes,9,rep,name=features"`
	Numa       *DomainNuma        `xml:"numa" json:"numa,omitempty" protobuf:"bytes,10,opt,name=numa"`
}

// +gogo:genproto=true
type DomainNuma struct {
	Cell          []DomainCell             `xml:"cell" json:"cell" protobuf:"bytes,1,rep,name=cell"`
	Interconnects *DomainNUMAInterconnects `xml:"interconnects" json:"interconnects,omitempty" protobuf:"bytes,2,opt,name=interconnects"`
}

// +gogo:genproto=true
type DomainCell struct {
	ID        *int32               `xml:"id,attr" json:"id,omitempty" protobuf:"varint,1,opt,name=id"`
	CPUs      string               `xml:"cpus,attr,omitempty" json:"cpus" protobuf:"bytes,2,opt,name=cpus"`
	Memory    int32                `xml:"memory,attr" json:"memory" protobuf:"varint,3,opt,name=memory"`
	Unit      string               `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,4,opt,name=unit"`
	MemAccess string               `xml:"memAccess,attr,omitempty" json:"memAccess,omitempty" protobuf:"bytes,5,opt,name=memAccess"`
	Discard   string               `xml:"discard,attr,omitempty" json:"discard,omitempty" protobuf:"bytes,6,opt,name=discard"`
	Distances *DomainCellDistances `xml:"distances" json:"distances,omitempty" protobuf:"bytes,7,opt,name=distances"`
	Caches    []DomainCellCache    `xml:"cache" json:"caches" protobuf:"bytes,8,rep,name=caches"`
}

// +gogo:genproto=true
type DomainCellDistances struct {
	Siblings []DomainCellSibling `xml:"sibling" json:"siblings" protobuf:"bytes,1,rep,name=siblings"`
}

// +gogo:genproto=true
type DomainCellSibling struct {
	ID    int32 `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Value int32 `xml:"value,attr" json:"value" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainCellCache struct {
	Level         int32               `xml:"level,attr" json:"level" protobuf:"varint,1,opt,name=level"`
	Associativity string              `xml:"associativity,attr" json:"associativity" protobuf:"bytes,2,opt,name=associativity"`
	Policy        string              `xml:"policy,attr" json:"policy" protobuf:"bytes,3,opt,name=policy"`
	Size_         DomainCellCacheSize `xml:"size" json:"size" protobuf:"bytes,4,opt,name=size"`
	Line          DomainCellCacheLine `xml:"line" json:"line" protobuf:"bytes,5,opt,name=line"`
}

// +gogo:genproto=true
type DomainCellCacheSize struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainCellCacheLine struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainNUMAInterconnects struct {
	Latencies  []DomainNUMAInterconnectLatency   `xml:"latency" json:"latencies" protobuf:"bytes,1,rep,name=latencies"`
	Bandwidths []DomainNUMAInterconnectBandwidth `xml:"bandwidth" json:"bandwidths" protobuf:"bytes,2,rep,name=bandwidths"`
}

// +gogo:genproto=true
type DomainNUMAInterconnectLatency struct {
	Initiator int32  `xml:"initiator,attr" json:"initiator" protobuf:"varint,1,opt,name=initiator"`
	Target    int32  `xml:"target,attr" json:"target" protobuf:"varint,2,opt,name=target"`
	Cache     int32  `xml:"cache,attr,omitempty" json:"cache,omitempty" protobuf:"varint,3,opt,name=cache"`
	Type      string `xml:"type,attr" json:"type" protobuf:"bytes,4,opt,name=type"`
	Value     int32  `xml:"value,attr" json:"value" protobuf:"varint,5,opt,name=value"`
}

// +gogo:genproto=true
type DomainNUMAInterconnectBandwidth struct {
	Initiator int32  `xml:"initiator,attr" json:"initiator" protobuf:"varint,1,opt,name=initiator"`
	Target    int32  `xml:"target,attr" json:"target" protobuf:"varint,2,opt,name=target"`
	Type      string `xml:"type,attr" json:"type" protobuf:"bytes,3,opt,name=type"`
	Value     int32  `xml:"value,attr" json:"value" protobuf:"varint,4,opt,name=value"`
	Unit      string `xml:"unit,attr" json:"unit" protobuf:"bytes,5,opt,name=unit"`
}

// +gogo:genproto=true
type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty" json:"offset,omitempty" protobuf:"bytes,1,opt,name=offset"`
	Basis      string        `xml:"basis,attr,omitempty" json:"basis,omitempty" protobuf:"bytes,2,opt,name=basis"`
	Adjustment string        `xml:"adjustment,attr,omitempty" json:"adjustment,omitempty" protobuf:"bytes,3,opt,name=adjustment"`
	TimeZone   string        `xml:"timezone,attr,omitempty" json:"time_zone,omitempty" protobuf:"bytes,4,opt,name=time_zone,json=timeZone"`
	Timer      []DomainTimer `xml:"timer" json:"timer" protobuf:"bytes,5,rep,name=timer"`
}

// +gogo:genproto=true
type DomainTimer struct {
	Name       string              `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Track      string              `xml:"track,attr,omitempty" json:"track,omitempty" protobuf:"bytes,2,opt,name=track"`
	TickPolicy string              `xml:"tickpolicy,attr,omitempty" json:"tickpolicy,omitempty" protobuf:"bytes,3,opt,name=tickpolicy"`
	CatchUp    *DomainTimerCatchUp `xml:"catchup" json:"catchup,omitempty" protobuf:"bytes,4,opt,name=catchup"`
	Frequency  int64               `xml:"frequency,attr,omitempty" json:"frequency,omitempty" protobuf:"varint,5,opt,name=frequency"`
	Mode       string              `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,6,opt,name=mode"`
	Present    string              `xml:"present,attr,omitempty" json:"present,omitempty" protobuf:"bytes,7,opt,name=present"`
}

// +gogo:genproto=true
type DomainTimerCatchUp struct {
	Threshold int32 `xml:"threshold,attr,omitempty" json:"threshold,omitempty" protobuf:"varint,1,opt,name=threshold"`
	Slew      int32 `xml:"slew,attr,omitempty" json:"slew,omitempty" protobuf:"varint,2,opt,name=slew"`
	Limit     int32 `xml:"limit,attr,omitempty" json:"limit,omitempty" protobuf:"varint,3,opt,name=limit"`
}

// +gogo:genproto=true
type DomainFeature struct {
}

// +gogo:genproto=true
type DomainFeatureState struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`
}

// +gogo:genproto=true
type DomainFeatureAPIC struct {
	EOI string `xml:"eoi,attr,omitempty" json:"eoi,omitempty" protobuf:"bytes,1,opt,name=eoi"`
}

// +gogo:genproto=true
type DomainFeatureHyperVVendorId struct {
	DomainFeatureState `json:",inline" protobuf:"bytes,1,opt,name=domainFeatureState"`
	Value              string `xml:"value,attr,omitempty" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureHyperVSpinlocks struct {
	DomainFeatureState `json:",inline" protobuf:"bytes,1,opt,name=domainFeatureState"`
	Retries            int32 `xml:"retries,attr,omitempty" json:"retries,omitempty" protobuf:"varint,2,opt,name=retries"`
}

// +gogo:genproto=true
type DomainFeatureHyperVSTimer struct {
	DomainFeatureState `json:",inline" protobuf:"bytes,1,opt,name=domainFeatureState"`
	Direct             *DomainFeatureState `xml:"direct" json:"direct,omitempty" protobuf:"bytes,2,opt,name=direct"`
}

// +gogo:genproto=true
type DomainFeatureHyperV struct {
	DomainFeature   `json:",inline" protobuf:"bytes,1,opt,name=domainFeature"`
	Relaxed         *DomainFeatureState           `xml:"relaxed" json:"relaxed,omitempty" protobuf:"bytes,2,opt,name=relaxed"`
	VAPIC           *DomainFeatureState           `xml:"vapic" json:"vapic,omitempty" protobuf:"bytes,3,opt,name=vapic"`
	Spinlocks       *DomainFeatureHyperVSpinlocks `xml:"spinlocks" json:"spinlocks,omitempty" protobuf:"bytes,4,opt,name=spinlocks"`
	VPIndex         *DomainFeatureState           `xml:"vpindex" json:"vpindex,omitempty" protobuf:"bytes,5,opt,name=vpindex"`
	Runtime         *DomainFeatureState           `xml:"runtime" json:"runtime,omitempty" protobuf:"bytes,6,opt,name=runtime"`
	Synic           *DomainFeatureState           `xml:"synic" json:"synic,omitempty" protobuf:"bytes,7,opt,name=synic"`
	STimer          *DomainFeatureHyperVSTimer    `xml:"stimer" json:"stimer,omitempty" protobuf:"bytes,8,opt,name=stimer"`
	Reset_          *DomainFeatureState           `xml:"reset" json:"reset,omitempty" protobuf:"bytes,9,opt,name=reset"`
	VendorId        *DomainFeatureHyperVVendorId  `xml:"vendor_id" json:"vendorId,omitempty" protobuf:"bytes,10,opt,name=vendorId"`
	Frequencies     *DomainFeatureState           `xml:"frequencies" json:"frequencies,omitempty" protobuf:"bytes,11,opt,name=frequencies"`
	ReEnlightenment *DomainFeatureState           `xml:"reenlightenment" json:"reenlightenment,omitempty" protobuf:"bytes,12,opt,name=reenlightenment"`
	TLBFlush        *DomainFeatureState           `xml:"tlbflush" json:"tlb_flush,omitempty" protobuf:"bytes,13,opt,name=tlb_flush,json=tlbFlush"`
	IPI             *DomainFeatureState           `xml:"ipi" json:"ipi,omitempty" protobuf:"bytes,14,opt,name=ipi"`
	EVMCS           *DomainFeatureState           `xml:"evmcs" json:"evmcs,omitempty" protobuf:"bytes,15,opt,name=evmcs"`
}

// +gogo:genproto=true
type DomainFeatureKVM struct {
	Hidden        *DomainFeatureState `xml:"hidden" json:"hidden,omitempty" protobuf:"bytes,1,opt,name=hidden"`
	HintDedicated *DomainFeatureState `xml:"hint-dedicated" json:"hintDedicated,omitempty" protobuf:"bytes,2,opt,name=hintDedicated"`
	PollControl   *DomainFeatureState `xml:"poll-control" json:"pollControl,omitempty" protobuf:"bytes,3,opt,name=pollControl"`
}

// +gogo:genproto=true
type DomainFeatureXenPassthrough struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`
	Mode  string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type DomainFeatureXenE820Host struct {
	State string `xml:"state,attr" json:"state" protobuf:"bytes,1,opt,name=state"`
}

// +gogo:genproto=true
type DomainFeatureXen struct {
	E820Host    *DomainFeatureXenE820Host    `xml:"e820_host" json:"e820Host,omitempty" protobuf:"bytes,1,opt,name=e820Host"`
	Passthrough *DomainFeatureXenPassthrough `xml:"passthrough" json:"passthrough,omitempty" protobuf:"bytes,2,opt,name=passthrough"`
}

// +gogo:genproto=true
type DomainFeatureGIC struct {
	Version string `xml:"version,attr,omitempty" json:"version,omitempty" protobuf:"bytes,1,opt,name=version"`
}

// +gogo:genproto=true
type DomainFeatureIOAPIC struct {
	Driver string `xml:"driver,attr,omitempty" json:"driver,omitempty" protobuf:"bytes,1,opt,name=driver"`
}

// +gogo:genproto=true
type DomainFeatureHPT struct {
	Resizing    string                    `xml:"resizing,attr,omitempty" json:"resizing,omitempty" protobuf:"bytes,1,opt,name=resizing"`
	MaxPageSize *DomainFeatureHPTPageSize `xml:"maxpagesize" json:"maxpagesize,omitempty" protobuf:"bytes,2,opt,name=maxpagesize"`
}

// +gogo:genproto=true
type DomainFeatureHPTPageSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,1,opt,name=unit"`
	Value string `xml:",chardata" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureSMM struct {
	State string                `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`
	TSeg  *DomainFeatureSMMTSeg `xml:"tseg" json:"tseg,omitempty" protobuf:"bytes,2,opt,name=tseg"`
}

// +gogo:genproto=true
type DomainFeatureSMMTSeg struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,1,opt,name=unit"`
	Value int32  `xml:",chardata" json:"value" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureCapability struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`
}

// +gogo:genproto=true
type DomainLaunchSecurity struct {
	SEV *DomainLaunchSecuritySEV `xml:"-" json:"sev,omitempty" protobuf:"bytes,1,opt,name=sev"`
}

// +gogo:genproto=true
type DomainLaunchSecuritySEV struct {
	CBitPos         *int32 `xml:"cbitpos" json:"cBitPos,omitempty" protobuf:"varint,1,opt,name=cBitPos"`
	ReducedPhysBits *int32 `xml:"reducedPhysBits" json:"reducedPhysBits,omitempty" protobuf:"varint,2,opt,name=reducedPhysBits"`
	Policy          *int32 `xml:"policy" json:"policy,omitempty" protobuf:"varint,3,opt,name=policy"`
	DHCert          string `xml:"dhCert" json:"dh_cert" protobuf:"bytes,4,opt,name=dh_cert,json=dhCert"`
	Session         string `xml:"session" json:"session" protobuf:"bytes,5,opt,name=session"`
}

// +gogo:genproto=true
type DomainFeatureCapabilities struct {
	Policy         string                   `xml:"policy,attr,omitempty" json:"policy,omitempty" protobuf:"bytes,1,opt,name=policy"`
	AuditControl   *DomainFeatureCapability `xml:"audit_control" json:"auditControl,omitempty" protobuf:"bytes,2,opt,name=auditControl"`
	AuditWrite     *DomainFeatureCapability `xml:"audit_write" json:"auditWrite,omitempty" protobuf:"bytes,3,opt,name=auditWrite"`
	BlockSuspend   *DomainFeatureCapability `xml:"block_suspend" json:"blockSuspend,omitempty" protobuf:"bytes,4,opt,name=blockSuspend"`
	Chown          *DomainFeatureCapability `xml:"chown" json:"chown,omitempty" protobuf:"bytes,5,opt,name=chown"`
	DACOverride    *DomainFeatureCapability `xml:"dac_override" json:"dacOverride,omitempty" protobuf:"bytes,6,opt,name=dacOverride"`
	DACReadSearch  *DomainFeatureCapability `xml:"dac_read_Search" json:"dacReadSearch,omitempty" protobuf:"bytes,7,opt,name=dacReadSearch"`
	FOwner         *DomainFeatureCapability `xml:"fowner" json:"fowner,omitempty" protobuf:"bytes,8,opt,name=fowner"`
	FSetID         *DomainFeatureCapability `xml:"fsetid" json:"fsetid,omitempty" protobuf:"bytes,9,opt,name=fsetid"`
	IPCLock        *DomainFeatureCapability `xml:"ipc_lock" json:"ipcLock,omitempty" protobuf:"bytes,10,opt,name=ipcLock"`
	IPCOwner       *DomainFeatureCapability `xml:"ipc_owner" json:"ipcOwner,omitempty" protobuf:"bytes,11,opt,name=ipcOwner"`
	Kill           *DomainFeatureCapability `xml:"kill" json:"kill,omitempty" protobuf:"bytes,12,opt,name=kill"`
	Lease          *DomainFeatureCapability `xml:"lease" json:"lease,omitempty" protobuf:"bytes,13,opt,name=lease"`
	LinuxImmutable *DomainFeatureCapability `xml:"linux_immutable" json:"linuxImmutable,omitempty" protobuf:"bytes,14,opt,name=linuxImmutable"`
	MACAdmin       *DomainFeatureCapability `xml:"mac_admin" json:"macAdmin,omitempty" protobuf:"bytes,15,opt,name=macAdmin"`
	MACOverride    *DomainFeatureCapability `xml:"mac_override" json:"macOverride,omitempty" protobuf:"bytes,16,opt,name=macOverride"`
	MkNod          *DomainFeatureCapability `xml:"mknod" json:"mknod,omitempty" protobuf:"bytes,17,opt,name=mknod"`
	NetAdmin       *DomainFeatureCapability `xml:"net_admin" json:"netAdmin,omitempty" protobuf:"bytes,18,opt,name=netAdmin"`
	NetBindService *DomainFeatureCapability `xml:"net_bind_service" json:"netBindService,omitempty" protobuf:"bytes,19,opt,name=netBindService"`
	NetBroadcast   *DomainFeatureCapability `xml:"net_broadcast" json:"netBroadcast,omitempty" protobuf:"bytes,20,opt,name=netBroadcast"`
	NetRaw         *DomainFeatureCapability `xml:"net_raw" json:"netraw,omitempty" protobuf:"bytes,21,opt,name=netraw"`
	SetGID         *DomainFeatureCapability `xml:"setgid" json:"setgid,omitempty" protobuf:"bytes,22,opt,name=setgid"`
	SetFCap        *DomainFeatureCapability `xml:"setfcap" json:"setfcap,omitempty" protobuf:"bytes,23,opt,name=setfcap"`
	SetPCap        *DomainFeatureCapability `xml:"setpcap" json:"setpcap,omitempty" protobuf:"bytes,24,opt,name=setpcap"`
	SetUID         *DomainFeatureCapability `xml:"setuid" json:"setuid,omitempty" protobuf:"bytes,25,opt,name=setuid"`
	SysAdmin       *DomainFeatureCapability `xml:"sys_admin" json:"admin,omitempty" protobuf:"bytes,26,opt,name=admin"`
	SysBoot        *DomainFeatureCapability `xml:"sys_boot" json:"boot,omitempty" protobuf:"bytes,27,opt,name=boot"`
	SysChRoot      *DomainFeatureCapability `xml:"sys_chroot" json:"chroot,omitempty" protobuf:"bytes,28,opt,name=chroot"`
	SysModule      *DomainFeatureCapability `xml:"sys_module" json:"module,omitempty" protobuf:"bytes,29,opt,name=module"`
	SysNice        *DomainFeatureCapability `xml:"sys_nice" json:"nice,omitempty" protobuf:"bytes,30,opt,name=nice"`
	SysPAcct       *DomainFeatureCapability `xml:"sys_pacct" json:"pacct,omitempty" protobuf:"bytes,31,opt,name=pacct"`
	SysPTrace      *DomainFeatureCapability `xml:"sys_ptrace" json:"ptrace,omitempty" protobuf:"bytes,32,opt,name=ptrace"`
	SysRawIO       *DomainFeatureCapability `xml:"sys_rawio" json:"rawio,omitempty" protobuf:"bytes,33,opt,name=rawio"`
	SysResource    *DomainFeatureCapability `xml:"sys_resource" json:"resource,omitempty" protobuf:"bytes,34,opt,name=resource"`
	SysTime        *DomainFeatureCapability `xml:"sys_time" json:"time,omitempty" protobuf:"bytes,35,opt,name=time"`
	SysTTYConfig   *DomainFeatureCapability `xml:"sys_tty_config" json:"ttyConfig,omitempty" protobuf:"bytes,36,opt,name=ttyConfig"`
	SysLog         *DomainFeatureCapability `xml:"syslog" json:"sysLog,omitempty" protobuf:"bytes,37,opt,name=sysLog"`
	WakeAlarm      *DomainFeatureCapability `xml:"wake_alarm" json:"wakeAlarm,omitempty" protobuf:"bytes,38,opt,name=wakeAlarm"`
}

// +gogo:genproto=true
type DomainFeatureMSRS struct {
	Unknown string `xml:"unknown,attr" json:"unknown" protobuf:"bytes,1,opt,name=unknown"`
}

// +gogo:genproto=true
type DomainFeatureCFPC struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureSBBC struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureIBS struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainFeatureList struct {
	PAE          *DomainFeature             `xml:"pae" json:"pae,omitempty" protobuf:"bytes,1,opt,name=pae"`
	ACPI         *DomainFeature             `xml:"acpi" json:"acpi,omitempty" protobuf:"bytes,2,opt,name=acpi"`
	APIC         *DomainFeatureAPIC         `xml:"apic" json:"apic,omitempty" protobuf:"bytes,3,opt,name=apic"`
	HAP          *DomainFeatureState        `xml:"hap" json:"hap,omitempty" protobuf:"bytes,4,opt,name=hap"`
	Viridian     *DomainFeature             `xml:"viridian" json:"viridian,omitempty" protobuf:"bytes,5,opt,name=viridian"`
	PrivNet      *DomainFeature             `xml:"privnet" json:"privnet,omitempty" protobuf:"bytes,6,opt,name=privnet"`
	HyperV       *DomainFeatureHyperV       `xml:"hyperv" json:"hyperv,omitempty" protobuf:"bytes,7,opt,name=hyperv"`
	KVM          *DomainFeatureKVM          `xml:"kvm" json:"kvm,omitempty" protobuf:"bytes,8,opt,name=kvm"`
	Xen          *DomainFeatureXen          `xml:"xen" json:"xen,omitempty" protobuf:"bytes,9,opt,name=xen"`
	PVSpinlock   *DomainFeatureState        `xml:"pvspinlock" json:"pvspinlock,omitempty" protobuf:"bytes,10,opt,name=pvspinlock"`
	PMU          *DomainFeatureState        `xml:"pmu" json:"pmu,omitempty" protobuf:"bytes,11,opt,name=pmu"`
	VMPort       *DomainFeatureState        `xml:"vmport" json:"vmport,omitempty" protobuf:"bytes,12,opt,name=vmport"`
	GIC          *DomainFeatureGIC          `xml:"gic" json:"gic,omitempty" protobuf:"bytes,13,opt,name=gic"`
	SMM          *DomainFeatureSMM          `xml:"smm" json:"smm,omitempty" protobuf:"bytes,14,opt,name=smm"`
	IOAPIC       *DomainFeatureIOAPIC       `xml:"ioapic" json:"ioapic,omitempty" protobuf:"bytes,15,opt,name=ioapic"`
	HPT          *DomainFeatureHPT          `xml:"hpt" json:"hpt,omitempty" protobuf:"bytes,16,opt,name=hpt"`
	HTM          *DomainFeatureState        `xml:"htm" json:"htm,omitempty" protobuf:"bytes,17,opt,name=htm"`
	NestedHV     *DomainFeatureState        `xml:"nested-hv" json:"nestedHv,omitempty" protobuf:"bytes,18,opt,name=nestedHv"`
	Capabilities *DomainFeatureCapabilities `xml:"capabilities" json:"capabilities,omitempty" protobuf:"bytes,19,opt,name=capabilities"`
	VMCoreInfo   *DomainFeatureState        `xml:"vmcoreinfo" json:"vmCoreInfo,omitempty" protobuf:"bytes,20,opt,name=vmCoreInfo"`
	MSRS         *DomainFeatureMSRS         `xml:"msrs" json:"msrs,omitempty" protobuf:"bytes,21,opt,name=msrs"`
	CCFAssist    *DomainFeatureState        `xml:"ccf-assist" json:"ccfAssist,omitempty" protobuf:"bytes,22,opt,name=ccfAssist"`
	CFPC         *DomainFeatureCFPC         `xml:"cfpc" json:"cfpc,omitempty" protobuf:"bytes,23,opt,name=cfpc"`
	SBBC         *DomainFeatureSBBC         `xml:"sbbc" json:"sbbc,omitempty" protobuf:"bytes,24,opt,name=sbbc"`
	IBS          *DomainFeatureIBS          `xml:"ibs" json:"ibs,omitempty" protobuf:"bytes,25,opt,name=ibs"`
}

// +gogo:genproto=true
type DomainCPUTuneShares struct {
	Value int32 `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainCPUTunePeriod struct {
	Value int64 `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainCPUTuneQuota struct {
	Value int64 `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainCPUTuneVCPUPin struct {
	VCPU   int32  `xml:"vcpu,attr" json:"vcpu" protobuf:"varint,1,opt,name=vcpu"`
	CPUSet string `xml:"cpuset,attr" json:"cpu_set" protobuf:"bytes,2,opt,name=cpu_set,json=cpuSet"`
}

// +gogo:genproto=true
type DomainCPUTuneEmulatorPin struct {
	CPUSet string `xml:"cpuset,attr" json:"cpuset" protobuf:"bytes,1,opt,name=cpuset"`
}

// +gogo:genproto=true
type DomainCPUTuneIOThreadPin struct {
	IOThread int32  `xml:"iothread,attr" json:"iothread" protobuf:"varint,1,opt,name=iothread"`
	CPUSet   string `xml:"cpuset,attr" json:"cpuset" protobuf:"bytes,2,opt,name=cpuset"`
}

// +gogo:genproto=true
type DomainCPUTuneVCPUSched struct {
	VCPUs     string `xml:"vcpus,attr" json:"vcpus" protobuf:"bytes,1,opt,name=vcpus"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty" protobuf:"bytes,2,opt,name=scheduler"`
	Priority  *int32 `xml:"priority,attr" json:"priority,omitempty" protobuf:"varint,3,opt,name=priority"`
}

// +gogo:genproto=true
type DomainCPUTuneIOThreadSched struct {
	IOThreads string `xml:"iothreads,attr" json:"iothreads" protobuf:"bytes,1,opt,name=iothreads"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty" protobuf:"bytes,2,opt,name=scheduler"`
	Priority  *int32 `xml:"priority,attr" json:"priority" protobuf:"varint,3,opt,name=priority"`
}

// +gogo:genproto=true
type DomainCPUTuneEmulatorSched struct {
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty" protobuf:"bytes,1,opt,name=scheduler"`
	Priority  *int32 `xml:"priority,attr" json:"priority,omitempty" protobuf:"varint,2,opt,name=priority"`
}

// +gogo:genproto=true
type DomainCPUCacheTune struct {
	VCPUs   string                      `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty" protobuf:"bytes,1,opt,name=vcpus"`
	Cache   []DomainCPUCacheTuneCache   `xml:"cache" json:"cache" protobuf:"bytes,2,rep,name=cache"`
	Monitor []DomainCPUCacheTuneMonitor `xml:"monitor" json:"monitor" protobuf:"bytes,3,rep,name=monitor"`
}

// +gogo:genproto=true
type DomainCPUCacheTuneCache struct {
	ID    int32  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Level int32  `xml:"level,attr" json:"level" protobuf:"varint,2,opt,name=level"`
	Type  string `xml:"type,attr" json:"type" protobuf:"bytes,3,opt,name=type"`
	Size_ int32  `xml:"size,attr" json:"size" protobuf:"varint,4,opt,name=size"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,5,opt,name=unit"`
}

// +gogo:genproto=true
type DomainCPUCacheTuneMonitor struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty" protobuf:"varint,1,opt,name=level"`
	VCPUs string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty" protobuf:"bytes,2,opt,name=vcpus"`
}

// +gogo:genproto=true
type DomainCPUMemoryTune struct {
	VCPUs   string                       `xml:"vcpus,attr" json:"vcp_us" protobuf:"bytes,1,opt,name=vcp_us,json=vcpUs"`
	Nodes   []DomainCPUMemoryTuneNode    `xml:"node" json:"nodes" protobuf:"bytes,2,rep,name=nodes"`
	Monitor []DomainCPUMemoryTuneMonitor `xml:"monitor" json:"monitor" protobuf:"bytes,3,rep,name=monitor"`
}

// +gogo:genproto=true
type DomainCPUMemoryTuneNode struct {
	ID        int32 `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Bandwidth int32 `xml:"bandwidth,attr" json:"bandwidth" protobuf:"varint,2,opt,name=bandwidth"`
}

// +gogo:genproto=true
type DomainCPUMemoryTuneMonitor struct {
	Level int32  `xml:"level,attr,omitempty" json:"level,omitempty" protobuf:"varint,1,opt,name=level"`
	VCPUs string `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty" protobuf:"bytes,2,opt,name=vcpus"`
}

// +gogo:genproto=true
type DomainCPUTune struct {
	Shares         *DomainCPUTuneShares         `xml:"shares" json:"shares,omitempty" protobuf:"bytes,1,opt,name=shares"`
	Period         *DomainCPUTunePeriod         `xml:"period" json:"period,omitempty" protobuf:"bytes,2,opt,name=period"`
	Quota          *DomainCPUTuneQuota          `xml:"quota" json:"quota,omitempty" protobuf:"bytes,3,opt,name=quota"`
	GlobalPeriod   *DomainCPUTunePeriod         `xml:"global_period" json:"globalPeriod,omitempty" protobuf:"bytes,4,opt,name=globalPeriod"`
	GlobalQuota    *DomainCPUTuneQuota          `xml:"global_quota" json:"globalQuota,omitempty" protobuf:"bytes,5,opt,name=globalQuota"`
	EmulatorPeriod *DomainCPUTunePeriod         `xml:"emulator_period" json:"emulatorPeriod,omitempty" protobuf:"bytes,6,opt,name=emulatorPeriod"`
	EmulatorQuota  *DomainCPUTuneQuota          `xml:"emulator_quota" json:"emulatorQuota,omitempty" protobuf:"bytes,7,opt,name=emulatorQuota"`
	IOThreadPeriod *DomainCPUTunePeriod         `xml:"iothread_period" json:"iothreadPeriod,omitempty" protobuf:"bytes,8,opt,name=iothreadPeriod"`
	IOThreadQuota  *DomainCPUTuneQuota          `xml:"iothread_quota" json:"iothreadQuota,omitempty" protobuf:"bytes,9,opt,name=iothreadQuota"`
	VCPUPin        []DomainCPUTuneVCPUPin       `xml:"vcpupin" json:"vcpupin,omitempty" protobuf:"bytes,10,rep,name=vcpupin"`
	EmulatorPin    *DomainCPUTuneEmulatorPin    `xml:"emulatorpin" json:"emulatorpin,omitempty" protobuf:"bytes,11,opt,name=emulatorpin"`
	IOThreadPin    []DomainCPUTuneIOThreadPin   `xml:"iothreadpin" json:"iothreadpin,omitempty" protobuf:"bytes,12,rep,name=iothreadpin"`
	VCPUSched      []DomainCPUTuneVCPUSched     `xml:"vcpusched" json:"vcpusched,omitempty" protobuf:"bytes,13,rep,name=vcpusched"`
	EmulatorSched  *DomainCPUTuneEmulatorSched  `xml:"emulatorsched" json:"emulatorsched,omitempty" protobuf:"bytes,14,opt,name=emulatorsched"`
	IOThreadSched  []DomainCPUTuneIOThreadSched `xml:"iothreadsched" json:"iothreadsched,omitempty" protobuf:"bytes,15,rep,name=iothreadsched"`
	CacheTune      []DomainCPUCacheTune         `xml:"cachetune" json:"cachetune,omitempty" protobuf:"bytes,16,rep,name=cachetune"`
	MemoryTune     []DomainCPUMemoryTune        `xml:"memorytune" json:"memorytune,omitempty" protobuf:"bytes,17,rep,name=memorytune"`
}

// +gogo:genproto=true
type DomainQEMUCommandlineArg struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainQEMUCommandlineEnv struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainQEMUCommandline struct {
	XMLName xml.Name                   `xml:"http://libvirt.org/schemas/domain/qemu/1.0 commandline" json:"-"`
	Args    []DomainQEMUCommandlineArg `xml:"arg" json:"args" protobuf:"bytes,1,rep,name=args"`
	Envs    []DomainQEMUCommandlineEnv `xml:"env" json:"envs" protobuf:"bytes,2,rep,name=envs"`
}

// +gogo:genproto=true
type DomainQEMUCapabilitiesEntry struct {
	Name string `xml:"capability,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainQEMUCapabilities struct {
	XMLName xml.Name                      `xml:"http://libvirt.org/schemas/domain/qemu/1.0 capabilities" json:"-"`
	Add     []DomainQEMUCapabilitiesEntry `xml:"add" json:"add" protobuf:"bytes,1,rep,name=add"`
	Del     []DomainQEMUCapabilitiesEntry `xml:"del" json:"del" protobuf:"bytes,2,rep,name=del"`
}

// +gogo:genproto=true
type DomainQEMUDeprecation struct {
	XMLName  xml.Name `xml:"http://libvirt.org/schemas/domain/qemu/1.0 deprecation" json:"-"`
	Behavior string   `xml:"behavior,attr,omitempty" json:"behavior,omitempty" protobuf:"bytes,1,opt,name=behavior"`
}

// +gogo:genproto=true
type DomainLXCNamespace struct {
	XMLName  xml.Name               `xml:"http://libvirt.org/schemas/domain/lxc/1.0 namespace" json:"-"`
	ShareNet *DomainLXCNamespaceMap `xml:"sharenet" json:"sharenet,omitempty" protobuf:"bytes,1,opt,name=sharenet"`
	ShareIPC *DomainLXCNamespaceMap `xml:"shareipc" json:"shareipc,omitempty" protobuf:"bytes,2,opt,name=shareipc"`
	ShareUTS *DomainLXCNamespaceMap `xml:"shareuts" json:"shareuts,omitempty" protobuf:"bytes,3,opt,name=shareuts"`
}

// +gogo:genproto=true
type DomainLXCNamespaceMap struct {
	Type  string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainBHyveCommandlineArg struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainBHyveCommandlineEnv struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type DomainBHyveCommandline struct {
	XMLName xml.Name                    `xml:"http://libvirt.org/schemas/domain/bhyve/1.0 commandline" json:"-"`
	Args    []DomainBHyveCommandlineArg `xml:"arg" json:"args" protobuf:"bytes,1,rep,name=args"`
	Envs    []DomainBHyveCommandlineEnv `xml:"env" json:"envs" protobuf:"bytes,2,rep,name=envs"`
}

// +gogo:genproto=true
type DomainXenCommandlineArg struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainXenCommandline struct {
	XMLName xml.Name                  `xml:"http://libvirt.org/schemas/domain/xen/1.0 commandline" json:"-"`
	Args    []DomainXenCommandlineArg `xml:"arg" json:"args" protobuf:"bytes,1,rep,name=args"`
}

// +gogo:genproto=true
type DomainBlockIOTune struct {
	Weight int32                     `xml:"weight,omitempty" json:"weight" protobuf:"varint,1,opt,name=weight"`
	Device []DomainBlockIOTuneDevice `xml:"device" json:"device" protobuf:"bytes,2,rep,name=device"`
}

// +gogo:genproto=true
type DomainBlockIOTuneDevice struct {
	Path          string `xml:"path" json:"path" protobuf:"bytes,1,opt,name=path"`
	Weight        int32  `xml:"weight,omitempty" json:"weight,omitempty" protobuf:"varint,2,opt,name=weight"`
	ReadIopsSec   int32  `xml:"read_iops_sec,omitempty" json:"readIopsSec,omitempty" protobuf:"varint,3,opt,name=readIopsSec"`
	WriteIopsSec  int32  `xml:"write_iops_sec,omitempty" json:"writeIopsSec,omitempty" protobuf:"varint,4,opt,name=writeIopsSec"`
	ReadBytesSec  int32  `xml:"read_bytes_sec,omitempty" json:"readBytesSec,omitempty" protobuf:"varint,5,opt,name=readBytesSec"`
	WriteBytesSec int32  `xml:"write_bytes_sec,omitempty" json:"writeBytesSec,omitempty" protobuf:"varint,6,opt,name=writeBytesSec"`
}

// +gogo:genproto=true
type DomainPM struct {
	SuspendToMem  *DomainPMPolicy `xml:"suspend-to-mem" json:"suspendToMem,omitempty" protobuf:"bytes,1,opt,name=suspendToMem"`
	SuspendToDisk *DomainPMPolicy `xml:"suspend-to-disk" json:"suspendToDisk,omitempty" protobuf:"bytes,2,opt,name=suspendToDisk"`
}

// +gogo:genproto=true
type DomainPMPolicy struct {
	Enabled string `xml:"enabled,attr" json:"enabled" protobuf:"bytes,1,opt,name=enabled"`
}

// +gogo:genproto=true
type DomainSecLabel struct {
	Type       string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Model      string `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,2,opt,name=model"`
	Relabel    string `xml:"relabel,attr,omitempty" json:"relabel,omitempty" protobuf:"bytes,3,opt,name=relabel"`
	Label      string `xml:"label,omitempty" json:"label,omitempty" protobuf:"bytes,4,opt,name=label"`
	ImageLabel string `xml:"imagelabel,omitempty" json:"imageLabel,omitempty" protobuf:"bytes,5,opt,name=imageLabel"`
	BaseLabel  string `xml:"baselabel,omitempty" json:"baseLabel,omitempty" protobuf:"bytes,6,opt,name=baseLabel"`
}

// +gogo:genproto=true
type DomainDeviceSecLabel struct {
	Model     string `xml:"model,attr,omitempty" json:"model,omitempty" protobuf:"bytes,1,opt,name=model"`
	LabelSkip string `xml:"labelskip,attr,omitempty" json:"labelSkip,omitempty" protobuf:"bytes,2,opt,name=labelSkip"`
	Relabel   string `xml:"relabel,attr,omitempty" json:"relabel,omitempty" protobuf:"bytes,3,opt,name=relabel"`
	Label     string `xml:"label,omitempty" json:"label,omitempty" protobuf:"bytes,4,opt,name=label"`
}

// +gogo:genproto=true
type DomainNUMATune struct {
	Memory   *DomainNUMATuneMemory   `xml:"memory" json:"memory" protobuf:"bytes,1,opt,name=memory"`
	MemNodes []DomainNUMATuneMemNode `xml:"memnode" json:"memNodes" protobuf:"bytes,2,rep,name=memNodes"`
}

// +gogo:genproto=true
type DomainNUMATuneMemory struct {
	Mode      string `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
	Nodeset   string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty" protobuf:"bytes,2,opt,name=nodeset"`
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty" protobuf:"bytes,3,opt,name=placement"`
}

// +gogo:genproto=true
type DomainNUMATuneMemNode struct {
	CellID  int32  `xml:"cellid,attr" json:"cellId,omitempty" protobuf:"varint,1,opt,name=cellId"`
	Mode    string `xml:"mode,attr" json:"mode,omitempty" protobuf:"bytes,2,opt,name=mode"`
	Nodeset string `xml:"nodeset,attr" json:"nodeset,omitempty" protobuf:"bytes,3,opt,name=nodeset"`
}

// +gogo:genproto=true
type DomainIOThreadIDs struct {
	IOThreads []DomainIOThread `xml:"iothread" json:"iothreads" protobuf:"bytes,1,rep,name=iothreads"`
}

// +gogo:genproto=true
type DomainIOThread struct {
	ID int32 `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
}

// +gogo:genproto=true
type DomainKeyWrap struct {
	Ciphers []DomainKeyWrapCipher `xml:"cipher" json:"ciphers" protobuf:"bytes,1,rep,name=ciphers"`
}

// +gogo:genproto=true
type DomainKeyWrapCipher struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	State string `xml:"state,attr" json:"state" protobuf:"bytes,2,opt,name=state"`
}

// +gogo:genproto=true
type DomainIDMap struct {
	UIDs []DomainIDMapRange `xml:"uid" json:"uid" protobuf:"bytes,1,rep,name=uid"`
	GIDs []DomainIDMapRange `xml:"gid" json:"gid" protobuf:"bytes,2,rep,name=gid"`
}

// +gogo:genproto=true
type DomainIDMapRange struct {
	Start  int32 `xml:"start,attr" json:"start" protobuf:"varint,1,opt,name=start"`
	Target int32 `xml:"target,attr" json:"target" protobuf:"varint,2,opt,name=target"`
	Count  int32 `xml:"count,attr" json:"count" protobuf:"varint,3,opt,name=count"`
}

// +gogo:genproto=true
type DomainMemoryTuneLimit struct {
	Value int64  `xml:",chardata" json:"value" protobuf:"varint,1,opt,name=value"`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type DomainMemoryTune struct {
	HardLimit     *DomainMemoryTuneLimit `xml:"hard_limit" json:"hardLimit,omitempty" protobuf:"bytes,1,opt,name=hardLimit"`
	SoftLimit     *DomainMemoryTuneLimit `xml:"soft_limit" json:"softLimit,omitempty" protobuf:"bytes,2,opt,name=softLimit"`
	MinGuarantee  *DomainMemoryTuneLimit `xml:"min_guarantee" json:"minGuarantee,omitempty" protobuf:"bytes,3,opt,name=minGuarantee"`
	SwapHardLimit *DomainMemoryTuneLimit `xml:"swap_hard_limit" json:"swapHardLimit,omitempty" protobuf:"bytes,4,opt,name=swapHardLimit"`
}

// +gogo:genproto=true
type DomainMetadata struct {
	XML string `xml:",innerxml" json:",inline" protobuf:"bytes,1,opt,name=xML"`
}

// +gogo:genproto=true
type DomainVMWareDataCenterPath struct {
	XMLName xml.Name `xml:"http://libvirt.org/schemas/domain/vmware/1.0 datacenterpath" json:"-"`
	Value   string   `xml:",chardata" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// +gogo:genproto=true
type DomainPerf struct {
	Events []DomainPerfEvent `xml:"event" json:"events" protobuf:"bytes,1,rep,name=events"`
}

// +gogo:genproto=true
type DomainPerfEvent struct {
	Name    string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Enabled string `xml:"enabled,attr" json:"enabled" protobuf:"bytes,2,opt,name=enabled"`
}

// +gogo:genproto=true
type DomainGenID struct {
	Value string `xml:",chardata" json:"value" protobuf:"bytes,1,opt,name=value"`
}

// Domain NB, try to keep the order of fields in this struct
// matching the order of XML elements that libvirt
// will generate when dumping XML.
// +gogo:deepcopy-gen=true
// +gogo:genproto=true
type Domain struct {
	XMLName        xml.Name              `xml:"domain" json:"-"`
	Type           string                `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	ID             *int32                `xml:"id,attr" json:"id" protobuf:"varint,2,opt,name=id"`
	Name           string                `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	UUID           string                `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,4,opt,name=uuid"`
	GenID          *DomainGenID          `xml:"genid" json:"genId,omitempty" protobuf:"bytes,5,opt,name=genId"`
	Title          string                `xml:"title,omitempty" json:"title,omitempty" protobuf:"bytes,6,opt,name=title"`
	Description    string                `xml:"description,omitempty" json:"description,omitempty" protobuf:"bytes,7,opt,name=description"`
	Metadata       *DomainMetadata       `xml:"metadata" json:"metadata,omitempty" protobuf:"bytes,8,opt,name=metadata"`
	MaximumMemory  *DomainMaxMemory      `xml:"maxMemory" json:"maxMemory,omitempty" protobuf:"bytes,9,opt,name=maxMemory"`
	Memory         *DomainMemory         `xml:"memory" json:"memory,omitempty" protobuf:"bytes,10,opt,name=memory"`
	CurrentMemory  *DomainCurrentMemory  `xml:"currentMemory" json:"currentMemory,omitempty" protobuf:"bytes,11,opt,name=currentMemory"`
	BlockIOTune    *DomainBlockIOTune    `xml:"blkiotune" json:"blockIoTune,omitempty" protobuf:"bytes,12,opt,name=blockIoTune"`
	MemoryTune     *DomainMemoryTune     `xml:"memtune" json:"memTune,omitempty" protobuf:"bytes,13,opt,name=memTune"`
	MemoryBacking  *DomainMemoryBacking  `xml:"memoryBacking" json:"memoryBacking,omitempty" protobuf:"bytes,14,opt,name=memoryBacking"`
	VCPU           *DomainVCPU           `xml:"vcpu" json:"vcpu,omitempty" protobuf:"bytes,15,opt,name=vcpu"`
	VCPUs          *DomainVCPUs          `xml:"vcpus" json:"vcpus,omitempty" protobuf:"bytes,16,opt,name=vcpus"`
	IOThreads      int32                 `xml:"iothreads,omitempty" json:"iothreads,omitempty" protobuf:"varint,17,opt,name=iothreads"`
	IOThreadIDs    *DomainIOThreadIDs    `xml:"iothreadids" json:"iothreadids,omitempty" protobuf:"bytes,18,opt,name=iothreadids"`
	CPUTune        *DomainCPUTune        `xml:"cputune" json:"cputune,omitempty" protobuf:"bytes,19,opt,name=cputune"`
	NUMATune       *DomainNUMATune       `xml:"numatune" json:"numatune,omitempty" protobuf:"bytes,20,opt,name=numatune"`
	Resource       *DomainResource       `xml:"resource" json:"resource,omitempty" protobuf:"bytes,21,opt,name=resource"`
	SysInfo        []DomainSysInfo       `xml:"sysinfo" json:"sysinfo" protobuf:"bytes,22,rep,name=sysinfo"`
	Bootloader     string                `xml:"bootloader,omitempty" json:"bootloader,omitempty" protobuf:"bytes,23,opt,name=bootloader"`
	BootloaderArgs string                `xml:"bootloader_args,omitempty" json:"bootloaderArgs,omitempty" protobuf:"bytes,24,opt,name=bootloaderArgs"`
	OS             *DomainOS             `xml:"os" json:"os,omitempty" protobuf:"bytes,25,opt,name=os"`
	IDMap          *DomainIDMap          `xml:"idmap" json:"idMap,omitempty" protobuf:"bytes,26,opt,name=idMap"`
	Features       *DomainFeatureList    `xml:"features" json:"features,omitempty" protobuf:"bytes,27,opt,name=features"`
	CPU            *DomainCPU            `xml:"cpu" json:"cpu,omitempty" protobuf:"bytes,28,opt,name=cpu"`
	Clock          *DomainClock          `xml:"clock" json:"clock,omitempty" protobuf:"bytes,29,opt,name=clock"`
	OnPoweroff     string                `xml:"on_poweroff,omitempty" json:"onPoweroff,omitempty" protobuf:"bytes,30,opt,name=onPoweroff"`
	OnReboot       string                `xml:"on_reboot,omitempty" json:"onReboot,omitempty" protobuf:"bytes,31,opt,name=onReboot"`
	OnCrash        string                `xml:"on_crash,omitempty" json:"onCrash,omitempty" protobuf:"bytes,32,opt,name=onCrash"`
	PM             *DomainPM             `xml:"pm" json:"pm,omitempty" protobuf:"bytes,33,opt,name=pm"`
	Perf           *DomainPerf           `xml:"perf" json:"perf,omitempty" protobuf:"bytes,34,opt,name=perf"`
	Devices        *DomainDeviceList     `xml:"devices" json:"devices,omitempty" protobuf:"bytes,35,opt,name=devices"`
	SecLabel       []DomainSecLabel      `xml:"seclabel" json:"seclabel" protobuf:"bytes,36,rep,name=seclabel"`
	KeyWrap        *DomainKeyWrap        `xml:"keywrap" json:"keywrap,omitempty" protobuf:"bytes,37,opt,name=keywrap"`
	LaunchSecurity *DomainLaunchSecurity `xml:"launchSecurity" json:"launchSecurity,omitempty" protobuf:"bytes,38,opt,name=launchSecurity"`

	/* Hypervisor namespaces must all be last */
	QEMUCommandline      *DomainQEMUCommandline      `json:"qemuCommandline,omitempty" protobuf:"bytes,39,opt,name=qemuCommandline"`
	QEMUCapabilities     *DomainQEMUCapabilities     `json:"qemuCapabilities,omitempty" protobuf:"bytes,40,opt,name=qemuCapabilities"`
	QEMUDeprecation      *DomainQEMUDeprecation      `json:"qemuDeprecation,omitempty" protobuf:"bytes,41,opt,name=qemuDeprecation"`
	LXCNamespace         *DomainLXCNamespace         `json:"lxcNamespace,omitempty" protobuf:"bytes,42,opt,name=lxcNamespace"`
	BHyveCommandline     *DomainBHyveCommandline     `json:"bHyveCommandline,omitempty" protobuf:"bytes,43,opt,name=bHyveCommandline"`
	VMWareDataCenterPath *DomainVMWareDataCenterPath `json:"vmWareDataCenterPath,omitempty" protobuf:"bytes,44,opt,name=vmWareDataCenterPath"`
	XenCommandline       *DomainXenCommandline       `json:"xenCommandline,omitempty" protobuf:"bytes,45,opt,name=xenCommandline"`
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
