//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

// AFDCustomDomainsClientCreateResponse contains the response from method AFDCustomDomainsClient.BeginCreate.
type AFDCustomDomainsClientCreateResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	AFDDomain
}

// AFDCustomDomainsClientDeleteResponse contains the response from method AFDCustomDomainsClient.BeginDelete.
type AFDCustomDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDCustomDomainsClientGetResponse contains the response from method AFDCustomDomainsClient.Get.
type AFDCustomDomainsClientGetResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	AFDDomain
}

// AFDCustomDomainsClientListByProfileResponse contains the response from method AFDCustomDomainsClient.NewListByProfilePager.
type AFDCustomDomainsClientListByProfileResponse struct {
	// Result of the request to list domains. It contains a list of domain objects and a URL link to get the next set of results.
	AFDDomainListResult
}

// AFDCustomDomainsClientRefreshValidationTokenResponse contains the response from method AFDCustomDomainsClient.BeginRefreshValidationToken.
type AFDCustomDomainsClientRefreshValidationTokenResponse struct {
	// placeholder for future response values
}

// AFDCustomDomainsClientUpdateResponse contains the response from method AFDCustomDomainsClient.BeginUpdate.
type AFDCustomDomainsClientUpdateResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	AFDDomain
}

// AFDEndpointsClientCreateResponse contains the response from method AFDEndpointsClient.BeginCreate.
type AFDEndpointsClientCreateResponse struct {
	// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as
	// origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
	AFDEndpoint
}

// AFDEndpointsClientDeleteResponse contains the response from method AFDEndpointsClient.BeginDelete.
type AFDEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDEndpointsClientGetResponse contains the response from method AFDEndpointsClient.Get.
type AFDEndpointsClientGetResponse struct {
	// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as
	// origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
	AFDEndpoint
}

// AFDEndpointsClientListByProfileResponse contains the response from method AFDEndpointsClient.NewListByProfilePager.
type AFDEndpointsClientListByProfileResponse struct {
	// Result of the request to list endpoints. It contains a list of endpoint objects and a URL link to get the next set of results.
	AFDEndpointListResult
}

// AFDEndpointsClientListResourceUsageResponse contains the response from method AFDEndpointsClient.NewListResourceUsagePager.
type AFDEndpointsClientListResourceUsageResponse struct {
	// The list usages operation response.
	UsagesListResult
}

// AFDEndpointsClientPurgeContentResponse contains the response from method AFDEndpointsClient.BeginPurgeContent.
type AFDEndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// AFDEndpointsClientUpdateResponse contains the response from method AFDEndpointsClient.BeginUpdate.
type AFDEndpointsClientUpdateResponse struct {
	// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as
	// origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
	AFDEndpoint
}

// AFDEndpointsClientValidateCustomDomainResponse contains the response from method AFDEndpointsClient.ValidateCustomDomain.
type AFDEndpointsClientValidateCustomDomainResponse struct {
	// Output of custom domain validation.
	ValidateCustomDomainOutput
}

// AFDOriginGroupsClientCreateResponse contains the response from method AFDOriginGroupsClient.BeginCreate.
type AFDOriginGroupsClientCreateResponse struct {
	// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure
	// Front Door.
	AFDOriginGroup
}

// AFDOriginGroupsClientDeleteResponse contains the response from method AFDOriginGroupsClient.BeginDelete.
type AFDOriginGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDOriginGroupsClientGetResponse contains the response from method AFDOriginGroupsClient.Get.
type AFDOriginGroupsClientGetResponse struct {
	// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure
	// Front Door.
	AFDOriginGroup
}

// AFDOriginGroupsClientListByProfileResponse contains the response from method AFDOriginGroupsClient.NewListByProfilePager.
type AFDOriginGroupsClientListByProfileResponse struct {
	// Result of the request to list origin groups. It contains a list of origin groups objects and a URL link to get the next
	// set of results.
	AFDOriginGroupListResult
}

// AFDOriginGroupsClientListResourceUsageResponse contains the response from method AFDOriginGroupsClient.NewListResourceUsagePager.
type AFDOriginGroupsClientListResourceUsageResponse struct {
	// The list usages operation response.
	UsagesListResult
}

