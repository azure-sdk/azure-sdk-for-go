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

// TransactionsClient contains the methods for the Transactions group.
// Don't use this type directly, use NewTransactionsClient() instead.
type TransactionsClient struct {
	internal *arm.Client
}

// NewTransactionsClient creates a new instance of TransactionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTransactionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TransactionsClient, error) {
	cl, err := arm.NewClient(moduleName+".TransactionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TransactionsClient{
		internal: cl,
	}
	return client, nil
}

// NewListByInvoicePager - Lists the transactions for an invoice. Transactions include purchases, refunds and Azure usage
// charges.
//
// Generated from API version 2020-05-01
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
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInvoiceCreateRequest(ctx, billingAccountName, invoiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TransactionsClientListByInvoiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TransactionsClientListByInvoiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TransactionsClientListByInvoiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInvoiceHandleResponse(resp)
		},
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
	reqQP.Set("api-version", "2020-05-01")
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
