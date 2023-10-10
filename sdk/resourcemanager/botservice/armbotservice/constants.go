//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice

const (
	moduleName    = "armbotservice"
	moduleVersion = "v1.2.0-beta.1"
)

// AccessMode - Access Mode of the resource association
type AccessMode string

const (
	AccessModeAudit    AccessMode = "Audit"
	AccessModeEnforced AccessMode = "Enforced"
	AccessModeLearning AccessMode = "Learning"
)

// PossibleAccessModeValues returns the possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{
		AccessModeAudit,
		AccessModeEnforced,
		AccessModeLearning,
	}
}

type ChannelName string

const (
	ChannelNameAcsChatChannel          ChannelName = "AcsChatChannel"
	ChannelNameAlexaChannel            ChannelName = "AlexaChannel"
	ChannelNameDirectLineChannel       ChannelName = "DirectLineChannel"
	ChannelNameDirectLineSpeechChannel ChannelName = "DirectLineSpeechChannel"
	ChannelNameEmailChannel            ChannelName = "EmailChannel"
	ChannelNameFacebookChannel         ChannelName = "FacebookChannel"
	ChannelNameKikChannel              ChannelName = "KikChannel"
	ChannelNameLineChannel             ChannelName = "LineChannel"
	ChannelNameM365Extensions          ChannelName = "M365Extensions"
	ChannelNameMsTeamsChannel          ChannelName = "MsTeamsChannel"
	ChannelNameOmnichannel             ChannelName = "Omnichannel"
	ChannelNameOutlookChannel          ChannelName = "OutlookChannel"
	ChannelNameSearchAssistant         ChannelName = "SearchAssistant"
	ChannelNameSkypeChannel            ChannelName = "SkypeChannel"
	ChannelNameSlackChannel            ChannelName = "SlackChannel"
	ChannelNameSmsChannel              ChannelName = "SmsChannel"
	ChannelNameTelegramChannel         ChannelName = "TelegramChannel"
	ChannelNameTelephonyChannel        ChannelName = "TelephonyChannel"
	ChannelNameWebChatChannel          ChannelName = "WebChatChannel"
)

// PossibleChannelNameValues returns the possible values for the ChannelName const type.
func PossibleChannelNameValues() []ChannelName {
	return []ChannelName{
		ChannelNameAcsChatChannel,
		ChannelNameAlexaChannel,
		ChannelNameDirectLineChannel,
		ChannelNameDirectLineSpeechChannel,
		ChannelNameEmailChannel,
		ChannelNameFacebookChannel,
		ChannelNameKikChannel,
		ChannelNameLineChannel,
		ChannelNameM365Extensions,
		ChannelNameMsTeamsChannel,
		ChannelNameOmnichannel,
		ChannelNameOutlookChannel,
		ChannelNameSearchAssistant,
		ChannelNameSkypeChannel,
		ChannelNameSlackChannel,
		ChannelNameSmsChannel,
		ChannelNameTelegramChannel,
		ChannelNameTelephonyChannel,
		ChannelNameWebChatChannel,
	}
}

// EmailChannelAuthMethod - Email channel auth method. 0 Password (Default); 1 Graph.
type EmailChannelAuthMethod float32

const (
	// EmailChannelAuthMethodGraph - Modern authentication.
	EmailChannelAuthMethodGraph EmailChannelAuthMethod = 1
	// EmailChannelAuthMethodPassword - Basic authentication.
	EmailChannelAuthMethodPassword EmailChannelAuthMethod = 0
)

// PossibleEmailChannelAuthMethodValues returns the possible values for the EmailChannelAuthMethod const type.
func PossibleEmailChannelAuthMethodValues() []EmailChannelAuthMethod {
	return []EmailChannelAuthMethod{
		EmailChannelAuthMethodGraph,
		EmailChannelAuthMethodPassword,
	}
}

// Key - Determines which key is to be regenerated
type Key string

const (
	KeyKey1 Key = "key1"
	KeyKey2 Key = "key2"
)

