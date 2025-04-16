// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

// AccessControlRecordsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccessControlRecordsClient.BeginCreateOrUpdate
// method.
type AccessControlRecordsClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AccessControlRecordsClientBeginDeleteOptions contains the optional parameters for the AccessControlRecordsClient.BeginDelete
// method.
type AccessControlRecordsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// AccessControlRecordsClientGetOptions contains the optional parameters for the AccessControlRecordsClient.Get method.
type AccessControlRecordsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccessControlRecordsClientListByManagerOptions contains the optional parameters for the AccessControlRecordsClient.NewListByManagerPager
// method.
type AccessControlRecordsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientClearOptions contains the optional parameters for the AlertsClient.Clear method.
type AlertsClientClearOptions struct {
	// placeholder for future optional parameters
}

// AlertsClientListByManagerOptions contains the optional parameters for the AlertsClient.NewListByManagerPager method.
type AlertsClientListByManagerOptions struct {
	// OData Filter options
	Filter *string
}

// AlertsClientSendTestEmailOptions contains the optional parameters for the AlertsClient.SendTestEmail method.
type AlertsClientSendTestEmailOptions struct {
	// placeholder for future optional parameters
}

// AvailableProviderOperationsClientListOptions contains the optional parameters for the AvailableProviderOperationsClient.NewListPager
// method.
type AvailableProviderOperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// BackupScheduleGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupScheduleGroupsClient.BeginCreateOrUpdate
// method.
type BackupScheduleGroupsClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BackupScheduleGroupsClientBeginDeleteOptions contains the optional parameters for the BackupScheduleGroupsClient.BeginDelete
// method.
type BackupScheduleGroupsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BackupScheduleGroupsClientGetOptions contains the optional parameters for the BackupScheduleGroupsClient.Get method.
type BackupScheduleGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupScheduleGroupsClientListByDeviceOptions contains the optional parameters for the BackupScheduleGroupsClient.NewListByDevicePager
// method.
type BackupScheduleGroupsClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientBeginCloneOptions contains the optional parameters for the BackupsClient.BeginClone method.
type BackupsClientBeginCloneOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
type BackupsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// BackupsClientListByDeviceOptions contains the optional parameters for the BackupsClient.NewListByDevicePager method.
type BackupsClientListByDeviceOptions struct {
	// OData Filter options
	Filter *string

	// Set to true if you need backups which can be used for failover.
	ForFailover *bool
}

// BackupsClientListByManagerOptions contains the optional parameters for the BackupsClient.NewListByManagerPager method.
type BackupsClientListByManagerOptions struct {
	// OData Filter options
	Filter *string
}

// ChapSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ChapSettingsClient.BeginCreateOrUpdate
// method.
type ChapSettingsClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ChapSettingsClientBeginDeleteOptions contains the optional parameters for the ChapSettingsClient.BeginDelete method.
type ChapSettingsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ChapSettingsClientGetOptions contains the optional parameters for the ChapSettingsClient.Get method.
type ChapSettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ChapSettingsClientListByDeviceOptions contains the optional parameters for the ChapSettingsClient.NewListByDevicePager
// method.
type ChapSettingsClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientBeginCreateOrUpdateAlertSettingsOptions contains the optional parameters for the DevicesClient.BeginCreateOrUpdateAlertSettings
// method.
type DevicesClientBeginCreateOrUpdateAlertSettingsOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginCreateOrUpdateSecuritySettingsOptions contains the optional parameters for the DevicesClient.BeginCreateOrUpdateSecuritySettings
// method.
type DevicesClientBeginCreateOrUpdateSecuritySettingsOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginDeactivateOptions contains the optional parameters for the DevicesClient.BeginDeactivate method.
type DevicesClientBeginDeactivateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginDeleteOptions contains the optional parameters for the DevicesClient.BeginDelete method.
type DevicesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginDownloadUpdatesOptions contains the optional parameters for the DevicesClient.BeginDownloadUpdates method.
type DevicesClientBeginDownloadUpdatesOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginFailoverOptions contains the optional parameters for the DevicesClient.BeginFailover method.
type DevicesClientBeginFailoverOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginInstallUpdatesOptions contains the optional parameters for the DevicesClient.BeginInstallUpdates method.
type DevicesClientBeginInstallUpdatesOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginPatchOptions contains the optional parameters for the DevicesClient.BeginPatch method.
type DevicesClientBeginPatchOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientBeginScanForUpdatesOptions contains the optional parameters for the DevicesClient.BeginScanForUpdates method.
type DevicesClientBeginScanForUpdatesOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// DevicesClientGetAlertSettingsOptions contains the optional parameters for the DevicesClient.GetAlertSettings method.
type DevicesClientGetAlertSettingsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientGetNetworkSettingsOptions contains the optional parameters for the DevicesClient.GetNetworkSettings method.
type DevicesClientGetNetworkSettingsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
type DevicesClientGetOptions struct {
	// Specify $expand=details to populate additional fields related to the device.
	Expand *string
}

