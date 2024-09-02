//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

// AttachedNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the AttachedNetworksClient.BeginCreateOrUpdate
// method.
type AttachedNetworksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AttachedNetworksClientBeginDeleteOptions contains the optional parameters for the AttachedNetworksClient.BeginDelete method.
type AttachedNetworksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AttachedNetworksClientGetByDevCenterOptions contains the optional parameters for the AttachedNetworksClient.GetByDevCenter
// method.
type AttachedNetworksClientGetByDevCenterOptions struct {
	// placeholder for future optional parameters
}

// AttachedNetworksClientGetByProjectOptions contains the optional parameters for the AttachedNetworksClient.GetByProject
// method.
type AttachedNetworksClientGetByProjectOptions struct {
	// placeholder for future optional parameters
}

// AttachedNetworksClientListByDevCenterOptions contains the optional parameters for the AttachedNetworksClient.NewListByDevCenterPager
// method.
type AttachedNetworksClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// AttachedNetworksClientListByProjectOptions contains the optional parameters for the AttachedNetworksClient.NewListByProjectPager
// method.
type AttachedNetworksClientListByProjectOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// CatalogsClientBeginConnectOptions contains the optional parameters for the CatalogsClient.BeginConnect method.
type CatalogsClientBeginConnectOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientBeginCreateOrUpdateOptions contains the optional parameters for the CatalogsClient.BeginCreateOrUpdate method.
type CatalogsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientBeginDeleteOptions contains the optional parameters for the CatalogsClient.BeginDelete method.
type CatalogsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientBeginSyncOptions contains the optional parameters for the CatalogsClient.BeginSync method.
type CatalogsClientBeginSyncOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientBeginUpdateOptions contains the optional parameters for the CatalogsClient.BeginUpdate method.
type CatalogsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CatalogsClientGetOptions contains the optional parameters for the CatalogsClient.Get method.
type CatalogsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientGetSyncErrorDetailsOptions contains the optional parameters for the CatalogsClient.GetSyncErrorDetails method.
type CatalogsClientGetSyncErrorDetailsOptions struct {
	// placeholder for future optional parameters
}

// CatalogsClientListByDevCenterOptions contains the optional parameters for the CatalogsClient.NewListByDevCenterPager method.
type CatalogsClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
// method.
type CheckNameAvailabilityClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CheckScopedNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckScopedNameAvailabilityClient.Execute
// method.
type CheckScopedNameAvailabilityClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CurationProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the CurationProfilesClient.BeginCreateOrUpdate
// method.
type CurationProfilesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CurationProfilesClientBeginDeleteOptions contains the optional parameters for the CurationProfilesClient.BeginDelete method.
type CurationProfilesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CurationProfilesClientBeginUpdateOptions contains the optional parameters for the CurationProfilesClient.BeginUpdate method.
type CurationProfilesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CurationProfilesClientGetOptions contains the optional parameters for the CurationProfilesClient.Get method.
type CurationProfilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CurationProfilesClientListByDevCenterOptions contains the optional parameters for the CurationProfilesClient.NewListByDevCenterPager
// method.
type CurationProfilesClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// CustomizationTasksClientGetErrorDetailsOptions contains the optional parameters for the CustomizationTasksClient.GetErrorDetails
// method.
type CustomizationTasksClientGetErrorDetailsOptions struct {
	// placeholder for future optional parameters
}

// CustomizationTasksClientGetOptions contains the optional parameters for the CustomizationTasksClient.Get method.
type CustomizationTasksClientGetOptions struct {
	// placeholder for future optional parameters
}

// CustomizationTasksClientListByCatalogOptions contains the optional parameters for the CustomizationTasksClient.NewListByCatalogPager
// method.
type CustomizationTasksClientListByCatalogOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// DevBoxDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginCreateOrUpdate
// method.
type DevBoxDefinitionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevBoxDefinitionsClientBeginDeleteOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginDelete
// method.
type DevBoxDefinitionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevBoxDefinitionsClientBeginUpdateOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginUpdate
// method.
type DevBoxDefinitionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevBoxDefinitionsClientGetByProjectOptions contains the optional parameters for the DevBoxDefinitionsClient.GetByProject
// method.
type DevBoxDefinitionsClientGetByProjectOptions struct {
	// placeholder for future optional parameters
}

