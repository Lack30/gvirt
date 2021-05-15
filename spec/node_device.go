package spec

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// +gogo:genproto=true
type NodeDevice struct {
	XMLName    xml.Name             `xml:"device" json:"-"`
	Name       string               `xml:"name" json:"name" protobuf:"bytes,1,opt,name=name"`
	Path       string               `xml:"path,omitempty" json:"path,omitempty" protobuf:"bytes,2,opt,name=path"`
	DevNodes   []NodeDeviceDevNode  `xml:"devnode" json:"devnodes" protobuf:"bytes,3,rep,name=devnodes"`
	Parent     string               `xml:"parent,omitempty" json:"parent,omitempty" protobuf:"bytes,4,opt,name=parent"`
	Driver     *NodeDeviceDriver    `xml:"driver" json:"driver,omitempty" protobuf:"bytes,5,opt,name=driver"`
	Capability NodeDeviceCapability `xml:"capability" json:"capability" protobuf:"bytes,6,opt,name=capability"`
}

// +gogo:genproto=true
type NodeDeviceDevNode struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Path string `xml:",chardata" json:"path" protobuf:"bytes,2,opt,name=path"`
}

// +gogo:genproto=true
type NodeDeviceDriver struct {
	Name string `xml:"name" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type NodeDeviceCapability struct {
	System     *NodeDeviceSystemCapability     `json:"system,omitempty" protobuf:"bytes,1,opt,name=system"`
	PCI        *NodeDevicePCICapability        `json:"pci,omitempty" protobuf:"bytes,2,opt,name=pci"`
	USB        *NodeDeviceUSBCapability        `json:"usb,omitempty" protobuf:"bytes,3,opt,name=usb"`
	USBDevice  *NodeDeviceUSBDeviceCapability  `json:"usbDevice,omitempty" protobuf:"bytes,4,opt,name=usbDevice"`
	Net        *NodeDeviceNetCapability        `json:"net,omitempty" protobuf:"bytes,5,opt,name=net"`
	SCSIHost   *NodeDeviceSCSIHostCapability   `json:"scsiHost,omitempty" protobuf:"bytes,6,opt,name=scsiHost"`
	SCSITarget *NodeDeviceSCSITargetCapability `json:"scsiTarget,omitempty" protobuf:"bytes,7,opt,name=scsiTarget"`
	SCSI       *NodeDeviceSCSICapability       `json:"scsi,omitempty" protobuf:"bytes,8,opt,name=scsi"`
	Storage    *NodeDeviceStorageCapability    `json:"storage,omitempty" protobuf:"bytes,9,opt,name=storage"`
	DRM        *NodeDeviceDRMCapability        `json:"drm,omitempty" protobuf:"bytes,10,opt,name=drm"`
	CCW        *NodeDeviceCCWCapability        `json:"ccw,omitempty" protobuf:"bytes,11,opt,name=ccw"`
	MDev       *NodeDeviceMDevCapability       `json:"mDev,omitempty" protobuf:"bytes,12,opt,name=mDev"`
	CSS        *NodeDeviceCSSCapability        `json:"css,omitempty" protobuf:"bytes,13,opt,name=css"`
	APQueue    *NodeDeviceAPQueueCapability    `json:"apQueue,omitempty" protobuf:"bytes,14,opt,name=apQueue"`
	APCard     *NodeDeviceAPCardCapability     `json:"apCard,omitempty" protobuf:"bytes,15,opt,name=apCard"`
	APMatrix   *NodeDeviceAPMatrixCapability   `json:"apMatrix,omitempty" protobuf:"bytes,16,opt,name=apMatrix"`
}

// +gogo:genproto=true
type NodeDeviceIDName struct {
	ID   string `xml:"id,attr" json:"id" protobuf:"bytes,1,opt,name=id"`
	Name string `xml:",chardata" json:"name" protobuf:"bytes,2,opt,name=name"`
}

// +gogo:genproto=true
type NodeDevicePCIExpress struct {
	Links []NodeDevicePCIExpressLink `xml:"link" json:"links" protobuf:"bytes,1,rep,name=links"`
}

// +gogo:genproto=true
type NodeDevicePCIExpressLink struct {
	Validity string  `xml:"validity,attr,omitempty" json:"validity,omitempty" protobuf:"bytes,1,opt,name=validity"`
	Speed    float64 `xml:"speed,attr,omitempty" json:"speed,omitempty" protobuf:"fixed64,2,opt,name=speed"`
	Port     *uint   `xml:"port,attr" json:"port" protobuf:"varint,3,opt,name=port"`
	Width    *uint   `xml:"width,attr" json:"width" protobuf:"varint,4,opt,name=width"`
}

// +gogo:genproto=true
type NodeDeviceIOMMUGroup struct {
	Number  int                    `xml:"number,attr" json:"number" protobuf:"varint,1,opt,name=number"`
	Address []NodeDevicePCIAddress `xml:"address" json:"address" protobuf:"bytes,2,rep,name=address"`
}

// +gogo:genproto=true
type NodeDeviceNUMA struct {
	Node int `xml:"node,attr" json:"node" protobuf:"varint,1,opt,name=node"`
}

// +gogo:genproto=true
type NodeDevicePCICapability struct {
	Class        string                       `xml:"class,omitempty" json:"class,omitempty" protobuf:"bytes,1,opt,name=class"`
	Domain       *uint                        `xml:"domain" json:"domain,omitempty" protobuf:"varint,2,opt,name=domain"`
	Bus          *uint                        `xml:"bus" json:"bus,omitempty" protobuf:"varint,3,opt,name=bus"`
	Slot         *uint                        `xml:"slot" json:"slot,omitempty" protobuf:"varint,4,opt,name=slot"`
	Function     *uint                        `xml:"function" json:"function,omitempty" protobuf:"varint,5,opt,name=function"`
	Product      NodeDeviceIDName             `xml:"product,omitempty" json:"product,omitempty" protobuf:"bytes,6,opt,name=product"`
	Vendor       NodeDeviceIDName             `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,7,opt,name=vendor"`
	IOMMUGroup   *NodeDeviceIOMMUGroup        `xml:"iommuGroup" json:"iommuGroup,omitempty" protobuf:"bytes,8,opt,name=iommuGroup"`
	NUMA         *NodeDeviceNUMA              `xml:"numa" json:"numa,omitempty" protobuf:"bytes,9,opt,name=numa"`
	PCIExpress   *NodeDevicePCIExpress        `xml:"pci-express" json:"pciExpress,omitempty" protobuf:"bytes,10,opt,name=pciExpress"`
	Capabilities []NodeDevicePCISubCapability `xml:"capability" json:"capabilities,omitempty" protobuf:"bytes,11,rep,name=capabilities"`
}

