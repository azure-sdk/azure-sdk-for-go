//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import "encoding/json"

func unmarshalAssetReferenceBaseClassification(rawMsg json.RawMessage) (AssetReferenceBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AssetReferenceBaseClassification
	switch m["referenceType"] {
	case string(ReferenceTypeDataPath):
		b = &DataPathAssetReference{}
	case string(ReferenceTypeID):
		b = &IDAssetReference{}
	case string(ReferenceTypeOutputPath):
		b = &OutputPathAssetReference{}
	default:
		b = &AssetReferenceBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAutoMLVerticalClassification(rawMsg json.RawMessage) (AutoMLVerticalClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AutoMLVerticalClassification
	switch m["taskType"] {
	case string(TaskTypeClassification):
		b = &Classification{}
	case string(TaskTypeForecasting):
		b = &Forecasting{}
	case string(TaskTypeImageClassification):
		b = &ImageClassification{}
	case string(TaskTypeImageClassificationMultilabel):
		b = &ImageClassificationMultilabel{}
	case string(TaskTypeImageInstanceSegmentation):
		b = &ImageInstanceSegmentation{}
	case string(TaskTypeImageObjectDetection):
		b = &ImageObjectDetection{}
	case string(TaskTypeRegression):
		b = &Regression{}
	case string(TaskTypeTextClassification):
		b = &TextClassification{}
	case string(TaskTypeTextClassificationMultilabel):
		b = &TextClassificationMultilabel{}
	case string(TaskTypeTextNER):
		b = &TextNer{}
	default:
		b = &AutoMLVertical{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalBatchDeploymentConfigurationClassification(rawMsg json.RawMessage) (BatchDeploymentConfigurationClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BatchDeploymentConfigurationClassification
	switch m["deploymentConfigurationType"] {
	case string(BatchDeploymentConfigurationTypePipelineComponent):
		b = &BatchPipelineComponentDeploymentConfiguration{}
	default:
		b = &BatchDeploymentConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalComputeClassification(rawMsg json.RawMessage) (ComputeClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AKS{}
	case string(ComputeTypeAmlCompute):
		b = &AmlCompute{}
	case string(ComputeTypeComputeInstance):
		b = &ComputeInstance{}
	case string(ComputeTypeDataFactory):
		b = &DataFactory{}
	case string(ComputeTypeDataLakeAnalytics):
		b = &DataLakeAnalytics{}
	case string(ComputeTypeDatabricks):
		b = &Databricks{}
	case string(ComputeTypeHDInsight):
		b = &HDInsight{}
	case string(ComputeTypeKubernetes):
		b = &Kubernetes{}
	case string(ComputeTypeSynapseSpark):
		b = &SynapseSpark{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachine{}
	default:
		b = &Compute{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalComputeSecretsClassification(rawMsg json.RawMessage) (ComputeSecretsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ComputeSecretsClassification
	switch m["computeType"] {
	case string(ComputeTypeAKS):
		b = &AksComputeSecrets{}
	case string(ComputeTypeDatabricks):
		b = &DatabricksComputeSecrets{}
	case string(ComputeTypeVirtualMachine):
		b = &VirtualMachineSecrets{}
	default:
		b = &ComputeSecrets{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataDriftMetricThresholdBaseClassification(rawMsg json.RawMessage) (DataDriftMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataDriftMetricThresholdBaseClassification
	switch m["dataType"] {
	case string(MonitoringFeatureDataTypeCategorical):
		b = &CategoricalDataDriftMetricThreshold{}
	case string(MonitoringFeatureDataTypeNumerical):
		b = &NumericalDataDriftMetricThreshold{}
	default:
		b = &DataDriftMetricThresholdBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataDriftMetricThresholdBaseClassificationArray(rawMsg json.RawMessage) ([]DataDriftMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataDriftMetricThresholdBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataDriftMetricThresholdBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataImportSourceClassification(rawMsg json.RawMessage) (DataImportSourceClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataImportSourceClassification
	switch m["sourceType"] {
	case string(DataImportSourceTypeDatabase):
		b = &DatabaseSource{}
	case string(DataImportSourceTypeFileSystem):
		b = &FileSystemSource{}
	default:
		b = &DataImportSource{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataQualityMetricThresholdBaseClassification(rawMsg json.RawMessage) (DataQualityMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataQualityMetricThresholdBaseClassification
	switch m["dataType"] {
	case string(MonitoringFeatureDataTypeCategorical):
		b = &CategoricalDataQualityMetricThreshold{}
	case string(MonitoringFeatureDataTypeNumerical):
		b = &NumericalDataQualityMetricThreshold{}
	default:
		b = &DataQualityMetricThresholdBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataQualityMetricThresholdBaseClassificationArray(rawMsg json.RawMessage) ([]DataQualityMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataQualityMetricThresholdBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataQualityMetricThresholdBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDataReferenceCredentialClassification(rawMsg json.RawMessage) (DataReferenceCredentialClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataReferenceCredentialClassification
	switch m["credentialType"] {
	case string(DataReferenceCredentialTypeDockerCredentials):
		b = &DockerCredential{}
	case string(DataReferenceCredentialTypeManagedIdentity):
		b = &ManagedIdentityCredential{}
	case string(DataReferenceCredentialTypeNoCredentials):
		b = &AnonymousAccessCredential{}
	case string(DataReferenceCredentialTypeSAS):
		b = &SASCredential{}
	default:
		b = &DataReferenceCredential{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDataVersionBasePropertiesClassification(rawMsg json.RawMessage) (DataVersionBasePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataVersionBasePropertiesClassification
	switch m["dataType"] {
	case string(DataTypeDataImport):
		b = &DataImport{}
	case string(DataTypeMltable):
		b = &MLTableData{}
	case string(DataTypeURIFile):
		b = &URIFileDataVersion{}
	case string(DataTypeURIFolder):
		b = &URIFolderDataVersion{}
	default:
		b = &DataVersionBaseProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatastoreCredentialsClassification(rawMsg json.RawMessage) (DatastoreCredentialsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastoreCredentialsClassification
	switch m["credentialsType"] {
	case string(CredentialsTypeAccountKey):
		b = &AccountKeyDatastoreCredentials{}
	case string(CredentialsTypeCertificate):
		b = &CertificateDatastoreCredentials{}
	case string(CredentialsTypeNone):
		b = &NoneDatastoreCredentials{}
	case string(CredentialsTypeSas):
		b = &SasDatastoreCredentials{}
	case string(CredentialsTypeServicePrincipal):
		b = &ServicePrincipalDatastoreCredentials{}
	default:
		b = &DatastoreCredentials{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatastorePropertiesClassification(rawMsg json.RawMessage) (DatastorePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastorePropertiesClassification
	switch m["datastoreType"] {
	case string(DatastoreTypeAzureBlob):
		b = &AzureBlobDatastore{}
	case string(DatastoreTypeAzureDataLakeGen1):
		b = &AzureDataLakeGen1Datastore{}
	case string(DatastoreTypeAzureDataLakeGen2):
		b = &AzureDataLakeGen2Datastore{}
	case string(DatastoreTypeAzureFile):
		b = &AzureFileDatastore{}
	case string(DatastoreTypeOneLake):
		b = &OneLakeDatastore{}
	default:
		b = &DatastoreProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDatastoreSecretsClassification(rawMsg json.RawMessage) (DatastoreSecretsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DatastoreSecretsClassification
	switch m["secretsType"] {
	case string(SecretsTypeAccountKey):
		b = &AccountKeyDatastoreSecrets{}
	case string(SecretsTypeCertificate):
		b = &CertificateDatastoreSecrets{}
	case string(SecretsTypeSas):
		b = &SasDatastoreSecrets{}
	case string(SecretsTypeServicePrincipal):
		b = &ServicePrincipalDatastoreSecrets{}
	default:
		b = &DatastoreSecrets{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalDistributionConfigurationClassification(rawMsg json.RawMessage) (DistributionConfigurationClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DistributionConfigurationClassification
	switch m["distributionType"] {
	case string(DistributionTypeMpi):
		b = &Mpi{}
	case string(DistributionTypePyTorch):
		b = &PyTorch{}
	case string(DistributionTypeTensorFlow):
		b = &TensorFlow{}
	default:
		b = &DistributionConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEarlyTerminationPolicyClassification(rawMsg json.RawMessage) (EarlyTerminationPolicyClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EarlyTerminationPolicyClassification
	switch m["policyType"] {
	case string(EarlyTerminationPolicyTypeBandit):
		b = &BanditPolicy{}
	case string(EarlyTerminationPolicyTypeMedianStopping):
		b = &MedianStoppingPolicy{}
	case string(EarlyTerminationPolicyTypeTruncationSelection):
		b = &TruncationSelectionPolicy{}
	default:
		b = &EarlyTerminationPolicy{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalForecastHorizonClassification(rawMsg json.RawMessage) (ForecastHorizonClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ForecastHorizonClassification
	switch m["mode"] {
	case string(ForecastHorizonModeAuto):
		b = &AutoForecastHorizon{}
	case string(ForecastHorizonModeCustom):
		b = &CustomForecastHorizon{}
	default:
		b = &ForecastHorizon{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalIdentityConfigurationClassification(rawMsg json.RawMessage) (IdentityConfigurationClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b IdentityConfigurationClassification
	switch m["identityType"] {
	case string(IdentityConfigurationTypeAMLToken):
		b = &AmlToken{}
	case string(IdentityConfigurationTypeManaged):
		b = &ManagedIdentity{}
	case string(IdentityConfigurationTypeUserIdentity):
		b = &UserIdentity{}
	default:
		b = &IdentityConfiguration{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalJobBasePropertiesClassification(rawMsg json.RawMessage) (JobBasePropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobBasePropertiesClassification
	switch m["jobType"] {
	case string(JobTypeAutoML):
		b = &AutoMLJob{}
	case string(JobTypeCommand):
		b = &CommandJob{}
	case string(JobTypePipeline):
		b = &PipelineJob{}
	case string(JobTypeSpark):
		b = &SparkJob{}
	case string(JobTypeSweep):
		b = &SweepJob{}
	default:
		b = &JobBaseProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalJobInputClassification(rawMsg json.RawMessage) (JobInputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobInputClassification
	switch m["jobInputType"] {
	case string(JobInputTypeCustomModel):
		b = &CustomModelJobInput{}
	case string(JobInputTypeLiteral):
		b = &LiteralJobInput{}
	case string(JobInputTypeMlflowModel):
		b = &MLFlowModelJobInput{}
	case string(JobInputTypeMltable):
		b = &MLTableJobInput{}
	case string(JobInputTypeTritonModel):
		b = &TritonModelJobInput{}
	case string(JobInputTypeURIFile):
		b = &URIFileJobInput{}
	case string(JobInputTypeURIFolder):
		b = &URIFolderJobInput{}
	default:
		b = &JobInput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalJobInputClassificationMap(rawMsg json.RawMessage) (map[string]JobInputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobInputClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalJobOutputClassification(rawMsg json.RawMessage) (JobOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobOutputClassification
	switch m["jobOutputType"] {
	case string(JobOutputTypeCustomModel):
		b = &CustomModelJobOutput{}
	case string(JobOutputTypeMlflowModel):
		b = &MLFlowModelJobOutput{}
	case string(JobOutputTypeMltable):
		b = &MLTableJobOutput{}
	case string(JobOutputTypeTritonModel):
		b = &TritonModelJobOutput{}
	case string(JobOutputTypeURIFile):
		b = &URIFileJobOutput{}
	case string(JobOutputTypeURIFolder):
		b = &URIFolderJobOutput{}
	default:
		b = &JobOutput{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalJobOutputClassificationMap(rawMsg json.RawMessage) (map[string]JobOutputClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]JobOutputClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalJobOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalMonitorComputeConfigurationBaseClassification(rawMsg json.RawMessage) (MonitorComputeConfigurationBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MonitorComputeConfigurationBaseClassification
	switch m["computeType"] {
	case string(MonitorComputeTypeServerlessSpark):
		b = &MonitorServerlessSparkCompute{}
	default:
		b = &MonitorComputeConfigurationBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMonitorComputeIdentityBaseClassification(rawMsg json.RawMessage) (MonitorComputeIdentityBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MonitorComputeIdentityBaseClassification
	switch m["computeIdentityType"] {
	case string(MonitorComputeIdentityTypeAmlToken):
		b = &AmlTokenComputeIdentity{}
	case string(MonitorComputeIdentityTypeManagedIdentity):
		b = &ManagedComputeIdentity{}
	default:
		b = &MonitorComputeIdentityBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMonitoringFeatureFilterBaseClassification(rawMsg json.RawMessage) (MonitoringFeatureFilterBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MonitoringFeatureFilterBaseClassification
	switch m["filterType"] {
	case string(MonitoringFeatureFilterTypeAllFeatures):
		b = &AllFeatures{}
	case string(MonitoringFeatureFilterTypeFeatureSubset):
		b = &FeatureSubset{}
	case string(MonitoringFeatureFilterTypeTopNByAttribution):
		b = &TopNFeaturesByAttribution{}
	default:
		b = &MonitoringFeatureFilterBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMonitoringInputDataBaseClassification(rawMsg json.RawMessage) (MonitoringInputDataBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MonitoringInputDataBaseClassification
	switch m["inputDataType"] {
	case string(MonitoringInputDataTypeFixed):
		b = &FixedInputData{}
	case string(MonitoringInputDataTypeRolling):
		b = &RollingInputData{}
	case string(MonitoringInputDataTypeStatic):
		b = &StaticInputData{}
	default:
		b = &MonitoringInputDataBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMonitoringInputDataBaseClassificationArray(rawMsg json.RawMessage) ([]MonitoringInputDataBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MonitoringInputDataBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMonitoringInputDataBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMonitoringInputDataBaseClassificationMap(rawMsg json.RawMessage) (map[string]MonitoringInputDataBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]MonitoringInputDataBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalMonitoringInputDataBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalMonitoringSignalBaseClassification(rawMsg json.RawMessage) (MonitoringSignalBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MonitoringSignalBaseClassification
	switch m["signalType"] {
	case string(MonitoringSignalTypeCustom):
		b = &CustomMonitoringSignal{}
	case string(MonitoringSignalTypeDataDrift):
		b = &DataDriftMonitoringSignal{}
	case string(MonitoringSignalTypeDataQuality):
		b = &DataQualityMonitoringSignal{}
	case string(MonitoringSignalTypeFeatureAttributionDrift):
		b = &FeatureAttributionDriftMonitoringSignal{}
	case string(MonitoringSignalTypePredictionDrift):
		b = &PredictionDriftMonitoringSignal{}
	default:
		b = &MonitoringSignalBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalMonitoringSignalBaseClassificationMap(rawMsg json.RawMessage) (map[string]MonitoringSignalBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]MonitoringSignalBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalMonitoringSignalBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalNCrossValidationsClassification(rawMsg json.RawMessage) (NCrossValidationsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NCrossValidationsClassification
	switch m["mode"] {
	case string(NCrossValidationsModeAuto):
		b = &AutoNCrossValidations{}
	case string(NCrossValidationsModeCustom):
		b = &CustomNCrossValidations{}
	default:
		b = &NCrossValidations{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNodesClassification(rawMsg json.RawMessage) (NodesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NodesClassification
	switch m["nodesValueType"] {
	case string(NodesValueTypeAll):
		b = &AllNodes{}
	default:
		b = &Nodes{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOneLakeArtifactClassification(rawMsg json.RawMessage) (OneLakeArtifactClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OneLakeArtifactClassification
	switch m["artifactType"] {
	case string(OneLakeArtifactTypeLakeHouse):
		b = &LakeHouseArtifact{}
	default:
		b = &OneLakeArtifact{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOnlineDeploymentPropertiesClassification(rawMsg json.RawMessage) (OnlineDeploymentPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OnlineDeploymentPropertiesClassification
	switch m["endpointComputeType"] {
	case string(EndpointComputeTypeKubernetes):
		b = &KubernetesOnlineDeployment{}
	case string(EndpointComputeTypeManaged):
		b = &ManagedOnlineDeployment{}
	default:
		b = &OnlineDeploymentProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOnlineScaleSettingsClassification(rawMsg json.RawMessage) (OnlineScaleSettingsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OnlineScaleSettingsClassification
	switch m["scaleType"] {
	case string(ScaleTypeDefault):
		b = &DefaultScaleSettings{}
	case string(ScaleTypeTargetUtilization):
		b = &TargetUtilizationScaleSettings{}
	default:
		b = &OnlineScaleSettings{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOutboundRuleClassification(rawMsg json.RawMessage) (OutboundRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OutboundRuleClassification
	switch m["type"] {
	case string(RuleTypeFQDN):
		b = &FqdnOutboundRule{}
	case string(RuleTypePrivateEndpoint):
		b = &PrivateEndpointOutboundRule{}
	case string(RuleTypeServiceTag):
		b = &ServiceTagOutboundRule{}
	default:
		b = &OutboundRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOutboundRuleClassificationMap(rawMsg json.RawMessage) (map[string]OutboundRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]OutboundRuleClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalOutboundRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalPendingUploadCredentialDtoClassification(rawMsg json.RawMessage) (PendingUploadCredentialDtoClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PendingUploadCredentialDtoClassification
	switch m["credentialType"] {
	case string(PendingUploadCredentialTypeSAS):
		b = &SASCredentialDto{}
	default:
		b = &PendingUploadCredentialDto{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPredictionDriftMetricThresholdBaseClassification(rawMsg json.RawMessage) (PredictionDriftMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PredictionDriftMetricThresholdBaseClassification
	switch m["dataType"] {
	case string(MonitoringFeatureDataTypeCategorical):
		b = &CategoricalPredictionDriftMetricThreshold{}
	case string(MonitoringFeatureDataTypeNumerical):
		b = &NumericalPredictionDriftMetricThreshold{}
	default:
		b = &PredictionDriftMetricThresholdBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPredictionDriftMetricThresholdBaseClassificationArray(rawMsg json.RawMessage) ([]PredictionDriftMetricThresholdBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PredictionDriftMetricThresholdBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPredictionDriftMetricThresholdBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSamplingAlgorithmClassification(rawMsg json.RawMessage) (SamplingAlgorithmClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SamplingAlgorithmClassification
	switch m["samplingAlgorithmType"] {
	case string(SamplingAlgorithmTypeBayesian):
		b = &BayesianSamplingAlgorithm{}
	case string(SamplingAlgorithmTypeGrid):
		b = &GridSamplingAlgorithm{}
	case string(SamplingAlgorithmTypeRandom):
		b = &RandomSamplingAlgorithm{}
	default:
		b = &SamplingAlgorithm{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalScheduleActionBaseClassification(rawMsg json.RawMessage) (ScheduleActionBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ScheduleActionBaseClassification
	switch m["actionType"] {
	case string(ScheduleActionTypeCreateJob):
		b = &JobScheduleAction{}
	case string(ScheduleActionTypeCreateMonitor):
		b = &CreateMonitorAction{}
	case string(ScheduleActionTypeImportData):
		b = &ImportDataAction{}
	case string(ScheduleActionTypeInvokeBatchEndpoint):
		b = &EndpointScheduleAction{}
	default:
		b = &ScheduleActionBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSeasonalityClassification(rawMsg json.RawMessage) (SeasonalityClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SeasonalityClassification
	switch m["mode"] {
	case string(SeasonalityModeAuto):
		b = &AutoSeasonality{}
	case string(SeasonalityModeCustom):
		b = &CustomSeasonality{}
	default:
		b = &Seasonality{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSparkJobEntryClassification(rawMsg json.RawMessage) (SparkJobEntryClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SparkJobEntryClassification
	switch m["sparkJobEntryType"] {
	case string(SparkJobEntryTypeSparkJobPythonEntry):
		b = &SparkJobPythonEntry{}
	case string(SparkJobEntryTypeSparkJobScalaEntry):
		b = &SparkJobScalaEntry{}
	default:
		b = &SparkJobEntry{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTargetLagsClassification(rawMsg json.RawMessage) (TargetLagsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetLagsClassification
	switch m["mode"] {
	case string(TargetLagsModeAuto):
		b = &AutoTargetLags{}
	case string(TargetLagsModeCustom):
		b = &CustomTargetLags{}
	default:
		b = &TargetLags{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTargetRollingWindowSizeClassification(rawMsg json.RawMessage) (TargetRollingWindowSizeClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TargetRollingWindowSizeClassification
	switch m["mode"] {
	case string(TargetRollingWindowSizeModeAuto):
		b = &AutoTargetRollingWindowSize{}
	case string(TargetRollingWindowSizeModeCustom):
		b = &CustomTargetRollingWindowSize{}
	default:
		b = &TargetRollingWindowSize{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTriggerBaseClassification(rawMsg json.RawMessage) (TriggerBaseClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerBaseClassification
	switch m["triggerType"] {
	case string(TriggerTypeCron):
		b = &CronTrigger{}
	case string(TriggerTypeRecurrence):
		b = &RecurrenceTrigger{}
	default:
		b = &TriggerBase{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWebhookClassification(rawMsg json.RawMessage) (WebhookClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WebhookClassification
	switch m["webhookType"] {
	case string(WebhookTypeAzureDevOps):
		b = &AzureDevOpsWebhook{}
	default:
		b = &Webhook{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWebhookClassificationMap(rawMsg json.RawMessage) (map[string]WebhookClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]WebhookClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalWebhookClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalWorkspaceConnectionPropertiesV2Classification(rawMsg json.RawMessage) (WorkspaceConnectionPropertiesV2Classification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkspaceConnectionPropertiesV2Classification
	switch m["authType"] {
	case string(ConnectionAuthTypeAAD):
		b = &AADAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeAccessKey):
		b = &AccessKeyAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeAccountKey):
		b = &AccountKeyAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeAPIKey):
		b = &APIKeyAuthWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeCustomKeys):
		b = &CustomKeysWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeManagedIdentity):
		b = &ManagedIdentityAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeNone):
		b = &NoneAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeOAuth2):
		b = &OAuth2AuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypePAT):
		b = &PATAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeSAS):
		b = &SASAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeServicePrincipal):
		b = &ServicePrincipalAuthTypeWorkspaceConnectionProperties{}
	case string(ConnectionAuthTypeUsernamePassword):
		b = &UsernamePasswordAuthTypeWorkspaceConnectionProperties{}
	default:
		b = &WorkspaceConnectionPropertiesV2{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
