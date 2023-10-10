//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

// AvailableWorkloadProfilesClientGetResponse contains the response from method AvailableWorkloadProfilesClient.NewGetPager.
type AvailableWorkloadProfilesClientGetResponse struct {
	// Collection of available workload profiles in the location.
	AvailableWorkloadProfilesCollection
}

// BillingMetersClientGetResponse contains the response from method BillingMetersClient.Get.
type BillingMetersClientGetResponse struct {
	// Collection of billing meters.
	BillingMeterCollection
}

// BuildersClientCreateOrUpdateResponse contains the response from method BuildersClient.BeginCreateOrUpdate.
type BuildersClientCreateOrUpdateResponse struct {
	// Information about the SourceToCloud builder resource.
	BuilderResource
}

// BuildersClientDeleteResponse contains the response from method BuildersClient.BeginDelete.
type BuildersClientDeleteResponse struct {
	// placeholder for future response values
}

// BuildersClientGetResponse contains the response from method BuildersClient.Get.
type BuildersClientGetResponse struct {
	// Information about the SourceToCloud builder resource.
	BuilderResource
}

// BuildersClientListByResourceGroupResponse contains the response from method BuildersClient.NewListByResourceGroupPager.
type BuildersClientListByResourceGroupResponse struct {
	// The response of a BuilderResource list operation.
	BuilderCollection
}

// BuildersClientListBySubscriptionResponse contains the response from method BuildersClient.NewListBySubscriptionPager.
type BuildersClientListBySubscriptionResponse struct {
	// The response of a BuilderResource list operation.
	BuilderCollection
}

// BuildersClientUpdateResponse contains the response from method BuildersClient.BeginUpdate.
type BuildersClientUpdateResponse struct {
	// Information about the SourceToCloud builder resource.
	BuilderResource
}

// BuildsClientCreateOrUpdateResponse contains the response from method BuildsClient.BeginCreateOrUpdate.
type BuildsClientCreateOrUpdateResponse struct {
	// Information pertaining to an individual build.
	BuildResource
}

// BuildsClientDeleteResponse contains the response from method BuildsClient.BeginDelete.
type BuildsClientDeleteResponse struct {
	// placeholder for future response values
}

// BuildsClientGetResponse contains the response from method BuildsClient.Get.
type BuildsClientGetResponse struct {
	// Information pertaining to an individual build.
	BuildResource
}

// BuildsClientListAuthTokenResponse contains the response from method BuildsClient.ListAuthToken.
type BuildsClientListAuthTokenResponse struct {
	// Build Auth Token.
	BuildToken
}

// BuildsClientListByBuilderResourceResponse contains the response from method BuildsClient.NewListByBuilderResourcePager.
type BuildsClientListByBuilderResourceResponse struct {
	// The response of a BuildResource list operation.
	BuildCollection
}

// CertificatesClientCreateOrUpdateResponse contains the response from method CertificatesClient.CreateOrUpdate.
type CertificatesClientCreateOrUpdateResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// CertificatesClientDeleteResponse contains the response from method CertificatesClient.Delete.
type CertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificatesClientGetResponse contains the response from method CertificatesClient.Get.
type CertificatesClientGetResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// CertificatesClientListResponse contains the response from method CertificatesClient.NewListPager.
type CertificatesClientListResponse struct {
	// Collection of Certificates.
	CertificateCollection
}

// CertificatesClientUpdateResponse contains the response from method CertificatesClient.Update.
type CertificatesClientUpdateResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsCertificatesClient.CreateOrUpdate.
type ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// ConnectedEnvironmentsCertificatesClientDeleteResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Delete.
type ConnectedEnvironmentsCertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsCertificatesClientGetResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Get.
type ConnectedEnvironmentsCertificatesClientGetResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// ConnectedEnvironmentsCertificatesClientListResponse contains the response from method ConnectedEnvironmentsCertificatesClient.NewListPager.
type ConnectedEnvironmentsCertificatesClientListResponse struct {
	// Collection of Certificates.
	CertificateCollection
}

// ConnectedEnvironmentsCertificatesClientUpdateResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Update.
type ConnectedEnvironmentsCertificatesClientUpdateResponse struct {
	// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
	Certificate
}

// ConnectedEnvironmentsClientCheckNameAvailabilityResponse contains the response from method ConnectedEnvironmentsClient.CheckNameAvailability.
type ConnectedEnvironmentsClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// ConnectedEnvironmentsClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsClient.BeginCreateOrUpdate.
type ConnectedEnvironmentsClientCreateOrUpdateResponse struct {
	// An environment for Kubernetes cluster specialized for web workloads by Azure App Service
	ConnectedEnvironment
}

