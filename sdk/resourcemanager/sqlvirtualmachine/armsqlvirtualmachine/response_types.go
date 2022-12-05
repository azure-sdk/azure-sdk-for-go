//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine

// AvailabilityGroupListenersClientCreateOrUpdateResponse contains the response from method AvailabilityGroupListenersClient.CreateOrUpdate.
type AvailabilityGroupListenersClientCreateOrUpdateResponse struct {
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientDeleteResponse contains the response from method AvailabilityGroupListenersClient.Delete.
type AvailabilityGroupListenersClientDeleteResponse struct {
	// placeholder for future response values
}

// AvailabilityGroupListenersClientGetResponse contains the response from method AvailabilityGroupListenersClient.Get.
type AvailabilityGroupListenersClientGetResponse struct {
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientListByGroupResponse contains the response from method AvailabilityGroupListenersClient.ListByGroup.
type AvailabilityGroupListenersClientListByGroupResponse struct {
	AvailabilityGroupListenerListResult
}

// GroupsClientCreateOrUpdateResponse contains the response from method GroupsClient.CreateOrUpdate.
type GroupsClientCreateOrUpdateResponse struct {
	Group
}

// GroupsClientDeleteResponse contains the response from method GroupsClient.Delete.
type GroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupsClientGetResponse contains the response from method GroupsClient.Get.
type GroupsClientGetResponse struct {
	Group
}

// GroupsClientListByResourceGroupResponse contains the response from method GroupsClient.ListByResourceGroup.
type GroupsClientListByResourceGroupResponse struct {
	GroupListResult
}

// GroupsClientListResponse contains the response from method GroupsClient.List.
type GroupsClientListResponse struct {
	GroupListResult
}

// GroupsClientUpdateResponse contains the response from method GroupsClient.Update.
type GroupsClientUpdateResponse struct {
	Group
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// SQLVirtualMachinesClientCreateOrUpdateResponse contains the response from method SQLVirtualMachinesClient.CreateOrUpdate.
type SQLVirtualMachinesClientCreateOrUpdateResponse struct {
	SQLVirtualMachine
}

// SQLVirtualMachinesClientDeleteResponse contains the response from method SQLVirtualMachinesClient.Delete.
type SQLVirtualMachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientGetResponse contains the response from method SQLVirtualMachinesClient.Get.
type SQLVirtualMachinesClientGetResponse struct {
	SQLVirtualMachine
}

// SQLVirtualMachinesClientListByResourceGroupResponse contains the response from method SQLVirtualMachinesClient.ListByResourceGroup.
type SQLVirtualMachinesClientListByResourceGroupResponse struct {
	ListResult
}

// SQLVirtualMachinesClientListBySQLVMGroupResponse contains the response from method SQLVirtualMachinesClient.ListBySQLVMGroup.
type SQLVirtualMachinesClientListBySQLVMGroupResponse struct {
	ListResult
}

// SQLVirtualMachinesClientListResponse contains the response from method SQLVirtualMachinesClient.List.
type SQLVirtualMachinesClientListResponse struct {
	ListResult
}

// SQLVirtualMachinesClientRedeployResponse contains the response from method SQLVirtualMachinesClient.Redeploy.
type SQLVirtualMachinesClientRedeployResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientStartAssessmentResponse contains the response from method SQLVirtualMachinesClient.StartAssessment.
type SQLVirtualMachinesClientStartAssessmentResponse struct {
	// placeholder for future response values
}

// SQLVirtualMachinesClientUpdateResponse contains the response from method SQLVirtualMachinesClient.Update.
type SQLVirtualMachinesClientUpdateResponse struct {
	SQLVirtualMachine
}

// TroubleshootClientTroubleshootResponse contains the response from method TroubleshootClient.Troubleshoot.
type TroubleshootClientTroubleshootResponse struct {
	SQLVMTroubleshooting
}
