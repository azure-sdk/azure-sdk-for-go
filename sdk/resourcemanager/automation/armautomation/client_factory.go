//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
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

func (c *ClientFactory) NewAccountClient() *AccountClient {
	subClient, _ := NewAccountClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewActivityClient() *ActivityClient {
	subClient, _ := NewActivityClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAgentRegistrationInformationClient() *AgentRegistrationInformationClient {
	subClient, _ := NewAgentRegistrationInformationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificateClient() *CertificateClient {
	subClient, _ := NewCertificateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectionClient() *ConnectionClient {
	subClient, _ := NewConnectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectionTypeClient() *ConnectionTypeClient {
	subClient, _ := NewConnectionTypeClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCredentialClient() *CredentialClient {
	subClient, _ := NewCredentialClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeletedAutomationAccountsClient() *DeletedAutomationAccountsClient {
	subClient, _ := NewDeletedAutomationAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDscCompilationJobClient() *DscCompilationJobClient {
	subClient, _ := NewDscCompilationJobClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDscCompilationJobStreamClient() *DscCompilationJobStreamClient {
	subClient, _ := NewDscCompilationJobStreamClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDscConfigurationClient() *DscConfigurationClient {
	subClient, _ := NewDscConfigurationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDscNodeClient() *DscNodeClient {
	subClient, _ := NewDscNodeClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDscNodeConfigurationClient() *DscNodeConfigurationClient {
	subClient, _ := NewDscNodeConfigurationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFieldsClient() *FieldsClient {
	subClient, _ := NewFieldsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHybridRunbookWorkerGroupClient() *HybridRunbookWorkerGroupClient {
	subClient, _ := NewHybridRunbookWorkerGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHybridRunbookWorkersClient() *HybridRunbookWorkersClient {
	subClient, _ := NewHybridRunbookWorkersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobClient() *JobClient {
	subClient, _ := NewJobClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobScheduleClient() *JobScheduleClient {
	subClient, _ := NewJobScheduleClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobStreamClient() *JobStreamClient {
	subClient, _ := NewJobStreamClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewKeysClient() *KeysClient {
	subClient, _ := NewKeysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLinkedWorkspaceClient() *LinkedWorkspaceClient {
	subClient, _ := NewLinkedWorkspaceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewModuleClient() *ModuleClient {
	subClient, _ := NewModuleClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNodeCountInformationClient() *NodeCountInformationClient {
	subClient, _ := NewNodeCountInformationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNodeReportsClient() *NodeReportsClient {
	subClient, _ := NewNodeReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewObjectDataTypesClient() *ObjectDataTypesClient {
	subClient, _ := NewObjectDataTypesClient(c.subscriptionID, c.credential, c.options)
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

func (c *ClientFactory) NewPython2PackageClient() *Python2PackageClient {
	subClient, _ := NewPython2PackageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPython3PackageClient() *Python3PackageClient {
	subClient, _ := NewPython3PackageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRunbookClient() *RunbookClient {
	subClient, _ := NewRunbookClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRunbookDraftClient() *RunbookDraftClient {
	subClient, _ := NewRunbookDraftClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScheduleClient() *ScheduleClient {
	subClient, _ := NewScheduleClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSoftwareUpdateConfigurationMachineRunsClient() *SoftwareUpdateConfigurationMachineRunsClient {
	subClient, _ := NewSoftwareUpdateConfigurationMachineRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSoftwareUpdateConfigurationRunsClient() *SoftwareUpdateConfigurationRunsClient {
	subClient, _ := NewSoftwareUpdateConfigurationRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSoftwareUpdateConfigurationsClient() *SoftwareUpdateConfigurationsClient {
	subClient, _ := NewSoftwareUpdateConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSourceControlClient() *SourceControlClient {
	subClient, _ := NewSourceControlClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSourceControlSyncJobClient() *SourceControlSyncJobClient {
	subClient, _ := NewSourceControlSyncJobClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSourceControlSyncJobStreamsClient() *SourceControlSyncJobStreamsClient {
	subClient, _ := NewSourceControlSyncJobStreamsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStatisticsClient() *StatisticsClient {
	subClient, _ := NewStatisticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTestJobClient() *TestJobClient {
	subClient, _ := NewTestJobClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTestJobStreamsClient() *TestJobStreamsClient {
	subClient, _ := NewTestJobStreamsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVariableClient() *VariableClient {
	subClient, _ := NewVariableClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWatcherClient() *WatcherClient {
	subClient, _ := NewWatcherClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebhookClient() *WebhookClient {
	subClient, _ := NewWebhookClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
