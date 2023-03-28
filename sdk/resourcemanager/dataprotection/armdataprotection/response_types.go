//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection

// BackupInstancesClientAdhocBackupResponse contains the response from method BackupInstancesClient.BeginAdhocBackup.
type BackupInstancesClientAdhocBackupResponse struct {
	OperationJobExtendedInfo
}

// BackupInstancesClientCreateOrUpdateResponse contains the response from method BackupInstancesClient.BeginCreateOrUpdate.
type BackupInstancesClientCreateOrUpdateResponse struct {
	BackupInstanceResource
}

// BackupInstancesClientDeleteResponse contains the response from method BackupInstancesClient.BeginDelete.
type BackupInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientGetBackupInstanceOperationResultResponse contains the response from method BackupInstancesClient.GetBackupInstanceOperationResult.
type BackupInstancesClientGetBackupInstanceOperationResultResponse struct {
	BackupInstanceResource
}

// BackupInstancesClientGetResponse contains the response from method BackupInstancesClient.Get.
type BackupInstancesClientGetResponse struct {
	BackupInstanceResource
}

// BackupInstancesClientListResponse contains the response from method BackupInstancesClient.NewListPager.
type BackupInstancesClientListResponse struct {
	BackupInstanceResourceList
}

// BackupInstancesClientResumeBackupsResponse contains the response from method BackupInstancesClient.BeginResumeBackups.
type BackupInstancesClientResumeBackupsResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientResumeProtectionResponse contains the response from method BackupInstancesClient.BeginResumeProtection.
type BackupInstancesClientResumeProtectionResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientStopProtectionResponse contains the response from method BackupInstancesClient.BeginStopProtection.
type BackupInstancesClientStopProtectionResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientSuspendBackupsResponse contains the response from method BackupInstancesClient.BeginSuspendBackups.
type BackupInstancesClientSuspendBackupsResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientSyncBackupInstanceResponse contains the response from method BackupInstancesClient.BeginSyncBackupInstance.
type BackupInstancesClientSyncBackupInstanceResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientTriggerRehydrateResponse contains the response from method BackupInstancesClient.BeginTriggerRehydrate.
type BackupInstancesClientTriggerRehydrateResponse struct {
	// placeholder for future response values
}

// BackupInstancesClientTriggerRestoreResponse contains the response from method BackupInstancesClient.BeginTriggerRestore.
type BackupInstancesClientTriggerRestoreResponse struct {
	OperationJobExtendedInfo
}

// BackupInstancesClientValidateForBackupResponse contains the response from method BackupInstancesClient.BeginValidateForBackup.
type BackupInstancesClientValidateForBackupResponse struct {
	OperationJobExtendedInfo
}

// BackupInstancesClientValidateForRestoreResponse contains the response from method BackupInstancesClient.BeginValidateForRestore.
type BackupInstancesClientValidateForRestoreResponse struct {
	OperationJobExtendedInfo
}

// BackupPoliciesClientCreateOrUpdateResponse contains the response from method BackupPoliciesClient.CreateOrUpdate.
type BackupPoliciesClientCreateOrUpdateResponse struct {
	BaseBackupPolicyResource
}