// DevicesClientGetTimeSettingsOptions contains the optional parameters for the DevicesClient.GetTimeSettings method.
type DevicesClientGetTimeSettingsOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientGetUpdateSummaryOptions contains the optional parameters for the DevicesClient.GetUpdateSummary method.
type DevicesClientGetUpdateSummaryOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListByManagerOptions contains the optional parameters for the DevicesClient.NewListByManagerPager method.
type DevicesClientListByManagerOptions struct {
	// Specify $expand=details to populate additional fields related to the device.
	Expand *string
}

// DevicesClientListFailoverTargetOptions contains the optional parameters for the DevicesClient.NewListFailoverTargetPager
// method.
type DevicesClientListFailoverTargetOptions struct {
	// Specify $expand=details to populate additional fields related to the device.
	Expand *string
}

// DevicesClientListMetricDefinitionOptions contains the optional parameters for the DevicesClient.NewListMetricDefinitionPager
// method.
type DevicesClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// DevicesClientListMetricsOptions contains the optional parameters for the DevicesClient.NewListMetricsPager method.
type DevicesClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// FileServersClientBeginBackupNowOptions contains the optional parameters for the FileServersClient.BeginBackupNow method.
type FileServersClientBeginBackupNowOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// FileServersClientBeginCreateOrUpdateOptions contains the optional parameters for the FileServersClient.BeginCreateOrUpdate
// method.
type FileServersClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// FileServersClientBeginDeleteOptions contains the optional parameters for the FileServersClient.BeginDelete method.
type FileServersClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// FileServersClientGetOptions contains the optional parameters for the FileServersClient.Get method.
type FileServersClientGetOptions struct {
	// placeholder for future optional parameters
}

// FileServersClientListByDeviceOptions contains the optional parameters for the FileServersClient.NewListByDevicePager method.
type FileServersClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// FileServersClientListByManagerOptions contains the optional parameters for the FileServersClient.NewListByManagerPager
// method.
type FileServersClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// FileServersClientListMetricDefinitionOptions contains the optional parameters for the FileServersClient.NewListMetricDefinitionPager
// method.
type FileServersClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// FileServersClientListMetricsOptions contains the optional parameters for the FileServersClient.NewListMetricsPager method.
type FileServersClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// FileSharesClientBeginCreateOrUpdateOptions contains the optional parameters for the FileSharesClient.BeginCreateOrUpdate
// method.
type FileSharesClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// FileSharesClientBeginDeleteOptions contains the optional parameters for the FileSharesClient.BeginDelete method.
type FileSharesClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// FileSharesClientGetOptions contains the optional parameters for the FileSharesClient.Get method.
type FileSharesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FileSharesClientListByDeviceOptions contains the optional parameters for the FileSharesClient.NewListByDevicePager method.
type FileSharesClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// FileSharesClientListByFileServerOptions contains the optional parameters for the FileSharesClient.NewListByFileServerPager
// method.
type FileSharesClientListByFileServerOptions struct {
	// placeholder for future optional parameters
}

// FileSharesClientListMetricDefinitionOptions contains the optional parameters for the FileSharesClient.NewListMetricDefinitionPager
// method.
type FileSharesClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// FileSharesClientListMetricsOptions contains the optional parameters for the FileSharesClient.NewListMetricsPager method.
type FileSharesClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// IscsiDisksClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiDisksClient.BeginCreateOrUpdate
// method.
type IscsiDisksClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// IscsiDisksClientBeginDeleteOptions contains the optional parameters for the IscsiDisksClient.BeginDelete method.
type IscsiDisksClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// IscsiDisksClientGetOptions contains the optional parameters for the IscsiDisksClient.Get method.
type IscsiDisksClientGetOptions struct {
	// placeholder for future optional parameters
}

// IscsiDisksClientListByDeviceOptions contains the optional parameters for the IscsiDisksClient.NewListByDevicePager method.
type IscsiDisksClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// IscsiDisksClientListByIscsiServerOptions contains the optional parameters for the IscsiDisksClient.NewListByIscsiServerPager
// method.
type IscsiDisksClientListByIscsiServerOptions struct {
	// placeholder for future optional parameters
}

// IscsiDisksClientListMetricDefinitionOptions contains the optional parameters for the IscsiDisksClient.NewListMetricDefinitionPager
// method.
type IscsiDisksClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// IscsiDisksClientListMetricsOptions contains the optional parameters for the IscsiDisksClient.NewListMetricsPager method.
type IscsiDisksClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// IscsiServersClientBeginBackupNowOptions contains the optional parameters for the IscsiServersClient.BeginBackupNow method.
type IscsiServersClientBeginBackupNowOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// IscsiServersClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiServersClient.BeginCreateOrUpdate
// method.
type IscsiServersClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// IscsiServersClientBeginDeleteOptions contains the optional parameters for the IscsiServersClient.BeginDelete method.
type IscsiServersClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// IscsiServersClientGetOptions contains the optional parameters for the IscsiServersClient.Get method.
type IscsiServersClientGetOptions struct {
	// placeholder for future optional parameters
}

