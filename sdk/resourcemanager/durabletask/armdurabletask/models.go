// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdurabletask

import "time"

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
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
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// Scheduler - A Durable Task Scheduler resource
type Scheduler struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *SchedulerProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The name of the Scheduler
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SchedulerListResult - The response of a Scheduler list operation.
type SchedulerListResult struct {
	// REQUIRED; The Scheduler items on this page
	Value []*Scheduler

	// The link to the next page of items
	NextLink *string
}

// SchedulerProperties - Details of the Scheduler
type SchedulerProperties struct {
	// REQUIRED; IP allow list for durable task scheduler. Values can be IPv4, IPv6 or CIDR
	IPAllowlist []*string

	// REQUIRED; SKU of the durable task scheduler
	SKU *SchedulerSKU

	// READ-ONLY; URL of the durable task scheduler
	Endpoint *string

	// READ-ONLY; The status of the last operation
	ProvisioningState *ProvisioningState
}

// SchedulerSKU - The SKU (Stock Keeping Unit) assigned to this durable task scheduler
type SchedulerSKU struct {
	// REQUIRED; The name of the SKU
	Name *string

	// The SKU capacity. This allows scale out/in for the resource and impacts zone redundancy
	Capacity *int32

	// READ-ONLY; Indicates whether the current SKU configuration is zone redundant
	RedundancyState *RedundancyState
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

// TaskHub - A Task Hub resource belonging to the scheduler
type TaskHub struct {
	// The resource-specific properties for this resource.
	Properties *TaskHubProperties

	// READ-ONLY; The name of the TaskHub
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TaskHubListResult - The response of a TaskHub list operation.
type TaskHubListResult struct {
	// REQUIRED; The TaskHub items on this page
	Value []*TaskHub

	// The link to the next page of items
	NextLink *string
}

// TaskHubProperties - The properties of Task Hub
type TaskHubProperties struct {
	// READ-ONLY; URL of the durable task scheduler dashboard
	DashboardURL *string

	// READ-ONLY; The status of the last operation
	ProvisioningState *ProvisioningState
}