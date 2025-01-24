//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

// AccountFiltersClientCreateOrUpdateOptions contains the optional parameters for the AccountFiltersClient.CreateOrUpdate
// method.
type AccountFiltersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AccountFiltersClientDeleteOptions contains the optional parameters for the AccountFiltersClient.Delete method.
type AccountFiltersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AccountFiltersClientGetOptions contains the optional parameters for the AccountFiltersClient.Get method.
type AccountFiltersClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountFiltersClientListOptions contains the optional parameters for the AccountFiltersClient.NewListPager method.
type AccountFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// AccountFiltersClientUpdateOptions contains the optional parameters for the AccountFiltersClient.Update method.
type AccountFiltersClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AssetFiltersClientCreateOrUpdateOptions contains the optional parameters for the AssetFiltersClient.CreateOrUpdate method.
type AssetFiltersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AssetFiltersClientDeleteOptions contains the optional parameters for the AssetFiltersClient.Delete method.
type AssetFiltersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AssetFiltersClientGetOptions contains the optional parameters for the AssetFiltersClient.Get method.
type AssetFiltersClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssetFiltersClientListOptions contains the optional parameters for the AssetFiltersClient.NewListPager method.
type AssetFiltersClientListOptions struct {
	// placeholder for future optional parameters
}

// AssetFiltersClientUpdateOptions contains the optional parameters for the AssetFiltersClient.Update method.
type AssetFiltersClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// AssetTrackOperationResultsClientGetOptions contains the optional parameters for the AssetTrackOperationResultsClient.Get
// method.
type AssetTrackOperationResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssetTrackOperationStatusesClientGetOptions contains the optional parameters for the AssetTrackOperationStatusesClient.Get
// method.
type AssetTrackOperationStatusesClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientCreateOrUpdateOptions contains the optional parameters for the AssetsClient.CreateOrUpdate method.
type AssetsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientDeleteOptions contains the optional parameters for the AssetsClient.Delete method.
type AssetsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientGetEncryptionKeyOptions contains the optional parameters for the AssetsClient.GetEncryptionKey method.
type AssetsClientGetEncryptionKeyOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientGetOptions contains the optional parameters for the AssetsClient.Get method.
type AssetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientListContainerSasOptions contains the optional parameters for the AssetsClient.ListContainerSas method.
type AssetsClientListContainerSasOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientListOptions contains the optional parameters for the AssetsClient.NewListPager method.
type AssetsClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string

	// Specifies a non-negative integer n that limits the number of items returned from a collection. The service returns the
	// number of available items up to but not greater than the specified value n.
	Top *int32
}

// AssetsClientListStreamingLocatorsOptions contains the optional parameters for the AssetsClient.ListStreamingLocators method.
type AssetsClientListStreamingLocatorsOptions struct {
	// placeholder for future optional parameters
}

// AssetsClientUpdateOptions contains the optional parameters for the AssetsClient.Update method.
type AssetsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
type ClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
type ClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ClientDeleteOptions contains the optional parameters for the Client.Delete method.
type ClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListBySubscriptionOptions contains the optional parameters for the Client.NewListBySubscriptionPager method.
type ClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ClientListEdgePoliciesOptions contains the optional parameters for the Client.ListEdgePolicies method.
type ClientListEdgePoliciesOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.NewListPager method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// ClientSyncStorageKeysOptions contains the optional parameters for the Client.SyncStorageKeys method.
type ClientSyncStorageKeysOptions struct {
	// placeholder for future optional parameters
}

// ContentKeyPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ContentKeyPoliciesClient.CreateOrUpdate
// method.
type ContentKeyPoliciesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ContentKeyPoliciesClientDeleteOptions contains the optional parameters for the ContentKeyPoliciesClient.Delete method.
type ContentKeyPoliciesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ContentKeyPoliciesClientGetOptions contains the optional parameters for the ContentKeyPoliciesClient.Get method.
type ContentKeyPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsOptions contains the optional parameters for the ContentKeyPoliciesClient.GetPolicyPropertiesWithSecrets
// method.
type ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsOptions struct {
	// placeholder for future optional parameters
}

// ContentKeyPoliciesClientListOptions contains the optional parameters for the ContentKeyPoliciesClient.NewListPager method.
type ContentKeyPoliciesClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string

	// Specifies a non-negative integer n that limits the number of items returned from a collection. The service returns the
	// number of available items up to but not greater than the specified value n.
	Top *int32
}

// ContentKeyPoliciesClientUpdateOptions contains the optional parameters for the ContentKeyPoliciesClient.Update method.
type ContentKeyPoliciesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// JobsClientCancelJobOptions contains the optional parameters for the JobsClient.CancelJob method.
type JobsClientCancelJobOptions struct {
	// placeholder for future optional parameters
}

