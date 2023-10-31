//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

// AppAttachPackageClientCreateOrUpdateOptions contains the optional parameters for the AppAttachPackageClient.CreateOrUpdate
// method.
type AppAttachPackageClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AppAttachPackageClientDeleteOptions contains the optional parameters for the AppAttachPackageClient.Delete method.
type AppAttachPackageClientDeleteOptions struct {
	// Force flag to delete App Attach package.
	Force *bool
}

// AppAttachPackageClientGetOptions contains the optional parameters for the AppAttachPackageClient.Get method.
type AppAttachPackageClientGetOptions struct {
	// placeholder for future optional parameters
}

// AppAttachPackageClientListByResourceGroupOptions contains the optional parameters for the AppAttachPackageClient.NewListByResourceGroupPager
// method.
type AppAttachPackageClientListByResourceGroupOptions struct {
	// OData filter expression. Valid properties for filtering are package name and host pool.
	Filter *string
}

// AppAttachPackageClientListBySubscriptionOptions contains the optional parameters for the AppAttachPackageClient.NewListBySubscriptionPager
// method.
type AppAttachPackageClientListBySubscriptionOptions struct {
	// OData filter expression. Valid properties for filtering are package name, host pool, and resource group.
	Filter *string
}

// AppAttachPackageClientUpdateOptions contains the optional parameters for the AppAttachPackageClient.Update method.
type AppAttachPackageClientUpdateOptions struct {
	// Object containing App Attach Package definition.
	AppAttachPackagePatch *AppAttachPackagePatch
}