// AFDOriginGroupsClientUpdateResponse contains the response from method AFDOriginGroupsClient.BeginUpdate.
type AFDOriginGroupsClientUpdateResponse struct {
	// AFDOrigin group comprising of origins is used for load balancing to origins when the content cannot be served from Azure
	// Front Door.
	AFDOriginGroup
}

// AFDOriginsClientCreateResponse contains the response from method AFDOriginsClient.BeginCreate.
type AFDOriginsClientCreateResponse struct {
	// Azure Front Door origin is the source of the content being delivered via Azure Front Door. When the edge nodes represented
	// by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	AFDOrigin
}

// AFDOriginsClientDeleteResponse contains the response from method AFDOriginsClient.BeginDelete.
type AFDOriginsClientDeleteResponse struct {
	// placeholder for future response values
}

// AFDOriginsClientGetResponse contains the response from method AFDOriginsClient.Get.
type AFDOriginsClientGetResponse struct {
	// Azure Front Door origin is the source of the content being delivered via Azure Front Door. When the edge nodes represented
	// by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	AFDOrigin
}

// AFDOriginsClientListByOriginGroupResponse contains the response from method AFDOriginsClient.NewListByOriginGroupPager.
type AFDOriginsClientListByOriginGroupResponse struct {
	// Result of the request to list origins. It contains a list of origin objects and a URL link to get the next set of results.
	AFDOriginListResult
}

// AFDOriginsClientUpdateResponse contains the response from method AFDOriginsClient.BeginUpdate.
type AFDOriginsClientUpdateResponse struct {
	// Azure Front Door origin is the source of the content being delivered via Azure Front Door. When the edge nodes represented
	// by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	AFDOrigin
}

// AFDProfilesClientCheckEndpointNameAvailabilityResponse contains the response from method AFDProfilesClient.CheckEndpointNameAvailability.
type AFDProfilesClientCheckEndpointNameAvailabilityResponse struct {
	// Output of check name availability API.
	CheckEndpointNameAvailabilityOutput
}

// AFDProfilesClientCheckHostNameAvailabilityResponse contains the response from method AFDProfilesClient.CheckHostNameAvailability.
type AFDProfilesClientCheckHostNameAvailabilityResponse struct {
	// Output of check name availability API.
	CheckNameAvailabilityOutput
}

// AFDProfilesClientListResourceUsageResponse contains the response from method AFDProfilesClient.NewListResourceUsagePager.
type AFDProfilesClientListResourceUsageResponse struct {
	// The list usages operation response.
	UsagesListResult
}

// AFDProfilesClientUpgradeResponse contains the response from method AFDProfilesClient.BeginUpgrade.
type AFDProfilesClientUpgradeResponse struct {
	// A profile is a logical grouping of endpoints that share the same settings.
	Profile
}

// AFDProfilesClientValidateSecretResponse contains the response from method AFDProfilesClient.ValidateSecret.
type AFDProfilesClientValidateSecretResponse struct {
	// Output of the validated secret.
	ValidateSecretOutput
}

// CustomDomainsClientCreateResponse contains the response from method CustomDomainsClient.BeginCreate.
type CustomDomainsClientCreateResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	CustomDomain
}

// CustomDomainsClientDeleteResponse contains the response from method CustomDomainsClient.BeginDelete.
type CustomDomainsClientDeleteResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	CustomDomain
}

// CustomDomainsClientDisableCustomHTTPSResponse contains the response from method CustomDomainsClient.BeginDisableCustomHTTPS.
type CustomDomainsClientDisableCustomHTTPSResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	CustomDomain
}

// CustomDomainsClientEnableCustomHTTPSResponse contains the response from method CustomDomainsClient.BeginEnableCustomHTTPS.
type CustomDomainsClientEnableCustomHTTPSResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	CustomDomain
}

// CustomDomainsClientGetResponse contains the response from method CustomDomainsClient.Get.
type CustomDomainsClientGetResponse struct {
	// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
	CustomDomain
}

// CustomDomainsClientListByEndpointResponse contains the response from method CustomDomainsClient.NewListByEndpointPager.
type CustomDomainsClientListByEndpointResponse struct {
	// Result of the request to list custom domains. It contains a list of custom domain objects and a URL link to get the next
	// set of results.
	CustomDomainListResult
}

