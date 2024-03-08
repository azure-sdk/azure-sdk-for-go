//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// NewActionGroupsClient creates a new instance of ActionGroupsClient.
func (c *ClientFactory) NewActionGroupsClient() *ActionGroupsClient {
	return &ActionGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewActivityLogAlertsClient creates a new instance of ActivityLogAlertsClient.
func (c *ClientFactory) NewActivityLogAlertsClient() *ActivityLogAlertsClient {
	return &ActivityLogAlertsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewActivityLogsClient creates a new instance of ActivityLogsClient.
func (c *ClientFactory) NewActivityLogsClient() *ActivityLogsClient {
	return &ActivityLogsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAlertRuleIncidentsClient creates a new instance of AlertRuleIncidentsClient.
func (c *ClientFactory) NewAlertRuleIncidentsClient() *AlertRuleIncidentsClient {
	return &AlertRuleIncidentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAlertRulesClient creates a new instance of AlertRulesClient.
func (c *ClientFactory) NewAlertRulesClient() *AlertRulesClient {
	return &AlertRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAutoscaleSettingsClient creates a new instance of AutoscaleSettingsClient.
func (c *ClientFactory) NewAutoscaleSettingsClient() *AutoscaleSettingsClient {
	return &AutoscaleSettingsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAzureMonitorWorkspacesClient creates a new instance of AzureMonitorWorkspacesClient.
func (c *ClientFactory) NewAzureMonitorWorkspacesClient() *AzureMonitorWorkspacesClient {
	return &AzureMonitorWorkspacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBaselinesClient creates a new instance of BaselinesClient.
func (c *ClientFactory) NewBaselinesClient() *BaselinesClient {
	return &BaselinesClient{
		internal: c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		internal: c.internal,
	}
}

// NewDataCollectionEndpointsClient creates a new instance of DataCollectionEndpointsClient.
func (c *ClientFactory) NewDataCollectionEndpointsClient() *DataCollectionEndpointsClient {
	return &DataCollectionEndpointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDataCollectionRuleAssociationsClient creates a new instance of DataCollectionRuleAssociationsClient.
func (c *ClientFactory) NewDataCollectionRuleAssociationsClient() *DataCollectionRuleAssociationsClient {
	return &DataCollectionRuleAssociationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDataCollectionRulesClient creates a new instance of DataCollectionRulesClient.
func (c *ClientFactory) NewDataCollectionRulesClient() *DataCollectionRulesClient {
	return &DataCollectionRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiagnosticSettingsCategoryClient creates a new instance of DiagnosticSettingsCategoryClient.
func (c *ClientFactory) NewDiagnosticSettingsCategoryClient() *DiagnosticSettingsCategoryClient {
	return &DiagnosticSettingsCategoryClient{
		internal: c.internal,
	}
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient.
func (c *ClientFactory) NewDiagnosticSettingsClient() *DiagnosticSettingsClient {
	return &DiagnosticSettingsClient{
		internal: c.internal,
	}
}

// NewEventCategoriesClient creates a new instance of EventCategoriesClient.
func (c *ClientFactory) NewEventCategoriesClient() *EventCategoriesClient {
	return &EventCategoriesClient{
		internal: c.internal,
	}
}

// NewLogProfilesClient creates a new instance of LogProfilesClient.
func (c *ClientFactory) NewLogProfilesClient() *LogProfilesClient {
	return &LogProfilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetricAlertsClient creates a new instance of MetricAlertsClient.
func (c *ClientFactory) NewMetricAlertsClient() *MetricAlertsClient {
	return &MetricAlertsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetricAlertsStatusClient creates a new instance of MetricAlertsStatusClient.
func (c *ClientFactory) NewMetricAlertsStatusClient() *MetricAlertsStatusClient {
	return &MetricAlertsStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetricDefinitionsClient creates a new instance of MetricDefinitionsClient.
func (c *ClientFactory) NewMetricDefinitionsClient() *MetricDefinitionsClient {
	return &MetricDefinitionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewMetricNamespacesClient creates a new instance of MetricNamespacesClient.
func (c *ClientFactory) NewMetricNamespacesClient() *MetricNamespacesClient {
	return &MetricNamespacesClient{
		internal: c.internal,
	}
}

// NewMetricsClient creates a new instance of MetricsClient.
func (c *ClientFactory) NewMetricsClient() *MetricsClient {
	return &MetricsClient{
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

// NewOperationsForMonitorClient creates a new instance of OperationsForMonitorClient.
func (c *ClientFactory) NewOperationsForMonitorClient() *OperationsForMonitorClient {
	return &OperationsForMonitorClient{
		internal: c.internal,
	}
}

// NewPredictiveMetricClient creates a new instance of PredictiveMetricClient.
func (c *ClientFactory) NewPredictiveMetricClient() *PredictiveMetricClient {
	return &PredictiveMetricClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
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

// NewPrivateLinkScopeOperationStatusClient creates a new instance of PrivateLinkScopeOperationStatusClient.
func (c *ClientFactory) NewPrivateLinkScopeOperationStatusClient() *PrivateLinkScopeOperationStatusClient {
	return &PrivateLinkScopeOperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkScopedResourcesClient creates a new instance of PrivateLinkScopedResourcesClient.
func (c *ClientFactory) NewPrivateLinkScopedResourcesClient() *PrivateLinkScopedResourcesClient {
	return &PrivateLinkScopedResourcesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewPrivateLinkScopesClient creates a new instance of PrivateLinkScopesClient.
func (c *ClientFactory) NewPrivateLinkScopesClient() *PrivateLinkScopesClient {
	return &PrivateLinkScopesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewScheduledQueryRulesClient creates a new instance of ScheduledQueryRulesClient.
func (c *ClientFactory) NewScheduledQueryRulesClient() *ScheduledQueryRulesClient {
	return &ScheduledQueryRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTenantActionGroupsClient creates a new instance of TenantActionGroupsClient.
func (c *ClientFactory) NewTenantActionGroupsClient() *TenantActionGroupsClient {
	return &TenantActionGroupsClient{
		internal: c.internal,
	}
}

// NewTenantActivityLogsClient creates a new instance of TenantActivityLogsClient.
func (c *ClientFactory) NewTenantActivityLogsClient() *TenantActivityLogsClient {
	return &TenantActivityLogsClient{
		internal: c.internal,
	}
}

// NewVMInsightsClient creates a new instance of VMInsightsClient.
func (c *ClientFactory) NewVMInsightsClient() *VMInsightsClient {
	return &VMInsightsClient{
		internal: c.internal,
	}
}
