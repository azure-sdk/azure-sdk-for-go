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

// SmfDeploymentsClient contains the methods for the SmfDeployments group.
// Don't use this type directly, use NewSmfDeploymentsClient() instead.
type SmfDeploymentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSmfDeploymentsClient creates a new instance of SmfDeploymentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSmfDeploymentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SmfDeploymentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SmfDeploymentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - smfDeploymentName - The name of the SmfDeployment
//   - resource - Resource create parameters.
//   - options - SmfDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the SmfDeploymentsClient.BeginCreateOrUpdate
//     method.
func (client *SmfDeploymentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, smfDeploymentName string, resource SmfDeploymentResource, options *SmfDeploymentsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SmfDeploymentsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, smfDeploymentName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SmfDeploymentsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SmfDeploymentsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *SmfDeploymentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, smfDeploymentName string, resource SmfDeploymentResource, options *SmfDeploymentsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SmfDeploymentsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, smfDeploymentName, resource, options)
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
func (client *SmfDeploymentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, smfDeploymentName string, resource SmfDeploymentResource, options *SmfDeploymentsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/smfDeployments/{smfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if smfDeploymentName == "" {
		return nil, errors.New("parameter smfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{smfDeploymentName}", url.PathEscape(smfDeploymentName))
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

// BeginDelete - Delete a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - smfDeploymentName - The name of the SmfDeployment
//   - options - SmfDeploymentsClientBeginDeleteOptions contains the optional parameters for the SmfDeploymentsClient.BeginDelete
//     method.
func (client *SmfDeploymentsClient) BeginDelete(ctx context.Context, resourceGroupName string, smfDeploymentName string, options *SmfDeploymentsClientBeginDeleteOptions) (*runtime.Poller[SmfDeploymentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, smfDeploymentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SmfDeploymentsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SmfDeploymentsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
func (client *SmfDeploymentsClient) deleteOperation(ctx context.Context, resourceGroupName string, smfDeploymentName string, options *SmfDeploymentsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SmfDeploymentsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, smfDeploymentName, options)
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
func (client *SmfDeploymentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, smfDeploymentName string, options *SmfDeploymentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/smfDeployments/{smfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if smfDeploymentName == "" {
		return nil, errors.New("parameter smfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{smfDeploymentName}", url.PathEscape(smfDeploymentName))
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

// Get - Get a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - smfDeploymentName - The name of the SmfDeployment
//   - options - SmfDeploymentsClientGetOptions contains the optional parameters for the SmfDeploymentsClient.Get method.
func (client *SmfDeploymentsClient) Get(ctx context.Context, resourceGroupName string, smfDeploymentName string, options *SmfDeploymentsClientGetOptions) (SmfDeploymentsClientGetResponse, error) {
	var err error
	const operationName = "SmfDeploymentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, smfDeploymentName, options)
	if err != nil {
		return SmfDeploymentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SmfDeploymentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SmfDeploymentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SmfDeploymentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, smfDeploymentName string, options *SmfDeploymentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/smfDeployments/{smfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if smfDeploymentName == "" {
		return nil, errors.New("parameter smfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{smfDeploymentName}", url.PathEscape(smfDeploymentName))
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
func (client *SmfDeploymentsClient) getHandleResponse(resp *http.Response) (SmfDeploymentsClientGetResponse, error) {
	result := SmfDeploymentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SmfDeploymentResource); err != nil {
		return SmfDeploymentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Session Management Function Deployments by Resource Group.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SmfDeploymentsClientListByResourceGroupOptions contains the optional parameters for the SmfDeploymentsClient.NewListByResourceGroupPager
//     method.
func (client *SmfDeploymentsClient) NewListByResourceGroupPager(resourceGroupName string, options *SmfDeploymentsClientListByResourceGroupOptions) *runtime.Pager[SmfDeploymentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SmfDeploymentsClientListByResourceGroupResponse]{
		More: func(page SmfDeploymentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SmfDeploymentsClientListByResourceGroupResponse) (SmfDeploymentsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SmfDeploymentsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SmfDeploymentsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SmfDeploymentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SmfDeploymentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/smfDeployments"
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
func (client *SmfDeploymentsClient) listByResourceGroupHandleResponse(resp *http.Response) (SmfDeploymentsClientListByResourceGroupResponse, error) {
	result := SmfDeploymentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SmfDeploymentResourceListResult); err != nil {
		return SmfDeploymentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Session Management Function Deployments by Subscription ID.
//
// Generated from API version 2023-10-15-preview
//   - options - SmfDeploymentsClientListBySubscriptionOptions contains the optional parameters for the SmfDeploymentsClient.NewListBySubscriptionPager
//     method.
func (client *SmfDeploymentsClient) NewListBySubscriptionPager(options *SmfDeploymentsClientListBySubscriptionOptions) *runtime.Pager[SmfDeploymentsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SmfDeploymentsClientListBySubscriptionResponse]{
		More: func(page SmfDeploymentsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SmfDeploymentsClientListBySubscriptionResponse) (SmfDeploymentsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SmfDeploymentsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SmfDeploymentsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SmfDeploymentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SmfDeploymentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobilePacketCore/smfDeployments"
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
func (client *SmfDeploymentsClient) listBySubscriptionHandleResponse(resp *http.Response) (SmfDeploymentsClientListBySubscriptionResponse, error) {
	result := SmfDeploymentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SmfDeploymentResourceListResult); err != nil {
		return SmfDeploymentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Update a SmfDeploymentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - smfDeploymentName - The name of the SmfDeployment
//   - properties - The resource properties to be updated.
//   - options - SmfDeploymentsClientUpdateTagsOptions contains the optional parameters for the SmfDeploymentsClient.UpdateTags
//     method.
func (client *SmfDeploymentsClient) UpdateTags(ctx context.Context, resourceGroupName string, smfDeploymentName string, properties SmfDeploymentResourceTagsUpdate, options *SmfDeploymentsClientUpdateTagsOptions) (SmfDeploymentsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "SmfDeploymentsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, smfDeploymentName, properties, options)
	if err != nil {
		return SmfDeploymentsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SmfDeploymentsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SmfDeploymentsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SmfDeploymentsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, smfDeploymentName string, properties SmfDeploymentResourceTagsUpdate, options *SmfDeploymentsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobilePacketCore/smfDeployments/{smfDeploymentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if smfDeploymentName == "" {
		return nil, errors.New("parameter smfDeploymentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{smfDeploymentName}", url.PathEscape(smfDeploymentName))
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
func (client *SmfDeploymentsClient) updateTagsHandleResponse(resp *http.Response) (SmfDeploymentsClientUpdateTagsResponse, error) {
	result := SmfDeploymentsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SmfDeploymentResource); err != nil {
		return SmfDeploymentsClientUpdateTagsResponse{}, err
	}
	return result, nil
}