// ConnectedEnvironmentsClientDeleteResponse contains the response from method ConnectedEnvironmentsClient.BeginDelete.
type ConnectedEnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsClientGetResponse contains the response from method ConnectedEnvironmentsClient.Get.
type ConnectedEnvironmentsClientGetResponse struct {
	// An environment for Kubernetes cluster specialized for web workloads by Azure App Service
	ConnectedEnvironment
}

// ConnectedEnvironmentsClientListByResourceGroupResponse contains the response from method ConnectedEnvironmentsClient.NewListByResourceGroupPager.
type ConnectedEnvironmentsClientListByResourceGroupResponse struct {
	// Collection of connectedEnvironments
	ConnectedEnvironmentCollection
}

// ConnectedEnvironmentsClientListBySubscriptionResponse contains the response from method ConnectedEnvironmentsClient.NewListBySubscriptionPager.
type ConnectedEnvironmentsClientListBySubscriptionResponse struct {
	// Collection of connectedEnvironments
	ConnectedEnvironmentCollection
}

// ConnectedEnvironmentsClientUpdateResponse contains the response from method ConnectedEnvironmentsClient.Update.
type ConnectedEnvironmentsClientUpdateResponse struct {
	// An environment for Kubernetes cluster specialized for web workloads by Azure App Service
	ConnectedEnvironment
}

// ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.CreateOrUpdate.
type ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse struct {
	// Dapr Component.
	DaprComponent
}

// ConnectedEnvironmentsDaprComponentsClientDeleteResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.Delete.
type ConnectedEnvironmentsDaprComponentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsDaprComponentsClientGetResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.Get.
type ConnectedEnvironmentsDaprComponentsClientGetResponse struct {
	// Dapr Component.
	DaprComponent
}

// ConnectedEnvironmentsDaprComponentsClientListResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.NewListPager.
type ConnectedEnvironmentsDaprComponentsClientListResponse struct {
	// Dapr Components ARM resource.
	DaprComponentsCollection
}

// ConnectedEnvironmentsDaprComponentsClientListSecretsResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.ListSecrets.
type ConnectedEnvironmentsDaprComponentsClientListSecretsResponse struct {
	// Dapr component Secrets Collection for ListSecrets Action.
	DaprSecretsCollection
}

// ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsStoragesClient.CreateOrUpdate.
type ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse struct {
	// Storage resource for connectedEnvironment.
	ConnectedEnvironmentStorage
}

// ConnectedEnvironmentsStoragesClientDeleteResponse contains the response from method ConnectedEnvironmentsStoragesClient.Delete.
type ConnectedEnvironmentsStoragesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsStoragesClientGetResponse contains the response from method ConnectedEnvironmentsStoragesClient.Get.
type ConnectedEnvironmentsStoragesClientGetResponse struct {
	// Storage resource for connectedEnvironment.
	ConnectedEnvironmentStorage
}

// ConnectedEnvironmentsStoragesClientListResponse contains the response from method ConnectedEnvironmentsStoragesClient.List.
type ConnectedEnvironmentsStoragesClientListResponse struct {
	// Collection of Storage for Environments
	ConnectedEnvironmentStoragesCollection
}

// ContainerAppsAPIClientGetCustomDomainVerificationIDResponse contains the response from method ContainerAppsAPIClient.GetCustomDomainVerificationID.
type ContainerAppsAPIClientGetCustomDomainVerificationIDResponse struct {
	// Custom domain verification Id of a subscription
	Value *string
}

// ContainerAppsAPIClientJobExecutionResponse contains the response from method ContainerAppsAPIClient.JobExecution.
type ContainerAppsAPIClientJobExecutionResponse struct {
	// Container Apps Job execution.
	JobExecution
}

// ContainerAppsAuthConfigsClientCreateOrUpdateResponse contains the response from method ContainerAppsAuthConfigsClient.CreateOrUpdate.
type ContainerAppsAuthConfigsClientCreateOrUpdateResponse struct {
	// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
	AuthConfig
}

// ContainerAppsAuthConfigsClientDeleteResponse contains the response from method ContainerAppsAuthConfigsClient.Delete.
type ContainerAppsAuthConfigsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsAuthConfigsClientGetResponse contains the response from method ContainerAppsAuthConfigsClient.Get.
type ContainerAppsAuthConfigsClientGetResponse struct {
	// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
	AuthConfig
}

