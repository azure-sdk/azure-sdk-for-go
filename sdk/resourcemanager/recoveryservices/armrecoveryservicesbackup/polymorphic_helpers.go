// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import "encoding/json"

func unmarshalBackupEngineBaseClassification(rawMsg json.RawMessage) (BackupEngineBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupEngineBaseClassification
	switch m["backupEngineType"] {
	case string(BackupEngineTypeAzureBackupServerEngine):
		b = &AzureBackupServerEngine{}
	case string(BackupEngineTypeDpmBackupEngine):
		b = &DpmBackupEngine{}
	default:
		b = &BackupEngineBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalBackupRequestClassification(rawMsg json.RawMessage) (BackupRequestClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupRequestClassification
	switch m["objectType"] {
	case "AzureFileShareBackupRequest":
		b = &AzureFileShareBackupRequest{}
	case "AzureWorkloadBackupRequest":
		b = &AzureWorkloadBackupRequest{}
	case "IaasVMBackupRequest":
		b = &IaasVMBackupRequest{}
	default:
		b = &BackupRequest{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalILRRequestClassification(rawMsg json.RawMessage) (ILRRequestClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ILRRequestClassification
	switch m["objectType"] {
	case "AzureFileShareProvisionILRRequest":
		b = &AzureFileShareProvisionILRRequest{}
	case "IaasVMILRRegistrationRequest":
		b = &IaasVMILRRegistrationRequest{}
	default:
		b = &ILRRequest{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalJobClassification(rawMsg json.RawMessage) (JobClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobClassification
	switch m["jobType"] {
	case "AzureIaaSVMJob":
		b = &AzureIaaSVMJob{}
	case "AzureIaaSVMJobV2":
		b = &AzureIaaSVMJobV2{}
	case "AzureStorageJob":
		b = &AzureStorageJob{}
	case "AzureWorkloadJob":
		b = &AzureWorkloadJob{}
	case "DpmJob":
		b = &DpmJob{}
	case "MabJob":
		b = &MabJob{}
	case "VaultJob":
		b = &VaultJob{}
	default:
		b = &Job{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOperationResultInfoBaseClassification(rawMsg json.RawMessage) (OperationResultInfoBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OperationResultInfoBaseClassification
	switch m["objectType"] {
	case "ExportJobsOperationResultInfo":
		b = &ExportJobsOperationResultInfo{}
	case "OperationResultInfo":
		b = &OperationResultInfo{}
	default:
		b = &OperationResultInfoBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOperationStatusExtendedInfoClassification(rawMsg json.RawMessage) (OperationStatusExtendedInfoClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OperationStatusExtendedInfoClassification
	switch m["objectType"] {
	case "OperationStatusJobExtendedInfo":
		b = &OperationStatusJobExtendedInfo{}
	case "OperationStatusJobsExtendedInfo":
		b = &OperationStatusJobsExtendedInfo{}
	case "OperationStatusProvisionILRExtendedInfo":
		b = &OperationStatusProvisionILRExtendedInfo{}
	case "OperationStatusValidateOperationExtendedInfo":
		b = &OperationStatusValidateOperationExtendedInfo{}
	default:
		b = &OperationStatusExtendedInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectableContainerClassification(rawMsg json.RawMessage) (ProtectableContainerClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectableContainerClassification
	switch m["protectableContainerType"] {
	case string(ProtectableContainerTypeStorageContainer):
		b = &AzureStorageProtectableContainer{}
	case string(ProtectableContainerTypeVMAppContainer):
		b = &AzureVMAppContainerProtectableContainer{}
	default:
		b = &ProtectableContainer{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectedItemClassification(rawMsg json.RawMessage) (ProtectedItemClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectedItemClassification
	switch m["protectedItemType"] {
	case "AzureFileShareProtectedItem":
		b = &AzureFileshareProtectedItem{}
	case "AzureIaaSVMProtectedItem":
		b = &AzureIaaSVMProtectedItem{}
	case "AzureVmWorkloadAnyDatabase":
		b = &AzureVMWorkloadAnyDatabaseProtectedItem{}
	case "AzureVmWorkloadOracleDatabase":
		b = &AzureVMWorkloadOracleDatabaseProtectedItem{}
	case "AzureVmWorkloadProtectedItem":
		b = &AzureVMWorkloadProtectedItem{}
	case "AzureVmWorkloadSAPAseDatabase":
		b = &AzureVMWorkloadSAPAseDatabaseProtectedItem{}
	case "AzureVmWorkloadSAPHanaDBInstance":
		b = &AzureVMWorkloadSAPHanaDBInstanceProtectedItem{}
	case "AzureVmWorkloadSAPHanaDatabase":
		b = &AzureVMWorkloadSAPHanaDatabaseProtectedItem{}
	case "AzureVmWorkloadSQLDatabase":
		b = &AzureVMWorkloadSQLDatabaseProtectedItem{}
	case "DPMProtectedItem":
		b = &DPMProtectedItem{}
	case "GenericProtectedItem":
		b = &GenericProtectedItem{}
	case "MabFileFolderProtectedItem":
		b = &MabFileFolderProtectedItem{}
	case "Microsoft.ClassicCompute/virtualMachines":
		b = &AzureIaaSClassicComputeVMProtectedItem{}
	case "Microsoft.Compute/virtualMachines":
		b = &AzureIaaSComputeVMProtectedItem{}
	case "Microsoft.Sql/servers/databases":
		b = &AzureSQLProtectedItem{}
	default:
		b = &ProtectedItem{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectionContainerClassification(rawMsg json.RawMessage) (ProtectionContainerClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectionContainerClassification
	switch m["containerType"] {
	case string(ProtectableContainerTypeAzureBackupServerContainer):
		b = &AzureBackupServerContainer{}
	case string(ProtectableContainerTypeAzureSQLContainer):
		b = &AzureSQLContainer{}
	case string(ProtectableContainerTypeAzureWorkloadContainer):
		b = &AzureWorkloadContainer{}
	case string(ProtectableContainerTypeDPMContainer):
		b = &DpmContainer{}
	case string(ProtectableContainerTypeGenericContainer):
		b = &GenericContainer{}
	case string(ProtectableContainerTypeIaasVMContainer):
		b = &IaaSVMContainer{}
	case string(ProtectableContainerTypeMicrosoftClassicComputeVirtualMachines):
		b = &AzureIaaSClassicComputeVMContainer{}
	case string(ProtectableContainerTypeMicrosoftComputeVirtualMachines):
		b = &AzureIaaSComputeVMContainer{}
	case string(ProtectableContainerTypeSQLAGWorkLoadContainer):
		b = &AzureSQLAGWorkloadContainerProtectionContainer{}
	case string(ProtectableContainerTypeStorageContainer):
		b = &AzureStorageContainer{}
	case string(ProtectableContainerTypeVMAppContainer):
		b = &AzureVMAppContainerProtectionContainer{}
	case string(ProtectableContainerTypeWindows):
		b = &MabContainer{}
	default:
		b = &ProtectionContainer{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectionIntentClassification(rawMsg json.RawMessage) (ProtectionIntentClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectionIntentClassification
	switch m["protectionIntentItemType"] {
	case "ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent":
		b = &AzureWorkloadAutoProtectionIntent{}
	case "ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent":
		b = &AzureWorkloadSQLAutoProtectionIntent{}
	case string(ProtectionIntentItemTypeAzureResourceItem):
		b = &AzureResourceProtectionIntent{}
	case string(ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent):
		b = &AzureWorkloadContainerAutoProtectionIntent{}
	case string(ProtectionIntentItemTypeRecoveryServiceVaultItem):
		b = &AzureRecoveryServiceVaultProtectionIntent{}
	default:
		b = &ProtectionIntent{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectionPolicyClassification(rawMsg json.RawMessage) (ProtectionPolicyClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectionPolicyClassification
	switch m["backupManagementType"] {
	case "AzureIaasVM":
		b = &AzureIaaSVMProtectionPolicy{}
	case "AzureSql":
		b = &AzureSQLProtectionPolicy{}
	case "AzureStorage":
		b = &AzureFileShareProtectionPolicy{}
	case "AzureWorkload":
		b = &AzureVMWorkloadProtectionPolicy{}
	case "GenericProtectionPolicy":
		b = &GenericProtectionPolicy{}
	case "MAB":
		b = &MabProtectionPolicy{}
	default:
		b = &ProtectionPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRecoveryPointClassification(rawMsg json.RawMessage) (RecoveryPointClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPointClassification
	switch m["objectType"] {
	case "AzureFileShareRecoveryPoint":
		b = &AzureFileShareRecoveryPoint{}
	case "AzureWorkloadAnyDatabasePointInTimeRecoveryPoint":
		b = &AzureWorkloadAnyDatabasePointInTimeRecoveryPoint{}
	case "AzureWorkloadAnyDatabaseRecoveryPoint":
		b = &AzureWorkloadAnyDatabaseRecoveryPoint{}
	case "AzureWorkloadOraclePointInTimeRecoveryPoint":
		b = &AzureWorkloadOraclePointInTimeRecoveryPoint{}
	case "AzureWorkloadOracleRecoveryPoint":
		b = &AzureWorkloadOracleRecoveryPoint{}
	case "AzureWorkloadPointInTimeRecoveryPoint":
		b = &AzureWorkloadPointInTimeRecoveryPoint{}
	case "AzureWorkloadRecoveryPoint":
		b = &AzureWorkloadRecoveryPoint{}
	case "AzureWorkloadSAPAsePointInTimeRecoveryPoint":
		b = &AzureWorkloadSAPAsePointInTimeRecoveryPoint{}
	case "AzureWorkloadSAPAseRecoveryPoint":
		b = &AzureWorkloadSAPAseRecoveryPoint{}
	case "AzureWorkloadSAPHanaPointInTimeRecoveryPoint":
		b = &AzureWorkloadSAPHanaPointInTimeRecoveryPoint{}
	case "AzureWorkloadSAPHanaRecoveryPoint":
		b = &AzureWorkloadSAPHanaRecoveryPoint{}
	case "AzureWorkloadSQLPointInTimeRecoveryPoint":
		b = &AzureWorkloadSQLPointInTimeRecoveryPoint{}
	case "AzureWorkloadSQLRecoveryPoint":
		b = &AzureWorkloadSQLRecoveryPoint{}
	case "GenericRecoveryPoint":
		b = &GenericRecoveryPoint{}
	case "IaasVMRecoveryPoint":
		b = &IaasVMRecoveryPoint{}
	default:
		b = &RecoveryPoint{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRestoreRequestClassification(rawMsg json.RawMessage) (RestoreRequestClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RestoreRequestClassification
	switch m["objectType"] {
	case "AzureFileShareRestoreRequest":
		b = &AzureFileShareRestoreRequest{}
	case "AzureWorkloadAnyDatabasePointInTimeRestoreRequest":
		b = &AzureWorkloadAnyDatabasePointInTimeRestoreRequest{}
	case "AzureWorkloadAnyDatabasePointInTimeRestoreWithRehydrateRequest":
		b = &AzureWorkloadAnyDatabasePointInTimeRestoreWithRehydrateRequest{}
	case "AzureWorkloadAnyDatabaseRestoreRequest":
		b = &AzureWorkloadAnyDatabaseRestoreRequest{}
	case "AzureWorkloadAnyDatabaseRestoreWithRehydrateRequest":
		b = &AzureWorkloadAnyDatabaseRestoreWithRehydrateRequest{}
	case "AzureWorkloadOraclePointInTimeRestoreRequest":
		b = &AzureWorkloadOraclePointInTimeRestoreRequest{}
	case "AzureWorkloadOraclePointInTimeRestoreWithRehydrateRequest":
		b = &AzureWorkloadOraclePointInTimeRestoreWithRehydrateRequest{}
	case "AzureWorkloadOracleRestoreRequest":
		b = &AzureWorkloadOracleRestoreRequest{}
	case "AzureWorkloadOracleRestoreWithRehydrateRequest":
		b = &AzureWorkloadOracleRestoreWithRehydrateRequest{}
	case "AzureWorkloadPointInTimeRestoreRequest":
		b = &AzureWorkloadPointInTimeRestoreRequest{}
	case "AzureWorkloadRestoreRequest":
		b = &AzureWorkloadRestoreRequest{}
	case "AzureWorkloadSAPAsePointInTimeRestoreRequest":
		b = &AzureWorkloadSAPAsePointInTimeRestoreRequest{}
	case "AzureWorkloadSAPAsePointInTimeRestoreWithRehydrateRequest":
		b = &AzureWorkloadSAPAsePointInTimeRestoreWithRehydrateRequest{}
	case "AzureWorkloadSAPAseRestoreRequest":
		b = &AzureWorkloadSAPAseRestoreRequest{}
	case "AzureWorkloadSAPAseRestoreWithRehydrateRequest":
		b = &AzureWorkloadSAPAseRestoreWithRehydrateRequest{}
	case "AzureWorkloadSAPHanaPointInTimeRestoreRequest":
		b = &AzureWorkloadSAPHanaPointInTimeRestoreRequest{}
	case "AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest":
		b = &AzureWorkloadSAPHanaPointInTimeRestoreWithRehydrateRequest{}
	case "AzureWorkloadSAPHanaRestoreRequest":
		b = &AzureWorkloadSAPHanaRestoreRequest{}
	case "AzureWorkloadSAPHanaRestoreWithRehydrateRequest":
		b = &AzureWorkloadSAPHanaRestoreWithRehydrateRequest{}
	case "AzureWorkloadSQLPointInTimeRestoreRequest":
		b = &AzureWorkloadSQLPointInTimeRestoreRequest{}
	case "AzureWorkloadSQLPointInTimeRestoreWithRehydrateRequest":
		b = &AzureWorkloadSQLPointInTimeRestoreWithRehydrateRequest{}
	case "AzureWorkloadSQLRestoreRequest":
		b = &AzureWorkloadSQLRestoreRequest{}
	case "AzureWorkloadSQLRestoreWithRehydrateRequest":
		b = &AzureWorkloadSQLRestoreWithRehydrateRequest{}
	case "IaasVMRestoreRequest":
		b = &IaasVMRestoreRequest{}
	case "IaasVMRestoreWithRehydrationRequest":
		b = &IaasVMRestoreWithRehydrationRequest{}
	default:
		b = &RestoreRequest{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRetentionPolicyClassification(rawMsg json.RawMessage) (RetentionPolicyClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RetentionPolicyClassification
	switch m["retentionPolicyType"] {
	case "LongTermRetentionPolicy":
		b = &LongTermRetentionPolicy{}
	case "SimpleRetentionPolicy":
		b = &SimpleRetentionPolicy{}
	default:
		b = &RetentionPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSchedulePolicyClassification(rawMsg json.RawMessage) (SchedulePolicyClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SchedulePolicyClassification
	switch m["schedulePolicyType"] {
	case "LogSchedulePolicy":
		b = &LogSchedulePolicy{}
	case "LongTermSchedulePolicy":
		b = &LongTermSchedulePolicy{}
	case "SimpleSchedulePolicy":
		b = &SimpleSchedulePolicy{}
	case "SimpleSchedulePolicyV2":
		b = &SimpleSchedulePolicyV2{}
	default:
		b = &SchedulePolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTieringCostInfoClassification(rawMsg json.RawMessage) (TieringCostInfoClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TieringCostInfoClassification
	switch m["objectType"] {
	case "TieringCostRehydrationInfo":
		b = &TieringCostRehydrationInfo{}
	case "TieringCostSavingInfo":
		b = &TieringCostSavingInfo{}
	default:
		b = &TieringCostInfo{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalValidateOperationRequestClassification(rawMsg json.RawMessage) (ValidateOperationRequestClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ValidateOperationRequestClassification
	switch m["objectType"] {
	case "ValidateIaasVMRestoreOperationRequest":
		b = &ValidateIaasVMRestoreOperationRequest{}
	case "ValidateRestoreOperationRequest":
		b = &ValidateRestoreOperationRequest{}
	default:
		b = &ValidateOperationRequest{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalVaultStorageConfigOperationResultResponseClassification(rawMsg json.RawMessage) (VaultStorageConfigOperationResultResponseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VaultStorageConfigOperationResultResponseClassification
	switch m["objectType"] {
	case "PrepareDataMoveResponse":
		b = &PrepareDataMoveResponse{}
	default:
		b = &VaultStorageConfigOperationResultResponse{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWorkloadItemClassification(rawMsg json.RawMessage) (WorkloadItemClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkloadItemClassification
	switch m["workloadItemType"] {
	case "AnyDataBase":
		b = &AzureVMWorkloadAnyDatabaseWorkloadItem{}
	case "AzureVmWorkloadItem":
		b = &AzureVMWorkloadItem{}
	case "OracleDataBase":
		b = &AzureVMWorkloadOracleDatabaseWorkloadItem{}
	case "SAPAseDatabase":
		b = &AzureVMWorkloadSAPAseDatabaseWorkloadItem{}
	case "SAPAseSystem":
		b = &AzureVMWorkloadSAPAseSystemWorkloadItem{}
	case "SAPHanaDatabase":
		b = &AzureVMWorkloadSAPHanaDatabaseWorkloadItem{}
	case "SAPHanaSystem":
		b = &AzureVMWorkloadSAPHanaSystemWorkloadItem{}
	case "SQLDataBase":
		b = &AzureVMWorkloadSQLDatabaseWorkloadItem{}
	case "SQLInstance":
		b = &AzureVMWorkloadSQLInstanceWorkloadItem{}
	default:
		b = &WorkloadItem{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWorkloadProtectableItemClassification(rawMsg json.RawMessage) (WorkloadProtectableItemClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkloadProtectableItemClassification
	switch m["protectableItemType"] {
	case "AnyDatabase":
		b = &AzureVMWorkloadAnyDatabaseProtectableItem{}
	case "AzureFileShare":
		b = &AzureFileShareProtectableItem{}
	case "AzureVmWorkloadProtectableItem":
		b = &AzureVMWorkloadProtectableItem{}
	case "HanaHSRContainer":
		b = &AzureVMWorkloadSAPHanaHSRProtectableItem{}
	case "HanaScaleoutContainer":
		b = &AzureVMWorkloadSAPHanaScaleoutProtectableItem{}
	case "IaaSVMProtectableItem":
		b = &IaaSVMProtectableItem{}
	case "Microsoft.ClassicCompute/virtualMachines":
		b = &AzureIaaSClassicComputeVMProtectableItem{}
	case "Microsoft.Compute/virtualMachines":
		b = &AzureIaaSComputeVMProtectableItem{}
	case "OracleDatabase":
		b = &AzureVMWorkloadOracleDatabaseProtectableItem{}
	case "SAPAseSystem":
		b = &AzureVMWorkloadSAPAseSystemProtectableItem{}
	case "SAPHanaDBInstance":
		b = &AzureVMWorkloadSAPHanaDBInstance{}
	case "SAPHanaDatabase":
		b = &AzureVMWorkloadSAPHanaDatabaseProtectableItem{}
	case "SAPHanaSystem":
		b = &AzureVMWorkloadSAPHanaSystemProtectableItem{}
	case "SQLAvailabilityGroupContainer":
		b = &AzureVMWorkloadSQLAvailabilityGroupProtectableItem{}
	case "SQLDataBase":
		b = &AzureVMWorkloadSQLDatabaseProtectableItem{}
	case "SQLInstance":
		b = &AzureVMWorkloadSQLInstanceProtectableItem{}
	default:
		b = &WorkloadProtectableItem{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
