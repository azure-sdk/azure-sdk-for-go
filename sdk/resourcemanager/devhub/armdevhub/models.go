//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevhub

import "time"

// ACR - Information on the azure container registry
type ACR struct {
	// ACR registry
	AcrRegistryName *string `json:"acrRegistryName,omitempty"`

	// ACR repository
	AcrRepositoryName *string `json:"acrRepositoryName,omitempty"`

	// ACR resource group
	AcrResourceGroup *string `json:"acrResourceGroup,omitempty"`

	// ACR subscription id
	AcrSubscriptionID *string `json:"acrSubscriptionId,omitempty"`
}

// ArtifactGenerationProperties - Properties used for generating artifacts such as Dockerfiles and manifests.
type ArtifactGenerationProperties struct {
	// The name of the app.
	AppName *string `json:"appName,omitempty"`

	// The version of the language image used for building the code in the generated dockerfile.
	BuilderVersion *string `json:"builderVersion,omitempty"`

	// The mode of generation to be used for generating Dockerfiles.
	DockerfileGenerationMode *DockerfileGenerationMode `json:"dockerfileGenerationMode,omitempty"`

	// The directory to output the generated Dockerfile to.
	DockerfileOutputDirectory *string `json:"dockerfileOutputDirectory,omitempty"`

	// The programming language used.
	GenerationLanguage *GenerationLanguage `json:"generationLanguage,omitempty"`

	// The name of the image to be generated.
	ImageName *string `json:"imageName,omitempty"`

	// The tag to apply to the generated image.
	ImageTag *string `json:"imageTag,omitempty"`

	// The version of the language image used for execution in the generated dockerfile.
	LanguageVersion *string `json:"languageVersion,omitempty"`

	// The mode of generation to be used for generating Manifest.
	ManifestGenerationMode *ManifestGenerationMode `json:"manifestGenerationMode,omitempty"`

	// The directory to output the generated manifests to.
	ManifestOutputDirectory *string `json:"manifestOutputDirectory,omitempty"`

	// Determines the type of manifests to be generated.
	ManifestType *GenerationManifestType `json:"manifestType,omitempty"`

	// The namespace to deploy the application to.
	Namespace *string `json:"namespace,omitempty"`

	// The port the application is exposed on.
	Port *string `json:"port,omitempty"`
}

// DeleteWorkflowResponse - delete response if content must be provided on delete operation
type DeleteWorkflowResponse struct {
	// delete status message
	Status *string `json:"status,omitempty"`
}

type DeploymentProperties struct {
	// Helm chart directory path in repository.
	HelmChartPath *string `json:"helmChartPath,omitempty"`

	// Helm Values.yaml file location in repository.
	HelmValues            *string   `json:"helmValues,omitempty"`
	KubeManifestLocations []*string `json:"kubeManifestLocations,omitempty"`

	// Determines the type of manifests within the repository.
	ManifestType *ManifestType `json:"manifestType,omitempty"`

	// Manifest override values.
	Overrides map[string]*string `json:"overrides,omitempty"`
}

// DeveloperHubServiceClientGeneratePreviewArtifactsOptions contains the optional parameters for the DeveloperHubServiceClient.GeneratePreviewArtifacts
// method.
type DeveloperHubServiceClientGeneratePreviewArtifactsOptions struct {
	Parameters *ArtifactGenerationProperties
}

// DeveloperHubServiceClientGitHubOAuthCallbackOptions contains the optional parameters for the DeveloperHubServiceClient.GitHubOAuthCallback
// method.
type DeveloperHubServiceClientGitHubOAuthCallbackOptions struct {
	// placeholder for future optional parameters
}

// DeveloperHubServiceClientGitHubOAuthOptions contains the optional parameters for the DeveloperHubServiceClient.GitHubOAuth
// method.
type DeveloperHubServiceClientGitHubOAuthOptions struct {
	Parameters *GitHubOAuthCallRequest
}

// DeveloperHubServiceClientListGitHubOAuthOptions contains the optional parameters for the DeveloperHubServiceClient.ListGitHubOAuth
// method.
type DeveloperHubServiceClientListGitHubOAuthOptions struct {
	// placeholder for future optional parameters
}

// GeneratePreviewArtifactsResponse - Dockerfile and manifest artifacts generated as a preview are returned as a map
type GeneratePreviewArtifactsResponse struct {
	// An example generated file
	PathToFile1WithExtension *string `json:"path/to/file1.withExtension,omitempty"`

	// A second example file
	PathToFile2WithExtension *string `json:"path/to/file2.withExtension,omitempty"`
}

// GitHubOAuthCallRequest - GitHubOAuth request object
type GitHubOAuthCallRequest struct {
	// The URL the client will redirect to on successful authentication. If empty, no redirect will occur.
	RedirectURL *string `json:"redirectUrl,omitempty"`
}

// GitHubOAuthInfoResponse - URL used to authorize the Developer Hub GitHub App
type GitHubOAuthInfoResponse struct {
	// URL for authorizing the Developer Hub GitHub App
	AuthURL *string `json:"authURL,omitempty"`

	// OAuth token used to make calls to GitHub
	Token *string `json:"token,omitempty"`
}

