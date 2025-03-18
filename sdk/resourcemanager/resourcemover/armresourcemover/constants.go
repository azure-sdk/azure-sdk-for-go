// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
	moduleVersion = "v1.4.0-beta.1"
)

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

type DependencyLevel string

const (
	DependencyLevelDescendant DependencyLevel = "Descendant"
	DependencyLevelDirect     DependencyLevel = "Direct"
)

// PossibleDependencyLevelValues returns the possible values for the DependencyLevel const type.
func PossibleDependencyLevelValues() []DependencyLevel {
	return []DependencyLevel{
		DependencyLevelDescendant,
		DependencyLevelDirect,
	}
}

// DependencyType - Defines the dependency type.
type DependencyType string

const (
	DependencyTypeRequiredForMove    DependencyType = "RequiredForMove"
	DependencyTypeRequiredForPrepare DependencyType = "RequiredForPrepare"
)

// PossibleDependencyTypeValues returns the possible values for the DependencyType const type.
func PossibleDependencyTypeValues() []DependencyType {
	return []DependencyType{
		DependencyTypeRequiredForMove,
		DependencyTypeRequiredForPrepare,
	}
}

// JobName - Defines the job name.
type JobName string

const (
	JobNameInitialSync JobName = "InitialSync"
)

// PossibleJobNameValues returns the possible values for the JobName const type.
func PossibleJobNameValues() []JobName {
	return []JobName{
		JobNameInitialSync,
	}
}

// MoveResourceInputType - Defines the move resource input type.
type MoveResourceInputType string

const (
	MoveResourceInputTypeMoveResourceID       MoveResourceInputType = "MoveResourceId"
	MoveResourceInputTypeMoveResourceSourceID MoveResourceInputType = "MoveResourceSourceId"
)

// PossibleMoveResourceInputTypeValues returns the possible values for the MoveResourceInputType const type.
func PossibleMoveResourceInputTypeValues() []MoveResourceInputType {
	return []MoveResourceInputType{
		MoveResourceInputTypeMoveResourceID,
		MoveResourceInputTypeMoveResourceSourceID,
	}
}

// MoveState - Defines the MoveResource states.
type MoveState string

const (
	MoveStateAssignmentPending     MoveState = "AssignmentPending"
	MoveStateCommitFailed          MoveState = "CommitFailed"
	MoveStateCommitInProgress      MoveState = "CommitInProgress"
	MoveStateCommitPending         MoveState = "CommitPending"
	MoveStateCommitted             MoveState = "Committed"
	MoveStateDeleteSourcePending   MoveState = "DeleteSourcePending"
	MoveStateDiscardFailed         MoveState = "DiscardFailed"
	MoveStateDiscardInProgress     MoveState = "DiscardInProgress"
	MoveStateMoveFailed            MoveState = "MoveFailed"
	MoveStateMoveInProgress        MoveState = "MoveInProgress"
	MoveStateMovePending           MoveState = "MovePending"
	MoveStatePrepareFailed         MoveState = "PrepareFailed"
	MoveStatePrepareInProgress     MoveState = "PrepareInProgress"
	MoveStatePreparePending        MoveState = "PreparePending"
	MoveStateResourceMoveCompleted MoveState = "ResourceMoveCompleted"
)

// PossibleMoveStateValues returns the possible values for the MoveState const type.
func PossibleMoveStateValues() []MoveState {
	return []MoveState{
		MoveStateAssignmentPending,
		MoveStateCommitFailed,
		MoveStateCommitInProgress,
		MoveStateCommitPending,
		MoveStateCommitted,
		MoveStateDeleteSourcePending,
		MoveStateDiscardFailed,
		MoveStateDiscardInProgress,
		MoveStateMoveFailed,
		MoveStateMoveInProgress,
		MoveStateMovePending,
		MoveStatePrepareFailed,
		MoveStatePrepareInProgress,
		MoveStatePreparePending,
		MoveStateResourceMoveCompleted,
	}
}

// MoveType - Defines the MoveType.
type MoveType string

const (
	MoveTypeAvailabilitySetToVMSSFlex MoveType = "AvailabilitySetToVMSSFlex"
	MoveTypeRegionToRegion            MoveType = "RegionToRegion"
	MoveTypeRegionToZone              MoveType = "RegionToZone"
)

// PossibleMoveTypeValues returns the possible values for the MoveType const type.
func PossibleMoveTypeValues() []MoveType {
	return []MoveType{
		MoveTypeAvailabilitySetToVMSSFlex,
		MoveTypeRegionToRegion,
		MoveTypeRegionToZone,
	}
}

// ProvisioningState - Defines the provisioning states.
type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResolutionType - Defines the resolution type.
type ResolutionType string

const (
	ResolutionTypeAutomatic ResolutionType = "Automatic"
	ResolutionTypeManual    ResolutionType = "Manual"
)

// PossibleResolutionTypeValues returns the possible values for the ResolutionType const type.
func PossibleResolutionTypeValues() []ResolutionType {
	return []ResolutionType{
		ResolutionTypeAutomatic,
		ResolutionTypeManual,
	}
}

// ResourceIdentityType - The type of identity used for the resource mover service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// TargetAvailabilityZone - Gets or sets the target availability zone.
type TargetAvailabilityZone string

const (
	TargetAvailabilityZoneNA    TargetAvailabilityZone = "NA"
	TargetAvailabilityZoneOne   TargetAvailabilityZone = "1"
	TargetAvailabilityZoneThree TargetAvailabilityZone = "3"
	TargetAvailabilityZoneTwo   TargetAvailabilityZone = "2"
)

// PossibleTargetAvailabilityZoneValues returns the possible values for the TargetAvailabilityZone const type.
func PossibleTargetAvailabilityZoneValues() []TargetAvailabilityZone {
	return []TargetAvailabilityZone{
		TargetAvailabilityZoneNA,
		TargetAvailabilityZoneOne,
		TargetAvailabilityZoneThree,
		TargetAvailabilityZoneTwo,
	}
}

// ZoneRedundant - Defines the zone redundant resource setting.
type ZoneRedundant string

const (
	ZoneRedundantDisable ZoneRedundant = "Disable"
	ZoneRedundantEnable  ZoneRedundant = "Enable"
)

// PossibleZoneRedundantValues returns the possible values for the ZoneRedundant const type.
func PossibleZoneRedundantValues() []ZoneRedundant {
	return []ZoneRedundant{
		ZoneRedundantDisable,
		ZoneRedundantEnable,
	}
}
