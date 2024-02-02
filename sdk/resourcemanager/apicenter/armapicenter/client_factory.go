//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapicenter

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewAPIDefinitionsClient creates a new instance of APIDefinitionsClient.
func (c *ClientFactory) NewAPIDefinitionsClient() *APIDefinitionsClient {
	subClient, _ := NewAPIDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAPIVersionsClient creates a new instance of APIVersionsClient.
func (c *ClientFactory) NewAPIVersionsClient() *APIVersionsClient {
	subClient, _ := NewAPIVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewApisClient creates a new instance of ApisClient.
func (c *ClientFactory) NewApisClient() *ApisClient {
	subClient, _ := NewApisClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeploymentsClient creates a new instance of DeploymentsClient.
func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	subClient, _ := NewDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient.
func (c *ClientFactory) NewEnvironmentsClient() *EnvironmentsClient {
	subClient, _ := NewEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMetadataSchemasClient creates a new instance of MetadataSchemasClient.
func (c *ClientFactory) NewMetadataSchemasClient() *MetadataSchemasClient {
	subClient, _ := NewMetadataSchemasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	subClient, _ := NewWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
