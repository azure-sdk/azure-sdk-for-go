//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

func (c *ClientFactory) NewContainerAppsAuthConfigsClient() *ContainerAppsAuthConfigsClient {
	subClient, _ := NewContainerAppsAuthConfigsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAvailableWorkloadProfilesClient() *AvailableWorkloadProfilesClient {
	subClient, _ := NewAvailableWorkloadProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBillingMetersClient() *BillingMetersClient {
	subClient, _ := NewBillingMetersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBuildersClient() *BuildersClient {
	subClient, _ := NewBuildersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBuildsClient() *BuildsClient {
	subClient, _ := NewBuildsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectedEnvironmentsClient() *ConnectedEnvironmentsClient {
	subClient, _ := NewConnectedEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectedEnvironmentsCertificatesClient() *ConnectedEnvironmentsCertificatesClient {
	subClient, _ := NewConnectedEnvironmentsCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectedEnvironmentsDaprComponentsClient() *ConnectedEnvironmentsDaprComponentsClient {
	subClient, _ := NewConnectedEnvironmentsDaprComponentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectedEnvironmentsStoragesClient() *ConnectedEnvironmentsStoragesClient {
	subClient, _ := NewConnectedEnvironmentsStoragesClient(c.subscriptionID, c.credential, c.options)
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

func (c *ClientFactory) NewContainerAppsRevisionReplicasClient() *ContainerAppsRevisionReplicasClient {
	subClient, _ := NewContainerAppsRevisionReplicasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainerAppsDiagnosticsClient() *ContainerAppsDiagnosticsClient {
	subClient, _ := NewContainerAppsDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedEnvironmentDiagnosticsClient() *ManagedEnvironmentDiagnosticsClient {
	subClient, _ := NewManagedEnvironmentDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedEnvironmentsDiagnosticsClient() *ManagedEnvironmentsDiagnosticsClient {
	subClient, _ := NewManagedEnvironmentsDiagnosticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsExecutionsClient() *JobsExecutionsClient {
	subClient, _ := NewJobsExecutionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainerAppsAPIClient() *ContainerAppsAPIClient {
	subClient, _ := NewContainerAppsAPIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedEnvironmentsClient() *ManagedEnvironmentsClient {
	subClient, _ := NewManagedEnvironmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	subClient, _ := NewCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedCertificatesClient() *ManagedCertificatesClient {
	subClient, _ := NewManagedCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNamespacesClient() *NamespacesClient {
	subClient, _ := NewNamespacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDaprComponentsClient() *DaprComponentsClient {
	subClient, _ := NewDaprComponentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedEnvironmentsStoragesClient() *ManagedEnvironmentsStoragesClient {
	subClient, _ := NewManagedEnvironmentsStoragesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainerAppsSourceControlsClient() *ContainerAppsSourceControlsClient {
	subClient, _ := NewContainerAppsSourceControlsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
