package spec

import "encoding/xml"

type CapsHostCPUTopology struct {
	Sockets int32 `xml:"sockets,attr" json:"sockets"`
	Cores   int32 `xml:"cores,attr" json:"cores"`
	Threads int32 `xml:"threads,attr" json:"threads"`
}

type CapsHostCPUFeatureFlag struct {
	Name string `xml:"name,attr" json:"name"`
}

type CapsHostCPUPageSize struct {
	Size int32  `xml:"size,attr" json:"size"`
	Unit string `xml:"unit,attr" json:"unit"`
}

type CapsHostCPUMicrocode struct {
	Version int32 `xml:"version,attr" json:"version"`
}

type CapsHostCPU struct {
	XMLName      xml.Name                 `xml:"cpu" json:"-"`
	Arch         string                   `xml:"arch,omitempty" json:"arch,omitempty"`
	Model        string                   `xml:"model,omitempty" json:"model,omitempty"`
	Vendor       string                   `xml:"vendor,omitempty" json:"vendor,omitempty"`
	Topology     *CapsHostCPUTopology     `xml:"topology" json:"topology"`
	FeatureFlags []CapsHostCPUFeatureFlag `xml:"feature" json:"feature"`
	Features     *CapsHostCPUFeatures     `xml:"features" json:"features"`
	PageSizes    []CapsHostCPUPageSize    `xml:"pages" json:"page"`
	Microcode    *CapsHostCPUMicrocode    `xml:"microcode" json:"microcode"`
}

type CapsHostCPUFeature struct {
}

type CapsHostCPUFeatures struct {
	PAE    *CapsHostCPUFeature `xml:"pae" json:"pae"`
	NonPAE *CapsHostCPUFeature `xml:"nonpae" json:"nonPae"`
	SVM    *CapsHostCPUFeature `xml:"svm" json:"svm"`
	VMX    *CapsHostCPUFeature `xml:"vmx" json:"vmx"`
}

type CapsHostNUMAMemory struct {
	Size int64  `xml:",chardata" json:"size"`
	Unit string `xml:"unit,attr" json:"unit"`
}

type CapsHostNUMAPageInfo struct {
	Size  int32  `xml:"size,attr" json:"size"`
	Unit  string `xml:"unit,attr" json:"unit"`
	Count int64  `xml:",chardata" json:"count"`
}

type CapsHostNUMACPU struct {
	ID       int32  `xml:"id,attr" json:"id"`
	SocketID *int32 `xml:"socket_id,attr" json:"socketId"`
	DieID    *int32 `xml:"die_id,attr" json:"dieId"`
	CoreID   *int32 `xml:"core_id,attr" json:"coreId"`
	Siblings string `xml:"siblings,attr,omitempty" json:"siblings"`
}

type CapsHostNUMASibling struct {
	ID    int32 `xml:"id,attr" json:"id"`
	Value int32 `xml:"value,attr" json:"value"`
}

type CapsHostNUMACell struct {
	ID        int32                  `xml:"id,attr" json:"id"`
	Memory    *CapsHostNUMAMemory    `xml:"memory" json:"memory"`
	PageInfo  []CapsHostNUMAPageInfo `xml:"pages" json:"pageInfo"`
	Distances *CapsHostNUMADistances `xml:"distances" json:"distances"`
	CPUS      *CapsHostNUMACPUs      `xml:"cpus" json:"cpus"`
}

type CapsHostNUMADistances struct {
	Siblings []CapsHostNUMASibling `xml:"sibling" json:"sibling"`
}

type CapsHostNUMACPUs struct {
	Num  int32             `xml:"num,attr,omitempty" json:"num,omitempty"`
	CPUs []CapsHostNUMACPU `xml:"cpu" json:"cpus"`
}

type CapsHostNUMATopology struct {
	Cells *CapsHostNUMACells `xml:"cells" json:"cells"`
}

type CapsHostNUMACells struct {
	Num   int32              `xml:"num,attr,omitempty" json:"num,omitempty"`
	Cells []CapsHostNUMACell `xml:"cell" json:"cells"`
}

