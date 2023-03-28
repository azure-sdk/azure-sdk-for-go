//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement

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

// BenefitUtilizationSummariesClient contains the methods for the BenefitUtilizationSummaries group.
// Don't use this type directly, use NewBenefitUtilizationSummariesClient() instead.
type BenefitUtilizationSummariesClient struct {
	internal *arm.Client
}

// NewBenefitUtilizationSummariesClient creates a new instance of BenefitUtilizationSummariesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBenefitUtilizationSummariesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BenefitUtilizationSummariesClient, error) {
	cl, err := arm.NewClient(moduleName+".BenefitUtilizationSummariesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BenefitUtilizationSummariesClient{
		internal: cl,
	}
	return client, nil
}

// NewListByBillingAccountIDPager - Lists savings plan utilization summaries for the enterprise agreement scope. Supported
// at grain values: 'Daily' and 'Monthly'.
//
// Generated from API version 2023-04-01-preview
//   - billingAccountID - Billing account ID
//   - options - BenefitUtilizationSummariesClientListByBillingAccountIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager
//     method.
func (client *BenefitUtilizationSummariesClient) NewListByBillingAccountIDPager(billingAccountID string, options *BenefitUtilizationSummariesClientListByBillingAccountIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingAccountIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[BenefitUtilizationSummariesClientListByBillingAccountIDResponse]{
		More: func(page BenefitUtilizationSummariesClientListByBillingAccountIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BenefitUtilizationSummariesClientListByBillingAccountIDResponse) (BenefitUtilizationSummariesClientListByBillingAccountIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingAccountIDCreateRequest(ctx, billingAccountID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BenefitUtilizationSummariesClientListByBillingAccountIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BenefitUtilizationSummariesClientListByBillingAccountIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BenefitUtilizationSummariesClientListByBillingAccountIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingAccountIDHandleResponse(resp)
		},
	})
}

// listByBillingAccountIDCreateRequest creates the ListByBillingAccountID request.
func (client *BenefitUtilizationSummariesClient) listByBillingAccountIDCreateRequest(ctx context.Context, billingAccountID string, options *BenefitUtilizationSummariesClientListByBillingAccountIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/benefitUtilizationSummaries"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.GrainParameter != nil {
		reqQP.Set("grainParameter", string(*options.GrainParameter))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountIDHandleResponse handles the ListByBillingAccountID response.
func (client *BenefitUtilizationSummariesClient) listByBillingAccountIDHandleResponse(resp *http.Response) (BenefitUtilizationSummariesClientListByBillingAccountIDResponse, error) {
	result := BenefitUtilizationSummariesClientListByBillingAccountIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitUtilizationSummariesListResult); err != nil {
		return BenefitUtilizationSummariesClientListByBillingAccountIDResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfileIDPager - Lists savings plan utilization summaries for billing profile. Supported at grain values:
// 'Daily' and 'Monthly'.
//
// Generated from API version 2023-04-01-preview
//   - billingAccountID - Billing account ID
//   - billingProfileID - Billing profile ID.
//   - options - BenefitUtilizationSummariesClientListByBillingProfileIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager
//     method.
func (client *BenefitUtilizationSummariesClient) NewListByBillingProfileIDPager(billingAccountID string, billingProfileID string, options *BenefitUtilizationSummariesClientListByBillingProfileIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingProfileIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[BenefitUtilizationSummariesClientListByBillingProfileIDResponse]{
		More: func(page BenefitUtilizationSummariesClientListByBillingProfileIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BenefitUtilizationSummariesClientListByBillingProfileIDResponse) (BenefitUtilizationSummariesClientListByBillingProfileIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBillingProfileIDCreateRequest(ctx, billingAccountID, billingProfileID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BenefitUtilizationSummariesClientListByBillingProfileIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BenefitUtilizationSummariesClientListByBillingProfileIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BenefitUtilizationSummariesClientListByBillingProfileIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBillingProfileIDHandleResponse(resp)
		},
	})
}

// listByBillingProfileIDCreateRequest creates the ListByBillingProfileID request.
func (client *BenefitUtilizationSummariesClient) listByBillingProfileIDCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, options *BenefitUtilizationSummariesClientListByBillingProfileIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.CostManagement/benefitUtilizationSummaries"
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
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.GrainParameter != nil {
		reqQP.Set("grainParameter", string(*options.GrainParameter))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileIDHandleResponse handles the ListByBillingProfileID response.
func (client *BenefitUtilizationSummariesClient) listByBillingProfileIDHandleResponse(resp *http.Response) (BenefitUtilizationSummariesClientListByBillingProfileIDResponse, error) {
	result := BenefitUtilizationSummariesClientListByBillingProfileIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitUtilizationSummariesListResult); err != nil {
		return BenefitUtilizationSummariesClientListByBillingProfileIDResponse{}, err
	}
	return result, nil
}

// NewListBySavingsPlanIDPager - Lists the savings plan utilization summaries for daily or monthly grain.
//
// Generated from API version 2023-04-01-preview
//   - savingsPlanOrderID - Savings plan order ID.
//   - savingsPlanID - Savings plan ID.
//   - options - BenefitUtilizationSummariesClientListBySavingsPlanIDOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager
//     method.
func (client *BenefitUtilizationSummariesClient) NewListBySavingsPlanIDPager(savingsPlanOrderID string, savingsPlanID string, options *BenefitUtilizationSummariesClientListBySavingsPlanIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[BenefitUtilizationSummariesClientListBySavingsPlanIDResponse]{
		More: func(page BenefitUtilizationSummariesClientListBySavingsPlanIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BenefitUtilizationSummariesClientListBySavingsPlanIDResponse) (BenefitUtilizationSummariesClientListBySavingsPlanIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySavingsPlanIDCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BenefitUtilizationSummariesClientListBySavingsPlanIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BenefitUtilizationSummariesClientListBySavingsPlanIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BenefitUtilizationSummariesClientListBySavingsPlanIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySavingsPlanIDHandleResponse(resp)
		},
	})
}

// listBySavingsPlanIDCreateRequest creates the ListBySavingsPlanID request.
func (client *BenefitUtilizationSummariesClient) listBySavingsPlanIDCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, options *BenefitUtilizationSummariesClientListBySavingsPlanIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}/providers/Microsoft.CostManagement/benefitUtilizationSummaries"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.GrainParameter != nil {
		reqQP.Set("grainParameter", string(*options.GrainParameter))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySavingsPlanIDHandleResponse handles the ListBySavingsPlanID response.
