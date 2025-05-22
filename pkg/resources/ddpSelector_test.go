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
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/resources"
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types/mocks"
)

var _ = Describe("DdpSelector", func() {
	Describe("DDP selector", func() {
		Context("initializing", func() {
			It("should populate vendors array", func() {
				profiles := []string{"GTPv1-C", "PPPoE"}
				sel := resources.NewDdpSelector(profiles)
				fmt.Printf("%#v", sel)
				// Expect(sel.GetDPProfiles()).To(ConsistOf(profiles))
			})
		})
		Context("filtering", func() {
			It("should return devices matching DDP profiles", func() {
				profiles := []string{"GTP"}
				sel := resources.NewDdpSelector(profiles)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetPciAddr").Return("0000:01:10.0")
				dev0.On("GetDDPProfiles").Return("GTP")

				dev1 := mocks.PciNetDevice{}
				dev1.On("GetPciAddr").Return("0000:01:10.1")
				dev1.On("GetDDPProfiles").Return("PPPoE")

				dev2 := mocks.PciNetDevice{}
				dev2.On("GetPciAddr").Return("0000:01:10.2")
				dev2.On("GetDDPProfiles").Return("")

				in := []types.HostDevice{&dev0, &dev1}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).NotTo(ContainElement(&dev1))
				Expect(filtered).NotTo(ContainElement(&dev2))
			})
		})
	})
})
