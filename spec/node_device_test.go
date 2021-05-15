package spec

import (
	"strings"
	"testing"
)

var pciDomain int32 = 1
var pciBus int32 = 21
var pciSlot int32 = 10
var pciFunc int32 = 50

var NodeDeviceTestData = []struct {
	Object *NodeDevice
	XML    []string
}{
	{
		Object: &NodeDevice{
			Name:   "pci_0000_81_00_0",
			Parent: "pci_0000_80_01_0",
			Driver: &NodeDeviceDriver{
				Name: "ixgbe",
			},
			Capability: NodeDeviceCapability{
				PCI: &NodeDevicePCICapability{
					Domain:   &pciDomain,
					Bus:      &pciBus,
					Slot:     &pciSlot,
					Function: &pciFunc,
					Product: NodeDeviceIDName{
						ID:   "0x1528",
						Name: "Ethernet Controller 10-Gigabit X540-AT2",
					},
					Vendor: NodeDeviceIDName{
						ID:   "0x8086",
						Name: "Intel Corporation",
					},
					IOMMUGroup: &NodeDeviceIOMMUGroup{
						Number: 3,
					},
					NUMA: &NodeDeviceNUMA{
						Node: 1,
					},
					Capabilities: []NodeDevicePCISubCapability{
						NodeDevicePCISubCapability{
							VirtFunctions: &NodeDevicePCIVirtFunctionsCapability{
								MaxCount: 63,
							},
						},
					},
				},
			},
		},
		XML: []string{
			`<device>`,
			`  <name>pci_0000_81_00_0</name>`,
			`  <parent>pci_0000_80_01_0</parent>`,
			`  <driver>`,
			`    <name>ixgbe</name>`,
			`  </driver>`,
			`  <capability type="pci">`,
			`    <domain>1</domain>`,
			`    <bus>21</bus>`,
			`    <slot>10</slot>`,
			`    <function>50</function>`,
			`    <product id="0x1528">Ethernet Controller 10-Gigabit X540-AT2</product>`,
			`    <vendor id="0x8086">Intel Corporation</vendor>`,
			`    <iommuGroup number="3"></iommuGroup>`,
			`    <numa node="1"></numa>`,
			`    <capability type="virt_functions" maxCount="63"></capability>`,
			`  </capability>`,
			`</device>`,
		},
	},
}

func TestNodeDevice(t *testing.T) {
	for _, test := range NodeDeviceTestData {
		expect := strings.Join(test.XML, "\n")

		nodeDevice := NodeDevice{}
		err := nodeDevice.Unmarshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		doc, err := nodeDevice.MarshalX()
		if err != nil {
			t.Fatal(err)
		}

		if doc != expect {
			t.Fatal("Bad xml:\n", string(doc), "\n does not match\n", expect, "\n")
		}
	}
}
