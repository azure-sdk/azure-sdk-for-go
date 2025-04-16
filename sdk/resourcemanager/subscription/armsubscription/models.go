// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import "time"

// AcceptOwnershipRequest - The parameters required to accept subscription ownership.
type AcceptOwnershipRequest struct {
	// Accept subscription ownership request properties.
	Properties *AcceptOwnershipRequestProperties
}

// AcceptOwnershipRequestProperties - Accept subscription ownership request properties.
type AcceptOwnershipRequestProperties struct {
	// REQUIRED; The friendly name of the subscription.
	DisplayName *string

	// Management group Id for the subscription.
	ManagementGroupID *string

	// Tags for the subscription
	Tags map[string]*string
}

// AcceptOwnershipStatusResponse - Subscription Accept Ownership Response
type AcceptOwnershipStatusResponse struct {
	// The display name of the subscription.
	DisplayName *string

	// Tenant Id of the subscription
	SubscriptionTenantID *string

	// Tags for the subscription
	Tags map[string]*string

	// READ-ONLY; The accept ownership state of the resource.
	AcceptOwnershipState *AcceptOwnership

	// READ-ONLY; UPN of the billing owner
	BillingOwner *string

	// READ-ONLY; The provisioning state of the resource.
	ProvisioningState *Provisioning

	// READ-ONLY; Newly created subscription Id.
	SubscriptionID *string
}

// AliasListResult - The list of aliases.
type AliasListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The list of alias.
	Value []*AliasResponse
}

// AliasResponse - Subscription Information with the alias.
type AliasResponse struct {
	// Subscription Alias response properties.
	Properties *AliasResponseProperties

	// READ-ONLY; Fully qualified ID for the alias resource.
	ID *string

	// READ-ONLY; Alias ID.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type, Microsoft.Subscription/aliases.
	Type *string
}

// AliasResponseProperties - Put subscription creation result properties.
type AliasResponseProperties struct {
	// Billing scope of the subscription. For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed -
	// /billingAccounts/{billingAccountName}/customers/{customerName} For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string

	// Created Time
	CreatedTime *string

	// The display name of the subscription.
	DisplayName *string

	// The Management Group Id.
	ManagementGroupID *string

	// The provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// Reseller Id
	ResellerID *string

	// Owner Id of the subscription
	SubscriptionOwnerID *string

	// Tags for the subscription
	Tags map[string]*string

	// The workload type of the subscription. It can be either Production or DevTest.
	Workload *Workload

	// READ-ONLY; The accept ownership state of the resource.
	AcceptOwnershipState *AcceptOwnership

	// READ-ONLY; Url to accept ownership of the subscription.
	AcceptOwnershipURL *string

	// READ-ONLY; Newly created subscription Id.
	SubscriptionID *string
}

// BillingAccountPoliciesResponse - Billing account policies information.
type BillingAccountPoliciesResponse struct {
	// Billing account policies response properties.
	Properties *BillingAccountPoliciesResponseProperties

	// READ-ONLY; Fully qualified ID for the policy.
	ID *string

	// READ-ONLY; Policy name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// BillingAccountPoliciesResponseProperties - Put billing account policies response properties.
type BillingAccountPoliciesResponseProperties struct {
	// Determine if the transfers are allowed for the billing account
	AllowTransfers *bool

	// Service tenant for the billing account.
	ServiceTenants []*ServiceTenantResponse
}

// CanceledSubscriptionID - The ID of the canceled subscription
type CanceledSubscriptionID struct {
	// READ-ONLY; The ID of the canceled subscription
	SubscriptionID *string
}

// CreationResult - The created subscription object.
type CreationResult struct {
	// The link to the new subscription. Use this link to check the status of subscription creation operation.
	SubscriptionLink *string
}

// EnabledSubscriptionID - The ID of the subscriptions that is being enabled
type EnabledSubscriptionID struct {
	// READ-ONLY; The ID of the subscriptions that is being enabled
	SubscriptionID *string
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Error code
	Code *string

	// Error message indicating why the operation failed.
	Message *string
}

// ErrorResponseBody - Error response indicates that the service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponseBody struct {
	// Error code
	Code *string

	// The details of the error.
	Error *ErrorResponse

	// Error message indicating why the operation failed.
	Message *string
}

// GetTenantPolicyListResponse - Tenant policy information list.
type GetTenantPolicyListResponse struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string

	// READ-ONLY; The list of tenant policies.
	Value []*GetTenantPolicyResponse
}

