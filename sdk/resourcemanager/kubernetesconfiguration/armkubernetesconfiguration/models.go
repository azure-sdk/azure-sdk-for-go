//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkubernetesconfiguration

import "time"

// AzureBlobDefinition - Parameters to reconcile to the AzureBlob source kind type.
type AzureBlobDefinition struct {
	// The account key (shared key) to access the storage account
	AccountKey *string

	// The Azure Blob container name to sync from the url endpoint for the flux configuration.
	ContainerName *string

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// Parameters to authenticate using a Managed Identity.
	ManagedIdentity *ManagedIdentityDefinition

	// The Shared Access token to access the storage container
	SasToken *string

	// Parameters to authenticate using Service Principal.
	ServicePrincipal *ServicePrincipalDefinition

	// The interval at which to re-reconcile the cluster Azure Blob source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster Azure Blob source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration Azure Blob storage account.
	URL *string
}

// AzureBlobPatchDefinition - Parameters to reconcile to the AzureBlob source kind type.
type AzureBlobPatchDefinition struct {
	// The account key (shared key) to access the storage account
	AccountKey *string

	// The Azure Blob container name to sync from the url endpoint for the flux configuration.
	ContainerName *string

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// Parameters to authenticate using a Managed Identity.
	ManagedIdentity *ManagedIdentityPatchDefinition

	// The Shared Access token to access the storage container
	SasToken *string

	// Parameters to authenticate using Service Principal.
	ServicePrincipal *ServicePrincipalPatchDefinition

	// The interval at which to re-reconcile the cluster Azure Blob source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster Azure Blob source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration Azure Blob storage account.
	URL *string
}

// BucketDefinition - Parameters to reconcile to the Bucket source kind type.
type BucketDefinition struct {
	// Plaintext access key used to securely access the S3 bucket
	AccessKey *string

	// The bucket name to sync from the url endpoint for the flux configuration.
	BucketName *string

	// Specify whether to use insecure communication when puling data from the S3 bucket.
	Insecure *bool

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// The interval at which to re-reconcile the cluster bucket source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster bucket source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration S3 bucket.
	URL *string
}

// BucketPatchDefinition - Parameters to reconcile to the Bucket source kind type.
type BucketPatchDefinition struct {
	// Plaintext access key used to securely access the S3 bucket
	AccessKey *string

	// The bucket name to sync from the url endpoint for the flux configuration.
	BucketName *string

	// Specify whether to use insecure communication when puling data from the S3 bucket.
	Insecure *bool

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// The interval at which to re-reconcile the cluster bucket source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster bucket source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration S3 bucket.
	URL *string
}

