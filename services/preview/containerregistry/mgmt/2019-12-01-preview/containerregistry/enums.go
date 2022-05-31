package containerregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Action enumerates the values for action.
type Action string

const (
	// Allow ...
	Allow Action = "Allow"
)

// PossibleActionValues returns an array of possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{Allow}
}

// ActionsRequired enumerates the values for actions required.
type ActionsRequired string

const (
	// None ...
	None ActionsRequired = "None"
	// Recreate ...
	Recreate ActionsRequired = "Recreate"
)

// PossibleActionsRequiredValues returns an array of possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{None, Recreate}
}

// Architecture enumerates the values for architecture.
type Architecture string

const (
	// Amd64 ...
	Amd64 Architecture = "amd64"
	// Arm ...
	Arm Architecture = "arm"
	// Arm64 ...
	Arm64 Architecture = "arm64"
	// ThreeEightSix ...
	ThreeEightSix Architecture = "386"
	// X86 ...
	X86 Architecture = "x86"
)

// PossibleArchitectureValues returns an array of possible values for the Architecture const type.
func PossibleArchitectureValues() []Architecture {
	return []Architecture{Amd64, Arm, Arm64, ThreeEightSix, X86}
}

// BaseImageDependencyType enumerates the values for base image dependency type.
type BaseImageDependencyType string

const (
	// BuildTime ...
	BuildTime BaseImageDependencyType = "BuildTime"
	// RunTime ...
	RunTime BaseImageDependencyType = "RunTime"
)

// PossibleBaseImageDependencyTypeValues returns an array of possible values for the BaseImageDependencyType const type.
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return []BaseImageDependencyType{BuildTime, RunTime}
}

// BaseImageTriggerType enumerates the values for base image trigger type.
type BaseImageTriggerType string

const (
	// All ...
	All BaseImageTriggerType = "All"
	// Runtime ...
	Runtime BaseImageTriggerType = "Runtime"
)

// PossibleBaseImageTriggerTypeValues returns an array of possible values for the BaseImageTriggerType const type.
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return []BaseImageTriggerType{All, Runtime}
}

// ConnectionStatus enumerates the values for connection status.
type ConnectionStatus string

const (
	// Approved ...
	Approved ConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected ConnectionStatus = "Disconnected"
	// Pending ...
	Pending ConnectionStatus = "Pending"
	// Rejected ...
	Rejected ConnectionStatus = "Rejected"
)

// PossibleConnectionStatusValues returns an array of possible values for the ConnectionStatus const type.
func PossibleConnectionStatusValues() []ConnectionStatus {
	return []ConnectionStatus{Approved, Disconnected, Pending, Rejected}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// DefaultActionAllow ...
	DefaultActionAllow DefaultAction = "Allow"
	// DefaultActionDeny ...
	DefaultActionDeny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{DefaultActionAllow, DefaultActionDeny}
}

// EncryptionStatus enumerates the values for encryption status.
type EncryptionStatus string

const (
	// Disabled ...
	Disabled EncryptionStatus = "disabled"
	// Enabled ...
	Enabled EncryptionStatus = "enabled"
)

// PossibleEncryptionStatusValues returns an array of possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{Disabled, Enabled}
}

// ImportMode enumerates the values for import mode.
type ImportMode string

const (
	// Force ...
	Force ImportMode = "Force"
	// NoForce ...
	NoForce ImportMode = "NoForce"
)

// PossibleImportModeValues returns an array of possible values for the ImportMode const type.
func PossibleImportModeValues() []ImportMode {
	return []ImportMode{Force, NoForce}
}

// LastModifiedByType enumerates the values for last modified by type.
type LastModifiedByType string

const (
	// LastModifiedByTypeApplication ...
	LastModifiedByTypeApplication LastModifiedByType = "Application"
	// LastModifiedByTypeKey ...
	LastModifiedByTypeKey LastModifiedByType = "Key"
	// LastModifiedByTypeManagedIdentity ...
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	// LastModifiedByTypeUser ...
	LastModifiedByTypeUser LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns an array of possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{LastModifiedByTypeApplication, LastModifiedByTypeKey, LastModifiedByTypeManagedIdentity, LastModifiedByTypeUser}
}

// NetworkRuleBypassOptions enumerates the values for network rule bypass options.
type NetworkRuleBypassOptions string

