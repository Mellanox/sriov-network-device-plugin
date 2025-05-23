// Copyright 2025 sriov-network-device-plugin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources_test

import (
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/resources"
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeviceSelectors", func() {
	Describe("vendor selector", func() {
		/*Context("initializing", func() {
			It("should populate vendors array", func() {
				vendors := []string{"8086", "15b3"}
				sel := resources.NewVendorSelector(vendors) // FIXME: .(*vendorSelector)
				// FIXME: Expect(sel.vendors).To(ConsistOf(vendors))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching vendor ID", func() {
				vendors := []string{"8086"}
				sel := resources.NewVendorSelector(vendors)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetVendor").Return("8086")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetVendor").Return("15b3")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
			})
		})
	})
	Describe("device selector", func() {
		/*Context("initializing", func() {
			It("should populate devices array", func() {
				devices := []string{"10ed", "154c"}
				sel := NewDeviceSelector(devices).(*deviceSelector)
				Expect(sel.devices).To(ConsistOf(devices))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching device code", func() {
				devices := []string{"10ed"}
				sel := resources.NewDeviceSelector(devices)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetDeviceCode").Return("10ed")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetDeviceCode").Return("154c")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
			})
		})
	})
	Describe("driver selector", func() {
		/*Context("initializing", func() {
			It("should populate drivers array", func() {
				drivers := []string{"vfio-pci", "igb_uio"}
				sel := resources.NewDriverSelector(drivers).(*driverSelector)
				Expect(sel.drivers).To(ConsistOf(drivers))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching driver name", func() {
				drivers := []string{"vfio-pci"}
				sel := resources.NewDriverSelector(drivers)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetDriver").Return("vfio-pci")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetDriver").Return("i40evf")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
			})
		})
	})
	Describe("pciAddress selector", func() {
		/*Context("initializing", func() {
			It("should populate pciAddresses array", func() {
				pciAddr := []string{"0000:03:02.0", "0000:03:02.1"}
				sel := resources.NewPciAddressSelector(pciAddr).(*pciAddressSelector)
				Expect(sel.pciAddresses).To(ConsistOf(pciAddr))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching on VF pci addresses", func() {
				pciAddr := []string{"0000:03:02.0", "0000:03:02.1"}
				sel := resources.NewPciAddressSelector(pciAddr)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetPciAddr").Return("0000:03:02.0")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetPciAddr").Return("0000:03:02.1")
				dev2 := mocks.PciNetDevice{}
				dev2.On("GetPciAddr").Return("0000:03:02.2")
				dev3 := mocks.PciNetDevice{}
				dev3.On("GetPciAddr").Return("0000:03:02.3")

				in := []types.HostDevice{&dev0, &dev1, &dev2, &dev3}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).To(ContainElement(&dev1))
				Expect(filtered).NotTo(ContainElement(&dev2))
				Expect(filtered).NotTo(ContainElement(&dev3))
			})
		})
	})
	Describe("pfName selector", func() {
		/*Context("initializing", func() {
			It("should populate ifnames array", func() {
				pfNames := []string{"ens0", "eth0"}
				sel := resources.NewPfNameSelector(pfNames).(*pfNameSelector)
				Expect(sel.pfNames).To(ConsistOf(pfNames))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching interface PF name", func() {
				netDevs := []string{"ens0", "ens2f0#1", "ens2f1#0,3-5,7"}
				sel := resources.NewPfNameSelector(netDevs)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetPfNetName").Return("ens0")
				dev0.On("GetFuncID").Return(-1)
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetPfNetName").Return("eth0")
				dev1.On("GetFuncID").Return(-1)
				dev2 := mocks.PciNetDevice{}
				dev2.On("GetPfNetName").Return("ens2f0")
				dev2.On("GetFuncID").Return(1)
				dev3 := mocks.PciNetDevice{}
				dev3.On("GetPfNetName").Return("ens2f1")
				dev3.On("GetFuncID").Return(0)
				dev4 := mocks.PciNetDevice{}
				dev4.On("GetPfNetName").Return("ens2f1")
				dev4.On("GetFuncID").Return(1)
				dev5 := mocks.PciNetDevice{}
				dev5.On("GetPfNetName").Return("ens2f1")
				dev5.On("GetFuncID").Return(2)
				dev6 := mocks.PciNetDevice{}
				dev6.On("GetPfNetName").Return("ens2f1")
				dev6.On("GetFuncID").Return(3)
				dev7 := mocks.PciNetDevice{}
				dev7.On("GetPfNetName").Return("ens2f1")
				dev7.On("GetFuncID").Return(4)
				dev8 := mocks.PciNetDevice{}
				dev8.On("GetPfNetName").Return("ens2f1")
				dev8.On("GetFuncID").Return(5)
				dev9 := mocks.PciNetDevice{}
				dev9.On("GetPfNetName").Return("ens2f1")
				dev9.On("GetFuncID").Return(6)
				dev10 := mocks.PciNetDevice{}
				dev10.On("GetPfNetName").Return("ens2f1")
				dev10.On("GetFuncID").Return(7)

				in := []types.HostDevice{&dev0, &dev1, &dev2,
					&dev3, &dev4, &dev5,
					&dev6, &dev7, &dev8,
					&dev9, &dev10}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
				Expect(filtered).To(ContainElement(&dev2))
				Expect(filtered).To(ContainElement(&dev3))
				Expect(filtered).NotTo(ContainElement(&dev4))
				Expect(filtered).NotTo(ContainElement(&dev5))
				Expect(filtered).To(ContainElement(&dev6))
				Expect(filtered).To(ContainElement(&dev7))
				Expect(filtered).To(ContainElement(&dev8))
				Expect(filtered).NotTo(ContainElement(&dev9))
				Expect(filtered).To(ContainElement(&dev10))
			})
		})
	})

	Describe("rootDevice selector", func() {
		Context("filtering", func() {
			It("should return devices matching interface PF address", func() {
				netDevs := []string{"0000:86:00.0", "0000:86:00.1#1", "0000:86:00.2#0-2,5,7"}
				sel := resources.NewRootDeviceSelector(netDevs)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetPfPciAddr").Return("0000:86:00.0")
				dev0.On("GetFuncID").Return(-1)
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetPfPciAddr").Return("0000:a0:00.0")
				dev1.On("GetFuncID").Return(-1)
				dev2 := mocks.PciNetDevice{}
				dev2.On("GetPfPciAddr").Return("0000:86:00.1")
				dev2.On("GetFuncID").Return(1)
				dev3 := mocks.PciNetDevice{}
				dev3.On("GetPfPciAddr").Return("0000:86:00.2")
				dev3.On("GetFuncID").Return(0)
				dev4 := mocks.PciNetDevice{}
				dev4.On("GetPfPciAddr").Return("0000:86:00.2")
				dev4.On("GetFuncID").Return(1)
				dev5 := mocks.PciNetDevice{}
				dev5.On("GetPfPciAddr").Return("0000:86:00.2")
				dev5.On("GetFuncID").Return(2)
				dev6 := mocks.PciNetDevice{}
				dev6.On("GetPfPciAddr").Return("0000:86:00.2")
				dev6.On("GetFuncID").Return(3)
				dev7 := mocks.PciNetDevice{}
				dev7.On("GetPfPciAddr").Return("0000:86:00.2")
				dev7.On("GetFuncID").Return(4)
				dev8 := mocks.PciNetDevice{}
				dev8.On("GetPfPciAddr").Return("0000:86:00.2")
				dev8.On("GetFuncID").Return(5)
				dev9 := mocks.PciNetDevice{}
				dev9.On("GetPfPciAddr").Return("0000:86:00.2")
				dev9.On("GetFuncID").Return(6)
				dev10 := mocks.PciNetDevice{}
				dev10.On("GetPfPciAddr").Return("0000:86:00.2")
				dev10.On("GetFuncID").Return(7)

				in := []types.HostDevice{&dev0, &dev1, &dev2,
					&dev3, &dev4, &dev5,
					&dev6, &dev7, &dev8,
					&dev9, &dev10}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
				Expect(filtered).To(ContainElement(&dev2))
				Expect(filtered).To(ContainElement(&dev3))
				Expect(filtered).To(ContainElement(&dev4))
				Expect(filtered).To(ContainElement(&dev5))
				Expect(filtered).NotTo(ContainElement(&dev6))
				Expect(filtered).NotTo(ContainElement(&dev7))
				Expect(filtered).To(ContainElement(&dev8))
				Expect(filtered).NotTo(ContainElement(&dev9))
				Expect(filtered).To(ContainElement(&dev10))
			})
		})
	})

	Describe("linkType selector", func() {
		/*Context("initializing", func() {
			It("should populate linkTypes array", func() {
				linkTypes := []string{"ether"}
				sel := NewLinkTypeSelector(linkTypes).(*linkTypeSelector)
				Expect(sel.linkTypes).To(ConsistOf(linkTypes))
			})
		})*/
		Context("filtering", func() {
			It("should return devices matching the correct link type", func() {
				linkTypes := []string{"ether"}
				sel := resources.NewLinkTypeSelector(linkTypes)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetLinkType").Return("ether")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetLinkType").Return("infiniband")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
			})
		})
	})
	Describe("auxTypes selector", func() {
		Context("filtering", func() {
			It("should return devices matching the correct aux type", func() {
				auxTypes := []string{"bar", "baz"}
				sel := resources.NewAuxTypeSelector(auxTypes)

				dev0 := mocks.AuxNetDevice{}
				dev0.On("GetAuxType").Return("bar")
				dev1 := mocks.AuxNetDevice{}
				dev1.On("GetAuxType").Return("baz")
				dev2 := mocks.AuxNetDevice{}
				dev2.On("GetAuxType").Return("")

				in := []types.HostDevice{&dev0, &dev1, &dev2}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).To(ContainElement(&dev1))
				Expect(filtered).NotTo(ContainElement(&dev2))
			})
		})
	})

	Describe("acpiIndex selector", func() {
		Context("filtering", func() {
			It("should return devices matching the correct acpi index", func() {
				acpiIndexes := []string{"101"}
				sel := resources.NewAcpiIndexSelector(acpiIndexes)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetAcpiIndex").Return("101")
				dev1 := mocks.PciNetDevice{}
				dev1.On("GetAcpiIndex").Return("102")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
			})
		})
	})
})
