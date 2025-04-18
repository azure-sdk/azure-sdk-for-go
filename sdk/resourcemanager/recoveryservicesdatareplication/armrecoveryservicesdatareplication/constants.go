// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	moduleVersion = "v0.2.0"
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

// HealthStatus - Gets or sets the fabric health.
type HealthStatus string

const (
	// HealthStatusCritical - Critical Status.
	HealthStatusCritical HealthStatus = "Critical"
	// HealthStatusNormal - Healthy Status.
	HealthStatusNormal HealthStatus = "Normal"
	// HealthStatusWarning - Warning Status.
	HealthStatusWarning HealthStatus = "Warning"
)

// PossibleHealthStatusValues returns the possible values for the HealthStatus const type.
func PossibleHealthStatusValues() []HealthStatus {
	return []HealthStatus{
		HealthStatusCritical,
		HealthStatusNormal,
		HealthStatusWarning,
	}
}

// JobObjectType - Gets or sets the object type.
type JobObjectType string

const (
	// JobObjectTypeAvsDiskPool - AVS disk pool.
	JobObjectTypeAvsDiskPool JobObjectType = "AvsDiskPool"
	// JobObjectTypeFabric - Fabric level job.
	JobObjectTypeFabric JobObjectType = "Fabric"
	// JobObjectTypeFabricAgent - Fabric agent level workflow.
	JobObjectTypeFabricAgent JobObjectType = "FabricAgent"
	// JobObjectTypePolicy - Policy level job.
	JobObjectTypePolicy JobObjectType = "Policy"
	// JobObjectTypeProtectedItem - Protected item level job.
	JobObjectTypeProtectedItem JobObjectType = "ProtectedItem"
	// JobObjectTypeRecoveryPlan - Recovery plan level job.
	JobObjectTypeRecoveryPlan JobObjectType = "RecoveryPlan"
	// JobObjectTypeReplicationExtension - Replication extension level job.
	JobObjectTypeReplicationExtension JobObjectType = "ReplicationExtension"
	// JobObjectTypeVault - Vault level job.
	JobObjectTypeVault JobObjectType = "Vault"
)

// PossibleJobObjectTypeValues returns the possible values for the JobObjectType const type.
func PossibleJobObjectTypeValues() []JobObjectType {
	return []JobObjectType{
		JobObjectTypeAvsDiskPool,
		JobObjectTypeFabric,
		JobObjectTypeFabricAgent,
		JobObjectTypePolicy,
		JobObjectTypeProtectedItem,
		JobObjectTypeRecoveryPlan,
		JobObjectTypeReplicationExtension,
		JobObjectTypeVault,
	}
}

// JobState - Gets or sets the job state.
type JobState string

const (
	// JobStateCancelled - Job has been cancelled.
	JobStateCancelled JobState = "Cancelled"
	// JobStateCancelling - Job cancellation is in progress.
	JobStateCancelling JobState = "Cancelling"
	// JobStateCompletedWithErrors - Job has completed with errors.
	JobStateCompletedWithErrors JobState = "CompletedWithErrors"
	// JobStateCompletedWithInformation - Job has completed with information.
	JobStateCompletedWithInformation JobState = "CompletedWithInformation"
	// JobStateCompletedWithWarnings - Job has completed with warnings.
	JobStateCompletedWithWarnings JobState = "CompletedWithWarnings"
	// JobStateFailed - Job failed.
	JobStateFailed JobState = "Failed"
	// JobStatePending - Job has not been started.
	JobStatePending JobState = "Pending"
	// JobStateStarted - Job is in progress.
	JobStateStarted JobState = "Started"
	// JobStateSucceeded - Job has completed successfully.
	JobStateSucceeded JobState = "Succeeded"
)

// PossibleJobStateValues returns the possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{
		JobStateCancelled,
		JobStateCancelling,
		JobStateCompletedWithErrors,
		JobStateCompletedWithInformation,
		JobStateCompletedWithWarnings,
		JobStateFailed,
		JobStatePending,
		JobStateStarted,
		JobStateSucceeded,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
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

// PrivateEndpointConnectionStatus - Gets or sets the status.
type PrivateEndpointConnectionStatus string

