//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armselfhelp

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	scope                   string
	diagnosticsResourceName string
	filter                  *string
	skiptoken               *string
	credential              azcore.TokenCredential
	options                 *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - scope - This is an extension resource provider and only resource level extension is supported at the moment.
//   - diagnosticsResourceName - Unique resource name for insight resources
//   - filter - Can be used to filter solutionIds by 'ProblemClassificationId'. The filter supports only 'and' and 'eq' operators.
//     Example: $filter=ProblemClassificationId eq '1ddda5b4-cf6c-4d4f-91ad-bc38ab0e811e'
//     and ProblemClassificationId eq '0a9673c2-7af6-4e19-90d3-4ee2461076d9'.
//   - skiptoken - Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a
//     nextLink element, the value of the nextLink element will include a skiptoken parameter that
//     specifies a starting point to use for subsequent calls.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(scope string, diagnosticsResourceName string, filter *string, skiptoken *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		scope: scope, diagnosticsResourceName: diagnosticsResourceName, filter: filter, skiptoken: skiptoken, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiagnosticsClient() *DiagnosticsClient {
	subClient, _ := NewDiagnosticsClient(c.scope, c.diagnosticsResourceName, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiscoverySolutionClient() *DiscoverySolutionClient {
	subClient, _ := NewDiscoverySolutionClient(c.scope, c.filter, c.skiptoken, c.credential, c.options)
	return subClient
}
