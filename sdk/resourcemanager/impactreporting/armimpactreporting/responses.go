// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armimpactreporting

// ConnectorsClientCreateOrUpdateResponse contains the response from method ConnectorsClient.BeginCreateOrUpdate.
type ConnectorsClientCreateOrUpdateResponse struct {
	// A connector is a resource that can be used to proactively report impacts against workloads in Azure to Microsoft.
	Connector
}

// ConnectorsClientDeleteResponse contains the response from method ConnectorsClient.Delete.
type ConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectorsClientGetResponse contains the response from method ConnectorsClient.Get.
type ConnectorsClientGetResponse struct {
	// A connector is a resource that can be used to proactively report impacts against workloads in Azure to Microsoft.
	Connector
}

// ConnectorsClientListBySubscriptionResponse contains the response from method ConnectorsClient.NewListBySubscriptionPager.
type ConnectorsClientListBySubscriptionResponse struct {
	// The response of a Connector list operation.
	ConnectorListResult
}

// ConnectorsClientUpdateResponse contains the response from method ConnectorsClient.Update.
type ConnectorsClientUpdateResponse struct {
	// A connector is a resource that can be used to proactively report impacts against workloads in Azure to Microsoft.
	Connector
}

// ImpactCategoriesClientGetResponse contains the response from method ImpactCategoriesClient.Get.
type ImpactCategoriesClientGetResponse struct {
	// ImpactCategory resource
	ImpactCategory
}

// ImpactCategoriesClientListBySubscriptionResponse contains the response from method ImpactCategoriesClient.NewListBySubscriptionPager.
type ImpactCategoriesClientListBySubscriptionResponse struct {
	// The response of a ImpactCategory list operation.
	ImpactCategoryListResult
}

// InsightsClientCreateResponse contains the response from method InsightsClient.Create.
type InsightsClientCreateResponse struct {
	// Insight resource
	Insight
}

// InsightsClientDeleteResponse contains the response from method InsightsClient.Delete.
type InsightsClientDeleteResponse struct {
	// placeholder for future response values
}

// InsightsClientGetResponse contains the response from method InsightsClient.Get.
type InsightsClientGetResponse struct {
	// Insight resource
	Insight
}

// InsightsClientListBySubscriptionResponse contains the response from method InsightsClient.NewListBySubscriptionPager.
type InsightsClientListBySubscriptionResponse struct {
	// The response of a Insight list operation.
	InsightListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// WorkloadImpactsClientCreateResponse contains the response from method WorkloadImpactsClient.BeginCreate.
type WorkloadImpactsClientCreateResponse struct {
	// Workload Impact properties
	WorkloadImpact
}

// WorkloadImpactsClientDeleteResponse contains the response from method WorkloadImpactsClient.Delete.
type WorkloadImpactsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkloadImpactsClientGetResponse contains the response from method WorkloadImpactsClient.Get.
type WorkloadImpactsClientGetResponse struct {
	// Workload Impact properties
	WorkloadImpact
}

// WorkloadImpactsClientListBySubscriptionResponse contains the response from method WorkloadImpactsClient.NewListBySubscriptionPager.
type WorkloadImpactsClientListBySubscriptionResponse struct {
	// The response of a WorkloadImpact list operation.
	WorkloadImpactListResult
}
