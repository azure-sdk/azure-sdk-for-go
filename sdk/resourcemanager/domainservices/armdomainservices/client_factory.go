// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

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
//   - subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
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

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDomainServiceOperationsClient creates a new instance of DomainServiceOperationsClient.
func (c *ClientFactory) NewDomainServiceOperationsClient() *DomainServiceOperationsClient {
	return &DomainServiceOperationsClient{
		internal: c.internal,
	}
}

// NewOuContainerClient creates a new instance of OuContainerClient.
func (c *ClientFactory) NewOuContainerClient() *OuContainerClient {
	return &OuContainerClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOuContainerOperationsClient creates a new instance of OuContainerOperationsClient.
func (c *ClientFactory) NewOuContainerOperationsClient() *OuContainerOperationsClient {
	return &OuContainerOperationsClient{
		internal: c.internal,
	}
}