// GetTenantPolicyResponse - Tenant policy Information.
type GetTenantPolicyResponse struct {
	// Tenant policy properties.
	Properties *TenantPolicy

	// READ-ONLY; Policy Id.
	ID *string

	// READ-ONLY; Policy name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// ListResult - Subscription list operation response.
type ListResult struct {
	// The URL to get the next set of results.
	NextLink *string

	// An array of subscriptions.
	Value []*Subscription
}

// Location information.
type Location struct {
	// READ-ONLY; The display name of the location.
	DisplayName *string

	// READ-ONLY; The fully qualified ID of the location. For example, /subscriptions/00000000-0000-0000-0000-000000000000/locations/westus.
	ID *string

	// READ-ONLY; The latitude of the location.
	Latitude *string

	// READ-ONLY; The longitude of the location.
	Longitude *string

	// READ-ONLY; The location name.
	Name *string

	// READ-ONLY; The subscription ID.
	SubscriptionID *string
}

// LocationListResult - Location list operation response.
type LocationListResult struct {
	// An array of locations.
	Value []*Location
}

// Name - The new name of the subscription.
type Name struct {
	// New subscription name
	SubscriptionName *string
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Localized friendly description for the operation
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Subscription
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - Result of the request to list operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of operations.
	Value []*Operation
}

// Policies - Subscription policies.
type Policies struct {
	// READ-ONLY; The subscription location placement ID. The ID indicates which regions are visible for a subscription. For example,
	// a subscription with a location placement Id of Public_2014-09-01 has access to Azure
	// public regions.
	LocationPlacementID *string

	// READ-ONLY; The subscription quota ID.
	QuotaID *string

	// READ-ONLY; The subscription spending limit.
	SpendingLimit *SpendingLimit
}

// PutAliasRequest - The parameters required to create a new subscription.
type PutAliasRequest struct {
	// Put alias request properties.
	Properties *PutAliasRequestProperties
}

// PutAliasRequestAdditionalProperties - Put subscription additional properties.
type PutAliasRequestAdditionalProperties struct {
	// Management group Id for the subscription.
	ManagementGroupID *string

	// Owner Id of the subscription
	SubscriptionOwnerID *string

	// Tenant Id of the subscription
	SubscriptionTenantID *string

	// Tags for the subscription
	Tags map[string]*string
}

// PutAliasRequestProperties - Put subscription properties.
type PutAliasRequestProperties struct {
	// Put alias request additional properties.
	AdditionalProperties *PutAliasRequestAdditionalProperties

	// Billing scope of the subscription. For CustomerLed and FieldLed - /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed -
	// /billingAccounts/{billingAccountName}/customers/{customerName} For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string

	// The friendly name of the subscription.
	DisplayName *string

	// Reseller Id
	ResellerID *string

	// This parameter can be used to create alias for existing subscription Id
	SubscriptionID *string

	// The workload type of the subscription. It can be either Production or DevTest.
	Workload *Workload
}

// PutTenantPolicyRequestProperties - Put tenant policy request properties.
type PutTenantPolicyRequestProperties struct {
	// Blocks the entering of subscriptions into user's tenant.
	BlockSubscriptionsIntoTenant *bool

	// Blocks the leaving of subscriptions from user's tenant.
	BlockSubscriptionsLeavingTenant *bool

	// List of user objectIds that are exempted from the set subscription tenant policies for the user's tenant.
	ExemptedPrincipals []*string
}

// RenamedSubscriptionID - The ID of the subscriptions that is being renamed
type RenamedSubscriptionID struct {
	// READ-ONLY; The ID of the subscriptions that is being renamed
	SubscriptionID *string
}

// ServiceTenantResponse - Billing account service tenant.
type ServiceTenantResponse struct {
	// Service tenant id.
	TenantID *string

	// Service tenant name.
	TenantName *string
}

// Subscription information.
type Subscription struct {
	// The authorization source of the request. Valid values are one or more combinations of Legacy, RoleBased, Bypassed, Direct
	// and Management. For example, 'Legacy, RoleBased'.
	AuthorizationSource *string

	// The subscription policies.
	SubscriptionPolicies *Policies

	// Tags for the subscription
	Tags map[string]*string

	// READ-ONLY; The subscription display name.
	DisplayName *string

	// READ-ONLY; The fully qualified ID for the subscription. For example, /subscriptions/00000000-0000-0000-0000-000000000000.
	ID *string

	// READ-ONLY; The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
	State *SubscriptionState

	// READ-ONLY; The subscription ID.
	SubscriptionID *string

	// READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TenantIDDescription - Tenant Id information.
type TenantIDDescription struct {
	// READ-ONLY; The country/region name of the address for the tenant.
	Country *string

	// READ-ONLY; The Country/region abbreviation for the tenant.
	CountryCode *string

	// READ-ONLY; The default domain for the tenant.
	DefaultDomain *string

	// READ-ONLY; The display name of the tenant.
	DisplayName *string

	// READ-ONLY; The list of domains for the tenant.
	Domains *string

	// READ-ONLY; The fully qualified ID of the tenant. For example, /tenants/00000000-0000-0000-0000-000000000000.
	ID *string

	// READ-ONLY; The category of the tenant. Possible values are TenantCategoryHome,TenantCategoryProjectedBy,TenantCategoryManagedBy
	TenantCategory *string

	// READ-ONLY; The tenant ID. For example, 00000000-0000-0000-0000-000000000000.
	TenantID *string

	// READ-ONLY; The tenant type. Only available for Home tenant category.
	TenantType *string
}

// TenantListResult - Tenant Ids information.
type TenantListResult struct {
	// REQUIRED; The URL to use for getting the next set of results.
	NextLink *string

	// An array of tenants.
	Value []*TenantIDDescription
}

// TenantPolicy - Tenant policy.
type TenantPolicy struct {
	// Blocks the entering of subscriptions into user's tenant.
	BlockSubscriptionsIntoTenant *bool

	// Blocks the leaving of subscriptions from user's tenant.
	BlockSubscriptionsLeavingTenant *bool

	// List of user objectIds that are exempted from the set subscription tenant policies for the user's tenant.
	ExemptedPrincipals []*string

	// READ-ONLY; Policy Id.
	PolicyID *string
}
