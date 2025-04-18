// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armspringappdiscovery

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

// NewErrorSummariesClient creates a new instance of ErrorSummariesClient.
func (c *ClientFactory) NewErrorSummariesClient() *ErrorSummariesClient {
	return &ErrorSummariesClient{
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

// NewSpringbootappsClient creates a new instance of SpringbootappsClient.
func (c *ClientFactory) NewSpringbootappsClient() *SpringbootappsClient {
	return &SpringbootappsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSpringbootserversClient creates a new instance of SpringbootserversClient.
func (c *ClientFactory) NewSpringbootserversClient() *SpringbootserversClient {
	return &SpringbootserversClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSpringbootsitesClient creates a new instance of SpringbootsitesClient.
func (c *ClientFactory) NewSpringbootsitesClient() *SpringbootsitesClient {
	return &SpringbootsitesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSummariesClient creates a new instance of SummariesClient.
func (c *ClientFactory) NewSummariesClient() *SummariesClient {
	return &SummariesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
