// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

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

// NewAPIDefinitionsClient creates a new instance of APIDefinitionsClient.
func (c *ClientFactory) NewAPIDefinitionsClient() *APIDefinitionsClient {
	return &APIDefinitionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAPIVersionsClient creates a new instance of APIVersionsClient.
func (c *ClientFactory) NewAPIVersionsClient() *APIVersionsClient {
	return &APIVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApisClient creates a new instance of ApisClient.
func (c *ClientFactory) NewApisClient() *ApisClient {
	return &ApisClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedServicesClient creates a new instance of DeletedServicesClient.
func (c *ClientFactory) NewDeletedServicesClient() *DeletedServicesClient {
	return &DeletedServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeploymentsClient creates a new instance of DeploymentsClient.
func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	return &DeploymentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient.
func (c *ClientFactory) NewEnvironmentsClient() *EnvironmentsClient {
	return &EnvironmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetadataSchemasClient creates a new instance of MetadataSchemasClient.
func (c *ClientFactory) NewMetadataSchemasClient() *MetadataSchemasClient {
	return &MetadataSchemasClient{
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

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	return &ServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	return &WorkspacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
