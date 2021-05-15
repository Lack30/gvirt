package spec

import "encoding/xml"

type NetworkBridge struct {
	Name            string `xml:"name,attr,omitempty" json:"name,omitempty"`
	STP             string `xml:"stp,attr,omitempty" json:"stp,omitempty"`
	Delay           string `xml:"delay,attr,omitempty" json:"delay,omitempty"`
	MACTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty"`
	Zone            string `xml:"zone,attr,omitempty" json:"zone,omitempty"`
}

type NetworkVirtualPort struct {
	Params *NetworkVirtualPortParams `xml:"parameters" json:"parameters,omitempty"`
}

type NetworkVirtualPortParams struct {
	Any          *NetworkVirtualPortParamsAny          `xml:"-" json:"any,omitempty"`
	VEPA8021QBG  *NetworkVirtualPortParamsVEPA8021QBG  `xml:"-" json:"vepa8021QBG,omitempty"`
	VNTag8011QBH *NetworkVirtualPortParamsVNTag8021QBH `xml:"-" json:"vntag8011QBH,omitempty"`
	OpenVSwitch  *NetworkVirtualPortParamsOpenVSwitch  `xml:"-" json:"openvswitch,omitempty"`
	MidoNet      *NetworkVirtualPortParamsMidoNet      `xml:"-" json:"midoNet,omitempty"`
}

type NetworkVirtualPortParamsAny struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerId,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeId,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeIdVersion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
}

type NetworkVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerId,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeId,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeIdVersion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty"`
}

type NetworkVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
}

type NetworkVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileId,omitempty"`
}

type NetworkVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty"`
}

type NetworkDomain struct {
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty"`
	LocalOnly string `xml:"localOnly,attr,omitempty" json:"localOnly,omitempty"`
}

type NetworkForwardNATAddress struct {
	Start string `xml:"start,attr" json:"start"`
	End   string `xml:"end,attr" json:"end"`
}

type NetworkForwardNATPort struct {
	Start uint `xml:"start,attr" json:"start"`
	End   uint `xml:"end,attr" json:"end"`
}

type NetworkForwardNAT struct {
	IPv6      string                     `xml:"ipv6,attr,omitempty" json:"ipv6,omitempty"`
	Addresses []NetworkForwardNATAddress `xml:"address" json:"address"`
	Ports     []NetworkForwardNATPort    `xml:"port" json:"ports"`
}

type NetworkForward struct {
	Mode       string                    `xml:"mode,attr,omitempty" json:"mode,omitempty"`
	Dev        string                    `xml:"dev,attr,omitempty" json:"dev,omitempty"`
	Managed    string                    `xml:"managed,attr,omitempty" json:"managed,omitempty"`
	Driver     *NetworkForwardDriver     `xml:"driver" json:"driver,omitempty"`
	PFs        []NetworkForwardPF        `xml:"pf" json:"pf"`
	NAT        *NetworkForwardNAT        `xml:"nat" json:"nat,omitempty"`
	Interfaces []NetworkForwardInterface `xml:"interface" json:"interface"`
	Addresses  []NetworkForwardAddress   `xml:"address" json:"address"`
}

type NetworkForwardDriver struct {
	Name string `xml:"name,attr" json:"name"`
}

type NetworkForwardPF struct {
	Dev string `xml:"dev,attr" json:"dev"`
}

type NetworkForwardAddress struct {
	PCI *NetworkForwardAddressPCI `xml:"-" json:"pci,omitempty"`
}

type NetworkForwardAddressPCI struct {
	Domain   *uint `xml:"domain,attr" json:"domain,omitempty"`
	Bus      *uint `xml:"bus,attr" json:"bus,omitempty"`
	Slot     *uint `xml:"slot,attr" json:"slot,omitempty"`
	Function *uint `xml:"function,attr" json:"function,omitempty"`
}

type NetworkForwardInterface struct {
	XMLName xml.Name `xml:"interface" json:"-"`
	Dev     string   `xml:"dev,attr,omitempty" json:"dev,omitempty"`
}

type NetworkMAC struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
}

type NetworkDHCPRange struct {
	XMLName xml.Name          `xml:"range" json:"-"`
	Start   string            `xml:"start,attr,omitempty" json:"start,omitempty"`
	End     string            `xml:"end,attr,omitempty" json:"end,omitempty"`
	Lease   *NetworkDHCPLease `xml:"lease" json:"lease,omitempty"`
}

type NetworkDHCPLease struct {
	Expiry uint   `xml:"expiry,attr" json:"expiry"`
	Unit   string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
}

type NetworkDHCPHost struct {
	XMLName xml.Name          `xml:"host" json:"-"`
	ID      string            `xml:"id,attr,omitempty" json:"id,omitempty"`
	MAC     string            `xml:"mac,attr,omitempty" json:"mac,omitempty"`
	Name    string            `xml:"name,attr,omitempty" json:"name,omitempty"`
	IP      string            `xml:"ip,attr,omitempty" json:"ip,omitempty"`
	Lease   *NetworkDHCPLease `xml:"lease" json:"lease,omitempty"`
}

