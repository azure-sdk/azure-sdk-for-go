// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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

// JitRequestsClient contains the methods for the JitRequests group.
// Don't use this type directly, use NewJitRequestsClient() instead.
type JitRequestsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJitRequestsClient creates a new instance of JitRequestsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJitRequestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JitRequestsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JitRequestsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the JIT request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jitRequestName - The name of the JIT request.
//   - parameters - Parameters supplied to the update JIT request.
//   - options - JitRequestsClientBeginCreateOrUpdateOptions contains the optional parameters for the JitRequestsClient.BeginCreateOrUpdate
//     method.
func (client *JitRequestsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, jitRequestName string, parameters JitRequestDefinition, options *JitRequestsClientBeginCreateOrUpdateOptions) (*runtime.Poller[JitRequestsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, jitRequestName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JitRequestsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JitRequestsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the JIT request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
func (client *JitRequestsClient) createOrUpdate(ctx context.Context, resourceGroupName string, jitRequestName string, parameters JitRequestDefinition, options *JitRequestsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "JitRequestsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, jitRequestName, parameters, options)
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
func (client *JitRequestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, jitRequestName string, parameters JitRequestDefinition, _ *JitRequestsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/jitRequests/{jitRequestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jitRequestName == "" {
		return nil, errors.New("parameter jitRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jitRequestName}", url.PathEscape(jitRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes the JIT request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jitRequestName - The name of the JIT request.
//   - options - JitRequestsClientDeleteOptions contains the optional parameters for the JitRequestsClient.Delete method.
func (client *JitRequestsClient) Delete(ctx context.Context, resourceGroupName string, jitRequestName string, options *JitRequestsClientDeleteOptions) (JitRequestsClientDeleteResponse, error) {
	var err error
	const operationName = "JitRequestsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jitRequestName, options)
	if err != nil {
		return JitRequestsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JitRequestsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return JitRequestsClientDeleteResponse{}, err
	}
	return JitRequestsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JitRequestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jitRequestName string, _ *JitRequestsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/jitRequests/{jitRequestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jitRequestName == "" {
		return nil, errors.New("parameter jitRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jitRequestName}", url.PathEscape(jitRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the JIT request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jitRequestName - The name of the JIT request.
//   - options - JitRequestsClientGetOptions contains the optional parameters for the JitRequestsClient.Get method.
func (client *JitRequestsClient) Get(ctx context.Context, resourceGroupName string, jitRequestName string, options *JitRequestsClientGetOptions) (JitRequestsClientGetResponse, error) {
	var err error
	const operationName = "JitRequestsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, jitRequestName, options)
	if err != nil {
		return JitRequestsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JitRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return JitRequestsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JitRequestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jitRequestName string, _ *JitRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/jitRequests/{jitRequestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jitRequestName == "" {
		return nil, errors.New("parameter jitRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jitRequestName}", url.PathEscape(jitRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JitRequestsClient) getHandleResponse(resp *http.Response) (JitRequestsClientGetResponse, error) {
	result := JitRequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JitRequestDefinition); err != nil {
		return JitRequestsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all JIT requests within the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - JitRequestsClientListByResourceGroupOptions contains the optional parameters for the JitRequestsClient.ListByResourceGroup
//     method.
func (client *JitRequestsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *JitRequestsClientListByResourceGroupOptions) (JitRequestsClientListByResourceGroupResponse, error) {
	var err error
	const operationName = "JitRequestsClient.ListByResourceGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return JitRequestsClientListByResourceGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JitRequestsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JitRequestsClientListByResourceGroupResponse{}, err
	}
	resp, err := client.listByResourceGroupHandleResponse(httpResp)
	return resp, err
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *JitRequestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *JitRequestsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/jitRequests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *JitRequestsClient) listByResourceGroupHandleResponse(resp *http.Response) (JitRequestsClientListByResourceGroupResponse, error) {
	result := JitRequestsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JitRequestDefinitionListResult); err != nil {
		return JitRequestsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all JIT requests within the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - options - JitRequestsClientListBySubscriptionOptions contains the optional parameters for the JitRequestsClient.ListBySubscription
//     method.
func (client *JitRequestsClient) ListBySubscription(ctx context.Context, options *JitRequestsClientListBySubscriptionOptions) (JitRequestsClientListBySubscriptionResponse, error) {
	var err error
	const operationName = "JitRequestsClient.ListBySubscription"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return JitRequestsClientListBySubscriptionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JitRequestsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JitRequestsClientListBySubscriptionResponse{}, err
	}
	resp, err := client.listBySubscriptionHandleResponse(httpResp)
	return resp, err
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *JitRequestsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *JitRequestsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/jitRequests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *JitRequestsClient) listBySubscriptionHandleResponse(resp *http.Response) (JitRequestsClientListBySubscriptionResponse, error) {
	result := JitRequestsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JitRequestDefinitionListResult); err != nil {
		return JitRequestsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates the JIT request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - jitRequestName - The name of the JIT request.
//   - parameters - Parameters supplied to the update JIT request.
//   - options - JitRequestsClientUpdateOptions contains the optional parameters for the JitRequestsClient.Update method.
func (client *JitRequestsClient) Update(ctx context.Context, resourceGroupName string, jitRequestName string, parameters JitRequestPatchable, options *JitRequestsClientUpdateOptions) (JitRequestsClientUpdateResponse, error) {
	var err error
	const operationName = "JitRequestsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jitRequestName, parameters, options)
	if err != nil {
		return JitRequestsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JitRequestsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JitRequestsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *JitRequestsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jitRequestName string, parameters JitRequestPatchable, _ *JitRequestsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/jitRequests/{jitRequestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jitRequestName == "" {
		return nil, errors.New("parameter jitRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jitRequestName}", url.PathEscape(jitRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *JitRequestsClient) updateHandleResponse(resp *http.Response) (JitRequestsClientUpdateResponse, error) {
	result := JitRequestsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JitRequestDefinition); err != nil {
		return JitRequestsClientUpdateResponse{}, err
	}
	return result, nil
}
