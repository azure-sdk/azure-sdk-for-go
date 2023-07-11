//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armoffazurespringboot

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

// SpringbootserversClient contains the methods for the Springbootservers group.
// Don't use this type directly, use NewSpringbootserversClient() instead.
type SpringbootserversClient struct {
	internal              *arm.Client
	siteName              string
	subscriptionID        string
	springbootserversName string
}

// NewSpringbootserversClient creates a new instance of SpringbootserversClient with the specified values.
//   - siteName - The springbootsites name.
//   - subscriptionID - The ID of the target subscription.
//   - springbootserversName - The springbootservers name.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSpringbootserversClient(siteName string, subscriptionID string, springbootserversName string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SpringbootserversClient, error) {
	cl, err := arm.NewClient(moduleName+".SpringbootserversClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SpringbootserversClient{
		siteName:              siteName,
		subscriptionID:        subscriptionID,
		springbootserversName: springbootserversName,
		internal:              cl,
	}
	return client, nil
}

// CreateOrUpdate - Create springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - springbootservers - Create a springbootservers payload.
//   - options - SpringbootserversClientCreateOrUpdateOptions contains the optional parameters for the SpringbootserversClient.CreateOrUpdate
//     method.
func (client *SpringbootserversClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, springbootservers SpringbootserversModel, options *SpringbootserversClientCreateOrUpdateOptions) (SpringbootserversClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, springbootservers, options)
	if err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SpringbootserversClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SpringbootserversClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, springbootservers SpringbootserversModel, options *SpringbootserversClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.springbootserversName == "" {
		return nil, errors.New("parameter client.springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(client.springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, springbootservers)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SpringbootserversClient) createOrUpdateHandleResponse(resp *http.Response) (SpringbootserversClientCreateOrUpdateResponse, error) {
	result := SpringbootserversClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SpringbootserversClientDeleteOptions contains the optional parameters for the SpringbootserversClient.Delete
//     method.
func (client *SpringbootserversClient) Delete(ctx context.Context, resourceGroupName string, options *SpringbootserversClientDeleteOptions) (SpringbootserversClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return SpringbootserversClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SpringbootserversClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *SpringbootserversClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, options *SpringbootserversClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.springbootserversName == "" {
		return nil, errors.New("parameter client.springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(client.springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *SpringbootserversClient) deleteHandleResponse(resp *http.Response) (SpringbootserversClientDeleteResponse, error) {
	result := SpringbootserversClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - List springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SpringbootserversClientGetOptions contains the optional parameters for the SpringbootserversClient.Get method.
func (client *SpringbootserversClient) Get(ctx context.Context, resourceGroupName string, options *SpringbootserversClientGetOptions) (SpringbootserversClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SpringbootserversClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SpringbootserversClient) getCreateRequest(ctx context.Context, resourceGroupName string, options *SpringbootserversClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.springbootserversName == "" {
		return nil, errors.New("parameter client.springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(client.springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SpringbootserversClient) getHandleResponse(resp *http.Response) (SpringbootserversClientGetResponse, error) {
	result := SpringbootserversClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List springbootservers resource by resourceGroup
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SpringbootserversClientListByResourceGroupOptions contains the optional parameters for the SpringbootserversClient.NewListByResourceGroupPager
//     method.
func (client *SpringbootserversClient) NewListByResourceGroupPager(resourceGroupName string, options *SpringbootserversClientListByResourceGroupOptions) *runtime.Pager[SpringbootserversClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpringbootserversClientListByResourceGroupResponse]{
		More: func(page SpringbootserversClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpringbootserversClientListByResourceGroupResponse) (SpringbootserversClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SpringbootserversClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SpringbootserversClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SpringbootserversClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SpringbootserversClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SpringbootserversClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
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
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SpringbootserversClient) listByResourceGroupHandleResponse(resp *http.Response) (SpringbootserversClientListByResourceGroupResponse, error) {
	result := SpringbootserversClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversListResult); err != nil {
		return SpringbootserversClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List springbootservers resource by subscription
//
// Generated from API version 2023-01-01-preview
//   - options - SpringbootserversClientListBySubscriptionOptions contains the optional parameters for the SpringbootserversClient.NewListBySubscriptionPager
//     method.
func (client *SpringbootserversClient) NewListBySubscriptionPager(options *SpringbootserversClientListBySubscriptionOptions) *runtime.Pager[SpringbootserversClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpringbootserversClientListBySubscriptionResponse]{
		More: func(page SpringbootserversClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpringbootserversClientListBySubscriptionResponse) (SpringbootserversClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SpringbootserversClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SpringbootserversClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SpringbootserversClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SpringbootserversClient) listBySubscriptionCreateRequest(ctx context.Context, options *SpringbootserversClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SpringbootserversClient) listBySubscriptionHandleResponse(resp *http.Response) (SpringbootserversClientListBySubscriptionResponse, error) {
	result := SpringbootserversClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversListResult); err != nil {
		return SpringbootserversClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update springbootservers resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - springbootservers - Update a springbootservers payload.
//   - options - SpringbootserversClientUpdateOptions contains the optional parameters for the SpringbootserversClient.Update
//     method.
func (client *SpringbootserversClient) Update(ctx context.Context, resourceGroupName string, springbootservers SpringbootserversPatch, options *SpringbootserversClientUpdateOptions) (SpringbootserversClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, springbootservers, options)
	if err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SpringbootserversClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SpringbootserversClient) updateCreateRequest(ctx context.Context, resourceGroupName string, springbootservers SpringbootserversPatch, options *SpringbootserversClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/springbootservers/{springbootserversName}"
	if client.siteName == "" {
		return nil, errors.New("parameter client.siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(client.siteName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.springbootserversName == "" {
		return nil, errors.New("parameter client.springbootserversName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{springbootserversName}", url.PathEscape(client.springbootserversName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, springbootservers)
}

// updateHandleResponse handles the Update response.
func (client *SpringbootserversClient) updateHandleResponse(resp *http.Response) (SpringbootserversClientUpdateResponse, error) {
	result := SpringbootserversClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpringbootserversModel); err != nil {
		return SpringbootserversClientUpdateResponse{}, err
	}
	return result, nil
}
