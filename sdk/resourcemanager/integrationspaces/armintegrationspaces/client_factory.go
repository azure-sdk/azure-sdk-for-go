//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armintegrationspaces

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

// NewApplicationResourcesClient creates a new instance of ApplicationResourcesClient.
func (c *ClientFactory) NewApplicationResourcesClient() *ApplicationResourcesClient {
	return &ApplicationResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewApplicationsClient creates a new instance of ApplicationsClient.
func (c *ClientFactory) NewApplicationsClient() *ApplicationsClient {
	return &ApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBusinessProcessVersionsClient creates a new instance of BusinessProcessVersionsClient.
func (c *ClientFactory) NewBusinessProcessVersionsClient() *BusinessProcessVersionsClient {
	return &BusinessProcessVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBusinessProcessesClient creates a new instance of BusinessProcessesClient.
func (c *ClientFactory) NewBusinessProcessesClient() *BusinessProcessesClient {
	return &BusinessProcessesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewInfrastructureResourcesClient creates a new instance of InfrastructureResourcesClient.
func (c *ClientFactory) NewInfrastructureResourcesClient() *InfrastructureResourcesClient {
	return &InfrastructureResourcesClient{
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

// NewSpacesClient creates a new instance of SpacesClient.
func (c *ClientFactory) NewSpacesClient() *SpacesClient {
	return &SpacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
