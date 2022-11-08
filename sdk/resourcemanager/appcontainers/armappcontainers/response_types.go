//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

// AvailableWorkloadProfilesClientGetResponse contains the response from method AvailableWorkloadProfilesClient.Get.
type AvailableWorkloadProfilesClientGetResponse struct {
	AvailableWorkloadProfilesCollection
}

// BillingMetersClientGetResponse contains the response from method BillingMetersClient.Get.
type BillingMetersClientGetResponse struct {
	BillingMeterCollection
}

// CertificatesClientCreateOrUpdateResponse contains the response from method CertificatesClient.CreateOrUpdate.
type CertificatesClientCreateOrUpdateResponse struct {
	Certificate
}

// CertificatesClientDeleteResponse contains the response from method CertificatesClient.Delete.
type CertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificatesClientGetResponse contains the response from method CertificatesClient.Get.
type CertificatesClientGetResponse struct {
	Certificate
}

// CertificatesClientListResponse contains the response from method CertificatesClient.List.
type CertificatesClientListResponse struct {
	CertificateCollection
}

// CertificatesClientUpdateResponse contains the response from method CertificatesClient.Update.
type CertificatesClientUpdateResponse struct {
	Certificate
}

// ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsCertificatesClient.CreateOrUpdate.
type ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse struct {
	Certificate
}

// ConnectedEnvironmentsCertificatesClientDeleteResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Delete.
type ConnectedEnvironmentsCertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsCertificatesClientGetResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Get.
type ConnectedEnvironmentsCertificatesClientGetResponse struct {
	Certificate
}

// ConnectedEnvironmentsCertificatesClientListResponse contains the response from method ConnectedEnvironmentsCertificatesClient.List.
type ConnectedEnvironmentsCertificatesClientListResponse struct {
	CertificateCollection
}

// ConnectedEnvironmentsCertificatesClientUpdateResponse contains the response from method ConnectedEnvironmentsCertificatesClient.Update.
type ConnectedEnvironmentsCertificatesClientUpdateResponse struct {
	Certificate
}

// ConnectedEnvironmentsClientCheckNameAvailabilityResponse contains the response from method ConnectedEnvironmentsClient.CheckNameAvailability.
type ConnectedEnvironmentsClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResponse
}

// ConnectedEnvironmentsClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsClient.CreateOrUpdate.
type ConnectedEnvironmentsClientCreateOrUpdateResponse struct {
	ConnectedEnvironment
}

// ConnectedEnvironmentsClientDeleteResponse contains the response from method ConnectedEnvironmentsClient.Delete.
type ConnectedEnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsClientGetResponse contains the response from method ConnectedEnvironmentsClient.Get.
type ConnectedEnvironmentsClientGetResponse struct {
	ConnectedEnvironment
}

// ConnectedEnvironmentsClientListByResourceGroupResponse contains the response from method ConnectedEnvironmentsClient.ListByResourceGroup.
type ConnectedEnvironmentsClientListByResourceGroupResponse struct {
	ConnectedEnvironmentCollection
}

// ConnectedEnvironmentsClientListBySubscriptionResponse contains the response from method ConnectedEnvironmentsClient.ListBySubscription.
type ConnectedEnvironmentsClientListBySubscriptionResponse struct {
	ConnectedEnvironmentCollection
}

// ConnectedEnvironmentsClientUpdateResponse contains the response from method ConnectedEnvironmentsClient.Update.
type ConnectedEnvironmentsClientUpdateResponse struct {
	ConnectedEnvironment
}

// ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.CreateOrUpdate.
type ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse struct {
	DaprComponent
}

// ConnectedEnvironmentsDaprComponentsClientDeleteResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.Delete.
type ConnectedEnvironmentsDaprComponentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsDaprComponentsClientGetResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.Get.
type ConnectedEnvironmentsDaprComponentsClientGetResponse struct {
	DaprComponent
}

