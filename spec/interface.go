package spec

import "encoding/xml"

type Interface struct {
	XMLName  xml.Name            `xml:"interface" json:"-"`
	Name     string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	Start    *InterfaceStart     `xml:"start" json:"start,omitempty"`
	MTU      *InterfaceMTU       `xml:"mtu" json:"mtu,omitempty"`
	Protocol []InterfaceProtocol `xml:"protocol" json:"protocol"`
	Link     *InterfaceLink      `xml:"link" json:"link,omitempty"`
	MAC      *InterfaceMAC       `xml:"mac" json:"mac,omitempty"`
	Bond     *InterfaceBond      `xml:"bond" json:"bond,omitempty"`
	Bridge   *InterfaceBridge    `xml:"bridge" json:"bridge,omitempty"`
	VLAN     *InterfaceVLAN      `xml:"vlan" json:"vlan,omitempty"`
}

type InterfaceStart struct {
	Mode string `xml:"mode,attr" json:"mode"`
}

type InterfaceMTU struct {
	Size uint `xml:"size,attr" json:"size"`
}

type InterfaceProtocol struct {
	Family   string             `xml:"family,attr,omitempty" json:"family,omitempty"`
	AutoConf *InterfaceAutoConf `xml:"autoconf" json:"autoConf,omitempty"`
	DHCP     *InterfaceDHCP     `xml:"dhcp" json:"dhcp,omitempty"`
	IPs      []InterfaceIP      `xml:"ip" json:"ip"`
	Route    []InterfaceRoute   `xml:"route" json:"route"`
}

type InterfaceAutoConf struct {
}

type InterfaceDHCP struct {
	PeerDNS string `xml:"peerdns,attr,omitempty" json:"peerDNS,omitempty"`
}

type InterfaceIP struct {
	Address string `xml:"address,attr" json:"address"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
}

type InterfaceRoute struct {
	Gateway string `xml:"gateway,attr" json:"gateway"`
}

type InterfaceLink struct {
	Speed uint   `xml:"speed,attr,omitempty" json:"speed,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty"`
}

type InterfaceMAC struct {
	Address string `xml:"address,attr" json:"address"`
}

type InterfaceBond struct {
	Mode       string               `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	ARPMon     *InterfaceBondARPMon `xml:"arpmon" json:"arpmon,omitempty"`
	MIIMon     *InterfaceBondMIIMon `xml:"miimon" json:"miimon,omitempty"`
	Interfaces []Interface          `xml:"interface" json:"interfaces"`
}

type InterfaceBondARPMon struct {
	Interval uint   `xml:"interval,attr,omitempty" json:"interval,omitempty"`
	Target   string `xml:"target,attr,omitempty" json:"target,omitempty"`
	Validate string `xml:"validate,attr,omitempty" json:"validate,omitempty"`
}

type InterfaceBondMIIMon struct {
	Freq    uint   `xml:"freq,attr,omitempty" json:"freq,omitempty"`
	UpDelay uint   `xml:"updelay,attr,omitempty" json:"updelay,omitempty"`
	Carrier string `xml:"carrier,attr,omitempty" json:"carrier,omitempty"`
}

type InterfaceBridge struct {
	STP        string      `xml:"stp,attr,omitempty" json:"stp,omitempty"`
	Delay      *float64    `xml:"delay,attr" json:"delay,omitempty"`
	Interfaces []Interface `xml:"interface" json:"interfaces"`
}

type InterfaceVLAN struct {
	Tag       *uint      `xml:"tag,attr" json:"tag,omitempty"`
	Interface *Interface `xml:"interface" json:"interface,omitempty"`
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
