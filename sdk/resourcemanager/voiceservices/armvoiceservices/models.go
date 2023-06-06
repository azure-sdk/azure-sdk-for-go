//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvoiceservices

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is not available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// CommunicationsGateway - A CommunicationsGateway resource
type CommunicationsGateway struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *CommunicationsGatewayProperties

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

// CommunicationsGatewayListResult - The response of a CommunicationsGateway list operation.
type CommunicationsGatewayListResult struct {
	// REQUIRED; The CommunicationsGateway items on this page
	Value []*CommunicationsGateway

	// The link to the next page of items
	NextLink *string
}

// CommunicationsGatewayProperties - Details of the CommunicationsGateway resource.
type CommunicationsGatewayProperties struct {
	// REQUIRED; Voice codecs to support
	Codecs []*TeamsCodecs

	// REQUIRED; How to connect back to the operator network, e.g. MAPS
	Connectivity *Connectivity

	// REQUIRED; How to handle 911 calls
	E911Type *E911Type

	// REQUIRED; What platforms to support
	Platforms []*CommunicationsPlatform

	// REQUIRED; The regions in which to deploy the resources needed for Teams Calling
	ServiceLocations []*ServiceRegionProperties

	// Details of API bridge functionality, if required
	APIBridge any

	// The scope at which the auto-generated domain name can be re-used
	AutoGeneratedDomainNameLabelScope *AutoGeneratedDomainNameLabelScope

	// A list of dial strings used for emergency calling.
	EmergencyDialStrings []*string

	// Whether an integrated Mobile Control Point is in use.
	IntegratedMcpEnabled *bool

	// Whether an on-premises Mobile Control Point is in use.
	OnPremMcpEnabled *bool

	// This number is used in Teams Phone Mobile scenarios for access to the voicemail IVR from the native dialer.
	TeamsVoicemailPilotNumber *string

	// READ-ONLY; The autogenerated label used as part of the FQDNs for accessing the Communications Gateway
	AutoGeneratedDomainNameLabel *string

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The current status of the deployment.
	Status *Status
}

// CommunicationsGatewayUpdate - The type used for update operations of the CommunicationsGateway.
type CommunicationsGatewayUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// Resource tags.
	Tags map[string]*string
}

// CommunicationsGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the CommunicationsGatewaysClient.BeginCreateOrUpdate
// method.
type CommunicationsGatewaysClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommunicationsGatewaysClientBeginDeleteOptions contains the optional parameters for the CommunicationsGatewaysClient.BeginDelete
// method.
type CommunicationsGatewaysClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CommunicationsGatewaysClientGetOptions contains the optional parameters for the CommunicationsGatewaysClient.Get method.
type CommunicationsGatewaysClientGetOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsGatewaysClientListByResourceGroupOptions contains the optional parameters for the CommunicationsGatewaysClient.NewListByResourceGroupPager
// method.
type CommunicationsGatewaysClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsGatewaysClientListBySubscriptionOptions contains the optional parameters for the CommunicationsGatewaysClient.NewListBySubscriptionPager
// method.
type CommunicationsGatewaysClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// CommunicationsGatewaysClientUpdateOptions contains the optional parameters for the CommunicationsGatewaysClient.Update
// method.
type CommunicationsGatewaysClientUpdateOptions struct {
	// placeholder for future optional parameters
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

// NameAvailabilityClientCheckLocalOptions contains the optional parameters for the NameAvailabilityClient.CheckLocal method.
type NameAvailabilityClientCheckLocalOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrimaryRegionProperties - The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionProperties struct {
	// REQUIRED; IP address to use to contact the operator network from this region
	OperatorAddresses []*string

	// The allowed source IP address or CIDR ranges for media
	AllowedMediaSourceAddressPrefixes []*string

	// The allowed source IP address or CIDR ranges for signaling
	AllowedSignalingSourceAddressPrefixes []*string

	// IP address to use to contact the ESRP from this region
	EsrpAddresses []*string
}

// ServiceRegionProperties - The service region configuration needed for Teams Callings.
type ServiceRegionProperties struct {
	// REQUIRED; The name of the region in which the resources needed for Teams Calling will be deployed.
	Name *string

	// REQUIRED; The configuration used in this region as primary, and other regions as backup.
	PrimaryRegionProperties *PrimaryRegionProperties
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

// TestLine - A TestLine resource
type TestLine struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *TestLineProperties

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

// TestLineListResult - The response of a TestLine list operation.
type TestLineListResult struct {
	// REQUIRED; The TestLine items on this page
	Value []*TestLine

	// The link to the next page of items
	NextLink *string
}

// TestLineProperties - Details of the TestLine resource.
type TestLineProperties struct {
	// REQUIRED; The phone number
	PhoneNumber *string

	// REQUIRED; Purpose of this test line, e.g. automated or manual testing
	Purpose *TestLinePurpose

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ProvisioningState
}

// TestLineUpdate - The type used for update operations of the TestLine.
type TestLineUpdate struct {
	// Resource tags.
	Tags map[string]*string
}

// TestLinesClientBeginCreateOrUpdateOptions contains the optional parameters for the TestLinesClient.BeginCreateOrUpdate
// method.
type TestLinesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TestLinesClientBeginDeleteOptions contains the optional parameters for the TestLinesClient.BeginDelete method.
type TestLinesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TestLinesClientGetOptions contains the optional parameters for the TestLinesClient.Get method.
type TestLinesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TestLinesClientListByCommunicationsGatewayOptions contains the optional parameters for the TestLinesClient.NewListByCommunicationsGatewayPager
// method.
type TestLinesClientListByCommunicationsGatewayOptions struct {
	// placeholder for future optional parameters
}

// TestLinesClientUpdateOptions contains the optional parameters for the TestLinesClient.Update method.
type TestLinesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
