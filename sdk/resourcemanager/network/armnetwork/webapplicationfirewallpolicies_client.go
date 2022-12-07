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

// WebApplicationFirewallPoliciesClient contains the methods for the WebApplicationFirewallPolicies group.
// Don't use this type directly, use NewWebApplicationFirewallPoliciesClient() instead.
type WebApplicationFirewallPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebApplicationFirewallPoliciesClient creates a new instance of WebApplicationFirewallPoliciesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWebApplicationFirewallPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebApplicationFirewallPoliciesClient, error) {
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
	client := &WebApplicationFirewallPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// parameters - Policy to be created.
// options - WebApplicationFirewallPoliciesClientCreateOrUpdateOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.CreateOrUpdate
// method.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesClientCreateOrUpdateOptions) (WebApplicationFirewallPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return WebApplicationFirewallPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebApplicationFirewallPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WebApplicationFirewallPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebApplicationFirewallPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy, options *WebApplicationFirewallPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebApplicationFirewallPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesClientCreateOrUpdateResponse, error) {
	result := WebApplicationFirewallPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicy); err != nil {
		return WebApplicationFirewallPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes Policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// options - WebApplicationFirewallPoliciesClientBeginDeleteOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.BeginDelete
// method.
func (client *WebApplicationFirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesClientBeginDeleteOptions) (*runtime.Poller[WebApplicationFirewallPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, policyName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WebApplicationFirewallPoliciesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WebApplicationFirewallPoliciesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes Policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
func (client *WebApplicationFirewallPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebApplicationFirewallPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve protection policy with specified name within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// policyName - The name of the policy.
// options - WebApplicationFirewallPoliciesClientGetOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.Get
// method.
func (client *WebApplicationFirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesClientGetOptions) (WebApplicationFirewallPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return WebApplicationFirewallPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebApplicationFirewallPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebApplicationFirewallPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebApplicationFirewallPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *WebApplicationFirewallPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebApplicationFirewallPoliciesClient) getHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesClientGetResponse, error) {
	result := WebApplicationFirewallPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicy); err != nil {
		return WebApplicationFirewallPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the protection policies within a resource group.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group.
// options - WebApplicationFirewallPoliciesClientListOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.List
// method.
func (client *WebApplicationFirewallPoliciesClient) NewListPager(resourceGroupName string, options *WebApplicationFirewallPoliciesClientListOptions) *runtime.Pager[WebApplicationFirewallPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebApplicationFirewallPoliciesClientListResponse]{
		More: func(page WebApplicationFirewallPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebApplicationFirewallPoliciesClientListResponse) (WebApplicationFirewallPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebApplicationFirewallPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebApplicationFirewallPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebApplicationFirewallPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WebApplicationFirewallPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *WebApplicationFirewallPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
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
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebApplicationFirewallPoliciesClient) listHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesClientListResponse, error) {
	result := WebApplicationFirewallPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicyListResult); err != nil {
		return WebApplicationFirewallPoliciesClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the WAF policies in a subscription.
// Generated from API version 2022-06-01-preview
// options - WebApplicationFirewallPoliciesClientListAllOptions contains the optional parameters for the WebApplicationFirewallPoliciesClient.ListAll
// method.
func (client *WebApplicationFirewallPoliciesClient) NewListAllPager(options *WebApplicationFirewallPoliciesClientListAllOptions) *runtime.Pager[WebApplicationFirewallPoliciesClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebApplicationFirewallPoliciesClientListAllResponse]{
		More: func(page WebApplicationFirewallPoliciesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebApplicationFirewallPoliciesClientListAllResponse) (WebApplicationFirewallPoliciesClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WebApplicationFirewallPoliciesClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WebApplicationFirewallPoliciesClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WebApplicationFirewallPoliciesClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *WebApplicationFirewallPoliciesClient) listAllCreateRequest(ctx context.Context, options *WebApplicationFirewallPoliciesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *WebApplicationFirewallPoliciesClient) listAllHandleResponse(resp *http.Response) (WebApplicationFirewallPoliciesClientListAllResponse, error) {
	result := WebApplicationFirewallPoliciesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationFirewallPolicyListResult); err != nil {
		return WebApplicationFirewallPoliciesClientListAllResponse{}, err
	}
	return result, nil
}
