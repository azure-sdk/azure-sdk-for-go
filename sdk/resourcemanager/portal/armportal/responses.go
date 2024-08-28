//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

// DashboardsClientCreateOrUpdateResponse contains the response from method DashboardsClient.CreateOrUpdate.
type DashboardsClientCreateOrUpdateResponse struct {
	// The shared dashboard resource definition.
	Dashboard
}

// DashboardsClientDeleteResponse contains the response from method DashboardsClient.Delete.
type DashboardsClientDeleteResponse struct {
	// placeholder for future response values
}

// DashboardsClientGetResponse contains the response from method DashboardsClient.Get.
type DashboardsClientGetResponse struct {
	// The shared dashboard resource definition.
	Dashboard
}

// DashboardsClientListByResourceGroupResponse contains the response from method DashboardsClient.NewListByResourceGroupPager.
type DashboardsClientListByResourceGroupResponse struct {
	// The response of a Dashboard list operation.
	DashboardListResult
}

// DashboardsClientListBySubscriptionResponse contains the response from method DashboardsClient.NewListBySubscriptionPager.
type DashboardsClientListBySubscriptionResponse struct {
	// The response of a Dashboard list operation.
	DashboardListResult
}

// DashboardsClientUpdateResponse contains the response from method DashboardsClient.Update.
type DashboardsClientUpdateResponse struct {
	// The shared dashboard resource definition.
	Dashboard
}

// ListTenantConfigurationViolationsClientListResponse contains the response from method ListTenantConfigurationViolationsClient.NewListPager.
type ListTenantConfigurationViolationsClientListResponse struct {
	// List of list of items that violate tenant's configuration.
	ViolationsList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// TenantConfigurationsClientCreateResponse contains the response from method TenantConfigurationsClient.Create.
type TenantConfigurationsClientCreateResponse struct {
	// The tenant configuration resource definition.
	Configuration
}

// TenantConfigurationsClientDeleteResponse contains the response from method TenantConfigurationsClient.Delete.
type TenantConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// TenantConfigurationsClientGetResponse contains the response from method TenantConfigurationsClient.Get.
type TenantConfigurationsClientGetResponse struct {
	// The tenant configuration resource definition.
	Configuration
}

// TenantConfigurationsClientListResponse contains the response from method TenantConfigurationsClient.NewListPager.
type TenantConfigurationsClientListResponse struct {
	// The response of a Configuration list operation.
	ConfigurationListResult
}