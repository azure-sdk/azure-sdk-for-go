//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	moduleVersion = "v1.3.0-beta.3"
)

// Action - The action of IP ACL rule.
type Action string

const (
	ActionAllow Action = "Allow"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionAllow,
	}
}

// ActionsRequired - A message indicating if changes on the service provider require any updates on the consumer.
type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

// PossibleActionsRequiredValues returns the possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{
		ActionsRequiredNone,
		ActionsRequiredRecreate,
	}
}

// ActivationStatus - The activation status of the connected registry.
type ActivationStatus string

const (
	ActivationStatusActive   ActivationStatus = "Active"
	ActivationStatusInactive ActivationStatus = "Inactive"
)

// PossibleActivationStatusValues returns the possible values for the ActivationStatus const type.
func PossibleActivationStatusValues() []ActivationStatus {
	return []ActivationStatus{
		ActivationStatusActive,
		ActivationStatusInactive,
	}
}

// Architecture - The OS architecture.
type Architecture string

const (
	ArchitectureAmd64                 Architecture = "amd64"
	ArchitectureArm                   Architecture = "arm"
	ArchitectureArm64                 Architecture = "arm64"
	ArchitectureThreeHundredEightySix Architecture = "386"
	ArchitectureX86                   Architecture = "x86"
)

// PossibleArchitectureValues returns the possible values for the Architecture const type.
func PossibleArchitectureValues() []Architecture {
	return []Architecture{
		ArchitectureAmd64,
		ArchitectureArm,
		ArchitectureArm64,
		ArchitectureThreeHundredEightySix,
		ArchitectureX86,
	}
}

// AuditLogStatus - Indicates whether audit logs are enabled on the connected registry.
type AuditLogStatus string

const (
	AuditLogStatusDisabled AuditLogStatus = "Disabled"
	AuditLogStatusEnabled  AuditLogStatus = "Enabled"
)

// PossibleAuditLogStatusValues returns the possible values for the AuditLogStatus const type.
func PossibleAuditLogStatusValues() []AuditLogStatus {
	return []AuditLogStatus{
		AuditLogStatusDisabled,
		AuditLogStatusEnabled,
	}
}

// AutoGeneratedDomainNameLabelScope - The auto generated domain name label of the container registry. This value defaults
// to "Unsecure".
type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScopeNoReuse            AutoGeneratedDomainNameLabelScope = "NoReuse"
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse  AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	AutoGeneratedDomainNameLabelScopeTenantReuse        AutoGeneratedDomainNameLabelScope = "TenantReuse"
	AutoGeneratedDomainNameLabelScopeUnsecure           AutoGeneratedDomainNameLabelScope = "Unsecure"
)

// PossibleAutoGeneratedDomainNameLabelScopeValues returns the possible values for the AutoGeneratedDomainNameLabelScope const type.
func PossibleAutoGeneratedDomainNameLabelScopeValues() []AutoGeneratedDomainNameLabelScope {
	return []AutoGeneratedDomainNameLabelScope{
		AutoGeneratedDomainNameLabelScopeNoReuse,
		AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		AutoGeneratedDomainNameLabelScopeTenantReuse,
		AutoGeneratedDomainNameLabelScopeUnsecure,
	}
}

// AzureADAuthenticationAsArmPolicyStatus - The value that indicates whether the policy is enabled or not.
type AzureADAuthenticationAsArmPolicyStatus string

const (
	AzureADAuthenticationAsArmPolicyStatusDisabled AzureADAuthenticationAsArmPolicyStatus = "disabled"
	AzureADAuthenticationAsArmPolicyStatusEnabled  AzureADAuthenticationAsArmPolicyStatus = "enabled"
)

