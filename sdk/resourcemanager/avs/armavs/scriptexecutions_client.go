//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// ScriptExecutionsClient contains the methods for the ScriptExecutions group.
// Don't use this type directly, use NewScriptExecutionsClient() instead.
type ScriptExecutionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScriptExecutionsClient creates a new instance of ScriptExecutionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScriptExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScriptExecutionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScriptExecutionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a ScriptExecution
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - scriptExecutionName - Name of the script cmdlet.
//   - scriptExecution - Resource create parameters.
//   - options - ScriptExecutionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ScriptExecutionsClient.BeginCreateOrUpdate
//     method.
func (client *ScriptExecutionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ScriptExecutionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptExecution, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ScriptExecutionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ScriptExecutionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a ScriptExecution
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ScriptExecutionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ScriptExecutionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptExecution, options)
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
func (client *ScriptExecutionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, scriptExecution); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a ScriptExecution
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - scriptExecutionName - Name of the script cmdlet.
//   - options - ScriptExecutionsClientBeginDeleteOptions contains the optional parameters for the ScriptExecutionsClient.BeginDelete
//     method.
func (client *ScriptExecutionsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*runtime.Poller[ScriptExecutionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ScriptExecutionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ScriptExecutionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a ScriptExecution
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *ScriptExecutionsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ScriptExecutionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
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
func (client *ScriptExecutionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
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

// Get - Get a ScriptExecution
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - scriptExecutionName - Name of the script cmdlet.
//   - options - ScriptExecutionsClientGetOptions contains the optional parameters for the ScriptExecutionsClient.Get method.
func (client *ScriptExecutionsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetOptions) (ScriptExecutionsClientGetResponse, error) {
	var err error
	const operationName = "ScriptExecutionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
	if err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScriptExecutionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScriptExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScriptExecutionsClient) getHandleResponse(resp *http.Response) (ScriptExecutionsClientGetResponse, error) {
	result := ScriptExecutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecution); err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	return result, nil
}

// GetExecutionLogs - Return the logs for a script execution resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - scriptExecutionName - Name of the script cmdlet.
//   - options - ScriptExecutionsClientGetExecutionLogsOptions contains the optional parameters for the ScriptExecutionsClient.GetExecutionLogs
//     method.
func (client *ScriptExecutionsClient) GetExecutionLogs(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetExecutionLogsOptions) (ScriptExecutionsClientGetExecutionLogsResponse, error) {
	var err error
	const operationName = "ScriptExecutionsClient.GetExecutionLogs"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getExecutionLogsCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
	if err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	resp, err := client.getExecutionLogsHandleResponse(httpResp)
	return resp, err
}

// getExecutionLogsCreateRequest creates the GetExecutionLogs request.
func (client *ScriptExecutionsClient) getExecutionLogsCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetExecutionLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}/getExecutionLogs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ScriptOutputStreamType != nil {
		if err := runtime.MarshalAsJSON(req, options.ScriptOutputStreamType); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// getExecutionLogsHandleResponse handles the GetExecutionLogs response.
func (client *ScriptExecutionsClient) getExecutionLogsHandleResponse(resp *http.Response) (ScriptExecutionsClientGetExecutionLogsResponse, error) {
	result := ScriptExecutionsClientGetExecutionLogsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecution); err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	return result, nil
}

// NewListPager - List ScriptExecution resources by PrivateCloud
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - ScriptExecutionsClientListOptions contains the optional parameters for the ScriptExecutionsClient.NewListPager
//     method.
func (client *ScriptExecutionsClient) NewListPager(resourceGroupName string, privateCloudName string, options *ScriptExecutionsClientListOptions) *runtime.Pager[ScriptExecutionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScriptExecutionsClientListResponse]{
		More: func(page ScriptExecutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScriptExecutionsClientListResponse) (ScriptExecutionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScriptExecutionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			}, nil)
			if err != nil {
				return ScriptExecutionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScriptExecutionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *ScriptExecutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScriptExecutionsClient) listHandleResponse(resp *http.Response) (ScriptExecutionsClientListResponse, error) {
	result := ScriptExecutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecutionsList); err != nil {
		return ScriptExecutionsClientListResponse{}, err
	}
	return result, nil
}
