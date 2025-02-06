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

// NatGatewaysClient contains the methods for the NatGateways group.
// Don't use this type directly, use NewNatGatewaysClient() instead.
type NatGatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNatGatewaysClient creates a new instance of NatGatewaysClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNatGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NatGatewaysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NatGatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - natGatewayName - The name of the nat gateway.
//   - parameters - Parameters supplied to the create or update nat gateway operation.
//   - options - NatGatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the NatGatewaysClient.BeginCreateOrUpdate
//     method.
func (client *NatGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[NatGatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, natGatewayName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NatGatewaysClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NatGatewaysClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *NatGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NatGatewaysClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NatGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway, options *NatGatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - natGatewayName - The name of the nat gateway.
//   - options - NatGatewaysClientBeginDeleteOptions contains the optional parameters for the NatGatewaysClient.BeginDelete method.
func (client *NatGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*runtime.Poller[NatGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, natGatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NatGatewaysClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NatGatewaysClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified nat gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *NatGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NatGatewaysClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, natGatewayName, options)
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
func (client *NatGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified nat gateway in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - natGatewayName - The name of the nat gateway.
//   - options - NatGatewaysClientGetOptions contains the optional parameters for the NatGatewaysClient.Get method.
func (client *NatGatewaysClient) Get(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientGetOptions) (NatGatewaysClientGetResponse, error) {
	var err error
	const operationName = "NatGatewaysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, natGatewayName, options)
	if err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NatGatewaysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NatGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, options *NatGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NatGatewaysClient) getHandleResponse(resp *http.Response) (NatGatewaysClientGetResponse, error) {
	result := NatGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGateway); err != nil {
		return NatGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all nat gateways in a resource group.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - options - NatGatewaysClientListOptions contains the optional parameters for the NatGatewaysClient.NewListPager method.
func (client *NatGatewaysClient) NewListPager(resourceGroupName string, options *NatGatewaysClientListOptions) *runtime.Pager[NatGatewaysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NatGatewaysClientListResponse]{
		More: func(page NatGatewaysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NatGatewaysClientListResponse) (NatGatewaysClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NatGatewaysClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return NatGatewaysClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *NatGatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *NatGatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways"
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
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NatGatewaysClient) listHandleResponse(resp *http.Response) (NatGatewaysClientListResponse, error) {
	result := NatGatewaysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGatewayListResult); err != nil {
		return NatGatewaysClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the Nat Gateways in a subscription.
//
// Generated from API version 2024-07-01
//   - options - NatGatewaysClientListAllOptions contains the optional parameters for the NatGatewaysClient.NewListAllPager method.
func (client *NatGatewaysClient) NewListAllPager(options *NatGatewaysClientListAllOptions) *runtime.Pager[NatGatewaysClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[NatGatewaysClientListAllResponse]{
		More: func(page NatGatewaysClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NatGatewaysClientListAllResponse) (NatGatewaysClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NatGatewaysClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return NatGatewaysClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *NatGatewaysClient) listAllCreateRequest(ctx context.Context, options *NatGatewaysClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *NatGatewaysClient) listAllHandleResponse(resp *http.Response) (NatGatewaysClientListAllResponse, error) {
	result := NatGatewaysClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGatewayListResult); err != nil {
		return NatGatewaysClientListAllResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates nat gateway tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - natGatewayName - The name of the nat gateway.
//   - parameters - Parameters supplied to update nat gateway tags.
//   - options - NatGatewaysClientUpdateTagsOptions contains the optional parameters for the NatGatewaysClient.UpdateTags method.
func (client *NatGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysClientUpdateTagsOptions) (NatGatewaysClientUpdateTagsResponse, error) {
	var err error
	const operationName = "NatGatewaysClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, natGatewayName, parameters, options)
	if err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *NatGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject, options *NatGatewaysClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if natGatewayName == "" {
		return nil, errors.New("parameter natGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{natGatewayName}", url.PathEscape(natGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *NatGatewaysClient) updateTagsHandleResponse(resp *http.Response) (NatGatewaysClientUpdateTagsResponse, error) {
	result := NatGatewaysClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NatGateway); err != nil {
		return NatGatewaysClientUpdateTagsResponse{}, err
	}
	return result, nil
}
