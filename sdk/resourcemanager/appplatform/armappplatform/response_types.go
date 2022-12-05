//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

// APIPortalCustomDomainsClientCreateOrUpdateResponse contains the response from method APIPortalCustomDomainsClient.CreateOrUpdate.
type APIPortalCustomDomainsClientCreateOrUpdateResponse struct {
	APIPortalCustomDomainResource
}

// APIPortalCustomDomainsClientDeleteResponse contains the response from method APIPortalCustomDomainsClient.Delete.
type APIPortalCustomDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// APIPortalCustomDomainsClientGetResponse contains the response from method APIPortalCustomDomainsClient.Get.
type APIPortalCustomDomainsClientGetResponse struct {
	APIPortalCustomDomainResource
}

// APIPortalCustomDomainsClientListResponse contains the response from method APIPortalCustomDomainsClient.List.
type APIPortalCustomDomainsClientListResponse struct {
	APIPortalCustomDomainResourceCollection
}

// APIPortalsClientCreateOrUpdateResponse contains the response from method APIPortalsClient.CreateOrUpdate.
type APIPortalsClientCreateOrUpdateResponse struct {
	APIPortalResource
}

// APIPortalsClientDeleteResponse contains the response from method APIPortalsClient.Delete.
type APIPortalsClientDeleteResponse struct {
	// placeholder for future response values
}

// APIPortalsClientGetResponse contains the response from method APIPortalsClient.Get.
type APIPortalsClientGetResponse struct {
	APIPortalResource
}

// APIPortalsClientListResponse contains the response from method APIPortalsClient.List.
type APIPortalsClientListResponse struct {
	APIPortalResourceCollection
}

// APIPortalsClientValidateDomainResponse contains the response from method APIPortalsClient.ValidateDomain.
type APIPortalsClientValidateDomainResponse struct {
	CustomDomainValidateResult
}

// AppsClientCreateOrUpdateResponse contains the response from method AppsClient.CreateOrUpdate.
type AppsClientCreateOrUpdateResponse struct {
	AppResource
}

// AppsClientDeleteResponse contains the response from method AppsClient.Delete.
type AppsClientDeleteResponse struct {
	// placeholder for future response values
}

// AppsClientGetResourceUploadURLResponse contains the response from method AppsClient.GetResourceUploadURL.
type AppsClientGetResourceUploadURLResponse struct {
	ResourceUploadDefinition
}

// AppsClientGetResponse contains the response from method AppsClient.Get.
type AppsClientGetResponse struct {
	AppResource
}

// AppsClientListResponse contains the response from method AppsClient.List.
type AppsClientListResponse struct {
	AppResourceCollection
}

// AppsClientSetActiveDeploymentsResponse contains the response from method AppsClient.SetActiveDeployments.
type AppsClientSetActiveDeploymentsResponse struct {
	AppResource
}

// AppsClientUpdateResponse contains the response from method AppsClient.Update.
type AppsClientUpdateResponse struct {
	AppResource
}

// AppsClientValidateDomainResponse contains the response from method AppsClient.ValidateDomain.
type AppsClientValidateDomainResponse struct {
	CustomDomainValidateResult
}

// BindingsClientCreateOrUpdateResponse contains the response from method BindingsClient.CreateOrUpdate.
type BindingsClientCreateOrUpdateResponse struct {
	BindingResource
}

// BindingsClientDeleteResponse contains the response from method BindingsClient.Delete.
type BindingsClientDeleteResponse struct {
	// placeholder for future response values
}

// BindingsClientGetResponse contains the response from method BindingsClient.Get.
type BindingsClientGetResponse struct {
	BindingResource
}

// BindingsClientListResponse contains the response from method BindingsClient.List.
type BindingsClientListResponse struct {
	BindingResourceCollection
}

// BindingsClientUpdateResponse contains the response from method BindingsClient.Update.
type BindingsClientUpdateResponse struct {
	BindingResource
}

// BuildServiceAgentPoolClientGetResponse contains the response from method BuildServiceAgentPoolClient.Get.
type BuildServiceAgentPoolClientGetResponse struct {
	BuildServiceAgentPoolResource
}

// BuildServiceAgentPoolClientListResponse contains the response from method BuildServiceAgentPoolClient.List.
type BuildServiceAgentPoolClientListResponse struct {
	BuildServiceAgentPoolResourceCollection
}

