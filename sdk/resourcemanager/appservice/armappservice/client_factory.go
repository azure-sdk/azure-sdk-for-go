//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewCertificatesClient creates a new instance of CertificatesClient.
func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	subClient, _ := NewCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewContainerAppsClient creates a new instance of ContainerAppsClient.
func (c *ClientFactory) NewContainerAppsClient() *ContainerAppsClient {
	subClient, _ := NewContainerAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewContainerAppsRevisionsClient creates a new instance of ContainerAppsRevisionsClient.
func (c *ClientFactory) NewContainerAppsRevisionsClient() *ContainerAppsRevisionsClient {
	subClient, _ := NewContainerAppsRevisionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeletedWebAppsClient creates a new instance of DeletedWebAppsClient.
func (c *ClientFactory) NewDeletedWebAppsClient() *DeletedWebAppsClient {
	subClient, _ := NewDeletedWebAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDiagnosticsClient creates a new instance of DiagnosticsClient.
func (c *ClientFactory) NewDiagnosticsClient() *DiagnosticsClient {
	subClient, _ := NewDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient.
func (c *ClientFactory) NewEnvironmentsClient() *EnvironmentsClient {
	subClient, _ := NewEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewGlobalClient creates a new instance of GlobalClient.
func (c *ClientFactory) NewGlobalClient() *GlobalClient {
	subClient, _ := NewGlobalClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewKubeEnvironmentsClient creates a new instance of KubeEnvironmentsClient.
func (c *ClientFactory) NewKubeEnvironmentsClient() *KubeEnvironmentsClient {
	subClient, _ := NewKubeEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPlansClient creates a new instance of PlansClient.
func (c *ClientFactory) NewPlansClient() *PlansClient {
	subClient, _ := NewPlansClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewProviderClient creates a new instance of ProviderClient.
func (c *ClientFactory) NewProviderClient() *ProviderClient {
	subClient, _ := NewProviderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRecommendationsClient creates a new instance of RecommendationsClient.
func (c *ClientFactory) NewRecommendationsClient() *RecommendationsClient {
	subClient, _ := NewRecommendationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewResourceHealthMetadataClient creates a new instance of ResourceHealthMetadataClient.
func (c *ClientFactory) NewResourceHealthMetadataClient() *ResourceHealthMetadataClient {
	subClient, _ := NewResourceHealthMetadataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewStaticSitesClient creates a new instance of StaticSitesClient.
func (c *ClientFactory) NewStaticSitesClient() *StaticSitesClient {
	subClient, _ := NewStaticSitesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWebAppsClient creates a new instance of WebAppsClient.
func (c *ClientFactory) NewWebAppsClient() *WebAppsClient {
	subClient, _ := NewWebAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWebSiteManagementClient creates a new instance of WebSiteManagementClient.
func (c *ClientFactory) NewWebSiteManagementClient() *WebSiteManagementClient {
	subClient, _ := NewWebSiteManagementClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowRunActionRepetitionsClient creates a new instance of WorkflowRunActionRepetitionsClient.
func (c *ClientFactory) NewWorkflowRunActionRepetitionsClient() *WorkflowRunActionRepetitionsClient {
	subClient, _ := NewWorkflowRunActionRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowRunActionRepetitionsRequestHistoriesClient creates a new instance of WorkflowRunActionRepetitionsRequestHistoriesClient.
func (c *ClientFactory) NewWorkflowRunActionRepetitionsRequestHistoriesClient() *WorkflowRunActionRepetitionsRequestHistoriesClient {
	subClient, _ := NewWorkflowRunActionRepetitionsRequestHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowRunActionScopeRepetitionsClient creates a new instance of WorkflowRunActionScopeRepetitionsClient.
func (c *ClientFactory) NewWorkflowRunActionScopeRepetitionsClient() *WorkflowRunActionScopeRepetitionsClient {
	subClient, _ := NewWorkflowRunActionScopeRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowRunActionsClient creates a new instance of WorkflowRunActionsClient.
func (c *ClientFactory) NewWorkflowRunActionsClient() *WorkflowRunActionsClient {
	subClient, _ := NewWorkflowRunActionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowRunsClient creates a new instance of WorkflowRunsClient.
func (c *ClientFactory) NewWorkflowRunsClient() *WorkflowRunsClient {
	subClient, _ := NewWorkflowRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowTriggerHistoriesClient creates a new instance of WorkflowTriggerHistoriesClient.
func (c *ClientFactory) NewWorkflowTriggerHistoriesClient() *WorkflowTriggerHistoriesClient {
	subClient, _ := NewWorkflowTriggerHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowTriggersClient creates a new instance of WorkflowTriggersClient.
func (c *ClientFactory) NewWorkflowTriggersClient() *WorkflowTriggersClient {
	subClient, _ := NewWorkflowTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowVersionsClient creates a new instance of WorkflowVersionsClient.
func (c *ClientFactory) NewWorkflowVersionsClient() *WorkflowVersionsClient {
	subClient, _ := NewWorkflowVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowsClient creates a new instance of WorkflowsClient.
func (c *ClientFactory) NewWorkflowsClient() *WorkflowsClient {
	subClient, _ := NewWorkflowsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
