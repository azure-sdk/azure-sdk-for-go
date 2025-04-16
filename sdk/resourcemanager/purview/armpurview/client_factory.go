// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

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
//   - subscriptionID - The subscription identifier
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

// NewAccountsClient creates a new instance of AccountsClient.
func (c *ClientFactory) NewAccountsClient() *AccountsClient {
	return &AccountsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDefaultAccountsClient creates a new instance of DefaultAccountsClient.
func (c *ClientFactory) NewDefaultAccountsClient() *DefaultAccountsClient {
	return &DefaultAccountsClient{
		internal: c.internal,
	}
}

// NewFeaturesClient creates a new instance of FeaturesClient.
func (c *ClientFactory) NewFeaturesClient() *FeaturesClient {
	return &FeaturesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIngestionPrivateEndpointConnectionsClient creates a new instance of IngestionPrivateEndpointConnectionsClient.
func (c *ClientFactory) NewIngestionPrivateEndpointConnectionsClient() *IngestionPrivateEndpointConnectionsClient {
	return &IngestionPrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewKafkaConfigurationsClient creates a new instance of KafkaConfigurationsClient.
func (c *ClientFactory) NewKafkaConfigurationsClient() *KafkaConfigurationsClient {
	return &KafkaConfigurationsClient{
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

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	return &UsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
