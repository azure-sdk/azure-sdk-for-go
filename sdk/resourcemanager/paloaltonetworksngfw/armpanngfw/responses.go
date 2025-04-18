// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

// CertificateObjectGlobalRulestackClientCreateOrUpdateResponse contains the response from method CertificateObjectGlobalRulestackClient.BeginCreateOrUpdate.
type CertificateObjectGlobalRulestackClientCreateOrUpdateResponse struct {
	// GlobalRulestack Certificate Object
	CertificateObjectGlobalRulestackResource
}

// CertificateObjectGlobalRulestackClientDeleteResponse contains the response from method CertificateObjectGlobalRulestackClient.BeginDelete.
type CertificateObjectGlobalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificateObjectGlobalRulestackClientGetResponse contains the response from method CertificateObjectGlobalRulestackClient.Get.
type CertificateObjectGlobalRulestackClientGetResponse struct {
	// GlobalRulestack Certificate Object
	CertificateObjectGlobalRulestackResource
}

// CertificateObjectGlobalRulestackClientListResponse contains the response from method CertificateObjectGlobalRulestackClient.NewListPager.
type CertificateObjectGlobalRulestackClientListResponse struct {
	// The response of a CertificateObjectGlobalRulestackResource list operation.
	CertificateObjectGlobalRulestackResourceListResult
}

// CertificateObjectLocalRulestackClientCreateOrUpdateResponse contains the response from method CertificateObjectLocalRulestackClient.BeginCreateOrUpdate.
type CertificateObjectLocalRulestackClientCreateOrUpdateResponse struct {
	// LocalRulestack Certificate Object
	CertificateObjectLocalRulestackResource
}

// CertificateObjectLocalRulestackClientDeleteResponse contains the response from method CertificateObjectLocalRulestackClient.BeginDelete.
type CertificateObjectLocalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificateObjectLocalRulestackClientGetResponse contains the response from method CertificateObjectLocalRulestackClient.Get.
type CertificateObjectLocalRulestackClientGetResponse struct {
	// LocalRulestack Certificate Object
	CertificateObjectLocalRulestackResource
}

// CertificateObjectLocalRulestackClientListByLocalRulestacksResponse contains the response from method CertificateObjectLocalRulestackClient.NewListByLocalRulestacksPager.
type CertificateObjectLocalRulestackClientListByLocalRulestacksResponse struct {
	// The response of a CertificateObjectLocalRulestackResource list operation.
	CertificateObjectLocalRulestackResourceListResult
}

// FirewallStatusClientGetResponse contains the response from method FirewallStatusClient.Get.
type FirewallStatusClientGetResponse struct {
	// Firewall Status
	FirewallStatusResource
}

// FirewallStatusClientListByFirewallsResponse contains the response from method FirewallStatusClient.NewListByFirewallsPager.
type FirewallStatusClientListByFirewallsResponse struct {
	// The response of a FirewallStatusResource list operation.
	FirewallStatusResourceListResult
}

// FirewallsClientCreateOrUpdateResponse contains the response from method FirewallsClient.BeginCreateOrUpdate.
type FirewallsClientCreateOrUpdateResponse struct {
	// PaloAltoNetworks Firewall
	FirewallResource
}

// FirewallsClientDeleteResponse contains the response from method FirewallsClient.BeginDelete.
type FirewallsClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallsClientGetGlobalRulestackResponse contains the response from method FirewallsClient.GetGlobalRulestack.
type FirewallsClientGetGlobalRulestackResponse struct {
	// PAN Rulestack Describe Object
	GlobalRulestackInfo
}

// FirewallsClientGetLogProfileResponse contains the response from method FirewallsClient.GetLogProfile.
type FirewallsClientGetLogProfileResponse struct {
	// Log Settings for Firewall
	LogSettings
}

// FirewallsClientGetResponse contains the response from method FirewallsClient.Get.
type FirewallsClientGetResponse struct {
	// PaloAltoNetworks Firewall
	FirewallResource
}

// FirewallsClientGetSupportInfoResponse contains the response from method FirewallsClient.GetSupportInfo.
type FirewallsClientGetSupportInfoResponse struct {
	// Support information for the resource
	SupportInfo
}

// FirewallsClientListByResourceGroupResponse contains the response from method FirewallsClient.NewListByResourceGroupPager.
type FirewallsClientListByResourceGroupResponse struct {
	// The response of a FirewallResource list operation.
	FirewallResourceListResult
}