const (
	// NetworkRuleBypassOptionsAzureServices ...
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	// NetworkRuleBypassOptionsNone ...
	NetworkRuleBypassOptionsNone NetworkRuleBypassOptions = "None"
)

// PossibleNetworkRuleBypassOptionsValues returns an array of possible values for the NetworkRuleBypassOptions const type.
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{NetworkRuleBypassOptionsAzureServices, NetworkRuleBypassOptionsNone}
}

// OS enumerates the values for os.
type OS string

const (
	// Linux ...
	Linux OS = "Linux"
	// Windows ...
	Windows OS = "Windows"
)

// PossibleOSValues returns an array of possible values for the OS const type.
func PossibleOSValues() []OS {
	return []OS{Linux, Windows}
}

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password ...
	Password PasswordName = "password"
	// Password2 ...
	Password2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns an array of possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{Password, Password2}
}

// PipelineOptions enumerates the values for pipeline options.
type PipelineOptions string

const (
	// ContinueOnErrors ...
	ContinueOnErrors PipelineOptions = "ContinueOnErrors"
	// DeleteSourceBlobOnSuccess ...
	DeleteSourceBlobOnSuccess PipelineOptions = "DeleteSourceBlobOnSuccess"
	// OverwriteBlobs ...
	OverwriteBlobs PipelineOptions = "OverwriteBlobs"
	// OverwriteTags ...
	OverwriteTags PipelineOptions = "OverwriteTags"
)

// PossiblePipelineOptionsValues returns an array of possible values for the PipelineOptions const type.
func PossiblePipelineOptionsValues() []PipelineOptions {
	return []PipelineOptions{ContinueOnErrors, DeleteSourceBlobOnSuccess, OverwriteBlobs, OverwriteTags}
}

// PipelineRunSourceType enumerates the values for pipeline run source type.
type PipelineRunSourceType string

const (
	// AzureStorageBlob ...
	AzureStorageBlob PipelineRunSourceType = "AzureStorageBlob"
)

// PossiblePipelineRunSourceTypeValues returns an array of possible values for the PipelineRunSourceType const type.
func PossiblePipelineRunSourceTypeValues() []PipelineRunSourceType {
	return []PipelineRunSourceType{AzureStorageBlob}
}

// PipelineRunTargetType enumerates the values for pipeline run target type.
type PipelineRunTargetType string

const (
	// PipelineRunTargetTypeAzureStorageBlob ...
	PipelineRunTargetTypeAzureStorageBlob PipelineRunTargetType = "AzureStorageBlob"
)

// PossiblePipelineRunTargetTypeValues returns an array of possible values for the PipelineRunTargetType const type.
func PossiblePipelineRunTargetTypeValues() []PipelineRunTargetType {
	return []PipelineRunTargetType{PipelineRunTargetTypeAzureStorageBlob}
}

// PipelineSourceType enumerates the values for pipeline source type.
type PipelineSourceType string

const (
	// AzureStorageBlobContainer ...
	AzureStorageBlobContainer PipelineSourceType = "AzureStorageBlobContainer"
)

// PossiblePipelineSourceTypeValues returns an array of possible values for the PipelineSourceType const type.
func PossiblePipelineSourceTypeValues() []PipelineSourceType {
	return []PipelineSourceType{AzureStorageBlobContainer}
}

// PolicyStatus enumerates the values for policy status.
type PolicyStatus string

const (
	// PolicyStatusDisabled ...
	PolicyStatusDisabled PolicyStatus = "disabled"
	// PolicyStatusEnabled ...
	PolicyStatusEnabled PolicyStatus = "enabled"
)

// PossiblePolicyStatusValues returns an array of possible values for the PolicyStatus const type.
func PossiblePolicyStatusValues() []PolicyStatus {
	return []PolicyStatus{PolicyStatusDisabled, PolicyStatusEnabled}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Creating, Deleting, Failed, Succeeded, Updating}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled ...
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	// PublicNetworkAccessEnabled ...
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{PublicNetworkAccessDisabled, PublicNetworkAccessEnabled}
}

// RegistryUsageUnit enumerates the values for registry usage unit.
type RegistryUsageUnit string

const (
	// Bytes ...
	Bytes RegistryUsageUnit = "Bytes"
	// Count ...
	Count RegistryUsageUnit = "Count"
)

