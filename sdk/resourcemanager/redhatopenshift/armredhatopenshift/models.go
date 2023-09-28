//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

import "time"

// APIServerProfile represents an API server profile.
type APIServerProfile struct {
	// The IP of the cluster API server.
	IP *string

	// The URL to access the cluster API server.
	URL *string

	// API server visibility.
	Visibility *Visibility
}

// CloudError represents a cloud error.
type CloudError struct {
	// An error response from the service.
	Error *CloudErrorBody
}

// CloudErrorBody represents the body of a cloud error.
type CloudErrorBody struct {
	// An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string

	// A list of additional details about the error.
	Details []*CloudErrorBody

	// A message describing the error, intended to be suitable for display in a user interface.
	Message *string

	// The target of the particular error. For example, the name of the property in error.
	Target *string
}

// ClusterProfile represents a cluster profile.
type ClusterProfile struct {
	// The domain for the cluster.
	Domain *string

	// If FIPS validated crypto modules are used
	FipsValidatedModules *FipsValidatedModules

	// The pull secret for the cluster.
	PullSecret *string

	// The ID of the cluster resource group.
	ResourceGroupID *string

	// The version of the cluster.
	Version *string
}

// ConsoleProfile represents a console profile.
type ConsoleProfile struct {
	// The URL to access the cluster console.
	URL *string
}

// Display represents the display details of an operation.
type Display struct {
	// Friendly name of the operation.
	Description *string

	// Operation type: read, write, delete, listKeys/action, etc.
	Operation *string

	// Friendly name of the resource provider.
	Provider *string

	// Resource type on which the operation is performed.
	Resource *string
}

// IngressProfile represents an ingress profile.
type IngressProfile struct {
	// The IP of the ingress.
	IP *string

	// The ingress profile name.
	Name *string

	// Ingress visibility.
	Visibility *Visibility
}

// MachinePool represents a MachinePool
type MachinePool struct {
	// The MachinePool Properties
	Properties *MachinePoolProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// MachinePoolList represents a list of MachinePools
type MachinePoolList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of Machine Pools.
	Value []*MachinePool
}

// MachinePoolProperties represents the properties of a MachinePool
type MachinePoolProperties struct {
	Resources *string
}

// MachinePoolUpdate - MachinePool represents a MachinePool
type MachinePoolUpdate struct {
	// The MachinePool Properties
	Properties *MachinePoolProperties

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData
}

// MasterProfile represents a master profile.
type MasterProfile struct {
	// The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetID *string

	// Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost

	// The Azure resource ID of the master subnet.
	SubnetID *string

	// The size of the master VMs.
	VMSize *string
}

// NetworkProfile represents a network profile.
type NetworkProfile struct {
	// The OutboundType used for egress traffic.
	OutboundType *OutboundType

	// The CIDR used for OpenShift/Kubernetes Pods.
	PodCidr *string

	// Specifies whether subnets are pre-attached with an NSG
	PreconfiguredNSG *PreconfiguredNSG

	// The CIDR used for OpenShift/Kubernetes Services.
	ServiceCidr *string
}

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The cluster properties.
	Properties *OpenShiftClusterProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OpenShiftClusterAdminKubeconfig represents an OpenShift cluster's admin kubeconfig.
type OpenShiftClusterAdminKubeconfig struct {
	// The base64-encoded kubeconfig file.
	Kubeconfig *string
}

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
type OpenShiftClusterCredentials struct {
	// The password for the kubeadmin user.
	KubeadminPassword *string

	// The username for the kubeadmin user.
	KubeadminUsername *string
}

// OpenShiftClusterList represents a list of OpenShift clusters.
type OpenShiftClusterList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of OpenShift clusters.
	Value []*OpenShiftCluster
}

// OpenShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties struct {
	// The cluster API server profile.
	ApiserverProfile *APIServerProfile

	// The cluster profile.
	ClusterProfile *ClusterProfile

	// The console profile.
	ConsoleProfile *ConsoleProfile

	// The cluster ingress profiles.
	IngressProfiles []*IngressProfile

	// The cluster master profile.
	MasterProfile *MasterProfile

	// The cluster network profile.
	NetworkProfile *NetworkProfile

	// The cluster provisioning state.
	ProvisioningState *ProvisioningState

	// The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile

	// The cluster worker profiles.
	WorkerProfiles []*WorkerProfile

	// READ-ONLY; The cluster worker profiles status.
	WorkerProfilesStatus []*WorkerProfile
}

// OpenShiftClusterUpdate - OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftClusterUpdate struct {
	// The cluster properties.
	Properties *OpenShiftClusterProperties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData
}

// OpenShiftVersion represents an OpenShift version that can be installed.
type OpenShiftVersion struct {
	// The properties for the OpenShiftVersion resource.
	Properties *OpenShiftVersionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// OpenShiftVersionList represents a List of available versions.
type OpenShiftVersionList struct {
	// Next Link to next operation.
	NextLink *string

	// The List of available versions.
	Value []*OpenShiftVersion
}

// OpenShiftVersionProperties represents the properties of an OpenShiftVersion.
type OpenShiftVersionProperties struct {
	// Version represents the version to create the cluster at.
	Version *string
}

// Operation represents an RP operation.
type Operation struct {
	// The object that describes the operation.
	Display *Display

	// Operation name: {provider}/{resource}/{operation}.
	Name *string

	// Sources of requests to this operation. Comma separated list with valid values user or system, e.g. "user,system".
	Origin *string
}

// OperationList represents an RP operation list.
type OperationList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// List of operations supported by the resource provider.
	Value []*Operation
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Secret represents a secret.
type Secret struct {
	// The Secret Properties
	Properties *SecretProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SecretList represents a list of Secrets
type SecretList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of secrets.
	Value []*Secret
}

// SecretProperties represents the properties of a Secret
type SecretProperties struct {
	// The Secrets Resources.
	SecretResources *string
}

// SecretUpdate - Secret represents a secret.
type SecretUpdate struct {
	// The Secret Properties
	Properties *SecretProperties

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData
}

// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile struct {
	// The client ID used for the cluster.
	ClientID *string

	// The client secret used for the cluster.
	ClientSecret *string
}

// SyncIdentityProvider represents a SyncIdentityProvider
type SyncIdentityProvider struct {
	// The SyncIdentityProvider Properties
	Properties *SyncIdentityProviderProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SyncIdentityProviderList - SyncSetList represents a list of SyncSets
type SyncIdentityProviderList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of sync identity providers
	Value []*SyncIdentityProvider
}

// SyncIdentityProviderProperties - SyncSetProperties represents the properties of a SyncSet
type SyncIdentityProviderProperties struct {
	Resources *string
}

// SyncIdentityProviderUpdate - SyncIdentityProvider represents a SyncIdentityProvider
type SyncIdentityProviderUpdate struct {
	// The SyncIdentityProvider Properties
	Properties *SyncIdentityProviderProperties

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData
}

// SyncSet represents a SyncSet for an Azure Red Hat OpenShift Cluster.
type SyncSet struct {
	// The Syncsets properties
	Properties *SyncSetProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SyncSetList represents a list of SyncSets
type SyncSetList struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of syncsets.
	Value []*SyncSet
}

// SyncSetProperties represents the properties of a SyncSet
type SyncSetProperties struct {
	// Resources represents the SyncSets configuration.
	Resources *string
}

// SyncSetUpdate - SyncSet represents a SyncSet for an Azure Red Hat OpenShift Cluster.
type SyncSetUpdate struct {
	// The Syncsets properties
	Properties *SyncSetProperties

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WorkerProfile represents a worker profile.
type WorkerProfile struct {
	// The number of worker VMs.
	Count *int32

	// The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetID *string

	// The disk size of the worker VMs.
	DiskSizeGB *int32

	// Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost

	// The worker profile name.
	Name *string

	// The Azure resource ID of the worker subnet.
	SubnetID *string

	// The size of the worker VMs.
	VMSize *string
}
