package spec

import "encoding/xml"

// +gogo:genproto=true
type Empty struct {
}

// +gogo:genproto=true
type StorageVolumeSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,1,opt,name=unit"`
	Value int64  `xml:",chardata" json:"value" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type StorageVolumeTargetPermissions struct {
	Owner string `xml:"owner,omitempty" json:"owner,omitempty" protobuf:"bytes,1,opt,name=owner"`
	Group string `xml:"group,omitempty" json:"group,omitempty" protobuf:"bytes,2,opt,name=group"`
	Mode  string `xml:"mode,omitempty" json:"mode,omitempty" protobuf:"bytes,3,opt,name=mode"`
	Label string `xml:"label,omitempty" json:"label,omitempty" protobuf:"bytes,4,opt,name=label"`
}

// +gogo:genproto=true
type StorageVolumeTargetFeature struct {
	LazyRefcounts *Empty `xml:"lazy_refcounts" json:"lazyRefcounts,omitempty" protobuf:"bytes,1,opt,name=lazyRefcounts"`
}

// +gogo:genproto=true
type StorageVolumeTargetFormat struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type StorageVolumeTargetTimestamps struct {
	Atime string `xml:"atime" json:"atime" protobuf:"bytes,1,opt,name=atime"`
	Mtime string `xml:"mtime" json:"mtime" protobuf:"bytes,2,opt,name=mtime"`
	Ctime string `xml:"ctime" json:"ctime" protobuf:"bytes,3,opt,name=ctime"`
}

// +gogo:genproto=true
type StorageVolumeTarget struct {
	Path        string                          `xml:"path,omitempty" json:"path" protobuf:"bytes,1,opt,name=path"`
	Format      *StorageVolumeTargetFormat      `xml:"format" json:"format,omitempty" protobuf:"bytes,2,opt,name=format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions" json:"permissions,omitempty" protobuf:"bytes,3,opt,name=permissions"`
	Timestamps  *StorageVolumeTargetTimestamps  `xml:"timestamps" json:"timestamps,omitempty" protobuf:"bytes,4,opt,name=timestamps"`
	Compat      string                          `xml:"compat,omitempty" json:"compat,omitempty" protobuf:"bytes,5,opt,name=compat"`
	NoCOW       *Empty                          `xml:"nocow" json:"nocow,omitempty" protobuf:"bytes,6,opt,name=nocow"`
	Features    []StorageVolumeTargetFeature    `xml:"features" json:"features" protobuf:"bytes,7,rep,name=features"`
	Encryption  *StorageEncryption              `xml:"encryption" json:"encryption,omitempty" protobuf:"bytes,8,opt,name=encryption"`
}

// +gogo:genproto=true
type StorageVolumeBackingStore struct {
	Path        string                          `xml:"path" json:"path" protobuf:"bytes,1,opt,name=path"`
	Format      *StorageVolumeTargetFormat      `xml:"format" json:"format,omitempty" protobuf:"bytes,2,opt,name=format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions" json:"permissions,omitempty" protobuf:"bytes,3,opt,name=permissions"`
}

// +gogo:genproto=true
type StorageVolume struct {
	XMLName      xml.Name                   `xml:"volume" json:"-"`
	Type         string                     `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Name         string                     `xml:"name" json:"name" protobuf:"bytes,2,opt,name=name"`
	Key          string                     `xml:"key,omitempty" json:"key" protobuf:"bytes,3,opt,name=key"`
	Allocation   *StorageVolumeSize         `xml:"allocation" json:"allocation,omitempty" protobuf:"bytes,4,opt,name=allocation"`
	Capacity     *StorageVolumeSize         `xml:"capacity" json:"capacity,omitempty" protobuf:"bytes,5,opt,name=capacity"`
	Physical     *StorageVolumeSize         `xml:"physical" json:"physical,omitempty" protobuf:"bytes,6,opt,name=physical"`
	Target       *StorageVolumeTarget       `xml:"target" json:"target,omitempty" protobuf:"bytes,7,opt,name=target"`
	BackingStore *StorageVolumeBackingStore `xml:"backingStore" json:"backingStore,omitempty" protobuf:"bytes,8,opt,name=backingStore"`
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
