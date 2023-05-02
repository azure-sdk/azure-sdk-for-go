//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

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

// QuotaByPeriodKeysClient contains the methods for the QuotaByPeriodKeys group.
// Don't use this type directly, use NewQuotaByPeriodKeysClient() instead.
type QuotaByPeriodKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQuotaByPeriodKeysClient creates a new instance of QuotaByPeriodKeysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotaByPeriodKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotaByPeriodKeysClient, error) {
	cl, err := arm.NewClient(moduleName+".QuotaByPeriodKeysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QuotaByPeriodKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the value of the quota counter associated with the counter-key in the policy for the specific period in service
// instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - quotaCounterKey - Quota counter key identifier.This is the result of expression defined in counter-key attribute of the
//     quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
//     accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible by "ba" key
//   - quotaPeriodKey - Quota period key identifier.
//   - options - QuotaByPeriodKeysClientGetOptions contains the optional parameters for the QuotaByPeriodKeysClient.Get method.
func (client *QuotaByPeriodKeysClient) Get(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, options *QuotaByPeriodKeysClientGetOptions) (QuotaByPeriodKeysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, quotaCounterKey, quotaPeriodKey, options)
	if err != nil {
		return QuotaByPeriodKeysClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QuotaByPeriodKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotaByPeriodKeysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QuotaByPeriodKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, options *QuotaByPeriodKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if quotaCounterKey == "" {
		return nil, errors.New("parameter quotaCounterKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaCounterKey}", url.PathEscape(quotaCounterKey))
	if quotaPeriodKey == "" {
		return nil, errors.New("parameter quotaPeriodKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaPeriodKey}", url.PathEscape(quotaPeriodKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QuotaByPeriodKeysClient) getHandleResponse(resp *http.Response) (QuotaByPeriodKeysClientGetResponse, error) {
	result := QuotaByPeriodKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaCounterContract); err != nil {
		return QuotaByPeriodKeysClientGetResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing quota counter value in the specified service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - quotaCounterKey - Quota counter key identifier.This is the result of expression defined in counter-key attribute of the
//     quota-by-key policy.For Example, if you specify counter-key="boo" in the policy, then it’s
//     accessible by "boo" counter key. But if it’s defined as counter-key="@("b"+"a")" then it will be accessible by "ba" key
//   - quotaPeriodKey - Quota period key identifier.
//   - parameters - The value of the Quota counter to be applied on the specified period.
//   - options - QuotaByPeriodKeysClientUpdateOptions contains the optional parameters for the QuotaByPeriodKeysClient.Update
//     method.
func (client *QuotaByPeriodKeysClient) Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, parameters QuotaCounterValueUpdateContract, options *QuotaByPeriodKeysClientUpdateOptions) (QuotaByPeriodKeysClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, quotaCounterKey, quotaPeriodKey, parameters, options)
	if err != nil {
		return QuotaByPeriodKeysClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QuotaByPeriodKeysClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotaByPeriodKeysClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *QuotaByPeriodKeysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, parameters QuotaCounterValueUpdateContract, options *QuotaByPeriodKeysClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if quotaCounterKey == "" {
		return nil, errors.New("parameter quotaCounterKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaCounterKey}", url.PathEscape(quotaCounterKey))
	if quotaPeriodKey == "" {
		return nil, errors.New("parameter quotaPeriodKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaPeriodKey}", url.PathEscape(quotaPeriodKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *QuotaByPeriodKeysClient) updateHandleResponse(resp *http.Response) (QuotaByPeriodKeysClientUpdateResponse, error) {
	result := QuotaByPeriodKeysClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaCounterContract); err != nil {
		return QuotaByPeriodKeysClientUpdateResponse{}, err
	}
	return result, nil
}
