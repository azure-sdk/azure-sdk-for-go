//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

// BillingInfoClientGetResponse contains the response from method BillingInfoClient.Get.
type BillingInfoClientGetResponse struct {
	// Marketplace Subscription and Organization details to which resource gets billed into.
	BillingInfoResponse
}

// CreationSupportedClientGetResponse contains the response from method CreationSupportedClient.Get.
type CreationSupportedClientGetResponse struct {
	// Datadog resource can be created or not.
	CreateResourceSupportedResponse
}

// CreationSupportedClientListResponse contains the response from method CreationSupportedClient.NewListPager.
type CreationSupportedClientListResponse struct {
	CreateResourceSupportedResponseList
}

// MarketplaceAgreementsClientCreateOrUpdateResponse contains the response from method MarketplaceAgreementsClient.CreateOrUpdate.
type MarketplaceAgreementsClientCreateOrUpdateResponse struct {
	AgreementResource
}

// MarketplaceAgreementsClientListResponse contains the response from method MarketplaceAgreementsClient.NewListPager.
type MarketplaceAgreementsClientListResponse struct {
	// Response of a list operation.
	AgreementResourceListResponse
}

// MonitoredSubscriptionsClientCreateorUpdateResponse contains the response from method MonitoredSubscriptionsClient.BeginCreateorUpdate.
type MonitoredSubscriptionsClientCreateorUpdateResponse struct {
	// The request to update subscriptions needed to be monitored by the Datadog monitor resource.
	MonitoredSubscriptionProperties
}

// MonitoredSubscriptionsClientDeleteResponse contains the response from method MonitoredSubscriptionsClient.BeginDelete.
type MonitoredSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitoredSubscriptionsClientGetResponse contains the response from method MonitoredSubscriptionsClient.Get.
type MonitoredSubscriptionsClientGetResponse struct {
	// The request to update subscriptions needed to be monitored by the Datadog monitor resource.
	MonitoredSubscriptionProperties
}

// MonitoredSubscriptionsClientListResponse contains the response from method MonitoredSubscriptionsClient.NewListPager.
type MonitoredSubscriptionsClientListResponse struct {
	MonitoredSubscriptionPropertiesList
}

// MonitoredSubscriptionsClientUpdateResponse contains the response from method MonitoredSubscriptionsClient.BeginUpdate.
type MonitoredSubscriptionsClientUpdateResponse struct {
	// The request to update subscriptions needed to be monitored by the Datadog monitor resource.
	MonitoredSubscriptionProperties
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.BeginCreate.
type MonitorsClientCreateResponse struct {
	MonitorResource
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.BeginDelete.
type MonitorsClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitorsClientGetDefaultKeyResponse contains the response from method MonitorsClient.GetDefaultKey.
type MonitorsClientGetDefaultKeyResponse struct {
	APIKey
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorResource
}

// MonitorsClientListAPIKeysResponse contains the response from method MonitorsClient.NewListAPIKeysPager.
type MonitorsClientListAPIKeysResponse struct {
	// Response of a list operation.
	APIKeyListResponse
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.NewListByResourceGroupPager.
type MonitorsClientListByResourceGroupResponse struct {
	// Response of a list operation.
	MonitorResourceListResponse
}

// MonitorsClientListHostsResponse contains the response from method MonitorsClient.NewListHostsPager.
type MonitorsClientListHostsResponse struct {
	// Response of a list operation.
	HostListResponse
}

// MonitorsClientListLinkedResourcesResponse contains the response from method MonitorsClient.NewListLinkedResourcesPager.
type MonitorsClientListLinkedResourcesResponse struct {
	// Response of a list operation.
	LinkedResourceListResponse
}

// MonitorsClientListMonitoredResourcesResponse contains the response from method MonitorsClient.NewListMonitoredResourcesPager.
type MonitorsClientListMonitoredResourcesResponse struct {
	// Response of a list operation.
	MonitoredResourceListResponse
}

// MonitorsClientListResponse contains the response from method MonitorsClient.NewListPager.
type MonitorsClientListResponse struct {
	// Response of a list operation.
	MonitorResourceListResponse
}

// MonitorsClientRefreshSetPasswordLinkResponse contains the response from method MonitorsClient.RefreshSetPasswordLink.
type MonitorsClientRefreshSetPasswordLinkResponse struct {
	SetPasswordLink
}

// MonitorsClientSetDefaultKeyResponse contains the response from method MonitorsClient.SetDefaultKey.
type MonitorsClientSetDefaultKeyResponse struct {
	// placeholder for future response values
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.BeginUpdate.
type MonitorsClientUpdateResponse struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of GET request to list the Microsoft.Datadog operations.
	OperationListResult
}

// SingleSignOnConfigurationsClientCreateOrUpdateResponse contains the response from method SingleSignOnConfigurationsClient.BeginCreateOrUpdate.
type SingleSignOnConfigurationsClientCreateOrUpdateResponse struct {
	SingleSignOnResource
}

// SingleSignOnConfigurationsClientGetResponse contains the response from method SingleSignOnConfigurationsClient.Get.
type SingleSignOnConfigurationsClientGetResponse struct {
	SingleSignOnResource
}

// SingleSignOnConfigurationsClientListResponse contains the response from method SingleSignOnConfigurationsClient.NewListPager.
type SingleSignOnConfigurationsClientListResponse struct {
	// Response of a list operation.
	SingleSignOnResourceListResponse
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	// Capture logs and metrics of Azure resources based on ARM tags.
	MonitoringTagRules
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	// Capture logs and metrics of Azure resources based on ARM tags.
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.NewListPager.
type TagRulesClientListResponse struct {
	// Response of a list operation.
	MonitoringTagRulesListResponse
}
