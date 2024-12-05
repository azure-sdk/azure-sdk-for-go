// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armworkloadssapvirtualinstance

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
	moduleVersion = "v1.0.0"
)

// ApplicationServerVirtualMachineType - Defines the type of application server VM.
type ApplicationServerVirtualMachineType string

const (
	// ApplicationServerVirtualMachineTypeActive - Active Application server vm type.
	ApplicationServerVirtualMachineTypeActive ApplicationServerVirtualMachineType = "Active"
	// ApplicationServerVirtualMachineTypeStandby - Standby Application server vm type.
	ApplicationServerVirtualMachineTypeStandby ApplicationServerVirtualMachineType = "Standby"
	// ApplicationServerVirtualMachineTypeUnknown - Unknown Application server vm type.
	ApplicationServerVirtualMachineTypeUnknown ApplicationServerVirtualMachineType = "Unknown"
)

// PossibleApplicationServerVirtualMachineTypeValues returns the possible values for the ApplicationServerVirtualMachineType const type.
func PossibleApplicationServerVirtualMachineTypeValues() []ApplicationServerVirtualMachineType {
	return []ApplicationServerVirtualMachineType{
		ApplicationServerVirtualMachineTypeActive,
		ApplicationServerVirtualMachineTypeStandby,
		ApplicationServerVirtualMachineTypeUnknown,
	}
}

// CentralServerVirtualMachineType - Defines the type of central server VM.
type CentralServerVirtualMachineType string

const (
	// CentralServerVirtualMachineTypeASCS - ASCS Central server vm type.
	CentralServerVirtualMachineTypeASCS CentralServerVirtualMachineType = "ASCS"
	// CentralServerVirtualMachineTypeERS - ERS Central server vm type.
	CentralServerVirtualMachineTypeERS CentralServerVirtualMachineType = "ERS"
	// CentralServerVirtualMachineTypeERSInactive - ERSInactive Central server vm type.
	CentralServerVirtualMachineTypeERSInactive CentralServerVirtualMachineType = "ERSInactive"
	// CentralServerVirtualMachineTypePrimary - Primary central server vm.
	CentralServerVirtualMachineTypePrimary CentralServerVirtualMachineType = "Primary"
	// CentralServerVirtualMachineTypeSecondary - Secondary central server vm.
	CentralServerVirtualMachineTypeSecondary CentralServerVirtualMachineType = "Secondary"
	// CentralServerVirtualMachineTypeStandby - Standby Central server vm type.
	CentralServerVirtualMachineTypeStandby CentralServerVirtualMachineType = "Standby"
	// CentralServerVirtualMachineTypeUnknown - Central server vm type unknown.
	CentralServerVirtualMachineTypeUnknown CentralServerVirtualMachineType = "Unknown"
)

