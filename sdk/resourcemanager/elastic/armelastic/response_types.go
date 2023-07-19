//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic

// AllTrafficFiltersClientListResponse contains the response from method AllTrafficFiltersClient.List.
type AllTrafficFiltersClientListResponse struct {
	TrafficFilterResponse
}

// AssociateTrafficFilterClientAssociateResponse contains the response from method AssociateTrafficFilterClient.BeginAssociate.
type AssociateTrafficFilterClientAssociateResponse struct {
	// placeholder for future response values
}

// BillingInfoClientGetResponse contains the response from method BillingInfoClient.Get.
type BillingInfoClientGetResponse struct {
	BillingInfoResponse
}

// ConnectedPartnerResourcesClientListResponse contains the response from method ConnectedPartnerResourcesClient.NewListPager.
type ConnectedPartnerResourcesClientListResponse struct {
	ConnectedPartnerResourcesListResponse
}

// CreateAndAssociateIPFilterClientCreateResponse contains the response from method CreateAndAssociateIPFilterClient.BeginCreate.
type CreateAndAssociateIPFilterClientCreateResponse struct {
	// placeholder for future response values
}

// CreateAndAssociatePLFilterClientCreateResponse contains the response from method CreateAndAssociatePLFilterClient.BeginCreate.
type CreateAndAssociatePLFilterClientCreateResponse struct {
	// placeholder for future response values
}

// DeploymentInfoClientListResponse contains the response from method DeploymentInfoClient.List.
type DeploymentInfoClientListResponse struct {
	DeploymentInfoResponse
}

// DetachAndDeleteTrafficFilterClientDeleteResponse contains the response from method DetachAndDeleteTrafficFilterClient.Delete.
type DetachAndDeleteTrafficFilterClientDeleteResponse struct {
	// placeholder for future response values
}

// DetachTrafficFilterClientUpdateResponse contains the response from method DetachTrafficFilterClient.BeginUpdate.
type DetachTrafficFilterClientUpdateResponse struct {
	// placeholder for future response values
}

// ExternalUserClientCreateOrUpdateResponse contains the response from method ExternalUserClient.CreateOrUpdate.
type ExternalUserClientCreateOrUpdateResponse struct {
	ExternalUserCreationResponse
}

// ListAssociatedTrafficFiltersClientListResponse contains the response from method ListAssociatedTrafficFiltersClient.List.
type ListAssociatedTrafficFiltersClientListResponse struct {
	TrafficFilterResponse
}

// MonitorClientUpgradeResponse contains the response from method MonitorClient.BeginUpgrade.
type MonitorClientUpgradeResponse struct {
	// placeholder for future response values
}

// MonitoredResourcesClientListResponse contains the response from method MonitoredResourcesClient.NewListPager.
type MonitoredResourcesClientListResponse struct {
	MonitoredResourceListResponse
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.BeginCreate.
type MonitorsClientCreateResponse struct {
	MonitorResource
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.BeginDelete.
type MonitorsClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorResource
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.NewListByResourceGroupPager.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientListResponse contains the response from method MonitorsClient.NewListPager.
type MonitorsClientListResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// OrganizationsClientGetAPIKeyResponse contains the response from method OrganizationsClient.GetAPIKey.
type OrganizationsClientGetAPIKeyResponse struct {
	UserAPIKeyResponse
}

// OrganizationsClientGetElasticToAzureSubscriptionMappingResponse contains the response from method OrganizationsClient.GetElasticToAzureSubscriptionMapping.
type OrganizationsClientGetElasticToAzureSubscriptionMappingResponse struct {
	OrganizationToAzureSubscriptionMappingResponse
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	MonitoringTagRules
}

// TagRulesClientDeleteResponse contains the response from method TagRulesClient.BeginDelete.
type TagRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.NewListPager.
type TagRulesClientListResponse struct {
	MonitoringTagRulesListResponse
}

// TrafficFiltersClientDeleteResponse contains the response from method TrafficFiltersClient.Delete.
type TrafficFiltersClientDeleteResponse struct {
	// placeholder for future response values
}

// UpgradableVersionsClientDetailsResponse contains the response from method UpgradableVersionsClient.Details.
type UpgradableVersionsClientDetailsResponse struct {
	UpgradableVersionsList
}

// VMCollectionClientUpdateResponse contains the response from method VMCollectionClient.Update.
type VMCollectionClientUpdateResponse struct {
	// placeholder for future response values
}

// VMHostClientListResponse contains the response from method VMHostClient.NewListPager.
type VMHostClientListResponse struct {
	VMHostListResponse
}

// VMIngestionClientDetailsResponse contains the response from method VMIngestionClient.Details.
type VMIngestionClientDetailsResponse struct {
	VMIngestionDetailsResponse
}

// VersionsClientListResponse contains the response from method VersionsClient.NewListPager.
type VersionsClientListResponse struct {
	VersionsListResponse
}
