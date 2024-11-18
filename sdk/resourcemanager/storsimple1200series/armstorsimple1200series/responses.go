//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

// AccessControlRecordsClientCreateOrUpdateResponse contains the response from method AccessControlRecordsClient.BeginCreateOrUpdate.
type AccessControlRecordsClientCreateOrUpdateResponse struct {
	// The access control record
	AccessControlRecord
}

// AccessControlRecordsClientDeleteResponse contains the response from method AccessControlRecordsClient.BeginDelete.
type AccessControlRecordsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccessControlRecordsClientGetResponse contains the response from method AccessControlRecordsClient.Get.
type AccessControlRecordsClientGetResponse struct {
	// The access control record
	AccessControlRecord
}

// AccessControlRecordsClientListByManagerResponse contains the response from method AccessControlRecordsClient.NewListByManagerPager.
type AccessControlRecordsClientListByManagerResponse struct {
	// Collection of AccessControlRecords
	AccessControlRecordList
}

// AlertsClientClearResponse contains the response from method AlertsClient.Clear.
type AlertsClientClearResponse struct {
	// placeholder for future response values
}

// AlertsClientListByManagerResponse contains the response from method AlertsClient.NewListByManagerPager.
type AlertsClientListByManagerResponse struct {
	// Collection of Alerts
	AlertList
}

// AlertsClientSendTestEmailResponse contains the response from method AlertsClient.SendTestEmail.
type AlertsClientSendTestEmailResponse struct {
	// placeholder for future response values
}

// AvailableProviderOperationsClientListResponse contains the response from method AvailableProviderOperationsClient.NewListPager.
type AvailableProviderOperationsClientListResponse struct {
	// Class for set of operations used for discovery of available provider operations.
	AvailableProviderOperations
}

// BackupScheduleGroupsClientCreateOrUpdateResponse contains the response from method BackupScheduleGroupsClient.BeginCreateOrUpdate.
type BackupScheduleGroupsClientCreateOrUpdateResponse struct {
	// The Backup Schedule Group
	BackupScheduleGroup
}

// BackupScheduleGroupsClientDeleteResponse contains the response from method BackupScheduleGroupsClient.BeginDelete.
type BackupScheduleGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupScheduleGroupsClientGetResponse contains the response from method BackupScheduleGroupsClient.Get.
type BackupScheduleGroupsClientGetResponse struct {
	// The Backup Schedule Group
	BackupScheduleGroup
}

// BackupScheduleGroupsClientListByDeviceResponse contains the response from method BackupScheduleGroupsClient.NewListByDevicePager.
type BackupScheduleGroupsClientListByDeviceResponse struct {
	// The list response of backup schedule groups
	BackupScheduleGroupList
}

// BackupsClientCloneResponse contains the response from method BackupsClient.BeginClone.
type BackupsClientCloneResponse struct {
	// placeholder for future response values
}

// BackupsClientDeleteResponse contains the response from method BackupsClient.BeginDelete.
type BackupsClientDeleteResponse struct {
	// placeholder for future response values
}

// BackupsClientListByDeviceResponse contains the response from method BackupsClient.NewListByDevicePager.
type BackupsClientListByDeviceResponse struct {
	// Collection of backups
	BackupList
}

// BackupsClientListByManagerResponse contains the response from method BackupsClient.NewListByManagerPager.
type BackupsClientListByManagerResponse struct {
	// Collection of backups
	BackupList
}

// ChapSettingsClientCreateOrUpdateResponse contains the response from method ChapSettingsClient.BeginCreateOrUpdate.
type ChapSettingsClientCreateOrUpdateResponse struct {
	// Challenge-Handshake Authentication Protocol (CHAP) setting
	ChapSettings
}

// ChapSettingsClientDeleteResponse contains the response from method ChapSettingsClient.BeginDelete.
type ChapSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ChapSettingsClientGetResponse contains the response from method ChapSettingsClient.Get.
type ChapSettingsClientGetResponse struct {
	// Challenge-Handshake Authentication Protocol (CHAP) setting
	ChapSettings
}

// ChapSettingsClientListByDeviceResponse contains the response from method ChapSettingsClient.NewListByDevicePager.
type ChapSettingsClientListByDeviceResponse struct {
	// Collection of Chap setting entities
	ChapSettingsList
}

// DevicesClientCreateOrUpdateAlertSettingsResponse contains the response from method DevicesClient.BeginCreateOrUpdateAlertSettings.
type DevicesClientCreateOrUpdateAlertSettingsResponse struct {
	// AlertSettings on the device which represents how alerts will be processed
	AlertSettings
}

