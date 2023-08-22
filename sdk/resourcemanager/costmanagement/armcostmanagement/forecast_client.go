//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ForecastClient contains the methods for the Forecast group.
// Don't use this type directly, use NewForecastClient() instead.
type ForecastClient struct {
	internal *arm.Client
}

// NewForecastClient creates a new instance of ForecastClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewForecastClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ForecastClient, error) {
	cl, err := arm.NewClient(moduleName+".ForecastClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ForecastClient{
		internal: cl,
	}
	return client, nil
}

// ExternalCloudProviderUsage - Lists the forecast charges for external cloud provider type defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - externalCloudProviderType - The external cloud provider type associated with dimension/query operations. This includes
//     'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated account.
//   - externalCloudProviderID - This can be '{externalSubscriptionId}' for linked account or '{externalBillingAccountId}' for
//     consolidated account used with dimension/query operations.
//   - parameters - Parameters supplied to the CreateOrUpdate Forecast Config operation.
//   - options - ForecastClientExternalCloudProviderUsageOptions contains the optional parameters for the ForecastClient.ExternalCloudProviderUsage
//     method.
func (client *ForecastClient) ExternalCloudProviderUsage(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters ForecastDefinition, options *ForecastClientExternalCloudProviderUsageOptions) (ForecastClientExternalCloudProviderUsageResponse, error) {
	var err error
	req, err := client.externalCloudProviderUsageCreateRequest(ctx, externalCloudProviderType, externalCloudProviderID, parameters, options)
	if err != nil {
		return ForecastClientExternalCloudProviderUsageResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ForecastClientExternalCloudProviderUsageResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ForecastClientExternalCloudProviderUsageResponse{}, err
	}
	resp, err := client.externalCloudProviderUsageHandleResponse(httpResp)
	return resp, err
}

// externalCloudProviderUsageCreateRequest creates the ExternalCloudProviderUsage request.
func (client *ForecastClient) externalCloudProviderUsageCreateRequest(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters ForecastDefinition, options *ForecastClientExternalCloudProviderUsageOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/forecast"
	if externalCloudProviderType == "" {
		return nil, errors.New("parameter externalCloudProviderType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderType}", url.PathEscape(string(externalCloudProviderType)))
	if externalCloudProviderID == "" {
		return nil, errors.New("parameter externalCloudProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderId}", url.PathEscape(externalCloudProviderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// externalCloudProviderUsageHandleResponse handles the ExternalCloudProviderUsage response.
func (client *ForecastClient) externalCloudProviderUsageHandleResponse(resp *http.Response) (ForecastClientExternalCloudProviderUsageResponse, error) {
	result := ForecastClientExternalCloudProviderUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForecastResult); err != nil {
		return ForecastClientExternalCloudProviderUsageResponse{}, err
	}
	return result, nil
}

// Usage - Lists the forecast charges for scope defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - scope - The scope associated with forecast operations. This includes '/subscriptions/{subscriptionId}/' for subscription
//     scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
//     resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
//     scope,
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
//     for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
//     for invoiceSection scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
//   - parameters - Parameters supplied to the CreateOrUpdate Forecast Config operation.
//   - options - ForecastClientUsageOptions contains the optional parameters for the ForecastClient.Usage method.
func (client *ForecastClient) Usage(ctx context.Context, scope string, parameters ForecastDefinition, options *ForecastClientUsageOptions) (ForecastClientUsageResponse, error) {
	var err error
	req, err := client.usageCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return ForecastClientUsageResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ForecastClientUsageResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ForecastClientUsageResponse{}, err
	}
	resp, err := client.usageHandleResponse(httpResp)
	return resp, err
}

// usageCreateRequest creates the Usage request.
func (client *ForecastClient) usageCreateRequest(ctx context.Context, scope string, parameters ForecastDefinition, options *ForecastClientUsageOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/forecast"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// usageHandleResponse handles the Usage response.
func (client *ForecastClient) usageHandleResponse(resp *http.Response) (ForecastClientUsageResponse, error) {
	result := ForecastClientUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ForecastResult); err != nil {
		return ForecastClientUsageResponse{}, err
	}
	return result, nil
}
