//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapicenter

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID     string
	workspaceName      string
	apiName            string
	versionName        string
	definitionName     string
	deploymentName     string
	environmentName    string
	metadataSchemaName string
	filter             *string
	credential         azcore.TokenCredential
	options            *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - deploymentName - The name of the API deployment.
//   - environmentName - The name of the environment.
//   - metadataSchemaName - The name of the metadata schema.
//   - filter - OData filter parameter.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, workspaceName string, apiName string, versionName string, definitionName string, deploymentName string, environmentName string, metadataSchemaName string, filter *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, workspaceName: workspaceName, apiName: apiName, versionName: versionName, definitionName: definitionName, deploymentName: deploymentName, environmentName: environmentName, metadataSchemaName: metadataSchemaName, filter: filter, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAPIDefinitionsClient() *APIDefinitionsClient {
	subClient, _ := NewAPIDefinitionsClient(c.subscriptionID, c.workspaceName, c.apiName, c.versionName, c.definitionName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIVersionsClient() *APIVersionsClient {
	subClient, _ := NewAPIVersionsClient(c.subscriptionID, c.workspaceName, c.apiName, c.versionName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewApisClient() *ApisClient {
	subClient, _ := NewApisClient(c.subscriptionID, c.workspaceName, c.apiName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	subClient, _ := NewDeploymentsClient(c.subscriptionID, c.workspaceName, c.apiName, c.deploymentName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEnvironmentsClient() *EnvironmentsClient {
	subClient, _ := NewEnvironmentsClient(c.subscriptionID, c.workspaceName, c.environmentName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetadataSchemasClient() *MetadataSchemasClient {
	subClient, _ := NewMetadataSchemasClient(c.subscriptionID, c.metadataSchemaName, c.filter, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	subClient, _ := NewWorkspacesClient(c.subscriptionID, c.workspaceName, c.filter, c.credential, c.options)
	return subClient
}
