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
	"strings"
)

// FetchCrossRegionRestoreJobClient contains the methods for the FetchCrossRegionRestoreJob group.
// Don't use this type directly, use NewFetchCrossRegionRestoreJobClient() instead.
type FetchCrossRegionRestoreJobClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFetchCrossRegionRestoreJobClient creates a new instance of FetchCrossRegionRestoreJobClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFetchCrossRegionRestoreJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FetchCrossRegionRestoreJobClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FetchCrossRegionRestoreJobClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Fetches the Cross Region Restore Job
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - location - The name of the Azure region.
//   - parameters - Request body for operation
//   - options - FetchCrossRegionRestoreJobClientGetOptions contains the optional parameters for the FetchCrossRegionRestoreJobClient.Get
//     method.
func (client *FetchCrossRegionRestoreJobClient) Get(ctx context.Context, resourceGroupName string, location string, parameters CrossRegionRestoreJobRequest, options *FetchCrossRegionRestoreJobClientGetOptions) (FetchCrossRegionRestoreJobClientGetResponse, error) {
	var err error
	const operationName = "FetchCrossRegionRestoreJobClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, location, parameters, options)
	if err != nil {
		return FetchCrossRegionRestoreJobClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FetchCrossRegionRestoreJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FetchCrossRegionRestoreJobClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FetchCrossRegionRestoreJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, location string, parameters CrossRegionRestoreJobRequest, _ *FetchCrossRegionRestoreJobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/locations/{location}/fetchCrossRegionRestoreJob"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FetchCrossRegionRestoreJobClient) getHandleResponse(resp *http.Response) (FetchCrossRegionRestoreJobClientGetResponse, error) {
	result := FetchCrossRegionRestoreJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupJobResource); err != nil {
		return FetchCrossRegionRestoreJobClientGetResponse{}, err
	}
	return result, nil
}
