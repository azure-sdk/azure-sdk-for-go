//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

const (
	moduleName    = "armsqlvirtualmachine"
	moduleVersion = "v0.6.0"
)

// AssessmentDayOfWeek - Day of the week to run assessment.
type AssessmentDayOfWeek string

const (
	AssessmentDayOfWeekMonday    AssessmentDayOfWeek = "Monday"
	AssessmentDayOfWeekTuesday   AssessmentDayOfWeek = "Tuesday"
	AssessmentDayOfWeekWednesday AssessmentDayOfWeek = "Wednesday"
	AssessmentDayOfWeekThursday  AssessmentDayOfWeek = "Thursday"
	AssessmentDayOfWeekFriday    AssessmentDayOfWeek = "Friday"
	AssessmentDayOfWeekSaturday  AssessmentDayOfWeek = "Saturday"
	AssessmentDayOfWeekSunday    AssessmentDayOfWeek = "Sunday"
)

// PossibleAssessmentDayOfWeekValues returns the possible values for the AssessmentDayOfWeek const type.
func PossibleAssessmentDayOfWeekValues() []AssessmentDayOfWeek {
	return []AssessmentDayOfWeek{
		AssessmentDayOfWeekMonday,
		AssessmentDayOfWeekTuesday,
		AssessmentDayOfWeekWednesday,
		AssessmentDayOfWeekThursday,
		AssessmentDayOfWeekFriday,
		AssessmentDayOfWeekSaturday,
		AssessmentDayOfWeekSunday,
	}
}

type AutoBackupDaysOfWeek string

const (
	AutoBackupDaysOfWeekFriday    AutoBackupDaysOfWeek = "Friday"
	AutoBackupDaysOfWeekMonday    AutoBackupDaysOfWeek = "Monday"
	AutoBackupDaysOfWeekSaturday  AutoBackupDaysOfWeek = "Saturday"
	AutoBackupDaysOfWeekSunday    AutoBackupDaysOfWeek = "Sunday"
	AutoBackupDaysOfWeekThursday  AutoBackupDaysOfWeek = "Thursday"
	AutoBackupDaysOfWeekTuesday   AutoBackupDaysOfWeek = "Tuesday"
	AutoBackupDaysOfWeekWednesday AutoBackupDaysOfWeek = "Wednesday"
)

// PossibleAutoBackupDaysOfWeekValues returns the possible values for the AutoBackupDaysOfWeek const type.
func PossibleAutoBackupDaysOfWeekValues() []AutoBackupDaysOfWeek {
	return []AutoBackupDaysOfWeek{
		AutoBackupDaysOfWeekFriday,
		AutoBackupDaysOfWeekMonday,
		AutoBackupDaysOfWeekSaturday,
		AutoBackupDaysOfWeekSunday,
		AutoBackupDaysOfWeekThursday,
		AutoBackupDaysOfWeekTuesday,
		AutoBackupDaysOfWeekWednesday,
	}
}

// BackupScheduleType - Backup schedule type.
type BackupScheduleType string

const (
	BackupScheduleTypeAutomated BackupScheduleType = "Automated"
	BackupScheduleTypeManual    BackupScheduleType = "Manual"
)

// PossibleBackupScheduleTypeValues returns the possible values for the BackupScheduleType const type.
func PossibleBackupScheduleTypeValues() []BackupScheduleType {
	return []BackupScheduleType{
		BackupScheduleTypeAutomated,
		BackupScheduleTypeManual,
	}
}

// ClusterConfiguration - Cluster type.
type ClusterConfiguration string

const (
	ClusterConfigurationDomainful ClusterConfiguration = "Domainful"
)

// PossibleClusterConfigurationValues returns the possible values for the ClusterConfiguration const type.
func PossibleClusterConfigurationValues() []ClusterConfiguration {
	return []ClusterConfiguration{
		ClusterConfigurationDomainful,
	}
}

// ClusterManagerType - Type of cluster manager: Windows Server Failover Cluster (WSFC), implied by the scale type of the
// group and the OS type.
type ClusterManagerType string

const (
	ClusterManagerTypeWSFC ClusterManagerType = "WSFC"
)

// PossibleClusterManagerTypeValues returns the possible values for the ClusterManagerType const type.
func PossibleClusterManagerTypeValues() []ClusterManagerType {
	return []ClusterManagerType{
		ClusterManagerTypeWSFC,
	}
}

// ClusterSubnetType - Cluster subnet type.
type ClusterSubnetType string

const (
	ClusterSubnetTypeMultiSubnet  ClusterSubnetType = "MultiSubnet"
	ClusterSubnetTypeSingleSubnet ClusterSubnetType = "SingleSubnet"
)

// PossibleClusterSubnetTypeValues returns the possible values for the ClusterSubnetType const type.
func PossibleClusterSubnetTypeValues() []ClusterSubnetType {
	return []ClusterSubnetType{
		ClusterSubnetTypeMultiSubnet,
		ClusterSubnetTypeSingleSubnet,
	}
}

