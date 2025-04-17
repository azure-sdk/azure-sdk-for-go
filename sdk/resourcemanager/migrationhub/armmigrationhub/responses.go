// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationhub

// DatabaseInstancesControllerClientGetDatabaseInstanceResponse contains the response from method DatabaseInstancesControllerClient.GetDatabaseInstance.
type DatabaseInstancesControllerClientGetDatabaseInstanceResponse struct {
	// DatabaseInstance REST resource.
	DatabaseInstance
}

// DatabaseInstancesControllerClientListDatabaseInstancesResponse contains the response from method DatabaseInstancesControllerClient.NewListDatabaseInstancesPager.
type DatabaseInstancesControllerClientListDatabaseInstancesResponse struct {
	// Collection of database instances.
	DatabaseInstanceCollection
}

// DatabasesControllerClientGetDatabaseResponse contains the response from method DatabasesControllerClient.GetDatabase.
type DatabasesControllerClientGetDatabaseResponse struct {
	// Database REST resource.
	Database
}

// DatabasesControllerClientListDatabasesResponse contains the response from method DatabasesControllerClient.NewListDatabasesPager.
type DatabasesControllerClientListDatabasesResponse struct {
	// Collection of databases.
	DatabaseCollection
}

// EventsControllerClientDeleteResponse contains the response from method EventsControllerClient.Delete.
type EventsControllerClientDeleteResponse struct {
	// Anything
	Interface any
}

// EventsControllerClientGetEventResponse contains the response from method EventsControllerClient.GetEvent.
type EventsControllerClientGetEventResponse struct {
	// MigrateEvent REST resource.
	MigrateEvent
}

// EventsControllerClientListEventsResponse contains the response from method EventsControllerClient.NewListEventsPager.
type EventsControllerClientListEventsResponse struct {
	// Collection of events.
	EventCollection
}

// MachinesControllerClientGetMachineResponse contains the response from method MachinesControllerClient.GetMachine.
type MachinesControllerClientGetMachineResponse struct {
	// Machine REST resource.
	Machine
}

// MachinesControllerClientListMachinesResponse contains the response from method MachinesControllerClient.NewListMachinesPager.
type MachinesControllerClientListMachinesResponse struct {
	// Collection of machines.
	MachineCollection
}

// MigrateProjectsControllerClientDeleteMigrateProjectResponse contains the response from method MigrateProjectsControllerClient.DeleteMigrateProject.
type MigrateProjectsControllerClientDeleteMigrateProjectResponse struct {
	// placeholder for future response values
}

// MigrateProjectsControllerClientGetMigrateProjectResponse contains the response from method MigrateProjectsControllerClient.GetMigrateProject.
type MigrateProjectsControllerClientGetMigrateProjectResponse struct {
	// Migrate project.
	MigrateProject
}

// MigrateProjectsControllerClientGetToolRegistrationDetailsResponse contains the response from method MigrateProjectsControllerClient.GetToolRegistrationDetails.
type MigrateProjectsControllerClientGetToolRegistrationDetailsResponse struct {
	RegistrationDetailsResponse
}

// MigrateProjectsControllerClientPatchMigrateProjectResponse contains the response from method MigrateProjectsControllerClient.PatchMigrateProject.
type MigrateProjectsControllerClientPatchMigrateProjectResponse struct {
	// Migrate project.
	MigrateProject
}

