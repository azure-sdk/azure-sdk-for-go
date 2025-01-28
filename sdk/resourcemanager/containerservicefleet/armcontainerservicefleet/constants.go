//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet"
	moduleVersion = "v1.3.0-beta.2"
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

// AutoUpgradeNodeImageSelectionType - The node image upgrade type.
type AutoUpgradeNodeImageSelectionType string

const (
	// AutoUpgradeNodeImageSelectionTypeConsistent - The image versions to upgrade nodes to are selected as described below: for
	// each node pool in managed clusters affected by the update run, the system selects the latest image version such that it
	// is available across all other node pools (in all other clusters) of the same image type. As a result, all node pools of
	// the same image type will be upgraded to the same image version. For example, if the latest image version for image type
	// 'AKSUbuntu-1804gen2containerd' is 'AKSUbuntu-1804gen2containerd-2021.10.12' for a node pool in cluster A in region X, and
	// is 'AKSUbuntu-1804gen2containerd-2021.10.17' for a node pool in cluster B in region Y, the system will upgrade both node
	// pools to image version 'AKSUbuntu-1804gen2containerd-2021.10.12'.
	AutoUpgradeNodeImageSelectionTypeConsistent AutoUpgradeNodeImageSelectionType = "Consistent"
	// AutoUpgradeNodeImageSelectionTypeLatest - Use the latest image version when upgrading nodes. Clusters may use different
	// image versions (e.g., 'AKSUbuntu-1804gen2containerd-2021.10.12' and 'AKSUbuntu-1804gen2containerd-2021.10.19') because,
	// for example, the latest available version is different in different regions.
	AutoUpgradeNodeImageSelectionTypeLatest AutoUpgradeNodeImageSelectionType = "Latest"
)

// PossibleAutoUpgradeNodeImageSelectionTypeValues returns the possible values for the AutoUpgradeNodeImageSelectionType const type.
func PossibleAutoUpgradeNodeImageSelectionTypeValues() []AutoUpgradeNodeImageSelectionType {
	return []AutoUpgradeNodeImageSelectionType{
		AutoUpgradeNodeImageSelectionTypeConsistent,
		AutoUpgradeNodeImageSelectionTypeLatest,
	}
}

// AutoUpgradeProfileProvisioningState - The provisioning state of the AutoUpgradeProfile resource.
type AutoUpgradeProfileProvisioningState string

const (
	// AutoUpgradeProfileProvisioningStateCanceled - Resource creation was canceled.
	AutoUpgradeProfileProvisioningStateCanceled AutoUpgradeProfileProvisioningState = "Canceled"
	// AutoUpgradeProfileProvisioningStateFailed - Resource creation failed.
	AutoUpgradeProfileProvisioningStateFailed AutoUpgradeProfileProvisioningState = "Failed"
	// AutoUpgradeProfileProvisioningStateSucceeded - Resource has been created.
	AutoUpgradeProfileProvisioningStateSucceeded AutoUpgradeProfileProvisioningState = "Succeeded"
)

