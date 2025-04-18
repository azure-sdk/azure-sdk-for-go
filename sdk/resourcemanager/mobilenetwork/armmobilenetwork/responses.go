// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

// AttachedDataNetworksClientCreateOrUpdateResponse contains the response from method AttachedDataNetworksClient.BeginCreateOrUpdate.
type AttachedDataNetworksClientCreateOrUpdateResponse struct {
	// Attached data network resource. Must be created in the same location as its parent packet core data plane.
	AttachedDataNetwork
}

// AttachedDataNetworksClientDeleteResponse contains the response from method AttachedDataNetworksClient.BeginDelete.
type AttachedDataNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// AttachedDataNetworksClientGetResponse contains the response from method AttachedDataNetworksClient.Get.
type AttachedDataNetworksClientGetResponse struct {
	// Attached data network resource. Must be created in the same location as its parent packet core data plane.
	AttachedDataNetwork
}

// AttachedDataNetworksClientListByPacketCoreDataPlaneResponse contains the response from method AttachedDataNetworksClient.NewListByPacketCoreDataPlanePager.
type AttachedDataNetworksClientListByPacketCoreDataPlaneResponse struct {
	// Response for attached data network API service call.
	AttachedDataNetworkListResult
}

// AttachedDataNetworksClientUpdateTagsResponse contains the response from method AttachedDataNetworksClient.UpdateTags.
type AttachedDataNetworksClientUpdateTagsResponse struct {
	// Attached data network resource. Must be created in the same location as its parent packet core data plane.
	AttachedDataNetwork
}

// DataNetworksClientCreateOrUpdateResponse contains the response from method DataNetworksClient.BeginCreateOrUpdate.
type DataNetworksClientCreateOrUpdateResponse struct {
	// Data network resource. Must be created in the same location as its parent mobile network.
	DataNetwork
}

// DataNetworksClientDeleteResponse contains the response from method DataNetworksClient.BeginDelete.
type DataNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// DataNetworksClientGetResponse contains the response from method DataNetworksClient.Get.
type DataNetworksClientGetResponse struct {
	// Data network resource. Must be created in the same location as its parent mobile network.
	DataNetwork
}

// DataNetworksClientListByMobileNetworkResponse contains the response from method DataNetworksClient.NewListByMobileNetworkPager.
type DataNetworksClientListByMobileNetworkResponse struct {
	// Response for data network API service call.
	DataNetworkListResult
}

// DataNetworksClientUpdateTagsResponse contains the response from method DataNetworksClient.UpdateTags.
type DataNetworksClientUpdateTagsResponse struct {
	// Data network resource. Must be created in the same location as its parent mobile network.
	DataNetwork
}

// DiagnosticsPackagesClientCreateOrUpdateResponse contains the response from method DiagnosticsPackagesClient.BeginCreateOrUpdate.
type DiagnosticsPackagesClientCreateOrUpdateResponse struct {
	// Diagnostics package resource.
	DiagnosticsPackage
}

// DiagnosticsPackagesClientDeleteResponse contains the response from method DiagnosticsPackagesClient.BeginDelete.
type DiagnosticsPackagesClientDeleteResponse struct {
	// placeholder for future response values
}

// DiagnosticsPackagesClientGetResponse contains the response from method DiagnosticsPackagesClient.Get.
type DiagnosticsPackagesClientGetResponse struct {
	// Diagnostics package resource.
	DiagnosticsPackage
}

// DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse contains the response from method DiagnosticsPackagesClient.NewListByPacketCoreControlPlanePager.
type DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse struct {
	// Response for diagnostics package API service call.
	DiagnosticsPackageListResult
}

// ExtendedUeInformationClientGetResponse contains the response from method ExtendedUeInformationClient.Get.
type ExtendedUeInformationClientGetResponse struct {
	// Extended User Equipment (UE) information.
	ExtendedUeInfo
}

// MobileNetworksClientCreateOrUpdateResponse contains the response from method MobileNetworksClient.BeginCreateOrUpdate.
type MobileNetworksClientCreateOrUpdateResponse struct {
	// Mobile network resource.
	MobileNetwork
}

// MobileNetworksClientDeleteResponse contains the response from method MobileNetworksClient.BeginDelete.
type MobileNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// MobileNetworksClientGetResponse contains the response from method MobileNetworksClient.Get.
type MobileNetworksClientGetResponse struct {
	// Mobile network resource.
	MobileNetwork
}

// MobileNetworksClientListByResourceGroupResponse contains the response from method MobileNetworksClient.NewListByResourceGroupPager.
type MobileNetworksClientListByResourceGroupResponse struct {
	// Response for mobile networks API service call.
	ListResult
}

