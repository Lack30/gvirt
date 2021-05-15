package spec

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NWFilter struct {
	XMLName  xml.Name        `xml:"filter" json:"-"`
	Name     string          `xml:"name,attr" json:"name"`
	UUID     string          `xml:"uuid,omitempty" json:"uuid,omitempty"`
	Chain    string          `xml:"chain,attr,omitempty" json:"chain,omitempty"`
	Priority int             `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	Entries  []NWFilterEntry `json:"entries"`
}

type NWFilterEntry struct {
	Rule *NWFilterRule `json:"rule,omitempty"`
	Ref  *NWFilterRef  `json:"ref,omitempty"`
}

type NWFilterRef struct {
	Filter     string              `xml:"filter,attr" json:"filter"`
	Parameters []NWFilterParameter `xml:"parameter" json:"parameters"`
}

type NWFilterParameter struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

type NWFilterField struct {
	Var  string `json:"var"`
	Str  string `json:"str"`
	Uint *uint  `json:"uint,omitempty"`
}

type NWFilterRule struct {
	Action     string `xml:"action,attr,omitempty" json:"action,omitempty"`
	Direction  string `xml:"direction,attr,omitempty" json:"direction,omitempty"`
	Priority   int    `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	StateMatch string `xml:"statematch,attr,omitempty" json:"statematch,omitempty"`

	ARP         *NWFilterRuleARP         `xml:"arp" json:"arp,omitempty"`
	RARP        *NWFilterRuleRARP        `xml:"rarp" json:"rarp,omitempty"`
	MAC         *NWFilterRuleMAC         `xml:"mac" json:"mac,omitempty"`
	VLAN        *NWFilterRuleVLAN        `xml:"vlan" json:"vlan,omitempty"`
	STP         *NWFilterRuleSTP         `xml:"stp" json:"stp,omitempty"`
	IP          *NWFilterRuleIP          `xml:"ip" json:"ip,omitempty"`
	IPv6        *NWFilterRuleIPv6        `xml:"ipv6" json:"ipv6,omitempty"`
	TCP         *NWFilterRuleTCP         `xml:"tcp" json:"tcp,omitempty"`
	UDP         *NWFilterRuleUDP         `xml:"udp" json:"udp,omitempty"`
	UDPLite     *NWFilterRuleUDPLite     `xml:"udplite" json:"udplite,omitempty"`
	ESP         *NWFilterRuleESP         `xml:"esp" json:"esp,omitempty"`
	AH          *NWFilterRuleAH          `xml:"ah" json:"ah,omitempty"`
	SCTP        *NWFilterRuleSCTP        `xml:"sctp" json:"sctp,omitempty"`
	ICMP        *NWFilterRuleICMP        `xml:"icmp" json:"icmp,omitempty"`
	All         *NWFilterRuleAll         `xml:"all" json:"all,omitempty"`
	IGMP        *NWFilterRuleIGMP        `xml:"igmp" json:"igmp,omitempty"`
	TCPIPv6     *NWFilterRuleTCPIPv6     `xml:"tcp-ipv6" json:"tcpipv6,omitempty"`
	UDPIPv6     *NWFilterRuleUDPIPv6     `xml:"udp-ipv6" json:"udpipv6,omitempty"`
	UDPLiteIPv6 *NWFilterRuleUDPLiteIPv6 `xml:"udplite-ipv6" json:"udpliteipv6,omitempty"`
	ESPIPv6     *NWFilterRuleESPIPv6     `xml:"esp-ipv6" json:"espipv6,omitempty"`
	AHIPv6      *NWFilterRuleAHIPv6      `xml:"ah-ipv6" json:"ahipv6,omitempty"`
	SCTPIPv6    *NWFilterRuleSCTPIPv6    `xml:"sctp-ipv6" json:"sctpipv6,omitempty"`
	ICMPv6      *NWFilterRuleICMPIPv6    `xml:"icmpv6" json:"icmpv6,omitempty"`
	AllIPv6     *NWFilterRuleAllIPv6     `xml:"all-ipv6" json:"allipv6,omitempty"`
}