type CapsHostSecModelLabel struct {
	Type  string `xml:"type,attr" json:"type"`
	Value string `xml:",chardata" json:"value"`
}

type CapsHostSecModel struct {
	Name   string                  `xml:"model" json:"name"`
	DOI    string                  `xml:"doi" json:"doi"`
	Labels []CapsHostSecModelLabel `xml:"baselabel" json:"labels"`
}

type CapsHostMigrationFeatures struct {
	Live          *CapsHostMigrationLive          `xml:"live" json:"live"`
	URITransports *CapsHostMigrationURITransports `xml:"uri_transports" json:"uriTransports"`
}

type CapsHostMigrationLive struct {
}

type CapsHostMigrationURITransports struct {
	URI []string `xml:"uri_transport" json:"uri"`
}

type CapsHost struct {
	UUID              string                     `xml:"uuid,omitempty" json:"uuid,omitempty"`
	CPU               *CapsHostCPU               `xml:"cpu" json:"cpu"`
	PowerManagement   *CapsHostPowerManagement   `xml:"power_management" json:"powerManagement"`
	IOMMU             *CapsHostIOMMU             `xml:"iommu" json:"iommu"`
	MigrationFeatures *CapsHostMigrationFeatures `xml:"migration_features" json:"migrationFeatures"`
	NUMA              *CapsHostNUMATopology      `xml:"topology" json:"numa"`
	Cache             *CapsHostCache             `xml:"cache" json:"cache"`
	MemoryBandwidth   *CapsHostMemoryBandwidth   `xml:"memory_bandwidth" json:"memoryBandwidth"`
	SecModel          []CapsHostSecModel         `xml:"secmodel" json:"secModel"`
}

type CapsHostPowerManagement struct {
	SuspendMem    *CapsHostPowerManagementMode `xml:"suspend_mem" json:"suspendMem"`
	SuspendDisk   *CapsHostPowerManagementMode `xml:"suspend_disk" json:"suspendDisk"`
	SuspendHybrid *CapsHostPowerManagementMode `xml:"suspend_hybrid" json:"suspendHybrid"`
}

type CapsHostPowerManagementMode struct {
}

type CapsHostIOMMU struct {
	Support string `xml:"support,attr" json:"support"`
}

type CapsHostCache struct {
	Banks   []CapsHostCacheBank   `xml:"bank" json:"banks"`
	Monitor *CapsHostCacheMonitor `xml:"monitor" json:"monitor"`
}

type CapsHostCacheBank struct {
	ID      int32                  `xml:"id,attr" json:"id"`
	Level   int32                  `xml:"level,attr" json:"level"`
	Type    string                 `xml:"type,attr" json:"type"`
	Size    int32                  `xml:"size,attr" json:"size"`
	Unit    string                 `xml:"unit,attr" json:"unit"`
	CPUs    string                 `xml:"cpus,attr" json:"cpus"`
	Control []CapsHostCacheControl `xml:"control" json:"control"`
}

type CapsHostCacheMonitor struct {
	Level          int32                         `xml:"level,attr,omitempty" json:"level,omitempty"`
	ResueThreshold int32                         `xml:"reuseThreshold,attr,omitempty" json:"resueThreshold,omitempty"`
	MaxMonitors    int32                         `xml:"maxMonitors,attr" json:"maxMonitors"`
	Features       []CapsHostCacheMonitorFeature `xml:"feature" json:"features"`
}

type CapsHostCacheMonitorFeature struct {
	Name string `xml:"name,attr" json:"name"`
}

type CapsHostCacheControl struct {
	Granularity int32  `xml:"granularity,attr" json:"granularity"`
	Min         int32  `xml:"min,attr,omitempty" json:"min,omitempty"`
	Unit        string `xml:"unit,attr" json:"unit"`
	Type        string `xml:"type,attr" json:"type"`
	MaxAllows   int32  `xml:"maxAllocs,attr" json:"maxAllows"`
}

type CapsHostMemoryBandwidth struct {
	Nodes   []CapsHostMemoryBandwidthNode   `xml:"node" json:"nodes"`
	Monitor *CapsHostMemoryBandwidthMonitor `xml:"monitor" json:"monitor"`
}