// PossibleCentralServerVirtualMachineTypeValues returns the possible values for the CentralServerVirtualMachineType const type.
func PossibleCentralServerVirtualMachineTypeValues() []CentralServerVirtualMachineType {
	return []CentralServerVirtualMachineType{
		CentralServerVirtualMachineTypeASCS,
		CentralServerVirtualMachineTypeERS,
		CentralServerVirtualMachineTypeERSInactive,
		CentralServerVirtualMachineTypePrimary,
		CentralServerVirtualMachineTypeSecondary,
		CentralServerVirtualMachineTypeStandby,
		CentralServerVirtualMachineTypeUnknown,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DiskSKUName - Defines the disk sku name.
type DiskSKUName string

const (
	// DiskSKUNamePremiumLRS - Premium_LRS Disk SKU.
	DiskSKUNamePremiumLRS DiskSKUName = "Premium_LRS"
	// DiskSKUNamePremiumV2LRS - PremiumV2_LRS Disk SKU.
	DiskSKUNamePremiumV2LRS DiskSKUName = "PremiumV2_LRS"
	// DiskSKUNamePremiumZRS - Premium_ZRS Disk SKU.
	DiskSKUNamePremiumZRS DiskSKUName = "Premium_ZRS"
	// DiskSKUNameStandardLRS - Standard LRS Disk SKU.
	DiskSKUNameStandardLRS DiskSKUName = "Standard_LRS"
	// DiskSKUNameStandardSSDLRS - StandardSSD_LRS Disk SKU.
	DiskSKUNameStandardSSDLRS DiskSKUName = "StandardSSD_LRS"
	// DiskSKUNameStandardSSDZRS - StandardSSD_ZRS Disk SKU.
	DiskSKUNameStandardSSDZRS DiskSKUName = "StandardSSD_ZRS"
	// DiskSKUNameUltraSSDLRS - UltraSSD_LRS Disk SKU.
	DiskSKUNameUltraSSDLRS DiskSKUName = "UltraSSD_LRS"
)

// PossibleDiskSKUNameValues returns the possible values for the DiskSKUName const type.
func PossibleDiskSKUNameValues() []DiskSKUName {
	return []DiskSKUName{
		DiskSKUNamePremiumLRS,
		DiskSKUNamePremiumV2LRS,
		DiskSKUNamePremiumZRS,
		DiskSKUNameStandardLRS,
		DiskSKUNameStandardSSDLRS,
		DiskSKUNameStandardSSDZRS,
		DiskSKUNameUltraSSDLRS,
	}
}

// EnqueueReplicationServerType - Defines the type of Enqueue Replication Server.
type EnqueueReplicationServerType string

const (
	// EnqueueReplicationServerTypeEnqueueReplicator1 - Enqueue Replication server type 1.
	EnqueueReplicationServerTypeEnqueueReplicator1 EnqueueReplicationServerType = "EnqueueReplicator1"
	// EnqueueReplicationServerTypeEnqueueReplicator2 - Enqueue Replication server type 2.
	EnqueueReplicationServerTypeEnqueueReplicator2 EnqueueReplicationServerType = "EnqueueReplicator2"
)

// PossibleEnqueueReplicationServerTypeValues returns the possible values for the EnqueueReplicationServerType const type.
func PossibleEnqueueReplicationServerTypeValues() []EnqueueReplicationServerType {
	return []EnqueueReplicationServerType{
		EnqueueReplicationServerTypeEnqueueReplicator1,
		EnqueueReplicationServerTypeEnqueueReplicator2,
	}
}

// FileShareConfigurationType - The type of file share config.
type FileShareConfigurationType string

const (
	// FileShareConfigurationTypeCreateAndMount - Fileshare will be created and mounted by service.
	FileShareConfigurationTypeCreateAndMount FileShareConfigurationType = "CreateAndMount"
	// FileShareConfigurationTypeMount - Existing fileshare provided will be mounted by service.
	FileShareConfigurationTypeMount FileShareConfigurationType = "Mount"
	// FileShareConfigurationTypeSkip - Skip creating the file share.
	FileShareConfigurationTypeSkip FileShareConfigurationType = "Skip"
)

// PossibleFileShareConfigurationTypeValues returns the possible values for the FileShareConfigurationType const type.
func PossibleFileShareConfigurationTypeValues() []FileShareConfigurationType {
	return []FileShareConfigurationType{
		FileShareConfigurationTypeCreateAndMount,
		FileShareConfigurationTypeMount,
		FileShareConfigurationTypeSkip,
	}
}

// ManagedResourcesNetworkAccessType - Defines the network access type for managed resources.
type ManagedResourcesNetworkAccessType string

const (
	// ManagedResourcesNetworkAccessTypePrivate - Managed resources will be deployed with public network access disabled.
	ManagedResourcesNetworkAccessTypePrivate ManagedResourcesNetworkAccessType = "Private"
	// ManagedResourcesNetworkAccessTypePublic - Managed resources will be deployed with public network access enabled.
	ManagedResourcesNetworkAccessTypePublic ManagedResourcesNetworkAccessType = "Public"
)

// PossibleManagedResourcesNetworkAccessTypeValues returns the possible values for the ManagedResourcesNetworkAccessType const type.
func PossibleManagedResourcesNetworkAccessTypeValues() []ManagedResourcesNetworkAccessType {
	return []ManagedResourcesNetworkAccessType{
		ManagedResourcesNetworkAccessTypePrivate,
		ManagedResourcesNetworkAccessTypePublic,
	}
}

// NamingPatternType - The pattern type to be used for resource naming.
type NamingPatternType string

const (
	// NamingPatternTypeFullResourceName - Full resource names that will be created by service.
	NamingPatternTypeFullResourceName NamingPatternType = "FullResourceName"
)

// PossibleNamingPatternTypeValues returns the possible values for the NamingPatternType const type.
func PossibleNamingPatternTypeValues() []NamingPatternType {
	return []NamingPatternType{
		NamingPatternTypeFullResourceName,
	}
}

// OSType - The OS Type
type OSType string

const (
	// OSTypeLinux - Linux OS Type.
	OSTypeLinux OSType = "Linux"
	// OSTypeWindows - Windows OS Type.
	OSTypeWindows OSType = "Windows"
)

// PossibleOSTypeValues returns the possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{
		OSTypeLinux,
		OSTypeWindows,
	}
}

