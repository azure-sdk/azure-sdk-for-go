//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

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

// NewDomainsClient creates a new instance of DomainsClient.
func (c *ClientFactory) NewDomainsClient() *DomainsClient {
	return &DomainsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEmailServicesClient creates a new instance of EmailServicesClient.
func (c *ClientFactory) NewEmailServicesClient() *EmailServicesClient {
	return &EmailServicesClient{
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

// NewSenderUsernamesClient creates a new instance of SenderUsernamesClient.
func (c *ClientFactory) NewSenderUsernamesClient() *SenderUsernamesClient {
	return &SenderUsernamesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	return &ServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSuppressionListAddressesClient creates a new instance of SuppressionListAddressesClient.
func (c *ClientFactory) NewSuppressionListAddressesClient() *SuppressionListAddressesClient {
	return &SuppressionListAddressesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSuppressionListsClient creates a new instance of SuppressionListsClient.
func (c *ClientFactory) NewSuppressionListsClient() *SuppressionListsClient {
	return &SuppressionListsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