// ContainerAppsAuthConfigsClientListByContainerAppResponse contains the response from method ContainerAppsAuthConfigsClient.NewListByContainerAppPager.
type ContainerAppsAuthConfigsClientListByContainerAppResponse struct {
	// AuthConfig collection ARM resource.
	AuthConfigCollection
}

// ContainerAppsClientCreateOrUpdateResponse contains the response from method ContainerAppsClient.BeginCreateOrUpdate.
type ContainerAppsClientCreateOrUpdateResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsClientDeleteResponse contains the response from method ContainerAppsClient.BeginDelete.
type ContainerAppsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsClientGetAuthTokenResponse contains the response from method ContainerAppsClient.GetAuthToken.
type ContainerAppsClientGetAuthTokenResponse struct {
	// Container App Auth Token.
	ContainerAppAuthToken
}

// ContainerAppsClientGetResponse contains the response from method ContainerAppsClient.Get.
type ContainerAppsClientGetResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsClientListByResourceGroupResponse contains the response from method ContainerAppsClient.NewListByResourceGroupPager.
type ContainerAppsClientListByResourceGroupResponse struct {
	// Container App collection ARM resource.
	ContainerAppCollection
}

// ContainerAppsClientListBySubscriptionResponse contains the response from method ContainerAppsClient.NewListBySubscriptionPager.
type ContainerAppsClientListBySubscriptionResponse struct {
	// Container App collection ARM resource.
	ContainerAppCollection
}

// ContainerAppsClientListCustomHostNameAnalysisResponse contains the response from method ContainerAppsClient.ListCustomHostNameAnalysis.
type ContainerAppsClientListCustomHostNameAnalysisResponse struct {
	// Custom domain analysis.
	CustomHostnameAnalysisResult
}

// ContainerAppsClientListSecretsResponse contains the response from method ContainerAppsClient.ListSecrets.
type ContainerAppsClientListSecretsResponse struct {
	// Container App Secrets Collection ARM resource.
	SecretsCollection
}

// ContainerAppsClientStartResponse contains the response from method ContainerAppsClient.BeginStart.
type ContainerAppsClientStartResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsClientStopResponse contains the response from method ContainerAppsClient.BeginStop.
type ContainerAppsClientStopResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsClientUpdateResponse contains the response from method ContainerAppsClient.BeginUpdate.
type ContainerAppsClientUpdateResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsDiagnosticsClientGetDetectorResponse contains the response from method ContainerAppsDiagnosticsClient.GetDetector.
type ContainerAppsDiagnosticsClientGetDetectorResponse struct {
	// Diagnostics data for a resource.
	Diagnostics
}

// ContainerAppsDiagnosticsClientGetRevisionResponse contains the response from method ContainerAppsDiagnosticsClient.GetRevision.
type ContainerAppsDiagnosticsClientGetRevisionResponse struct {
	// Container App Revision.
	Revision
}

// ContainerAppsDiagnosticsClientGetRootResponse contains the response from method ContainerAppsDiagnosticsClient.GetRoot.
type ContainerAppsDiagnosticsClientGetRootResponse struct {
	// Container App.
	ContainerApp
}

// ContainerAppsDiagnosticsClientListDetectorsResponse contains the response from method ContainerAppsDiagnosticsClient.NewListDetectorsPager.
type ContainerAppsDiagnosticsClientListDetectorsResponse struct {
	// Diagnostics data collection for a resource.
	DiagnosticsCollection
}

// ContainerAppsDiagnosticsClientListRevisionsResponse contains the response from method ContainerAppsDiagnosticsClient.NewListRevisionsPager.
type ContainerAppsDiagnosticsClientListRevisionsResponse struct {
	// Container App Revisions collection ARM resource.
	RevisionCollection
}

// ContainerAppsRevisionReplicasClientGetReplicaResponse contains the response from method ContainerAppsRevisionReplicasClient.GetReplica.
type ContainerAppsRevisionReplicasClientGetReplicaResponse struct {
	// Container App Revision Replica.
	Replica
}

// ContainerAppsRevisionReplicasClientListReplicasResponse contains the response from method ContainerAppsRevisionReplicasClient.ListReplicas.
type ContainerAppsRevisionReplicasClientListReplicasResponse struct {
	// Container App Revision Replicas collection ARM resource.
	ReplicaCollection
}

// ContainerAppsRevisionsClientActivateRevisionResponse contains the response from method ContainerAppsRevisionsClient.ActivateRevision.
type ContainerAppsRevisionsClientActivateRevisionResponse struct {
	// placeholder for future response values
}