// +gogo:genproto=true
type NodeDevicePCIAddress struct {
	Domain   *uint `xml:"domain,attr" json:"domain,omitempty" protobuf:"varint,1,opt,name=domain"`
	Bus      *uint `xml:"bus,attr" json:"bus,omitempty" protobuf:"varint,2,opt,name=bus"`
	Slot     *uint `xml:"slot,attr" json:"slot,omitempty" protobuf:"varint,3,opt,name=slot"`
	Function *uint `xml:"function,attr" json:"function,omitempty" protobuf:"varint,4,opt,name=function"`
}

// +gogo:genproto=true
type NodeDevicePCISubCapability struct {
	VirtFunctions *NodeDevicePCIVirtFunctionsCapability `json:"virtFunctions,omitempty" protobuf:"bytes,1,opt,name=virtFunctions"`
	PhysFunction  *NodeDevicePCIPhysFunctionCapability  `json:"physFunction,omitempty" protobuf:"bytes,2,opt,name=physFunction"`
	MDevTypes     *NodeDevicePCIMDevTypesCapability     `json:"mDevTypes,omitempty" protobuf:"bytes,3,opt,name=mDevTypes"`
	Bridge        *NodeDevicePCIBridgeCapability        `json:"bridge,omitempty" protobuf:"bytes,4,opt,name=bridge"`
}

// +gogo:genproto=true
type NodeDevicePCIVirtFunctionsCapability struct {
	Address  []NodeDevicePCIAddress `xml:"address,omitempty" json:"address,omitempty" protobuf:"bytes,1,rep,name=address"`
	MaxCount int                    `xml:"maxCount,attr,omitempty" json:"maxCount,omitempty" protobuf:"varint,2,opt,name=maxCount"`
}

