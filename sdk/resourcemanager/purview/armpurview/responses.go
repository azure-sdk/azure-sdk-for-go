//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

// AccountsClientAddRootCollectionAdminResponse contains the response from method AccountsClient.AddRootCollectionAdmin.
type AccountsClientAddRootCollectionAdminResponse struct {
	// placeholder for future response values
}

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	// The response payload for CheckNameAvailability API
	CheckNameAvailabilityResult
}

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.BeginCreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	// Account resource
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// Account resource
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// Paged list of Account resources
	AccountList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	// Paged list of Account resources
	AccountList
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	// The Purview Account access keys.
	AccessKeys
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	// Account resource
	Account
}

// DefaultAccountsClientGetResponse contains the response from method DefaultAccountsClient.Get.
type DefaultAccountsClientGetResponse struct {
	// Payload to get and set the default account in the given scope
	DefaultAccountPayload
}

// DefaultAccountsClientRemoveResponse contains the response from method DefaultAccountsClient.Remove.
type DefaultAccountsClientRemoveResponse struct {
	// placeholder for future response values
}

// DefaultAccountsClientSetResponse contains the response from method DefaultAccountsClient.Set.
type DefaultAccountsClientSetResponse struct {
	// Payload to get and set the default account in the given scope
	DefaultAccountPayload
}

// FeaturesClientAccountGetResponse contains the response from method FeaturesClient.AccountGet.
type FeaturesClientAccountGetResponse struct {
	// List of features with enabled status
	BatchFeatureStatus
}

// FeaturesClientSubscriptionGetResponse contains the response from method FeaturesClient.SubscriptionGet.
type FeaturesClientSubscriptionGetResponse struct {
	// List of features with enabled status
	BatchFeatureStatus
}

// IngestionPrivateEndpointConnectionsClientListResponse contains the response from method IngestionPrivateEndpointConnectionsClient.NewListPager.
type IngestionPrivateEndpointConnectionsClientListResponse struct {
	// Paged list of private endpoint connections
	PrivateEndpointConnectionList
}

// IngestionPrivateEndpointConnectionsClientUpdateStatusResponse contains the response from method IngestionPrivateEndpointConnectionsClient.UpdateStatus.
type IngestionPrivateEndpointConnectionsClientUpdateStatusResponse struct {
	// A private endpoint connection status update response class.
	PrivateEndpointConnectionStatusUpdateResponse
}

// KafkaConfigurationsClientCreateOrUpdateResponse contains the response from method KafkaConfigurationsClient.CreateOrUpdate.
type KafkaConfigurationsClientCreateOrUpdateResponse struct {
	// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
	KafkaConfiguration
}

// KafkaConfigurationsClientDeleteResponse contains the response from method KafkaConfigurationsClient.Delete.
type KafkaConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// KafkaConfigurationsClientGetResponse contains the response from method KafkaConfigurationsClient.Get.
type KafkaConfigurationsClientGetResponse struct {
	// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
	KafkaConfiguration
}

// KafkaConfigurationsClientListByAccountResponse contains the response from method KafkaConfigurationsClient.NewListByAccountPager.
type KafkaConfigurationsClientListByAccountResponse struct {
	// Paged list of kafka configuration resources
	KafkaConfigurationList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Paged list of operation resources
	OperationList
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// A private endpoint connection class.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// A private endpoint connection class.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByAccountResponse contains the response from method PrivateEndpointConnectionsClient.NewListByAccountPager.
type PrivateEndpointConnectionsClientListByAccountResponse struct {
	// Paged list of private endpoint connections
	PrivateEndpointConnectionList
}

// PrivateLinkResourcesClientGetByGroupIDResponse contains the response from method PrivateLinkResourcesClient.GetByGroupID.
type PrivateLinkResourcesClientGetByGroupIDResponse struct {
	// A privately linkable resource.
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByAccountResponse contains the response from method PrivateLinkResourcesClient.NewListByAccountPager.
type PrivateLinkResourcesClientListByAccountResponse struct {
	// Paged list of private link resources
	PrivateLinkResourceList
}

// UsagesClientGetResponse contains the response from method UsagesClient.Get.
type UsagesClientGetResponse struct {
	// List of usage information
	UsageList
}
