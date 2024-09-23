//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// BuildersClient contains the methods for the Builders group.
// Don't use this type directly, use NewBuildersClient() instead.
type BuildersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBuildersClient creates a new instance of BuildersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBuildersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BuildersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BuildersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - builderName - The name of the builder.
//   - builderEnvelope - Resource create parameters.
//   - options - BuildersClientBeginCreateOrUpdateOptions contains the optional parameters for the BuildersClient.BeginCreateOrUpdate
//     method.
func (client *BuildersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResource, options *BuildersClientBeginCreateOrUpdateOptions) (*runtime.Poller[BuildersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, builderName, builderEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BuildersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BuildersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
func (client *BuildersClient) createOrUpdate(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResource, options *BuildersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BuildersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, builderName, builderEnvelope, options)
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
func (client *BuildersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResource, options *BuildersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, builderEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - builderName - The name of the builder.
//   - options - BuildersClientBeginDeleteOptions contains the optional parameters for the BuildersClient.BeginDelete method.
func (client *BuildersClient) BeginDelete(ctx context.Context, resourceGroupName string, builderName string, options *BuildersClientBeginDeleteOptions) (*runtime.Poller[BuildersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, builderName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BuildersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BuildersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
func (client *BuildersClient) deleteOperation(ctx context.Context, resourceGroupName string, builderName string, options *BuildersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BuildersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, builderName, options)
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
func (client *BuildersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, builderName string, options *BuildersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - builderName - The name of the builder.
//   - options - BuildersClientGetOptions contains the optional parameters for the BuildersClient.Get method.
func (client *BuildersClient) Get(ctx context.Context, resourceGroupName string, builderName string, options *BuildersClientGetOptions) (BuildersClientGetResponse, error) {
	var err error
	const operationName = "BuildersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, builderName, options)
	if err != nil {
		return BuildersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BuildersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BuildersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BuildersClient) getCreateRequest(ctx context.Context, resourceGroupName string, builderName string, options *BuildersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BuildersClient) getHandleResponse(resp *http.Response) (BuildersClientGetResponse, error) {
	result := BuildersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuilderResource); err != nil {
		return BuildersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List BuilderResource resources by resource group
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - BuildersClientListByResourceGroupOptions contains the optional parameters for the BuildersClient.NewListByResourceGroupPager
//     method.
func (client *BuildersClient) NewListByResourceGroupPager(resourceGroupName string, options *BuildersClientListByResourceGroupOptions) *runtime.Pager[BuildersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[BuildersClientListByResourceGroupResponse]{
		More: func(page BuildersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BuildersClientListByResourceGroupResponse) (BuildersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BuildersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return BuildersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *BuildersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *BuildersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders"
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
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *BuildersClient) listByResourceGroupHandleResponse(resp *http.Response) (BuildersClientListByResourceGroupResponse, error) {
	result := BuildersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuilderCollection); err != nil {
		return BuildersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List BuilderResource resources by subscription ID
//
// Generated from API version 2024-08-02-preview
//   - options - BuildersClientListBySubscriptionOptions contains the optional parameters for the BuildersClient.NewListBySubscriptionPager
//     method.
func (client *BuildersClient) NewListBySubscriptionPager(options *BuildersClientListBySubscriptionOptions) *runtime.Pager[BuildersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[BuildersClientListBySubscriptionResponse]{
		More: func(page BuildersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BuildersClientListBySubscriptionResponse) (BuildersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BuildersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return BuildersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *BuildersClient) listBySubscriptionCreateRequest(ctx context.Context, options *BuildersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/builders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *BuildersClient) listBySubscriptionHandleResponse(resp *http.Response) (BuildersClientListBySubscriptionResponse, error) {
	result := BuildersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuilderCollection); err != nil {
		return BuildersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - builderName - The name of the builder.
//   - builderEnvelope - The resource properties to be updated.
//   - options - BuildersClientBeginUpdateOptions contains the optional parameters for the BuildersClient.BeginUpdate method.
func (client *BuildersClient) BeginUpdate(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResourceUpdate, options *BuildersClientBeginUpdateOptions) (*runtime.Poller[BuildersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, builderName, builderEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BuildersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BuildersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a BuilderResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
func (client *BuildersClient) update(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResourceUpdate, options *BuildersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BuildersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, builderName, builderEnvelope, options)
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
func (client *BuildersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, builderName string, builderEnvelope BuilderResourceUpdate, options *BuildersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, builderEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}
