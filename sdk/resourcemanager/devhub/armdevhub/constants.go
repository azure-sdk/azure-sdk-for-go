// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevhub

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
	moduleVersion = "v0.6.0"
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

// AuthorizationStatus - Determines the authorization status of requests.
type AuthorizationStatus string

const (
	// AuthorizationStatusAuthorized - Requests authorized successfully
	AuthorizationStatusAuthorized AuthorizationStatus = "Authorized"
	// AuthorizationStatusError - Requests returned other error response
	AuthorizationStatusError AuthorizationStatus = "Error"
	// AuthorizationStatusNotFound - Requests returned NotFound response
	AuthorizationStatusNotFound AuthorizationStatus = "NotFound"
)

// PossibleAuthorizationStatusValues returns the possible values for the AuthorizationStatus const type.
func PossibleAuthorizationStatusValues() []AuthorizationStatus {
	return []AuthorizationStatus{
		AuthorizationStatusAuthorized,
		AuthorizationStatusError,
		AuthorizationStatusNotFound,
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

// DockerfileGenerationMode - The mode of generation to be used for generating Dockerfiles.
type DockerfileGenerationMode string

const (
	// DockerfileGenerationModeDisabled - Dockerfiles will not be generated
	DockerfileGenerationModeDisabled DockerfileGenerationMode = "disabled"
	// DockerfileGenerationModeEnabled - Dockerfiles will be generated
	DockerfileGenerationModeEnabled DockerfileGenerationMode = "enabled"
)

// PossibleDockerfileGenerationModeValues returns the possible values for the DockerfileGenerationMode const type.
func PossibleDockerfileGenerationModeValues() []DockerfileGenerationMode {
	return []DockerfileGenerationMode{
		DockerfileGenerationModeDisabled,
		DockerfileGenerationModeEnabled,
	}
}

// GenerationLanguage - The programming language used.
type GenerationLanguage string

const (
	// GenerationLanguageClojure - clojure language
	GenerationLanguageClojure GenerationLanguage = "clojure"
	// GenerationLanguageCsharp - csharp language
	GenerationLanguageCsharp GenerationLanguage = "csharp"
	// GenerationLanguageErlang - erlang language
	GenerationLanguageErlang GenerationLanguage = "erlang"
	// GenerationLanguageGo - go language
	GenerationLanguageGo GenerationLanguage = "go"
	// GenerationLanguageGomodule - gomodule language
	GenerationLanguageGomodule GenerationLanguage = "gomodule"
	// GenerationLanguageGradle - gradle language
	GenerationLanguageGradle GenerationLanguage = "gradle"
	// GenerationLanguageJava - java language
	GenerationLanguageJava GenerationLanguage = "java"
	// GenerationLanguageJavascript - javascript language
	GenerationLanguageJavascript GenerationLanguage = "javascript"
	// GenerationLanguagePhp - php language
	GenerationLanguagePhp GenerationLanguage = "php"
	// GenerationLanguagePython - python language
	GenerationLanguagePython GenerationLanguage = "python"
	// GenerationLanguageRuby - ruby language
	GenerationLanguageRuby GenerationLanguage = "ruby"
	// GenerationLanguageRust - rust language
	GenerationLanguageRust GenerationLanguage = "rust"
	// GenerationLanguageSwift - swift language
	GenerationLanguageSwift GenerationLanguage = "swift"
)

// PossibleGenerationLanguageValues returns the possible values for the GenerationLanguage const type.
func PossibleGenerationLanguageValues() []GenerationLanguage {
	return []GenerationLanguage{
		GenerationLanguageClojure,
		GenerationLanguageCsharp,
		GenerationLanguageErlang,
		GenerationLanguageGo,
		GenerationLanguageGomodule,
		GenerationLanguageGradle,
		GenerationLanguageJava,
		GenerationLanguageJavascript,
		GenerationLanguagePhp,
		GenerationLanguagePython,
		GenerationLanguageRuby,
		GenerationLanguageRust,
		GenerationLanguageSwift,
	}
}

// GenerationManifestType - Determines the type of manifests to be generated.
type GenerationManifestType string

const (
	// GenerationManifestTypeHelm - Helm manifests
	GenerationManifestTypeHelm GenerationManifestType = "helm"
	// GenerationManifestTypeKube - Kubernetes manifests
	GenerationManifestTypeKube GenerationManifestType = "kube"
)

// PossibleGenerationManifestTypeValues returns the possible values for the GenerationManifestType const type.
func PossibleGenerationManifestTypeValues() []GenerationManifestType {
	return []GenerationManifestType{
		GenerationManifestTypeHelm,
		GenerationManifestTypeKube,
	}
}

// ManifestGenerationMode - The mode of generation to be used for generating Manifest.
type ManifestGenerationMode string

const (
	// ManifestGenerationModeDisabled - Manifests will not be generated
	ManifestGenerationModeDisabled ManifestGenerationMode = "disabled"
	// ManifestGenerationModeEnabled - Manifests will be generated
	ManifestGenerationModeEnabled ManifestGenerationMode = "enabled"
)

// PossibleManifestGenerationModeValues returns the possible values for the ManifestGenerationMode const type.
func PossibleManifestGenerationModeValues() []ManifestGenerationMode {
	return []ManifestGenerationMode{
		ManifestGenerationModeDisabled,
		ManifestGenerationModeEnabled,
	}
}

// ManifestType - Determines the type of manifests within the repository.
type ManifestType string

const (
	// ManifestTypeHelm - Repositories using helm
	ManifestTypeHelm ManifestType = "helm"
	// ManifestTypeKube - Repositories using kubernetes manifests
	ManifestTypeKube ManifestType = "kube"
)

// PossibleManifestTypeValues returns the possible values for the ManifestType const type.
func PossibleManifestTypeValues() []ManifestType {
	return []ManifestType{
		ManifestTypeHelm,
		ManifestTypeKube,
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

// ParameterKind - The type of the template parameter.
type ParameterKind string

const (
	// ParameterKindAzureContainerRegistry - azure container registry name
	ParameterKindAzureContainerRegistry ParameterKind = "azureContainerRegistry"
	// ParameterKindAzureKeyvaultURI - azure keyvault uri
	ParameterKindAzureKeyvaultURI ParameterKind = "azureKeyvaultUri"
	// ParameterKindAzureManagedCluster - azure managed cluster name
	ParameterKindAzureManagedCluster ParameterKind = "azureManagedCluster"
	// ParameterKindAzureResourceGroup - azure resource group
	ParameterKindAzureResourceGroup ParameterKind = "azureResourceGroup"
	// ParameterKindAzureServiceConnection - azure service connection
	ParameterKindAzureServiceConnection ParameterKind = "azureServiceConnection"
	// ParameterKindClusterResourceType - cluster resource type
	ParameterKindClusterResourceType ParameterKind = "clusterResourceType"
	// ParameterKindContainerImageName - container image name
	ParameterKindContainerImageName ParameterKind = "containerImageName"
	// ParameterKindContainerImageVersion - container image version
	ParameterKindContainerImageVersion ParameterKind = "containerImageVersion"
	// ParameterKindDirPath - directory path
	ParameterKindDirPath ParameterKind = "dirPath"
	// ParameterKindDockerFileName - dockerfile name
	ParameterKindDockerFileName ParameterKind = "dockerFileName"
	// ParameterKindEnvVarMap - environment variables in the form of a json object
	ParameterKindEnvVarMap ParameterKind = "envVarMap"
	// ParameterKindFilePath - file path
	ParameterKindFilePath ParameterKind = "filePath"
	// ParameterKindFlag - boolean flag
	ParameterKindFlag ParameterKind = "flag"
	// ParameterKindHelmChartOverrides - helm overrides in the form of a string key1=value1,key2=value2
	ParameterKindHelmChartOverrides ParameterKind = "helmChartOverrides"
	// ParameterKindImagePullPolicy - kubernetes deployment image pull policy
	ParameterKindImagePullPolicy ParameterKind = "imagePullPolicy"
	// ParameterKindIngressHostName - kubernetes ingress host name
	ParameterKindIngressHostName ParameterKind = "ingressHostName"
	// ParameterKindKubernetesNamespace - kubernetes namespace
	ParameterKindKubernetesNamespace ParameterKind = "kubernetesNamespace"
	// ParameterKindKubernetesProbeDelay - kubernetes probe delay
	ParameterKindKubernetesProbeDelay ParameterKind = "kubernetesProbeDelay"
	// ParameterKindKubernetesProbeHTTPPath - kubernetes probe http path
	ParameterKindKubernetesProbeHTTPPath ParameterKind = "kubernetesProbeHttpPath"
	// ParameterKindKubernetesProbePeriod - kubernetes probe period
	ParameterKindKubernetesProbePeriod ParameterKind = "kubernetesProbePeriod"
	// ParameterKindKubernetesProbeThreshold - kubernetes probe threshold
	ParameterKindKubernetesProbeThreshold ParameterKind = "kubernetesProbeThreshold"
	// ParameterKindKubernetesProbeTimeout - kubernetes probe timeout
	ParameterKindKubernetesProbeTimeout ParameterKind = "kubernetesProbeTimeout"
	// ParameterKindKubernetesProbeType - kubernetes probe type
	ParameterKindKubernetesProbeType ParameterKind = "kubernetesProbeType"
	// ParameterKindKubernetesResourceLimit - kubernetes resource limit
	ParameterKindKubernetesResourceLimit ParameterKind = "kubernetesResourceLimit"
	// ParameterKindKubernetesResourceName - kubernetes resource name
	ParameterKindKubernetesResourceName ParameterKind = "kubernetesResourceName"
	// ParameterKindKubernetesResourceRequest - kubernetes resource request
	ParameterKindKubernetesResourceRequest ParameterKind = "kubernetesResourceRequest"
	// ParameterKindLabel - kubernetes label value
	ParameterKindLabel ParameterKind = "label"
	// ParameterKindPort - service port
	ParameterKindPort ParameterKind = "port"
	// ParameterKindReplicaCount - kubernetes replica count
	ParameterKindReplicaCount ParameterKind = "replicaCount"
	// ParameterKindRepositoryBranch - repository branch name
	ParameterKindRepositoryBranch ParameterKind = "repositoryBranch"
	// ParameterKindResourceLimit - kubernetes resource limit
	ParameterKindResourceLimit ParameterKind = "resourceLimit"
	// ParameterKindScalingResourceType - kubernetes scaling resource type
	ParameterKindScalingResourceType ParameterKind = "scalingResourceType"
	// ParameterKindScalingResourceUtilization - kubernetes resource utilization type
	ParameterKindScalingResourceUtilization ParameterKind = "scalingResourceUtilization"
	// ParameterKindWorkflowAuthType - workflow authentication type
	ParameterKindWorkflowAuthType ParameterKind = "workflowAuthType"
	// ParameterKindWorkflowName - workflow name
	ParameterKindWorkflowName ParameterKind = "workflowName"
)

// PossibleParameterKindValues returns the possible values for the ParameterKind const type.
func PossibleParameterKindValues() []ParameterKind {
	return []ParameterKind{
		ParameterKindAzureContainerRegistry,
		ParameterKindAzureKeyvaultURI,
		ParameterKindAzureManagedCluster,
		ParameterKindAzureResourceGroup,
		ParameterKindAzureServiceConnection,
		ParameterKindClusterResourceType,
		ParameterKindContainerImageName,
		ParameterKindContainerImageVersion,
		ParameterKindDirPath,
		ParameterKindDockerFileName,
		ParameterKindEnvVarMap,
		ParameterKindFilePath,
		ParameterKindFlag,
		ParameterKindHelmChartOverrides,
		ParameterKindImagePullPolicy,
		ParameterKindIngressHostName,
		ParameterKindKubernetesNamespace,
		ParameterKindKubernetesProbeDelay,
		ParameterKindKubernetesProbeHTTPPath,
		ParameterKindKubernetesProbePeriod,
		ParameterKindKubernetesProbeThreshold,
		ParameterKindKubernetesProbeTimeout,
		ParameterKindKubernetesProbeType,
		ParameterKindKubernetesResourceLimit,
		ParameterKindKubernetesResourceName,
		ParameterKindKubernetesResourceRequest,
		ParameterKindLabel,
		ParameterKindPort,
		ParameterKindReplicaCount,
		ParameterKindRepositoryBranch,
		ParameterKindResourceLimit,
		ParameterKindScalingResourceType,
		ParameterKindScalingResourceUtilization,
		ParameterKindWorkflowAuthType,
		ParameterKindWorkflowName,
	}
}

// ParameterType - The type of the template parameter.
type ParameterType string

const (
	// ParameterTypeBool - boolean parameter type.
	ParameterTypeBool ParameterType = "bool"
	// ParameterTypeFloat - float parameter type.
	ParameterTypeFloat ParameterType = "float"
	// ParameterTypeInt - int parameter type.
	ParameterTypeInt ParameterType = "int"
	// ParameterTypeObject - object parameter type.
	ParameterTypeObject ParameterType = "object"
	// ParameterTypeString - string parameter type.
	ParameterTypeString ParameterType = "string"
)

// PossibleParameterTypeValues returns the possible values for the ParameterType const type.
func PossibleParameterTypeValues() []ParameterType {
	return []ParameterType{
		ParameterTypeBool,
		ParameterTypeFloat,
		ParameterTypeInt,
		ParameterTypeObject,
		ParameterTypeString,
	}
}

// PullRequestStatus - The status of the Pull Request submitted against the users repository.
type PullRequestStatus string

const (
	// PullRequestStatusMerged - Pull Request merged into repository.
	PullRequestStatusMerged PullRequestStatus = "merged"
	// PullRequestStatusRemoved - Workflow no longer found within repository.
	PullRequestStatusRemoved PullRequestStatus = "removed"
	// PullRequestStatusSubmitted - Pull Request submitted to repository.
	PullRequestStatusSubmitted PullRequestStatus = "submitted"
	// PullRequestStatusUnknown - Pull Request state unknown.
	PullRequestStatusUnknown PullRequestStatus = "unknown"
)

// PossiblePullRequestStatusValues returns the possible values for the PullRequestStatus const type.
func PossiblePullRequestStatusValues() []PullRequestStatus {
	return []PullRequestStatus{
		PullRequestStatusMerged,
		PullRequestStatusRemoved,
		PullRequestStatusSubmitted,
		PullRequestStatusUnknown,
	}
}

// QuickStartTemplateType - Determines the authorization status of requests.
type QuickStartTemplateType string

const (
	// QuickStartTemplateTypeHCI - The template use quick start template of HCI
	QuickStartTemplateTypeHCI QuickStartTemplateType = "HCI"
	// QuickStartTemplateTypeHCIAKS - The template use quick start template of HCI and AKS
	QuickStartTemplateTypeHCIAKS QuickStartTemplateType = "HCIAKS"
	// QuickStartTemplateTypeHCIARCVM - The template use quick start template of HCI and ArcVM
	QuickStartTemplateTypeHCIARCVM QuickStartTemplateType = "HCIARCVM"
	// QuickStartTemplateTypeNone - The template has not use quick start template
	QuickStartTemplateTypeNone QuickStartTemplateType = "None"
)

// PossibleQuickStartTemplateTypeValues returns the possible values for the QuickStartTemplateType const type.
func PossibleQuickStartTemplateTypeValues() []QuickStartTemplateType {
	return []QuickStartTemplateType{
		QuickStartTemplateTypeHCI,
		QuickStartTemplateTypeHCIAKS,
		QuickStartTemplateTypeHCIARCVM,
		QuickStartTemplateTypeNone,
	}
}

// RepositoryProviderType - The status of the Pull Request submitted against the users repository.
type RepositoryProviderType string

const (
	// RepositoryProviderTypeAdo - ADO repository provider type.
	RepositoryProviderTypeAdo RepositoryProviderType = "ado"
	// RepositoryProviderTypeGithub - GitHub repository provider type.
	RepositoryProviderTypeGithub RepositoryProviderType = "github"
)

// PossibleRepositoryProviderTypeValues returns the possible values for the RepositoryProviderType const type.
func PossibleRepositoryProviderTypeValues() []RepositoryProviderType {
	return []RepositoryProviderType{
		RepositoryProviderTypeAdo,
		RepositoryProviderTypeGithub,
	}
}

// TemplateType - The type of the template.
type TemplateType string

const (
	// TemplateTypeDeployment - Deployment template type.
	TemplateTypeDeployment TemplateType = "deployment"
	// TemplateTypeDockerfile - Dockerfile template type.
	TemplateTypeDockerfile TemplateType = "dockerfile"
	// TemplateTypeManifest - Manifest template type.
	TemplateTypeManifest TemplateType = "manifest"
	// TemplateTypeWorkflow - Workflow template type.
	TemplateTypeWorkflow TemplateType = "workflow"
)

// PossibleTemplateTypeValues returns the possible values for the TemplateType const type.
func PossibleTemplateTypeValues() []TemplateType {
	return []TemplateType{
		TemplateTypeDeployment,
		TemplateTypeDockerfile,
		TemplateTypeManifest,
		TemplateTypeWorkflow,
	}
}

// WorkflowRunStatus - Describes the status of the workflow run
type WorkflowRunStatus string

const (
	// WorkflowRunStatusCompleted - Workflow run is completed
	WorkflowRunStatusCompleted WorkflowRunStatus = "completed"
	// WorkflowRunStatusInprogress - Workflow run is inprogress
	WorkflowRunStatusInprogress WorkflowRunStatus = "inprogress"
	// WorkflowRunStatusQueued - Workflow run is queued
	WorkflowRunStatusQueued WorkflowRunStatus = "queued"
)

// PossibleWorkflowRunStatusValues returns the possible values for the WorkflowRunStatus const type.
func PossibleWorkflowRunStatusValues() []WorkflowRunStatus {
	return []WorkflowRunStatus{
		WorkflowRunStatusCompleted,
		WorkflowRunStatusInprogress,
		WorkflowRunStatusQueued,
	}
}