// MigrateProjectsControllerClientPutMigrateProjectResponse contains the response from method MigrateProjectsControllerClient.PutMigrateProject.
type MigrateProjectsControllerClientPutMigrateProjectResponse struct {
	// Migrate project.
	MigrateProject

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// MigrateProjectsControllerClientRefreshSummaryResponse contains the response from method MigrateProjectsControllerClient.RefreshSummary.
type MigrateProjectsControllerClientRefreshSummaryResponse struct {
	// Class representing the refresh summary status of the migrate project.
	RefreshSummaryResult
}

// MigrateProjectsControllerClientRegisterToolResponse contains the response from method MigrateProjectsControllerClient.RegisterTool.
type MigrateProjectsControllerClientRegisterToolResponse struct {
	// Class representing the registration status of a tool with the migrate project.
	RegistrationResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PrivateEndpointConnectionControllerClientDeletePrivateEndpointConnectionResponse contains the response from method PrivateEndpointConnectionControllerClient.DeletePrivateEndpointConnection.
type PrivateEndpointConnectionControllerClientDeletePrivateEndpointConnectionResponse struct {
	Value *Enum
}

// PrivateEndpointConnectionControllerClientGetPrivateEndpointConnectionResponse contains the response from method PrivateEndpointConnectionControllerClient.GetPrivateEndpointConnection.
type PrivateEndpointConnectionControllerClientGetPrivateEndpointConnectionResponse struct {
	// REST model used to encapsulate the user visible state of a PrivateEndpoint.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionControllerClientPutPrivateEndpointConnectionResponse contains the response from method PrivateEndpointConnectionControllerClient.PutPrivateEndpointConnection.
type PrivateEndpointConnectionControllerClientPutPrivateEndpointConnectionResponse struct {
	// REST model used to encapsulate the user visible state of a PrivateEndpoint.
	PrivateEndpointConnection

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// PrivateEndpointConnectionProxyControllerClientCreateResponse contains the response from method PrivateEndpointConnectionProxyControllerClient.Create.
type PrivateEndpointConnectionProxyControllerClientCreateResponse struct {
	// Defines Private endpoint proxy resource.
	PrivateEndpointConnectionProxy

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// PrivateEndpointConnectionProxyControllerClientDeleteResponse contains the response from method PrivateEndpointConnectionProxyControllerClient.Delete.
type PrivateEndpointConnectionProxyControllerClientDeleteResponse struct {
	// Anything
	Interface any
}

// PrivateEndpointConnectionProxyControllerClientGetResponse contains the response from method PrivateEndpointConnectionProxyControllerClient.Get.
type PrivateEndpointConnectionProxyControllerClientGetResponse struct {
	// Defines Private endpoint proxy resource.
	PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxyControllerClientListPrivateEndpointConnectionProxiesResponse contains the response from method
// PrivateEndpointConnectionProxyControllerClient.NewListPrivateEndpointConnectionProxiesPager.
type PrivateEndpointConnectionProxyControllerClientListPrivateEndpointConnectionProxiesResponse struct {
	// Collection of PrivateLink proxy resources.
	PrivateEndpointConnectionProxyCollection
}

// PrivateEndpointConnectionProxyControllerClientValidateResponse contains the response from method PrivateEndpointConnectionProxyControllerClient.Validate.
type PrivateEndpointConnectionProxyControllerClientValidateResponse struct {
	// Defines Private endpoint proxy resource.
	PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionsControllerClientGetPrivateEndpointConnectionsResponse contains the response from method PrivateEndpointConnectionsControllerClient.NewGetPrivateEndpointConnectionsPager.
type PrivateEndpointConnectionsControllerClientGetPrivateEndpointConnectionsResponse struct {
	// Collection of PrivateLink resources.
	PrivateEndpointConnectionCollection
}

// PrivateLinkResourceControllerClientGetPrivateLinkResourceResponse contains the response from method PrivateLinkResourceControllerClient.GetPrivateLinkResource.
type PrivateLinkResourceControllerClientGetPrivateLinkResourceResponse struct {
	// Private link resource.
	PrivateLinkResource
}

// PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse contains the response from method PrivateLinkResourceControllerClient.NewGetPrivateLinkResourcesPager.
type PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse struct {
	// Collection of private link resources.
	PrivateLinkResourceCollection
}

// ProjectsClientListBySubscriptionResponse contains the response from method ProjectsClient.NewListBySubscriptionPager.
type ProjectsClientListBySubscriptionResponse struct {
	// List of projects.
	ProjectResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientListResponse contains the response from method ProjectsClient.NewListPager.
type ProjectsClientListResponse struct {
	// List of projects.
	ProjectResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// SolutionsControllerClientCleanupDataResponse contains the response from method SolutionsControllerClient.CleanupData.
type SolutionsControllerClientCleanupDataResponse struct {
	// Anything
	Interface any
}

// SolutionsControllerClientCreateResponse contains the response from method SolutionsControllerClient.Create.
type SolutionsControllerClientCreateResponse struct {
	// Solution REST Resource.
	Solution

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// SolutionsControllerClientDeleteSolutionResponse contains the response from method SolutionsControllerClient.DeleteSolution.
type SolutionsControllerClientDeleteSolutionResponse struct {
	// Anything
	Interface any
}

// SolutionsControllerClientGetConfigResponse contains the response from method SolutionsControllerClient.GetConfig.
type SolutionsControllerClientGetConfigResponse struct {
	// Class representing the config for the solution in the migrate project.
	SolutionConfig
}

// SolutionsControllerClientGetSolutionResponse contains the response from method SolutionsControllerClient.GetSolution.
type SolutionsControllerClientGetSolutionResponse struct {
	// Solution REST Resource.
	Solution
}

// SolutionsControllerClientListSolutionsResponse contains the response from method SolutionsControllerClient.NewListSolutionsPager.
type SolutionsControllerClientListSolutionsResponse struct {
	// Collection of solutions.
	SolutionsCollection
}

// SolutionsControllerClientUpdateResponse contains the response from method SolutionsControllerClient.Update.
type SolutionsControllerClientUpdateResponse struct {
	// Solution REST Resource.
	Solution
}

// VirtualDesktopUserControllerClientGetVirtualDesktopUserResponse contains the response from method VirtualDesktopUserControllerClient.GetVirtualDesktopUser.
type VirtualDesktopUserControllerClientGetVirtualDesktopUserResponse struct {
	// Class representing virtual desktop user.
	VirtualDesktopUser
}

// VirtualDesktopUserControllerClientListVirtualDesktopUsersResponse contains the response from method VirtualDesktopUserControllerClient.NewListVirtualDesktopUsersPager.
type VirtualDesktopUserControllerClientListVirtualDesktopUsersResponse struct {
	// Collection of virtual desktop users.
	VirtualDesktopUserCollection
}

// WebServersControllerClientGetWebServerResponse contains the response from method WebServersControllerClient.GetWebServer.
type WebServersControllerClientGetWebServerResponse struct {
	// Class representing a web server.
	WebServer
}

// WebServersControllerClientListWebServersResponse contains the response from method WebServersControllerClient.NewListWebServersPager.
type WebServersControllerClientListWebServersResponse struct {
	// Collection of web servers.
	WebServerCollection
}

// WebSitesControllerClientGetWebSiteResponse contains the response from method WebSitesControllerClient.GetWebSite.
type WebSitesControllerClientGetWebSiteResponse struct {
	// Class representing a web site.
	WebSite
}

// WebSitesControllerClientListWebSitesResponse contains the response from method WebSitesControllerClient.NewListWebSitesPager.
type WebSitesControllerClientListWebSitesResponse struct {
	// Collection of web sites.
	WebSiteCollection
}
