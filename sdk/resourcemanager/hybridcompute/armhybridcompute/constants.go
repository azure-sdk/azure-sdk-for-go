//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	moduleVersion = "v2.0.0-beta.3"
)

// AgentConfigurationMode - Name of configuration mode to use. Modes are pre-defined configurations of security controls,
// extension allowlists and guest configuration, maintained by Microsoft.
type AgentConfigurationMode string

const (
	AgentConfigurationModeFull    AgentConfigurationMode = "full"
	AgentConfigurationModeMonitor AgentConfigurationMode = "monitor"
)

// PossibleAgentConfigurationModeValues returns the possible values for the AgentConfigurationMode const type.
func PossibleAgentConfigurationModeValues() []AgentConfigurationMode {
	return []AgentConfigurationMode{
		AgentConfigurationModeFull,
		AgentConfigurationModeMonitor,
	}
}

// ArcKindEnum - Indicates which kind of Arc machine placement on-premises, such as HCI, SCVMM or VMware etc.
type ArcKindEnum string

const (
	ArcKindEnumAVS    ArcKindEnum = "AVS"
	ArcKindEnumAWS    ArcKindEnum = "AWS"
	ArcKindEnumEPS    ArcKindEnum = "EPS"
	ArcKindEnumGCP    ArcKindEnum = "GCP"
	ArcKindEnumHCI    ArcKindEnum = "HCI"
	ArcKindEnumSCVMM  ArcKindEnum = "SCVMM"
	ArcKindEnumVMware ArcKindEnum = "VMware"
)

// PossibleArcKindEnumValues returns the possible values for the ArcKindEnum const type.
func PossibleArcKindEnumValues() []ArcKindEnum {
	return []ArcKindEnum{
		ArcKindEnumAVS,
		ArcKindEnumAWS,
		ArcKindEnumEPS,
		ArcKindEnumGCP,
		ArcKindEnumHCI,
		ArcKindEnumSCVMM,
		ArcKindEnumVMware,
	}
}

// AssessmentModeTypes - Specifies the assessment mode.
type AssessmentModeTypes string

const (
	AssessmentModeTypesAutomaticByPlatform AssessmentModeTypes = "AutomaticByPlatform"
	AssessmentModeTypesImageDefault        AssessmentModeTypes = "ImageDefault"
)

