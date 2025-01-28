//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

import "time"

// Account - An Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The SKU of this account.
	SKU *SKU

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Get or Set Kind property.
	Kind *Kind

	// The map account properties.
	Properties *AccountProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AccountKeys - The set of keys which can be used to access the Maps REST APIs. Two keys are provided for key rotation without
// interruption.
type AccountKeys struct {
	// READ-ONLY; The primary key for accessing the Maps REST APIs.
	PrimaryKey *string

	// READ-ONLY; The last updated date and time of the primary key.
	PrimaryKeyLastUpdated *string

	// READ-ONLY; The secondary key for accessing the Maps REST APIs.
	SecondaryKey *string

	// READ-ONLY; The last updated date and time of the secondary key.
	SecondaryKeyLastUpdated *string
}

// AccountProperties - Additional Maps account properties
type AccountProperties struct {
	// Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no CorsRule
	// elements are included in the request body, all CORS rules will be deleted, and
	// CORS will be disabled for the Blob service.
	Cors *CorsRules

	// Allows toggle functionality on Azure Policy to disable Azure Maps local authentication support. This will disable Shared
	// Keys and Shared Access Signature Token authentication from any usage.
	DisableLocalAuth *bool

	// All encryption configuration for a resource.
	Encryption *Encryption

	// The array of associated resources to the Maps account. Linked resource in the array cannot individually update, you must
	// update all linked resources in the array together. These resources may be used
	// on operations on the Azure Maps REST API. Access is controlled by the Maps Account Managed Identity(s) permissions to those
	// resource(s).
	LinkedResources []*LinkedResource

	// List of additional data processing regions for the Maps Account, which may result in requests being processed in another
	// geography. Some features or results may be restricted to specific regions. By
	// default, Maps REST APIs process requests according to the account location or the geographic scope [https://learn.microsoft.com/azure/azure-maps/geographic-scope].
	Locations []*LocationsItem

	// READ-ONLY; The provisioning state of the Maps account resource, Account updates can only be performed on terminal states.
	// Terminal states: Succeeded and Failed
	ProvisioningState *string

	// READ-ONLY; A unique identifier for the Maps Account
	UniqueID *string
}

// AccountSasParameters - Parameters used to create an account Shared Access Signature (SAS) token. The REST API access control
// is provided by Azure Maps Role Based Access (RBAC) identity and access.
type AccountSasParameters struct {
	// REQUIRED; The date time offset of when the token validity expires. For example "2017-05-24T10:42:03.1567373Z". Maximum
	// duration allowed is 24 hours between start and expiry.
	Expiry *string

	// REQUIRED; Required parameter which represents the desired maximum request per second to allowed for the given SAS token.
	// This does not guarantee perfect accuracy in measurements but provides application safe
	// guards of abuse with eventual enforcement.
	MaxRatePerSecond *int32

	// REQUIRED; The principal Id also known as the object Id of a User Assigned Managed Identity currently assigned to the Maps
	// Account. To assign a Managed Identity of the account, use operation Create or Update an
	// assign a User Assigned Identity resource Id.
	PrincipalID *string

	// REQUIRED; The Maps account key to use for signing. Picking primaryKey or secondaryKey will use the Maps account Shared
	// Keys, and using managedIdentity will use the auto-renewed private key to sign the SAS.
	SigningKey *SigningKey

	// REQUIRED; The date time offset of when the token validity begins. For example "2017-05-24T10:42:03.1567373Z". Maximum duration
	// allowed is 24 hours between start and expiry.
	Start *string

	// Optional, allows control of which region locations are permitted access to Azure Maps REST APIs with the SAS token. Example:
	// "eastus", "westus2". Omitting this parameter will allow all region
	// locations to be accessible.
	Regions []*string
}

// AccountSasToken - A new Sas token which can be used to access the Maps REST APIs and is controlled by the specified Managed
// identity permissions on Azure (IAM) Role Based Access Control.
type AccountSasToken struct {
	// READ-ONLY; The shared access signature access token.
	AccountSasToken *string
}

// AccountUpdateParameters - Parameters used to update an existing Maps Account.
type AccountUpdateParameters struct {
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// Get or Set Kind property.
	Kind *Kind

	// The map account properties.
	Properties *AccountProperties

	// The SKU of this account.
	SKU *SKU

	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a
	// resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string
}