// BuildServiceAgentPoolClientUpdatePutResponse contains the response from method BuildServiceAgentPoolClient.UpdatePut.
type BuildServiceAgentPoolClientUpdatePutResponse struct {
	BuildServiceAgentPoolResource
}

// BuildServiceBuilderClientCreateOrUpdateResponse contains the response from method BuildServiceBuilderClient.CreateOrUpdate.
type BuildServiceBuilderClientCreateOrUpdateResponse struct {
	BuilderResource
}

// BuildServiceBuilderClientDeleteResponse contains the response from method BuildServiceBuilderClient.Delete.
type BuildServiceBuilderClientDeleteResponse struct {
	// placeholder for future response values
}

// BuildServiceBuilderClientGetResponse contains the response from method BuildServiceBuilderClient.Get.
type BuildServiceBuilderClientGetResponse struct {
	BuilderResource
}

// BuildServiceBuilderClientListDeploymentsResponse contains the response from method BuildServiceBuilderClient.ListDeployments.
type BuildServiceBuilderClientListDeploymentsResponse struct {
	DeploymentList
}

// BuildServiceBuilderClientListResponse contains the response from method BuildServiceBuilderClient.List.
type BuildServiceBuilderClientListResponse struct {
	BuilderResourceCollection
}

// BuildServiceClientCreateOrUpdateBuildResponse contains the response from method BuildServiceClient.CreateOrUpdateBuild.
type BuildServiceClientCreateOrUpdateBuildResponse struct {
	Build
}

// BuildServiceClientGetBuildResponse contains the response from method BuildServiceClient.GetBuild.
type BuildServiceClientGetBuildResponse struct {
	Build
}

// BuildServiceClientGetBuildResultLogResponse contains the response from method BuildServiceClient.GetBuildResultLog.
type BuildServiceClientGetBuildResultLogResponse struct {
	BuildResultLog
}

// BuildServiceClientGetBuildResultResponse contains the response from method BuildServiceClient.GetBuildResult.
type BuildServiceClientGetBuildResultResponse struct {
	BuildResult
}

// BuildServiceClientGetBuildServiceResponse contains the response from method BuildServiceClient.GetBuildService.
type BuildServiceClientGetBuildServiceResponse struct {
	BuildService
}

// BuildServiceClientGetResourceUploadURLResponse contains the response from method BuildServiceClient.GetResourceUploadURL.
type BuildServiceClientGetResourceUploadURLResponse struct {
	ResourceUploadDefinition
}

// BuildServiceClientGetSupportedBuildpackResponse contains the response from method BuildServiceClient.GetSupportedBuildpack.
type BuildServiceClientGetSupportedBuildpackResponse struct {
	SupportedBuildpackResource
}

// BuildServiceClientGetSupportedStackResponse contains the response from method BuildServiceClient.GetSupportedStack.
type BuildServiceClientGetSupportedStackResponse struct {
	SupportedStackResource
}

// BuildServiceClientListBuildResultsResponse contains the response from method BuildServiceClient.ListBuildResults.
type BuildServiceClientListBuildResultsResponse struct {
	BuildResultCollection
}

// BuildServiceClientListBuildServicesResponse contains the response from method BuildServiceClient.ListBuildServices.
type BuildServiceClientListBuildServicesResponse struct {
	BuildServiceCollection
}

// BuildServiceClientListBuildsResponse contains the response from method BuildServiceClient.ListBuilds.
type BuildServiceClientListBuildsResponse struct {
	BuildCollection
}

// BuildServiceClientListSupportedBuildpacksResponse contains the response from method BuildServiceClient.ListSupportedBuildpacks.
type BuildServiceClientListSupportedBuildpacksResponse struct {
	SupportedBuildpacksCollection
}

// BuildServiceClientListSupportedStacksResponse contains the response from method BuildServiceClient.ListSupportedStacks.
type BuildServiceClientListSupportedStacksResponse struct {
	SupportedStacksCollection
}

// BuildpackBindingClientCreateOrUpdateResponse contains the response from method BuildpackBindingClient.CreateOrUpdate.
type BuildpackBindingClientCreateOrUpdateResponse struct {
	BuildpackBindingResource
}

