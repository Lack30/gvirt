package spec

import "encoding/xml"

// +gogo:genproto=true
type StoragePoolSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,1,opt,name=unit"`
	Value uint64 `xml:",chardata" json:"value,omitempty" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type StoragePoolTargetPermissions struct {
	Owner string `xml:"owner,omitempty" json:"owner,omitempty" protobuf:"bytes,1,opt,name=owner"`
	Group string `xml:"group,omitempty" json:"group,omitempty" protobuf:"bytes,2,opt,name=group"`
	Mode  string `xml:"mode,omitempty" json:"mode,omitempty" protobuf:"bytes,3,opt,name=mode"`
	Label string `xml:"label,omitempty" json:"label,omitempty" protobuf:"bytes,4,opt,name=label"`
}

// +gogo:genproto=true
type StoragePoolTargetTimestamps struct {
	Atime string `xml:"atime" json:"atime,omitempty" protobuf:"bytes,1,opt,name=atime"`
	Mtime string `xml:"mtime" json:"mtime,omitempty" protobuf:"bytes,2,opt,name=mtime"`
	Ctime string `xml:"ctime" json:"ctime,omitempty" protobuf:"bytes,3,opt,name=ctime"`
}

// +gogo:genproto=true
type StoragePoolTarget struct {
	Path        string                        `xml:"path,omitempty" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	Permissions *StoragePoolTargetPermissions `xml:"permissions" json:"permissions,omitempty" protobuf:"bytes,2,opt,name=permissions"`
	Timestamps  *StoragePoolTargetTimestamps  `xml:"timestamps" json:"timestamps,omitempty" protobuf:"bytes,3,opt,name=timestamps"`
	Encryption  *StorageEncryption            `xml:"encryption" json:"encryption,omitempty" protobuf:"bytes,4,opt,name=encryption"`
}

// +gogo:genproto=true
type StoragePoolSourceFormat struct {
	Type string `xml:"type,attr" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type StoragePoolSourceProtocol struct {
	Version string `xml:"ver,attr" json:"version,omitempty" protobuf:"bytes,1,opt,name=version"`
}

// +gogo:genproto=true
type StoragePoolSourceHost struct {
	Name string `xml:"name,attr" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Port string `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"bytes,2,opt,name=port"`
}

// +gogo:genproto=true
type StoragePoolSourceDevice struct {
	Path          string                              `xml:"path,attr" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	PartSeparator string                              `xml:"part_separator,attr,omitempty" json:"part_separator,omitempty" protobuf:"bytes,2,opt,name=part_separator,json=partSeparator"`
	FreeExtents   []StoragePoolSourceDeviceFreeExtent `xml:"freeExtent" json:"freeExtents,omitempty" protobuf:"bytes,3,rep,name=freeExtents"`
}

// +gogo:genproto=true
type StoragePoolSourceDeviceFreeExtent struct {
	Start uint64 `xml:"start,attr" json:"start,omitempty" protobuf:"varint,1,opt,name=start"`
	End   uint64 `xml:"end,attr" json:"end,omitempty" protobuf:"varint,2,opt,name=end"`
}

// +gogo:genproto=true
type StoragePoolSourceAuthSecret struct {
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty" protobuf:"bytes,1,opt,name=usage"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty" protobuf:"bytes,2,opt,name=uuid"`
}

// +gogo:genproto=true
type StoragePoolSourceAuth struct {
	Type     string                       `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Username string                       `xml:"username,attr,omitempty" json:"username,omitempty" protobuf:"bytes,2,opt,name=username"`
	Secret   *StoragePoolSourceAuthSecret `xml:"secret,omitempty" json:"secret,omitempty" protobuf:"bytes,3,opt,name=secret"`
}

// +gogo:genproto=true
type StoragePoolSourceVendor struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type StoragePoolSourceProduct struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type StoragePoolPCIAddress struct {
	Domain   *uint `xml:"domain,attr,omitempty" json:"domain,omitempty" protobuf:"varint,1,opt,name=domain"`
	Bus      *uint `xml:"bus,attr,omitempty" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Slot     *uint `xml:"slot,attr,omitempty" json:"slot,omitempty" protobuf:"varint,3,opt,name=slot"`
	Function *uint `xml:"function,attr,omitempty" json:"function,omitempty" protobuf:"varint,4,opt,name=function"`
}

// +gogo:genproto=true
type StoragePoolSourceAdapterParentAddr struct {
	UniqueID uint64                 `xml:"unique_id,attr,omitempty" json:"uniqueId,omitempty" protobuf:"varint,1,opt,name=uniqueId"`
	Address  *StoragePoolPCIAddress `xml:"address,omitempty" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
}

// +gogo:genproto=true
type StoragePoolSourceAdapter struct {
	Type       string                              `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Name       string                              `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	Parent     string                              `xml:"parent,attr,omitempty" json:"parent,omitempty" protobuf:"bytes,3,opt,name=parent"`
	Managed    string                              `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,4,opt,name=managed"`
	WWNN       string                              `xml:"wwnn,attr,omitempty" json:"wwnn,omitempty" protobuf:"bytes,5,opt,name=wwnn"`
	WWPN       string                              `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty" protobuf:"bytes,6,opt,name=wwpn"`
	ParentAddr *StoragePoolSourceAdapterParentAddr `xml:"parentaddr" json:"parentAddr,omitempty" protobuf:"bytes,7,opt,name=parentAddr"`
}

// +gogo:genproto=true
type StoragePoolSourceDir struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
}

