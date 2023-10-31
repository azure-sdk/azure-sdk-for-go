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

// PoliciesClient contains the methods for the Policies group.
// Don't use this type directly, use NewPoliciesClient() instead.
type PoliciesClient struct {
	internal *arm.Client
}

// NewPoliciesClient creates a new instance of PoliciesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPoliciesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".PoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PoliciesClient{
		internal: cl,
	}
	return client, nil
}

// GetByBillingProfile - Lists the policies for a billing profile. This operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - PoliciesClientGetByBillingProfileOptions contains the optional parameters for the PoliciesClient.GetByBillingProfile
//     method.
func (client *PoliciesClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, options *PoliciesClientGetByBillingProfileOptions) (PoliciesClientGetByBillingProfileResponse, error) {
	var err error
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return PoliciesClientGetByBillingProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PoliciesClientGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PoliciesClientGetByBillingProfileResponse{}, err
	}
	resp, err := client.getByBillingProfileHandleResponse(httpResp)
	return resp, err
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *PoliciesClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *PoliciesClientGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *PoliciesClient) getByBillingProfileHandleResponse(resp *http.Response) (PoliciesClientGetByBillingProfileResponse, error) {
	result := PoliciesClientGetByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Policy); err != nil {
		return PoliciesClientGetByBillingProfileResponse{}, err
	}
	return result, nil
}

// GetByCustomer - Lists the policies for a customer. This operation is supported only for billing accounts with agreement
// type Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - customerName - The ID that uniquely identifies a customer.
//   - options - PoliciesClientGetByCustomerOptions contains the optional parameters for the PoliciesClient.GetByCustomer method.
func (client *PoliciesClient) GetByCustomer(ctx context.Context, billingAccountName string, customerName string, options *PoliciesClientGetByCustomerOptions) (PoliciesClientGetByCustomerResponse, error) {
	var err error
	req, err := client.getByCustomerCreateRequest(ctx, billingAccountName, customerName, options)
	if err != nil {
		return PoliciesClientGetByCustomerResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PoliciesClientGetByCustomerResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PoliciesClientGetByCustomerResponse{}, err
	}
	resp, err := client.getByCustomerHandleResponse(httpResp)
	return resp, err
}

// getByCustomerCreateRequest creates the GetByCustomer request.
func (client *PoliciesClient) getByCustomerCreateRequest(ctx context.Context, billingAccountName string, customerName string, options *PoliciesClientGetByCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default"
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
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByCustomerHandleResponse handles the GetByCustomer response.
func (client *PoliciesClient) getByCustomerHandleResponse(resp *http.Response) (PoliciesClientGetByCustomerResponse, error) {
	result := PoliciesClientGetByCustomerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerPolicy); err != nil {
		return PoliciesClientGetByCustomerResponse{}, err
	}
	return result, nil
}

// Update - Updates the policies for a billing profile. This operation is supported only for billing accounts with agreement
// type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - parameters - Request parameters that are provided to the update policies operation.
//   - options - PoliciesClientUpdateOptions contains the optional parameters for the PoliciesClient.Update method.
func (client *PoliciesClient) Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy, options *PoliciesClientUpdateOptions) (PoliciesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return PoliciesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PoliciesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *PoliciesClient) updateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy, options *PoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PoliciesClient) updateHandleResponse(resp *http.Response) (PoliciesClientUpdateResponse, error) {
	result := PoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Policy); err != nil {
		return PoliciesClientUpdateResponse{}, err
	}
	return result, nil
}

// UpdateCustomer - Updates the policies for a customer. This operation is supported only for billing accounts with agreement
// type Microsoft Partner Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - customerName - The ID that uniquely identifies a customer.
//   - parameters - Request parameters that are provided to the update policies operation.
//   - options - PoliciesClientUpdateCustomerOptions contains the optional parameters for the PoliciesClient.UpdateCustomer method.
func (client *PoliciesClient) UpdateCustomer(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy, options *PoliciesClientUpdateCustomerOptions) (PoliciesClientUpdateCustomerResponse, error) {
	var err error
	req, err := client.updateCustomerCreateRequest(ctx, billingAccountName, customerName, parameters, options)
	if err != nil {
		return PoliciesClientUpdateCustomerResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PoliciesClientUpdateCustomerResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PoliciesClientUpdateCustomerResponse{}, err
	}
	resp, err := client.updateCustomerHandleResponse(httpResp)
	return resp, err
}

// updateCustomerCreateRequest creates the UpdateCustomer request.
func (client *PoliciesClient) updateCustomerCreateRequest(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy, options *PoliciesClientUpdateCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateCustomerHandleResponse handles the UpdateCustomer response.
func (client *PoliciesClient) updateCustomerHandleResponse(resp *http.Response) (PoliciesClientUpdateCustomerResponse, error) {
	result := PoliciesClientUpdateCustomerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerPolicy); err != nil {
		return PoliciesClientUpdateCustomerResponse{}, err
	}
	return result, nil
}
