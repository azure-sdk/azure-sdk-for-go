// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// NewExtensionsClient creates a new instance of ExtensionsClient.
func (c *ClientFactory) NewExtensionsClient() *ExtensionsClient {
	return &ExtensionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFluxConfigOperationStatusClient creates a new instance of FluxConfigOperationStatusClient.
func (c *ClientFactory) NewFluxConfigOperationStatusClient() *FluxConfigOperationStatusClient {
	return &FluxConfigOperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFluxConfigurationsClient creates a new instance of FluxConfigurationsClient.
func (c *ClientFactory) NewFluxConfigurationsClient() *FluxConfigurationsClient {
	return &FluxConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusClient creates a new instance of OperationStatusClient.
func (c *ClientFactory) NewOperationStatusClient() *OperationStatusClient {
	return &OperationStatusClient{
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

// NewSourceControlConfigurationsClient creates a new instance of SourceControlConfigurationsClient.
func (c *ClientFactory) NewSourceControlConfigurationsClient() *SourceControlConfigurationsClient {
	return &SourceControlConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
