//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageactions

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// StorageTaskAssignmentClientListResponse contains the response from method StorageTaskAssignmentClient.NewListPager.
type StorageTaskAssignmentClientListResponse struct {
	// The response from the List Storage Tasks operation.
	PagedStorageTaskAssignment
}

// StorageTasksClientCreateResponse contains the response from method StorageTasksClient.BeginCreate.
type StorageTasksClientCreateResponse struct {
	// Represents Storage Task.
	StorageTask
}

// StorageTasksClientDeleteResponse contains the response from method StorageTasksClient.BeginDelete.
type StorageTasksClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageTasksClientGetResponse contains the response from method StorageTasksClient.Get.
type StorageTasksClientGetResponse struct {
	// Represents Storage Task.
	StorageTask
}

// StorageTasksClientListByResourceGroupResponse contains the response from method StorageTasksClient.NewListByResourceGroupPager.
type StorageTasksClientListByResourceGroupResponse struct {
	// The response of a StorageTask list operation.
	StorageTaskListResult
}

// StorageTasksClientListBySubscriptionResponse contains the response from method StorageTasksClient.NewListBySubscriptionPager.
type StorageTasksClientListBySubscriptionResponse struct {
	// The response of a StorageTask list operation.
	StorageTaskListResult
}

// StorageTasksClientPreviewActionsResponse contains the response from method StorageTasksClient.PreviewActions.
type StorageTasksClientPreviewActionsResponse struct {
	// Storage Task Preview Action.
	StorageTaskPreviewAction
}

// StorageTasksClientUpdateResponse contains the response from method StorageTasksClient.BeginUpdate.
type StorageTasksClientUpdateResponse struct {
	// Represents Storage Task.
	StorageTask
}

// StorageTasksReportClientListResponse contains the response from method StorageTasksReportClient.NewListPager.
type StorageTasksReportClientListResponse struct {
	// The response of a StorageTaskReportInstance list operation.
	StorageTaskReportInstanceListResult
}