// PossibleAzureADAuthenticationAsArmPolicyStatusValues returns the possible values for the AzureADAuthenticationAsArmPolicyStatus const type.
func PossibleAzureADAuthenticationAsArmPolicyStatusValues() []AzureADAuthenticationAsArmPolicyStatus {
	return []AzureADAuthenticationAsArmPolicyStatus{
		AzureADAuthenticationAsArmPolicyStatusDisabled,
		AzureADAuthenticationAsArmPolicyStatusEnabled,
	}
}

// BaseImageDependencyType - The type of the base image dependency.
type BaseImageDependencyType string

const (
	BaseImageDependencyTypeBuildTime BaseImageDependencyType = "BuildTime"
	BaseImageDependencyTypeRunTime   BaseImageDependencyType = "RunTime"
)

// PossibleBaseImageDependencyTypeValues returns the possible values for the BaseImageDependencyType const type.
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return []BaseImageDependencyType{
		BaseImageDependencyTypeBuildTime,
		BaseImageDependencyTypeRunTime,
	}
}

// BaseImageTriggerType - The type of the auto trigger for base image dependency updates.
type BaseImageTriggerType string

const (
	BaseImageTriggerTypeAll     BaseImageTriggerType = "All"
	BaseImageTriggerTypeRuntime BaseImageTriggerType = "Runtime"
)

// PossibleBaseImageTriggerTypeValues returns the possible values for the BaseImageTriggerType const type.
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return []BaseImageTriggerType{
		BaseImageTriggerTypeAll,
		BaseImageTriggerTypeRuntime,
	}
}

// CertificateType - The type of certificate location.
type CertificateType string

const (
	CertificateTypeLocalDirectory CertificateType = "LocalDirectory"
)

// PossibleCertificateTypeValues returns the possible values for the CertificateType const type.
func PossibleCertificateTypeValues() []CertificateType {
	return []CertificateType{
		CertificateTypeLocalDirectory,
	}
}

// ConnectedRegistryMode - The mode of the connected registry resource that indicates the permissions of the registry.
type ConnectedRegistryMode string

const (
	ConnectedRegistryModeMirror    ConnectedRegistryMode = "Mirror"
	ConnectedRegistryModeReadOnly  ConnectedRegistryMode = "ReadOnly"
	ConnectedRegistryModeReadWrite ConnectedRegistryMode = "ReadWrite"
	ConnectedRegistryModeRegistry  ConnectedRegistryMode = "Registry"
)

// PossibleConnectedRegistryModeValues returns the possible values for the ConnectedRegistryMode const type.
func PossibleConnectedRegistryModeValues() []ConnectedRegistryMode {
	return []ConnectedRegistryMode{
		ConnectedRegistryModeMirror,
		ConnectedRegistryModeReadOnly,
		ConnectedRegistryModeReadWrite,
		ConnectedRegistryModeRegistry,
	}
}

// ConnectionState - The current connection state of the connected registry.
type ConnectionState string

const (
	ConnectionStateOffline   ConnectionState = "Offline"
	ConnectionStateOnline    ConnectionState = "Online"
	ConnectionStateSyncing   ConnectionState = "Syncing"
	ConnectionStateUnhealthy ConnectionState = "Unhealthy"
)

// PossibleConnectionStateValues returns the possible values for the ConnectionState const type.
func PossibleConnectionStateValues() []ConnectionState {
	return []ConnectionState{
		ConnectionStateOffline,
		ConnectionStateOnline,
		ConnectionStateSyncing,
		ConnectionStateUnhealthy,
	}
}

// ConnectionStatus - The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