type NetworkBootp struct {
	File   string `xml:"file,attr,omitempty" json:"file,omitempty"`
	Server string `xml:"server,attr,omitempty" json:"server,omitempty"`
}

type NetworkDHCP struct {
	Ranges []NetworkDHCPRange `xml:"range" json:"range"`
	Hosts  []NetworkDHCPHost  `xml:"host" json:"host"`
	Bootp  []NetworkBootp     `xml:"bootp" json:"bootp"`
}

type NetworkIP struct {
	Address  string       `xml:"address,attr,omitempty" json:"address,omitempty"`
	Family   string       `xml:"family,attr,omitempty" json:"family,omitempty"`
	Netmask  string       `xml:"netmask,attr,omitempty" json:"netmask,omitempty"`
	Prefix   uint         `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	LocalPtr string       `xml:"localPtr,attr,omitempty" json:"localPtr,omitempty"`
	DHCP     *NetworkDHCP `xml:"dhcp" json:"dhcp,omitempty"`
	TFTP     *NetworkTFTP `xml:"tftp" json:"tftp,omitempty"`
}

type NetworkTFTP struct {
	Root string `xml:"root,attr,omitempty" json:"root,omitempty"`
}

type NetworkRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty"`
	Metric  string `xml:"metric,attr,omitempty" json:"metric,omitempty"`
}

type NetworkDNSForwarder struct {
	Domain string `xml:"domain,attr,omitempty" json:"domain,omitempty"`
	Addr   string `xml:"addr,attr,omitempty" json:"addr,omitempty"`
}

type NetworkDNSTXT struct {
	XMLName xml.Name `xml:"txt" json:"-"`
	Name    string   `xml:"name,attr" json:"name"`
	Value   string   `xml:"value,attr" json:"value"`
}

type NetworkDNSHostHostname struct {
	Hostname string `xml:",chardata" json:"hostname"`
}

type NetworkDNSHost struct {
	XMLName   xml.Name                 `xml:"host" json:"-"`
	IP        string                   `xml:"ip,attr" json:"ip"`
	Hostnames []NetworkDNSHostHostname `xml:"hostname" json:"hostnames"`
}

type NetworkDNSSRV struct {
	XMLName  xml.Name `xml:"srv" json:"-"`
	Service  string   `xml:"service,attr,omitempty" json:"service,omitempty"`
	Protocol string   `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	Target   string   `xml:"target,attr,omitempty" json:"target,omitempty"`
	Port     uint     `xml:"port,attr,omitempty" json:"port,omitempty"`
	Priority uint     `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	Weight   uint     `xml:"weight,attr,omitempty" json:"weight,omitempty"`
	Domain   string   `xml:"domain,attr,omitempty" json:"domain,omitempty"`
}

type NetworkDNS struct {
	Enable            string                `xml:"enable,attr,omitempty" json:"enable,omitempty"`
	ForwardPlainNames string                `xml:"forwardPlainNames,attr,omitempty" json:"forwardPlainNames,omitempty"`
	Forwarders        []NetworkDNSForwarder `xml:"forwarder" json:"forwarders"`
	TXTs              []NetworkDNSTXT       `xml:"txt" json:"txt"`
	Host              []NetworkDNSHost      `xml:"host" json:"host"`
	SRVs              []NetworkDNSSRV       `xml:"srv" json:"srv"`
}

type NetworkMetadata struct {
	XML string `xml:",innerxml" json:",inline"`
}

type NetworkMTU struct {
	Size uint `xml:"size,attr" json:"size"`
}

type Network struct {
	XMLName             xml.Name            `xml:"network" json:"-"`
	IPv6                string              `xml:"ipv6,attr,omitempty" json:"ipv6,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty"`
	Name                string              `xml:"name,omitempty" json:"name,omitempty"`
	UUID                string              `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Metadata            *NetworkMetadata    `xml:"metadata" json:"metadata,omitempty"`
	Forward             *NetworkForward     `xml:"forward" json:"forward,omitempty"`
	Bridge              *NetworkBridge      `xml:"bridge" json:"bridge,omitempty"`
	MTU                 *NetworkMTU         `xml:"mtu" json:"mtu,omitempty"`
	MAC                 *NetworkMAC         `xml:"mac" json:"mac,omitempty"`
	Domain              *NetworkDomain      `xml:"domain" json:"domain,omitempty"`
	DNS                 *NetworkDNS         `xml:"dns" json:"dns,omitempty"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan,omitempty"`
	Bandwidth           *NetworkBandwidth   `xml:"bandwidth" json:"bandwidth,omitempty"`
	PortOptions         *NetworkPortOptions `xml:"port" json:"portOptions,omitempty"`
	IPs                 []NetworkIP         `xml:"ip" json:"ip"`
	Routes              []NetworkRoute      `xml:"route" json:"route"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualPort"`
	PortGroups          []NetworkPortGroup  `xml:"portgroup" json:"portGroup"`

	DnsmasqOptions *NetworkDnsmasqOptions `json:"dnsmasqOptions,omitempty"`
}

type NetworkPortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty"`
}

type NetworkPortGroup struct {
	XMLName             xml.Name            `xml:"portgroup" json:"-"`
	Name                string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	Default             string              `xml:"default,attr,omitempty" json:"default,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan,omitempty"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualPort,omitempty"`
}

type NetworkVLAN struct {
	Trunk string           `xml:"trunk,attr,omitempty" json:"trunk,omitempty"`
	Tags  []NetworkVLANTag `xml:"tag" json:"tag"`
}

type NetworkVLANTag struct {
	ID         uint   `xml:"id,attr" json:"id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty"`
}

type NetworkBandwidthParams struct {
	Average *uint `xml:"average,attr" json:"average"`
	Peak    *uint `xml:"peak,attr" json:"peak"`
	Burst   *uint `xml:"burst,attr" json:"burst"`
	Floor   *uint `xml:"floor,attr" json:"floor"`
}

type NetworkBandwidth struct {
	ClassID  uint                    `xml:"classID,attr,omitempty" json:"classId"`
	Inbound  *NetworkBandwidthParams `xml:"inbound" json:"inbound,omitempty"`
	Outbound *NetworkBandwidthParams `xml:"outbound" json:"outbound,omitempty"`
}

type NetworkDnsmasqOptions struct {
	XMLName xml.Name               `xml:"http://libvirt.org/schemas/network/dnsmasq/1.0 options" json:"-"`
	Option  []NetworkDnsmasqOption `xml:"option" json:"option"`
}

type NetworkDnsmasqOption struct {
	Value string `xml:"value,attr" json:"value"`
}

func (a *NetworkVirtualPortParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "parameters"
	if a.Any != nil {
		return e.EncodeElement(a.Any, start)
	} else if a.VEPA8021QBG != nil {
		return e.EncodeElement(a.VEPA8021QBG, start)
	} else if a.VNTag8011QBH != nil {
		return e.EncodeElement(a.VNTag8011QBH, start)
	} else if a.OpenVSwitch != nil {
		return e.EncodeElement(a.OpenVSwitch, start)
	} else if a.MidoNet != nil {
		return e.EncodeElement(a.MidoNet, start)
	}
	return nil
}

func (a *NetworkVirtualPortParams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Any != nil {
		return d.DecodeElement(a.Any, &start)
	} else if a.VEPA8021QBG != nil {
		return d.DecodeElement(a.VEPA8021QBG, &start)
	} else if a.VNTag8011QBH != nil {
		return d.DecodeElement(a.VNTag8011QBH, &start)
	} else if a.OpenVSwitch != nil {
		return d.DecodeElement(a.OpenVSwitch, &start)
	} else if a.MidoNet != nil {
		return d.DecodeElement(a.MidoNet, &start)
	}
	return nil
}

type networkVirtualPort NetworkVirtualPort

func (a *NetworkVirtualPort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "virtualport"
	if a.Params != nil {
		if a.Params.Any != nil {
			/* no type attr wanted */
		} else if a.Params.VEPA8021QBG != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbg",
			})
		} else if a.Params.VNTag8011QBH != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbh",
			})
		} else if a.Params.OpenVSwitch != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "openvswitch",
			})
		} else if a.Params.MidoNet != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "midonet",
			})
		}
	}
	vp := networkVirtualPort(*a)
	return e.EncodeElement(&vp, start)
}

func (a *NetworkVirtualPort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	a.Params = &NetworkVirtualPortParams{}
	if !ok {
		var any NetworkVirtualPortParamsAny
		a.Params.Any = &any
	} else if typ == "802.1Qbg" {
		var vepa NetworkVirtualPortParamsVEPA8021QBG
		a.Params.VEPA8021QBG = &vepa
	} else if typ == "802.1Qbh" {
		var vntag NetworkVirtualPortParamsVNTag8021QBH
		a.Params.VNTag8011QBH = &vntag
	} else if typ == "openvswitch" {
		var ovs NetworkVirtualPortParamsOpenVSwitch
		a.Params.OpenVSwitch = &ovs
	} else if typ == "midonet" {
		var mido NetworkVirtualPortParamsMidoNet
		a.Params.MidoNet = &mido
	}

	vp := networkVirtualPort(*a)
	err := d.DecodeElement(&vp, &start)
	if err != nil {
		return err
	}
	*a = NetworkVirtualPort(vp)
	return nil
}

func (a *NetworkForwardAddressPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *NetworkForwardAddressPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (a *NetworkForwardAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(a.PCI, start)
	} else {
		return nil
	}
}

func (a *NetworkForwardAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
			break
		}
	}
	if typ == "" {
		d.Skip()
		return nil
	}

	if typ == "pci" {
		a.PCI = &NetworkForwardAddressPCI{}
		return d.DecodeElement(a.PCI, &start)
	}

	return nil
}

func (s *NetworkDHCPHost) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPHost) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSHost) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSHost) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkPortGroup) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkPortGroup) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSTXT) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSTXT) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSSRV) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSSRV) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDHCPRange) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPRange) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkForwardInterface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkForwardInterface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *Network) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Network) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
