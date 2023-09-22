//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

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

func (c *ClientFactory) NewClustersClient() *ClustersClient {
	subClient, _ := NewClustersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFunctionsClient() *FunctionsClient {
	subClient, _ := NewFunctionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInputsClient() *InputsClient {
	subClient, _ := NewInputsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOutputsClient() *OutputsClient {
	subClient, _ := NewOutputsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointsClient() *PrivateEndpointsClient {
	subClient, _ := NewPrivateEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStreamingJobsClient() *StreamingJobsClient {
	subClient, _ := NewStreamingJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSubscriptionsClient() *SubscriptionsClient {
	subClient, _ := NewSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTransformationsClient() *TransformationsClient {
	subClient, _ := NewTransformationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