// DevicesClientCreateOrUpdateSecuritySettingsResponse contains the response from method DevicesClient.BeginCreateOrUpdateSecuritySettings.
type DevicesClientCreateOrUpdateSecuritySettingsResponse struct {
	// placeholder for future response values
}

// DevicesClientDeactivateResponse contains the response from method DevicesClient.BeginDeactivate.
type DevicesClientDeactivateResponse struct {
	// placeholder for future response values
}

// DevicesClientDeleteResponse contains the response from method DevicesClient.BeginDelete.
type DevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DevicesClientDownloadUpdatesResponse contains the response from method DevicesClient.BeginDownloadUpdates.
type DevicesClientDownloadUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientFailoverResponse contains the response from method DevicesClient.BeginFailover.
type DevicesClientFailoverResponse struct {
	// placeholder for future response values
}

// DevicesClientGetAlertSettingsResponse contains the response from method DevicesClient.GetAlertSettings.
type DevicesClientGetAlertSettingsResponse struct {
	// AlertSettings on the device which represents how alerts will be processed
	AlertSettings
}

// DevicesClientGetNetworkSettingsResponse contains the response from method DevicesClient.GetNetworkSettings.
type DevicesClientGetNetworkSettingsResponse struct {
	// The NetworkSettings of a device
	NetworkSettings
}

// DevicesClientGetResponse contains the response from method DevicesClient.Get.
type DevicesClientGetResponse struct {
	// Represents a StorSimple device object along with its properties
	Device
}

// DevicesClientGetTimeSettingsResponse contains the response from method DevicesClient.GetTimeSettings.
type DevicesClientGetTimeSettingsResponse struct {
	// The TimeSettings of a device
	TimeSettings
}

// DevicesClientGetUpdateSummaryResponse contains the response from method DevicesClient.GetUpdateSummary.
type DevicesClientGetUpdateSummaryResponse struct {
	// The updates profile
	Updates
}

// DevicesClientInstallUpdatesResponse contains the response from method DevicesClient.BeginInstallUpdates.
type DevicesClientInstallUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientListByManagerResponse contains the response from method DevicesClient.NewListByManagerPager.
type DevicesClientListByManagerResponse struct {
	// Collection of Devices
	DeviceList
}

// DevicesClientListFailoverTargetResponse contains the response from method DevicesClient.NewListFailoverTargetPager.
type DevicesClientListFailoverTargetResponse struct {
	// Collection of Devices
	DeviceList
}

// DevicesClientListMetricDefinitionResponse contains the response from method DevicesClient.NewListMetricDefinitionPager.
type DevicesClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// DevicesClientListMetricsResponse contains the response from method DevicesClient.NewListMetricsPager.
type DevicesClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// DevicesClientPatchResponse contains the response from method DevicesClient.BeginPatch.
type DevicesClientPatchResponse struct {
	// Represents a StorSimple device object along with its properties
	Device
}

// DevicesClientScanForUpdatesResponse contains the response from method DevicesClient.BeginScanForUpdates.
type DevicesClientScanForUpdatesResponse struct {
	// placeholder for future response values
}

// FileServersClientBackupNowResponse contains the response from method FileServersClient.BeginBackupNow.
type FileServersClientBackupNowResponse struct {
	// placeholder for future response values
}

// FileServersClientCreateOrUpdateResponse contains the response from method FileServersClient.BeginCreateOrUpdate.
type FileServersClientCreateOrUpdateResponse struct {
	// The file server.
	FileServer
}

// FileServersClientDeleteResponse contains the response from method FileServersClient.BeginDelete.
type FileServersClientDeleteResponse struct {
	// placeholder for future response values
}

// FileServersClientGetResponse contains the response from method FileServersClient.Get.
type FileServersClientGetResponse struct {
	// The file server.
	FileServer
}

// FileServersClientListByDeviceResponse contains the response from method FileServersClient.NewListByDevicePager.
type FileServersClientListByDeviceResponse struct {
	// Collection of file servers
	FileServerList
}

// FileServersClientListByManagerResponse contains the response from method FileServersClient.NewListByManagerPager.
type FileServersClientListByManagerResponse struct {
	// Collection of file servers
	FileServerList
}

// FileServersClientListMetricDefinitionResponse contains the response from method FileServersClient.NewListMetricDefinitionPager.
type FileServersClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// FileServersClientListMetricsResponse contains the response from method FileServersClient.NewListMetricsPager.
type FileServersClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// FileSharesClientCreateOrUpdateResponse contains the response from method FileSharesClient.BeginCreateOrUpdate.
type FileSharesClientCreateOrUpdateResponse struct {
	// The File Share.
	FileShare
}

// FileSharesClientDeleteResponse contains the response from method FileSharesClient.BeginDelete.
type FileSharesClientDeleteResponse struct {
	// placeholder for future response values
}

