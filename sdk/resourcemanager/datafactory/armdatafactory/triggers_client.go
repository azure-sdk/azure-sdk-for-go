// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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

// TriggersClient contains the methods for the Triggers group.
// Don't use this type directly, use NewTriggersClient() instead.
type TriggersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTriggersClient creates a new instance of TriggersClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTriggersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TriggersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TriggersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - trigger - Trigger resource definition.
//   - options - TriggersClientCreateOrUpdateOptions contains the optional parameters for the TriggersClient.CreateOrUpdate method.
func (client *TriggersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, trigger TriggerResource, options *TriggersClientCreateOrUpdateOptions) (TriggersClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TriggersClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, triggerName, trigger, options)
	if err != nil {
		return TriggersClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TriggersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, trigger TriggerResource, options *TriggersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, trigger); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TriggersClient) createOrUpdateHandleResponse(resp *http.Response) (TriggersClientCreateOrUpdateResponse, error) {
	result := TriggersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerResource); err != nil {
		return TriggersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientDeleteOptions contains the optional parameters for the TriggersClient.Delete method.
func (client *TriggersClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientDeleteOptions) (TriggersClientDeleteResponse, error) {
	var err error
	const operationName = "TriggersClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
	if err != nil {
		return TriggersClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientDeleteResponse{}, err
	}
	return TriggersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TriggersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientGetOptions contains the optional parameters for the TriggersClient.Get method.
func (client *TriggersClient) Get(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientGetOptions) (TriggersClientGetResponse, error) {
	var err error
	const operationName = "TriggersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
	if err != nil {
		return TriggersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotModified) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TriggersClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TriggersClient) getHandleResponse(resp *http.Response) (TriggersClientGetResponse, error) {
	result := TriggersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerResource); err != nil {
		return TriggersClientGetResponse{}, err
	}
	return result, nil
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientGetEventSubscriptionStatusOptions contains the optional parameters for the TriggersClient.GetEventSubscriptionStatus
//     method.
func (client *TriggersClient) GetEventSubscriptionStatus(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientGetEventSubscriptionStatusOptions) (TriggersClientGetEventSubscriptionStatusResponse, error) {
	var err error
	const operationName = "TriggersClient.GetEventSubscriptionStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEventSubscriptionStatusCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
	if err != nil {
		return TriggersClientGetEventSubscriptionStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientGetEventSubscriptionStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientGetEventSubscriptionStatusResponse{}, err
	}
	resp, err := client.getEventSubscriptionStatusHandleResponse(httpResp)
	return resp, err
}

// getEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *TriggersClient) getEventSubscriptionStatusCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientGetEventSubscriptionStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/getEventSubscriptionStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *TriggersClient) getEventSubscriptionStatusHandleResponse(resp *http.Response) (TriggersClientGetEventSubscriptionStatusResponse, error) {
	result := TriggersClientGetEventSubscriptionStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerSubscriptionOperationStatus); err != nil {
		return TriggersClientGetEventSubscriptionStatusResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists triggers.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - TriggersClientListByFactoryOptions contains the optional parameters for the TriggersClient.NewListByFactoryPager
//     method.
func (client *TriggersClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *TriggersClientListByFactoryOptions) *runtime.Pager[TriggersClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[TriggersClientListByFactoryResponse]{
		More: func(page TriggersClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TriggersClientListByFactoryResponse) (TriggersClientListByFactoryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TriggersClient.NewListByFactoryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			}, nil)
			if err != nil {
				return TriggersClientListByFactoryResponse{}, err
			}
			return client.listByFactoryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *TriggersClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, _ *TriggersClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *TriggersClient) listByFactoryHandleResponse(resp *http.Response) (TriggersClientListByFactoryResponse, error) {
	result := TriggersClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerListResponse); err != nil {
		return TriggersClientListByFactoryResponse{}, err
	}
	return result, nil
}

// QueryByFactory - Query triggers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - filterParameters - Parameters to filter the triggers.
//   - options - TriggersClientQueryByFactoryOptions contains the optional parameters for the TriggersClient.QueryByFactory method.
func (client *TriggersClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters TriggerFilterParameters, options *TriggersClientQueryByFactoryOptions) (TriggersClientQueryByFactoryResponse, error) {
	var err error
	const operationName = "TriggersClient.QueryByFactory"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, filterParameters, options)
	if err != nil {
		return TriggersClientQueryByFactoryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggersClientQueryByFactoryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggersClientQueryByFactoryResponse{}, err
	}
	resp, err := client.queryByFactoryHandleResponse(httpResp)
	return resp, err
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *TriggersClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, filterParameters TriggerFilterParameters, _ *TriggersClientQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/querytriggers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, filterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *TriggersClient) queryByFactoryHandleResponse(resp *http.Response) (TriggersClientQueryByFactoryResponse, error) {
	result := TriggersClientQueryByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerQueryResponse); err != nil {
		return TriggersClientQueryByFactoryResponse{}, err
	}
	return result, nil
}