// PossibleAssessmentModeTypesValues returns the possible values for the AssessmentModeTypes const type.
func PossibleAssessmentModeTypesValues() []AssessmentModeTypes {
	return []AssessmentModeTypes{
		AssessmentModeTypesAutomaticByPlatform,
		AssessmentModeTypesImageDefault,
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

// EsuEligibility - The ESU eligibility.
type EsuEligibility string

const (
	EsuEligibilityEligible   EsuEligibility = "Eligible"
	EsuEligibilityIneligible EsuEligibility = "Ineligible"
	EsuEligibilityUnknown    EsuEligibility = "Unknown"
)

// PossibleEsuEligibilityValues returns the possible values for the EsuEligibility const type.
func PossibleEsuEligibilityValues() []EsuEligibility {
	return []EsuEligibility{
		EsuEligibilityEligible,
		EsuEligibilityIneligible,
		EsuEligibilityUnknown,
	}
}

// EsuKeyState - The ESU key state.
type EsuKeyState string

const (
	EsuKeyStateActive   EsuKeyState = "Active"
	EsuKeyStateInactive EsuKeyState = "Inactive"
)

// PossibleEsuKeyStateValues returns the possible values for the EsuKeyState const type.
func PossibleEsuKeyStateValues() []EsuKeyState {
	return []EsuKeyState{
		EsuKeyStateActive,
		EsuKeyStateInactive,
	}
}

// EsuServerType - The server types for Esu.
type EsuServerType string

const (
	EsuServerTypeDatacenter EsuServerType = "Datacenter"
	EsuServerTypeStandard   EsuServerType = "Standard"
)

// PossibleEsuServerTypeValues returns the possible values for the EsuServerType const type.
func PossibleEsuServerTypeValues() []EsuServerType {
	return []EsuServerType{
		EsuServerTypeDatacenter,
		EsuServerTypeStandard,
	}
}

// ExecutionState - Script execution status.
type ExecutionState string

const (
	ExecutionStateCanceled  ExecutionState = "Canceled"
	ExecutionStateFailed    ExecutionState = "Failed"
	ExecutionStatePending   ExecutionState = "Pending"
	ExecutionStateRunning   ExecutionState = "Running"
	ExecutionStateSucceeded ExecutionState = "Succeeded"
	ExecutionStateTimedOut  ExecutionState = "TimedOut"
	ExecutionStateUnknown   ExecutionState = "Unknown"
)

// PossibleExecutionStateValues returns the possible values for the ExecutionState const type.
func PossibleExecutionStateValues() []ExecutionState {
	return []ExecutionState{
		ExecutionStateCanceled,
		ExecutionStateFailed,
		ExecutionStatePending,
		ExecutionStateRunning,
		ExecutionStateSucceeded,
		ExecutionStateTimedOut,
		ExecutionStateUnknown,
	}
}

// ExtensionsStatusLevelTypes - The level code.
type ExtensionsStatusLevelTypes string

const (
	ExtensionsStatusLevelTypesError   ExtensionsStatusLevelTypes = "Error"
	ExtensionsStatusLevelTypesInfo    ExtensionsStatusLevelTypes = "Info"
	ExtensionsStatusLevelTypesWarning ExtensionsStatusLevelTypes = "Warning"
)

// PossibleExtensionsStatusLevelTypesValues returns the possible values for the ExtensionsStatusLevelTypes const type.
func PossibleExtensionsStatusLevelTypesValues() []ExtensionsStatusLevelTypes {
	return []ExtensionsStatusLevelTypes{
		ExtensionsStatusLevelTypesError,
		ExtensionsStatusLevelTypesInfo,
		ExtensionsStatusLevelTypesWarning,
	}
}

// GatewayType - The type of the Gateway resource.
type GatewayType string

const (
	GatewayTypePublic GatewayType = "Public"
)

// PossibleGatewayTypeValues returns the possible values for the GatewayType const type.
func PossibleGatewayTypeValues() []GatewayType {
	return []GatewayType{
		GatewayTypePublic,
	}
}

// LastAttemptStatusEnum - Specifies the status of Agent Upgrade.
type LastAttemptStatusEnum string

const (
	LastAttemptStatusEnumFailed  LastAttemptStatusEnum = "Failed"
	LastAttemptStatusEnumSuccess LastAttemptStatusEnum = "Success"
)

// PossibleLastAttemptStatusEnumValues returns the possible values for the LastAttemptStatusEnum const type.
func PossibleLastAttemptStatusEnumValues() []LastAttemptStatusEnum {
	return []LastAttemptStatusEnum{
		LastAttemptStatusEnumFailed,
		LastAttemptStatusEnumSuccess,
	}
}

// LicenseAssignmentState - Describes the license assignment state (Assigned or NotAssigned).
type LicenseAssignmentState string

const (
	LicenseAssignmentStateAssigned    LicenseAssignmentState = "Assigned"
	LicenseAssignmentStateNotAssigned LicenseAssignmentState = "NotAssigned"
)

// PossibleLicenseAssignmentStateValues returns the possible values for the LicenseAssignmentState const type.
func PossibleLicenseAssignmentStateValues() []LicenseAssignmentState {
	return []LicenseAssignmentState{
		LicenseAssignmentStateAssigned,
		LicenseAssignmentStateNotAssigned,
	}
}

// LicenseCoreType - Describes the license core type (pCore or vCore).
type LicenseCoreType string

const (
	LicenseCoreTypePCore LicenseCoreType = "pCore"
	LicenseCoreTypeVCore LicenseCoreType = "vCore"
)

// PossibleLicenseCoreTypeValues returns the possible values for the LicenseCoreType const type.
func PossibleLicenseCoreTypeValues() []LicenseCoreType {
	return []LicenseCoreType{
		LicenseCoreTypePCore,
		LicenseCoreTypeVCore,
	}
}

// LicenseEdition - Describes the edition of the license. The values are either Standard or Datacenter.
type LicenseEdition string

const (
	LicenseEditionDatacenter LicenseEdition = "Datacenter"
	LicenseEditionStandard   LicenseEdition = "Standard"
)

// PossibleLicenseEditionValues returns the possible values for the LicenseEdition const type.
func PossibleLicenseEditionValues() []LicenseEdition {
	return []LicenseEdition{
		LicenseEditionDatacenter,
		LicenseEditionStandard,
	}
}

// LicenseProfileProductType - The product type of the license.
type LicenseProfileProductType string

const (
	LicenseProfileProductTypeWindowsIoTEnterprise LicenseProfileProductType = "WindowsIoTEnterprise"
	LicenseProfileProductTypeWindowsServer        LicenseProfileProductType = "WindowsServer"
)

// PossibleLicenseProfileProductTypeValues returns the possible values for the LicenseProfileProductType const type.
func PossibleLicenseProfileProductTypeValues() []LicenseProfileProductType {
	return []LicenseProfileProductType{
		LicenseProfileProductTypeWindowsIoTEnterprise,
		LicenseProfileProductTypeWindowsServer,
	}
}

// LicenseProfileSubscriptionStatus - Subscription status of the OS or Product feature.
type LicenseProfileSubscriptionStatus string

const (
	LicenseProfileSubscriptionStatusDisabled LicenseProfileSubscriptionStatus = "Disabled"
	LicenseProfileSubscriptionStatusEnabled  LicenseProfileSubscriptionStatus = "Enabled"
	LicenseProfileSubscriptionStatusEnabling LicenseProfileSubscriptionStatus = "Enabling"
	LicenseProfileSubscriptionStatusUnknown  LicenseProfileSubscriptionStatus = "Unknown"
)

// PossibleLicenseProfileSubscriptionStatusValues returns the possible values for the LicenseProfileSubscriptionStatus const type.
func PossibleLicenseProfileSubscriptionStatusValues() []LicenseProfileSubscriptionStatus {
	return []LicenseProfileSubscriptionStatus{
		LicenseProfileSubscriptionStatusDisabled,
		LicenseProfileSubscriptionStatusEnabled,
		LicenseProfileSubscriptionStatusEnabling,
		LicenseProfileSubscriptionStatusUnknown,
	}
}

// LicenseState - Describes the state of the license.
type LicenseState string

const (
	LicenseStateActivated   LicenseState = "Activated"
	LicenseStateDeactivated LicenseState = "Deactivated"
)

// PossibleLicenseStateValues returns the possible values for the LicenseState const type.
func PossibleLicenseStateValues() []LicenseState {
	return []LicenseState{
		LicenseStateActivated,
		LicenseStateDeactivated,
	}
}

// LicenseStatus - The license status.
type LicenseStatus string

const (
	LicenseStatusExtendedGrace   LicenseStatus = "ExtendedGrace"
	LicenseStatusLicensed        LicenseStatus = "Licensed"
	LicenseStatusNonGenuineGrace LicenseStatus = "NonGenuineGrace"
	LicenseStatusNotification    LicenseStatus = "Notification"
	LicenseStatusOOBGrace        LicenseStatus = "OOBGrace"
	LicenseStatusOOTGrace        LicenseStatus = "OOTGrace"
	LicenseStatusUnlicensed      LicenseStatus = "Unlicensed"
)

// PossibleLicenseStatusValues returns the possible values for the LicenseStatus const type.
func PossibleLicenseStatusValues() []LicenseStatus {
	return []LicenseStatus{
		LicenseStatusExtendedGrace,
		LicenseStatusLicensed,
		LicenseStatusNonGenuineGrace,
		LicenseStatusNotification,
		LicenseStatusOOBGrace,
		LicenseStatusOOTGrace,
		LicenseStatusUnlicensed,
	}
}

// LicenseTarget - Describes the license target server.
type LicenseTarget string

const (
	LicenseTargetWindowsServer2012   LicenseTarget = "Windows Server 2012"
	LicenseTargetWindowsServer2012R2 LicenseTarget = "Windows Server 2012 R2"
)

// PossibleLicenseTargetValues returns the possible values for the LicenseTarget const type.
func PossibleLicenseTargetValues() []LicenseTarget {
	return []LicenseTarget{
		LicenseTargetWindowsServer2012,
		LicenseTargetWindowsServer2012R2,
	}
}

// LicenseType - The type of the license resource.
type LicenseType string

const (
	LicenseTypeESU LicenseType = "ESU"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeESU,
	}
}

// OsType - The operating system type of the machine.
type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeLinux,
		OsTypeWindows,
	}
}

