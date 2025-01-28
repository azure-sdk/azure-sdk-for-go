//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ReservationTransactionsClient contains the methods for the ReservationTransactions group.
// Don't use this type directly, use NewReservationTransactionsClient() instead.
type ReservationTransactionsClient struct {
	internal *arm.Client
}

// NewReservationTransactionsClient creates a new instance of ReservationTransactionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReservationTransactionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationTransactionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationTransactionsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - List of transactions for reserved instances on billing account scope. Note: The refund transactions are
// posted along with its purchase transaction (i.e. in the purchase billing month). For example,
// The refund is requested in May 2021. This refund transaction will have event date as May 2021 but the billing month as
// April 2020 when the reservation purchase was made. Note: ARM has a payload size
// limit of 12MB, so currently callers get 400 when the response size exceeds the ARM limit. In such cases, API call should
// be made with smaller date ranges.
//
// Generated from API version 2024-08-01
//   - billingAccountID - BillingAccount ID
//   - options - ReservationTransactionsClientListOptions contains the optional parameters for the ReservationTransactionsClient.NewListPager
//     method.
func (client *ReservationTransactionsClient) NewListPager(billingAccountID string, options *ReservationTransactionsClientListOptions) *runtime.Pager[ReservationTransactionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationTransactionsClientListResponse]{
		More: func(page ReservationTransactionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationTransactionsClientListResponse) (ReservationTransactionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReservationTransactionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, billingAccountID, options)
			}, nil)
			if err != nil {
				return ReservationTransactionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReservationTransactionsClient) listCreateRequest(ctx context.Context, billingAccountID string, options *ReservationTransactionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/reservationTransactions"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-08-01")
	if options != nil && options.PreviewMarkupPercentage != nil {
		reqQP.Set("previewMarkupPercentage", strconv.FormatFloat(*options.PreviewMarkupPercentage, 'f', -1, 64))
	}
	if options != nil && options.UseMarkupIfPartner != nil {
		reqQP.Set("useMarkupIfPartner", strconv.FormatBool(*options.UseMarkupIfPartner))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReservationTransactionsClient) listHandleResponse(resp *http.Response) (ReservationTransactionsClientListResponse, error) {
	result := ReservationTransactionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationTransactionsListResult); err != nil {
		return ReservationTransactionsClientListResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - List of transactions for reserved instances on billing profile scope. The refund transactions
// are posted along with its purchase transaction (i.e. in the purchase billing month). For example, The
// refund is requested in May 2021. This refund transaction will have event date as May 2021 but the billing month as April
// 2020 when the reservation purchase was made. Note: ARM has a payload size limit
// of 12MB, so currently callers get 400 when the response size exceeds the ARM limit. In such cases, API call should be made
// with smaller date ranges.
//
// Generated from API version 2024-08-01
//   - billingAccountID - BillingAccount ID
//   - billingProfileID - Azure Billing Profile ID.
//   - options - ReservationTransactionsClientListByBillingProfileOptions contains the optional parameters for the ReservationTransactionsClient.NewListByBillingProfilePager
//     method.
func (client *ReservationTransactionsClient) NewListByBillingProfilePager(billingAccountID string, billingProfileID string, options *ReservationTransactionsClientListByBillingProfileOptions) *runtime.Pager[ReservationTransactionsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationTransactionsClientListByBillingProfileResponse]{
		More: func(page ReservationTransactionsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationTransactionsClientListByBillingProfileResponse) (ReservationTransactionsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReservationTransactionsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountID, billingProfileID, options)
			}, nil)
			if err != nil {
				return ReservationTransactionsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *ReservationTransactionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, options *ReservationTransactionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/reservationTransactions"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *ReservationTransactionsClient) listByBillingProfileHandleResponse(resp *http.Response) (ReservationTransactionsClientListByBillingProfileResponse, error) {
	result := ReservationTransactionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModernReservationTransactionsListResult); err != nil {
		return ReservationTransactionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}