// ConnectedEnvironmentsDaprComponentsClientListResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.List.
type ConnectedEnvironmentsDaprComponentsClientListResponse struct {
	DaprComponentsCollection
}

// ConnectedEnvironmentsDaprComponentsClientListSecretsResponse contains the response from method ConnectedEnvironmentsDaprComponentsClient.ListSecrets.
type ConnectedEnvironmentsDaprComponentsClientListSecretsResponse struct {
	DaprSecretsCollection
}

// ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse contains the response from method ConnectedEnvironmentsStoragesClient.CreateOrUpdate.
type ConnectedEnvironmentsStoragesClientCreateOrUpdateResponse struct {
	ConnectedEnvironmentStorage
}

// ConnectedEnvironmentsStoragesClientDeleteResponse contains the response from method ConnectedEnvironmentsStoragesClient.Delete.
type ConnectedEnvironmentsStoragesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectedEnvironmentsStoragesClientGetResponse contains the response from method ConnectedEnvironmentsStoragesClient.Get.
type ConnectedEnvironmentsStoragesClientGetResponse struct {
	ConnectedEnvironmentStorage
}

// ConnectedEnvironmentsStoragesClientListResponse contains the response from method ConnectedEnvironmentsStoragesClient.List.
type ConnectedEnvironmentsStoragesClientListResponse struct {
	ConnectedEnvironmentStoragesCollection
}

// ContainerAppsAuthConfigsClientCreateOrUpdateResponse contains the response from method ContainerAppsAuthConfigsClient.CreateOrUpdate.
type ContainerAppsAuthConfigsClientCreateOrUpdateResponse struct {
	AuthConfig
}

// ContainerAppsAuthConfigsClientDeleteResponse contains the response from method ContainerAppsAuthConfigsClient.Delete.
type ContainerAppsAuthConfigsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsAuthConfigsClientGetResponse contains the response from method ContainerAppsAuthConfigsClient.Get.
type ContainerAppsAuthConfigsClientGetResponse struct {
	AuthConfig
}

// ContainerAppsAuthConfigsClientListByContainerAppResponse contains the response from method ContainerAppsAuthConfigsClient.ListByContainerApp.
type ContainerAppsAuthConfigsClientListByContainerAppResponse struct {
	AuthConfigCollection
}

// ContainerAppsClientCreateOrUpdateResponse contains the response from method ContainerAppsClient.CreateOrUpdate.
type ContainerAppsClientCreateOrUpdateResponse struct {
	ContainerApp
}

// ContainerAppsClientDeleteResponse contains the response from method ContainerAppsClient.Delete.
type ContainerAppsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsClientGetAuthTokenResponse contains the response from method ContainerAppsClient.GetAuthToken.
type ContainerAppsClientGetAuthTokenResponse struct {
	ContainerAppAuthToken
}

// ContainerAppsClientGetResponse contains the response from method ContainerAppsClient.Get.
type ContainerAppsClientGetResponse struct {
	ContainerApp
}

// ContainerAppsClientListByResourceGroupResponse contains the response from method ContainerAppsClient.ListByResourceGroup.
type ContainerAppsClientListByResourceGroupResponse struct {
	ContainerAppCollection
}

// ContainerAppsClientListBySubscriptionResponse contains the response from method ContainerAppsClient.ListBySubscription.
type ContainerAppsClientListBySubscriptionResponse struct {
	ContainerAppCollection
}

// ContainerAppsClientListCustomHostNameAnalysisResponse contains the response from method ContainerAppsClient.ListCustomHostNameAnalysis.
type ContainerAppsClientListCustomHostNameAnalysisResponse struct {
	CustomHostnameAnalysisResult
}

// ContainerAppsClientListSecretsResponse contains the response from method ContainerAppsClient.ListSecrets.
type ContainerAppsClientListSecretsResponse struct {
	SecretsCollection
}

