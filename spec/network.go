package spec

import "encoding/xml"

// +gogo:genproto=true
type NetworkBridge struct {
	Name            string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	STP             string `xml:"stp,attr,omitempty" json:"stp,omitempty" protobuf:"bytes,2,opt,name=stp"`
	Delay           string `xml:"delay,attr,omitempty" json:"delay,omitempty" protobuf:"bytes,3,opt,name=delay"`
	MACTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty" protobuf:"bytes,4,opt,name=macTableManager"`
	Zone            string `xml:"zone,attr,omitempty" json:"zone,omitempty" protobuf:"bytes,5,opt,name=zone"`
}

// +gogo:genproto=true
type NetworkVirtualPort struct {
	Params *NetworkVirtualPortParams `xml:"parameters" json:"parameters,omitempty" protobuf:"bytes,1,opt,name=parameters"`
}

// +gogo:genproto=true
type NetworkVirtualPortParams struct {
	Any          *NetworkVirtualPortParamsAny          `xml:"-" json:"any,omitempty" protobuf:"bytes,1,opt,name=any"`
	VEPA8021QBG  *NetworkVirtualPortParamsVEPA8021QBG  `xml:"-" json:"vepa8021QBG,omitempty" protobuf:"bytes,2,opt,name=vepa8021QBG"`
	VNTag8011QBH *NetworkVirtualPortParamsVNTag8021QBH `xml:"-" json:"vntag8011QBH,omitempty" protobuf:"bytes,3,opt,name=vntag8011QBH"`
	OpenVSwitch  *NetworkVirtualPortParamsOpenVSwitch  `xml:"-" json:"openvswitch,omitempty" protobuf:"bytes,4,opt,name=openvswitch"`
	MidoNet      *NetworkVirtualPortParamsMidoNet      `xml:"-" json:"midoNet,omitempty" protobuf:"bytes,5,opt,name=midoNet"`
}

// +gogo:genproto=true
type NetworkVirtualPortParamsAny struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty" protobuf:"varint,1,opt,name=managerId"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty" protobuf:"varint,2,opt,name=typeId"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty" protobuf:"varint,3,opt,name=typeIdVersion"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty" protobuf:"bytes,4,opt,name=instanceId"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,5,opt,name=profileId"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,6,opt,name=interfaceId"`
}

// +gogo:genproto=true
type NetworkVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *int32 `xml:"managerid,attr" json:"managerId,omitempty" protobuf:"varint,1,opt,name=managerId"`
	TypeID        *int32 `xml:"typeid,attr" json:"typeId,omitempty" protobuf:"varint,2,opt,name=typeId"`
	TypeIDVersion *int32 `xml:"typeidversion,attr" json:"typeIdVersion,omitempty" protobuf:"varint,3,opt,name=typeIdVersion"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceId,omitempty" protobuf:"bytes,4,opt,name=instanceId"`
}

// +gogo:genproto=true
type NetworkVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,1,opt,name=profileId"`
}

// +gogo:genproto=true
type NetworkVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,1,opt,name=interfaceId"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileId,omitempty" protobuf:"bytes,2,opt,name=profileId"`
}

// +gogo:genproto=true
type NetworkVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceId,omitempty" protobuf:"bytes,1,opt,name=interfaceId"`
}

// +gogo:genproto=true
type NetworkDomain struct {
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	LocalOnly string `xml:"localOnly,attr,omitempty" json:"localOnly,omitempty" protobuf:"bytes,2,opt,name=localOnly"`
}

// +gogo:genproto=true
type NetworkForwardNATAddress struct {
	Start string `xml:"start,attr" json:"start" protobuf:"bytes,1,opt,name=start"`
	End   string `xml:"end,attr" json:"end" protobuf:"bytes,2,opt,name=end"`
}

// +gogo:genproto=true
type NetworkForwardNATPort struct {
	Start int32 `xml:"start,attr" json:"start" protobuf:"varint,1,opt,name=start"`
	End   int32 `xml:"end,attr" json:"end" protobuf:"varint,2,opt,name=end"`
}

