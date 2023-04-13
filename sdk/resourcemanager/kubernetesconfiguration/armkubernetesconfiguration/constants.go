//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkubernetesconfiguration

const (
	moduleName    = "armkubernetesconfiguration"
	moduleVersion = "v2.0.0"
)

// AKSIdentityType - The identity type.
type AKSIdentityType string

const (
	AKSIdentityTypeSystemAssigned AKSIdentityType = "SystemAssigned"
	AKSIdentityTypeUserAssigned   AKSIdentityType = "UserAssigned"
)

// PossibleAKSIdentityTypeValues returns the possible values for the AKSIdentityType const type.
func PossibleAKSIdentityTypeValues() []AKSIdentityType {
	return []AKSIdentityType{
		AKSIdentityTypeSystemAssigned,
		AKSIdentityTypeUserAssigned,
	}
}

// ComplianceStateType - The compliance state of the configuration.
type ComplianceStateType string

const (
	ComplianceStateTypeCompliant    ComplianceStateType = "Compliant"
	ComplianceStateTypeFailed       ComplianceStateType = "Failed"
	ComplianceStateTypeInstalled    ComplianceStateType = "Installed"
	ComplianceStateTypeNoncompliant ComplianceStateType = "Noncompliant"
	ComplianceStateTypePending      ComplianceStateType = "Pending"
)

