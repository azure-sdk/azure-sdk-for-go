//go:build go1.18
// +build go1.18

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
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
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

func (c *ClientFactory) NewAvailabilitySetsClient() *AvailabilitySetsClient {
	subClient, _ := NewAvailabilitySetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCloudsClient() *CloudsClient {
	subClient, _ := NewCloudsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGuestAgentsClient() *GuestAgentsClient {
	subClient, _ := NewGuestAgentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHybridIdentityMetadatasClient() *HybridIdentityMetadatasClient {
	subClient, _ := NewHybridIdentityMetadatasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInventoryItemsClient() *InventoryItemsClient {
	subClient, _ := NewInventoryItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMachineExtensionsClient() *MachineExtensionsClient {
	subClient, _ := NewMachineExtensionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachineTemplatesClient() *VirtualMachineTemplatesClient {
	subClient, _ := NewVirtualMachineTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachinesClient() *VirtualMachinesClient {
	subClient, _ := NewVirtualMachinesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworksClient() *VirtualNetworksClient {
	subClient, _ := NewVirtualNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVmmServersClient() *VmmServersClient {
	subClient, _ := NewVmmServersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
