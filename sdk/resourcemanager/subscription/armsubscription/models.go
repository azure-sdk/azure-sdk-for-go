//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscription

// ClientBeginCancelOptions contains the optional parameters for the Client.BeginCancel method.
type ClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginEnableOptions contains the optional parameters for the Client.BeginEnable method.
type ClientBeginEnableOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginOperationResultsOptions contains the optional parameters for the Client.BeginOperationResults method.
type ClientBeginOperationResultsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ErrorResponseBody - Error response indicates that the service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponseBody struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// The details of the error.
	Error *ErrorResponse `json:"error,omitempty"`

	// Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ListResult - Subscription list operation response.
type ListResult struct {
	// The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of subscriptions.
	Value []*Subscription `json:"value,omitempty"`
}

// Location information.
type Location struct {
	// READ-ONLY; The display name of the location.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The latitude of the location.
	Latitude *string `json:"latitude,omitempty" azure:"ro"`

	// READ-ONLY; The longitude of the location.
	Longitude *string `json:"longitude,omitempty" azure:"ro"`

	// READ-ONLY; The location name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`
}

// LocationListResult - Location list operation response.
type LocationListResult struct {
	// An array of locations.
	Value []*Location `json:"value,omitempty"`
}

// Policies - Subscription policies.
type Policies struct {
	// READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example,
	// a subscription with a location placement Id of Public_2014-09-01 has access to Azure
	// public regions.
	LocationPlacementID *string `json:"locationPlacementId,omitempty" azure:"ro"`

	// READ-ONLY; The subscription quota ID.
	QuotaID *string `json:"quotaId,omitempty" azure:"ro"`

	// READ-ONLY; The subscription spending limit.
	SpendingLimit *SpendingLimit `json:"spendingLimit,omitempty" azure:"ro"`
}

// Subscription information.
type Subscription struct {
	// The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct
	// and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string `json:"authorizationSource,omitempty"`

	// The subscription policies.
	SubscriptionPolicies *Policies `json:"subscriptionPolicies,omitempty"`

	// Tags for the subscription
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The subscription display name.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State *SubscriptionState `json:"state,omitempty" azure:"ro"`

	// READ-ONLY; The subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// SubscriptionsClientGetOptions contains the optional parameters for the SubscriptionsClient.Get method.
type SubscriptionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionsClientListLocationsOptions contains the optional parameters for the SubscriptionsClient.NewListLocationsPager
// method.
type SubscriptionsClientListLocationsOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionsClientListOptions contains the optional parameters for the SubscriptionsClient.NewListPager method.
type SubscriptionsClientListOptions struct {
	// placeholder for future optional parameters
}

// TenantIDDescription - Tenant Id information.
type TenantIDDescription struct {
	// READ-ONLY; The country/region name of the address for the tenant.
	Country *string `json:"country,omitempty" azure:"ro"`

	// READ-ONLY; The Country/region abbreviation for the tenant.
	CountryCode *string `json:"countryCode,omitempty" azure:"ro"`

	// READ-ONLY; The default domain for the tenant.
	DefaultDomain *string `json:"defaultDomain,omitempty" azure:"ro"`

	// READ-ONLY; The display name of the tenant.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The list of domains for the tenant.
	Domains *string `json:"domains,omitempty" azure:"ro"`

	// READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The category of the tenant. Possible values are TenantCategoryHome,TenantCategoryProjectedBy,TenantCategoryManagedBy
	TenantCategory *string `json:"tenantCategory,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant type. Only available for Home tenant category.
	TenantType *string `json:"tenantType,omitempty" azure:"ro"`
}

// TenantListResult - Tenant Ids information.
type TenantListResult struct {
	// REQUIRED; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// An array of tenants.
	Value []*TenantIDDescription `json:"value,omitempty"`
}

// TenantsClientListOptions contains the optional parameters for the TenantsClient.NewListPager method.
type TenantsClientListOptions struct {
	// placeholder for future optional parameters
}