// +gogo:genproto=true
type NetworkForwardNAT struct {
	IPv6      string                     `xml:"ipv6,attr,omitempty" json:"ipv6,omitempty" protobuf:"bytes,1,opt,name=ipv6"`
	Addresses []NetworkForwardNATAddress `xml:"address" json:"address" protobuf:"bytes,2,rep,name=address"`
	Ports     []NetworkForwardNATPort    `xml:"port" json:"ports" protobuf:"bytes,3,rep,name=ports"`
}

// +gogo:genproto=true
type NetworkForward struct {
	Mode       string                    `xml:"mode,attr,omitempty" json:"mode,omitempty" protobuf:"bytes,1,opt,name=mode"`
	Dev        string                    `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,2,opt,name=dev"`
	Managed    string                    `xml:"managed,attr,omitempty" json:"managed,omitempty" protobuf:"bytes,3,opt,name=managed"`
	Driver     *NetworkForwardDriver     `xml:"driver" json:"driver,omitempty" protobuf:"bytes,4,opt,name=driver"`
	PFs        []NetworkForwardPF        `xml:"pf" json:"pf" protobuf:"bytes,5,rep,name=pf"`
	NAT        *NetworkForwardNAT        `xml:"nat" json:"nat,omitempty" protobuf:"bytes,6,opt,name=nat"`
	Interfaces []NetworkForwardInterface `xml:"interface" json:"interface" protobuf:"bytes,7,rep,name=interface"`
	Addresses  []NetworkForwardAddress   `xml:"address" json:"address" protobuf:"bytes,8,rep,name=address"`
}

// +gogo:genproto=true
type NetworkForwardDriver struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type NetworkForwardPF struct {
	Dev string `xml:"dev,attr" json:"dev" protobuf:"bytes,1,opt,name=dev"`
}

// +gogo:genproto=true
type NetworkForwardAddress struct {
	PCI *NetworkForwardAddressPCI `xml:"-" json:"pci,omitempty" protobuf:"bytes,1,opt,name=pci"`
}

