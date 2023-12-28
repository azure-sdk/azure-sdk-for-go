//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

// AgentVersionClientGetOptions contains the optional parameters for the AgentVersionClient.Get method.
type AgentVersionClientGetOptions struct {
	// placeholder for future optional parameters
}

// AgentVersionClientListOptions contains the optional parameters for the AgentVersionClient.List method.
type AgentVersionClientListOptions struct {
	// placeholder for future optional parameters
}

// ExtensionMetadataClientGetOptions contains the optional parameters for the ExtensionMetadataClient.Get method.
type ExtensionMetadataClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExtensionMetadataClientListOptions contains the optional parameters for the ExtensionMetadataClient.NewListPager method.
type ExtensionMetadataClientListOptions struct {
	// placeholder for future optional parameters
}

// HybridIdentityMetadataClientGetOptions contains the optional parameters for the HybridIdentityMetadataClient.Get method.
type HybridIdentityMetadataClientGetOptions struct {
	// placeholder for future optional parameters
}

// HybridIdentityMetadataClientListByMachinesOptions contains the optional parameters for the HybridIdentityMetadataClient.NewListByMachinesPager
// method.
type HybridIdentityMetadataClientListByMachinesOptions struct {
	// placeholder for future optional parameters
}

// LicenseProfilesClientBeginCreateOrUpdateOptions contains the optional parameters for the LicenseProfilesClient.BeginCreateOrUpdate
// method.
type LicenseProfilesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicenseProfilesClientBeginDeleteOptions contains the optional parameters for the LicenseProfilesClient.BeginDelete method.
type LicenseProfilesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicenseProfilesClientBeginUpdateOptions contains the optional parameters for the LicenseProfilesClient.BeginUpdate method.
type LicenseProfilesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicenseProfilesClientGetOptions contains the optional parameters for the LicenseProfilesClient.Get method.
type LicenseProfilesClientGetOptions struct {
	// placeholder for future optional parameters
}

// LicenseProfilesClientListOptions contains the optional parameters for the LicenseProfilesClient.NewListPager method.
type LicenseProfilesClientListOptions struct {
	// placeholder for future optional parameters
}

// LicensesClientBeginCreateOrUpdateOptions contains the optional parameters for the LicensesClient.BeginCreateOrUpdate method.
type LicensesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicensesClientBeginDeleteOptions contains the optional parameters for the LicensesClient.BeginDelete method.
type LicensesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicensesClientBeginUpdateOptions contains the optional parameters for the LicensesClient.BeginUpdate method.
type LicensesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicensesClientBeginValidateLicenseOptions contains the optional parameters for the LicensesClient.BeginValidateLicense
// method.
type LicensesClientBeginValidateLicenseOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LicensesClientGetOptions contains the optional parameters for the LicensesClient.Get method.
type LicensesClientGetOptions struct {
	// placeholder for future optional parameters
}

// LicensesClientListByResourceGroupOptions contains the optional parameters for the LicensesClient.NewListByResourceGroupPager
// method.
type LicensesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// LicensesClientListBySubscriptionOptions contains the optional parameters for the LicensesClient.NewListBySubscriptionPager
// method.
type LicensesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MachineExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the MachineExtensionsClient.BeginCreateOrUpdate
// method.
type MachineExtensionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineExtensionsClientBeginDeleteOptions contains the optional parameters for the MachineExtensionsClient.BeginDelete
// method.
type MachineExtensionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineExtensionsClientBeginUpdateOptions contains the optional parameters for the MachineExtensionsClient.BeginUpdate
// method.
type MachineExtensionsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineExtensionsClientGetOptions contains the optional parameters for the MachineExtensionsClient.Get method.
type MachineExtensionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MachineExtensionsClientListOptions contains the optional parameters for the MachineExtensionsClient.NewListPager method.
type MachineExtensionsClientListOptions struct {
	// The expand expression to apply on the operation.
	Expand *string
}

// MachineRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the MachineRunCommandsClient.BeginCreateOrUpdate
// method.
type MachineRunCommandsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineRunCommandsClientBeginDeleteOptions contains the optional parameters for the MachineRunCommandsClient.BeginDelete
// method.
type MachineRunCommandsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineRunCommandsClientBeginUpdateOptions contains the optional parameters for the MachineRunCommandsClient.BeginUpdate
// method.
type MachineRunCommandsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachineRunCommandsClientGetOptions contains the optional parameters for the MachineRunCommandsClient.Get method.
type MachineRunCommandsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MachineRunCommandsClientListOptions contains the optional parameters for the MachineRunCommandsClient.NewListPager method.
type MachineRunCommandsClientListOptions struct {
	// The expand expression to apply on the operation.
	Expand *string
}

