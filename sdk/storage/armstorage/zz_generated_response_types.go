// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"net/http"
	"time"
)

// BlobContainerResponse is the response envelope for operations that return a BlobContainer type.
type BlobContainerResponse struct {
	// Properties of the blob container, including Id, resource name, resource type, Etag.
	BlobContainer *BlobContainer

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlobInventoryPolicyResponse is the response envelope for operations that return a BlobInventoryPolicy type.
type BlobInventoryPolicyResponse struct {
	// The storage account blob inventory policy.
	BlobInventoryPolicy *BlobInventoryPolicy

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlobRestoreStatusPollerResponse is the response envelope for operations that asynchronously return a BlobRestoreStatus type.
type BlobRestoreStatusPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (BlobRestoreStatusResponse, error)

	// Poller contains an initialized poller.
	Poller BlobRestoreStatusPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlobRestoreStatusResponse is the response envelope for operations that return a BlobRestoreStatus type.
type BlobRestoreStatusResponse struct {
	// Blob restore status.
	BlobRestoreStatus *BlobRestoreStatus

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlobServiceItemsResponse is the response envelope for operations that return a BlobServiceItems type.
type BlobServiceItemsResponse struct {
	BlobServiceItems *BlobServiceItems

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlobServicePropertiesResponse is the response envelope for operations that return a BlobServiceProperties type.
type BlobServicePropertiesResponse struct {
	// The properties of a storage account’s Blob service.
	BlobServiceProperties *BlobServiceProperties

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// CheckNameAvailabilityResultResponse is the response envelope for operations that return a CheckNameAvailabilityResult type.
type CheckNameAvailabilityResultResponse struct {
	// The CheckNameAvailability operation response.
	CheckNameAvailabilityResult *CheckNameAvailabilityResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeletedAccountListResultResponse is the response envelope for operations that return a DeletedAccountListResult type.
type DeletedAccountListResultResponse struct {
	// The response from the List Deleted Accounts operation.
	DeletedAccountListResult *DeletedAccountListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeletedAccountResponse is the response envelope for operations that return a DeletedAccount type.
type DeletedAccountResponse struct {
	// Deleted storage account
	DeletedAccount *DeletedAccount

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EncryptionScopeListResultResponse is the response envelope for operations that return a EncryptionScopeListResult type.
type EncryptionScopeListResultResponse struct {
	// List of encryption scopes requested, and if paging is required, a URL to the next page of encryption scopes.
	EncryptionScopeListResult *EncryptionScopeListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EncryptionScopeResponse is the response envelope for operations that return a EncryptionScope type.
type EncryptionScopeResponse struct {
	// The Encryption Scope resource.
	EncryptionScope *EncryptionScope

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FileServiceItemsResponse is the response envelope for operations that return a FileServiceItems type.
type FileServiceItemsResponse struct {
	FileServiceItems *FileServiceItems

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FileServicePropertiesResponse is the response envelope for operations that return a FileServiceProperties type.
type FileServicePropertiesResponse struct {
	// The properties of File services in storage account.
	FileServiceProperties *FileServiceProperties

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FileShareItemsResponse is the response envelope for operations that return a FileShareItems type.
type FileShareItemsResponse struct {
	// Response schema. Contains list of shares returned, and if paging is requested or required, a URL to next page of shares.
	FileShareItems *FileShareItems

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FileShareResponse is the response envelope for operations that return a FileShare type.
type FileShareResponse struct {
	// Properties of the file share, including Id, resource name, resource type, Etag.
	FileShare *FileShare

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPPollerResponse contains the asynchronous HTTP response from the call to the service endpoint.
type HTTPPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*http.Response, error)

	// Poller contains an initialized poller.
	Poller HTTPPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ImmutabilityPolicyResponse is the response envelope for operations that return a ImmutabilityPolicy type.
type ImmutabilityPolicyResponse struct {
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy *ImmutabilityPolicy

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LeaseContainerResponseResponse is the response envelope for operations that return a LeaseContainerResponse type.
type LeaseContainerResponseResponse struct {
	// Lease Container response schema.
	LeaseContainerResponse *LeaseContainerResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LeaseShareResponseResponse is the response envelope for operations that return a LeaseShareResponse type.
type LeaseShareResponseResponse struct {
	// ETag contains the information returned from the ETag header response.
	ETag *string

	// Lease Share response schema.
	LeaseShareResponse *LeaseShareResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// LegalHoldResponse is the response envelope for operations that return a LegalHold type.
type LegalHoldResponse struct {
	// The LegalHold property of a blob container.
	LegalHold *LegalHold

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListAccountSasResponseResponse is the response envelope for operations that return a ListAccountSasResponse type.
type ListAccountSasResponseResponse struct {
	// The List SAS credentials operation response.
	ListAccountSasResponse *ListAccountSasResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListBlobInventoryPolicyResponse is the response envelope for operations that return a ListBlobInventoryPolicy type.
type ListBlobInventoryPolicyResponse struct {
	// List of blob inventory policies returned.
	ListBlobInventoryPolicy *ListBlobInventoryPolicy

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListContainerItemsResponse is the response envelope for operations that return a ListContainerItems type.
type ListContainerItemsResponse struct {
	// Response schema. Contains list of blobs returned, and if paging is requested or required, a URL to next page of containers.
	ListContainerItems *ListContainerItems

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListQueueResourceResponse is the response envelope for operations that return a ListQueueResource type.
type ListQueueResourceResponse struct {
	// Response schema. Contains list of queues returned
	ListQueueResource *ListQueueResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListQueueServicesResponse is the response envelope for operations that return a ListQueueServices type.
type ListQueueServicesResponse struct {
	ListQueueServices *ListQueueServices

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListServiceSasResponseResponse is the response envelope for operations that return a ListServiceSasResponse type.
type ListServiceSasResponseResponse struct {
	// The List service SAS credentials operation response.
	ListServiceSasResponse *ListServiceSasResponse

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListTableResourceResponse is the response envelope for operations that return a ListTableResource type.
type ListTableResourceResponse struct {
	// Response schema. Contains list of tables returned
	ListTableResource *ListTableResource

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListTableServicesResponse is the response envelope for operations that return a ListTableServices type.
type ListTableServicesResponse struct {
	ListTableServices *ListTableServices

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ManagementPolicyResponse is the response envelope for operations that return a ManagementPolicy type.
type ManagementPolicyResponse struct {
	// The Get Storage Account ManagementPolicies operation response.
	ManagementPolicy *ManagementPolicy

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectReplicationPoliciesResponse is the response envelope for operations that return a ObjectReplicationPolicies type.
type ObjectReplicationPoliciesResponse struct {
	// List storage account object replication policies.
	ObjectReplicationPolicies *ObjectReplicationPolicies

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectReplicationPolicyResponse is the response envelope for operations that return a ObjectReplicationPolicy type.
type ObjectReplicationPolicyResponse struct {
	// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
	ObjectReplicationPolicy *ObjectReplicationPolicy

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationListResultResponse is the response envelope for operations that return a OperationListResult type.
type OperationListResultResponse struct {
	// Result of the request to list Storage operations. It contains a list of operations and a URL link to get the next set of results.
	OperationListResult *OperationListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionListResultResponse is the response envelope for operations that return a PrivateEndpointConnectionListResult type.
type PrivateEndpointConnectionListResultResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResult *PrivateEndpointConnectionListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateEndpointConnectionResponse is the response envelope for operations that return a PrivateEndpointConnection type.
type PrivateEndpointConnectionResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection *PrivateEndpointConnection

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourceListResultResponse is the response envelope for operations that return a PrivateLinkResourceListResult type.
type PrivateLinkResourceListResultResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult *PrivateLinkResourceListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// QueueServicePropertiesResponse is the response envelope for operations that return a QueueServiceProperties type.
type QueueServicePropertiesResponse struct {
	// The properties of a storage account’s Queue service.
	QueueServiceProperties *QueueServiceProperties

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountListKeysResultResponse is the response envelope for operations that return a StorageAccountListKeysResult type.
type StorageAccountListKeysResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The response from the ListKeys operation.
	StorageAccountListKeysResult *StorageAccountListKeysResult
}

// StorageAccountListResultResponse is the response envelope for operations that return a StorageAccountListResult type.
type StorageAccountListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The response from the List Storage Accounts operation.
	StorageAccountListResult *StorageAccountListResult
}

// StorageAccountPollerResponse is the response envelope for operations that asynchronously return a StorageAccount type.
type StorageAccountPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (StorageAccountResponse, error)

	// Poller contains an initialized poller.
	Poller StorageAccountPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// StorageAccountResponse is the response envelope for operations that return a StorageAccount type.
type StorageAccountResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The storage account.
	StorageAccount *StorageAccount
}

// StorageQueueResponse is the response envelope for operations that return a StorageQueue type.
type StorageQueueResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse  *http.Response
	StorageQueue *StorageQueue
}

// StorageSKUListResultResponse is the response envelope for operations that return a StorageSKUListResult type.
type StorageSKUListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The response from the List Storage SKUs operation.
	StorageSKUListResult *StorageSKUListResult
}

// TableResponse is the response envelope for operations that return a Table type.
type TableResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Properties of the table, including Id, resource name, resource type.
	Table *Table
}

// TableServicePropertiesResponse is the response envelope for operations that return a TableServiceProperties type.
type TableServicePropertiesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The properties of a storage account’s Table service.
	TableServiceProperties *TableServiceProperties
}

// UsageListResultResponse is the response envelope for operations that return a UsageListResult type.
type UsageListResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// The response from the List Usages operation.
	UsageListResult *UsageListResult
}