const (
	// PrivateEndpointConnectionStatusApproved - Approved Status.
	PrivateEndpointConnectionStatusApproved PrivateEndpointConnectionStatus = "Approved"
	// PrivateEndpointConnectionStatusDisconnected - Disconnected Status.
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = "Disconnected"
	// PrivateEndpointConnectionStatusPending - Pending Status.
	PrivateEndpointConnectionStatusPending PrivateEndpointConnectionStatus = "Pending"
	// PrivateEndpointConnectionStatusRejected - Rejected Status.
	PrivateEndpointConnectionStatusRejected PrivateEndpointConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointConnectionStatusValues returns the possible values for the PrivateEndpointConnectionStatus const type.
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return []PrivateEndpointConnectionStatus{
		PrivateEndpointConnectionStatusApproved,
		PrivateEndpointConnectionStatusDisconnected,
		PrivateEndpointConnectionStatusPending,
		PrivateEndpointConnectionStatusRejected,
	}
}

// ProtectedItemActiveLocation - Gets or sets the location of the protected item.
type ProtectedItemActiveLocation string

const (
	// ProtectedItemActiveLocationPrimary - Protected item is active on Primary.
	ProtectedItemActiveLocationPrimary ProtectedItemActiveLocation = "Primary"
	// ProtectedItemActiveLocationRecovery - Protected item is active on Recovery.
	ProtectedItemActiveLocationRecovery ProtectedItemActiveLocation = "Recovery"
)

// PossibleProtectedItemActiveLocationValues returns the possible values for the ProtectedItemActiveLocation const type.
func PossibleProtectedItemActiveLocationValues() []ProtectedItemActiveLocation {
	return []ProtectedItemActiveLocation{
		ProtectedItemActiveLocationPrimary,
		ProtectedItemActiveLocationRecovery,
	}
}

// ProtectionState - Gets or sets the protection state.
type ProtectionState string