// Accounts - A list of Maps Accounts.
type Accounts struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; a Maps Account.
	Value []*Account
}

// CorsRule - Specifies a CORS rule for the Map Account.
type CorsRule struct {
	// REQUIRED; Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*" to allow
	// all domains
	AllowedOrigins []*string
}

// CorsRules - Sets the CORS rules. You can include up to five CorsRule elements in the request.
type CorsRules struct {
	// The list of CORS rules. You can include up to five CorsRule elements in the request.
	CorsRules []*CorsRule
}

// Creator - An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type Creator struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The Creator resource properties.
	Properties *CreatorProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CreatorList - A list of Creator resources.
type CreatorList struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; a Creator account.
	Value []*Creator
}

// CreatorProperties - Creator resource properties
type CreatorProperties struct {
	// REQUIRED; The storage units to be allocated. Integer values from 1 to 100, inclusive.
	StorageUnits *int32

	// The consumed storage unit size in bytes for the creator resource.
	ConsumedStorageUnitSizeInBytes *int32

	// The total allocated storage unit size in bytes for the creator resource.
	TotalStorageUnitSizeInBytes *int32

	// READ-ONLY; The state of the resource provisioning, terminal states: Succeeded, Failed, Canceled
	ProvisioningState *string
}

// CreatorUpdateParameters - Parameters used to update an existing Creator resource.
type CreatorUpdateParameters struct {
	// Creator resource properties.
	Properties *CreatorProperties

	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a
	// resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string
}

// Encryption - All encryption configuration for a resource.
type Encryption struct {
	// All Customer-managed key encryption properties for the resource.
	CustomerManagedKeyEncryption *EncryptionCustomerManagedKeyEncryption

	// (Optional) Discouraged to include in resource definition. Only needed where it is possible to disable platform (AKA infrastructure)
	// encryption. Azure SQL TDE is an example of this. Values are enabled
	// and disabled.
	InfrastructureEncryption *InfrastructureEncryption
}

// EncryptionCustomerManagedKeyEncryption - All Customer-managed key encryption properties for the resource.
type EncryptionCustomerManagedKeyEncryption struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	KeyEncryptionKeyIdentity *EncryptionCustomerManagedKeyEncryptionKeyIdentity

	// key encryption key Url, versioned or unversioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	// or https://contosovault.vault.azure.net/keys/contosokek.
	KeyEncryptionKeyURL *string
}

// EncryptionCustomerManagedKeyEncryptionKeyIdentity - All identity configuration for Customer-managed key settings defining
// which identity should be used to auth to Key Vault.
type EncryptionCustomerManagedKeyEncryptionKeyIdentity struct {
	// delegated identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups//providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
	// Mutually exclusive with identityType systemAssignedIdentity and userAssignedIdentity - internal use only.
	DelegatedIdentityClientID *string

	// application client identity to use for accessing key encryption key Url in a different tenant. Ex: f83c6b1b-4d34-47e4-bb34-9d83df58b540
	FederatedClientID *string

	// The type of identity to use. Values can be systemAssignedIdentity, userAssignedIdentity, or delegatedResourceIdentity.
	IdentityType *EncryptionCustomerManagedKeyEncryptionKeyIdentityType

	// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/
	// /providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId. Mutually exclusive with identityType systemAssignedIdentity.
	UserAssignedIdentityResourceID *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// KeySpecification - Whether the operation refers to the primary or secondary key.
type KeySpecification struct {
	// REQUIRED; Whether the operation refers to the primary or secondary key.
	KeyType *KeyType
}

// LinkedResource - Linked resource is reference to a resource deployed in an Azure subscription, add the linked resource
// uniqueName value as an optional parameter for operations on Azure Maps Geospatial REST APIs.
type LinkedResource struct {
	// REQUIRED; ARM resource id in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/accounts/{storageName}'.
	ID *string

	// REQUIRED; A provided name which uniquely identifies the linked resource.
	UniqueName *string
}

// LocationsItem - Data processing location.
type LocationsItem struct {
	// REQUIRED; The location name.
	LocationName *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
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
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SKU - The SKU of the Maps Account.
type SKU struct {
	// REQUIRED; The name of the SKU, in standard format (such as G2).
	Name *Name

	// READ-ONLY; Gets the sku tier. This is based on the SKU name.
	Tier *string
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

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
