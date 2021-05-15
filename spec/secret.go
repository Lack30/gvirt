package spec

import "encoding/xml"

type SecretUsage struct {
	Type   string `xml:"type,attr"`
	Volume string `xml:"volume,omitempty"`
	Name   string `xml:"name,omitempty"`
	Target string `xml:"target,omitempty"`
}

type Secret struct {
	XMLName     xml.Name     `xml:"secret"`
	Ephemeral   string       `xml:"ephemeral,attr,omitempty"`
	Private     string       `xml:"private,attr,omitempty"`
	Description string       `xml:"description,omitempty"`
	UUID        string       `xml:"uuid,omitempty"`
	Usage       *SecretUsage `xml:"usage"`
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
