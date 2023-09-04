//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

const (
	moduleName    = "armservicelinker"
	moduleVersion = "v2.0.0-beta.1"
)

type AccessKeyPermissions string

const (
	AccessKeyPermissionsListen AccessKeyPermissions = "Listen"
	AccessKeyPermissionsManage AccessKeyPermissions = "Manage"
	AccessKeyPermissionsRead   AccessKeyPermissions = "Read"
	AccessKeyPermissionsSend   AccessKeyPermissions = "Send"
	AccessKeyPermissionsWrite  AccessKeyPermissions = "Write"
)

// PossibleAccessKeyPermissionsValues returns the possible values for the AccessKeyPermissions const type.
func PossibleAccessKeyPermissionsValues() []AccessKeyPermissions {
	return []AccessKeyPermissions{
		AccessKeyPermissionsListen,
		AccessKeyPermissionsManage,
		AccessKeyPermissionsRead,
		AccessKeyPermissionsSend,
		AccessKeyPermissionsWrite,
	}
}

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeEnable   ActionType = "enable"
	ActionTypeInternal ActionType = "Internal"
	ActionTypeOptOut   ActionType = "optOut"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeEnable,
		ActionTypeInternal,
		ActionTypeOptOut,
	}
}

// AllowType - Whether to allow firewall rules.
type AllowType string

const (
	AllowTypeFalse AllowType = "false"
	AllowTypeTrue  AllowType = "true"
)

// PossibleAllowTypeValues returns the possible values for the AllowType const type.
func PossibleAllowTypeValues() []AllowType {
	return []AllowType{
		AllowTypeFalse,
		AllowTypeTrue,
	}
}

// AuthType - The authentication type.
type AuthType string

const (
	AuthTypeAccessKey                   AuthType = "accessKey"
	AuthTypeSecret                      AuthType = "secret"
	AuthTypeServicePrincipalCertificate AuthType = "servicePrincipalCertificate"
	AuthTypeServicePrincipalSecret      AuthType = "servicePrincipalSecret"
	AuthTypeSystemAssignedIdentity      AuthType = "systemAssignedIdentity"
	AuthTypeUserAccount                 AuthType = "userAccount"
	AuthTypeUserAssignedIdentity        AuthType = "userAssignedIdentity"
)

// PossibleAuthTypeValues returns the possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{
		AuthTypeAccessKey,
		AuthTypeSecret,
		AuthTypeServicePrincipalCertificate,
		AuthTypeServicePrincipalSecret,
		AuthTypeSystemAssignedIdentity,
		AuthTypeUserAccount,
		AuthTypeUserAssignedIdentity,
	}
}

// AzureResourceType - The azure resource type.
type AzureResourceType string

const (
	AzureResourceTypeKeyVault AzureResourceType = "KeyVault"
)

// PossibleAzureResourceTypeValues returns the possible values for the AzureResourceType const type.
func PossibleAzureResourceTypeValues() []AzureResourceType {
	return []AzureResourceType{
		AzureResourceTypeKeyVault,
	}
}

// ClientType - The application client type
type ClientType string

const (
	ClientTypeDapr            ClientType = "dapr"
	ClientTypeDjango          ClientType = "django"
	ClientTypeDotnet          ClientType = "dotnet"
	ClientTypeGo              ClientType = "go"
	ClientTypeJava            ClientType = "java"
	ClientTypeKafkaSpringBoot ClientType = "kafka-springBoot"
	ClientTypeNodejs          ClientType = "nodejs"
	ClientTypeNone            ClientType = "none"
	ClientTypePhp             ClientType = "php"
	ClientTypePython          ClientType = "python"
	ClientTypeRuby            ClientType = "ruby"
	ClientTypeSpringBoot      ClientType = "springBoot"
)