// PossibleConnectionStatusValues returns the possible values for the ConnectionStatus const type.
func PossibleConnectionStatusValues() []ConnectionStatus {
	return []ConnectionStatus{
		ConnectionStatusApproved,
		ConnectionStatusDisconnected,
		ConnectionStatusPending,
		ConnectionStatusRejected,
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

// CredentialHealthStatus - The health status of credential.
type CredentialHealthStatus string

const (
	CredentialHealthStatusHealthy   CredentialHealthStatus = "Healthy"
	CredentialHealthStatusUnhealthy CredentialHealthStatus = "Unhealthy"
)

// PossibleCredentialHealthStatusValues returns the possible values for the CredentialHealthStatus const type.
func PossibleCredentialHealthStatusValues() []CredentialHealthStatus {
	return []CredentialHealthStatus{
		CredentialHealthStatusHealthy,
		CredentialHealthStatusUnhealthy,
	}
}

// CredentialName - The name of the credential.
type CredentialName string

const (
	CredentialNameCredential1 CredentialName = "Credential1"
)

// PossibleCredentialNameValues returns the possible values for the CredentialName const type.
func PossibleCredentialNameValues() []CredentialName {
	return []CredentialName{
		CredentialNameCredential1,
	}
}

// DefaultAction - The default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// EncryptionStatus - Indicates whether or not the encryption is enabled for container registry.
type EncryptionStatus string

const (
	EncryptionStatusDisabled EncryptionStatus = "disabled"
	EncryptionStatusEnabled  EncryptionStatus = "enabled"
)

// PossibleEncryptionStatusValues returns the possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{
		EncryptionStatusDisabled,
		EncryptionStatusEnabled,
	}
}

// ExportPolicyStatus - The value that indicates whether the policy is enabled or not.
type ExportPolicyStatus string

const (
	ExportPolicyStatusDisabled ExportPolicyStatus = "disabled"
	ExportPolicyStatusEnabled  ExportPolicyStatus = "enabled"
)

// PossibleExportPolicyStatusValues returns the possible values for the ExportPolicyStatus const type.
func PossibleExportPolicyStatusValues() []ExportPolicyStatus {
	return []ExportPolicyStatus{
		ExportPolicyStatusDisabled,
		ExportPolicyStatusEnabled,
	}
}

// ImportMode - When Force, any existing target tags will be overwritten. When NoForce, any existing target tags will fail
// the operation before any copying begins.
type ImportMode string

const (
	ImportModeForce   ImportMode = "Force"
	ImportModeNoForce ImportMode = "NoForce"
)

// PossibleImportModeValues returns the possible values for the ImportMode const type.
func PossibleImportModeValues() []ImportMode {
	return []ImportMode{
		ImportModeForce,
		ImportModeNoForce,
	}
}

// LastModifiedByType - The type of identity that last modified the resource.
type LastModifiedByType string

const (
	LastModifiedByTypeApplication     LastModifiedByType = "Application"
	LastModifiedByTypeKey             LastModifiedByType = "Key"
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	LastModifiedByTypeUser            LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns the possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{
		LastModifiedByTypeApplication,
		LastModifiedByTypeKey,
		LastModifiedByTypeManagedIdentity,
		LastModifiedByTypeUser,
	}
}

// LogLevel - The verbosity of logs persisted on the connected registry.
type LogLevel string

const (
	LogLevelDebug       LogLevel = "Debug"
	LogLevelError       LogLevel = "Error"
	LogLevelInformation LogLevel = "Information"
	LogLevelNone        LogLevel = "None"
	LogLevelWarning     LogLevel = "Warning"
)

// PossibleLogLevelValues returns the possible values for the LogLevel const type.
func PossibleLogLevelValues() []LogLevel {
	return []LogLevel{
		LogLevelDebug,
		LogLevelError,
		LogLevelInformation,
		LogLevelNone,
		LogLevelWarning,
	}
}

// MetadataSearch - Determines whether registry artifacts are indexed for metadata search.
type MetadataSearch string

const (
	MetadataSearchDisabled MetadataSearch = "Disabled"
	MetadataSearchEnabled  MetadataSearch = "Enabled"
)

// PossibleMetadataSearchValues returns the possible values for the MetadataSearch const type.
func PossibleMetadataSearchValues() []MetadataSearch {
	return []MetadataSearch{
		MetadataSearchDisabled,
		MetadataSearchEnabled,
	}
}

