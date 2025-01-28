//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	moduleVersion = "v2.0.0"
)

// BillingFrequency - The billing frequency.
type BillingFrequency string

const (
	BillingFrequencyMonth   BillingFrequency = "Month"
	BillingFrequencyQuarter BillingFrequency = "Quarter"
	BillingFrequencyYear    BillingFrequency = "Year"
)

// PossibleBillingFrequencyValues returns the possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{
		BillingFrequencyMonth,
		BillingFrequencyQuarter,
		BillingFrequencyYear,
	}
}

// BudgetOperatorType - The operator to use for comparison.
type BudgetOperatorType string

const (
	BudgetOperatorTypeIn BudgetOperatorType = "In"
)

// PossibleBudgetOperatorTypeValues returns the possible values for the BudgetOperatorType const type.
func PossibleBudgetOperatorTypeValues() []BudgetOperatorType {
	return []BudgetOperatorType{
		BudgetOperatorTypeIn,
	}
}

// CategoryType - The category of the budget, whether the budget tracks cost or usage.
type CategoryType string

const (
	CategoryTypeCost CategoryType = "Cost"
)

// PossibleCategoryTypeValues returns the possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{
		CategoryTypeCost,
	}
}

// ChargeSummaryKind - Specifies the kind of charge summary.
type ChargeSummaryKind string

const (
	ChargeSummaryKindLegacy ChargeSummaryKind = "legacy"
	ChargeSummaryKindModern ChargeSummaryKind = "modern"
)

// PossibleChargeSummaryKindValues returns the possible values for the ChargeSummaryKind const type.
func PossibleChargeSummaryKindValues() []ChargeSummaryKind {
	return []ChargeSummaryKind{
		ChargeSummaryKindLegacy,
		ChargeSummaryKindModern,
	}
}

// CultureCode - Language in which the recipient will receive the notification
type CultureCode string

const (
	CultureCodeCsCz CultureCode = "cs-cz"
	CultureCodeDaDk CultureCode = "da-dk"
	CultureCodeDeDe CultureCode = "de-de"
	CultureCodeEnGb CultureCode = "en-gb"
	CultureCodeEnUs CultureCode = "en-us"
	CultureCodeEsEs CultureCode = "es-es"
	CultureCodeFrFr CultureCode = "fr-fr"
	CultureCodeHuHu CultureCode = "hu-hu"
	CultureCodeItIt CultureCode = "it-it"
	CultureCodeJaJp CultureCode = "ja-jp"
	CultureCodeKoKr CultureCode = "ko-kr"
	CultureCodeNbNo CultureCode = "nb-no"
	CultureCodeNlNl CultureCode = "nl-nl"
	CultureCodePlPl CultureCode = "pl-pl"
	CultureCodePtBr CultureCode = "pt-br"
	CultureCodePtPt CultureCode = "pt-pt"
	CultureCodeRuRu CultureCode = "ru-ru"
	CultureCodeSvSe CultureCode = "sv-se"
	CultureCodeTrTr CultureCode = "tr-tr"
	CultureCodeZhCn CultureCode = "zh-cn"
	CultureCodeZhTw CultureCode = "zh-tw"
)

// PossibleCultureCodeValues returns the possible values for the CultureCode const type.
func PossibleCultureCodeValues() []CultureCode {
	return []CultureCode{
		CultureCodeCsCz,
		CultureCodeDaDk,
		CultureCodeDeDe,
		CultureCodeEnGb,
		CultureCodeEnUs,
		CultureCodeEsEs,
		CultureCodeFrFr,
		CultureCodeHuHu,
		CultureCodeItIt,
		CultureCodeJaJp,
		CultureCodeKoKr,
		CultureCodeNbNo,
		CultureCodeNlNl,
		CultureCodePlPl,
		CultureCodePtBr,
		CultureCodePtPt,
		CultureCodeRuRu,
		CultureCodeSvSe,
		CultureCodeTrTr,
		CultureCodeZhCn,
		CultureCodeZhTw,
	}
}

type Datagrain string

const (
	// DatagrainDailyGrain - Daily grain of data
	DatagrainDailyGrain Datagrain = "daily"
	// DatagrainMonthlyGrain - Monthly grain of data
	DatagrainMonthlyGrain Datagrain = "monthly"
)

// PossibleDatagrainValues returns the possible values for the Datagrain const type.
func PossibleDatagrainValues() []Datagrain {
	return []Datagrain{
		DatagrainDailyGrain,
		DatagrainMonthlyGrain,
	}
}

// EventType - Identifies the type of the event.
type EventType string

