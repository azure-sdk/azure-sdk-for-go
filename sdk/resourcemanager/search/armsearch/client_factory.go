// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	internal *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		internal: internal,
	}, nil
}

// NewAdminKeysClient creates a new instance of AdminKeysClient.
func (c *ClientFactory) NewAdminKeysClient(subscriptionID string) *AdminKeysClient {
	return &AdminKeysClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewManagementClient creates a new instance of ManagementClient.
func (c *ClientFactory) NewManagementClient(subscriptionID string) *ManagementClient {
	return &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient.
func (c *ClientFactory) NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string) *NetworkSecurityPerimeterConfigurationsClient {
	return &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewOfferingsClient creates a new instance of OfferingsClient.
func (c *ClientFactory) NewOfferingsClient() *OfferingsClient {
	return &OfferingsClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient(subscriptionID string) *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient(subscriptionID string) *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewQueryKeysClient creates a new instance of QueryKeysClient.
func (c *ClientFactory) NewQueryKeysClient(subscriptionID string) *QueryKeysClient {
	return &QueryKeysClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient(subscriptionID string) *ServicesClient {
	return &ServicesClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient.
func (c *ClientFactory) NewSharedPrivateLinkResourcesClient(subscriptionID string) *SharedPrivateLinkResourcesClient {
	return &SharedPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient(subscriptionID string) *UsagesClient {
	return &UsagesClient{
		subscriptionID: subscriptionID,
		internal:       c.internal,
	}
}
