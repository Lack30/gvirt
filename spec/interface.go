package spec

import "encoding/xml"

// +gogo:genproto=true
type Interface struct {
	XMLName  xml.Name            `xml:"interface" json:"-"`
	Name     string              `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Start    *InterfaceStart     `xml:"start" json:"start,omitempty" protobuf:"bytes,2,opt,name=start"`
	MTU      *InterfaceMTU       `xml:"mtu" json:"mtu,omitempty" protobuf:"bytes,3,opt,name=mtu"`
	Protocol []InterfaceProtocol `xml:"protocol" json:"protocol" protobuf:"bytes,4,rep,name=protocol"`
	Link     *InterfaceLink      `xml:"link" json:"link,omitempty" protobuf:"bytes,5,opt,name=link"`
	MAC      *InterfaceMAC       `xml:"mac" json:"mac,omitempty" protobuf:"bytes,6,opt,name=mac"`
	Bond     *InterfaceBond      `xml:"bond" json:"bond,omitempty" protobuf:"bytes,7,opt,name=bond"`
	Bridge   *InterfaceBridge    `xml:"bridge" json:"bridge,omitempty" protobuf:"bytes,8,opt,name=bridge"`
	VLAN     *InterfaceVLAN      `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,9,opt,name=vlan"`
}

// +gogo:genproto=true
type InterfaceStart struct {
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,1,opt,name=mode"`
}

// +gogo:genproto=true
type InterfaceMTU struct {
	Size_ int32 `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
}

// +gogo:genproto=true
type InterfaceProtocol struct {
	Family   string             `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,1,opt,name=family"`
	AutoConf *InterfaceAutoConf `xml:"autoconf" json:"autoConf,omitempty" protobuf:"bytes,2,opt,name=autoConf"`
	DHCP     *InterfaceDHCP     `xml:"dhcp" json:"dhcp,omitempty" protobuf:"bytes,3,opt,name=dhcp"`
	IPs      []InterfaceIP      `xml:"ip" json:"ip" protobuf:"bytes,4,rep,name=ip"`
	Route    []InterfaceRoute   `xml:"route" json:"route" protobuf:"bytes,5,rep,name=route"`
}

// +gogo:genproto=true
type InterfaceAutoConf struct {
}

// +gogo:genproto=true
type InterfaceDHCP struct {
	PeerDNS string `xml:"peerdns,attr,omitempty" json:"peerDNS,omitempty" protobuf:"bytes,1,opt,name=peerDNS"`
}

// +gogo:genproto=true
type InterfaceIP struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty" protobuf:"varint,2,opt,name=prefix"`
}

// +gogo:genproto=true
type InterfaceRoute struct {
	Gateway string `xml:"gateway,attr" json:"gateway" protobuf:"bytes,1,opt,name=gateway"`
}

// +gogo:genproto=true
type InterfaceLink struct {
	Speed int32  `xml:"speed,attr,omitempty" json:"speed,omitempty" protobuf:"varint,1,opt,name=speed"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,2,opt,name=state"`
}

// +gogo:genproto=true
type InterfaceMAC struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type InterfaceBond struct {
	Mode       string               `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
	ARPMon     *InterfaceBondARPMon `xml:"arpmon" json:"arpmon,omitempty" protobuf:"bytes,2,opt,name=arpmon"`
	MIIMon     *InterfaceBondMIIMon `xml:"miimon" json:"miimon,omitempty" protobuf:"bytes,3,opt,name=miimon"`
	Interfaces []Interface          `xml:"interface" json:"interfaces" protobuf:"bytes,4,rep,name=interfaces"`
}

// +gogo:genproto=true
type InterfaceBondARPMon struct {
	Interval int32  `xml:"interval,attr,omitempty" json:"interval,omitempty" protobuf:"varint,1,opt,name=interval"`
	Target   string `xml:"target,attr,omitempty" json:"target,omitempty" protobuf:"bytes,2,opt,name=target"`
	Validate string `xml:"validate,attr,omitempty" json:"validate,omitempty" protobuf:"bytes,3,opt,name=validate"`
}

// +gogo:genproto=true
type InterfaceBondMIIMon struct {
	Freq    int32  `xml:"freq,attr,omitempty" json:"freq,omitempty" protobuf:"varint,1,opt,name=freq"`
	UpDelay int32  `xml:"updelay,attr,omitempty" json:"updelay,omitempty" protobuf:"varint,2,opt,name=updelay"`
	Carrier string `xml:"carrier,attr,omitempty" json:"carrier,omitempty" protobuf:"bytes,3,opt,name=carrier"`
}

// +gogo:genproto=true
type InterfaceBridge struct {
	STP        string      `xml:"stp,attr,omitempty" json:"stp,omitempty" protobuf:"bytes,1,opt,name=stp"`
	Delay      *float64    `xml:"delay,attr" json:"delay,omitempty" protobuf:"fixed64,2,opt,name=delay"`
	Interfaces []Interface `xml:"interface" json:"interfaces" protobuf:"bytes,3,rep,name=interfaces"`
}

// +gogo:genproto=true
type InterfaceVLAN struct {
	Tag       *int32     `xml:"tag,attr" json:"tag,omitempty" protobuf:"varint,1,opt,name=tag"`
	Interface *Interface `xml:"interface" json:"interface,omitempty" protobuf:"bytes,2,opt,name=interface"`
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

func (s *Interface) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Interface) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
