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
	"github.com/k8snetworkplumbingwg/sriovnet"
)

// SriovnetProvider is a wrapper type over sriovnet library
type SriovnetProvider interface {
	GetUplinkRepresentor(vfPciAddress string) (string, error)
	GetUplinkRepresentorFromAux(auxDev string) (string, error)
	GetPfPciFromAux(auxDev string) (string, error)
	GetSfIndexByAuxDev(auxDev string) (int, error)
	GetNetDevicesFromAux(auxDev string) ([]string, error)
	GetAuxNetDevicesFromPci(pciAddr string) ([]string, error)
	GetDefaultPKeyFromPci(pciAddr string) (string, error)
}

type defaultSriovnetProvider struct {
}

var sriovnetProvider SriovnetProvider = &defaultSriovnetProvider{}

// SetSriovnetProviderInst method would be used by unit tests in other packages
func SetSriovnetProviderInst(inst SriovnetProvider) {
	sriovnetProvider = inst
}

// GetSriovnetProvider will be invoked by functions in other packages that would need access to the sriovnet library methods.
func GetSriovnetProvider() SriovnetProvider {
	return sriovnetProvider
}

func (defaultSriovnetProvider) GetUplinkRepresentor(vfPciAddress string) (string, error) {
	return sriovnet.GetUplinkRepresentor(vfPciAddress)
}

func (defaultSriovnetProvider) GetUplinkRepresentorFromAux(auxDev string) (string, error) {
	return sriovnet.GetUplinkRepresentorFromAux(auxDev)
}

func (defaultSriovnetProvider) GetPfPciFromAux(auxDev string) (string, error) {
	return sriovnet.GetPfPciFromAux(auxDev)
}

func (defaultSriovnetProvider) GetSfIndexByAuxDev(auxDev string) (int, error) {
	return sriovnet.GetSfIndexByAuxDev(auxDev)
}

func (defaultSriovnetProvider) GetNetDevicesFromAux(auxDev string) ([]string, error) {
	return sriovnet.GetNetDevicesFromAux(auxDev)
}

func (defaultSriovnetProvider) GetAuxNetDevicesFromPci(pciAddr string) ([]string, error) {
	return sriovnet.GetAuxNetDevicesFromPci(pciAddr)
}

func (defaultSriovnetProvider) GetDefaultPKeyFromPci(pciAddr string) (string, error) {
	return sriovnet.GetDefaultPKeyFromPci(pciAddr)
}