// EdgeNodesClientListResponse contains the response from method EdgeNodesClient.NewListPager.
type EdgeNodesClientListResponse struct {
	// Result of the request to list CDN edgenodes. It contains a list of ip address group and a URL link to get the next set
	// of results.
	EdgenodeResult
}

// EndpointsClientCreateResponse contains the response from method EndpointsClient.BeginCreate.
type EndpointsClientCreateResponse struct {
	// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content
	// caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
	Endpoint
}

// EndpointsClientDeleteResponse contains the response from method EndpointsClient.BeginDelete.
type EndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// EndpointsClientGetResponse contains the response from method EndpointsClient.Get.
type EndpointsClientGetResponse struct {
	// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content
	// caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
	Endpoint
}

// EndpointsClientListByProfileResponse contains the response from method EndpointsClient.NewListByProfilePager.
type EndpointsClientListByProfileResponse struct {
	// Result of the request to list endpoints. It contains a list of endpoint objects and a URL link to get the next set of results.
	EndpointListResult
}

// EndpointsClientListResourceUsageResponse contains the response from method EndpointsClient.NewListResourceUsagePager.
type EndpointsClientListResourceUsageResponse struct {
	// Output of check resource usage API.
	ResourceUsageListResult
}

// EndpointsClientLoadContentResponse contains the response from method EndpointsClient.BeginLoadContent.
type EndpointsClientLoadContentResponse struct {
	// placeholder for future response values
}

// EndpointsClientPurgeContentResponse contains the response from method EndpointsClient.BeginPurgeContent.
type EndpointsClientPurgeContentResponse struct {
	// placeholder for future response values
}

// EndpointsClientStartResponse contains the response from method EndpointsClient.BeginStart.
type EndpointsClientStartResponse struct {
	// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content
	// caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
	Endpoint
}

// EndpointsClientStopResponse contains the response from method EndpointsClient.BeginStop.
type EndpointsClientStopResponse struct {
	// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content
	// caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
	Endpoint
}

// EndpointsClientUpdateResponse contains the response from method EndpointsClient.BeginUpdate.
type EndpointsClientUpdateResponse struct {
	// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content
	// caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
	Endpoint
}

// EndpointsClientValidateCustomDomainResponse contains the response from method EndpointsClient.ValidateCustomDomain.
type EndpointsClientValidateCustomDomainResponse struct {
	// Output of custom domain validation.
	ValidateCustomDomainOutput
}

// LogAnalyticsClientGetLogAnalyticsLocationsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsLocations.
type LogAnalyticsClientGetLogAnalyticsLocationsResponse struct {
	// Continents Response
	ContinentsResponse
}

// LogAnalyticsClientGetLogAnalyticsMetricsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsMetrics.
type LogAnalyticsClientGetLogAnalyticsMetricsResponse struct {
	// Metrics Response
	MetricsResponse
}

// LogAnalyticsClientGetLogAnalyticsRankingsResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsRankings.
type LogAnalyticsClientGetLogAnalyticsRankingsResponse struct {
	// Rankings Response
	RankingsResponse
}

// LogAnalyticsClientGetLogAnalyticsResourcesResponse contains the response from method LogAnalyticsClient.GetLogAnalyticsResources.
type LogAnalyticsClientGetLogAnalyticsResourcesResponse struct {
	// Resources Response
	ResourcesResponse
}

// LogAnalyticsClientGetWafLogAnalyticsMetricsResponse contains the response from method LogAnalyticsClient.GetWafLogAnalyticsMetrics.
type LogAnalyticsClientGetWafLogAnalyticsMetricsResponse struct {
	// Waf Metrics Response
	WafMetricsResponse
}

// LogAnalyticsClientGetWafLogAnalyticsRankingsResponse contains the response from method LogAnalyticsClient.GetWafLogAnalyticsRankings.
type LogAnalyticsClientGetWafLogAnalyticsRankingsResponse struct {
	// Waf Rankings Response
	WafRankingsResponse
}

// ManagedRuleSetsClientListResponse contains the response from method ManagedRuleSetsClient.NewListPager.
type ManagedRuleSetsClientListResponse struct {
	// List of managed rule set definitions available for use in a policy.
	ManagedRuleSetDefinitionList
}

