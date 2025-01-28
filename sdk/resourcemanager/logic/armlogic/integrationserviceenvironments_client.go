//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationServiceEnvironmentsClient contains the methods for the IntegrationServiceEnvironments group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentsClient() instead.
type IntegrationServiceEnvironmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationServiceEnvironmentsClient creates a new instance of IntegrationServiceEnvironmentsClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationServiceEnvironmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationServiceEnvironmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - integrationServiceEnvironment - The integration service environment.
//   - options - IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.BeginCreateOrUpdate
//     method.
func (client *IntegrationServiceEnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IntegrationServiceEnvironmentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IntegrationServiceEnvironmentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IntegrationServiceEnvironmentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
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
func (client *IntegrationServiceEnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, integrationServiceEnvironment); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - options - IntegrationServiceEnvironmentsClientDeleteOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.Delete
//     method.
func (client *IntegrationServiceEnvironmentsClient) Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientDeleteOptions) (IntegrationServiceEnvironmentsClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationServiceEnvironmentsClientDeleteResponse{}, err
	}
	return IntegrationServiceEnvironmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationServiceEnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - options - IntegrationServiceEnvironmentsClientGetOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.Get
//     method.
func (client *IntegrationServiceEnvironmentsClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientGetOptions) (IntegrationServiceEnvironmentsClientGetResponse, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationServiceEnvironmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationServiceEnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationServiceEnvironmentsClient) getHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsClientGetResponse, error) {
	result := IntegrationServiceEnvironmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironment); err != nil {
		return IntegrationServiceEnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of integration service environments by resource group.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - options - IntegrationServiceEnvironmentsClientListByResourceGroupOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.NewListByResourceGroupPager
//     method.
func (client *IntegrationServiceEnvironmentsClient) NewListByResourceGroupPager(resourceGroup string, options *IntegrationServiceEnvironmentsClientListByResourceGroupOptions) *runtime.Pager[IntegrationServiceEnvironmentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationServiceEnvironmentsClientListByResourceGroupResponse]{
		More: func(page IntegrationServiceEnvironmentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentsClientListByResourceGroupResponse) (IntegrationServiceEnvironmentsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationServiceEnvironmentsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
			}, nil)
			if err != nil {
				return IntegrationServiceEnvironmentsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IntegrationServiceEnvironmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *IntegrationServiceEnvironmentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IntegrationServiceEnvironmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsClientListByResourceGroupResponse, error) {
	result := IntegrationServiceEnvironmentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentListResult); err != nil {
		return IntegrationServiceEnvironmentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of integration service environments by subscription.
//
// Generated from API version 2019-05-01
//   - options - IntegrationServiceEnvironmentsClientListBySubscriptionOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.NewListBySubscriptionPager
//     method.
func (client *IntegrationServiceEnvironmentsClient) NewListBySubscriptionPager(options *IntegrationServiceEnvironmentsClientListBySubscriptionOptions) *runtime.Pager[IntegrationServiceEnvironmentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationServiceEnvironmentsClientListBySubscriptionResponse]{
		More: func(page IntegrationServiceEnvironmentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentsClientListBySubscriptionResponse) (IntegrationServiceEnvironmentsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationServiceEnvironmentsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return IntegrationServiceEnvironmentsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IntegrationServiceEnvironmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *IntegrationServiceEnvironmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationServiceEnvironments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IntegrationServiceEnvironmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsClientListBySubscriptionResponse, error) {
	result := IntegrationServiceEnvironmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentListResult); err != nil {
		return IntegrationServiceEnvironmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Restart - Restarts an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - options - IntegrationServiceEnvironmentsClientRestartOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.Restart
//     method.
func (client *IntegrationServiceEnvironmentsClient) Restart(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientRestartOptions) (IntegrationServiceEnvironmentsClientRestartResponse, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentsClient.Restart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsClientRestartResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsClientRestartResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationServiceEnvironmentsClientRestartResponse{}, err
	}
	return IntegrationServiceEnvironmentsClientRestartResponse{}, nil
}

// restartCreateRequest creates the Restart request.
func (client *IntegrationServiceEnvironmentsClient) restartCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsClientRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Updates an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - integrationServiceEnvironment - The integration service environment.
//   - options - IntegrationServiceEnvironmentsClientBeginUpdateOptions contains the optional parameters for the IntegrationServiceEnvironmentsClient.BeginUpdate
//     method.
func (client *IntegrationServiceEnvironmentsClient) BeginUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginUpdateOptions) (*runtime.Poller[IntegrationServiceEnvironmentsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IntegrationServiceEnvironmentsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IntegrationServiceEnvironmentsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates an integration service environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentsClient) update(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
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

// updateCreateRequest creates the Update request.
func (client *IntegrationServiceEnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, integrationServiceEnvironment); err != nil {
		return nil, err
	}
	return req, nil
}
