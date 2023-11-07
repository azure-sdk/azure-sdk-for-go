//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

// DicomServicesClientCreateOrUpdateResponse contains the response from method DicomServicesClient.BeginCreateOrUpdate.
type DicomServicesClientCreateOrUpdateResponse struct {
	// The description of Dicom Service
	DicomService
}

// DicomServicesClientDeleteResponse contains the response from method DicomServicesClient.BeginDelete.
type DicomServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DicomServicesClientGetResponse contains the response from method DicomServicesClient.Get.
type DicomServicesClientGetResponse struct {
	// The description of Dicom Service
	DicomService
}

// DicomServicesClientListByWorkspaceResponse contains the response from method DicomServicesClient.NewListByWorkspacePager.
type DicomServicesClientListByWorkspaceResponse struct {
	// The collection of Dicom Services.
	DicomServiceCollection
}

// DicomServicesClientUpdateResponse contains the response from method DicomServicesClient.BeginUpdate.
type DicomServicesClientUpdateResponse struct {
	// The description of Dicom Service
	DicomService
}

// FhirDestinationsClientListByIotConnectorResponse contains the response from method FhirDestinationsClient.NewListByIotConnectorPager.
type FhirDestinationsClientListByIotConnectorResponse struct {
	// A collection of IoT Connector FHIR destinations.
	IotFhirDestinationCollection
}

// FhirServicesClientCreateOrUpdateResponse contains the response from method FhirServicesClient.BeginCreateOrUpdate.
type FhirServicesClientCreateOrUpdateResponse struct {
	// The description of Fhir Service
	FhirService
}

// FhirServicesClientDeleteResponse contains the response from method FhirServicesClient.BeginDelete.
type FhirServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// FhirServicesClientGetResponse contains the response from method FhirServicesClient.Get.
type FhirServicesClientGetResponse struct {
	// The description of Fhir Service
	FhirService
}

// FhirServicesClientListByWorkspaceResponse contains the response from method FhirServicesClient.NewListByWorkspacePager.
type FhirServicesClientListByWorkspaceResponse struct {
	// A collection of Fhir services.
	FhirServiceCollection
}

// FhirServicesClientUpdateResponse contains the response from method FhirServicesClient.BeginUpdate.
type FhirServicesClientUpdateResponse struct {
	// The description of Fhir Service
	FhirService
}

// IotConnectorFhirDestinationClientCreateOrUpdateResponse contains the response from method IotConnectorFhirDestinationClient.BeginCreateOrUpdate.
type IotConnectorFhirDestinationClientCreateOrUpdateResponse struct {
	// IoT Connector FHIR destination definition.
	IotFhirDestination
}

// IotConnectorFhirDestinationClientDeleteResponse contains the response from method IotConnectorFhirDestinationClient.BeginDelete.
type IotConnectorFhirDestinationClientDeleteResponse struct {
	// placeholder for future response values
}

// IotConnectorFhirDestinationClientGetResponse contains the response from method IotConnectorFhirDestinationClient.Get.
type IotConnectorFhirDestinationClientGetResponse struct {
	// IoT Connector FHIR destination definition.
	IotFhirDestination
}

// IotConnectorsClientCreateOrUpdateResponse contains the response from method IotConnectorsClient.BeginCreateOrUpdate.
type IotConnectorsClientCreateOrUpdateResponse struct {
	// IoT Connector definition.
	IotConnector
}

// IotConnectorsClientDeleteResponse contains the response from method IotConnectorsClient.BeginDelete.
type IotConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// IotConnectorsClientGetResponse contains the response from method IotConnectorsClient.Get.
type IotConnectorsClientGetResponse struct {
	// IoT Connector definition.
	IotConnector
}

// IotConnectorsClientListByWorkspaceResponse contains the response from method IotConnectorsClient.NewListByWorkspacePager.
type IotConnectorsClientListByWorkspaceResponse struct {
	// A collection of IoT Connectors.
	IotConnectorCollection
}

// IotConnectorsClientUpdateResponse contains the response from method IotConnectorsClient.BeginUpdate.
type IotConnectorsClientUpdateResponse struct {
	// IoT Connector definition.
	IotConnector
}

// OperationResultsClientGetResponse contains the response from method OperationResultsClient.Get.
type OperationResultsClientGetResponse struct {
	// The properties indicating the operation result of an operation on a service.
	OperationResultsDescription
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Available operations of the service
	ListOperations
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionDescription
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionDescription
}

// PrivateEndpointConnectionsClientListByServiceResponse contains the response from method PrivateEndpointConnectionsClient.NewListByServicePager.
type PrivateEndpointConnectionsClientListByServiceResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResultDescription
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateLinkResourceDescription
}

// PrivateLinkResourcesClientListByServiceResponse contains the response from method PrivateLinkResourcesClient.ListByService.
type PrivateLinkResourcesClientListByServiceResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResultDescription
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// The properties indicating whether a given service name is available.
	ServicesNameAvailabilityInfo
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// The description of the service.
	ServicesDescription
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// The description of the service.
	ServicesDescription
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// A list of service description objects with a next link.
	ServicesDescriptionListResult
}

// ServicesClientListResponse contains the response from method ServicesClient.NewListPager.
type ServicesClientListResponse struct {
	// A list of service description objects with a next link.
	ServicesDescriptionListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.BeginUpdate.
type ServicesClientUpdateResponse struct {
	// The description of the service.
	ServicesDescription
}

// WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type WorkspacePrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionDescription
}

// WorkspacePrivateEndpointConnectionsClientDeleteResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.BeginDelete.
type WorkspacePrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacePrivateEndpointConnectionsClientGetResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.Get.
type WorkspacePrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnectionDescription
}

// WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse contains the response from method WorkspacePrivateEndpointConnectionsClient.NewListByWorkspacePager.
type WorkspacePrivateEndpointConnectionsClientListByWorkspaceResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResultDescription
}

// WorkspacePrivateLinkResourcesClientGetResponse contains the response from method WorkspacePrivateLinkResourcesClient.Get.
type WorkspacePrivateLinkResourcesClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateLinkResourceDescription
}

// WorkspacePrivateLinkResourcesClientListByWorkspaceResponse contains the response from method WorkspacePrivateLinkResourcesClient.NewListByWorkspacePager.
type WorkspacePrivateLinkResourcesClientListByWorkspaceResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResultDescription
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.BeginCreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	// Workspace resource.
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.BeginDelete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	// Workspace resource.
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.NewListByResourceGroupPager.
type WorkspacesClientListByResourceGroupResponse struct {
	// Collection of workspace object with a next link
	WorkspaceList
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.NewListBySubscriptionPager.
type WorkspacesClientListBySubscriptionResponse struct {
	// Collection of workspace object with a next link
	WorkspaceList
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.BeginUpdate.
type WorkspacesClientUpdateResponse struct {
	// Workspace resource.
	Workspace
}
