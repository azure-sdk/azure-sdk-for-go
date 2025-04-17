// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

// MoveCollectionsClientBulkRemoveResponse contains the response from method MoveCollectionsClient.BeginBulkRemove.
type MoveCollectionsClientBulkRemoveResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientCommitResponse contains the response from method MoveCollectionsClient.BeginCommit.
type MoveCollectionsClientCommitResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientCreateResponse contains the response from method MoveCollectionsClient.Create.
type MoveCollectionsClientCreateResponse struct {
	// Define the move collection.
	MoveCollection
}

// MoveCollectionsClientDeleteResponse contains the response from method MoveCollectionsClient.BeginDelete.
type MoveCollectionsClientDeleteResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientDiscardResponse contains the response from method MoveCollectionsClient.BeginDiscard.
type MoveCollectionsClientDiscardResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientGetResponse contains the response from method MoveCollectionsClient.Get.
type MoveCollectionsClientGetResponse struct {
	// Define the move collection.
	MoveCollection
}

// MoveCollectionsClientInitiateMoveResponse contains the response from method MoveCollectionsClient.BeginInitiateMove.
type MoveCollectionsClientInitiateMoveResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientListMoveCollectionsByResourceGroupResponse contains the response from method MoveCollectionsClient.NewListMoveCollectionsByResourceGroupPager.
type MoveCollectionsClientListMoveCollectionsByResourceGroupResponse struct {
	// Defines the collection of move collections.
	MoveCollectionResultList
}

// MoveCollectionsClientListMoveCollectionsBySubscriptionResponse contains the response from method MoveCollectionsClient.NewListMoveCollectionsBySubscriptionPager.
type MoveCollectionsClientListMoveCollectionsBySubscriptionResponse struct {
	// Defines the collection of move collections.
	MoveCollectionResultList
}

// MoveCollectionsClientListRequiredForResponse contains the response from method MoveCollectionsClient.ListRequiredFor.
type MoveCollectionsClientListRequiredForResponse struct {
	// Required for resources collection.
	RequiredForResourcesCollection
}

// MoveCollectionsClientPrepareResponse contains the response from method MoveCollectionsClient.BeginPrepare.
type MoveCollectionsClientPrepareResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientResolveDependenciesResponse contains the response from method MoveCollectionsClient.BeginResolveDependencies.
type MoveCollectionsClientResolveDependenciesResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveCollectionsClientUpdateResponse contains the response from method MoveCollectionsClient.Update.
type MoveCollectionsClientUpdateResponse struct {
	// Define the move collection.
	MoveCollection
}

// MoveResourcesClientCreateResponse contains the response from method MoveResourcesClient.BeginCreate.
type MoveResourcesClientCreateResponse struct {
	// Defines the move resource.
	MoveResource
}

// MoveResourcesClientDeleteResponse contains the response from method MoveResourcesClient.BeginDelete.
type MoveResourcesClientDeleteResponse struct {
	// Operation status REST resource.
	OperationStatus
}

// MoveResourcesClientGetResponse contains the response from method MoveResourcesClient.Get.
type MoveResourcesClientGetResponse struct {
	// Defines the move resource.
	MoveResource
}

// MoveResourcesClientListResponse contains the response from method MoveResourcesClient.NewListPager.
type MoveResourcesClientListResponse struct {
	// Defines the collection of move resources.
	MoveResourceCollection
}

// OperationsDiscoveryClientGetResponse contains the response from method OperationsDiscoveryClient.Get.
type OperationsDiscoveryClientGetResponse struct {
	// Collection of ClientDiscovery details.
	OperationsDiscoveryCollection
}

// UnresolvedDependenciesClientGetResponse contains the response from method UnresolvedDependenciesClient.NewGetPager.
type UnresolvedDependenciesClientGetResponse struct {
	// Unresolved dependency collection.
	UnresolvedDependencyCollection
}