// SAPConfigurationType - The configuration Type.
type SAPConfigurationType string

const (
	// SAPConfigurationTypeDeployment - SAP system will be deployed by service. No OS configurations will be done.
	SAPConfigurationTypeDeployment SAPConfigurationType = "Deployment"
	// SAPConfigurationTypeDeploymentWithOSConfig - SAP system will be deployed by service. OS configurations will be done.
	SAPConfigurationTypeDeploymentWithOSConfig SAPConfigurationType = "DeploymentWithOSConfig"
	// SAPConfigurationTypeDiscovery - Existing SAP system will be registered.
	SAPConfigurationTypeDiscovery SAPConfigurationType = "Discovery"
)

// PossibleSAPConfigurationTypeValues returns the possible values for the SAPConfigurationType const type.
func PossibleSAPConfigurationTypeValues() []SAPConfigurationType {
	return []SAPConfigurationType{
		SAPConfigurationTypeDeployment,
		SAPConfigurationTypeDeploymentWithOSConfig,
		SAPConfigurationTypeDiscovery,
	}
}

// SAPDatabaseScaleMethod - The database scale method.
type SAPDatabaseScaleMethod string

const (
	// SAPDatabaseScaleMethodScaleUp - ScaleUp Hana Database deployment type
	SAPDatabaseScaleMethodScaleUp SAPDatabaseScaleMethod = "ScaleUp"
)

// PossibleSAPDatabaseScaleMethodValues returns the possible values for the SAPDatabaseScaleMethod const type.
func PossibleSAPDatabaseScaleMethodValues() []SAPDatabaseScaleMethod {
	return []SAPDatabaseScaleMethod{
		SAPDatabaseScaleMethodScaleUp,
	}
}

// SAPDatabaseType - Defines the supported SAP Database types.
type SAPDatabaseType string

const (
	// SAPDatabaseTypeDB2 - DB2 database type of the SAP system.
	SAPDatabaseTypeDB2 SAPDatabaseType = "DB2"
	// SAPDatabaseTypeHANA - HANA Database type of SAP system.
	SAPDatabaseTypeHANA SAPDatabaseType = "HANA"
)

// PossibleSAPDatabaseTypeValues returns the possible values for the SAPDatabaseType const type.
func PossibleSAPDatabaseTypeValues() []SAPDatabaseType {
	return []SAPDatabaseType{
		SAPDatabaseTypeDB2,
		SAPDatabaseTypeHANA,
	}
}

// SAPDeploymentType - The type of SAP deployment, single server or Three tier.
type SAPDeploymentType string

const (
	// SAPDeploymentTypeSingleServer - SAP Single server deployment type.
	SAPDeploymentTypeSingleServer SAPDeploymentType = "SingleServer"
	// SAPDeploymentTypeThreeTier - SAP Distributed deployment type.
	SAPDeploymentTypeThreeTier SAPDeploymentType = "ThreeTier"
)

// PossibleSAPDeploymentTypeValues returns the possible values for the SAPDeploymentType const type.
func PossibleSAPDeploymentTypeValues() []SAPDeploymentType {
	return []SAPDeploymentType{
		SAPDeploymentTypeSingleServer,
		SAPDeploymentTypeThreeTier,
	}
}

// SAPEnvironmentType - Defines the environment type - Production/Non Production.
type SAPEnvironmentType string

const (
	// SAPEnvironmentTypeNonProd - Non Production SAP system.
	SAPEnvironmentTypeNonProd SAPEnvironmentType = "NonProd"
	// SAPEnvironmentTypeProd - Production SAP system.
	SAPEnvironmentTypeProd SAPEnvironmentType = "Prod"
)

// PossibleSAPEnvironmentTypeValues returns the possible values for the SAPEnvironmentType const type.
func PossibleSAPEnvironmentTypeValues() []SAPEnvironmentType {
	return []SAPEnvironmentType{
		SAPEnvironmentTypeNonProd,
		SAPEnvironmentTypeProd,
	}
}

// SAPHealthState - Defines the health of SAP Instances.
type SAPHealthState string

