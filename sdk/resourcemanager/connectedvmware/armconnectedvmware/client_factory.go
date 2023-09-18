//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armconnectedvmware

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewClustersClient() *ClustersClient {
	subClient, _ := NewClustersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatastoresClient() *DatastoresClient {
	subClient, _ := NewDatastoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHostsClient() *HostsClient {
	subClient, _ := NewHostsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInventoryItemsClient() *InventoryItemsClient {
	subClient, _ := NewInventoryItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourcePoolsClient() *ResourcePoolsClient {
	subClient, _ := NewResourcePoolsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVCentersClient() *VCentersClient {
	subClient, _ := NewVCentersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVMInstanceGuestAgentsClient() *VMInstanceGuestAgentsClient {
	subClient, _ := NewVMInstanceGuestAgentsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVMInstanceHybridIdentityMetadataClient() *VMInstanceHybridIdentityMetadataClient {
	subClient, _ := NewVMInstanceHybridIdentityMetadataClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachineInstancesClient() *VirtualMachineInstancesClient {
	subClient, _ := NewVirtualMachineInstancesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachineTemplatesClient() *VirtualMachineTemplatesClient {
	subClient, _ := NewVirtualMachineTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworksClient() *VirtualNetworksClient {
	subClient, _ := NewVirtualNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
