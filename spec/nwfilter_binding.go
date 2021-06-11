package spec

import "encoding/xml"

// +gogo:genproto=true
type NWFilterBinding struct {
	XMLName   xml.Name                  `xml:"filterbinding" json:"-"`
	Owner     *NWFilterBindingOwner     `xml:"owner" json:"owner,omitempty" protobuf:"bytes,1,opt,name=owner"`
	PortDev   *NWFilterBindingPortDev   `xml:"portdev" json:"portDev,omitempty" protobuf:"bytes,2,opt,name=portDev"`
	MAC       *NWFilterBindingMAC       `xml:"mac" json:"mac,omitempty" protobuf:"bytes,3,opt,name=mac"`
	FilterRef *NWFilterBindingFilterRef `xml:"filterref" json:"filterref,omitempty" protobuf:"bytes,4,opt,name=filterref"`
}

// +gogo:genproto=true
type NWFilterBindingOwner struct {
	UUID string `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,1,opt,name=uuid"`
	Name string `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type NWFilterBindingPortDev struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type NWFilterBindingMAC struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type NWFilterBindingFilterRef struct {
	Filter     string                       `xml:"filter,attr" json:"filter" protobuf:"bytes,1,opt,name=filter"`
	Parameters []NWFilterBindingFilterParam `xml:"parameter" json:"parameters" protobuf:"bytes,2,rep,name=parameters"`
}

// +gogo:genproto=true
type NWFilterBindingFilterParam struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
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
