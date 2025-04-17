// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaddons

// CanonicalSupportPlanInfoDefinition - Definition object with the properties of a canonical plan
type CanonicalSupportPlanInfoDefinition struct {
	// Flag to indicate if this support plan type is currently enabled for the subscription.
	Enabled *bool

	// The one time charge status for the subscription.
	OneTimeCharge *OneTimeCharge

	// Support plan type.
	SupportPlanType *SupportPlanType
}

// CanonicalSupportPlanProperties - The properties of the Canonical support plan.
type CanonicalSupportPlanProperties struct {
	// The provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// CanonicalSupportPlanResponseEnvelope - The status of the Canonical support plan.
type CanonicalSupportPlanResponseEnvelope struct {
	// REQUIRED; Describes Canonical support plan type and status.
	Properties *CanonicalSupportPlanProperties

	// READ-ONLY; The id of the ARM resource, e.g. "/subscriptions/{id}/providers/Microsoft.Addons/supportProvider/{supportProviderName}/supportPlanTypes/{planTypeName}".
	ID *string

	// READ-ONLY; The name of the Canonical support plan, i.e. "essential", "standard" or "advanced".
	Name *string

	// READ-ONLY; Microsoft.Addons/supportProvider
	Type *string
}

// ErrorDefinition - Error description and code explaining why an operation failed.
type ErrorDefinition struct {
	// REQUIRED; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// REQUIRED; Description of the error.
	Message *string
}

// OperationListValue - List of supported operations.
type OperationListValue struct {
	// List of supported operations.
	Value []*OperationsDefinition
}

// OperationsDefinition - Definition object with the name and properties of an operation.
type OperationsDefinition struct {
	// Display object with properties of the operation.
	Display *OperationsDisplayDefinition

	// Name of the operation.
	Name *string
}

// OperationsDisplayDefinition - Display object with properties of the operation.
type OperationsDisplayDefinition struct {
	// Description of the operation.
	Description *string

	// Short description of the operation.
	Operation *string

	// Resource provider of the operation.
	Provider *string

	// Resource for the operation.
	Resource *string
}