// ContainerAppsClientUpdateResponse contains the response from method ContainerAppsClient.Update.
type ContainerAppsClientUpdateResponse struct {
	ContainerApp
}

// ContainerAppsDiagnosticsClientGetDetectorResponse contains the response from method ContainerAppsDiagnosticsClient.GetDetector.
type ContainerAppsDiagnosticsClientGetDetectorResponse struct {
	Diagnostics
}

// ContainerAppsDiagnosticsClientGetRevisionResponse contains the response from method ContainerAppsDiagnosticsClient.GetRevision.
type ContainerAppsDiagnosticsClientGetRevisionResponse struct {
	Revision
}

// ContainerAppsDiagnosticsClientGetRootResponse contains the response from method ContainerAppsDiagnosticsClient.GetRoot.
type ContainerAppsDiagnosticsClientGetRootResponse struct {
	ContainerApp
}

// ContainerAppsDiagnosticsClientListDetectorsResponse contains the response from method ContainerAppsDiagnosticsClient.ListDetectors.
type ContainerAppsDiagnosticsClientListDetectorsResponse struct {
	DiagnosticsCollection
}

// ContainerAppsDiagnosticsClientListRevisionsResponse contains the response from method ContainerAppsDiagnosticsClient.ListRevisions.
type ContainerAppsDiagnosticsClientListRevisionsResponse struct {
	RevisionCollection
}

// ContainerAppsRevisionReplicasClientGetReplicaResponse contains the response from method ContainerAppsRevisionReplicasClient.GetReplica.
type ContainerAppsRevisionReplicasClientGetReplicaResponse struct {
	Replica
}

// ContainerAppsRevisionReplicasClientListReplicasResponse contains the response from method ContainerAppsRevisionReplicasClient.ListReplicas.
type ContainerAppsRevisionReplicasClientListReplicasResponse struct {
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
	Revision
}

// ContainerAppsRevisionsClientListRevisionsResponse contains the response from method ContainerAppsRevisionsClient.ListRevisions.
type ContainerAppsRevisionsClientListRevisionsResponse struct {
	RevisionCollection
}

// ContainerAppsRevisionsClientRestartRevisionResponse contains the response from method ContainerAppsRevisionsClient.RestartRevision.
type ContainerAppsRevisionsClientRestartRevisionResponse struct {
	// placeholder for future response values
}

// ContainerAppsSourceControlsClientCreateOrUpdateResponse contains the response from method ContainerAppsSourceControlsClient.CreateOrUpdate.
type ContainerAppsSourceControlsClientCreateOrUpdateResponse struct {
	SourceControl
}

// ContainerAppsSourceControlsClientDeleteResponse contains the response from method ContainerAppsSourceControlsClient.Delete.
type ContainerAppsSourceControlsClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainerAppsSourceControlsClientGetResponse contains the response from method ContainerAppsSourceControlsClient.Get.
type ContainerAppsSourceControlsClientGetResponse struct {
	SourceControl
}

// ContainerAppsSourceControlsClientListByContainerAppResponse contains the response from method ContainerAppsSourceControlsClient.ListByContainerApp.
type ContainerAppsSourceControlsClientListByContainerAppResponse struct {
	SourceControlCollection
}

// DaprComponentsClientCreateOrUpdateResponse contains the response from method DaprComponentsClient.CreateOrUpdate.
type DaprComponentsClientCreateOrUpdateResponse struct {
	DaprComponent
}

// DaprComponentsClientDeleteResponse contains the response from method DaprComponentsClient.Delete.
type DaprComponentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DaprComponentsClientGetResponse contains the response from method DaprComponentsClient.Get.
type DaprComponentsClientGetResponse struct {
	DaprComponent
}

// DaprComponentsClientListResponse contains the response from method DaprComponentsClient.List.
type DaprComponentsClientListResponse struct {
	DaprComponentsCollection
}

