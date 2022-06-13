//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestack

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LinkedSubscriptionsClient contains the methods for the LinkedSubscriptions group.
// Don't use this type directly, use NewLinkedSubscriptionsClient() instead.
type LinkedSubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLinkedSubscriptionsClient creates a new instance of LinkedSubscriptionsClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkedSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedSubscriptionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LinkedSubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a linked subscription resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// resourceGroup - Name of the resource group.
// linkedSubscriptionName - Name of the Linked Subscription resource.
// resource - Linked subscription resource parameter.
// options - LinkedSubscriptionsClientCreateOrUpdateOptions contains the optional parameters for the LinkedSubscriptionsClient.CreateOrUpdate
// method.
func (client *LinkedSubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroup string, linkedSubscriptionName string, resource LinkedSubscriptionParameter, options *LinkedSubscriptionsClientCreateOrUpdateOptions) (LinkedSubscriptionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroup, linkedSubscriptionName, resource, options)
	if err != nil {
		return LinkedSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return LinkedSubscriptionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroup string, linkedSubscriptionName string, resource LinkedSubscriptionParameter, options *LinkedSubscriptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/linkedSubscriptions/{linkedSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if linkedSubscriptionName == "" {
		return nil, errors.New("parameter linkedSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedSubscriptionName}", url.PathEscape(linkedSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LinkedSubscriptionsClient) createOrUpdateHandleResponse(resp *http.Response) (LinkedSubscriptionsClientCreateOrUpdateResponse, error) {
	result := LinkedSubscriptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedSubscription); err != nil {
		return LinkedSubscriptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the requested Linked Subscription resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// resourceGroup - Name of the resource group.
// linkedSubscriptionName - Name of the Linked Subscription resource.
// options - LinkedSubscriptionsClientDeleteOptions contains the optional parameters for the LinkedSubscriptionsClient.Delete
// method.
func (client *LinkedSubscriptionsClient) Delete(ctx context.Context, resourceGroup string, linkedSubscriptionName string, options *LinkedSubscriptionsClientDeleteOptions) (LinkedSubscriptionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, linkedSubscriptionName, options)
	if err != nil {
		return LinkedSubscriptionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return LinkedSubscriptionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return LinkedSubscriptionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LinkedSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroup string, linkedSubscriptionName string, options *LinkedSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/linkedSubscriptions/{linkedSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if linkedSubscriptionName == "" {
		return nil, errors.New("parameter linkedSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedSubscriptionName}", url.PathEscape(linkedSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the properties of a Linked Subscription resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// resourceGroup - Name of the resource group.
// linkedSubscriptionName - Name of the Linked Subscription resource.
// options - LinkedSubscriptionsClientGetOptions contains the optional parameters for the LinkedSubscriptionsClient.Get method.
func (client *LinkedSubscriptionsClient) Get(ctx context.Context, resourceGroup string, linkedSubscriptionName string, options *LinkedSubscriptionsClientGetOptions) (LinkedSubscriptionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, linkedSubscriptionName, options)
	if err != nil {
		return LinkedSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedSubscriptionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroup string, linkedSubscriptionName string, options *LinkedSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/linkedSubscriptions/{linkedSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if linkedSubscriptionName == "" {
		return nil, errors.New("parameter linkedSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedSubscriptionName}", url.PathEscape(linkedSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedSubscriptionsClient) getHandleResponse(resp *http.Response) (LinkedSubscriptionsClientGetResponse, error) {
	result := LinkedSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedSubscription); err != nil {
		return LinkedSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns a list of all linked subscriptions under current resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// resourceGroup - Name of the resource group.
// options - LinkedSubscriptionsClientListByResourceGroupOptions contains the optional parameters for the LinkedSubscriptionsClient.ListByResourceGroup
// method.
func (client *LinkedSubscriptionsClient) NewListByResourceGroupPager(resourceGroup string, options *LinkedSubscriptionsClientListByResourceGroupOptions) *runtime.Pager[LinkedSubscriptionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedSubscriptionsClientListByResourceGroupResponse]{
		More: func(page LinkedSubscriptionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkedSubscriptionsClientListByResourceGroupResponse) (LinkedSubscriptionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkedSubscriptionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LinkedSubscriptionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *LinkedSubscriptionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/linkedSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LinkedSubscriptionsClient) listByResourceGroupHandleResponse(resp *http.Response) (LinkedSubscriptionsClientListByResourceGroupResponse, error) {
	result := LinkedSubscriptionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedSubscriptionsList); err != nil {
		return LinkedSubscriptionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns a list of all linked subscriptions under current subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// options - LinkedSubscriptionsClientListBySubscriptionOptions contains the optional parameters for the LinkedSubscriptionsClient.ListBySubscription
// method.
func (client *LinkedSubscriptionsClient) NewListBySubscriptionPager(options *LinkedSubscriptionsClientListBySubscriptionOptions) *runtime.Pager[LinkedSubscriptionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedSubscriptionsClientListBySubscriptionResponse]{
		More: func(page LinkedSubscriptionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkedSubscriptionsClientListBySubscriptionResponse) (LinkedSubscriptionsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkedSubscriptionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LinkedSubscriptionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *LinkedSubscriptionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStack/linkedSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LinkedSubscriptionsClient) listBySubscriptionHandleResponse(resp *http.Response) (LinkedSubscriptionsClientListBySubscriptionResponse, error) {
	result := LinkedSubscriptionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedSubscriptionsList); err != nil {
		return LinkedSubscriptionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patch a Linked Subscription resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-01-preview
// resourceGroup - Name of the resource group.
// linkedSubscriptionName - Name of the Linked Subscription resource.
// resource - Linked subscription resource parameter.
// options - LinkedSubscriptionsClientUpdateOptions contains the optional parameters for the LinkedSubscriptionsClient.Update
// method.
func (client *LinkedSubscriptionsClient) Update(ctx context.Context, resourceGroup string, linkedSubscriptionName string, resource LinkedSubscriptionParameter, options *LinkedSubscriptionsClientUpdateOptions) (LinkedSubscriptionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroup, linkedSubscriptionName, resource, options)
	if err != nil {
		return LinkedSubscriptionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedSubscriptionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedSubscriptionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *LinkedSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroup string, linkedSubscriptionName string, resource LinkedSubscriptionParameter, options *LinkedSubscriptionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/linkedSubscriptions/{linkedSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if linkedSubscriptionName == "" {
		return nil, errors.New("parameter linkedSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedSubscriptionName}", url.PathEscape(linkedSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// updateHandleResponse handles the Update response.
func (client *LinkedSubscriptionsClient) updateHandleResponse(resp *http.Response) (LinkedSubscriptionsClientUpdateResponse, error) {
	result := LinkedSubscriptionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedSubscription); err != nil {
		return LinkedSubscriptionsClientUpdateResponse{}, err
	}
	return result, nil
}
