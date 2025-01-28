//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

// AggregatedCostClientGetByManagementGroupResponse contains the response from method AggregatedCostClient.GetByManagementGroup.
type AggregatedCostClientGetByManagementGroupResponse struct {
	// A management group aggregated cost resource.
	ManagementGroupAggregatedCostResult
}

// AggregatedCostClientGetForBillingPeriodByManagementGroupResponse contains the response from method AggregatedCostClient.GetForBillingPeriodByManagementGroup.
type AggregatedCostClientGetForBillingPeriodByManagementGroupResponse struct {
	// A management group aggregated cost resource.
	ManagementGroupAggregatedCostResult
}

// BalancesClientGetByBillingAccountResponse contains the response from method BalancesClient.GetByBillingAccount.
type BalancesClientGetByBillingAccountResponse struct {
	// A balance resource.
	Balance
}

// BalancesClientGetForBillingPeriodByBillingAccountResponse contains the response from method BalancesClient.GetForBillingPeriodByBillingAccount.
type BalancesClientGetForBillingPeriodByBillingAccountResponse struct {
	// A balance resource.
	Balance
}

// BudgetsClientCreateOrUpdateResponse contains the response from method BudgetsClient.CreateOrUpdate.
type BudgetsClientCreateOrUpdateResponse struct {
	// A budget resource.
	Budget
}

// BudgetsClientDeleteResponse contains the response from method BudgetsClient.Delete.
type BudgetsClientDeleteResponse struct {
	// placeholder for future response values
}

// BudgetsClientGetResponse contains the response from method BudgetsClient.Get.
type BudgetsClientGetResponse struct {
	// A budget resource.
	Budget
}

// BudgetsClientListResponse contains the response from method BudgetsClient.NewListPager.
type BudgetsClientListResponse struct {
	// Result of listing budgets. It contains a list of available budgets in the scope provided.
	BudgetsListResult
}

// ChargesClientListResponse contains the response from method ChargesClient.List.
type ChargesClientListResponse struct {
	// Result of listing charge summary.
	ChargesListResult
}

// CreditsClientGetResponse contains the response from method CreditsClient.Get.
type CreditsClientGetResponse struct {
	// A credit summary resource.
	CreditSummary
}

// EventsClientListByBillingAccountResponse contains the response from method EventsClient.NewListByBillingAccountPager.
type EventsClientListByBillingAccountResponse struct {
	// Result of listing event summary.
	Events
}

// EventsClientListByBillingProfileResponse contains the response from method EventsClient.NewListByBillingProfilePager.
type EventsClientListByBillingProfileResponse struct {
	// Result of listing event summary.
	Events
}

// LotsClientListByBillingAccountResponse contains the response from method LotsClient.NewListByBillingAccountPager.
type LotsClientListByBillingAccountResponse struct {
	// Result of listing lot summary.
	Lots
}

// LotsClientListByBillingProfileResponse contains the response from method LotsClient.NewListByBillingProfilePager.
type LotsClientListByBillingProfileResponse struct {
	// Result of listing lot summary.
	Lots
}

// LotsClientListByCustomerResponse contains the response from method LotsClient.NewListByCustomerPager.
type LotsClientListByCustomerResponse struct {
	// Result of listing lot summary.
	Lots
}

// MarketplacesClientListResponse contains the response from method MarketplacesClient.NewListPager.
type MarketplacesClientListResponse struct {
	// Result of listing marketplaces. It contains a list of available marketplaces in reverse chronological order by billing
	// period.
	MarketplacesListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of listing consumption operations. It contains a list of operations and a URL link to get the next set of results.
	OperationListResult
}

// PriceSheetClientDownloadByBillingAccountPeriodResponse contains the response from method PriceSheetClient.BeginDownloadByBillingAccountPeriod.
type PriceSheetClientDownloadByBillingAccountPeriodResponse struct {
	// The status of the long running operation.
	OperationStatus
}

// PriceSheetClientGetByBillingPeriodResponse contains the response from method PriceSheetClient.GetByBillingPeriod.
type PriceSheetClientGetByBillingPeriodResponse struct {
	// An pricesheet resource.
	PriceSheetResult
}

// PriceSheetClientGetResponse contains the response from method PriceSheetClient.Get.
type PriceSheetClientGetResponse struct {
	// An pricesheet resource.
	PriceSheetResult
}

// ReservationRecommendationDetailsClientGetResponse contains the response from method ReservationRecommendationDetailsClient.Get.
type ReservationRecommendationDetailsClientGetResponse struct {
	// Reservation recommendation details.
	ReservationRecommendationDetailsModel
}

// ReservationRecommendationsClientListResponse contains the response from method ReservationRecommendationsClient.NewListPager.
type ReservationRecommendationsClientListResponse struct {
	// Result of listing reservation recommendations.
	ReservationRecommendationsListResult
}

// ReservationTransactionsClientListByBillingProfileResponse contains the response from method ReservationTransactionsClient.NewListByBillingProfilePager.
type ReservationTransactionsClientListByBillingProfileResponse struct {
	// Result of listing reservation recommendations.
	ModernReservationTransactionsListResult
}

// ReservationTransactionsClientListResponse contains the response from method ReservationTransactionsClient.NewListPager.
type ReservationTransactionsClientListResponse struct {
	// Result of listing reservation recommendations.
	ReservationTransactionsListResult
}

// ReservationsDetailsClientListByReservationOrderAndReservationResponse contains the response from method ReservationsDetailsClient.NewListByReservationOrderAndReservationPager.
type ReservationsDetailsClientListByReservationOrderAndReservationResponse struct {
	// Result of listing reservation details.
	ReservationDetailsListResult
}

// ReservationsDetailsClientListByReservationOrderResponse contains the response from method ReservationsDetailsClient.NewListByReservationOrderPager.
type ReservationsDetailsClientListByReservationOrderResponse struct {
	// Result of listing reservation details.
	ReservationDetailsListResult
}

// ReservationsDetailsClientListResponse contains the response from method ReservationsDetailsClient.NewListPager.
type ReservationsDetailsClientListResponse struct {
	// Result of listing reservation details.
	ReservationDetailsListResult
}

// ReservationsSummariesClientListByReservationOrderAndReservationResponse contains the response from method ReservationsSummariesClient.NewListByReservationOrderAndReservationPager.
type ReservationsSummariesClientListByReservationOrderAndReservationResponse struct {
	// Result of listing reservation summaries.
	ReservationSummariesListResult
}

// ReservationsSummariesClientListByReservationOrderResponse contains the response from method ReservationsSummariesClient.NewListByReservationOrderPager.
type ReservationsSummariesClientListByReservationOrderResponse struct {
	// Result of listing reservation summaries.
	ReservationSummariesListResult
}

// ReservationsSummariesClientListResponse contains the response from method ReservationsSummariesClient.NewListPager.
type ReservationsSummariesClientListResponse struct {
	// Result of listing reservation summaries.
	ReservationSummariesListResult
}

// TagsClientGetResponse contains the response from method TagsClient.Get.
type TagsClientGetResponse struct {
	// A resource listing all tags.
	TagsResult
}

// UsageDetailsClientListResponse contains the response from method UsageDetailsClient.NewListPager.
type UsageDetailsClientListResponse struct {
	// Result of listing usage details. It contains a list of available usage details in reverse chronological order by billing
	// period.
	UsageDetailsListResult
}
