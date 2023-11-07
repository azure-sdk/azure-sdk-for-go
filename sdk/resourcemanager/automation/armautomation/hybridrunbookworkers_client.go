//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// HybridRunbookWorkersClient contains the methods for the HybridRunbookWorkers group.
// Don't use this type directly, use NewHybridRunbookWorkersClient() instead.
type HybridRunbookWorkersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHybridRunbookWorkersClient creates a new instance of HybridRunbookWorkersClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHybridRunbookWorkersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridRunbookWorkersClient, error) {
	cl, err := arm.NewClient(moduleName+".HybridRunbookWorkersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HybridRunbookWorkersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - hybridRunbookWorkerGroupName - The hybrid runbook worker group name
//   - hybridRunbookWorkerID - The hybrid runbook worker id
//   - hybridRunbookWorkerCreationParameters - The create or update parameters for hybrid runbook worker.
//   - options - HybridRunbookWorkersClientCreateOptions contains the optional parameters for the HybridRunbookWorkersClient.Create
//     method.
func (client *HybridRunbookWorkersClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters, options *HybridRunbookWorkersClientCreateOptions) (HybridRunbookWorkersClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerCreationParameters, options)
	if err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *HybridRunbookWorkersClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters, options *HybridRunbookWorkersClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hybridRunbookWorkerCreationParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *HybridRunbookWorkersClient) createHandleResponse(resp *http.Response) (HybridRunbookWorkersClientCreateResponse, error) {
	result := HybridRunbookWorkersClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorker); err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - hybridRunbookWorkerGroupName - The hybrid runbook worker group name
//   - hybridRunbookWorkerID - The hybrid runbook worker id
//   - options - HybridRunbookWorkersClientDeleteOptions contains the optional parameters for the HybridRunbookWorkersClient.Delete
//     method.
func (client *HybridRunbookWorkersClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientDeleteOptions) (HybridRunbookWorkersClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, options)
	if err != nil {
		return HybridRunbookWorkersClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridRunbookWorkersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HybridRunbookWorkersClientDeleteResponse{}, err
	}
	return HybridRunbookWorkersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridRunbookWorkersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - hybridRunbookWorkerGroupName - The hybrid runbook worker group name
//   - hybridRunbookWorkerID - The hybrid runbook worker id
//   - options - HybridRunbookWorkersClientGetOptions contains the optional parameters for the HybridRunbookWorkersClient.Get
//     method.
func (client *HybridRunbookWorkersClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientGetOptions) (HybridRunbookWorkersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, options)
	if err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HybridRunbookWorkersClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridRunbookWorkersClient) getHandleResponse(resp *http.Response) (HybridRunbookWorkersClientGetResponse, error) {
	result := HybridRunbookWorkersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorker); err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHybridRunbookWorkerGroupPager - Retrieve a list of hybrid runbook workers.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - hybridRunbookWorkerGroupName - The hybrid runbook worker group name
//   - options - HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions contains the optional parameters for the HybridRunbookWorkersClient.NewListByHybridRunbookWorkerGroupPager
//     method.
func (client *HybridRunbookWorkersClient) NewListByHybridRunbookWorkerGroupPager(resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions) *runtime.Pager[HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse]{
		More: func(page HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse) (HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHybridRunbookWorkerGroupCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHybridRunbookWorkerGroupHandleResponse(resp)
		},
	})
}

// listByHybridRunbookWorkerGroupCreateRequest creates the ListByHybridRunbookWorkerGroup request.
func (client *HybridRunbookWorkersClient) listByHybridRunbookWorkerGroupCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHybridRunbookWorkerGroupHandleResponse handles the ListByHybridRunbookWorkerGroup response.
func (client *HybridRunbookWorkersClient) listByHybridRunbookWorkerGroupHandleResponse(resp *http.Response) (HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse, error) {
	result := HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkersListResult); err != nil {
		return HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}, err
	}
	return result, nil
}

// Move - Move a hybrid worker to a different group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - hybridRunbookWorkerGroupName - The hybrid runbook worker group name
//   - hybridRunbookWorkerID - The hybrid runbook worker id
//   - hybridRunbookWorkerMoveParameters - The hybrid runbook worker move parameters
//   - options - HybridRunbookWorkersClientMoveOptions contains the optional parameters for the HybridRunbookWorkersClient.Move
//     method.
func (client *HybridRunbookWorkersClient) Move(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters, options *HybridRunbookWorkersClientMoveOptions) (HybridRunbookWorkersClientMoveResponse, error) {
	var err error
	req, err := client.moveCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerMoveParameters, options)
	if err != nil {
		return HybridRunbookWorkersClientMoveResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridRunbookWorkersClientMoveResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HybridRunbookWorkersClientMoveResponse{}, err
	}
	return HybridRunbookWorkersClientMoveResponse{}, nil
}

// moveCreateRequest creates the Move request.
func (client *HybridRunbookWorkersClient) moveCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters, options *HybridRunbookWorkersClientMoveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}/move"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hybridRunbookWorkerMoveParameters); err != nil {
		return nil, err
	}
	return req, nil
}
