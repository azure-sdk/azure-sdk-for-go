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
	"strconv"
	"strings"
	"time"
)

// TransactionsClient contains the methods for the Transactions group.
// Don't use this type directly, use NewTransactionsClient() instead.
type TransactionsClient struct {
	internal *arm.Client
}

// NewTransactionsClient creates a new instance of TransactionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTransactionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TransactionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TransactionsClient{
		internal: cl,
	}
	return client, nil
}

// GetTransactionSummaryByInvoice - Gets the transaction summary for an invoice. Transactions include purchases, refunds and
// Azure usage charges.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - invoiceName - The ID that uniquely identifies an invoice.
//   - options - TransactionsClientGetTransactionSummaryByInvoiceOptions contains the optional parameters for the TransactionsClient.GetTransactionSummaryByInvoice
//     method.
func (client *TransactionsClient) GetTransactionSummaryByInvoice(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientGetTransactionSummaryByInvoiceOptions) (TransactionsClientGetTransactionSummaryByInvoiceResponse, error) {
	var err error
	const operationName = "TransactionsClient.GetTransactionSummaryByInvoice"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTransactionSummaryByInvoiceCreateRequest(ctx, billingAccountName, invoiceName, options)
	if err != nil {
		return TransactionsClientGetTransactionSummaryByInvoiceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransactionsClientGetTransactionSummaryByInvoiceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TransactionsClientGetTransactionSummaryByInvoiceResponse{}, err
	}
	resp, err := client.getTransactionSummaryByInvoiceHandleResponse(httpResp)
	return resp, err
}

// getTransactionSummaryByInvoiceCreateRequest creates the GetTransactionSummaryByInvoice request.
func (client *TransactionsClient) getTransactionSummaryByInvoiceCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientGetTransactionSummaryByInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/transactionSummary"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTransactionSummaryByInvoiceHandleResponse handles the GetTransactionSummaryByInvoice response.
func (client *TransactionsClient) getTransactionSummaryByInvoiceHandleResponse(resp *http.Response) (TransactionsClientGetTransactionSummaryByInvoiceResponse, error) {
	result := TransactionsClientGetTransactionSummaryByInvoiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionSummary); err != nil {
		return TransactionsClientGetTransactionSummaryByInvoiceResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the billed or unbilled transactions by billing profile name for given start and end
// date. Transactions include purchases, refunds and Azure usage charges. Unbilled transactions are listed under
// pending invoice Id and do not include tax. Tax is added to the amount once an invoice is generated.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - periodStartDate - The start date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - periodEndDate - The end date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - typeParam - The type of transaction.
//   - options - TransactionsClientListByBillingProfileOptions contains the optional parameters for the TransactionsClient.NewListByBillingProfilePager
//     method.
func (client *TransactionsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByBillingProfileOptions) *runtime.Pager[TransactionsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransactionsClientListByBillingProfileResponse]{
		More: func(page TransactionsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransactionsClientListByBillingProfileResponse) (TransactionsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TransactionsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, periodStartDate, periodEndDate, typeParam, options)
			}, nil)
			if err != nil {
				return TransactionsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *TransactionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/transactions"
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
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	reqQP.Set("periodEndDate", periodEndDate.Format("2006-01-02"))
	reqQP.Set("periodStartDate", periodStartDate.Format("2006-01-02"))
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(*options.Top, 10))
	}
	reqQP.Set("type", string(typeParam))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *TransactionsClient) listByBillingProfileHandleResponse(resp *http.Response) (TransactionsClientListByBillingProfileResponse, error) {
	result := TransactionsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionListResult); err != nil {
		return TransactionsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByCustomerPager - Lists the billed or unbilled transactions by customer id for given start date and end date. Transactions
// include purchases, refunds and Azure usage charges. Unbilled transactions are listed under
// pending invoice Id and do not include tax. Tax is added to the amount once an invoice is generated.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - customerName - The ID that uniquely identifies a customer.
//   - periodStartDate - The start date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - periodEndDate - The end date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - typeParam - The type of transaction.
//   - options - TransactionsClientListByCustomerOptions contains the optional parameters for the TransactionsClient.NewListByCustomerPager
//     method.
func (client *TransactionsClient) NewListByCustomerPager(billingAccountName string, billingProfileName string, customerName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByCustomerOptions) *runtime.Pager[TransactionsClientListByCustomerResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransactionsClientListByCustomerResponse]{
		More: func(page TransactionsClientListByCustomerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransactionsClientListByCustomerResponse) (TransactionsClientListByCustomerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TransactionsClient.NewListByCustomerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCustomerCreateRequest(ctx, billingAccountName, billingProfileName, customerName, periodStartDate, periodEndDate, typeParam, options)
			}, nil)
			if err != nil {
				return TransactionsClientListByCustomerResponse{}, err
			}
			return client.listByCustomerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCustomerCreateRequest creates the ListByCustomer request.
func (client *TransactionsClient) listByCustomerCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, customerName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transactions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	reqQP.Set("periodEndDate", periodEndDate.Format("2006-01-02"))
	reqQP.Set("periodStartDate", periodStartDate.Format("2006-01-02"))
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(*options.Top, 10))
	}
	reqQP.Set("type", string(typeParam))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCustomerHandleResponse handles the ListByCustomer response.
func (client *TransactionsClient) listByCustomerHandleResponse(resp *http.Response) (TransactionsClientListByCustomerResponse, error) {
	result := TransactionsClientListByCustomerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionListResult); err != nil {
		return TransactionsClientListByCustomerResponse{}, err
	}
	return result, nil
}

