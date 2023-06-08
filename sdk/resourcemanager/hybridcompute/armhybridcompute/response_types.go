//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

// AgentVersionClientGetResponse contains the response from method AgentVersionClient.Get.
type AgentVersionClientGetResponse struct {
	AgentVersion
}

// AgentVersionClientListResponse contains the response from method AgentVersionClient.List.
type AgentVersionClientListResponse struct {
	AgentVersionsList
}

// ExtensionMetadataClientGetResponse contains the response from method ExtensionMetadataClient.Get.
type ExtensionMetadataClientGetResponse struct {
	ExtensionValue
}

// ExtensionMetadataClientListResponse contains the response from method ExtensionMetadataClient.NewListPager.
type ExtensionMetadataClientListResponse struct {
	ExtensionValueListResult
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListByMachinesResponse contains the response from method HybridIdentityMetadataClient.NewListByMachinesPager.
type HybridIdentityMetadataClientListByMachinesResponse struct {
	HybridIdentityMetadataList
}

// MachineExtensionsClientCreateOrUpdateResponse contains the response from method MachineExtensionsClient.BeginCreateOrUpdate.
type MachineExtensionsClientCreateOrUpdateResponse struct {
	MachineExtension
}

// MachineExtensionsClientDeleteResponse contains the response from method MachineExtensionsClient.BeginDelete.
type MachineExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineExtensionsClientGetResponse contains the response from method MachineExtensionsClient.Get.
type MachineExtensionsClientGetResponse struct {
	MachineExtension
}

// MachineExtensionsClientListResponse contains the response from method MachineExtensionsClient.NewListPager.
type MachineExtensionsClientListResponse struct {
	MachineExtensionsListResult
}

// MachineExtensionsClientUpdateResponse contains the response from method MachineExtensionsClient.BeginUpdate.
type MachineExtensionsClientUpdateResponse struct {
	MachineExtension
}

// MachineRunCommandsClientCreateOrUpdateResponse contains the response from method MachineRunCommandsClient.BeginCreateOrUpdate.
type MachineRunCommandsClientCreateOrUpdateResponse struct {
	MachineRunCommand
}

// MachineRunCommandsClientDeleteResponse contains the response from method MachineRunCommandsClient.BeginDelete.
type MachineRunCommandsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineRunCommandsClientGetResponse contains the response from method MachineRunCommandsClient.Get.
type MachineRunCommandsClientGetResponse struct {
	MachineRunCommand
}

// MachineRunCommandsClientListResponse contains the response from method MachineRunCommandsClient.NewListPager.
type MachineRunCommandsClientListResponse struct {
	MachineRunCommandsListResult
}

// MachineRunCommandsClientUpdateResponse contains the response from method MachineRunCommandsClient.BeginUpdate.
type MachineRunCommandsClientUpdateResponse struct {
	MachineRunCommand
}

// MachinesClientAssessPatchesResponse contains the response from method MachinesClient.BeginAssessPatches.
type MachinesClientAssessPatchesResponse struct {
	MachineAssessPatchesResult
}

// MachinesClientCreateOrUpdateResponse contains the response from method MachinesClient.CreateOrUpdate.
type MachinesClientCreateOrUpdateResponse struct {
	Machine
}

// MachinesClientDeleteResponse contains the response from method MachinesClient.Delete.
type MachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	Machine
}

// MachinesClientInstallPatchesResponse contains the response from method MachinesClient.BeginInstallPatches.
type MachinesClientInstallPatchesResponse struct {
	MachineInstallPatchesResult
}

// MachinesClientListByResourceGroupResponse contains the response from method MachinesClient.NewListByResourceGroupPager.
type MachinesClientListByResourceGroupResponse struct {
	MachineListResult
}

// MachinesClientListBySubscriptionResponse contains the response from method MachinesClient.NewListBySubscriptionPager.
type MachinesClientListBySubscriptionResponse struct {
	MachineListResult
}

// MachinesClientUpdateResponse contains the response from method MachinesClient.Update.
type MachinesClientUpdateResponse struct {
	Machine
}

// ManagementClientUpgradeExtensionsResponse contains the response from method ManagementClient.BeginUpgradeExtensions.
type ManagementClientUpgradeExtensionsResponse struct {
	// placeholder for future response values
}

// NetworkProfileClientGetResponse contains the response from method NetworkProfileClient.Get.
type NetworkProfileClientGetResponse struct {
	NetworkProfile
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.NewListByPrivateLinkScopePager.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	PrivateLinkResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	PrivateLinkScope
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.BeginDelete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	PrivateLinkScope
}

// PrivateLinkScopesClientGetValidationDetailsForMachineResponse contains the response from method PrivateLinkScopesClient.GetValidationDetailsForMachine.
type PrivateLinkScopesClientGetValidationDetailsForMachineResponse struct {
	PrivateLinkScopeValidationDetails
}

// PrivateLinkScopesClientGetValidationDetailsResponse contains the response from method PrivateLinkScopesClient.GetValidationDetails.
type PrivateLinkScopesClientGetValidationDetailsResponse struct {
	PrivateLinkScopeValidationDetails
}

// PrivateLinkScopesClientListByResourceGroupResponse contains the response from method PrivateLinkScopesClient.NewListByResourceGroupPager.
type PrivateLinkScopesClientListByResourceGroupResponse struct {
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.NewListPager.
type PrivateLinkScopesClientListResponse struct {
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	PrivateLinkScope
}
