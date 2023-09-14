//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// RaiBlocklistItemClient contains the methods for the RaiBlocklistItem group.
// Don't use this type directly, use NewRaiBlocklistItemClient() instead.
type RaiBlocklistItemClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRaiBlocklistItemClient creates a new instance of RaiBlocklistItemClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRaiBlocklistItemClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RaiBlocklistItemClient, error) {
	cl, err := arm.NewClient(moduleName+".RaiBlocklistItemClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RaiBlocklistItemClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Update the state of specified blocklist item associated with the Azure OpenAI account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - raiBlocklistItemName - The name of the RaiBlocklist Item associated with the custom blocklist
//   - raiBlocklistItem - Properties describing the custom blocklist.
//   - options - RaiBlocklistItemClientCreateOrUpdateOptions contains the optional parameters for the RaiBlocklistItemClient.CreateOrUpdate
//     method.
func (client *RaiBlocklistItemClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, raiBlocklistItem RaiBlocklistItem, options *RaiBlocklistItemClientCreateOrUpdateOptions) (RaiBlocklistItemClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, raiBlocklistItemName, raiBlocklistItem, options)
	if err != nil {
		return RaiBlocklistItemClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiBlocklistItemClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RaiBlocklistItemClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RaiBlocklistItemClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, raiBlocklistItem RaiBlocklistItem, options *RaiBlocklistItemClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}/raiBlocklistItems/{raiBlocklistItemName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	if raiBlocklistItemName == "" {
		return nil, errors.New("parameter raiBlocklistItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistItemName}", url.PathEscape(raiBlocklistItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, raiBlocklistItem); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RaiBlocklistItemClient) createOrUpdateHandleResponse(resp *http.Response) (RaiBlocklistItemClientCreateOrUpdateResponse, error) {
	result := RaiBlocklistItemClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklistItem); err != nil {
		return RaiBlocklistItemClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified blocklist Item associated with the custom blocklist.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - raiBlocklistItemName - The name of the RaiBlocklist Item associated with the custom blocklist
//   - options - RaiBlocklistItemClientBeginDeleteOptions contains the optional parameters for the RaiBlocklistItemClient.BeginDelete
//     method.
func (client *RaiBlocklistItemClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, options *RaiBlocklistItemClientBeginDeleteOptions) (*runtime.Poller[RaiBlocklistItemClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, raiBlocklistName, raiBlocklistItemName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RaiBlocklistItemClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[RaiBlocklistItemClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified blocklist Item associated with the custom blocklist.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *RaiBlocklistItemClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, options *RaiBlocklistItemClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, raiBlocklistItemName, options)
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
func (client *RaiBlocklistItemClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, options *RaiBlocklistItemClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}/raiBlocklistItems/{raiBlocklistItemName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	if raiBlocklistItemName == "" {
		return nil, errors.New("parameter raiBlocklistItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistItemName}", url.PathEscape(raiBlocklistItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified custom blocklist Item associated with the custom blocklist.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - raiBlocklistName - The name of the RaiBlocklist associated with the Cognitive Services Account
//   - raiBlocklistItemName - The name of the RaiBlocklist Item associated with the custom blocklist
//   - options - RaiBlocklistItemClientGetOptions contains the optional parameters for the RaiBlocklistItemClient.Get method.
func (client *RaiBlocklistItemClient) Get(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, options *RaiBlocklistItemClientGetOptions) (RaiBlocklistItemClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, raiBlocklistName, raiBlocklistItemName, options)
	if err != nil {
		return RaiBlocklistItemClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RaiBlocklistItemClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RaiBlocklistItemClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RaiBlocklistItemClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklistItemName string, options *RaiBlocklistItemClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/raiBlocklists/{raiBlocklistName}/raiBlocklistItems/{raiBlocklistItemName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	if raiBlocklistItemName == "" {
		return nil, errors.New("parameter raiBlocklistItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistItemName}", url.PathEscape(raiBlocklistItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RaiBlocklistItemClient) getHandleResponse(resp *http.Response) (RaiBlocklistItemClientGetResponse, error) {
	result := RaiBlocklistItemClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklistItem); err != nil {
		return RaiBlocklistItemClientGetResponse{}, err
	}
	return result, nil
}
