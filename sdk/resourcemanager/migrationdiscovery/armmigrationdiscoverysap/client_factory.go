// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationdiscoverysap

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

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewSapDiscoverySitesClient creates a new instance of SapDiscoverySitesClient.
func (c *ClientFactory) NewSapDiscoverySitesClient() *SapDiscoverySitesClient {
	return &SapDiscoverySitesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSapInstancesClient creates a new instance of SapInstancesClient.
func (c *ClientFactory) NewSapInstancesClient() *SapInstancesClient {
	return &SapInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerInstancesClient creates a new instance of ServerInstancesClient.
func (c *ClientFactory) NewServerInstancesClient() *ServerInstancesClient {
	return &ServerInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