// BuildpackBindingClientDeleteResponse contains the response from method BuildpackBindingClient.Delete.
type BuildpackBindingClientDeleteResponse struct {
	// placeholder for future response values
}

// BuildpackBindingClientGetResponse contains the response from method BuildpackBindingClient.Get.
type BuildpackBindingClientGetResponse struct {
	BuildpackBindingResource
}

// BuildpackBindingClientListResponse contains the response from method BuildpackBindingClient.List.
type BuildpackBindingClientListResponse struct {
	BuildpackBindingResourceCollection
}

// CertificatesClientCreateOrUpdateResponse contains the response from method CertificatesClient.CreateOrUpdate.
type CertificatesClientCreateOrUpdateResponse struct {
	CertificateResource
}

// CertificatesClientDeleteResponse contains the response from method CertificatesClient.Delete.
type CertificatesClientDeleteResponse struct {
	// placeholder for future response values
}

// CertificatesClientGetResponse contains the response from method CertificatesClient.Get.
type CertificatesClientGetResponse struct {
	CertificateResource
}

// CertificatesClientListResponse contains the response from method CertificatesClient.List.
type CertificatesClientListResponse struct {
	CertificateResourceCollection
}

// ConfigServersClientGetResponse contains the response from method ConfigServersClient.Get.
type ConfigServersClientGetResponse struct {
	ConfigServerResource
}

// ConfigServersClientUpdatePatchResponse contains the response from method ConfigServersClient.UpdatePatch.
type ConfigServersClientUpdatePatchResponse struct {
	ConfigServerResource
}

// ConfigServersClientUpdatePutResponse contains the response from method ConfigServersClient.UpdatePut.
type ConfigServersClientUpdatePutResponse struct {
	ConfigServerResource
}

// ConfigServersClientValidateResponse contains the response from method ConfigServersClient.Validate.
type ConfigServersClientValidateResponse struct {
	ConfigServerSettingsValidateResult
}

// ConfigurationServicesClientCreateOrUpdateResponse contains the response from method ConfigurationServicesClient.CreateOrUpdate.
type ConfigurationServicesClientCreateOrUpdateResponse struct {
	ConfigurationServiceResource
}

// ConfigurationServicesClientDeleteResponse contains the response from method ConfigurationServicesClient.Delete.
type ConfigurationServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationServicesClientGetResponse contains the response from method ConfigurationServicesClient.Get.
type ConfigurationServicesClientGetResponse struct {
	ConfigurationServiceResource
}

// ConfigurationServicesClientListResponse contains the response from method ConfigurationServicesClient.List.
type ConfigurationServicesClientListResponse struct {
	ConfigurationServiceResourceCollection
}

// ConfigurationServicesClientValidateResponse contains the response from method ConfigurationServicesClient.Validate.
type ConfigurationServicesClientValidateResponse struct {
	ConfigurationServiceSettingsValidateResult
}

// CustomDomainsClientCreateOrUpdateResponse contains the response from method CustomDomainsClient.CreateOrUpdate.
type CustomDomainsClientCreateOrUpdateResponse struct {
	CustomDomainResource
}

// CustomDomainsClientDeleteResponse contains the response from method CustomDomainsClient.Delete.
type CustomDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// CustomDomainsClientGetResponse contains the response from method CustomDomainsClient.Get.
type CustomDomainsClientGetResponse struct {
	CustomDomainResource
}

// CustomDomainsClientListResponse contains the response from method CustomDomainsClient.List.
type CustomDomainsClientListResponse struct {
	CustomDomainResourceCollection
}

// CustomDomainsClientUpdateResponse contains the response from method CustomDomainsClient.Update.
type CustomDomainsClientUpdateResponse struct {
	CustomDomainResource
}

// DeploymentsClientCreateOrUpdateResponse contains the response from method DeploymentsClient.CreateOrUpdate.
type DeploymentsClientCreateOrUpdateResponse struct {
	DeploymentResource
}

// DeploymentsClientDeleteResponse contains the response from method DeploymentsClient.Delete.
type DeploymentsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentsClientDisableRemoteDebuggingResponse contains the response from method DeploymentsClient.DisableRemoteDebugging.
type DeploymentsClientDisableRemoteDebuggingResponse struct {
	RemoteDebugging
}

