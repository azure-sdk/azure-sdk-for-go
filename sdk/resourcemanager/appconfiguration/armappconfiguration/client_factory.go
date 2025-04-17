// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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
//   - subscriptionID - The Microsoft Azure subscription ID.
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

// NewConfigurationStoresClient creates a new instance of ConfigurationStoresClient.
func (c *ClientFactory) NewConfigurationStoresClient() *ConfigurationStoresClient {
	return &ConfigurationStoresClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKeyValuesClient creates a new instance of KeyValuesClient.
func (c *ClientFactory) NewKeyValuesClient() *KeyValuesClient {
	return &KeyValuesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
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

// NewReplicasClient creates a new instance of ReplicasClient.
func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	return &ReplicasClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSnapshotsClient creates a new instance of SnapshotsClient.
func (c *ClientFactory) NewSnapshotsClient() *SnapshotsClient {
	return &SnapshotsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
