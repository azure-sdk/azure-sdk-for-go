//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

const (
	moduleName    = "armdesktopvirtualization"
	moduleVersion = "v2.1.0-beta.1"
)

// ApplicationGroupType - Resource Type of ApplicationGroup.
type ApplicationGroupType string

const (
	ApplicationGroupTypeDesktop   ApplicationGroupType = "Desktop"
	ApplicationGroupTypeRemoteApp ApplicationGroupType = "RemoteApp"
)

// PossibleApplicationGroupTypeValues returns the possible values for the ApplicationGroupType const type.
func PossibleApplicationGroupTypeValues() []ApplicationGroupType {
	return []ApplicationGroupType{
		ApplicationGroupTypeDesktop,
		ApplicationGroupTypeRemoteApp,
	}
}

// ApplicationType - Application type of application.
type ApplicationType string

const (
	ApplicationTypeDesktop   ApplicationType = "Desktop"
	ApplicationTypeRemoteApp ApplicationType = "RemoteApp"
)

// PossibleApplicationTypeValues returns the possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{
		ApplicationTypeDesktop,
		ApplicationTypeRemoteApp,
	}
}

// CommandLineSetting - Specifies whether this published application can be launched with command line arguments provided
// by the client, command line arguments specified at publish time, or no command line arguments at all.
type CommandLineSetting string

const (
	CommandLineSettingAllow      CommandLineSetting = "Allow"
	CommandLineSettingDoNotAllow CommandLineSetting = "DoNotAllow"
	CommandLineSettingRequire    CommandLineSetting = "Require"
)