// DevBoxDefinitionsClientGetOptions contains the optional parameters for the DevBoxDefinitionsClient.Get method.
type DevBoxDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DevBoxDefinitionsClientListByDevCenterOptions contains the optional parameters for the DevBoxDefinitionsClient.NewListByDevCenterPager
// method.
type DevBoxDefinitionsClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// DevBoxDefinitionsClientListByProjectOptions contains the optional parameters for the DevBoxDefinitionsClient.NewListByProjectPager
// method.
type DevBoxDefinitionsClientListByProjectOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// DevCentersClientBeginCreateOrUpdateOptions contains the optional parameters for the DevCentersClient.BeginCreateOrUpdate
// method.
type DevCentersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevCentersClientBeginDeleteOptions contains the optional parameters for the DevCentersClient.BeginDelete method.
type DevCentersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevCentersClientBeginUpdateOptions contains the optional parameters for the DevCentersClient.BeginUpdate method.
type DevCentersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DevCentersClientGetOptions contains the optional parameters for the DevCentersClient.Get method.
type DevCentersClientGetOptions struct {
	// placeholder for future optional parameters
}

// DevCentersClientListByResourceGroupOptions contains the optional parameters for the DevCentersClient.NewListByResourceGroupPager
// method.
type DevCentersClientListByResourceGroupOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// DevCentersClientListBySubscriptionOptions contains the optional parameters for the DevCentersClient.NewListBySubscriptionPager
// method.
type DevCentersClientListBySubscriptionOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// EncryptionSetsClientBeginCreateOrUpdateOptions contains the optional parameters for the EncryptionSetsClient.BeginCreateOrUpdate
// method.
type EncryptionSetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EncryptionSetsClientBeginDeleteOptions contains the optional parameters for the EncryptionSetsClient.BeginDelete method.
type EncryptionSetsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EncryptionSetsClientBeginUpdateOptions contains the optional parameters for the EncryptionSetsClient.BeginUpdate method.
type EncryptionSetsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// EncryptionSetsClientGetOptions contains the optional parameters for the EncryptionSetsClient.Get method.
type EncryptionSetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// EncryptionSetsClientListOptions contains the optional parameters for the EncryptionSetsClient.NewListPager method.
type EncryptionSetsClientListOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// EnvironmentDefinitionsClientGetByProjectCatalogOptions contains the optional parameters for the EnvironmentDefinitionsClient.GetByProjectCatalog
// method.
type EnvironmentDefinitionsClientGetByProjectCatalogOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentDefinitionsClientGetErrorDetailsOptions contains the optional parameters for the EnvironmentDefinitionsClient.GetErrorDetails
// method.
type EnvironmentDefinitionsClientGetErrorDetailsOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentDefinitionsClientGetOptions contains the optional parameters for the EnvironmentDefinitionsClient.Get method.
type EnvironmentDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentDefinitionsClientListByCatalogOptions contains the optional parameters for the EnvironmentDefinitionsClient.NewListByCatalogPager
// method.
type EnvironmentDefinitionsClientListByCatalogOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// EnvironmentDefinitionsClientListByProjectCatalogOptions contains the optional parameters for the EnvironmentDefinitionsClient.NewListByProjectCatalogPager
// method.
type EnvironmentDefinitionsClientListByProjectCatalogOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentTypesClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentTypesClient.CreateOrUpdate
// method.
type EnvironmentTypesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentTypesClientDeleteOptions contains the optional parameters for the EnvironmentTypesClient.Delete method.
type EnvironmentTypesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentTypesClientGetOptions contains the optional parameters for the EnvironmentTypesClient.Get method.
type EnvironmentTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentTypesClientListByDevCenterOptions contains the optional parameters for the EnvironmentTypesClient.NewListByDevCenterPager
// method.
type EnvironmentTypesClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// EnvironmentTypesClientUpdateOptions contains the optional parameters for the EnvironmentTypesClient.Update method.
type EnvironmentTypesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// GalleriesClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleriesClient.BeginCreateOrUpdate
// method.
type GalleriesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GalleriesClientBeginDeleteOptions contains the optional parameters for the GalleriesClient.BeginDelete method.
type GalleriesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// GalleriesClientGetOptions contains the optional parameters for the GalleriesClient.Get method.
type GalleriesClientGetOptions struct {
	// placeholder for future optional parameters
}

