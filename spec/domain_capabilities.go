package spec

import "encoding/xml"

// +gogo:genproto=true
type DomainCaps struct {
	XMLName   xml.Name             `xml:"domainCapabilities" json:"-"`
	Path      string               `xml:"path" json:"path" protobuf:"bytes,1,opt,name=path"`
	Domain    string               `xml:"domain" json:"domain" protobuf:"bytes,2,opt,name=domain"`
	Machine   string               `xml:"machine,omitempty" json:"machine,omitempty" protobuf:"bytes,3,opt,name=machine"`
	Arch      string               `xml:"arch" json:"arch" protobuf:"bytes,4,opt,name=arch"`
	VCPU      *DomainCapsVCPU      `xml:"vcpu" json:"vcpu,omitempty" protobuf:"bytes,5,opt,name=vcpu"`
	IOThreads *DomainCapsIOThreads `xml:"iothreads" json:"iothreads,omitempty" protobuf:"bytes,6,opt,name=iothreads"`
	OS        *DomainCapsOS        `xml:"os" json:"os,omitempty" protobuf:"bytes,7,opt,name=os"`
	CPU       *DomainCapsCPU       `xml:"cpu" json:"cpu,omitempty" protobuf:"bytes,8,opt,name=cpu"`
	Devices   *DomainCapsDevices   `xml:"devices" json:"devices,omitempty" protobuf:"bytes,9,opt,name=devices"`
	Features  *DomainCapsFeatures  `xml:"features" json:"features,omitempty" protobuf:"bytes,10,opt,name=features"`
}

// +gogo:genproto=true
type DomainCapsVCPU struct {
	Max int32 `xml:"max,attr" json:"max" protobuf:"varint,1,opt,name=max"`
}

// +gogo:genproto=true
type DomainCapsOS struct {
	Supported string              `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
	Loader    *DomainCapsOSLoader `xml:"loader" json:"loader,omitempty" protobuf:"bytes,2,opt,name=loader"`
	Enums     []DomainCapsEnum    `xml:"enum" json:"enums" protobuf:"bytes,3,rep,name=enums"`
}

// +gogo:genproto=true
type DomainCapsOSLoader struct {
	Supported string           `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
	Values    []string         `xml:"value" json:"values" protobuf:"bytes,2,rep,name=values"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums" protobuf:"bytes,3,rep,name=enums"`
}

// +gogo:genproto=true
type DomainCapsIOThreads struct {
	Supported string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
}

// +gogo:genproto=true
type DomainCapsCPU struct {
	Modes []DomainCapsCPUMode `xml:"mode" json:"modes" protobuf:"bytes,1,rep,name=modes"`
}

// +gogo:genproto=true
type DomainCapsCPUMode struct {
	Name      string                 `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Supported string                 `xml:"supported,attr" json:"supported" protobuf:"bytes,2,opt,name=supported"`
	Models    []DomainCapsCPUModel   `xml:"model" json:"models" protobuf:"bytes,3,rep,name=models"`
	Vendor    string                 `xml:"vendor,omitempty" json:"vendor" protobuf:"bytes,4,opt,name=vendor"`
	Features  []DomainCapsCPUFeature `xml:"feature" json:"features" protobuf:"bytes,5,rep,name=features"`
	Enums     []DomainCapsEnum       `xml:"enum" json:"enums" protobuf:"bytes,6,rep,name=enums"`
}

// +gogo:genproto=true
type DomainCapsCPUModel struct {
	Name       string `xml:",chardata" json:"name" protobuf:"bytes,1,opt,name=name"`
	Usable     string `xml:"usable,attr,omitempty" json:"usable,omitempty" protobuf:"bytes,2,opt,name=usable"`
	Fallback   string `xml:"fallback,attr,omitempty" json:"fallback,omitempty" protobuf:"bytes,3,opt,name=fallback"`
	Deprecated string `xml:"deprecated,attr,omitempty" json:"deprecated,omitempty" protobuf:"bytes,4,opt,name=deprecated"`
}

// +gogo:genproto=true
type DomainCapsCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty" json:"policy,omitempty" protobuf:"bytes,1,opt,name=policy"`
	Name   string `xml:"name,attr" json:"name" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type DomainCapsEnum struct {
	Name   string   `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Values []string `xml:"value" json:"values" protobuf:"bytes,2,rep,name=values"`
}

// +gogo:genproto=true
type DomainCapsDevices struct {
	Disk     *DomainCapsDevice `xml:"disk" json:"disk,omitempty" protobuf:"bytes,1,opt,name=disk"`
	Graphics *DomainCapsDevice `xml:"graphics" json:"graphics,omitempty" protobuf:"bytes,2,opt,name=graphics"`
	Video    *DomainCapsDevice `xml:"video" json:"video,omitempty" protobuf:"bytes,3,opt,name=video"`
	HostDev  *DomainCapsDevice `xml:"hostdev" json:"hostdev,omitempty" protobuf:"bytes,4,opt,name=hostdev"`
	RNG      *DomainCapsDevice `xml:"rng" json:"rng,omitempty" protobuf:"bytes,5,opt,name=rng"`
}

// +gogo:genproto=true
type DomainCapsDevice struct {
	Supported string           `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums" protobuf:"bytes,2,rep,name=enums"`
}

// +gogo:genproto=true
type DomainCapsFeatures struct {
	GIC               *DomainCapsFeatureGIC               `xml:"gic" json:"gic,omitempty" protobuf:"bytes,1,opt,name=gic"`
	VMCoreInfo        *DomainCapsFeatureVMCoreInfo        `xml:"vmcoreinfo" json:"vmCoreInfo,omitempty" protobuf:"bytes,2,opt,name=vmCoreInfo"`
	GenID             *DomainCapsFeatureGenID             `xml:"genid" json:"genId,omitempty" protobuf:"bytes,3,opt,name=genId"`
	BackingStoreInput *DomainCapsFeatureBackingStoreInput `xml:"backingStoreInput" json:"backingStoreInput,omitempty" protobuf:"bytes,4,opt,name=backingStoreInput"`
	Backup            *DomainCapsFeatureBackup            `xml:"backup" json:"backup,omitempty" protobuf:"bytes,5,opt,name=backup"`
	SEV               *DomainCapsFeatureSEV               `xml:"sev" json:"sev,omitempty" protobuf:"bytes,6,opt,name=sev"`
}

// +gogo:genproto=true
type DomainCapsFeatureGIC struct {
	Supported string           `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums" protobuf:"bytes,2,rep,name=enums"`
}

// +gogo:genproto=true
type DomainCapsFeatureVMCoreInfo struct {
	Supported string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
}

// +gogo:genproto=true
type DomainCapsFeatureGenID struct {
	Supported string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
}

// +gogo:genproto=true
type DomainCapsFeatureBackingStoreInput struct {
	Supported string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
}

// +gogo:genproto=true
type DomainCapsFeatureBackup struct {
	Supported string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
}

// +gogo:genproto=true
type DomainCapsFeatureSEV struct {
	Supported       string `xml:"supported,attr" json:"supported" protobuf:"bytes,1,opt,name=supported"`
	CBitPos         int32  `xml:"cbitpos,omitempty" json:"cBitPos,omitempty" protobuf:"varint,2,opt,name=cBitPos"`
	ReducedPhysBits int32  `xml:"reducedPhysBits,omitempty" json:"reducedPhysBits,omitempty" protobuf:"varint,3,opt,name=reducedPhysBits"`
}

func (c *DomainCaps) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *DomainCaps) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