// AppAttachPackageInfoClientImportOptions contains the optional parameters for the AppAttachPackageInfoClient.NewImportPager
// method.
type AppAttachPackageInfoClientImportOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupsClientCreateOrUpdateOptions contains the optional parameters for the ApplicationGroupsClient.CreateOrUpdate
// method.
type ApplicationGroupsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupsClientDeleteOptions contains the optional parameters for the ApplicationGroupsClient.Delete method.
type ApplicationGroupsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupsClientGetOptions contains the optional parameters for the ApplicationGroupsClient.Get method.
type ApplicationGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationGroupsClientListByResourceGroupOptions contains the optional parameters for the ApplicationGroupsClient.NewListByResourceGroupPager
// method.
type ApplicationGroupsClientListByResourceGroupOptions struct {
	// OData filter expression. Valid properties for filtering are applicationGroupType.
	Filter *string

	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ApplicationGroupsClientListBySubscriptionOptions contains the optional parameters for the ApplicationGroupsClient.NewListBySubscriptionPager
// method.
type ApplicationGroupsClientListBySubscriptionOptions struct {
	// OData filter expression. Valid properties for filtering are applicationGroupType.
	Filter *string
}

// ApplicationGroupsClientUpdateOptions contains the optional parameters for the ApplicationGroupsClient.Update method.
type ApplicationGroupsClientUpdateOptions struct {
	// Object containing ApplicationGroup definitions.
	ApplicationGroup *ApplicationGroupPatch
}

// ApplicationsClientCreateOrUpdateOptions contains the optional parameters for the ApplicationsClient.CreateOrUpdate method.
type ApplicationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsClientDeleteOptions contains the optional parameters for the ApplicationsClient.Delete method.
type ApplicationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsClientGetOptions contains the optional parameters for the ApplicationsClient.Get method.
type ApplicationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ApplicationsClientListOptions contains the optional parameters for the ApplicationsClient.NewListPager method.
type ApplicationsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ApplicationsClientUpdateOptions contains the optional parameters for the ApplicationsClient.Update method.
type ApplicationsClientUpdateOptions struct {
	// Object containing Application definitions.
	Application *ApplicationPatch
}

// DesktopsClientGetOptions contains the optional parameters for the DesktopsClient.Get method.
type DesktopsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DesktopsClientListOptions contains the optional parameters for the DesktopsClient.NewListPager method.
type DesktopsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// DesktopsClientUpdateOptions contains the optional parameters for the DesktopsClient.Update method.
type DesktopsClientUpdateOptions struct {
	// Object containing Desktop definitions.
	Desktop *DesktopPatch
}

// HostPoolsClientCreateOrUpdateOptions contains the optional parameters for the HostPoolsClient.CreateOrUpdate method.
type HostPoolsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// HostPoolsClientDeleteOptions contains the optional parameters for the HostPoolsClient.Delete method.
type HostPoolsClientDeleteOptions struct {
	// Force flag to delete sessionHost.
	Force *bool
}

// HostPoolsClientGetOptions contains the optional parameters for the HostPoolsClient.Get method.
type HostPoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// HostPoolsClientListByResourceGroupOptions contains the optional parameters for the HostPoolsClient.NewListByResourceGroupPager
// method.
type HostPoolsClientListByResourceGroupOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// HostPoolsClientListOptions contains the optional parameters for the HostPoolsClient.NewListPager method.
type HostPoolsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// HostPoolsClientRetrieveRegistrationTokenOptions contains the optional parameters for the HostPoolsClient.RetrieveRegistrationToken
// method.
type HostPoolsClientRetrieveRegistrationTokenOptions struct {
	// placeholder for future optional parameters
}

// HostPoolsClientUpdateOptions contains the optional parameters for the HostPoolsClient.Update method.
type HostPoolsClientUpdateOptions struct {
	// Object containing HostPool definitions.
	HostPool *HostPoolPatch
}

// MSIXPackagesClientCreateOrUpdateOptions contains the optional parameters for the MSIXPackagesClient.CreateOrUpdate method.
type MSIXPackagesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// MSIXPackagesClientDeleteOptions contains the optional parameters for the MSIXPackagesClient.Delete method.
type MSIXPackagesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MSIXPackagesClientGetOptions contains the optional parameters for the MSIXPackagesClient.Get method.
type MSIXPackagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// MSIXPackagesClientListOptions contains the optional parameters for the MSIXPackagesClient.NewListPager method.
type MSIXPackagesClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// MSIXPackagesClientUpdateOptions contains the optional parameters for the MSIXPackagesClient.Update method.
type MSIXPackagesClientUpdateOptions struct {
	// Object containing MSIX Package definitions.
	MsixPackage *MSIXPackagePatch
}

// MsixImagesClientExpandOptions contains the optional parameters for the MsixImagesClient.NewExpandPager method.
type MsixImagesClientExpandOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientDeleteByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.DeleteByHostPool
// method.
type PrivateEndpointConnectionsClientDeleteByHostPoolOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientDeleteByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.DeleteByWorkspace
// method.
type PrivateEndpointConnectionsClientDeleteByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.GetByHostPool
// method.
type PrivateEndpointConnectionsClientGetByHostPoolOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.GetByWorkspace
// method.
type PrivateEndpointConnectionsClientGetByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByHostPoolPager
// method.
type PrivateEndpointConnectionsClientListByHostPoolOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// PrivateEndpointConnectionsClientListByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByWorkspacePager
// method.
type PrivateEndpointConnectionsClientListByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientUpdateByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.UpdateByHostPool
// method.
type PrivateEndpointConnectionsClientUpdateByHostPoolOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientUpdateByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.UpdateByWorkspace
// method.
type PrivateEndpointConnectionsClientUpdateByWorkspaceOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByHostPoolOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByHostPoolPager
// method.
type PrivateLinkResourcesClientListByHostPoolOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// PrivateLinkResourcesClientListByWorkspaceOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByWorkspacePager
// method.
type PrivateLinkResourcesClientListByWorkspaceOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlanPersonalSchedulesClientCreateOptions contains the optional parameters for the ScalingPlanPersonalSchedulesClient.Create
// method.
type ScalingPlanPersonalSchedulesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPersonalSchedulesClientDeleteOptions contains the optional parameters for the ScalingPlanPersonalSchedulesClient.Delete
// method.
type ScalingPlanPersonalSchedulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPersonalSchedulesClientGetOptions contains the optional parameters for the ScalingPlanPersonalSchedulesClient.Get
// method.
type ScalingPlanPersonalSchedulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPersonalSchedulesClientListOptions contains the optional parameters for the ScalingPlanPersonalSchedulesClient.NewListPager
// method.
type ScalingPlanPersonalSchedulesClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlanPersonalSchedulesClientUpdateOptions contains the optional parameters for the ScalingPlanPersonalSchedulesClient.Update
// method.
type ScalingPlanPersonalSchedulesClientUpdateOptions struct {
	// Object containing ScalingPlanPersonalSchedule definitions.
	ScalingPlanSchedule *ScalingPlanPersonalSchedulePatch
}

// ScalingPlanPooledSchedulesClientCreateOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Create
// method.
type ScalingPlanPooledSchedulesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPooledSchedulesClientDeleteOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Delete
// method.
type ScalingPlanPooledSchedulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPooledSchedulesClientGetOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Get
// method.
type ScalingPlanPooledSchedulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlanPooledSchedulesClientListOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.NewListPager
// method.
type ScalingPlanPooledSchedulesClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlanPooledSchedulesClientUpdateOptions contains the optional parameters for the ScalingPlanPooledSchedulesClient.Update
// method.
type ScalingPlanPooledSchedulesClientUpdateOptions struct {
	// Object containing ScalingPlanPooledSchedule definitions.
	ScalingPlanSchedule *ScalingPlanPooledSchedulePatch
}

// ScalingPlansClientCreateOptions contains the optional parameters for the ScalingPlansClient.Create method.
type ScalingPlansClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlansClientDeleteOptions contains the optional parameters for the ScalingPlansClient.Delete method.
type ScalingPlansClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlansClientGetOptions contains the optional parameters for the ScalingPlansClient.Get method.
type ScalingPlansClientGetOptions struct {
	// placeholder for future optional parameters
}

// ScalingPlansClientListByHostPoolOptions contains the optional parameters for the ScalingPlansClient.NewListByHostPoolPager
// method.
type ScalingPlansClientListByHostPoolOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlansClientListByResourceGroupOptions contains the optional parameters for the ScalingPlansClient.NewListByResourceGroupPager
// method.
type ScalingPlansClientListByResourceGroupOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlansClientListBySubscriptionOptions contains the optional parameters for the ScalingPlansClient.NewListBySubscriptionPager
// method.
type ScalingPlansClientListBySubscriptionOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// ScalingPlansClientUpdateOptions contains the optional parameters for the ScalingPlansClient.Update method.
type ScalingPlansClientUpdateOptions struct {
	// Object containing scaling plan definitions.
	ScalingPlan *ScalingPlanPatch
}

// SessionHostsClientDeleteOptions contains the optional parameters for the SessionHostsClient.Delete method.
type SessionHostsClientDeleteOptions struct {
	// Force flag to force sessionHost deletion even when userSession exists.
	Force *bool
}

// SessionHostsClientGetOptions contains the optional parameters for the SessionHostsClient.Get method.
type SessionHostsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SessionHostsClientListOptions contains the optional parameters for the SessionHostsClient.NewListPager method.
type SessionHostsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// SessionHostsClientUpdateOptions contains the optional parameters for the SessionHostsClient.Update method.
type SessionHostsClientUpdateOptions struct {
	// Force flag to update assign, unassign or reassign personal desktop.
	Force *bool

	// Object containing SessionHost definitions.
	SessionHost *SessionHostPatch
}

// StartMenuItemsClientListOptions contains the optional parameters for the StartMenuItemsClient.NewListPager method.
type StartMenuItemsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// UserSessionsClientDeleteOptions contains the optional parameters for the UserSessionsClient.Delete method.
type UserSessionsClientDeleteOptions struct {
	// Force flag to login off userSession.
	Force *bool
}

// UserSessionsClientDisconnectOptions contains the optional parameters for the UserSessionsClient.Disconnect method.
type UserSessionsClientDisconnectOptions struct {
	// placeholder for future optional parameters
}

// UserSessionsClientGetOptions contains the optional parameters for the UserSessionsClient.Get method.
type UserSessionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// UserSessionsClientListByHostPoolOptions contains the optional parameters for the UserSessionsClient.NewListByHostPoolPager
// method.
type UserSessionsClientListByHostPoolOptions struct {
	// OData filter expression. Valid properties for filtering are userprincipalname and sessionstate.
	Filter *string

	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// UserSessionsClientListOptions contains the optional parameters for the UserSessionsClient.NewListPager method.
type UserSessionsClientListOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// UserSessionsClientSendMessageOptions contains the optional parameters for the UserSessionsClient.SendMessage method.
type UserSessionsClientSendMessageOptions struct {
	// Object containing message includes title and message body
	SendMessage *SendMessage
}

// WorkspacesClientCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.CreateOrUpdate method.
type WorkspacesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientDeleteOptions contains the optional parameters for the WorkspacesClient.Delete method.
type WorkspacesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
type WorkspacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.NewListByResourceGroupPager
// method.
type WorkspacesClientListByResourceGroupOptions struct {
	// Initial number of items to skip.
	InitialSkip *int32

	// Indicates whether the collection is descending.
	IsDescending *bool

	// Number of items per page.
	PageSize *int32
}

// WorkspacesClientListBySubscriptionOptions contains the optional parameters for the WorkspacesClient.NewListBySubscriptionPager
// method.
type WorkspacesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientUpdateOptions contains the optional parameters for the WorkspacesClient.Update method.
type WorkspacesClientUpdateOptions struct {
	// Object containing Workspace definitions.
	Workspace *WorkspacePatch
}
