//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

// AgentPoolsClientAbortLatestOperationResponse contains the response from method AgentPoolsClient.BeginAbortLatestOperation.
type AgentPoolsClientAbortLatestOperationResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientCreateOrUpdateResponse contains the response from method AgentPoolsClient.BeginCreateOrUpdate.
type AgentPoolsClientCreateOrUpdateResponse struct {
	// Agent Pool.
	AgentPool
}

// AgentPoolsClientDeleteResponse contains the response from method AgentPoolsClient.BeginDelete.
type AgentPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientGetAvailableAgentPoolVersionsResponse contains the response from method AgentPoolsClient.GetAvailableAgentPoolVersions.
type AgentPoolsClientGetAvailableAgentPoolVersionsResponse struct {
	// The list of available versions for an agent pool.
	AgentPoolAvailableVersions
}

// AgentPoolsClientGetResponse contains the response from method AgentPoolsClient.Get.
type AgentPoolsClientGetResponse struct {
	// Agent Pool.
	AgentPool
}

// AgentPoolsClientGetUpgradeProfileResponse contains the response from method AgentPoolsClient.GetUpgradeProfile.
type AgentPoolsClientGetUpgradeProfileResponse struct {
	// The list of available upgrades for an agent pool.
	AgentPoolUpgradeProfile
}

// AgentPoolsClientListResponse contains the response from method AgentPoolsClient.NewListPager.
type AgentPoolsClientListResponse struct {
	// The response from the List Agent Pools operation.
	AgentPoolListResult
}

// AgentPoolsClientUpgradeNodeImageVersionResponse contains the response from method AgentPoolsClient.BeginUpgradeNodeImageVersion.
type AgentPoolsClientUpgradeNodeImageVersionResponse struct {
	// Agent Pool.
	AgentPool
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	// A machine. Contains details about the underlying virtual machine. A machine may be visible here but not in kubectl get
	// nodes; if so it may be because the machine has not been registered with the Kubernetes API Server yet.
	Machine
}

// MachinesClientListResponse contains the response from method MachinesClient.NewListPager.
type MachinesClientListResponse struct {
	// The response from the List Machines operation.
	MachineListResult
}