const (
	// SAPHealthStateDegraded - SAP System health is degraded.
	SAPHealthStateDegraded SAPHealthState = "Degraded"
	// SAPHealthStateHealthy - SAP System health is healthy.
	SAPHealthStateHealthy SAPHealthState = "Healthy"
	// SAPHealthStateUnhealthy - SAP System is unhealthy.
	SAPHealthStateUnhealthy SAPHealthState = "Unhealthy"
	// SAPHealthStateUnknown - SAP System health is unknown.
	SAPHealthStateUnknown SAPHealthState = "Unknown"
)

// PossibleSAPHealthStateValues returns the possible values for the SAPHealthState const type.
func PossibleSAPHealthStateValues() []SAPHealthState {
	return []SAPHealthState{
		SAPHealthStateDegraded,
		SAPHealthStateHealthy,
		SAPHealthStateUnhealthy,
		SAPHealthStateUnknown,
	}
}

// SAPHighAvailabilityType - The high availability type (AvailabilitySet or AvailabilityZone).
type SAPHighAvailabilityType string

const (
	// SAPHighAvailabilityTypeAvailabilitySet - HA deployment with availability sets.
	SAPHighAvailabilityTypeAvailabilitySet SAPHighAvailabilityType = "AvailabilitySet"
	// SAPHighAvailabilityTypeAvailabilityZone - HA deployment with availability zones.
	SAPHighAvailabilityTypeAvailabilityZone SAPHighAvailabilityType = "AvailabilityZone"
)

// PossibleSAPHighAvailabilityTypeValues returns the possible values for the SAPHighAvailabilityType const type.
func PossibleSAPHighAvailabilityTypeValues() []SAPHighAvailabilityType {
	return []SAPHighAvailabilityType{
		SAPHighAvailabilityTypeAvailabilitySet,
		SAPHighAvailabilityTypeAvailabilityZone,
	}
}

// SAPProductType - Defines the SAP Product type.
type SAPProductType string

const (
	// SAPProductTypeECC - SAP Product ECC.
	SAPProductTypeECC SAPProductType = "ECC"
	// SAPProductTypeOther - SAP Products other than the ones listed.
	SAPProductTypeOther SAPProductType = "Other"
	// SAPProductTypeS4HANA - SAP Product S4HANA.
	SAPProductTypeS4HANA SAPProductType = "S4HANA"
)

// PossibleSAPProductTypeValues returns the possible values for the SAPProductType const type.
func PossibleSAPProductTypeValues() []SAPProductType {
	return []SAPProductType{
		SAPProductTypeECC,
		SAPProductTypeOther,
		SAPProductTypeS4HANA,
	}
}

// SAPSoftwareInstallationType - The SAP software installation Type.
type SAPSoftwareInstallationType string

const (
	// SAPSoftwareInstallationTypeExternal - External software installation type.
	SAPSoftwareInstallationTypeExternal SAPSoftwareInstallationType = "External"
	// SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig - SAP Install without OS Config.
	SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig SAPSoftwareInstallationType = "SAPInstallWithoutOSConfig"
	// SAPSoftwareInstallationTypeServiceInitiated - SAP Install managed by service.
	SAPSoftwareInstallationTypeServiceInitiated SAPSoftwareInstallationType = "ServiceInitiated"
)

// PossibleSAPSoftwareInstallationTypeValues returns the possible values for the SAPSoftwareInstallationType const type.
func PossibleSAPSoftwareInstallationTypeValues() []SAPSoftwareInstallationType {
	return []SAPSoftwareInstallationType{
		SAPSoftwareInstallationTypeExternal,
		SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig,
		SAPSoftwareInstallationTypeServiceInitiated,
	}
}

// SAPVirtualInstanceIdentityType - Type of managed service identity (where only None and UserAssigned types are allowed).
type SAPVirtualInstanceIdentityType string

const (
	// SAPVirtualInstanceIdentityTypeNone - No managed identity.
	SAPVirtualInstanceIdentityTypeNone SAPVirtualInstanceIdentityType = "None"
	// SAPVirtualInstanceIdentityTypeUserAssigned - User assigned managed identity.
	SAPVirtualInstanceIdentityTypeUserAssigned SAPVirtualInstanceIdentityType = "UserAssigned"
)

