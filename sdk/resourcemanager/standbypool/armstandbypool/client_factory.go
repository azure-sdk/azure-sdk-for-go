// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armstandbypool

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

// NewStandbyContainerGroupPoolsClient creates a new instance of StandbyContainerGroupPoolsClient.
func (c *ClientFactory) NewStandbyContainerGroupPoolsClient() *StandbyContainerGroupPoolsClient {
	return &StandbyContainerGroupPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewStandbyVirtualMachinePoolsClient creates a new instance of StandbyVirtualMachinePoolsClient.
func (c *ClientFactory) NewStandbyVirtualMachinePoolsClient() *StandbyVirtualMachinePoolsClient {
	return &StandbyVirtualMachinePoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewStandbyVirtualMachinesClient creates a new instance of StandbyVirtualMachinesClient.
func (c *ClientFactory) NewStandbyVirtualMachinesClient() *StandbyVirtualMachinesClient {
	return &StandbyVirtualMachinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
