package spec

import "encoding/xml"

type SecretUsage struct {
	Type   string `xml:"type,attr" json:"type"`
	Volume string `xml:"volume,omitempty" json:"volume,omitempty"`
	Name   string `xml:"name,omitempty" json:"name,omitempty"`
	Target string `xml:"target,omitempty" json:"target,omitempty"`
}

type Secret struct {
	XMLName     xml.Name     `xml:"secret" json:"-"`
	Ephemeral   string       `xml:"ephemeral,attr,omitempty" json:"ephemeral,omitempty"`
	Private     string       `xml:"private,attr,omitempty" json:"private,omitempty"`
	Description string       `xml:"description,omitempty" json:"description,omitempty"`
	UUID        string       `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Usage       *SecretUsage `xml:"usage" json:"usage"`
}

func (s *Secret) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Secret) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
