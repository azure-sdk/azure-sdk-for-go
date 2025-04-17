// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

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
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
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

// NewDisasterRecoveryConfigsClient creates a new instance of DisasterRecoveryConfigsClient.
func (c *ClientFactory) NewDisasterRecoveryConfigsClient() *DisasterRecoveryConfigsClient {
	return &DisasterRecoveryConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMigrationConfigsClient creates a new instance of MigrationConfigsClient.
func (c *ClientFactory) NewMigrationConfigsClient() *MigrationConfigsClient {
	return &MigrationConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNamespacesClient creates a new instance of NamespacesClient.
func (c *ClientFactory) NewNamespacesClient() *NamespacesClient {
	return &NamespacesClient{
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

// NewQueuesClient creates a new instance of QueuesClient.
func (c *ClientFactory) NewQueuesClient() *QueuesClient {
	return &QueuesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRulesClient creates a new instance of RulesClient.
func (c *ClientFactory) NewRulesClient() *RulesClient {
	return &RulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient.
func (c *ClientFactory) NewSubscriptionsClient() *SubscriptionsClient {
	return &SubscriptionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTopicsClient creates a new instance of TopicsClient.
func (c *ClientFactory) NewTopicsClient() *TopicsClient {
	return &TopicsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
