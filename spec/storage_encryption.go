package spec

type StorageEncryptionSecret struct {
	Type string `xml:"type,attr" json:"type"`
	UUID string `xml:"uuid,attr" json:"uuid"`
}

type StorageEncryptionCipher struct {
	Name string `xml:"name,attr" json:"name"`
	Size uint64 `xml:"size,attr" json:"size"`
	Mode string `xml:"mode,attr" json:"mode"`
	Hash string `xml:"hash,attr" json:"hash"`
}

type StorageEncryptionIvgen struct {
	Name string `xml:"name,attr" json:"name"`
	Hash string `xml:"hash,attr" json:"hash"`
}

type StorageEncryption struct {
	Format string                   `xml:"format,attr" json:"format"`
	Secret *StorageEncryptionSecret `xml:"secret" json:"secret,omitempty"`
	Cipher *StorageEncryptionCipher `xml:"cipher" json:"cipher,omitempty"`
	Ivgen  *StorageEncryptionIvgen  `xml:"ivgen" json:"ivgen,omitempty"`
}
