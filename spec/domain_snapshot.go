package spec

import "encoding/xml"

// +gogo:genproto=true
type DomainSnapshotDisk struct {
	Name     string            `xml:"name,attr" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Snapshot string            `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty" protobuf:"bytes,2,opt,name=snapshot"`
	Driver   *DomainDiskDriver `xml:"driver" json:"driver,omitempty" protobuf:"bytes,3,opt,name=driver"`
	Source   *DomainDiskSource `xml:"source" json:"source,omitempty" protobuf:"bytes,4,opt,name=source"`
}

// +gogo:genproto=true
type DomainSnapshotDisks struct {
	Disks []DomainSnapshotDisk `xml:"disk" json:"disks" protobuf:"bytes,1,rep,name=disks"`
}

// +gogo:genproto=true
type DomainSnapshotMemory struct {
	Snapshot string `xml:"snapshot,attr" json:"snapshot" protobuf:"bytes,1,opt,name=snapshot"`
	File     string `xml:"file,attr,omitempty" json:"file,omitempty" protobuf:"bytes,2,opt,name=file"`
}

// +gogo:genproto=true
type DomainSnapshotParent struct {
	Name string `xml:"name" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type DomainSnapshot struct {
	XMLName      xml.Name              `xml:"domainsnapshot" json:"-"`
	Name         string                `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Description  string                `xml:"description,omitempty" json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	State        string                `xml:"state,omitempty" json:"state,omitempty" protobuf:"bytes,3,opt,name=state"`
	CreationTime string                `xml:"creationTime,omitempty" json:"creationTime,omitempty" protobuf:"bytes,4,opt,name=creationTime"`
	Parent       *DomainSnapshotParent `xml:"parent" json:"parent,omitempty" protobuf:"bytes,5,opt,name=parent"`
	Memory       *DomainSnapshotMemory `xml:"memory" json:"memory,omitempty" protobuf:"bytes,6,opt,name=memory"`
	Disks        *DomainSnapshotDisks  `xml:"disks" json:"disks,omitempty" protobuf:"bytes,7,opt,name=disks"`
	Domain       *Domain               `xml:"domain" json:"domain,omitempty" protobuf:"bytes,8,opt,name=domain"`
	Active       *int32                `xml:"active" json:"active,omitempty" protobuf:"varint,9,opt,name=active"`
}

type domainSnapshotDisk DomainSnapshotDisk

func (a *DomainSnapshotDisk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
		}
	}
	disk := domainSnapshotDisk(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainSnapshotDisk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	}
	disk := domainSnapshotDisk(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainSnapshotDisk(disk)
	return nil
}

func (s *DomainSnapshot) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *DomainSnapshot) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