const (
	// ProtectionStateCancelFailoverFailedOnPrimary - Cancel failover failed on the primary side.
	ProtectionStateCancelFailoverFailedOnPrimary ProtectionState = "CancelFailoverFailedOnPrimary"
	// ProtectionStateCancelFailoverFailedOnRecovery - Cancel failover failed on the recovery side.
	ProtectionStateCancelFailoverFailedOnRecovery ProtectionState = "CancelFailoverFailedOnRecovery"
	// ProtectionStateCancelFailoverInProgressOnPrimary - Cancel failover is in progress on the primary side.
	ProtectionStateCancelFailoverInProgressOnPrimary ProtectionState = "CancelFailoverInProgressOnPrimary"
	// ProtectionStateCancelFailoverInProgressOnRecovery - Cancel failover is in progress on the recovery side.
	ProtectionStateCancelFailoverInProgressOnRecovery ProtectionState = "CancelFailoverInProgressOnRecovery"
	// ProtectionStateCancelFailoverStatesBegin - Begin marker for cancel failover states.
	ProtectionStateCancelFailoverStatesBegin ProtectionState = "CancelFailoverStatesBegin"
	// ProtectionStateCancelFailoverStatesEnd - End marker for cancel failover states.
	ProtectionStateCancelFailoverStatesEnd ProtectionState = "CancelFailoverStatesEnd"
	// ProtectionStateChangeRecoveryPointCompleted - Change recovery point has been completed successfully.
	ProtectionStateChangeRecoveryPointCompleted ProtectionState = "ChangeRecoveryPointCompleted"
	// ProtectionStateChangeRecoveryPointFailed - Change recovery point has failed.
	ProtectionStateChangeRecoveryPointFailed ProtectionState = "ChangeRecoveryPointFailed"
	// ProtectionStateChangeRecoveryPointInitiated - Change recovery point has been initiated..
	ProtectionStateChangeRecoveryPointInitiated ProtectionState = "ChangeRecoveryPointInitiated"
	// ProtectionStateChangeRecoveryPointStatesBegin - Begin marker for change recovery point states.
	ProtectionStateChangeRecoveryPointStatesBegin ProtectionState = "ChangeRecoveryPointStatesBegin"
	// ProtectionStateChangeRecoveryPointStatesEnd - End marker for change recovery point states.
	ProtectionStateChangeRecoveryPointStatesEnd ProtectionState = "ChangeRecoveryPointStatesEnd"
	// ProtectionStateCommitFailoverCompleted - Commit failover has been completed successfully.
	ProtectionStateCommitFailoverCompleted ProtectionState = "CommitFailoverCompleted"
	// ProtectionStateCommitFailoverFailedOnPrimary - Commit failover failed on the primary side.
	ProtectionStateCommitFailoverFailedOnPrimary ProtectionState = "CommitFailoverFailedOnPrimary"
	// ProtectionStateCommitFailoverFailedOnRecovery - Commit failover failed on the recovery side.
	ProtectionStateCommitFailoverFailedOnRecovery ProtectionState = "CommitFailoverFailedOnRecovery"
	// ProtectionStateCommitFailoverInProgressOnPrimary - Commit failover is in progress on the primary side.
	ProtectionStateCommitFailoverInProgressOnPrimary ProtectionState = "CommitFailoverInProgressOnPrimary"
	// ProtectionStateCommitFailoverInProgressOnRecovery - Commit failover is in progress on the recovery side.
	ProtectionStateCommitFailoverInProgressOnRecovery ProtectionState = "CommitFailoverInProgressOnRecovery"
	// ProtectionStateCommitFailoverStatesBegin - Begin marker for commit failover states.
	ProtectionStateCommitFailoverStatesBegin ProtectionState = "CommitFailoverStatesBegin"
	// ProtectionStateCommitFailoverStatesEnd - End marker for commit failover states.
	ProtectionStateCommitFailoverStatesEnd ProtectionState = "CommitFailoverStatesEnd"
	// ProtectionStateDisablingFailed - Disable protection failed.
	ProtectionStateDisablingFailed ProtectionState = "DisablingFailed"
	// ProtectionStateDisablingProtection - Disabling protection is in progress.
	ProtectionStateDisablingProtection ProtectionState = "DisablingProtection"
	// ProtectionStateEnablingFailed - Enable protection failed.
	ProtectionStateEnablingFailed ProtectionState = "EnablingFailed"
	// ProtectionStateEnablingProtection - Enable protection is in progress.
	ProtectionStateEnablingProtection ProtectionState = "EnablingProtection"
	// ProtectionStateInitialReplicationCompletedOnPrimary - Initial replication has completed on the primary side.
	ProtectionStateInitialReplicationCompletedOnPrimary ProtectionState = "InitialReplicationCompletedOnPrimary"
	// ProtectionStateInitialReplicationCompletedOnRecovery - Initial replication has completed on the recovery side.
	ProtectionStateInitialReplicationCompletedOnRecovery ProtectionState = "InitialReplicationCompletedOnRecovery"
	// ProtectionStateInitialReplicationFailed - Initial replication failed and would need to be started again.
	ProtectionStateInitialReplicationFailed ProtectionState = "InitialReplicationFailed"
	// ProtectionStateInitialReplicationInProgress - Initial replication is in progress.
	ProtectionStateInitialReplicationInProgress ProtectionState = "InitialReplicationInProgress"
	// ProtectionStateInitialReplicationStatesBegin - Begin marker for initial replication states.
	ProtectionStateInitialReplicationStatesBegin ProtectionState = "InitialReplicationStatesBegin"
	// ProtectionStateInitialReplicationStatesEnd - End marker for initial replication states.
	ProtectionStateInitialReplicationStatesEnd ProtectionState = "InitialReplicationStatesEnd"
	// ProtectionStateMarkedForDeletion - Disabling protection succeeded. This is a transient state before the protected item
	// is deleted.
	ProtectionStateMarkedForDeletion ProtectionState = "MarkedForDeletion"
	// ProtectionStatePlannedFailoverCompleted - Planned failover has been completed successfully.
	ProtectionStatePlannedFailoverCompleted ProtectionState = "PlannedFailoverCompleted"
	// ProtectionStatePlannedFailoverCompleting - Planned failover preparing protected entities is in progress.
	ProtectionStatePlannedFailoverCompleting ProtectionState = "PlannedFailoverCompleting"
	// ProtectionStatePlannedFailoverCompletionFailed - Planned failover preparing protected entities failed.
	ProtectionStatePlannedFailoverCompletionFailed ProtectionState = "PlannedFailoverCompletionFailed"
	// ProtectionStatePlannedFailoverFailed - Planned failover initiation failed.
	ProtectionStatePlannedFailoverFailed ProtectionState = "PlannedFailoverFailed"
	// ProtectionStatePlannedFailoverInitiated - Planned failover has been initiated.
	ProtectionStatePlannedFailoverInitiated ProtectionState = "PlannedFailoverInitiated"
	// ProtectionStatePlannedFailoverTransitionStatesBegin - Begin marker for planned failover transition states.
	ProtectionStatePlannedFailoverTransitionStatesBegin ProtectionState = "PlannedFailoverTransitionStatesBegin"
	// ProtectionStatePlannedFailoverTransitionStatesEnd - End marker for planned failover transition states.
	ProtectionStatePlannedFailoverTransitionStatesEnd ProtectionState = "PlannedFailoverTransitionStatesEnd"
	// ProtectionStateProtected - Protected item is protected and replication is on-going. Any issues with replication will be
	// surfaced separately via the health property and will not affect the state.
	ProtectionStateProtected ProtectionState = "Protected"
	// ProtectionStateProtectedStatesBegin - Begin marker for protected steady-state states.
	ProtectionStateProtectedStatesBegin ProtectionState = "ProtectedStatesBegin"
	// ProtectionStateProtectedStatesEnd - End marker for protected steady-state states.
	ProtectionStateProtectedStatesEnd ProtectionState = "ProtectedStatesEnd"
	// ProtectionStateReprotectFailed - Reprotect has failed.
	ProtectionStateReprotectFailed ProtectionState = "ReprotectFailed"
	// ProtectionStateReprotectInitiated - Reprotect has been initiated.
	ProtectionStateReprotectInitiated ProtectionState = "ReprotectInitiated"
	// ProtectionStateReprotectStatesBegin - Begin marker for reprotect states.
	ProtectionStateReprotectStatesBegin ProtectionState = "ReprotectStatesBegin"
	// ProtectionStateReprotectStatesEnd - End marker for reprotect states.
	ProtectionStateReprotectStatesEnd ProtectionState = "ReprotectStatesEnd"
	// ProtectionStateUnplannedFailoverCompleted - Unplanned failover preparing protected entities is in progress.
	ProtectionStateUnplannedFailoverCompleted ProtectionState = "UnplannedFailoverCompleted"
	// ProtectionStateUnplannedFailoverCompleting - Unplanned failover preparing protected entities is in progress.
	ProtectionStateUnplannedFailoverCompleting ProtectionState = "UnplannedFailoverCompleting"
	// ProtectionStateUnplannedFailoverCompletionFailed - Unplanned failover preparing protected entities failed.
	ProtectionStateUnplannedFailoverCompletionFailed ProtectionState = "UnplannedFailoverCompletionFailed"
	// ProtectionStateUnplannedFailoverFailed - Unplanned failover initiation failed.
	ProtectionStateUnplannedFailoverFailed ProtectionState = "UnplannedFailoverFailed"
	// ProtectionStateUnplannedFailoverInitiated - Unplanned failover has been initiated.
	ProtectionStateUnplannedFailoverInitiated ProtectionState = "UnplannedFailoverInitiated"
	// ProtectionStateUnplannedFailoverTransitionStatesBegin - Begin marker for unplanned failover transition states.
	ProtectionStateUnplannedFailoverTransitionStatesBegin ProtectionState = "UnplannedFailoverTransitionStatesBegin"
	// ProtectionStateUnplannedFailoverTransitionStatesEnd - End marker for unplanned failover transition states.
	ProtectionStateUnplannedFailoverTransitionStatesEnd ProtectionState = "UnplannedFailoverTransitionStatesEnd"
	// ProtectionStateUnprotectedStatesBegin - Begin marker for unprotected states.
	ProtectionStateUnprotectedStatesBegin ProtectionState = "UnprotectedStatesBegin"
	// ProtectionStateUnprotectedStatesEnd - End marker for unprotected states.
	ProtectionStateUnprotectedStatesEnd ProtectionState = "UnprotectedStatesEnd"
)

