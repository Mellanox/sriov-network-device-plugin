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

package resources

import (
	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
)

type ddpSelector struct {
	profiles []string
}

// NewDdpSelector returns a DeviceSelector interface to filter devices based on available DDP profile
func NewDdpSelector(profiles []string) types.DeviceSelector {
	return &ddpSelector{profiles: profiles}
}

func (ds *ddpSelector) Filter(inDevices []types.HostDevice) []types.HostDevice {
	filteredList := make([]types.HostDevice, 0)

	for _, dev := range inDevices {
		ddpProfile := dev.(types.PciNetDevice).GetDDPProfiles()
		if ddpProfile != "" && contains(ds.profiles, ddpProfile) {
			filteredList = append(filteredList, dev)
		}
	}

	return filteredList
}