// NetworkRuleBypassOptions - Whether to allow trusted Azure services to access a network restricted registry.
type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

// PossibleNetworkRuleBypassOptionsValues returns the possible values for the NetworkRuleBypassOptions const type.
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{
		NetworkRuleBypassOptionsAzureServices,
		NetworkRuleBypassOptionsNone,
	}
}

// OS - The OS of agent machine
type OS string

const (
	OSLinux   OS = "Linux"
	OSWindows OS = "Windows"
)

// PossibleOSValues returns the possible values for the OS const type.
func PossibleOSValues() []OS {
	return []OS{
		OSLinux,
		OSWindows,
	}
}

// PackageSourceType - The type of package source for a archive.
type PackageSourceType string

const (
	PackageSourceTypeRemote PackageSourceType = "remote"
)

// PossiblePackageSourceTypeValues returns the possible values for the PackageSourceType const type.
func PossiblePackageSourceTypeValues() []PackageSourceType {
	return []PackageSourceType{
		PackageSourceTypeRemote,
	}
}

// PasswordName - The password name.
type PasswordName string

const (
	PasswordNamePassword  PasswordName = "password"
	PasswordNamePassword2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns the possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{
		PasswordNamePassword,
		PasswordNamePassword2,
	}
}

type PipelineOptions string

const (
	PipelineOptionsContinueOnErrors          PipelineOptions = "ContinueOnErrors"
	PipelineOptionsDeleteSourceBlobOnSuccess PipelineOptions = "DeleteSourceBlobOnSuccess"
	PipelineOptionsOverwriteBlobs            PipelineOptions = "OverwriteBlobs"
	PipelineOptionsOverwriteTags             PipelineOptions = "OverwriteTags"
)

// PossiblePipelineOptionsValues returns the possible values for the PipelineOptions const type.
func PossiblePipelineOptionsValues() []PipelineOptions {
	return []PipelineOptions{
		PipelineOptionsContinueOnErrors,
		PipelineOptionsDeleteSourceBlobOnSuccess,
		PipelineOptionsOverwriteBlobs,
		PipelineOptionsOverwriteTags,
	}
}

// PipelineRunSourceType - The type of the source.
type PipelineRunSourceType string

const (
	PipelineRunSourceTypeAzureStorageBlob PipelineRunSourceType = "AzureStorageBlob"
)

// PossiblePipelineRunSourceTypeValues returns the possible values for the PipelineRunSourceType const type.
func PossiblePipelineRunSourceTypeValues() []PipelineRunSourceType {
	return []PipelineRunSourceType{
		PipelineRunSourceTypeAzureStorageBlob,
	}
}

// PipelineRunTargetType - The type of the target.
type PipelineRunTargetType string

const (
	PipelineRunTargetTypeAzureStorageBlob PipelineRunTargetType = "AzureStorageBlob"
)

// PossiblePipelineRunTargetTypeValues returns the possible values for the PipelineRunTargetType const type.
func PossiblePipelineRunTargetTypeValues() []PipelineRunTargetType {
	return []PipelineRunTargetType{
		PipelineRunTargetTypeAzureStorageBlob,
	}
}

// PipelineSourceType - The type of source for the import pipeline.
type PipelineSourceType string

const (
	PipelineSourceTypeAzureStorageBlobContainer PipelineSourceType = "AzureStorageBlobContainer"
)

// PossiblePipelineSourceTypeValues returns the possible values for the PipelineSourceType const type.
func PossiblePipelineSourceTypeValues() []PipelineSourceType {
	return []PipelineSourceType{
		PipelineSourceTypeAzureStorageBlobContainer,
	}
}

// PolicyStatus - The value that indicates whether the policy is enabled or not.
type PolicyStatus string

const (
	PolicyStatusDisabled PolicyStatus = "disabled"
	PolicyStatusEnabled  PolicyStatus = "enabled"
)