// PossibleProtectionStateValues returns the possible values for the ProtectionState const type.
func PossibleProtectionStateValues() []ProtectionState {
	return []ProtectionState{
		ProtectionStateCancelFailoverFailedOnPrimary,
		ProtectionStateCancelFailoverFailedOnRecovery,
		ProtectionStateCancelFailoverInProgressOnPrimary,
		ProtectionStateCancelFailoverInProgressOnRecovery,
		ProtectionStateCancelFailoverStatesBegin,
		ProtectionStateCancelFailoverStatesEnd,
		ProtectionStateChangeRecoveryPointCompleted,
		ProtectionStateChangeRecoveryPointFailed,
		ProtectionStateChangeRecoveryPointInitiated,
		ProtectionStateChangeRecoveryPointStatesBegin,
		ProtectionStateChangeRecoveryPointStatesEnd,
		ProtectionStateCommitFailoverCompleted,
		ProtectionStateCommitFailoverFailedOnPrimary,
		ProtectionStateCommitFailoverFailedOnRecovery,
		ProtectionStateCommitFailoverInProgressOnPrimary,
		ProtectionStateCommitFailoverInProgressOnRecovery,
		ProtectionStateCommitFailoverStatesBegin,
		ProtectionStateCommitFailoverStatesEnd,
		ProtectionStateDisablingFailed,
		ProtectionStateDisablingProtection,
		ProtectionStateEnablingFailed,
		ProtectionStateEnablingProtection,
		ProtectionStateInitialReplicationCompletedOnPrimary,
		ProtectionStateInitialReplicationCompletedOnRecovery,
		ProtectionStateInitialReplicationFailed,
		ProtectionStateInitialReplicationInProgress,
		ProtectionStateInitialReplicationStatesBegin,
		ProtectionStateInitialReplicationStatesEnd,
		ProtectionStateMarkedForDeletion,
		ProtectionStatePlannedFailoverCompleted,
		ProtectionStatePlannedFailoverCompleting,
		ProtectionStatePlannedFailoverCompletionFailed,
		ProtectionStatePlannedFailoverFailed,
		ProtectionStatePlannedFailoverInitiated,
		ProtectionStatePlannedFailoverTransitionStatesBegin,
		ProtectionStatePlannedFailoverTransitionStatesEnd,
		ProtectionStateProtected,
		ProtectionStateProtectedStatesBegin,
		ProtectionStateProtectedStatesEnd,
		ProtectionStateReprotectFailed,
		ProtectionStateReprotectInitiated,
		ProtectionStateReprotectStatesBegin,
		ProtectionStateReprotectStatesEnd,
		ProtectionStateUnplannedFailoverCompleted,
		ProtectionStateUnplannedFailoverCompleting,
		ProtectionStateUnplannedFailoverCompletionFailed,
		ProtectionStateUnplannedFailoverFailed,
		ProtectionStateUnplannedFailoverInitiated,
		ProtectionStateUnplannedFailoverTransitionStatesBegin,
		ProtectionStateUnplannedFailoverTransitionStatesEnd,
		ProtectionStateUnprotectedStatesBegin,
		ProtectionStateUnprotectedStatesEnd,
	}
}

