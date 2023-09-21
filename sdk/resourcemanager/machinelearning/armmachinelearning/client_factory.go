//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

func (c *ClientFactory) NewBatchDeploymentsClient() *BatchDeploymentsClient {
	subClient, _ := NewBatchDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBatchEndpointsClient() *BatchEndpointsClient {
	subClient, _ := NewBatchEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCodeContainersClient() *CodeContainersClient {
	subClient, _ := NewCodeContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCodeVersionsClient() *CodeVersionsClient {
	subClient, _ := NewCodeVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentContainersClient() *ComponentContainersClient {
	subClient, _ := NewComponentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentVersionsClient() *ComponentVersionsClient {
	subClient, _ := NewComponentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComputeClient() *ComputeClient {
	subClient, _ := NewComputeClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataContainersClient() *DataContainersClient {
	subClient, _ := NewDataContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataVersionsClient() *DataVersionsClient {
	subClient, _ := NewDataVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatastoresClient() *DatastoresClient {
	subClient, _ := NewDatastoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEnvironmentContainersClient() *EnvironmentContainersClient {
	subClient, _ := NewEnvironmentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEnvironmentVersionsClient() *EnvironmentVersionsClient {
	subClient, _ := NewEnvironmentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedNetworkProvisionsClient() *ManagedNetworkProvisionsClient {
	subClient, _ := NewManagedNetworkProvisionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedNetworkSettingsRuleClient() *ManagedNetworkSettingsRuleClient {
	subClient, _ := NewManagedNetworkSettingsRuleClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewModelContainersClient() *ModelContainersClient {
	subClient, _ := NewModelContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewModelVersionsClient() *ModelVersionsClient {
	subClient, _ := NewModelVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOnlineDeploymentsClient() *OnlineDeploymentsClient {
	subClient, _ := NewOnlineDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOnlineEndpointsClient() *OnlineEndpointsClient {
	subClient, _ := NewOnlineEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewQuotasClient() *QuotasClient {
	subClient, _ := NewQuotasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistriesClient() *RegistriesClient {
	subClient, _ := NewRegistriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryCodeContainersClient() *RegistryCodeContainersClient {
	subClient, _ := NewRegistryCodeContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryCodeVersionsClient() *RegistryCodeVersionsClient {
	subClient, _ := NewRegistryCodeVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryComponentContainersClient() *RegistryComponentContainersClient {
	subClient, _ := NewRegistryComponentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryComponentVersionsClient() *RegistryComponentVersionsClient {
	subClient, _ := NewRegistryComponentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryDataContainersClient() *RegistryDataContainersClient {
	subClient, _ := NewRegistryDataContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryDataReferencesClient() *RegistryDataReferencesClient {
	subClient, _ := NewRegistryDataReferencesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryDataVersionsClient() *RegistryDataVersionsClient {
	subClient, _ := NewRegistryDataVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryEnvironmentContainersClient() *RegistryEnvironmentContainersClient {
	subClient, _ := NewRegistryEnvironmentContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryEnvironmentVersionsClient() *RegistryEnvironmentVersionsClient {
	subClient, _ := NewRegistryEnvironmentVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryModelContainersClient() *RegistryModelContainersClient {
	subClient, _ := NewRegistryModelContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistryModelVersionsClient() *RegistryModelVersionsClient {
	subClient, _ := NewRegistryModelVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSchedulesClient() *SchedulesClient {
	subClient, _ := NewSchedulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachineSizesClient() *VirtualMachineSizesClient {
	subClient, _ := NewVirtualMachineSizesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceConnectionsClient() *WorkspaceConnectionsClient {
	subClient, _ := NewWorkspaceConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceFeaturesClient() *WorkspaceFeaturesClient {
	subClient, _ := NewWorkspaceFeaturesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	subClient, _ := NewWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
