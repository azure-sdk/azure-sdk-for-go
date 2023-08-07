//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

// AgentPoolsClientBeginAbortLatestOperationOptions contains the optional parameters for the AgentPoolsClient.BeginAbortLatestOperation
// method.
type AgentPoolsClientBeginAbortLatestOperationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginCreateOrUpdate
// method.
type AgentPoolsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
type AgentPoolsClientBeginDeleteOptions struct {
	// ignore-pod-disruption-budget=true to delete those pods on a node without considering Pod Disruption Budget
	IgnorePodDisruptionBudget *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientBeginUpgradeNodeImageVersionOptions contains the optional parameters for the AgentPoolsClient.BeginUpgradeNodeImageVersion
// method.
type AgentPoolsClientBeginUpgradeNodeImageVersionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AgentPoolsClientGetAvailableAgentPoolVersionsOptions contains the optional parameters for the AgentPoolsClient.GetAvailableAgentPoolVersions
// method.
type AgentPoolsClientGetAvailableAgentPoolVersionsOptions struct {
	// placeholder for future optional parameters
}

// AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
type AgentPoolsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AgentPoolsClientGetUpgradeProfileOptions contains the optional parameters for the AgentPoolsClient.GetUpgradeProfile method.
type AgentPoolsClientGetUpgradeProfileOptions struct {
	// placeholder for future optional parameters
}

// AgentPoolsClientListOptions contains the optional parameters for the AgentPoolsClient.NewListPager method.
type AgentPoolsClientListOptions struct {
	// placeholder for future optional parameters
}

// MachineClientGetOptions contains the optional parameters for the MachineClient.Get method.
type MachineClientGetOptions struct {
	// placeholder for future optional parameters
}

// MachinesClientListOptions contains the optional parameters for the MachinesClient.NewListPager method.
type MachinesClientListOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the MaintenanceConfigurationsClient.CreateOrUpdate
// method.
type MaintenanceConfigurationsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceConfigurationsClientDeleteOptions contains the optional parameters for the MaintenanceConfigurationsClient.Delete
// method.
type MaintenanceConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceConfigurationsClientGetOptions contains the optional parameters for the MaintenanceConfigurationsClient.Get
// method.
type MaintenanceConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MaintenanceConfigurationsClientListByManagedClusterOptions contains the optional parameters for the MaintenanceConfigurationsClient.NewListByManagedClusterPager
// method.
type MaintenanceConfigurationsClientListByManagedClusterOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientCreateOrUpdateOptions contains the optional parameters for the ManagedClusterSnapshotsClient.CreateOrUpdate
// method.
type ManagedClusterSnapshotsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientDeleteOptions contains the optional parameters for the ManagedClusterSnapshotsClient.Delete
// method.
type ManagedClusterSnapshotsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientGetOptions contains the optional parameters for the ManagedClusterSnapshotsClient.Get method.
type ManagedClusterSnapshotsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientListByResourceGroupOptions contains the optional parameters for the ManagedClusterSnapshotsClient.NewListByResourceGroupPager
// method.
type ManagedClusterSnapshotsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientListOptions contains the optional parameters for the ManagedClusterSnapshotsClient.NewListPager
// method.
type ManagedClusterSnapshotsClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagedClusterSnapshotsClientUpdateTagsOptions contains the optional parameters for the ManagedClusterSnapshotsClient.UpdateTags
// method.
type ManagedClusterSnapshotsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientBeginAbortLatestOperationOptions contains the optional parameters for the ManagedClustersClient.BeginAbortLatestOperation
// method.
type ManagedClustersClientBeginAbortLatestOperationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedClustersClient.BeginCreateOrUpdate
// method.
type ManagedClustersClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginDeleteOptions contains the optional parameters for the ManagedClustersClient.BeginDelete method.
type ManagedClustersClientBeginDeleteOptions struct {
	// ignore-pod-disruption-budget=true to delete those pods on a node without considering Pod Disruption Budget
	IgnorePodDisruptionBudget *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginResetAADProfileOptions contains the optional parameters for the ManagedClustersClient.BeginResetAADProfile
// method.
type ManagedClustersClientBeginResetAADProfileOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginResetServicePrincipalProfileOptions contains the optional parameters for the ManagedClustersClient.BeginResetServicePrincipalProfile
// method.
type ManagedClustersClientBeginResetServicePrincipalProfileOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginRotateClusterCertificatesOptions contains the optional parameters for the ManagedClustersClient.BeginRotateClusterCertificates
// method.
type ManagedClustersClientBeginRotateClusterCertificatesOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions contains the optional parameters for the ManagedClustersClient.BeginRotateServiceAccountSigningKeys
// method.
type ManagedClustersClientBeginRotateServiceAccountSigningKeysOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginRunCommandOptions contains the optional parameters for the ManagedClustersClient.BeginRunCommand
// method.
type ManagedClustersClientBeginRunCommandOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginStartOptions contains the optional parameters for the ManagedClustersClient.BeginStart method.
type ManagedClustersClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginStopOptions contains the optional parameters for the ManagedClustersClient.BeginStop method.
type ManagedClustersClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientBeginUpdateTagsOptions contains the optional parameters for the ManagedClustersClient.BeginUpdateTags
// method.
type ManagedClustersClientBeginUpdateTagsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ManagedClustersClientGetAccessProfileOptions contains the optional parameters for the ManagedClustersClient.GetAccessProfile
// method.
type ManagedClustersClientGetAccessProfileOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientGetCommandResultOptions contains the optional parameters for the ManagedClustersClient.GetCommandResult
// method.
type ManagedClustersClientGetCommandResultOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientGetMeshRevisionProfileOptions contains the optional parameters for the ManagedClustersClient.GetMeshRevisionProfile
// method.
type ManagedClustersClientGetMeshRevisionProfileOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientGetMeshUpgradeProfileOptions contains the optional parameters for the ManagedClustersClient.GetMeshUpgradeProfile
// method.
type ManagedClustersClientGetMeshUpgradeProfileOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientGetOSOptionsOptions contains the optional parameters for the ManagedClustersClient.GetOSOptions method.
type ManagedClustersClientGetOSOptionsOptions struct {
	// The resource type for which the OS options needs to be returned
	ResourceType *string
}

// ManagedClustersClientGetOptions contains the optional parameters for the ManagedClustersClient.Get method.
type ManagedClustersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientGetUpgradeProfileOptions contains the optional parameters for the ManagedClustersClient.GetUpgradeProfile
// method.
type ManagedClustersClientGetUpgradeProfileOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListByResourceGroupOptions contains the optional parameters for the ManagedClustersClient.NewListByResourceGroupPager
// method.
type ManagedClustersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListClusterAdminCredentialsOptions contains the optional parameters for the ManagedClustersClient.ListClusterAdminCredentials
// method.
type ManagedClustersClientListClusterAdminCredentialsOptions struct {
	// server fqdn type for credentials to be returned
	ServerFqdn *string
}

// ManagedClustersClientListClusterMonitoringUserCredentialsOptions contains the optional parameters for the ManagedClustersClient.ListClusterMonitoringUserCredentials
// method.
type ManagedClustersClientListClusterMonitoringUserCredentialsOptions struct {
	// server fqdn type for credentials to be returned
	ServerFqdn *string
}

// ManagedClustersClientListClusterUserCredentialsOptions contains the optional parameters for the ManagedClustersClient.ListClusterUserCredentials
// method.
type ManagedClustersClientListClusterUserCredentialsOptions struct {
	// Only apply to AAD clusters, specifies the format of returned kubeconfig. Format 'azure' will return azure auth-provider
	// kubeconfig; format 'exec' will return exec format kubeconfig, which requires
	// kubelogin binary in the path.
	Format *Format

	// server fqdn type for credentials to be returned
	ServerFqdn *string
}

// ManagedClustersClientListKubernetesVersionsOptions contains the optional parameters for the ManagedClustersClient.ListKubernetesVersions
// method.
type ManagedClustersClientListKubernetesVersionsOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListMeshRevisionProfilesOptions contains the optional parameters for the ManagedClustersClient.NewListMeshRevisionProfilesPager
// method.
type ManagedClustersClientListMeshRevisionProfilesOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListMeshUpgradeProfilesOptions contains the optional parameters for the ManagedClustersClient.NewListMeshUpgradeProfilesPager
// method.
type ManagedClustersClientListMeshUpgradeProfilesOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListOptions contains the optional parameters for the ManagedClustersClient.NewListPager method.
type ManagedClustersClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagedClustersClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the ManagedClustersClient.NewListOutboundNetworkDependenciesEndpointsPager
// method.
type ManagedClustersClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
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

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Update
// method.
type PrivateEndpointConnectionsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.List method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// ResolvePrivateLinkServiceIDClientPOSTOptions contains the optional parameters for the ResolvePrivateLinkServiceIDClient.POST
// method.
type ResolvePrivateLinkServiceIDClientPOSTOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientCreateOrUpdateOptions contains the optional parameters for the SnapshotsClient.CreateOrUpdate method.
type SnapshotsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientDeleteOptions contains the optional parameters for the SnapshotsClient.Delete method.
type SnapshotsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientGetOptions contains the optional parameters for the SnapshotsClient.Get method.
type SnapshotsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientListByResourceGroupOptions contains the optional parameters for the SnapshotsClient.NewListByResourceGroupPager
// method.
type SnapshotsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.NewListPager method.
type SnapshotsClientListOptions struct {
	// placeholder for future optional parameters
}

// SnapshotsClientUpdateTagsOptions contains the optional parameters for the SnapshotsClient.UpdateTags method.
type SnapshotsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.BeginCreateOrUpdate
// method.
type TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrustedAccessRoleBindingsClientBeginDeleteOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.BeginDelete
// method.
type TrustedAccessRoleBindingsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrustedAccessRoleBindingsClientGetOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.Get
// method.
type TrustedAccessRoleBindingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TrustedAccessRoleBindingsClientListOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.NewListPager
// method.
type TrustedAccessRoleBindingsClientListOptions struct {
	// placeholder for future optional parameters
}

// TrustedAccessRolesClientListOptions contains the optional parameters for the TrustedAccessRolesClient.NewListPager method.
type TrustedAccessRolesClientListOptions struct {
	// placeholder for future optional parameters
}