// PossibleRegistryUsageUnitValues returns an array of possible values for the RegistryUsageUnit const type.
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return []RegistryUsageUnit{Bytes, Count}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone ...
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned ...
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeSystemAssignedUserAssigned ...
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// ResourceIdentityTypeUserAssigned ...
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{ResourceIdentityTypeNone, ResourceIdentityTypeSystemAssigned, ResourceIdentityTypeSystemAssignedUserAssigned, ResourceIdentityTypeUserAssigned}
}

// RunStatus enumerates the values for run status.
type RunStatus string

const (
	// RunStatusCanceled ...
	RunStatusCanceled RunStatus = "Canceled"
	// RunStatusError ...
	RunStatusError RunStatus = "Error"
	// RunStatusFailed ...
	RunStatusFailed RunStatus = "Failed"
	// RunStatusQueued ...
	RunStatusQueued RunStatus = "Queued"
	// RunStatusRunning ...
	RunStatusRunning RunStatus = "Running"
	// RunStatusStarted ...
	RunStatusStarted RunStatus = "Started"
	// RunStatusSucceeded ...
	RunStatusSucceeded RunStatus = "Succeeded"
	// RunStatusTimeout ...
	RunStatusTimeout RunStatus = "Timeout"
)

// PossibleRunStatusValues returns an array of possible values for the RunStatus const type.
func PossibleRunStatusValues() []RunStatus {
	return []RunStatus{RunStatusCanceled, RunStatusError, RunStatusFailed, RunStatusQueued, RunStatusRunning, RunStatusStarted, RunStatusSucceeded, RunStatusTimeout}
}

// RunType enumerates the values for run type.
type RunType string

const (
	// AutoBuild ...
	AutoBuild RunType = "AutoBuild"
	// AutoRun ...
	AutoRun RunType = "AutoRun"
	// QuickBuild ...
	QuickBuild RunType = "QuickBuild"
	// QuickRun ...
	QuickRun RunType = "QuickRun"
)

// PossibleRunTypeValues returns an array of possible values for the RunType const type.
func PossibleRunTypeValues() []RunType {
	return []RunType{AutoBuild, AutoRun, QuickBuild, QuickRun}
}

// SecretObjectType enumerates the values for secret object type.
type SecretObjectType string

const (
	// Opaque ...
	Opaque SecretObjectType = "Opaque"
	// Vaultsecret ...
	Vaultsecret SecretObjectType = "Vaultsecret"
)

// PossibleSecretObjectTypeValues returns an array of possible values for the SecretObjectType const type.
func PossibleSecretObjectTypeValues() []SecretObjectType {
	return []SecretObjectType{Opaque, Vaultsecret}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Classic ...
	Classic SkuName = "Classic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Classic, Premium, Standard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierClassic ...
	SkuTierClassic SkuTier = "Classic"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierClassic, SkuTierPremium, SkuTierStandard}
}

// SourceControlType enumerates the values for source control type.
type SourceControlType string

const (
	// Github ...
	Github SourceControlType = "Github"
	// VisualStudioTeamService ...
	VisualStudioTeamService SourceControlType = "VisualStudioTeamService"
)

// PossibleSourceControlTypeValues returns an array of possible values for the SourceControlType const type.
func PossibleSourceControlTypeValues() []SourceControlType {
	return []SourceControlType{Github, VisualStudioTeamService}
}

// SourceRegistryLoginMode enumerates the values for source registry login mode.
type SourceRegistryLoginMode string

const (
	// SourceRegistryLoginModeDefault ...
	SourceRegistryLoginModeDefault SourceRegistryLoginMode = "Default"
	// SourceRegistryLoginModeNone ...
	SourceRegistryLoginModeNone SourceRegistryLoginMode = "None"
)

// PossibleSourceRegistryLoginModeValues returns an array of possible values for the SourceRegistryLoginMode const type.
func PossibleSourceRegistryLoginModeValues() []SourceRegistryLoginMode {
	return []SourceRegistryLoginMode{SourceRegistryLoginModeDefault, SourceRegistryLoginModeNone}
}

// SourceTriggerEvent enumerates the values for source trigger event.
type SourceTriggerEvent string

const (
	// Commit ...
	Commit SourceTriggerEvent = "commit"
	// Pullrequest ...
	Pullrequest SourceTriggerEvent = "pullrequest"
)

// PossibleSourceTriggerEventValues returns an array of possible values for the SourceTriggerEvent const type.
func PossibleSourceTriggerEventValues() []SourceTriggerEvent {
	return []SourceTriggerEvent{Commit, Pullrequest}
}

// TaskStatus enumerates the values for task status.
type TaskStatus string

