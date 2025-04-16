// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

// AcceleratorAuthSettingClassification provides polymorphic access to related types.
// Call the interface's GetAcceleratorAuthSetting() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AcceleratorAuthSetting, *AcceleratorBasicAuthSetting, *AcceleratorPublicSetting, *AcceleratorSSHSetting
type AcceleratorAuthSettingClassification interface {
	// GetAcceleratorAuthSetting returns the AcceleratorAuthSetting content of the underlying type.
	GetAcceleratorAuthSetting() *AcceleratorAuthSetting
}

// CertificatePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetCertificateProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CertificateProperties, *ContentCertificateProperties, *KeyVaultCertificateProperties
type CertificatePropertiesClassification interface {
	// GetCertificateProperties returns the CertificateProperties content of the underlying type.
	GetCertificateProperties() *CertificateProperties
}

// ContainerRegistryCredentialsClassification provides polymorphic access to related types.
// Call the interface's GetContainerRegistryCredentials() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ContainerRegistryBasicCredentials, *ContainerRegistryCredentials
type ContainerRegistryCredentialsClassification interface {
	// GetContainerRegistryCredentials returns the ContainerRegistryCredentials content of the underlying type.
	GetContainerRegistryCredentials() *ContainerRegistryCredentials
}

// CustomPersistentDiskPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetCustomPersistentDiskProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureFileVolume, *CustomPersistentDiskProperties
type CustomPersistentDiskPropertiesClassification interface {
	// GetCustomPersistentDiskProperties returns the CustomPersistentDiskProperties content of the underlying type.
	GetCustomPersistentDiskProperties() *CustomPersistentDiskProperties
}

// GatewayResponseCachePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetGatewayResponseCacheProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *GatewayLocalResponseCachePerInstanceProperties, *GatewayLocalResponseCachePerRouteProperties, *GatewayResponseCacheProperties
type GatewayResponseCachePropertiesClassification interface {
	// GetGatewayResponseCacheProperties returns the GatewayResponseCacheProperties content of the underlying type.
	GetGatewayResponseCacheProperties() *GatewayResponseCacheProperties
}

// JobTriggerConfigClassification provides polymorphic access to related types.
// Call the interface's GetJobTriggerConfig() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *JobTriggerConfig, *ManualJobTriggerConfig
type JobTriggerConfigClassification interface {
	// GetJobTriggerConfig returns the JobTriggerConfig content of the underlying type.
	GetJobTriggerConfig() *JobTriggerConfig
}

// MaintenanceScheduleConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetMaintenanceScheduleConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MaintenanceScheduleConfiguration, *WeeklyMaintenanceScheduleConfiguration
type MaintenanceScheduleConfigurationClassification interface {
	// GetMaintenanceScheduleConfiguration returns the MaintenanceScheduleConfiguration content of the underlying type.
	GetMaintenanceScheduleConfiguration() *MaintenanceScheduleConfiguration
}

// ProbeActionClassification provides polymorphic access to related types.
// Call the interface's GetProbeAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ExecAction, *HTTPGetAction, *ProbeAction, *TCPSocketAction
type ProbeActionClassification interface {
	// GetProbeAction returns the ProbeAction content of the underlying type.
	GetProbeAction() *ProbeAction
}

// StoragePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetStorageProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *StorageAccount, *StorageProperties
type StoragePropertiesClassification interface {
	// GetStorageProperties returns the StorageProperties content of the underlying type.
	GetStorageProperties() *StorageProperties
}

// UploadedUserSourceInfoClassification provides polymorphic access to related types.
// Call the interface's GetUploadedUserSourceInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *JarUploadedUserSourceInfo, *NetCoreZipUploadedUserSourceInfo, *SourceUploadedUserSourceInfo, *UploadedUserSourceInfo,
// - *WarUploadedUserSourceInfo
type UploadedUserSourceInfoClassification interface {
	UserSourceInfoClassification
	// GetUploadedUserSourceInfo returns the UploadedUserSourceInfo content of the underlying type.
	GetUploadedUserSourceInfo() *UploadedUserSourceInfo
}

// UserSourceInfoClassification provides polymorphic access to related types.
// Call the interface's GetUserSourceInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BuildResultUserSourceInfo, *CustomContainerUserSourceInfo, *JarUploadedUserSourceInfo, *NetCoreZipUploadedUserSourceInfo,
// - *SourceUploadedUserSourceInfo, *UploadedUserSourceInfo, *UserSourceInfo, *WarUploadedUserSourceInfo
type UserSourceInfoClassification interface {
	// GetUserSourceInfo returns the UserSourceInfo content of the underlying type.
	GetUserSourceInfo() *UserSourceInfo
}