// GalleriesClientListByDevCenterOptions contains the optional parameters for the GalleriesClient.NewListByDevCenterPager
// method.
type GalleriesClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ImageVersionsClientGetByProjectOptions contains the optional parameters for the ImageVersionsClient.GetByProject method.
type ImageVersionsClientGetByProjectOptions struct {
	// placeholder for future optional parameters
}

// ImageVersionsClientGetOptions contains the optional parameters for the ImageVersionsClient.Get method.
type ImageVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ImageVersionsClientListByImageOptions contains the optional parameters for the ImageVersionsClient.NewListByImagePager
// method.
type ImageVersionsClientListByImageOptions struct {
	// placeholder for future optional parameters
}

// ImageVersionsClientListByProjectOptions contains the optional parameters for the ImageVersionsClient.NewListByProjectPager
// method.
type ImageVersionsClientListByProjectOptions struct {
	// placeholder for future optional parameters
}

// ImagesClientGetByProjectOptions contains the optional parameters for the ImagesClient.GetByProject method.
type ImagesClientGetByProjectOptions struct {
	// placeholder for future optional parameters
}

// ImagesClientGetOptions contains the optional parameters for the ImagesClient.Get method.
type ImagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ImagesClientListByDevCenterOptions contains the optional parameters for the ImagesClient.NewListByDevCenterPager method.
type ImagesClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ImagesClientListByGalleryOptions contains the optional parameters for the ImagesClient.NewListByGalleryPager method.
type ImagesClientListByGalleryOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ImagesClientListByProjectOptions contains the optional parameters for the ImagesClient.NewListByProjectPager method.
type ImagesClientListByProjectOptions struct {
	// placeholder for future optional parameters
}

// NetworkConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkConnectionsClient.BeginCreateOrUpdate
// method.
type NetworkConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkConnectionsClientBeginDeleteOptions contains the optional parameters for the NetworkConnectionsClient.BeginDelete
// method.
type NetworkConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkConnectionsClientBeginRunHealthChecksOptions contains the optional parameters for the NetworkConnectionsClient.BeginRunHealthChecks
// method.
type NetworkConnectionsClientBeginRunHealthChecksOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkConnectionsClientBeginUpdateOptions contains the optional parameters for the NetworkConnectionsClient.BeginUpdate
// method.
type NetworkConnectionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkConnectionsClientGetHealthDetailsOptions contains the optional parameters for the NetworkConnectionsClient.GetHealthDetails
// method.
type NetworkConnectionsClientGetHealthDetailsOptions struct {
	// placeholder for future optional parameters
}

// NetworkConnectionsClientGetOptions contains the optional parameters for the NetworkConnectionsClient.Get method.
type NetworkConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// NetworkConnectionsClientListByResourceGroupOptions contains the optional parameters for the NetworkConnectionsClient.NewListByResourceGroupPager
// method.
type NetworkConnectionsClientListByResourceGroupOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// NetworkConnectionsClientListBySubscriptionOptions contains the optional parameters for the NetworkConnectionsClient.NewListBySubscriptionPager
// method.
type NetworkConnectionsClientListBySubscriptionOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// NetworkConnectionsClientListHealthDetailsOptions contains the optional parameters for the NetworkConnectionsClient.NewListHealthDetailsPager
// method.
type NetworkConnectionsClientListHealthDetailsOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the NetworkConnectionsClient.NewListOutboundNetworkDependenciesEndpointsPager
// method.
type NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// OperationStatusesClientGetOptions contains the optional parameters for the OperationStatusesClient.Get method.
type OperationStatusesClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PlanMembersClientBeginCreateOrUpdateOptions contains the optional parameters for the PlanMembersClient.BeginCreateOrUpdate
// method.
type PlanMembersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlanMembersClientBeginDeleteOptions contains the optional parameters for the PlanMembersClient.BeginDelete method.
type PlanMembersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlanMembersClientBeginUpdateOptions contains the optional parameters for the PlanMembersClient.BeginUpdate method.
type PlanMembersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlanMembersClientGetOptions contains the optional parameters for the PlanMembersClient.Get method.
type PlanMembersClientGetOptions struct {
	// placeholder for future optional parameters
}

