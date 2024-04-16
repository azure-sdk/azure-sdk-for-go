//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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
//   - subscriptionID - The ID of the target subscription.
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

// NewExtensionMetadataClient creates a new instance of ExtensionMetadataClient.
func (c *ClientFactory) NewExtensionMetadataClient() *ExtensionMetadataClient {
	return &ExtensionMetadataClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGatewaysClient creates a new instance of GatewaysClient.
func (c *ClientFactory) NewGatewaysClient() *GatewaysClient {
	return &GatewaysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMachineExtensionsClient creates a new instance of MachineExtensionsClient.
func (c *ClientFactory) NewMachineExtensionsClient() *MachineExtensionsClient {
	return &MachineExtensionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMachineRunCommandsClient creates a new instance of MachineRunCommandsClient.
func (c *ClientFactory) NewMachineRunCommandsClient() *MachineRunCommandsClient {
	return &MachineRunCommandsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMachinesClient creates a new instance of MachinesClient.
func (c *ClientFactory) NewMachinesClient() *MachinesClient {
	return &MachinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagementClient creates a new instance of ManagementClient.
func (c *ClientFactory) NewManagementClient() *ManagementClient {
	return &ManagementClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkProfileClient creates a new instance of NetworkProfileClient.
func (c *ClientFactory) NewNetworkProfileClient() *NetworkProfileClient {
	return &NetworkProfileClient{
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

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkScopesClient creates a new instance of PrivateLinkScopesClient.
func (c *ClientFactory) NewPrivateLinkScopesClient() *PrivateLinkScopesClient {
	return &PrivateLinkScopesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
