package spec

import "encoding/xml"

// +gogo:genproto=true
type SecretUsage struct {
	Type   string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Volume string `xml:"volume,omitempty" json:"volume,omitempty" protobuf:"bytes,2,opt,name=volume"`
	Name   string `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	Target string `xml:"target,omitempty" json:"target,omitempty" protobuf:"bytes,4,opt,name=target"`
}

// +gogo:genproto=true
type Secret struct {
	XMLName     xml.Name     `xml:"secret" json:"-"`
	Ephemeral   string       `xml:"ephemeral,attr,omitempty" json:"ephemeral,omitempty" protobuf:"bytes,1,opt,name=ephemeral"`
	Private     string       `xml:"private,attr,omitempty" json:"private,omitempty" protobuf:"bytes,2,opt,name=private"`
	Description string       `xml:"description,omitempty" json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	UUID        string       `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,4,opt,name=uuid"`
	Usage       *SecretUsage `xml:"usage" json:"usage" protobuf:"bytes,5,opt,name=usage"`
}

func (s *Secret) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Secret) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
