//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReservationsClient contains the methods for the Reservations group.
// Don't use this type directly, use NewReservationsClient() instead.
type ReservationsClient struct {
	internal *arm.Client
}

// NewReservationsClient creates a new instance of ReservationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReservationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReservationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationsClient{
		internal: cl,
	}
	return client, nil
}

// NewListByBillingAccountPager - Lists the reservations for a billing account and the roll up counts of reservations group
// by provisioning states.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - ReservationsClientListByBillingAccountOptions contains the optional parameters for the ReservationsClient.NewListByBillingAccountPager
//     method.
func (client *ReservationsClient) NewListByBillingAccountPager(billingAccountName string, options *ReservationsClientListByBillingAccountOptions) *runtime.Pager[ReservationsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsClientListByBillingAccountResponse]{
		More: func(page ReservationsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsClientListByBillingAccountResponse) (ReservationsClientListByBillingAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsClientListByBillingAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReservationsClientListByBillingAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsClientListByBillingAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *ReservationsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *ReservationsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/reservations"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.RefreshSummary != nil {
		reqQP.Set("refreshSummary", *options.RefreshSummary)
	}
	if options != nil && options.SelectedState != nil {
		reqQP.Set("selectedState", *options.SelectedState)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *ReservationsClient) listByBillingAccountHandleResponse(resp *http.Response) (ReservationsClientListByBillingAccountResponse, error) {
	result := ReservationsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationsListResult); err != nil {
		return ReservationsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the reservations for a billing profile and the roll up counts of reservations group
// by provisioning state.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - ReservationsClientListByBillingProfileOptions contains the optional parameters for the ReservationsClient.NewListByBillingProfilePager
//     method.
func (client *ReservationsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *ReservationsClientListByBillingProfileOptions) *runtime.Pager[ReservationsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReservationsClientListByBillingProfileResponse]{
		More: func(page ReservationsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReservationsClientListByBillingProfileResponse) (ReservationsClientListByBillingProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReservationsClientListByBillingProfileResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReservationsClientListByBillingProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReservationsClientListByBillingProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *ReservationsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *ReservationsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/reservations"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.RefreshSummary != nil {
		reqQP.Set("refreshSummary", *options.RefreshSummary)
	}
	if options != nil && options.SelectedState != nil {
		reqQP.Set("selectedState", *options.SelectedState)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *ReservationsClient) listByBillingProfileHandleResponse(resp *http.Response) (ReservationsClientListByBillingProfileResponse, error) {
	result := ReservationsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationsListResult); err != nil {
		return ReservationsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}
