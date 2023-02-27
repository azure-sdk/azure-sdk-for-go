//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armprofessionalservice

const (
	moduleName    = "armprofessionalservice"
	moduleVersion = "v0.1.0"
)

// PaymentChannelType - The Payment channel for the ProfessionalServiceSubscription.
type PaymentChannelType string

const (
	PaymentChannelTypeCustomerDelegated     PaymentChannelType = "CustomerDelegated"
	PaymentChannelTypeSubscriptionDelegated PaymentChannelType = "SubscriptionDelegated"
)

// PossiblePaymentChannelTypeValues returns the possible values for the PaymentChannelType const type.
func PossiblePaymentChannelTypeValues() []PaymentChannelType {
	return []PaymentChannelType{
		PaymentChannelTypeCustomerDelegated,
		PaymentChannelTypeSubscriptionDelegated,
	}
}

// ProfessionalServiceResourceStatus - The ProfessionalService Subscription Status.
type ProfessionalServiceResourceStatus string

const (
	ProfessionalServiceResourceStatusNotStarted              ProfessionalServiceResourceStatus = "NotStarted"
	ProfessionalServiceResourceStatusPendingFulfillmentStart ProfessionalServiceResourceStatus = "PendingFulfillmentStart"
	ProfessionalServiceResourceStatusSubscribed              ProfessionalServiceResourceStatus = "Subscribed"
	ProfessionalServiceResourceStatusSuspended               ProfessionalServiceResourceStatus = "Suspended"
	ProfessionalServiceResourceStatusUnsubscribed            ProfessionalServiceResourceStatus = "Unsubscribed"
)

// PossibleProfessionalServiceResourceStatusValues returns the possible values for the ProfessionalServiceResourceStatus const type.
func PossibleProfessionalServiceResourceStatusValues() []ProfessionalServiceResourceStatus {
	return []ProfessionalServiceResourceStatus{
		ProfessionalServiceResourceStatusNotStarted,
		ProfessionalServiceResourceStatusPendingFulfillmentStart,
		ProfessionalServiceResourceStatusSubscribed,
		ProfessionalServiceResourceStatusSuspended,
		ProfessionalServiceResourceStatusUnsubscribed,
	}
}
