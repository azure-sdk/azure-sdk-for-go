//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

// CapabilitiesClientCreateOrUpdateResponse contains the response from method CapabilitiesClient.CreateOrUpdate.
type CapabilitiesClientCreateOrUpdateResponse struct {
	// Model that represents a Capability resource.
	Capability
}

// CapabilitiesClientDeleteResponse contains the response from method CapabilitiesClient.Delete.
type CapabilitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// CapabilitiesClientGetResponse contains the response from method CapabilitiesClient.Get.
type CapabilitiesClientGetResponse struct {
	// Model that represents a Capability resource.
	Capability
}

// CapabilitiesClientListResponse contains the response from method CapabilitiesClient.NewListPager.
type CapabilitiesClientListResponse struct {
	// Model that represents a list of Capability resources and a link for pagination.
	CapabilityListResult
}

// CapabilityTypesClientGetResponse contains the response from method CapabilityTypesClient.Get.
type CapabilityTypesClientGetResponse struct {
	// Model that represents a Capability Type resource.
	CapabilityType
}

// CapabilityTypesClientListResponse contains the response from method CapabilityTypesClient.NewListPager.
type CapabilityTypesClientListResponse struct {
	// Model that represents a list of Capability Type resources and a link for pagination.
	CapabilityTypeListResult
}

// ExperimentsClientCancelResponse contains the response from method ExperimentsClient.BeginCancel.
type ExperimentsClientCancelResponse struct {
	// placeholder for future response values
}

// ExperimentsClientCreateOrUpdateResponse contains the response from method ExperimentsClient.BeginCreateOrUpdate.
type ExperimentsClientCreateOrUpdateResponse struct {
	// Model that represents a Experiment resource.
	Experiment
}

// ExperimentsClientDeleteResponse contains the response from method ExperimentsClient.BeginDelete.
type ExperimentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExperimentsClientGetExecutionDetailsResponse contains the response from method ExperimentsClient.GetExecutionDetails.
type ExperimentsClientGetExecutionDetailsResponse struct {
	// Model that represents the execution details of a Experiment.
	ExperimentExecutionDetails
}

// ExperimentsClientGetResponse contains the response from method ExperimentsClient.Get.
type ExperimentsClientGetResponse struct {
	// Model that represents a Experiment resource.
	Experiment
}

// ExperimentsClientGetStatusResponse contains the response from method ExperimentsClient.GetStatus.
type ExperimentsClientGetStatusResponse struct {
	// Model that represents the status of a Experiment.
	ExperimentStatus
}

// ExperimentsClientListAllResponse contains the response from method ExperimentsClient.NewListAllPager.
type ExperimentsClientListAllResponse struct {
	// Model that represents a list of Experiment resources and a link for pagination.
	ExperimentListResult
}

// ExperimentsClientListAllStatusesResponse contains the response from method ExperimentsClient.NewListAllStatusesPager.
type ExperimentsClientListAllStatusesResponse struct {
	// Model that represents a list of Experiment statuses and a link for pagination.
	ExperimentStatusListResult
}

// ExperimentsClientListExecutionDetailsResponse contains the response from method ExperimentsClient.NewListExecutionDetailsPager.
type ExperimentsClientListExecutionDetailsResponse struct {
	// Model that represents a list of Experiment execution details and a link for pagination.
	ExperimentExecutionDetailsListResult
}

// ExperimentsClientListResponse contains the response from method ExperimentsClient.NewListPager.
type ExperimentsClientListResponse struct {
	// Model that represents a list of Experiment resources and a link for pagination.
	ExperimentListResult
}

// ExperimentsClientStartResponse contains the response from method ExperimentsClient.BeginStart.
type ExperimentsClientStartResponse struct {
	// placeholder for future response values
}

// ExperimentsClientUpdateResponse contains the response from method ExperimentsClient.BeginUpdate.
type ExperimentsClientUpdateResponse struct {
	// Model that represents a Experiment resource.
	Experiment
}