func (client *BenefitUtilizationSummariesClient) listBySavingsPlanIDHandleResponse(resp *http.Response) (BenefitUtilizationSummariesClientListBySavingsPlanIDResponse, error) {
	result := BenefitUtilizationSummariesClientListBySavingsPlanIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitUtilizationSummariesListResult); err != nil {
		return BenefitUtilizationSummariesClientListBySavingsPlanIDResponse{}, err
	}
	return result, nil
}

// NewListBySavingsPlanOrderPager - Lists the savings plan utilization summaries for daily or monthly grain.
//
// Generated from API version 2023-04-01-preview
//   - savingsPlanOrderID - Savings plan order ID.
//   - options - BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions contains the optional parameters for the BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager
//     method.
func (client *BenefitUtilizationSummariesClient) NewListBySavingsPlanOrderPager(savingsPlanOrderID string, options *BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse] {
	return runtime.NewPager(runtime.PagingHandler[BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse]{
		More: func(page BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse) (BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySavingsPlanOrderCreateRequest(ctx, savingsPlanOrderID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySavingsPlanOrderHandleResponse(resp)
		},
	})
}

// listBySavingsPlanOrderCreateRequest creates the ListBySavingsPlanOrder request.
func (client *BenefitUtilizationSummariesClient) listBySavingsPlanOrderCreateRequest(ctx context.Context, savingsPlanOrderID string, options *BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/providers/Microsoft.CostManagement/benefitUtilizationSummaries"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.GrainParameter != nil {
		reqQP.Set("grainParameter", string(*options.GrainParameter))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySavingsPlanOrderHandleResponse handles the ListBySavingsPlanOrder response.
func (client *BenefitUtilizationSummariesClient) listBySavingsPlanOrderHandleResponse(resp *http.Response) (BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse, error) {
	result := BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BenefitUtilizationSummariesListResult); err != nil {
		return BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse{}, err
	}
	return result, nil
}