type NWFilterRuleCommonMAC struct {
	SrcMACAddr NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty"`
	SrcMACMask NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcMacMask,omitempty"`
	DstMACAddr NWFilterField `xml:"dstmacaddr,attr,omitempty" json:"dstMacAddr,omitempty"`
	DstMACMask NWFilterField `xml:"dstmacmask,attr,omitempty" json:"dstMacMask,omitempty"`
}

type NWFilterRuleCommonIP struct {
	SrcMACAddr     NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty"`
	SrcIPAddr      NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty"`
	SrcIPMask      NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty"`
	DstIPAddr      NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty"`
	DstIPMask      NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty"`
	SrcIPFrom      NWFilterField `xml:"srcipfrom,attr,omitempty" json:"srcIpFrom,omitempty"`
	SrcIPTo        NWFilterField `xml:"srcipto,attr,omitempty" json:"srcIpTo,omitempty"`
	DstIPFrom      NWFilterField `xml:"dstipfrom,attr,omitempty" json:"dstIpFrom,omitempty"`
	DstIPTo        NWFilterField `xml:"dstipto,attr,omitempty" json:"dstIpTo,omitempty"`
	DSCP           NWFilterField `xml:"dscp,attr" json:"dscp"`
	ConnLimitAbove NWFilterField `xml:"connlimit-above,attr" json:"connLimitAbove,omitempty"`
	State          NWFilterField `xml:"state,attr,omitempty" json:"state,omitempty"`
	IPSet          NWFilterField `xml:"ipset,attr,omitempty" json:"ipSet,omitempty"`
	IPSetFlags     NWFilterField `xml:"ipsetflags,attr,omitempty" json:"ipSetFlags,omitempty"`
}

type NWFilterRuleCommonPort struct {
	SrcPortStart NWFilterField `xml:"srcportstart,attr" json:"srcPortStart"`
	SrcPortEnd   NWFilterField `xml:"srcportend,attr" json:"srcPortEnd"`
	DstPortStart NWFilterField `xml:"dstportstart,attr" json:"dstPortStart"`
	DstPortEnd   NWFilterField `xml:"dstportend,attr" json:"dstPortEnd"`
}