// ProvisioningState - Gets or sets the provisioning state of the email configuration.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation has been canceled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Resource is being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted - Resource has been deleted.
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting - Resource is being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource creation/update succeeded.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Resource is being updated.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RecoveryPointType - Gets or sets the recovery point type.
type RecoveryPointType string

const (
	// RecoveryPointTypeApplicationConsistent - Application consistent recovery point.
	RecoveryPointTypeApplicationConsistent RecoveryPointType = "ApplicationConsistent"
	// RecoveryPointTypeCrashConsistent - Crash consistent recovery point.
	RecoveryPointTypeCrashConsistent RecoveryPointType = "CrashConsistent"
)

// PossibleRecoveryPointTypeValues returns the possible values for the RecoveryPointType const type.
func PossibleRecoveryPointTypeValues() []RecoveryPointType {
	return []RecoveryPointType{
		RecoveryPointTypeApplicationConsistent,
		RecoveryPointTypeCrashConsistent,
	}
}

// ReplicationVaultType - Gets or sets the type of vault.
type ReplicationVaultType string

const (
	// ReplicationVaultTypeDisasterRecovery - Disaster recovery vault.
	ReplicationVaultTypeDisasterRecovery ReplicationVaultType = "DisasterRecovery"
	// ReplicationVaultTypeMigrate - Migrate vault.
	ReplicationVaultTypeMigrate ReplicationVaultType = "Migrate"
)

