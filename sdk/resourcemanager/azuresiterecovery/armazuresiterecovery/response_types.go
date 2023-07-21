//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuresiterecovery

// DraClientCreateResponse contains the response from method DraClient.BeginCreate.
type DraClientCreateResponse struct {
	// Dra model.
	DraModel
}

// DraClientDeleteResponse contains the response from method DraClient.BeginDelete.
type DraClientDeleteResponse struct {
	// placeholder for future response values
}

// DraClientGetResponse contains the response from method DraClient.Get.
type DraClientGetResponse struct {
	// Dra model.
	DraModel
}

// DraClientListResponse contains the response from method DraClient.NewListPager.
type DraClientListResponse struct {
	// Dra model collection.
	DraModelCollection
}

// DraOperationStatusClientGetResponse contains the response from method DraOperationStatusClient.Get.
type DraOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// EmailConfigurationClientCreateResponse contains the response from method EmailConfigurationClient.Create.
type EmailConfigurationClientCreateResponse struct {
	// Email configuration model.
	EmailConfigurationModel
}

// EmailConfigurationClientGetResponse contains the response from method EmailConfigurationClient.Get.
type EmailConfigurationClientGetResponse struct {
	// Email configuration model.
	EmailConfigurationModel
}

// EmailConfigurationClientListResponse contains the response from method EmailConfigurationClient.NewListPager.
type EmailConfigurationClientListResponse struct {
	// Email configuration model collection.
	EmailConfigurationModelCollection
}

// EventClientGetResponse contains the response from method EventClient.Get.
type EventClientGetResponse struct {
	// Event model.
	EventModel
}

// EventClientListResponse contains the response from method EventClient.NewListPager.
type EventClientListResponse struct {
	// Event model collection.
	EventModelCollection
}

// FabricClientCreateResponse contains the response from method FabricClient.BeginCreate.
type FabricClientCreateResponse struct {
	// Fabric model.
	FabricModel
}

// FabricClientDeleteResponse contains the response from method FabricClient.BeginDelete.
type FabricClientDeleteResponse struct {
	// placeholder for future response values
}

// FabricClientGetResponse contains the response from method FabricClient.Get.
type FabricClientGetResponse struct {
	// Fabric model.
	FabricModel
}

// FabricClientListBySubscriptionResponse contains the response from method FabricClient.NewListBySubscriptionPager.
type FabricClientListBySubscriptionResponse struct {
	// Fabric model collection.
	FabricModelCollection
}

// FabricClientListResponse contains the response from method FabricClient.NewListPager.
type FabricClientListResponse struct {
	// Fabric model collection.
	FabricModelCollection
}

// FabricClientUpdateResponse contains the response from method FabricClient.BeginUpdate.
type FabricClientUpdateResponse struct {
	// Fabric model.
	FabricModel
}

// FabricOperationsStatusClientGetResponse contains the response from method FabricOperationsStatusClient.Get.
type FabricOperationsStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// ManagementServiceAPIClientCheckNameAvailabilityResponse contains the response from method ManagementServiceAPIClient.CheckNameAvailability.
type ManagementServiceAPIClientCheckNameAvailabilityResponse struct {
	// Check name availability response model.
	CheckNameAvailabilityResponseModel
}

// ManagementServiceAPIClientDeploymentPreflightResponse contains the response from method ManagementServiceAPIClient.DeploymentPreflight.
type ManagementServiceAPIClientDeploymentPreflightResponse struct {
	// Deployment preflight model.
	DeploymentPreflightModel
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// Array of OperationModel
	OperationModelArray []*OperationModel
}

// PolicyClientCreateResponse contains the response from method PolicyClient.BeginCreate.
type PolicyClientCreateResponse struct {
	// Policy model.
	PolicyModel
}

// PolicyClientDeleteResponse contains the response from method PolicyClient.BeginDelete.
type PolicyClientDeleteResponse struct {
	// placeholder for future response values
}

// PolicyClientGetResponse contains the response from method PolicyClient.Get.
type PolicyClientGetResponse struct {
	// Policy model.
	PolicyModel
}

// PolicyClientListResponse contains the response from method PolicyClient.NewListPager.
type PolicyClientListResponse struct {
	// Policy model collection.
	PolicyModelCollection
}