type NWFilterRuleARP struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	HWType                NWFilterField `xml:"hwtype,attr" json:"hwType"`
	ProtocolType          NWFilterField `xml:"protocoltype,attr" json:"protocolType"`
	OpCode                NWFilterField `xml:"opcode,attr,omitempty" json:"opCode,omitempty"`
	ARPSrcMACAddr         NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpSrcMacAddr,omitempty"`
	ARPDstMACAddr         NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpDstMacAddr,omitempty"`
	ARPSrcIPAddr          NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpSrcIpAddr,omitempty"`
	ARPSrcIPMask          NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpSrcIpMask,omitempty"`
	ARPDstIPAddr          NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpDstIpAddr,omitempty"`
	ARPDstIPMask          NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpDstIpMask,omitempty"`
	Gratuitous            NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleRARP struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	HWType                NWFilterField `xml:"hwtype,attr" json:"hwType"`
	ProtocolType          NWFilterField `xml:"protocoltype,attr" json:"protocolType"`
	OpCode                NWFilterField `xml:"opcode,attr,omitempty" json:"opCode,omitempty"`
	ARPSrcMACAddr         NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpSrcMacAddr,omitempty"`
	ARPDstMACAddr         NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpDstMacAddr,omitempty"`
	ARPSrcIPAddr          NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpSrcIpAddr,omitempty"`
	ARPSrcIPMask          NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpSrcIpMask,omitempty"`
	ARPDstIPAddr          NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpDstIpAddr,omitempty"`
	ARPDstIPMask          NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpDstIpMask,omitempty"`
	Gratuitous            NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleMAC struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	ProtocolID NWFilterField `xml:"protocolid,attr,omitempty" json:"protocolId,omitempty"`
	Comment    string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleVLAN struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	VLANID        NWFilterField `xml:"vlanid,attr,omitempty" json:"vlanid,omitempty"`
	EncapProtocol NWFilterField `xml:"encap-protocol,attr,omitempty" json:"encapProtocol"`
	Comment       string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleSTP struct {
	Match             NWFilterField `xml:"match,attr,omitempty" json:"match,omitempty"`
	SrcMACAddr        NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty"`
	SrcMACMask        NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcMacMask,omitempty"`
	Type              NWFilterField `xml:"type,attr" json:"type"`
	Flags             NWFilterField `xml:"flags,attr" json:"flags"`
	RootPriority      NWFilterField `xml:"root-priority,attr" json:"rootPriority"`
	RootPriorityHi    NWFilterField `xml:"root-priority-hi,attr" json:"rootPriorityHi"`
	RootAddress       NWFilterField `xml:"root-address,attr,omitempty" json:"rootAddress,omitempty"`
	RootAddressMask   NWFilterField `xml:"root-address-mask,attr,omitempty" json:"rootAddressMask,omitempty"`
	RootCost          NWFilterField `xml:"root-cost,attr" json:"rootCost"`
	RootCostHi        NWFilterField `xml:"root-cost-hi,attr" json:"rootCostHi"`
	SenderPriority    NWFilterField `xml:"sender-priority,attr" json:"senderPriority"`
	SenderPriorityHi  NWFilterField `xml:"sender-priority-hi,attr" json:"senderPriorityHi"`
	SenderAddress     NWFilterField `xml:"sender-address,attr,omitempty" json:"senderAddress,omitempty"`
	SenderAddressMask NWFilterField `xml:"sender-address-mask,attr,omitempty" json:"senderAddressMask,omitempty"`
	Port              NWFilterField `xml:"port,attr" json:"port"`
	PortHi            NWFilterField `xml:"port-hi,attr" json:"portHi"`
	Age               NWFilterField `xml:"age,attr" json:"age"`
	AgeHi             NWFilterField `xml:"age-hi,attr" json:"ageHi"`
	MaxAge            NWFilterField `xml:"max-age,attr" json:"maxAge"`
	MaxAgeHi          NWFilterField `xml:"max-age-hi,attr" json:"maxAgeHi"`
	HelloTime         NWFilterField `xml:"hello-time,attr" json:"helloTime"`
	HelloTimeHi       NWFilterField `xml:"hello-time-hi,attr" json:"helloTimeHi"`
	ForwardDelay      NWFilterField `xml:"forward-delay,attr" json:"forwardDelay,omitempty"`
	ForwardDelayHi    NWFilterField `xml:"forward-delay-hi,attr" json:"forwardDelayHi,omitempty"`
	Comment           string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleIP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	NWFilterRuleCommonPort `json:",inline"`
	DSCP    NWFilterField `xml:"dscp,attr" json:"dscp,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonMAC `json:",inline"`
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty"`
	NWFilterRuleCommonPort `json:",inline"`
	Type    NWFilterField `xml:"type,attr" json:"type"`
	TypeEnd NWFilterField `xml:"typeend,attr" json:"typeEnd"`
	Code    NWFilterField `xml:"code,attr" json:"code"`
	CodeEnd NWFilterField `xml:"codeend,attr" json:"codeEnd"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment"`
}

