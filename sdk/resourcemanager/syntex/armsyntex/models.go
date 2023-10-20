//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsyntex

import "time"

// DocumentProcessorList - Document processor list
type DocumentProcessorList struct {
	// READ-ONLY; Link to retrieve next page of results.
	NextLink *string

	// READ-ONLY; Array of document processors.
	Value []*DocumentProcessorResource
}

// DocumentProcessorPatchableProperties - Document processor properties that can be updated
type DocumentProcessorPatchableProperties struct {
	// Syntex properties for PATCH.
	Properties *DocumentProcessorPatchablePropertiesProperties

	// Resource tags.
	Tags map[string]*string
}

// DocumentProcessorPatchablePropertiesProperties - Syntex properties for PATCH.
type DocumentProcessorPatchablePropertiesProperties struct {
	// The ID (GUID) of an SharePoint Online (SPO) tenant associated with this document processor resource
	SpoTenantID *string

	// The URL of an SharePoint Online (SPO) tenant associated with this document processor resource
	SpoTenantURL *string
}

// DocumentProcessorProperties - Document processor properties
type DocumentProcessorProperties struct {
	// REQUIRED; The ID (GUID) of an SharePoint Online (SPO) tenant associated with this document processor resource
	SpoTenantID *string

	// REQUIRED; The URL of an SharePoint Online (SPO) tenant associated with this document processor resource
	SpoTenantURL *string

	// READ-ONLY; The managed resource provisioning state.
	ProvisioningState *ProvisioningState
}

// DocumentProcessorResource - Document processor details
type DocumentProcessorResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Document processor properties.
	Properties *DocumentProcessorProperties

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
