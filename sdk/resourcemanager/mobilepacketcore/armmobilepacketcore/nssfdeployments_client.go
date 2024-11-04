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

// NssfDeploymentsClient contains the methods for the NssfDeployments group.
// Don't use this type directly, use NewNssfDeploymentsClient() instead.
type NssfDeploymentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNssfDeploymentsClient creates a new instance of NssfDeploymentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNssfDeploymentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NssfDeploymentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NssfDeploymentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - nssfDeploymentName - The name of the NssfDeployment
//   - resource - Resource create parameters.
//   - options - NssfDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the NssfDeploymentsClient.BeginCreateOrUpdate
//     method.
func (client *NssfDeploymentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, nssfDeploymentName string, resource NssfDeploymentResource, options *NssfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NssfDeploymentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, nssfDeploymentName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NssfDeploymentsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NssfDeploymentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *NssfDeploymentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, nssfDeploymentName string, resource NssfDeploymentResource, options *NssfDeploymentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NssfDeploymentsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, nssfDeploymentName, resource, options)
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
func (client *NssfDeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, nssfDeploymentName string, resource NssfDeploymentResource, options *NssfDeploymentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/nssfDeployments/{nssfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if nssfDeploymentName == "" {
		return nil, errors.New("parameter nssfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nssfDeploymentName}", url.PathEscape(nssfDeploymentName))
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

// BeginDelete - Delete a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - nssfDeploymentName - The name of the NssfDeployment
//   - options - NssfDeploymentsClientBeginDeleteOptions contains the optional parameters for the NssfDeploymentsClient.BeginDelete
//     method.
func (client *NssfDeploymentsClient) BeginDelete(ctx context.Context, resourceGroupName string, nssfDeploymentName string, options *NssfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[NssfDeploymentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, nssfDeploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NssfDeploymentsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NssfDeploymentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *NssfDeploymentsClient) deleteOperation(ctx context.Context, resourceGroupName string, nssfDeploymentName string, options *NssfDeploymentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NssfDeploymentsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, nssfDeploymentName, options)
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
func (client *NssfDeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, nssfDeploymentName string, options *NssfDeploymentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/nssfDeployments/{nssfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if nssfDeploymentName == "" {
		return nil, errors.New("parameter nssfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nssfDeploymentName}", url.PathEscape(nssfDeploymentName))
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

// Get - Get a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - nssfDeploymentName - The name of the NssfDeployment
//   - options - NssfDeploymentsClientGetOptions contains the optional parameters for the NssfDeploymentsClient.Get method.
func (client *NssfDeploymentsClient) Get(ctx context.Context, resourceGroupName string, nssfDeploymentName string, options *NssfDeploymentsClientGetOptions) (NssfDeploymentsClientGetResponse, error) {
	var err error
	const operationName = "NssfDeploymentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, nssfDeploymentName, options)
	if err != nil {
		return NssfDeploymentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NssfDeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NssfDeploymentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NssfDeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, nssfDeploymentName string, options *NssfDeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/nssfDeployments/{nssfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if nssfDeploymentName == "" {
		return nil, errors.New("parameter nssfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nssfDeploymentName}", url.PathEscape(nssfDeploymentName))
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
func (client *NssfDeploymentsClient) getHandleResponse(resp *http.Response) (NssfDeploymentsClientGetResponse, error) {
	result := NssfDeploymentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NssfDeploymentResource); err != nil {
		return NssfDeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Network Slice Selection Function Deployments by Resource Group.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NssfDeploymentsClientListByResourceGroupOptions contains the optional parameters for the NssfDeploymentsClient.NewListByResourceGroupPager
//     method.
func (client *NssfDeploymentsClient) NewListByResourceGroupPager(resourceGroupName string, options *NssfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[NssfDeploymentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NssfDeploymentsClientListByResourceGroupResponse]{
		More: func(page NssfDeploymentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NssfDeploymentsClientListByResourceGroupResponse) (NssfDeploymentsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NssfDeploymentsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return NssfDeploymentsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NssfDeploymentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NssfDeploymentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/nssfDeployments"
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
func (client *NssfDeploymentsClient) listByResourceGroupHandleResponse(resp *http.Response) (NssfDeploymentsClientListByResourceGroupResponse, error) {
	result := NssfDeploymentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NssfDeploymentResourceListResult); err != nil {
		return NssfDeploymentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Network Slice Selection Function Deployments by Subscription ID.
//
// Generated from API version 2023-10-15-preview
//   - options - NssfDeploymentsClientListBySubscriptionOptions contains the optional parameters for the NssfDeploymentsClient.NewListBySubscriptionPager
//     method.
func (client *NssfDeploymentsClient) NewListBySubscriptionPager(options *NssfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[NssfDeploymentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NssfDeploymentsClientListBySubscriptionResponse]{
		More: func(page NssfDeploymentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NssfDeploymentsClientListBySubscriptionResponse) (NssfDeploymentsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NssfDeploymentsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return NssfDeploymentsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NssfDeploymentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *NssfDeploymentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobilePacketCore/nssfDeployments"
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
func (client *NssfDeploymentsClient) listBySubscriptionHandleResponse(resp *http.Response) (NssfDeploymentsClientListBySubscriptionResponse, error) {
	result := NssfDeploymentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NssfDeploymentResourceListResult); err != nil {
		return NssfDeploymentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a NssfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - nssfDeploymentName - The name of the NssfDeployment
//   - properties - The resource properties to be updated.
//   - options - NssfDeploymentsClientUpdateTagsOptions contains the optional parameters for the NssfDeploymentsClient.UpdateTags
//     method.
func (client *NssfDeploymentsClient) UpdateTags(ctx context.Context, resourceGroupName string, nssfDeploymentName string, properties NssfDeploymentResourceTagsUpdate, options *NssfDeploymentsClientUpdateTagsOptions) (NssfDeploymentsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "NssfDeploymentsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, nssfDeploymentName, properties, options)
	if err != nil {
		return NssfDeploymentsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NssfDeploymentsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NssfDeploymentsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NssfDeploymentsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, nssfDeploymentName string, properties NssfDeploymentResourceTagsUpdate, options *NssfDeploymentsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/nssfDeployments/{nssfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if nssfDeploymentName == "" {
		return nil, errors.New("parameter nssfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nssfDeploymentName}", url.PathEscape(nssfDeploymentName))
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
func (client *NssfDeploymentsClient) updateTagsHandleResponse(resp *http.Response) (NssfDeploymentsClientUpdateTagsResponse, error) {
	result := NssfDeploymentsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NssfDeploymentResource); err != nil {
		return NssfDeploymentsClientUpdateTagsResponse{}, err
	}
	return result, nil
}