// FirewallsClientListBySubscriptionResponse contains the response from method FirewallsClient.NewListBySubscriptionPager.
type FirewallsClientListBySubscriptionResponse struct {
	// The response of a FirewallResource list operation.
	FirewallResourceListResult
}

// FirewallsClientSaveLogProfileResponse contains the response from method FirewallsClient.SaveLogProfile.
type FirewallsClientSaveLogProfileResponse struct {
	// placeholder for future response values
}

// FirewallsClientUpdateResponse contains the response from method FirewallsClient.Update.
type FirewallsClientUpdateResponse struct {
	// PaloAltoNetworks Firewall
	FirewallResource
}

// FqdnListGlobalRulestackClientCreateOrUpdateResponse contains the response from method FqdnListGlobalRulestackClient.BeginCreateOrUpdate.
type FqdnListGlobalRulestackClientCreateOrUpdateResponse struct {
	// GlobalRulestack fqdnList
	FqdnListGlobalRulestackResource
}

// FqdnListGlobalRulestackClientDeleteResponse contains the response from method FqdnListGlobalRulestackClient.BeginDelete.
type FqdnListGlobalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// FqdnListGlobalRulestackClientGetResponse contains the response from method FqdnListGlobalRulestackClient.Get.
type FqdnListGlobalRulestackClientGetResponse struct {
	// GlobalRulestack fqdnList
	FqdnListGlobalRulestackResource
}

// FqdnListGlobalRulestackClientListResponse contains the response from method FqdnListGlobalRulestackClient.NewListPager.
type FqdnListGlobalRulestackClientListResponse struct {
	// The response of a FqdnListGlobalRulestackResource list operation.
	FqdnListGlobalRulestackResourceListResult
}

// FqdnListLocalRulestackClientCreateOrUpdateResponse contains the response from method FqdnListLocalRulestackClient.BeginCreateOrUpdate.
type FqdnListLocalRulestackClientCreateOrUpdateResponse struct {
	// LocalRulestack fqdnList
	FqdnListLocalRulestackResource
}

// FqdnListLocalRulestackClientDeleteResponse contains the response from method FqdnListLocalRulestackClient.BeginDelete.
type FqdnListLocalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// FqdnListLocalRulestackClientGetResponse contains the response from method FqdnListLocalRulestackClient.Get.
type FqdnListLocalRulestackClientGetResponse struct {
	// LocalRulestack fqdnList
	FqdnListLocalRulestackResource
}

// FqdnListLocalRulestackClientListByLocalRulestacksResponse contains the response from method FqdnListLocalRulestackClient.NewListByLocalRulestacksPager.
type FqdnListLocalRulestackClientListByLocalRulestacksResponse struct {
	// The response of a FqdnListLocalRulestackResource list operation.
	FqdnListLocalRulestackResourceListResult
}

// GlobalRulestackClientCommitResponse contains the response from method GlobalRulestackClient.BeginCommit.
type GlobalRulestackClientCommitResponse struct {
	// placeholder for future response values
}

// GlobalRulestackClientCreateOrUpdateResponse contains the response from method GlobalRulestackClient.BeginCreateOrUpdate.
type GlobalRulestackClientCreateOrUpdateResponse struct {
	// PaloAltoNetworks GlobalRulestack
	GlobalRulestackResource
}

// GlobalRulestackClientDeleteResponse contains the response from method GlobalRulestackClient.BeginDelete.
type GlobalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// GlobalRulestackClientGetChangeLogResponse contains the response from method GlobalRulestackClient.GetChangeLog.
type GlobalRulestackClientGetChangeLogResponse struct {
	// Changelog list
	Changelog
}

// GlobalRulestackClientGetResponse contains the response from method GlobalRulestackClient.Get.
type GlobalRulestackClientGetResponse struct {
	// PaloAltoNetworks GlobalRulestack
	GlobalRulestackResource
}

// GlobalRulestackClientListAdvancedSecurityObjectsResponse contains the response from method GlobalRulestackClient.ListAdvancedSecurityObjects.
type GlobalRulestackClientListAdvancedSecurityObjectsResponse struct {
	// advanced security object
	AdvSecurityObjectListResponse
}