// +gogo:genproto=true
type NodeDevicePCIPhysFunctionCapability struct {
	Address NodeDevicePCIAddress `xml:"address,omitempty" json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
}

// +gogo:genproto=true
type NodeDevicePCIMDevTypesCapability struct {
	Types []NodeDeviceMDevType `xml:"type" json:"types" protobuf:"bytes,1,rep,name=types"`
}

// +gogo:genproto=true
type NodeDeviceMDevType struct {
	ID                 string `xml:"id,attr" json:"id" protobuf:"bytes,1,opt,name=id"`
	Name               string `xml:"name" json:"name" protobuf:"bytes,2,opt,name=name"`
	DeviceAPI          string `xml:"deviceAPI" json:"deviceAPI" protobuf:"bytes,3,opt,name=deviceAPI"`
	AvailableInstances uint   `xml:"availableInstances" json:"availableInstances" protobuf:"varint,4,opt,name=availableInstances"`
}

// +gogo:genproto=true
type NodeDevicePCIBridgeCapability struct {
}

// +gogo:genproto=true
type NodeDeviceSystemHardware struct {
	Vendor  string `xml:"vendor" json:"vendor" protobuf:"bytes,1,opt,name=vendor"`
	Version string `xml:"version" json:"version" protobuf:"bytes,2,opt,name=version"`
	Serial  string `xml:"serial" json:"serial" protobuf:"bytes,3,opt,name=serial"`
	UUID    string `xml:"uuid" json:"uuid" protobuf:"bytes,4,opt,name=uuid"`
}

// +gogo:genproto=true
type NodeDeviceSystemFirmware struct {
	Vendor      string `xml:"vendor" json:"vendor" protobuf:"bytes,1,opt,name=vendor"`
	Version     string `xml:"version" json:"version" protobuf:"bytes,2,opt,name=version"`
	ReleaseData string `xml:"release_date" json:"releaseData" protobuf:"bytes,3,opt,name=releaseData"`
}

// +gogo:genproto=true
type NodeDeviceSystemCapability struct {
	Product  string                    `xml:"product,omitempty" json:"product" protobuf:"bytes,1,opt,name=product"`
	Hardware *NodeDeviceSystemHardware `xml:"hardware" json:"hardware,omitempty" protobuf:"bytes,2,opt,name=hardware"`
	Firmware *NodeDeviceSystemFirmware `xml:"firmware" json:"firmware,omitempty" protobuf:"bytes,3,opt,name=firmware"`
}

// +gogo:genproto=true
type NodeDeviceUSBDeviceCapability struct {
	Bus     int              `xml:"bus" json:"bus" protobuf:"varint,1,opt,name=bus"`
	Device  int              `xml:"device" json:"device" protobuf:"varint,2,opt,name=device"`
	Product NodeDeviceIDName `xml:"product,omitempty" json:"product,omitempty" protobuf:"bytes,3,opt,name=product"`
	Vendor  NodeDeviceIDName `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,4,opt,name=vendor"`
}

// +gogo:genproto=true
type NodeDeviceUSBCapability struct {
	Number      int    `xml:"number" json:"number" protobuf:"varint,1,opt,name=number"`
	Class       int    `xml:"class" json:"class" protobuf:"varint,2,opt,name=class"`
	Subclass    int    `xml:"subclass" json:"subclass" protobuf:"varint,3,opt,name=subclass"`
	Protocol    int    `xml:"protocol" json:"protocol" protobuf:"varint,4,opt,name=protocol"`
	Description string `xml:"description,omitempty" json:"description,omitempty" protobuf:"bytes,5,opt,name=description"`
}

// +gogo:genproto=true
type NodeDeviceNetOffloadFeatures struct {
	Name string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
}

// +gogo:genproto=true
type NodeDeviceNetLink struct {
	State string `xml:"state,attr" json:"state" protobuf:"bytes,1,opt,name=state"`
	Speed string `xml:"speed,attr,omitempty" json:"speed,omitempty" protobuf:"bytes,2,opt,name=speed"`
}

// +gogo:genproto=true
type NodeDeviceNetSubCapability struct {
	Wireless80211 *NodeDeviceNet80211Capability `json:"wireless80211" protobuf:"bytes,1,opt,name=wireless80211"`
	Ethernet80203 *NodeDeviceNet80203Capability `json:"ethernet80203" protobuf:"bytes,2,opt,name=ethernet80203"`
}

// +gogo:genproto=true
type NodeDeviceNet80211Capability struct {
}

// +gogo:genproto=true
type NodeDeviceNet80203Capability struct {
}

// +gogo:genproto=true
type NodeDeviceNetCapability struct {
	Interface  string                         `xml:"interface" json:"interface" protobuf:"bytes,1,opt,name=interface"`
	Address    string                         `xml:"address" json:"address" protobuf:"bytes,2,opt,name=address"`
	Link       *NodeDeviceNetLink             `xml:"link" json:"link,omitempty" protobuf:"bytes,3,opt,name=link"`
	Features   []NodeDeviceNetOffloadFeatures `xml:"feature,omitempty" json:"features" protobuf:"bytes,4,rep,name=features"`
	Capability []NodeDeviceNetSubCapability   `xml:"capability" json:"capability" protobuf:"bytes,5,rep,name=capability"`
}

