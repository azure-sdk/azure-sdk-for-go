//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

// AvailabilityGroupListenersClientCreateOrUpdateResponse contains the response from method AvailabilityGroupListenersClient.BeginCreateOrUpdate.
type AvailabilityGroupListenersClientCreateOrUpdateResponse struct {
	// A SQL Server availability group listener.
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientDeleteResponse contains the response from method AvailabilityGroupListenersClient.BeginDelete.
type AvailabilityGroupListenersClientDeleteResponse struct {
	// placeholder for future response values
}

// AvailabilityGroupListenersClientGetResponse contains the response from method AvailabilityGroupListenersClient.Get.
type AvailabilityGroupListenersClientGetResponse struct {
	// A SQL Server availability group listener.
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientListByGroupResponse contains the response from method AvailabilityGroupListenersClient.NewListByGroupPager.
type AvailabilityGroupListenersClientListByGroupResponse struct {
	// A list of availability group listeners.
	AvailabilityGroupListenerListResult
}

// GroupsClientCreateOrUpdateResponse contains the response from method GroupsClient.BeginCreateOrUpdate.
type GroupsClientCreateOrUpdateResponse struct {
	// A SQL virtual machine group.
	Group
}

// GroupsClientDeleteResponse contains the response from method GroupsClient.BeginDelete.
type GroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupsClientGetResponse contains the response from method GroupsClient.Get.
type GroupsClientGetResponse struct {
	// A SQL virtual machine group.
	Group
}

// GroupsClientListByResourceGroupResponse contains the response from method GroupsClient.NewListByResourceGroupPager.
type GroupsClientListByResourceGroupResponse struct {
	// A list of SQL virtual machine groups.
	GroupListResult
}

// GroupsClientListResponse contains the response from method GroupsClient.NewListPager.
type GroupsClientListResponse struct {
	// A list of SQL virtual machine groups.
	GroupListResult
}

// GroupsClientUpdateResponse contains the response from method GroupsClient.BeginUpdate.
type GroupsClientUpdateResponse struct {
	// A SQL virtual machine group.
	Group
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list SQL operations.
	OperationListResult
}

// SQLVirtualMachinesClientCreateOrUpdateResponse contains the response from method SQLVirtualMachinesClient.BeginCreateOrUpdate.
type SQLVirtualMachinesClientCreateOrUpdateResponse struct {
	// A SQL virtual machine.
	SQLVirtualMachine
}

// SQLVirtualMachinesClientDeleteResponse contains the response from method SQLVirtualMachinesClient.BeginDelete.
type SQLVirtualMachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientGetResponse contains the response from method SQLVirtualMachinesClient.Get.
type SQLVirtualMachinesClientGetResponse struct {
	// A SQL virtual machine.
	SQLVirtualMachine
}

// SQLVirtualMachinesClientListByResourceGroupResponse contains the response from method SQLVirtualMachinesClient.NewListByResourceGroupPager.
type SQLVirtualMachinesClientListByResourceGroupResponse struct {
	// A list of SQL virtual machines.
	ListResult
}

// SQLVirtualMachinesClientListBySQLVMGroupResponse contains the response from method SQLVirtualMachinesClient.NewListBySQLVMGroupPager.
type SQLVirtualMachinesClientListBySQLVMGroupResponse struct {
	// A list of SQL virtual machines.
	ListResult
}

// SQLVirtualMachinesClientListResponse contains the response from method SQLVirtualMachinesClient.NewListPager.
type SQLVirtualMachinesClientListResponse struct {
	// A list of SQL virtual machines.
	ListResult
}

// SQLVirtualMachinesClientRedeployResponse contains the response from method SQLVirtualMachinesClient.BeginRedeploy.
type SQLVirtualMachinesClientRedeployResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientStartAssessmentResponse contains the response from method SQLVirtualMachinesClient.BeginStartAssessment.
type SQLVirtualMachinesClientStartAssessmentResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientUpdateResponse contains the response from method SQLVirtualMachinesClient.BeginUpdate.
type SQLVirtualMachinesClientUpdateResponse struct {
	// A SQL virtual machine.
	SQLVirtualMachine
}

// TroubleshootClientTroubleshootResponse contains the response from method TroubleshootClient.BeginTroubleshoot.
type TroubleshootClientTroubleshootResponse struct {
	// Details required for SQL VM troubleshooting
	SQLVMTroubleshooting
}
