//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

// AttachedNetworksClientCreateOrUpdateResponse contains the response from method AttachedNetworksClient.BeginCreateOrUpdate.
type AttachedNetworksClientCreateOrUpdateResponse struct {
	// Represents an attached NetworkConnection.
	AttachedNetworkConnection
}

// AttachedNetworksClientDeleteResponse contains the response from method AttachedNetworksClient.BeginDelete.
type AttachedNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// AttachedNetworksClientGetByDevCenterResponse contains the response from method AttachedNetworksClient.GetByDevCenter.
type AttachedNetworksClientGetByDevCenterResponse struct {
	// Represents an attached NetworkConnection.
	AttachedNetworkConnection
}

// AttachedNetworksClientGetByProjectResponse contains the response from method AttachedNetworksClient.GetByProject.
type AttachedNetworksClientGetByProjectResponse struct {
	// Represents an attached NetworkConnection.
	AttachedNetworkConnection
}

// AttachedNetworksClientListByDevCenterResponse contains the response from method AttachedNetworksClient.NewListByDevCenterPager.
type AttachedNetworksClientListByDevCenterResponse struct {
	// Results of the Attached Networks list operation.
	AttachedNetworkListResult
}

// AttachedNetworksClientListByProjectResponse contains the response from method AttachedNetworksClient.NewListByProjectPager.
type AttachedNetworksClientListByProjectResponse struct {
	// Results of the Attached Networks list operation.
	AttachedNetworkListResult
}

// CatalogsClientCreateOrUpdateResponse contains the response from method CatalogsClient.BeginCreateOrUpdate.
type CatalogsClientCreateOrUpdateResponse struct {
	// Represents a catalog.
	Catalog
}

// CatalogsClientDeleteResponse contains the response from method CatalogsClient.BeginDelete.
type CatalogsClientDeleteResponse struct {
	// placeholder for future response values
}

// CatalogsClientGetResponse contains the response from method CatalogsClient.Get.
type CatalogsClientGetResponse struct {
	// Represents a catalog.
	Catalog
}

// CatalogsClientListByDevCenterResponse contains the response from method CatalogsClient.NewListByDevCenterPager.
type CatalogsClientListByDevCenterResponse struct {
	// Results of the catalog list operation.
	CatalogListResult
}

// CatalogsClientSyncResponse contains the response from method CatalogsClient.BeginSync.
type CatalogsClientSyncResponse struct {
	// placeholder for future response values
}