// ContainerAppsRevisionsClientDeactivateRevisionResponse contains the response from method ContainerAppsRevisionsClient.DeactivateRevision.
type ContainerAppsRevisionsClientDeactivateRevisionResponse struct {
	// placeholder for future response values
}

// ContainerAppsRevisionsClientGetRevisionResponse contains the response from method ContainerAppsRevisionsClient.GetRevision.
type ContainerAppsRevisionsClientGetRevisionResponse struct {
	// Container App Revision.
	Revision
}

// ContainerAppsRevisionsClientListRevisionsResponse contains the response from method ContainerAppsRevisionsClient.NewListRevisionsPager.
type ContainerAppsRevisionsClientListRevisionsResponse struct {
	// Container App Revisions collection ARM resource.
	RevisionCollection
}

// ContainerAppsRevisionsClientRestartRevisionResponse contains the response from method ContainerAppsRevisionsClient.RestartRevision.
type ContainerAppsRevisionsClientRestartRevisionResponse struct {
	// placeholder for future response values
}

// ContainerAppsSourceControlsClientCreateOrUpdateResponse contains the response from method ContainerAppsSourceControlsClient.BeginCreateOrUpdate.
type ContainerAppsSourceControlsClientCreateOrUpdateResponse struct {
	// Container App SourceControl.
	SourceControl
}

// ContainerAppsSourceControlsClientDeleteResponse contains the response from method ContainerAppsSourceControlsClient.BeginDelete.
type ContainerAppsSourceControlsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsSourceControlsClientGetResponse contains the response from method ContainerAppsSourceControlsClient.Get.
type ContainerAppsSourceControlsClientGetResponse struct {
	// Container App SourceControl.
	SourceControl
}

// ContainerAppsSourceControlsClientListByContainerAppResponse contains the response from method ContainerAppsSourceControlsClient.NewListByContainerAppPager.
type ContainerAppsSourceControlsClientListByContainerAppResponse struct {
	// SourceControl collection ARM resource.
	SourceControlCollection
}

// DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse contains the response from method DaprComponentResiliencyPoliciesClient.CreateOrUpdate.
type DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse struct {
	// Dapr Component Resiliency Policy.
	DaprComponentResiliencyPolicy
}

// DaprComponentResiliencyPoliciesClientDeleteResponse contains the response from method DaprComponentResiliencyPoliciesClient.Delete.
type DaprComponentResiliencyPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// DaprComponentResiliencyPoliciesClientGetResponse contains the response from method DaprComponentResiliencyPoliciesClient.Get.
type DaprComponentResiliencyPoliciesClientGetResponse struct {
	// Dapr Component Resiliency Policy.
	DaprComponentResiliencyPolicy
}

// DaprComponentResiliencyPoliciesClientListResponse contains the response from method DaprComponentResiliencyPoliciesClient.NewListPager.
type DaprComponentResiliencyPoliciesClientListResponse struct {
	// Dapr Component Resiliency Policies ARM resource.
	DaprComponentResiliencyPoliciesCollection
}

// DaprComponentsClientCreateOrUpdateResponse contains the response from method DaprComponentsClient.CreateOrUpdate.
type DaprComponentsClientCreateOrUpdateResponse struct {
	// Dapr Component.
	DaprComponent
}

// DaprComponentsClientDeleteResponse contains the response from method DaprComponentsClient.Delete.
type DaprComponentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DaprComponentsClientGetResponse contains the response from method DaprComponentsClient.Get.
type DaprComponentsClientGetResponse struct {
	// Dapr Component.
	DaprComponent
}

// DaprComponentsClientListResponse contains the response from method DaprComponentsClient.NewListPager.
type DaprComponentsClientListResponse struct {
	// Dapr Components ARM resource.
	DaprComponentsCollection
}

// DaprComponentsClientListSecretsResponse contains the response from method DaprComponentsClient.ListSecrets.
type DaprComponentsClientListSecretsResponse struct {
	// Dapr component Secrets Collection for ListSecrets Action.
	DaprSecretsCollection
}

// DaprSubscriptionsClientCreateOrUpdateResponse contains the response from method DaprSubscriptionsClient.CreateOrUpdate.
type DaprSubscriptionsClientCreateOrUpdateResponse struct {
	// Dapr PubSub Event Subscription.
	DaprSubscription
}

