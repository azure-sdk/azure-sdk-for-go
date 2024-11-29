// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armneonpostgres

import "time"

// CompanyDetails - Company details for an organization
type CompanyDetails struct {
// Business phone number of the company
	BusinessPhone *string

// Company name
	CompanyName *string

// Country name of the company
	Country *string

// Domain of the user
	Domain *string

// Number of employees in the company
	NumberOfEmployees *int64

// Office address of the company
	OfficeAddress *string
}

// MarketplaceDetails - Marketplace details for an organization
type MarketplaceDetails struct {
// REQUIRED; Offer details for the marketplace that is selected by the user
	OfferDetails *OfferDetails

// SaaS subscription id for the the marketplace offer
	SubscriptionID *string

// Marketplace subscription status
	SubscriptionStatus *MarketplaceSubscriptionStatus
}

// OfferDetails - Offer details for the marketplace that is selected by the user
type OfferDetails struct {
// REQUIRED; Offer Id for the marketplace offer
	OfferID *string

// REQUIRED; Plan Id for the marketplace offer
	PlanID *string

// REQUIRED; Publisher Id for the marketplace offer
	PublisherID *string

// Plan Name for the marketplace offer
	PlanName *string

// Term Id for the marketplace offer
	TermID *string

// Term Name for the marketplace offer
	TermUnit *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

// READ-ONLY; Localized display information for this particular operation.
	Display *OperationDisplay

// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
// Resource Manager/control-plane operations.
	IsDataAction *bool

// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
// Machine", "Restart Virtual Machine".
	Operation *string

// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
// Compute".
	Provider *string

// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
// REQUIRED; The Operation items on this page
	Value []*Operation

// The link to the next page of items
	NextLink *string
}

// OrganizationProperties - Properties specific to Data Organization resource
type OrganizationProperties struct {
// REQUIRED; Details of the company.
	CompanyDetails *CompanyDetails

// REQUIRED; Marketplace details of the resource.
	MarketplaceDetails *MarketplaceDetails

// REQUIRED; Details of the user.
	UserDetails *UserDetails

// Organization properties
	PartnerOrganizationProperties *PartnerOrganizationProperties

// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// OrganizationResource - Organization Resource by Neon
type OrganizationResource struct {
// REQUIRED; The geo-location where the resource lives
	Location *string

// The resource-specific properties for this resource.
	Properties *OrganizationProperties

// Resource tags.
	Tags map[string]*string

// READ-ONLY; Name of the Neon Organizations resource
	Name *string

// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OrganizationResourceListResult - The response of a OrganizationResource list operation.
type OrganizationResourceListResult struct {
// REQUIRED; The OrganizationResource items on this page
	Value []*OrganizationResource

// The link to the next page of items
	NextLink *string
}

// PartnerOrganizationProperties - Properties specific to Partner's organization
type PartnerOrganizationProperties struct {
// REQUIRED; Organization name in partner's system
	OrganizationName *string

// Organization Id in partner's system
	OrganizationID *string

// Single Sign On properties for the organization
	SingleSignOnProperties *SingleSignOnProperties
}

// SingleSignOnProperties - Properties specific to Single Sign On Resource
type SingleSignOnProperties struct {
// List of AAD domains fetched from Microsoft Graph for user.
	AADDomains []*string

// AAD enterprise application Id used to setup SSO
	EnterpriseAppID *string

// State of the Single Sign On for the organization
	SingleSignOnState *SingleSignOnStates

// URL for SSO to be used by the partner to redirect the user to their system
	SingleSignOnURL *string
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

// UserDetails - User details for an organization
type UserDetails struct {
// Email address of the user
	EmailAddress *string

// First name of the user
	FirstName *string

// Last name of the user
	LastName *string

// User's phone number
	PhoneNumber *string

// User's principal name
	Upn *string
}

