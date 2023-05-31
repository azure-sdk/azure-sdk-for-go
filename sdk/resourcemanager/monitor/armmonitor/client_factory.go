//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

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

func (c *ClientFactory) NewAzureMonitorWorkspacesClient() *AzureMonitorWorkspacesClient {
	subClient, _ := NewAzureMonitorWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsForMonitorClient() *OperationsForMonitorClient {
	subClient, _ := NewOperationsForMonitorClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAutoscaleSettingsClient() *AutoscaleSettingsClient {
	subClient, _ := NewAutoscaleSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPredictiveMetricClient() *PredictiveMetricClient {
	subClient, _ := NewPredictiveMetricClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertRuleIncidentsClient() *AlertRuleIncidentsClient {
	subClient, _ := NewAlertRuleIncidentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertRulesClient() *AlertRulesClient {
	subClient, _ := NewAlertRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLogProfilesClient() *LogProfilesClient {
	subClient, _ := NewLogProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiagnosticSettingsClient() *DiagnosticSettingsClient {
	subClient, _ := NewDiagnosticSettingsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiagnosticSettingsCategoryClient() *DiagnosticSettingsCategoryClient {
	subClient, _ := NewDiagnosticSettingsCategoryClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewActionGroupsClient() *ActionGroupsClient {
	subClient, _ := NewActionGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantActionGroupsClient() *TenantActionGroupsClient {
	subClient, _ := NewTenantActionGroupsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewActivityLogsClient() *ActivityLogsClient {
	subClient, _ := NewActivityLogsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEventCategoriesClient() *EventCategoriesClient {
	subClient, _ := NewEventCategoriesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantActivityLogsClient() *TenantActivityLogsClient {
	subClient, _ := NewTenantActivityLogsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetricDefinitionsClient() *MetricDefinitionsClient {
	subClient, _ := NewMetricDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetricsClient() *MetricsClient {
	subClient, _ := NewMetricsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBaselinesClient() *BaselinesClient {
	subClient, _ := NewBaselinesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetricAlertsClient() *MetricAlertsClient {
	subClient, _ := NewMetricAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetricAlertsStatusClient() *MetricAlertsStatusClient {
	subClient, _ := NewMetricAlertsStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScheduledQueryRulesClient() *ScheduledQueryRulesClient {
	subClient, _ := NewScheduledQueryRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetricNamespacesClient() *MetricNamespacesClient {
	subClient, _ := NewMetricNamespacesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVMInsightsClient() *VMInsightsClient {
	subClient, _ := NewVMInsightsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkScopesClient() *PrivateLinkScopesClient {
	subClient, _ := NewPrivateLinkScopesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkScopeOperationStatusClient() *PrivateLinkScopeOperationStatusClient {
	subClient, _ := NewPrivateLinkScopeOperationStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkScopedResourcesClient() *PrivateLinkScopedResourcesClient {
	subClient, _ := NewPrivateLinkScopedResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewActivityLogAlertsClient() *ActivityLogAlertsClient {
	subClient, _ := NewActivityLogAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataCollectionEndpointsClient() *DataCollectionEndpointsClient {
	subClient, _ := NewDataCollectionEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataCollectionRuleAssociationsClient() *DataCollectionRuleAssociationsClient {
	subClient, _ := NewDataCollectionRuleAssociationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataCollectionRulesClient() *DataCollectionRulesClient {
	subClient, _ := NewDataCollectionRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