// +gogo:genproto=true
type NodeDeviceSCSIVPortOpsCapability struct {
	VPorts    int `xml:"vports,omitempty" json:"vports,omitempty" protobuf:"varint,1,opt,name=vports"`
	MaxVPorts int `xml:"maxvports,omitempty" json:"maxvports,omitempty" protobuf:"varint,2,opt,name=maxvports"`
}

// +gogo:genproto=true
type NodeDeviceSCSIFCHostCapability struct {
	WWNN      string `xml:"wwnn,omitempty" json:"wwnn,omitempty" protobuf:"bytes,1,opt,name=wwnn"`
	WWPN      string `xml:"wwpn,omitempty" json:"wwpn,omitempty" protobuf:"bytes,2,opt,name=wwpn"`
	FabricWWN string `xml:"fabric_wwn,omitempty" json:"fabricWwn,omitempty" protobuf:"bytes,3,opt,name=fabricWwn"`
}

// +gogo:genproto=true
type NodeDeviceSCSIHostSubCapability struct {
	VPortOps *NodeDeviceSCSIVPortOpsCapability `json:"vportops,omitempty" protobuf:"bytes,1,opt,name=vportops"`
	FCHost   *NodeDeviceSCSIFCHostCapability   `json:"fchost,omitempty" protobuf:"bytes,2,opt,name=fchost"`
}

// +gogo:genproto=true
type NodeDeviceSCSIHostCapability struct {
	Host       uint                              `xml:"host" json:"host" protobuf:"varint,1,opt,name=host"`
	UniqueID   *uint                             `xml:"unique_id" json:"uniqueId,omitempty" protobuf:"varint,2,opt,name=uniqueId"`
	Capability []NodeDeviceSCSIHostSubCapability `xml:"capability" json:"capability" protobuf:"bytes,3,rep,name=capability"`
}

// +gogo:genproto=true
type NodeDeviceSCSITargetCapability struct {
	Target     string                              `xml:"target" json:"target" protobuf:"bytes,1,opt,name=target"`
	Capability []NodeDeviceSCSITargetSubCapability `xml:"capability" json:"capability" protobuf:"bytes,2,rep,name=capability"`
}

// +gogo:genproto=true
type NodeDeviceSCSITargetSubCapability struct {
	FCRemotePort *NodeDeviceSCSIFCRemotePortCapability `json:"fcRemotePort,omitempty" protobuf:"bytes,1,opt,name=fcRemotePort"`
}

// +gogo:genproto=true
type NodeDeviceSCSIFCRemotePortCapability struct {
	RPort string `xml:"rport" json:"rport" protobuf:"bytes,1,opt,name=rport"`
	WWPN  string `xml:"wwpn" json:"wwpn" protobuf:"bytes,2,opt,name=wwpn"`
}

// +gogo:genproto=true
type NodeDeviceSCSICapability struct {
	Host   int    `xml:"host" json:"host" protobuf:"varint,1,opt,name=host"`
	Bus    int    `xml:"bus" json:"bus" protobuf:"varint,2,opt,name=bus"`
	Target int    `xml:"target" json:"target" protobuf:"varint,3,opt,name=target"`
	Lun    int    `xml:"lun" json:"lun" protobuf:"varint,4,opt,name=lun"`
	Type   string `xml:"type" json:"type" protobuf:"bytes,5,opt,name=type"`
}

// +gogo:genproto=true
type NodeDeviceStorageSubCapability struct {
	Removable *NodeDeviceStorageRemovableCapability `json:"removable,omitempty" protobuf:"bytes,1,opt,name=removable"`
}

// +gogo:genproto=true
type NodeDeviceStorageRemovableCapability struct {
	MediaAvailable   *uint  `xml:"media_available" json:"mediaAvailable,omitempty" protobuf:"varint,1,opt,name=mediaAvailable"`
	MediaSize        *uint  `xml:"media_size" json:"mediaSize,omitempty" protobuf:"varint,2,opt,name=mediaSize"`
	MediaLabel       string `xml:"media_label,omitempty" json:"mediaLabel,omitempty" protobuf:"bytes,3,opt,name=mediaLabel"`
	LogicalBlockSize *uint  `xml:"logical_block_size" json:"logicalBlockSize,omitempty" protobuf:"varint,4,opt,name=logicalBlockSize"`
	NumBlocks        *uint  `xml:"num_blocks" json:"numBlocks,omitempty" protobuf:"varint,5,opt,name=numBlocks"`
}

