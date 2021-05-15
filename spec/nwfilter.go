package spec

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// +gogo:genproto=true
type NWFilter struct {
	XMLName  xml.Name        `xml:"filter" json:"-"`
	Name     string          `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	UUID     string          `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,2,opt,name=uuid"`
	Chain    string          `xml:"chain,attr,omitempty" json:"chain,omitempty" protobuf:"bytes,3,opt,name=chain"`
	Priority int             `xml:"priority,attr,omitempty" json:"priority,omitempty" protobuf:"varint,4,opt,name=priority"`
	Entries  []NWFilterEntry `json:"entries" protobuf:"bytes,5,rep,name=entries"`
}

// +gogo:genproto=true
type NWFilterEntry struct {
	Rule *NWFilterRule `json:"rule,omitempty" protobuf:"bytes,1,opt,name=rule"`
	Ref  *NWFilterRef  `json:"ref,omitempty" protobuf:"bytes,2,opt,name=ref"`
}

// +gogo:genproto=true
type NWFilterRef struct {
	Filter     string              `xml:"filter,attr" json:"filter" protobuf:"bytes,1,opt,name=filter"`
	Parameters []NWFilterParameter `xml:"parameter" json:"parameters" protobuf:"bytes,2,rep,name=parameters"`
}

// +gogo:genproto=true
type NWFilterParameter struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type NWFilterField struct {
	Var  string `json:"var" protobuf:"bytes,1,opt,name=var"`
	Str  string `json:"str" protobuf:"bytes,2,opt,name=str"`
	Uint *uint  `json:"uint,omitempty" protobuf:"varint,3,opt,name=uint"`
}

// +gogo:genproto=true
type NWFilterRule struct {
	Action     string `xml:"action,attr,omitempty" json:"action,omitempty" protobuf:"bytes,1,opt,name=action"`
	Direction  string `xml:"direction,attr,omitempty" json:"direction,omitempty" protobuf:"bytes,2,opt,name=direction"`
	Priority   int    `xml:"priority,attr,omitempty" json:"priority,omitempty" protobuf:"varint,3,opt,name=priority"`
	StateMatch string `xml:"statematch,attr,omitempty" json:"statematch,omitempty" protobuf:"bytes,4,opt,name=statematch"`

	ARP         *NWFilterRuleARP         `xml:"arp" json:"arp,omitempty" protobuf:"bytes,5,opt,name=arp"`
	RARP        *NWFilterRuleRARP        `xml:"rarp" json:"rarp,omitempty" protobuf:"bytes,6,opt,name=rarp"`
	MAC         *NWFilterRuleMAC         `xml:"mac" json:"mac,omitempty" protobuf:"bytes,7,opt,name=mac"`
	VLAN        *NWFilterRuleVLAN        `xml:"vlan" json:"vlan,omitempty" protobuf:"bytes,8,opt,name=vlan"`
	STP         *NWFilterRuleSTP         `xml:"stp" json:"stp,omitempty" protobuf:"bytes,9,opt,name=stp"`
	IP          *NWFilterRuleIP          `xml:"ip" json:"ip,omitempty" protobuf:"bytes,10,opt,name=ip"`
	IPv6        *NWFilterRuleIPv6        `xml:"ipv6" json:"ipv6,omitempty" protobuf:"bytes,11,opt,name=ipv6"`
	TCP         *NWFilterRuleTCP         `xml:"tcp" json:"tcp,omitempty" protobuf:"bytes,12,opt,name=tcp"`
	UDP         *NWFilterRuleUDP         `xml:"udp" json:"udp,omitempty" protobuf:"bytes,13,opt,name=udp"`
	UDPLite     *NWFilterRuleUDPLite     `xml:"udplite" json:"udplite,omitempty" protobuf:"bytes,14,opt,name=udplite"`
	ESP         *NWFilterRuleESP         `xml:"esp" json:"esp,omitempty" protobuf:"bytes,15,opt,name=esp"`
	AH          *NWFilterRuleAH          `xml:"ah" json:"ah,omitempty" protobuf:"bytes,16,opt,name=ah"`
	SCTP        *NWFilterRuleSCTP        `xml:"sctp" json:"sctp,omitempty" protobuf:"bytes,17,opt,name=sctp"`
	ICMP        *NWFilterRuleICMP        `xml:"icmp" json:"icmp,omitempty" protobuf:"bytes,18,opt,name=icmp"`
	All         *NWFilterRuleAll         `xml:"all" json:"all,omitempty" protobuf:"bytes,19,opt,name=all"`
	IGMP        *NWFilterRuleIGMP        `xml:"igmp" json:"igmp,omitempty" protobuf:"bytes,20,opt,name=igmp"`
	TCPIPv6     *NWFilterRuleTCPIPv6     `xml:"tcp-ipv6" json:"tcpipv6,omitempty" protobuf:"bytes,21,opt,name=tcpipv6"`
	UDPIPv6     *NWFilterRuleUDPIPv6     `xml:"udp-ipv6" json:"udpipv6,omitempty" protobuf:"bytes,22,opt,name=udpipv6"`
	UDPLiteIPv6 *NWFilterRuleUDPLiteIPv6 `xml:"udplite-ipv6" json:"udpliteipv6,omitempty" protobuf:"bytes,23,opt,name=udpliteipv6"`
	ESPIPv6     *NWFilterRuleESPIPv6     `xml:"esp-ipv6" json:"espipv6,omitempty" protobuf:"bytes,24,opt,name=espipv6"`
	AHIPv6      *NWFilterRuleAHIPv6      `xml:"ah-ipv6" json:"ahipv6,omitempty" protobuf:"bytes,25,opt,name=ahipv6"`
	SCTPIPv6    *NWFilterRuleSCTPIPv6    `xml:"sctp-ipv6" json:"sctpipv6,omitempty" protobuf:"bytes,26,opt,name=sctpipv6"`
	ICMPv6      *NWFilterRuleICMPIPv6    `xml:"icmpv6" json:"icmpv6,omitempty" protobuf:"bytes,27,opt,name=icmpv6"`
	AllIPv6     *NWFilterRuleAllIPv6     `xml:"all-ipv6" json:"allipv6,omitempty" protobuf:"bytes,28,opt,name=allipv6"`
}

