package spec

import "encoding/xml"

// +gogo:genproto=true
type CapsHostCPUTopology struct {
	Sockets int32 `xml:"sockets,attr" json:"sockets" protobuf:"varint,1,opt,name=sockets"`
	Cores   int32 `xml:"cores,attr" json:"cores" protobuf:"varint,2,opt,name=cores"`
	Threads int32 `xml:"threads,attr" json:"threads" protobuf:"varint,3,opt,name=threads"`
}

// +gogo:genproto=true
type CapsHostCPUFeatureFlag struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type CapsHostCPUPageSize struct {
	Size_ int32  `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type CapsHostCPUMicrocode struct {
	Version int32 `xml:"version,attr" json:"version" protobuf:"varint,1,opt,name=version"`
}

// +gogo:genproto=true
type CapsHostCPU struct {
	XMLName      xml.Name                 `xml:"cpu" json:"-"`
	Arch         string                   `xml:"arch,omitempty" json:"arch,omitempty" protobuf:"bytes,1,opt,name=arch"`
	Model        string                   `xml:"model,omitempty" json:"model,omitempty" protobuf:"bytes,2,opt,name=model"`
	Vendor       string                   `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,3,opt,name=vendor"`
	Topology     *CapsHostCPUTopology     `xml:"topology" json:"topology" protobuf:"bytes,4,opt,name=topology"`
	FeatureFlags []CapsHostCPUFeatureFlag `xml:"feature" json:"feature" protobuf:"bytes,5,rep,name=feature"`
	Features     *CapsHostCPUFeatures     `xml:"features" json:"features" protobuf:"bytes,6,opt,name=features"`
	PageSizes    []CapsHostCPUPageSize    `xml:"pages" json:"page" protobuf:"bytes,7,rep,name=page"`
	Microcode    *CapsHostCPUMicrocode    `xml:"microcode" json:"microcode" protobuf:"bytes,8,opt,name=microcode"`
}

// +gogo:genproto=true
type CapsHostCPUFeature struct {
}

// +gogo:genproto=true
type CapsHostCPUFeatures struct {
	PAE    *CapsHostCPUFeature `xml:"pae" json:"pae" protobuf:"bytes,1,opt,name=pae"`
	NonPAE *CapsHostCPUFeature `xml:"nonpae" json:"nonPae" protobuf:"bytes,2,opt,name=nonPae"`
	SVM    *CapsHostCPUFeature `xml:"svm" json:"svm" protobuf:"bytes,3,opt,name=svm"`
	VMX    *CapsHostCPUFeature `xml:"vmx" json:"vmx" protobuf:"bytes,4,opt,name=vmx"`
}

// +gogo:genproto=true
type CapsHostNUMAMemory struct {
	Size_ int64  `xml:",chardata" json:"size" protobuf:"varint,1,opt,name=size"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,2,opt,name=unit"`
}

// +gogo:genproto=true
type CapsHostNUMAPageInfo struct {
	Size_ int32  `xml:"size,attr" json:"size" protobuf:"varint,1,opt,name=size"`
	Unit  string `xml:"unit,attr" json:"unit" protobuf:"bytes,2,opt,name=unit"`
	Count int64  `xml:",chardata" json:"count" protobuf:"varint,3,opt,name=count"`
}

// +gogo:genproto=true
type CapsHostNUMACPU struct {
	ID       int32  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	SocketID *int32 `xml:"socket_id,attr" json:"socketId" protobuf:"varint,2,opt,name=socketId"`
	DieID    *int32 `xml:"die_id,attr" json:"dieId" protobuf:"varint,3,opt,name=dieId"`
	CoreID   *int32 `xml:"core_id,attr" json:"coreId" protobuf:"varint,4,opt,name=coreId"`
	Siblings string `xml:"siblings,attr,omitempty" json:"siblings" protobuf:"bytes,5,opt,name=siblings"`
}

