//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvideoindexer

import "time"

// AccessToken - Azure Video Indexer access token.
type AccessToken struct {
	// READ-ONLY; The access token.
	AccessToken *string `json:"accessToken,omitempty" azure:"ro"`
}

// Account - An Azure Video Indexer account.
type Account struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// List of account properties
	Properties *AccountPropertiesForPutRequest `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountCheckNameAvailabilityParameters - The parameters used to check the availability of the Video Indexer account name.
type AccountCheckNameAvailabilityParameters struct {
	// REQUIRED; The VideoIndexer account name.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The type of resource, Microsoft.VideoIndexer/accounts
	Type *Type `json:"type,omitempty"`
}

// AccountList - The list operation response, that contains the data pools and their properties.
type AccountList struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; List of accounts and their properties.
	Value []*Account `json:"value,omitempty" azure:"ro"`
}

// AccountPatch - Azure Video Indexer account
type AccountPatch struct {
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// List of account properties
	Properties *AccountPropertiesForPatchRequest `json:"properties,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// AccountPropertiesForPatchRequest - Azure Video Indexer account properties
type AccountPropertiesForPatchRequest struct {
	// The media services details
	MediaServices *MediaServicesForPatchRequest `json:"mediaServices,omitempty"`

	// READ-ONLY; The account's data-plane ID
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; Gets the status of the account at the time the operation was called.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The account's tenant id
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// AccountPropertiesForPutRequest - Azure Video Indexer account properties
type AccountPropertiesForPutRequest struct {
	// The account's data-plane ID. This can be set only when connecting an existing classic account
	AccountID *string `json:"accountId,omitempty"`

	// The media services details
	MediaServices *MediaServicesForPutRequest `json:"mediaServices,omitempty"`

	// READ-ONLY; The account's name
	AccountName *string `json:"accountName,omitempty" azure:"ro"`

	// READ-ONLY; Gets the status of the account at the time the operation was called.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The account's tenant id
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`

	// READ-ONLY; An integer representing the total seconds that have been indexed on the account
	TotalSecondsIndexed *int32 `json:"totalSecondsIndexed,omitempty" azure:"ro"`
}

// AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
type AccountsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientCreateOrUpdateOptions contains the optional parameters for the AccountsClient.CreateOrUpdate method.
type AccountsClientCreateOrUpdateOptions struct {
	// The parameters to provide for the Azure Video Indexer account.
	Parameters *Account
}

// AccountsClientDeleteOptions contains the optional parameters for the AccountsClient.Delete method.
type AccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.ListByResourceGroup method.
type AccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListOptions contains the optional parameters for the AccountsClient.List method.
type AccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientUpdateOptions contains the optional parameters for the AccountsClient.Update method.
type AccountsClientUpdateOptions struct {
	// The parameters to provide for the current Azure Video Indexer account.
	Parameters *AccountPatch
}

// CheckNameAvailabilityResult - The CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	// READ-ONLY; Gets an error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Gets a boolean value that indicates whether the name is available for you to use. If true, the name is available.
	// If false, the name has already been taken.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; Gets the reason that a Video Indexer account name could not be used. The Reason element is only returned if
	// NameAvailable is false.
	Reason *Reason `json:"reason,omitempty" azure:"ro"`
}

// ClassicAccount - An Azure Video Indexer classic account.
type ClassicAccount struct {
	// The account's location
	Location *string `json:"location,omitempty"`

	// List of classic account properties
	Properties *ClassicAccountProperties `json:"properties,omitempty"`

	// READ-ONLY; The account's id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The account's name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// ClassicAccountMediaServices - Azure Video Indexer classic account properties
type ClassicAccountMediaServices struct {
	// The aad application id
	AADApplicationID *string `json:"aadApplicationId,omitempty"`

	// The aad tenant id
	AADTenantID *string `json:"aadTenantId,omitempty"`

	// Represents wether the media services is connected or not
	Connected *bool `json:"connected,omitempty"`

	// Represents if the media services event grid is connected or not
	EventGridProviderRegistered *bool `json:"eventGridProviderRegistered,omitempty"`

	// The media services name
	Name *string `json:"name,omitempty"`

	// The resource group that the media services belong to
	ResourceGroup *string `json:"resourceGroup,omitempty"`

	// Represents wether the media services streaming endpoint has started
	StreamingEndpointStarted *bool `json:"streamingEndpointStarted,omitempty"`

	// The media services subscriptionId
	SubscriptionID *string `json:"subscriptionId,omitempty"`
}

// ClassicAccountProperties - Azure Video Indexer classic account properties
type ClassicAccountProperties struct {
	// The media services details
	MediaServices *ClassicAccountMediaServices `json:"mediaServices,omitempty"`
}

// ClassicAccountSlim - An Azure Video Indexer classic account.
type ClassicAccountSlim struct {
	// READ-ONLY; The account's id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The account's location
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The account's name
	Name *string `json:"name,omitempty" azure:"ro"`
}

// ClassicAccountsClientGetDetailsOptions contains the optional parameters for the ClassicAccountsClient.GetDetails method.
type ClassicAccountsClientGetDetailsOptions struct {
	// placeholder for future optional parameters
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// GenerateAccessTokenParameters - Access token generation request's parameters
type GenerateAccessTokenParameters struct {
	// REQUIRED; The requested permission
	PermissionType *PermissionType `json:"permissionType,omitempty"`

	// REQUIRED; The requested media type
	Scope *Scope `json:"scope,omitempty"`

	// The project ID
	ProjectID *string `json:"projectId,omitempty"`

	// The video ID
	VideoID *string `json:"videoId,omitempty"`
}

// GenerateClientAccessTokenOptions contains the optional parameters for the GenerateClient.AccessToken method.
type GenerateClientAccessTokenOptions struct {
	// The parameters for generating access token
	Parameters *GenerateAccessTokenParameters
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MediaServicesForPatchRequest - The media services details
type MediaServicesForPatchRequest struct {
	// The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// MediaServicesForPutRequest - The media services details
type MediaServicesForPutRequest struct {
	// The media services resource id
	ResourceID *string `json:"resourceId,omitempty"`

	// The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

// Operation detail payload
type Operation struct {
	// READ-ONLY; Indicates the action type.
	ActionType *string `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Display of the operation
	Display *OperationDisplay `json:"display,omitempty" azure:"ro"`

	// READ-ONLY; Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; Name of the operation
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Origin of the operation
	Origin *string `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Operation display payload
type OperationDisplay struct {
	// READ-ONLY; Localized friendly description for the operation
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Localized friendly name for the operation
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Resource provider of the operation
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource of the operation
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - Available operations of the service.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the Resource Provider.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// Tags - Resource tags
type Tags struct {
	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}

// UserClassicAccountList - The list of user classic accounts.
type UserClassicAccountList struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; List of classic account names and their location.
	Value []*ClassicAccountSlim `json:"value,omitempty" azure:"ro"`
}

// UserClassicAccountsClientListOptions contains the optional parameters for the UserClassicAccountsClient.List method.
type UserClassicAccountsClientListOptions struct {
	// placeholder for future optional parameters
}