// MobileNetworksClientListBySubscriptionResponse contains the response from method MobileNetworksClient.NewListBySubscriptionPager.
type MobileNetworksClientListBySubscriptionResponse struct {
	// Response for mobile networks API service call.
	ListResult
}

// MobileNetworksClientListSimGroupsResponse contains the response from method MobileNetworksClient.NewListSimGroupsPager.
type MobileNetworksClientListSimGroupsResponse struct {
	// Response for list SIM groups API service call.
	SimGroupListResult
}

// MobileNetworksClientUpdateTagsResponse contains the response from method MobileNetworksClient.UpdateTags.
type MobileNetworksClientUpdateTagsResponse struct {
	// Mobile network resource.
	MobileNetwork
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// List of the operations.
	OperationList
}

// PacketCapturesClientCreateOrUpdateResponse contains the response from method PacketCapturesClient.BeginCreateOrUpdate.
type PacketCapturesClientCreateOrUpdateResponse struct {
	// Packet capture session resource.
	PacketCapture
}

// PacketCapturesClientDeleteResponse contains the response from method PacketCapturesClient.BeginDelete.
type PacketCapturesClientDeleteResponse struct {
	// placeholder for future response values
}

// PacketCapturesClientGetResponse contains the response from method PacketCapturesClient.Get.
type PacketCapturesClientGetResponse struct {
	// Packet capture session resource.
	PacketCapture
}

// PacketCapturesClientListByPacketCoreControlPlaneResponse contains the response from method PacketCapturesClient.NewListByPacketCoreControlPlanePager.
type PacketCapturesClientListByPacketCoreControlPlaneResponse struct {
	// Response for packet capture API service call.
	PacketCaptureListResult
}

// PacketCapturesClientStopResponse contains the response from method PacketCapturesClient.BeginStop.
type PacketCapturesClientStopResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// PacketCoreControlPlaneVersionsClientGetBySubscriptionResponse contains the response from method PacketCoreControlPlaneVersionsClient.GetBySubscription.
type PacketCoreControlPlaneVersionsClientGetBySubscriptionResponse struct {
	// Packet core control plane version resource.
	PacketCoreControlPlaneVersion
}

// PacketCoreControlPlaneVersionsClientGetResponse contains the response from method PacketCoreControlPlaneVersionsClient.Get.
type PacketCoreControlPlaneVersionsClientGetResponse struct {
	// Packet core control plane version resource.
	PacketCoreControlPlaneVersion
}

// PacketCoreControlPlaneVersionsClientListBySubscriptionResponse contains the response from method PacketCoreControlPlaneVersionsClient.NewListBySubscriptionPager.
type PacketCoreControlPlaneVersionsClientListBySubscriptionResponse struct {
	// Response for packet core control plane version API service call.
	PacketCoreControlPlaneVersionListResult
}

// PacketCoreControlPlaneVersionsClientListResponse contains the response from method PacketCoreControlPlaneVersionsClient.NewListPager.
type PacketCoreControlPlaneVersionsClientListResponse struct {
	// Response for packet core control plane version API service call.
	PacketCoreControlPlaneVersionListResult
}

// PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse contains the response from method PacketCoreControlPlanesClient.BeginCollectDiagnosticsPackage.
type PacketCoreControlPlanesClientCollectDiagnosticsPackageResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientCreateOrUpdateResponse contains the response from method PacketCoreControlPlanesClient.BeginCreateOrUpdate.
type PacketCoreControlPlanesClientCreateOrUpdateResponse struct {
	// Packet core control plane resource.
	PacketCoreControlPlane
}

// PacketCoreControlPlanesClientDeleteResponse contains the response from method PacketCoreControlPlanesClient.BeginDelete.
type PacketCoreControlPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// PacketCoreControlPlanesClientGetResponse contains the response from method PacketCoreControlPlanesClient.Get.
type PacketCoreControlPlanesClientGetResponse struct {
	// Packet core control plane resource.
	PacketCoreControlPlane
}

// PacketCoreControlPlanesClientListByResourceGroupResponse contains the response from method PacketCoreControlPlanesClient.NewListByResourceGroupPager.
type PacketCoreControlPlanesClientListByResourceGroupResponse struct {
	// Response for packet core control planes API service call.
	PacketCoreControlPlaneListResult
}

// PacketCoreControlPlanesClientListBySubscriptionResponse contains the response from method PacketCoreControlPlanesClient.NewListBySubscriptionPager.
type PacketCoreControlPlanesClientListBySubscriptionResponse struct {
	// Response for packet core control planes API service call.
	PacketCoreControlPlaneListResult
}

