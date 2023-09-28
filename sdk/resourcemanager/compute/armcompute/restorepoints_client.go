//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// RestorePointsClient contains the methods for the RestorePoints group.
// Don't use this type directly, use NewRestorePointsClient() instead.
type RestorePointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorePointsClient creates a new instance of RestorePointsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorePointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorePointsClient, error) {
	cl, err := arm.NewClient(moduleName+".RestorePointsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorePointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to create the restore point. Updating properties of an existing restore point is not allowed
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the restore point collection.
//   - restorePointName - The name of the restore point.
//   - parameters - Parameters supplied to the Create restore point operation.
//   - options - RestorePointsClientBeginCreateOptions contains the optional parameters for the RestorePointsClient.BeginCreate
//     method.
func (client *RestorePointsClient) BeginCreate(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters RestorePoint, options *RestorePointsClientBeginCreateOptions) (*runtime.Poller[RestorePointsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, restorePointCollectionName, restorePointName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RestorePointsClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RestorePointsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - The operation to create the restore point. Updating properties of an existing restore point is not allowed
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *RestorePointsClient) create(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters RestorePoint, options *RestorePointsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, restorePointCollectionName, restorePointName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *RestorePointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, parameters RestorePoint, options *RestorePointsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if restorePointName == "" {
		return nil, errors.New("parameter restorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointName}", url.PathEscape(restorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete the restore point.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the Restore Point Collection.
//   - restorePointName - The name of the restore point.
//   - options - RestorePointsClientBeginDeleteOptions contains the optional parameters for the RestorePointsClient.BeginDelete
//     method.
func (client *RestorePointsClient) BeginDelete(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *RestorePointsClientBeginDeleteOptions) (*runtime.Poller[RestorePointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, restorePointCollectionName, restorePointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[RestorePointsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RestorePointsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete the restore point.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *RestorePointsClient) deleteOperation(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *RestorePointsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, restorePointCollectionName, restorePointName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RestorePointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *RestorePointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if restorePointName == "" {
		return nil, errors.New("parameter restorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointName}", url.PathEscape(restorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get the restore point.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - restorePointCollectionName - The name of the restore point collection.
//   - restorePointName - The name of the restore point.
//   - options - RestorePointsClientGetOptions contains the optional parameters for the RestorePointsClient.Get method.
func (client *RestorePointsClient) Get(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *RestorePointsClientGetOptions) (RestorePointsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, restorePointCollectionName, restorePointName, options)
	if err != nil {
		return RestorePointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorePointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RestorePointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RestorePointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, restorePointCollectionName string, restorePointName string, options *RestorePointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/restorePointCollections/{restorePointCollectionName}/restorePoints/{restorePointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if restorePointCollectionName == "" {
		return nil, errors.New("parameter restorePointCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointCollectionName}", url.PathEscape(restorePointCollectionName))
	if restorePointName == "" {
		return nil, errors.New("parameter restorePointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorePointName}", url.PathEscape(restorePointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorePointsClient) getHandleResponse(resp *http.Response) (RestorePointsClientGetResponse, error) {
	result := RestorePointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorePoint); err != nil {
		return RestorePointsClientGetResponse{}, err
	}
	return result, nil
}
