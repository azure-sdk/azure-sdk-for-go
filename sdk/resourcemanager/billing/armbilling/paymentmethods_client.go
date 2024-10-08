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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// PaymentMethodsClient contains the methods for the PaymentMethods group.
// Don't use this type directly, use NewPaymentMethodsClient() instead.
type PaymentMethodsClient struct {
	internal *arm.Client
}

// NewPaymentMethodsClient creates a new instance of PaymentMethodsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPaymentMethodsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PaymentMethodsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PaymentMethodsClient{
		internal: cl,
	}
	return client, nil
}

// DeleteByUser - Deletes a payment method owned by the caller.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - paymentMethodName - The ID that uniquely identifies a payment method.
//   - options - PaymentMethodsClientDeleteByUserOptions contains the optional parameters for the PaymentMethodsClient.DeleteByUser
//     method.
func (client *PaymentMethodsClient) DeleteByUser(ctx context.Context, paymentMethodName string, options *PaymentMethodsClientDeleteByUserOptions) (PaymentMethodsClientDeleteByUserResponse, error) {
	var err error
	const operationName = "PaymentMethodsClient.DeleteByUser"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByUserCreateRequest(ctx, paymentMethodName, options)
	if err != nil {
		return PaymentMethodsClientDeleteByUserResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PaymentMethodsClientDeleteByUserResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PaymentMethodsClientDeleteByUserResponse{}, err
	}
	return PaymentMethodsClientDeleteByUserResponse{}, nil
}