// PacketCoreControlPlanesClientReinstallResponse contains the response from method PacketCoreControlPlanesClient.BeginReinstall.
type PacketCoreControlPlanesClientReinstallResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientRollbackResponse contains the response from method PacketCoreControlPlanesClient.BeginRollback.
type PacketCoreControlPlanesClientRollbackResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// PacketCoreControlPlanesClientUpdateTagsResponse contains the response from method PacketCoreControlPlanesClient.UpdateTags.
type PacketCoreControlPlanesClientUpdateTagsResponse struct {
	// Packet core control plane resource.
	PacketCoreControlPlane
}

// PacketCoreDataPlanesClientCreateOrUpdateResponse contains the response from method PacketCoreDataPlanesClient.BeginCreateOrUpdate.
type PacketCoreDataPlanesClientCreateOrUpdateResponse struct {
	// Packet core data plane resource. Must be created in the same location as its parent packet core control plane.
	PacketCoreDataPlane
}

// PacketCoreDataPlanesClientDeleteResponse contains the response from method PacketCoreDataPlanesClient.BeginDelete.
type PacketCoreDataPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// PacketCoreDataPlanesClientGetResponse contains the response from method PacketCoreDataPlanesClient.Get.
type PacketCoreDataPlanesClientGetResponse struct {
	// Packet core data plane resource. Must be created in the same location as its parent packet core control plane.
	PacketCoreDataPlane
}

// PacketCoreDataPlanesClientListByPacketCoreControlPlaneResponse contains the response from method PacketCoreDataPlanesClient.NewListByPacketCoreControlPlanePager.
type PacketCoreDataPlanesClientListByPacketCoreControlPlaneResponse struct {
	// Response for packet core data planes API service call.
	PacketCoreDataPlaneListResult
}

// PacketCoreDataPlanesClientUpdateTagsResponse contains the response from method PacketCoreDataPlanesClient.UpdateTags.
type PacketCoreDataPlanesClientUpdateTagsResponse struct {
	// Packet core data plane resource. Must be created in the same location as its parent packet core control plane.
	PacketCoreDataPlane
}

// RoutingInfoClientGetResponse contains the response from method RoutingInfoClient.Get.
type RoutingInfoClientGetResponse struct {
	// Routing information
	RoutingInfoModel
}

// RoutingInfoClientListResponse contains the response from method RoutingInfoClient.NewListPager.
type RoutingInfoClientListResponse struct {
	// Response for the list routing information API service call.
	RoutingInfoListResult
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// Service resource. Must be created in the same location as its parent mobile network.
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// Service resource. Must be created in the same location as its parent mobile network.
	Service
}

// ServicesClientListByMobileNetworkResponse contains the response from method ServicesClient.NewListByMobileNetworkPager.
type ServicesClientListByMobileNetworkResponse struct {
	// Response for services API service call.
	ServiceListResult
}

// ServicesClientUpdateTagsResponse contains the response from method ServicesClient.UpdateTags.
type ServicesClientUpdateTagsResponse struct {
	// Service resource. Must be created in the same location as its parent mobile network.
	Service
}

// SimGroupsClientCreateOrUpdateResponse contains the response from method SimGroupsClient.BeginCreateOrUpdate.
type SimGroupsClientCreateOrUpdateResponse struct {
	// SIM group resource.
	SimGroup
}

// SimGroupsClientDeleteResponse contains the response from method SimGroupsClient.BeginDelete.
type SimGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// SimGroupsClientGetResponse contains the response from method SimGroupsClient.Get.
type SimGroupsClientGetResponse struct {
	// SIM group resource.
	SimGroup
}

// SimGroupsClientListByResourceGroupResponse contains the response from method SimGroupsClient.NewListByResourceGroupPager.
type SimGroupsClientListByResourceGroupResponse struct {
	// Response for list SIM groups API service call.
	SimGroupListResult
}

// SimGroupsClientListBySubscriptionResponse contains the response from method SimGroupsClient.NewListBySubscriptionPager.
type SimGroupsClientListBySubscriptionResponse struct {
	// Response for list SIM groups API service call.
	SimGroupListResult
}

// SimGroupsClientUpdateTagsResponse contains the response from method SimGroupsClient.UpdateTags.
type SimGroupsClientUpdateTagsResponse struct {
	// SIM group resource.
	SimGroup
}

// SimPoliciesClientCreateOrUpdateResponse contains the response from method SimPoliciesClient.BeginCreateOrUpdate.
type SimPoliciesClientCreateOrUpdateResponse struct {
	// SIM policy resource.
	SimPolicy
}

