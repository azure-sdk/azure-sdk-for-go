//go:build go1.18
// +build go1.18

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

func (c *ClientFactory) NewSpringbootsitesClient() *SpringbootsitesClient {
	subClient, _ := NewSpringbootsitesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSpringbootserversClient() *SpringbootserversClient {
	subClient, _ := NewSpringbootserversClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSpringbootappsClient() *SpringbootappsClient {
	subClient, _ := NewSpringbootappsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSummariesClient() *SummariesClient {
	subClient, _ := NewSummariesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewErrorSummariesClient() *ErrorSummariesClient {
	subClient, _ := NewErrorSummariesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
