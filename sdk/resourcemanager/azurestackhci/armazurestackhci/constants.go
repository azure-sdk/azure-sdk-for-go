//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

const (
	moduleName    = "armazurestackhci"
	moduleVersion = "v1.2.0-beta.1"
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

// ArcSettingAggregateState - Aggregate state of Arc agent across the nodes in this HCI cluster.
type ArcSettingAggregateState string

const (
	ArcSettingAggregateStateAccepted           ArcSettingAggregateState = "Accepted"
	ArcSettingAggregateStateCanceled           ArcSettingAggregateState = "Canceled"
	ArcSettingAggregateStateConnected          ArcSettingAggregateState = "Connected"
	ArcSettingAggregateStateCreating           ArcSettingAggregateState = "Creating"
	ArcSettingAggregateStateDeleted            ArcSettingAggregateState = "Deleted"
	ArcSettingAggregateStateDeleting           ArcSettingAggregateState = "Deleting"
	ArcSettingAggregateStateDisableInProgress  ArcSettingAggregateState = "DisableInProgress"
	ArcSettingAggregateStateDisconnected       ArcSettingAggregateState = "Disconnected"
	ArcSettingAggregateStateError              ArcSettingAggregateState = "Error"
	ArcSettingAggregateStateFailed             ArcSettingAggregateState = "Failed"
	ArcSettingAggregateStateInProgress         ArcSettingAggregateState = "InProgress"
	ArcSettingAggregateStateMoving             ArcSettingAggregateState = "Moving"
	ArcSettingAggregateStateNotSpecified       ArcSettingAggregateState = "NotSpecified"
	ArcSettingAggregateStatePartiallyConnected ArcSettingAggregateState = "PartiallyConnected"
	ArcSettingAggregateStatePartiallySucceeded ArcSettingAggregateState = "PartiallySucceeded"
	ArcSettingAggregateStateProvisioning       ArcSettingAggregateState = "Provisioning"
	ArcSettingAggregateStateSucceeded          ArcSettingAggregateState = "Succeeded"
	ArcSettingAggregateStateUpdating           ArcSettingAggregateState = "Updating"
)

// PossibleArcSettingAggregateStateValues returns the possible values for the ArcSettingAggregateState const type.
func PossibleArcSettingAggregateStateValues() []ArcSettingAggregateState {
	return []ArcSettingAggregateState{
		ArcSettingAggregateStateAccepted,
		ArcSettingAggregateStateCanceled,
		ArcSettingAggregateStateConnected,
		ArcSettingAggregateStateCreating,
		ArcSettingAggregateStateDeleted,
		ArcSettingAggregateStateDeleting,
		ArcSettingAggregateStateDisableInProgress,
		ArcSettingAggregateStateDisconnected,
		ArcSettingAggregateStateError,
		ArcSettingAggregateStateFailed,
		ArcSettingAggregateStateInProgress,
		ArcSettingAggregateStateMoving,
		ArcSettingAggregateStateNotSpecified,
		ArcSettingAggregateStatePartiallyConnected,
		ArcSettingAggregateStatePartiallySucceeded,
		ArcSettingAggregateStateProvisioning,
		ArcSettingAggregateStateSucceeded,
		ArcSettingAggregateStateUpdating,
	}
}

// AvailabilityType - Indicates the way the update content can be downloaded.
type AvailabilityType string

const (
	AvailabilityTypeLocal  AvailabilityType = "Local"
	AvailabilityTypeNotify AvailabilityType = "Notify"
	AvailabilityTypeOnline AvailabilityType = "Online"
)

// PossibleAvailabilityTypeValues returns the possible values for the AvailabilityType const type.
func PossibleAvailabilityTypeValues() []AvailabilityType {
	return []AvailabilityType{
		AvailabilityTypeLocal,
		AvailabilityTypeNotify,
		AvailabilityTypeOnline,
	}
}

// ClusterNodeType - The node type of all the nodes of the cluster.
type ClusterNodeType string

const (
	ClusterNodeTypeFirstParty ClusterNodeType = "FirstParty"
	ClusterNodeTypeThirdParty ClusterNodeType = "ThirdParty"
)

// PossibleClusterNodeTypeValues returns the possible values for the ClusterNodeType const type.
func PossibleClusterNodeTypeValues() []ClusterNodeType {
	return []ClusterNodeType{
		ClusterNodeTypeFirstParty,
		ClusterNodeTypeThirdParty,
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

// DiagnosticLevel - Desired level of diagnostic data emitted by the cluster.
type DiagnosticLevel string

const (
	DiagnosticLevelBasic    DiagnosticLevel = "Basic"
	DiagnosticLevelEnhanced DiagnosticLevel = "Enhanced"
	DiagnosticLevelOff      DiagnosticLevel = "Off"
)

// PossibleDiagnosticLevelValues returns the possible values for the DiagnosticLevel const type.
func PossibleDiagnosticLevelValues() []DiagnosticLevel {
	return []DiagnosticLevel{
		DiagnosticLevelBasic,
		DiagnosticLevelEnhanced,
		DiagnosticLevelOff,
	}
}

// ExtensionAggregateState - Aggregate state of Arc Extensions across the nodes in this HCI cluster.
type ExtensionAggregateState string

const (
	ExtensionAggregateStateAccepted                       ExtensionAggregateState = "Accepted"
	ExtensionAggregateStateCanceled                       ExtensionAggregateState = "Canceled"
	ExtensionAggregateStateConnected                      ExtensionAggregateState = "Connected"
	ExtensionAggregateStateCreating                       ExtensionAggregateState = "Creating"
	ExtensionAggregateStateDeleted                        ExtensionAggregateState = "Deleted"
	ExtensionAggregateStateDeleting                       ExtensionAggregateState = "Deleting"
	ExtensionAggregateStateDisconnected                   ExtensionAggregateState = "Disconnected"
	ExtensionAggregateStateError                          ExtensionAggregateState = "Error"
	ExtensionAggregateStateFailed                         ExtensionAggregateState = "Failed"
	ExtensionAggregateStateInProgress                     ExtensionAggregateState = "InProgress"
	ExtensionAggregateStateMoving                         ExtensionAggregateState = "Moving"
	ExtensionAggregateStateNotSpecified                   ExtensionAggregateState = "NotSpecified"
	ExtensionAggregateStatePartiallyConnected             ExtensionAggregateState = "PartiallyConnected"
	ExtensionAggregateStatePartiallySucceeded             ExtensionAggregateState = "PartiallySucceeded"
	ExtensionAggregateStateProvisioning                   ExtensionAggregateState = "Provisioning"
	ExtensionAggregateStateSucceeded                      ExtensionAggregateState = "Succeeded"
	ExtensionAggregateStateUpdating                       ExtensionAggregateState = "Updating"
	ExtensionAggregateStateUpgradeFailedRollbackSucceeded ExtensionAggregateState = "UpgradeFailedRollbackSucceeded"
)

// PossibleExtensionAggregateStateValues returns the possible values for the ExtensionAggregateState const type.
func PossibleExtensionAggregateStateValues() []ExtensionAggregateState {
	return []ExtensionAggregateState{
		ExtensionAggregateStateAccepted,
		ExtensionAggregateStateCanceled,
		ExtensionAggregateStateConnected,
		ExtensionAggregateStateCreating,
		ExtensionAggregateStateDeleted,
		ExtensionAggregateStateDeleting,
		ExtensionAggregateStateDisconnected,
		ExtensionAggregateStateError,
		ExtensionAggregateStateFailed,
		ExtensionAggregateStateInProgress,
		ExtensionAggregateStateMoving,
		ExtensionAggregateStateNotSpecified,
		ExtensionAggregateStatePartiallyConnected,
		ExtensionAggregateStatePartiallySucceeded,
		ExtensionAggregateStateProvisioning,
		ExtensionAggregateStateSucceeded,
		ExtensionAggregateStateUpdating,
		ExtensionAggregateStateUpgradeFailedRollbackSucceeded,
	}
}

// ExtensionManagedBy - Indicates if the extension is managed by azure or the user.
type ExtensionManagedBy string

const (
	ExtensionManagedByAzure ExtensionManagedBy = "Azure"
	ExtensionManagedByUser  ExtensionManagedBy = "User"
)

// PossibleExtensionManagedByValues returns the possible values for the ExtensionManagedBy const type.
func PossibleExtensionManagedByValues() []ExtensionManagedBy {
	return []ExtensionManagedBy{
		ExtensionManagedByAzure,
		ExtensionManagedByUser,
	}
}

type HealthState string

const (
	HealthStateError      HealthState = "Error"
	HealthStateFailure    HealthState = "Failure"
	HealthStateInProgress HealthState = "InProgress"
	HealthStateSuccess    HealthState = "Success"
	HealthStateUnknown    HealthState = "Unknown"
	HealthStateWarning    HealthState = "Warning"
)

// PossibleHealthStateValues returns the possible values for the HealthState const type.
func PossibleHealthStateValues() []HealthState {
	return []HealthState{
		HealthStateError,
		HealthStateFailure,
		HealthStateInProgress,
		HealthStateSuccess,
		HealthStateUnknown,
		HealthStateWarning,
	}
}

// ImdsAttestation - IMDS attestation status of the cluster.
type ImdsAttestation string

const (
	ImdsAttestationDisabled ImdsAttestation = "Disabled"
	ImdsAttestationEnabled  ImdsAttestation = "Enabled"
)

// PossibleImdsAttestationValues returns the possible values for the ImdsAttestation const type.
func PossibleImdsAttestationValues() []ImdsAttestation {
	return []ImdsAttestation{
		ImdsAttestationDisabled,
		ImdsAttestationEnabled,
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

// NodeArcState - State of Arc agent in this node.
type NodeArcState string

const (
	NodeArcStateAccepted           NodeArcState = "Accepted"
	NodeArcStateCanceled           NodeArcState = "Canceled"
	NodeArcStateConnected          NodeArcState = "Connected"
	NodeArcStateCreating           NodeArcState = "Creating"
	NodeArcStateDeleted            NodeArcState = "Deleted"
	NodeArcStateDeleting           NodeArcState = "Deleting"
	NodeArcStateDisableInProgress  NodeArcState = "DisableInProgress"
	NodeArcStateDisconnected       NodeArcState = "Disconnected"
	NodeArcStateError              NodeArcState = "Error"
	NodeArcStateFailed             NodeArcState = "Failed"
	NodeArcStateInProgress         NodeArcState = "InProgress"
	NodeArcStateMoving             NodeArcState = "Moving"
	NodeArcStateNotSpecified       NodeArcState = "NotSpecified"
	NodeArcStatePartiallyConnected NodeArcState = "PartiallyConnected"
	NodeArcStatePartiallySucceeded NodeArcState = "PartiallySucceeded"
	NodeArcStateProvisioning       NodeArcState = "Provisioning"
	NodeArcStateSucceeded          NodeArcState = "Succeeded"
	NodeArcStateUpdating           NodeArcState = "Updating"
)

// PossibleNodeArcStateValues returns the possible values for the NodeArcState const type.
func PossibleNodeArcStateValues() []NodeArcState {
	return []NodeArcState{
		NodeArcStateAccepted,
		NodeArcStateCanceled,
		NodeArcStateConnected,
		NodeArcStateCreating,
		NodeArcStateDeleted,
		NodeArcStateDeleting,
		NodeArcStateDisableInProgress,
		NodeArcStateDisconnected,
		NodeArcStateError,
		NodeArcStateFailed,
		NodeArcStateInProgress,
		NodeArcStateMoving,
		NodeArcStateNotSpecified,
		NodeArcStatePartiallyConnected,
		NodeArcStatePartiallySucceeded,
		NodeArcStateProvisioning,
		NodeArcStateSucceeded,
		NodeArcStateUpdating,
	}
}

// NodeExtensionState - State of Arc Extension in this node.
type NodeExtensionState string

const (
	NodeExtensionStateAccepted           NodeExtensionState = "Accepted"
	NodeExtensionStateCanceled           NodeExtensionState = "Canceled"
	NodeExtensionStateConnected          NodeExtensionState = "Connected"
	NodeExtensionStateCreating           NodeExtensionState = "Creating"
	NodeExtensionStateDeleted            NodeExtensionState = "Deleted"
	NodeExtensionStateDeleting           NodeExtensionState = "Deleting"
	NodeExtensionStateDisconnected       NodeExtensionState = "Disconnected"
	NodeExtensionStateError              NodeExtensionState = "Error"
	NodeExtensionStateFailed             NodeExtensionState = "Failed"
	NodeExtensionStateInProgress         NodeExtensionState = "InProgress"
	NodeExtensionStateMoving             NodeExtensionState = "Moving"
	NodeExtensionStateNotSpecified       NodeExtensionState = "NotSpecified"
	NodeExtensionStatePartiallyConnected NodeExtensionState = "PartiallyConnected"
	NodeExtensionStatePartiallySucceeded NodeExtensionState = "PartiallySucceeded"
	NodeExtensionStateProvisioning       NodeExtensionState = "Provisioning"
	NodeExtensionStateSucceeded          NodeExtensionState = "Succeeded"
	NodeExtensionStateUpdating           NodeExtensionState = "Updating"
)

// PossibleNodeExtensionStateValues returns the possible values for the NodeExtensionState const type.
func PossibleNodeExtensionStateValues() []NodeExtensionState {
	return []NodeExtensionState{
		NodeExtensionStateAccepted,
		NodeExtensionStateCanceled,
		NodeExtensionStateConnected,
		NodeExtensionStateCreating,
		NodeExtensionStateDeleted,
		NodeExtensionStateDeleting,
		NodeExtensionStateDisconnected,
		NodeExtensionStateError,
		NodeExtensionStateFailed,
		NodeExtensionStateInProgress,
		NodeExtensionStateMoving,
		NodeExtensionStateNotSpecified,
		NodeExtensionStatePartiallyConnected,
		NodeExtensionStatePartiallySucceeded,
		NodeExtensionStateProvisioning,
		NodeExtensionStateSucceeded,
		NodeExtensionStateUpdating,
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

// ProvisioningState - Provisioning state of the ArcSetting proxy resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted           ProvisioningState = "Accepted"
	ProvisioningStateCanceled           ProvisioningState = "Canceled"
	ProvisioningStateConnected          ProvisioningState = "Connected"
	ProvisioningStateCreating           ProvisioningState = "Creating"
	ProvisioningStateDeleted            ProvisioningState = "Deleted"
	ProvisioningStateDeleting           ProvisioningState = "Deleting"
	ProvisioningStateDisableInProgress  ProvisioningState = "DisableInProgress"
	ProvisioningStateDisconnected       ProvisioningState = "Disconnected"
	ProvisioningStateError              ProvisioningState = "Error"
	ProvisioningStateFailed             ProvisioningState = "Failed"
	ProvisioningStateInProgress         ProvisioningState = "InProgress"
	ProvisioningStateMoving             ProvisioningState = "Moving"
	ProvisioningStateNotSpecified       ProvisioningState = "NotSpecified"
	ProvisioningStatePartiallyConnected ProvisioningState = "PartiallyConnected"
	ProvisioningStatePartiallySucceeded ProvisioningState = "PartiallySucceeded"
	ProvisioningStateProvisioning       ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded          ProvisioningState = "Succeeded"
	ProvisioningStateUpdating           ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateConnected,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateDisableInProgress,
		ProvisioningStateDisconnected,
		ProvisioningStateError,
		ProvisioningStateFailed,
		ProvisioningStateInProgress,
		ProvisioningStateMoving,
		ProvisioningStateNotSpecified,
		ProvisioningStatePartiallyConnected,
		ProvisioningStatePartiallySucceeded,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

type RebootRequirement string

const (
	RebootRequirementFalse   RebootRequirement = "False"
	RebootRequirementTrue    RebootRequirement = "True"
	RebootRequirementUnknown RebootRequirement = "Unknown"
)

// PossibleRebootRequirementValues returns the possible values for the RebootRequirement const type.
func PossibleRebootRequirementValues() []RebootRequirement {
	return []RebootRequirement{
		RebootRequirementFalse,
		RebootRequirementTrue,
		RebootRequirementUnknown,
	}
}

// Severity - Severity of the result (Critical, Warning, Informational, Hidden). This answers how important the result is.
// Critical is the only update-blocking severity.
type Severity string

const (
	SeverityCritical      Severity = "Critical"
	SeverityHidden        Severity = "Hidden"
	SeverityInformational Severity = "Informational"
	SeverityWarning       Severity = "Warning"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeverityCritical,
		SeverityHidden,
		SeverityInformational,
		SeverityWarning,
	}
}

// SoftwareAssuranceIntent - Customer Intent for Software Assurance Benefit.
type SoftwareAssuranceIntent string

const (
	SoftwareAssuranceIntentDisable SoftwareAssuranceIntent = "Disable"
	SoftwareAssuranceIntentEnable  SoftwareAssuranceIntent = "Enable"
)

// PossibleSoftwareAssuranceIntentValues returns the possible values for the SoftwareAssuranceIntent const type.
func PossibleSoftwareAssuranceIntentValues() []SoftwareAssuranceIntent {
	return []SoftwareAssuranceIntent{
		SoftwareAssuranceIntentDisable,
		SoftwareAssuranceIntentEnable,
	}
}

// SoftwareAssuranceStatus - Status of the Software Assurance for the cluster.
type SoftwareAssuranceStatus string

const (
	SoftwareAssuranceStatusDisabled SoftwareAssuranceStatus = "Disabled"
	SoftwareAssuranceStatusEnabled  SoftwareAssuranceStatus = "Enabled"
)

// PossibleSoftwareAssuranceStatusValues returns the possible values for the SoftwareAssuranceStatus const type.
func PossibleSoftwareAssuranceStatusValues() []SoftwareAssuranceStatus {
	return []SoftwareAssuranceStatus{
		SoftwareAssuranceStatusDisabled,
		SoftwareAssuranceStatusEnabled,
	}
}

// State - State of the update as it relates to this stamp.
type State string

const (
	StateDownloadFailed                                State = "DownloadFailed"
	StateDownloading                                   State = "Downloading"
	StateHasPrerequisite                               State = "HasPrerequisite"
	StateHealthCheckFailed                             State = "HealthCheckFailed"
	StateHealthChecking                                State = "HealthChecking"
	StateInstallationFailed                            State = "InstallationFailed"
	StateInstalled                                     State = "Installed"
	StateInstalling                                    State = "Installing"
	StateInvalid                                       State = "Invalid"
	StateNotApplicableBecauseAnotherUpdateIsInProgress State = "NotApplicableBecauseAnotherUpdateIsInProgress"
	StateObsolete                                      State = "Obsolete"
	StatePreparationFailed                             State = "PreparationFailed"
	StatePreparing                                     State = "Preparing"
	StateReady                                         State = "Ready"
	StateReadyToInstall                                State = "ReadyToInstall"
	StateRecalled                                      State = "Recalled"
	StateScanFailed                                    State = "ScanFailed"
	StateScanInProgress                                State = "ScanInProgress"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDownloadFailed,
		StateDownloading,
		StateHasPrerequisite,
		StateHealthCheckFailed,
		StateHealthChecking,
		StateInstallationFailed,
		StateInstalled,
		StateInstalling,
		StateInvalid,
		StateNotApplicableBecauseAnotherUpdateIsInProgress,
		StateObsolete,
		StatePreparationFailed,
		StatePreparing,
		StateReady,
		StateReadyToInstall,
		StateRecalled,
		StateScanFailed,
		StateScanInProgress,
	}
}

// Status - Status of the cluster agent.
type Status string

const (
	StatusConnectedRecently    Status = "ConnectedRecently"
	StatusDisconnected         Status = "Disconnected"
	StatusError                Status = "Error"
	StatusFailed               Status = "Failed"
	StatusInProgress           Status = "InProgress"
	StatusNotConnectedRecently Status = "NotConnectedRecently"
	StatusNotSpecified         Status = "NotSpecified"
	StatusNotYetRegistered     Status = "NotYetRegistered"
	StatusSucceeded            Status = "Succeeded"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusConnectedRecently,
		StatusDisconnected,
		StatusError,
		StatusFailed,
		StatusInProgress,
		StatusNotConnectedRecently,
		StatusNotSpecified,
		StatusNotYetRegistered,
		StatusSucceeded,
	}
}

// StatusLevelTypes - The level code.
type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

// PossibleStatusLevelTypesValues returns the possible values for the StatusLevelTypes const type.
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return []StatusLevelTypes{
		StatusLevelTypesError,
		StatusLevelTypesInfo,
		StatusLevelTypesWarning,
	}
}

// UpdateRunPropertiesState - State of the update run.
type UpdateRunPropertiesState string

const (
	UpdateRunPropertiesStateFailed     UpdateRunPropertiesState = "Failed"
	UpdateRunPropertiesStateInProgress UpdateRunPropertiesState = "InProgress"
	UpdateRunPropertiesStateSucceeded  UpdateRunPropertiesState = "Succeeded"
	UpdateRunPropertiesStateUnknown    UpdateRunPropertiesState = "Unknown"
)

// PossibleUpdateRunPropertiesStateValues returns the possible values for the UpdateRunPropertiesState const type.
func PossibleUpdateRunPropertiesStateValues() []UpdateRunPropertiesState {
	return []UpdateRunPropertiesState{
		UpdateRunPropertiesStateFailed,
		UpdateRunPropertiesStateInProgress,
		UpdateRunPropertiesStateSucceeded,
		UpdateRunPropertiesStateUnknown,
	}
}

// UpdateSummariesPropertiesState - Overall update state of the stamp.
type UpdateSummariesPropertiesState string

const (
	UpdateSummariesPropertiesStateAppliedSuccessfully   UpdateSummariesPropertiesState = "AppliedSuccessfully"
	UpdateSummariesPropertiesStateNeedsAttention        UpdateSummariesPropertiesState = "NeedsAttention"
	UpdateSummariesPropertiesStatePreparationFailed     UpdateSummariesPropertiesState = "PreparationFailed"
	UpdateSummariesPropertiesStatePreparationInProgress UpdateSummariesPropertiesState = "PreparationInProgress"
	UpdateSummariesPropertiesStateUnknown               UpdateSummariesPropertiesState = "Unknown"
	UpdateSummariesPropertiesStateUpdateAvailable       UpdateSummariesPropertiesState = "UpdateAvailable"
	UpdateSummariesPropertiesStateUpdateFailed          UpdateSummariesPropertiesState = "UpdateFailed"
	UpdateSummariesPropertiesStateUpdateInProgress      UpdateSummariesPropertiesState = "UpdateInProgress"
)

// PossibleUpdateSummariesPropertiesStateValues returns the possible values for the UpdateSummariesPropertiesState const type.
func PossibleUpdateSummariesPropertiesStateValues() []UpdateSummariesPropertiesState {
	return []UpdateSummariesPropertiesState{
		UpdateSummariesPropertiesStateAppliedSuccessfully,
		UpdateSummariesPropertiesStateNeedsAttention,
		UpdateSummariesPropertiesStatePreparationFailed,
		UpdateSummariesPropertiesStatePreparationInProgress,
		UpdateSummariesPropertiesStateUnknown,
		UpdateSummariesPropertiesStateUpdateAvailable,
		UpdateSummariesPropertiesStateUpdateFailed,
		UpdateSummariesPropertiesStateUpdateInProgress,
	}
}

// WindowsServerSubscription - Desired state of Windows Server Subscription.
type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled WindowsServerSubscription = "Disabled"
	WindowsServerSubscriptionEnabled  WindowsServerSubscription = "Enabled"
)

// PossibleWindowsServerSubscriptionValues returns the possible values for the WindowsServerSubscription const type.
func PossibleWindowsServerSubscriptionValues() []WindowsServerSubscription {
	return []WindowsServerSubscription{
		WindowsServerSubscriptionDisabled,
		WindowsServerSubscriptionEnabled,
	}
}