// PossibleAutoUpgradeProfileProvisioningStateValues returns the possible values for the AutoUpgradeProfileProvisioningState const type.
func PossibleAutoUpgradeProfileProvisioningStateValues() []AutoUpgradeProfileProvisioningState {
	return []AutoUpgradeProfileProvisioningState{
		AutoUpgradeProfileProvisioningStateCanceled,
		AutoUpgradeProfileProvisioningStateFailed,
		AutoUpgradeProfileProvisioningStateSucceeded,
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

// FleetMemberProvisioningState - The provisioning state of the last accepted operation.
type FleetMemberProvisioningState string

const (
	// FleetMemberProvisioningStateCanceled - Resource creation was canceled.
	FleetMemberProvisioningStateCanceled FleetMemberProvisioningState = "Canceled"
	// FleetMemberProvisioningStateFailed - Resource creation failed.
	FleetMemberProvisioningStateFailed FleetMemberProvisioningState = "Failed"
	// FleetMemberProvisioningStateJoining - The provisioning state of a member joining a fleet.
	FleetMemberProvisioningStateJoining FleetMemberProvisioningState = "Joining"
	// FleetMemberProvisioningStateLeaving - The provisioning state of a member leaving a fleet.
	FleetMemberProvisioningStateLeaving FleetMemberProvisioningState = "Leaving"
	// FleetMemberProvisioningStateSucceeded - Resource has been created.
	FleetMemberProvisioningStateSucceeded FleetMemberProvisioningState = "Succeeded"
	// FleetMemberProvisioningStateUpdating - The provisioning state of a member being updated.
	FleetMemberProvisioningStateUpdating FleetMemberProvisioningState = "Updating"
)

// PossibleFleetMemberProvisioningStateValues returns the possible values for the FleetMemberProvisioningState const type.
func PossibleFleetMemberProvisioningStateValues() []FleetMemberProvisioningState {
	return []FleetMemberProvisioningState{
		FleetMemberProvisioningStateCanceled,
		FleetMemberProvisioningStateFailed,
		FleetMemberProvisioningStateJoining,
		FleetMemberProvisioningStateLeaving,
		FleetMemberProvisioningStateSucceeded,
		FleetMemberProvisioningStateUpdating,
	}
}

// FleetProvisioningState - The provisioning state of the last accepted operation.
type FleetProvisioningState string

const (
	// FleetProvisioningStateCanceled - Resource creation was canceled.
	FleetProvisioningStateCanceled FleetProvisioningState = "Canceled"
	// FleetProvisioningStateCreating - The provisioning state of a fleet being created.
	FleetProvisioningStateCreating FleetProvisioningState = "Creating"
	// FleetProvisioningStateDeleting - The provisioning state of a fleet being deleted.
	FleetProvisioningStateDeleting FleetProvisioningState = "Deleting"
	// FleetProvisioningStateFailed - Resource creation failed.
	FleetProvisioningStateFailed FleetProvisioningState = "Failed"
	// FleetProvisioningStateSucceeded - Resource has been created.
	FleetProvisioningStateSucceeded FleetProvisioningState = "Succeeded"
	// FleetProvisioningStateUpdating - The provisioning state of a fleet being updated.
	FleetProvisioningStateUpdating FleetProvisioningState = "Updating"
)

// PossibleFleetProvisioningStateValues returns the possible values for the FleetProvisioningState const type.
func PossibleFleetProvisioningStateValues() []FleetProvisioningState {
	return []FleetProvisioningState{
		FleetProvisioningStateCanceled,
		FleetProvisioningStateCreating,
		FleetProvisioningStateDeleting,
		FleetProvisioningStateFailed,
		FleetProvisioningStateSucceeded,
		FleetProvisioningStateUpdating,
	}
}

// FleetUpdateStrategyProvisioningState - The provisioning state of the UpdateStrategy resource.
type FleetUpdateStrategyProvisioningState string

const (
	// FleetUpdateStrategyProvisioningStateCanceled - Resource creation was canceled.
	FleetUpdateStrategyProvisioningStateCanceled FleetUpdateStrategyProvisioningState = "Canceled"
	// FleetUpdateStrategyProvisioningStateFailed - Resource creation failed.
	FleetUpdateStrategyProvisioningStateFailed FleetUpdateStrategyProvisioningState = "Failed"
	// FleetUpdateStrategyProvisioningStateSucceeded - Resource has been created.
	FleetUpdateStrategyProvisioningStateSucceeded FleetUpdateStrategyProvisioningState = "Succeeded"
)

// PossibleFleetUpdateStrategyProvisioningStateValues returns the possible values for the FleetUpdateStrategyProvisioningState const type.
func PossibleFleetUpdateStrategyProvisioningStateValues() []FleetUpdateStrategyProvisioningState {
	return []FleetUpdateStrategyProvisioningState{
		FleetUpdateStrategyProvisioningStateCanceled,
		FleetUpdateStrategyProvisioningStateFailed,
		FleetUpdateStrategyProvisioningStateSucceeded,
	}
}

// ManagedClusterUpgradeType - The type of upgrade to perform when targeting ManagedClusters.
type ManagedClusterUpgradeType string

const (
	// ManagedClusterUpgradeTypeControlPlaneOnly - ControlPlaneOnly upgrades only targets the KubernetesVersion of the ManagedClusters
	// and will not be applied to the AgentPool. Requires the ManagedClusterUpgradeSpec.KubernetesVersion property to be set.
	ManagedClusterUpgradeTypeControlPlaneOnly ManagedClusterUpgradeType = "ControlPlaneOnly"
	// ManagedClusterUpgradeTypeFull - Full upgrades the control plane and all agent pools of the target ManagedClusters. Requires
	// the ManagedClusterUpgradeSpec.KubernetesVersion property to be set.
	ManagedClusterUpgradeTypeFull ManagedClusterUpgradeType = "Full"
	// ManagedClusterUpgradeTypeNodeImageOnly - NodeImageOnly upgrades only the node images of the target ManagedClusters. Requires
	// the ManagedClusterUpgradeSpec.KubernetesVersion property to NOT be set.
	ManagedClusterUpgradeTypeNodeImageOnly ManagedClusterUpgradeType = "NodeImageOnly"
)

// PossibleManagedClusterUpgradeTypeValues returns the possible values for the ManagedClusterUpgradeType const type.
func PossibleManagedClusterUpgradeTypeValues() []ManagedClusterUpgradeType {
	return []ManagedClusterUpgradeType{
		ManagedClusterUpgradeTypeControlPlaneOnly,
		ManagedClusterUpgradeTypeFull,
		ManagedClusterUpgradeTypeNodeImageOnly,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// NodeImageSelectionType - The node image upgrade type.
type NodeImageSelectionType string

const (
	// NodeImageSelectionTypeConsistent - The image versions to upgrade nodes to are selected as described below: for each node
	// pool in managed clusters affected by the update run, the system selects the latest image version such that it is available
	// across all other node pools (in all other clusters) of the same image type. As a result, all node pools of the same image
	// type will be upgraded to the same image version. For example, if the latest image version for image type 'AKSUbuntu-1804gen2containerd'
	// is 'AKSUbuntu-1804gen2containerd-2021.10.12' for a node pool in cluster A in region X, and is 'AKSUbuntu-1804gen2containerd-2021.10.17'
	// for a node pool in cluster B in region Y, the system will upgrade both node pools to image version 'AKSUbuntu-1804gen2containerd-2021.10.12'.
	NodeImageSelectionTypeConsistent NodeImageSelectionType = "Consistent"
	// NodeImageSelectionTypeCustom - Upgrade the nodes to the custom image versions. When set, update run will use node image
	// versions provided in customNodeImageVersions to upgrade the nodes. If set, customNodeImageVersions must not be empty.
	NodeImageSelectionTypeCustom NodeImageSelectionType = "Custom"
	// NodeImageSelectionTypeLatest - Use the latest image version when upgrading nodes. Clusters may use different image versions
	// (e.g., 'AKSUbuntu-1804gen2containerd-2021.10.12' and 'AKSUbuntu-1804gen2containerd-2021.10.19') because, for example, the
	// latest available version is different in different regions.
	NodeImageSelectionTypeLatest NodeImageSelectionType = "Latest"
)

// PossibleNodeImageSelectionTypeValues returns the possible values for the NodeImageSelectionType const type.
func PossibleNodeImageSelectionTypeValues() []NodeImageSelectionType {
	return []NodeImageSelectionType{
		NodeImageSelectionTypeConsistent,
		NodeImageSelectionTypeCustom,
		NodeImageSelectionTypeLatest,
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

// TargetType - The target type of a skip request.
type TargetType string

const (
	// TargetTypeAfterStageWait - Skip the update of the after stage wait of a certain stage.
	TargetTypeAfterStageWait TargetType = "AfterStageWait"
	// TargetTypeGroup - Skip the update of a group.
	TargetTypeGroup TargetType = "Group"
	// TargetTypeMember - Skip the update of a member.
	TargetTypeMember TargetType = "Member"
	// TargetTypeStage - Skip the update of an entire stage including the after stage wait.
	TargetTypeStage TargetType = "Stage"
)

// PossibleTargetTypeValues returns the possible values for the TargetType const type.
func PossibleTargetTypeValues() []TargetType {
	return []TargetType{
		TargetTypeAfterStageWait,
		TargetTypeGroup,
		TargetTypeMember,
		TargetTypeStage,
	}
}

// UpdateRunProvisioningState - The provisioning state of the UpdateRun resource.
type UpdateRunProvisioningState string

const (
	// UpdateRunProvisioningStateCanceled - Resource creation was canceled.
	UpdateRunProvisioningStateCanceled UpdateRunProvisioningState = "Canceled"
	// UpdateRunProvisioningStateFailed - Resource creation failed.
	UpdateRunProvisioningStateFailed UpdateRunProvisioningState = "Failed"
	// UpdateRunProvisioningStateSucceeded - Resource has been created.
	UpdateRunProvisioningStateSucceeded UpdateRunProvisioningState = "Succeeded"
)

// PossibleUpdateRunProvisioningStateValues returns the possible values for the UpdateRunProvisioningState const type.
func PossibleUpdateRunProvisioningStateValues() []UpdateRunProvisioningState {
	return []UpdateRunProvisioningState{
		UpdateRunProvisioningStateCanceled,
		UpdateRunProvisioningStateFailed,
		UpdateRunProvisioningStateSucceeded,
	}
}

// UpdateState - The state of the UpdateRun, UpdateStage, UpdateGroup, or MemberUpdate.
type UpdateState string

const (
	// UpdateStateCompleted - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that has completed.
	UpdateStateCompleted UpdateState = "Completed"
	// UpdateStateFailed - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that has failed.
	UpdateStateFailed UpdateState = "Failed"
	// UpdateStateNotStarted - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that has not been started.
	UpdateStateNotStarted UpdateState = "NotStarted"
	// UpdateStateRunning - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that is running.
	UpdateStateRunning UpdateState = "Running"
	// UpdateStateSkipped - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that has been skipped.
	UpdateStateSkipped UpdateState = "Skipped"
	// UpdateStateStopped - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that has stopped.
	UpdateStateStopped UpdateState = "Stopped"
	// UpdateStateStopping - The state of an UpdateRun/UpdateStage/UpdateGroup/MemberUpdate that is being stopped.
	UpdateStateStopping UpdateState = "Stopping"
)

// PossibleUpdateStateValues returns the possible values for the UpdateState const type.
func PossibleUpdateStateValues() []UpdateState {
	return []UpdateState{
		UpdateStateCompleted,
		UpdateStateFailed,
		UpdateStateNotStarted,
		UpdateStateRunning,
		UpdateStateSkipped,
		UpdateStateStopped,
		UpdateStateStopping,
	}
}

// UpgradeChannel - Configuration of how auto upgrade will be run.
type UpgradeChannel string

const (
	// UpgradeChannelNodeImage - Upgrade node image version of the clusters.
	UpgradeChannelNodeImage UpgradeChannel = "NodeImage"
	// UpgradeChannelRapid - Upgrades the clusters kubernetes version to the latest supported patch release on the latest supported
	// minor version.
	UpgradeChannelRapid UpgradeChannel = "Rapid"
	// UpgradeChannelStable - Upgrades the clusters kubernetes version to the latest supported patch release on minor version
	// N-1, where N is the latest supported minor version.
	// For example, if a cluster runs version 1.17.7 and versions 1.17.9, 1.18.4, 1.18.6, and 1.19.1 are available, the cluster
	// upgrades to 1.18.6.
	UpgradeChannelStable UpgradeChannel = "Stable"
)

// PossibleUpgradeChannelValues returns the possible values for the UpgradeChannel const type.
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return []UpgradeChannel{
		UpgradeChannelNodeImage,
		UpgradeChannelRapid,
		UpgradeChannelStable,
	}
}
