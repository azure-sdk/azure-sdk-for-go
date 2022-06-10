//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtimeseriesinsights

// AccessPoliciesClientCreateOrUpdateResponse contains the response from method AccessPoliciesClient.CreateOrUpdate.
type AccessPoliciesClientCreateOrUpdateResponse struct {
	AccessPolicyResource
}

// AccessPoliciesClientDeleteResponse contains the response from method AccessPoliciesClient.Delete.
type AccessPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// AccessPoliciesClientGetResponse contains the response from method AccessPoliciesClient.Get.
type AccessPoliciesClientGetResponse struct {
	AccessPolicyResource
}

// AccessPoliciesClientListByEnvironmentResponse contains the response from method AccessPoliciesClient.ListByEnvironment.
type AccessPoliciesClientListByEnvironmentResponse struct {
	AccessPolicyListResponse
}

// AccessPoliciesClientUpdateResponse contains the response from method AccessPoliciesClient.Update.
type AccessPoliciesClientUpdateResponse struct {
	AccessPolicyResource
}

// EnvironmentsClientCreateOrUpdateResponse contains the response from method EnvironmentsClient.CreateOrUpdate.
type EnvironmentsClientCreateOrUpdateResponse struct {
	EnvironmentResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnvironmentsClientCreateOrUpdateResponse.
func (e *EnvironmentsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEnvironmentResourceClassification(data)
	if err != nil {
		return err
	}
	e.EnvironmentResourceClassification = res
	return nil
}

// EnvironmentsClientDeleteResponse contains the response from method EnvironmentsClient.Delete.
type EnvironmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// EnvironmentsClientGetResponse contains the response from method EnvironmentsClient.Get.
type EnvironmentsClientGetResponse struct {
	EnvironmentResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnvironmentsClientGetResponse.
func (e *EnvironmentsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEnvironmentResourceClassification(data)
	if err != nil {
		return err
	}
	e.EnvironmentResourceClassification = res
	return nil
}

// EnvironmentsClientListByResourceGroupResponse contains the response from method EnvironmentsClient.ListByResourceGroup.
type EnvironmentsClientListByResourceGroupResponse struct {
	EnvironmentListResponse
}

// EnvironmentsClientListBySubscriptionResponse contains the response from method EnvironmentsClient.ListBySubscription.
type EnvironmentsClientListBySubscriptionResponse struct {
	EnvironmentListResponse
}

// EnvironmentsClientUpdateResponse contains the response from method EnvironmentsClient.Update.
type EnvironmentsClientUpdateResponse struct {
	EnvironmentResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EnvironmentsClientUpdateResponse.
func (e *EnvironmentsClientUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEnvironmentResourceClassification(data)
	if err != nil {
		return err
	}
	e.EnvironmentResourceClassification = res
	return nil
}

// EventSourcesClientCreateOrUpdateResponse contains the response from method EventSourcesClient.CreateOrUpdate.
type EventSourcesClientCreateOrUpdateResponse struct {
	EventSourceResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventSourcesClientCreateOrUpdateResponse.
func (e *EventSourcesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEventSourceResourceClassification(data)
	if err != nil {
		return err
	}
	e.EventSourceResourceClassification = res
	return nil
}

// EventSourcesClientDeleteResponse contains the response from method EventSourcesClient.Delete.
type EventSourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// EventSourcesClientGetResponse contains the response from method EventSourcesClient.Get.
type EventSourcesClientGetResponse struct {
	EventSourceResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventSourcesClientGetResponse.
func (e *EventSourcesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEventSourceResourceClassification(data)
	if err != nil {
		return err
	}
	e.EventSourceResourceClassification = res
	return nil
}

// EventSourcesClientListByEnvironmentResponse contains the response from method EventSourcesClient.ListByEnvironment.
type EventSourcesClientListByEnvironmentResponse struct {
	EventSourceListResponse
}

// EventSourcesClientUpdateResponse contains the response from method EventSourcesClient.Update.
type EventSourcesClientUpdateResponse struct {
	EventSourceResourceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventSourcesClientUpdateResponse.
func (e *EventSourcesClientUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEventSourceResourceClassification(data)
	if err != nil {
		return err
	}
	e.EventSourceResourceClassification = res
	return nil
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByEnvironmentResponse contains the response from method PrivateEndpointConnectionsClient.ListByEnvironment.
type PrivateEndpointConnectionsClientListByEnvironmentResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListSupportedResponse contains the response from method PrivateLinkResourcesClient.ListSupported.
type PrivateLinkResourcesClientListSupportedResponse struct {
	PrivateLinkResourceListResult
}

// ReferenceDataSetsClientCreateOrUpdateResponse contains the response from method ReferenceDataSetsClient.CreateOrUpdate.
type ReferenceDataSetsClientCreateOrUpdateResponse struct {
	ReferenceDataSetResource
}

// ReferenceDataSetsClientDeleteResponse contains the response from method ReferenceDataSetsClient.Delete.
type ReferenceDataSetsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReferenceDataSetsClientGetResponse contains the response from method ReferenceDataSetsClient.Get.
type ReferenceDataSetsClientGetResponse struct {
	ReferenceDataSetResource
}

// ReferenceDataSetsClientListByEnvironmentResponse contains the response from method ReferenceDataSetsClient.ListByEnvironment.
type ReferenceDataSetsClientListByEnvironmentResponse struct {
	ReferenceDataSetListResponse
}

// ReferenceDataSetsClientUpdateResponse contains the response from method ReferenceDataSetsClient.Update.
type ReferenceDataSetsClientUpdateResponse struct {
	ReferenceDataSetResource
}
