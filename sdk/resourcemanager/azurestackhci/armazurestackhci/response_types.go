//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

// GalleryImagesClientCreateOrUpdateResponse contains the response from method GalleryImagesClient.BeginCreateOrUpdate.
type GalleryImagesClientCreateOrUpdateResponse struct {
	// The gallery images resource definition.
	GalleryImages
}

// GalleryImagesClientDeleteResponse contains the response from method GalleryImagesClient.BeginDelete.
type GalleryImagesClientDeleteResponse struct {
	// placeholder for future response values
}

// GalleryImagesClientGetResponse contains the response from method GalleryImagesClient.Get.
type GalleryImagesClientGetResponse struct {
	// The gallery images resource definition.
	GalleryImages
}

// GalleryImagesClientListAllResponse contains the response from method GalleryImagesClient.NewListAllPager.
type GalleryImagesClientListAllResponse struct {
	// List of gallery images.
	GalleryImagesListResult
}

// GalleryImagesClientListResponse contains the response from method GalleryImagesClient.NewListPager.
type GalleryImagesClientListResponse struct {
	// List of gallery images.
	GalleryImagesListResult
}

// GalleryImagesClientUpdateResponse contains the response from method GalleryImagesClient.BeginUpdate.
type GalleryImagesClientUpdateResponse struct {
	// The gallery images resource definition.
	GalleryImages
}

// GuestAgentClientCreateResponse contains the response from method GuestAgentClient.BeginCreate.
type GuestAgentClientCreateResponse struct {
	// Defines the GuestAgent.
	GuestAgent
}

// GuestAgentClientDeleteResponse contains the response from method GuestAgentClient.BeginDelete.
type GuestAgentClientDeleteResponse struct {
	// placeholder for future response values
}

// GuestAgentClientGetResponse contains the response from method GuestAgentClient.Get.
type GuestAgentClientGetResponse struct {
	// Defines the GuestAgent.
	GuestAgent
}

