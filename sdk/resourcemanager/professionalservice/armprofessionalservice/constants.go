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
	moduleVersion = "v1.0.0"
)

// ProfessionalServiceResourceStatus - The ProfessionalService Subscription Status.
type ProfessionalServiceResourceStatus string

const (
	ProfessionalServiceResourceStatusNotStarted   ProfessionalServiceResourceStatus = "NotStarted"
	ProfessionalServiceResourceStatusSubscribed   ProfessionalServiceResourceStatus = "Subscribed"
	ProfessionalServiceResourceStatusSuspended    ProfessionalServiceResourceStatus = "Suspended"
	ProfessionalServiceResourceStatusUnsubscribed ProfessionalServiceResourceStatus = "Unsubscribed"
)

// PossibleProfessionalServiceResourceStatusValues returns the possible values for the ProfessionalServiceResourceStatus const type.
func PossibleProfessionalServiceResourceStatusValues() []ProfessionalServiceResourceStatus {
	return []ProfessionalServiceResourceStatus{
		ProfessionalServiceResourceStatusNotStarted,
		ProfessionalServiceResourceStatusSubscribed,
		ProfessionalServiceResourceStatusSuspended,
		ProfessionalServiceResourceStatusUnsubscribed,
	}
}