// OperationStatusesClientGetResponse contains the response from method OperationStatusesClient.Get.
type OperationStatusesClientGetResponse struct {
	// The status of operation.
	OperationStatus
}

// OperationsClientListAllResponse contains the response from method OperationsClient.NewListAllPager.
type OperationsClientListAllResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PrivateAccessesClientCreateOrUpdateResponse contains the response from method PrivateAccessesClient.CreateOrUpdate.
type PrivateAccessesClientCreateOrUpdateResponse struct {
	// PrivateAccesses tracked resource.
	PrivateAccess
}

// PrivateAccessesClientDeleteAPrivateEndpointConnectionResponse contains the response from method PrivateAccessesClient.BeginDeleteAPrivateEndpointConnection.
type PrivateAccessesClientDeleteAPrivateEndpointConnectionResponse struct {
	// placeholder for future response values
}

// PrivateAccessesClientDeleteResponse contains the response from method PrivateAccessesClient.Delete.
type PrivateAccessesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateAccessesClientGetAPrivateEndpointConnectionResponse contains the response from method PrivateAccessesClient.GetAPrivateEndpointConnection.
type PrivateAccessesClientGetAPrivateEndpointConnectionResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateAccessesClientGetPrivateLinkResourcesResponse contains the response from method PrivateAccessesClient.GetPrivateLinkResources.
type PrivateAccessesClientGetPrivateLinkResourcesResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// PrivateAccessesClientGetResponse contains the response from method PrivateAccessesClient.Get.
type PrivateAccessesClientGetResponse struct {
	// PrivateAccesses tracked resource.
	PrivateAccess
}

// PrivateAccessesClientListAllResponse contains the response from method PrivateAccessesClient.NewListAllPager.
type PrivateAccessesClientListAllResponse struct {
	// Model that represents a list of private access resources and a link for pagination.
	PrivateAccessListResult
}

// PrivateAccessesClientListPrivateEndpointConnectionsResponse contains the response from method PrivateAccessesClient.NewListPrivateEndpointConnectionsPager.
type PrivateAccessesClientListPrivateEndpointConnectionsResponse struct {
	// A list of private link resources
	PrivateEndpointConnectionListResult
}

// PrivateAccessesClientListResponse contains the response from method PrivateAccessesClient.NewListPager.
type PrivateAccessesClientListResponse struct {
	// Model that represents a list of private access resources and a link for pagination.
	PrivateAccessListResult
}

// PrivateAccessesClientUpdateAPrivateEndpointConnectionResponse contains the response from method PrivateAccessesClient.BeginUpdateAPrivateEndpointConnection.
type PrivateAccessesClientUpdateAPrivateEndpointConnectionResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateAccessesClientUpdateResponse contains the response from method PrivateAccessesClient.Update.
type PrivateAccessesClientUpdateResponse struct {
	// PrivateAccesses tracked resource.
	PrivateAccess
}

// TargetTypesClientGetResponse contains the response from method TargetTypesClient.Get.
type TargetTypesClientGetResponse struct {
	// Model that represents a Target Type resource.
	TargetType
}

// TargetTypesClientListResponse contains the response from method TargetTypesClient.NewListPager.
type TargetTypesClientListResponse struct {
	// Model that represents a list of Target Type resources and a link for pagination.
	TargetTypeListResult
}

// TargetsClientCreateOrUpdateResponse contains the response from method TargetsClient.CreateOrUpdate.
type TargetsClientCreateOrUpdateResponse struct {
	// Model that represents a Target resource.
	Target
}

// TargetsClientDeleteResponse contains the response from method TargetsClient.Delete.
type TargetsClientDeleteResponse struct {
	// placeholder for future response values
}

// TargetsClientGetResponse contains the response from method TargetsClient.Get.
type TargetsClientGetResponse struct {
	// Model that represents a Target resource.
	Target
}

// TargetsClientListResponse contains the response from method TargetsClient.NewListPager.
type TargetsClientListResponse struct {
	// Model that represents a list of Target resources and a link for pagination.
	TargetListResult
}