// JobsClientCreateOptions contains the optional parameters for the JobsClient.Create method.
type JobsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// JobsClientDeleteOptions contains the optional parameters for the JobsClient.Delete method.
type JobsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
type JobsClientGetOptions struct {
	// placeholder for future optional parameters
}

// JobsClientListOptions contains the optional parameters for the JobsClient.NewListPager method.
type JobsClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string
}

// JobsClientUpdateOptions contains the optional parameters for the JobsClient.Update method.
type JobsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// LiveEventsClientAsyncOperationOptions contains the optional parameters for the LiveEventsClient.AsyncOperation method.
type LiveEventsClientAsyncOperationOptions struct {
	// placeholder for future optional parameters
}

// LiveEventsClientBeginAllocateOptions contains the optional parameters for the LiveEventsClient.BeginAllocate method.
type LiveEventsClientBeginAllocateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginCreateOptions contains the optional parameters for the LiveEventsClient.BeginCreate method.
type LiveEventsClientBeginCreateOptions struct {
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginDeleteOptions contains the optional parameters for the LiveEventsClient.BeginDelete method.
type LiveEventsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginListGetStatusOptions contains the optional parameters for the LiveEventsClient.BeginListGetStatus
// method.
type LiveEventsClientBeginListGetStatusOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginListGetStreamEventsOptions contains the optional parameters for the LiveEventsClient.BeginListGetStreamEvents
// method.
type LiveEventsClientBeginListGetStreamEventsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginListGetTrackIngestHeartbeatsOptions contains the optional parameters for the LiveEventsClient.BeginListGetTrackIngestHeartbeats
// method.
type LiveEventsClientBeginListGetTrackIngestHeartbeatsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginResetOptions contains the optional parameters for the LiveEventsClient.BeginReset method.
type LiveEventsClientBeginResetOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginStartOptions contains the optional parameters for the LiveEventsClient.BeginStart method.
type LiveEventsClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginStopOptions contains the optional parameters for the LiveEventsClient.BeginStop method.
type LiveEventsClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientBeginUpdateOptions contains the optional parameters for the LiveEventsClient.BeginUpdate method.
type LiveEventsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveEventsClientGetOptions contains the optional parameters for the LiveEventsClient.Get method.
type LiveEventsClientGetOptions struct {
	// placeholder for future optional parameters
}

// LiveEventsClientListOptions contains the optional parameters for the LiveEventsClient.NewListPager method.
type LiveEventsClientListOptions struct {
	// placeholder for future optional parameters
}

// LiveEventsClientOperationLocationOptions contains the optional parameters for the LiveEventsClient.OperationLocation method.
type LiveEventsClientOperationLocationOptions struct {
	// placeholder for future optional parameters
}

// LiveOutputsClientAsyncOperationOptions contains the optional parameters for the LiveOutputsClient.AsyncOperation method.
type LiveOutputsClientAsyncOperationOptions struct {
	// placeholder for future optional parameters
}

// LiveOutputsClientBeginCreateOptions contains the optional parameters for the LiveOutputsClient.BeginCreate method.
type LiveOutputsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveOutputsClientBeginDeleteOptions contains the optional parameters for the LiveOutputsClient.BeginDelete method.
type LiveOutputsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LiveOutputsClientGetOptions contains the optional parameters for the LiveOutputsClient.Get method.
type LiveOutputsClientGetOptions struct {
	// placeholder for future optional parameters
}

// LiveOutputsClientListOptions contains the optional parameters for the LiveOutputsClient.NewListPager method.
type LiveOutputsClientListOptions struct {
	// placeholder for future optional parameters
}

// LiveOutputsClientOperationLocationOptions contains the optional parameters for the LiveOutputsClient.OperationLocation
// method.
type LiveOutputsClientOperationLocationOptions struct {
	// placeholder for future optional parameters
}

// LocationsClientCheckNameAvailabilityOptions contains the optional parameters for the LocationsClient.CheckNameAvailability
// method.
type LocationsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// OperationResultsClientGetOptions contains the optional parameters for the OperationResultsClient.Get method.
type OperationResultsClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationStatusesClientGetOptions contains the optional parameters for the OperationStatusesClient.Get method.
type OperationStatusesClientGetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
// method.
type PrivateEndpointConnectionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Delete
// method.
type PrivateEndpointConnectionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.List method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// StreamingEndpointsClientAsyncOperationOptions contains the optional parameters for the StreamingEndpointsClient.AsyncOperation
// method.
type StreamingEndpointsClientAsyncOperationOptions struct {
	// placeholder for future optional parameters
}

// StreamingEndpointsClientBeginCreateOptions contains the optional parameters for the StreamingEndpointsClient.BeginCreate
// method.
type StreamingEndpointsClientBeginCreateOptions struct {
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientBeginDeleteOptions contains the optional parameters for the StreamingEndpointsClient.BeginDelete
// method.
type StreamingEndpointsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientBeginScaleOptions contains the optional parameters for the StreamingEndpointsClient.BeginScale
// method.
type StreamingEndpointsClientBeginScaleOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientBeginStartOptions contains the optional parameters for the StreamingEndpointsClient.BeginStart
// method.
type StreamingEndpointsClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientBeginStopOptions contains the optional parameters for the StreamingEndpointsClient.BeginStop method.
type StreamingEndpointsClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientBeginUpdateOptions contains the optional parameters for the StreamingEndpointsClient.BeginUpdate
// method.
type StreamingEndpointsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StreamingEndpointsClientGetOptions contains the optional parameters for the StreamingEndpointsClient.Get method.
type StreamingEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StreamingEndpointsClientListOptions contains the optional parameters for the StreamingEndpointsClient.NewListPager method.
type StreamingEndpointsClientListOptions struct {
	// placeholder for future optional parameters
}

// StreamingEndpointsClientOperationLocationOptions contains the optional parameters for the StreamingEndpointsClient.OperationLocation
// method.
type StreamingEndpointsClientOperationLocationOptions struct {
	// placeholder for future optional parameters
}

// StreamingEndpointsClientSKUsOptions contains the optional parameters for the StreamingEndpointsClient.SKUs method.
type StreamingEndpointsClientSKUsOptions struct {
	// placeholder for future optional parameters
}

// StreamingLocatorsClientCreateOptions contains the optional parameters for the StreamingLocatorsClient.Create method.
type StreamingLocatorsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// StreamingLocatorsClientDeleteOptions contains the optional parameters for the StreamingLocatorsClient.Delete method.
type StreamingLocatorsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// StreamingLocatorsClientGetOptions contains the optional parameters for the StreamingLocatorsClient.Get method.
type StreamingLocatorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StreamingLocatorsClientListContentKeysOptions contains the optional parameters for the StreamingLocatorsClient.ListContentKeys
// method.
type StreamingLocatorsClientListContentKeysOptions struct {
	// placeholder for future optional parameters
}

// StreamingLocatorsClientListOptions contains the optional parameters for the StreamingLocatorsClient.NewListPager method.
type StreamingLocatorsClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string

	// Specifies a non-negative integer n that limits the number of items returned from a collection. The service returns the
	// number of available items up to but not greater than the specified value n.
	Top *int32
}

// StreamingLocatorsClientListPathsOptions contains the optional parameters for the StreamingLocatorsClient.ListPaths method.
type StreamingLocatorsClientListPathsOptions struct {
	// placeholder for future optional parameters
}

// StreamingPoliciesClientCreateOptions contains the optional parameters for the StreamingPoliciesClient.Create method.
type StreamingPoliciesClientCreateOptions struct {
	// placeholder for future optional parameters
}

// StreamingPoliciesClientDeleteOptions contains the optional parameters for the StreamingPoliciesClient.Delete method.
type StreamingPoliciesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// StreamingPoliciesClientGetOptions contains the optional parameters for the StreamingPoliciesClient.Get method.
type StreamingPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// StreamingPoliciesClientListOptions contains the optional parameters for the StreamingPoliciesClient.NewListPager method.
type StreamingPoliciesClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string

	// Specifies a non-negative integer n that limits the number of items returned from a collection. The service returns the
	// number of available items up to but not greater than the specified value n.
	Top *int32
}

