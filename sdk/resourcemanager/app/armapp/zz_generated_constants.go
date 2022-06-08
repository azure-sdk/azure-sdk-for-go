//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapp

const (
	moduleName    = "armapp"
	moduleVersion = "v0.1.0"
)

// AccessMode - Access mode for storage
type AccessMode string

const (
	AccessModeReadOnly  AccessMode = "ReadOnly"
	AccessModeReadWrite AccessMode = "ReadWrite"
)

// PossibleAccessModeValues returns the possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{
		AccessModeReadOnly,
		AccessModeReadWrite,
	}
}

// ActiveRevisionsMode - ActiveRevisionsMode controls how active revisions are handled for the Container app:Multiple: multiple
// revisions can be active.Single: Only one revision can be active at a time. Revision weights can
// not be used in this mode. If no value if provided, this is the default.
type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple ActiveRevisionsMode = "Multiple"
	ActiveRevisionsModeSingle   ActiveRevisionsMode = "Single"
)

// PossibleActiveRevisionsModeValues returns the possible values for the ActiveRevisionsMode const type.
func PossibleActiveRevisionsModeValues() []ActiveRevisionsMode {
	return []ActiveRevisionsMode{
		ActiveRevisionsModeMultiple,
		ActiveRevisionsModeSingle,
	}
}

// AppProtocol - Tells Dapr which protocol your application is using. Valid options are http and grpc. Default is http
type AppProtocol string

const (
	AppProtocolGrpc AppProtocol = "grpc"
	AppProtocolHTTP AppProtocol = "http"
)

// PossibleAppProtocolValues returns the possible values for the AppProtocol const type.
func PossibleAppProtocolValues() []AppProtocol {
	return []AppProtocol{
		AppProtocolGrpc,
		AppProtocolHTTP,
	}
}

// BindingType - Custom Domain binding type.
type BindingType string

const (
	BindingTypeDisabled   BindingType = "Disabled"
	BindingTypeSniEnabled BindingType = "SniEnabled"
)

// PossibleBindingTypeValues returns the possible values for the BindingType const type.
func PossibleBindingTypeValues() []BindingType {
	return []BindingType{
		BindingTypeDisabled,
		BindingTypeSniEnabled,
	}
}

// CertificateProvisioningState - Provisioning state of the certificate.
type CertificateProvisioningState string

const (
	CertificateProvisioningStateCanceled     CertificateProvisioningState = "Canceled"
	CertificateProvisioningStateDeleteFailed CertificateProvisioningState = "DeleteFailed"
	CertificateProvisioningStateFailed       CertificateProvisioningState = "Failed"
	CertificateProvisioningStatePending      CertificateProvisioningState = "Pending"
	CertificateProvisioningStateSucceeded    CertificateProvisioningState = "Succeeded"
)

