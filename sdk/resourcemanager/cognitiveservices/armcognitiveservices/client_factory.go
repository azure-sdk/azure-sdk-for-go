//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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
//   - subscriptionID - The ID of the target subscription.
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

func (c *ClientFactory) NewAccountsClient() *AccountsClient {
	subClient, _ := NewAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCommitmentPlansClient() *CommitmentPlansClient {
	subClient, _ := NewCommitmentPlansClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCommitmentTiersClient() *CommitmentTiersClient {
	subClient, _ := NewCommitmentTiersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeletedAccountsClient() *DeletedAccountsClient {
	subClient, _ := NewDeletedAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	subClient, _ := NewDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEncryptionScopesClient() *EncryptionScopesClient {
	subClient, _ := NewEncryptionScopesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagementClient() *ManagementClient {
	subClient, _ := NewManagementClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewModelsClient() *ModelsClient {
	subClient, _ := NewModelsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRaiBlocklistItemsClient() *RaiBlocklistItemsClient {
	subClient, _ := NewRaiBlocklistItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRaiBlocklistsClient() *RaiBlocklistsClient {
	subClient, _ := NewRaiBlocklistsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRaiContentFiltersClient() *RaiContentFiltersClient {
	subClient, _ := NewRaiContentFiltersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRaiPoliciesClient() *RaiPoliciesClient {
	subClient, _ := NewRaiPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourceSKUsClient() *ResourceSKUsClient {
	subClient, _ := NewResourceSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
