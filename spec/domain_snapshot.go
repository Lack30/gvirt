package spec

import "encoding/xml"

type DomainSnapshotDisk struct {
	Name     string            `xml:"name,attr" json:"name,omitempty"`
	Snapshot string            `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty"`
	Driver   *DomainDiskDriver `xml:"driver" json:"driver,omitempty"`
	Source   *DomainDiskSource `xml:"source" json:"source,omitempty"`
}

type DomainSnapshotDisks struct {
	Disks []DomainSnapshotDisk `xml:"disk" json:"disks"`
}

type DomainSnapshotMemory struct {
	Snapshot string `xml:"snapshot,attr" json:"snapshot"`
	File     string `xml:"file,attr,omitempty" json:"file,omitempty"`
}

type DomainSnapshotParent struct {
	Name string `xml:"name" json:"name"`
}

type DomainSnapshot struct {
	XMLName      xml.Name              `xml:"domainsnapshot" json:"-"`
	Name         string                `xml:"name,omitempty" json:"name,omitempty"`
	Description  string                `xml:"description,omitempty" json:"description,omitempty"`
	State        string                `xml:"state,omitempty" json:"state,omitempty"`
	CreationTime string                `xml:"creationTime,omitempty" json:"creationTime,omitempty"`
	Parent       *DomainSnapshotParent `xml:"parent" json:"parent,omitempty"`
	Memory       *DomainSnapshotMemory `xml:"memory" json:"memory,omitempty"`
	Disks        *DomainSnapshotDisks  `xml:"disks" json:"disks,omitempty"`
	Domain       *Domain               `xml:"domain" json:"domain,omitempty"`
	Active       *int32                `xml:"active" json:"active,omitempty"`
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