// GlobalRulestackClientListAppIDsResponse contains the response from method GlobalRulestackClient.ListAppIDs.
type GlobalRulestackClientListAppIDsResponse struct {
	ListAppIDResponse
}

// GlobalRulestackClientListCountriesResponse contains the response from method GlobalRulestackClient.ListCountries.
type GlobalRulestackClientListCountriesResponse struct {
	// Countries Response Object
	CountriesResponse
}

// GlobalRulestackClientListFirewallsResponse contains the response from method GlobalRulestackClient.ListFirewalls.
type GlobalRulestackClientListFirewallsResponse struct {
	// List firewalls response
	ListFirewallsResponse
}

// GlobalRulestackClientListPredefinedURLCategoriesResponse contains the response from method GlobalRulestackClient.ListPredefinedURLCategories.
type GlobalRulestackClientListPredefinedURLCategoriesResponse struct {
	// predefined url categories response
	PredefinedURLCategoriesResponse
}

// GlobalRulestackClientListResponse contains the response from method GlobalRulestackClient.NewListPager.
type GlobalRulestackClientListResponse struct {
	// The response of a GlobalRulestackResource list operation.
	GlobalRulestackResourceListResult
}

// GlobalRulestackClientListSecurityServicesResponse contains the response from method GlobalRulestackClient.ListSecurityServices.
type GlobalRulestackClientListSecurityServicesResponse struct {
	// Security services list response
	SecurityServicesResponse
}

// GlobalRulestackClientRevertResponse contains the response from method GlobalRulestackClient.Revert.
type GlobalRulestackClientRevertResponse struct {
	// placeholder for future response values
}

// GlobalRulestackClientUpdateResponse contains the response from method GlobalRulestackClient.Update.
type GlobalRulestackClientUpdateResponse struct {
	// PaloAltoNetworks GlobalRulestack
	GlobalRulestackResource
}

// LocalRulesClientCreateOrUpdateResponse contains the response from method LocalRulesClient.BeginCreateOrUpdate.
type LocalRulesClientCreateOrUpdateResponse struct {
	// LocalRulestack rule list
	LocalRulesResource
}

// LocalRulesClientDeleteResponse contains the response from method LocalRulesClient.BeginDelete.
type LocalRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// LocalRulesClientGetCountersResponse contains the response from method LocalRulesClient.GetCounters.
type LocalRulesClientGetCountersResponse struct {
	// Rule counter
	RuleCounter
}

// LocalRulesClientGetResponse contains the response from method LocalRulesClient.Get.
type LocalRulesClientGetResponse struct {
	// LocalRulestack rule list
	LocalRulesResource
}

// LocalRulesClientListByLocalRulestacksResponse contains the response from method LocalRulesClient.NewListByLocalRulestacksPager.
type LocalRulesClientListByLocalRulestacksResponse struct {
	// The response of a LocalRulesResource list operation.
	LocalRulesResourceListResult
}

// LocalRulesClientRefreshCountersResponse contains the response from method LocalRulesClient.RefreshCounters.
type LocalRulesClientRefreshCountersResponse struct {
	// placeholder for future response values
}

// LocalRulesClientResetCountersResponse contains the response from method LocalRulesClient.ResetCounters.
type LocalRulesClientResetCountersResponse struct {
	// Rule counter reset
	RuleCounterReset
}

// LocalRulestacksClientCommitResponse contains the response from method LocalRulestacksClient.BeginCommit.
type LocalRulestacksClientCommitResponse struct {
	// placeholder for future response values
}

// LocalRulestacksClientCreateOrUpdateResponse contains the response from method LocalRulestacksClient.BeginCreateOrUpdate.
type LocalRulestacksClientCreateOrUpdateResponse struct {
	// PaloAltoNetworks LocalRulestack
	LocalRulestackResource
}

// LocalRulestacksClientDeleteResponse contains the response from method LocalRulestacksClient.BeginDelete.
type LocalRulestacksClientDeleteResponse struct {
	// placeholder for future response values
}

// LocalRulestacksClientGetChangeLogResponse contains the response from method LocalRulestacksClient.GetChangeLog.
type LocalRulestacksClientGetChangeLogResponse struct {
	// Changelog list
	Changelog
}

// LocalRulestacksClientGetResponse contains the response from method LocalRulestacksClient.Get.
type LocalRulestacksClientGetResponse struct {
	// PaloAltoNetworks LocalRulestack
	LocalRulestackResource
}

