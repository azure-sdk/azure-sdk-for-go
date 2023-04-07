//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatadog

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

// TagRulesClient contains the methods for the TagRules group.
// Don't use this type directly, use NewTagRulesClient() instead.
type TagRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTagRulesClient creates a new instance of TagRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTagRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".TagRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TagRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - ruleSetName - Rule set name
//   - options - TagRulesClientCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.CreateOrUpdate method.
func (client *TagRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientCreateOrUpdateOptions) (TagRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagRulesClient) createOrUpdateHandleResponse(resp *http.Response) (TagRulesClientCreateOrUpdateResponse, error) {
	result := TagRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - ruleSetName - Rule set name
//   - options - TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
func (client *TagRulesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (TagRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagRulesClient) getHandleResponse(resp *http.Response) (TagRulesClientGetResponse, error) {
	result := TagRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the tag rules for a given monitor resource.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - options - TagRulesClientListOptions contains the optional parameters for the TagRulesClient.NewListPager method.
func (client *TagRulesClient) NewListPager(resourceGroupName string, monitorName string, options *TagRulesClientListOptions) *runtime.Pager[TagRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagRulesClientListResponse]{
		More: func(page TagRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagRulesClientListResponse) (TagRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TagRulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TagRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TagRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TagRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *TagRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/tagRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagRulesClient) listHandleResponse(resp *http.Response) (TagRulesClientListResponse, error) {
	result := TagRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRulesListResponse); err != nil {
		return TagRulesClientListResponse{}, err
	}
	return result, nil
}