// SimPoliciesClientDeleteResponse contains the response from method SimPoliciesClient.BeginDelete.
type SimPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// SimPoliciesClientGetResponse contains the response from method SimPoliciesClient.Get.
type SimPoliciesClientGetResponse struct {
	// SIM policy resource.
	SimPolicy
}

// SimPoliciesClientListByMobileNetworkResponse contains the response from method SimPoliciesClient.NewListByMobileNetworkPager.
type SimPoliciesClientListByMobileNetworkResponse struct {
	// Response for SIM policies API service call.
	SimPolicyListResult
}

// SimPoliciesClientUpdateTagsResponse contains the response from method SimPoliciesClient.UpdateTags.
type SimPoliciesClientUpdateTagsResponse struct {
	// SIM policy resource.
	SimPolicy
}

// SimsClientBulkDeleteResponse contains the response from method SimsClient.BeginBulkDelete.
type SimsClientBulkDeleteResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// SimsClientBulkUploadEncryptedResponse contains the response from method SimsClient.BeginBulkUploadEncrypted.
type SimsClientBulkUploadEncryptedResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// SimsClientBulkUploadResponse contains the response from method SimsClient.BeginBulkUpload.
type SimsClientBulkUploadResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// SimsClientCloneResponse contains the response from method SimsClient.BeginClone.
type SimsClientCloneResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// SimsClientCreateOrUpdateResponse contains the response from method SimsClient.BeginCreateOrUpdate.
type SimsClientCreateOrUpdateResponse struct {
	// SIM resource.
	Sim
}

// SimsClientDeleteResponse contains the response from method SimsClient.BeginDelete.
type SimsClientDeleteResponse struct {
	// placeholder for future response values
}

// SimsClientGetResponse contains the response from method SimsClient.Get.
type SimsClientGetResponse struct {
	// SIM resource.
	Sim
}

// SimsClientListByGroupResponse contains the response from method SimsClient.NewListByGroupPager.
type SimsClientListByGroupResponse struct {
	// Response for list SIMs API service call.
	SimListResult
}

// SimsClientMoveResponse contains the response from method SimsClient.BeginMove.
type SimsClientMoveResponse struct {
	// The current status of an async operation.
	AsyncOperationStatus
}

// SitesClientCreateOrUpdateResponse contains the response from method SitesClient.BeginCreateOrUpdate.
type SitesClientCreateOrUpdateResponse struct {
	// Site resource. Must be created in the same location as its parent mobile network.
	Site
}

// SitesClientDeletePacketCoreResponse contains the response from method SitesClient.BeginDeletePacketCore.
type SitesClientDeletePacketCoreResponse struct {
	// placeholder for future response values
}

// SitesClientDeleteResponse contains the response from method SitesClient.BeginDelete.
type SitesClientDeleteResponse struct {
	// placeholder for future response values
}

// SitesClientGetResponse contains the response from method SitesClient.Get.
type SitesClientGetResponse struct {
	// Site resource. Must be created in the same location as its parent mobile network.
	Site
}

// SitesClientListByMobileNetworkResponse contains the response from method SitesClient.NewListByMobileNetworkPager.
type SitesClientListByMobileNetworkResponse struct {
	// Response for sites API service call.
	SiteListResult
}

// SitesClientUpdateTagsResponse contains the response from method SitesClient.UpdateTags.
type SitesClientUpdateTagsResponse struct {
	// Site resource. Must be created in the same location as its parent mobile network.
	Site
}

// SlicesClientCreateOrUpdateResponse contains the response from method SlicesClient.BeginCreateOrUpdate.
type SlicesClientCreateOrUpdateResponse struct {
	// Network slice resource. Must be created in the same location as its parent mobile network.
	Slice
}

// SlicesClientDeleteResponse contains the response from method SlicesClient.BeginDelete.
type SlicesClientDeleteResponse struct {
	// placeholder for future response values
}

// SlicesClientGetResponse contains the response from method SlicesClient.Get.
type SlicesClientGetResponse struct {
	// Network slice resource. Must be created in the same location as its parent mobile network.
	Slice
}

// SlicesClientListByMobileNetworkResponse contains the response from method SlicesClient.NewListByMobileNetworkPager.
type SlicesClientListByMobileNetworkResponse struct {
	// Response for network slice API service call.
	SliceListResult
}

// SlicesClientUpdateTagsResponse contains the response from method SlicesClient.UpdateTags.
type SlicesClientUpdateTagsResponse struct {
	// Network slice resource. Must be created in the same location as its parent mobile network.
	Slice
}

// UeInformationClientListResponse contains the response from method UeInformationClient.NewListPager.
type UeInformationClientListResponse struct {
	// Response for packet core list UEs API call.
	UeInfoList
}