// +gogo:genproto=true
type NodeDeviceStorageCapability struct {
	Block            string                           `xml:"block,omitempty" json:"block,omitempty" protobuf:"bytes,1,opt,name=block"`
	Bus              string                           `xml:"bus,omitempty" json:"bus,omitempty" protobuf:"bytes,2,opt,name=bus"`
	DriverType       string                           `xml:"drive_type,omitempty" json:"driverType,omitempty" protobuf:"bytes,3,opt,name=driverType"`
	Model            string                           `xml:"model,omitempty" json:"model,omitempty" protobuf:"bytes,4,opt,name=model"`
	Vendor           string                           `xml:"vendor,omitempty" json:"vendor,omitempty" protobuf:"bytes,5,opt,name=vendor"`
	Serial           string                           `xml:"serial,omitempty" json:"serial,omitempty" protobuf:"bytes,6,opt,name=serial"`
	Size             *uint                            `xml:"size" json:"size,omitempty"`
	LogicalBlockSize *uint                            `xml:"logical_block_size" json:"logicalBlockSize,omitempty" protobuf:"varint,8,opt,name=logicalBlockSize"`
	NumBlocks        *uint                            `xml:"num_blocks" json:"numBlocks,omitempty" protobuf:"varint,9,opt,name=numBlocks"`
	Capability       []NodeDeviceStorageSubCapability `xml:"capability" json:"capability" protobuf:"bytes,10,rep,name=capability"`
}

// +gogo:genproto=true
type NodeDeviceDRMCapability struct {
	Type string `xml:"type" json:"type" protobuf:"bytes,1,opt,name=type"`
}

// +gogo:genproto=true
type NodeDeviceCCWCapability struct {
	CSSID *uint `xml:"cssid" json:"cssid,omitempty" protobuf:"varint,1,opt,name=cssid"`
	SSID  *uint `xml:"ssid" json:"ssid,omitempty" protobuf:"varint,2,opt,name=ssid"`
	DevNo *uint `xml:"devno" json:"devno,omitempty" protobuf:"varint,3,opt,name=devno"`
}