// PossiblePolicyStatusValues returns the possible values for the PolicyStatus const type.
func PossiblePolicyStatusValues() []PolicyStatus {
	return []PolicyStatus{
		PolicyStatusDisabled,
		PolicyStatusEnabled,
	}
}

// ProvisioningState - The provisioning state of this agent pool
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether or not public network access is allowed for the container registry.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// RegistryUsageUnit - The unit of measurement.
type RegistryUsageUnit string

const (
	RegistryUsageUnitBytes RegistryUsageUnit = "Bytes"
	RegistryUsageUnitCount RegistryUsageUnit = "Count"
)

// PossibleRegistryUsageUnitValues returns the possible values for the RegistryUsageUnit const type.
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return []RegistryUsageUnit{
		RegistryUsageUnitBytes,
		RegistryUsageUnitCount,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// RunStatus - The current status of the run.
type RunStatus string

const (
	RunStatusCanceled  RunStatus = "Canceled"
	RunStatusError     RunStatus = "Error"
	RunStatusFailed    RunStatus = "Failed"
	RunStatusQueued    RunStatus = "Queued"
	RunStatusRunning   RunStatus = "Running"
	RunStatusStarted   RunStatus = "Started"
	RunStatusSucceeded RunStatus = "Succeeded"
	RunStatusTimeout   RunStatus = "Timeout"
)

// PossibleRunStatusValues returns the possible values for the RunStatus const type.
func PossibleRunStatusValues() []RunStatus {
	return []RunStatus{
		RunStatusCanceled,
		RunStatusError,
		RunStatusFailed,
		RunStatusQueued,
		RunStatusRunning,
		RunStatusStarted,
		RunStatusSucceeded,
		RunStatusTimeout,
	}
}

// RunType - The type of run.
type RunType string

const (
	RunTypeAutoBuild  RunType = "AutoBuild"
	RunTypeAutoRun    RunType = "AutoRun"
	RunTypeQuickBuild RunType = "QuickBuild"
	RunTypeQuickRun   RunType = "QuickRun"
)

// PossibleRunTypeValues returns the possible values for the RunType const type.
func PossibleRunTypeValues() []RunType {
	return []RunType{
		RunTypeAutoBuild,
		RunTypeAutoRun,
		RunTypeQuickBuild,
		RunTypeQuickRun,
	}
}

// SKUName - The SKU name of the container registry. Required for registry creation.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameClassic  SKUName = "Classic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameClassic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// SKUTier - The SKU tier based on the SKU name.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierClassic  SKUTier = "Classic"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierClassic,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// SecretObjectType - The type of the secret object which determines how the value of the secret object has to be interpreted.
type SecretObjectType string

const (
	SecretObjectTypeOpaque      SecretObjectType = "Opaque"
	SecretObjectTypeVaultsecret SecretObjectType = "Vaultsecret"
)

// PossibleSecretObjectTypeValues returns the possible values for the SecretObjectType const type.
func PossibleSecretObjectTypeValues() []SecretObjectType {
	return []SecretObjectType{
		SecretObjectTypeOpaque,
		SecretObjectTypeVaultsecret,
	}
}

// SourceControlType - The type of source control service.
type SourceControlType string

const (
	SourceControlTypeGithub                  SourceControlType = "Github"
	SourceControlTypeVisualStudioTeamService SourceControlType = "VisualStudioTeamService"
)

// PossibleSourceControlTypeValues returns the possible values for the SourceControlType const type.
func PossibleSourceControlTypeValues() []SourceControlType {
	return []SourceControlType{
		SourceControlTypeGithub,
		SourceControlTypeVisualStudioTeamService,
	}
}

// SourceRegistryLoginMode - The authentication mode which determines the source registry login scope. The credentials for
// the source registry will be generated using the given scope. These credentials will be used to login to
// the source registry during the run.
type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = "Default"
	SourceRegistryLoginModeNone    SourceRegistryLoginMode = "None"
)

