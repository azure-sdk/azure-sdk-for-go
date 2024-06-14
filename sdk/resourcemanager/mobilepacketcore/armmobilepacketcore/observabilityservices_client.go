//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilepacketcore

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

// ObservabilityServicesClient contains the methods for the ObservabilityServices group.
// Don't use this type directly, use NewObservabilityServicesClient() instead.
type ObservabilityServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewObservabilityServicesClient creates a new instance of ObservabilityServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewObservabilityServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ObservabilityServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ObservabilityServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - observabilityServiceName - The name of the Observability Service
//   - resource - Resource create parameters.
//   - options - ObservabilityServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ObservabilityServicesClient.BeginCreateOrUpdate
//     method.
func (client *ObservabilityServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, observabilityServiceName string, resource ObservabilityServiceResource, options *ObservabilityServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ObservabilityServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, observabilityServiceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ObservabilityServicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ObservabilityServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *ObservabilityServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, observabilityServiceName string, resource ObservabilityServiceResource, options *ObservabilityServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ObservabilityServicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, observabilityServiceName, resource, options)
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
func (client *ObservabilityServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, observabilityServiceName string, resource ObservabilityServiceResource, options *ObservabilityServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/observabilityServices/{observabilityServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if observabilityServiceName == "" {
		return nil, errors.New("parameter observabilityServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{observabilityServiceName}", url.PathEscape(observabilityServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - observabilityServiceName - The name of the Observability Service
//   - options - ObservabilityServicesClientBeginDeleteOptions contains the optional parameters for the ObservabilityServicesClient.BeginDelete
//     method.
func (client *ObservabilityServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, observabilityServiceName string, options *ObservabilityServicesClientBeginDeleteOptions) (*runtime.Poller[ObservabilityServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, observabilityServiceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ObservabilityServicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ObservabilityServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *ObservabilityServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, observabilityServiceName string, options *ObservabilityServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ObservabilityServicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, observabilityServiceName, options)
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
func (client *ObservabilityServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, observabilityServiceName string, options *ObservabilityServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/observabilityServices/{observabilityServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if observabilityServiceName == "" {
		return nil, errors.New("parameter observabilityServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{observabilityServiceName}", url.PathEscape(observabilityServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - observabilityServiceName - The name of the Observability Service
//   - options - ObservabilityServicesClientGetOptions contains the optional parameters for the ObservabilityServicesClient.Get
//     method.
func (client *ObservabilityServicesClient) Get(ctx context.Context, resourceGroupName string, observabilityServiceName string, options *ObservabilityServicesClientGetOptions) (ObservabilityServicesClientGetResponse, error) {
	var err error
	const operationName = "ObservabilityServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, observabilityServiceName, options)
	if err != nil {
		return ObservabilityServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObservabilityServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObservabilityServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ObservabilityServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, observabilityServiceName string, options *ObservabilityServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/observabilityServices/{observabilityServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if observabilityServiceName == "" {
		return nil, errors.New("parameter observabilityServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{observabilityServiceName}", url.PathEscape(observabilityServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObservabilityServicesClient) getHandleResponse(resp *http.Response) (ObservabilityServicesClientGetResponse, error) {
	result := ObservabilityServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObservabilityServiceResource); err != nil {
		return ObservabilityServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Observability Services by Resource Group.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ObservabilityServicesClientListByResourceGroupOptions contains the optional parameters for the ObservabilityServicesClient.NewListByResourceGroupPager
//     method.
func (client *ObservabilityServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *ObservabilityServicesClientListByResourceGroupOptions) *runtime.Pager[ObservabilityServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObservabilityServicesClientListByResourceGroupResponse]{
		More: func(page ObservabilityServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ObservabilityServicesClientListByResourceGroupResponse) (ObservabilityServicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ObservabilityServicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ObservabilityServicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ObservabilityServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ObservabilityServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/observabilityServices"
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
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ObservabilityServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (ObservabilityServicesClientListByResourceGroupResponse, error) {
	result := ObservabilityServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObservabilityServiceResourceListResult); err != nil {
		return ObservabilityServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Observability Services by Subscription ID.
//
// Generated from API version 2023-10-15-preview
//   - options - ObservabilityServicesClientListBySubscriptionOptions contains the optional parameters for the ObservabilityServicesClient.NewListBySubscriptionPager
//     method.
func (client *ObservabilityServicesClient) NewListBySubscriptionPager(options *ObservabilityServicesClientListBySubscriptionOptions) *runtime.Pager[ObservabilityServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObservabilityServicesClientListBySubscriptionResponse]{
		More: func(page ObservabilityServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ObservabilityServicesClientListBySubscriptionResponse) (ObservabilityServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ObservabilityServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ObservabilityServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ObservabilityServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ObservabilityServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobilePacketCore/observabilityServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ObservabilityServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ObservabilityServicesClientListBySubscriptionResponse, error) {
	result := ObservabilityServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObservabilityServiceResourceListResult); err != nil {
		return ObservabilityServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a ObservabilityServiceResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - observabilityServiceName - The name of the Observability Service
//   - properties - The resource properties to be updated.
//   - options - ObservabilityServicesClientUpdateTagsOptions contains the optional parameters for the ObservabilityServicesClient.UpdateTags
//     method.
func (client *ObservabilityServicesClient) UpdateTags(ctx context.Context, resourceGroupName string, observabilityServiceName string, properties ObservabilityServiceResourceTagsUpdate, options *ObservabilityServicesClientUpdateTagsOptions) (ObservabilityServicesClientUpdateTagsResponse, error) {
	var err error
	const operationName = "ObservabilityServicesClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, observabilityServiceName, properties, options)
	if err != nil {
		return ObservabilityServicesClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObservabilityServicesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObservabilityServicesClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ObservabilityServicesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, observabilityServiceName string, properties ObservabilityServiceResourceTagsUpdate, options *ObservabilityServicesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/observabilityServices/{observabilityServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if observabilityServiceName == "" {
		return nil, errors.New("parameter observabilityServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{observabilityServiceName}", url.PathEscape(observabilityServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ObservabilityServicesClient) updateTagsHandleResponse(resp *http.Response) (ObservabilityServicesClientUpdateTagsResponse, error) {
	result := ObservabilityServicesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObservabilityServiceResource); err != nil {
		return ObservabilityServicesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
