// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// HypervOperationsStatusControllerClient contains the methods for the HypervOperationsStatusController group.
// Don't use this type directly, use NewHypervOperationsStatusControllerClient() instead.
type HypervOperationsStatusControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHypervOperationsStatusControllerClient creates a new instance of HypervOperationsStatusControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHypervOperationsStatusControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HypervOperationsStatusControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HypervOperationsStatusControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetHypervOperationsStatus - Method to get operation status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - operationStatusName - Operation status Arm Name.
//   - options - HypervOperationsStatusControllerClientGetHypervOperationsStatusOptions contains the optional parameters for the
//     HypervOperationsStatusControllerClient.GetHypervOperationsStatus method.
func (client *HypervOperationsStatusControllerClient) GetHypervOperationsStatus(ctx context.Context, resourceGroupName string, siteName string, operationStatusName string, options *HypervOperationsStatusControllerClientGetHypervOperationsStatusOptions) (HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse, error) {
	var err error
	const operationName = "HypervOperationsStatusControllerClient.GetHypervOperationsStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getHypervOperationsStatusCreateRequest(ctx, resourceGroupName, siteName, operationStatusName, options)
	if err != nil {
		return HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse{}, err
	}
	resp, err := client.getHypervOperationsStatusHandleResponse(httpResp)
	return resp, err
}

// getHypervOperationsStatusCreateRequest creates the GetHypervOperationsStatus request.
func (client *HypervOperationsStatusControllerClient) getHypervOperationsStatusCreateRequest(ctx context.Context, resourceGroupName string, siteName string, operationStatusName string, _ *HypervOperationsStatusControllerClientGetHypervOperationsStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/operationsStatus/{operationStatusName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if operationStatusName == "" {
		return nil, errors.New("parameter operationStatusName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationStatusName}", url.PathEscape(operationStatusName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHypervOperationsStatusHandleResponse handles the GetHypervOperationsStatus response.
func (client *HypervOperationsStatusControllerClient) getHypervOperationsStatusHandleResponse(resp *http.Response) (HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse, error) {
	result := HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return HypervOperationsStatusControllerClientGetHypervOperationsStatusResponse{}, err
	}
	return result, nil
}
