package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OperationStatusValues enumerates the values for operation status values.
type OperationStatusValues string

const (
	// Canceled ...
	Canceled OperationStatusValues = "Canceled"
	// Failed ...
	Failed OperationStatusValues = "Failed"
	// InProgress ...
	InProgress OperationStatusValues = "InProgress"
	// Invalid ...
	Invalid OperationStatusValues = "Invalid"
	// Succeeded ...
	Succeeded OperationStatusValues = "Succeeded"
)

// PossibleOperationStatusValuesValues returns an array of possible values for the OperationStatusValues const type.
func PossibleOperationStatusValuesValues() []OperationStatusValues {
	return []OperationStatusValues{Canceled, Failed, InProgress, Invalid, Succeeded}
}

// PrivateEndpointConnectionStatus enumerates the values for private endpoint connection status.
type PrivateEndpointConnectionStatus string

const (
	// Approved ...
	Approved PrivateEndpointConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected PrivateEndpointConnectionStatus = "Disconnected"
	// Pending ...
	Pending PrivateEndpointConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateEndpointConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointConnectionStatusValues returns an array of possible values for the PrivateEndpointConnectionStatus const type.
func PossiblePrivateEndpointConnectionStatusValues() []PrivateEndpointConnectionStatus {
	return []PrivateEndpointConnectionStatus{Approved, Disconnected, Pending, Rejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStatePending ...
	ProvisioningStatePending ProvisioningState = "Pending"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStatePending, ProvisioningStateSucceeded}
}