// GitHubOAuthListResponse - The response from List GitHubOAuth operation.
type GitHubOAuthListResponse struct {
	// Singleton list response containing one GitHubOAuthResponse response
	Value []*GitHubOAuthResponse `json:"value,omitempty"`
}

// GitHubOAuthProperties - The response from List GitHubOAuth operation.
type GitHubOAuthProperties struct {
	// user making request
	Username *string `json:"username,omitempty"`
}

// GitHubOAuthResponse - Singleton response of GitHubOAuth containing
type GitHubOAuthResponse struct {
	// Properties of a workflow.
	Properties *GitHubOAuthProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// GitHubWorkflowProfile - GitHub Workflow Profile
type GitHubWorkflowProfile struct {
	// Information on the azure container registry
	Acr *ACR `json:"acr,omitempty"`

	// The Azure Kubernetes Cluster Resource the application will be deployed to.
	AksResourceID *string `json:"aksResourceId,omitempty"`

	// Repository Branch Name
	BranchName           *string               `json:"branchName,omitempty"`
	DeploymentProperties *DeploymentProperties `json:"deploymentProperties,omitempty"`

	// Path to Dockerfile Build Context within the repository.
	DockerBuildContext *string `json:"dockerBuildContext,omitempty"`

	// Path to the Dockerfile within the repository.
	Dockerfile      *string      `json:"dockerfile,omitempty"`
	LastWorkflowRun *WorkflowRun `json:"lastWorkflowRun,omitempty"`

	// Kubernetes namespace the application is deployed to.
	Namespace *string `json:"namespace,omitempty"`

	// The fields needed for OIDC with GitHub.
	OidcCredentials *GitHubWorkflowProfileOidcCredentials `json:"oidcCredentials,omitempty"`

	// Repository Name
	RepositoryName *string `json:"repositoryName,omitempty"`

	// Repository Owner
	RepositoryOwner *string `json:"repositoryOwner,omitempty"`

	// READ-ONLY; Determines the authorization status of requests.
	AuthStatus *AuthorizationStatus `json:"authStatus,omitempty" azure:"ro"`

	// READ-ONLY; The status of the Pull Request submitted against the users repository.
	PrStatus *PullRequestStatus `json:"prStatus,omitempty" azure:"ro"`

	// READ-ONLY; The URL to the Pull Request submitted against the users repository.
	PrURL *string `json:"prURL,omitempty" azure:"ro"`

	// READ-ONLY; The number associated with the submitted pull request.
	PullNumber *int32 `json:"pullNumber,omitempty" azure:"ro"`
}

// GitHubWorkflowProfileOidcCredentials - The fields needed for OIDC with GitHub.
type GitHubWorkflowProfileOidcCredentials struct {
	// Azure Application Client ID
	AzureClientID *string `json:"azureClientId,omitempty"`

	// Azure Directory (tenant) ID
	AzureTenantID *string `json:"azureTenantId,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TagsObject - Resource tags.
type TagsObject struct {
	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`
}

// Workflow - Resource representation of a workflow
type Workflow struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Properties of a workflow.
	Properties *WorkflowProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// WorkflowClientCreateOrUpdateOptions contains the optional parameters for the WorkflowClient.CreateOrUpdate method.
type WorkflowClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WorkflowClientDeleteOptions contains the optional parameters for the WorkflowClient.Delete method.
type WorkflowClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkflowClientGetOptions contains the optional parameters for the WorkflowClient.Get method.
type WorkflowClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkflowClientListByResourceGroupOptions contains the optional parameters for the WorkflowClient.NewListByResourceGroupPager
// method.
type WorkflowClientListByResourceGroupOptions struct {
}

// WorkflowClientListOptions contains the optional parameters for the WorkflowClient.NewListPager method.
type WorkflowClientListOptions struct {
	// placeholder for future optional parameters
}

// WorkflowClientUpdateTagsOptions contains the optional parameters for the WorkflowClient.UpdateTags method.
type WorkflowClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// WorkflowListResult - The response from List Workflows operation.
type WorkflowListResult struct {
	// The list of workflows.
	Value []*Workflow `json:"value,omitempty"`

	// READ-ONLY; The URL to the next set of workflow results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// WorkflowProperties - Workflow properties
type WorkflowProperties struct {
	// Properties used for generating artifacts such as Dockerfiles and manifests.
	ArtifactGenerationProperties *ArtifactGenerationProperties `json:"artifactGenerationProperties,omitempty"`

	// Profile of a github workflow.
	GithubWorkflowProfile *GitHubWorkflowProfile `json:"githubWorkflowProfile,omitempty"`
}

type WorkflowRun struct {
	// READ-ONLY; The timestamp of the last workflow run.
	LastRunAt *time.Time `json:"lastRunAt,omitempty" azure:"ro"`

	// READ-ONLY; Describes if the workflow run succeeded.
	Succeeded *bool `json:"succeeded,omitempty" azure:"ro"`

	// READ-ONLY; URL to the run of the workflow.
	WorkflowRunURL *string `json:"workflowRunURL,omitempty" azure:"ro"`
}
