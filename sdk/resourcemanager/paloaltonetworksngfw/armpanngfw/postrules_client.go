//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

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

// PostRulesClient contains the methods for the PostRules group.
// Don't use this type directly, use NewPostRulesClient() instead.
type PostRulesClient struct {
	internal *arm.Client
}

// NewPostRulesClient creates a new instance of PostRulesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPostRulesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PostRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".PostRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PostRulesClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a PostRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - resource - Resource create parameters.
//   - options - PostRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the PostRulesClient.BeginCreateOrUpdate
//     method.
func (client *PostRulesClient) BeginCreateOrUpdate(ctx context.Context, globalRulestackName string, priority string, resource PostRulesResource, options *PostRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PostRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, globalRulestackName, priority, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PostRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PostRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a PostRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *PostRulesClient) createOrUpdate(ctx context.Context, globalRulestackName string, priority string, resource PostRulesResource, options *PostRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, globalRulestackName, priority, resource, options)
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
func (client *PostRulesClient) createOrUpdateCreateRequest(ctx context.Context, globalRulestackName string, priority string, resource PostRulesResource, options *PostRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a PostRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - options - PostRulesClientBeginDeleteOptions contains the optional parameters for the PostRulesClient.BeginDelete method.
func (client *PostRulesClient) BeginDelete(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientBeginDeleteOptions) (*runtime.Poller[PostRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, globalRulestackName, priority, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PostRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PostRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a PostRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *PostRulesClient) deleteOperation(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, globalRulestackName, priority, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PostRulesClient) deleteCreateRequest(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a PostRulesResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - options - PostRulesClientGetOptions contains the optional parameters for the PostRulesClient.Get method.
func (client *PostRulesClient) Get(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientGetOptions) (PostRulesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, globalRulestackName, priority, options)
	if err != nil {
		return PostRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PostRulesClient) getCreateRequest(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PostRulesClient) getHandleResponse(resp *http.Response) (PostRulesClientGetResponse, error) {
	result := PostRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostRulesResource); err != nil {
		return PostRulesClientGetResponse{}, err
	}
	return result, nil
}

// GetCounters - Get counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - options - PostRulesClientGetCountersOptions contains the optional parameters for the PostRulesClient.GetCounters method.
func (client *PostRulesClient) GetCounters(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientGetCountersOptions) (PostRulesClientGetCountersResponse, error) {
	var err error
	req, err := client.getCountersCreateRequest(ctx, globalRulestackName, priority, options)
	if err != nil {
		return PostRulesClientGetCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostRulesClientGetCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostRulesClientGetCountersResponse{}, err
	}
	resp, err := client.getCountersHandleResponse(httpResp)
	return resp, err
}

// getCountersCreateRequest creates the GetCounters request.
func (client *PostRulesClient) getCountersCreateRequest(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientGetCountersOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}/getCounters"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCountersHandleResponse handles the GetCounters response.
func (client *PostRulesClient) getCountersHandleResponse(resp *http.Response) (PostRulesClientGetCountersResponse, error) {
	result := PostRulesClientGetCountersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCounter); err != nil {
		return PostRulesClientGetCountersResponse{}, err
	}
	return result, nil
}

// NewListPager - List PostRulesResource resources by Tenant
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - options - PostRulesClientListOptions contains the optional parameters for the PostRulesClient.NewListPager method.
func (client *PostRulesClient) NewListPager(globalRulestackName string, options *PostRulesClientListOptions) *runtime.Pager[PostRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PostRulesClientListResponse]{
		More: func(page PostRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PostRulesClientListResponse) (PostRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, globalRulestackName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PostRulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PostRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PostRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PostRulesClient) listCreateRequest(ctx context.Context, globalRulestackName string, options *PostRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PostRulesClient) listHandleResponse(resp *http.Response) (PostRulesClientListResponse, error) {
	result := PostRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PostRulesResourceListResult); err != nil {
		return PostRulesClientListResponse{}, err
	}
	return result, nil
}

// RefreshCounters - Refresh counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - options - PostRulesClientRefreshCountersOptions contains the optional parameters for the PostRulesClient.RefreshCounters
//     method.
func (client *PostRulesClient) RefreshCounters(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientRefreshCountersOptions) (PostRulesClientRefreshCountersResponse, error) {
	var err error
	req, err := client.refreshCountersCreateRequest(ctx, globalRulestackName, priority, options)
	if err != nil {
		return PostRulesClientRefreshCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostRulesClientRefreshCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PostRulesClientRefreshCountersResponse{}, err
	}
	return PostRulesClientRefreshCountersResponse{}, nil
}

// refreshCountersCreateRequest creates the RefreshCounters request.
func (client *PostRulesClient) refreshCountersCreateRequest(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientRefreshCountersOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}/refreshCounters"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ResetCounters - Reset counters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - globalRulestackName - GlobalRulestack resource name
//   - priority - Post Rule priority
//   - options - PostRulesClientResetCountersOptions contains the optional parameters for the PostRulesClient.ResetCounters method.
func (client *PostRulesClient) ResetCounters(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientResetCountersOptions) (PostRulesClientResetCountersResponse, error) {
	var err error
	req, err := client.resetCountersCreateRequest(ctx, globalRulestackName, priority, options)
	if err != nil {
		return PostRulesClientResetCountersResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PostRulesClientResetCountersResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PostRulesClientResetCountersResponse{}, err
	}
	resp, err := client.resetCountersHandleResponse(httpResp)
	return resp, err
}

// resetCountersCreateRequest creates the ResetCounters request.
func (client *PostRulesClient) resetCountersCreateRequest(ctx context.Context, globalRulestackName string, priority string, options *PostRulesClientResetCountersOptions) (*policy.Request, error) {
	urlPath := "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/{globalRulestackName}/postRules/{priority}/resetCounters"
	if globalRulestackName == "" {
		return nil, errors.New("parameter globalRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalRulestackName}", url.PathEscape(globalRulestackName))
	if priority == "" {
		return nil, errors.New("parameter priority cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{priority}", url.PathEscape(priority))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	if options != nil && options.FirewallName != nil {
		reqQP.Set("firewallName", *options.FirewallName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// resetCountersHandleResponse handles the ResetCounters response.
func (client *PostRulesClient) resetCountersHandleResponse(resp *http.Response) (PostRulesClientResetCountersResponse, error) {
	result := PostRulesClientResetCountersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleCounterReset); err != nil {
		return PostRulesClientResetCountersResponse{}, err
	}
	return result, nil
}
