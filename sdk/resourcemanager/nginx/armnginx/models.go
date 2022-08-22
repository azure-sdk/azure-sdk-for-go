//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnginx

import "time"

type Certificate struct {
	Location   *string                `json:"location,omitempty"`
	Properties *CertificateProperties `json:"properties,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY
	Type *string `json:"type,omitempty" azure:"ro"`
}

type CertificateListResponse struct {
	NextLink *string        `json:"nextLink,omitempty"`
	Value    []*Certificate `json:"value,omitempty"`
}

type CertificateProperties struct {
	CertificateVirtualPath *string            `json:"certificateVirtualPath,omitempty"`
	KeyVaultSecretID       *string            `json:"keyVaultSecretId,omitempty"`
	KeyVirtualPath         *string            `json:"keyVirtualPath,omitempty"`
	ProvisioningState      *ProvisioningState `json:"provisioningState,omitempty"`
}

// CertificatesClientBeginCreateOptions contains the optional parameters for the CertificatesClient.BeginCreate method.
type CertificatesClientBeginCreateOptions struct {
	// The certificate
	Body *Certificate
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CertificatesClientBeginDeleteOptions contains the optional parameters for the CertificatesClient.BeginDelete method.
type CertificatesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CertificatesClientGetOptions contains the optional parameters for the CertificatesClient.Get method.
type CertificatesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CertificatesClientListOptions contains the optional parameters for the CertificatesClient.List method.
type CertificatesClientListOptions struct {
	// placeholder for future optional parameters
}

type Configuration struct {
	Location   *string                  `json:"location,omitempty"`
	Properties *ConfigurationProperties `json:"properties,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY
	Type *string `json:"type,omitempty" azure:"ro"`
}

type ConfigurationFile struct {
	Content     *string `json:"content,omitempty"`
	VirtualPath *string `json:"virtualPath,omitempty"`
}

// ConfigurationListResponse - Response of a list operation.
type ConfigurationListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// Results of a list operation.
	Value []*Configuration `json:"value,omitempty"`
}

type ConfigurationPackage struct {
	Data *string `json:"data,omitempty"`
}

type ConfigurationProperties struct {
	Files             []*ConfigurationFile  `json:"files,omitempty"`
	Package           *ConfigurationPackage `json:"package,omitempty"`
	ProvisioningState *ProvisioningState    `json:"provisioningState,omitempty"`
	RootFile          *string               `json:"rootFile,omitempty"`
}

// ConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginCreateOrUpdate
// method.
type ConfigurationsClientBeginCreateOrUpdateOptions struct {
	// The Nginx configuration
	Body *Configuration
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientBeginDeleteOptions contains the optional parameters for the ConfigurationsClient.BeginDelete method.
type ConfigurationsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
type ConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientListOptions contains the optional parameters for the ConfigurationsClient.List method.
type ConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

type Deployment struct {
	Identity   *IdentityProperties   `json:"identity,omitempty"`
	Location   *string               `json:"location,omitempty"`
	Properties *DeploymentProperties `json:"properties,omitempty"`
	SKU        *ResourceSKU          `json:"sku,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY
	Type *string `json:"type,omitempty" azure:"ro"`
}

type DeploymentListResponse struct {
	NextLink *string       `json:"nextLink,omitempty"`
	Value    []*Deployment `json:"value,omitempty"`
}

type DeploymentProperties struct {
	EnableDiagnosticsSupport *bool `json:"enableDiagnosticsSupport,omitempty"`

	// The managed resource group to deploy VNet injection related network resources.
	ManagedResourceGroup *string            `json:"managedResourceGroup,omitempty"`
	NetworkProfile       *NetworkProfile    `json:"networkProfile,omitempty"`
	ProvisioningState    *ProvisioningState `json:"provisioningState,omitempty"`

	// READ-ONLY; The IP address of the deployment.
	IPAddress *string `json:"ipAddress,omitempty" azure:"ro"`

	// READ-ONLY
	NginxVersion *string `json:"nginxVersion,omitempty" azure:"ro"`
}

type DeploymentUpdateParameters struct {
	Identity   *IdentityProperties         `json:"identity,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Properties *DeploymentUpdateProperties `json:"properties,omitempty"`
	SKU        *ResourceSKU                `json:"sku,omitempty"`

	// Dictionary of
	Tags map[string]*string `json:"tags,omitempty"`
}

type DeploymentUpdateProperties struct {
	EnableDiagnosticsSupport *bool `json:"enableDiagnosticsSupport,omitempty"`
}

// DeploymentsClientBeginCreateOptions contains the optional parameters for the DeploymentsClient.BeginCreate method.
type DeploymentsClientBeginCreateOptions struct {
	Body *Deployment
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientBeginDeleteOptions contains the optional parameters for the DeploymentsClient.BeginDelete method.
type DeploymentsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientBeginUpdateOptions contains the optional parameters for the DeploymentsClient.BeginUpdate method.
type DeploymentsClientBeginUpdateOptions struct {
	Body *DeploymentUpdateParameters
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DeploymentsClientGetOptions contains the optional parameters for the DeploymentsClient.Get method.
type DeploymentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientListByResourceGroupOptions contains the optional parameters for the DeploymentsClient.ListByResourceGroup
// method.
type DeploymentsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DeploymentsClientListOptions contains the optional parameters for the DeploymentsClient.List method.
type DeploymentsClientListOptions struct {
	// placeholder for future optional parameters
}

type ErrorResponseBody struct {
	Code    *string              `json:"code,omitempty"`
	Details []*ErrorResponseBody `json:"details,omitempty"`
	Message *string              `json:"message,omitempty"`
	Target  *string              `json:"target,omitempty"`
}

type FrontendIPConfiguration struct {
	PrivateIPAddresses []*PrivateIPAddress `json:"privateIPAddresses,omitempty"`
	PublicIPAddresses  []*PublicIPAddress  `json:"publicIPAddresses,omitempty"`
}

type IdentityProperties struct {
	Type *IdentityType `json:"type,omitempty"`

	// Dictionary of
	UserAssignedIdentities map[string]*UserIdentityProperties `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

type NetworkInterfaceConfiguration struct {
	SubnetID *string `json:"subnetId,omitempty"`
}

type NetworkProfile struct {
	FrontEndIPConfiguration       *FrontendIPConfiguration       `json:"frontEndIPConfiguration,omitempty"`
	NetworkInterfaceConfiguration *NetworkInterfaceConfiguration `json:"networkInterfaceConfiguration,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write deployments'.
	Description *string `json:"description,omitempty"`

	// Operation type, e.g., read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Nginx.NginxPlus
	Provider *string `json:"provider,omitempty"`

	// Type on which the operation is performed, e.g., 'deployments'.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of GET request to list Nginx.NginxPlus operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by the Nginx.NginxPlus provider.
	Value []*OperationResult `json:"value,omitempty"`
}

// OperationResult - A Nginx.NginxPlus REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is a data action
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

type PrivateIPAddress struct {
	PrivateIPAddress          *string                         `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *NginxPrivateIPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`
	SubnetID                  *string                         `json:"subnetId,omitempty"`
}

type PublicIPAddress struct {
	ID *string `json:"id,omitempty"`
}

type ResourceProviderDefaultErrorResponse struct {
	Error *ErrorResponseBody `json:"error,omitempty"`
}

type ResourceSKU struct {
	// REQUIRED; Name of the SKU.
	Name *string `json:"name,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

type UserIdentityProperties struct {
	// READ-ONLY
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}