// +gogo:genproto=true
type NetworkForwardAddressPCI struct {
	Domain   *int32 `xml:"domain,attr" json:"domain,omitempty" protobuf:"varint,1,opt,name=domain"`
	Bus      *int32 `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Slot     *int32 `xml:"slot,attr" json:"slot,omitempty" protobuf:"varint,3,opt,name=slot"`
	Function *int32 `xml:"function,attr" json:"function,omitempty" protobuf:"varint,4,opt,name=function"`
}

// +gogo:genproto=true
type NetworkForwardInterface struct {
	XMLName xml.Name `xml:"interface" json:"-"`
	Dev     string   `xml:"dev,attr,omitempty" json:"dev,omitempty" protobuf:"bytes,1,opt,name=dev"`
}

// +gogo:genproto=true
type NetworkMAC struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type NetworkDHCPRange struct {
	XMLName xml.Name          `xml:"range" json:"-"`
	Start   string            `xml:"start,attr,omitempty" json:"start,omitempty" protobuf:"bytes,1,opt,name=start"`
	End     string            `xml:"end,attr,omitempty" json:"end,omitempty" protobuf:"bytes,2,opt,name=end"`
	Lease   *NetworkDHCPLease `xml:"lease" json:"lease,omitempty" protobuf:"bytes,3,opt,name=lease"`
}

// +gogo:genproto=true
type NetworkDHCPLease struct {
	Expiry int32  `xml:"expiry,attr" json:"expiry" protobuf:"varint,1,opt,name=expiry"`
	Unit   string `xml:"unit,attr,omitempty" json:"unit,omitempty" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type NetworkDHCPHost struct {
	XMLName xml.Name          `xml:"host" json:"-"`
	ID      string            `xml:"id,attr,omitempty" json:"id,omitempty" protobuf:"bytes,1,opt,name=id"`
	MAC     string            `xml:"mac,attr,omitempty" json:"mac,omitempty" protobuf:"bytes,2,opt,name=mac"`
	Name    string            `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	IP      string            `xml:"ip,attr,omitempty" json:"ip,omitempty" protobuf:"bytes,4,opt,name=ip"`
	Lease   *NetworkDHCPLease `xml:"lease" json:"lease,omitempty" protobuf:"bytes,5,opt,name=lease"`
}

// +gogo:genproto=true
type NetworkBootp struct {
	File   string `xml:"file,attr,omitempty" json:"file,omitempty" protobuf:"bytes,1,opt,name=file"`
	Server string `xml:"server,attr,omitempty" json:"server,omitempty" protobuf:"bytes,2,opt,name=server"`
}

// +gogo:genproto=true
type NetworkDHCP struct {
	Ranges []NetworkDHCPRange `xml:"range" json:"range" protobuf:"bytes,1,rep,name=range"`
	Hosts  []NetworkDHCPHost  `xml:"host" json:"host" protobuf:"bytes,2,rep,name=host"`
	Bootp  []NetworkBootp     `xml:"bootp" json:"bootp" protobuf:"bytes,3,rep,name=bootp"`
}

// +gogo:genproto=true
type NetworkIP struct {
	Address  string       `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	Family   string       `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,2,opt,name=family"`
	Netmask  string       `xml:"netmask,attr,omitempty" json:"netmask,omitempty" protobuf:"bytes,3,opt,name=netmask"`
	Prefix   int32        `xml:"prefix,attr,omitempty" json:"prefix,omitempty" protobuf:"varint,4,opt,name=prefix"`
	LocalPtr string       `xml:"localPtr,attr,omitempty" json:"localPtr,omitempty" protobuf:"bytes,5,opt,name=localPtr"`
	DHCP     *NetworkDHCP `xml:"dhcp" json:"dhcp,omitempty" protobuf:"bytes,6,opt,name=dhcp"`
	TFTP     *NetworkTFTP `xml:"tftp" json:"tftp,omitempty" protobuf:"bytes,7,opt,name=tftp"`
}

// +gogo:genproto=true
type NetworkTFTP struct {
	Root string `xml:"root,attr,omitempty" json:"root,omitempty" protobuf:"bytes,1,opt,name=root"`
}

// +gogo:genproto=true
type NetworkRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" protobuf:"bytes,1,opt,name=family"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty" protobuf:"bytes,3,opt,name=netmask"`
	Prefix  int32  `xml:"prefix,attr,omitempty" json:"prefix,omitempty" protobuf:"varint,4,opt,name=prefix"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty" protobuf:"bytes,5,opt,name=gateway"`
	Metric  string `xml:"metric,attr,omitempty" json:"metric,omitempty" protobuf:"bytes,6,opt,name=metric"`
}

// +gogo:genproto=true
type NetworkDNSForwarder struct {
	Domain string `xml:"domain,attr,omitempty" json:"domain,omitempty" protobuf:"bytes,1,opt,name=domain"`
	Addr   string `xml:"addr,attr,omitempty" json:"addr,omitempty" protobuf:"bytes,2,opt,name=addr"`
}