// PatchModeTypes - Specifies the patch mode.
type PatchModeTypes string

const (
	PatchModeTypesAutomaticByOS       PatchModeTypes = "AutomaticByOS"
	PatchModeTypesAutomaticByPlatform PatchModeTypes = "AutomaticByPlatform"
	PatchModeTypesImageDefault        PatchModeTypes = "ImageDefault"
	PatchModeTypesManual              PatchModeTypes = "Manual"
)

// PossiblePatchModeTypesValues returns the possible values for the PatchModeTypes const type.
func PossiblePatchModeTypesValues() []PatchModeTypes {
	return []PatchModeTypes{
		PatchModeTypesAutomaticByOS,
		PatchModeTypesAutomaticByPlatform,
		PatchModeTypesImageDefault,
		PatchModeTypesManual,
	}
}

// PatchOperationStartedBy - Indicates if operation was triggered by user or by platform.
type PatchOperationStartedBy string

const (
	PatchOperationStartedByPlatform PatchOperationStartedBy = "Platform"
	PatchOperationStartedByUser     PatchOperationStartedBy = "User"
)

// PossiblePatchOperationStartedByValues returns the possible values for the PatchOperationStartedBy const type.
func PossiblePatchOperationStartedByValues() []PatchOperationStartedBy {
	return []PatchOperationStartedBy{
		PatchOperationStartedByPlatform,
		PatchOperationStartedByUser,
	}
}

