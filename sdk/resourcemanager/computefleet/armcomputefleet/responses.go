// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcomputefleet

// FleetsClientCreateOrUpdateResponse contains the response from method FleetsClient.BeginCreateOrUpdate.
type FleetsClientCreateOrUpdateResponse struct {
	// An Compute Fleet resource
	Fleet
}

// FleetsClientDeleteResponse contains the response from method FleetsClient.BeginDelete.
type FleetsClientDeleteResponse struct {
	// placeholder for future response values
}

// FleetsClientGetResponse contains the response from method FleetsClient.Get.
type FleetsClientGetResponse struct {
	// An Compute Fleet resource
	Fleet
}

// FleetsClientListByResourceGroupResponse contains the response from method FleetsClient.NewListByResourceGroupPager.
type FleetsClientListByResourceGroupResponse struct {
	// The response of a Fleet list operation.
	FleetListResult
}

// FleetsClientListBySubscriptionResponse contains the response from method FleetsClient.NewListBySubscriptionPager.
type FleetsClientListBySubscriptionResponse struct {
	// The response of a Fleet list operation.
	FleetListResult
}

// FleetsClientListVirtualMachineScaleSetsResponse contains the response from method FleetsClient.NewListVirtualMachineScaleSetsPager.
type FleetsClientListVirtualMachineScaleSetsResponse struct {
	// The response of a VirtualMachineScaleSet list operation.
	VirtualMachineScaleSetListResult
}

// FleetsClientUpdateResponse contains the response from method FleetsClient.BeginUpdate.
type FleetsClientUpdateResponse struct {
	// An Compute Fleet resource
	Fleet
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