// PlanMembersClientListOptions contains the optional parameters for the PlanMembersClient.NewListPager method.
type PlanMembersClientListOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// PlansClientBeginCreateOrUpdateOptions contains the optional parameters for the PlansClient.BeginCreateOrUpdate method.
type PlansClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlansClientBeginDeleteOptions contains the optional parameters for the PlansClient.BeginDelete method.
type PlansClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlansClientBeginUpdateOptions contains the optional parameters for the PlansClient.BeginUpdate method.
type PlansClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PlansClientGetOptions contains the optional parameters for the PlansClient.Get method.
type PlansClientGetOptions struct {
	// placeholder for future optional parameters
}

// PlansClientListByResourceGroupOptions contains the optional parameters for the PlansClient.NewListByResourceGroupPager
// method.
type PlansClientListByResourceGroupOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// PlansClientListBySubscriptionOptions contains the optional parameters for the PlansClient.NewListBySubscriptionPager method.
type PlansClientListBySubscriptionOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// PoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the PoolsClient.BeginCreateOrUpdate method.
type PoolsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PoolsClientBeginDeleteOptions contains the optional parameters for the PoolsClient.BeginDelete method.
type PoolsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PoolsClientBeginRunHealthChecksOptions contains the optional parameters for the PoolsClient.BeginRunHealthChecks method.
type PoolsClientBeginRunHealthChecksOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PoolsClientBeginUpdateOptions contains the optional parameters for the PoolsClient.BeginUpdate method.
type PoolsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PoolsClientGetOptions contains the optional parameters for the PoolsClient.Get method.
type PoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PoolsClientListByProjectOptions contains the optional parameters for the PoolsClient.NewListByProjectPager method.
type PoolsClientListByProjectOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectAllowedEnvironmentTypesClientGetOptions contains the optional parameters for the ProjectAllowedEnvironmentTypesClient.Get
// method.
type ProjectAllowedEnvironmentTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectAllowedEnvironmentTypesClientListOptions contains the optional parameters for the ProjectAllowedEnvironmentTypesClient.NewListPager
// method.
type ProjectAllowedEnvironmentTypesClientListOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectCatalogEnvironmentDefinitionsClientGetErrorDetailsOptions contains the optional parameters for the ProjectCatalogEnvironmentDefinitionsClient.GetErrorDetails
// method.
type ProjectCatalogEnvironmentDefinitionsClientGetErrorDetailsOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogImageDefinitionBuildClientBeginCancelOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.BeginCancel
// method.
type ProjectCatalogImageDefinitionBuildClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogImageDefinitionBuildClientGetBuildDetailsOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.GetBuildDetails
// method.
type ProjectCatalogImageDefinitionBuildClientGetBuildDetailsOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogImageDefinitionBuildClientGetOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildClient.Get
// method.
type ProjectCatalogImageDefinitionBuildClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogImageDefinitionBuildsClientListByImageDefinitionOptions contains the optional parameters for the ProjectCatalogImageDefinitionBuildsClient.NewListByImageDefinitionPager
// method.
type ProjectCatalogImageDefinitionBuildsClientListByImageDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogImageDefinitionsClientBeginBuildImageOptions contains the optional parameters for the ProjectCatalogImageDefinitionsClient.BeginBuildImage
// method.
type ProjectCatalogImageDefinitionsClientBeginBuildImageOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogImageDefinitionsClientGetByProjectCatalogOptions contains the optional parameters for the ProjectCatalogImageDefinitionsClient.GetByProjectCatalog
// method.
type ProjectCatalogImageDefinitionsClientGetByProjectCatalogOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogImageDefinitionsClientListByProjectCatalogOptions contains the optional parameters for the ProjectCatalogImageDefinitionsClient.NewListByProjectCatalogPager
// method.
type ProjectCatalogImageDefinitionsClientListByProjectCatalogOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectCatalogsClientBeginConnectOptions contains the optional parameters for the ProjectCatalogsClient.BeginConnect method.
type ProjectCatalogsClientBeginConnectOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProjectCatalogsClient.BeginCreateOrUpdate
// method.
type ProjectCatalogsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogsClientBeginDeleteOptions contains the optional parameters for the ProjectCatalogsClient.BeginDelete method.
type ProjectCatalogsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogsClientBeginPatchOptions contains the optional parameters for the ProjectCatalogsClient.BeginPatch method.
type ProjectCatalogsClientBeginPatchOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogsClientBeginSyncOptions contains the optional parameters for the ProjectCatalogsClient.BeginSync method.
type ProjectCatalogsClientBeginSyncOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectCatalogsClientGetOptions contains the optional parameters for the ProjectCatalogsClient.Get method.
type ProjectCatalogsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogsClientGetSyncErrorDetailsOptions contains the optional parameters for the ProjectCatalogsClient.GetSyncErrorDetails
// method.
type ProjectCatalogsClientGetSyncErrorDetailsOptions struct {
	// placeholder for future optional parameters
}

