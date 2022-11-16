//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconsumption

// DownloadProperties - The properties of the price sheet download.
type DownloadProperties struct {
	// READ-ONLY; The link (url) to download the pricesheet.
	DownloadURL *string `json:"downloadUrl,omitempty" azure:"ro"`

	// READ-ONLY; Download link validity.
	ValidTill *string `json:"validTill,omitempty" azure:"ro"`
}

// ErrorDetails - The details of the error.
type ErrorDetails struct {
	// READ-ONLY; Error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message indicating why the operation failed.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request. The reason is provided
// in the error message.
// Some Error responses:
// * 429 TooManyRequests - Request is throttled. Retry after waiting for the time specified in the "x-ms-ratelimit-microsoft.consumption-retry-after"
// header.
//
// * 503 ServiceUnavailable - Service is temporarily unavailable. Retry after waiting for the time specified in the "Retry-After"
// header.
type ErrorResponse struct {
	// The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// MeterDetails - The properties of the meter detail.
type MeterDetails struct {
	// READ-ONLY; The category of the meter, for example, 'Cloud services', 'Networking', etc..
	MeterCategory *string `json:"meterCategory,omitempty" azure:"ro"`

	// READ-ONLY; The location in which the Azure service is available.
	MeterLocation *string `json:"meterLocation,omitempty" azure:"ro"`

	// READ-ONLY; The name of the meter, within the given meter category
	MeterName *string `json:"meterName,omitempty" azure:"ro"`

	// READ-ONLY; The subcategory of the meter, for example, 'A6 Cloud services', 'ExpressRoute (IXP)', etc..
	MeterSubCategory *string `json:"meterSubCategory,omitempty" azure:"ro"`

	// READ-ONLY; The pretax listing price.
	PretaxStandardRate *float64 `json:"pretaxStandardRate,omitempty" azure:"ro"`

	// READ-ONLY; The name of the service.
	ServiceName *string `json:"serviceName,omitempty" azure:"ro"`

	// READ-ONLY; The service tier.
	ServiceTier *string `json:"serviceTier,omitempty" azure:"ro"`

	// READ-ONLY; The total included quantity associated with the offer.
	TotalIncludedQuantity *float64 `json:"totalIncludedQuantity,omitempty" azure:"ro"`

	// READ-ONLY; The unit in which the meter consumption is charged, for example, 'Hours', 'GB', etc.
	Unit *string `json:"unit,omitempty" azure:"ro"`
}

// Operation - A Consumption REST API operation.
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Operation Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty" azure:"ro"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Description of the operation.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Service provider: Microsoft.Consumption.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource on which the operation is performed: UsageDetail, etc.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - Result of listing consumption operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of consumption operations supported by the Microsoft.Consumption resource provider.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PriceSheetClientGetByBillingPeriodOptions contains the optional parameters for the PriceSheetClient.GetByBillingPeriod
// method.
type PriceSheetClientGetByBillingPeriodOptions struct {
	// May be used to expand the properties/meterDetails within a price sheet. By default, these fields are not included when
	// returning price sheet.
	Expand *string
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results to the top N results.
	Top *int32
}

// PriceSheetClientGetOptions contains the optional parameters for the PriceSheetClient.Get method.
type PriceSheetClientGetOptions struct {
	// May be used to expand the properties/meterDetails within a price sheet. By default, these fields are not included when
	// returning price sheet.
	Expand *string
	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string
	// May be used to limit the number of results to the top N results.
	Top *int32
}

// PriceSheetModel - price sheet result. It contains the pricesheet associated with billing period
type PriceSheetModel struct {
	// READ-ONLY; Pricesheet download details.
	Download *MeterDetails `json:"download,omitempty" azure:"ro"`

	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Price sheet
	Pricesheets []*PriceSheetProperties `json:"pricesheets,omitempty" azure:"ro"`
}

// PriceSheetProperties - The properties of the price sheet.
type PriceSheetProperties struct {
	// READ-ONLY; The id of the billing period resource that the usage belongs to.
	BillingPeriodID *string `json:"billingPeriodId,omitempty" azure:"ro"`

	// READ-ONLY; Currency Code
	CurrencyCode *string `json:"currencyCode,omitempty" azure:"ro"`

	// READ-ONLY; Included quality for an offer
	IncludedQuantity *float64 `json:"includedQuantity,omitempty" azure:"ro"`

	// READ-ONLY; The details about the meter. By default this is not populated, unless it's specified in $expand.
	MeterDetails *MeterDetails `json:"meterDetails,omitempty" azure:"ro"`

	// READ-ONLY; The meter id (GUID)
	MeterID *string `json:"meterId,omitempty" azure:"ro"`

	// READ-ONLY; Offer Id
	OfferID *string `json:"offerId,omitempty" azure:"ro"`

	// READ-ONLY; Part Number
	PartNumber *string `json:"partNumber,omitempty" azure:"ro"`

	// READ-ONLY; Unit of measure
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty" azure:"ro"`

	// READ-ONLY; Unit Price
	UnitPrice *float64 `json:"unitPrice,omitempty" azure:"ro"`
}

// PriceSheetResult - An pricesheet resource.
type PriceSheetResult struct {
	// price sheet result. It contains the pricesheet associated with billing period
	Properties *PriceSheetModel `json:"properties,omitempty"`

	// READ-ONLY; The etag for the resource.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; The full qualified ARM ID of an event.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The ID that uniquely identifies an event.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - The Resource model definition.
type Resource struct {
	// READ-ONLY; The etag for the resource.
	Etag *string `json:"etag,omitempty" azure:"ro"`

	// READ-ONLY; The full qualified ARM ID of an event.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The ID that uniquely identifies an event.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource tags.
	Tags map[string]*string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}
