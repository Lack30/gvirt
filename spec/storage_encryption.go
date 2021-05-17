package spec

// +gogo:genproto=true
type StorageEncryptionSecret struct {
	Type string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	UUID string `xml:"uuid,attr" json:"uuid" protobuf:"bytes,2,opt,name=uuid"`
}

// +gogo:genproto=true
type StorageEncryptionCipher struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Size int64  `xml:"size,attr" json:"size" protobuf:"varint,5,opt,name=size"`
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,3,opt,name=mode"`
	Hash string `xml:"hash,attr" json:"hash" protobuf:"bytes,4,opt,name=hash"`
}

// +gogo:genproto=true
type StorageEncryptionIvgen struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Hash string `xml:"hash,attr" json:"hash" protobuf:"bytes,2,opt,name=hash"`
}

// +gogo:genproto=true
type StorageEncryption struct {
	Format string                   `xml:"format,attr" json:"format" protobuf:"bytes,1,opt,name=format"`
	Secret *StorageEncryptionSecret `xml:"secret" json:"secret,omitempty" protobuf:"bytes,2,opt,name=secret"`
	Cipher *StorageEncryptionCipher `xml:"cipher" json:"cipher,omitempty" protobuf:"bytes,3,opt,name=cipher"`
	Ivgen  *StorageEncryptionIvgen  `xml:"ivgen" json:"ivgen,omitempty" protobuf:"bytes,4,opt,name=ivgen"`
}
