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
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewClustersClient creates a new instance of ClustersClient.
func (c *ClientFactory) NewClustersClient() *ClustersClient {
	return &ClustersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDatastoresClient creates a new instance of DatastoresClient.
func (c *ClientFactory) NewDatastoresClient() *DatastoresClient {
	return &DatastoresClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewHostsClient creates a new instance of HostsClient.
func (c *ClientFactory) NewHostsClient() *HostsClient {
	return &HostsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewInventoryItemsClient creates a new instance of InventoryItemsClient.
func (c *ClientFactory) NewInventoryItemsClient() *InventoryItemsClient {
	return &InventoryItemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewResourcePoolsClient creates a new instance of ResourcePoolsClient.
func (c *ClientFactory) NewResourcePoolsClient() *ResourcePoolsClient {
	return &ResourcePoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVCentersClient creates a new instance of VCentersClient.
func (c *ClientFactory) NewVCentersClient() *VCentersClient {
	return &VCentersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVMInstanceGuestAgentsClient creates a new instance of VMInstanceGuestAgentsClient.
func (c *ClientFactory) NewVMInstanceGuestAgentsClient() *VMInstanceGuestAgentsClient {
	return &VMInstanceGuestAgentsClient{
		internal: c.internal,
	}
}

// NewVMInstanceHybridIdentityMetadataClient creates a new instance of VMInstanceHybridIdentityMetadataClient.
func (c *ClientFactory) NewVMInstanceHybridIdentityMetadataClient() *VMInstanceHybridIdentityMetadataClient {
	return &VMInstanceHybridIdentityMetadataClient{
		internal: c.internal,
	}
}

// NewVirtualMachineInstancesClient creates a new instance of VirtualMachineInstancesClient.
func (c *ClientFactory) NewVirtualMachineInstancesClient() *VirtualMachineInstancesClient {
	return &VirtualMachineInstancesClient{
		internal: c.internal,
	}
}

// NewVirtualMachineTemplatesClient creates a new instance of VirtualMachineTemplatesClient.
func (c *ClientFactory) NewVirtualMachineTemplatesClient() *VirtualMachineTemplatesClient {
	return &VirtualMachineTemplatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient.
func (c *ClientFactory) NewVirtualNetworksClient() *VirtualNetworksClient {
	return &VirtualNetworksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