// PossibleComplianceStateTypeValues returns the possible values for the ComplianceStateType const type.
func PossibleComplianceStateTypeValues() []ComplianceStateType {
	return []ComplianceStateType{
		ComplianceStateTypeCompliant,
		ComplianceStateTypeFailed,
		ComplianceStateTypeInstalled,
		ComplianceStateTypeNoncompliant,
		ComplianceStateTypePending,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
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

// FluxComplianceState - Compliance state of the cluster object.
type FluxComplianceState string

const (
	FluxComplianceStateCompliant    FluxComplianceState = "Compliant"
	FluxComplianceStateNonCompliant FluxComplianceState = "Non-Compliant"
	FluxComplianceStatePending      FluxComplianceState = "Pending"
	FluxComplianceStateSuspended    FluxComplianceState = "Suspended"
	FluxComplianceStateUnknown      FluxComplianceState = "Unknown"
)

// PossibleFluxComplianceStateValues returns the possible values for the FluxComplianceState const type.
func PossibleFluxComplianceStateValues() []FluxComplianceState {
	return []FluxComplianceState{
		FluxComplianceStateCompliant,
		FluxComplianceStateNonCompliant,
		FluxComplianceStatePending,
		FluxComplianceStateSuspended,
		FluxComplianceStateUnknown,
	}
}

type KubernetesClusterResourceName string

const (
	KubernetesClusterResourceNameConnectedClusters   KubernetesClusterResourceName = "connectedClusters"
	KubernetesClusterResourceNameManagedClusters     KubernetesClusterResourceName = "managedClusters"
	KubernetesClusterResourceNameProvisionedClusters KubernetesClusterResourceName = "provisionedClusters"
)

// PossibleKubernetesClusterResourceNameValues returns the possible values for the KubernetesClusterResourceName const type.
func PossibleKubernetesClusterResourceNameValues() []KubernetesClusterResourceName {
	return []KubernetesClusterResourceName{
		KubernetesClusterResourceNameConnectedClusters,
		KubernetesClusterResourceNameManagedClusters,
		KubernetesClusterResourceNameProvisionedClusters,
	}
}

type KubernetesClusterResourceProviderName string

const (
	KubernetesClusterResourceProviderNameMicrosoftContainerService       KubernetesClusterResourceProviderName = "Microsoft.ContainerService"
	KubernetesClusterResourceProviderNameMicrosoftHybridContainerService KubernetesClusterResourceProviderName = "Microsoft.HybridContainerService"
	KubernetesClusterResourceProviderNameMicrosoftKubernetes             KubernetesClusterResourceProviderName = "Microsoft.Kubernetes"
)

// PossibleKubernetesClusterResourceProviderNameValues returns the possible values for the KubernetesClusterResourceProviderName const type.
func PossibleKubernetesClusterResourceProviderNameValues() []KubernetesClusterResourceProviderName {
	return []KubernetesClusterResourceProviderName{
		KubernetesClusterResourceProviderNameMicrosoftContainerService,
		KubernetesClusterResourceProviderNameMicrosoftHybridContainerService,
		KubernetesClusterResourceProviderNameMicrosoftKubernetes,
	}
}

// KustomizationValidationType - Specify whether to validate the Kubernetes objects referenced in the Kustomization before
// applying them to the cluster.
type KustomizationValidationType string

const (
	KustomizationValidationTypeClient KustomizationValidationType = "client"
	KustomizationValidationTypeNone   KustomizationValidationType = "none"
	KustomizationValidationTypeServer KustomizationValidationType = "server"
)

// PossibleKustomizationValidationTypeValues returns the possible values for the KustomizationValidationType const type.
func PossibleKustomizationValidationTypeValues() []KustomizationValidationType {
	return []KustomizationValidationType{
		KustomizationValidationTypeClient,
		KustomizationValidationTypeNone,
		KustomizationValidationTypeServer,
	}
}

// LevelType - Level of the status.
type LevelType string

const (
	LevelTypeError       LevelType = "Error"
	LevelTypeInformation LevelType = "Information"
	LevelTypeWarning     LevelType = "Warning"
)

// PossibleLevelTypeValues returns the possible values for the LevelType const type.
func PossibleLevelTypeValues() []LevelType {
	return []LevelType{
		LevelTypeError,
		LevelTypeInformation,
		LevelTypeWarning,
	}
}

// MessageLevelType - Level of the message.
type MessageLevelType string

const (
	MessageLevelTypeError       MessageLevelType = "Error"
	MessageLevelTypeInformation MessageLevelType = "Information"
	MessageLevelTypeWarning     MessageLevelType = "Warning"
)

// PossibleMessageLevelTypeValues returns the possible values for the MessageLevelType const type.
func PossibleMessageLevelTypeValues() []MessageLevelType {
	return []MessageLevelType{
		MessageLevelTypeError,
		MessageLevelTypeInformation,
		MessageLevelTypeWarning,
	}
}

// OperatorScopeType - Scope at which the operator will be installed.
type OperatorScopeType string

const (
	OperatorScopeTypeCluster   OperatorScopeType = "cluster"
	OperatorScopeTypeNamespace OperatorScopeType = "namespace"
)

// PossibleOperatorScopeTypeValues returns the possible values for the OperatorScopeType const type.
func PossibleOperatorScopeTypeValues() []OperatorScopeType {
	return []OperatorScopeType{
		OperatorScopeTypeCluster,
		OperatorScopeTypeNamespace,
	}
}

// OperatorType - Type of the operator
type OperatorType string

const (
	OperatorTypeFlux OperatorType = "Flux"
)

// PossibleOperatorTypeValues returns the possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{
		OperatorTypeFlux,
	}
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ProvisioningStateType - The provisioning state of the resource provider.
type ProvisioningStateType string

const (
	ProvisioningStateTypeAccepted  ProvisioningStateType = "Accepted"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeRunning   ProvisioningStateType = "Running"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
)

// PossibleProvisioningStateTypeValues returns the possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{
		ProvisioningStateTypeAccepted,
		ProvisioningStateTypeDeleting,
		ProvisioningStateTypeFailed,
		ProvisioningStateTypeRunning,
		ProvisioningStateTypeSucceeded,
	}
}

// ScopeType - Scope at which the configuration will be installed.
type ScopeType string

const (
	ScopeTypeCluster   ScopeType = "cluster"
	ScopeTypeNamespace ScopeType = "namespace"
)

// PossibleScopeTypeValues returns the possible values for the ScopeType const type.
func PossibleScopeTypeValues() []ScopeType {
	return []ScopeType{
		ScopeTypeCluster,
		ScopeTypeNamespace,
	}
}

// SourceKindType - Source Kind to pull the configuration data from.
type SourceKindType string

const (
	SourceKindTypeAzureBlob     SourceKindType = "AzureBlob"
	SourceKindTypeBucket        SourceKindType = "Bucket"
	SourceKindTypeGitRepository SourceKindType = "GitRepository"
)

// PossibleSourceKindTypeValues returns the possible values for the SourceKindType const type.
func PossibleSourceKindTypeValues() []SourceKindType {
	return []SourceKindType{
		SourceKindTypeAzureBlob,
		SourceKindTypeBucket,
		SourceKindTypeGitRepository,
	}
}
