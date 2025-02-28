// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdeviceregistry

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
	moduleVersion = "v1.0.1"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AuthenticationMethod - The method to authenticate the user of the client at the server.
type AuthenticationMethod string

const (
	// AuthenticationMethodAnonymous - The user authentication method is anonymous.
	AuthenticationMethodAnonymous AuthenticationMethod = "Anonymous"
	// AuthenticationMethodCertificate - The user authentication method is an x509 certificate.
	AuthenticationMethodCertificate AuthenticationMethod = "Certificate"
	// AuthenticationMethodUsernamePassword - The user authentication method is a username and password.
	AuthenticationMethodUsernamePassword AuthenticationMethod = "UsernamePassword"
)

// PossibleAuthenticationMethodValues returns the possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodAnonymous,
		AuthenticationMethodCertificate,
		AuthenticationMethodUsernamePassword,
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

// DataPointObservabilityMode - Defines the data point observability mode.
type DataPointObservabilityMode string

const (
	// DataPointObservabilityModeCounter - Map as counter to OpenTelemetry.
	DataPointObservabilityModeCounter DataPointObservabilityMode = "Counter"
	// DataPointObservabilityModeGauge - Map as gauge to OpenTelemetry.
	DataPointObservabilityModeGauge DataPointObservabilityMode = "Gauge"
	// DataPointObservabilityModeHistogram - Map as histogram to OpenTelemetry.
	DataPointObservabilityModeHistogram DataPointObservabilityMode = "Histogram"
	// DataPointObservabilityModeLog - Map as log to OpenTelemetry.
	DataPointObservabilityModeLog DataPointObservabilityMode = "Log"
	// DataPointObservabilityModeNone - No mapping to OpenTelemetry.
	DataPointObservabilityModeNone DataPointObservabilityMode = "None"
)

// PossibleDataPointObservabilityModeValues returns the possible values for the DataPointObservabilityMode const type.
func PossibleDataPointObservabilityModeValues() []DataPointObservabilityMode {
	return []DataPointObservabilityMode{
		DataPointObservabilityModeCounter,
		DataPointObservabilityModeGauge,
		DataPointObservabilityModeHistogram,
		DataPointObservabilityModeLog,
		DataPointObservabilityModeNone,
	}
}

// EventObservabilityMode - Defines the event observability mode.
type EventObservabilityMode string

const (
	// EventObservabilityModeLog - Map as log to OpenTelemetry.
	EventObservabilityModeLog EventObservabilityMode = "Log"
	// EventObservabilityModeNone - No mapping to OpenTelemetry.
	EventObservabilityModeNone EventObservabilityMode = "None"
)

// PossibleEventObservabilityModeValues returns the possible values for the EventObservabilityMode const type.
func PossibleEventObservabilityModeValues() []EventObservabilityMode {
	return []EventObservabilityMode{
		EventObservabilityModeLog,
		EventObservabilityModeNone,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - The provisioning status of the resource.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource has been accepted by the server.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource is deleting.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// TopicRetainType - Topic retain types.
type TopicRetainType string

const (
	// TopicRetainTypeKeep - Retain the messages.
	TopicRetainTypeKeep TopicRetainType = "Keep"
	// TopicRetainTypeNever - Never retain messages.
	TopicRetainTypeNever TopicRetainType = "Never"
)

// PossibleTopicRetainTypeValues returns the possible values for the TopicRetainType const type.
func PossibleTopicRetainTypeValues() []TopicRetainType {
	return []TopicRetainType{
		TopicRetainTypeKeep,
		TopicRetainTypeNever,
	}
}