// DeploymentsClientEnableRemoteDebuggingResponse contains the response from method DeploymentsClient.EnableRemoteDebugging.
type DeploymentsClientEnableRemoteDebuggingResponse struct {
	RemoteDebugging
}

// DeploymentsClientGenerateHeapDumpResponse contains the response from method DeploymentsClient.GenerateHeapDump.
type DeploymentsClientGenerateHeapDumpResponse struct {
	// placeholder for future response values
}

// DeploymentsClientGenerateThreadDumpResponse contains the response from method DeploymentsClient.GenerateThreadDump.
type DeploymentsClientGenerateThreadDumpResponse struct {
	// placeholder for future response values
}

// DeploymentsClientGetLogFileURLResponse contains the response from method DeploymentsClient.GetLogFileURL.
type DeploymentsClientGetLogFileURLResponse struct {
	LogFileURLResponse
}

// DeploymentsClientGetRemoteDebuggingConfigResponse contains the response from method DeploymentsClient.GetRemoteDebuggingConfig.
type DeploymentsClientGetRemoteDebuggingConfigResponse struct {
	RemoteDebugging
}

// DeploymentsClientGetResponse contains the response from method DeploymentsClient.Get.
type DeploymentsClientGetResponse struct {
	DeploymentResource
}

// DeploymentsClientListForClusterResponse contains the response from method DeploymentsClient.ListForCluster.
type DeploymentsClientListForClusterResponse struct {
	DeploymentResourceCollection
}

// DeploymentsClientListResponse contains the response from method DeploymentsClient.List.
type DeploymentsClientListResponse struct {
	DeploymentResourceCollection
}

// DeploymentsClientRestartResponse contains the response from method DeploymentsClient.Restart.
type DeploymentsClientRestartResponse struct {
	// placeholder for future response values
}

// DeploymentsClientStartJFRResponse contains the response from method DeploymentsClient.StartJFR.
type DeploymentsClientStartJFRResponse struct {
	// placeholder for future response values
}

// DeploymentsClientStartResponse contains the response from method DeploymentsClient.Start.
type DeploymentsClientStartResponse struct {
	// placeholder for future response values
}

// DeploymentsClientStopResponse contains the response from method DeploymentsClient.Stop.
type DeploymentsClientStopResponse struct {
	// placeholder for future response values
}

// DeploymentsClientUpdateResponse contains the response from method DeploymentsClient.Update.
type DeploymentsClientUpdateResponse struct {
	DeploymentResource
}

// GatewayCustomDomainsClientCreateOrUpdateResponse contains the response from method GatewayCustomDomainsClient.CreateOrUpdate.
type GatewayCustomDomainsClientCreateOrUpdateResponse struct {
	GatewayCustomDomainResource
}

// GatewayCustomDomainsClientDeleteResponse contains the response from method GatewayCustomDomainsClient.Delete.
type GatewayCustomDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// GatewayCustomDomainsClientGetResponse contains the response from method GatewayCustomDomainsClient.Get.
type GatewayCustomDomainsClientGetResponse struct {
	GatewayCustomDomainResource
}

// GatewayCustomDomainsClientListResponse contains the response from method GatewayCustomDomainsClient.List.
type GatewayCustomDomainsClientListResponse struct {
	GatewayCustomDomainResourceCollection
}

// GatewayRouteConfigsClientCreateOrUpdateResponse contains the response from method GatewayRouteConfigsClient.CreateOrUpdate.
type GatewayRouteConfigsClientCreateOrUpdateResponse struct {
	GatewayRouteConfigResource
}

// GatewayRouteConfigsClientDeleteResponse contains the response from method GatewayRouteConfigsClient.Delete.
type GatewayRouteConfigsClientDeleteResponse struct {
	// placeholder for future response values
}

// GatewayRouteConfigsClientGetResponse contains the response from method GatewayRouteConfigsClient.Get.
type GatewayRouteConfigsClientGetResponse struct {
	GatewayRouteConfigResource
}

// GatewayRouteConfigsClientListResponse contains the response from method GatewayRouteConfigsClient.List.
type GatewayRouteConfigsClientListResponse struct {
	GatewayRouteConfigResourceCollection
}

// GatewaysClientCreateOrUpdateResponse contains the response from method GatewaysClient.CreateOrUpdate.
type GatewaysClientCreateOrUpdateResponse struct {
	GatewayResource
}