// PolicyOperationStatusClientGetResponse contains the response from method PolicyOperationStatusClient.Get.
type PolicyOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// ProtectedItemClientCreateResponse contains the response from method ProtectedItemClient.BeginCreate.
type ProtectedItemClientCreateResponse struct {
	// Protected item model.
	ProtectedItemModel
}

// ProtectedItemClientDeleteResponse contains the response from method ProtectedItemClient.BeginDelete.
type ProtectedItemClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectedItemClientGetResponse contains the response from method ProtectedItemClient.Get.
type ProtectedItemClientGetResponse struct {
	// Protected item model.
	ProtectedItemModel
}

// ProtectedItemClientListResponse contains the response from method ProtectedItemClient.NewListPager.
type ProtectedItemClientListResponse struct {
	// Protected item model collection.
	ProtectedItemModelCollection
}

// ProtectedItemClientPlannedFailoverResponse contains the response from method ProtectedItemClient.BeginPlannedFailover.
type ProtectedItemClientPlannedFailoverResponse struct {
	// Planned failover model.
	PlannedFailoverModel
}

// ProtectedItemOperationStatusClientGetResponse contains the response from method ProtectedItemOperationStatusClient.Get.
type ProtectedItemOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// RecoveryPointsClientGetResponse contains the response from method RecoveryPointsClient.Get.
type RecoveryPointsClientGetResponse struct {
	// Recovery point model.
	RecoveryPointModel
}

// RecoveryPointsClientListResponse contains the response from method RecoveryPointsClient.NewListPager.
type RecoveryPointsClientListResponse struct {
	// Recovery point model collection.
	RecoveryPointModelCollection
}

// ReplicationExtensionClientCreateResponse contains the response from method ReplicationExtensionClient.BeginCreate.
type ReplicationExtensionClientCreateResponse struct {
	// Replication extension model.
	ReplicationExtensionModel
}

// ReplicationExtensionClientDeleteResponse contains the response from method ReplicationExtensionClient.BeginDelete.
type ReplicationExtensionClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationExtensionClientGetResponse contains the response from method ReplicationExtensionClient.Get.
type ReplicationExtensionClientGetResponse struct {
	// Replication extension model.
	ReplicationExtensionModel
}

// ReplicationExtensionClientListResponse contains the response from method ReplicationExtensionClient.NewListPager.
type ReplicationExtensionClientListResponse struct {
	// Replication extension model collection.
	ReplicationExtensionModelCollection
}

// ReplicationExtensionOperationStatusClientGetResponse contains the response from method ReplicationExtensionOperationStatusClient.Get.
type ReplicationExtensionOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// VaultClientCreateResponse contains the response from method VaultClient.BeginCreate.
type VaultClientCreateResponse struct {
	// Vault model.
	VaultModel
}

// VaultClientDeleteResponse contains the response from method VaultClient.BeginDelete.
type VaultClientDeleteResponse struct {
	// placeholder for future response values
}

// VaultClientGetResponse contains the response from method VaultClient.Get.
type VaultClientGetResponse struct {
	// Vault model.
	VaultModel
}

// VaultClientListBySubscriptionResponse contains the response from method VaultClient.NewListBySubscriptionPager.
type VaultClientListBySubscriptionResponse struct {
	// Vault model collection.
	VaultModelCollection
}

// VaultClientListResponse contains the response from method VaultClient.NewListPager.
type VaultClientListResponse struct {
	// Vault model collection.
	VaultModelCollection
}

// VaultClientUpdateResponse contains the response from method VaultClient.BeginUpdate.
type VaultClientUpdateResponse struct {
	// Vault model.
	VaultModel
}

// VaultOperationStatusClientGetResponse contains the response from method VaultOperationStatusClient.Get.
type VaultOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}

// WorkflowClientGetResponse contains the response from method WorkflowClient.Get.
type WorkflowClientGetResponse struct {
	// Workflow model.
	WorkflowModel
}

// WorkflowClientListResponse contains the response from method WorkflowClient.NewListPager.
type WorkflowClientListResponse struct {
	// Workflow model collection.
	WorkflowModelCollection
}

// WorkflowOperationStatusClientGetResponse contains the response from method WorkflowOperationStatusClient.Get.
type WorkflowOperationStatusClientGetResponse struct {
	// Defines the operation status.
	OperationStatus
}
