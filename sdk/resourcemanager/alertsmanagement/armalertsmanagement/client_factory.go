// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	issueName      string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - issueName - The name of the IssueResource
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, issueName string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		issueName:      issueName,
		internal:       internal,
	}, nil
}

// NewAlertProcessingRulesClient creates a new instance of AlertProcessingRulesClient.
func (c *ClientFactory) NewAlertProcessingRulesClient() *AlertProcessingRulesClient {
	return &AlertProcessingRulesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAlertRuleRecommendationsClient creates a new instance of AlertRuleRecommendationsClient.
func (c *ClientFactory) NewAlertRuleRecommendationsClient() *AlertRuleRecommendationsClient {
	return &AlertRuleRecommendationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewAlertsClient creates a new instance of AlertsClient.
func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	return &AlertsClient{
		internal: c.internal,
	}
}

// NewIssueClient creates a new instance of IssueClient.
func (c *ClientFactory) NewIssueClient() *IssueClient {
	return &IssueClient{
		issueName: c.issueName,
		internal:  c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewPrometheusRuleGroupsClient creates a new instance of PrometheusRuleGroupsClient.
func (c *ClientFactory) NewPrometheusRuleGroupsClient() *PrometheusRuleGroupsClient {
	return &PrometheusRuleGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSmartGroupsClient creates a new instance of SmartGroupsClient.
func (c *ClientFactory) NewSmartGroupsClient() *SmartGroupsClient {
	return &SmartGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
