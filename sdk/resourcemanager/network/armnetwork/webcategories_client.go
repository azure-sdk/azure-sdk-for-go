//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// WebCategoriesClient contains the methods for the WebCategories group.
// Don't use this type directly, use NewWebCategoriesClient() instead.
type WebCategoriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebCategoriesClient creates a new instance of WebCategoriesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebCategoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebCategoriesClient, error) {
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
	client := &WebCategoriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the specified Azure Web Category.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - name - The name of the azureWebCategory.
//   - options - WebCategoriesClientGetOptions contains the optional parameters for the WebCategoriesClient.Get method.
func (client *WebCategoriesClient) Get(ctx context.Context, name string, options *WebCategoriesClientGetOptions) (WebCategoriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return WebCategoriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebCategoriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebCategoriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebCategoriesClient) getCreateRequest(ctx context.Context, name string, options *WebCategoriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebCategoriesClient) getHandleResponse(resp *http.Response) (WebCategoriesClientGetResponse, error) {
	result := WebCategoriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureWebCategory); err != nil {
		return WebCategoriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all the Azure Web Categories in a subscription.
//
// Generated from API version 2022-11-01
//   - options - WebCategoriesClientListBySubscriptionOptions contains the optional parameters for the WebCategoriesClient.NewListBySubscriptionPager
//     method.
func (client *WebCategoriesClient) NewListBySubscriptionPager(options *WebCategoriesClientListBySubscriptionOptions) *runtime.Pager[WebCategoriesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebCategoriesClientListBySubscriptionResponse]{
		More: func(page WebCategoriesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebCategoriesClientListBySubscriptionResponse) (WebCategoriesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebCategoriesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebCategoriesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebCategoriesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WebCategoriesClient) listBySubscriptionCreateRequest(ctx context.Context, options *WebCategoriesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WebCategoriesClient) listBySubscriptionHandleResponse(resp *http.Response) (WebCategoriesClientListBySubscriptionResponse, error) {
	result := WebCategoriesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureWebCategoryListResult); err != nil {
		return WebCategoriesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