// FileSharesClientGetResponse contains the response from method FileSharesClient.Get.
type FileSharesClientGetResponse struct {
	// The File Share.
	FileShare
}

// FileSharesClientListByDeviceResponse contains the response from method FileSharesClient.NewListByDevicePager.
type FileSharesClientListByDeviceResponse struct {
	// Collection of file shares
	FileShareList
}

// FileSharesClientListByFileServerResponse contains the response from method FileSharesClient.NewListByFileServerPager.
type FileSharesClientListByFileServerResponse struct {
	// Collection of file shares
	FileShareList
}

// FileSharesClientListMetricDefinitionResponse contains the response from method FileSharesClient.NewListMetricDefinitionPager.
type FileSharesClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// FileSharesClientListMetricsResponse contains the response from method FileSharesClient.NewListMetricsPager.
type FileSharesClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// IscsiDisksClientCreateOrUpdateResponse contains the response from method IscsiDisksClient.BeginCreateOrUpdate.
type IscsiDisksClientCreateOrUpdateResponse struct {
	// The iSCSI disk.
	ISCSIDisk
}

// IscsiDisksClientDeleteResponse contains the response from method IscsiDisksClient.BeginDelete.
type IscsiDisksClientDeleteResponse struct {
	// placeholder for future response values
}

// IscsiDisksClientGetResponse contains the response from method IscsiDisksClient.Get.
type IscsiDisksClientGetResponse struct {
	// The iSCSI disk.
	ISCSIDisk
}

// IscsiDisksClientListByDeviceResponse contains the response from method IscsiDisksClient.NewListByDevicePager.
type IscsiDisksClientListByDeviceResponse struct {
	// Collection of Iscsi disk
	ISCSIDiskList
}

// IscsiDisksClientListByIscsiServerResponse contains the response from method IscsiDisksClient.NewListByIscsiServerPager.
type IscsiDisksClientListByIscsiServerResponse struct {
	// Collection of Iscsi disk
	ISCSIDiskList
}

// IscsiDisksClientListMetricDefinitionResponse contains the response from method IscsiDisksClient.NewListMetricDefinitionPager.
type IscsiDisksClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// IscsiDisksClientListMetricsResponse contains the response from method IscsiDisksClient.NewListMetricsPager.
type IscsiDisksClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// IscsiServersClientBackupNowResponse contains the response from method IscsiServersClient.BeginBackupNow.
type IscsiServersClientBackupNowResponse struct {
	// placeholder for future response values
}

// IscsiServersClientCreateOrUpdateResponse contains the response from method IscsiServersClient.BeginCreateOrUpdate.
type IscsiServersClientCreateOrUpdateResponse struct {
	// The iSCSI server.
	ISCSIServer
}

// IscsiServersClientDeleteResponse contains the response from method IscsiServersClient.BeginDelete.
type IscsiServersClientDeleteResponse struct {
	// placeholder for future response values
}

// IscsiServersClientGetResponse contains the response from method IscsiServersClient.Get.
type IscsiServersClientGetResponse struct {
	// The iSCSI server.
	ISCSIServer
}

// IscsiServersClientListByDeviceResponse contains the response from method IscsiServersClient.NewListByDevicePager.
type IscsiServersClientListByDeviceResponse struct {
	// Collection of Iscsi servers
	ISCSIServerList
}

// IscsiServersClientListByManagerResponse contains the response from method IscsiServersClient.NewListByManagerPager.
type IscsiServersClientListByManagerResponse struct {
	// Collection of Iscsi servers
	ISCSIServerList
}

// IscsiServersClientListMetricDefinitionResponse contains the response from method IscsiServersClient.NewListMetricDefinitionPager.
type IscsiServersClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// IscsiServersClientListMetricsResponse contains the response from method IscsiServersClient.NewListMetricsPager.
type IscsiServersClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	// The Job.
	Job
}

// JobsClientListByDeviceResponse contains the response from method JobsClient.NewListByDevicePager.
type JobsClientListByDeviceResponse struct {
	// Collection of jobs
	JobList
}

// JobsClientListByManagerResponse contains the response from method JobsClient.NewListByManagerPager.
type JobsClientListByManagerResponse struct {
	// Collection of jobs
	JobList
}

// ManagersClientCreateExtendedInfoResponse contains the response from method ManagersClient.CreateExtendedInfo.
type ManagersClientCreateExtendedInfoResponse struct {
	// The extended info of the manager.
	ManagerExtendedInfo
}

// ManagersClientCreateOrUpdateResponse contains the response from method ManagersClient.CreateOrUpdate.
type ManagersClientCreateOrUpdateResponse struct {
	// The StorSimple Manager
	Manager
}

