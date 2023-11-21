//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancersClient, error) {
	cl, err := arm.NewClient(moduleName+".LoadBalancersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - parameters - Parameters supplied to the create or update load balancer operation.
//   - options - LoadBalancersClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadBalancersClient.BeginCreateOrUpdate
//     method.
func (client *LoadBalancersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*runtime.Poller[LoadBalancersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LoadBalancersClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
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
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancersClientBeginDeleteOptions contains the optional parameters for the LoadBalancersClient.BeginDelete
//     method.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*runtime.Poller[LoadBalancersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
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
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancersClientGetOptions contains the optional parameters for the LoadBalancersClient.Get method.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (LoadBalancersClientGetResponse, error) {
	var err error
	const operationName = "LoadBalancersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancersClient) getHandleResponse(resp *http.Response) (LoadBalancersClientGetResponse, error) {
	result := LoadBalancersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the load balancers in a resource group.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - options - LoadBalancersClientListOptions contains the optional parameters for the LoadBalancersClient.NewListPager method.
func (client *LoadBalancersClient) NewListPager(resourceGroupName string, options *LoadBalancersClientListOptions) *runtime.Pager[LoadBalancersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancersClientListResponse]{
		More: func(page LoadBalancersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancersClientListResponse) (LoadBalancersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancersClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LoadBalancersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LoadBalancersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancersClient) listHandleResponse(resp *http.Response) (LoadBalancersClientListResponse, error) {
	result := LoadBalancersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the load balancers in a subscription.
//
// Generated from API version 2023-09-01-preview
//   - options - LoadBalancersClientListAllOptions contains the optional parameters for the LoadBalancersClient.NewListAllPager
//     method.
func (client *LoadBalancersClient) NewListAllPager(options *LoadBalancersClientListAllOptions) *runtime.Pager[LoadBalancersClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancersClientListAllResponse]{
		More: func(page LoadBalancersClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancersClientListAllResponse) (LoadBalancersClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancersClient.NewListAllPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancersClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LoadBalancersClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancersClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *LoadBalancersClient) listAllCreateRequest(ctx context.Context, options *LoadBalancersClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *LoadBalancersClient) listAllHandleResponse(resp *http.Response) (LoadBalancersClientListAllResponse, error) {
	result := LoadBalancersClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListAllResponse{}, err
	}
	return result, nil
}

// BeginListInboundNatRulePortMappings - List of inbound NAT rule port mappings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - groupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - backendPoolName - The name of the load balancer backend address pool.
//   - parameters - Query inbound NAT rule port mapping request.
//   - options - LoadBalancersClientBeginListInboundNatRulePortMappingsOptions contains the optional parameters for the LoadBalancersClient.BeginListInboundNatRulePortMappings
//     method.
func (client *LoadBalancersClient) BeginListInboundNatRulePortMappings(ctx context.Context, groupName string, loadBalancerName string, backendPoolName string, parameters QueryInboundNatRulePortMappingRequest, options *LoadBalancersClientBeginListInboundNatRulePortMappingsOptions) (*runtime.Poller[LoadBalancersClientListInboundNatRulePortMappingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listInboundNatRulePortMappings(ctx, groupName, loadBalancerName, backendPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientListInboundNatRulePortMappingsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientListInboundNatRulePortMappingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ListInboundNatRulePortMappings - List of inbound NAT rule port mappings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LoadBalancersClient) listInboundNatRulePortMappings(ctx context.Context, groupName string, loadBalancerName string, backendPoolName string, parameters QueryInboundNatRulePortMappingRequest, options *LoadBalancersClientBeginListInboundNatRulePortMappingsOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginListInboundNatRulePortMappings"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listInboundNatRulePortMappingsCreateRequest(ctx, groupName, loadBalancerName, backendPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// listInboundNatRulePortMappingsCreateRequest creates the ListInboundNatRulePortMappings request.
func (client *LoadBalancersClient) listInboundNatRulePortMappingsCreateRequest(ctx context.Context, groupName string, loadBalancerName string, backendPoolName string, parameters QueryInboundNatRulePortMappingRequest, options *LoadBalancersClientBeginListInboundNatRulePortMappingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendPoolName}/queryInboundNatRulePortMapping"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if backendPoolName == "" {
		return nil, errors.New("parameter backendPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backendPoolName}", url.PathEscape(backendPoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// MigrateToIPBased - Migrate load balancer to IP Based
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - groupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancersClientMigrateToIPBasedOptions contains the optional parameters for the LoadBalancersClient.MigrateToIPBased
//     method.
func (client *LoadBalancersClient) MigrateToIPBased(ctx context.Context, groupName string, loadBalancerName string, options *LoadBalancersClientMigrateToIPBasedOptions) (LoadBalancersClientMigrateToIPBasedResponse, error) {
	var err error
	const operationName = "LoadBalancersClient.MigrateToIPBased"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.migrateToIPBasedCreateRequest(ctx, groupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientMigrateToIPBasedResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersClientMigrateToIPBasedResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancersClientMigrateToIPBasedResponse{}, err
	}
	resp, err := client.migrateToIPBasedHandleResponse(httpResp)
	return resp, err
}

// migrateToIPBasedCreateRequest creates the MigrateToIPBased request.
func (client *LoadBalancersClient) migrateToIPBasedCreateRequest(ctx context.Context, groupName string, loadBalancerName string, options *LoadBalancersClientMigrateToIPBasedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/migrateToIpBased"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// migrateToIPBasedHandleResponse handles the MigrateToIPBased response.
func (client *LoadBalancersClient) migrateToIPBasedHandleResponse(resp *http.Response) (LoadBalancersClientMigrateToIPBasedResponse, error) {
	result := LoadBalancersClientMigrateToIPBasedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigratedPools); err != nil {
		return LoadBalancersClientMigrateToIPBasedResponse{}, err
	}
	return result, nil
}

// BeginSwapPublicIPAddresses - Swaps VIPs between two load balancers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - location - The region where load balancers are located at.
//   - parameters - Parameters that define which VIPs should be swapped.
//   - options - LoadBalancersClientBeginSwapPublicIPAddressesOptions contains the optional parameters for the LoadBalancersClient.BeginSwapPublicIPAddresses
//     method.
func (client *LoadBalancersClient) BeginSwapPublicIPAddresses(ctx context.Context, location string, parameters LoadBalancerVipSwapRequest, options *LoadBalancersClientBeginSwapPublicIPAddressesOptions) (*runtime.Poller[LoadBalancersClientSwapPublicIPAddressesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.swapPublicIPAddresses(ctx, location, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientSwapPublicIPAddressesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientSwapPublicIPAddressesResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// SwapPublicIPAddresses - Swaps VIPs between two load balancers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *LoadBalancersClient) swapPublicIPAddresses(ctx context.Context, location string, parameters LoadBalancerVipSwapRequest, options *LoadBalancersClientBeginSwapPublicIPAddressesOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginSwapPublicIPAddresses"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.swapPublicIPAddressesCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// swapPublicIPAddressesCreateRequest creates the SwapPublicIPAddresses request.
func (client *LoadBalancersClient) swapPublicIPAddressesCreateRequest(ctx context.Context, location string, parameters LoadBalancerVipSwapRequest, options *LoadBalancersClientBeginSwapPublicIPAddressesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/setLoadBalancerFrontendPublicIpAddresses"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateTags - Updates a load balancer tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - parameters - Parameters supplied to update load balancer tags.
//   - options - LoadBalancersClientUpdateTagsOptions contains the optional parameters for the LoadBalancersClient.UpdateTags
//     method.
func (client *LoadBalancersClient) UpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientUpdateTagsOptions) (LoadBalancersClientUpdateTagsResponse, error) {
	var err error
	const operationName = "LoadBalancersClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LoadBalancersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *LoadBalancersClient) updateTagsHandleResponse(resp *http.Response) (LoadBalancersClientUpdateTagsResponse, error) {
	result := LoadBalancersClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientUpdateTagsResponse{}, err
	}
	return result, nil
}