type CapsHostMemoryBandwidthNode struct {
	ID      int32                               `xml:"id,attr" json:"id"`
	CPUs    string                              `xml:"cpus,attr" json:"cpus"`
	Control *CapsHostMemoryBandwidthNodeControl `xml:"control" json:"control"`
}

type CapsHostMemoryBandwidthNodeControl struct {
	Granularity int32 `xml:"granularity,attr" json:"granularity"`
	Min         int32 `xml:"min,attr" json:"min"`
	MaxAllocs   int32 `xml:"maxAllocs,attr" json:"maxAllocs"`
}

type CapsHostMemoryBandwidthMonitor struct {
	MaxMonitors int32                                   `xml:"maxMonitors,attr" json:"max_monitors"`
	Features    []CapsHostMemoryBandwidthMonitorFeature `xml:"feature" json:"features"`
}

type CapsHostMemoryBandwidthMonitorFeature struct {
	Name string `xml:"name,attr" json:"name"`
}

type CapsGuestMachine struct {
	Name      string `xml:",chardata" json:"name"`
	MaxCPUs   int32  `xml:"maxCpus,attr,omitempty" json:"maxCpus,omitempty"`
	Canonical string `xml:"canonical,attr,omitempty" json:"canonical,omitempty"`
}

type CapsGuestDomain struct {
	Type     string             `xml:"type,attr" json:"type"`
	Emulator string             `xml:"emulator,omitempty" json:"emulator,omitempty"`
	Machines []CapsGuestMachine `xml:"machine" json:"machines"`
}

type CapsGuestArch struct {
	Name     string             `xml:"name,attr" json:"name"`
	WordSize string             `xml:"wordsize" json:"wordSize"`
	Emulator string             `xml:"emulator" json:"emulator"`
	Loader   string             `xml:"loader,omitempty" json:"loader,omitempty"`
	Machines []CapsGuestMachine `xml:"machine" json:"machines"`
	Domains  []CapsGuestDomain  `xml:"domain" json:"domains"`
}

type CapsGuestFeatureCPUSelection struct {
}

type CapsGuestFeatureDeviceBoot struct {
}

type CapsGuestFeaturePAE struct {
}

type CapsGuestFeatureNonPAE struct {
}

type CapsGuestFeatureDiskSnapshot struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty"`
}

type CapsGuestFeatureAPIC struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty"`
}

type CapsGuestFeatureACPI struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty"`
}

type CapsGuestFeatureIA64BE struct {
}

type CapsGuestFeatures struct {
	CPUSelection *CapsGuestFeatureCPUSelection `xml:"cpuselection" json:"cpuSelection"`
	DeviceBoot   *CapsGuestFeatureDeviceBoot   `xml:"deviceboot" json:"deviceBoot"`
	DiskSnapshot *CapsGuestFeatureDiskSnapshot `xml:"disksnapshot" json:"diskSnapshot"`
	PAE          *CapsGuestFeaturePAE          `xml:"pae" json:"pae"`
	NonPAE       *CapsGuestFeatureNonPAE       `xml:"nonpae" json:"nonPae"`
	APIC         *CapsGuestFeatureAPIC         `xml:"apic" json:"apic"`
	ACPI         *CapsGuestFeatureACPI         `xml:"acpi" json:"acpi"`
	IA64BE       *CapsGuestFeatureIA64BE       `xml:"ia64_be" json:"ia64_be"`
}

type CapsGuest struct {
	OSType   string             `xml:"os_type" json:"osType"`
	Arch     CapsGuestArch      `xml:"arch" json:"arch"`
	Features *CapsGuestFeatures `xml:"features" json:"features"`
}

type Caps struct {
	XMLName xml.Name    `xml:"capabilities" json:"-"`
	Host    CapsHost    `xml:"host" json:"host"`
	Guests  []CapsGuest `xml:"guest" json:"guests"`
}

func (c *CapsHostCPU) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *CapsHostCPU) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (c *Caps) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *Caps) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
