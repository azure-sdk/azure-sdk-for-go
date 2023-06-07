//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtestbase

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

// BillingHubServiceClient contains the methods for the BillingHubService group.
// Don't use this type directly, use NewBillingHubServiceClient() instead.
type BillingHubServiceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBillingHubServiceClient creates a new instance of BillingHubServiceClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBillingHubServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BillingHubServiceClient, error) {
	cl, err := arm.NewClient(moduleName+".BillingHubServiceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BillingHubServiceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetFreeHourBalance -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - BillingHubServiceClientGetFreeHourBalanceOptions contains the optional parameters for the BillingHubServiceClient.GetFreeHourBalance
//     method.
func (client *BillingHubServiceClient) GetFreeHourBalance(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *BillingHubServiceClientGetFreeHourBalanceOptions) (BillingHubServiceClientGetFreeHourBalanceResponse, error) {
	req, err := client.getFreeHourBalanceCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return BillingHubServiceClientGetFreeHourBalanceResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BillingHubServiceClientGetFreeHourBalanceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingHubServiceClientGetFreeHourBalanceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getFreeHourBalanceHandleResponse(resp)
}

// getFreeHourBalanceCreateRequest creates the GetFreeHourBalance request.
func (client *BillingHubServiceClient) getFreeHourBalanceCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *BillingHubServiceClientGetFreeHourBalanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/getFreeHourBalance"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFreeHourBalanceHandleResponse handles the GetFreeHourBalance response.
func (client *BillingHubServiceClient) getFreeHourBalanceHandleResponse(resp *http.Response) (BillingHubServiceClientGetFreeHourBalanceResponse, error) {
	result := BillingHubServiceClientGetFreeHourBalanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingHubGetFreeHourBalanceResponse); err != nil {
		return BillingHubServiceClientGetFreeHourBalanceResponse{}, err
	}
	return result, nil
}

// GetUsage -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - BillingHubServiceClientGetUsageOptions contains the optional parameters for the BillingHubServiceClient.GetUsage
//     method.
func (client *BillingHubServiceClient) GetUsage(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *BillingHubServiceClientGetUsageOptions) (BillingHubServiceClientGetUsageResponse, error) {
	req, err := client.getUsageCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return BillingHubServiceClientGetUsageResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BillingHubServiceClientGetUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingHubServiceClientGetUsageResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUsageHandleResponse(resp)
}

// getUsageCreateRequest creates the GetUsage request.
func (client *BillingHubServiceClient) getUsageCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *BillingHubServiceClientGetUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/getUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.GetUsageRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.GetUsageRequest)
	}
	return req, nil
}

// getUsageHandleResponse handles the GetUsage response.
func (client *BillingHubServiceClient) getUsageHandleResponse(resp *http.Response) (BillingHubServiceClientGetUsageResponse, error) {
	result := BillingHubServiceClientGetUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingHubGetUsageResponse); err != nil {
		return BillingHubServiceClientGetUsageResponse{}, err
	}
	return result, nil
}