// ManagementClientCheckEndpointNameAvailabilityResponse contains the response from method ManagementClient.CheckEndpointNameAvailability.
type ManagementClientCheckEndpointNameAvailabilityResponse struct {
	// Output of check name availability API.
	CheckEndpointNameAvailabilityOutput
}

// ManagementClientCheckNameAvailabilityResponse contains the response from method ManagementClient.CheckNameAvailability.
type ManagementClientCheckNameAvailabilityResponse struct {
	// Output of check name availability API.
	CheckNameAvailabilityOutput
}

// ManagementClientCheckNameAvailabilityWithSubscriptionResponse contains the response from method ManagementClient.CheckNameAvailabilityWithSubscription.
type ManagementClientCheckNameAvailabilityWithSubscriptionResponse struct {
	// Output of check name availability API.
	CheckNameAvailabilityOutput
}

// ManagementClientValidateProbeResponse contains the response from method ManagementClient.ValidateProbe.
type ManagementClientValidateProbeResponse struct {
	// Output of the validate probe API.
	ValidateProbeOutput
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list CDN operations. It contains a list of operations and a URL link to get the next set of results.
	OperationsListResult
}

// OriginGroupsClientCreateResponse contains the response from method OriginGroupsClient.BeginCreate.
type OriginGroupsClientCreateResponse struct {
	// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
	OriginGroup
}

// OriginGroupsClientDeleteResponse contains the response from method OriginGroupsClient.BeginDelete.
type OriginGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// OriginGroupsClientGetResponse contains the response from method OriginGroupsClient.Get.
type OriginGroupsClientGetResponse struct {
	// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
	OriginGroup
}

// OriginGroupsClientListByEndpointResponse contains the response from method OriginGroupsClient.NewListByEndpointPager.
type OriginGroupsClientListByEndpointResponse struct {
	// Result of the request to list origin groups. It contains a list of origin groups objects and a URL link to get the next
	// set of results.
	OriginGroupListResult
}

// OriginGroupsClientUpdateResponse contains the response from method OriginGroupsClient.BeginUpdate.
type OriginGroupsClientUpdateResponse struct {
	// Origin group comprising of origins is used for load balancing to origins when the content cannot be served from CDN.
	OriginGroup
}

// OriginsClientCreateResponse contains the response from method OriginsClient.BeginCreate.
type OriginsClientCreateResponse struct {
	// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not
	// have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	Origin
}

// OriginsClientDeleteResponse contains the response from method OriginsClient.BeginDelete.
type OriginsClientDeleteResponse struct {
	// placeholder for future response values
}

// OriginsClientGetResponse contains the response from method OriginsClient.Get.
type OriginsClientGetResponse struct {
	// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not
	// have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	Origin
}

// OriginsClientListByEndpointResponse contains the response from method OriginsClient.NewListByEndpointPager.
type OriginsClientListByEndpointResponse struct {
	// Result of the request to list origins. It contains a list of origin objects and a URL link to get the next set of results.
	OriginListResult
}

// OriginsClientUpdateResponse contains the response from method OriginsClient.BeginUpdate.
type OriginsClientUpdateResponse struct {
	// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not
	// have the requested content cached, they attempt to fetch it from one or more of the configured origins.
	Origin
}

// PoliciesClientCreateOrUpdateResponse contains the response from method PoliciesClient.BeginCreateOrUpdate.
type PoliciesClientCreateOrUpdateResponse struct {
	// Defines web application firewall policy for Azure CDN.
	WebApplicationFirewallPolicy
}

// PoliciesClientDeleteResponse contains the response from method PoliciesClient.Delete.
type PoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// PoliciesClientGetResponse contains the response from method PoliciesClient.Get.
type PoliciesClientGetResponse struct {
	// Defines web application firewall policy for Azure CDN.
	WebApplicationFirewallPolicy
}

// PoliciesClientListResponse contains the response from method PoliciesClient.NewListPager.
type PoliciesClientListResponse struct {
	// Defines a list of WebApplicationFirewallPolicies for Azure CDN. It contains a list of WebApplicationFirewallPolicy objects
	// and a URL link to get the next set of results.
	WebApplicationFirewallPolicyList
}

