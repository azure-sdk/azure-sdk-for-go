//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

// AgentVersionClientGetResponse contains the response from method AgentVersionClient.Get.
type AgentVersionClientGetResponse struct {
	// Describes properties of Agent Version.
	AgentVersion
}

// AgentVersionClientListResponse contains the response from method AgentVersionClient.List.
type AgentVersionClientListResponse struct {
	// Describes AgentVersions List.
	AgentVersionsList
}

// ExtensionMetadataClientGetResponse contains the response from method ExtensionMetadataClient.Get.
type ExtensionMetadataClientGetResponse struct {
	// Describes a Extension Metadata
	ExtensionValue
}

// ExtensionMetadataClientListResponse contains the response from method ExtensionMetadataClient.NewListPager.
type ExtensionMetadataClientListResponse struct {
	// The List Extension Metadata response.
	ExtensionValueListResult
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	// Defines the HybridIdentityMetadata.
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListByMachinesResponse contains the response from method HybridIdentityMetadataClient.NewListByMachinesPager.
type HybridIdentityMetadataClientListByMachinesResponse struct {
	// List of HybridIdentityMetadata.
	HybridIdentityMetadataList
}

// LicenseProfilesClientCreateOrUpdateResponse contains the response from method LicenseProfilesClient.BeginCreateOrUpdate.
type LicenseProfilesClientCreateOrUpdateResponse struct {
	// Describes a license profile in a hybrid machine.
	LicenseProfile
}

// LicenseProfilesClientDeleteResponse contains the response from method LicenseProfilesClient.BeginDelete.
type LicenseProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// LicenseProfilesClientGetResponse contains the response from method LicenseProfilesClient.Get.
type LicenseProfilesClientGetResponse struct {
	// Describes a license profile in a hybrid machine.
	LicenseProfile
}

// LicenseProfilesClientListResponse contains the response from method LicenseProfilesClient.NewListPager.
type LicenseProfilesClientListResponse struct {
	// The List hybrid machine license profile operation response.
	LicenseProfilesListResult
}

// LicenseProfilesClientUpdateResponse contains the response from method LicenseProfilesClient.BeginUpdate.
type LicenseProfilesClientUpdateResponse struct {
	// Describes a license profile in a hybrid machine.
	LicenseProfile
}

// LicensesClientCreateOrUpdateResponse contains the response from method LicensesClient.BeginCreateOrUpdate.
type LicensesClientCreateOrUpdateResponse struct {
	// Describes a license in a hybrid machine.
	License
}

// LicensesClientDeleteResponse contains the response from method LicensesClient.BeginDelete.
type LicensesClientDeleteResponse struct {
	// placeholder for future response values
}

// LicensesClientGetResponse contains the response from method LicensesClient.Get.
type LicensesClientGetResponse struct {
	// Describes a license in a hybrid machine.
	License
}

// LicensesClientListByResourceGroupResponse contains the response from method LicensesClient.NewListByResourceGroupPager.
type LicensesClientListByResourceGroupResponse struct {
	// The List license operation response.
	LicensesListResult
}

// LicensesClientListBySubscriptionResponse contains the response from method LicensesClient.NewListBySubscriptionPager.
type LicensesClientListBySubscriptionResponse struct {
	// The List license operation response.
	LicensesListResult
}

// LicensesClientUpdateResponse contains the response from method LicensesClient.BeginUpdate.
type LicensesClientUpdateResponse struct {
	// Describes a license in a hybrid machine.
	License
}

// LicensesClientValidateLicenseResponse contains the response from method LicensesClient.BeginValidateLicense.
type LicensesClientValidateLicenseResponse struct {
	// Describes a license in a hybrid machine.
	License
}

// MachineExtensionsClientCreateOrUpdateResponse contains the response from method MachineExtensionsClient.BeginCreateOrUpdate.
type MachineExtensionsClientCreateOrUpdateResponse struct {
	// Describes a Machine Extension.
	MachineExtension
}

// MachineExtensionsClientDeleteResponse contains the response from method MachineExtensionsClient.BeginDelete.
type MachineExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineExtensionsClientGetResponse contains the response from method MachineExtensionsClient.Get.
type MachineExtensionsClientGetResponse struct {
	// Describes a Machine Extension.
	MachineExtension
}

// MachineExtensionsClientListResponse contains the response from method MachineExtensionsClient.NewListPager.
type MachineExtensionsClientListResponse struct {
	// Describes the Machine Extensions List Result.
	MachineExtensionsListResult
}

// MachineExtensionsClientUpdateResponse contains the response from method MachineExtensionsClient.BeginUpdate.
type MachineExtensionsClientUpdateResponse struct {
	// Describes a Machine Extension.
	MachineExtension
}

// MachineRunCommandsClientCreateOrUpdateResponse contains the response from method MachineRunCommandsClient.BeginCreateOrUpdate.
type MachineRunCommandsClientCreateOrUpdateResponse struct {
	// Describes a Run Command
	MachineRunCommand
}

// MachineRunCommandsClientDeleteResponse contains the response from method MachineRunCommandsClient.BeginDelete.
type MachineRunCommandsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineRunCommandsClientGetResponse contains the response from method MachineRunCommandsClient.Get.
type MachineRunCommandsClientGetResponse struct {
	// Describes a Run Command
	MachineRunCommand
}

// MachineRunCommandsClientListResponse contains the response from method MachineRunCommandsClient.NewListPager.
type MachineRunCommandsClientListResponse struct {
	// Describes the Run Commands List Result.
	MachineRunCommandsListResult
}

// MachineRunCommandsClientUpdateResponse contains the response from method MachineRunCommandsClient.BeginUpdate.
type MachineRunCommandsClientUpdateResponse struct {
	// Describes a Run Command
	MachineRunCommand
}

// MachinesClientAssessPatchesResponse contains the response from method MachinesClient.BeginAssessPatches.
type MachinesClientAssessPatchesResponse struct {
	// Describes the properties of an AssessPatches result.
	MachineAssessPatchesResult
}

// MachinesClientCreateOrUpdateResponse contains the response from method MachinesClient.CreateOrUpdate.
type MachinesClientCreateOrUpdateResponse struct {
	// Describes a hybrid machine.
	Machine
}

// MachinesClientDeleteResponse contains the response from method MachinesClient.Delete.
type MachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	// Describes a hybrid machine.
	Machine
}