// MachinesClientBeginAssessPatchesOptions contains the optional parameters for the MachinesClient.BeginAssessPatches method.
type MachinesClientBeginAssessPatchesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachinesClientBeginInstallPatchesOptions contains the optional parameters for the MachinesClient.BeginInstallPatches method.
type MachinesClientBeginInstallPatchesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MachinesClientCreateOrUpdateOptions contains the optional parameters for the MachinesClient.CreateOrUpdate method.
type MachinesClientCreateOrUpdateOptions struct {
	// Expands referenced resources.
	Expand *string
}

// MachinesClientDeleteOptions contains the optional parameters for the MachinesClient.Delete method.
type MachinesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MachinesClientGetOptions contains the optional parameters for the MachinesClient.Get method.
type MachinesClientGetOptions struct {
	// The expand expression to apply on the operation.
	Expand *InstanceViewTypes
}

// MachinesClientListByResourceGroupOptions contains the optional parameters for the MachinesClient.NewListByResourceGroupPager
// method.
type MachinesClientListByResourceGroupOptions struct {
	// Expands referenced resources.
	Expand *string
}

// MachinesClientListBySubscriptionOptions contains the optional parameters for the MachinesClient.NewListBySubscriptionPager
// method.
type MachinesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MachinesClientUpdateOptions contains the optional parameters for the MachinesClient.Update method.
type MachinesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientBeginUpgradeExtensionsOptions contains the optional parameters for the ManagementClient.BeginUpgradeExtensions
// method.
type ManagementClientBeginUpgradeExtensionsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the NetworkConfigurationsClient.CreateOrUpdate
// method.
type NetworkConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// NetworkConfigurationsClientGetOptions contains the optional parameters for the NetworkConfigurationsClient.Get method.
type NetworkConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// NetworkConfigurationsClientUpdateOptions contains the optional parameters for the NetworkConfigurationsClient.Update method.
type NetworkConfigurationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// NetworkProfileClientGetOptions contains the optional parameters for the NetworkProfileClient.Get method.
type NetworkProfileClientGetOptions struct {
	// placeholder for future optional parameters
}

// NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions contains the optional parameters for
// the NetworkSecurityPerimeterConfigurationsClient.BeginReconcileForPrivateLinkScope method.
type NetworkSecurityPerimeterConfigurationsClientBeginReconcileForPrivateLinkScopeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.GetByPrivateLinkScope
// method.
type NetworkSecurityPerimeterConfigurationsClientGetByPrivateLinkScopeOptions struct {
	// placeholder for future optional parameters
}

// NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.NewListByPrivateLinkScopePager
// method.
type NetworkSecurityPerimeterConfigurationsClientListByPrivateLinkScopeOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager
// method.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByPrivateLinkScopeOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByPrivateLinkScopePager
// method.
type PrivateLinkResourcesClientListByPrivateLinkScopeOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientBeginDeleteOptions contains the optional parameters for the PrivateLinkScopesClient.BeginDelete
// method.
type PrivateLinkScopesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateLinkScopesClientCreateOrUpdateOptions contains the optional parameters for the PrivateLinkScopesClient.CreateOrUpdate
// method.
type PrivateLinkScopesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientGetOptions contains the optional parameters for the PrivateLinkScopesClient.Get method.
type PrivateLinkScopesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientGetValidationDetailsForMachineOptions contains the optional parameters for the PrivateLinkScopesClient.GetValidationDetailsForMachine
// method.
type PrivateLinkScopesClientGetValidationDetailsForMachineOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientGetValidationDetailsOptions contains the optional parameters for the PrivateLinkScopesClient.GetValidationDetails
// method.
type PrivateLinkScopesClientGetValidationDetailsOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientListByResourceGroupOptions contains the optional parameters for the PrivateLinkScopesClient.NewListByResourceGroupPager
// method.
type PrivateLinkScopesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientListOptions contains the optional parameters for the PrivateLinkScopesClient.NewListPager method.
type PrivateLinkScopesClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkScopesClientUpdateTagsOptions contains the optional parameters for the PrivateLinkScopesClient.UpdateTags method.
type PrivateLinkScopesClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}
