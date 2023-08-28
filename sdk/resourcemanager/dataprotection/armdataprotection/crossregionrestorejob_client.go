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
	"strings"
)

// CrossRegionRestoreJobClient contains the methods for the CrossRegionRestoreJob group.
// Don't use this type directly, use NewCrossRegionRestoreJobClient() instead.
type CrossRegionRestoreJobClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCrossRegionRestoreJobClient creates a new instance of CrossRegionRestoreJobClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCrossRegionRestoreJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CrossRegionRestoreJobClient, error) {
	cl, err := arm.NewClient(moduleName+".CrossRegionRestoreJobClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CrossRegionRestoreJobClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - Request body for operation
//   - options - CrossRegionRestoreJobClientGetOptions contains the optional parameters for the CrossRegionRestoreJobClient.Get
//     method.
func (client *CrossRegionRestoreJobClient) Get(ctx context.Context, resourceGroupName string, location string, parameters CrossRegionRestoreJobRequest, options *CrossRegionRestoreJobClientGetOptions) (CrossRegionRestoreJobClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, location, parameters, options)
	if err != nil {
		return CrossRegionRestoreJobClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CrossRegionRestoreJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CrossRegionRestoreJobClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CrossRegionRestoreJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, location string, parameters CrossRegionRestoreJobRequest, options *CrossRegionRestoreJobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/locations/{location}/fetchCrossRegionRestoreJob"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CrossRegionRestoreJobClient) getHandleResponse(resp *http.Response) (CrossRegionRestoreJobClientGetResponse, error) {
	result := CrossRegionRestoreJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupJobResource); err != nil {
		return CrossRegionRestoreJobClientGetResponse{}, err
	}
	return result, nil
}
