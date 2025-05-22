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
package netdevice

import (
	nettypes "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	nadutils "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/utils"

	"github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
)

// nadutils implements types.NadUtils interface
// It's purpose is to wrap the utilities provided by github.com/k8snetworkplumbingwg/network-attachment-definition-client
// in order to make mocking easy for Unit Tests
type nadUtils struct {
}

func (nu *nadUtils) SaveDeviceInfoFile(resourceName, deviceID string, devInfo *nettypes.DeviceInfo) error {
	return nadutils.SaveDeviceInfoForDP(resourceName, deviceID, devInfo)
}

func (nu *nadUtils) CleanDeviceInfoFile(resourceName, deviceID string) error {
	return nadutils.CleanDeviceInfoForDP(resourceName, deviceID)
}

// NewNadUtils returns a new NadUtils
func NewNadUtils() types.NadUtils {
	return &nadUtils{}
}