// MachinesClientInstallPatchesResponse contains the response from method MachinesClient.BeginInstallPatches.
type MachinesClientInstallPatchesResponse struct {
	// The result summary of an installation operation.
	MachineInstallPatchesResult
}

// MachinesClientListByResourceGroupResponse contains the response from method MachinesClient.NewListByResourceGroupPager.
type MachinesClientListByResourceGroupResponse struct {
	// The List hybrid machine operation response.
	MachineListResult
}

// MachinesClientListBySubscriptionResponse contains the response from method MachinesClient.NewListBySubscriptionPager.
type MachinesClientListBySubscriptionResponse struct {
	// The List hybrid machine operation response.
	MachineListResult
}

// MachinesClientUpdateResponse contains the response from method MachinesClient.Update.
type MachinesClientUpdateResponse struct {
	// Describes a hybrid machine.
	Machine
}

// ManagementClientUpgradeExtensionsResponse contains the response from method ManagementClient.BeginUpgradeExtensions.
type ManagementClientUpgradeExtensionsResponse struct {
	// placeholder for future response values
}

// NetworkConfigurationsClientCreateOrUpdateResponse contains the response from method NetworkConfigurationsClient.CreateOrUpdate.
type NetworkConfigurationsClientCreateOrUpdateResponse struct {
	NetworkConfiguration
}

// NetworkConfigurationsClientGetResponse contains the response from method NetworkConfigurationsClient.Get.
type NetworkConfigurationsClientGetResponse struct {
	NetworkConfiguration
}

// NetworkConfigurationsClientUpdateResponse contains the response from method NetworkConfigurationsClient.Update.
type NetworkConfigurationsClientUpdateResponse struct {
	NetworkConfiguration
}

// NetworkProfileClientGetResponse contains the response from method NetworkProfileClient.Get.
type NetworkProfileClientGetResponse struct {
	// Describes the network information on this machine.
	NetworkProfile
}

// NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope.
type NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeResponse struct {
	// Properties that define a Network Security Perimeter resource.
	NetworkSecurityPerimeterConfiguration
}

// NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager.
type NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeResponse struct {
	// A list of network security perimeter configurations.
	NetworkSecurityPerimeterConfigurationListResult
}

// NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope.
type NetworkSecurityPerimeterConfigurationsClientReconcileForPrivateLinkScopeResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The List Compute Operation operation response.
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	// A list of private endpoint connections.
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.NewListByPrivateLinkScopePager.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	// An Azure Arc PrivateLinkScope definition.
	PrivateLinkScope
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.BeginDelete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	// An Azure Arc PrivateLinkScope definition.
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
	// Describes the list of Azure Arc PrivateLinkScope resources.
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.NewListPager.
type PrivateLinkScopesClientListResponse struct {
	// Describes the list of Azure Arc PrivateLinkScope resources.
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	// An Azure Arc PrivateLinkScope definition.
	PrivateLinkScope
}