const (
	// TaskStatusDisabled ...
	TaskStatusDisabled TaskStatus = "Disabled"
	// TaskStatusEnabled ...
	TaskStatusEnabled TaskStatus = "Enabled"
)

// PossibleTaskStatusValues returns an array of possible values for the TaskStatus const type.
func PossibleTaskStatusValues() []TaskStatus {
	return []TaskStatus{TaskStatusDisabled, TaskStatusEnabled}
}

// TokenCertificateName enumerates the values for token certificate name.
type TokenCertificateName string

const (
	// Certificate1 ...
	Certificate1 TokenCertificateName = "certificate1"
	// Certificate2 ...
	Certificate2 TokenCertificateName = "certificate2"
)

// PossibleTokenCertificateNameValues returns an array of possible values for the TokenCertificateName const type.
func PossibleTokenCertificateNameValues() []TokenCertificateName {
	return []TokenCertificateName{Certificate1, Certificate2}
}

// TokenPasswordName enumerates the values for token password name.
type TokenPasswordName string

const (
	// TokenPasswordNamePassword1 ...
	TokenPasswordNamePassword1 TokenPasswordName = "password1"
	// TokenPasswordNamePassword2 ...
	TokenPasswordNamePassword2 TokenPasswordName = "password2"
)

// PossibleTokenPasswordNameValues returns an array of possible values for the TokenPasswordName const type.
func PossibleTokenPasswordNameValues() []TokenPasswordName {
	return []TokenPasswordName{TokenPasswordNamePassword1, TokenPasswordNamePassword2}
}

// TokenStatus enumerates the values for token status.
type TokenStatus string

const (
	// TokenStatusDisabled ...
	TokenStatusDisabled TokenStatus = "disabled"
	// TokenStatusEnabled ...
	TokenStatusEnabled TokenStatus = "enabled"
)

// PossibleTokenStatusValues returns an array of possible values for the TokenStatus const type.
func PossibleTokenStatusValues() []TokenStatus {
	return []TokenStatus{TokenStatusDisabled, TokenStatusEnabled}
}

// TokenType enumerates the values for token type.
type TokenType string

const (
	// OAuth ...
	OAuth TokenType = "OAuth"
	// PAT ...
	PAT TokenType = "PAT"
)

// PossibleTokenTypeValues returns an array of possible values for the TokenType const type.
func PossibleTokenTypeValues() []TokenType {
	return []TokenType{OAuth, PAT}
}

// TriggerStatus enumerates the values for trigger status.
type TriggerStatus string

const (
	// TriggerStatusDisabled ...
	TriggerStatusDisabled TriggerStatus = "Disabled"
	// TriggerStatusEnabled ...
	TriggerStatusEnabled TriggerStatus = "Enabled"
)

// PossibleTriggerStatusValues returns an array of possible values for the TriggerStatus const type.
func PossibleTriggerStatusValues() []TriggerStatus {
	return []TriggerStatus{TriggerStatusDisabled, TriggerStatusEnabled}
}

// TrustPolicyType enumerates the values for trust policy type.
type TrustPolicyType string

const (
	// Notary ...
	Notary TrustPolicyType = "Notary"
)

// PossibleTrustPolicyTypeValues returns an array of possible values for the TrustPolicyType const type.
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return []TrustPolicyType{Notary}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeDockerBuildRequest ...
	TypeDockerBuildRequest Type = "DockerBuildRequest"
	// TypeEncodedTaskRunRequest ...
	TypeEncodedTaskRunRequest Type = "EncodedTaskRunRequest"
	// TypeFileTaskRunRequest ...
	TypeFileTaskRunRequest Type = "FileTaskRunRequest"
	// TypeRunRequest ...
	TypeRunRequest Type = "RunRequest"
	// TypeTaskRunRequest ...
	TypeTaskRunRequest Type = "TaskRunRequest"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeDockerBuildRequest, TypeEncodedTaskRunRequest, TypeFileTaskRunRequest, TypeRunRequest, TypeTaskRunRequest}
}

// TypeBasicTaskStepProperties enumerates the values for type basic task step properties.
type TypeBasicTaskStepProperties string

const (
	// TypeDocker ...
	TypeDocker TypeBasicTaskStepProperties = "Docker"
	// TypeEncodedTask ...
	TypeEncodedTask TypeBasicTaskStepProperties = "EncodedTask"
	// TypeFileTask ...
	TypeFileTask TypeBasicTaskStepProperties = "FileTask"
	// TypeTaskStepProperties ...
	TypeTaskStepProperties TypeBasicTaskStepProperties = "TaskStepProperties"
)

