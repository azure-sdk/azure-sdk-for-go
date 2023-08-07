//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import "time"

// ConfigDiagnostics - Configuration Diagnostics
type ConfigDiagnostics struct {
	// Last domain configuration diagnostics DateTime
	LastExecuted *time.Time

	// List of Configuration Diagnostics validator results.
	ValidatorResults []*ConfigDiagnosticsValidatorResult
}

// ConfigDiagnosticsValidatorResult - Config Diagnostics validator result data
type ConfigDiagnosticsValidatorResult struct {
	// List of resource config validation issues.
	Issues []*ConfigDiagnosticsValidatorResultIssue

	// Replica set location and subnet name
	ReplicaSetSubnetDisplayName *string

	// Status for individual validator after running diagnostics.
	Status *Status

	// Validator identifier
	ValidatorID *string
}

// ConfigDiagnosticsValidatorResultIssue - Specific issue for a particular config diagnostics validator
type ConfigDiagnosticsValidatorResultIssue struct {
	// List of domain resource property name or values used to compose a rich description.
	DescriptionParams []*string

	// Validation issue identifier.
	ID *string
}

// ContainerAccount - Container Account Description
type ContainerAccount struct {
	// The account name
	AccountName *string

	// The account password
	Password *string

	// The account spn
	Spn *string
}

// DomainSecuritySettings - Domain Security Settings
type DomainSecuritySettings struct {
	// A flag to determine whether or not ChannelBinding is enabled or disabled.
	ChannelBinding *ChannelBinding

	// A flag to determine whether or not KerberosArmoring is enabled or disabled.
	KerberosArmoring *KerberosArmoring

	// A flag to determine whether or not KerberosRc4Encryption is enabled or disabled.
	KerberosRc4Encryption *KerberosRc4Encryption

	// A flag to determine whether or not LdapSigning is enabled or disabled.
	LdapSigning *LdapSigning

	// A flag to determine whether or not NtlmV1 is enabled or disabled.
	NtlmV1 *NtlmV1

	// A flag to determine whether or not SyncKerberosPasswords is enabled or disabled.
	SyncKerberosPasswords *SyncKerberosPasswords

	// A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.
	SyncNtlmPasswords *SyncNtlmPasswords

	// A flag to determine whether or not SyncOnPremPasswords is enabled or disabled.
	SyncOnPremPasswords *SyncOnPremPasswords

	// A flag to determine whether or not TlsV1 is enabled or disabled.
	TLSV1 *TLSV1
}

// DomainService - Domain service.
type DomainService struct {
	// Resource etag
	Etag *string

	// Resource location
	Location *string

	// Domain service properties
	Properties *DomainServiceProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Resource type
	Type *string
}

