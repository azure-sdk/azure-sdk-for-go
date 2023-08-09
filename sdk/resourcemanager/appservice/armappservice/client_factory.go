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
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewCertificateOrdersClient() *CertificateOrdersClient {
	subClient, _ := NewCertificateOrdersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificateOrdersDiagnosticsClient() *CertificateOrdersDiagnosticsClient {
	subClient, _ := NewCertificateOrdersDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificateRegistrationProviderClient() *CertificateRegistrationProviderClient {
	subClient, _ := NewCertificateRegistrationProviderClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	subClient, _ := NewCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainerAppsClient() *ContainerAppsClient {
	subClient, _ := NewContainerAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainerAppsRevisionsClient() *ContainerAppsRevisionsClient {
	subClient, _ := NewContainerAppsRevisionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeletedWebAppsClient() *DeletedWebAppsClient {
	subClient, _ := NewDeletedWebAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiagnosticsClient() *DiagnosticsClient {
	subClient, _ := NewDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDomainRegistrationProviderClient() *DomainRegistrationProviderClient {
	subClient, _ := NewDomainRegistrationProviderClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDomainsClient() *DomainsClient {
	subClient, _ := NewDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEnvironmentsClient() *EnvironmentsClient {
	subClient, _ := NewEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGlobalClient() *GlobalClient {
	subClient, _ := NewGlobalClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewKubeEnvironmentsClient() *KubeEnvironmentsClient {
	subClient, _ := NewKubeEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPlansClient() *PlansClient {
	subClient, _ := NewPlansClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProviderClient() *ProviderClient {
	subClient, _ := NewProviderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRecommendationsClient() *RecommendationsClient {
	subClient, _ := NewRecommendationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourceHealthMetadataClient() *ResourceHealthMetadataClient {
	subClient, _ := NewResourceHealthMetadataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStaticSitesClient() *StaticSitesClient {
	subClient, _ := NewStaticSitesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTopLevelDomainsClient() *TopLevelDomainsClient {
	subClient, _ := NewTopLevelDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebAppsClient() *WebAppsClient {
	subClient, _ := NewWebAppsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebSiteManagementClient() *WebSiteManagementClient {
	subClient, _ := NewWebSiteManagementClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionRepetitionsClient() *WorkflowRunActionRepetitionsClient {
	subClient, _ := NewWorkflowRunActionRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionRepetitionsRequestHistoriesClient() *WorkflowRunActionRepetitionsRequestHistoriesClient {
	subClient, _ := NewWorkflowRunActionRepetitionsRequestHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionScopeRepetitionsClient() *WorkflowRunActionScopeRepetitionsClient {
	subClient, _ := NewWorkflowRunActionScopeRepetitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunActionsClient() *WorkflowRunActionsClient {
	subClient, _ := NewWorkflowRunActionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowRunsClient() *WorkflowRunsClient {
	subClient, _ := NewWorkflowRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowTriggerHistoriesClient() *WorkflowTriggerHistoriesClient {
	subClient, _ := NewWorkflowTriggerHistoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowTriggersClient() *WorkflowTriggersClient {
	subClient, _ := NewWorkflowTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowVersionsClient() *WorkflowVersionsClient {
	subClient, _ := NewWorkflowVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkflowsClient() *WorkflowsClient {
	subClient, _ := NewWorkflowsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
