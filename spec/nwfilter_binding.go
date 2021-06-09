package spec

import "encoding/xml"


type NWFilterBinding struct {
	XMLName   xml.Name                  `xml:"filterbinding" json:"-"`
	Owner     *NWFilterBindingOwner     `xml:"owner" json:"owner,omitempty"`
	PortDev   *NWFilterBindingPortDev   `xml:"portdev" json:"portDev,omitempty"`
	MAC       *NWFilterBindingMAC       `xml:"mac" json:"mac,omitempty"`
	FilterRef *NWFilterBindingFilterRef `xml:"filterref" json:"filterref,omitempty"`
}


type NWFilterBindingOwner struct {
	UUID string `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}


type NWFilterBindingPortDev struct {
	Name string `xml:"name,attr" json:"name"`
}


type NWFilterBindingMAC struct {
	Address string `xml:"address,attr" json:"address"`
}


type NWFilterBindingFilterRef struct {
	Filter     string                       `xml:"filter,attr" json:"filter"`
	Parameters []NWFilterBindingFilterParam `xml:"parameter" json:"parameters"`
}


type NWFilterBindingFilterParam struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

func (s *NWFilterBinding) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NWFilterBinding) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
