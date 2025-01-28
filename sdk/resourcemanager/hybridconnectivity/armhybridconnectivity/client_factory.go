//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridconnectivity

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

// NewEndpointsClient creates a new instance of EndpointsClient.
func (c *ClientFactory) NewEndpointsClient() *EndpointsClient {
	return &EndpointsClient{
		internal: c.internal,
	}
}

// NewGenerateAwsTemplateClient creates a new instance of GenerateAwsTemplateClient.
func (c *ClientFactory) NewGenerateAwsTemplateClient() *GenerateAwsTemplateClient {
	return &GenerateAwsTemplateClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewInventoryClient creates a new instance of InventoryClient.
func (c *ClientFactory) NewInventoryClient() *InventoryClient {
	return &InventoryClient{
		internal: c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPublicCloudConnectorsClient creates a new instance of PublicCloudConnectorsClient.
func (c *ClientFactory) NewPublicCloudConnectorsClient() *PublicCloudConnectorsClient {
	return &PublicCloudConnectorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServiceConfigurationsClient creates a new instance of ServiceConfigurationsClient.
func (c *ClientFactory) NewServiceConfigurationsClient() *ServiceConfigurationsClient {
	return &ServiceConfigurationsClient{
		internal: c.internal,
	}
}

// NewSolutionConfigurationsClient creates a new instance of SolutionConfigurationsClient.
func (c *ClientFactory) NewSolutionConfigurationsClient() *SolutionConfigurationsClient {
	return &SolutionConfigurationsClient{
		internal: c.internal,
	}
}

// NewSolutionTypesClient creates a new instance of SolutionTypesClient.
func (c *ClientFactory) NewSolutionTypesClient() *SolutionTypesClient {
	return &SolutionTypesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