// MaintenanceConfigurationsClientCreateOrUpdateResponse contains the response from method MaintenanceConfigurationsClient.CreateOrUpdate.
type MaintenanceConfigurationsClientCreateOrUpdateResponse struct {
	// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned
	// maintenance.
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientDeleteResponse contains the response from method MaintenanceConfigurationsClient.Delete.
type MaintenanceConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// MaintenanceConfigurationsClientGetResponse contains the response from method MaintenanceConfigurationsClient.Get.
type MaintenanceConfigurationsClientGetResponse struct {
	// See [planned maintenance](https://docs.microsoft.com/azure/aks/planned-maintenance) for more information about planned
	// maintenance.
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientListByManagedClusterResponse contains the response from method MaintenanceConfigurationsClient.NewListByManagedClusterPager.
type MaintenanceConfigurationsClientListByManagedClusterResponse struct {
	// The response from the List maintenance configurations operation.
	MaintenanceConfigurationListResult
}

// ManagedClusterSnapshotsClientCreateOrUpdateResponse contains the response from method ManagedClusterSnapshotsClient.CreateOrUpdate.
type ManagedClusterSnapshotsClientCreateOrUpdateResponse struct {
	// A managed cluster snapshot resource.
	ManagedClusterSnapshot
}

// ManagedClusterSnapshotsClientDeleteResponse contains the response from method ManagedClusterSnapshotsClient.Delete.
type ManagedClusterSnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClusterSnapshotsClientGetResponse contains the response from method ManagedClusterSnapshotsClient.Get.
type ManagedClusterSnapshotsClientGetResponse struct {
	// A managed cluster snapshot resource.
	ManagedClusterSnapshot
}

// ManagedClusterSnapshotsClientListByResourceGroupResponse contains the response from method ManagedClusterSnapshotsClient.NewListByResourceGroupPager.
type ManagedClusterSnapshotsClientListByResourceGroupResponse struct {
	// The response from the List Managed Cluster Snapshots operation.
	ManagedClusterSnapshotListResult
}

// ManagedClusterSnapshotsClientListResponse contains the response from method ManagedClusterSnapshotsClient.NewListPager.
type ManagedClusterSnapshotsClientListResponse struct {
	// The response from the List Managed Cluster Snapshots operation.
	ManagedClusterSnapshotListResult
}

// ManagedClusterSnapshotsClientUpdateTagsResponse contains the response from method ManagedClusterSnapshotsClient.UpdateTags.
type ManagedClusterSnapshotsClientUpdateTagsResponse struct {
	// A managed cluster snapshot resource.
	ManagedClusterSnapshot
}

// ManagedClustersClientAbortLatestOperationResponse contains the response from method ManagedClustersClient.BeginAbortLatestOperation.
type ManagedClustersClientAbortLatestOperationResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientCreateOrUpdateResponse contains the response from method ManagedClustersClient.BeginCreateOrUpdate.
type ManagedClustersClientCreateOrUpdateResponse struct {
	// Managed cluster.
	ManagedCluster
}

// ManagedClustersClientDeleteResponse contains the response from method ManagedClustersClient.BeginDelete.
type ManagedClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientGetAccessProfileResponse contains the response from method ManagedClustersClient.GetAccessProfile.
type ManagedClustersClientGetAccessProfileResponse struct {
	// Managed cluster Access Profile.
	ManagedClusterAccessProfile
}

// ManagedClustersClientGetCommandResultResponse contains the response from method ManagedClustersClient.GetCommandResult.
type ManagedClustersClientGetCommandResultResponse struct {
	// run command result.
	RunCommandResult

	// Location contains the information returned from the Location header response.
	Location *string
}

// ManagedClustersClientGetGuardrailsVersionsResponse contains the response from method ManagedClustersClient.GetGuardrailsVersions.
type ManagedClustersClientGetGuardrailsVersionsResponse struct {
	// Available Guardrails Version
	GuardrailsAvailableVersion
}

// ManagedClustersClientGetMeshRevisionProfileResponse contains the response from method ManagedClustersClient.GetMeshRevisionProfile.
type ManagedClustersClientGetMeshRevisionProfileResponse struct {
	// Mesh revision profile for a mesh.
	MeshRevisionProfile
}

// ManagedClustersClientGetMeshUpgradeProfileResponse contains the response from method ManagedClustersClient.GetMeshUpgradeProfile.
type ManagedClustersClientGetMeshUpgradeProfileResponse struct {
	// Upgrade profile for given mesh.
	MeshUpgradeProfile
}

// ManagedClustersClientGetOSOptionsResponse contains the response from method ManagedClustersClient.GetOSOptions.
type ManagedClustersClientGetOSOptionsResponse struct {
	// The OS option profile.
	OSOptionProfile
}

// ManagedClustersClientGetResponse contains the response from method ManagedClustersClient.Get.
type ManagedClustersClientGetResponse struct {
	// Managed cluster.
	ManagedCluster
}

// ManagedClustersClientGetUpgradeProfileResponse contains the response from method ManagedClustersClient.GetUpgradeProfile.
type ManagedClustersClientGetUpgradeProfileResponse struct {
	// The list of available upgrades for compute pools.
	ManagedClusterUpgradeProfile
}

// ManagedClustersClientListByResourceGroupResponse contains the response from method ManagedClustersClient.NewListByResourceGroupPager.
type ManagedClustersClientListByResourceGroupResponse struct {
	// The response from the List Managed Clusters operation.
	ManagedClusterListResult
}

// ManagedClustersClientListClusterAdminCredentialsResponse contains the response from method ManagedClustersClient.ListClusterAdminCredentials.
type ManagedClustersClientListClusterAdminCredentialsResponse struct {
	// The list credential result response.
	CredentialResults
}

// ManagedClustersClientListClusterMonitoringUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterMonitoringUserCredentials.
type ManagedClustersClientListClusterMonitoringUserCredentialsResponse struct {
	// The list credential result response.
	CredentialResults
}

// ManagedClustersClientListClusterUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterUserCredentials.
type ManagedClustersClientListClusterUserCredentialsResponse struct {
	// The list credential result response.
	CredentialResults
}

// ManagedClustersClientListGuardrailsVersionsResponse contains the response from method ManagedClustersClient.NewListGuardrailsVersionsPager.
type ManagedClustersClientListGuardrailsVersionsResponse struct {
	// Hold values properties, which is array of GuardrailsVersions
	GuardrailsAvailableVersionsList
}

// ManagedClustersClientListKubernetesVersionsResponse contains the response from method ManagedClustersClient.ListKubernetesVersions.
type ManagedClustersClientListKubernetesVersionsResponse struct {
	// Hold values properties, which is array of KubernetesVersion
	KubernetesVersionListResult
}

// ManagedClustersClientListMeshRevisionProfilesResponse contains the response from method ManagedClustersClient.NewListMeshRevisionProfilesPager.
type ManagedClustersClientListMeshRevisionProfilesResponse struct {
	// Holds an array of MeshRevisionsProfiles
	MeshRevisionProfileList
}

// ManagedClustersClientListMeshUpgradeProfilesResponse contains the response from method ManagedClustersClient.NewListMeshUpgradeProfilesPager.
type ManagedClustersClientListMeshUpgradeProfilesResponse struct {
	// Holds an array of MeshUpgradeProfiles
	MeshUpgradeProfileList
}

// ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method ManagedClustersClient.NewListOutboundNetworkDependenciesEndpointsPager.
type ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse struct {
	// Collection of OutboundEnvironmentEndpoint
	OutboundEnvironmentEndpointCollection
}

// ManagedClustersClientListResponse contains the response from method ManagedClustersClient.NewListPager.
type ManagedClustersClientListResponse struct {
	// The response from the List Managed Clusters operation.
	ManagedClusterListResult
}

// ManagedClustersClientResetAADProfileResponse contains the response from method ManagedClustersClient.BeginResetAADProfile.
type ManagedClustersClientResetAADProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientResetServicePrincipalProfileResponse contains the response from method ManagedClustersClient.BeginResetServicePrincipalProfile.
type ManagedClustersClientResetServicePrincipalProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateClusterCertificatesResponse contains the response from method ManagedClustersClient.BeginRotateClusterCertificates.
type ManagedClustersClientRotateClusterCertificatesResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateServiceAccountSigningKeysResponse contains the response from method ManagedClustersClient.BeginRotateServiceAccountSigningKeys.
type ManagedClustersClientRotateServiceAccountSigningKeysResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRunCommandResponse contains the response from method ManagedClustersClient.BeginRunCommand.
type ManagedClustersClientRunCommandResponse struct {
	// run command result.
	RunCommandResult
}

// ManagedClustersClientStartResponse contains the response from method ManagedClustersClient.BeginStart.
type ManagedClustersClientStartResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientStopResponse contains the response from method ManagedClustersClient.BeginStop.
type ManagedClustersClientStopResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientUpdateTagsResponse contains the response from method ManagedClustersClient.BeginUpdateTags.
type ManagedClustersClientUpdateTagsResponse struct {
	// Managed cluster.
	ManagedCluster
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The List Operation response.
	OperationListResult
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

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	// A list of private endpoint connections
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	// A list of private link resources
	PrivateLinkResourcesListResult
}

// ResolvePrivateLinkServiceIDClientPOSTResponse contains the response from method ResolvePrivateLinkServiceIDClient.POST.
type ResolvePrivateLinkServiceIDClientPOSTResponse struct {
	// A private link resource
	PrivateLinkResource
}

// SnapshotsClientCreateOrUpdateResponse contains the response from method SnapshotsClient.CreateOrUpdate.
type SnapshotsClientCreateOrUpdateResponse struct {
	// A node pool snapshot resource.
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.Delete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	// A node pool snapshot resource.
	Snapshot
}

// SnapshotsClientListByResourceGroupResponse contains the response from method SnapshotsClient.NewListByResourceGroupPager.
type SnapshotsClientListByResourceGroupResponse struct {
	// The response from the List Snapshots operation.
	SnapshotListResult
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.NewListPager.
type SnapshotsClientListResponse struct {
	// The response from the List Snapshots operation.
	SnapshotListResult
}

// SnapshotsClientUpdateTagsResponse contains the response from method SnapshotsClient.UpdateTags.
type SnapshotsClientUpdateTagsResponse struct {
	// A node pool snapshot resource.
	Snapshot
}

// TrustedAccessRoleBindingsClientCreateOrUpdateResponse contains the response from method TrustedAccessRoleBindingsClient.BeginCreateOrUpdate.
type TrustedAccessRoleBindingsClientCreateOrUpdateResponse struct {
	// Defines binding between a resource and role
	TrustedAccessRoleBinding
}

// TrustedAccessRoleBindingsClientDeleteResponse contains the response from method TrustedAccessRoleBindingsClient.BeginDelete.
type TrustedAccessRoleBindingsClientDeleteResponse struct {
	// placeholder for future response values
}

// TrustedAccessRoleBindingsClientGetResponse contains the response from method TrustedAccessRoleBindingsClient.Get.
type TrustedAccessRoleBindingsClientGetResponse struct {
	// Defines binding between a resource and role
	TrustedAccessRoleBinding
}

// TrustedAccessRoleBindingsClientListResponse contains the response from method TrustedAccessRoleBindingsClient.NewListPager.
type TrustedAccessRoleBindingsClientListResponse struct {
	// List of trusted access role bindings
	TrustedAccessRoleBindingListResult
}

// TrustedAccessRolesClientListResponse contains the response from method TrustedAccessRolesClient.NewListPager.
type TrustedAccessRolesClientListResponse struct {
	// List of trusted access roles
	TrustedAccessRoleListResult
}

// UsagesClientListResponse contains the response from method UsagesClient.List.
type UsagesClientListResponse struct {
	// The List Usages operation response.
	ListUsagesResult
}