// PossibleCertificateProvisioningStateValues returns the possible values for the CertificateProvisioningState const type.
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return []CertificateProvisioningState{
		CertificateProvisioningStateCanceled,
		CertificateProvisioningStateDeleteFailed,
		CertificateProvisioningStateFailed,
		CertificateProvisioningStatePending,
		CertificateProvisioningStateSucceeded,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// ContainerAppProvisioningState - Provisioning state of the Container App.
type ContainerAppProvisioningState string

const (
	ContainerAppProvisioningStateCanceled   ContainerAppProvisioningState = "Canceled"
	ContainerAppProvisioningStateFailed     ContainerAppProvisioningState = "Failed"
	ContainerAppProvisioningStateInProgress ContainerAppProvisioningState = "InProgress"
	ContainerAppProvisioningStateSucceeded  ContainerAppProvisioningState = "Succeeded"
)

// PossibleContainerAppProvisioningStateValues returns the possible values for the ContainerAppProvisioningState const type.
func PossibleContainerAppProvisioningStateValues() []ContainerAppProvisioningState {
	return []ContainerAppProvisioningState{
		ContainerAppProvisioningStateCanceled,
		ContainerAppProvisioningStateFailed,
		ContainerAppProvisioningStateInProgress,
		ContainerAppProvisioningStateSucceeded,
	}
}

// CookieExpirationConvention - The convention used when determining the session cookie's expiration.
type CookieExpirationConvention string

const (
	CookieExpirationConventionFixedTime               CookieExpirationConvention = "FixedTime"
	CookieExpirationConventionIdentityProviderDerived CookieExpirationConvention = "IdentityProviderDerived"
)

// PossibleCookieExpirationConventionValues returns the possible values for the CookieExpirationConvention const type.
func PossibleCookieExpirationConventionValues() []CookieExpirationConvention {
	return []CookieExpirationConvention{
		CookieExpirationConventionFixedTime,
		CookieExpirationConventionIdentityProviderDerived,
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

// DNSVerificationTestResult - DNS verification test result.
type DNSVerificationTestResult string

const (
	DNSVerificationTestResultPassed  DNSVerificationTestResult = "Passed"
	DNSVerificationTestResultFailed  DNSVerificationTestResult = "Failed"
	DNSVerificationTestResultSkipped DNSVerificationTestResult = "Skipped"
)

// PossibleDNSVerificationTestResultValues returns the possible values for the DNSVerificationTestResult const type.
func PossibleDNSVerificationTestResultValues() []DNSVerificationTestResult {
	return []DNSVerificationTestResult{
		DNSVerificationTestResultPassed,
		DNSVerificationTestResultFailed,
		DNSVerificationTestResultSkipped,
	}
}

// EnvironmentProvisioningState - Provisioning state of the Environment.
type EnvironmentProvisioningState string

const (
	EnvironmentProvisioningStateCanceled                      EnvironmentProvisioningState = "Canceled"
	EnvironmentProvisioningStateFailed                        EnvironmentProvisioningState = "Failed"
	EnvironmentProvisioningStateInfrastructureSetupComplete   EnvironmentProvisioningState = "InfrastructureSetupComplete"
	EnvironmentProvisioningStateInfrastructureSetupInProgress EnvironmentProvisioningState = "InfrastructureSetupInProgress"
	EnvironmentProvisioningStateInitializationInProgress      EnvironmentProvisioningState = "InitializationInProgress"
	EnvironmentProvisioningStateScheduledForDelete            EnvironmentProvisioningState = "ScheduledForDelete"
	EnvironmentProvisioningStateSucceeded                     EnvironmentProvisioningState = "Succeeded"
	EnvironmentProvisioningStateUpgradeFailed                 EnvironmentProvisioningState = "UpgradeFailed"
	EnvironmentProvisioningStateUpgradeRequested              EnvironmentProvisioningState = "UpgradeRequested"
	EnvironmentProvisioningStateWaiting                       EnvironmentProvisioningState = "Waiting"
)

// PossibleEnvironmentProvisioningStateValues returns the possible values for the EnvironmentProvisioningState const type.
func PossibleEnvironmentProvisioningStateValues() []EnvironmentProvisioningState {
	return []EnvironmentProvisioningState{
		EnvironmentProvisioningStateCanceled,
		EnvironmentProvisioningStateFailed,
		EnvironmentProvisioningStateInfrastructureSetupComplete,
		EnvironmentProvisioningStateInfrastructureSetupInProgress,
		EnvironmentProvisioningStateInitializationInProgress,
		EnvironmentProvisioningStateScheduledForDelete,
		EnvironmentProvisioningStateSucceeded,
		EnvironmentProvisioningStateUpgradeFailed,
		EnvironmentProvisioningStateUpgradeRequested,
		EnvironmentProvisioningStateWaiting,
	}
}

// ForwardProxyConvention - The convention used to determine the url of the request made.
type ForwardProxyConvention string

const (
	ForwardProxyConventionNoProxy  ForwardProxyConvention = "NoProxy"
	ForwardProxyConventionStandard ForwardProxyConvention = "Standard"
	ForwardProxyConventionCustom   ForwardProxyConvention = "Custom"
)

// PossibleForwardProxyConventionValues returns the possible values for the ForwardProxyConvention const type.
func PossibleForwardProxyConventionValues() []ForwardProxyConvention {
	return []ForwardProxyConvention{
		ForwardProxyConventionNoProxy,
		ForwardProxyConventionStandard,
		ForwardProxyConventionCustom,
	}
}

// IngressTransportMethod - Ingress transport protocol
type IngressTransportMethod string

const (
	IngressTransportMethodAuto  IngressTransportMethod = "auto"
	IngressTransportMethodHTTP  IngressTransportMethod = "http"
	IngressTransportMethodHTTP2 IngressTransportMethod = "http2"
)

// PossibleIngressTransportMethodValues returns the possible values for the IngressTransportMethod const type.
func PossibleIngressTransportMethodValues() []IngressTransportMethod {
	return []IngressTransportMethod{
		IngressTransportMethodAuto,
		IngressTransportMethodHTTP,
		IngressTransportMethodHTTP2,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// RevisionHealthState - Current health State of the revision
type RevisionHealthState string

const (
	RevisionHealthStateHealthy   RevisionHealthState = "Healthy"
	RevisionHealthStateNone      RevisionHealthState = "None"
	RevisionHealthStateUnhealthy RevisionHealthState = "Unhealthy"
)

// PossibleRevisionHealthStateValues returns the possible values for the RevisionHealthState const type.
func PossibleRevisionHealthStateValues() []RevisionHealthState {
	return []RevisionHealthState{
		RevisionHealthStateHealthy,
		RevisionHealthStateNone,
		RevisionHealthStateUnhealthy,
	}
}

// RevisionProvisioningState - Current provisioning State of the revision
type RevisionProvisioningState string

const (
	RevisionProvisioningStateDeprovisioned  RevisionProvisioningState = "Deprovisioned"
	RevisionProvisioningStateDeprovisioning RevisionProvisioningState = "Deprovisioning"
	RevisionProvisioningStateFailed         RevisionProvisioningState = "Failed"
	RevisionProvisioningStateProvisioned    RevisionProvisioningState = "Provisioned"
	RevisionProvisioningStateProvisioning   RevisionProvisioningState = "Provisioning"
)

// PossibleRevisionProvisioningStateValues returns the possible values for the RevisionProvisioningState const type.
func PossibleRevisionProvisioningStateValues() []RevisionProvisioningState {
	return []RevisionProvisioningState{
		RevisionProvisioningStateDeprovisioned,
		RevisionProvisioningStateDeprovisioning,
		RevisionProvisioningStateFailed,
		RevisionProvisioningStateProvisioned,
		RevisionProvisioningStateProvisioning,
	}
}

// Scheme - Scheme to use for connecting to the host. Defaults to HTTP.
type Scheme string

const (
	SchemeHTTP  Scheme = "HTTP"
	SchemeHTTPS Scheme = "HTTPS"
)

// PossibleSchemeValues returns the possible values for the Scheme const type.
func PossibleSchemeValues() []Scheme {
	return []Scheme{
		SchemeHTTP,
		SchemeHTTPS,
	}
}

// SourceControlOperationState - Current provisioning State of the operation
type SourceControlOperationState string

const (
	SourceControlOperationStateCanceled   SourceControlOperationState = "Canceled"
	SourceControlOperationStateFailed     SourceControlOperationState = "Failed"
	SourceControlOperationStateInProgress SourceControlOperationState = "InProgress"
	SourceControlOperationStateSucceeded  SourceControlOperationState = "Succeeded"
)

// PossibleSourceControlOperationStateValues returns the possible values for the SourceControlOperationState const type.
func PossibleSourceControlOperationStateValues() []SourceControlOperationState {
	return []SourceControlOperationState{
		SourceControlOperationStateCanceled,
		SourceControlOperationStateFailed,
		SourceControlOperationStateInProgress,
		SourceControlOperationStateSucceeded,
	}
}

// StorageType - Storage type for the volume. If not provided, use EmptyDir.
type StorageType string

const (
	StorageTypeAzureFile StorageType = "AzureFile"
	StorageTypeEmptyDir  StorageType = "EmptyDir"
)

// PossibleStorageTypeValues returns the possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{
		StorageTypeAzureFile,
		StorageTypeEmptyDir,
	}
}

// Type - The type of probe.
type Type string

const (
	TypeLiveness  Type = "Liveness"
	TypeReadiness Type = "Readiness"
	TypeStartup   Type = "Startup"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeLiveness,
		TypeReadiness,
		TypeStartup,
	}
}

// UnauthenticatedClientActionV2 - The action to take when an unauthenticated client attempts to access the app.
type UnauthenticatedClientActionV2 string

const (
	UnauthenticatedClientActionV2RedirectToLoginPage UnauthenticatedClientActionV2 = "RedirectToLoginPage"
	UnauthenticatedClientActionV2AllowAnonymous      UnauthenticatedClientActionV2 = "AllowAnonymous"
	UnauthenticatedClientActionV2Return401           UnauthenticatedClientActionV2 = "Return401"
	UnauthenticatedClientActionV2Return403           UnauthenticatedClientActionV2 = "Return403"
)

// PossibleUnauthenticatedClientActionV2Values returns the possible values for the UnauthenticatedClientActionV2 const type.
func PossibleUnauthenticatedClientActionV2Values() []UnauthenticatedClientActionV2 {
	return []UnauthenticatedClientActionV2{
		UnauthenticatedClientActionV2RedirectToLoginPage,
		UnauthenticatedClientActionV2AllowAnonymous,
		UnauthenticatedClientActionV2Return401,
		UnauthenticatedClientActionV2Return403,
	}
}
