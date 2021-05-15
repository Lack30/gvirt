package spec

import "encoding/xml"

type NWFilterBinding struct {
	XMLName   xml.Name                  `xml:"filterbinding"`
	Owner     *NWFilterBindingOwner     `xml:"owner"`
	PortDev   *NWFilterBindingPortDev   `xml:"portdev"`
	MAC       *NWFilterBindingMAC       `xml:"mac"`
	FilterRef *NWFilterBindingFilterRef `xml:"filterref"`
}

type NWFilterBindingOwner struct {
	UUID string `xml:"uuid,omitempty"`
	Name string `xml:"name,omitempty"`
}

type NWFilterBindingPortDev struct {
	Name string `xml:"name,attr"`
}

type NWFilterBindingMAC struct {
	Address string `xml:"address,attr"`
}

type NWFilterBindingFilterRef struct {
	Filter     string                       `xml:"filter,attr"`
	Parameters []NWFilterBindingFilterParam `xml:"parameter"`
}

type NWFilterBindingFilterParam struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (s *NWFilterBinding) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NWFilterBinding) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
