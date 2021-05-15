package spec

import (
	"encoding/xml"
	"fmt"
)

// +gogo:genproto=true
type NetworkPort struct {
	XMLName     xml.Name                `xml:"networkport" json:"-"`
	UUID        string                  `xml:"uuid,omitempty" json:"uuid" protobuf:"bytes,1,opt,name=uuid"`
	Owner       *NetworkPortOwner       `xml:"owner" json:"owner,omitempty" protobuf:"bytes,2,opt,name=owner"`
	MAC         *NetworkPortMAC         `xml:"mac" json:"mac,omitempty" protobuf:"bytes,3,opt,name=mac"`
	Group       string                  `xml:"group,omitempty" json:"group,omitempty" protobuf:"bytes,4,opt,name=group"`
	Bandwidth   *NetworkBandwidth       `xml:"bandwidth" json:"bandwidth,omitempty" protobuf:"bytes,5,opt,name=bandwidth"`
	VLAN        *NetworkPortVLAN        `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,6,opt,name=vlan"`
	PortOptions *NetworkPortPortOptions `xml:"port" json:"portOptions,omitempty" protobuf:"bytes,7,opt,name=portOptions"`
	VirtualPort *NetworkVirtualPort     `xml:"virtualport" json:"virtualPort,omitempty" protobuf:"bytes,8,opt,name=virtualPort"`
	RXFilters   *NetworkPortRXFilters   `xml:"rxfilters" json:"rxFilters,omitempty" protobuf:"bytes,9,opt,name=rxFilters"`
	Plug        *NetworkPortPlug        `xml:"plug" json:"plug,omitempty" protobuf:"bytes,10,opt,name=plug"`
}

// +gogo:genproto=true
type NetworkPortPortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty" protobuf:"bytes,1,opt,name=isolated"`
}

// +gogo:genproto=true
type NetworkPortVLAN struct {
	Trunk string               `xml:"trunk,attr,omitempty" json:"trunk,omitempty" protobuf:"bytes,1,opt,name=trunk"`
	Tags  []NetworkPortVLANTag `xml:"tag" json:"tags" protobuf:"bytes,2,rep,name=tags"`
}

// +gogo:genproto=true
type NetworkPortVLANTag struct {
	ID         uint   `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty" protobuf:"bytes,2,opt,name=nativeMode"`
}

// +gogo:genproto=true
type NetworkPortOwner struct {
	UUID string `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,1,opt,name=uuid"`
	Name string `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type NetworkPortMAC struct {
	Address string `xml:"address,attr" json:"address" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type NetworkPortRXFilters struct {
	TrustGuest string `xml:"trustGuest,attr" json:"trustGuest" protobuf:"bytes,1,opt,name=trustGuest"`
}

// +gogo:genproto=true
type NetworkPortPlug struct {
	Bridge     *NetworkPortPlugBridge     `xml:"-" json:"bridge,omitempty" protobuf:"bytes,1,opt,name=bridge"`
	Network    *NetworkPortPlugNetwork    `xml:"-" json:"network,omitempty" protobuf:"bytes,2,opt,name=network"`
	Direct     *NetworkPortPlugDirect     `xml:"-" json:"direct,omitempty" protobuf:"bytes,3,opt,name=direct"`
	HostDevPCI *NetworkPortPlugHostDevPCI `xml:"-" json:"hostDevPCI,omitempty" protobuf:"bytes,4,opt,name=hostDevPCI"`
}

// +gogo:genproto=true
type NetworkPortPlugBridge struct {
	Bridge          string `xml:"bridge,attr" json:"bridge" protobuf:"bytes,1,opt,name=bridge"`
	MacTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty" protobuf:"bytes,2,opt,name=macTableManager"`
}

// +gogo:genproto=true
type NetworkPortPlugNetwork struct {
	Bridge          string `xml:"bridge,attr" json:"bridge" protobuf:"bytes,1,opt,name=bridge"`
	MacTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty" protobuf:"bytes,2,opt,name=macTableManager"`
}

// +gogo:genproto=true
type NetworkPortPlugDirect struct {
	Dev  string `xml:"dev,attr" json:"dev" protobuf:"bytes,1,opt,name=dev"`
	Mode string `xml:"mode,attr" json:"mode" protobuf:"bytes,2,opt,name=mode"`
}

// +gogo:genproto=true
type NetworkPortPlugHostDevPCI struct {
	Managed string                            `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,1,opt,name=managed"`
	Driver  *NetworkPortPlugHostDevPCIDriver  `xml:"driver" json:"driver,omitempty" protobuf:"bytes,2,opt,name=driver"`
	Address *NetworkPortPlugHostDevPCIAddress `xml:"address" json:"address,omitempty" protobuf:"bytes,3,opt,name=address"`
}

// +gogo:genproto=true
type NetworkPortPlugHostDevPCIDriver struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type NetworkPortPlugHostDevPCIAddress struct {
	Domain   *uint `xml:"domain,attr" json:"domain" protobuf:"varint,1,opt,name=domain"`
	Bus      *uint `xml:"bus,attr" json:"bus" protobuf:"varint,2,opt,name=bus"`
	Slot     *uint `xml:"slot,attr" json:"slot" protobuf:"varint,3,opt,name=slot"`
	Function *uint `xml:"function,attr" json:"function" protobuf:"varint,4,opt,name=function"`
}

func (a *NetworkPortPlugHostDevPCIAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *NetworkPortPlugHostDevPCIAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "domain" {
			if err := unmarshalUintAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalUintAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshalUintAttr(attr.Value, &a.Function, 0); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (p *NetworkPortPlug) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "plug"
	if p.Bridge != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "bridge",
		})
		return e.EncodeElement(p.Bridge, start)
	} else if p.Network != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "network",
		})
		return e.EncodeElement(p.Network, start)
	} else if p.Direct != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "direct",
		})
		return e.EncodeElement(p.Direct, start)
	} else if p.HostDevPCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "hostdev-pci",
		})
		return e.EncodeElement(p.HostDevPCI, start)
	}
	return nil
}

func (p *NetworkPortPlug) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing type attribute on plug")
	} else if typ == "bridge" {
		var pb NetworkPortPlugBridge
		if err := d.DecodeElement(&pb, &start); err != nil {
			return err
		}
		p.Bridge = &pb
	} else if typ == "network" {
		var pn NetworkPortPlugNetwork
		if err := d.DecodeElement(&pn, &start); err != nil {
			return err
		}
		p.Network = &pn
	} else if typ == "direct" {
		var pd NetworkPortPlugDirect
		if err := d.DecodeElement(&pd, &start); err != nil {
			return err
		}
		p.Direct = &pd
	} else if typ == "hostdev-pci" {
		var ph NetworkPortPlugHostDevPCI
		if err := d.DecodeElement(&ph, &start); err != nil {
			return err
		}
		p.HostDevPCI = &ph
	}
	d.Skip()
	return nil
}

func (s *NetworkPort) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkPort) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
