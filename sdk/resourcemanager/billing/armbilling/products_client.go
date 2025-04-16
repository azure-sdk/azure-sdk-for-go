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
)

// ProductsClient contains the methods for the Products group.
// Don't use this type directly, use NewProductsClient() instead.
type ProductsClient struct {
	internal *arm.Client
}

// NewProductsClient creates a new instance of ProductsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProductsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProductsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets a product by ID. The operation is supported only for billing accounts with agreement type Microsoft Customer
// Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - productName - The ID that uniquely identifies a product.
//   - options - ProductsClientGetOptions contains the optional parameters for the ProductsClient.Get method.
func (client *ProductsClient) Get(ctx context.Context, billingAccountName string, productName string, options *ProductsClientGetOptions) (ProductsClientGetResponse, error) {
	var err error
	const operationName = "ProductsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, billingAccountName, productName, options)
	if err != nil {
		return ProductsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProductsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProductsClient) getCreateRequest(ctx context.Context, billingAccountName string, productName string, _ *ProductsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProductsClient) getHandleResponse(resp *http.Response) (ProductsClientGetResponse, error) {
	result := ProductsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return ProductsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the products for a billing account. These don't include products billed based on usage.
// The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or
// Microsoft Partner Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - ProductsClientListByBillingAccountOptions contains the optional parameters for the ProductsClient.NewListByBillingAccountPager
//     method.
func (client *ProductsClient) NewListByBillingAccountPager(billingAccountName string, options *ProductsClientListByBillingAccountOptions) *runtime.Pager[ProductsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientListByBillingAccountResponse]{
		More: func(page ProductsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientListByBillingAccountResponse) (ProductsClientListByBillingAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsClient.NewListByBillingAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			}, nil)
			if err != nil {
				return ProductsClientListByBillingAccountResponse{}, err
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *ProductsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *ProductsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
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

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *ProductsClient) listByBillingAccountHandleResponse(resp *http.Response) (ProductsClientListByBillingAccountResponse, error) {
	result := ProductsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductListResult); err != nil {
		return ProductsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists the products for a billing profile. These don't include products billed based on usage.
// The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or
// Microsoft Partner Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - ProductsClientListByBillingProfileOptions contains the optional parameters for the ProductsClient.NewListByBillingProfilePager
//     method.
func (client *ProductsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *ProductsClientListByBillingProfileOptions) *runtime.Pager[ProductsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientListByBillingProfileResponse]{
		More: func(page ProductsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientListByBillingProfileResponse) (ProductsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			}, nil)
			if err != nil {
				return ProductsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *ProductsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *ProductsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/products"
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

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *ProductsClient) listByBillingProfileHandleResponse(resp *http.Response) (ProductsClientListByBillingProfileResponse, error) {
	result := ProductsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductListResult); err != nil {
		return ProductsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByCustomerPager - Lists the products for a customer. These don't include products billed based on usage.The operation
// is supported only for billing accounts with agreement type Microsoft Partner Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - customerName - The ID that uniquely identifies a customer.
//   - options - ProductsClientListByCustomerOptions contains the optional parameters for the ProductsClient.NewListByCustomerPager
//     method.
func (client *ProductsClient) NewListByCustomerPager(billingAccountName string, customerName string, options *ProductsClientListByCustomerOptions) *runtime.Pager[ProductsClientListByCustomerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientListByCustomerResponse]{
		More: func(page ProductsClientListByCustomerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientListByCustomerResponse) (ProductsClientListByCustomerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsClient.NewListByCustomerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCustomerCreateRequest(ctx, billingAccountName, customerName, options)
			}, nil)
			if err != nil {
				return ProductsClientListByCustomerResponse{}, err
			}
			return client.listByCustomerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCustomerCreateRequest creates the ListByCustomer request.
func (client *ProductsClient) listByCustomerCreateRequest(ctx context.Context, billingAccountName string, customerName string, options *ProductsClientListByCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/products"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
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

// listByCustomerHandleResponse handles the ListByCustomer response.
func (client *ProductsClient) listByCustomerHandleResponse(resp *http.Response) (ProductsClientListByCustomerResponse, error) {
	result := ProductsClientListByCustomerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductListResult); err != nil {
		return ProductsClientListByCustomerResponse{}, err
	}
	return result, nil
}

// NewListByInvoiceSectionPager - Lists the products for an invoice section. These don't include products billed based on
// usage. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - invoiceSectionName - The ID that uniquely identifies an invoice section.
//   - options - ProductsClientListByInvoiceSectionOptions contains the optional parameters for the ProductsClient.NewListByInvoiceSectionPager
//     method.
func (client *ProductsClient) NewListByInvoiceSectionPager(billingAccountName string, billingProfileName string, invoiceSectionName string, options *ProductsClientListByInvoiceSectionOptions) *runtime.Pager[ProductsClientListByInvoiceSectionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsClientListByInvoiceSectionResponse]{
		More: func(page ProductsClientListByInvoiceSectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsClientListByInvoiceSectionResponse) (ProductsClientListByInvoiceSectionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsClient.NewListByInvoiceSectionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
			}, nil)
			if err != nil {
				return ProductsClientListByInvoiceSectionResponse{}, err
			}
			return client.listByInvoiceSectionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInvoiceSectionCreateRequest creates the ListByInvoiceSection request.
func (client *ProductsClient) listByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *ProductsClientListByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/products"
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

// listByInvoiceSectionHandleResponse handles the ListByInvoiceSection response.
func (client *ProductsClient) listByInvoiceSectionHandleResponse(resp *http.Response) (ProductsClientListByInvoiceSectionResponse, error) {
	result := ProductsClientListByInvoiceSectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductListResult); err != nil {
		return ProductsClientListByInvoiceSectionResponse{}, err
	}
	return result, nil
}

// BeginMove - Moves a product's charges to a new invoice section. The new invoice section must belong to the same billing
// profile as the existing invoice section. This operation is supported only for products that
// are purchased with a recurring charge and for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - productName - The ID that uniquely identifies a product.
//   - parameters - The properties of the product to initiate a transfer.
//   - options - ProductsClientBeginMoveOptions contains the optional parameters for the ProductsClient.BeginMove method.
func (client *ProductsClient) BeginMove(ctx context.Context, billingAccountName string, productName string, parameters MoveProductRequest, options *ProductsClientBeginMoveOptions) (*runtime.Poller[ProductsClientMoveResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.move(ctx, billingAccountName, productName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProductsClientMoveResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProductsClientMoveResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Move - Moves a product's charges to a new invoice section. The new invoice section must belong to the same billing profile
// as the existing invoice section. This operation is supported only for products that
// are purchased with a recurring charge and for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ProductsClient) move(ctx context.Context, billingAccountName string, productName string, parameters MoveProductRequest, options *ProductsClientBeginMoveOptions) (*http.Response, error) {
	var err error
	const operationName = "ProductsClient.BeginMove"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.moveCreateRequest(ctx, billingAccountName, productName, parameters, options)
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

// moveCreateRequest creates the Move request.
func (client *ProductsClient) moveCreateRequest(ctx context.Context, billingAccountName string, productName string, parameters MoveProductRequest, _ *ProductsClientBeginMoveOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}/move"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Update - Updates the properties of a Product. Currently, auto renew can be updated. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - productName - The ID that uniquely identifies a product.
//   - parameters - A product.
//   - options - ProductsClientUpdateOptions contains the optional parameters for the ProductsClient.Update method.
func (client *ProductsClient) Update(ctx context.Context, billingAccountName string, productName string, parameters ProductPatch, options *ProductsClientUpdateOptions) (ProductsClientUpdateResponse, error) {
	var err error
	const operationName = "ProductsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, billingAccountName, productName, parameters, options)
	if err != nil {
		return ProductsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProductsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ProductsClient) updateCreateRequest(ctx context.Context, billingAccountName string, productName string, parameters ProductPatch, _ *ProductsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ProductsClient) updateHandleResponse(resp *http.Response) (ProductsClientUpdateResponse, error) {
	result := ProductsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Product); err != nil {
		return ProductsClientUpdateResponse{}, err
	}
	return result, nil
}

// ValidateMoveEligibility - Validates if a product's charges can be moved to a new invoice section. This operation is supported
// only for products that are purchased with a recurring charge and for billing accounts with agreement
// type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - productName - The ID that uniquely identifies a product.
//   - parameters - The properties of the product to initiate a transfer.
//   - options - ProductsClientValidateMoveEligibilityOptions contains the optional parameters for the ProductsClient.ValidateMoveEligibility
//     method.
func (client *ProductsClient) ValidateMoveEligibility(ctx context.Context, billingAccountName string, productName string, parameters MoveProductRequest, options *ProductsClientValidateMoveEligibilityOptions) (ProductsClientValidateMoveEligibilityResponse, error) {
	var err error
	const operationName = "ProductsClient.ValidateMoveEligibility"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.validateMoveEligibilityCreateRequest(ctx, billingAccountName, productName, parameters, options)
	if err != nil {
		return ProductsClientValidateMoveEligibilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductsClientValidateMoveEligibilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProductsClientValidateMoveEligibilityResponse{}, err
	}
	resp, err := client.validateMoveEligibilityHandleResponse(httpResp)
	return resp, err
}

// validateMoveEligibilityCreateRequest creates the ValidateMoveEligibility request.
func (client *ProductsClient) validateMoveEligibilityCreateRequest(ctx context.Context, billingAccountName string, productName string, parameters MoveProductRequest, _ *ProductsClientValidateMoveEligibilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}/validateMoveEligibility"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// validateMoveEligibilityHandleResponse handles the ValidateMoveEligibility response.
func (client *ProductsClient) validateMoveEligibilityHandleResponse(resp *http.Response) (ProductsClientValidateMoveEligibilityResponse, error) {
	result := ProductsClientValidateMoveEligibilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MoveProductEligibilityResult); err != nil {
		return ProductsClientValidateMoveEligibilityResponse{}, err
	}
	return result, nil
}