// CatalogsClientUpdateResponse contains the response from method CatalogsClient.BeginUpdate.
type CatalogsClientUpdateResponse struct {
	// Represents a catalog.
	Catalog
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// DevBoxDefinitionsClientCreateOrUpdateResponse contains the response from method DevBoxDefinitionsClient.BeginCreateOrUpdate.
type DevBoxDefinitionsClientCreateOrUpdateResponse struct {
	// Represents a definition for a Developer Machine.
	DevBoxDefinition
}

// DevBoxDefinitionsClientDeleteResponse contains the response from method DevBoxDefinitionsClient.BeginDelete.
type DevBoxDefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DevBoxDefinitionsClientGetByProjectResponse contains the response from method DevBoxDefinitionsClient.GetByProject.
type DevBoxDefinitionsClientGetByProjectResponse struct {
	// Represents a definition for a Developer Machine.
	DevBoxDefinition
}

// DevBoxDefinitionsClientGetResponse contains the response from method DevBoxDefinitionsClient.Get.
type DevBoxDefinitionsClientGetResponse struct {
	// Represents a definition for a Developer Machine.
	DevBoxDefinition
}

// DevBoxDefinitionsClientListByDevCenterResponse contains the response from method DevBoxDefinitionsClient.NewListByDevCenterPager.
type DevBoxDefinitionsClientListByDevCenterResponse struct {
	// Results of the Dev Box definition list operation.
	DevBoxDefinitionListResult
}

// DevBoxDefinitionsClientListByProjectResponse contains the response from method DevBoxDefinitionsClient.NewListByProjectPager.
type DevBoxDefinitionsClientListByProjectResponse struct {
	// Results of the Dev Box definition list operation.
	DevBoxDefinitionListResult
}

// DevBoxDefinitionsClientUpdateResponse contains the response from method DevBoxDefinitionsClient.BeginUpdate.
type DevBoxDefinitionsClientUpdateResponse struct {
	// Represents a definition for a Developer Machine.
	DevBoxDefinition
}

// DevCentersClientCreateOrUpdateResponse contains the response from method DevCentersClient.BeginCreateOrUpdate.
type DevCentersClientCreateOrUpdateResponse struct {
	// Represents a devcenter resource.
	DevCenter
}

// DevCentersClientDeleteResponse contains the response from method DevCentersClient.BeginDelete.
type DevCentersClientDeleteResponse struct {
	// placeholder for future response values
}

// DevCentersClientGetResponse contains the response from method DevCentersClient.Get.
type DevCentersClientGetResponse struct {
	// Represents a devcenter resource.
	DevCenter
}

// DevCentersClientListByResourceGroupResponse contains the response from method DevCentersClient.NewListByResourceGroupPager.
type DevCentersClientListByResourceGroupResponse struct {
	// Result of the list devcenters operation
	ListResult
}

// DevCentersClientListBySubscriptionResponse contains the response from method DevCentersClient.NewListBySubscriptionPager.
type DevCentersClientListBySubscriptionResponse struct {
	// Result of the list devcenters operation
	ListResult
}

// DevCentersClientUpdateResponse contains the response from method DevCentersClient.BeginUpdate.
type DevCentersClientUpdateResponse struct {
	// Represents a devcenter resource.
	DevCenter
}

// EnvironmentTypesClientCreateOrUpdateResponse contains the response from method EnvironmentTypesClient.CreateOrUpdate.
type EnvironmentTypesClientCreateOrUpdateResponse struct {
	// Represents an environment type.
	EnvironmentType
}

// EnvironmentTypesClientDeleteResponse contains the response from method EnvironmentTypesClient.Delete.
type EnvironmentTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// EnvironmentTypesClientGetResponse contains the response from method EnvironmentTypesClient.Get.
type EnvironmentTypesClientGetResponse struct {
	// Represents an environment type.
	EnvironmentType
}

// EnvironmentTypesClientListByDevCenterResponse contains the response from method EnvironmentTypesClient.NewListByDevCenterPager.
type EnvironmentTypesClientListByDevCenterResponse struct {
	// Result of the environment type list operation.
	EnvironmentTypeListResult
}

// EnvironmentTypesClientUpdateResponse contains the response from method EnvironmentTypesClient.Update.
type EnvironmentTypesClientUpdateResponse struct {
	// Represents an environment type.
	EnvironmentType
}

// GalleriesClientCreateOrUpdateResponse contains the response from method GalleriesClient.BeginCreateOrUpdate.
type GalleriesClientCreateOrUpdateResponse struct {
	// Represents a gallery.
	Gallery
}

// GalleriesClientDeleteResponse contains the response from method GalleriesClient.BeginDelete.
type GalleriesClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleriesClientGetResponse contains the response from method GalleriesClient.Get.
type GalleriesClientGetResponse struct {
	// Represents a gallery.
	Gallery
}

// GalleriesClientListByDevCenterResponse contains the response from method GalleriesClient.NewListByDevCenterPager.
type GalleriesClientListByDevCenterResponse struct {
	// Results of the gallery list operation.
	GalleryListResult
}

// ImageVersionsClientGetResponse contains the response from method ImageVersionsClient.Get.
type ImageVersionsClientGetResponse struct {
	// Represents an image version.
	ImageVersion
}

// ImageVersionsClientListByImageResponse contains the response from method ImageVersionsClient.NewListByImagePager.
type ImageVersionsClientListByImageResponse struct {
	// Results of the image version list operation.
	ImageVersionListResult
}

// ImagesClientGetResponse contains the response from method ImagesClient.Get.
type ImagesClientGetResponse struct {
	// Represents an image.
	Image
}

// ImagesClientListByDevCenterResponse contains the response from method ImagesClient.NewListByDevCenterPager.
type ImagesClientListByDevCenterResponse struct {
	// Results of the image list operation.
	ImageListResult
}

// ImagesClientListByGalleryResponse contains the response from method ImagesClient.NewListByGalleryPager.
type ImagesClientListByGalleryResponse struct {
	// Results of the image list operation.
	ImageListResult
}

// NetworkConnectionsClientCreateOrUpdateResponse contains the response from method NetworkConnectionsClient.BeginCreateOrUpdate.
type NetworkConnectionsClientCreateOrUpdateResponse struct {
	// Network related settings
	NetworkConnection
}

// NetworkConnectionsClientDeleteResponse contains the response from method NetworkConnectionsClient.BeginDelete.
type NetworkConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkConnectionsClientGetHealthDetailsResponse contains the response from method NetworkConnectionsClient.GetHealthDetails.
type NetworkConnectionsClientGetHealthDetailsResponse struct {
	// Health Check details.
	HealthCheckStatusDetails
}

// NetworkConnectionsClientGetResponse contains the response from method NetworkConnectionsClient.Get.
type NetworkConnectionsClientGetResponse struct {
	// Network related settings
	NetworkConnection
}

// NetworkConnectionsClientListByResourceGroupResponse contains the response from method NetworkConnectionsClient.NewListByResourceGroupPager.
type NetworkConnectionsClientListByResourceGroupResponse struct {
	// Result of the network connection list operation.
	NetworkConnectionListResult
}

// NetworkConnectionsClientListBySubscriptionResponse contains the response from method NetworkConnectionsClient.NewListBySubscriptionPager.
type NetworkConnectionsClientListBySubscriptionResponse struct {
	// Result of the network connection list operation.
	NetworkConnectionListResult
}

// NetworkConnectionsClientListHealthDetailsResponse contains the response from method NetworkConnectionsClient.NewListHealthDetailsPager.
type NetworkConnectionsClientListHealthDetailsResponse struct {
	// Result of the network health check list operation.
	HealthCheckStatusDetailsListResult
}

// NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method NetworkConnectionsClient.NewListOutboundNetworkDependenciesEndpointsPager.
type NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse struct {
	// Values returned by the List operation.
	OutboundEnvironmentEndpointCollection
}

// NetworkConnectionsClientRunHealthChecksResponse contains the response from method NetworkConnectionsClient.BeginRunHealthChecks.
type NetworkConnectionsClientRunHealthChecksResponse struct {
	// placeholder for future response values
}

// NetworkConnectionsClientUpdateResponse contains the response from method NetworkConnectionsClient.BeginUpdate.
type NetworkConnectionsClientUpdateResponse struct {
	// Network related settings
	NetworkConnection
}

// OperationStatusesClientGetResponse contains the response from method OperationStatusesClient.Get.
type OperationStatusesClientGetResponse struct {
	// The current status of an async operation
	OperationStatus
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PoolsClientCreateOrUpdateResponse contains the response from method PoolsClient.BeginCreateOrUpdate.
type PoolsClientCreateOrUpdateResponse struct {
	// A pool of Virtual Machines.
	Pool
}

// PoolsClientDeleteResponse contains the response from method PoolsClient.BeginDelete.
type PoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// PoolsClientGetResponse contains the response from method PoolsClient.Get.
type PoolsClientGetResponse struct {
	// A pool of Virtual Machines.
	Pool
}

// PoolsClientListByProjectResponse contains the response from method PoolsClient.NewListByProjectPager.
type PoolsClientListByProjectResponse struct {
	// Results of the machine pool list operation.
	PoolListResult
}

// PoolsClientRunHealthChecksResponse contains the response from method PoolsClient.BeginRunHealthChecks.
type PoolsClientRunHealthChecksResponse struct {
	// placeholder for future response values
}

// PoolsClientUpdateResponse contains the response from method PoolsClient.BeginUpdate.
type PoolsClientUpdateResponse struct {
	// A pool of Virtual Machines.
	Pool
}

// ProjectAllowedEnvironmentTypesClientGetResponse contains the response from method ProjectAllowedEnvironmentTypesClient.Get.
type ProjectAllowedEnvironmentTypesClientGetResponse struct {
	// Represents an allowed environment type.
	AllowedEnvironmentType
}

// ProjectAllowedEnvironmentTypesClientListResponse contains the response from method ProjectAllowedEnvironmentTypesClient.NewListPager.
type ProjectAllowedEnvironmentTypesClientListResponse struct {
	// Result of the allowed environment type list operation.
	AllowedEnvironmentTypeListResult
}

// ProjectEnvironmentTypesClientCreateOrUpdateResponse contains the response from method ProjectEnvironmentTypesClient.CreateOrUpdate.
type ProjectEnvironmentTypesClientCreateOrUpdateResponse struct {
	// Represents an environment type.
	ProjectEnvironmentType
}

// ProjectEnvironmentTypesClientDeleteResponse contains the response from method ProjectEnvironmentTypesClient.Delete.
type ProjectEnvironmentTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectEnvironmentTypesClientGetResponse contains the response from method ProjectEnvironmentTypesClient.Get.
type ProjectEnvironmentTypesClientGetResponse struct {
	// Represents an environment type.
	ProjectEnvironmentType
}

// ProjectEnvironmentTypesClientListResponse contains the response from method ProjectEnvironmentTypesClient.NewListPager.
type ProjectEnvironmentTypesClientListResponse struct {
	// Result of the project environment type list operation.
	ProjectEnvironmentTypeListResult
}

// ProjectEnvironmentTypesClientUpdateResponse contains the response from method ProjectEnvironmentTypesClient.Update.
type ProjectEnvironmentTypesClientUpdateResponse struct {
	// Represents an environment type.
	ProjectEnvironmentType
}

// ProjectsClientCreateOrUpdateResponse contains the response from method ProjectsClient.BeginCreateOrUpdate.
type ProjectsClientCreateOrUpdateResponse struct {
	// Represents a project resource.
	Project
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.BeginDelete.
type ProjectsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	// Represents a project resource.
	Project
}

// ProjectsClientListByResourceGroupResponse contains the response from method ProjectsClient.NewListByResourceGroupPager.
type ProjectsClientListByResourceGroupResponse struct {
	// Results of the project list operation.
	ProjectListResult
}

// ProjectsClientListBySubscriptionResponse contains the response from method ProjectsClient.NewListBySubscriptionPager.
type ProjectsClientListBySubscriptionResponse struct {
	// Results of the project list operation.
	ProjectListResult
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.BeginUpdate.
type ProjectsClientUpdateResponse struct {
	// Represents a project resource.
	Project
}

// SKUsClientListBySubscriptionResponse contains the response from method SKUsClient.NewListBySubscriptionPager.
type SKUsClientListBySubscriptionResponse struct {
	// Results of the Microsoft.DevCenter SKU list operation.
	SKUListResult
}

// SchedulesClientCreateOrUpdateResponse contains the response from method SchedulesClient.BeginCreateOrUpdate.
type SchedulesClientCreateOrUpdateResponse struct {
	// Represents a Schedule to execute a task.
	Schedule
}

// SchedulesClientDeleteResponse contains the response from method SchedulesClient.BeginDelete.
type SchedulesClientDeleteResponse struct {
	// placeholder for future response values
}

// SchedulesClientGetResponse contains the response from method SchedulesClient.Get.
type SchedulesClientGetResponse struct {
	// Represents a Schedule to execute a task.
	Schedule
}

// SchedulesClientListByPoolResponse contains the response from method SchedulesClient.NewListByPoolPager.
type SchedulesClientListByPoolResponse struct {
	// Result of the schedule list operation.
	ScheduleListResult
}

// SchedulesClientUpdateResponse contains the response from method SchedulesClient.BeginUpdate.
type SchedulesClientUpdateResponse struct {
	// Represents a Schedule to execute a task.
	Schedule
}

// UsagesClientListByLocationResponse contains the response from method UsagesClient.NewListByLocationPager.
type UsagesClientListByLocationResponse struct {
	// List of Core Usages.
	ListUsagesResult
}