// ManagersClientDeleteExtendedInfoResponse contains the response from method ManagersClient.DeleteExtendedInfo.
type ManagersClientDeleteExtendedInfoResponse struct {
	// placeholder for future response values
}

// ManagersClientDeleteResponse contains the response from method ManagersClient.Delete.
type ManagersClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagersClientGetEncryptionKeyResponse contains the response from method ManagersClient.GetEncryptionKey.
type ManagersClientGetEncryptionKeyResponse struct {
	// This class can be used as the Type for any secret entity represented as Value, ValueCertificateThumbprint, EncryptionAlgorithm.
	// In this case, "Value" is a secret and the "valueThumbprint" represents the certificate thumbprint of the value. The algorithm
	// field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
	SymmetricEncryptedSecret
}

// ManagersClientGetEncryptionSettingsResponse contains the response from method ManagersClient.GetEncryptionSettings.
type ManagersClientGetEncryptionSettingsResponse struct {
	// The EncryptionSettings
	EncryptionSettings
}

// ManagersClientGetExtendedInfoResponse contains the response from method ManagersClient.GetExtendedInfo.
type ManagersClientGetExtendedInfoResponse struct {
	// The extended info of the manager.
	ManagerExtendedInfo
}

// ManagersClientGetResponse contains the response from method ManagersClient.Get.
type ManagersClientGetResponse struct {
	// The StorSimple Manager
	Manager
}

// ManagersClientListByResourceGroupResponse contains the response from method ManagersClient.NewListByResourceGroupPager.
type ManagersClientListByResourceGroupResponse struct {
	// List of StorSimple Managers under a particular resourceGroup
	ManagerList
}

// ManagersClientListMetricDefinitionResponse contains the response from method ManagersClient.NewListMetricDefinitionPager.
type ManagersClientListMetricDefinitionResponse struct {
	// List of metric definition
	MetricDefinitionList
}

// ManagersClientListMetricsResponse contains the response from method ManagersClient.NewListMetricsPager.
type ManagersClientListMetricsResponse struct {
	// Collection of metrics
	MetricList
}

// ManagersClientListResponse contains the response from method ManagersClient.NewListPager.
type ManagersClientListResponse struct {
	// List of StorSimple Managers under a particular resourceGroup
	ManagerList
}

// ManagersClientUpdateExtendedInfoResponse contains the response from method ManagersClient.UpdateExtendedInfo.
type ManagersClientUpdateExtendedInfoResponse struct {
	// The extended info of the manager.
	ManagerExtendedInfo
}

// ManagersClientUpdateResponse contains the response from method ManagersClient.Update.
type ManagersClientUpdateResponse struct {
	// The StorSimple Manager
	Manager
}

// ManagersClientUploadRegistrationCertificateResponse contains the response from method ManagersClient.UploadRegistrationCertificate.
type ManagersClientUploadRegistrationCertificateResponse struct {
	// Upload Certificate Response from IDM
	UploadCertificateResponse
}

// StorageAccountCredentialsClientCreateOrUpdateResponse contains the response from method StorageAccountCredentialsClient.BeginCreateOrUpdate.
type StorageAccountCredentialsClientCreateOrUpdateResponse struct {
	// The storage account credential
	StorageAccountCredential
}

// StorageAccountCredentialsClientDeleteResponse contains the response from method StorageAccountCredentialsClient.BeginDelete.
type StorageAccountCredentialsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountCredentialsClientGetResponse contains the response from method StorageAccountCredentialsClient.Get.
type StorageAccountCredentialsClientGetResponse struct {
	// The storage account credential
	StorageAccountCredential
}

// StorageAccountCredentialsClientListByManagerResponse contains the response from method StorageAccountCredentialsClient.NewListByManagerPager.
type StorageAccountCredentialsClientListByManagerResponse struct {
	// Collection of Storage account credential entities
	StorageAccountCredentialList
}

// StorageDomainsClientCreateOrUpdateResponse contains the response from method StorageDomainsClient.BeginCreateOrUpdate.
type StorageDomainsClientCreateOrUpdateResponse struct {
	// The storage domain.
	StorageDomain
}

// StorageDomainsClientDeleteResponse contains the response from method StorageDomainsClient.BeginDelete.
type StorageDomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageDomainsClientGetResponse contains the response from method StorageDomainsClient.Get.
type StorageDomainsClientGetResponse struct {
	// The storage domain.
	StorageDomain
}

// StorageDomainsClientListByManagerResponse contains the response from method StorageDomainsClient.NewListByManagerPager.
type StorageDomainsClientListByManagerResponse struct {
	// Collection of storage domains
	StorageDomainList
}