// IscsiServersClientListByDeviceOptions contains the optional parameters for the IscsiServersClient.NewListByDevicePager
// method.
type IscsiServersClientListByDeviceOptions struct {
	// placeholder for future optional parameters
}

// IscsiServersClientListByManagerOptions contains the optional parameters for the IscsiServersClient.NewListByManagerPager
// method.
type IscsiServersClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// IscsiServersClientListMetricDefinitionOptions contains the optional parameters for the IscsiServersClient.NewListMetricDefinitionPager
// method.
type IscsiServersClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// IscsiServersClientListMetricsOptions contains the optional parameters for the IscsiServersClient.NewListMetricsPager method.
type IscsiServersClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
type JobsClientGetOptions struct {
	// placeholder for future optional parameters
}

// JobsClientListByDeviceOptions contains the optional parameters for the JobsClient.NewListByDevicePager method.
type JobsClientListByDeviceOptions struct {
	// OData Filter options
	Filter *string
}

// JobsClientListByManagerOptions contains the optional parameters for the JobsClient.NewListByManagerPager method.
type JobsClientListByManagerOptions struct {
	// OData Filter options
	Filter *string
}

// ManagersClientCreateExtendedInfoOptions contains the optional parameters for the ManagersClient.CreateExtendedInfo method.
type ManagersClientCreateExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientCreateOrUpdateOptions contains the optional parameters for the ManagersClient.CreateOrUpdate method.
type ManagersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientDeleteExtendedInfoOptions contains the optional parameters for the ManagersClient.DeleteExtendedInfo method.
type ManagersClientDeleteExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientDeleteOptions contains the optional parameters for the ManagersClient.Delete method.
type ManagersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetEncryptionKeyOptions contains the optional parameters for the ManagersClient.GetEncryptionKey method.
type ManagersClientGetEncryptionKeyOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetEncryptionSettingsOptions contains the optional parameters for the ManagersClient.GetEncryptionSettings
// method.
type ManagersClientGetEncryptionSettingsOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetExtendedInfoOptions contains the optional parameters for the ManagersClient.GetExtendedInfo method.
type ManagersClientGetExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientGetOptions contains the optional parameters for the ManagersClient.Get method.
type ManagersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListByResourceGroupOptions contains the optional parameters for the ManagersClient.NewListByResourceGroupPager
// method.
type ManagersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListMetricDefinitionOptions contains the optional parameters for the ManagersClient.NewListMetricDefinitionPager
// method.
type ManagersClientListMetricDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientListMetricsOptions contains the optional parameters for the ManagersClient.NewListMetricsPager method.
type ManagersClientListMetricsOptions struct {
	// OData Filter options
	Filter *string
}

// ManagersClientListOptions contains the optional parameters for the ManagersClient.NewListPager method.
type ManagersClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientUpdateExtendedInfoOptions contains the optional parameters for the ManagersClient.UpdateExtendedInfo method.
type ManagersClientUpdateExtendedInfoOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientUpdateOptions contains the optional parameters for the ManagersClient.Update method.
type ManagersClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ManagersClientUploadRegistrationCertificateOptions contains the optional parameters for the ManagersClient.UploadRegistrationCertificate
// method.
type ManagersClientUploadRegistrationCertificateOptions struct {
	// placeholder for future optional parameters
}

// StorageAccountCredentialsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginCreateOrUpdate
// method.
type StorageAccountCredentialsClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// StorageAccountCredentialsClientBeginDeleteOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginDelete
// method.
type StorageAccountCredentialsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// StorageAccountCredentialsClientGetOptions contains the optional parameters for the StorageAccountCredentialsClient.Get
// method.
type StorageAccountCredentialsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageAccountCredentialsClientListByManagerOptions contains the optional parameters for the StorageAccountCredentialsClient.NewListByManagerPager
// method.
type StorageAccountCredentialsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}

// StorageDomainsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageDomainsClient.BeginCreateOrUpdate
// method.
type StorageDomainsClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// StorageDomainsClientBeginDeleteOptions contains the optional parameters for the StorageDomainsClient.BeginDelete method.
type StorageDomainsClientBeginDeleteOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// StorageDomainsClientGetOptions contains the optional parameters for the StorageDomainsClient.Get method.
type StorageDomainsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageDomainsClientListByManagerOptions contains the optional parameters for the StorageDomainsClient.NewListByManagerPager
// method.
type StorageDomainsClientListByManagerOptions struct {
	// placeholder for future optional parameters
}