// +gogo:genproto=true
type NodeDeviceMDevCapability struct {
	Type       *NodeDeviceMDevCapabilityType   `xml:"type" json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	IOMMUGroup *NodeDeviceIOMMUGroup           `xml:"iommuGroup" json:"iommuGroup,omitempty" protobuf:"bytes,2,opt,name=iommuGroup"`
	UUID       string                          `xml:"uuid,omitempty" json:"uuid,omitempty" protobuf:"bytes,3,opt,name=uuid"`
	Attrs      []NodeDeviceMDevCapabilityAttrs `xml:"attr,omitempty" json:"attrs,omitempty" protobuf:"bytes,4,rep,name=attrs"`
}

// +gogo:genproto=true
type NodeDeviceMDevCapabilityType struct {
	ID string `xml:"id,attr" json:"id" protobuf:"bytes,1,opt,name=id"`
}

// +gogo:genproto=true
type NodeDeviceMDevCapabilityAttrs struct {
	Name  string `xml:"name,attr" json:"name" protobuf:"bytes,1,opt,name=name"`
	Value string `xml:"value,attr" json:"value" protobuf:"bytes,2,opt,name=value"`
}

// +gogo:genproto=true
type NodeDeviceCSSCapability struct {
	CSSID        *uint                        `xml:"cssid" json:"cssid,omitempty" protobuf:"varint,1,opt,name=cssid"`
	SSID         *uint                        `xml:"ssid" json:"ssid,omitempty" protobuf:"varint,2,opt,name=ssid"`
	DevNo        *uint                        `xml:"devno" json:"devno,omitempty" protobuf:"varint,3,opt,name=devno"`
	Capabilities []NodeDeviceCSSSubCapability `xml:"capability" json:"capabilities" protobuf:"bytes,4,rep,name=capabilities"`
}

// +gogo:genproto=true
type NodeDeviceCSSSubCapability struct {
	MDevTypes *NodeDeviceCSSMDevTypesCapability `json:"mDevTypes,omitempty" protobuf:"bytes,1,opt,name=mDevTypes"`
}

// +gogo:genproto=true
type NodeDeviceCSSMDevTypesCapability struct {
	Types []NodeDeviceMDevType `xml:"type" json:"types" protobuf:"bytes,1,rep,name=types"`
}

// +gogo:genproto=true
type NodeDeviceAPQueueCapability struct {
	APAdapter string `xml:"ap-adapter" json:"apAdapter" protobuf:"bytes,1,opt,name=apAdapter"`
	APDomain  string `xml:"ap-domain" json:"apDomain" protobuf:"bytes,2,opt,name=apDomain"`
}

// +gogo:genproto=true
type NodeDeviceAPCardCapability struct {
	APAdapter string `xml:"ap-adapter" json:"apAdapter" protobuf:"bytes,1,opt,name=apAdapter"`
}

// +gogo:genproto=true
type NodeDeviceAPMatrixCapability struct {
	Capabilities []NodeDeviceAPMatrixSubCapability `xml:"capability" json:"capabilities" protobuf:"bytes,1,rep,name=capabilities"`
}

// +gogo:genproto=true
type NodeDeviceAPMatrixSubCapability struct {
	MDevTypes *NodeDeviceAPMatrixMDevTypesCapability `json:"mDevTypes" protobuf:"bytes,1,opt,name=mDevTypes"`
}

// +gogo:genproto=true
type NodeDeviceAPMatrixMDevTypesCapability struct {
	Types []NodeDeviceMDevType `xml:"type" json:"types" protobuf:"bytes,1,rep,name=types"`
}

func (a *NodeDevicePCIAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *NodeDevicePCIAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (c *NodeDeviceCSSSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "mdev_types":
		var mdevTypesCaps NodeDeviceCSSMDevTypesCapability
		if err := d.DecodeElement(&mdevTypesCaps, &start); err != nil {
			return err
		}
		c.MDevTypes = &mdevTypesCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceCSSSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.MDevTypes != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev_types",
		})
		return e.EncodeElement(c.MDevTypes, start)
	}
	return nil
}

func (c *NodeDeviceCCWCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if c.CSSID != nil {
		cssid := xml.StartElement{
			Name: xml.Name{Local: "cssid"},
		}
		e.EncodeToken(cssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.CSSID)))
		e.EncodeToken(cssid.End())
	}
	if c.SSID != nil {
		ssid := xml.StartElement{
			Name: xml.Name{Local: "ssid"},
		}
		e.EncodeToken(ssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.SSID)))
		e.EncodeToken(ssid.End())
	}
	if c.DevNo != nil {
		devno := xml.StartElement{
			Name: xml.Name{Local: "devno"},
		}
		e.EncodeToken(devno)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *c.DevNo)))
		e.EncodeToken(devno.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (c *NodeDeviceCCWCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			cdata, err := d.Token()
			if err != nil {
				return err
			}

			if tok.Name.Local != "cssid" &&
				tok.Name.Local != "ssid" &&
				tok.Name.Local != "devno" {
				continue
			}

			chardata, ok := cdata.(xml.CharData)
			if !ok {
				return fmt.Errorf("Expected text for CCW '%s'", tok.Name.Local)
			}

			valstr := strings.TrimPrefix(string(chardata), "0x")
			val, err := strconv.ParseUint(valstr, 16, 64)
			if err != nil {
				return err
			}

			vali := uint(val)
			if tok.Name.Local == "cssid" {
				c.CSSID = &vali
			} else if tok.Name.Local == "ssid" {
				c.SSID = &vali
			} else if tok.Name.Local == "devno" {
				c.DevNo = &vali
			}
		}
	}
	return nil
}

func (c *NodeDeviceCSSCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if c.CSSID != nil {
		cssid := xml.StartElement{
			Name: xml.Name{Local: "cssid"},
		}
		e.EncodeToken(cssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.CSSID)))
		e.EncodeToken(cssid.End())
	}
	if c.SSID != nil {
		ssid := xml.StartElement{
			Name: xml.Name{Local: "ssid"},
		}
		e.EncodeToken(ssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.SSID)))
		e.EncodeToken(ssid.End())
	}
	if c.DevNo != nil {
		devno := xml.StartElement{
			Name: xml.Name{Local: "devno"},
		}
		e.EncodeToken(devno)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *c.DevNo)))
		e.EncodeToken(devno.End())
	}
	if c.Capabilities != nil {
		for _, subcap := range c.Capabilities {
			start := xml.StartElement{
				Name: xml.Name{Local: "capability"},
			}
			e.EncodeElement(&subcap, start)
		}
	}
	e.EncodeToken(start.End())
	return nil
}

func (c *NodeDeviceCSSCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			cdata, err := d.Token()
			if err != nil {
				return err
			}

			if tok.Name.Local == "capability" {
				subcap := &NodeDeviceCSSSubCapability{}
				err := d.DecodeElement(subcap, &tok)
				if err != nil {
					return err
				}
				c.Capabilities = append(c.Capabilities, *subcap)
				continue
			}

			if tok.Name.Local != "cssid" &&
				tok.Name.Local != "ssid" &&
				tok.Name.Local != "devno" {
				continue
			}

			chardata, ok := cdata.(xml.CharData)
			if !ok {
				return fmt.Errorf("Expected text for CSS '%s'", tok.Name.Local)
			}

			valstr := strings.TrimPrefix(string(chardata), "0x")
			val, err := strconv.ParseUint(valstr, 16, 64)
			if err != nil {
				return err
			}

			vali := uint(val)
			if tok.Name.Local == "cssid" {
				c.CSSID = &vali
			} else if tok.Name.Local == "ssid" {
				c.SSID = &vali
			} else if tok.Name.Local == "devno" {
				c.DevNo = &vali
			}
		}
	}
	return nil
}

func (c *NodeDevicePCISubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "virt_functions":
		var virtFuncCaps NodeDevicePCIVirtFunctionsCapability
		if err := d.DecodeElement(&virtFuncCaps, &start); err != nil {
			return err
		}
		c.VirtFunctions = &virtFuncCaps
	case "phys_function":
		var physFuncCaps NodeDevicePCIPhysFunctionCapability
		if err := d.DecodeElement(&physFuncCaps, &start); err != nil {
			return err
		}
		c.PhysFunction = &physFuncCaps
	case "mdev_types":
		var mdevTypeCaps NodeDevicePCIMDevTypesCapability
		if err := d.DecodeElement(&mdevTypeCaps, &start); err != nil {
			return err
		}
		c.MDevTypes = &mdevTypeCaps
	case "pci-bridge":
		var bridgeCaps NodeDevicePCIBridgeCapability
		if err := d.DecodeElement(&bridgeCaps, &start); err != nil {
			return err
		}
		c.Bridge = &bridgeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDevicePCISubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.VirtFunctions != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virt_functions",
		})
		return e.EncodeElement(c.VirtFunctions, start)
	} else if c.PhysFunction != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "phys_function",
		})
		return e.EncodeElement(c.PhysFunction, start)
	} else if c.MDevTypes != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev_types",
		})
		return e.EncodeElement(c.MDevTypes, start)
	} else if c.Bridge != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci-bridge",
		})
		return e.EncodeElement(c.Bridge, start)
	}
	return nil
}

func (c *NodeDeviceSCSITargetSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "fc_remote_port":
		var fcCaps NodeDeviceSCSIFCRemotePortCapability
		if err := d.DecodeElement(&fcCaps, &start); err != nil {
			return err
		}
		c.FCRemotePort = &fcCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceSCSITargetSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.FCRemotePort != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "fc_remote_port",
		})
		return e.EncodeElement(c.FCRemotePort, start)
	}
	return nil
}

func (c *NodeDeviceSCSIHostSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "fc_host":
		var fcCaps NodeDeviceSCSIFCHostCapability
		if err := d.DecodeElement(&fcCaps, &start); err != nil {
			return err
		}
		c.FCHost = &fcCaps
	case "vport_ops":
		var vportCaps NodeDeviceSCSIVPortOpsCapability
		if err := d.DecodeElement(&vportCaps, &start); err != nil {
			return err
		}
		c.VPortOps = &vportCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceSCSIHostSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.FCHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "fc_host",
		})
		return e.EncodeElement(c.FCHost, start)
	} else if c.VPortOps != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "vport_ops",
		})
		return e.EncodeElement(c.VPortOps, start)
	}
	return nil
}

func (c *NodeDeviceStorageSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "removable":
		var removeCaps NodeDeviceStorageRemovableCapability
		if err := d.DecodeElement(&removeCaps, &start); err != nil {
			return err
		}
		c.Removable = &removeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceStorageSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Removable != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "removable",
		})
		return e.EncodeElement(c.Removable, start)
	}
	return nil
}

func (c *NodeDeviceNetSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "80211":
		var wlanCaps NodeDeviceNet80211Capability
		if err := d.DecodeElement(&wlanCaps, &start); err != nil {
			return err
		}
		c.Wireless80211 = &wlanCaps
	case "80203":
		var ethCaps NodeDeviceNet80203Capability
		if err := d.DecodeElement(&ethCaps, &start); err != nil {
			return err
		}
		c.Ethernet80203 = &ethCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceNetSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Wireless80211 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "80211",
		})
		return e.EncodeElement(c.Wireless80211, start)
	} else if c.Ethernet80203 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "80203",
		})
		return e.EncodeElement(c.Ethernet80203, start)
	}
	return nil
}

func (c *NodeDeviceAPMatrixSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "mdev_types":
		var mdevTypeCaps NodeDeviceAPMatrixMDevTypesCapability
		if err := d.DecodeElement(&mdevTypeCaps, &start); err != nil {
			return err
		}
		c.MDevTypes = &mdevTypeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceAPMatrixSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.MDevTypes != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev_types",
		})
		return e.EncodeElement(c.MDevTypes, start)
	}
	return nil
}

func (c *NodeDeviceCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "pci":
		var pciCaps NodeDevicePCICapability
		if err := d.DecodeElement(&pciCaps, &start); err != nil {
			return err
		}
		c.PCI = &pciCaps
	case "system":
		var systemCaps NodeDeviceSystemCapability
		if err := d.DecodeElement(&systemCaps, &start); err != nil {
			return err
		}
		c.System = &systemCaps
	case "usb_device":
		var usbdevCaps NodeDeviceUSBDeviceCapability
		if err := d.DecodeElement(&usbdevCaps, &start); err != nil {
			return err
		}
		c.USBDevice = &usbdevCaps
	case "usb":
		var usbCaps NodeDeviceUSBCapability
		if err := d.DecodeElement(&usbCaps, &start); err != nil {
			return err
		}
		c.USB = &usbCaps
	case "net":
		var netCaps NodeDeviceNetCapability
		if err := d.DecodeElement(&netCaps, &start); err != nil {
			return err
		}
		c.Net = &netCaps
	case "scsi_host":
		var scsiHostCaps NodeDeviceSCSIHostCapability
		if err := d.DecodeElement(&scsiHostCaps, &start); err != nil {
			return err
		}
		c.SCSIHost = &scsiHostCaps
	case "scsi_target":
		var scsiTargetCaps NodeDeviceSCSITargetCapability
		if err := d.DecodeElement(&scsiTargetCaps, &start); err != nil {
			return err
		}
		c.SCSITarget = &scsiTargetCaps
	case "scsi":
		var scsiCaps NodeDeviceSCSICapability
		if err := d.DecodeElement(&scsiCaps, &start); err != nil {
			return err
		}
		c.SCSI = &scsiCaps
	case "storage":
		var storageCaps NodeDeviceStorageCapability
		if err := d.DecodeElement(&storageCaps, &start); err != nil {
			return err
		}
		c.Storage = &storageCaps
	case "drm":
		var drmCaps NodeDeviceDRMCapability
		if err := d.DecodeElement(&drmCaps, &start); err != nil {
			return err
		}
		c.DRM = &drmCaps
	case "ccw":
		var ccwCaps NodeDeviceCCWCapability
		if err := d.DecodeElement(&ccwCaps, &start); err != nil {
			return err
		}
		c.CCW = &ccwCaps
	case "mdev":
		var mdevCaps NodeDeviceMDevCapability
		if err := d.DecodeElement(&mdevCaps, &start); err != nil {
			return err
		}
		c.MDev = &mdevCaps
	case "css":
		var cssCaps NodeDeviceCSSCapability
		if err := d.DecodeElement(&cssCaps, &start); err != nil {
			return err
		}
		c.CSS = &cssCaps
	case "ap_queue":
		var apCaps NodeDeviceAPQueueCapability
		if err := d.DecodeElement(&apCaps, &start); err != nil {
			return err
		}
		c.APQueue = &apCaps
	case "ap_matrix":
		var apCaps NodeDeviceAPMatrixCapability
		if err := d.DecodeElement(&apCaps, &start); err != nil {
			return err
		}
		c.APMatrix = &apCaps
	case "ap_card":
		var apCaps NodeDeviceAPCardCapability
		if err := d.DecodeElement(&apCaps, &start); err != nil {
			return err
		}
		c.APCard = &apCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(c.PCI, start)
	} else if c.System != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "system",
		})
		return e.EncodeElement(c.System, start)
	} else if c.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return e.EncodeElement(c.USB, start)
	} else if c.USBDevice != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb_device",
		})
		return e.EncodeElement(c.USBDevice, start)
	} else if c.Net != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		return e.EncodeElement(c.Net, start)
	} else if c.SCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		return e.EncodeElement(c.SCSI, start)
	} else if c.SCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		return e.EncodeElement(c.SCSIHost, start)
	} else if c.SCSITarget != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_target",
		})
		return e.EncodeElement(c.SCSITarget, start)
	} else if c.Storage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		return e.EncodeElement(c.Storage, start)
	} else if c.DRM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drm",
		})
		return e.EncodeElement(c.DRM, start)
	} else if c.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return e.EncodeElement(c.CCW, start)
	} else if c.MDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		return e.EncodeElement(c.MDev, start)
	} else if c.CSS != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "css",
		})
		return e.EncodeElement(c.CSS, start)
	} else if c.APQueue != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ap_queue",
		})
		return e.EncodeElement(c.APQueue, start)
	} else if c.APCard != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ap_card",
		})
		return e.EncodeElement(c.APCard, start)
	} else if c.APMatrix != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ap_matrix",
		})
		return e.EncodeElement(c.APMatrix, start)
	}
	return nil
}

func (c *NodeDevice) UnmarshalX(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *NodeDevice) MarshalX() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
