//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAvailabilitySetsClient creates a new instance of AvailabilitySetsClient.
func (c *ClientFactory) NewAvailabilitySetsClient() *AvailabilitySetsClient {
	return &AvailabilitySetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCapacityReservationGroupsClient creates a new instance of CapacityReservationGroupsClient.
func (c *ClientFactory) NewCapacityReservationGroupsClient() *CapacityReservationGroupsClient {
	return &CapacityReservationGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCapacityReservationsClient creates a new instance of CapacityReservationsClient.
func (c *ClientFactory) NewCapacityReservationsClient() *CapacityReservationsClient {
	return &CapacityReservationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServiceOperatingSystemsClient creates a new instance of CloudServiceOperatingSystemsClient.
func (c *ClientFactory) NewCloudServiceOperatingSystemsClient() *CloudServiceOperatingSystemsClient {
	return &CloudServiceOperatingSystemsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServiceRoleInstancesClient creates a new instance of CloudServiceRoleInstancesClient.
func (c *ClientFactory) NewCloudServiceRoleInstancesClient() *CloudServiceRoleInstancesClient {
	return &CloudServiceRoleInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServiceRolesClient creates a new instance of CloudServiceRolesClient.
func (c *ClientFactory) NewCloudServiceRolesClient() *CloudServiceRolesClient {
	return &CloudServiceRolesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServicesClient creates a new instance of CloudServicesClient.
func (c *ClientFactory) NewCloudServicesClient() *CloudServicesClient {
	return &CloudServicesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCloudServicesUpdateDomainClient creates a new instance of CloudServicesUpdateDomainClient.
func (c *ClientFactory) NewCloudServicesUpdateDomainClient() *CloudServicesUpdateDomainClient {
	return &CloudServicesUpdateDomainClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCommunityGalleriesClient creates a new instance of CommunityGalleriesClient.
func (c *ClientFactory) NewCommunityGalleriesClient() *CommunityGalleriesClient {
	return &CommunityGalleriesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCommunityGalleryImageVersionsClient creates a new instance of CommunityGalleryImageVersionsClient.
func (c *ClientFactory) NewCommunityGalleryImageVersionsClient() *CommunityGalleryImageVersionsClient {
	return &CommunityGalleryImageVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCommunityGalleryImagesClient creates a new instance of CommunityGalleryImagesClient.
func (c *ClientFactory) NewCommunityGalleryImagesClient() *CommunityGalleryImagesClient {
	return &CommunityGalleryImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDedicatedHostGroupsClient creates a new instance of DedicatedHostGroupsClient.
func (c *ClientFactory) NewDedicatedHostGroupsClient() *DedicatedHostGroupsClient {
	return &DedicatedHostGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDedicatedHostsClient creates a new instance of DedicatedHostsClient.
func (c *ClientFactory) NewDedicatedHostsClient() *DedicatedHostsClient {
	return &DedicatedHostsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiskAccessesClient creates a new instance of DiskAccessesClient.
func (c *ClientFactory) NewDiskAccessesClient() *DiskAccessesClient {
	return &DiskAccessesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiskEncryptionSetsClient creates a new instance of DiskEncryptionSetsClient.
func (c *ClientFactory) NewDiskEncryptionSetsClient() *DiskEncryptionSetsClient {
	return &DiskEncryptionSetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDiskRestorePointClient creates a new instance of DiskRestorePointClient.
func (c *ClientFactory) NewDiskRestorePointClient() *DiskRestorePointClient {
	return &DiskRestorePointClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDisksClient creates a new instance of DisksClient.
func (c *ClientFactory) NewDisksClient() *DisksClient {
	return &DisksClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleriesClient creates a new instance of GalleriesClient.
func (c *ClientFactory) NewGalleriesClient() *GalleriesClient {
	return &GalleriesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryApplicationVersionsClient creates a new instance of GalleryApplicationVersionsClient.
func (c *ClientFactory) NewGalleryApplicationVersionsClient() *GalleryApplicationVersionsClient {
	return &GalleryApplicationVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryApplicationsClient creates a new instance of GalleryApplicationsClient.
func (c *ClientFactory) NewGalleryApplicationsClient() *GalleryApplicationsClient {
	return &GalleryApplicationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryImageVersionsClient creates a new instance of GalleryImageVersionsClient.
func (c *ClientFactory) NewGalleryImageVersionsClient() *GalleryImageVersionsClient {
	return &GalleryImageVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryImagesClient creates a new instance of GalleryImagesClient.
func (c *ClientFactory) NewGalleryImagesClient() *GalleryImagesClient {
	return &GalleryImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryInVMAccessControlProfileVersionsClient creates a new instance of GalleryInVMAccessControlProfileVersionsClient.
func (c *ClientFactory) NewGalleryInVMAccessControlProfileVersionsClient() *GalleryInVMAccessControlProfileVersionsClient {
	return &GalleryInVMAccessControlProfileVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGalleryInVMAccessControlProfilesClient creates a new instance of GalleryInVMAccessControlProfilesClient.
func (c *ClientFactory) NewGalleryInVMAccessControlProfilesClient() *GalleryInVMAccessControlProfilesClient {
	return &GalleryInVMAccessControlProfilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewGallerySharingProfileClient creates a new instance of GallerySharingProfileClient.
func (c *ClientFactory) NewGallerySharingProfileClient() *GallerySharingProfileClient {
	return &GallerySharingProfileClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewImagesClient creates a new instance of ImagesClient.
func (c *ClientFactory) NewImagesClient() *ImagesClient {
	return &ImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient.
func (c *ClientFactory) NewLogAnalyticsClient() *LogAnalyticsClient {
	return &LogAnalyticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewProximityPlacementGroupsClient creates a new instance of ProximityPlacementGroupsClient.
func (c *ClientFactory) NewProximityPlacementGroupsClient() *ProximityPlacementGroupsClient {
	return &ProximityPlacementGroupsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient.
func (c *ClientFactory) NewResourceSKUsClient() *ResourceSKUsClient {
	return &ResourceSKUsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRestorePointCollectionsClient creates a new instance of RestorePointCollectionsClient.
func (c *ClientFactory) NewRestorePointCollectionsClient() *RestorePointCollectionsClient {
	return &RestorePointCollectionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRestorePointsClient creates a new instance of RestorePointsClient.
func (c *ClientFactory) NewRestorePointsClient() *RestorePointsClient {
	return &RestorePointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSSHPublicKeysClient creates a new instance of SSHPublicKeysClient.
func (c *ClientFactory) NewSSHPublicKeysClient() *SSHPublicKeysClient {
	return &SSHPublicKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSharedGalleriesClient creates a new instance of SharedGalleriesClient.
func (c *ClientFactory) NewSharedGalleriesClient() *SharedGalleriesClient {
	return &SharedGalleriesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSharedGalleryImageVersionsClient creates a new instance of SharedGalleryImageVersionsClient.
func (c *ClientFactory) NewSharedGalleryImageVersionsClient() *SharedGalleryImageVersionsClient {
	return &SharedGalleryImageVersionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSharedGalleryImagesClient creates a new instance of SharedGalleryImagesClient.
func (c *ClientFactory) NewSharedGalleryImagesClient() *SharedGalleryImagesClient {
	return &SharedGalleryImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSnapshotsClient creates a new instance of SnapshotsClient.
func (c *ClientFactory) NewSnapshotsClient() *SnapshotsClient {
	return &SnapshotsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewSoftDeletedResourceClient creates a new instance of SoftDeletedResourceClient.
func (c *ClientFactory) NewSoftDeletedResourceClient() *SoftDeletedResourceClient {
	return &SoftDeletedResourceClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewUsageClient creates a new instance of UsageClient.
func (c *ClientFactory) NewUsageClient() *UsageClient {
	return &UsageClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineExtensionImagesClient creates a new instance of VirtualMachineExtensionImagesClient.
func (c *ClientFactory) NewVirtualMachineExtensionImagesClient() *VirtualMachineExtensionImagesClient {
	return &VirtualMachineExtensionImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineExtensionsClient creates a new instance of VirtualMachineExtensionsClient.
func (c *ClientFactory) NewVirtualMachineExtensionsClient() *VirtualMachineExtensionsClient {
	return &VirtualMachineExtensionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineImagesClient creates a new instance of VirtualMachineImagesClient.
func (c *ClientFactory) NewVirtualMachineImagesClient() *VirtualMachineImagesClient {
	return &VirtualMachineImagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineImagesEdgeZoneClient creates a new instance of VirtualMachineImagesEdgeZoneClient.
func (c *ClientFactory) NewVirtualMachineImagesEdgeZoneClient() *VirtualMachineImagesEdgeZoneClient {
	return &VirtualMachineImagesEdgeZoneClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineRunCommandsClient creates a new instance of VirtualMachineRunCommandsClient.
func (c *ClientFactory) NewVirtualMachineRunCommandsClient() *VirtualMachineRunCommandsClient {
	return &VirtualMachineRunCommandsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetExtensionsClient creates a new instance of VirtualMachineScaleSetExtensionsClient.
func (c *ClientFactory) NewVirtualMachineScaleSetExtensionsClient() *VirtualMachineScaleSetExtensionsClient {
	return &VirtualMachineScaleSetExtensionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetRollingUpgradesClient creates a new instance of VirtualMachineScaleSetRollingUpgradesClient.
func (c *ClientFactory) NewVirtualMachineScaleSetRollingUpgradesClient() *VirtualMachineScaleSetRollingUpgradesClient {
	return &VirtualMachineScaleSetRollingUpgradesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetVMExtensionsClient creates a new instance of VirtualMachineScaleSetVMExtensionsClient.
func (c *ClientFactory) NewVirtualMachineScaleSetVMExtensionsClient() *VirtualMachineScaleSetVMExtensionsClient {
	return &VirtualMachineScaleSetVMExtensionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates a new instance of VirtualMachineScaleSetVMRunCommandsClient.
func (c *ClientFactory) NewVirtualMachineScaleSetVMRunCommandsClient() *VirtualMachineScaleSetVMRunCommandsClient {
	return &VirtualMachineScaleSetVMRunCommandsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetVMsClient creates a new instance of VirtualMachineScaleSetVMsClient.
func (c *ClientFactory) NewVirtualMachineScaleSetVMsClient() *VirtualMachineScaleSetVMsClient {
	return &VirtualMachineScaleSetVMsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineScaleSetsClient creates a new instance of VirtualMachineScaleSetsClient.
func (c *ClientFactory) NewVirtualMachineScaleSetsClient() *VirtualMachineScaleSetsClient {
	return &VirtualMachineScaleSetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient.
func (c *ClientFactory) NewVirtualMachineSizesClient() *VirtualMachineSizesClient {
	return &VirtualMachineSizesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient.
func (c *ClientFactory) NewVirtualMachinesClient() *VirtualMachinesClient {
	return &VirtualMachinesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