// NewListByInvoicePager - Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage
// charges.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - invoiceName - The ID that uniquely identifies an invoice.
//   - options - TransactionsClientListByInvoiceOptions contains the optional parameters for the TransactionsClient.NewListByInvoicePager
//     method.
func (client *TransactionsClient) NewListByInvoicePager(billingAccountName string, invoiceName string, options *TransactionsClientListByInvoiceOptions) *runtime.Pager[TransactionsClientListByInvoiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransactionsClientListByInvoiceResponse]{
		More: func(page TransactionsClientListByInvoiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransactionsClientListByInvoiceResponse) (TransactionsClientListByInvoiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TransactionsClient.NewListByInvoicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInvoiceCreateRequest(ctx, billingAccountName, invoiceName, options)
			}, nil)
			if err != nil {
				return TransactionsClientListByInvoiceResponse{}, err
			}
			return client.listByInvoiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInvoiceCreateRequest creates the ListByInvoice request.
func (client *TransactionsClient) listByInvoiceCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientListByInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/transactions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(*options.Top, 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInvoiceHandleResponse handles the ListByInvoice response.
func (client *TransactionsClient) listByInvoiceHandleResponse(resp *http.Response) (TransactionsClientListByInvoiceResponse, error) {
	result := TransactionsClientListByInvoiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionListResult); err != nil {
		return TransactionsClientListByInvoiceResponse{}, err
	}
	return result, nil
}

// NewListByInvoiceSectionPager - Lists the billed or unbilled transactions by invoice section name for given start date and
// end date. Transactions include purchases, refunds and Azure usage charges. Unbilled transactions are listed
// under pending invoice Id and do not include tax. Tax is added to the amount once an invoice is generated.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - periodStartDate - The start date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - periodEndDate - The end date to fetch the transactions. The date should be specified in MM-DD-YYYY format.
//   - typeParam - The type of transaction.
//   - options - TransactionsClientListByInvoiceSectionOptions contains the optional parameters for the TransactionsClient.NewListByInvoiceSectionPager
//     method.
func (client *TransactionsClient) NewListByInvoiceSectionPager(billingAccountName string, billingProfileName string, invoiceSectionName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByInvoiceSectionOptions) *runtime.Pager[TransactionsClientListByInvoiceSectionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransactionsClientListByInvoiceSectionResponse]{
		More: func(page TransactionsClientListByInvoiceSectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransactionsClientListByInvoiceSectionResponse) (TransactionsClientListByInvoiceSectionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TransactionsClient.NewListByInvoiceSectionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, periodStartDate, periodEndDate, typeParam, options)
			}, nil)
			if err != nil {
				return TransactionsClientListByInvoiceSectionResponse{}, err
			}
			return client.listByInvoiceSectionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInvoiceSectionCreateRequest creates the ListByInvoiceSection request.
func (client *TransactionsClient) listByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, periodStartDate time.Time, periodEndDate time.Time, typeParam TransactionType, options *TransactionsClientListByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transactions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	if options != nil && options.Count != nil {
		reqQP.Set("count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("orderBy", *options.OrderBy)
	}
	reqQP.Set("periodEndDate", periodEndDate.Format("2006-01-02"))
	reqQP.Set("periodStartDate", periodStartDate.Format("2006-01-02"))
	if options != nil && options.Search != nil {
		reqQP.Set("search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("skip", strconv.FormatInt(*options.Skip, 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(*options.Top, 10))
	}
	reqQP.Set("type", string(typeParam))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInvoiceSectionHandleResponse handles the ListByInvoiceSection response.
func (client *TransactionsClient) listByInvoiceSectionHandleResponse(resp *http.Response) (TransactionsClientListByInvoiceSectionResponse, error) {
	result := TransactionsClientListByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransactionListResult); err != nil {
		return TransactionsClientListByInvoiceSectionResponse{}, err
	}
	return result, nil
}

// BeginTransactionsDownloadByInvoice - Gets a URL to download the transactions document for an invoice. The operation is
// supported for billing accounts with agreement type Enterprise Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - invoiceName - The ID that uniquely identifies an invoice.
//   - options - TransactionsClientBeginTransactionsDownloadByInvoiceOptions contains the optional parameters for the TransactionsClient.BeginTransactionsDownloadByInvoice
//     method.
func (client *TransactionsClient) BeginTransactionsDownloadByInvoice(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientBeginTransactionsDownloadByInvoiceOptions) (*runtime.Poller[TransactionsClientTransactionsDownloadByInvoiceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.transactionsDownloadByInvoice(ctx, billingAccountName, invoiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TransactionsClientTransactionsDownloadByInvoiceResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TransactionsClientTransactionsDownloadByInvoiceResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TransactionsDownloadByInvoice - Gets a URL to download the transactions document for an invoice. The operation is supported
// for billing accounts with agreement type Enterprise Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *TransactionsClient) transactionsDownloadByInvoice(ctx context.Context, billingAccountName string, invoiceName string, options *TransactionsClientBeginTransactionsDownloadByInvoiceOptions) (*http.Response, error) {
	var err error
	const operationName = "TransactionsClient.BeginTransactionsDownloadByInvoice"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.transactionsDownloadByInvoiceCreateRequest(ctx, billingAccountName, invoiceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// transactionsDownloadByInvoiceCreateRequest creates the TransactionsDownloadByInvoice request.
func (client *TransactionsClient) transactionsDownloadByInvoiceCreateRequest(ctx context.Context, billingAccountName string, invoiceName string, _ *TransactionsClientBeginTransactionsDownloadByInvoiceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices/{invoiceName}/transactionsDownload"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if invoiceName == "" {
		return nil, errors.New("parameter invoiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceName}", url.PathEscape(invoiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