// +gogo:genproto=true
type NWFilterRuleCommonMAC struct {
	SrcMACAddr NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty" protobuf:"bytes,1,opt,name=srcMacAddr"`
	SrcMACMask NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcMacMask,omitempty" protobuf:"bytes,2,opt,name=srcMacMask"`
	DstMACAddr NWFilterField `xml:"dstmacaddr,attr,omitempty" json:"dstMacAddr,omitempty" protobuf:"bytes,3,opt,name=dstMacAddr"`
	DstMACMask NWFilterField `xml:"dstmacmask,attr,omitempty" json:"dstMacMask,omitempty" protobuf:"bytes,4,opt,name=dstMacMask"`
}

// +gogo:genproto=true
type NWFilterRuleCommonIP struct {
	SrcMACAddr     NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty" protobuf:"bytes,1,opt,name=srcMacAddr"`
	SrcIPAddr      NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty" protobuf:"bytes,2,opt,name=srcIpAddr"`
	SrcIPMask      NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty" protobuf:"bytes,3,opt,name=srcIpMask"`
	DstIPAddr      NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty" protobuf:"bytes,4,opt,name=dstIpAddr"`
	DstIPMask      NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty" protobuf:"bytes,5,opt,name=dstIpMask"`
	SrcIPFrom      NWFilterField `xml:"srcipfrom,attr,omitempty" json:"srcIpFrom,omitempty" protobuf:"bytes,6,opt,name=srcIpFrom"`
	SrcIPTo        NWFilterField `xml:"srcipto,attr,omitempty" json:"srcIpTo,omitempty" protobuf:"bytes,7,opt,name=srcIpTo"`
	DstIPFrom      NWFilterField `xml:"dstipfrom,attr,omitempty" json:"dstIpFrom,omitempty" protobuf:"bytes,8,opt,name=dstIpFrom"`
	DstIPTo        NWFilterField `xml:"dstipto,attr,omitempty" json:"dstIpTo,omitempty" protobuf:"bytes,9,opt,name=dstIpTo"`
	DSCP           NWFilterField `xml:"dscp,attr" json:"dscp" protobuf:"bytes,10,opt,name=dscp"`
	ConnLimitAbove NWFilterField `xml:"connlimit-above,attr" json:"connLimitAbove,omitempty" protobuf:"bytes,11,opt,name=connLimitAbove"`
	State          NWFilterField `xml:"state,attr,omitempty" json:"state,omitempty" protobuf:"bytes,12,opt,name=state"`
	IPSet          NWFilterField `xml:"ipset,attr,omitempty" json:"ipSet,omitempty" protobuf:"bytes,13,opt,name=ipSet"`
	IPSetFlags     NWFilterField `xml:"ipsetflags,attr,omitempty" json:"ipSetFlags,omitempty" protobuf:"bytes,14,opt,name=ipSetFlags"`
}