const (
	EventTypeCreditExpired        EventType = "CreditExpired"
	EventTypeNewCredit            EventType = "NewCredit"
	EventTypePendingAdjustments   EventType = "PendingAdjustments"
	EventTypePendingCharges       EventType = "PendingCharges"
	EventTypePendingExpiredCredit EventType = "PendingExpiredCredit"
	EventTypePendingNewCredit     EventType = "PendingNewCredit"
	EventTypeSettledCharges       EventType = "SettledCharges"
	EventTypeUnKnown              EventType = "UnKnown"
)

// PossibleEventTypeValues returns the possible values for the EventType const type.
func PossibleEventTypeValues() []EventType {
	return []EventType{
		EventTypeCreditExpired,
		EventTypeNewCredit,
		EventTypePendingAdjustments,
		EventTypePendingCharges,
		EventTypePendingExpiredCredit,
		EventTypePendingNewCredit,
		EventTypeSettledCharges,
		EventTypeUnKnown,
	}
}

type LookBackPeriod string

const (
	// LookBackPeriodLast07Days - Use 7 days of data for recommendations
	LookBackPeriodLast07Days LookBackPeriod = "Last7Days"
	// LookBackPeriodLast30Days - Use 30 days of data for recommendations
	LookBackPeriodLast30Days LookBackPeriod = "Last30Days"
	// LookBackPeriodLast60Days - Use 60 days of data for recommendations
	LookBackPeriodLast60Days LookBackPeriod = "Last60Days"
)

// PossibleLookBackPeriodValues returns the possible values for the LookBackPeriod const type.
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return []LookBackPeriod{
		LookBackPeriodLast07Days,
		LookBackPeriodLast30Days,
		LookBackPeriodLast60Days,
	}
}

// LotSource - The source of the lot.
type LotSource string

const (
	LotSourceConsumptionCommitment LotSource = "ConsumptionCommitment"
	LotSourcePromotionalCredit     LotSource = "PromotionalCredit"
	LotSourcePurchasedCredit       LotSource = "PurchasedCredit"
)

// PossibleLotSourceValues returns the possible values for the LotSource const type.
func PossibleLotSourceValues() []LotSource {
	return []LotSource{
		LotSourceConsumptionCommitment,
		LotSourcePromotionalCredit,
		LotSourcePurchasedCredit,
	}
}

type Metrictype string

const (
	// MetrictypeActualCostMetricType - Actual cost data.
	MetrictypeActualCostMetricType Metrictype = "actualcost"
	// MetrictypeAmortizedCostMetricType - Amortized cost data.
	MetrictypeAmortizedCostMetricType Metrictype = "amortizedcost"
	// MetrictypeUsageMetricType - Usage data.
	MetrictypeUsageMetricType Metrictype = "usage"
)

// PossibleMetrictypeValues returns the possible values for the Metrictype const type.
func PossibleMetrictypeValues() []Metrictype {
	return []Metrictype{
		MetrictypeActualCostMetricType,
		MetrictypeAmortizedCostMetricType,
		MetrictypeUsageMetricType,
	}
}

// OperationStatusType - The status of the long running operation.
type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

// PossibleOperationStatusTypeValues returns the possible values for the OperationStatusType const type.
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return []OperationStatusType{
		OperationStatusTypeCompleted,
		OperationStatusTypeFailed,
		OperationStatusTypeRunning,
	}
}

// OperatorType - The comparison operator.
type OperatorType string

const (
	// OperatorTypeEqualTo - Alert will be triggered if the evaluated cost is the same as threshold value. Note: It’s not recommended
	// to use this OperatorType as there’s low chance of cost being exactly the same as threshold value, leading to missing of
	// your alert. This OperatorType will be deprecated in future.
	OperatorTypeEqualTo OperatorType = "EqualTo"
	// OperatorTypeGreaterThan - Alert will be triggered if the evaluated cost is greater than the threshold value. Note: This
	// is the recommended OperatorType while configuring Budget Alert.
	OperatorTypeGreaterThan OperatorType = "GreaterThan"
	// OperatorTypeGreaterThanOrEqualTo - Alert will be triggered if the evaluated cost is greater than or equal to the threshold
	// value.
	OperatorTypeGreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns the possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{
		OperatorTypeEqualTo,
		OperatorTypeGreaterThan,
		OperatorTypeGreaterThanOrEqualTo,
	}
}

// OrganizationType - The organization type of the lot.
type OrganizationType string

const (
	// OrganizationTypeContributorOrganizationType - Contributor organization type for Multi-Entity consumption commitment.
	OrganizationTypeContributorOrganizationType OrganizationType = "Contributor"
	// OrganizationTypePrimaryOrganizationType - Primary organization type for Multi-Entity consumption commitment.
	OrganizationTypePrimaryOrganizationType OrganizationType = "Primary"
)