// +gogo:genproto=true
type CapsHostNUMASibling struct {
	ID    int32 `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Value int32 `xml:"value,attr" json:"value" protobuf:"varint,2,opt,name=value"`
}

// +gogo:genproto=true
type CapsHostNUMACell struct {
	ID        int32                  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Memory    *CapsHostNUMAMemory    `xml:"memory" json:"memory" protobuf:"bytes,2,opt,name=memory"`
	PageInfo  []CapsHostNUMAPageInfo `xml:"pages" json:"pageInfo" protobuf:"bytes,3,rep,name=pageInfo"`
	Distances *CapsHostNUMADistances `xml:"distances" json:"distances" protobuf:"bytes,4,opt,name=distances"`
	CPUS      *CapsHostNUMACPUs      `xml:"cpus" json:"cpus" protobuf:"bytes,5,opt,name=cpus"`
}

// +gogo:genproto=true
type CapsHostNUMADistances struct {
	Siblings []CapsHostNUMASibling `xml:"sibling" json:"sibling" protobuf:"bytes,1,rep,name=sibling"`
}

// +gogo:genproto=true
type CapsHostNUMACPUs struct {
	Num  int32             `xml:"num,attr,omitempty" json:"num,omitempty" protobuf:"varint,1,opt,name=num"`
	CPUs []CapsHostNUMACPU `xml:"cpu" json:"cpus" protobuf:"bytes,2,rep,name=cpus"`
}

// +gogo:genproto=true
type CapsHostNUMATopology struct {
	Cells *CapsHostNUMACells `xml:"cells" json:"cells" protobuf:"bytes,1,opt,name=cells"`
}

// +gogo:genproto=true
type CapsHostNUMACells struct {
	Num   int32              `xml:"num,attr,omitempty" json:"num,omitempty" protobuf:"varint,1,opt,name=num"`
	Cells []CapsHostNUMACell `xml:"cell" json:"cells" protobuf:"bytes,2,rep,name=cells"`
}

// +gogo:genproto=true
type CapsHostSecModelLabel struct {
	Type  string `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Value string `xml:",chardata" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type CapsHostSecModel struct {
	Name   string                  `xml:"model" json:"name" protobuf:"bytes,1,opt,name=name"`
	DOI    string                  `xml:"doi" json:"doi" protobuf:"bytes,2,opt,name=doi"`
	Labels []CapsHostSecModelLabel `xml:"baselabel" json:"labels" protobuf:"bytes,3,rep,name=labels"`
}

// +gogo:genproto=true
type CapsHostMigrationFeatures struct {
	Live          *CapsHostMigrationLive          `xml:"live" json:"live" protobuf:"bytes,1,opt,name=live"`
	URITransports *CapsHostMigrationURITransports `xml:"uri_transports" json:"uriTransports" protobuf:"bytes,2,opt,name=uriTransports"`
}

// +gogo:genproto=true
type CapsHostMigrationLive struct {
}

// +gogo:genproto=true
type CapsHostMigrationURITransports struct {
	URI []string `xml:"uri_transport" json:"uri" protobuf:"bytes,1,rep,name=uri"`
}

// +gogo:genproto=true
type CapsHost struct {
	UUID              string                     `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,1,opt,name=uuid"`
	CPU               *CapsHostCPU               `xml:"cpu" json:"cpu" protobuf:"bytes,2,opt,name=cpu"`
	PowerManagement   *CapsHostPowerManagement   `xml:"power_management" json:"powerManagement" protobuf:"bytes,3,opt,name=powerManagement"`
	IOMMU             *CapsHostIOMMU             `xml:"iommu" json:"iommu" protobuf:"bytes,4,opt,name=iommu"`
	MigrationFeatures *CapsHostMigrationFeatures `xml:"migration_features" json:"migrationFeatures" protobuf:"bytes,5,opt,name=migrationFeatures"`
	NUMA              *CapsHostNUMATopology      `xml:"topology" json:"numa" protobuf:"bytes,6,opt,name=numa"`
	Cache             *CapsHostCache             `xml:"cache" json:"cache" protobuf:"bytes,7,opt,name=cache"`
	MemoryBandwidth   *CapsHostMemoryBandwidth   `xml:"memory_bandwidth" json:"memoryBandwidth" protobuf:"bytes,8,opt,name=memoryBandwidth"`
	SecModel          []CapsHostSecModel         `xml:"secmodel" json:"secModel" protobuf:"bytes,9,rep,name=secModel"`
}

// +gogo:genproto=true
type CapsHostPowerManagement struct {
	SuspendMem    *CapsHostPowerManagementMode `xml:"suspend_mem" json:"suspendMem" protobuf:"bytes,1,opt,name=suspendMem"`
	SuspendDisk   *CapsHostPowerManagementMode `xml:"suspend_disk" json:"suspendDisk" protobuf:"bytes,2,opt,name=suspendDisk"`
	SuspendHybrid *CapsHostPowerManagementMode `xml:"suspend_hybrid" json:"suspendHybrid" protobuf:"bytes,3,opt,name=suspendHybrid"`
}

// +gogo:genproto=true
type CapsHostPowerManagementMode struct {
}

// +gogo:genproto=true
type CapsHostIOMMU struct {
	Support string `xml:"support,attr" json:"support" protobuf:"bytes,1,opt,name=support"`
}