// ComplianceStatus - Compliance Status details
type ComplianceStatus struct {
	// Datetime the configuration was last applied.
	LastConfigApplied *time.Time

	// Message from when the configuration was applied.
	Message *string

	// Level of the message.
	MessageLevel *MessageLevelType

	// READ-ONLY; The compliance state of the configuration.
	ComplianceState *ComplianceStateType
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Extension - The Extension object.
type Extension struct {
	// Identity of the Extension resource
	Identity *Identity

	// The plan information.
	Plan *Plan

	// Properties of an Extension resource
	Properties *ExtensionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ExtensionProperties - Properties of an Extension resource
type ExtensionProperties struct {
	// Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity *ExtensionPropertiesAksAssignedIdentity

	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool

	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings map[string]*string

	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]*string

	// Type of the Extension, of which this resource is an instance of. It must be one of the Extension Types registered with
	// Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType *string

	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion
	// is 'true'.
	ReleaseTrain *string

	// Scope at which the extension is installed.
	Scope *Scope

	// Status from this extension.
	Statuses []*ExtensionStatus

	// User-specified version of the extension for this extension to 'pin'. To use 'version', autoUpgradeMinorVersion must be
	// 'false'.
	Version *string

	// READ-ONLY; Currently installed version of the extension.
	CurrentVersion *string

	// READ-ONLY; Custom Location settings properties.
	CustomLocationSettings map[string]*string

	// READ-ONLY; Error information from the Agent - e.g. errors during installation.
	ErrorInfo *ErrorDetail

	// READ-ONLY; Flag to note if this extension is a system extension
	IsSystemExtension *bool

	// READ-ONLY; Uri of the Helm package
	PackageURI *string

	// READ-ONLY; Status of installation of this extension.
	ProvisioningState *ProvisioningState
}

// ExtensionPropertiesAksAssignedIdentity - Identity of the Extension resource in an AKS cluster
type ExtensionPropertiesAksAssignedIdentity struct {
	// The identity type.
	Type *AKSIdentityType

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// ExtensionStatus - Status from the extension.
type ExtensionStatus struct {
	// Status code provided by the Extension
	Code *string

	// Short description of status of the extension.
	DisplayStatus *string

	// Level of the status.
	Level *LevelType

	// Detailed message of the status from the Extension.
	Message *string

	// DateLiteral (per ISO8601) noting the time of installation status.
	Time *string
}

// ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
type ExtensionsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
type ExtensionsClientBeginDeleteOptions struct {
	// Delete the extension resource in Azure - not the normal asynchronous delete.
	ForceDelete *bool
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
type ExtensionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
type ExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsClientListOptions contains the optional parameters for the ExtensionsClient.NewListPager method.
type ExtensionsClientListOptions struct {
	// placeholder for future optional parameters
}

// ExtensionsList - Result of the request to list Extensions. It contains a list of Extension objects and a URL link to get
// the next set of results.
type ExtensionsList struct {
	// READ-ONLY; URL to get the next set of extension objects, if any.
	NextLink *string

	// READ-ONLY; List of Extensions within a Kubernetes cluster.
	Value []*Extension
}

// FluxConfigOperationStatusClientGetOptions contains the optional parameters for the FluxConfigOperationStatusClient.Get
// method.
type FluxConfigOperationStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// FluxConfiguration - The Flux Configuration object returned in Get & Put response.
type FluxConfiguration struct {
	// Properties to create a Flux Configuration resource
	Properties *FluxConfigurationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FluxConfigurationPatch - The Flux Configuration Patch Request object.
type FluxConfigurationPatch struct {
	// Updatable properties of an Flux Configuration Patch Request
	Properties *FluxConfigurationPatchProperties
}

// FluxConfigurationPatchProperties - Updatable properties of an Flux Configuration Patch Request
type FluxConfigurationPatchProperties struct {
	// Parameters to reconcile to the AzureBlob source kind type.
	AzureBlob *AzureBlobPatchDefinition

	// Parameters to reconcile to the Bucket source kind type.
	Bucket *BucketPatchDefinition

	// Key-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]*string

	// Parameters to reconcile to the GitRepository source kind type.
	GitRepository *GitRepositoryPatchDefinition

	// Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
	Kustomizations map[string]*KustomizationPatchDefinition

	// Source Kind to pull the configuration data from.
	SourceKind *SourceKindType

	// Whether this configuration should suspend its reconciliation of its kustomizations and sources.
	Suspend *bool
}

// FluxConfigurationProperties - Properties to create a Flux Configuration resource
type FluxConfigurationProperties struct {
	// Parameters to reconcile to the AzureBlob source kind type.
	AzureBlob *AzureBlobDefinition

	// Parameters to reconcile to the Bucket source kind type.
	Bucket *BucketDefinition

	// Key-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]*string

	// Parameters to reconcile to the GitRepository source kind type.
	GitRepository *GitRepositoryDefinition

	// Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
	Kustomizations map[string]*KustomizationDefinition

	// The namespace to which this configuration is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and
	// period only.
	Namespace *string

	// Maximum duration to wait for flux configuration reconciliation. E.g PT1H, PT5M, P1D
	ReconciliationWaitDuration *string

	// Scope at which the operator will be installed.
	Scope *ScopeType

	// Source Kind to pull the configuration data from.
	SourceKind *SourceKindType

	// Whether this configuration should suspend its reconciliation of its kustomizations and sources.
	Suspend *bool

	// Whether flux configuration deployment should wait for cluster to reconcile the kustomizations.
	WaitForReconciliation *bool

	// READ-ONLY; Combined status of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed
	// objects.
	ComplianceState *FluxComplianceState

	// READ-ONLY; Error message returned to the user in the case of provisioning failure.
	ErrorMessage *string

	// READ-ONLY; Status of the creation of the fluxConfiguration.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Public Key associated with this fluxConfiguration (either generated within the cluster or provided by the user).
	RepositoryPublicKey *string

	// READ-ONLY; Branch and/or SHA of the source commit synced with the cluster.
	SourceSyncedCommitID *string

	// READ-ONLY; Datetime the fluxConfiguration synced its source on the cluster.
	SourceUpdatedAt *time.Time

	// READ-ONLY; Datetime the fluxConfiguration synced its status on the cluster with Azure.
	StatusUpdatedAt *time.Time

	// READ-ONLY; Statuses of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects
	// provisioned by the fluxConfiguration.
	Statuses []*ObjectStatusDefinition
}

// FluxConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginCreateOrUpdate
// method.
type FluxConfigurationsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientBeginDeleteOptions contains the optional parameters for the FluxConfigurationsClient.BeginDelete
// method.
type FluxConfigurationsClientBeginDeleteOptions struct {
	// Delete the extension resource in Azure - not the normal asynchronous delete.
	ForceDelete *bool
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientBeginUpdateOptions contains the optional parameters for the FluxConfigurationsClient.BeginUpdate
// method.
type FluxConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FluxConfigurationsClientGetOptions contains the optional parameters for the FluxConfigurationsClient.Get method.
type FluxConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// FluxConfigurationsClientListOptions contains the optional parameters for the FluxConfigurationsClient.NewListPager method.
type FluxConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// FluxConfigurationsList - Result of the request to list Flux Configurations. It contains a list of FluxConfiguration objects
// and a URL link to get the next set of results.
type FluxConfigurationsList struct {
	// READ-ONLY; URL to get the next set of configuration objects, if any.
	NextLink *string

	// READ-ONLY; List of Flux Configurations within a Kubernetes cluster.
	Value []*FluxConfiguration
}

// GitRepositoryDefinition - Parameters to reconcile to the GitRepository source kind type.
type GitRepositoryDefinition struct {
	// Base64-encoded HTTPS certificate authority contents used to access git private git repositories over HTTPS
	HTTPSCACert *string

	// Plaintext HTTPS username used to access private git repositories over HTTPS
	HTTPSUser *string

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// The source reference for the GitRepository object.
	RepositoryRef *RepositoryRefDefinition

	// Base64-encoded known_hosts value containing public SSH keys required to access private git repositories over SSH
	SSHKnownHosts *string

	// The interval at which to re-reconcile the cluster git repository source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster git repository source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration git repository.
	URL *string
}

// GitRepositoryPatchDefinition - Parameters to reconcile to the GitRepository source kind type.
type GitRepositoryPatchDefinition struct {
	// Base64-encoded HTTPS certificate authority contents used to access git private git repositories over HTTPS
	HTTPSCACert *string

	// Plaintext HTTPS username used to access private git repositories over HTTPS
	HTTPSUser *string

	// Name of a local secret on the Kubernetes cluster to use as the authentication secret rather than the managed or user-provided
	// configuration secrets.
	LocalAuthRef *string

	// The source reference for the GitRepository object.
	RepositoryRef *RepositoryRefDefinition

	// Base64-encoded known_hosts value containing public SSH keys required to access private git repositories over SSH
	SSHKnownHosts *string

	// The interval at which to re-reconcile the cluster git repository source with the remote.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the cluster git repository source with the remote.
	TimeoutInSeconds *int64

	// The URL to sync for the flux configuration git repository.
	URL *string
}

// HelmOperatorProperties - Properties for Helm operator.
type HelmOperatorProperties struct {
	// Values override for the operator Helm chart.
	ChartValues *string

	// Version of the operator Helm chart.
	ChartVersion *string
}

// HelmReleasePropertiesDefinition - Properties for HelmRelease objects
type HelmReleasePropertiesDefinition struct {
	// Total number of times that the HelmRelease failed to install or upgrade
	FailureCount *int64

	// The reference to the HelmChart object used as the source to this HelmRelease
	HelmChartRef *ObjectReferenceDefinition

	// Number of times that the HelmRelease failed to install
	InstallFailureCount *int64

	// The revision number of the last released object change
	LastRevisionApplied *int64

	// Number of times that the HelmRelease failed to upgrade
	UpgradeFailureCount *int64
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// KustomizationDefinition - The Kustomization defining how to reconcile the artifact pulled by the source type on the cluster.
type KustomizationDefinition struct {
	// Specifies other Kustomizations that this Kustomization depends on. This Kustomization will not reconcile until all dependencies
	// have completed their reconciliation.
	DependsOn []*string

	// Enable/disable re-creating Kubernetes resources on the cluster when patching fails due to an immutable field change.
	Force *bool

	// The path in the source reference to reconcile on the cluster.
	Path *string

	// Used for variable substitution for this Kustomization after kustomize build.
	PostBuild map[string]*PostBuildDefinition

	// Enable/disable garbage collections of Kubernetes objects created by this Kustomization.
	Prune *bool

	// The interval at which to re-reconcile the Kustomization on the cluster in the event of failure on reconciliation.
	RetryIntervalInSeconds *int64

	// The interval at which to re-reconcile the Kustomization on the cluster.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the Kustomization on the cluster.
	TimeoutInSeconds *int64

	// Enable/disable health check for all Kubernetes objects created by this Kustomization.
	Wait *bool

	// READ-ONLY; Name of the Kustomization, matching the key in the Kustomizations object map.
	Name *string
}

// KustomizationPatchDefinition - The Kustomization defining how to reconcile the artifact pulled by the source type on the
// cluster.
type KustomizationPatchDefinition struct {
	// Specifies other Kustomizations that this Kustomization depends on. This Kustomization will not reconcile until all dependencies
	// have completed their reconciliation.
	DependsOn []*string

	// Enable/disable re-creating Kubernetes resources on the cluster when patching fails due to an immutable field change.
	Force *bool

	// The path in the source reference to reconcile on the cluster.
	Path *string

	// Enable/disable garbage collections of Kubernetes objects created by this Kustomization.
	Prune *bool

	// The interval at which to re-reconcile the Kustomization on the cluster in the event of failure on reconciliation.
	RetryIntervalInSeconds *int64

	// The interval at which to re-reconcile the Kustomization on the cluster.
	SyncIntervalInSeconds *int64

	// The maximum time to attempt to reconcile the Kustomization on the cluster.
	TimeoutInSeconds *int64
}

// ManagedIdentityDefinition - Parameters to authenticate using a Managed Identity.
type ManagedIdentityDefinition struct {
	// The client Id for authenticating a Managed Identity.
	ClientID *string
}

// ManagedIdentityPatchDefinition - Parameters to authenticate using a Managed Identity.
type ManagedIdentityPatchDefinition struct {
	// The client Id for authenticating a Managed Identity.
	ClientID *string
}

// ObjectReferenceDefinition - Object reference to a Kubernetes object on a cluster
type ObjectReferenceDefinition struct {
	// Name of the object
	Name *string

	// Namespace of the object
	Namespace *string
}

// ObjectStatusConditionDefinition - Status condition of Kubernetes object
type ObjectStatusConditionDefinition struct {
	// Last time this status condition has changed
	LastTransitionTime *time.Time

	// A more verbose description of the object status condition
	Message *string

	// Reason for the specified status condition type status
	Reason *string

	// Status of the Kubernetes object condition type
	Status *string

	// Object status condition type for this object
	Type *string
}

// ObjectStatusDefinition - Statuses of objects deployed by the user-specified kustomizations from the git repository.
type ObjectStatusDefinition struct {
	// Object reference to the Kustomization that applied this object
	AppliedBy *ObjectReferenceDefinition

	// Compliance state of the applied object showing whether the applied object has come into a ready state on the cluster.
	ComplianceState *FluxComplianceState

	// Additional properties that are provided from objects of the HelmRelease kind
	HelmReleaseProperties *HelmReleasePropertiesDefinition

	// Kind of the applied object
	Kind *string

	// Name of the applied object
	Name *string

	// Namespace of the applied object
	Namespace *string

	// List of Kubernetes object status conditions present on the cluster
	StatusConditions []*ObjectStatusConditionDefinition
}

// OperationStatusClientGetOptions contains the optional parameters for the OperationStatusClient.Get method.
type OperationStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationStatusClientListOptions contains the optional parameters for the OperationStatusClient.NewListPager method.
type OperationStatusClientListOptions struct {
	// placeholder for future optional parameters
}

// OperationStatusList - The async operations in progress, in the cluster.
type OperationStatusList struct {
	// READ-ONLY; URL to get the next set of Operation Result objects, if any.
	NextLink *string

	// READ-ONLY; List of async operations in progress, in the cluster.
	Value []*OperationStatusResult
}

// OperationStatusResult - The current status of an async operation.
type OperationStatusResult struct {
	// REQUIRED; Operation status.
	Status *string

	// Fully qualified ID for the async operation.
	ID *string

	// Name of the async operation.
	Name *string

	// Additional information, if available.
	Properties map[string]*string

	// READ-ONLY; If present, details of the operation error.
	Error *ErrorDetail
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PatchExtension - The Extension Patch Request object.
type PatchExtension struct {
	// Updatable properties of an Extension Patch Request
	Properties *PatchExtensionProperties
}

// PatchExtensionProperties - Updatable properties of an Extension Patch Request
type PatchExtensionProperties struct {
	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool

	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings map[string]*string

	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]*string

	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion
	// is 'true'.
	ReleaseTrain *string

	// Version of the extension for this extension, if it is 'pinned' to a specific version. autoUpgradeMinorVersion must be 'false'.
	Version *string
}

// Plan for the resource.
type Plan struct {
	// REQUIRED; A user defined name of the 3rd Party Artifact that is being procured.
	Name *string

	// REQUIRED; The 3rd Party artifact that is being procured. E.g. NewRelic. Product maps to the OfferID specified for the artifact
	// at the time of Data Market onboarding.
	Product *string

	// REQUIRED; The publisher of the 3rd Party Artifact that is being bought. E.g. NewRelic
	Publisher *string

	// A publisher provided promotion code as provisioned in Data Market for the said product/artifact.
	PromotionCode *string

	// The version of the desired product/artifact.
	Version *string
}

// PostBuildDefinition - The postBuild definitions defining variable substitutions for this Kustomization after kustomize
// build.
type PostBuildDefinition struct {
	// Key/value pairs holding the variables to be substituted in this Kustomization.
	Substitute map[string]*string

	// Array of ConfigMaps/Secrets from which the variables are substituted for this Kustomization.
	SubstituteFrom []*SubstituteFromDefinition
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RepositoryRefDefinition - The source reference for the GitRepository object.
type RepositoryRefDefinition struct {
	// The git repository branch name to checkout.
	Branch *string

	// The commit SHA to checkout. This value must be combined with the branch name to be valid. This takes precedence over semver.
	Commit *string

	// The semver range used to match against git repository tags. This takes precedence over tag.
	Semver *string

	// The git repository tag name to checkout. This takes precedence over branch.
	Tag *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceProviderOperation - Supported operation of this resource provider.
type ResourceProviderOperation struct {
	// Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay

	// Operation name, in format of {provider}/{resource}/{operation}
	Name *string

	// READ-ONLY; The flag that indicates whether the operation applies to data plane.
	IsDataAction *bool

	// READ-ONLY; Origin of the operation
	Origin *string
}

// ResourceProviderOperationDisplay - Display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Description of this operation.
	Description *string

	// Type of operation: get, read, delete, etc.
	Operation *string

	// Resource provider: Microsoft KubernetesConfiguration.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// ResourceProviderOperationList - Result of the request to list operations.
type ResourceProviderOperationList struct {
	// List of operations supported by this resource provider.
	Value []*ResourceProviderOperation

	// READ-ONLY; URL to the next set of results, if any.
	NextLink *string
}

// Scope of the extension. It can be either Cluster or Namespace; but not both.
type Scope struct {
	// Specifies that the scope of the extension is Cluster
	Cluster *ScopeCluster

	// Specifies that the scope of the extension is Namespace
	Namespace *ScopeNamespace
}

// ScopeCluster - Specifies that the scope of the extension is Cluster
type ScopeCluster struct {
	// Namespace where the extension Release must be placed, for a Cluster scoped extension. If this namespace does not exist,
	// it will be created
	ReleaseNamespace *string
}

// ScopeNamespace - Specifies that the scope of the extension is Namespace
type ScopeNamespace struct {
	// Namespace where the extension will be created for an Namespace scoped extension. If this namespace does not exist, it will
	// be created
	TargetNamespace *string
}

// ServicePrincipalDefinition - Parameters to authenticate using Service Principal.
type ServicePrincipalDefinition struct {
	// Base64-encoded certificate used to authenticate a Service Principal
	ClientCertificate *string

	// The password for the certificate used to authenticate a Service Principal
	ClientCertificatePassword *string

	// Specifies whether to include x5c header in client claims when acquiring a token to enable subject name / issuer based authentication
	// for the Client Certificate
	ClientCertificateSendChain *bool

	// The client Id for authenticating a Service Principal.
	ClientID *string

	// The client secret for authenticating a Service Principal
	ClientSecret *string

	// The tenant Id for authenticating a Service Principal
	TenantID *string
}

// ServicePrincipalPatchDefinition - Parameters to authenticate using Service Principal.
type ServicePrincipalPatchDefinition struct {
	// Base64-encoded certificate used to authenticate a Service Principal
	ClientCertificate *string

	// The password for the certificate used to authenticate a Service Principal
	ClientCertificatePassword *string

	// Specifies whether to include x5c header in client claims when acquiring a token to enable subject name / issuer based authentication
	// for the Client Certificate
	ClientCertificateSendChain *bool

	// The client Id for authenticating a Service Principal.
	ClientID *string

	// The client secret for authenticating a Service Principal
	ClientSecret *string

	// The tenant Id for authenticating a Service Principal
	TenantID *string
}

// SourceControlConfiguration - The SourceControl Configuration object returned in Get & Put response.
type SourceControlConfiguration struct {
	// Properties to create a Source Control Configuration resource
	Properties *SourceControlConfigurationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SourceControlConfigurationList - Result of the request to list Source Control Configurations. It contains a list of SourceControlConfiguration
// objects and a URL link to get the next set of results.
type SourceControlConfigurationList struct {
	// READ-ONLY; URL to get the next set of configuration objects, if any.
	NextLink *string

	// READ-ONLY; List of Source Control Configurations within a Kubernetes cluster.
	Value []*SourceControlConfiguration
}

// SourceControlConfigurationProperties - Properties to create a Source Control Configuration resource
type SourceControlConfigurationProperties struct {
	// Name-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]*string

	// Option to enable Helm Operator for this git configuration.
	EnableHelmOperator *bool

	// Properties for Helm operator.
	HelmOperatorProperties *HelmOperatorProperties

	// Instance name of the operator - identifying the specific configuration.
	OperatorInstanceName *string

	// The namespace to which this operator is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period
	// only.
	OperatorNamespace *string

	// Any Parameters for the Operator instance in string format.
	OperatorParams *string

	// Scope at which the operator will be installed.
	OperatorScope *OperatorScopeType

	// Type of the operator
	OperatorType *OperatorType

	// Url of the SourceControl Repository.
	RepositoryURL *string

	// Base64-encoded known_hosts contents containing public SSH keys required to access private Git instances
	SSHKnownHostsContents *string

	// READ-ONLY; Compliance Status of the Configuration
	ComplianceStatus *ComplianceStatus

	// READ-ONLY; The provisioning state of the resource provider.
	ProvisioningState *ProvisioningStateType

	// READ-ONLY; Public Key associated with this SourceControl configuration (either generated within the cluster or provided
	// by the user).
	RepositoryPublicKey *string
}

// SourceControlConfigurationsClientBeginDeleteOptions contains the optional parameters for the SourceControlConfigurationsClient.BeginDelete
// method.
type SourceControlConfigurationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SourceControlConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the SourceControlConfigurationsClient.CreateOrUpdate
// method.
type SourceControlConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SourceControlConfigurationsClientGetOptions contains the optional parameters for the SourceControlConfigurationsClient.Get
// method.
type SourceControlConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SourceControlConfigurationsClientListOptions contains the optional parameters for the SourceControlConfigurationsClient.NewListPager
// method.
type SourceControlConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SubstituteFromDefinition - Array of ConfigMaps/Secrets from which the variables are substituted for this Kustomization.
type SubstituteFromDefinition struct {
	// Define whether it is ConfigMap or Secret that holds the variables to be used in substitution.
	Kind *string

	// Name of the ConfigMap/Secret that holds the variables to be used in substitution.
	Name *string

	// Set to True to proceed without ConfigMap/Secret, if it is not present.
	Optional *bool
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