// TracksClientBeginCreateOrUpdateOptions contains the optional parameters for the TracksClient.BeginCreateOrUpdate method.
type TracksClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TracksClientBeginDeleteOptions contains the optional parameters for the TracksClient.BeginDelete method.
type TracksClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TracksClientBeginUpdateOptions contains the optional parameters for the TracksClient.BeginUpdate method.
type TracksClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TracksClientBeginUpdateTrackDataOptions contains the optional parameters for the TracksClient.BeginUpdateTrackData method.
type TracksClientBeginUpdateTrackDataOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TracksClientGetOptions contains the optional parameters for the TracksClient.Get method.
type TracksClientGetOptions struct {
	// placeholder for future optional parameters
}

// TracksClientListOptions contains the optional parameters for the TracksClient.NewListPager method.
type TracksClientListOptions struct {
	// placeholder for future optional parameters
}

// TransformsClientCreateOrUpdateOptions contains the optional parameters for the TransformsClient.CreateOrUpdate method.
type TransformsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TransformsClientDeleteOptions contains the optional parameters for the TransformsClient.Delete method.
type TransformsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TransformsClientGetOptions contains the optional parameters for the TransformsClient.Get method.
type TransformsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TransformsClientListOptions contains the optional parameters for the TransformsClient.NewListPager method.
type TransformsClientListOptions struct {
	// Restricts the set of items returned.
	Filter *string

	// Specifies the key by which the result collection should be ordered.
	Orderby *string
}

// TransformsClientUpdateOptions contains the optional parameters for the TransformsClient.Update method.
type TransformsClientUpdateOptions struct {
	// placeholder for future optional parameters
}
