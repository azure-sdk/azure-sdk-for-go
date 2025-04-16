// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// PortalConfigClient contains the methods for the PortalConfig group.
// Don't use this type directly, use NewPortalConfigClient() instead.
type PortalConfigClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPortalConfigClient creates a new instance of PortalConfigClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPortalConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PortalConfigClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PortalConfigClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalConfigID - Portal configuration identifier.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update the developer portal configuration.
//   - options - PortalConfigClientCreateOrUpdateOptions contains the optional parameters for the PortalConfigClient.CreateOrUpdate
//     method.
func (client *PortalConfigClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientCreateOrUpdateOptions) (PortalConfigClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PortalConfigClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, ifMatch, parameters, options)
	if err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PortalConfigClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, _ *PortalConfigClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PortalConfigClient) createOrUpdateHandleResponse(resp *http.Response) (PortalConfigClientCreateOrUpdateResponse, error) {
	result := PortalConfigClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalConfigID - Portal configuration identifier.
//   - options - PortalConfigClientGetOptions contains the optional parameters for the PortalConfigClient.Get method.
func (client *PortalConfigClient) Get(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetOptions) (PortalConfigClientGetResponse, error) {
	var err error
	const operationName = "PortalConfigClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, options)
	if err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalConfigClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PortalConfigClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, _ *PortalConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PortalConfigClient) getHandleResponse(resp *http.Response) (PortalConfigClientGetResponse, error) {
	result := PortalConfigClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the developer portal configuration.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalConfigID - Portal configuration identifier.
//   - options - PortalConfigClientGetEntityTagOptions contains the optional parameters for the PortalConfigClient.GetEntityTag
//     method.
func (client *PortalConfigClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, options *PortalConfigClientGetEntityTagOptions) (PortalConfigClientGetEntityTagResponse, error) {
	var err error
	const operationName = "PortalConfigClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, options)
	if err != nil {
		return PortalConfigClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalConfigClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalConfigClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *PortalConfigClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, _ *PortalConfigClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *PortalConfigClient) getEntityTagHandleResponse(resp *http.Response) (PortalConfigClientGetEntityTagResponse, error) {
	result := PortalConfigClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByServicePager - Lists the developer portal configurations.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PortalConfigClientListByServiceOptions contains the optional parameters for the PortalConfigClient.NewListByServicePager
//     method.
func (client *PortalConfigClient) NewListByServicePager(resourceGroupName string, serviceName string, options *PortalConfigClientListByServiceOptions) *runtime.Pager[PortalConfigClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PortalConfigClientListByServiceResponse]{
		More: func(page PortalConfigClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PortalConfigClientListByServiceResponse) (PortalConfigClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PortalConfigClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			}, nil)
			if err != nil {
				return PortalConfigClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PortalConfigClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, _ *PortalConfigClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PortalConfigClient) listByServiceHandleResponse(resp *http.Response) (PortalConfigClientListByServiceResponse, error) {
	result := PortalConfigClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigCollection); err != nil {
		return PortalConfigClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Update the developer portal configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalConfigID - Portal configuration identifier.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update the developer portal configuration.
//   - options - PortalConfigClientUpdateOptions contains the optional parameters for the PortalConfigClient.Update method.
func (client *PortalConfigClient) Update(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, options *PortalConfigClientUpdateOptions) (PortalConfigClientUpdateResponse, error) {
	var err error
	const operationName = "PortalConfigClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, portalConfigID, ifMatch, parameters, options)
	if err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalConfigClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *PortalConfigClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalConfigID string, ifMatch string, parameters PortalConfigContract, _ *PortalConfigClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalconfigs/{portalConfigId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalConfigID == "" {
		return nil, errors.New("parameter portalConfigID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalConfigId}", url.PathEscape(portalConfigID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PortalConfigClient) updateHandleResponse(resp *http.Response) (PortalConfigClientUpdateResponse, error) {
	result := PortalConfigClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalConfigContract); err != nil {
		return PortalConfigClientUpdateResponse{}, err
	}
	return result, nil
}
