//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaddons

const (
	moduleName    = "armaddons"
	moduleVersion = "v1.0.0"
)

// OneTimeCharge - The one time charge status for the subscription.
type OneTimeCharge string

const (
	OneTimeChargeNo          OneTimeCharge = "no"
	OneTimeChargeOnEnabled   OneTimeCharge = "onEnabled"
	OneTimeChargeOnReenabled OneTimeCharge = "onReenabled"
)

// PossibleOneTimeChargeValues returns the possible values for the OneTimeCharge const type.
func PossibleOneTimeChargeValues() []OneTimeCharge {
	return []OneTimeCharge{
		OneTimeChargeNo,
		OneTimeChargeOnEnabled,
		OneTimeChargeOnReenabled,
	}
}

type PlanTypeName string

const (
	PlanTypeNameAdvanced  PlanTypeName = "Advanced"
	PlanTypeNameEssential PlanTypeName = "Essential"
	PlanTypeNameStandard  PlanTypeName = "Standard"
)

// PossiblePlanTypeNameValues returns the possible values for the PlanTypeName const type.
func PossiblePlanTypeNameValues() []PlanTypeName {
	return []PlanTypeName{
		PlanTypeNameAdvanced,
		PlanTypeNameEssential,
		PlanTypeNameStandard,
	}
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCancelled   ProvisioningState = "Cancelled"
	ProvisioningStateCancelling  ProvisioningState = "Cancelling"
	ProvisioningStateDowngrading ProvisioningState = "Downgrading"
	ProvisioningStateFailed      ProvisioningState = "Failed"
	ProvisioningStatePurchasing  ProvisioningState = "Purchasing"
	ProvisioningStateSucceeded   ProvisioningState = "Succeeded"
	ProvisioningStateUpgrading   ProvisioningState = "Upgrading"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCancelled,
		ProvisioningStateCancelling,
		ProvisioningStateDowngrading,
		ProvisioningStateFailed,
		ProvisioningStatePurchasing,
		ProvisioningStateSucceeded,
		ProvisioningStateUpgrading,
	}
}

// SupportPlanType - Support plan type.
type SupportPlanType string

const (
	SupportPlanTypeAdvanced  SupportPlanType = "advanced"
	SupportPlanTypeEssential SupportPlanType = "essential"
	SupportPlanTypeStandard  SupportPlanType = "standard"
)

// PossibleSupportPlanTypeValues returns the possible values for the SupportPlanType const type.
func PossibleSupportPlanTypeValues() []SupportPlanType {
	return []SupportPlanType{
		SupportPlanTypeAdvanced,
		SupportPlanTypeEssential,
		SupportPlanTypeStandard,
	}
}