// +gogo:genproto=true
type CapsHostCache struct {
	Banks   []CapsHostCacheBank   `xml:"bank" json:"banks" protobuf:"bytes,1,rep,name=banks"`
	Monitor *CapsHostCacheMonitor `xml:"monitor" json:"monitor" protobuf:"bytes,2,opt,name=monitor"`
}

// +gogo:genproto=true
type CapsHostCacheBank struct {
	ID      int32                  `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	Level   int32                  `xml:"level,attr" json:"level" protobuf:"varint,2,opt,name=level"`
	Type    string                 `xml:"type,attr" json:"type" protobuf:"bytes,3,opt,name=type"`
	Size_   int32                  `xml:"size,attr" json:"size" protobuf:"varint,4,opt,name=size"`
	Unit    string                 `xml:"unit,attr" json:"unit" protobuf:"bytes,5,opt,name=unit"`
	CPUs    string                 `xml:"cpus,attr" json:"cpus" protobuf:"bytes,6,opt,name=cpus"`
	Control []CapsHostCacheControl `xml:"control" json:"control" protobuf:"bytes,7,rep,name=control"`
}

// +gogo:genproto=true
type CapsHostCacheMonitor struct {
	Level          int32                         `xml:"level,attr,omitempty" json:"level,omitempty" protobuf:"varint,1,opt,name=level"`
	ResueThreshold int32                         `xml:"reuseThreshold,attr,omitempty" json:"resueThreshold,omitempty" protobuf:"varint,2,opt,name=resueThreshold"`
	MaxMonitors    int32                         `xml:"maxMonitors,attr" json:"maxMonitors" protobuf:"varint,3,opt,name=maxMonitors"`
	Features       []CapsHostCacheMonitorFeature `xml:"feature" json:"features" protobuf:"bytes,4,rep,name=features"`
}

// +gogo:genproto=true
type CapsHostCacheMonitorFeature struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type CapsHostCacheControl struct {
	Granularity int32  `xml:"granularity,attr" json:"granularity" protobuf:"varint,1,opt,name=granularity"`
	Min         int32  `xml:"min,attr,omitempty" json:"min,omitempty" protobuf:"varint,2,opt,name=min"`
	Unit        string `xml:"unit,attr" json:"unit" protobuf:"bytes,3,opt,name=unit"`
	Type        string `xml:"type,attr" json:"type" protobuf:"bytes,4,opt,name=type"`
	MaxAllows   int32  `xml:"maxAllocs,attr" json:"maxAllows" protobuf:"varint,5,opt,name=maxAllows"`
}

// +gogo:genproto=true
type CapsHostMemoryBandwidth struct {
	Nodes   []CapsHostMemoryBandwidthNode   `xml:"node" json:"nodes" protobuf:"bytes,1,rep,name=nodes"`
	Monitor *CapsHostMemoryBandwidthMonitor `xml:"monitor" json:"monitor" protobuf:"bytes,2,opt,name=monitor"`
}

// +gogo:genproto=true
type CapsHostMemoryBandwidthNode struct {
	ID      int32                               `xml:"id,attr" json:"id" protobuf:"varint,1,opt,name=id"`
	CPUs    string                              `xml:"cpus,attr" json:"cpus" protobuf:"bytes,2,opt,name=cpus"`
	Control *CapsHostMemoryBandwidthNodeControl `xml:"control" json:"control" protobuf:"bytes,3,opt,name=control"`
}

// +gogo:genproto=true
type CapsHostMemoryBandwidthNodeControl struct {
	Granularity int32 `xml:"granularity,attr" json:"granularity" protobuf:"varint,1,opt,name=granularity"`
	Min         int32 `xml:"min,attr" json:"min" protobuf:"varint,2,opt,name=min"`
	MaxAllocs   int32 `xml:"maxAllocs,attr" json:"maxAllocs" protobuf:"varint,3,opt,name=maxAllocs"`
}

// +gogo:genproto=true
type CapsHostMemoryBandwidthMonitor struct {
	MaxMonitors int32                                   `xml:"maxMonitors,attr" json:"max_monitors" protobuf:"varint,1,opt,name=max_monitors,json=maxMonitors"`
	Features    []CapsHostMemoryBandwidthMonitorFeature `xml:"feature" json:"features" protobuf:"bytes,2,rep,name=features"`
}

// +gogo:genproto=true
type CapsHostMemoryBandwidthMonitorFeature struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type CapsGuestMachine struct {
	Name      string `xml:",chardata" json:"name" protobuf:"bytes,1,opt,name=name"`
	MaxCPUs   int32  `xml:"maxCpus,attr,omitempty" json:"maxCpus,omitempty" protobuf:"varint,2,opt,name=maxCpus"`
	Canonical string `xml:"canonical,attr,omitempty" json:"canonical,omitempty" protobuf:"bytes,3,opt,name=canonical"`
}

// +gogo:genproto=true
type CapsGuestDomain struct {
	Type     string             `xml:"type,attr" json:"type" protobuf:"bytes,1,opt,name=type"`
	Emulator string             `xml:"emulator,omitempty" json:"emulator,omitempty" protobuf:"bytes,2,opt,name=emulator"`
	Machines []CapsGuestMachine `xml:"machine" json:"machines" protobuf:"bytes,3,rep,name=machines"`
}

// +gogo:genproto=true
type CapsGuestArch struct {
	Name     string             `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	WordSize string             `xml:"wordsize" json:"wordSize" protobuf:"bytes,2,opt,name=wordSize"`
	Emulator string             `xml:"emulator" json:"emulator" protobuf:"bytes,3,opt,name=emulator"`
	Loader   string             `xml:"loader,omitempty" json:"loader,omitempty" protobuf:"bytes,4,opt,name=loader"`
	Machines []CapsGuestMachine `xml:"machine" json:"machines" protobuf:"bytes,5,rep,name=machines"`
	Domains  []CapsGuestDomain  `xml:"domain" json:"domains" protobuf:"bytes,6,rep,name=domains"`
}