// +gogo:genproto=true
type NWFilterRuleCommonPort struct {
	SrcPortStart NWFilterField `xml:"srcportstart,attr" json:"srcPortStart" protobuf:"bytes,1,opt,name=srcPortStart"`
	SrcPortEnd   NWFilterField `xml:"srcportend,attr" json:"srcPortEnd" protobuf:"bytes,2,opt,name=srcPortEnd"`
	DstPortStart NWFilterField `xml:"dstportstart,attr" json:"dstPortStart" protobuf:"bytes,3,opt,name=dstPortStart"`
	DstPortEnd   NWFilterField `xml:"dstportend,attr" json:"dstPortEnd" protobuf:"bytes,4,opt,name=dstPortEnd"`
}

// +gogo:genproto=true
type NWFilterRuleARP struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	HWType                NWFilterField `xml:"hwtype,attr" json:"hwType" protobuf:"bytes,3,opt,name=hwType"`
	ProtocolType          NWFilterField `xml:"protocoltype,attr" json:"protocolType" protobuf:"bytes,4,opt,name=protocolType"`
	OpCode                NWFilterField `xml:"opcode,attr,omitempty" json:"opCode,omitempty" protobuf:"bytes,5,opt,name=opCode"`
	ARPSrcMACAddr         NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpSrcMacAddr,omitempty" protobuf:"bytes,6,opt,name=arpSrcMacAddr"`
	ARPDstMACAddr         NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpDstMacAddr,omitempty" protobuf:"bytes,7,opt,name=arpDstMacAddr"`
	ARPSrcIPAddr          NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpSrcIpAddr,omitempty" protobuf:"bytes,8,opt,name=arpSrcIpAddr"`
	ARPSrcIPMask          NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpSrcIpMask,omitempty" protobuf:"bytes,9,opt,name=arpSrcIpMask"`
	ARPDstIPAddr          NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpDstIpAddr,omitempty" protobuf:"bytes,10,opt,name=arpDstIpAddr"`
	ARPDstIPMask          NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpDstIpMask,omitempty" protobuf:"bytes,11,opt,name=arpDstIpMask"`
	Gratuitous            NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty" protobuf:"bytes,12,opt,name=gratuitous"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,13,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleRARP struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	HWType                NWFilterField `xml:"hwtype,attr" json:"hwType" protobuf:"bytes,3,opt,name=hwType"`
	ProtocolType          NWFilterField `xml:"protocoltype,attr" json:"protocolType" protobuf:"bytes,4,opt,name=protocolType"`
	OpCode                NWFilterField `xml:"opcode,attr,omitempty" json:"opCode,omitempty" protobuf:"bytes,5,opt,name=opCode"`
	ARPSrcMACAddr         NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpSrcMacAddr,omitempty" protobuf:"bytes,6,opt,name=arpSrcMacAddr"`
	ARPDstMACAddr         NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpDstMacAddr,omitempty" protobuf:"bytes,7,opt,name=arpDstMacAddr"`
	ARPSrcIPAddr          NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpSrcIpAddr,omitempty" protobuf:"bytes,8,opt,name=arpSrcIpAddr"`
	ARPSrcIPMask          NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpSrcIpMask,omitempty" protobuf:"bytes,9,opt,name=arpSrcIpMask"`
	ARPDstIPAddr          NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpDstIpAddr,omitempty" protobuf:"bytes,10,opt,name=arpDstIpAddr"`
	ARPDstIPMask          NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpDstIpMask,omitempty" protobuf:"bytes,11,opt,name=arpDstIpMask"`
	Gratuitous            NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty" protobuf:"bytes,12,opt,name=gratuitous"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,13,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleMAC struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	ProtocolID            NWFilterField `xml:"protocolid,attr,omitempty" json:"protocolId,omitempty" protobuf:"bytes,3,opt,name=protocolId"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,4,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleVLAN struct {
	Match                 string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	VLANID                NWFilterField `xml:"vlanid,attr,omitempty" json:"vlanid,omitempty" protobuf:"bytes,3,opt,name=vlanid"`
	EncapProtocol         NWFilterField `xml:"encap-protocol,attr,omitempty" json:"encapProtocol" protobuf:"bytes,4,opt,name=encapProtocol"`
	Comment               string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,5,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleSTP struct {
	Match             NWFilterField `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	SrcMACAddr        NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcMacAddr,omitempty" protobuf:"bytes,2,opt,name=srcMacAddr"`
	SrcMACMask        NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcMacMask,omitempty" protobuf:"bytes,3,opt,name=srcMacMask"`
	Type              NWFilterField `xml:"type,attr" json:"type" protobuf:"bytes,4,opt,name=type"`
	Flags             NWFilterField `xml:"flags,attr" json:"flags" protobuf:"bytes,5,opt,name=flags"`
	RootPriority      NWFilterField `xml:"root-priority,attr" json:"rootPriority" protobuf:"bytes,6,opt,name=rootPriority"`
	RootPriorityHi    NWFilterField `xml:"root-priority-hi,attr" json:"rootPriorityHi" protobuf:"bytes,7,opt,name=rootPriorityHi"`
	RootAddress       NWFilterField `xml:"root-address,attr,omitempty" json:"rootAddress,omitempty" protobuf:"bytes,8,opt,name=rootAddress"`
	RootAddressMask   NWFilterField `xml:"root-address-mask,attr,omitempty" json:"rootAddressMask,omitempty" protobuf:"bytes,9,opt,name=rootAddressMask"`
	RootCost          NWFilterField `xml:"root-cost,attr" json:"rootCost" protobuf:"bytes,10,opt,name=rootCost"`
	RootCostHi        NWFilterField `xml:"root-cost-hi,attr" json:"rootCostHi" protobuf:"bytes,11,opt,name=rootCostHi"`
	SenderPriority    NWFilterField `xml:"sender-priority,attr" json:"senderPriority" protobuf:"bytes,12,opt,name=senderPriority"`
	SenderPriorityHi  NWFilterField `xml:"sender-priority-hi,attr" json:"senderPriorityHi" protobuf:"bytes,13,opt,name=senderPriorityHi"`
	SenderAddress     NWFilterField `xml:"sender-address,attr,omitempty" json:"senderAddress,omitempty" protobuf:"bytes,14,opt,name=senderAddress"`
	SenderAddressMask NWFilterField `xml:"sender-address-mask,attr,omitempty" json:"senderAddressMask,omitempty" protobuf:"bytes,15,opt,name=senderAddressMask"`
	Port              NWFilterField `xml:"port,attr" json:"port" protobuf:"bytes,16,opt,name=port"`
	PortHi            NWFilterField `xml:"port-hi,attr" json:"portHi" protobuf:"bytes,17,opt,name=portHi"`
	Age               NWFilterField `xml:"age,attr" json:"age" protobuf:"bytes,18,opt,name=age"`
	AgeHi             NWFilterField `xml:"age-hi,attr" json:"ageHi" protobuf:"bytes,19,opt,name=ageHi"`
	MaxAge            NWFilterField `xml:"max-age,attr" json:"maxAge" protobuf:"bytes,20,opt,name=maxAge"`
	MaxAgeHi          NWFilterField `xml:"max-age-hi,attr" json:"maxAgeHi" protobuf:"bytes,21,opt,name=maxAgeHi"`
	HelloTime         NWFilterField `xml:"hello-time,attr" json:"helloTime" protobuf:"bytes,22,opt,name=helloTime"`
	HelloTimeHi       NWFilterField `xml:"hello-time-hi,attr" json:"helloTimeHi" protobuf:"bytes,23,opt,name=helloTimeHi"`
	ForwardDelay      NWFilterField `xml:"forward-delay,attr" json:"forwardDelay,omitempty" protobuf:"bytes,24,opt,name=forwardDelay"`
	ForwardDelayHi    NWFilterField `xml:"forward-delay-hi,attr" json:"forwardDelayHi,omitempty" protobuf:"bytes,25,opt,name=forwardDelayHi"`
	Comment           string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,26,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleIP struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC  `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	SrcIPAddr              NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty" protobuf:"bytes,3,opt,name=srcIpAddr"`
	SrcIPMask              NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty" protobuf:"bytes,4,opt,name=srcIpMask"`
	DstIPAddr              NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty" protobuf:"bytes,5,opt,name=dstIpAddr"`
	DstIPMask              NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty" protobuf:"bytes,6,opt,name=dstIpMask"`
	Protocol               NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty" protobuf:"bytes,7,opt,name=protocol"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,8,opt,name=nWFilterRuleCommonPort"`
	DSCP                   NWFilterField `xml:"dscp,attr" json:"dscp,omitempty" protobuf:"bytes,9,opt,name=dscp"`
	Comment                string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,10,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleIPv6 struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonMAC  `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonMAC"`
	SrcIPAddr              NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcIpAddr,omitempty" protobuf:"bytes,3,opt,name=srcIpAddr"`
	SrcIPMask              NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcIpMask,omitempty" protobuf:"bytes,4,opt,name=srcIpMask"`
	DstIPAddr              NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstIpAddr,omitempty" protobuf:"bytes,5,opt,name=dstIpAddr"`
	DstIPMask              NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstIpMask,omitempty" protobuf:"bytes,6,opt,name=dstIpMask"`
	Protocol               NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty" protobuf:"bytes,7,opt,name=protocol"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,8,opt,name=nWFilterRuleCommonPort"`
	Type                   NWFilterField `xml:"type,attr" json:"type" protobuf:"bytes,9,opt,name=type"`
	TypeEnd                NWFilterField `xml:"typeend,attr" json:"typeEnd" protobuf:"bytes,10,opt,name=typeEnd"`
	Code                   NWFilterField `xml:"code,attr" json:"code" protobuf:"bytes,11,opt,name=code"`
	CodeEnd                NWFilterField `xml:"codeend,attr" json:"codeEnd" protobuf:"bytes,12,opt,name=codeEnd"`
	Comment                string        `xml:"comment,attr,omitempty" json:"comment" protobuf:"bytes,13,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleTCP struct {
	Match                  string `xml:"match,attr,omitempty" json:"match" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Option                 NWFilterField `xml:"option,attr" json:"option" protobuf:"bytes,4,opt,name=option"`
	Flags                  NWFilterField `xml:"flags,attr,omitempty" json:"flags,omitempty" protobuf:"bytes,5,opt,name=flags"`
	Comment                string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,6,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleUDP struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Comment                string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,4,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleUDPLite struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleESP struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleAH struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleSCTP struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Comment                string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,4,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleICMP struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Type                 NWFilterField `xml:"type,attr" json:"type" protobuf:"bytes,3,opt,name=type"`
	Code                 NWFilterField `xml:"code,attr" json:"code" protobuf:"bytes,4,opt,name=code"`
	Comment              string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,5,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleAll struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleIGMP struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleTCPIPv6 struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Option                 NWFilterField `xml:"option,attr" json:"option" protobuf:"bytes,4,opt,name=option"`
	Comment                string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,5,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleUDPIPv6 struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Comment                string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,4,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleUDPLiteIPv6 struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleESPIPv6 struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleAHIPv6 struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleSCTPIPv6 struct {
	Match                  string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP   `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	NWFilterRuleCommonPort `json:",inline" protobuf:"bytes,3,opt,name=nWFilterRuleCommonPort"`
	Comment                string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,4,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleICMPIPv6 struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Type                 NWFilterField `xml:"type,attr" json:"type" protobuf:"bytes,3,opt,name=type"`
	Code                 NWFilterField `xml:"code,attr" json:"code" protobuf:"bytes,4,opt,name=code"`
	Comment              string        `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,5,opt,name=comment"`
}

// +gogo:genproto=true
type NWFilterRuleAllIPv6 struct {
	Match                string `xml:"match,attr,omitempty" json:"match,omitempty" protobuf:"bytes,1,opt,name=match"`
	NWFilterRuleCommonIP `json:",inline" protobuf:"bytes,2,opt,name=nWFilterRuleCommonIP"`
	Comment              string `xml:"comment,attr,omitempty" json:"comment,omitempty" protobuf:"bytes,3,opt,name=comment"`
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

func (s *NWFilter) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NWFilter) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
