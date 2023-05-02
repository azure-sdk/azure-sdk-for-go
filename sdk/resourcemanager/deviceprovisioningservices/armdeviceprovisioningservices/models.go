//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceprovisioningservices

import "time"

// AsyncOperationResult - Result of a long running operation.
type AsyncOperationResult struct {
	// Error message containing code, description and details
	Error *ErrorMessage

	// current status of a long running operation.
	Status *string
}

// CertificateBodyDescription - The JSON-serialized X509 Certificate.
type CertificateBodyDescription struct {
	// Base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
	Certificate *string

	// True indicates that the certificate will be created in verified state and proof of possession will not be required.
	IsVerified *bool
}

// CertificateListDescription - The JSON-serialized array of Certificate objects.
type CertificateListDescription struct {
	// The array of Certificate objects.
	Value []*CertificateResponse
}

// CertificateProperties - The description of an X509 CA Certificate.
type CertificateProperties struct {
	// base-64 representation of X509 certificate .cer file or just .pem file content.
	Certificate []byte

	// Determines whether certificate has been verified.
	IsVerified *bool

	// READ-ONLY; The certificate's creation date and time.
	Created *time.Time

	// READ-ONLY; The certificate's expiration date and time.
	Expiry *time.Time

	// READ-ONLY; The certificate's subject name.
	Subject *string

	// READ-ONLY; The certificate's thumbprint.
	Thumbprint *string

	// READ-ONLY; The certificate's last update date and time.
	Updated *time.Time
}

