//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// DataConnectorDefinitionsClient contains the methods for the DataConnectorDefinitions group.
// Don't use this type directly, use NewDataConnectorDefinitionsClient() instead.
type DataConnectorDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataConnectorDefinitionsClient creates a new instance of DataConnectorDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataConnectorDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataConnectorDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataConnectorDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the data connector definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorDefinitionName - The data connector definition name.
//   - connectorDefinitionInput - The data connector definition
//   - options - DataConnectorDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the DataConnectorDefinitionsClient.CreateOrUpdate
//     method.
func (client *DataConnectorDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, connectorDefinitionInput DataConnectorDefinitionClassification, options *DataConnectorDefinitionsClientCreateOrUpdateOptions) (DataConnectorDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DataConnectorDefinitionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorDefinitionName, connectorDefinitionInput, options)
	if err != nil {
		return DataConnectorDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataConnectorDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, connectorDefinitionInput DataConnectorDefinitionClassification, options *DataConnectorDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorDefinitions/{dataConnectorDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorDefinitionName == "" {
		return nil, errors.New("parameter dataConnectorDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorDefinitionName}", url.PathEscape(dataConnectorDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectorDefinitionInput); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataConnectorDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (DataConnectorDefinitionsClientCreateOrUpdateResponse, error) {
	result := DataConnectorDefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the data connector definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorDefinitionName - The data connector definition name.
//   - options - DataConnectorDefinitionsClientDeleteOptions contains the optional parameters for the DataConnectorDefinitionsClient.Delete
//     method.
func (client *DataConnectorDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, options *DataConnectorDefinitionsClientDeleteOptions) (DataConnectorDefinitionsClientDeleteResponse, error) {
	var err error
	const operationName = "DataConnectorDefinitionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorDefinitionName, options)
	if err != nil {
		return DataConnectorDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorDefinitionsClientDeleteResponse{}, err
	}
	return DataConnectorDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataConnectorDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, options *DataConnectorDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorDefinitions/{dataConnectorDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorDefinitionName == "" {
		return nil, errors.New("parameter dataConnectorDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorDefinitionName}", url.PathEscape(dataConnectorDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a data connector definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorDefinitionName - The data connector definition name.
//   - options - DataConnectorDefinitionsClientGetOptions contains the optional parameters for the DataConnectorDefinitionsClient.Get
//     method.
func (client *DataConnectorDefinitionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, options *DataConnectorDefinitionsClientGetOptions) (DataConnectorDefinitionsClientGetResponse, error) {
	var err error
	const operationName = "DataConnectorDefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorDefinitionName, options)
	if err != nil {
		return DataConnectorDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataConnectorDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorDefinitionName string, options *DataConnectorDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorDefinitions/{dataConnectorDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorDefinitionName == "" {
		return nil, errors.New("parameter dataConnectorDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorDefinitionName}", url.PathEscape(dataConnectorDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataConnectorDefinitionsClient) getHandleResponse(resp *http.Response) (DataConnectorDefinitionsClientGetResponse, error) {
	result := DataConnectorDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all data connector definitions.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - DataConnectorDefinitionsClientListOptions contains the optional parameters for the DataConnectorDefinitionsClient.NewListPager
//     method.
func (client *DataConnectorDefinitionsClient) NewListPager(resourceGroupName string, workspaceName string, options *DataConnectorDefinitionsClientListOptions) *runtime.Pager[DataConnectorDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataConnectorDefinitionsClientListResponse]{
		More: func(page DataConnectorDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataConnectorDefinitionsClientListResponse) (DataConnectorDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataConnectorDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return DataConnectorDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DataConnectorDefinitionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DataConnectorDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectorDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataConnectorDefinitionsClient) listHandleResponse(resp *http.Response) (DataConnectorDefinitionsClientListResponse, error) {
	result := DataConnectorDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectorDefinitionArmCollectionWrapper); err != nil {
		return DataConnectorDefinitionsClientListResponse{}, err
	}
	return result, nil
}