// +gogo:genproto=true
type StoragePoolSourceInitiator struct {
	IQN StoragePoolSourceInitiatorIQN `xml:"iqn,omitempty" json:"iqn,omitempty" protobuf:"bytes,1,opt,name=iqn"`
}

// +gogo:genproto=true
type StoragePoolSourceInitiatorIQN struct {
	Name string `xml:"name,attr,omitempty,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type StoragePoolSource struct {
	Name      string                      `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Dir       *StoragePoolSourceDir       `xml:"dir,omitempty" json:"dir,omitempty" protobuf:"bytes,2,opt,name=dir"`
	Host      []StoragePoolSourceHost     `xml:"host,omitempty" json:"host,omitempty" protobuf:"bytes,3,rep,name=host"`
	Device    []StoragePoolSourceDevice   `xml:"device,omitempty" json:"device,omitempty" protobuf:"bytes,4,rep,name=device"`
	Auth      *StoragePoolSourceAuth      `xml:"auth,omitempty" json:"auth,omitempty" protobuf:"bytes,5,opt,name=auth"`
	Vendor    *StoragePoolSourceVendor    `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,6,opt,name=vendor"`
	Product   *StoragePoolSourceProduct   `xml:"product,omitempty" json:"product,omitempty" protobuf:"bytes,7,opt,name=product"`
	Format    *StoragePoolSourceFormat    `xml:"format,omitempty" json:"format,omitempty" protobuf:"bytes,8,opt,name=format"`
	Protocol  *StoragePoolSourceProtocol  `xml:"protocol,omitempty" json:"protocol,omitempty" protobuf:"bytes,9,opt,name=protocol"`
	Adapter   *StoragePoolSourceAdapter   `xml:"adapter,omitempty" json:"adapter,omitempty" protobuf:"bytes,10,opt,name=adapter"`
	Initiator *StoragePoolSourceInitiator `xml:"initiator,omitempty" json:"initiator,omitempty" protobuf:"bytes,11,opt,name=initiator"`
}

// +gogo:genproto=true
type StoragePoolRefreshVol struct {
	Allocation string `xml:"allocation,attr,omitempty" json:"allocation,omitempty" protobuf:"bytes,1,opt,name=allocation"`
}

// +gogo:genproto=true
type StoragePoolRefresh struct {
	Volume StoragePoolRefreshVol `xml:"volume,omitempty" json:"volume,omitempty" protobuf:"bytes,1,opt,name=volume"`
}

// +gogo:genproto=true
type StoragePoolFeatures struct {
	COW StoragePoolFeatureCOW `xml:"cow,omitempty" json:"cow,omitempty" protobuf:"bytes,1,opt,name=cow"`
}

// +gogo:genproto=true
type StoragePoolFeatureCOW struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,1,opt,name=state"`
}

// +gogo:genproto=true
type StoragePool struct {
	XMLName    xml.Name             `xml:"pool" json:"-"`
	Type       string               `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Name       string               `xml:"name,omitempty" json:"name" protobuf:"bytes,2,opt,name=name"`
	UUID       string               `xml:"uuid,omitempty" json:"uuid" protobuf:"bytes,3,opt,name=uUID"`
	Allocation *StoragePoolSize     `xml:"allocation" json:"allocation,omitempty" protobuf:"bytes,4,opt,name=allocation"`
	Capacity   *StoragePoolSize     `xml:"capacity" json:"capacity,omitempty" protobuf:"bytes,5,opt,name=capacity"`
	Available  *StoragePoolSize     `xml:"available" json:"available,omitempty" protobuf:"bytes,6,opt,name=available"`
	Features   *StoragePoolFeatures `xml:"features" json:"features,omitempty" protobuf:"bytes,7,opt,name=features"`
	Target     *StoragePoolTarget   `xml:"target" json:"target,omitempty" protobuf:"bytes,8,opt,name=target"`
	Source     *StoragePoolSource   `xml:"source" json:"source,omitempty" protobuf:"bytes,9,opt,name=source"`
	Refresh    *StoragePoolRefresh  `xml:"refresh" json:"refresh,omitempty" protobuf:"bytes,10,opt,name=refresh"`

	/* Pool backend namespcaes must be last */
	FSCommandline  *StoragePoolFSCommandline  `json:"fSCommandline,omitempty" protobuf:"bytes,11,opt,name=fSCommandline"`
	RBDCommandline *StoragePoolRBDCommandline `json:"rBDCommandline,omitempty" protobuf:"bytes,12,opt,name=rBDCommandline"`
}

// +gogo:genproto=true
type StoragePoolFSCommandlineOption struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type StoragePoolFSCommandline struct {
	XMLName xml.Name                         `xml:"http://libvirt.org/schemas/storagepool/fs/1.0 mount_opts" json:"-"`
	Options []StoragePoolFSCommandlineOption `xml:"option" json:"options" protobuf:"bytes,1,rep,name=options"`
}

// +gogo:genproto=true
type StoragePoolRBDCommandlineOption struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type StoragePoolRBDCommandline struct {
	XMLName xml.Name                          `xml:"http://libvirt.org/schemas/storagepool/rbd/1.0 config_opts" json:"-"`
	Options []StoragePoolRBDCommandlineOption `xml:"option" json:"options" protobuf:"bytes,1,rep,name=options"`
}

func (a *StoragePoolPCIAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *StoragePoolPCIAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "domain" {
			if err := unmarshalUintAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalUintAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshalUintAttr(attr.Value, &a.Function, 0); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (s *StoragePool) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StoragePool) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
