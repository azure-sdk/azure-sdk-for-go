//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

// AccountsClientCreateResponse contains the response from method AccountsClient.BeginCreate.
type AccountsClientCreateResponse struct {
	// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// The list of cognitive services accounts operation response.
	AccountListResult
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	// The access keys for the cognitive services account.
	APIKeys
}

// AccountsClientListModelsResponse contains the response from method AccountsClient.NewListModelsPager.
type AccountsClientListModelsResponse struct {
	// The list of cognitive services accounts operation response.
	AccountModelListResult
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	// The list of cognitive services accounts operation response.
	AccountListResult
}

// AccountsClientListSKUsResponse contains the response from method AccountsClient.ListSKUs.
type AccountsClientListSKUsResponse struct {
	// The list of cognitive services accounts operation response.
	AccountSKUListResult
}

// AccountsClientListUsagesResponse contains the response from method AccountsClient.ListUsages.
type AccountsClientListUsagesResponse struct {
	// The response to a list usage request.
	UsageListResult
}

// AccountsClientRegenerateKeyResponse contains the response from method AccountsClient.RegenerateKey.
type AccountsClientRegenerateKeyResponse struct {
	// The access keys for the cognitive services account.
	APIKeys
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
	Account
}

// CommitmentPlansClientCreateOrUpdateAssociationResponse contains the response from method CommitmentPlansClient.BeginCreateOrUpdateAssociation.
type CommitmentPlansClientCreateOrUpdateAssociationResponse struct {
	// The commitment plan association.
	CommitmentPlanAccountAssociation
}

// CommitmentPlansClientCreateOrUpdatePlanResponse contains the response from method CommitmentPlansClient.BeginCreateOrUpdatePlan.
type CommitmentPlansClientCreateOrUpdatePlanResponse struct {
	// Cognitive Services account commitment plan.
	CommitmentPlan
}

// CommitmentPlansClientCreateOrUpdateResponse contains the response from method CommitmentPlansClient.CreateOrUpdate.
type CommitmentPlansClientCreateOrUpdateResponse struct {
	// Cognitive Services account commitment plan.
	CommitmentPlan
}

// CommitmentPlansClientDeleteAssociationResponse contains the response from method CommitmentPlansClient.BeginDeleteAssociation.
type CommitmentPlansClientDeleteAssociationResponse struct {
	// placeholder for future response values
}

// CommitmentPlansClientDeletePlanResponse contains the response from method CommitmentPlansClient.BeginDeletePlan.
type CommitmentPlansClientDeletePlanResponse struct {
	// placeholder for future response values
}

// CommitmentPlansClientDeleteResponse contains the response from method CommitmentPlansClient.BeginDelete.
type CommitmentPlansClientDeleteResponse struct {
	// placeholder for future response values
}

// CommitmentPlansClientGetAssociationResponse contains the response from method CommitmentPlansClient.GetAssociation.
type CommitmentPlansClientGetAssociationResponse struct {
	// The commitment plan association.
	CommitmentPlanAccountAssociation
}

// CommitmentPlansClientGetPlanResponse contains the response from method CommitmentPlansClient.GetPlan.
type CommitmentPlansClientGetPlanResponse struct {
	// Cognitive Services account commitment plan.
	CommitmentPlan
}

// CommitmentPlansClientGetResponse contains the response from method CommitmentPlansClient.Get.
type CommitmentPlansClientGetResponse struct {
	// Cognitive Services account commitment plan.
	CommitmentPlan
}

// CommitmentPlansClientListAssociationsResponse contains the response from method CommitmentPlansClient.NewListAssociationsPager.
type CommitmentPlansClientListAssociationsResponse struct {
	// The list of cognitive services Commitment Plan Account Association operation response.
	CommitmentPlanAccountAssociationListResult
}

// CommitmentPlansClientListPlansByResourceGroupResponse contains the response from method CommitmentPlansClient.NewListPlansByResourceGroupPager.
type CommitmentPlansClientListPlansByResourceGroupResponse struct {
	// The list of cognitive services accounts operation response.
	CommitmentPlanListResult
}

// CommitmentPlansClientListPlansBySubscriptionResponse contains the response from method CommitmentPlansClient.NewListPlansBySubscriptionPager.
type CommitmentPlansClientListPlansBySubscriptionResponse struct {
	// The list of cognitive services accounts operation response.
	CommitmentPlanListResult
}

// CommitmentPlansClientListResponse contains the response from method CommitmentPlansClient.NewListPager.
type CommitmentPlansClientListResponse struct {
	// The list of cognitive services accounts operation response.
	CommitmentPlanListResult
}

// CommitmentPlansClientUpdatePlanResponse contains the response from method CommitmentPlansClient.BeginUpdatePlan.
type CommitmentPlansClientUpdatePlanResponse struct {
	// Cognitive Services account commitment plan.
	CommitmentPlan
}

// CommitmentTiersClientListResponse contains the response from method CommitmentTiersClient.NewListPager.
type CommitmentTiersClientListResponse struct {
	// The list of cognitive services accounts operation response.
	CommitmentTierListResult
}

// DeletedAccountsClientGetResponse contains the response from method DeletedAccountsClient.Get.
type DeletedAccountsClientGetResponse struct {
	// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
	Account
}

// DeletedAccountsClientListResponse contains the response from method DeletedAccountsClient.NewListPager.
type DeletedAccountsClientListResponse struct {
	// The list of cognitive services accounts operation response.
	AccountListResult
}