// LocalRulestacksClientGetSupportInfoResponse contains the response from method LocalRulestacksClient.GetSupportInfo.
type LocalRulestacksClientGetSupportInfoResponse struct {
	// Support information for the resource
	SupportInfo
}

// LocalRulestacksClientListAdvancedSecurityObjectsResponse contains the response from method LocalRulestacksClient.ListAdvancedSecurityObjects.
type LocalRulestacksClientListAdvancedSecurityObjectsResponse struct {
	// advanced security object
	AdvSecurityObjectListResponse
}

// LocalRulestacksClientListAppIDsResponse contains the response from method LocalRulestacksClient.ListAppIDs.
type LocalRulestacksClientListAppIDsResponse struct {
	ListAppIDResponse
}

// LocalRulestacksClientListByResourceGroupResponse contains the response from method LocalRulestacksClient.NewListByResourceGroupPager.
type LocalRulestacksClientListByResourceGroupResponse struct {
	// The response of a LocalRulestackResource list operation.
	LocalRulestackResourceListResult
}

// LocalRulestacksClientListBySubscriptionResponse contains the response from method LocalRulestacksClient.NewListBySubscriptionPager.
type LocalRulestacksClientListBySubscriptionResponse struct {
	// The response of a LocalRulestackResource list operation.
	LocalRulestackResourceListResult
}

// LocalRulestacksClientListCountriesResponse contains the response from method LocalRulestacksClient.ListCountries.
type LocalRulestacksClientListCountriesResponse struct {
	// Countries Response Object
	CountriesResponse
}

// LocalRulestacksClientListFirewallsResponse contains the response from method LocalRulestacksClient.ListFirewalls.
type LocalRulestacksClientListFirewallsResponse struct {
	// List firewalls response
	ListFirewallsResponse
}

// LocalRulestacksClientListPredefinedURLCategoriesResponse contains the response from method LocalRulestacksClient.ListPredefinedURLCategories.
type LocalRulestacksClientListPredefinedURLCategoriesResponse struct {
	// predefined url categories response
	PredefinedURLCategoriesResponse
}

// LocalRulestacksClientListSecurityServicesResponse contains the response from method LocalRulestacksClient.ListSecurityServices.
type LocalRulestacksClientListSecurityServicesResponse struct {
	// Security services list response
	SecurityServicesResponse
}

// LocalRulestacksClientRevertResponse contains the response from method LocalRulestacksClient.Revert.
type LocalRulestacksClientRevertResponse struct {
	// placeholder for future response values
}

