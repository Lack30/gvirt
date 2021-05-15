package spec

import "encoding/xml"

type DomainCaps struct {
	XMLName   xml.Name             `xml:"domainCapabilities" json:"-"`
	Path      string               `xml:"path" json:"path"`
	Domain    string               `xml:"domain" json:"domain"`
	Machine   string               `xml:"machine,omitempty" json:"machine,omitempty"`
	Arch      string               `xml:"arch" json:"arch"`
	VCPU      *DomainCapsVCPU      `xml:"vcpu" json:"vcpu,omitempty"`
	IOThreads *DomainCapsIOThreads `xml:"iothreads" json:"iothreads,omitempty"`
	OS        *DomainCapsOS        `xml:"os" json:"os,omitempty"`
	CPU       *DomainCapsCPU       `xml:"cpu" json:"cpu,omitempty"`
	Devices   *DomainCapsDevices   `xml:"devices" json:"devices,omitempty"`
	Features  *DomainCapsFeatures  `xml:"features" json:"features,omitempty"`
}

type DomainCapsVCPU struct {
	Max uint `xml:"max,attr" json:"max"`
}

type DomainCapsOS struct {
	Supported string              `xml:"supported,attr" json:"supported"`
	Loader    *DomainCapsOSLoader `xml:"loader" json:"loader,omitempty"`
	Enums     []DomainCapsEnum    `xml:"enum" json:"enums"`
}

type DomainCapsOSLoader struct {
	Supported string           `xml:"supported,attr" json:"supported"`
	Values    []string         `xml:"value" json:"values"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums"`
}

type DomainCapsIOThreads struct {
	Supported string `xml:"supported,attr" json:"supported"`
}

type DomainCapsCPU struct {
	Modes []DomainCapsCPUMode `xml:"mode" json:"modes"`
}

type DomainCapsCPUMode struct {
	Name      string                 `xml:"name,attr" json:"name"`
	Supported string                 `xml:"supported,attr" json:"supported"`
	Models    []DomainCapsCPUModel   `xml:"model" json:"models"`
	Vendor    string                 `xml:"vendor,omitempty" json:"vendor"`
	Features  []DomainCapsCPUFeature `xml:"feature" json:"features"`
	Enums     []DomainCapsEnum       `xml:"enum" json:"enums"`
}

type DomainCapsCPUModel struct {
	Name       string `xml:",chardata" json:"name"`
	Usable     string `xml:"usable,attr,omitempty" json:"usable,omitempty"`
	Fallback   string `xml:"fallback,attr,omitempty" json:"fallback,omitempty"`
	Deprecated string `xml:"deprecated,attr,omitempty" json:"deprecated,omitempty"`
}

type DomainCapsCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty" json:"policy,omitempty"`
	Name   string `xml:"name,attr" json:"name"`
}

type DomainCapsEnum struct {
	Name   string   `xml:"name,attr" json:"name"`
	Values []string `xml:"value" json:"values"`
}

type DomainCapsDevices struct {
	Disk     *DomainCapsDevice `xml:"disk" json:"disk,omitempty"`
	Graphics *DomainCapsDevice `xml:"graphics" json:"graphics,omitempty"`
	Video    *DomainCapsDevice `xml:"video" json:"video,omitempty"`
	HostDev  *DomainCapsDevice `xml:"hostdev" json:"hostdev,omitempty"`
	RNG      *DomainCapsDevice `xml:"rng" json:"rng,omitempty"`
}

type DomainCapsDevice struct {
	Supported string           `xml:"supported,attr" json:"supported"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums"`
}

type DomainCapsFeatures struct {
	GIC               *DomainCapsFeatureGIC               `xml:"gic" json:"gic,omitempty"`
	VMCoreInfo        *DomainCapsFeatureVMCoreInfo        `xml:"vmcoreinfo" json:"vmCoreInfo,omitempty"`
	GenID             *DomainCapsFeatureGenID             `xml:"genid" json:"genId,omitempty"`
	BackingStoreInput *DomainCapsFeatureBackingStoreInput `xml:"backingStoreInput" json:"backingStoreInput,omitempty"`
	Backup            *DomainCapsFeatureBackup            `xml:"backup" json:"backup,omitempty"`
	SEV               *DomainCapsFeatureSEV               `xml:"sev" json:"sev,omitempty"`
}

type DomainCapsFeatureGIC struct {
	Supported string           `xml:"supported,attr" json:"supported"`
	Enums     []DomainCapsEnum `xml:"enum" json:"enums"`
}

type DomainCapsFeatureVMCoreInfo struct {
	Supported string `xml:"supported,attr" json:"supported"`
}

type DomainCapsFeatureGenID struct {
	Supported string `xml:"supported,attr" json:"supported"`
}

type DomainCapsFeatureBackingStoreInput struct {
	Supported string `xml:"supported,attr" json:"supported"`
}

type DomainCapsFeatureBackup struct {
	Supported string `xml:"supported,attr" json:"supported"`
}

type DomainCapsFeatureSEV struct {
	Supported       string `xml:"supported,attr" json:"supported"`
	CBitPos         uint   `xml:"cbitpos,omitempty" json:"cBitPos,omitempty"`
	ReducedPhysBits uint   `xml:"reducedPhysBits,omitempty" json:"reducedPhysBits,omitempty"`
}

func (c *DomainCaps) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *DomainCaps) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
