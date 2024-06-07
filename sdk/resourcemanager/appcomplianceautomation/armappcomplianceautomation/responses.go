//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcomplianceautomation

// EvidenceClientCreateOrUpdateResponse contains the response from method EvidenceClient.CreateOrUpdate.
type EvidenceClientCreateOrUpdateResponse struct {
	// A class represent an AppComplianceAutomation evidence resource.
	EvidenceResource
}

// EvidenceClientDeleteResponse contains the response from method EvidenceClient.Delete.
type EvidenceClientDeleteResponse struct {
	// placeholder for future response values
}

// EvidenceClientDownloadResponse contains the response from method EvidenceClient.Download.
type EvidenceClientDownloadResponse struct {
	// Object that includes all the possible response for the evidence file download operation.
	EvidenceFileDownloadResponse
}

// EvidenceClientGetResponse contains the response from method EvidenceClient.Get.
type EvidenceClientGetResponse struct {
	// A class represent an AppComplianceAutomation evidence resource.
	EvidenceResource
}

// EvidenceClientListByReportResponse contains the response from method EvidenceClient.NewListByReportPager.
type EvidenceClientListByReportResponse struct {
	// The response of a EvidenceResource list operation.
	EvidenceResourceListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ProviderActionsClientCheckNameAvailabilityResponse contains the response from method ProviderActionsClient.CheckNameAvailability.
type ProviderActionsClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// ProviderActionsClientGetCollectionCountResponse contains the response from method ProviderActionsClient.GetCollectionCount.
type ProviderActionsClientGetCollectionCountResponse struct {
	// The get collection count response.
	GetCollectionCountResponse
}

// ProviderActionsClientGetOverviewStatusResponse contains the response from method ProviderActionsClient.GetOverviewStatus.
type ProviderActionsClientGetOverviewStatusResponse struct {
	// The get overview status response.
	GetOverviewStatusResponse
}

// ProviderActionsClientListInUseStorageAccountsResponse contains the response from method ProviderActionsClient.ListInUseStorageAccounts.
type ProviderActionsClientListInUseStorageAccountsResponse struct {
	// Parameters for listing in use storage accounts operation. If subscription list is null, it will check the user's all subscriptions.
	ListInUseStorageAccountsResponse
}

// ProviderActionsClientOnboardResponse contains the response from method ProviderActionsClient.BeginOnboard.
type ProviderActionsClientOnboardResponse struct {
	// Success. The response indicates given subscriptions has been onboarded.
	OnboardResponse
}

// ProviderActionsClientTriggerEvaluationResponse contains the response from method ProviderActionsClient.BeginTriggerEvaluation.
type ProviderActionsClientTriggerEvaluationResponse struct {
	// Trigger evaluation response.
	TriggerEvaluationResponse
}

// ReportClientCreateOrUpdateResponse contains the response from method ReportClient.BeginCreateOrUpdate.
type ReportClientCreateOrUpdateResponse struct {
	// A class represent an AppComplianceAutomation report resource.
	ReportResource
}

// ReportClientDeleteResponse contains the response from method ReportClient.BeginDelete.
type ReportClientDeleteResponse struct {
	// placeholder for future response values
}

// ReportClientFixResponse contains the response from method ReportClient.BeginFix.
type ReportClientFixResponse struct {
	// Report fix result.
	ReportFixResult
}

// ReportClientGetResponse contains the response from method ReportClient.Get.
type ReportClientGetResponse struct {
	// A class represent an AppComplianceAutomation report resource.
	ReportResource
}

// ReportClientGetScopingQuestionsResponse contains the response from method ReportClient.GetScopingQuestions.
type ReportClientGetScopingQuestionsResponse struct {
	// Scoping question list.
	ScopingQuestions
}

// ReportClientListResponse contains the response from method ReportClient.NewListPager.
type ReportClientListResponse struct {
	// The response of a ReportResource list operation.
	ReportResourceListResult
}

// ReportClientNestedResourceCheckNameAvailabilityResponse contains the response from method ReportClient.NestedResourceCheckNameAvailability.
type ReportClientNestedResourceCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// ReportClientSyncCertRecordResponse contains the response from method ReportClient.BeginSyncCertRecord.
type ReportClientSyncCertRecordResponse struct {
	// Synchronize certification record response.
	SyncCertRecordResponse
}

// ReportClientUpdateResponse contains the response from method ReportClient.BeginUpdate.
type ReportClientUpdateResponse struct {
	// A class represent an AppComplianceAutomation report resource.
	ReportResource
}

// ReportClientVerifyResponse contains the response from method ReportClient.BeginVerify.
type ReportClientVerifyResponse struct {
	// Report health status verification result.
	ReportVerificationResult
}

// ScopingConfigurationClientCreateOrUpdateResponse contains the response from method ScopingConfigurationClient.CreateOrUpdate.
type ScopingConfigurationClientCreateOrUpdateResponse struct {
	// A class represent an AppComplianceAutomation scoping configuration resource.
	ScopingConfigurationResource
}

// ScopingConfigurationClientDeleteResponse contains the response from method ScopingConfigurationClient.Delete.
type ScopingConfigurationClientDeleteResponse struct {
	// placeholder for future response values
}

// ScopingConfigurationClientGetResponse contains the response from method ScopingConfigurationClient.Get.
type ScopingConfigurationClientGetResponse struct {
	// A class represent an AppComplianceAutomation scoping configuration resource.
	ScopingConfigurationResource
}

// ScopingConfigurationClientListResponse contains the response from method ScopingConfigurationClient.NewListPager.
type ScopingConfigurationClientListResponse struct {
	// The response of a ScopingConfigurationResource list operation.
	ScopingConfigurationResourceListResult
}

// SnapshotClientDownloadResponse contains the response from method SnapshotClient.BeginDownload.
type SnapshotClientDownloadResponse struct {
	// Object that includes all the possible response for the download operation.
	DownloadResponse
}

// SnapshotClientGetResponse contains the response from method SnapshotClient.Get.
type SnapshotClientGetResponse struct {
	// A class represent a AppComplianceAutomation snapshot resource.
	SnapshotResource
}

// SnapshotClientListResponse contains the response from method SnapshotClient.NewListPager.
type SnapshotClientListResponse struct {
	// The response of a SnapshotResource list operation.
	SnapshotResourceListResult
}

// WebhookClientCreateOrUpdateResponse contains the response from method WebhookClient.CreateOrUpdate.
type WebhookClientCreateOrUpdateResponse struct {
	// A class represent an AppComplianceAutomation webhook resource.
	WebhookResource
}

// WebhookClientDeleteResponse contains the response from method WebhookClient.Delete.
type WebhookClientDeleteResponse struct {
	// placeholder for future response values
}

// WebhookClientGetResponse contains the response from method WebhookClient.Get.
type WebhookClientGetResponse struct {
	// A class represent an AppComplianceAutomation webhook resource.
	WebhookResource
}

// WebhookClientListResponse contains the response from method WebhookClient.NewListPager.
type WebhookClientListResponse struct {
	// The response of a WebhookResource list operation.
	WebhookResourceListResult
}

// WebhookClientUpdateResponse contains the response from method WebhookClient.Update.
type WebhookClientUpdateResponse struct {
	// A class represent an AppComplianceAutomation webhook resource.
	WebhookResource
}