// PoliciesClientUpdateResponse contains the response from method PoliciesClient.BeginUpdate.
type PoliciesClientUpdateResponse struct {
	// Defines web application firewall policy for Azure CDN.
	WebApplicationFirewallPolicy
}

// ProfilesClientCanMigrateResponse contains the response from method ProfilesClient.BeginCanMigrate.
type ProfilesClientCanMigrateResponse struct {
	// Result for canMigrate operation.
	CanMigrateResult
}

// ProfilesClientCreateResponse contains the response from method ProfilesClient.BeginCreate.
type ProfilesClientCreateResponse struct {
	// A profile is a logical grouping of endpoints that share the same settings.
	Profile
}

// ProfilesClientDeleteResponse contains the response from method ProfilesClient.BeginDelete.
type ProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProfilesClientGenerateSsoURIResponse contains the response from method ProfilesClient.GenerateSsoURI.
type ProfilesClientGenerateSsoURIResponse struct {
	// The URI required to login to the supplemental portal from the Azure portal.
	SsoURI
}

// ProfilesClientGetResponse contains the response from method ProfilesClient.Get.
type ProfilesClientGetResponse struct {
	// A profile is a logical grouping of endpoints that share the same settings.
	Profile
}

// ProfilesClientListByResourceGroupResponse contains the response from method ProfilesClient.NewListByResourceGroupPager.
type ProfilesClientListByResourceGroupResponse struct {
	// Result of the request to list profiles. It contains a list of profile objects and a URL link to get the next set of results.
	ProfileListResult
}

// ProfilesClientListResourceUsageResponse contains the response from method ProfilesClient.NewListResourceUsagePager.
type ProfilesClientListResourceUsageResponse struct {
	// Output of check resource usage API.
	ResourceUsageListResult
}

// ProfilesClientListResponse contains the response from method ProfilesClient.NewListPager.
type ProfilesClientListResponse struct {
	// Result of the request to list profiles. It contains a list of profile objects and a URL link to get the next set of results.
	ProfileListResult
}

// ProfilesClientListSupportedOptimizationTypesResponse contains the response from method ProfilesClient.ListSupportedOptimizationTypes.
type ProfilesClientListSupportedOptimizationTypesResponse struct {
	// The result of the GetSupportedOptimizationTypes API
	SupportedOptimizationTypesListResult
}

// ProfilesClientMigrateResponse contains the response from method ProfilesClient.BeginMigrate.
type ProfilesClientMigrateResponse struct {
	// Result for migrate operation.
	MigrateResult
}

// ProfilesClientMigrationCommitResponse contains the response from method ProfilesClient.BeginMigrationCommit.
type ProfilesClientMigrationCommitResponse struct {
	// placeholder for future response values
}

// ProfilesClientUpdateResponse contains the response from method ProfilesClient.BeginUpdate.
type ProfilesClientUpdateResponse struct {
	// A profile is a logical grouping of endpoints that share the same settings.
	Profile
}

// ResourceUsageClientListResponse contains the response from method ResourceUsageClient.NewListPager.
type ResourceUsageClientListResponse struct {
	// Output of check resource usage API.
	ResourceUsageListResult
}

// RoutesClientCreateResponse contains the response from method RoutesClient.BeginCreate.
type RoutesClientCreateResponse struct {
	// Friendly Routes name mapping to the any Routes or secret related information.
	Route
}

// RoutesClientDeleteResponse contains the response from method RoutesClient.BeginDelete.
type RoutesClientDeleteResponse struct {
	// placeholder for future response values
}

// RoutesClientGetResponse contains the response from method RoutesClient.Get.
type RoutesClientGetResponse struct {
	// Friendly Routes name mapping to the any Routes or secret related information.
	Route
}

// RoutesClientListByEndpointResponse contains the response from method RoutesClient.NewListByEndpointPager.
type RoutesClientListByEndpointResponse struct {
	// Result of the request to list routes. It contains a list of route objects and a URL link to get the next set of results.
	RouteListResult
}

// RoutesClientUpdateResponse contains the response from method RoutesClient.BeginUpdate.
type RoutesClientUpdateResponse struct {
	// Friendly Routes name mapping to the any Routes or secret related information.
	Route
}