// DaprSubscriptionsClientDeleteResponse contains the response from method DaprSubscriptionsClient.Delete.
type DaprSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DaprSubscriptionsClientGetResponse contains the response from method DaprSubscriptionsClient.Get.
type DaprSubscriptionsClientGetResponse struct {
	// Dapr PubSub Event Subscription.
	DaprSubscription
}

// DaprSubscriptionsClientListResponse contains the response from method DaprSubscriptionsClient.NewListPager.
type DaprSubscriptionsClientListResponse struct {
	// Dapr Subscriptions ARM resource.
	DaprSubscriptionsCollection
}

// JobsClientCreateOrUpdateResponse contains the response from method JobsClient.BeginCreateOrUpdate.
type JobsClientCreateOrUpdateResponse struct {
	// Container App Job
	Job
}

// JobsClientDeleteResponse contains the response from method JobsClient.BeginDelete.
type JobsClientDeleteResponse struct {
	// placeholder for future response values
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	// Container App Job
	Job
}

// JobsClientListByResourceGroupResponse contains the response from method JobsClient.NewListByResourceGroupPager.
type JobsClientListByResourceGroupResponse struct {
	// Container Apps Jobs collection ARM resource.
	JobsCollection
}

// JobsClientListBySubscriptionResponse contains the response from method JobsClient.NewListBySubscriptionPager.
type JobsClientListBySubscriptionResponse struct {
	// Container Apps Jobs collection ARM resource.
	JobsCollection
}

// JobsClientListSecretsResponse contains the response from method JobsClient.ListSecrets.
type JobsClientListSecretsResponse struct {
	// Container Apps Job Secrets Collection ARM resource.
	JobSecretsCollection
}

// JobsClientStartResponse contains the response from method JobsClient.BeginStart.
type JobsClientStartResponse struct {
	// Container App's Job execution name.
	JobExecutionBase
}

// JobsClientStopExecutionResponse contains the response from method JobsClient.BeginStopExecution.
type JobsClientStopExecutionResponse struct {
	// placeholder for future response values
}

// JobsClientStopMultipleExecutionsResponse contains the response from method JobsClient.BeginStopMultipleExecutions.
type JobsClientStopMultipleExecutionsResponse struct {
	// Container App executions collection ARM resource.
	ContainerAppJobExecutions
}

// JobsClientUpdateResponse contains the response from method JobsClient.BeginUpdate.
type JobsClientUpdateResponse struct {
	// Container App Job
	Job
}

// JobsExecutionsClientListResponse contains the response from method JobsExecutionsClient.NewListPager.
type JobsExecutionsClientListResponse struct {
	// Container App executions collection ARM resource.
	ContainerAppJobExecutions
}

// ManagedCertificatesClientCreateOrUpdateResponse contains the response from method ManagedCertificatesClient.BeginCreateOrUpdate.
type ManagedCertificatesClientCreateOrUpdateResponse struct {
	// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
	ManagedCertificate
}

// ManagedCertificatesClientDeleteResponse contains the response from method ManagedCertificatesClient.Delete.
type ManagedCertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedCertificatesClientGetResponse contains the response from method ManagedCertificatesClient.Get.
type ManagedCertificatesClientGetResponse struct {
	// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
	ManagedCertificate
}

// ManagedCertificatesClientListResponse contains the response from method ManagedCertificatesClient.NewListPager.
type ManagedCertificatesClientListResponse struct {
	// Collection of Managed Certificates.
	ManagedCertificateCollection
}

// ManagedCertificatesClientUpdateResponse contains the response from method ManagedCertificatesClient.Update.
type ManagedCertificatesClientUpdateResponse struct {
	// Managed certificates used for Custom Domain bindings of Container Apps in a Managed Environment
	ManagedCertificate
}

// ManagedEnvironmentDiagnosticsClientGetDetectorResponse contains the response from method ManagedEnvironmentDiagnosticsClient.GetDetector.
type ManagedEnvironmentDiagnosticsClientGetDetectorResponse struct {
	// Diagnostics data for a resource.
	Diagnostics
}

// ManagedEnvironmentDiagnosticsClientListDetectorsResponse contains the response from method ManagedEnvironmentDiagnosticsClient.ListDetectors.
type ManagedEnvironmentDiagnosticsClientListDetectorsResponse struct {
	// Diagnostics data collection for a resource.
	DiagnosticsCollection
}

// ManagedEnvironmentUsagesClientListResponse contains the response from method ManagedEnvironmentUsagesClient.NewListPager.
type ManagedEnvironmentUsagesClientListResponse struct {
	ListUsagesResult
}