// PossibleOrganizationTypeValues returns the possible values for the OrganizationType const type.
func PossibleOrganizationTypeValues() []OrganizationType {
	return []OrganizationType{
		OrganizationTypeContributorOrganizationType,
		OrganizationTypePrimaryOrganizationType,
	}
}

// PricingModelType - Identifier that indicates how the meter is priced.
type PricingModelType string

const (
	PricingModelTypeOnDemand    PricingModelType = "On Demand"
	PricingModelTypeReservation PricingModelType = "Reservation"
	PricingModelTypeSpot        PricingModelType = "Spot"
)

// PossiblePricingModelTypeValues returns the possible values for the PricingModelType const type.
func PossiblePricingModelTypeValues() []PricingModelType {
	return []PricingModelType{
		PricingModelTypeOnDemand,
		PricingModelTypeReservation,
		PricingModelTypeSpot,
	}
}

// ReservationRecommendationKind - Specifies the kind of reservation recommendation.
type ReservationRecommendationKind string

const (
	ReservationRecommendationKindLegacy ReservationRecommendationKind = "legacy"
	ReservationRecommendationKindModern ReservationRecommendationKind = "modern"
)

// PossibleReservationRecommendationKindValues returns the possible values for the ReservationRecommendationKind const type.
func PossibleReservationRecommendationKindValues() []ReservationRecommendationKind {
	return []ReservationRecommendationKind{
		ReservationRecommendationKindLegacy,
		ReservationRecommendationKindModern,
	}
}

type Scope string

const (
	ScopeShared Scope = "Shared"
	ScopeSingle Scope = "Single"
)

// PossibleScopeValues returns the possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{
		ScopeShared,
		ScopeSingle,
	}
}

// Status - The status of the lot.
type Status string

const (
	StatusActive   Status = "Active"
	StatusCanceled Status = "Canceled"
	StatusComplete Status = "Complete"
	StatusExpired  Status = "Expired"
	StatusInactive Status = "Inactive"
	StatusNone     Status = "None"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusActive,
		StatusCanceled,
		StatusComplete,
		StatusExpired,
		StatusInactive,
		StatusNone,
	}
}

type Term string

const (
	// TermP1M - 1 month reservation term
	TermP1M Term = "P1M"
	// TermP1Y - 1 year reservation term
	TermP1Y Term = "P1Y"
	// TermP3Y - 3 year reservation term
	TermP3Y Term = "P3Y"
)

// PossibleTermValues returns the possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{
		TermP1M,
		TermP1Y,
		TermP3Y,
	}
}

// ThresholdType - The type of threshold
type ThresholdType string

const (
	// ThresholdTypeActual - Actual costs budget alerts notify when the actual accrued cost exceeds the allocated budget .
	ThresholdTypeActual ThresholdType = "Actual"
	// ThresholdTypeForecasted - Forecasted costs budget alerts provide advanced notification that your spending trends are likely
	// to exceed your allocated budget, as it relies on forecasted cost predictions.
	ThresholdTypeForecasted ThresholdType = "Forecasted"
)

// PossibleThresholdTypeValues returns the possible values for the ThresholdType const type.
func PossibleThresholdTypeValues() []ThresholdType {
	return []ThresholdType{
		ThresholdTypeActual,
		ThresholdTypeForecasted,
	}
}

// TimeGrainType - The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth,
// BillingQuarter, and BillingAnnual are only supported by WD customers
type TimeGrainType string

const (
	TimeGrainTypeAnnually       TimeGrainType = "Annually"
	TimeGrainTypeBillingAnnual  TimeGrainType = "BillingAnnual"
	TimeGrainTypeBillingMonth   TimeGrainType = "BillingMonth"
	TimeGrainTypeBillingQuarter TimeGrainType = "BillingQuarter"
	TimeGrainTypeMonthly        TimeGrainType = "Monthly"
	TimeGrainTypeQuarterly      TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns the possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{
		TimeGrainTypeAnnually,
		TimeGrainTypeBillingAnnual,
		TimeGrainTypeBillingMonth,
		TimeGrainTypeBillingQuarter,
		TimeGrainTypeMonthly,
		TimeGrainTypeQuarterly,
	}
}

// UsageDetailsKind - Specifies the kind of usage details.
type UsageDetailsKind string

const (
	UsageDetailsKindLegacy UsageDetailsKind = "legacy"
	UsageDetailsKindModern UsageDetailsKind = "modern"
)

// PossibleUsageDetailsKindValues returns the possible values for the UsageDetailsKind const type.
func PossibleUsageDetailsKindValues() []UsageDetailsKind {
	return []UsageDetailsKind{
		UsageDetailsKindLegacy,
		UsageDetailsKindModern,
	}
}
