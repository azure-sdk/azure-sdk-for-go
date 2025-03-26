// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
	moduleVersion = "v2.0.0"
)

// SubscriptionFeatureRegistrationApprovalType - The feature approval type.
type SubscriptionFeatureRegistrationApprovalType string

const (
	SubscriptionFeatureRegistrationApprovalTypeApprovalRequired SubscriptionFeatureRegistrationApprovalType = "ApprovalRequired"
	SubscriptionFeatureRegistrationApprovalTypeAutoApproval     SubscriptionFeatureRegistrationApprovalType = "AutoApproval"
	SubscriptionFeatureRegistrationApprovalTypeNotSpecified     SubscriptionFeatureRegistrationApprovalType = "NotSpecified"
)

// PossibleSubscriptionFeatureRegistrationApprovalTypeValues returns the possible values for the SubscriptionFeatureRegistrationApprovalType const type.
func PossibleSubscriptionFeatureRegistrationApprovalTypeValues() []SubscriptionFeatureRegistrationApprovalType {
	return []SubscriptionFeatureRegistrationApprovalType{
		SubscriptionFeatureRegistrationApprovalTypeApprovalRequired,
		SubscriptionFeatureRegistrationApprovalTypeAutoApproval,
		SubscriptionFeatureRegistrationApprovalTypeNotSpecified,
	}
}

// SubscriptionFeatureRegistrationState - The state.
type SubscriptionFeatureRegistrationState string

const (
	SubscriptionFeatureRegistrationStateNotRegistered SubscriptionFeatureRegistrationState = "NotRegistered"
	SubscriptionFeatureRegistrationStateNotSpecified  SubscriptionFeatureRegistrationState = "NotSpecified"
	SubscriptionFeatureRegistrationStatePending       SubscriptionFeatureRegistrationState = "Pending"
	SubscriptionFeatureRegistrationStateRegistered    SubscriptionFeatureRegistrationState = "Registered"
	SubscriptionFeatureRegistrationStateUnregistered  SubscriptionFeatureRegistrationState = "Unregistered"
)

// PossibleSubscriptionFeatureRegistrationStateValues returns the possible values for the SubscriptionFeatureRegistrationState const type.
func PossibleSubscriptionFeatureRegistrationStateValues() []SubscriptionFeatureRegistrationState {
	return []SubscriptionFeatureRegistrationState{
		SubscriptionFeatureRegistrationStateNotRegistered,
		SubscriptionFeatureRegistrationStateNotSpecified,
		SubscriptionFeatureRegistrationStatePending,
		SubscriptionFeatureRegistrationStateRegistered,
		SubscriptionFeatureRegistrationStateUnregistered,
	}
}
