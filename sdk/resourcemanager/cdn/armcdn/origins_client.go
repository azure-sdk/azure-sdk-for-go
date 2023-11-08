//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// OriginsClient contains the methods for the Origins group.
// Don't use this type directly, use NewOriginsClient() instead.
type OriginsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOriginsClient creates a new instance of OriginsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOriginsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OriginsClient, error) {
	cl, err := arm.NewClient(moduleName+".OriginsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OriginsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new origin within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originName - Name of the origin that is unique within the endpoint.
//   - origin - Origin properties
//   - options - OriginsClientBeginCreateOptions contains the optional parameters for the OriginsClient.BeginCreate method.
func (client *OriginsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*runtime.Poller[OriginsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[OriginsClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a new origin within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OriginsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, origin, options)
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

// createCreateRequest creates the Create request.
func (client *OriginsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, origin Origin, options *OriginsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, origin); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originName - Name of the origin which is unique within the endpoint.
//   - options - OriginsClientBeginDeleteOptions contains the optional parameters for the OriginsClient.BeginDelete method.
func (client *OriginsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*runtime.Poller[OriginsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, originName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[OriginsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OriginsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
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
func (client *OriginsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originName - Name of the origin which is unique within the endpoint.
//   - options - OriginsClientGetOptions contains the optional parameters for the OriginsClient.Get method.
func (client *OriginsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientGetOptions) (OriginsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, options)
	if err != nil {
		return OriginsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OriginsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OriginsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OriginsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, options *OriginsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OriginsClient) getHandleResponse(resp *http.Response) (OriginsClientGetResponse, error) {
	result := OriginsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Origin); err != nil {
		return OriginsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointPager - Lists all of the existing origins within an endpoint.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - options - OriginsClientListByEndpointOptions contains the optional parameters for the OriginsClient.NewListByEndpointPager
//     method.
func (client *OriginsClient) NewListByEndpointPager(resourceGroupName string, profileName string, endpointName string, options *OriginsClientListByEndpointOptions) *runtime.Pager[OriginsClientListByEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[OriginsClientListByEndpointResponse]{
		More: func(page OriginsClientListByEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OriginsClientListByEndpointResponse) (OriginsClientListByEndpointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OriginsClientListByEndpointResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OriginsClientListByEndpointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OriginsClientListByEndpointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEndpointHandleResponse(resp)
		},
	})
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *OriginsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *OriginsClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *OriginsClient) listByEndpointHandleResponse(resp *http.Response) (OriginsClientListByEndpointResponse, error) {
	result := OriginsClientListByEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginListResult); err != nil {
		return OriginsClientListByEndpointResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originName - Name of the origin which is unique within the endpoint.
//   - originUpdateProperties - Origin properties
//   - options - OriginsClientBeginUpdateOptions contains the optional parameters for the OriginsClient.BeginUpdate method.
func (client *OriginsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*runtime.Poller[OriginsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[OriginsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[OriginsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing origin within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OriginsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, originName, originUpdateProperties, options)
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

// updateCreateRequest creates the Update request.
func (client *OriginsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, options *OriginsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originName == "" {
		return nil, errors.New("parameter originName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originName}", url.PathEscape(originName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, originUpdateProperties); err != nil {
		return nil, err
	}
	return req, nil
}
