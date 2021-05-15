package spec

import "encoding/xml"

type StoragePoolSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value uint64 `xml:",chardata" json:"value,omitempty"`
}

type StoragePoolTargetPermissions struct {
	Owner string `xml:"owner,omitempty" json:"owner,omitempty"`
	Group string `xml:"group,omitempty" json:"group,omitempty"`
	Mode  string `xml:"mode,omitempty" json:"mode,omitempty"`
	Label string `xml:"label,omitempty" json:"label,omitempty"`
}

type StoragePoolTargetTimestamps struct {
	Atime string `xml:"atime" json:"atime,omitempty"`
	Mtime string `xml:"mtime" json:"mtime,omitempty"`
	Ctime string `xml:"ctime" json:"ctime,omitempty"`
}

type StoragePoolTarget struct {
	Path        string                        `xml:"path,omitempty" json:"path,omitempty"`
	Permissions *StoragePoolTargetPermissions `xml:"permissions" json:"permissions,omitempty"`
	Timestamps  *StoragePoolTargetTimestamps  `xml:"timestamps" json:"timestamps,omitempty"`
	Encryption  *StorageEncryption            `xml:"encryption" json:"encryption,omitempty"`
}

type StoragePoolSourceFormat struct {
	Type string `xml:"type,attr" json:"type,omitempty"`
}

type StoragePoolSourceProtocol struct {
	Version string `xml:"ver,attr" json:"version,omitempty"`
}

type StoragePoolSourceHost struct {
	Name string `xml:"name,attr" json:"name,omitempty"`
	Port string `xml:"port,attr,omitempty" json:"port,omitempty"`
}

type StoragePoolSourceDevice struct {
	Path          string                              `xml:"path,attr" json:"path,omitempty"`
	PartSeparator string                              `xml:"part_separator,attr,omitempty" json:"part_separator,omitempty"`
	FreeExtents   []StoragePoolSourceDeviceFreeExtent `xml:"freeExtent" json:"freeExtents,omitempty"`
}

type StoragePoolSourceDeviceFreeExtent struct {
	Start uint64 `xml:"start,attr" json:"start,omitempty"`
	End   uint64 `xml:"end,attr" json:"end,omitempty"`
}

type StoragePoolSourceAuthSecret struct {
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
}

type StoragePoolSourceAuth struct {
	Type     string                       `xml:"type,attr,omitempty" json:"type,omitempty"`
	Username string                       `xml:"username,attr,omitempty" json:"username,omitempty"`
	Secret   *StoragePoolSourceAuthSecret `xml:"secret,omitempty" json:"secret,omitempty"`
}

type StoragePoolSourceVendor struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type StoragePoolSourceProduct struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

type StoragePoolPCIAddress struct {
	Domain   *uint `xml:"domain,attr,omitempty" json:"domain,omitempty"`
	Bus      *uint `xml:"bus,attr,omitempty" json:"bus,omitempty"`
	Slot     *uint `xml:"slot,attr,omitempty" json:"slot,omitempty"`
	Function *uint `xml:"function,attr,omitempty" json:"function,omitempty"`
}

type StoragePoolSourceAdapterParentAddr struct {
	UniqueID uint64                 `xml:"unique_id,attr,omitempty" json:"uniqueId,omitempty"`
	Address  *StoragePoolPCIAddress `xml:"address,omitempty" json:"address,omitempty"`
}

type StoragePoolSourceAdapter struct {
	Type       string                              `xml:"type,attr,omitempty" json:"type,omitempty"`
	Name       string                              `xml:"name,attr,omitempty" json:"name,omitempty"`
	Parent     string                              `xml:"parent,attr,omitempty" json:"parent,omitempty"`
	Managed    string                              `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	WWNN       string                              `xml:"wwnn,attr,omitempty" json:"wwnn,omitempty"`
	WWPN       string                              `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty"`
	ParentAddr *StoragePoolSourceAdapterParentAddr `xml:"parentaddr" json:"parentAddr,omitempty"`
}

type StoragePoolSourceDir struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty"`
}

type StoragePoolSourceInitiator struct {
	IQN StoragePoolSourceInitiatorIQN `xml:"iqn,omitempty" json:"iqn,omitempty"`
}

type StoragePoolSourceInitiatorIQN struct {
	Name string `xml:"name,attr,omitempty,omitempty" json:"name,omitempty"`
}

type StoragePoolSource struct {
	Name      string                      `xml:"name,omitempty" json:"name,omitempty"`
	Dir       *StoragePoolSourceDir       `xml:"dir,omitempty" json:"dir,omitempty"`
	Host      []StoragePoolSourceHost     `xml:"host,omitempty" json:"host,omitempty"`
	Device    []StoragePoolSourceDevice   `xml:"device,omitempty" json:"device,omitempty"`
	Auth      *StoragePoolSourceAuth      `xml:"auth,omitempty" json:"auth,omitempty"`
	Vendor    *StoragePoolSourceVendor    `xml:"vendor,omitempty" json:"vendor,omitempty"`
	Product   *StoragePoolSourceProduct   `xml:"product,omitempty" json:"product,omitempty"`
	Format    *StoragePoolSourceFormat    `xml:"format,omitempty" json:"format,omitempty"`
	Protocol  *StoragePoolSourceProtocol  `xml:"protocol,omitempty" json:"protocol,omitempty"`
	Adapter   *StoragePoolSourceAdapter   `xml:"adapter,omitempty" json:"adapter,omitempty"`
	Initiator *StoragePoolSourceInitiator `xml:"initiator,omitempty" json:"initiator,omitempty"`
}

type StoragePoolRefreshVol struct {
	Allocation string `xml:"allocation,attr,omitempty" json:"allocation,omitempty"`
}

type StoragePoolRefresh struct {
	Volume StoragePoolRefreshVol `xml:"volume,omitempty" json:"volume,omitempty"`
}

type StoragePoolFeatures struct {
	COW StoragePoolFeatureCOW `xml:"cow,omitempty" json:"cow,omitempty"`
}

type StoragePoolFeatureCOW struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty"`
}

type StoragePool struct {
	XMLName    xml.Name             `xml:"pool"`
	Type       string               `xml:"type,attr"`
	Name       string               `xml:"name,omitempty"`
	UUID       string               `xml:"uuid,omitempty"`
	Allocation *StoragePoolSize     `xml:"allocation"`
	Capacity   *StoragePoolSize     `xml:"capacity"`
	Available  *StoragePoolSize     `xml:"available"`
	Features   *StoragePoolFeatures `xml:"features"`
	Target     *StoragePoolTarget   `xml:"target"`
	Source     *StoragePoolSource   `xml:"source"`
	Refresh    *StoragePoolRefresh  `xml:"refresh"`

	/* Pool backend namespcaes must be last */
	FSCommandline  *StoragePoolFSCommandline
	RBDCommandline *StoragePoolRBDCommandline
}

type StoragePoolFSCommandlineOption struct {
	Name string `xml:"name,attr"`
}

type StoragePoolFSCommandline struct {
	XMLName xml.Name                         `xml:"http://libvirt.org/schemas/storagepool/fs/1.0 mount_opts"`
	Options []StoragePoolFSCommandlineOption `xml:"option"`
}

type StoragePoolRBDCommandlineOption struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type StoragePoolRBDCommandline struct {
	XMLName xml.Name                          `xml:"http://libvirt.org/schemas/storagepool/rbd/1.0 config_opts"`
	Options []StoragePoolRBDCommandlineOption `xml:"option"`
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

func (s *StoragePool) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StoragePool) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}