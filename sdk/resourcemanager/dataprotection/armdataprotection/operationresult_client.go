//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

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

// OperationResultClient contains the methods for the OperationResult group.
// Don't use this type directly, use NewOperationResultClient() instead.
type OperationResultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationResultClient creates a new instance of OperationResultClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationResultClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationResultClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationResultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the operation result for a resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - options - OperationResultClientGetOptions contains the optional parameters for the OperationResultClient.Get method.
func (client *OperationResultClient) Get(ctx context.Context, operationID string, location string, options *OperationResultClientGetOptions) (OperationResultClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, operationID, location, options)
	if err != nil {
		return OperationResultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return OperationResultClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OperationResultClient) getCreateRequest(ctx context.Context, operationID string, location string, options *OperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataProtection/locations/{location}/operationResults/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationResultClient) getHandleResponse(resp *http.Response) (OperationResultClientGetResponse, error) {
	result := OperationResultClientGetResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return OperationResultClientGetResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationJobExtendedInfo); err != nil {
		return OperationResultClientGetResponse{}, err
	}
	return result, nil
}
