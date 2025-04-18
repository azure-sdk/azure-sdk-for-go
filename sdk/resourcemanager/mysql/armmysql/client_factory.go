// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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
//   - subscriptionID - The ID of the target subscription.
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

// NewAdvisorsClient creates a new instance of AdvisorsClient.
func (c *ClientFactory) NewAdvisorsClient() *AdvisorsClient {
	return &AdvisorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient.
func (c *ClientFactory) NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient {
	return &CheckNameAvailabilityClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient.
func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	return &ConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDatabasesClient creates a new instance of DatabasesClient.
func (c *ClientFactory) NewDatabasesClient() *DatabasesClient {
	return &DatabasesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient.
func (c *ClientFactory) NewFirewallRulesClient() *FirewallRulesClient {
	return &FirewallRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocationBasedPerformanceTierClient creates a new instance of LocationBasedPerformanceTierClient.
func (c *ClientFactory) NewLocationBasedPerformanceTierClient() *LocationBasedPerformanceTierClient {
	return &LocationBasedPerformanceTierClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocationBasedRecommendedActionSessionsOperationStatusClient creates a new instance of LocationBasedRecommendedActionSessionsOperationStatusClient.
func (c *ClientFactory) NewLocationBasedRecommendedActionSessionsOperationStatusClient() *LocationBasedRecommendedActionSessionsOperationStatusClient {
	return &LocationBasedRecommendedActionSessionsOperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLocationBasedRecommendedActionSessionsResultClient creates a new instance of LocationBasedRecommendedActionSessionsResultClient.
func (c *ClientFactory) NewLocationBasedRecommendedActionSessionsResultClient() *LocationBasedRecommendedActionSessionsResultClient {
	return &LocationBasedRecommendedActionSessionsResultClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLogFilesClient creates a new instance of LogFilesClient.
func (c *ClientFactory) NewLogFilesClient() *LogFilesClient {
	return &LogFilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagementClient creates a new instance of ManagementClient.
func (c *ClientFactory) NewManagementClient() *ManagementClient {
	return &ManagementClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	return &PrivateLinkResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewQueryTextsClient creates a new instance of QueryTextsClient.
func (c *ClientFactory) NewQueryTextsClient() *QueryTextsClient {
	return &QueryTextsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRecommendedActionsClient creates a new instance of RecommendedActionsClient.
func (c *ClientFactory) NewRecommendedActionsClient() *RecommendedActionsClient {
	return &RecommendedActionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRecoverableServersClient creates a new instance of RecoverableServersClient.
func (c *ClientFactory) NewRecoverableServersClient() *RecoverableServersClient {
	return &RecoverableServersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewReplicasClient creates a new instance of ReplicasClient.
func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	return &ReplicasClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerAdministratorsClient creates a new instance of ServerAdministratorsClient.
func (c *ClientFactory) NewServerAdministratorsClient() *ServerAdministratorsClient {
	return &ServerAdministratorsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerBasedPerformanceTierClient creates a new instance of ServerBasedPerformanceTierClient.
func (c *ClientFactory) NewServerBasedPerformanceTierClient() *ServerBasedPerformanceTierClient {
	return &ServerBasedPerformanceTierClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerKeysClient creates a new instance of ServerKeysClient.
func (c *ClientFactory) NewServerKeysClient() *ServerKeysClient {
	return &ServerKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerParametersClient creates a new instance of ServerParametersClient.
func (c *ClientFactory) NewServerParametersClient() *ServerParametersClient {
	return &ServerParametersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServerSecurityAlertPoliciesClient creates a new instance of ServerSecurityAlertPoliciesClient.
func (c *ClientFactory) NewServerSecurityAlertPoliciesClient() *ServerSecurityAlertPoliciesClient {
	return &ServerSecurityAlertPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewServersClient creates a new instance of ServersClient.
func (c *ClientFactory) NewServersClient() *ServersClient {
	return &ServersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTopQueryStatisticsClient creates a new instance of TopQueryStatisticsClient.
func (c *ClientFactory) NewTopQueryStatisticsClient() *TopQueryStatisticsClient {
	return &TopQueryStatisticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient.
func (c *ClientFactory) NewVirtualNetworkRulesClient() *VirtualNetworkRulesClient {
	return &VirtualNetworkRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewWaitStatisticsClient creates a new instance of WaitStatisticsClient.
func (c *ClientFactory) NewWaitStatisticsClient() *WaitStatisticsClient {
	return &WaitStatisticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
