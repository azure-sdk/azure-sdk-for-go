// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabasefleetmanager

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasefleetmanager/armdatabasefleetmanager"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AzureProvisioningState - The provisioning state of the resource.
type AzureProvisioningState string

const (
	// AzureProvisioningStateAccepted - Request on the resource has been accepted.
	AzureProvisioningStateAccepted AzureProvisioningState = "Accepted"
	// AzureProvisioningStateCanceled - Resource creation was canceled.
	AzureProvisioningStateCanceled AzureProvisioningState = "Canceled"
	// AzureProvisioningStateFailed - Resource creation failed.
	AzureProvisioningStateFailed AzureProvisioningState = "Failed"
	// AzureProvisioningStateProvisioning - Resource is provisioning.
	AzureProvisioningStateProvisioning AzureProvisioningState = "Provisioning"
	// AzureProvisioningStateSucceeded - Resource has been created.
	AzureProvisioningStateSucceeded AzureProvisioningState = "Succeeded"
)

// PossibleAzureProvisioningStateValues returns the possible values for the AzureProvisioningState const type.
func PossibleAzureProvisioningStateValues() []AzureProvisioningState {
	return []AzureProvisioningState{
		AzureProvisioningStateAccepted,
		AzureProvisioningStateCanceled,
		AzureProvisioningStateFailed,
		AzureProvisioningStateProvisioning,
		AzureProvisioningStateSucceeded,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DatabaseCreateMode - Create mode. Available options: Default - Create a database. Copy - Copy the source database (source
// database name must be specified) PointInTimeRestore - Create a database by restoring source
// database from a point in time (source database name and restore from time must be specified)
type DatabaseCreateMode string

const (
	// DatabaseCreateModeCopy - Copy the source database (source database name must be specified).
	DatabaseCreateModeCopy DatabaseCreateMode = "Copy"
	// DatabaseCreateModeDefault - Create a database.
	DatabaseCreateModeDefault DatabaseCreateMode = "Default"
	// DatabaseCreateModePointInTimeRestore - Create a database by restoring source database from a point in time (source database
	// name and restore from time must be specified).
	DatabaseCreateModePointInTimeRestore DatabaseCreateMode = "PointInTimeRestore"
)

// PossibleDatabaseCreateModeValues returns the possible values for the DatabaseCreateMode const type.
func PossibleDatabaseCreateModeValues() []DatabaseCreateMode {
	return []DatabaseCreateMode{
		DatabaseCreateModeCopy,
		DatabaseCreateModeDefault,
		DatabaseCreateModePointInTimeRestore,
	}
}

// IdentityType - Identity type of the main principal.
type IdentityType string

const (
	// IdentityTypeNone - No identity.
	IdentityTypeNone IdentityType = "None"
	// IdentityTypeUserAssigned - User assigned identity.
	IdentityTypeUserAssigned IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeUserAssigned,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PrincipalType - Principal type of the authorized principal.
type PrincipalType string

const (
	// PrincipalTypeApplication - Application principal type.
	PrincipalTypeApplication PrincipalType = "Application"
	// PrincipalTypeUser - User principal type.
	PrincipalTypeUser PrincipalType = "User"
)

// PossiblePrincipalTypeValues returns the possible values for the PrincipalType const type.
func PossiblePrincipalTypeValues() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeApplication,
		PrincipalTypeUser,
	}
}

// ResourceType - Resource type of the destination tier override.
type ResourceType string

const (
	// ResourceTypeDatabase - Database resource type.
	ResourceTypeDatabase ResourceType = "Database"
	// ResourceTypePool - Elastic pool resource type.
	ResourceTypePool ResourceType = "Pool"
)

// PossibleResourceTypeValues returns the possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{
		ResourceTypeDatabase,
		ResourceTypePool,
	}
}

// ZoneRedundancy - Status of zone redundancy in a tier.
type ZoneRedundancy string

const (
	// ZoneRedundancyDisabled - Zone redundancy disabled.
	ZoneRedundancyDisabled ZoneRedundancy = "Disabled"
	// ZoneRedundancyEnabled - Zone redundancy enabled.
	ZoneRedundancyEnabled ZoneRedundancy = "Enabled"
)

// PossibleZoneRedundancyValues returns the possible values for the ZoneRedundancy const type.
func PossibleZoneRedundancyValues() []ZoneRedundancy {
	return []ZoneRedundancy{
		ZoneRedundancyDisabled,
		ZoneRedundancyEnabled,
	}
}