// Commit - Replica commit mode in availability group.
type Commit string

const (
	CommitASYNCHRONOUSCOMMIT Commit = "ASYNCHRONOUS_COMMIT"
	CommitSYNCHRONOUSCOMMIT  Commit = "SYNCHRONOUS_COMMIT"
)

// PossibleCommitValues returns the possible values for the Commit const type.
func PossibleCommitValues() []Commit {
	return []Commit{
		CommitASYNCHRONOUSCOMMIT,
		CommitSYNCHRONOUSCOMMIT,
	}
}

// ConnectivityType - SQL Server connectivity option.
type ConnectivityType string

const (
	ConnectivityTypeLOCAL   ConnectivityType = "LOCAL"
	ConnectivityTypePRIVATE ConnectivityType = "PRIVATE"
	ConnectivityTypePUBLIC  ConnectivityType = "PUBLIC"
)

// PossibleConnectivityTypeValues returns the possible values for the ConnectivityType const type.
func PossibleConnectivityTypeValues() []ConnectivityType {
	return []ConnectivityType{
		ConnectivityTypeLOCAL,
		ConnectivityTypePRIVATE,
		ConnectivityTypePUBLIC,
	}
}

// DayOfWeek - Day of week to apply the patch on.
type DayOfWeek string

const (
	DayOfWeekEveryday  DayOfWeek = "Everyday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekEveryday,
		DayOfWeekMonday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
		DayOfWeekThursday,
		DayOfWeekFriday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
	}
}

// DiskConfigurationType - Disk configuration to apply to SQL Server.
type DiskConfigurationType string

const (
	DiskConfigurationTypeADD    DiskConfigurationType = "ADD"
	DiskConfigurationTypeEXTEND DiskConfigurationType = "EXTEND"
	DiskConfigurationTypeNEW    DiskConfigurationType = "NEW"
)

// PossibleDiskConfigurationTypeValues returns the possible values for the DiskConfigurationType const type.
func PossibleDiskConfigurationTypeValues() []DiskConfigurationType {
	return []DiskConfigurationType{
		DiskConfigurationTypeADD,
		DiskConfigurationTypeEXTEND,
		DiskConfigurationTypeNEW,
	}
}

// Failover - Replica failover mode in availability group.
type Failover string

const (
	FailoverAUTOMATIC Failover = "AUTOMATIC"
	FailoverMANUAL    Failover = "MANUAL"
)

// PossibleFailoverValues returns the possible values for the Failover const type.
func PossibleFailoverValues() []Failover {
	return []Failover{
		FailoverAUTOMATIC,
		FailoverMANUAL,
	}
}

// FullBackupFrequencyType - Frequency of full backups. In both cases, full backups begin during the next scheduled time window.
type FullBackupFrequencyType string

const (
	FullBackupFrequencyTypeDaily  FullBackupFrequencyType = "Daily"
	FullBackupFrequencyTypeWeekly FullBackupFrequencyType = "Weekly"
)

// PossibleFullBackupFrequencyTypeValues returns the possible values for the FullBackupFrequencyType const type.
func PossibleFullBackupFrequencyTypeValues() []FullBackupFrequencyType {
	return []FullBackupFrequencyType{
		FullBackupFrequencyTypeDaily,
		FullBackupFrequencyTypeWeekly,
	}
}

// IdentityTypeWithNone - The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure
// Active Directory principal for the resource.
type IdentityTypeWithNone string

const (
	IdentityTypeWithNoneNone           IdentityTypeWithNone = "None"
	IdentityTypeWithNoneSystemAssigned IdentityTypeWithNone = "SystemAssigned"
)

// PossibleIdentityTypeWithNoneValues returns the possible values for the IdentityTypeWithNone const type.
func PossibleIdentityTypeWithNoneValues() []IdentityTypeWithNone {
	return []IdentityTypeWithNone{
		IdentityTypeWithNoneNone,
		IdentityTypeWithNoneSystemAssigned,
	}
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginSystem OperationOrigin = "system"
	OperationOriginUser   OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginSystem,
		OperationOriginUser,
	}
}

// ReadableSecondary - Replica readable secondary mode in availability group.
type ReadableSecondary string

const (
	ReadableSecondaryALL      ReadableSecondary = "ALL"
	ReadableSecondaryNO       ReadableSecondary = "NO"
	ReadableSecondaryREADONLY ReadableSecondary = "READ_ONLY"
)

// PossibleReadableSecondaryValues returns the possible values for the ReadableSecondary const type.
func PossibleReadableSecondaryValues() []ReadableSecondary {
	return []ReadableSecondary{
		ReadableSecondaryALL,
		ReadableSecondaryNO,
		ReadableSecondaryREADONLY,
	}
}

