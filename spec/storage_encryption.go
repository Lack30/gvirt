package spec

type StorageEncryptionSecret struct {
	Type string `xml:"type,attr"`
	UUID string `xml:"uuid,attr"`
}

type StorageEncryptionCipher struct {
	Name string `xml:"name,attr"`
	Size uint64 `xml:"size,attr"`
	Mode string `xml:"mode,attr"`
	Hash string `xml:"hash,attr"`
}

type StorageEncryptionIvgen struct {
	Name string `xml:"name,attr"`
	Hash string `xml:"hash,attr"`
}

type StorageEncryption struct {
	Format string                   `xml:"format,attr"`
	Secret *StorageEncryptionSecret `xml:"secret"`
	Cipher *StorageEncryptionCipher `xml:"cipher"`
	Ivgen  *StorageEncryptionIvgen  `xml:"ivgen"`
}