// ProjectCatalogsClientListOptions contains the optional parameters for the ProjectCatalogsClient.NewListPager method.
type ProjectCatalogsClientListOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectEnvironmentTypesClientCreateOrUpdateOptions contains the optional parameters for the ProjectEnvironmentTypesClient.CreateOrUpdate
// method.
type ProjectEnvironmentTypesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProjectEnvironmentTypesClientDeleteOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Delete
// method.
type ProjectEnvironmentTypesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ProjectEnvironmentTypesClientGetOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Get method.
type ProjectEnvironmentTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectEnvironmentTypesClientListOptions contains the optional parameters for the ProjectEnvironmentTypesClient.NewListPager
// method.
type ProjectEnvironmentTypesClientListOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectEnvironmentTypesClientUpdateOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Update
// method.
type ProjectEnvironmentTypesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProjectPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ProjectPoliciesClient.BeginCreateOrUpdate
// method.
type ProjectPoliciesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectPoliciesClientBeginDeleteOptions contains the optional parameters for the ProjectPoliciesClient.BeginDelete method.
type ProjectPoliciesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectPoliciesClientBeginUpdateOptions contains the optional parameters for the ProjectPoliciesClient.BeginUpdate method.
type ProjectPoliciesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectPoliciesClientGetOptions contains the optional parameters for the ProjectPoliciesClient.Get method.
type ProjectPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectPoliciesClientListByDevCenterOptions contains the optional parameters for the ProjectPoliciesClient.NewListByDevCenterPager
// method.
type ProjectPoliciesClientListByDevCenterOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectsClientBeginCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.BeginCreateOrUpdate method.
type ProjectsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectsClientBeginDeleteOptions contains the optional parameters for the ProjectsClient.BeginDelete method.
type ProjectsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectsClientBeginUpdateOptions contains the optional parameters for the ProjectsClient.BeginUpdate method.
type ProjectsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ProjectsClientGetInheritedSettingsOptions contains the optional parameters for the ProjectsClient.GetInheritedSettings
// method.
type ProjectsClientGetInheritedSettingsOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
type ProjectsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProjectsClientListByResourceGroupOptions contains the optional parameters for the ProjectsClient.NewListByResourceGroupPager
// method.
type ProjectsClientListByResourceGroupOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// ProjectsClientListBySubscriptionOptions contains the optional parameters for the ProjectsClient.NewListBySubscriptionPager
// method.
type ProjectsClientListBySubscriptionOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SKUsClientListByProjectOptions contains the optional parameters for the SKUsClient.NewListByProjectPager method.
type SKUsClientListByProjectOptions struct {
	// placeholder for future optional parameters
}

// SKUsClientListBySubscriptionOptions contains the optional parameters for the SKUsClient.NewListBySubscriptionPager method.
type SKUsClientListBySubscriptionOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SchedulesClientBeginCreateOrUpdateOptions contains the optional parameters for the SchedulesClient.BeginCreateOrUpdate
// method.
type SchedulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SchedulesClientBeginDeleteOptions contains the optional parameters for the SchedulesClient.BeginDelete method.
type SchedulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SchedulesClientBeginUpdateOptions contains the optional parameters for the SchedulesClient.BeginUpdate method.
type SchedulesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SchedulesClientGetOptions contains the optional parameters for the SchedulesClient.Get method.
type SchedulesClientGetOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// SchedulesClientListByPoolOptions contains the optional parameters for the SchedulesClient.NewListByPoolPager method.
type SchedulesClientListByPoolOptions struct {
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int32
}

// UsagesClientListByLocationOptions contains the optional parameters for the UsagesClient.NewListByLocationPager method.
type UsagesClientListByLocationOptions struct {
	// placeholder for future optional parameters
}