// PossibleReplicationVaultTypeValues returns the possible values for the ReplicationVaultType const type.
func PossibleReplicationVaultTypeValues() []ReplicationVaultType {
	return []ReplicationVaultType{
		ReplicationVaultTypeDisasterRecovery,
		ReplicationVaultTypeMigrate,
	}
}

// ResynchronizationState - Gets or sets the resynchronization state.
type ResynchronizationState string

const (
	// ResynchronizationStateNone - Resynchronization is not active.
	ResynchronizationStateNone ResynchronizationState = "None"
	// ResynchronizationStateResynchronizationCompleted - Resynchronization has been completed successfully.
	ResynchronizationStateResynchronizationCompleted ResynchronizationState = "ResynchronizationCompleted"
	// ResynchronizationStateResynchronizationFailed - Resynchronization has failed and would need to be started again.
	ResynchronizationStateResynchronizationFailed ResynchronizationState = "ResynchronizationFailed"
	// ResynchronizationStateResynchronizationInitiated - Resynchronization has been initiated.
	ResynchronizationStateResynchronizationInitiated ResynchronizationState = "ResynchronizationInitiated"
)

// PossibleResynchronizationStateValues returns the possible values for the ResynchronizationState const type.
func PossibleResynchronizationStateValues() []ResynchronizationState {
	return []ResynchronizationState{
		ResynchronizationStateNone,
		ResynchronizationStateResynchronizationCompleted,
		ResynchronizationStateResynchronizationFailed,
		ResynchronizationStateResynchronizationInitiated,
	}
}

// TaskState - Gets or sets the task state.
type TaskState string

const (
	// TaskStateCancelled - Task has been cancelled.
	TaskStateCancelled TaskState = "Cancelled"
	// TaskStateFailed - Task failed.
	TaskStateFailed TaskState = "Failed"
	// TaskStatePending - Task has not been started.
	TaskStatePending TaskState = "Pending"
	// TaskStateSkipped - Task has been skipped.
	TaskStateSkipped TaskState = "Skipped"
	// TaskStateStarted - Task is in progress.
	TaskStateStarted TaskState = "Started"
	// TaskStateSucceeded - Task has completed successfully.
	TaskStateSucceeded TaskState = "Succeeded"
)

// PossibleTaskStateValues returns the possible values for the TaskState const type.
func PossibleTaskStateValues() []TaskState {
	return []TaskState{
		TaskStateCancelled,
		TaskStateFailed,
		TaskStatePending,
		TaskStateSkipped,
		TaskStateStarted,
		TaskStateSucceeded,
	}
}

// TestFailoverState - Gets or sets the test failover state.
type TestFailoverState string

const (
	// TestFailoverStateMarkedForDeletion - Test failover cleanup has completed/failed. This is a transient state before the state
	// is moved back to None.
	TestFailoverStateMarkedForDeletion TestFailoverState = "MarkedForDeletion"
	// TestFailoverStateNone - Test failover is not active.
	TestFailoverStateNone TestFailoverState = "None"
	// TestFailoverStateTestFailoverCleanupCompleting - Cleaning up test protected entities is in progress.
	TestFailoverStateTestFailoverCleanupCompleting TestFailoverState = "TestFailoverCleanupCompleting"
	// TestFailoverStateTestFailoverCleanupInitiated - Test failover cleanup has been initiated.
	TestFailoverStateTestFailoverCleanupInitiated TestFailoverState = "TestFailoverCleanupInitiated"
	// TestFailoverStateTestFailoverCompleted - Test failover has been completed successfully.
	TestFailoverStateTestFailoverCompleted TestFailoverState = "TestFailoverCompleted"
	// TestFailoverStateTestFailoverCompleting - Preparing test protected entities is in progress.
	TestFailoverStateTestFailoverCompleting TestFailoverState = "TestFailoverCompleting"
	// TestFailoverStateTestFailoverCompletionFailed - Preparing test protected entities failed.
	TestFailoverStateTestFailoverCompletionFailed TestFailoverState = "TestFailoverCompletionFailed"
	// TestFailoverStateTestFailoverFailed - Test failover initiation failed..
	TestFailoverStateTestFailoverFailed TestFailoverState = "TestFailoverFailed"
	// TestFailoverStateTestFailoverInitiated - Test failover has been initiated.
	TestFailoverStateTestFailoverInitiated TestFailoverState = "TestFailoverInitiated"
)