// BeginStart - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientBeginStartOptions contains the optional parameters for the TriggersClient.BeginStart method.
func (client *TriggersClient) BeginStart(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginStartOptions) (*runtime.Poller[TriggersClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, factoryName, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TriggersClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TriggersClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *TriggersClient) start(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "TriggersClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *TriggersClient) startCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientBeginStopOptions contains the optional parameters for the TriggersClient.BeginStop method.
func (client *TriggersClient) BeginStop(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginStopOptions) (*runtime.Poller[TriggersClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, factoryName, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TriggersClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TriggersClientStopResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Stop - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *TriggersClient) stop(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginStopOptions) (*http.Response, error) {
	var err error
	const operationName = "TriggersClient.BeginStop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// stopCreateRequest creates the Stop request.
func (client *TriggersClient) stopCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginSubscribeToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientBeginSubscribeToEventsOptions contains the optional parameters for the TriggersClient.BeginSubscribeToEvents
//     method.
func (client *TriggersClient) BeginSubscribeToEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginSubscribeToEventsOptions) (*runtime.Poller[TriggersClientSubscribeToEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.subscribeToEvents(ctx, resourceGroupName, factoryName, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TriggersClientSubscribeToEventsResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TriggersClientSubscribeToEventsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SubscribeToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *TriggersClient) subscribeToEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginSubscribeToEventsOptions) (*http.Response, error) {
	var err error
	const operationName = "TriggersClient.BeginSubscribeToEvents"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.subscribeToEventsCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
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

// subscribeToEventsCreateRequest creates the SubscribeToEvents request.
func (client *TriggersClient) subscribeToEventsCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientBeginSubscribeToEventsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/subscribeToEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUnsubscribeFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - triggerName - The trigger name.
//   - options - TriggersClientBeginUnsubscribeFromEventsOptions contains the optional parameters for the TriggersClient.BeginUnsubscribeFromEvents
//     method.
func (client *TriggersClient) BeginUnsubscribeFromEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginUnsubscribeFromEventsOptions) (*runtime.Poller[TriggersClientUnsubscribeFromEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.unsubscribeFromEvents(ctx, resourceGroupName, factoryName, triggerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TriggersClientUnsubscribeFromEventsResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TriggersClientUnsubscribeFromEventsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UnsubscribeFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *TriggersClient) unsubscribeFromEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *TriggersClientBeginUnsubscribeFromEventsOptions) (*http.Response, error) {
	var err error
	const operationName = "TriggersClient.BeginUnsubscribeFromEvents"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unsubscribeFromEventsCreateRequest(ctx, resourceGroupName, factoryName, triggerName, options)
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

// unsubscribeFromEventsCreateRequest creates the UnsubscribeFromEvents request.
func (client *TriggersClient) unsubscribeFromEventsCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, _ *TriggersClientBeginUnsubscribeFromEventsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/unsubscribeFromEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
