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

// IntegrationAccountSchemasClient contains the methods for the IntegrationAccountSchemas group.
// Don't use this type directly, use NewIntegrationAccountSchemasClient() instead.
type IntegrationAccountSchemasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountSchemasClient creates a new instance of IntegrationAccountSchemasClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountSchemasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountSchemasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountSchemasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - schemaName - The integration account schema name.
//   - resource - The integration account schema.
//   - options - IntegrationAccountSchemasClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountSchemasClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountSchemasClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, resource IntegrationAccountSchema, options *IntegrationAccountSchemasClientCreateOrUpdateOptions) (IntegrationAccountSchemasClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountSchemasClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, resource, options)
	if err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountSchemasClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, resource IntegrationAccountSchema, _ *IntegrationAccountSchemasClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
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
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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
func (client *IntegrationAccountSchemasClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountSchemasClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchema); err != nil {
		return IntegrationAccountSchemasClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - schemaName - The integration account schema name.
//   - options - IntegrationAccountSchemasClientDeleteOptions contains the optional parameters for the IntegrationAccountSchemasClient.Delete
//     method.
func (client *IntegrationAccountSchemasClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientDeleteOptions) (IntegrationAccountSchemasClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountSchemasClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, options)
	if err != nil {
		return IntegrationAccountSchemasClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountSchemasClientDeleteResponse{}, err
	}
	return IntegrationAccountSchemasClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountSchemasClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, _ *IntegrationAccountSchemasClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
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
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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

// Get - Gets an integration account schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - schemaName - The integration account schema name.
//   - options - IntegrationAccountSchemasClientGetOptions contains the optional parameters for the IntegrationAccountSchemasClient.Get
//     method.
func (client *IntegrationAccountSchemasClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, options *IntegrationAccountSchemasClientGetOptions) (IntegrationAccountSchemasClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountSchemasClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, options)
	if err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountSchemasClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, _ *IntegrationAccountSchemasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}"
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
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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
func (client *IntegrationAccountSchemasClient) getHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientGetResponse, error) {
	result := IntegrationAccountSchemasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchema); err != nil {
		return IntegrationAccountSchemasClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account schemas.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountSchemasClientListOptions contains the optional parameters for the IntegrationAccountSchemasClient.NewListPager
//     method.
func (client *IntegrationAccountSchemasClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountSchemasClientListOptions) *runtime.Pager[IntegrationAccountSchemasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountSchemasClientListResponse]{
		More: func(page IntegrationAccountSchemasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountSchemasClientListResponse) (IntegrationAccountSchemasClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountSchemasClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			}, nil)
			if err != nil {
				return IntegrationAccountSchemasClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountSchemasClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountSchemasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas"
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
func (client *IntegrationAccountSchemasClient) listHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientListResponse, error) {
	result := IntegrationAccountSchemasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountSchemaListResult); err != nil {
		return IntegrationAccountSchemasClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - integrationAccountName - The integration account name.
//   - schemaName - The integration account schema name.
//   - options - IntegrationAccountSchemasClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountSchemasClient.ListContentCallbackURL
//     method.
func (client *IntegrationAccountSchemasClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, body GetCallbackURLParameters, options *IntegrationAccountSchemasClientListContentCallbackURLOptions) (IntegrationAccountSchemasClientListContentCallbackURLResponse, error) {
	var err error
	const operationName = "IntegrationAccountSchemasClient.ListContentCallbackURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, schemaName, body, options)
	if err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.listContentCallbackURLHandleResponse(httpResp)
	return resp, err
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountSchemasClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, schemaName string, body GetCallbackURLParameters, _ *IntegrationAccountSchemasClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}/listContentCallbackUrl"
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
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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
func (client *IntegrationAccountSchemasClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountSchemasClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountSchemasClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountSchemasClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
