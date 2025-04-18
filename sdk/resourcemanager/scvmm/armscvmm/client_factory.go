// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// NewAvailabilitySetsClient creates a new instance of AvailabilitySetsClient.
func (c *ClientFactory) NewAvailabilitySetsClient() *AvailabilitySetsClient {
	return &AvailabilitySetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudsClient creates a new instance of CloudsClient.
func (c *ClientFactory) NewCloudsClient() *CloudsClient {
	return &CloudsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGuestAgentsClient creates a new instance of GuestAgentsClient.
func (c *ClientFactory) NewGuestAgentsClient() *GuestAgentsClient {
	return &GuestAgentsClient{
		internal: c.internal,
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

// NewVMInstanceHybridIdentityMetadatasClient creates a new instance of VMInstanceHybridIdentityMetadatasClient.
func (c *ClientFactory) NewVMInstanceHybridIdentityMetadatasClient() *VMInstanceHybridIdentityMetadatasClient {
	return &VMInstanceHybridIdentityMetadatasClient{
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

// NewVmmServersClient creates a new instance of VmmServersClient.
func (c *ClientFactory) NewVmmServersClient() *VmmServersClient {
	return &VmmServersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
