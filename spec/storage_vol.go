package spec

import "encoding/xml"

type StorageVolumeSize struct {
	Unit  string `xml:"unit,attr,omitempty"`
	Value uint64 `xml:",chardata"`
}

type StorageVolumeTargetPermissions struct {
	Owner string `xml:"owner,omitempty"`
	Group string `xml:"group,omitempty"`
	Mode  string `xml:"mode,omitempty"`
	Label string `xml:"label,omitempty"`
}

type StorageVolumeTargetFeature struct {
	LazyRefcounts *struct{} `xml:"lazy_refcounts"`
}

type StorageVolumeTargetFormat struct {
	Type string `xml:"type,attr"`
}

type StorageVolumeTargetTimestamps struct {
	Atime string `xml:"atime"`
	Mtime string `xml:"mtime"`
	Ctime string `xml:"ctime"`
}

type StorageVolumeTarget struct {
	Path        string                          `xml:"path,omitempty"`
	Format      *StorageVolumeTargetFormat      `xml:"format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions"`
	Timestamps  *StorageVolumeTargetTimestamps  `xml:"timestamps"`
	Compat      string                          `xml:"compat,omitempty"`
	NoCOW       *struct{}                       `xml:"nocow"`
	Features    []StorageVolumeTargetFeature    `xml:"features"`
	Encryption  *StorageEncryption              `xml:"encryption"`
}

type StorageVolumeBackingStore struct {
	Path        string                          `xml:"path"`
	Format      *StorageVolumeTargetFormat      `xml:"format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions"`
}

type StorageVolume struct {
	XMLName      xml.Name                   `xml:"volume"`
	Type         string                     `xml:"type,attr,omitempty"`
	Name         string                     `xml:"name"`
	Key          string                     `xml:"key,omitempty"`
	Allocation   *StorageVolumeSize         `xml:"allocation"`
	Capacity     *StorageVolumeSize         `xml:"capacity"`
	Physical     *StorageVolumeSize         `xml:"physical"`
	Target       *StorageVolumeTarget       `xml:"target"`
	BackingStore *StorageVolumeBackingStore `xml:"backingStore"`
}

func (s *StorageVolume) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StorageVolume) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