// PossibleSourceRegistryLoginModeValues returns the possible values for the SourceRegistryLoginMode const type.
func PossibleSourceRegistryLoginModeValues() []SourceRegistryLoginMode {
	return []SourceRegistryLoginMode{
		SourceRegistryLoginModeDefault,
		SourceRegistryLoginModeNone,
	}
}

type SourceTriggerEvent string

const (
	SourceTriggerEventCommit      SourceTriggerEvent = "commit"
	SourceTriggerEventPullrequest SourceTriggerEvent = "pullrequest"
)

// PossibleSourceTriggerEventValues returns the possible values for the SourceTriggerEvent const type.
func PossibleSourceTriggerEventValues() []SourceTriggerEvent {
	return []SourceTriggerEvent{
		SourceTriggerEventCommit,
		SourceTriggerEventPullrequest,
	}
}

// StepType - The type of the step.
type StepType string

const (
	StepTypeDocker      StepType = "Docker"
	StepTypeEncodedTask StepType = "EncodedTask"
	StepTypeFileTask    StepType = "FileTask"
)

// PossibleStepTypeValues returns the possible values for the StepType const type.
func PossibleStepTypeValues() []StepType {
	return []StepType{
		StepTypeDocker,
		StepTypeEncodedTask,
		StepTypeFileTask,
	}
}

// TLSStatus - Indicates whether HTTPS is enabled for the login server.
type TLSStatus string

const (
	TLSStatusDisabled TLSStatus = "Disabled"
	TLSStatusEnabled  TLSStatus = "Enabled"
)

// PossibleTLSStatusValues returns the possible values for the TLSStatus const type.
func PossibleTLSStatusValues() []TLSStatus {
	return []TLSStatus{
		TLSStatusDisabled,
		TLSStatusEnabled,
	}
}

// TaskStatus - The current status of task.
type TaskStatus string

const (
	TaskStatusDisabled TaskStatus = "Disabled"
	TaskStatusEnabled  TaskStatus = "Enabled"
)

// PossibleTaskStatusValues returns the possible values for the TaskStatus const type.
func PossibleTaskStatusValues() []TaskStatus {
	return []TaskStatus{
		TaskStatusDisabled,
		TaskStatusEnabled,
	}
}

type TokenCertificateName string

const (
	TokenCertificateNameCertificate1 TokenCertificateName = "certificate1"
	TokenCertificateNameCertificate2 TokenCertificateName = "certificate2"
)

// PossibleTokenCertificateNameValues returns the possible values for the TokenCertificateName const type.
func PossibleTokenCertificateNameValues() []TokenCertificateName {
	return []TokenCertificateName{
		TokenCertificateNameCertificate1,
		TokenCertificateNameCertificate2,
	}
}

// TokenPasswordName - The password name "password1" or "password2"
type TokenPasswordName string

const (
	TokenPasswordNamePassword1 TokenPasswordName = "password1"
	TokenPasswordNamePassword2 TokenPasswordName = "password2"
)

// PossibleTokenPasswordNameValues returns the possible values for the TokenPasswordName const type.
func PossibleTokenPasswordNameValues() []TokenPasswordName {
	return []TokenPasswordName{
		TokenPasswordNamePassword1,
		TokenPasswordNamePassword2,
	}
}

// TokenStatus - The status of the token example enabled or disabled.
type TokenStatus string

const (
	TokenStatusDisabled TokenStatus = "disabled"
	TokenStatusEnabled  TokenStatus = "enabled"
)

// PossibleTokenStatusValues returns the possible values for the TokenStatus const type.
func PossibleTokenStatusValues() []TokenStatus {
	return []TokenStatus{
		TokenStatusDisabled,
		TokenStatusEnabled,
	}
}

// TokenType - The type of Auth token.
type TokenType string