// PossibleKeyValues returns the possible values for the Key const type.
func PossibleKeyValues() []Key {
	return []Key{
		KeyKey1,
		KeyKey2,
	}
}

// Kind - Indicates the type of bot service
type Kind string

const (
	KindAzurebot Kind = "azurebot"
	KindBot      Kind = "bot"
	KindDesigner Kind = "designer"
	KindFunction Kind = "function"
	KindSdk      Kind = "sdk"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindAzurebot,
		KindBot,
		KindDesigner,
		KindFunction,
		KindSdk,
	}
}

// MsaAppType - Microsoft App Type for the bot
type MsaAppType string

const (
	MsaAppTypeMultiTenant     MsaAppType = "MultiTenant"
	MsaAppTypeSingleTenant    MsaAppType = "SingleTenant"
	MsaAppTypeUserAssignedMSI MsaAppType = "UserAssignedMSI"
)

// PossibleMsaAppTypeValues returns the possible values for the MsaAppType const type.
func PossibleMsaAppTypeValues() []MsaAppType {
	return []MsaAppType{
		MsaAppTypeMultiTenant,
		MsaAppTypeSingleTenant,
		MsaAppTypeUserAssignedMSI,
	}
}

// NspAccessRuleDirection - Direction of Access Rule
type NspAccessRuleDirection string

const (
	NspAccessRuleDirectionInbound  NspAccessRuleDirection = "Inbound"
	NspAccessRuleDirectionOutbound NspAccessRuleDirection = "Outbound"
)

// PossibleNspAccessRuleDirectionValues returns the possible values for the NspAccessRuleDirection const type.
func PossibleNspAccessRuleDirectionValues() []NspAccessRuleDirection {
	return []NspAccessRuleDirection{
		NspAccessRuleDirectionInbound,
		NspAccessRuleDirectionOutbound,
	}
}

// OperationResultStatus - The status of the operation being performed.
type OperationResultStatus string

const (
	OperationResultStatusCanceled  OperationResultStatus = "Canceled"
	OperationResultStatusFailed    OperationResultStatus = "Failed"
	OperationResultStatusRequested OperationResultStatus = "Requested"
	OperationResultStatusRunning   OperationResultStatus = "Running"
	OperationResultStatusSucceeded OperationResultStatus = "Succeeded"
)

// PossibleOperationResultStatusValues returns the possible values for the OperationResultStatus const type.
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return []OperationResultStatus{
		OperationResultStatusCanceled,
		OperationResultStatusFailed,
		OperationResultStatusRequested,
		OperationResultStatusRunning,
		OperationResultStatusSucceeded,
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

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether the bot is in an isolated network
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled           PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled            PublicNetworkAccess = "Enabled"
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
		PublicNetworkAccessSecuredByPerimeter,
	}
}

type RegenerateKeysChannelName string

const (
	RegenerateKeysChannelNameDirectLineChannel RegenerateKeysChannelName = "DirectLineChannel"
	RegenerateKeysChannelNameWebChatChannel    RegenerateKeysChannelName = "WebChatChannel"
)

// PossibleRegenerateKeysChannelNameValues returns the possible values for the RegenerateKeysChannelName const type.
func PossibleRegenerateKeysChannelNameValues() []RegenerateKeysChannelName {
	return []RegenerateKeysChannelName{
		RegenerateKeysChannelNameDirectLineChannel,
		RegenerateKeysChannelNameWebChatChannel,
	}
}

// SKUName - The name of SKU.
type SKUName string

const (
	SKUNameF0 SKUName = "F0"
	SKUNameS1 SKUName = "S1"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameF0,
		SKUNameS1,
	}
}

// SKUTier - Gets the sku tier. This is based on the SKU name.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierStandard,
	}
}

// Severity - Provisioning state of Network Security Perimeter configuration propagation
type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityWarning Severity = "Warning"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeverityError,
		SeverityWarning,
	}
}
