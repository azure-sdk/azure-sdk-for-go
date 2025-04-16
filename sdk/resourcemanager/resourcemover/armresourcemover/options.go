// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

// MoveCollectionsClientBeginBulkRemoveOptions contains the optional parameters for the MoveCollectionsClient.BeginBulkRemove
// method.
type MoveCollectionsClientBeginBulkRemoveOptions struct {
	Body *BulkRemoveRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginCommitOptions contains the optional parameters for the MoveCollectionsClient.BeginCommit method.
type MoveCollectionsClientBeginCommitOptions struct {
	Body *CommitRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginDeleteOptions contains the optional parameters for the MoveCollectionsClient.BeginDelete method.
type MoveCollectionsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginDiscardOptions contains the optional parameters for the MoveCollectionsClient.BeginDiscard method.
type MoveCollectionsClientBeginDiscardOptions struct {
	Body *DiscardRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginInitiateMoveOptions contains the optional parameters for the MoveCollectionsClient.BeginInitiateMove
// method.
type MoveCollectionsClientBeginInitiateMoveOptions struct {
	Body *ResourceMoveRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginPrepareOptions contains the optional parameters for the MoveCollectionsClient.BeginPrepare method.
type MoveCollectionsClientBeginPrepareOptions struct {
	Body *PrepareRequest

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientBeginResolveDependenciesOptions contains the optional parameters for the MoveCollectionsClient.BeginResolveDependencies
// method.
type MoveCollectionsClientBeginResolveDependenciesOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveCollectionsClientCreateOptions contains the optional parameters for the MoveCollectionsClient.Create method.
type MoveCollectionsClientCreateOptions struct {
	Body *MoveCollection
}

// MoveCollectionsClientGetOptions contains the optional parameters for the MoveCollectionsClient.Get method.
type MoveCollectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MoveCollectionsClientListMoveCollectionsByResourceGroupOptions contains the optional parameters for the MoveCollectionsClient.NewListMoveCollectionsByResourceGroupPager
// method.
type MoveCollectionsClientListMoveCollectionsByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MoveCollectionsClientListMoveCollectionsBySubscriptionOptions contains the optional parameters for the MoveCollectionsClient.NewListMoveCollectionsBySubscriptionPager
// method.
type MoveCollectionsClientListMoveCollectionsBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MoveCollectionsClientListRequiredForOptions contains the optional parameters for the MoveCollectionsClient.ListRequiredFor
// method.
type MoveCollectionsClientListRequiredForOptions struct {
	// placeholder for future optional parameters
}

// MoveCollectionsClientUpdateOptions contains the optional parameters for the MoveCollectionsClient.Update method.
type MoveCollectionsClientUpdateOptions struct {
	Body *UpdateMoveCollectionRequest
}

// MoveResourcesClientBeginCreateOptions contains the optional parameters for the MoveResourcesClient.BeginCreate method.
type MoveResourcesClientBeginCreateOptions struct {
	Body *MoveResource

	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveResourcesClientBeginDeleteOptions contains the optional parameters for the MoveResourcesClient.BeginDelete method.
type MoveResourcesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// MoveResourcesClientGetOptions contains the optional parameters for the MoveResourcesClient.Get method.
type MoveResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// MoveResourcesClientListOptions contains the optional parameters for the MoveResourcesClient.NewListPager method.
type MoveResourcesClientListOptions struct {
	// The filter to apply on the operation. For example, you can use $filter=Properties/ProvisioningState eq 'Succeeded'.
	Filter *string
}

// OperationsDiscoveryClientGetOptions contains the optional parameters for the OperationsDiscoveryClient.Get method.
type OperationsDiscoveryClientGetOptions struct {
	// placeholder for future optional parameters
}

// UnresolvedDependenciesClientGetOptions contains the optional parameters for the UnresolvedDependenciesClient.NewGetPager
// method.
type UnresolvedDependenciesClientGetOptions struct {
	// Defines the dependency level.
	DependencyLevel *DependencyLevel

	// The filter to apply on the operation. For example, $apply=filter(count eq 2).
	Filter *string

	// OData order by query option. For example, you can use $orderby=Count desc.
	Orderby *string
}
