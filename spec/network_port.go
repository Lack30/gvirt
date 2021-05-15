package spec

import (
	"encoding/xml"
	"fmt"
)

type NetworkPort struct {
	XMLName     xml.Name                `xml:"networkport" json:"-"`
	UUID        string                  `xml:"uuid,omitempty" json:"uuid"`
	Owner       *NetworkPortOwner       `xml:"owner" json:"owner,omitempty"`
	MAC         *NetworkPortMAC         `xml:"mac" json:"mac,omitempty"`
	Group       string                  `xml:"group,omitempty" json:"group,omitempty"`
	Bandwidth   *NetworkBandwidth       `xml:"bandwidth" json:"bandwidth,omitempty"`
	VLAN        *NetworkPortVLAN        `xml:"vlan" json:"vlan,omitempty"`
	PortOptions *NetworkPortPortOptions `xml:"port" json:"portOptions,omitempty"`
	VirtualPort *NetworkVirtualPort     `xml:"virtualport" json:"virtualPort,omitempty"`
	RXFilters   *NetworkPortRXFilters   `xml:"rxfilters" json:"rxFilters,omitempty"`
	Plug        *NetworkPortPlug        `xml:"plug" json:"plug,omitempty"`
}

type NetworkPortPortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty"`
}

type NetworkPortVLAN struct {
	Trunk string               `xml:"trunk,attr,omitempty" json:"trunk,omitempty"`
	Tags  []NetworkPortVLANTag `xml:"tag" json:"tags"`
}

type NetworkPortVLANTag struct {
	ID         uint   `xml:"id,attr" json:"id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty"`
}

type NetworkPortOwner struct {
	UUID string `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

type NetworkPortMAC struct {
	Address string `xml:"address,attr" json:"address"`
}

type NetworkPortRXFilters struct {
	TrustGuest string `xml:"trustGuest,attr" json:"trustGuest"`
}

type NetworkPortPlug struct {
	Bridge     *NetworkPortPlugBridge     `xml:"-" json:"bridge,omitempty"`
	Network    *NetworkPortPlugNetwork    `xml:"-" json:"network,omitempty"`
	Direct     *NetworkPortPlugDirect     `xml:"-" json:"direct,omitempty"`
	HostDevPCI *NetworkPortPlugHostDevPCI `xml:"-" json:"hostDevPCI,omitempty"`
}

type NetworkPortPlugBridge struct {
	Bridge          string `xml:"bridge,attr" json:"bridge"`
	MacTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty"`
}

type NetworkPortPlugNetwork struct {
	Bridge          string `xml:"bridge,attr" json:"bridge"`
	MacTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty"`
}

type NetworkPortPlugDirect struct {
	Dev  string `xml:"dev,attr" json:"dev"`
	Mode string `xml:"mode,attr" json:"mode"`
}

type NetworkPortPlugHostDevPCI struct {
	Managed string                            `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	Driver  *NetworkPortPlugHostDevPCIDriver  `xml:"driver" json:"driver,omitempty"`
	Address *NetworkPortPlugHostDevPCIAddress `xml:"address" json:"address,omitempty"`
}

type NetworkPortPlugHostDevPCIDriver struct {
	Name string `xml:"name,attr" json:"name"`
}

type NetworkPortPlugHostDevPCIAddress struct {
	Domain   *uint `xml:"domain,attr" json:"domain"`
	Bus      *uint `xml:"bus,attr" json:"bus"`
	Slot     *uint `xml:"slot,attr" json:"slot"`
	Function *uint `xml:"function,attr" json:"function"`
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

func (s *NetworkPort) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkPort) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
