//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapicenter

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

// APIDefinitionsClient contains the methods for the APIDefinitions group.
// Don't use this type directly, use NewAPIDefinitionsClient() instead.
type APIDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
	workspaceName  string
	apiName        string
	versionName    string
	filter         *string
	definitionName string
}

// NewAPIDefinitionsClient creates a new instance of APIDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - workspaceName - The name of the workspace.
//   - apiName - The name of the API.
//   - versionName - The name of the API version.
//   - definitionName - The name of the API definition.
//   - filter - OData filter parameter.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIDefinitionsClient(subscriptionID string, workspaceName string, apiName string, versionName string, definitionName string, filter *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".APIDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIDefinitionsClient{
		subscriptionID: subscriptionID,
		workspaceName:  workspaceName,
		apiName:        apiName,
		versionName:    versionName,
		definitionName: definitionName,
		filter:         filter,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing API definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - payload - API definition entity.
//   - options - APIDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the APIDefinitionsClient.CreateOrUpdate
//     method.
func (client *APIDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, payload APIDefinition, options *APIDefinitionsClientCreateOrUpdateOptions) (APIDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return APIDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, payload APIDefinition, options *APIDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (APIDefinitionsClientCreateOrUpdateResponse, error) {
	result := APIDefinitionsClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinition); err != nil {
		return APIDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified API definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - APIDefinitionsClientDeleteOptions contains the optional parameters for the APIDefinitionsClient.Delete method.
func (client *APIDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientDeleteOptions) (APIDefinitionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return APIDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientDeleteResponse{}, err
	}
	return APIDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ExportSpecification - Exports the API specification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - APIDefinitionsClientExportSpecificationOptions contains the optional parameters for the APIDefinitionsClient.ExportSpecification
//     method.
func (client *APIDefinitionsClient) ExportSpecification(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientExportSpecificationOptions) (APIDefinitionsClientExportSpecificationResponse, error) {
	var err error
	req, err := client.exportSpecificationCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return APIDefinitionsClientExportSpecificationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientExportSpecificationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientExportSpecificationResponse{}, err
	}
	resp, err := client.exportSpecificationHandleResponse(httpResp)
	return resp, err
}

// exportSpecificationCreateRequest creates the ExportSpecification request.
func (client *APIDefinitionsClient) exportSpecificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientExportSpecificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}/exportSpecification"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// exportSpecificationHandleResponse handles the ExportSpecification response.
func (client *APIDefinitionsClient) exportSpecificationHandleResponse(resp *http.Response) (APIDefinitionsClientExportSpecificationResponse, error) {
	result := APIDefinitionsClientExportSpecificationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APISpecExportResult); err != nil {
		return APIDefinitionsClientExportSpecificationResponse{}, err
	}
	return result, nil
}

// Get - Returns details of the API definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - APIDefinitionsClientGetOptions contains the optional parameters for the APIDefinitionsClient.Get method.
func (client *APIDefinitionsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientGetOptions) (APIDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return APIDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *APIDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIDefinitionsClient) getHandleResponse(resp *http.Response) (APIDefinitionsClientGetResponse, error) {
	result := APIDefinitionsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinition); err != nil {
		return APIDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified API definition exists.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - APIDefinitionsClientHeadOptions contains the optional parameters for the APIDefinitionsClient.Head method.
func (client *APIDefinitionsClient) Head(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientHeadOptions) (APIDefinitionsClientHeadResponse, error) {
	var err error
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return APIDefinitionsClientHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientHeadResponse{}, err
	}
	return APIDefinitionsClientHeadResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *APIDefinitionsClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ImportSpecification - Imports the API specification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - payload - The API specification source entity.
//   - options - APIDefinitionsClientImportSpecificationOptions contains the optional parameters for the APIDefinitionsClient.ImportSpecification
//     method.
func (client *APIDefinitionsClient) ImportSpecification(ctx context.Context, resourceGroupName string, serviceName string, payload APISpecImportSource, options *APIDefinitionsClientImportSpecificationOptions) (APIDefinitionsClientImportSpecificationResponse, error) {
	var err error
	req, err := client.importSpecificationCreateRequest(ctx, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return APIDefinitionsClientImportSpecificationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return APIDefinitionsClientImportSpecificationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return APIDefinitionsClientImportSpecificationResponse{}, err
	}
	return APIDefinitionsClientImportSpecificationResponse{}, nil
}

// importSpecificationCreateRequest creates the ImportSpecification request.
func (client *APIDefinitionsClient) importSpecificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, payload APISpecImportSource, options *APIDefinitionsClientImportSpecificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions/{definitionName}/importSpecification"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	if client.definitionName == "" {
		return nil, errors.New("parameter client.definitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{definitionName}", url.PathEscape(client.definitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListPager - Returns a collection of API definitions.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - APIDefinitionsClientListOptions contains the optional parameters for the APIDefinitionsClient.NewListPager method.
func (client *APIDefinitionsClient) NewListPager(resourceGroupName string, serviceName string, options *APIDefinitionsClientListOptions) *runtime.Pager[APIDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIDefinitionsClientListResponse]{
		More: func(page APIDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIDefinitionsClientListResponse) (APIDefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIDefinitionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return APIDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *APIDefinitionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/apis/{apiName}/versions/{versionName}/definitions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.apiName == "" {
		return nil, errors.New("parameter client.apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(client.apiName))
	if client.versionName == "" {
		return nil, errors.New("parameter client.versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(client.versionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	if client.filter != nil {
		reqQP.Set("$filter", *client.filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *APIDefinitionsClient) listHandleResponse(resp *http.Response) (APIDefinitionsClientListResponse, error) {
	result := APIDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIDefinitionCollection); err != nil {
		return APIDefinitionsClientListResponse{}, err
	}
	return result, nil
}
