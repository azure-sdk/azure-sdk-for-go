//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Check whether a workspace name is available
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - request - The check request
//   - options - OperationsClientCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.CheckNameAvailability
//     method.
func (client *OperationsClient) CheckNameAvailability(ctx context.Context, request CheckNameAvailabilityRequest, options *OperationsClientCheckNameAvailabilityOptions) (OperationsClientCheckNameAvailabilityResponse, error) {
	var err error
	req, err := client.checkNameAvailabilityCreateRequest(ctx, request, options)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *OperationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, request CheckNameAvailabilityRequest, options *OperationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *OperationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (OperationsClientCheckNameAvailabilityResponse, error) {
	result := OperationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// GetAzureAsyncHeaderResult - Get the status of an operation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - operationID - Operation ID
//   - options - OperationsClientGetAzureAsyncHeaderResultOptions contains the optional parameters for the OperationsClient.GetAzureAsyncHeaderResult
//     method.
func (client *OperationsClient) GetAzureAsyncHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string, options *OperationsClientGetAzureAsyncHeaderResultOptions) (OperationsClientGetAzureAsyncHeaderResultResponse, error) {
	var err error
	req, err := client.getAzureAsyncHeaderResultCreateRequest(ctx, resourceGroupName, workspaceName, operationID, options)
	if err != nil {
		return OperationsClientGetAzureAsyncHeaderResultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientGetAzureAsyncHeaderResultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientGetAzureAsyncHeaderResultResponse{}, err
	}
	resp, err := client.getAzureAsyncHeaderResultHandleResponse(httpResp)
	return resp, err
}

// getAzureAsyncHeaderResultCreateRequest creates the GetAzureAsyncHeaderResult request.
func (client *OperationsClient) getAzureAsyncHeaderResultCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, operationID string, options *OperationsClientGetAzureAsyncHeaderResultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationStatuses/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAzureAsyncHeaderResultHandleResponse handles the GetAzureAsyncHeaderResult response.
func (client *OperationsClient) getAzureAsyncHeaderResultHandleResponse(resp *http.Response) (OperationsClientGetAzureAsyncHeaderResultResponse, error) {
	result := OperationsClientGetAzureAsyncHeaderResultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationResource); err != nil {
		return OperationsClientGetAzureAsyncHeaderResultResponse{}, err
	}
	return result, nil
}

// GetLocationHeaderResult - Get the result of an operation
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - operationID - Operation ID
//   - options - OperationsClientGetLocationHeaderResultOptions contains the optional parameters for the OperationsClient.GetLocationHeaderResult
//     method.
func (client *OperationsClient) GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string, options *OperationsClientGetLocationHeaderResultOptions) (OperationsClientGetLocationHeaderResultResponse, error) {
	var err error
	req, err := client.getLocationHeaderResultCreateRequest(ctx, resourceGroupName, workspaceName, operationID, options)
	if err != nil {
		return OperationsClientGetLocationHeaderResultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientGetLocationHeaderResultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientGetLocationHeaderResultResponse{}, err
	}
	return OperationsClientGetLocationHeaderResultResponse{}, nil
}

// getLocationHeaderResultCreateRequest creates the GetLocationHeaderResult request.
func (client *OperationsClient) getLocationHeaderResultCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, operationID string, options *OperationsClientGetLocationHeaderResultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// List - Get all available operations
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - options - OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
func (client *OperationsClient) List(ctx context.Context, options *OperationsClientListOptions) (OperationsClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Synapse/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableRpOperationArray); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}