// PatchOperationStatus - The overall success or failure status of the operation. It remains "InProgress" until the operation
// completes. At that point it will become "Unknown", "Failed", "Succeeded", or
// "CompletedWithWarnings."
type PatchOperationStatus string

const (
	PatchOperationStatusCompletedWithWarnings PatchOperationStatus = "CompletedWithWarnings"
	PatchOperationStatusFailed                PatchOperationStatus = "Failed"
	PatchOperationStatusInProgress            PatchOperationStatus = "InProgress"
	PatchOperationStatusSucceeded             PatchOperationStatus = "Succeeded"
	PatchOperationStatusUnknown               PatchOperationStatus = "Unknown"
)

// PossiblePatchOperationStatusValues returns the possible values for the PatchOperationStatus const type.
func PossiblePatchOperationStatusValues() []PatchOperationStatus {
	return []PatchOperationStatus{
		PatchOperationStatusCompletedWithWarnings,
		PatchOperationStatusFailed,
		PatchOperationStatusInProgress,
		PatchOperationStatusSucceeded,
		PatchOperationStatusUnknown,
	}
}

// PatchServiceUsed - Specifies the patch service used for the operation.
type PatchServiceUsed string

const (
	PatchServiceUsedAPT     PatchServiceUsed = "APT"
	PatchServiceUsedUnknown PatchServiceUsed = "Unknown"
	PatchServiceUsedWU      PatchServiceUsed = "WU"
	PatchServiceUsedWUWSUS  PatchServiceUsed = "WU_WSUS"
	PatchServiceUsedYUM     PatchServiceUsed = "YUM"
	PatchServiceUsedZypper  PatchServiceUsed = "Zypper"
)

// PossiblePatchServiceUsedValues returns the possible values for the PatchServiceUsed const type.
func PossiblePatchServiceUsedValues() []PatchServiceUsed {
	return []PatchServiceUsed{
		PatchServiceUsedAPT,
		PatchServiceUsedUnknown,
		PatchServiceUsedWU,
		PatchServiceUsedWUWSUS,
		PatchServiceUsedYUM,
		PatchServiceUsedZypper,
	}
}

// ProgramYear - Describes the program year the volume license is for.
type ProgramYear string

const (
	ProgramYearYear1 ProgramYear = "Year 1"
	ProgramYearYear2 ProgramYear = "Year 2"
	ProgramYearYear3 ProgramYear = "Year 3"
)

// PossibleProgramYearValues returns the possible values for the ProgramYear const type.
func PossibleProgramYearValues() []ProgramYear {
	return []ProgramYear{
		ProgramYearYear1,
		ProgramYearYear2,
		ProgramYearYear3,
	}
}

// ProvisioningState - The provisioning state, which only appears in the response.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccessType - The network access policy to determine if Azure Arc agents can use public Azure Arc service endpoints.
// Defaults to disabled (access to Azure Arc services only via private link).
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled - Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet)
	// endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled - Allows Azure Arc agents to communicate with Azure Arc services over both public (internet)
	// and private endpoints.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
	// PublicNetworkAccessTypeSecuredByPerimeter - Azure Arc agent communication with Azure Arc services over public (internet)
	// is enforced by Network Security Perimeter (NSP)
	PublicNetworkAccessTypeSecuredByPerimeter PublicNetworkAccessType = "SecuredByPerimeter"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeDisabled,
		PublicNetworkAccessTypeEnabled,
		PublicNetworkAccessTypeSecuredByPerimeter,
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

// StatusTypes - The status of the hybrid machine agent.
type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

// PossibleStatusTypesValues returns the possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{
		StatusTypesConnected,
		StatusTypesDisconnected,
		StatusTypesError,
	}
}

type VMGuestPatchClassificationLinux string

const (
	VMGuestPatchClassificationLinuxCritical VMGuestPatchClassificationLinux = "Critical"
	VMGuestPatchClassificationLinuxOther    VMGuestPatchClassificationLinux = "Other"
	VMGuestPatchClassificationLinuxSecurity VMGuestPatchClassificationLinux = "Security"
)

