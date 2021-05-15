package spec

import "encoding/xml"

type Interface struct {
	XMLName  xml.Name            `xml:"interface"`
	Name     string              `xml:"name,attr,omitempty"`
	Start    *InterfaceStart     `xml:"start"`
	MTU      *InterfaceMTU       `xml:"mtu"`
	Protocol []InterfaceProtocol `xml:"protocol"`
	Link     *InterfaceLink      `xml:"link"`
	MAC      *InterfaceMAC       `xml:"mac"`
	Bond     *InterfaceBond      `xml:"bond"`
	Bridge   *InterfaceBridge    `xml:"bridge"`
	VLAN     *InterfaceVLAN      `xml:"vlan"`
}

type InterfaceStart struct {
	Mode string `xml:"mode,attr"`
}

type InterfaceMTU struct {
	Size uint `xml:"size,attr"`
}

type InterfaceProtocol struct {
	Family   string             `xml:"family,attr,omitempty"`
	AutoConf *InterfaceAutoConf `xml:"autoconf"`
	DHCP     *InterfaceDHCP     `xml:"dhcp"`
	IPs      []InterfaceIP      `xml:"ip"`
	Route    []InterfaceRoute   `xml:"route"`
}

type InterfaceAutoConf struct {
}

type InterfaceDHCP struct {
	PeerDNS string `xml:"peerdns,attr,omitempty"`
}

type InterfaceIP struct {
	Address string `xml:"address,attr"`
	Prefix  uint   `xml:"prefix,attr,omitempty"`
}

type InterfaceRoute struct {
	Gateway string `xml:"gateway,attr"`
}

type InterfaceLink struct {
	Speed uint   `xml:"speed,attr,omitempty"`
	State string `xml:"state,attr,omitempty"`
}

type InterfaceMAC struct {
	Address string `xml:"address,attr"`
}

type InterfaceBond struct {
	Mode       string               `xml:"mode,attr,omitempty"`
	ARPMon     *InterfaceBondARPMon `xml:"arpmon"`
	MIIMon     *InterfaceBondMIIMon `xml:"miimon"`
	Interfaces []Interface          `xml:"interface"`
}

type InterfaceBondARPMon struct {
	Interval uint   `xml:"interval,attr,omitempty"`
	Target   string `xml:"target,attr,omitempty"`
	Validate string `xml:"validate,attr,omitempty"`
}

type InterfaceBondMIIMon struct {
	Freq    uint   `xml:"freq,attr,omitempty"`
	UpDelay uint   `xml:"updelay,attr,omitempty"`
	Carrier string `xml:"carrier,attr,omitempty"`
}

type InterfaceBridge struct {
	STP        string      `xml:"stp,attr,omitempty"`
	Delay      *float64    `xml:"delay,attr"`
	Interfaces []Interface `xml:"interface"`
}

type InterfaceVLAN struct {
	Tag       *uint      `xml:"tag,attr"`
	Interface *Interface `xml:"interface"`
}

type interfaceDup Interface

func (s *Interface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "interface"

	typ := "ethernet"
	if s.Bond != nil {
		typ = "bond"
	} else if s.Bridge != nil {
		typ = "bridge"
	} else if s.VLAN != nil {
		typ = "vlan"
	}

	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "type"},
		Value: typ,
	})

	i := interfaceDup(*s)

	return e.EncodeElement(i, start)
}

func (s *Interface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Interface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
