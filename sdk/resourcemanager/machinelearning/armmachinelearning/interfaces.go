//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

// AssetReferenceBaseClassification provides polymorphic access to related types.
// Call the interface's GetAssetReferenceBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AssetReferenceBase, *DataPathAssetReference, *IDAssetReference, *OutputPathAssetReference
type AssetReferenceBaseClassification interface {
	// GetAssetReferenceBase returns the AssetReferenceBase content of the underlying type.
	GetAssetReferenceBase() *AssetReferenceBase
}

// AutoMLVerticalClassification provides polymorphic access to related types.
// Call the interface's GetAutoMLVertical() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoMLVertical, *Classification, *Forecasting, *ImageClassification, *ImageClassificationMultilabel, *ImageInstanceSegmentation,
// - *ImageObjectDetection, *Regression, *TextClassification, *TextClassificationMultilabel, *TextNer
type AutoMLVerticalClassification interface {
	// GetAutoMLVertical returns the AutoMLVertical content of the underlying type.
	GetAutoMLVertical() *AutoMLVertical
}

// BatchDeploymentConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetBatchDeploymentConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BatchDeploymentConfiguration, *BatchPipelineComponentDeploymentConfiguration
type BatchDeploymentConfigurationClassification interface {
	// GetBatchDeploymentConfiguration returns the BatchDeploymentConfiguration content of the underlying type.
	GetBatchDeploymentConfiguration() *BatchDeploymentConfiguration
}

// ComputeClassification provides polymorphic access to related types.
// Call the interface's GetCompute() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AKS, *AmlCompute, *Compute, *ComputeInstance, *DataFactory, *DataLakeAnalytics, *Databricks, *HDInsight, *Kubernetes,
// - *SynapseSpark, *VirtualMachine
type ComputeClassification interface {
	// GetCompute returns the Compute content of the underlying type.
	GetCompute() *Compute
}

// ComputeSecretsClassification provides polymorphic access to related types.
// Call the interface's GetComputeSecrets() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AksComputeSecrets, *ComputeSecrets, *DatabricksComputeSecrets, *VirtualMachineSecrets
type ComputeSecretsClassification interface {
	// GetComputeSecrets returns the ComputeSecrets content of the underlying type.
	GetComputeSecrets() *ComputeSecrets
}

// DataDriftMetricThresholdBaseClassification provides polymorphic access to related types.
// Call the interface's GetDataDriftMetricThresholdBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CategoricalDataDriftMetricThreshold, *DataDriftMetricThresholdBase, *NumericalDataDriftMetricThreshold
type DataDriftMetricThresholdBaseClassification interface {
	// GetDataDriftMetricThresholdBase returns the DataDriftMetricThresholdBase content of the underlying type.
	GetDataDriftMetricThresholdBase() *DataDriftMetricThresholdBase
}

// DataQualityMetricThresholdBaseClassification provides polymorphic access to related types.
// Call the interface's GetDataQualityMetricThresholdBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CategoricalDataQualityMetricThreshold, *DataQualityMetricThresholdBase, *NumericalDataQualityMetricThreshold
type DataQualityMetricThresholdBaseClassification interface {
	// GetDataQualityMetricThresholdBase returns the DataQualityMetricThresholdBase content of the underlying type.
	GetDataQualityMetricThresholdBase() *DataQualityMetricThresholdBase
}

// DataReferenceCredentialClassification provides polymorphic access to related types.
// Call the interface's GetDataReferenceCredential() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AnonymousAccessCredential, *DataReferenceCredential, *DockerCredential, *ManagedIdentityCredential, *SASCredential
type DataReferenceCredentialClassification interface {
	// GetDataReferenceCredential returns the DataReferenceCredential content of the underlying type.
	GetDataReferenceCredential() *DataReferenceCredential
}

// DataVersionBasePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetDataVersionBaseProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DataVersionBaseProperties, *MLTableData, *URIFileDataVersion, *URIFolderDataVersion
type DataVersionBasePropertiesClassification interface {
	// GetDataVersionBaseProperties returns the DataVersionBaseProperties content of the underlying type.
	GetDataVersionBaseProperties() *DataVersionBaseProperties
}

// DatastoreCredentialsClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreCredentials() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccountKeyDatastoreCredentials, *CertificateDatastoreCredentials, *DatastoreCredentials, *NoneDatastoreCredentials,
// - *SasDatastoreCredentials, *ServicePrincipalDatastoreCredentials
type DatastoreCredentialsClassification interface {
	// GetDatastoreCredentials returns the DatastoreCredentials content of the underlying type.
	GetDatastoreCredentials() *DatastoreCredentials
}

// DatastorePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureBlobDatastore, *AzureDataLakeGen1Datastore, *AzureDataLakeGen2Datastore, *AzureFileDatastore, *DatastoreProperties,
// - *OneLakeDatastore
type DatastorePropertiesClassification interface {
	// GetDatastoreProperties returns the DatastoreProperties content of the underlying type.
	GetDatastoreProperties() *DatastoreProperties
}

// DatastoreSecretsClassification provides polymorphic access to related types.
// Call the interface's GetDatastoreSecrets() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AccountKeyDatastoreSecrets, *CertificateDatastoreSecrets, *DatastoreSecrets, *SasDatastoreSecrets, *ServicePrincipalDatastoreSecrets
type DatastoreSecretsClassification interface {
	// GetDatastoreSecrets returns the DatastoreSecrets content of the underlying type.
	GetDatastoreSecrets() *DatastoreSecrets
}

// DistributionConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetDistributionConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DistributionConfiguration, *Mpi, *PyTorch, *TensorFlow
type DistributionConfigurationClassification interface {
	// GetDistributionConfiguration returns the DistributionConfiguration content of the underlying type.
	GetDistributionConfiguration() *DistributionConfiguration
}

// EarlyTerminationPolicyClassification provides polymorphic access to related types.
// Call the interface's GetEarlyTerminationPolicy() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BanditPolicy, *EarlyTerminationPolicy, *MedianStoppingPolicy, *TruncationSelectionPolicy
type EarlyTerminationPolicyClassification interface {
	// GetEarlyTerminationPolicy returns the EarlyTerminationPolicy content of the underlying type.
	GetEarlyTerminationPolicy() *EarlyTerminationPolicy
}

// FineTuningVerticalClassification provides polymorphic access to related types.
// Call the interface's GetFineTuningVertical() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureOpenAiFineTuning, *CustomModelFineTuning, *FineTuningVertical
type FineTuningVerticalClassification interface {
	// GetFineTuningVertical returns the FineTuningVertical content of the underlying type.
	GetFineTuningVertical() *FineTuningVertical
}

// ForecastHorizonClassification provides polymorphic access to related types.
// Call the interface's GetForecastHorizon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoForecastHorizon, *CustomForecastHorizon, *ForecastHorizon
type ForecastHorizonClassification interface {
	// GetForecastHorizon returns the ForecastHorizon content of the underlying type.
	GetForecastHorizon() *ForecastHorizon
}

// IdentityConfigurationClassification provides polymorphic access to related types.
// Call the interface's GetIdentityConfiguration() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmlToken, *IdentityConfiguration, *ManagedIdentity, *UserIdentity
type IdentityConfigurationClassification interface {
	// GetIdentityConfiguration returns the IdentityConfiguration content of the underlying type.
	GetIdentityConfiguration() *IdentityConfiguration
}

// JobBasePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetJobBaseProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoMLJob, *CommandJob, *FineTuningJob, *JobBaseProperties, *PipelineJob, *SparkJob, *SweepJob
type JobBasePropertiesClassification interface {
	// GetJobBaseProperties returns the JobBaseProperties content of the underlying type.
	GetJobBaseProperties() *JobBaseProperties
}

// JobInputClassification provides polymorphic access to related types.
// Call the interface's GetJobInput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomModelJobInput, *JobInput, *LiteralJobInput, *MLFlowModelJobInput, *MLTableJobInput, *TritonModelJobInput, *URIFileJobInput,
// - *URIFolderJobInput
type JobInputClassification interface {
	// GetJobInput returns the JobInput content of the underlying type.
	GetJobInput() *JobInput
}

// JobLimitsClassification provides polymorphic access to related types.
// Call the interface's GetJobLimits() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CommandJobLimits, *JobLimits, *SweepJobLimits
type JobLimitsClassification interface {
	// GetJobLimits returns the JobLimits content of the underlying type.
	GetJobLimits() *JobLimits
}

// JobOutputClassification provides polymorphic access to related types.
// Call the interface's GetJobOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomModelJobOutput, *JobOutput, *MLFlowModelJobOutput, *MLTableJobOutput, *TritonModelJobOutput, *URIFileJobOutput,
// - *URIFolderJobOutput
type JobOutputClassification interface {
	// GetJobOutput returns the JobOutput content of the underlying type.
	GetJobOutput() *JobOutput
}

// MonitorComputeConfigurationBaseClassification provides polymorphic access to related types.
// Call the interface's GetMonitorComputeConfigurationBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MonitorComputeConfigurationBase, *MonitorServerlessSparkCompute
type MonitorComputeConfigurationBaseClassification interface {
	// GetMonitorComputeConfigurationBase returns the MonitorComputeConfigurationBase content of the underlying type.
	GetMonitorComputeConfigurationBase() *MonitorComputeConfigurationBase
}