type NWFilterRuleTCP struct {
	Match string `xml:"match,attr,omitempty" json:"match"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Option  NWFilterField `xml:"option,attr" json:"option"`
	Flags   NWFilterField `xml:"flags,attr,omitempty" json:"flags,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleUDP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleUDPLite struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleESP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleAH struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleSCTP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleICMP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Type    NWFilterField `xml:"type,attr" json:"type"`
	Code    NWFilterField `xml:"code,attr" json:"code"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleAll struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleIGMP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleTCPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Option  NWFilterField `xml:"option,attr" json:"option"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleUDPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleUDPLiteIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleESPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleAHIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleSCTPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	NWFilterRuleCommonPort `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleICMPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Type    NWFilterField `xml:"type,attr" json:"type"`
	Code    NWFilterField `xml:"code,attr" json:"code"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

type NWFilterRuleAllIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty"`
	NWFilterRuleCommonIP `json:",inline"`
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty"`
}

func (s *NWFilterField) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s == nil {
		return xml.Attr{}, nil
	}
	if s.Str != "" {
		return xml.Attr{
			Name:  name,
			Value: s.Str,
		}, nil
	} else if s.Var != "" {
		return xml.Attr{
			Name:  name,
			Value: "$" + s.Str,
		}, nil
	} else if s.Uint != nil {
		return xml.Attr{
			Name:  name,
			Value: fmt.Sprintf("0x%x", *s.Uint),
		}, nil
	} else {
		return xml.Attr{}, nil
	}
}

func (s *NWFilterField) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}
	if attr.Value[0] == '$' {
		s.Var = attr.Value[1:]
	}
	if strings.HasPrefix(attr.Value, "0x") {
		val, err := strconv.ParseUint(attr.Value[2:], 16, 64)
		if err != nil {
			return err
		}
		uval := uint(val)
		s.Uint = &uval
	}
	s.Str = attr.Value
	return nil
}

func (a *NWFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "filter"
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "name"},
		Value: a.Name,
	})
	if a.Chain != "" {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "chain"},
			Value: a.Chain,
		})
	}
	if a.Priority != 0 {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "priority"},
			Value: fmt.Sprintf("%d", a.Priority),
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	if a.UUID != "" {
		uuid := xml.StartElement{
			Name: xml.Name{Local: "uuid"},
		}
		e.EncodeToken(uuid)
		e.EncodeToken(xml.CharData(a.UUID))
		e.EncodeToken(uuid.End())
	}

	for _, entry := range a.Entries {
		if entry.Rule != nil {
			rule := xml.StartElement{
				Name: xml.Name{Local: "rule"},
			}
			e.EncodeElement(entry.Rule, rule)
		} else if entry.Ref != nil {
			ref := xml.StartElement{
				Name: xml.Name{Local: "filterref"},
			}
			e.EncodeElement(entry.Ref, ref)
		}
	}

	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}

func (a *NWFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	name, ok := getAttr(start.Attr, "name")
	if !ok {
		return fmt.Errorf("Missing filter name")
	}
	a.Name = name
	a.Chain, _ = getAttr(start.Attr, "chain")
	prio, ok := getAttr(start.Attr, "priority")
	if ok {
		val, err := strconv.ParseInt(prio, 10, 64)
		if err != nil {
			return err
		}
		a.Priority = int(val)
	}

	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			{
				if tok.Name.Local == "uuid" {
					txt, err := d.Token()
					if err != nil {
						return err
					}

					txt2, ok := txt.(xml.CharData)
					if !ok {
						return fmt.Errorf("Expected UUID string")
					}
					a.UUID = string(txt2)
				} else if tok.Name.Local == "rule" {
					entry := NWFilterEntry{
						Rule: &NWFilterRule{},
					}

					d.DecodeElement(entry.Rule, &tok)

					a.Entries = append(a.Entries, entry)
				} else if tok.Name.Local == "filterref" {
					entry := NWFilterEntry{
						Ref: &NWFilterRef{},
					}

					d.DecodeElement(entry.Ref, &tok)

					a.Entries = append(a.Entries, entry)
				}
			}
		}

	}
	return nil
}

func (s *NWFilter) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NWFilter) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
