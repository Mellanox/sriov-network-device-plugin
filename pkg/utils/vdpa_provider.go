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

package utils

import (
	"fmt"

	"github.com/golang/glog"
	vdpa "github.com/k8snetworkplumbingwg/govdpa/pkg/kvdpa"
)

// VdpaProvider is a wrapper type over go-vdpa library
type VdpaProvider interface {
	GetVdpaDeviceByPci(pciAddr string) (vdpa.VdpaDevice, error)
}

type defaultVdpaProvider struct {
}

var vdpaProvider VdpaProvider = &defaultVdpaProvider{}

// SetVdpaProviderInst method would be used by unit tests in other packages
func SetVdpaProviderInst(inst VdpaProvider) {
	vdpaProvider = inst
}

// GetVdpaProvider will be invoked by functions in other packages that would need access to the vdpa library methods.
func GetVdpaProvider() VdpaProvider {
	return vdpaProvider
}

func (defaultVdpaProvider) GetVdpaDeviceByPci(pciAddr string) (vdpa.VdpaDevice, error) {
	// the govdpa library requires the pci address to include the "pci/" prefix
	fullPciAddr := "pci/" + pciAddr
	vdpaDevices, err := vdpa.GetVdpaDevicesByPciAddress(fullPciAddr)
	if err != nil {
		return nil, err
	}
	numVdpaDevices := len(vdpaDevices)
	if numVdpaDevices == 0 {
		return nil, fmt.Errorf("no vdpa device associated to pciAddress %s", pciAddr)
	}
	if numVdpaDevices > 1 {
		glog.Infof("More than one vDPA device found for pciAddress %s, returning the first one", pciAddr)
	}
	return vdpaDevices[0], nil
}