// MonitorComputeIdentityBaseClassification provides polymorphic access to related types.
// Call the interface's GetMonitorComputeIdentityBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AmlTokenComputeIdentity, *ManagedComputeIdentity, *MonitorComputeIdentityBase
type MonitorComputeIdentityBaseClassification interface {
	// GetMonitorComputeIdentityBase returns the MonitorComputeIdentityBase content of the underlying type.
	GetMonitorComputeIdentityBase() *MonitorComputeIdentityBase
}

// MonitoringFeatureFilterBaseClassification provides polymorphic access to related types.
// Call the interface's GetMonitoringFeatureFilterBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AllFeatures, *FeatureSubset, *MonitoringFeatureFilterBase, *TopNFeaturesByAttribution
type MonitoringFeatureFilterBaseClassification interface {
	// GetMonitoringFeatureFilterBase returns the MonitoringFeatureFilterBase content of the underlying type.
	GetMonitoringFeatureFilterBase() *MonitoringFeatureFilterBase
}

// MonitoringInputDataBaseClassification provides polymorphic access to related types.
// Call the interface's GetMonitoringInputDataBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FixedInputData, *MonitoringInputDataBase, *RollingInputData, *StaticInputData
type MonitoringInputDataBaseClassification interface {
	// GetMonitoringInputDataBase returns the MonitoringInputDataBase content of the underlying type.
	GetMonitoringInputDataBase() *MonitoringInputDataBase
}

// MonitoringSignalBaseClassification provides polymorphic access to related types.
// Call the interface's GetMonitoringSignalBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CustomMonitoringSignal, *DataDriftMonitoringSignal, *DataQualityMonitoringSignal, *FeatureAttributionDriftMonitoringSignal,
// - *MonitoringSignalBase, *PredictionDriftMonitoringSignal
type MonitoringSignalBaseClassification interface {
	// GetMonitoringSignalBase returns the MonitoringSignalBase content of the underlying type.
	GetMonitoringSignalBase() *MonitoringSignalBase
}

// NCrossValidationsClassification provides polymorphic access to related types.
// Call the interface's GetNCrossValidations() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoNCrossValidations, *CustomNCrossValidations, *NCrossValidations
type NCrossValidationsClassification interface {
	// GetNCrossValidations returns the NCrossValidations content of the underlying type.
	GetNCrossValidations() *NCrossValidations
}

// NodesClassification provides polymorphic access to related types.
// Call the interface's GetNodes() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AllNodes, *Nodes
type NodesClassification interface {
	// GetNodes returns the Nodes content of the underlying type.
	GetNodes() *Nodes
}

// OneLakeArtifactClassification provides polymorphic access to related types.
// Call the interface's GetOneLakeArtifact() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LakeHouseArtifact, *OneLakeArtifact
type OneLakeArtifactClassification interface {
	// GetOneLakeArtifact returns the OneLakeArtifact content of the underlying type.
	GetOneLakeArtifact() *OneLakeArtifact
}

// OnlineDeploymentPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetOnlineDeploymentProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *KubernetesOnlineDeployment, *ManagedOnlineDeployment, *OnlineDeploymentProperties
type OnlineDeploymentPropertiesClassification interface {
	// GetOnlineDeploymentProperties returns the OnlineDeploymentProperties content of the underlying type.
	GetOnlineDeploymentProperties() *OnlineDeploymentProperties
}

// OnlineScaleSettingsClassification provides polymorphic access to related types.
// Call the interface's GetOnlineScaleSettings() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DefaultScaleSettings, *OnlineScaleSettings, *TargetUtilizationScaleSettings
type OnlineScaleSettingsClassification interface {
	// GetOnlineScaleSettings returns the OnlineScaleSettings content of the underlying type.
	GetOnlineScaleSettings() *OnlineScaleSettings
}

// OutboundRuleClassification provides polymorphic access to related types.
// Call the interface's GetOutboundRule() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FqdnOutboundRule, *OutboundRule, *PrivateEndpointOutboundRule, *ServiceTagOutboundRule
type OutboundRuleClassification interface {
	// GetOutboundRule returns the OutboundRule content of the underlying type.
	GetOutboundRule() *OutboundRule
}

// PendingUploadCredentialDtoClassification provides polymorphic access to related types.
// Call the interface's GetPendingUploadCredentialDto() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PendingUploadCredentialDto, *SASCredentialDto
type PendingUploadCredentialDtoClassification interface {
	// GetPendingUploadCredentialDto returns the PendingUploadCredentialDto content of the underlying type.
	GetPendingUploadCredentialDto() *PendingUploadCredentialDto
}