// BackupPoliciesClientDeleteResponse contains the response from method BackupPoliciesClient.Delete.
type BackupPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupPoliciesClientGetResponse contains the response from method BackupPoliciesClient.Get.
type BackupPoliciesClientGetResponse struct {
	BaseBackupPolicyResource
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.NewListPager.
type BackupPoliciesClientListResponse struct {
	BaseBackupPolicyResourceList
}

// BackupVaultOperationResultsClientGetResponse contains the response from method BackupVaultOperationResultsClient.Get.
type BackupVaultOperationResultsClientGetResponse struct {
	BackupVaultResource
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string

	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// BackupVaultsClientCheckNameAvailabilityResponse contains the response from method BackupVaultsClient.CheckNameAvailability.
type BackupVaultsClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResult
}

// BackupVaultsClientCreateOrUpdateResponse contains the response from method BackupVaultsClient.BeginCreateOrUpdate.
type BackupVaultsClientCreateOrUpdateResponse struct {
	BackupVaultResource
}

// BackupVaultsClientDeleteResponse contains the response from method BackupVaultsClient.BeginDelete.
type BackupVaultsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupVaultsClientGetInResourceGroupResponse contains the response from method BackupVaultsClient.NewGetInResourceGroupPager.
type BackupVaultsClientGetInResourceGroupResponse struct {
	BackupVaultResourceList
}

// BackupVaultsClientGetInSubscriptionResponse contains the response from method BackupVaultsClient.NewGetInSubscriptionPager.
type BackupVaultsClientGetInSubscriptionResponse struct {
	BackupVaultResourceList
}

// BackupVaultsClientGetResponse contains the response from method BackupVaultsClient.Get.
type BackupVaultsClientGetResponse struct {
	BackupVaultResource
}

// BackupVaultsClientUpdateResponse contains the response from method BackupVaultsClient.BeginUpdate.
type BackupVaultsClientUpdateResponse struct {
	BackupVaultResource
}

// ClientCheckFeatureSupportResponse contains the response from method Client.CheckFeatureSupport.
type ClientCheckFeatureSupportResponse struct {
	FeatureValidationResponseBaseClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientCheckFeatureSupportResponse.
func (c *ClientCheckFeatureSupportResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalFeatureValidationResponseBaseClassification(data)
	if err != nil {
		return err
	}
	c.FeatureValidationResponseBaseClassification = res
	return nil
}

// ClientFetchSecondaryRPsResponse contains the response from method Client.NewFetchSecondaryRPsPager.
type ClientFetchSecondaryRPsResponse struct {
	AzureBackupRecoveryPointResourceList
}

// DeletedBackupInstancesClientGetResponse contains the response from method DeletedBackupInstancesClient.Get.
type DeletedBackupInstancesClientGetResponse struct {
	DeletedBackupInstanceResource
}

// DeletedBackupInstancesClientListResponse contains the response from method DeletedBackupInstancesClient.NewListPager.
type DeletedBackupInstancesClientListResponse struct {
	DeletedBackupInstanceResourceList
}

// DeletedBackupInstancesClientUndeleteResponse contains the response from method DeletedBackupInstancesClient.BeginUndelete.
type DeletedBackupInstancesClientUndeleteResponse struct {
	// placeholder for future response values
}

// DppResourceGuardProxyClientCreateOrUpdateResponse contains the response from method DppResourceGuardProxyClient.CreateOrUpdate.
type DppResourceGuardProxyClientCreateOrUpdateResponse struct {
	ResourceGuardProxyBaseResource
}

// DppResourceGuardProxyClientDeleteResponse contains the response from method DppResourceGuardProxyClient.Delete.
type DppResourceGuardProxyClientDeleteResponse struct {
	// placeholder for future response values
}

// DppResourceGuardProxyClientGetResponse contains the response from method DppResourceGuardProxyClient.Get.
type DppResourceGuardProxyClientGetResponse struct {
	ResourceGuardProxyBaseResource
}

// DppResourceGuardProxyClientListResponse contains the response from method DppResourceGuardProxyClient.NewListPager.
type DppResourceGuardProxyClientListResponse struct {
	ResourceGuardProxyBaseResourceList
}

// DppResourceGuardProxyClientUnlockDeleteResponse contains the response from method DppResourceGuardProxyClient.UnlockDelete.
type DppResourceGuardProxyClientUnlockDeleteResponse struct {
	UnlockDeleteResponse
}

// ExportJobsClientTriggerResponse contains the response from method ExportJobsClient.BeginTrigger.
type ExportJobsClientTriggerResponse struct {
	// placeholder for future response values
}

// ExportJobsOperationResultClientGetResponse contains the response from method ExportJobsOperationResultClient.Get.
type ExportJobsOperationResultClientGetResponse struct {
	ExportJobsResult
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	AzureBackupJobResource
}

// JobsClientListResponse contains the response from method JobsClient.NewListPager.
type JobsClientListResponse struct {
	AzureBackupJobResourceList
}

// JobsClientTriggerCancelResponse contains the response from method JobsClient.BeginTriggerCancel.
type JobsClientTriggerCancelResponse struct {
	// placeholder for future response values
}

// OperationResultClientGetResponse contains the response from method OperationResultClient.Get.
type OperationResultClientGetResponse struct {
	OperationJobExtendedInfo
	// AzureAsyncOperation contains the information returned from the Azure-AsyncOperation header response.
	AzureAsyncOperation *string

	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *int32
}

// OperationStatusBackupVaultContextClientGetResponse contains the response from method OperationStatusBackupVaultContextClient.Get.
type OperationStatusBackupVaultContextClientGetResponse struct {
	OperationResource
}

// OperationStatusClientGetResponse contains the response from method OperationStatusClient.Get.
type OperationStatusClientGetResponse struct {
	OperationResource
}

// OperationStatusResourceGroupContextClientGetResponse contains the response from method OperationStatusResourceGroupContextClient.Get.
type OperationStatusResourceGroupContextClientGetResponse struct {
	OperationResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	ClientDiscoveryResponse
}

// RecoveryPointsClientGetResponse contains the response from method RecoveryPointsClient.Get.
type RecoveryPointsClientGetResponse struct {
	AzureBackupRecoveryPointResource
}

// RecoveryPointsClientListResponse contains the response from method RecoveryPointsClient.NewListPager.
type RecoveryPointsClientListResponse struct {
	AzureBackupRecoveryPointResourceList
}

// ResourceGuardsClientDeleteResponse contains the response from method ResourceGuardsClient.Delete.
type ResourceGuardsClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceGuardsClientGetBackupSecurityPINRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetBackupSecurityPINRequestsObjectsPager.
type ResourceGuardsClientGetBackupSecurityPINRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientGetDefaultBackupSecurityPINRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultBackupSecurityPINRequestsObject.
type ResourceGuardsClientGetDefaultBackupSecurityPINRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDefaultDeleteProtectedItemRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultDeleteProtectedItemRequestsObject.
type ResourceGuardsClientGetDefaultDeleteProtectedItemRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDefaultDeleteResourceGuardProxyRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultDeleteResourceGuardProxyRequestsObject.
type ResourceGuardsClientGetDefaultDeleteResourceGuardProxyRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDefaultDisableSoftDeleteRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultDisableSoftDeleteRequestsObject.
type ResourceGuardsClientGetDefaultDisableSoftDeleteRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDefaultUpdateProtectedItemRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultUpdateProtectedItemRequestsObject.
type ResourceGuardsClientGetDefaultUpdateProtectedItemRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDefaultUpdateProtectionPolicyRequestsObjectResponse contains the response from method ResourceGuardsClient.GetDefaultUpdateProtectionPolicyRequestsObject.
type ResourceGuardsClientGetDefaultUpdateProtectionPolicyRequestsObjectResponse struct {
	DppBaseResource
}

// ResourceGuardsClientGetDeleteProtectedItemRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetDeleteProtectedItemRequestsObjectsPager.
type ResourceGuardsClientGetDeleteProtectedItemRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientGetDeleteResourceGuardProxyRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetDeleteResourceGuardProxyRequestsObjectsPager.
type ResourceGuardsClientGetDeleteResourceGuardProxyRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientGetDisableSoftDeleteRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetDisableSoftDeleteRequestsObjectsPager.
type ResourceGuardsClientGetDisableSoftDeleteRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientGetResourcesInResourceGroupResponse contains the response from method ResourceGuardsClient.NewGetResourcesInResourceGroupPager.
type ResourceGuardsClientGetResourcesInResourceGroupResponse struct {
	ResourceGuardResourceList
}

// ResourceGuardsClientGetResourcesInSubscriptionResponse contains the response from method ResourceGuardsClient.NewGetResourcesInSubscriptionPager.
type ResourceGuardsClientGetResourcesInSubscriptionResponse struct {
	ResourceGuardResourceList
}

// ResourceGuardsClientGetResponse contains the response from method ResourceGuardsClient.Get.
type ResourceGuardsClientGetResponse struct {
	ResourceGuardResource
}

// ResourceGuardsClientGetUpdateProtectedItemRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetUpdateProtectedItemRequestsObjectsPager.
type ResourceGuardsClientGetUpdateProtectedItemRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientGetUpdateProtectionPolicyRequestsObjectsResponse contains the response from method ResourceGuardsClient.NewGetUpdateProtectionPolicyRequestsObjectsPager.
type ResourceGuardsClientGetUpdateProtectionPolicyRequestsObjectsResponse struct {
	DppBaseResourceList
}

// ResourceGuardsClientPatchResponse contains the response from method ResourceGuardsClient.Patch.
type ResourceGuardsClientPatchResponse struct {
	ResourceGuardResource
}

// ResourceGuardsClientPutResponse contains the response from method ResourceGuardsClient.Put.
type ResourceGuardsClientPutResponse struct {
	ResourceGuardResource
}

// RestorableTimeRangesClientFindResponse contains the response from method RestorableTimeRangesClient.Find.
type RestorableTimeRangesClientFindResponse struct {
	AzureBackupFindRestorableTimeRangesResponseResource
}