// +gogo:genproto=true
type NetworkDNSTXT struct {
	XMLName xml.Name `xml:"txt" json:"-"`
	Name    string   `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value   string   `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type NetworkDNSHostHostname struct {
	Hostname string `xml:",chardata" json:"hostname" protobuf:"bytes,1,opt,name=hostname"`
}

// +gogo:genproto=true
type NetworkDNSHost struct {
	XMLName   xml.Name                 `xml:"host" json:"-"`
	IP        string                   `xml:"ip,attr" json:"ip" protobuf:"bytes,1,opt,name=ip"`
	Hostnames []NetworkDNSHostHostname `xml:"hostname" json:"hostnames" protobuf:"bytes,2,rep,name=hostnames"`
}

// +gogo:genproto=true
type NetworkDNSSRV struct {
	XMLName  xml.Name `xml:"srv" json:"-"`
	Service  string   `xml:"service,attr,omitempty" json:"service,omitempty" protobuf:"bytes,1,opt,name=service"`
	Protocol string   `xml:"protocol,attr,omitempty" json:"protocol,omitempty" protobuf:"bytes,2,opt,name=protocol"`
	Target   string   `xml:"target,attr,omitempty" json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
	Port     int32    `xml:"port,attr,omitempty" json:"port,omitempty" protobuf:"varint,4,opt,name=port"`
	Priority int32    `xml:"priority,attr,omitempty" json:"priority,omitempty" protobuf:"varint,5,opt,name=priority"`
	Weight   int32    `xml:"weight,attr,omitempty" json:"weight,omitempty" protobuf:"varint,6,opt,name=weight"`
	Domain   string   `xml:"domain,attr,omitempty" json:"domain,omitempty" protobuf:"bytes,7,opt,name=domain"`
}

// +gogo:genproto=true
type NetworkDNS struct {
	Enable            string                `xml:"enable,attr,omitempty" json:"enable,omitempty" protobuf:"bytes,1,opt,name=enable"`
	ForwardPlainNames string                `xml:"forwardPlainNames,attr,omitempty" json:"forwardPlainNames,omitempty" protobuf:"bytes,2,opt,name=forwardPlainNames"`
	Forwarders        []NetworkDNSForwarder `xml:"forwarder" json:"forwarders" protobuf:"bytes,3,rep,name=forwarders"`
	TXTs              []NetworkDNSTXT       `xml:"txt" json:"txt" protobuf:"bytes,4,rep,name=txt"`
	Host              []NetworkDNSHost      `xml:"host" json:"host" protobuf:"bytes,5,rep,name=host"`
	SRVs              []NetworkDNSSRV       `xml:"srv" json:"srv" protobuf:"bytes,6,rep,name=srv"`
}

// +gogo:genproto=true
type NetworkMetadata struct {
	XML string `xml:",innerxml" json:",inline" protobuf:"bytes,1,opt,name=xML"`
}

// +gogo:genproto=true
type NetworkMTU struct {
	Size int32 `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
}

// +gogo:genproto=true
type Network struct {
	XMLName             xml.Name            `xml:"network" json:"-"`
	IPv6                string              `xml:"ipv6,attr,omitempty" json:"ipv6,omitempty" protobuf:"bytes,1,opt,name=ipv6"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty" protobuf:"bytes,2,opt,name=trustGuestRxFilters"`
	Name                string              `xml:"name,omitempty" json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	UUID                string              `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,4,opt,name=uuid"`
	Metadata            *NetworkMetadata    `xml:"metadata" json:"metadata,omitempty" protobuf:"bytes,5,opt,name=metadata"`
	Forward             *NetworkForward     `xml:"forward" json:"forward,omitempty" protobuf:"bytes,6,opt,name=forward"`
	Bridge              *NetworkBridge      `xml:"bridge" json:"bridge,omitempty" protobuf:"bytes,7,opt,name=bridge"`
	MTU                 *NetworkMTU         `xml:"mtu" json:"mtu,omitempty" protobuf:"bytes,8,opt,name=mtu"`
	MAC                 *NetworkMAC         `xml:"mac" json:"mac,omitempty" protobuf:"bytes,9,opt,name=mac"`
	Domain              *NetworkDomain      `xml:"domain" json:"domain,omitempty" protobuf:"bytes,10,opt,name=domain"`
	DNS                 *NetworkDNS         `xml:"dns" json:"dns,omitempty" protobuf:"bytes,11,opt,name=dns"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,12,opt,name=vlan"`
	Bandwidth           *NetworkBandwidth   `xml:"bandwidth" json:"bandwidth,omitempty" protobuf:"bytes,13,opt,name=bandwidth"`
	PortOptions         *NetworkPortOptions `xml:"port" json:"portOptions,omitempty" protobuf:"bytes,14,opt,name=portOptions"`
	IPs                 []NetworkIP         `xml:"ip" json:"ip" protobuf:"bytes,15,rep,name=ip"`
	Routes              []NetworkRoute      `xml:"route" json:"route" protobuf:"bytes,16,rep,name=route"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualPort" protobuf:"bytes,17,opt,name=virtualPort"`
	PortGroups          []NetworkPortGroup  `xml:"portgroup" json:"portGroup" protobuf:"bytes,18,rep,name=portGroup"`

	DnsmasqOptions *NetworkDnsmasqOptions `json:"dnsmasqOptions,omitempty" protobuf:"bytes,19,opt,name=dnsmasqOptions"`
}

// +gogo:genproto=true
type NetworkPortOptions struct {
	Isolated string `xml:"isolated,attr,omitempty" json:"isolated,omitempty" protobuf:"bytes,1,opt,name=isolated"`
}

// +gogo:genproto=true
type NetworkPortGroup struct {
	XMLName             xml.Name            `xml:"portgroup" json:"-"`
	Name                string              `xml:"name,attr,omitempty" json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Default             string              `xml:"default,attr,omitempty" json:"default,omitempty" protobuf:"bytes,2,opt,name=default"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty" protobuf:"bytes,3,opt,name=trustGuestRxFilters"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,4,opt,name=vlan"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualPort,omitempty" protobuf:"bytes,5,opt,name=virtualPort"`
}

// +gogo:genproto=true
type NetworkVLAN struct {
	Trunk string           `xml:"trunk,attr,omitempty" json:"trunk,omitempty" protobuf:"bytes,1,opt,name=trunk"`
	Tags  []NetworkVLANTag `xml:"tag" json:"tag" protobuf:"bytes,2,rep,name=tag"`
}

// +gogo:genproto=true
type NetworkVLANTag struct {
	ID         int32  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty" protobuf:"bytes,2,opt,name=nativeMode"`
}

// +gogo:genproto=true
type NetworkBandwidthParams struct {
	Average *int32 `xml:"average,attr" json:"average" protobuf:"varint,1,opt,name=average"`
	Peak    *int32 `xml:"peak,attr" json:"peak" protobuf:"varint,2,opt,name=peak"`
	Burst   *int32 `xml:"burst,attr" json:"burst" protobuf:"varint,3,opt,name=burst"`
	Floor   *int32 `xml:"floor,attr" json:"floor" protobuf:"varint,4,opt,name=floor"`
}

// +gogo:genproto=true
type NetworkBandwidth struct {
	ClassID  int32                   `xml:"classID,attr,omitempty" json:"classId" protobuf:"varint,1,opt,name=classId"`
	Inbound  *NetworkBandwidthParams `xml:"inbound" json:"inbound,omitempty" protobuf:"bytes,2,opt,name=inbound"`
	Outbound *NetworkBandwidthParams `xml:"outbound" json:"outbound,omitempty" protobuf:"bytes,3,opt,name=outbound"`
}

// +gogo:genproto=true
type NetworkDnsmasqOptions struct {
	XMLName xml.Name               `xml:"http://libvirt.org/schemas/network/dnsmasq/1.0 options" json:"-"`
	Option  []NetworkDnsmasqOption `xml:"option" json:"option" protobuf:"bytes,1,rep,name=option"`
}

// +gogo:genproto=true
type NetworkDnsmasqOption struct {
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,1,opt,name=value"`
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
			if err := unmarshalIntAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalIntAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalIntAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshalIntAttr(attr.Value, &a.Function, 0); err != nil {
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

func (s *NetworkDHCPHost) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPHost) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSHost) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSHost) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkPortGroup) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkPortGroup) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSTXT) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSTXT) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSSRV) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSSRV) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDHCPRange) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPRange) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkForwardInterface) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkForwardInterface) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *Network) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Network) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