// DeletedAccountsClientPurgeResponse contains the response from method DeletedAccountsClient.BeginPurge.
type DeletedAccountsClientPurgeResponse struct {
	// placeholder for future response values
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.BeginCreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	// Cognitive Services account deployment.
	Deployment
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.BeginDelete.
type DeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	// Cognitive Services account deployment.
	Deployment
}

// DeploymentsClientListResponse contains the response from method DeploymentsClient.NewListPager.
type DeploymentsClientListResponse struct {
	// The list of cognitive services accounts operation response.
	DeploymentListResult
}

// DeploymentsClientListSKUsResponse contains the response from method DeploymentsClient.NewListSKUsPager.
type DeploymentsClientListSKUsResponse struct {
	// The list of cognitive services accounts operation response.
	DeploymentSKUListResult
}

// DeploymentsClientUpdateResponse contains the response from method DeploymentsClient.BeginUpdate.
type DeploymentsClientUpdateResponse struct {
	// Cognitive Services account deployment.
	Deployment
}

// EncryptionScopesClientCreateOrUpdateResponse contains the response from method EncryptionScopesClient.CreateOrUpdate.
type EncryptionScopesClientCreateOrUpdateResponse struct {
	// Cognitive Services EncryptionScope
	EncryptionScope
}

// EncryptionScopesClientDeleteResponse contains the response from method EncryptionScopesClient.BeginDelete.
type EncryptionScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// EncryptionScopesClientGetResponse contains the response from method EncryptionScopesClient.Get.
type EncryptionScopesClientGetResponse struct {
	// Cognitive Services EncryptionScope
	EncryptionScope
}

// EncryptionScopesClientListResponse contains the response from method EncryptionScopesClient.NewListPager.
type EncryptionScopesClientListResponse struct {
	// The list of cognitive services EncryptionScopes.
	EncryptionScopeListResult
}

// ManagementClientCheckDomainAvailabilityResponse contains the response from method ManagementClient.CheckDomainAvailability.
type ManagementClientCheckDomainAvailabilityResponse struct {
	// Domain availability.
	DomainAvailability
}

// ManagementClientCheckSKUAvailabilityResponse contains the response from method ManagementClient.CheckSKUAvailability.
type ManagementClientCheckSKUAvailabilityResponse struct {
	// Check SKU availability result list.
	SKUAvailabilityListResult
}

// ModelsClientListResponse contains the response from method ModelsClient.NewListPager.
type ModelsClientListResponse struct {
	// The list of cognitive services models.
	ModelListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	// A list of private endpoint connections
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// RaiBlocklistItemsClientCreateOrUpdateResponse contains the response from method RaiBlocklistItemsClient.CreateOrUpdate.
type RaiBlocklistItemsClientCreateOrUpdateResponse struct {
	// Cognitive Services RaiBlocklist Item.
	RaiBlocklistItem
}

// RaiBlocklistItemsClientDeleteResponse contains the response from method RaiBlocklistItemsClient.BeginDelete.
type RaiBlocklistItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// RaiBlocklistItemsClientGetResponse contains the response from method RaiBlocklistItemsClient.Get.
type RaiBlocklistItemsClientGetResponse struct {
	// Cognitive Services RaiBlocklist Item.
	RaiBlocklistItem
}

// RaiBlocklistItemsClientListResponse contains the response from method RaiBlocklistItemsClient.NewListPager.
type RaiBlocklistItemsClientListResponse struct {
	// The list of cognitive services RAI Blocklist Items.
	RaiBlockListItemsResult
}

// RaiBlocklistsClientCreateOrUpdateResponse contains the response from method RaiBlocklistsClient.CreateOrUpdate.
type RaiBlocklistsClientCreateOrUpdateResponse struct {
	// Cognitive Services RaiBlocklist.
	RaiBlocklist
}

// RaiBlocklistsClientDeleteResponse contains the response from method RaiBlocklistsClient.BeginDelete.
type RaiBlocklistsClientDeleteResponse struct {
	// placeholder for future response values
}

// RaiBlocklistsClientGetResponse contains the response from method RaiBlocklistsClient.Get.
type RaiBlocklistsClientGetResponse struct {
	// Cognitive Services RaiBlocklist.
	RaiBlocklist
}

// RaiBlocklistsClientListResponse contains the response from method RaiBlocklistsClient.NewListPager.
type RaiBlocklistsClientListResponse struct {
	// The list of cognitive services RAI Blocklists.
	RaiBlockListResult
}

// RaiContentFiltersClientListResponse contains the response from method RaiContentFiltersClient.NewListPager.
type RaiContentFiltersClientListResponse struct {
	// The list of Content Filters.
	RaiContentFilterListResult
}

// RaiPoliciesClientCreateOrUpdateResponse contains the response from method RaiPoliciesClient.CreateOrUpdate.
type RaiPoliciesClientCreateOrUpdateResponse struct {
	// Cognitive Services RaiPolicy.
	RaiPolicy
}

// RaiPoliciesClientDeleteResponse contains the response from method RaiPoliciesClient.BeginDelete.
type RaiPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// RaiPoliciesClientGetResponse contains the response from method RaiPoliciesClient.Get.
type RaiPoliciesClientGetResponse struct {
	// Cognitive Services RaiPolicy.
	RaiPolicy
}

// RaiPoliciesClientListResponse contains the response from method RaiPoliciesClient.NewListPager.
type RaiPoliciesClientListResponse struct {
	// The list of cognitive services RaiPolicies.
	RaiPolicyListResult
}

// ResourceSKUsClientListResponse contains the response from method ResourceSKUsClient.NewListPager.
type ResourceSKUsClientListResponse struct {
	// The Get Skus operation response.
	ResourceSKUListResult
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	// The response to a list usage request.
	UsageListResult
}
