//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredatatransfer

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

// FlowsClient contains the methods for the Flows group.
// Don't use this type directly, use NewFlowsClient() instead.
type FlowsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFlowsClient creates a new instance of FlowsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFlowsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FlowsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FlowsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - flow - Flow body
//   - options - FlowsClientBeginCreateOrUpdateOptions contains the optional parameters for the FlowsClient.BeginCreateOrUpdate
//     method.
func (client *FlowsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow Flow, options *FlowsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FlowsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, connectionName, flowName, flow, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) createOrUpdate(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow Flow, options *FlowsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectionName, flowName, flow, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FlowsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow Flow, options *FlowsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, flow); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - options - FlowsClientBeginDeleteOptions contains the optional parameters for the FlowsClient.BeginDelete method.
func (client *FlowsClient) BeginDelete(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDeleteOptions) (*runtime.Poller[FlowsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, connectionName, flowName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) deleteOperation(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectionName, flowName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FlowsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDisable - Disables the specified flow
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - options - FlowsClientBeginDisableOptions contains the optional parameters for the FlowsClient.BeginDisable method.
func (client *FlowsClient) BeginDisable(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDisableOptions) (*runtime.Poller[FlowsClientDisableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disable(ctx, resourceGroupName, connectionName, flowName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientDisableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientDisableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Disable - Disables the specified flow
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) disable(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDisableOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginDisable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disableCreateRequest(ctx, resourceGroupName, connectionName, flowName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// disableCreateRequest creates the Disable request.
func (client *FlowsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}/disable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginEnable - Enables the specified flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - options - FlowsClientBeginEnableOptions contains the optional parameters for the FlowsClient.BeginEnable method.
func (client *FlowsClient) BeginEnable(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginEnableOptions) (*runtime.Poller[FlowsClientEnableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.enable(ctx, resourceGroupName, connectionName, flowName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientEnableResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientEnableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Enable - Enables the specified flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) enable(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginEnableOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginEnable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.enableCreateRequest(ctx, resourceGroupName, connectionName, flowName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// enableCreateRequest creates the Enable request.
func (client *FlowsClient) enableCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientBeginEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}/enable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - options - FlowsClientGetOptions contains the optional parameters for the FlowsClient.Get method.
func (client *FlowsClient) Get(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientGetOptions) (FlowsClientGetResponse, error) {
	var err error
	const operationName = "FlowsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectionName, flowName, options)
	if err != nil {
		return FlowsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlowsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlowsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FlowsClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, options *FlowsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FlowsClient) getHandleResponse(resp *http.Response) (FlowsClientGetResponse, error) {
	result := FlowsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Flow); err != nil {
		return FlowsClientGetResponse{}, err
	}
	return result, nil
}

// BeginLink - Links the specified flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - flow - Flow body
//   - options - FlowsClientBeginLinkOptions contains the optional parameters for the FlowsClient.BeginLink method.
func (client *FlowsClient) BeginLink(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow ResourceBody, options *FlowsClientBeginLinkOptions) (*runtime.Poller[FlowsClientLinkResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.link(ctx, resourceGroupName, connectionName, flowName, flow, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientLinkResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientLinkResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Link - Links the specified flow.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) link(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow ResourceBody, options *FlowsClientBeginLinkOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginLink"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.linkCreateRequest(ctx, resourceGroupName, connectionName, flowName, flow, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// linkCreateRequest creates the Link request.
func (client *FlowsClient) linkCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow ResourceBody, options *FlowsClientBeginLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}/link"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, flow); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListByConnectionPager - Gets flows in a connection.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - options - FlowsClientListByConnectionOptions contains the optional parameters for the FlowsClient.NewListByConnectionPager
//     method.
func (client *FlowsClient) NewListByConnectionPager(resourceGroupName string, connectionName string, options *FlowsClientListByConnectionOptions) *runtime.Pager[FlowsClientListByConnectionResponse] {
	return runtime.NewPager(runtime.PagingHandler[FlowsClientListByConnectionResponse]{
		More: func(page FlowsClientListByConnectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FlowsClientListByConnectionResponse) (FlowsClientListByConnectionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FlowsClient.NewListByConnectionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByConnectionCreateRequest(ctx, resourceGroupName, connectionName, options)
			}, nil)
			if err != nil {
				return FlowsClientListByConnectionResponse{}, err
			}
			return client.listByConnectionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByConnectionCreateRequest creates the ListByConnection request.
func (client *FlowsClient) listByConnectionCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, options *FlowsClientListByConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConnectionHandleResponse handles the ListByConnection response.
func (client *FlowsClient) listByConnectionHandleResponse(resp *http.Response) (FlowsClientListByConnectionResponse, error) {
	result := FlowsClientListByConnectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlowsListResult); err != nil {
		return FlowsClientListByConnectionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - flowName - The name for the flow that is to be onboarded.
//   - flow - Flow body
//   - options - FlowsClientBeginUpdateOptions contains the optional parameters for the FlowsClient.BeginUpdate method.
func (client *FlowsClient) BeginUpdate(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow FlowsPatch, options *FlowsClientBeginUpdateOptions) (*runtime.Poller[FlowsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, connectionName, flowName, flow, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FlowsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FlowsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the flow resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-25
func (client *FlowsClient) update(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow FlowsPatch, options *FlowsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FlowsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, connectionName, flowName, flow, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *FlowsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, flowName string, flow FlowsPatch, options *FlowsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/flows/{flowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if flowName == "" {
		return nil, errors.New("parameter flowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flowName}", url.PathEscape(flowName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, flow); err != nil {
		return nil, err
	}
	return req, nil
}