// PossibleClientTypeValues returns the possible values for the ClientType const type.
func PossibleClientTypeValues() []ClientType {
	return []ClientType{
		ClientTypeDapr,
		ClientTypeDjango,
		ClientTypeDotnet,
		ClientTypeGo,
		ClientTypeJava,
		ClientTypeKafkaSpringBoot,
		ClientTypeNodejs,
		ClientTypeNone,
		ClientTypePhp,
		ClientTypePython,
		ClientTypeRuby,
		ClientTypeSpringBoot,
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

// DeleteOrUpdateBehavior - The cleanup behavior to indicate whether clean up operation when resource is deleted or updated
type DeleteOrUpdateBehavior string

const (
	DeleteOrUpdateBehaviorDefault       DeleteOrUpdateBehavior = "Default"
	DeleteOrUpdateBehaviorForcedCleanup DeleteOrUpdateBehavior = "ForcedCleanup"
)

// PossibleDeleteOrUpdateBehaviorValues returns the possible values for the DeleteOrUpdateBehavior const type.
func PossibleDeleteOrUpdateBehaviorValues() []DeleteOrUpdateBehavior {
	return []DeleteOrUpdateBehavior{
		DeleteOrUpdateBehaviorDefault,
		DeleteOrUpdateBehaviorForcedCleanup,
	}
}

// DryrunActionName - The name of action for you dryrun job.
type DryrunActionName string

const (
	DryrunActionNameCreateOrUpdate DryrunActionName = "createOrUpdate"
)

// PossibleDryrunActionNameValues returns the possible values for the DryrunActionName const type.
func PossibleDryrunActionNameValues() []DryrunActionName {
	return []DryrunActionName{
		DryrunActionNameCreateOrUpdate,
	}
}

// DryrunPrerequisiteResultType - The type of dryrun result.
type DryrunPrerequisiteResultType string

const (
	DryrunPrerequisiteResultTypeBasicError         DryrunPrerequisiteResultType = "basicError"
	DryrunPrerequisiteResultTypePermissionsMissing DryrunPrerequisiteResultType = "permissionsMissing"
)

// PossibleDryrunPrerequisiteResultTypeValues returns the possible values for the DryrunPrerequisiteResultType const type.
func PossibleDryrunPrerequisiteResultTypeValues() []DryrunPrerequisiteResultType {
	return []DryrunPrerequisiteResultType{
		DryrunPrerequisiteResultTypeBasicError,
		DryrunPrerequisiteResultTypePermissionsMissing,
	}
}

// DryrunPreviewOperationType - The operation type
type DryrunPreviewOperationType string

const (
	DryrunPreviewOperationTypeConfigAuth       DryrunPreviewOperationType = "configAuth"
	DryrunPreviewOperationTypeConfigConnection DryrunPreviewOperationType = "configConnection"
	DryrunPreviewOperationTypeConfigNetwork    DryrunPreviewOperationType = "configNetwork"
)

// PossibleDryrunPreviewOperationTypeValues returns the possible values for the DryrunPreviewOperationType const type.
func PossibleDryrunPreviewOperationTypeValues() []DryrunPreviewOperationType {
	return []DryrunPreviewOperationType{
		DryrunPreviewOperationTypeConfigAuth,
		DryrunPreviewOperationTypeConfigConnection,
		DryrunPreviewOperationTypeConfigNetwork,
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

// SecretType - The secret type.
type SecretType string

const (
	SecretTypeKeyVaultSecretReference SecretType = "keyVaultSecretReference"
	SecretTypeKeyVaultSecretURI       SecretType = "keyVaultSecretUri"
	SecretTypeRawValue                SecretType = "rawValue"
)

// PossibleSecretTypeValues returns the possible values for the SecretType const type.
func PossibleSecretTypeValues() []SecretType {
	return []SecretType{
		SecretTypeKeyVaultSecretReference,
		SecretTypeKeyVaultSecretURI,
		SecretTypeRawValue,
	}
}

// TargetServiceType - The target service type.
type TargetServiceType string

const (
	TargetServiceTypeAzureResource            TargetServiceType = "AzureResource"
	TargetServiceTypeConfluentBootstrapServer TargetServiceType = "ConfluentBootstrapServer"
	TargetServiceTypeConfluentSchemaRegistry  TargetServiceType = "ConfluentSchemaRegistry"
	TargetServiceTypeSelfHostedServer         TargetServiceType = "SelfHostedServer"
)

// PossibleTargetServiceTypeValues returns the possible values for the TargetServiceType const type.
func PossibleTargetServiceTypeValues() []TargetServiceType {
	return []TargetServiceType{
		TargetServiceTypeAzureResource,
		TargetServiceTypeConfluentBootstrapServer,
		TargetServiceTypeConfluentSchemaRegistry,
		TargetServiceTypeSelfHostedServer,
	}
}

// VNetSolutionType - Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypePrivateLink     VNetSolutionType = "privateLink"
	VNetSolutionTypeServiceEndpoint VNetSolutionType = "serviceEndpoint"
)

// PossibleVNetSolutionTypeValues returns the possible values for the VNetSolutionType const type.
func PossibleVNetSolutionTypeValues() []VNetSolutionType {
	return []VNetSolutionType{
		VNetSolutionTypePrivateLink,
		VNetSolutionTypeServiceEndpoint,
	}
}

// ValidationResultStatus - The result of validation
type ValidationResultStatus string

const (
	ValidationResultStatusFailure ValidationResultStatus = "failure"
	ValidationResultStatusSuccess ValidationResultStatus = "success"
	ValidationResultStatusWarning ValidationResultStatus = "warning"
)

// PossibleValidationResultStatusValues returns the possible values for the ValidationResultStatus const type.
func PossibleValidationResultStatusValues() []ValidationResultStatus {
	return []ValidationResultStatus{
		ValidationResultStatusFailure,
		ValidationResultStatusSuccess,
		ValidationResultStatusWarning,
	}
}