// ManagedEnvironmentsClientCreateOrUpdateResponse contains the response from method ManagedEnvironmentsClient.BeginCreateOrUpdate.
type ManagedEnvironmentsClientCreateOrUpdateResponse struct {
	// An environment for hosting container apps
	ManagedEnvironment
}

// ManagedEnvironmentsClientDeleteResponse contains the response from method ManagedEnvironmentsClient.BeginDelete.
type ManagedEnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedEnvironmentsClientGetAuthTokenResponse contains the response from method ManagedEnvironmentsClient.GetAuthToken.
type ManagedEnvironmentsClientGetAuthTokenResponse struct {
	// Environment Auth Token.
	EnvironmentAuthToken
}

// ManagedEnvironmentsClientGetResponse contains the response from method ManagedEnvironmentsClient.Get.
type ManagedEnvironmentsClientGetResponse struct {
	// An environment for hosting container apps
	ManagedEnvironment
}

// ManagedEnvironmentsClientListByResourceGroupResponse contains the response from method ManagedEnvironmentsClient.NewListByResourceGroupPager.
type ManagedEnvironmentsClientListByResourceGroupResponse struct {
	// Collection of Environments
	ManagedEnvironmentsCollection
}

// ManagedEnvironmentsClientListBySubscriptionResponse contains the response from method ManagedEnvironmentsClient.NewListBySubscriptionPager.
type ManagedEnvironmentsClientListBySubscriptionResponse struct {
	// Collection of Environments
	ManagedEnvironmentsCollection
}

// ManagedEnvironmentsClientListWorkloadProfileStatesResponse contains the response from method ManagedEnvironmentsClient.NewListWorkloadProfileStatesPager.
type ManagedEnvironmentsClientListWorkloadProfileStatesResponse struct {
	// Collection of workloadProfileStates
	WorkloadProfileStatesCollection
}

// ManagedEnvironmentsClientUpdateResponse contains the response from method ManagedEnvironmentsClient.BeginUpdate.
type ManagedEnvironmentsClientUpdateResponse struct {
	// An environment for hosting container apps
	ManagedEnvironment
}

// ManagedEnvironmentsDiagnosticsClientGetRootResponse contains the response from method ManagedEnvironmentsDiagnosticsClient.GetRoot.
type ManagedEnvironmentsDiagnosticsClientGetRootResponse struct {
	// An environment for hosting container apps
	ManagedEnvironment
}

// ManagedEnvironmentsStoragesClientCreateOrUpdateResponse contains the response from method ManagedEnvironmentsStoragesClient.CreateOrUpdate.
type ManagedEnvironmentsStoragesClientCreateOrUpdateResponse struct {
	// Storage resource for managedEnvironment.
	ManagedEnvironmentStorage
}

// ManagedEnvironmentsStoragesClientDeleteResponse contains the response from method ManagedEnvironmentsStoragesClient.Delete.
type ManagedEnvironmentsStoragesClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedEnvironmentsStoragesClientGetResponse contains the response from method ManagedEnvironmentsStoragesClient.Get.
type ManagedEnvironmentsStoragesClientGetResponse struct {
	// Storage resource for managedEnvironment.
	ManagedEnvironmentStorage
}

// ManagedEnvironmentsStoragesClientListResponse contains the response from method ManagedEnvironmentsStoragesClient.List.
type ManagedEnvironmentsStoragesClientListResponse struct {
	// Collection of Storage for Environments
	ManagedEnvironmentStoragesCollection
}

// NamespacesClientCheckNameAvailabilityResponse contains the response from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Available operations of the service
	AvailableOperations
}

// PatchesClientApplyResponse contains the response from method PatchesClient.BeginApply.
type PatchesClientApplyResponse struct {
	// Container App Patch
	PatchResource
}

// PatchesClientDeleteResponse contains the response from method PatchesClient.BeginDelete.
type PatchesClientDeleteResponse struct {
	// placeholder for future response values
}

// PatchesClientGetResponse contains the response from method PatchesClient.Get.
type PatchesClientGetResponse struct {
	// Container App Patch
	PatchResource
}

// PatchesClientListByBuilderResourceResponse contains the response from method PatchesClient.ListByBuilderResource.
type PatchesClientListByBuilderResourceResponse struct {
	// Container App SourceToCloud patch collection
	PatchCollection
}

// PatchesClientSkipConfigureResponse contains the response from method PatchesClient.BeginSkipConfigure.
type PatchesClientSkipConfigureResponse struct {
	// placeholder for future response values
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	ListUsagesResult
}
