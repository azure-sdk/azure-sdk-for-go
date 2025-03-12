// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationAccountMapsClient contains the methods for the IntegrationAccountMaps group.
// Don't use this type directly, use NewIntegrationAccountMapsClient() instead.
type IntegrationAccountMapsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountMapsClient creates a new instance of IntegrationAccountMapsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountMapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountMapsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountMapsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account map. If the map is larger than 4 MB, you need to store the map
// in an Azure blob and use the blob's Shared Access Signature (SAS) URL as the 'contentLink'
// property value.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - mapName - The integration account map name.
//   - resource - The integration account map.
//   - options - IntegrationAccountMapsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountMapsClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountMapsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, resource IntegrationAccountMap, options *IntegrationAccountMapsClientCreateOrUpdateOptions) (IntegrationAccountMapsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountMapsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, resource, options)
	if err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountMapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, resource IntegrationAccountMap, _ *IntegrationAccountMapsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountMapsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountMapsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountMapsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMap); err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account map.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - mapName - The integration account map name.
//   - options - IntegrationAccountMapsClientDeleteOptions contains the optional parameters for the IntegrationAccountMapsClient.Delete
//     method.
func (client *IntegrationAccountMapsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientDeleteOptions) (IntegrationAccountMapsClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountMapsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, options)
	if err != nil {
		return IntegrationAccountMapsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountMapsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountMapsClientDeleteResponse{}, err
	}
	return IntegrationAccountMapsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountMapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, _ *IntegrationAccountMapsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration account map.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - mapName - The integration account map name.
//   - options - IntegrationAccountMapsClientGetOptions contains the optional parameters for the IntegrationAccountMapsClient.Get
//     method.
func (client *IntegrationAccountMapsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientGetOptions) (IntegrationAccountMapsClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountMapsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, options)
	if err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountMapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, _ *IntegrationAccountMapsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountMapsClient) getHandleResponse(resp *http.Response) (IntegrationAccountMapsClientGetResponse, error) {
	result := IntegrationAccountMapsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMap); err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account maps.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountMapsClientListOptions contains the optional parameters for the IntegrationAccountMapsClient.NewListPager
//     method.
func (client *IntegrationAccountMapsClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountMapsClientListOptions) *runtime.Pager[IntegrationAccountMapsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountMapsClientListResponse]{
		More: func(page IntegrationAccountMapsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountMapsClientListResponse) (IntegrationAccountMapsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountMapsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			}, nil)
			if err != nil {
				return IntegrationAccountMapsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountMapsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountMapsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountMapsClient) listHandleResponse(resp *http.Response) (IntegrationAccountMapsClientListResponse, error) {
	result := IntegrationAccountMapsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMapListResult); err != nil {
		return IntegrationAccountMapsClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - mapName - The integration account map name.
//   - options - IntegrationAccountMapsClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountMapsClient.ListContentCallbackURL
//     method.
func (client *IntegrationAccountMapsClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, body GetCallbackURLParameters, options *IntegrationAccountMapsClientListContentCallbackURLOptions) (IntegrationAccountMapsClientListContentCallbackURLResponse, error) {
	var err error
	const operationName = "IntegrationAccountMapsClient.ListContentCallbackURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, body, options)
	if err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.listContentCallbackURLHandleResponse(httpResp)
	return resp, err
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountMapsClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, body GetCallbackURLParameters, _ *IntegrationAccountMapsClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountMapsClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountMapsClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountMapsClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
