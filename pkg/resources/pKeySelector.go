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

type pKeySelector struct {
	pKeys []string
}

// NewPKeySelector returns a DeviceSelector interface to filter devices based on available PKeys
func NewPKeySelector(pKeys []string) types.DeviceSelector {
	return &pKeySelector{pKeys: pKeys}
}

func (ds *pKeySelector) Filter(inDevices []types.HostDevice) []types.HostDevice {
	filteredList := make([]types.HostDevice, 0)

	for _, dev := range inDevices {
		pKey := dev.(types.PciNetDevice).GetPKey()
		if pKey != "" && contains(ds.pKeys, pKey) {
			filteredList = append(filteredList, dev)
		}
	}

	return filteredList
}