const (
	TokenTypeOAuth TokenType = "OAuth"
	TokenTypePAT   TokenType = "PAT"
)

// PossibleTokenTypeValues returns the possible values for the TokenType const type.
func PossibleTokenTypeValues() []TokenType {
	return []TokenType{
		TokenTypeOAuth,
		TokenTypePAT,
	}
}

// TriggerStatus - The current status of trigger.
type TriggerStatus string

const (
	TriggerStatusDisabled TriggerStatus = "Disabled"
	TriggerStatusEnabled  TriggerStatus = "Enabled"
)

// PossibleTriggerStatusValues returns the possible values for the TriggerStatus const type.
func PossibleTriggerStatusValues() []TriggerStatus {
	return []TriggerStatus{
		TriggerStatusDisabled,
		TriggerStatusEnabled,
	}
}

// TrustPolicyType - The type of trust policy.
type TrustPolicyType string

const (
	TrustPolicyTypeNotary TrustPolicyType = "Notary"
)

// PossibleTrustPolicyTypeValues returns the possible values for the TrustPolicyType const type.
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return []TrustPolicyType{
		TrustPolicyTypeNotary,
	}
}

// UpdateTriggerPayloadType - Type of Payload body for Base image update triggers.
type UpdateTriggerPayloadType string

const (
	UpdateTriggerPayloadTypeDefault UpdateTriggerPayloadType = "Default"
	UpdateTriggerPayloadTypeToken   UpdateTriggerPayloadType = "Token"
)

// PossibleUpdateTriggerPayloadTypeValues returns the possible values for the UpdateTriggerPayloadType const type.
func PossibleUpdateTriggerPayloadTypeValues() []UpdateTriggerPayloadType {
	return []UpdateTriggerPayloadType{
		UpdateTriggerPayloadTypeDefault,
		UpdateTriggerPayloadTypeToken,
	}
}

// Variant - Variant of the CPU.
type Variant string

const (
	VariantV6 Variant = "v6"
	VariantV7 Variant = "v7"
	VariantV8 Variant = "v8"
)

// PossibleVariantValues returns the possible values for the Variant const type.
func PossibleVariantValues() []Variant {
	return []Variant{
		VariantV6,
		VariantV7,
		VariantV8,
	}
}

type WebhookAction string

const (
	WebhookActionChartDelete WebhookAction = "chart_delete"
	WebhookActionChartPush   WebhookAction = "chart_push"
	WebhookActionDelete      WebhookAction = "delete"
	WebhookActionPush        WebhookAction = "push"
	WebhookActionQuarantine  WebhookAction = "quarantine"
)

// PossibleWebhookActionValues returns the possible values for the WebhookAction const type.
func PossibleWebhookActionValues() []WebhookAction {
	return []WebhookAction{
		WebhookActionChartDelete,
		WebhookActionChartPush,
		WebhookActionDelete,
		WebhookActionPush,
		WebhookActionQuarantine,
	}
}

// WebhookStatus - The status of the webhook at the time the operation was called.
type WebhookStatus string

const (
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusEnabled  WebhookStatus = "enabled"
)

// PossibleWebhookStatusValues returns the possible values for the WebhookStatus const type.
func PossibleWebhookStatusValues() []WebhookStatus {
	return []WebhookStatus{
		WebhookStatusDisabled,
		WebhookStatusEnabled,
	}
}

// ZoneRedundancy - Whether or not zone redundancy is enabled for this container registry
type ZoneRedundancy string

const (
	ZoneRedundancyDisabled ZoneRedundancy = "Disabled"
	ZoneRedundancyEnabled  ZoneRedundancy = "Enabled"
)

// PossibleZoneRedundancyValues returns the possible values for the ZoneRedundancy const type.
func PossibleZoneRedundancyValues() []ZoneRedundancy {
	return []ZoneRedundancy{
		ZoneRedundancyDisabled,
		ZoneRedundancyEnabled,
	}
}
