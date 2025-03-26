// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armworkloadssapvirtualinstance

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

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewSAPApplicationServerInstancesClient creates a new instance of SAPApplicationServerInstancesClient.
func (c *ClientFactory) NewSAPApplicationServerInstancesClient() *SAPApplicationServerInstancesClient {
	return &SAPApplicationServerInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSAPCentralServerInstancesClient creates a new instance of SAPCentralServerInstancesClient.
func (c *ClientFactory) NewSAPCentralServerInstancesClient() *SAPCentralServerInstancesClient {
	return &SAPCentralServerInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSAPDatabaseInstancesClient creates a new instance of SAPDatabaseInstancesClient.
func (c *ClientFactory) NewSAPDatabaseInstancesClient() *SAPDatabaseInstancesClient {
	return &SAPDatabaseInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSAPVirtualInstancesClient creates a new instance of SAPVirtualInstancesClient.
func (c *ClientFactory) NewSAPVirtualInstancesClient() *SAPVirtualInstancesClient {
	return &SAPVirtualInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