// GuestAgentsClientListResponse contains the response from method GuestAgentsClient.NewListPager.
type GuestAgentsClientListResponse struct {
	// List of GuestAgent.
	GuestAgentList
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	// Defines the HybridIdentityMetadata.
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListResponse contains the response from method HybridIdentityMetadataClient.NewListPager.
type HybridIdentityMetadataClientListResponse struct {
	// List of HybridIdentityMetadata.
	HybridIdentityMetadataList
}

// LogicalNetworksClientCreateOrUpdateResponse contains the response from method LogicalNetworksClient.BeginCreateOrUpdate.
type LogicalNetworksClientCreateOrUpdateResponse struct {
	// The logical network resource definition.
	LogicalNetworks
}

// LogicalNetworksClientDeleteResponse contains the response from method LogicalNetworksClient.BeginDelete.
type LogicalNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// LogicalNetworksClientGetResponse contains the response from method LogicalNetworksClient.Get.
type LogicalNetworksClientGetResponse struct {
	// The logical network resource definition.
	LogicalNetworks
}

// LogicalNetworksClientListAllResponse contains the response from method LogicalNetworksClient.NewListAllPager.
type LogicalNetworksClientListAllResponse struct {
	LogicalNetworksListResult
}

// LogicalNetworksClientListResponse contains the response from method LogicalNetworksClient.NewListPager.
type LogicalNetworksClientListResponse struct {
	LogicalNetworksListResult
}

// LogicalNetworksClientUpdateResponse contains the response from method LogicalNetworksClient.BeginUpdate.
type LogicalNetworksClientUpdateResponse struct {
	// The logical network resource definition.
	LogicalNetworks
}

// MarketplaceGalleryImagesClientCreateOrUpdateResponse contains the response from method MarketplaceGalleryImagesClient.BeginCreateOrUpdate.
type MarketplaceGalleryImagesClientCreateOrUpdateResponse struct {
	// The marketplace gallery image resource definition.
	MarketplaceGalleryImages
}

// MarketplaceGalleryImagesClientDeleteResponse contains the response from method MarketplaceGalleryImagesClient.BeginDelete.
type MarketplaceGalleryImagesClientDeleteResponse struct {
	// placeholder for future response values
}

// MarketplaceGalleryImagesClientGetResponse contains the response from method MarketplaceGalleryImagesClient.Get.
type MarketplaceGalleryImagesClientGetResponse struct {
	// The marketplace gallery image resource definition.
	MarketplaceGalleryImages
}

// MarketplaceGalleryImagesClientListAllResponse contains the response from method MarketplaceGalleryImagesClient.NewListAllPager.
type MarketplaceGalleryImagesClientListAllResponse struct {
	MarketplaceGalleryImagesListResult
}

// MarketplaceGalleryImagesClientListResponse contains the response from method MarketplaceGalleryImagesClient.NewListPager.
type MarketplaceGalleryImagesClientListResponse struct {
	MarketplaceGalleryImagesListResult
}

// MarketplaceGalleryImagesClientUpdateResponse contains the response from method MarketplaceGalleryImagesClient.BeginUpdate.
type MarketplaceGalleryImagesClientUpdateResponse struct {
	// The marketplace gallery image resource definition.
	MarketplaceGalleryImages
}

// NetworkInterfacesClientCreateOrUpdateResponse contains the response from method NetworkInterfacesClient.BeginCreateOrUpdate.
type NetworkInterfacesClientCreateOrUpdateResponse struct {
	// The network interface resource definition.
	NetworkInterfaces
}

// NetworkInterfacesClientDeleteResponse contains the response from method NetworkInterfacesClient.BeginDelete.
type NetworkInterfacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkInterfacesClientGetResponse contains the response from method NetworkInterfacesClient.Get.
type NetworkInterfacesClientGetResponse struct {
	// The network interface resource definition.
	NetworkInterfaces
}

// NetworkInterfacesClientListAllResponse contains the response from method NetworkInterfacesClient.NewListAllPager.
type NetworkInterfacesClientListAllResponse struct {
	NetworkInterfacesListResult
}

// NetworkInterfacesClientListResponse contains the response from method NetworkInterfacesClient.NewListPager.
type NetworkInterfacesClientListResponse struct {
	NetworkInterfacesListResult
}

// NetworkInterfacesClientUpdateResponse contains the response from method NetworkInterfacesClient.BeginUpdate.
type NetworkInterfacesClientUpdateResponse struct {
	// The network interface resource definition.
	NetworkInterfaces
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// StorageContainersClientCreateOrUpdateResponse contains the response from method StorageContainersClient.BeginCreateOrUpdate.
type StorageContainersClientCreateOrUpdateResponse struct {
	// The storage container resource definition.
	StorageContainers
}

// StorageContainersClientDeleteResponse contains the response from method StorageContainersClient.BeginDelete.
type StorageContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageContainersClientGetResponse contains the response from method StorageContainersClient.Get.
type StorageContainersClientGetResponse struct {
	// The storage container resource definition.
	StorageContainers
}

// StorageContainersClientListAllResponse contains the response from method StorageContainersClient.NewListAllPager.
type StorageContainersClientListAllResponse struct {
	StorageContainersListResult
}

// StorageContainersClientListResponse contains the response from method StorageContainersClient.NewListPager.
type StorageContainersClientListResponse struct {
	StorageContainersListResult
}

// StorageContainersClientUpdateResponse contains the response from method StorageContainersClient.BeginUpdate.
type StorageContainersClientUpdateResponse struct {
	// The storage container resource definition.
	StorageContainers
}

// VirtualHardDisksClientCreateOrUpdateResponse contains the response from method VirtualHardDisksClient.BeginCreateOrUpdate.
type VirtualHardDisksClientCreateOrUpdateResponse struct {
	// The virtual hard disk resource definition.
	VirtualHardDisks
}

// VirtualHardDisksClientDeleteResponse contains the response from method VirtualHardDisksClient.BeginDelete.
type VirtualHardDisksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualHardDisksClientGetResponse contains the response from method VirtualHardDisksClient.Get.
type VirtualHardDisksClientGetResponse struct {
	// The virtual hard disk resource definition.
	VirtualHardDisks
}

// VirtualHardDisksClientListAllResponse contains the response from method VirtualHardDisksClient.NewListAllPager.
type VirtualHardDisksClientListAllResponse struct {
	VirtualHardDisksListResult
}

// VirtualHardDisksClientListResponse contains the response from method VirtualHardDisksClient.NewListPager.
type VirtualHardDisksClientListResponse struct {
	VirtualHardDisksListResult
}

// VirtualHardDisksClientUpdateResponse contains the response from method VirtualHardDisksClient.BeginUpdate.
type VirtualHardDisksClientUpdateResponse struct {
	// The virtual hard disk resource definition.
	VirtualHardDisks
}

// VirtualMachineInstancesClientCreateOrUpdateResponse contains the response from method VirtualMachineInstancesClient.BeginCreateOrUpdate.
type VirtualMachineInstancesClientCreateOrUpdateResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}

// VirtualMachineInstancesClientDeleteResponse contains the response from method VirtualMachineInstancesClient.BeginDelete.
type VirtualMachineInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineInstancesClientGetResponse contains the response from method VirtualMachineInstancesClient.Get.
type VirtualMachineInstancesClientGetResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}

// VirtualMachineInstancesClientListResponse contains the response from method VirtualMachineInstancesClient.NewListPager.
type VirtualMachineInstancesClientListResponse struct {
	VirtualMachineInstanceListResult
}

// VirtualMachineInstancesClientRestartResponse contains the response from method VirtualMachineInstancesClient.BeginRestart.
type VirtualMachineInstancesClientRestartResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}

// VirtualMachineInstancesClientStartResponse contains the response from method VirtualMachineInstancesClient.BeginStart.
type VirtualMachineInstancesClientStartResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}

// VirtualMachineInstancesClientStopResponse contains the response from method VirtualMachineInstancesClient.BeginStop.
type VirtualMachineInstancesClientStopResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}

// VirtualMachineInstancesClientUpdateResponse contains the response from method VirtualMachineInstancesClient.BeginUpdate.
type VirtualMachineInstancesClientUpdateResponse struct {
	// The virtual machine instance resource definition.
	VirtualMachineInstance
}