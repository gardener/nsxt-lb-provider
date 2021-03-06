/*
 * Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package loadbalancer

import (
	"io"

	cloudprovider "k8s.io/cloud-provider"

	"github.com/gardener/nsxt-lb-provider/pkg/loadbalancer/config"
)

const (
	// ProviderName is the provider name
	ProviderName = "vsphere"
)

var (
	// Version contains the build version
	Version string
	// AppName contains the app name
	AppName string
)

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(configReader io.Reader) (cloudprovider.Interface, error) {
		cfg, err := config.ReadConfig(configReader)
		if err != nil {
			return nil, err
		}

		return newProvider(cfg)
	})
}
