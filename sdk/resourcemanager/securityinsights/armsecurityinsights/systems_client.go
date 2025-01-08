//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// SystemsClient contains the methods for the Systems group.
// Don't use this type directly, use NewSystemsClient() instead.
type SystemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSystemsClient creates a new instance of SystemsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSystemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SystemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientCreateOrUpdateOptions contains the optional parameters for the SystemsClient.CreateOrUpdate method.
func (client *SystemsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientCreateOrUpdateOptions) (SystemsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SystemsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
	if err != nil {
		return SystemsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SystemsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SystemToUpsert != nil {
		if err := runtime.MarshalAsJSON(req, *options.SystemToUpsert); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SystemsClient) createOrUpdateHandleResponse(resp *http.Response) (SystemsClientCreateOrUpdateResponse, error) {
	result := SystemsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemResource); err != nil {
		return SystemsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientDeleteOptions contains the optional parameters for the SystemsClient.Delete method.
func (client *SystemsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientDeleteOptions) (SystemsClientDeleteResponse, error) {
	var err error
	const operationName = "SystemsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
	if err != nil {
		return SystemsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SystemsClientDeleteResponse{}, err
	}
	return SystemsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SystemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientGetOptions contains the optional parameters for the SystemsClient.Get method.
func (client *SystemsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientGetOptions) (SystemsClientGetResponse, error) {
	var err error
	const operationName = "SystemsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
	if err != nil {
		return SystemsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SystemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemsClient) getHandleResponse(resp *http.Response) (SystemsClientGetResponse, error) {
	result := SystemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemResource); err != nil {
		return SystemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - ListAll the systems.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - options - SystemsClientListOptions contains the optional parameters for the SystemsClient.NewListPager method.
func (client *SystemsClient) NewListPager(resourceGroupName string, workspaceName string, agentResourceName string, options *SystemsClientListOptions) *runtime.Pager[SystemsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemsClientListResponse]{
		More: func(page SystemsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemsClientListResponse) (SystemsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SystemsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, options)
			}, nil)
			if err != nil {
				return SystemsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SystemsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, options *SystemsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SystemsClient) listHandleResponse(resp *http.Response) (SystemsClientListResponse, error) {
	result := SystemsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemsList); err != nil {
		return SystemsClientListResponse{}, err
	}
	return result, nil
}

// NewListActionsPager - List of actions for a business application system.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientListActionsOptions contains the optional parameters for the SystemsClient.NewListActionsPager method.
func (client *SystemsClient) NewListActionsPager(resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientListActionsOptions) *runtime.Pager[SystemsClientListActionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemsClientListActionsResponse]{
		More: func(page SystemsClientListActionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemsClientListActionsResponse) (SystemsClientListActionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SystemsClient.NewListActionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listActionsCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
			}, nil)
			if err != nil {
				return SystemsClientListActionsResponse{}, err
			}
			return client.listActionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listActionsCreateRequest creates the ListActions request.
func (client *SystemsClient) listActionsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientListActionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}/listActions"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listActionsHandleResponse handles the ListActions response.
func (client *SystemsClient) listActionsHandleResponse(resp *http.Response) (SystemsClientListActionsResponse, error) {
	result := SystemsClientListActionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListActionsResponse); err != nil {
		return SystemsClientListActionsResponse{}, err
	}
	return result, nil
}

// ReportActionStatus - Report the status of the action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientReportActionStatusOptions contains the optional parameters for the SystemsClient.ReportActionStatus
//     method.
func (client *SystemsClient) ReportActionStatus(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientReportActionStatusOptions) (SystemsClientReportActionStatusResponse, error) {
	var err error
	const operationName = "SystemsClient.ReportActionStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reportActionStatusCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
	if err != nil {
		return SystemsClientReportActionStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemsClientReportActionStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SystemsClientReportActionStatusResponse{}, err
	}
	return SystemsClientReportActionStatusResponse{}, nil
}

// reportActionStatusCreateRequest creates the ReportActionStatus request.
func (client *SystemsClient) reportActionStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientReportActionStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}/reportActionStatus"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// UndoAction - Undo action, based on the actionId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - systemResourceName - The name of the system.
//   - options - SystemsClientUndoActionOptions contains the optional parameters for the SystemsClient.UndoAction method.
func (client *SystemsClient) UndoAction(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientUndoActionOptions) (SystemsClientUndoActionResponse, error) {
	var err error
	const operationName = "SystemsClient.UndoAction"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.undoActionCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, systemResourceName, options)
	if err != nil {
		return SystemsClientUndoActionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemsClientUndoActionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SystemsClientUndoActionResponse{}, err
	}
	return SystemsClientUndoActionResponse{}, nil
}

// undoActionCreateRequest creates the UndoAction request.
func (client *SystemsClient) undoActionCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, systemResourceName string, options *SystemsClientUndoActionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}/systems/{systemResourceName}/undoAction"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	if systemResourceName == "" {
		return nil, errors.New("parameter systemResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemResourceName}", url.PathEscape(systemResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Payload != nil {
		if err := runtime.MarshalAsJSON(req, *options.Payload); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
