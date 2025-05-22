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

var _ = Describe("PKeySelector", func() {
	Describe("PKey selector", func() {
		Context("filtering", func() {
			It("should return devices matching given PKeys", func() {
				pKeys := []string{"0x1", "0x2"}
				sel := resources.NewPKeySelector(pKeys)

				dev0 := mocks.PciNetDevice{}
				dev0.On("GetPKey").Return("0x1")

				dev1 := mocks.PciNetDevice{}
				dev1.On("GetPKey").Return("0x2")

				dev2 := mocks.PciNetDevice{}
				dev2.On("GetPKey").Return("0x3")

				in := []types.HostDevice{&dev0, &dev1, &dev2}
				filtered := sel.Filter(in)

				Expect(filtered).To(ContainElement(&dev0))
				Expect(filtered).To(ContainElement(&dev1))
				Expect(filtered).NotTo(ContainElement(&dev2))
			})
		})
	})
})
