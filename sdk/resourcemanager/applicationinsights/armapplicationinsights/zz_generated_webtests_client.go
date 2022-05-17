//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// WebTestsClient contains the methods for the WebTests group.
// Don't use this type directly, use NewWebTestsClient() instead.
type WebTestsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebTestsClient creates a new instance of WebTestsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWebTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebTestsClient, error) {
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
	client := &WebTestsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// webTestDefinition - Properties that need to be specified to create or update an Application Insights web test definition.
// options - WebTestsClientCreateOrUpdateOptions contains the optional parameters for the WebTestsClient.CreateOrUpdate method.
func (client *WebTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsClientCreateOrUpdateOptions) (WebTestsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, webTestName, webTestDefinition, options)
	if err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WebTestsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, webTestDefinition)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebTestsClient) createOrUpdateHandleResponse(resp *http.Response) (WebTestsClientCreateOrUpdateResponse, error) {
	result := WebTestsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Application Insights web test.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// options - WebTestsClientDeleteOptions contains the optional parameters for the WebTestsClient.Delete method.
func (client *WebTestsClient) Delete(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientDeleteOptions) (WebTestsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WebTestsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WebTestsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// options - WebTestsClientGetOptions contains the optional parameters for the WebTestsClient.Get method.
func (client *WebTestsClient) Get(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientGetOptions) (WebTestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebTestsClient) getHandleResponse(resp *http.Response) (WebTestsClientGetResponse, error) {
	result := WebTestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all Application Insights web test alerts definitions within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// options - WebTestsClientListOptions contains the optional parameters for the WebTestsClient.List method.
func (client *WebTestsClient) NewListPager(options *WebTestsClientListOptions) *runtime.Pager[WebTestsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebTestsClientListResponse]{
		More: func(page WebTestsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebTestsClientListResponse) (WebTestsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebTestsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebTestsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebTestsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WebTestsClient) listCreateRequest(ctx context.Context, options *WebTestsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebTestsClient) listHandleResponse(resp *http.Response) (WebTestsClientListResponse, error) {
	result := WebTestsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListResponse{}, err
	}
	return result, nil
}

// NewListByComponentPager - Get all Application Insights web tests defined for the specified component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// componentName - The name of the Application Insights component resource.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WebTestsClientListByComponentOptions contains the optional parameters for the WebTestsClient.ListByComponent
// method.
func (client *WebTestsClient) NewListByComponentPager(componentName string, resourceGroupName string, options *WebTestsClientListByComponentOptions) *runtime.Pager[WebTestsClientListByComponentResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebTestsClientListByComponentResponse]{
		More: func(page WebTestsClientListByComponentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebTestsClientListByComponentResponse) (WebTestsClientListByComponentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByComponentCreateRequest(ctx, componentName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebTestsClientListByComponentResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebTestsClientListByComponentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebTestsClientListByComponentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByComponentHandleResponse(resp)
		},
	})
}

// listByComponentCreateRequest creates the ListByComponent request.
func (client *WebTestsClient) listByComponentCreateRequest(ctx context.Context, componentName string, resourceGroupName string, options *WebTestsClientListByComponentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests"
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByComponentHandleResponse handles the ListByComponent response.
func (client *WebTestsClient) listByComponentHandleResponse(resp *http.Response) (WebTestsClientListByComponentResponse, error) {
	result := WebTestsClientListByComponentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListByComponentResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all Application Insights web tests defined within a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WebTestsClientListByResourceGroupOptions contains the optional parameters for the WebTestsClient.ListByResourceGroup
// method.
func (client *WebTestsClient) NewListByResourceGroupPager(resourceGroupName string, options *WebTestsClientListByResourceGroupOptions) *runtime.Pager[WebTestsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebTestsClientListByResourceGroupResponse]{
		More: func(page WebTestsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebTestsClientListByResourceGroupResponse) (WebTestsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebTestsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebTestsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebTestsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WebTestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WebTestsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WebTestsClient) listByResourceGroupHandleResponse(resp *http.Response) (WebTestsClientListByResourceGroupResponse, error) {
	result := WebTestsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Creates or updates an Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// webTestTags - Updated tag information to set into the web test instance.
// options - WebTestsClientUpdateTagsOptions contains the optional parameters for the WebTestsClient.UpdateTags method.
func (client *WebTestsClient) UpdateTags(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsClientUpdateTagsOptions) (WebTestsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, webTestName, webTestTags, options)
	if err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *WebTestsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, webTestTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *WebTestsClient) updateTagsHandleResponse(resp *http.Response) (WebTestsClientUpdateTagsResponse, error) {
	result := WebTestsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