// CertificateResponse - The X509 Certificate.
type CertificateResponse struct {
	// properties of a certificate
	Properties *CertificateProperties

	// READ-ONLY; The entity tag.
	Etag *string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The name of the certificate.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// CustomerInitiatedFailoverInput - The DPS failover input.
type CustomerInitiatedFailoverInput struct {
	// REQUIRED; Region that failover is porting to
	FailoverRegion *string
}

// DpsCertificateClientCreateOrUpdateOptions contains the optional parameters for the DpsCertificateClient.CreateOrUpdate
// method.
type DpsCertificateClientCreateOrUpdateOptions struct {
	// ETag of the certificate. This is required to update an existing certificate, and ignored while creating a brand new certificate.
	IfMatch *string
}

// DpsCertificateClientDeleteOptions contains the optional parameters for the DpsCertificateClient.Delete method.
type DpsCertificateClientDeleteOptions struct {
	// Time the certificate is created.
	CertificateCreated *time.Time
	// Indicates if the certificate contains a private key.
	CertificateHasPrivateKey *bool
	// Indicates if certificate has been verified by owner of the private key.
	CertificateIsVerified *bool
	// Time the certificate is last updated.
	CertificateLastUpdated *time.Time
	// This is optional, and it is the Common Name of the certificate.
	CertificateName1 *string
	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string
	// A description that mentions the purpose of the certificate.
	CertificatePurpose *CertificatePurpose
	// Raw data within the certificate.
	CertificateRawBytes []byte
}

// DpsCertificateClientGenerateVerificationCodeOptions contains the optional parameters for the DpsCertificateClient.GenerateVerificationCode
// method.
type DpsCertificateClientGenerateVerificationCodeOptions struct {
	// Certificate creation time.
	CertificateCreated *time.Time
	// Indicates if the certificate contains private key.
	CertificateHasPrivateKey *bool
	// Indicates if the certificate has been verified by owner of the private key.
	CertificateIsVerified *bool
	// Certificate last updated time.
	CertificateLastUpdated *time.Time
	// Common Name for the certificate.
	CertificateName1 *string
	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string
	// Description mentioning the purpose of the certificate.
	CertificatePurpose *CertificatePurpose
	// Raw data of certificate.
	CertificateRawBytes []byte
}

// DpsCertificateClientGetOptions contains the optional parameters for the DpsCertificateClient.Get method.
type DpsCertificateClientGetOptions struct {
	// ETag of the certificate.
	IfMatch *string
}

// DpsCertificateClientListOptions contains the optional parameters for the DpsCertificateClient.List method.
type DpsCertificateClientListOptions struct {
	// placeholder for future optional parameters
}

// DpsCertificateClientVerifyCertificateOptions contains the optional parameters for the DpsCertificateClient.VerifyCertificate
// method.
type DpsCertificateClientVerifyCertificateOptions struct {
	// Certificate creation time.
	CertificateCreated *time.Time
	// Indicates if the certificate contains private key.
	CertificateHasPrivateKey *bool
	// Indicates if the certificate has been verified by owner of the private key.
	CertificateIsVerified *bool
	// Certificate last updated time.
	CertificateLastUpdated *time.Time
	// Common Name for the certificate.
	CertificateName1 *string
	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string
	// Describe the purpose of the certificate.
	CertificatePurpose *CertificatePurpose
	// Raw data of certificate.
	CertificateRawBytes []byte
}

// ErrorDetails - Error details.
type ErrorDetails struct {
	// READ-ONLY; The error code.
	Code *int32

	// READ-ONLY; The error details.
	Details *string

	// READ-ONLY; The HTTP status code.
	HTTPStatusCode *string

	// READ-ONLY; The error message.
	Message *string
}

// ErrorMessage - Error response containing message and code.
type ErrorMessage struct {
	// standard error code
	Code *string

	// detailed summary of error
	Details *string

	// standard error description
	Message *string
}

// GroupIDInformation - The group information for creating a private endpoint on a provisioning service
type GroupIDInformation struct {
	// REQUIRED; The properties for a group information object
	Properties *GroupIDInformationProperties

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// GroupIDInformationProperties - The properties for a group information object
type GroupIDInformationProperties struct {
	// The group id
	GroupID *string

	// The required members for a specific group id
	RequiredMembers []*string

	// The required DNS zones for a specific group id
	RequiredZoneNames []*string
}

// IPFilterRule - The IP filter rules for a provisioning Service.
type IPFilterRule struct {
	// REQUIRED; The desired action for requests captured by this rule.
	Action *IPFilterActionType

	// REQUIRED; The name of the IP filter rule.
	FilterName *string

	// REQUIRED; A string that contains the IP address range in CIDR notation for the rule.
	IPMask *string

	// Target for requests captured by this rule.
	Target *IPFilterTargetType
}

// IotDpsPropertiesDescription - the service specific properties of a provisioning service, including keys, linked iot hubs,
// current state, and system generated properties such as hostname and idScope
type IotDpsPropertiesDescription struct {
	// Allocation policy to be used by this provisioning service.
	AllocationPolicy *AllocationPolicy

	// List of authorization keys for a provisioning service.
	AuthorizationPolicies []*SharedAccessSignatureAuthorizationRuleAccessRightsDescription

	// The DPS failover input.
	DpsFailoverDescription *IotDpsPropertiesDescriptionDpsFailoverDescription

	// Indicates if the DPS instance has Customer Enabled Failover enabled.
	EnableCustomerInitiatedFailover *bool

	// The IP filter rules.
	IPFilterRules []*IPFilterRule

	// List of IoT hubs associated with this provisioning service.
	IotHubs []*IotHubDefinitionDescription

	// Portal endpoint to enable CORS for this provisioning service.
	PortalOperationsHostName *string

	// Private endpoint connections created on this IotHub
	PrivateEndpointConnections []*PrivateEndpointConnection

	// The ARM provisioning state of the provisioning service.
	ProvisioningState *string

	// Whether requests from Public Network are allowed
	PublicNetworkAccess *PublicNetworkAccess

	// Current state of the provisioning service.
	State *State

	// READ-ONLY; Device endpoint for this provisioning service.
	DeviceProvisioningHostName *string

	// READ-ONLY; Unique identifier of this provisioning service.
	IDScope *string

	// READ-ONLY; Service endpoint for provisioning service.
	ServiceOperationsHostName *string
}

// IotDpsPropertiesDescriptionDpsFailoverDescription - The DPS failover input.
type IotDpsPropertiesDescriptionDpsFailoverDescription struct {
	// REQUIRED; Region that failover is porting to
	FailoverRegion *string
}

// IotDpsResourceClientBeginCreateOrUpdateOptions contains the optional parameters for the IotDpsResourceClient.BeginCreateOrUpdate
// method.
type IotDpsResourceClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.BeginCreateOrUpdatePrivateEndpointConnection
// method.
type IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginDeleteOptions contains the optional parameters for the IotDpsResourceClient.BeginDelete method.
type IotDpsResourceClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.BeginDeletePrivateEndpointConnection
// method.
type IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginUpdateOptions contains the optional parameters for the IotDpsResourceClient.BeginUpdate method.
type IotDpsResourceClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions contains the optional parameters for the IotDpsResourceClient.CheckProvisioningServiceNameAvailability
// method.
type IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientFailoverOptions contains the optional parameters for the IotDpsResourceClient.Failover method.
type IotDpsResourceClientFailoverOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetOperationResultOptions contains the optional parameters for the IotDpsResourceClient.GetOperationResult
// method.
type IotDpsResourceClientGetOperationResultOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetOptions contains the optional parameters for the IotDpsResourceClient.Get method.
type IotDpsResourceClientGetOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetPrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.GetPrivateEndpointConnection
// method.
type IotDpsResourceClientGetPrivateEndpointConnectionOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetPrivateLinkResourcesOptions contains the optional parameters for the IotDpsResourceClient.GetPrivateLinkResources
// method.
type IotDpsResourceClientGetPrivateLinkResourcesOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListByResourceGroupOptions contains the optional parameters for the IotDpsResourceClient.NewListByResourceGroupPager
// method.
type IotDpsResourceClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListBySubscriptionOptions contains the optional parameters for the IotDpsResourceClient.NewListBySubscriptionPager
// method.
type IotDpsResourceClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListKeysForKeyNameOptions contains the optional parameters for the IotDpsResourceClient.ListKeysForKeyName
// method.
type IotDpsResourceClientListKeysForKeyNameOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListKeysOptions contains the optional parameters for the IotDpsResourceClient.NewListKeysPager method.
type IotDpsResourceClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListPrivateEndpointConnectionsOptions contains the optional parameters for the IotDpsResourceClient.ListPrivateEndpointConnections
// method.
type IotDpsResourceClientListPrivateEndpointConnectionsOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListPrivateLinkResourcesOptions contains the optional parameters for the IotDpsResourceClient.ListPrivateLinkResources
// method.
type IotDpsResourceClientListPrivateLinkResourcesOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListValidSKUsOptions contains the optional parameters for the IotDpsResourceClient.NewListValidSKUsPager
// method.
type IotDpsResourceClientListValidSKUsOptions struct {
	// placeholder for future optional parameters
}

// IotDpsSKUDefinition - Available SKUs of tier and units.
type IotDpsSKUDefinition struct {
	// Sku name.
	Name *IotDpsSKU
}

// IotDpsSKUDefinitionListResult - List of available SKUs.
type IotDpsSKUDefinitionListResult struct {
	// The list of SKUs
	Value []*IotDpsSKUDefinition

	// READ-ONLY; The next link.
	NextLink *string
}

// IotDpsSKUInfo - List of possible provisioning service SKUs.
type IotDpsSKUInfo struct {
	// The number of units to provision
	Capacity *int64

	// Sku name.
	Name *IotDpsSKU

	// READ-ONLY; Pricing tier name of the provisioning service.
	Tier *string
}

// IotHubDefinitionDescription - Description of the IoT hub.
type IotHubDefinitionDescription struct {
	// REQUIRED; Connection string of the IoT hub.
	ConnectionString *string

	// REQUIRED; ARM region of the IoT hub.
	Location *string

	// weight to apply for a given iot h.
	AllocationWeight *int32

	// flag for applying allocationPolicy or not for a given iot hub.
	ApplyAllocationPolicy *bool

	// READ-ONLY; Host name of the IoT hub.
	Name *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// NameAvailabilityInfo - Description of name availability.
type NameAvailabilityInfo struct {
	// message containing a detailed reason name is unavailable
	Message *string

	// specifies if a name is available or not
	NameAvailable *bool

	// specifies the reason a name is unavailable
	Reason *NameUnavailabilityReason
}

// Operation - Provisioning Service REST API operation.
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// READ-ONLY; Operation name: {provider}/{resource}/{read | write | action | delete}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Name of the operation.
	Operation *string

	// READ-ONLY; Service provider: Microsoft Devices.
	Provider *string

	// READ-ONLY; Resource Type: ProvisioningServices.
	Resource *string
}

// OperationInputs - Input values for operation results call.
type OperationInputs struct {
	// REQUIRED; The name of the Provisioning Service to check.
	Name *string
}

// OperationListResult - Result of the request to list provisioning service operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string

	// READ-ONLY; Provisioning service operations supported by the Microsoft.Devices resource provider.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpoint - The private endpoint property of a private endpoint connection
type PrivateEndpoint struct {
	// READ-ONLY; The resource identifier.
	ID *string
}

// PrivateEndpointConnection - The private endpoint connection of a provisioning service
type PrivateEndpointConnection struct {
	// REQUIRED; The properties of a private endpoint connection
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// PrivateEndpointConnectionProperties - The properties of a private endpoint connection
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; The current state of a private endpoint connection
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The private endpoint property of a private endpoint connection
	PrivateEndpoint *PrivateEndpoint
}

// PrivateLinkResources - The available private link resources for a provisioning service
type PrivateLinkResources struct {
	// The list of available private link resources for a provisioning service
	Value []*GroupIDInformation
}

// PrivateLinkServiceConnectionState - The current state of a private endpoint connection
type PrivateLinkServiceConnectionState struct {
	// REQUIRED; The description for the current state of a private endpoint connection
	Description *string

	// REQUIRED; The status of a private endpoint connection
	Status *PrivateLinkServiceConnectionStatus

	// Actions required for a private endpoint connection
	ActionsRequired *string
}

// ProvisioningServiceDescription - The description of the provisioning service.
type ProvisioningServiceDescription struct {
	// REQUIRED; The resource location.
	Location *string

	// REQUIRED; Service specific properties for a provisioning service
	Properties *IotDpsPropertiesDescription

	// REQUIRED; Sku info for a provisioning Service.
	SKU *IotDpsSKUInfo

	// The Etag field is not required. If it is provided in the response body, it must also be provided as a header per the normal
	// ETag convention.
	Etag *string

	// The managed identities for a provisioning service.
	Identity *ManagedServiceIdentity

	// The resource group of the resource.
	Resourcegroup *string

	// The subscription id of the resource.
	Subscriptionid *string

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// ProvisioningServiceDescriptionListResult - List of provisioning service descriptions.
type ProvisioningServiceDescriptionListResult struct {
	// List of provisioning service descriptions.
	Value []*ProvisioningServiceDescription

	// READ-ONLY; the next link
	NextLink *string
}

// Resource - The common properties of an Azure resource.
type Resource struct {
	// REQUIRED; The resource location.
	Location *string

	// The resource group of the resource.
	Resourcegroup *string

	// The subscription id of the resource.
	Subscriptionid *string

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// SharedAccessSignatureAuthorizationRuleAccessRightsDescription - Description of the shared access key.
type SharedAccessSignatureAuthorizationRuleAccessRightsDescription struct {
	// REQUIRED; Name of the key.
	KeyName *string

	// REQUIRED; Rights that this key has.
	Rights *AccessRightsDescription

	// Primary SAS key value.
	PrimaryKey *string

	// Secondary SAS key value.
	SecondaryKey *string
}

// SharedAccessSignatureAuthorizationRuleListResult - List of shared access keys.
type SharedAccessSignatureAuthorizationRuleListResult struct {
	// The list of shared access policies.
	Value []*SharedAccessSignatureAuthorizationRuleAccessRightsDescription

	// READ-ONLY; The next link.
	NextLink *string
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

// TagsResource - A container holding only the Tags for a resource, allowing the user to update the tags on a Provisioning
// Service instance.
type TagsResource struct {
	// Resource tags
	Tags map[string]*string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VerificationCodeRequest - The JSON-serialized leaf certificate
type VerificationCodeRequest struct {
	// base-64 representation of X509 certificate .cer file or just .pem file content.
	Certificate *string
}

// VerificationCodeResponse - Description of the response of the verification code.
type VerificationCodeResponse struct {
	Properties *VerificationCodeResponseProperties

	// READ-ONLY; Request etag.
	Etag *string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; Name of certificate.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

type VerificationCodeResponseProperties struct {
	// base-64 representation of X509 certificate .cer file or just .pem file content.
	Certificate []byte

	// Certificate created time.
	Created *string

	// Code expiry.
	Expiry *string

	// Indicate if the certificate is verified by owner of private key.
	IsVerified *bool

	// Certificate subject.
	Subject *string

	// Certificate thumbprint.
	Thumbprint *string

	// Certificate updated time.
	Updated *string

	// Verification code.
	VerificationCode *string
}