// PossibleTypeBasicTaskStepPropertiesValues returns an array of possible values for the TypeBasicTaskStepProperties const type.
func PossibleTypeBasicTaskStepPropertiesValues() []TypeBasicTaskStepProperties {
	return []TypeBasicTaskStepProperties{TypeDocker, TypeEncodedTask, TypeFileTask, TypeTaskStepProperties}
}

// TypeBasicTaskStepUpdateParameters enumerates the values for type basic task step update parameters.
type TypeBasicTaskStepUpdateParameters string

const (
	// TypeBasicTaskStepUpdateParametersTypeDocker ...
	TypeBasicTaskStepUpdateParametersTypeDocker TypeBasicTaskStepUpdateParameters = "Docker"
	// TypeBasicTaskStepUpdateParametersTypeEncodedTask ...
	TypeBasicTaskStepUpdateParametersTypeEncodedTask TypeBasicTaskStepUpdateParameters = "EncodedTask"
	// TypeBasicTaskStepUpdateParametersTypeFileTask ...
	TypeBasicTaskStepUpdateParametersTypeFileTask TypeBasicTaskStepUpdateParameters = "FileTask"
	// TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters ...
	TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters TypeBasicTaskStepUpdateParameters = "TaskStepUpdateParameters"
)

// PossibleTypeBasicTaskStepUpdateParametersValues returns an array of possible values for the TypeBasicTaskStepUpdateParameters const type.
func PossibleTypeBasicTaskStepUpdateParametersValues() []TypeBasicTaskStepUpdateParameters {
	return []TypeBasicTaskStepUpdateParameters{TypeBasicTaskStepUpdateParametersTypeDocker, TypeBasicTaskStepUpdateParametersTypeEncodedTask, TypeBasicTaskStepUpdateParametersTypeFileTask, TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters}
}

// UpdateTriggerPayloadType enumerates the values for update trigger payload type.
type UpdateTriggerPayloadType string

const (
	// UpdateTriggerPayloadTypeDefault ...
	UpdateTriggerPayloadTypeDefault UpdateTriggerPayloadType = "Default"
	// UpdateTriggerPayloadTypeToken ...
	UpdateTriggerPayloadTypeToken UpdateTriggerPayloadType = "Token"
)

// PossibleUpdateTriggerPayloadTypeValues returns an array of possible values for the UpdateTriggerPayloadType const type.
func PossibleUpdateTriggerPayloadTypeValues() []UpdateTriggerPayloadType {
	return []UpdateTriggerPayloadType{UpdateTriggerPayloadTypeDefault, UpdateTriggerPayloadTypeToken}
}

// Variant enumerates the values for variant.
type Variant string

const (
	// V6 ...
	V6 Variant = "v6"
	// V7 ...
	V7 Variant = "v7"
	// V8 ...
	V8 Variant = "v8"
)

// PossibleVariantValues returns an array of possible values for the Variant const type.
func PossibleVariantValues() []Variant {
	return []Variant{V6, V7, V8}
}

// WebhookAction enumerates the values for webhook action.
type WebhookAction string

const (
	// ChartDelete ...
	ChartDelete WebhookAction = "chart_delete"
	// ChartPush ...
	ChartPush WebhookAction = "chart_push"
	// Delete ...
	Delete WebhookAction = "delete"
	// Push ...
	Push WebhookAction = "push"
	// Quarantine ...
	Quarantine WebhookAction = "quarantine"
)

// PossibleWebhookActionValues returns an array of possible values for the WebhookAction const type.
func PossibleWebhookActionValues() []WebhookAction {
	return []WebhookAction{ChartDelete, ChartPush, Delete, Push, Quarantine}
}

// WebhookStatus enumerates the values for webhook status.
type WebhookStatus string

const (
	// WebhookStatusDisabled ...
	WebhookStatusDisabled WebhookStatus = "disabled"
	// WebhookStatusEnabled ...
	WebhookStatusEnabled WebhookStatus = "enabled"
)

// PossibleWebhookStatusValues returns an array of possible values for the WebhookStatus const type.
func PossibleWebhookStatusValues() []WebhookStatus {
	return []WebhookStatus{WebhookStatusDisabled, WebhookStatusEnabled}
}