// PossibleCommandLineSettingValues returns the possible values for the CommandLineSetting const type.
func PossibleCommandLineSettingValues() []CommandLineSetting {
	return []CommandLineSetting{
		CommandLineSettingAllow,
		CommandLineSettingDoNotAllow,
		CommandLineSettingRequire,
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

// DayOfWeek - Day of the week.
type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// HealthCheckName - Represents the name of the health check operation performed.
type HealthCheckName string

const (
	// HealthCheckNameAppAttachHealthCheck - Verifies that the AppAttachService is healthy (there were no issues during package
	// staging). The AppAttachService is used to enable the staging/registration (and eventual deregistration/destaging) of MSIX
	// apps that have been set up by the tenant admin. This checks whether the component had any failures during package staging.
	// Failures in staging will prevent some MSIX apps from working properly for the end user. If this check fails, it is non
	// fatal and the machine still can service connections, main issue may be certain apps will not work for end-users. (Currently
	// Enabled)
	HealthCheckNameAppAttachHealthCheck HealthCheckName = "AppAttachHealthCheck"
	// HealthCheckNameDomainJoinedCheck - Verifies the SessionHost is joined to a domain. If this check fails is classified as
	// fatal as no connection can succeed if the SessionHost is not joined to the domain. (Currently Enabled)
	HealthCheckNameDomainJoinedCheck HealthCheckName = "DomainJoinedCheck"
	// HealthCheckNameDomainReachable - Verifies the domain the SessionHost is joined to is still reachable. If this check fails
	// is classified as fatal as no connection can succeed if the domain the SessionHost is joined is not reachable at the time
	// of connection. (Currently Disabled)
	HealthCheckNameDomainReachable HealthCheckName = "DomainReachable"
	// HealthCheckNameDomainTrustCheck - Verifies the SessionHost is not experiencing domain trust issues that will prevent authentication
	// on SessionHost at connection time when session is created. If this check fails is classified as fatal as no connection
	// can succeed if we cannot reach the domain for authentication on the SessionHost. (Currently Enabled)
	HealthCheckNameDomainTrustCheck HealthCheckName = "DomainTrustCheck"
	// HealthCheckNameFSLogixHealthCheck - Verifies the FSLogix service is up and running to make sure users' profiles are loaded
	// in the session. If this check fails is classified as fatal as even if the connection can succeed, user experience is bad
	// as the user profile cannot be loaded and user will get a temporary profile in the session. (Currently Disabled)
	HealthCheckNameFSLogixHealthCheck HealthCheckName = "FSLogixHealthCheck"
	// HealthCheckNameMetaDataServiceCheck - Verifies the metadata service is accessible and return compute properties. (Currently
	// Enabled)
	HealthCheckNameMetaDataServiceCheck HealthCheckName = "MetaDataServiceCheck"
	// HealthCheckNameMonitoringAgentCheck - Verifies that the required Geneva agent is running. If this check fails, it is non
	// fatal and the machine still can service connections, main issue may be that monitoring agent is missing or running (possibly)
	// older version. (Currently Enabled)
	HealthCheckNameMonitoringAgentCheck HealthCheckName = "MonitoringAgentCheck"
	// HealthCheckNameSupportedEncryptionCheck - Verifies the value of SecurityLayer registration key. If the value is 0 (SecurityLayer.RDP)
	// this check fails with Error code = NativeMethodErrorCode.E_FAIL and is fatal. If the value is 1 (SecurityLayer.Negotiate)
	// this check fails with Error code = NativeMethodErrorCode.ERROR_SUCCESS and is non fatal. (Currently Disabled)
	HealthCheckNameSupportedEncryptionCheck HealthCheckName = "SupportedEncryptionCheck"
	// HealthCheckNameSxSStackListenerCheck - Verifies that the SxS stack is up and running so connections can succeed. If this
	// check fails is classified as fatal as no connection can succeed if the SxS stack is not ready. (Currently Enabled)
	HealthCheckNameSxSStackListenerCheck HealthCheckName = "SxSStackListenerCheck"
	// HealthCheckNameUrlsAccessibleCheck - Verifies that the required WVD service and Geneva URLs are reachable from the SessionHost.
	// These URLs are: RdTokenUri, RdBrokerURI, RdDiagnosticsUri and storage blob URLs for agent monitoring (geneva). If this
	// check fails, it is non fatal and the machine still can service connections, main issue may be that monitoring agent is
	// unable to store warm path data (logs, operations ...). (Currently Disabled)
	HealthCheckNameUrlsAccessibleCheck HealthCheckName = "UrlsAccessibleCheck"
	// HealthCheckNameWebRTCRedirectorCheck - Verifies whether the WebRTCRedirector component is healthy. The WebRTCRedirector
	// component is used to optimize video and audio performance in Microsoft Teams. This checks whether the component is still
	// running, and whether there is a higher version available. If this check fails, it is non fatal and the machine still can
	// service connections, main issue may be the WebRTCRedirector component has to be restarted or updated. (Currently Disabled)
	HealthCheckNameWebRTCRedirectorCheck HealthCheckName = "WebRTCRedirectorCheck"
)

// PossibleHealthCheckNameValues returns the possible values for the HealthCheckName const type.
func PossibleHealthCheckNameValues() []HealthCheckName {
	return []HealthCheckName{
		HealthCheckNameAppAttachHealthCheck,
		HealthCheckNameDomainJoinedCheck,
		HealthCheckNameDomainReachable,
		HealthCheckNameDomainTrustCheck,
		HealthCheckNameFSLogixHealthCheck,
		HealthCheckNameMetaDataServiceCheck,
		HealthCheckNameMonitoringAgentCheck,
		HealthCheckNameSupportedEncryptionCheck,
		HealthCheckNameSxSStackListenerCheck,
		HealthCheckNameUrlsAccessibleCheck,
		HealthCheckNameWebRTCRedirectorCheck,
	}
}

// HealthCheckResult - Represents the Health state of the health check we performed.
type HealthCheckResult string

const (
	// HealthCheckResultHealthCheckFailed - Health check failed.
	HealthCheckResultHealthCheckFailed HealthCheckResult = "HealthCheckFailed"
	// HealthCheckResultHealthCheckSucceeded - Health check passed.
	HealthCheckResultHealthCheckSucceeded HealthCheckResult = "HealthCheckSucceeded"
	// HealthCheckResultSessionHostShutdown - We received a Shutdown notification.
	HealthCheckResultSessionHostShutdown HealthCheckResult = "SessionHostShutdown"
	// HealthCheckResultUnknown - Health check result is not currently known.
	HealthCheckResultUnknown HealthCheckResult = "Unknown"
)

// PossibleHealthCheckResultValues returns the possible values for the HealthCheckResult const type.
func PossibleHealthCheckResultValues() []HealthCheckResult {
	return []HealthCheckResult{
		HealthCheckResultHealthCheckFailed,
		HealthCheckResultHealthCheckSucceeded,
		HealthCheckResultSessionHostShutdown,
		HealthCheckResultUnknown,
	}
}

// HostPoolType - HostPool type for desktop.
type HostPoolType string

const (
	// HostPoolTypeBYODesktop - Users assign their own machines, load balancing logic remains the same as Personal. PersonalDesktopAssignmentType
	// must be Direct.
	HostPoolTypeBYODesktop HostPoolType = "BYODesktop"
	// HostPoolTypePersonal - Users will be assigned a SessionHost either by administrators (PersonalDesktopAssignmentType = Direct)
	// or upon connecting to the pool (PersonalDesktopAssignmentType = Automatic). They will always be redirected to their assigned
	// SessionHost.
	HostPoolTypePersonal HostPoolType = "Personal"
	// HostPoolTypePooled - Users get a new (random) SessionHost every time it connects to the HostPool.
	HostPoolTypePooled HostPoolType = "Pooled"
)

// PossibleHostPoolTypeValues returns the possible values for the HostPoolType const type.
func PossibleHostPoolTypeValues() []HostPoolType {
	return []HostPoolType{
		HostPoolTypeBYODesktop,
		HostPoolTypePersonal,
		HostPoolTypePooled,
	}
}

// HostpoolPublicNetworkAccess - Enabled allows this resource to be accessed from both public and private networks, Disabled
// allows this resource to only be accessed via private endpoints
type HostpoolPublicNetworkAccess string

const (
	HostpoolPublicNetworkAccessDisabled                   HostpoolPublicNetworkAccess = "Disabled"
	HostpoolPublicNetworkAccessEnabled                    HostpoolPublicNetworkAccess = "Enabled"
	HostpoolPublicNetworkAccessEnabledForClientsOnly      HostpoolPublicNetworkAccess = "EnabledForClientsOnly"
	HostpoolPublicNetworkAccessEnabledForSessionHostsOnly HostpoolPublicNetworkAccess = "EnabledForSessionHostsOnly"
)

// PossibleHostpoolPublicNetworkAccessValues returns the possible values for the HostpoolPublicNetworkAccess const type.
func PossibleHostpoolPublicNetworkAccessValues() []HostpoolPublicNetworkAccess {
	return []HostpoolPublicNetworkAccess{
		HostpoolPublicNetworkAccessDisabled,
		HostpoolPublicNetworkAccessEnabled,
		HostpoolPublicNetworkAccessEnabledForClientsOnly,
		HostpoolPublicNetworkAccessEnabledForSessionHostsOnly,
	}
}

// LoadBalancerType - The type of the load balancer.
type LoadBalancerType string

const (
	LoadBalancerTypeBreadthFirst LoadBalancerType = "BreadthFirst"
	LoadBalancerTypeDepthFirst   LoadBalancerType = "DepthFirst"
	LoadBalancerTypePersistent   LoadBalancerType = "Persistent"
)

// PossibleLoadBalancerTypeValues returns the possible values for the LoadBalancerType const type.
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return []LoadBalancerType{
		LoadBalancerTypeBreadthFirst,
		LoadBalancerTypeDepthFirst,
		LoadBalancerTypePersistent,
	}
}

// PersonalDesktopAssignmentType - PersonalDesktopAssignment type for HostPool.
type PersonalDesktopAssignmentType string

const (
	PersonalDesktopAssignmentTypeAutomatic PersonalDesktopAssignmentType = "Automatic"
	PersonalDesktopAssignmentTypeDirect    PersonalDesktopAssignmentType = "Direct"
)

// PossiblePersonalDesktopAssignmentTypeValues returns the possible values for the PersonalDesktopAssignmentType const type.
func PossiblePersonalDesktopAssignmentTypeValues() []PersonalDesktopAssignmentType {
	return []PersonalDesktopAssignmentType{
		PersonalDesktopAssignmentTypeAutomatic,
		PersonalDesktopAssignmentTypeDirect,
	}
}

// PreferredAppGroupType - The type of preferred application group type, default to Desktop Application Group
type PreferredAppGroupType string

const (
	PreferredAppGroupTypeDesktop          PreferredAppGroupType = "Desktop"
	PreferredAppGroupTypeNone             PreferredAppGroupType = "None"
	PreferredAppGroupTypeRailApplications PreferredAppGroupType = "RailApplications"
)

// PossiblePreferredAppGroupTypeValues returns the possible values for the PreferredAppGroupType const type.
func PossiblePreferredAppGroupTypeValues() []PreferredAppGroupType {
	return []PreferredAppGroupType{
		PreferredAppGroupTypeDesktop,
		PreferredAppGroupTypeNone,
		PreferredAppGroupTypeRailApplications,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// PublicNetworkAccess - Enabled allows this resource to be accessed from both public and private networks, Disabled allows
// this resource to only be accessed via private endpoints
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// RegistrationTokenOperation - The type of resetting the token.
type RegistrationTokenOperation string

const (
	RegistrationTokenOperationDelete RegistrationTokenOperation = "Delete"
	RegistrationTokenOperationNone   RegistrationTokenOperation = "None"
	RegistrationTokenOperationUpdate RegistrationTokenOperation = "Update"
)

// PossibleRegistrationTokenOperationValues returns the possible values for the RegistrationTokenOperation const type.
func PossibleRegistrationTokenOperationValues() []RegistrationTokenOperation {
	return []RegistrationTokenOperation{
		RegistrationTokenOperationDelete,
		RegistrationTokenOperationNone,
		RegistrationTokenOperationUpdate,
	}
}

// RemoteApplicationType - Resource Type of Application.
type RemoteApplicationType string

const (
	RemoteApplicationTypeInBuilt         RemoteApplicationType = "InBuilt"
	RemoteApplicationTypeMsixApplication RemoteApplicationType = "MsixApplication"
)

// PossibleRemoteApplicationTypeValues returns the possible values for the RemoteApplicationType const type.
func PossibleRemoteApplicationTypeValues() []RemoteApplicationType {
	return []RemoteApplicationType{
		RemoteApplicationTypeInBuilt,
		RemoteApplicationTypeMsixApplication,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierFree     SKUTier = "Free"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierFree,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// SSOSecretType - The type of single sign on Secret Type.
type SSOSecretType string

const (
	SSOSecretTypeCertificate           SSOSecretType = "Certificate"
	SSOSecretTypeCertificateInKeyVault SSOSecretType = "CertificateInKeyVault"
	SSOSecretTypeSharedKey             SSOSecretType = "SharedKey"
	SSOSecretTypeSharedKeyInKeyVault   SSOSecretType = "SharedKeyInKeyVault"
)

// PossibleSSOSecretTypeValues returns the possible values for the SSOSecretType const type.
func PossibleSSOSecretTypeValues() []SSOSecretType {
	return []SSOSecretType{
		SSOSecretTypeCertificate,
		SSOSecretTypeCertificateInKeyVault,
		SSOSecretTypeSharedKey,
		SSOSecretTypeSharedKeyInKeyVault,
	}
}

// ScalingHostPoolType - HostPool type for desktop.
type ScalingHostPoolType string

const (
	// ScalingHostPoolTypePooled - Users get a new (random) SessionHost every time it connects to the HostPool.
	ScalingHostPoolTypePooled ScalingHostPoolType = "Pooled"
)

// PossibleScalingHostPoolTypeValues returns the possible values for the ScalingHostPoolType const type.
func PossibleScalingHostPoolTypeValues() []ScalingHostPoolType {
	return []ScalingHostPoolType{
		ScalingHostPoolTypePooled,
	}
}

type ScalingScheduleDaysOfWeekItem string

const (
	ScalingScheduleDaysOfWeekItemFriday    ScalingScheduleDaysOfWeekItem = "Friday"
	ScalingScheduleDaysOfWeekItemMonday    ScalingScheduleDaysOfWeekItem = "Monday"
	ScalingScheduleDaysOfWeekItemSaturday  ScalingScheduleDaysOfWeekItem = "Saturday"
	ScalingScheduleDaysOfWeekItemSunday    ScalingScheduleDaysOfWeekItem = "Sunday"
	ScalingScheduleDaysOfWeekItemThursday  ScalingScheduleDaysOfWeekItem = "Thursday"
	ScalingScheduleDaysOfWeekItemTuesday   ScalingScheduleDaysOfWeekItem = "Tuesday"
	ScalingScheduleDaysOfWeekItemWednesday ScalingScheduleDaysOfWeekItem = "Wednesday"
)

// PossibleScalingScheduleDaysOfWeekItemValues returns the possible values for the ScalingScheduleDaysOfWeekItem const type.
func PossibleScalingScheduleDaysOfWeekItemValues() []ScalingScheduleDaysOfWeekItem {
	return []ScalingScheduleDaysOfWeekItem{
		ScalingScheduleDaysOfWeekItemFriday,
		ScalingScheduleDaysOfWeekItemMonday,
		ScalingScheduleDaysOfWeekItemSaturday,
		ScalingScheduleDaysOfWeekItemSunday,
		ScalingScheduleDaysOfWeekItemThursday,
		ScalingScheduleDaysOfWeekItemTuesday,
		ScalingScheduleDaysOfWeekItemWednesday,
	}
}

// SessionHandlingOperation - Action to be taken after a user disconnect during the ramp up period.
type SessionHandlingOperation string

const (
	SessionHandlingOperationDeallocate SessionHandlingOperation = "Deallocate"
	SessionHandlingOperationHibernate  SessionHandlingOperation = "Hibernate"
	SessionHandlingOperationNone       SessionHandlingOperation = "None"
)

// PossibleSessionHandlingOperationValues returns the possible values for the SessionHandlingOperation const type.
func PossibleSessionHandlingOperationValues() []SessionHandlingOperation {
	return []SessionHandlingOperation{
		SessionHandlingOperationDeallocate,
		SessionHandlingOperationHibernate,
		SessionHandlingOperationNone,
	}
}

// SessionHostComponentUpdateType - The type of maintenance for session host components.
type SessionHostComponentUpdateType string

const (
	// SessionHostComponentUpdateTypeDefault - Agent and other agent side components are delivery schedule is controlled by WVD
	// Infra.
	SessionHostComponentUpdateTypeDefault SessionHostComponentUpdateType = "Default"
	// SessionHostComponentUpdateTypeScheduled - TenantAdmin have opted in for Scheduled Component Update feature.
	SessionHostComponentUpdateTypeScheduled SessionHostComponentUpdateType = "Scheduled"
)

// PossibleSessionHostComponentUpdateTypeValues returns the possible values for the SessionHostComponentUpdateType const type.
func PossibleSessionHostComponentUpdateTypeValues() []SessionHostComponentUpdateType {
	return []SessionHostComponentUpdateType{
		SessionHostComponentUpdateTypeDefault,
		SessionHostComponentUpdateTypeScheduled,
	}
}

// SessionHostLoadBalancingAlgorithm - Load balancing algorithm for ramp up period.
type SessionHostLoadBalancingAlgorithm string

const (
	SessionHostLoadBalancingAlgorithmBreadthFirst SessionHostLoadBalancingAlgorithm = "BreadthFirst"
	SessionHostLoadBalancingAlgorithmDepthFirst   SessionHostLoadBalancingAlgorithm = "DepthFirst"
)

// PossibleSessionHostLoadBalancingAlgorithmValues returns the possible values for the SessionHostLoadBalancingAlgorithm const type.
func PossibleSessionHostLoadBalancingAlgorithmValues() []SessionHostLoadBalancingAlgorithm {
	return []SessionHostLoadBalancingAlgorithm{
		SessionHostLoadBalancingAlgorithmBreadthFirst,
		SessionHostLoadBalancingAlgorithmDepthFirst,
	}
}

// SessionState - State of user session.
type SessionState string

const (
	SessionStateActive                 SessionState = "Active"
	SessionStateDisconnected           SessionState = "Disconnected"
	SessionStateLogOff                 SessionState = "LogOff"
	SessionStatePending                SessionState = "Pending"
	SessionStateUnknown                SessionState = "Unknown"
	SessionStateUserProfileDiskMounted SessionState = "UserProfileDiskMounted"
)

// PossibleSessionStateValues returns the possible values for the SessionState const type.
func PossibleSessionStateValues() []SessionState {
	return []SessionState{
		SessionStateActive,
		SessionStateDisconnected,
		SessionStateLogOff,
		SessionStatePending,
		SessionStateUnknown,
		SessionStateUserProfileDiskMounted,
	}
}

// SetStartVMOnConnect - The desired configuration of Start VM On Connect for the hostpool during the ramp up phase. If this
// is disabled, session hosts must be turned on using rampUpAutoStartHosts or by turning them on
// manually.
type SetStartVMOnConnect string

const (
	SetStartVMOnConnectDisable SetStartVMOnConnect = "Disable"
	SetStartVMOnConnectEnable  SetStartVMOnConnect = "Enable"
)

// PossibleSetStartVMOnConnectValues returns the possible values for the SetStartVMOnConnect const type.
func PossibleSetStartVMOnConnectValues() []SetStartVMOnConnect {
	return []SetStartVMOnConnect{
		SetStartVMOnConnectDisable,
		SetStartVMOnConnectEnable,
	}
}

// StartupBehavior - The desired startup behavior during the ramp up period for personal vms in the hostpool.
type StartupBehavior string

const (
	// StartupBehaviorAll - All personal session hosts in the hostpool will be started during ramp up.
	StartupBehaviorAll StartupBehavior = "All"
	// StartupBehaviorNone - Session hosts will not be started by the service. This setting depends on Start VM on Connect to
	// be enabled to start the session hosts.
	StartupBehaviorNone StartupBehavior = "None"
	// StartupBehaviorWithAssignedUser - Session hosts with an assigned user will be started during Ramp Up
	StartupBehaviorWithAssignedUser StartupBehavior = "WithAssignedUser"
)

// PossibleStartupBehaviorValues returns the possible values for the StartupBehavior const type.
func PossibleStartupBehaviorValues() []StartupBehavior {
	return []StartupBehavior{
		StartupBehaviorAll,
		StartupBehaviorNone,
		StartupBehaviorWithAssignedUser,
	}
}

// Status - Status for a SessionHost.
type Status string

const (
	// StatusAvailable - Session Host has passed all the health checks and is available to handle connections.
	StatusAvailable Status = "Available"
	// StatusDisconnected - The Session Host is unavailable because it is currently disconnected.
	StatusDisconnected Status = "Disconnected"
	// StatusDomainTrustRelationshipLost - SessionHost's domain trust relationship lost
	StatusDomainTrustRelationshipLost Status = "DomainTrustRelationshipLost"
	// StatusFSLogixNotHealthy - FSLogix is in an unhealthy state on the session host.
	StatusFSLogixNotHealthy Status = "FSLogixNotHealthy"
	// StatusNeedsAssistance - New status to inform admins that the health on their endpoint needs to be fixed. The connections
	// might not fail, as these issues are not fatal.
	StatusNeedsAssistance Status = "NeedsAssistance"
	// StatusNoHeartbeat - The Session Host is not heart beating.
	StatusNoHeartbeat Status = "NoHeartbeat"
	// StatusNotJoinedToDomain - SessionHost is not joined to domain.
	StatusNotJoinedToDomain Status = "NotJoinedToDomain"
	// StatusShutdown - Session Host is shutdown - RD Agent reported session host to be stopped or deallocated.
	StatusShutdown Status = "Shutdown"
	// StatusSxSStackListenerNotReady - SxS stack installed on the SessionHost is not ready to receive connections.
	StatusSxSStackListenerNotReady Status = "SxSStackListenerNotReady"
	// StatusUnavailable - Session Host is either turned off or has failed critical health checks which is causing service not
	// to be able to route connections to this session host. Note this replaces previous 'NoHeartBeat' status.
	StatusUnavailable Status = "Unavailable"
	// StatusUpgradeFailed - Session Host is unavailable because the critical component upgrade (agent, side-by-side stack, etc.)
	// failed.
	StatusUpgradeFailed Status = "UpgradeFailed"
	// StatusUpgrading - Session Host is unavailable because currently an upgrade of RDAgent/side-by-side stack is in progress.
	// Note: this state will be removed once the upgrade completes and the host is able to accept connections.
	StatusUpgrading Status = "Upgrading"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusAvailable,
		StatusDisconnected,
		StatusDomainTrustRelationshipLost,
		StatusFSLogixNotHealthy,
		StatusNeedsAssistance,
		StatusNoHeartbeat,
		StatusNotJoinedToDomain,
		StatusShutdown,
		StatusSxSStackListenerNotReady,
		StatusUnavailable,
		StatusUpgradeFailed,
		StatusUpgrading,
	}
}

// StopHostsWhen - Specifies when to stop hosts during ramp down period.
type StopHostsWhen string

const (
	StopHostsWhenZeroActiveSessions StopHostsWhen = "ZeroActiveSessions"
	StopHostsWhenZeroSessions       StopHostsWhen = "ZeroSessions"
)

// PossibleStopHostsWhenValues returns the possible values for the StopHostsWhen const type.
func PossibleStopHostsWhenValues() []StopHostsWhen {
	return []StopHostsWhen{
		StopHostsWhenZeroActiveSessions,
		StopHostsWhenZeroSessions,
	}
}

// UpdateState - Update state of a SessionHost.
type UpdateState string

const (
	UpdateStateFailed    UpdateState = "Failed"
	UpdateStateInitial   UpdateState = "Initial"
	UpdateStatePending   UpdateState = "Pending"
	UpdateStateStarted   UpdateState = "Started"
	UpdateStateSucceeded UpdateState = "Succeeded"
)

// PossibleUpdateStateValues returns the possible values for the UpdateState const type.
func PossibleUpdateStateValues() []UpdateState {
	return []UpdateState{
		UpdateStateFailed,
		UpdateStateInitial,
		UpdateStatePending,
		UpdateStateStarted,
		UpdateStateSucceeded,
	}
}