// Role - Replica Role in availability group.
type Role string

const (
	RolePRIMARY   Role = "PRIMARY"
	RoleSECONDARY Role = "SECONDARY"
)

// PossibleRoleValues returns the possible values for the Role const type.
func PossibleRoleValues() []Role {
	return []Role{
		RolePRIMARY,
		RoleSECONDARY,
	}
}

// SQLImageSKU - SQL Server edition type.
type SQLImageSKU string

const (
	SQLImageSKUDeveloper  SQLImageSKU = "Developer"
	SQLImageSKUEnterprise SQLImageSKU = "Enterprise"
	SQLImageSKUExpress    SQLImageSKU = "Express"
	SQLImageSKUStandard   SQLImageSKU = "Standard"
	SQLImageSKUWeb        SQLImageSKU = "Web"
)

// PossibleSQLImageSKUValues returns the possible values for the SQLImageSKU const type.
func PossibleSQLImageSKUValues() []SQLImageSKU {
	return []SQLImageSKU{
		SQLImageSKUDeveloper,
		SQLImageSKUEnterprise,
		SQLImageSKUExpress,
		SQLImageSKUStandard,
		SQLImageSKUWeb,
	}
}

// SQLManagementMode - SQL Server Management type.
type SQLManagementMode string

const (
	SQLManagementModeFull        SQLManagementMode = "Full"
	SQLManagementModeLightWeight SQLManagementMode = "LightWeight"
	SQLManagementModeNoAgent     SQLManagementMode = "NoAgent"
)

// PossibleSQLManagementModeValues returns the possible values for the SQLManagementMode const type.
func PossibleSQLManagementModeValues() []SQLManagementMode {
	return []SQLManagementMode{
		SQLManagementModeFull,
		SQLManagementModeLightWeight,
		SQLManagementModeNoAgent,
	}
}

// SQLServerLicenseType - SQL Server license type.
type SQLServerLicenseType string

const (
	SQLServerLicenseTypeAHUB SQLServerLicenseType = "AHUB"
	SQLServerLicenseTypeDR   SQLServerLicenseType = "DR"
	SQLServerLicenseTypePAYG SQLServerLicenseType = "PAYG"
)

// PossibleSQLServerLicenseTypeValues returns the possible values for the SQLServerLicenseType const type.
func PossibleSQLServerLicenseTypeValues() []SQLServerLicenseType {
	return []SQLServerLicenseType{
		SQLServerLicenseTypeAHUB,
		SQLServerLicenseTypeDR,
		SQLServerLicenseTypePAYG,
	}
}

// SQLVMGroupImageSKU - SQL image sku.
type SQLVMGroupImageSKU string

const (
	SQLVMGroupImageSKUDeveloper  SQLVMGroupImageSKU = "Developer"
	SQLVMGroupImageSKUEnterprise SQLVMGroupImageSKU = "Enterprise"
)

// PossibleSQLVMGroupImageSKUValues returns the possible values for the SQLVMGroupImageSKU const type.
func PossibleSQLVMGroupImageSKUValues() []SQLVMGroupImageSKU {
	return []SQLVMGroupImageSKU{
		SQLVMGroupImageSKUDeveloper,
		SQLVMGroupImageSKUEnterprise,
	}
}

// SQLWorkloadType - SQL Server workload type.
type SQLWorkloadType string

const (
	SQLWorkloadTypeDW      SQLWorkloadType = "DW"
	SQLWorkloadTypeGENERAL SQLWorkloadType = "GENERAL"
	SQLWorkloadTypeOLTP    SQLWorkloadType = "OLTP"
)

// PossibleSQLWorkloadTypeValues returns the possible values for the SQLWorkloadType const type.
func PossibleSQLWorkloadTypeValues() []SQLWorkloadType {
	return []SQLWorkloadType{
		SQLWorkloadTypeDW,
		SQLWorkloadTypeGENERAL,
		SQLWorkloadTypeOLTP,
	}
}

// ScaleType - Scale type.
type ScaleType string

const (
	ScaleTypeHA ScaleType = "HA"
)

// PossibleScaleTypeValues returns the possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{
		ScaleTypeHA,
	}
}

// StorageWorkloadType - Storage workload type.
type StorageWorkloadType string

const (
	StorageWorkloadTypeDW      StorageWorkloadType = "DW"
	StorageWorkloadTypeGENERAL StorageWorkloadType = "GENERAL"
	StorageWorkloadTypeOLTP    StorageWorkloadType = "OLTP"
)

// PossibleStorageWorkloadTypeValues returns the possible values for the StorageWorkloadType const type.
func PossibleStorageWorkloadTypeValues() []StorageWorkloadType {
	return []StorageWorkloadType{
		StorageWorkloadTypeDW,
		StorageWorkloadTypeGENERAL,
		StorageWorkloadTypeOLTP,
	}
}