// DaprComponentsClientListSecretsResponse contains the response from method DaprComponentsClient.ListSecrets.
type DaprComponentsClientListSecretsResponse struct {
	DaprSecretsCollection
}

// ManagedEnvironmentDiagnosticsClientGetDetectorResponse contains the response from method ManagedEnvironmentDiagnosticsClient.GetDetector.
type ManagedEnvironmentDiagnosticsClientGetDetectorResponse struct {
	Diagnostics
}

// ManagedEnvironmentDiagnosticsClientListDetectorsResponse contains the response from method ManagedEnvironmentDiagnosticsClient.ListDetectors.
type ManagedEnvironmentDiagnosticsClientListDetectorsResponse struct {
	DiagnosticsCollection
}

// ManagedEnvironmentsClientCreateOrUpdateResponse contains the response from method ManagedEnvironmentsClient.CreateOrUpdate.
type ManagedEnvironmentsClientCreateOrUpdateResponse struct {
	ManagedEnvironment
}

// ManagedEnvironmentsClientDeleteResponse contains the response from method ManagedEnvironmentsClient.Delete.
type ManagedEnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedEnvironmentsClientGetAuthTokenResponse contains the response from method ManagedEnvironmentsClient.GetAuthToken.
type ManagedEnvironmentsClientGetAuthTokenResponse struct {
	EnvironmentAuthToken
}

// ManagedEnvironmentsClientGetResponse contains the response from method ManagedEnvironmentsClient.Get.
type ManagedEnvironmentsClientGetResponse struct {
	ManagedEnvironment
}

// ManagedEnvironmentsClientListByResourceGroupResponse contains the response from method ManagedEnvironmentsClient.ListByResourceGroup.
type ManagedEnvironmentsClientListByResourceGroupResponse struct {
	ManagedEnvironmentsCollection
}

// ManagedEnvironmentsClientListBySubscriptionResponse contains the response from method ManagedEnvironmentsClient.ListBySubscription.
type ManagedEnvironmentsClientListBySubscriptionResponse struct {
	ManagedEnvironmentsCollection
}

// ManagedEnvironmentsClientListWorkloadProfileStatesResponse contains the response from method ManagedEnvironmentsClient.ListWorkloadProfileStates.
type ManagedEnvironmentsClientListWorkloadProfileStatesResponse struct {
	WorkloadProfileStatesCollection
}

// ManagedEnvironmentsClientUpdateResponse contains the response from method ManagedEnvironmentsClient.Update.
type ManagedEnvironmentsClientUpdateResponse struct {
	ManagedEnvironment
}

// ManagedEnvironmentsDiagnosticsClientGetRootResponse contains the response from method ManagedEnvironmentsDiagnosticsClient.GetRoot.
type ManagedEnvironmentsDiagnosticsClientGetRootResponse struct {
	ManagedEnvironment
}

// ManagedEnvironmentsStoragesClientCreateOrUpdateResponse contains the response from method ManagedEnvironmentsStoragesClient.CreateOrUpdate.
type ManagedEnvironmentsStoragesClientCreateOrUpdateResponse struct {
	ManagedEnvironmentStorage
}

// ManagedEnvironmentsStoragesClientDeleteResponse contains the response from method ManagedEnvironmentsStoragesClient.Delete.
type ManagedEnvironmentsStoragesClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedEnvironmentsStoragesClientGetResponse contains the response from method ManagedEnvironmentsStoragesClient.Get.
type ManagedEnvironmentsStoragesClientGetResponse struct {
	ManagedEnvironmentStorage
}

// ManagedEnvironmentsStoragesClientListResponse contains the response from method ManagedEnvironmentsStoragesClient.List.
type ManagedEnvironmentsStoragesClientListResponse struct {
	ManagedEnvironmentStoragesCollection
}

// NamespacesClientCheckNameAvailabilityResponse contains the response from method NamespacesClient.CheckNameAvailability.
type NamespacesClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	AvailableOperations
}