// +gogo:genproto=true
type CapsGuestFeatureCPUSelection struct {
}

// +gogo:genproto=true
type CapsGuestFeatureDeviceBoot struct {
}

// +gogo:genproto=true
type CapsGuestFeaturePAE struct {
}

// +gogo:genproto=true
type CapsGuestFeatureNonPAE struct {
}

// +gogo:genproto=true
type CapsGuestFeatureDiskSnapshot struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" protobuf:"bytes,1,opt,name=default"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" protobuf:"bytes,2,opt,name=toggle"`
}

// +gogo:genproto=true
type CapsGuestFeatureAPIC struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" protobuf:"bytes,1,opt,name=default"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" protobuf:"bytes,2,opt,name=toggle"`
}

// +gogo:genproto=true
type CapsGuestFeatureACPI struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" protobuf:"bytes,1,opt,name=default"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" protobuf:"bytes,2,opt,name=toggle"`
}

// +gogo:genproto=true
type CapsGuestFeatureIA64BE struct {
}

// +gogo:genproto=true
type CapsGuestFeatures struct {
	CPUSelection *CapsGuestFeatureCPUSelection `xml:"cpuselection" json:"cpuSelection" protobuf:"bytes,1,opt,name=cpuSelection"`
	DeviceBoot   *CapsGuestFeatureDeviceBoot   `xml:"deviceboot" json:"deviceBoot" protobuf:"bytes,2,opt,name=deviceBoot"`
	DiskSnapshot *CapsGuestFeatureDiskSnapshot `xml:"disksnapshot" json:"diskSnapshot" protobuf:"bytes,3,opt,name=diskSnapshot"`
	PAE          *CapsGuestFeaturePAE          `xml:"pae" json:"pae" protobuf:"bytes,4,opt,name=pae"`
	NonPAE       *CapsGuestFeatureNonPAE       `xml:"nonpae" json:"nonPae" protobuf:"bytes,5,opt,name=nonPae"`
	APIC         *CapsGuestFeatureAPIC         `xml:"apic" json:"apic" protobuf:"bytes,6,opt,name=apic"`
	ACPI         *CapsGuestFeatureACPI         `xml:"acpi" json:"acpi" protobuf:"bytes,7,opt,name=acpi"`
	IA64BE       *CapsGuestFeatureIA64BE       `xml:"ia64_be" json:"ia64_be" protobuf:"bytes,8,opt,name=ia64_be,json=ia64Be"`
}

// +gogo:genproto=true
type CapsGuest struct {
	OSType   string             `xml:"os_type" json:"osType" protobuf:"bytes,1,opt,name=osType"`
	Arch     CapsGuestArch      `xml:"arch" json:"arch" protobuf:"bytes,2,opt,name=arch"`
	Features *CapsGuestFeatures `xml:"features" json:"features" protobuf:"bytes,3,opt,name=features"`
}

// +gogo:genproto=true
type Caps struct {
	XMLName xml.Name    `xml:"capabilities" json:"-"`
	Host    CapsHost    `xml:"host" json:"host" protobuf:"bytes,1,opt,name=host"`
	Guests  []CapsGuest `xml:"guest" json:"guests" protobuf:"bytes,2,rep,name=guests"`
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