// PossibleTestFailoverStateValues returns the possible values for the TestFailoverState const type.
func PossibleTestFailoverStateValues() []TestFailoverState {
	return []TestFailoverState{
		TestFailoverStateMarkedForDeletion,
		TestFailoverStateNone,
		TestFailoverStateTestFailoverCleanupCompleting,
		TestFailoverStateTestFailoverCleanupInitiated,
		TestFailoverStateTestFailoverCompleted,
		TestFailoverStateTestFailoverCompleting,
		TestFailoverStateTestFailoverCompletionFailed,
		TestFailoverStateTestFailoverFailed,
		TestFailoverStateTestFailoverInitiated,
	}
}

// VMNicSelection - Gets or sets the selection type of the NIC.
type VMNicSelection string

const (
	// VMNicSelectionNotSelected - Not Selected.
	VMNicSelectionNotSelected VMNicSelection = "NotSelected"
	// VMNicSelectionSelectedByDefault - Default selection by ASR.
	VMNicSelectionSelectedByDefault VMNicSelection = "SelectedByDefault"
	// VMNicSelectionSelectedByUser - Selected by user.
	VMNicSelectionSelectedByUser VMNicSelection = "SelectedByUser"
	// VMNicSelectionSelectedByUserOverride - NIC configuration overridden by user. Differs from SelectedByUser in the sense that
	// the legacy SelectedByUser is used both for explicit modification by user and implicit approval of user if the settings
	// are used for TFO/FO. SelectedByUserOverride implies user overriding at least one of the configurations.
	VMNicSelectionSelectedByUserOverride VMNicSelection = "SelectedByUserOverride"
)

// PossibleVMNicSelectionValues returns the possible values for the VMNicSelection const type.
func PossibleVMNicSelectionValues() []VMNicSelection {
	return []VMNicSelection{
		VMNicSelectionNotSelected,
		VMNicSelectionSelectedByDefault,
		VMNicSelectionSelectedByUser,
		VMNicSelectionSelectedByUserOverride,
	}
}

// VMwareToAzureMigrateResyncState - Gets or sets the resync state.
type VMwareToAzureMigrateResyncState string

const (
	// VMwareToAzureMigrateResyncStateNone - None state.
	VMwareToAzureMigrateResyncStateNone VMwareToAzureMigrateResyncState = "None"
	// VMwareToAzureMigrateResyncStatePreparedForResynchronization - Prepared for resynchronization state.
	VMwareToAzureMigrateResyncStatePreparedForResynchronization VMwareToAzureMigrateResyncState = "PreparedForResynchronization"
	// VMwareToAzureMigrateResyncStateStartedResynchronization - Started resynchronization state.
	VMwareToAzureMigrateResyncStateStartedResynchronization VMwareToAzureMigrateResyncState = "StartedResynchronization"
)

// PossibleVMwareToAzureMigrateResyncStateValues returns the possible values for the VMwareToAzureMigrateResyncState const type.
func PossibleVMwareToAzureMigrateResyncStateValues() []VMwareToAzureMigrateResyncState {
	return []VMwareToAzureMigrateResyncState{
		VMwareToAzureMigrateResyncStateNone,
		VMwareToAzureMigrateResyncStatePreparedForResynchronization,
		VMwareToAzureMigrateResyncStateStartedResynchronization,
	}
}

// VaultIdentityType - Gets or sets the identityType which can be either SystemAssigned or None.
type VaultIdentityType string

const (
	// VaultIdentityTypeNone - No identity.
	VaultIdentityTypeNone VaultIdentityType = "None"
	// VaultIdentityTypeSystemAssigned - System assigned identity.
	VaultIdentityTypeSystemAssigned VaultIdentityType = "SystemAssigned"
	// VaultIdentityTypeUserAssigned - User assigned identity.
	VaultIdentityTypeUserAssigned VaultIdentityType = "UserAssigned"
)

// PossibleVaultIdentityTypeValues returns the possible values for the VaultIdentityType const type.
func PossibleVaultIdentityTypeValues() []VaultIdentityType {
	return []VaultIdentityType{
		VaultIdentityTypeNone,
		VaultIdentityTypeSystemAssigned,
		VaultIdentityTypeUserAssigned,
	}
}