// LocalRulestacksClientUpdateResponse contains the response from method LocalRulestacksClient.Update.
type LocalRulestacksClientUpdateResponse struct {
	// PaloAltoNetworks LocalRulestack
	LocalRulestackResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PaloAltoNetworksCloudngfwClientCreateProductSerialNumberResponse contains the response from method PaloAltoNetworksCloudngfwClient.CreateProductSerialNumber.
type PaloAltoNetworksCloudngfwClientCreateProductSerialNumberResponse struct {
	// Create Product Serial Number Request status
	ProductSerialNumberRequestStatus
}

// PaloAltoNetworksCloudngfwClientListCloudManagerTenantsResponse contains the response from method PaloAltoNetworksCloudngfwClient.ListCloudManagerTenants.
type PaloAltoNetworksCloudngfwClientListCloudManagerTenantsResponse struct {
	// Cloud Manager Tenant
	CloudManagerTenantList
}

// PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusResponse contains the response from method PaloAltoNetworksCloudngfwClient.ListProductSerialNumberStatus.
type PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusResponse struct {
	// Product serial and status for the service
	ProductSerialNumberStatus
}

// PaloAltoNetworksCloudngfwClientListSupportInfoResponse contains the response from method PaloAltoNetworksCloudngfwClient.ListSupportInfo.
type PaloAltoNetworksCloudngfwClientListSupportInfoResponse struct {
	// Support information for the service
	SupportInfoModel
}

// PostRulesClientCreateOrUpdateResponse contains the response from method PostRulesClient.BeginCreateOrUpdate.
type PostRulesClientCreateOrUpdateResponse struct {
	// PostRulestack rule list
	PostRulesResource
}

// PostRulesClientDeleteResponse contains the response from method PostRulesClient.BeginDelete.
type PostRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// PostRulesClientGetCountersResponse contains the response from method PostRulesClient.GetCounters.
type PostRulesClientGetCountersResponse struct {
	// Rule counter
	RuleCounter
}

// PostRulesClientGetResponse contains the response from method PostRulesClient.Get.
type PostRulesClientGetResponse struct {
	// PostRulestack rule list
	PostRulesResource
}

// PostRulesClientListResponse contains the response from method PostRulesClient.NewListPager.
type PostRulesClientListResponse struct {
	// The response of a PostRulesResource list operation.
	PostRulesResourceListResult
}

// PostRulesClientRefreshCountersResponse contains the response from method PostRulesClient.RefreshCounters.
type PostRulesClientRefreshCountersResponse struct {
	// placeholder for future response values
}

// PostRulesClientResetCountersResponse contains the response from method PostRulesClient.ResetCounters.
type PostRulesClientResetCountersResponse struct {
	// Rule counter reset
	RuleCounterReset
}

// PreRulesClientCreateOrUpdateResponse contains the response from method PreRulesClient.BeginCreateOrUpdate.
type PreRulesClientCreateOrUpdateResponse struct {
	// PreRulestack rule list
	PreRulesResource
}

// PreRulesClientDeleteResponse contains the response from method PreRulesClient.BeginDelete.
type PreRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// PreRulesClientGetCountersResponse contains the response from method PreRulesClient.GetCounters.
type PreRulesClientGetCountersResponse struct {
	// Rule counter
	RuleCounter
}

// PreRulesClientGetResponse contains the response from method PreRulesClient.Get.
type PreRulesClientGetResponse struct {
	// PreRulestack rule list
	PreRulesResource
}

// PreRulesClientListResponse contains the response from method PreRulesClient.NewListPager.
type PreRulesClientListResponse struct {
	// The response of a PreRulesResource list operation.
	PreRulesResourceListResult
}

// PreRulesClientRefreshCountersResponse contains the response from method PreRulesClient.RefreshCounters.
type PreRulesClientRefreshCountersResponse struct {
	// placeholder for future response values
}

// PreRulesClientResetCountersResponse contains the response from method PreRulesClient.ResetCounters.
type PreRulesClientResetCountersResponse struct {
	// Rule counter reset
	RuleCounterReset
}

// PrefixListGlobalRulestackClientCreateOrUpdateResponse contains the response from method PrefixListGlobalRulestackClient.BeginCreateOrUpdate.
type PrefixListGlobalRulestackClientCreateOrUpdateResponse struct {
	// GlobalRulestack prefixList
	PrefixListGlobalRulestackResource
}

// PrefixListGlobalRulestackClientDeleteResponse contains the response from method PrefixListGlobalRulestackClient.BeginDelete.
type PrefixListGlobalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// PrefixListGlobalRulestackClientGetResponse contains the response from method PrefixListGlobalRulestackClient.Get.
type PrefixListGlobalRulestackClientGetResponse struct {
	// GlobalRulestack prefixList
	PrefixListGlobalRulestackResource
}

// PrefixListGlobalRulestackClientListResponse contains the response from method PrefixListGlobalRulestackClient.NewListPager.
type PrefixListGlobalRulestackClientListResponse struct {
	// The response of a PrefixListGlobalRulestackResource list operation.
	PrefixListGlobalRulestackResourceListResult
}

// PrefixListLocalRulestackClientCreateOrUpdateResponse contains the response from method PrefixListLocalRulestackClient.BeginCreateOrUpdate.
type PrefixListLocalRulestackClientCreateOrUpdateResponse struct {
	// LocalRulestack prefixList
	PrefixListResource
}

// PrefixListLocalRulestackClientDeleteResponse contains the response from method PrefixListLocalRulestackClient.BeginDelete.
type PrefixListLocalRulestackClientDeleteResponse struct {
	// placeholder for future response values
}

// PrefixListLocalRulestackClientGetResponse contains the response from method PrefixListLocalRulestackClient.Get.
type PrefixListLocalRulestackClientGetResponse struct {
	// LocalRulestack prefixList
	PrefixListResource
}

// PrefixListLocalRulestackClientListByLocalRulestacksResponse contains the response from method PrefixListLocalRulestackClient.NewListByLocalRulestacksPager.
type PrefixListLocalRulestackClientListByLocalRulestacksResponse struct {
	// The response of a PrefixListResource list operation.
	PrefixListResourceListResult
}
