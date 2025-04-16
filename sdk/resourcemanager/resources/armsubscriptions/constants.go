// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscriptions

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	moduleVersion = "v1.3.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// LocationType - The location type.
type LocationType string

const (
	LocationTypeEdgeZone LocationType = "EdgeZone"
	LocationTypeRegion   LocationType = "Region"
)

// PossibleLocationTypeValues returns the possible values for the LocationType const type.
func PossibleLocationTypeValues() []LocationType {
	return []LocationType{
		LocationTypeEdgeZone,
		LocationTypeRegion,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
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

// RegionCategory - The category of the region.
type RegionCategory string

const (
	RegionCategoryExtended    RegionCategory = "Extended"
	RegionCategoryOther       RegionCategory = "Other"
	RegionCategoryRecommended RegionCategory = "Recommended"
)

// PossibleRegionCategoryValues returns the possible values for the RegionCategory const type.
func PossibleRegionCategoryValues() []RegionCategory {
	return []RegionCategory{
		RegionCategoryExtended,
		RegionCategoryOther,
		RegionCategoryRecommended,
	}
}

// RegionType - The type of the region.
type RegionType string

const (
	RegionTypeLogical  RegionType = "Logical"
	RegionTypePhysical RegionType = "Physical"
)

// PossibleRegionTypeValues returns the possible values for the RegionType const type.
func PossibleRegionTypeValues() []RegionType {
	return []RegionType{
		RegionTypeLogical,
		RegionTypePhysical,
	}
}

// ResourceNameStatus - Is the resource name Allowed or Reserved
type ResourceNameStatus string

const (
	ResourceNameStatusAllowed  ResourceNameStatus = "Allowed"
	ResourceNameStatusReserved ResourceNameStatus = "Reserved"
)

// PossibleResourceNameStatusValues returns the possible values for the ResourceNameStatus const type.
func PossibleResourceNameStatusValues() []ResourceNameStatus {
	return []ResourceNameStatus{
		ResourceNameStatusAllowed,
		ResourceNameStatusReserved,
	}
}

// SpendingLimit - The subscription spending limit.
type SpendingLimit string

const (
	SpendingLimitCurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
	SpendingLimitOff              SpendingLimit = "Off"
	SpendingLimitOn               SpendingLimit = "On"
)

// PossibleSpendingLimitValues returns the possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{
		SpendingLimitCurrentPeriodOff,
		SpendingLimitOff,
		SpendingLimitOn,
	}
}

// SubscriptionState - The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
type SubscriptionState string

const (
	SubscriptionStateDeleted  SubscriptionState = "Deleted"
	SubscriptionStateDisabled SubscriptionState = "Disabled"
	SubscriptionStateEnabled  SubscriptionState = "Enabled"
	SubscriptionStatePastDue  SubscriptionState = "PastDue"
	SubscriptionStateWarned   SubscriptionState = "Warned"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateDeleted,
		SubscriptionStateDisabled,
		SubscriptionStateEnabled,
		SubscriptionStatePastDue,
		SubscriptionStateWarned,
	}
}

// TenantCategory - Category of the tenant.
type TenantCategory string

const (
	TenantCategoryHome        TenantCategory = "Home"
	TenantCategoryManagedBy   TenantCategory = "ManagedBy"
	TenantCategoryProjectedBy TenantCategory = "ProjectedBy"
)

// PossibleTenantCategoryValues returns the possible values for the TenantCategory const type.
func PossibleTenantCategoryValues() []TenantCategory {
	return []TenantCategory{
		TenantCategoryHome,
		TenantCategoryManagedBy,
		TenantCategoryProjectedBy,
	}
}
