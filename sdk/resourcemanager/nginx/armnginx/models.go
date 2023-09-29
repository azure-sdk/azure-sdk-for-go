//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnginx

import "time"

type Certificate struct {
	Location   *string
	Properties *CertificateProperties

	// Dictionary of
	Tags map[string]*string

	// READ-ONLY
	ID *string

	// READ-ONLY
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY
	Type *string
}

type CertificateListResponse struct {
	NextLink *string
	Value    []*Certificate
}

type CertificateProperties struct {
	CertificateVirtualPath *string
	KeyVaultSecretID       *string
	KeyVirtualPath         *string

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

type Configuration struct {
	Location   *string
	Properties *ConfigurationProperties

	// Dictionary of
	Tags map[string]*string

	// READ-ONLY
	ID *string

	// READ-ONLY
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY
	Type *string
}

type ConfigurationFile struct {
	Content     *string
	VirtualPath *string
}

// ConfigurationListResponse - Response of a list operation.
type ConfigurationListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*Configuration
}

type ConfigurationPackage struct {
	Data *string
}

type ConfigurationProperties struct {
	Files          []*ConfigurationFile
	Package        *ConfigurationPackage
	ProtectedFiles []*ConfigurationFile
	RootFile       *string

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

type Deployment struct {
	Identity   *IdentityProperties
	Location   *string
	Properties *DeploymentProperties
	SKU        *ResourceSKU

	// Dictionary of
	Tags map[string]*string

	// READ-ONLY
	ID *string

	// READ-ONLY
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY
	Type *string
}

type DeploymentListResponse struct {
	NextLink *string
	Value    []*Deployment
}

type DeploymentProperties struct {
	EnableDiagnosticsSupport *bool
	Logging                  *Logging

	// The managed resource group to deploy VNet injection related network resources.
	ManagedResourceGroup *string
	NetworkProfile       *NetworkProfile

	// READ-ONLY; The IP address of the deployment.
	IPAddress *string

	// READ-ONLY
	NginxVersion *string

	// READ-ONLY
	ProvisioningState *ProvisioningState
}

type DeploymentUpdateParameters struct {
	Identity   *IdentityProperties
	Location   *string
	Properties *DeploymentUpdateProperties
	SKU        *ResourceSKU

	// Dictionary of
	Tags map[string]*string
}

type DeploymentUpdateProperties struct {
	EnableDiagnosticsSupport *bool
	Logging                  *Logging
}

type ErrorResponseBody struct {
	Code    *string
	Details []*ErrorResponseBody
	Message *string
	Target  *string
}

type FrontendIPConfiguration struct {
	PrivateIPAddresses []*PrivateIPAddress
	PublicIPAddresses  []*PublicIPAddress
}

type IdentityProperties struct {
	Type *IdentityType

	// Dictionary of
	UserAssignedIdentities map[string]*UserIdentityProperties

	// READ-ONLY
	PrincipalID *string

	// READ-ONLY
	TenantID *string
}

type Logging struct {
	StorageAccount *StorageAccount
}

type NetworkInterfaceConfiguration struct {
	SubnetID *string
}

type NetworkProfile struct {
	FrontEndIPConfiguration       *FrontendIPConfiguration
	NetworkInterfaceConfiguration *NetworkInterfaceConfiguration
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write deployments'.
	Description *string

	// Operation type, e.g., read, write, delete, etc.
	Operation *string

	// Service provider: Nginx.NginxPlus
	Provider *string

	// Type on which the operation is performed, e.g., 'deployments'.
	Resource *string
}

// OperationListResult - Result of GET request to list Nginx.NginxPlus operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of operations supported by the Nginx.NginxPlus provider.
	Value []*OperationResult
}

// OperationResult - A Nginx.NginxPlus REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

type PrivateIPAddress struct {
	PrivateIPAddress          *string
	PrivateIPAllocationMethod *NginxPrivateIPAllocationMethod
	SubnetID                  *string
}

type PublicIPAddress struct {
	ID *string
}

type ResourceProviderDefaultErrorResponse struct {
	Error *ErrorResponseBody
}

type ResourceSKU struct {
	// REQUIRED; Name of the SKU.
	Name *string
}

type StorageAccount struct {
	AccountName   *string
	ContainerName *string
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

type UserIdentityProperties struct {
	// READ-ONLY
	ClientID *string

	// READ-ONLY
	PrincipalID *string
}