// PredictionDriftMetricThresholdBaseClassification provides polymorphic access to related types.
// Call the interface's GetPredictionDriftMetricThresholdBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CategoricalPredictionDriftMetricThreshold, *NumericalPredictionDriftMetricThreshold, *PredictionDriftMetricThresholdBase
type PredictionDriftMetricThresholdBaseClassification interface {
	// GetPredictionDriftMetricThresholdBase returns the PredictionDriftMetricThresholdBase content of the underlying type.
	GetPredictionDriftMetricThresholdBase() *PredictionDriftMetricThresholdBase
}

// SamplingAlgorithmClassification provides polymorphic access to related types.
// Call the interface's GetSamplingAlgorithm() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *BayesianSamplingAlgorithm, *GridSamplingAlgorithm, *RandomSamplingAlgorithm, *SamplingAlgorithm
type SamplingAlgorithmClassification interface {
	// GetSamplingAlgorithm returns the SamplingAlgorithm content of the underlying type.
	GetSamplingAlgorithm() *SamplingAlgorithm
}

// ScheduleActionBaseClassification provides polymorphic access to related types.
// Call the interface's GetScheduleActionBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CreateMonitorAction, *EndpointScheduleAction, *JobScheduleAction, *ScheduleActionBase
type ScheduleActionBaseClassification interface {
	// GetScheduleActionBase returns the ScheduleActionBase content of the underlying type.
	GetScheduleActionBase() *ScheduleActionBase
}

// SeasonalityClassification provides polymorphic access to related types.
// Call the interface's GetSeasonality() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoSeasonality, *CustomSeasonality, *Seasonality
type SeasonalityClassification interface {
	// GetSeasonality returns the Seasonality content of the underlying type.
	GetSeasonality() *Seasonality
}

// SparkJobEntryClassification provides polymorphic access to related types.
// Call the interface's GetSparkJobEntry() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *SparkJobEntry, *SparkJobPythonEntry, *SparkJobScalaEntry
type SparkJobEntryClassification interface {
	// GetSparkJobEntry returns the SparkJobEntry content of the underlying type.
	GetSparkJobEntry() *SparkJobEntry
}

// TargetLagsClassification provides polymorphic access to related types.
// Call the interface's GetTargetLags() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoTargetLags, *CustomTargetLags, *TargetLags
type TargetLagsClassification interface {
	// GetTargetLags returns the TargetLags content of the underlying type.
	GetTargetLags() *TargetLags
}

// TargetRollingWindowSizeClassification provides polymorphic access to related types.
// Call the interface's GetTargetRollingWindowSize() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutoTargetRollingWindowSize, *CustomTargetRollingWindowSize, *TargetRollingWindowSize
type TargetRollingWindowSizeClassification interface {
	// GetTargetRollingWindowSize returns the TargetRollingWindowSize content of the underlying type.
	GetTargetRollingWindowSize() *TargetRollingWindowSize
}

// TriggerBaseClassification provides polymorphic access to related types.
// Call the interface's GetTriggerBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CronTrigger, *RecurrenceTrigger, *TriggerBase
type TriggerBaseClassification interface {
	// GetTriggerBase returns the TriggerBase content of the underlying type.
	GetTriggerBase() *TriggerBase
}

// WebhookClassification provides polymorphic access to related types.
// Call the interface's GetWebhook() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDevOpsWebhook, *Webhook
type WebhookClassification interface {
	// GetWebhook returns the Webhook content of the underlying type.
	GetWebhook() *Webhook
}

// WorkspaceConnectionPropertiesV2Classification provides polymorphic access to related types.
// Call the interface's GetWorkspaceConnectionPropertiesV2() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AADAuthTypeWorkspaceConnectionProperties, *APIKeyAuthWorkspaceConnectionProperties, *AccessKeyAuthTypeWorkspaceConnectionProperties,
// - *AccountKeyAuthTypeWorkspaceConnectionProperties, *CustomKeysWorkspaceConnectionProperties, *ManagedIdentityAuthTypeWorkspaceConnectionProperties,
// - *NoneAuthTypeWorkspaceConnectionProperties, *OAuth2AuthTypeWorkspaceConnectionProperties, *PATAuthTypeWorkspaceConnectionProperties,
// - *SASAuthTypeWorkspaceConnectionProperties, *ServicePrincipalAuthTypeWorkspaceConnectionProperties, *UsernamePasswordAuthTypeWorkspaceConnectionProperties,
// - *WorkspaceConnectionPropertiesV2
type WorkspaceConnectionPropertiesV2Classification interface {
	// GetWorkspaceConnectionPropertiesV2 returns the WorkspaceConnectionPropertiesV2 content of the underlying type.
	GetWorkspaceConnectionPropertiesV2() *WorkspaceConnectionPropertiesV2
}