// PossibleVMGuestPatchClassificationLinuxValues returns the possible values for the VMGuestPatchClassificationLinux const type.
func PossibleVMGuestPatchClassificationLinuxValues() []VMGuestPatchClassificationLinux {
	return []VMGuestPatchClassificationLinux{
		VMGuestPatchClassificationLinuxCritical,
		VMGuestPatchClassificationLinuxOther,
		VMGuestPatchClassificationLinuxSecurity,
	}
}

type VMGuestPatchClassificationWindows string

const (
	VMGuestPatchClassificationWindowsCritical     VMGuestPatchClassificationWindows = "Critical"
	VMGuestPatchClassificationWindowsDefinition   VMGuestPatchClassificationWindows = "Definition"
	VMGuestPatchClassificationWindowsFeaturePack  VMGuestPatchClassificationWindows = "FeaturePack"
	VMGuestPatchClassificationWindowsSecurity     VMGuestPatchClassificationWindows = "Security"
	VMGuestPatchClassificationWindowsServicePack  VMGuestPatchClassificationWindows = "ServicePack"
	VMGuestPatchClassificationWindowsTools        VMGuestPatchClassificationWindows = "Tools"
	VMGuestPatchClassificationWindowsUpdateRollUp VMGuestPatchClassificationWindows = "UpdateRollUp"
	VMGuestPatchClassificationWindowsUpdates      VMGuestPatchClassificationWindows = "Updates"
)

// PossibleVMGuestPatchClassificationWindowsValues returns the possible values for the VMGuestPatchClassificationWindows const type.
func PossibleVMGuestPatchClassificationWindowsValues() []VMGuestPatchClassificationWindows {
	return []VMGuestPatchClassificationWindows{
		VMGuestPatchClassificationWindowsCritical,
		VMGuestPatchClassificationWindowsDefinition,
		VMGuestPatchClassificationWindowsFeaturePack,
		VMGuestPatchClassificationWindowsSecurity,
		VMGuestPatchClassificationWindowsServicePack,
		VMGuestPatchClassificationWindowsTools,
		VMGuestPatchClassificationWindowsUpdateRollUp,
		VMGuestPatchClassificationWindowsUpdates,
	}
}

// VMGuestPatchRebootSetting - Defines when it is acceptable to reboot a VM during a software update operation.
type VMGuestPatchRebootSetting string

const (
	VMGuestPatchRebootSettingAlways     VMGuestPatchRebootSetting = "Always"
	VMGuestPatchRebootSettingIfRequired VMGuestPatchRebootSetting = "IfRequired"
	VMGuestPatchRebootSettingNever      VMGuestPatchRebootSetting = "Never"
)

// PossibleVMGuestPatchRebootSettingValues returns the possible values for the VMGuestPatchRebootSetting const type.
func PossibleVMGuestPatchRebootSettingValues() []VMGuestPatchRebootSetting {
	return []VMGuestPatchRebootSetting{
		VMGuestPatchRebootSettingAlways,
		VMGuestPatchRebootSettingIfRequired,
		VMGuestPatchRebootSettingNever,
	}
}

// VMGuestPatchRebootStatus - The reboot state of the VM following completion of the operation.
type VMGuestPatchRebootStatus string

const (
	VMGuestPatchRebootStatusCompleted VMGuestPatchRebootStatus = "Completed"
	VMGuestPatchRebootStatusFailed    VMGuestPatchRebootStatus = "Failed"
	VMGuestPatchRebootStatusNotNeeded VMGuestPatchRebootStatus = "NotNeeded"
	VMGuestPatchRebootStatusRequired  VMGuestPatchRebootStatus = "Required"
	VMGuestPatchRebootStatusStarted   VMGuestPatchRebootStatus = "Started"
	VMGuestPatchRebootStatusUnknown   VMGuestPatchRebootStatus = "Unknown"
)

// PossibleVMGuestPatchRebootStatusValues returns the possible values for the VMGuestPatchRebootStatus const type.
func PossibleVMGuestPatchRebootStatusValues() []VMGuestPatchRebootStatus {
	return []VMGuestPatchRebootStatus{
		VMGuestPatchRebootStatusCompleted,
		VMGuestPatchRebootStatusFailed,
		VMGuestPatchRebootStatusNotNeeded,
		VMGuestPatchRebootStatusRequired,
		VMGuestPatchRebootStatusStarted,
		VMGuestPatchRebootStatusUnknown,
	}
}