// PossibleSAPVirtualInstanceIdentityTypeValues returns the possible values for the SAPVirtualInstanceIdentityType const type.
func PossibleSAPVirtualInstanceIdentityTypeValues() []SAPVirtualInstanceIdentityType {
	return []SAPVirtualInstanceIdentityType{
		SAPVirtualInstanceIdentityTypeNone,
		SAPVirtualInstanceIdentityTypeUserAssigned,
	}
}

// SAPVirtualInstanceState - Defines the Virtual Instance for SAP state.
type SAPVirtualInstanceState string

const (
	// SAPVirtualInstanceStateACSSInstallationBlocked - ACSS installation cannot proceed.
	SAPVirtualInstanceStateACSSInstallationBlocked SAPVirtualInstanceState = "ACSSInstallationBlocked"
	// SAPVirtualInstanceStateDiscoveryFailed - Registration has failed.
	SAPVirtualInstanceStateDiscoveryFailed SAPVirtualInstanceState = "DiscoveryFailed"
	// SAPVirtualInstanceStateDiscoveryInProgress - Registration is in progress.
	SAPVirtualInstanceStateDiscoveryInProgress SAPVirtualInstanceState = "DiscoveryInProgress"
	// SAPVirtualInstanceStateDiscoveryPending - Registration has not started.
	SAPVirtualInstanceStateDiscoveryPending SAPVirtualInstanceState = "DiscoveryPending"
	// SAPVirtualInstanceStateInfrastructureDeploymentFailed - Infrastructure deployment has failed.
	SAPVirtualInstanceStateInfrastructureDeploymentFailed SAPVirtualInstanceState = "InfrastructureDeploymentFailed"
	// SAPVirtualInstanceStateInfrastructureDeploymentInProgress - Infrastructure deployment is in progress.
	SAPVirtualInstanceStateInfrastructureDeploymentInProgress SAPVirtualInstanceState = "InfrastructureDeploymentInProgress"
	// SAPVirtualInstanceStateInfrastructureDeploymentPending - Infrastructure is not yet deployed.
	SAPVirtualInstanceStateInfrastructureDeploymentPending SAPVirtualInstanceState = "InfrastructureDeploymentPending"
	// SAPVirtualInstanceStateRegistrationComplete - Registration is complete.
	SAPVirtualInstanceStateRegistrationComplete SAPVirtualInstanceState = "RegistrationComplete"
	// SAPVirtualInstanceStateSoftwareDetectionFailed - Software detection failed.
	SAPVirtualInstanceStateSoftwareDetectionFailed SAPVirtualInstanceState = "SoftwareDetectionFailed"
	// SAPVirtualInstanceStateSoftwareDetectionInProgress - Software detection is in progress.
	SAPVirtualInstanceStateSoftwareDetectionInProgress SAPVirtualInstanceState = "SoftwareDetectionInProgress"
	// SAPVirtualInstanceStateSoftwareInstallationFailed - Software installation failed.
	SAPVirtualInstanceStateSoftwareInstallationFailed SAPVirtualInstanceState = "SoftwareInstallationFailed"
	// SAPVirtualInstanceStateSoftwareInstallationInProgress - Software installation is in progress.
	SAPVirtualInstanceStateSoftwareInstallationInProgress SAPVirtualInstanceState = "SoftwareInstallationInProgress"
	// SAPVirtualInstanceStateSoftwareInstallationPending - Infrastructure deployment is successful. Software installation is
	// pending.
	SAPVirtualInstanceStateSoftwareInstallationPending SAPVirtualInstanceState = "SoftwareInstallationPending"
)

// PossibleSAPVirtualInstanceStateValues returns the possible values for the SAPVirtualInstanceState const type.
func PossibleSAPVirtualInstanceStateValues() []SAPVirtualInstanceState {
	return []SAPVirtualInstanceState{
		SAPVirtualInstanceStateACSSInstallationBlocked,
		SAPVirtualInstanceStateDiscoveryFailed,
		SAPVirtualInstanceStateDiscoveryInProgress,
		SAPVirtualInstanceStateDiscoveryPending,
		SAPVirtualInstanceStateInfrastructureDeploymentFailed,
		SAPVirtualInstanceStateInfrastructureDeploymentInProgress,
		SAPVirtualInstanceStateInfrastructureDeploymentPending,
		SAPVirtualInstanceStateRegistrationComplete,
		SAPVirtualInstanceStateSoftwareDetectionFailed,
		SAPVirtualInstanceStateSoftwareDetectionInProgress,
		SAPVirtualInstanceStateSoftwareInstallationFailed,
		SAPVirtualInstanceStateSoftwareInstallationInProgress,
		SAPVirtualInstanceStateSoftwareInstallationPending,
	}
}