// GatewaysClientDeleteResponse contains the response from method GatewaysClient.Delete.
type GatewaysClientDeleteResponse struct {
	// placeholder for future response values
}

// GatewaysClientGetResponse contains the response from method GatewaysClient.Get.
type GatewaysClientGetResponse struct {
	GatewayResource
}

// GatewaysClientListResponse contains the response from method GatewaysClient.List.
type GatewaysClientListResponse struct {
	GatewayResourceCollection
}

// GatewaysClientValidateDomainResponse contains the response from method GatewaysClient.ValidateDomain.
type GatewaysClientValidateDomainResponse struct {
	CustomDomainValidateResult
}

// MonitoringSettingsClientGetResponse contains the response from method MonitoringSettingsClient.Get.
type MonitoringSettingsClientGetResponse struct {
	MonitoringSettingResource
}

// MonitoringSettingsClientUpdatePatchResponse contains the response from method MonitoringSettingsClient.UpdatePatch.
type MonitoringSettingsClientUpdatePatchResponse struct {
	MonitoringSettingResource
}

// MonitoringSettingsClientUpdatePutResponse contains the response from method MonitoringSettingsClient.UpdatePut.
type MonitoringSettingsClientUpdatePutResponse struct {
	MonitoringSettingResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	AvailableOperations
}

// RuntimeVersionsClientListRuntimeVersionsResponse contains the response from method RuntimeVersionsClient.ListRuntimeVersions.
type RuntimeVersionsClientListRuntimeVersionsResponse struct {
	AvailableRuntimeVersions
}

// SKUsClientListResponse contains the response from method SKUsClient.List.
type SKUsClientListResponse struct {
	ResourceSKUCollection
}

// ServiceRegistriesClientCreateOrUpdateResponse contains the response from method ServiceRegistriesClient.CreateOrUpdate.
type ServiceRegistriesClientCreateOrUpdateResponse struct {
	ServiceRegistryResource
}

// ServiceRegistriesClientDeleteResponse contains the response from method ServiceRegistriesClient.Delete.
type ServiceRegistriesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceRegistriesClientGetResponse contains the response from method ServiceRegistriesClient.Get.
type ServiceRegistriesClientGetResponse struct {
	ServiceRegistryResource
}

// ServiceRegistriesClientListResponse contains the response from method ServiceRegistriesClient.List.
type ServiceRegistriesClientListResponse struct {
	ServiceRegistryResourceCollection
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	NameAvailability
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientDisableTestEndpointResponse contains the response from method ServicesClient.DisableTestEndpoint.
type ServicesClientDisableTestEndpointResponse struct {
	// placeholder for future response values
}

// ServicesClientEnableTestEndpointResponse contains the response from method ServicesClient.EnableTestEndpoint.
type ServicesClientEnableTestEndpointResponse struct {
	TestKeys
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	ServiceResource
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.ListBySubscription.
type ServicesClientListBySubscriptionResponse struct {
	ServiceResourceList
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	ServiceResourceList
}

// ServicesClientListTestKeysResponse contains the response from method ServicesClient.ListTestKeys.
type ServicesClientListTestKeysResponse struct {
	TestKeys
}

// ServicesClientRegenerateTestKeyResponse contains the response from method ServicesClient.RegenerateTestKey.
type ServicesClientRegenerateTestKeyResponse struct {
	TestKeys
}

// ServicesClientStartResponse contains the response from method ServicesClient.Start.
type ServicesClientStartResponse struct {
	// placeholder for future response values
}

// ServicesClientStopResponse contains the response from method ServicesClient.Stop.
type ServicesClientStopResponse struct {
	// placeholder for future response values
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	ServiceResource
}

// StoragesClientCreateOrUpdateResponse contains the response from method StoragesClient.CreateOrUpdate.
type StoragesClientCreateOrUpdateResponse struct {
	StorageResource
}

// StoragesClientDeleteResponse contains the response from method StoragesClient.Delete.
type StoragesClientDeleteResponse struct {
	// placeholder for future response values
}

// StoragesClientGetResponse contains the response from method StoragesClient.Get.
type StoragesClientGetResponse struct {
	StorageResource
}

// StoragesClientListResponse contains the response from method StoragesClient.List.
type StoragesClientListResponse struct {
	StorageResourceCollection
}
