package spec

import "encoding/xml"

type Empty struct {
}

type StorageVolumeSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value int64  `xml:",chardata" json:"value"`
}

type StorageVolumeTargetPermissions struct {
	Owner string `xml:"owner,omitempty" json:"owner,omitempty"`
	Group string `xml:"group,omitempty" json:"group,omitempty"`
	Mode  string `xml:"mode,omitempty" json:"mode,omitempty"`
	Label string `xml:"label,omitempty" json:"label,omitempty"`
}

type StorageVolumeTargetFeature struct {
	LazyRefcounts *Empty `xml:"lazy_refcounts" json:"lazyRefcounts,omitempty"`
}

type StorageVolumeTargetFormat struct {
	Type string `xml:"type,attr" json:"type"`
}

type StorageVolumeTargetTimestamps struct {
	Atime string `xml:"atime" json:"atime"`
	Mtime string `xml:"mtime" json:"mtime"`
	Ctime string `xml:"ctime" json:"ctime"`
}

type StorageVolumeTarget struct {
	Path        string                          `xml:"path,omitempty" json:"path"`
	Format      *StorageVolumeTargetFormat      `xml:"format" json:"format,omitempty"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions" json:"permissions,omitempty"`
	Timestamps  *StorageVolumeTargetTimestamps  `xml:"timestamps" json:"timestamps,omitempty"`
	Compat      string                          `xml:"compat,omitempty" json:"compat,omitempty"`
	NoCOW       *Empty                          `xml:"nocow" json:"nocow,omitempty"`
	Features    []StorageVolumeTargetFeature    `xml:"features" json:"features"`
	Encryption  *StorageEncryption              `xml:"encryption" json:"encryption,omitempty"`
}

type StorageVolumeBackingStore struct {
	Path        string                          `xml:"path" json:"path"`
	Format      *StorageVolumeTargetFormat      `xml:"format" json:"format,omitempty"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions" json:"permissions,omitempty"`
}

type StorageVolume struct {
	XMLName      xml.Name                   `xml:"volume" json:"-"`
	Type         string                     `xml:"type,attr,omitempty" json:"type,omitempty"`
	Name         string                     `xml:"name" json:"name"`
	Key          string                     `xml:"key,omitempty" json:"key"`
	Allocation   *StorageVolumeSize         `xml:"allocation" json:"allocation,omitempty"`
	Capacity     *StorageVolumeSize         `xml:"capacity" json:"capacity,omitempty"`
	Physical     *StorageVolumeSize         `xml:"physical" json:"physical,omitempty"`
	Target       *StorageVolumeTarget       `xml:"target" json:"target,omitempty"`
	BackingStore *StorageVolumeBackingStore `xml:"backingStore" json:"backingStore,omitempty"`
}

func (s *StorageVolume) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StorageVolume) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