// deleteByUserCreateRequest creates the DeleteByUser request.
func (client *PaymentMethodsClient) deleteByUserCreateRequest(ctx context.Context, paymentMethodName string, options *PaymentMethodsClientDeleteByUserOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/paymentMethods/{paymentMethodName}"
	if paymentMethodName == "" {
		return nil, errors.New("parameter paymentMethodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{paymentMethodName}", url.PathEscape(paymentMethodName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetByBillingAccount - Gets a payment method available for a billing account. The operation is supported only for billing
// accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - paymentMethodName - The ID that uniquely identifies a payment method.
//   - options - PaymentMethodsClientGetByBillingAccountOptions contains the optional parameters for the PaymentMethodsClient.GetByBillingAccount
//     method.
func (client *PaymentMethodsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, paymentMethodName string, options *PaymentMethodsClientGetByBillingAccountOptions) (PaymentMethodsClientGetByBillingAccountResponse, error) {
	var err error
	const operationName = "PaymentMethodsClient.GetByBillingAccount"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountName, paymentMethodName, options)
	if err != nil {
		return PaymentMethodsClientGetByBillingAccountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PaymentMethodsClientGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PaymentMethodsClientGetByBillingAccountResponse{}, err
	}
	resp, err := client.getByBillingAccountHandleResponse(httpResp)
	return resp, err
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *PaymentMethodsClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, paymentMethodName string, options *PaymentMethodsClientGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/paymentMethods/{paymentMethodName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if paymentMethodName == "" {
		return nil, errors.New("parameter paymentMethodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{paymentMethodName}", url.PathEscape(paymentMethodName))
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

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *PaymentMethodsClient) getByBillingAccountHandleResponse(resp *http.Response) (PaymentMethodsClientGetByBillingAccountResponse, error) {
	result := PaymentMethodsClientGetByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethod); err != nil {
		return PaymentMethodsClientGetByBillingAccountResponse{}, err
	}
	return result, nil
}

// GetByBillingProfile - Gets a payment method linked with a billing profile. The operation is supported only for billing
// accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - paymentMethodName - The ID that uniquely identifies a payment method.
//   - options - PaymentMethodsClientGetByBillingProfileOptions contains the optional parameters for the PaymentMethodsClient.GetByBillingProfile
//     method.
func (client *PaymentMethodsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, paymentMethodName string, options *PaymentMethodsClientGetByBillingProfileOptions) (PaymentMethodsClientGetByBillingProfileResponse, error) {
	var err error
	const operationName = "PaymentMethodsClient.GetByBillingProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, paymentMethodName, options)
	if err != nil {
		return PaymentMethodsClientGetByBillingProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PaymentMethodsClientGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PaymentMethodsClientGetByBillingProfileResponse{}, err
	}
	resp, err := client.getByBillingProfileHandleResponse(httpResp)
	return resp, err
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *PaymentMethodsClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, paymentMethodName string, options *PaymentMethodsClientGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/paymentMethodLinks/{paymentMethodName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if paymentMethodName == "" {
		return nil, errors.New("parameter paymentMethodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{paymentMethodName}", url.PathEscape(paymentMethodName))
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

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *PaymentMethodsClient) getByBillingProfileHandleResponse(resp *http.Response) (PaymentMethodsClientGetByBillingProfileResponse, error) {
	result := PaymentMethodsClientGetByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethodLink); err != nil {
		return PaymentMethodsClientGetByBillingProfileResponse{}, err
	}
	return result, nil
}

// GetByUser - Gets a payment method owned by the caller.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - paymentMethodName - The ID that uniquely identifies a payment method.
//   - options - PaymentMethodsClientGetByUserOptions contains the optional parameters for the PaymentMethodsClient.GetByUser
//     method.
func (client *PaymentMethodsClient) GetByUser(ctx context.Context, paymentMethodName string, options *PaymentMethodsClientGetByUserOptions) (PaymentMethodsClientGetByUserResponse, error) {
	var err error
	const operationName = "PaymentMethodsClient.GetByUser"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByUserCreateRequest(ctx, paymentMethodName, options)
	if err != nil {
		return PaymentMethodsClientGetByUserResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PaymentMethodsClientGetByUserResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PaymentMethodsClientGetByUserResponse{}, err
	}
	resp, err := client.getByUserHandleResponse(httpResp)
	return resp, err
}

// getByUserCreateRequest creates the GetByUser request.
func (client *PaymentMethodsClient) getByUserCreateRequest(ctx context.Context, paymentMethodName string, options *PaymentMethodsClientGetByUserOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/paymentMethods/{paymentMethodName}"
	if paymentMethodName == "" {
		return nil, errors.New("parameter paymentMethodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{paymentMethodName}", url.PathEscape(paymentMethodName))
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

// getByUserHandleResponse handles the GetByUser response.
func (client *PaymentMethodsClient) getByUserHandleResponse(resp *http.Response) (PaymentMethodsClientGetByUserResponse, error) {
	result := PaymentMethodsClientGetByUserResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethod); err != nil {
		return PaymentMethodsClientGetByUserResponse{}, err
	}
	return result, nil
}

// NewListByBillingAccountPager - Lists the payment methods available for a billing account. Along with the payment methods
// owned by the caller, these payment methods can be attached to a billing profile to make payments. The
// operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - PaymentMethodsClientListByBillingAccountOptions contains the optional parameters for the PaymentMethodsClient.NewListByBillingAccountPager
//     method.
func (client *PaymentMethodsClient) NewListByBillingAccountPager(billingAccountName string, options *PaymentMethodsClientListByBillingAccountOptions) *runtime.Pager[PaymentMethodsClientListByBillingAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[PaymentMethodsClientListByBillingAccountResponse]{
		More: func(page PaymentMethodsClientListByBillingAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PaymentMethodsClientListByBillingAccountResponse) (PaymentMethodsClientListByBillingAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PaymentMethodsClient.NewListByBillingAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
			}, nil)
			if err != nil {
				return PaymentMethodsClientListByBillingAccountResponse{}, err
			}
			return client.listByBillingAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *PaymentMethodsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *PaymentMethodsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/paymentMethods"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *PaymentMethodsClient) listByBillingAccountHandleResponse(resp *http.Response) (PaymentMethodsClientListByBillingAccountResponse, error) {
	result := PaymentMethodsClientListByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethodsListResult); err != nil {
		return PaymentMethodsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// NewListByBillingProfilePager - Lists payment methods attached to a billing profile. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - PaymentMethodsClientListByBillingProfileOptions contains the optional parameters for the PaymentMethodsClient.NewListByBillingProfilePager
//     method.
func (client *PaymentMethodsClient) NewListByBillingProfilePager(billingAccountName string, billingProfileName string, options *PaymentMethodsClientListByBillingProfileOptions) *runtime.Pager[PaymentMethodsClientListByBillingProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[PaymentMethodsClientListByBillingProfileResponse]{
		More: func(page PaymentMethodsClientListByBillingProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PaymentMethodsClientListByBillingProfileResponse) (PaymentMethodsClientListByBillingProfileResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PaymentMethodsClient.NewListByBillingProfilePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
			}, nil)
			if err != nil {
				return PaymentMethodsClientListByBillingProfileResponse{}, err
			}
			return client.listByBillingProfileHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *PaymentMethodsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *PaymentMethodsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/paymentMethodLinks"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *PaymentMethodsClient) listByBillingProfileHandleResponse(resp *http.Response) (PaymentMethodsClientListByBillingProfileResponse, error) {
	result := PaymentMethodsClientListByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethodLinksListResult); err != nil {
		return PaymentMethodsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}

// NewListByUserPager - Lists the payment methods owned by the caller.
//
// Generated from API version 2024-04-01
//   - options - PaymentMethodsClientListByUserOptions contains the optional parameters for the PaymentMethodsClient.NewListByUserPager
//     method.
func (client *PaymentMethodsClient) NewListByUserPager(options *PaymentMethodsClientListByUserOptions) *runtime.Pager[PaymentMethodsClientListByUserResponse] {
	return runtime.NewPager(runtime.PagingHandler[PaymentMethodsClientListByUserResponse]{
		More: func(page PaymentMethodsClientListByUserResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PaymentMethodsClientListByUserResponse) (PaymentMethodsClientListByUserResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PaymentMethodsClient.NewListByUserPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByUserCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PaymentMethodsClientListByUserResponse{}, err
			}
			return client.listByUserHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByUserCreateRequest creates the ListByUser request.
func (client *PaymentMethodsClient) listByUserCreateRequest(ctx context.Context, options *PaymentMethodsClientListByUserOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/paymentMethods"
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

// listByUserHandleResponse handles the ListByUser response.
func (client *PaymentMethodsClient) listByUserHandleResponse(resp *http.Response) (PaymentMethodsClientListByUserResponse, error) {
	result := PaymentMethodsClientListByUserResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PaymentMethodsListResult); err != nil {
		return PaymentMethodsClientListByUserResponse{}, err
	}
	return result, nil
}