// DomainServiceListResult - The response from the List Domain Services operation.
type DomainServiceListResult struct {
	// the list of domain services.
	Value []*DomainService

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// DomainServiceProperties - Properties of the Domain Service.
type DomainServiceProperties struct {
	// Configuration diagnostics data containing latest execution from client.
	ConfigDiagnostics *ConfigDiagnostics

	// Domain Configuration Type
	DomainConfigurationType *string

	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string

	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings

	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *FilteredSync

	// Secure LDAP Settings
	LdapsSettings *LdapsSettings

	// Notification Settings
	NotificationSettings *NotificationSettings

	// List of ReplicaSets
	ReplicaSets []*ReplicaSet

	// Resource Forest Settings
	ResourceForestSettings *ResourceForestSettings

	// Sku Type
	SKU *string

	// All or CloudOnly, All users in AAD are synced to AAD DS domain or only users actively syncing in the cloud
	SyncScope *SyncScope

	// READ-ONLY; Deployment Id
	DeploymentID *string

	// READ-ONLY; Migration Properties
	MigrationProperties *MigrationProperties

	// READ-ONLY; the current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string

	// READ-ONLY; The unique sync application id of the Azure AD Domain Services deployment.
	SyncApplicationID *string

	// READ-ONLY; SyncOwner ReplicaSet Id
	SyncOwner *string

	// READ-ONLY; Azure Active Directory Tenant Id
	TenantID *string

	// READ-ONLY; Data Model Version
	Version *int32
}

// ForestTrust - Forest Trust Setting
type ForestTrust struct {
	// Friendly Name
	FriendlyName *string

	// Remote Dns ips
	RemoteDNSIPs *string

	// Trust Direction
	TrustDirection *string

	// Trust Password
	TrustPassword *string

	// Trusted Domain FQDN
	TrustedDomainFqdn *string
}

// HealthAlert - Health Alert Description
type HealthAlert struct {
	// READ-ONLY; Health Alert Id
	ID *string

	// READ-ONLY; Health Alert Issue
	Issue *string

	// READ-ONLY; Health Alert Last Detected DateTime
	LastDetected *time.Time

	// READ-ONLY; Health Alert Name
	Name *string

	// READ-ONLY; Health Alert Raised DateTime
	Raised *time.Time

	// READ-ONLY; Health Alert TSG Link
	ResolutionURI *string

	// READ-ONLY; Health Alert Severity
	Severity *string
}

// HealthMonitor - Health Monitor Description
type HealthMonitor struct {
	// READ-ONLY; Health Monitor Details
	Details *string

	// READ-ONLY; Health Monitor Id
	ID *string

	// READ-ONLY; Health Monitor Name
	Name *string
}

// LdapsSettings - Secure LDAP Settings
type LdapsSettings struct {
	// A flag to determine whether or not Secure LDAP access over the internet is enabled or disabled.
	ExternalAccess *ExternalAccess

	// A flag to determine whether or not Secure LDAP is enabled or disabled.
	Ldaps *Ldaps

	// The certificate required to configure Secure LDAP. The parameter passed here should be a base64encoded representation of
	// the certificate pfx file.
	PfxCertificate *string

	// The password to decrypt the provided Secure LDAP certificate pfx file.
	PfxCertificatePassword *string

	// READ-ONLY; NotAfter DateTime of configure ldaps certificate.
	CertificateNotAfter *time.Time

	// READ-ONLY; Thumbprint of configure ldaps certificate.
	CertificateThumbprint *string

	// READ-ONLY; Public certificate used to configure secure ldap.
	PublicCertificate *string
}

// MigrationProgress - Migration Progress
type MigrationProgress struct {
	// Completion Percentage
	CompletionPercentage *float64

	// Progress Message
	ProgressMessage *string
}

// MigrationProperties - Migration Properties
type MigrationProperties struct {
	// READ-ONLY; Migration Progress
	MigrationProgress *MigrationProgress

	// READ-ONLY; Old Subnet Id
	OldSubnetID *string

	// READ-ONLY; Old Vnet Site Id
	OldVnetSiteID *string
}

// NotificationSettings - Settings for notification
type NotificationSettings struct {
	// The list of additional recipients
	AdditionalRecipients []*string

	// Should domain controller admins be notified
	NotifyDcAdmins *NotifyDcAdmins

	// Should global admins be notified
	NotifyGlobalAdmins *NotifyGlobalAdmins
}

// OperationDisplayInfo - The operation supported by Domain Services.
type OperationDisplayInfo struct {
	// The description of the operation.
	Description *string

	// The action that users can perform, based on their permission level.
	Operation *string

	// Service provider: Domain Services.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// OperationEntity - The operation supported by Domain Services.
type OperationEntity struct {
	// The operation supported by Domain Services.
	Display *OperationDisplayInfo

	// Operation name: {provider}/{resource}/{operation}.
	Name *string

	// The origin of the operation.
	Origin *string
}

// OperationEntityListResult - The list of domain service operation response.
type OperationEntityListResult struct {
	// The list of operations.
	Value []*OperationEntity

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// OuContainer - Resource for OuContainer.
type OuContainer struct {
	// Resource etag
	Etag *string

	// Resource location
	Location *string

	// OuContainer properties
	Properties *OuContainerProperties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Resource type
	Type *string
}

// OuContainerListResult - The response from the List OuContainer operation.
type OuContainerListResult struct {
	// The list of OuContainer.
	Value []*OuContainer

	// READ-ONLY; The continuation token for the next page of results.
	NextLink *string
}

// OuContainerProperties - Properties of the OuContainer.
type OuContainerProperties struct {
	// The list of container accounts
	Accounts []*ContainerAccount

	// READ-ONLY; The OuContainer name
	ContainerID *string

	// READ-ONLY; The Deployment id
	DeploymentID *string

	// READ-ONLY; Distinguished Name of OuContainer instance
	DistinguishedName *string

	// READ-ONLY; The domain name of Domain Services.
	DomainName *string

	// READ-ONLY; The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string

	// READ-ONLY; Status of OuContainer instance
	ServiceStatus *string

	// READ-ONLY; Azure Active Directory tenant id
	TenantID *string
}

// ReplicaSet - Replica Set Definition
type ReplicaSet struct {
	// Virtual network location
	Location *string

	// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will
	// be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetID *string

	// READ-ONLY; List of Domain Controller IP Address
	DomainControllerIPAddress []*string

	// READ-ONLY; External access ip address.
	ExternalAccessIPAddress *string

	// READ-ONLY; List of Domain Health Alerts
	HealthAlerts []*HealthAlert

	// READ-ONLY; Last domain evaluation run DateTime
	HealthLastEvaluated *time.Time

	// READ-ONLY; List of Domain Health Monitors
	HealthMonitors []*HealthMonitor

	// READ-ONLY; ReplicaSet Id
	ReplicaSetID *string

	// READ-ONLY; Status of Domain Service instance
	ServiceStatus *string

	// READ-ONLY; Virtual network site id
	VnetSiteID *string
}

// Resource - The Resource model definition.
type Resource struct {
	// Resource etag
	Etag *string

	// Resource location
	Location *string

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Resource type
	Type *string
}

// ResourceForestSettings - Settings for Resource Forest
type ResourceForestSettings struct {
	// Resource Forest
	ResourceForest *string

	// List of settings for Resource Forest
	Settings []*ForestTrust
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