// RuleSetsClientCreateResponse contains the response from method RuleSetsClient.Create.
type RuleSetsClientCreateResponse struct {
	// Friendly RuleSet name mapping to the any RuleSet or secret related information.
	RuleSet
}

// RuleSetsClientDeleteResponse contains the response from method RuleSetsClient.BeginDelete.
type RuleSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// RuleSetsClientGetResponse contains the response from method RuleSetsClient.Get.
type RuleSetsClientGetResponse struct {
	// Friendly RuleSet name mapping to the any RuleSet or secret related information.
	RuleSet
}

// RuleSetsClientListByProfileResponse contains the response from method RuleSetsClient.NewListByProfilePager.
type RuleSetsClientListByProfileResponse struct {
	// Result of the request to list rule sets. It contains a list of rule set objects and a URL link to get the next set of results.
	RuleSetListResult
}

// RuleSetsClientListResourceUsageResponse contains the response from method RuleSetsClient.NewListResourceUsagePager.
type RuleSetsClientListResourceUsageResponse struct {
	// The list usages operation response.
	UsagesListResult
}

// RulesClientCreateResponse contains the response from method RulesClient.BeginCreate.
type RulesClientCreateResponse struct {
	// Friendly Rules name mapping to the any Rules or secret related information.
	Rule
}

// RulesClientDeleteResponse contains the response from method RulesClient.BeginDelete.
type RulesClientDeleteResponse struct {
	// placeholder for future response values
}

// RulesClientGetResponse contains the response from method RulesClient.Get.
type RulesClientGetResponse struct {
	// Friendly Rules name mapping to the any Rules or secret related information.
	Rule
}

// RulesClientListByRuleSetResponse contains the response from method RulesClient.NewListByRuleSetPager.
type RulesClientListByRuleSetResponse struct {
	// Result of the request to list rules. It contains a list of rule objects and a URL link to get the next set of results.
	RuleListResult
}

// RulesClientUpdateResponse contains the response from method RulesClient.BeginUpdate.
type RulesClientUpdateResponse struct {
	// Friendly Rules name mapping to the any Rules or secret related information.
	Rule
}

// SecretsClientCreateResponse contains the response from method SecretsClient.BeginCreate.
type SecretsClientCreateResponse struct {
	// Friendly Secret name mapping to the any Secret or secret related information.
	Secret
}

// SecretsClientDeleteResponse contains the response from method SecretsClient.BeginDelete.
type SecretsClientDeleteResponse struct {
	// placeholder for future response values
}

// SecretsClientGetResponse contains the response from method SecretsClient.Get.
type SecretsClientGetResponse struct {
	// Friendly Secret name mapping to the any Secret or secret related information.
	Secret
}

// SecretsClientListByProfileResponse contains the response from method SecretsClient.NewListByProfilePager.
type SecretsClientListByProfileResponse struct {
	// Result of the request to list secrets. It contains a list of Secret objects and a URL link to get the next set of results.
	SecretListResult
}

// SecurityPoliciesClientCreateResponse contains the response from method SecurityPoliciesClient.BeginCreate.
type SecurityPoliciesClientCreateResponse struct {
	// SecurityPolicy association for AzureFrontDoor profile
	SecurityPolicy
}

// SecurityPoliciesClientDeleteResponse contains the response from method SecurityPoliciesClient.BeginDelete.
type SecurityPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SecurityPoliciesClientGetResponse contains the response from method SecurityPoliciesClient.Get.
type SecurityPoliciesClientGetResponse struct {
	// SecurityPolicy association for AzureFrontDoor profile
	SecurityPolicy
}

// SecurityPoliciesClientListByProfileResponse contains the response from method SecurityPoliciesClient.NewListByProfilePager.
type SecurityPoliciesClientListByProfileResponse struct {
	// Result of the request to list security policies. It contains a list of security policy objects and a URL link to get the
	// next set of results.
	SecurityPolicyListResult
}

// SecurityPoliciesClientPatchResponse contains the response from method SecurityPoliciesClient.BeginPatch.
type SecurityPoliciesClientPatchResponse struct {
	// SecurityPolicy association for AzureFrontDoor profile
	SecurityPolicy
}