// SAPVirtualInstanceStatus - Defines the SAP Instance status.
type SAPVirtualInstanceStatus string

const (
	// SAPVirtualInstanceStatusOffline - SAP system is offline.
	SAPVirtualInstanceStatusOffline SAPVirtualInstanceStatus = "Offline"
	// SAPVirtualInstanceStatusPartiallyRunning - SAP system is partially running.
	SAPVirtualInstanceStatusPartiallyRunning SAPVirtualInstanceStatus = "PartiallyRunning"
	// SAPVirtualInstanceStatusRunning - SAP system is running.
	SAPVirtualInstanceStatusRunning SAPVirtualInstanceStatus = "Running"
	// SAPVirtualInstanceStatusSoftShutdown - Soft shutdown of SAP system is initiated.
	SAPVirtualInstanceStatusSoftShutdown SAPVirtualInstanceStatus = "SoftShutdown"
	// SAPVirtualInstanceStatusStarting - SAP system is getting started.
	SAPVirtualInstanceStatusStarting SAPVirtualInstanceStatus = "Starting"
	// SAPVirtualInstanceStatusStopping - SAP system is being stopped.
	SAPVirtualInstanceStatusStopping SAPVirtualInstanceStatus = "Stopping"
	// SAPVirtualInstanceStatusUnavailable - SAP system status is unavailable.
	SAPVirtualInstanceStatusUnavailable SAPVirtualInstanceStatus = "Unavailable"
)

// PossibleSAPVirtualInstanceStatusValues returns the possible values for the SAPVirtualInstanceStatus const type.
func PossibleSAPVirtualInstanceStatusValues() []SAPVirtualInstanceStatus {
	return []SAPVirtualInstanceStatus{
		SAPVirtualInstanceStatusOffline,
		SAPVirtualInstanceStatusPartiallyRunning,
		SAPVirtualInstanceStatusRunning,
		SAPVirtualInstanceStatusSoftShutdown,
		SAPVirtualInstanceStatusStarting,
		SAPVirtualInstanceStatusStopping,
		SAPVirtualInstanceStatusUnavailable,
	}
}

// SapVirtualInstanceProvisioningState - Defines the provisioning states.
type SapVirtualInstanceProvisioningState string

const (
	// SapVirtualInstanceProvisioningStateCanceled - ACSS Canceled provisioning state.
	SapVirtualInstanceProvisioningStateCanceled SapVirtualInstanceProvisioningState = "Canceled"
	// SapVirtualInstanceProvisioningStateCreating - ACSS Creating provisioning state.
	SapVirtualInstanceProvisioningStateCreating SapVirtualInstanceProvisioningState = "Creating"
	// SapVirtualInstanceProvisioningStateDeleting - ACSS Deleting provisioning state.
	SapVirtualInstanceProvisioningStateDeleting SapVirtualInstanceProvisioningState = "Deleting"
	// SapVirtualInstanceProvisioningStateFailed - ACSS Failed provisioning state.
	SapVirtualInstanceProvisioningStateFailed SapVirtualInstanceProvisioningState = "Failed"
	// SapVirtualInstanceProvisioningStateSucceeded - ACSS succeeded provisioning state.
	SapVirtualInstanceProvisioningStateSucceeded SapVirtualInstanceProvisioningState = "Succeeded"
	// SapVirtualInstanceProvisioningStateUpdating - ACSS updating provisioning state.
	SapVirtualInstanceProvisioningStateUpdating SapVirtualInstanceProvisioningState = "Updating"
)

// PossibleSapVirtualInstanceProvisioningStateValues returns the possible values for the SapVirtualInstanceProvisioningState const type.
func PossibleSapVirtualInstanceProvisioningStateValues() []SapVirtualInstanceProvisioningState {
	return []SapVirtualInstanceProvisioningState{
		SapVirtualInstanceProvisioningStateCanceled,
		SapVirtualInstanceProvisioningStateCreating,
		SapVirtualInstanceProvisioningStateDeleting,
		SapVirtualInstanceProvisioningStateFailed,
		SapVirtualInstanceProvisioningStateSucceeded,
		SapVirtualInstanceProvisioningStateUpdating,
	}
}
