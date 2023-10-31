//go:build go1.18
// +build go1.18

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
	"strconv"
	"strings"
)

// WorkspaceGlobalSchemaClient contains the methods for the WorkspaceGlobalSchema group.
// Don't use this type directly, use NewWorkspaceGlobalSchemaClient() instead.
type WorkspaceGlobalSchemaClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceGlobalSchemaClient creates a new instance of WorkspaceGlobalSchemaClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceGlobalSchemaClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceGlobalSchemaClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceGlobalSchemaClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceGlobalSchemaClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates new or updates existing specified Schema of the workspace in an API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - schemaID - Schema id identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - WorkspaceGlobalSchemaClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspaceGlobalSchemaClient.BeginCreateOrUpdate
//     method.
func (client *WorkspaceGlobalSchemaClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, parameters GlobalSchemaContract, options *WorkspaceGlobalSchemaClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceGlobalSchemaClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, workspaceID, schemaID, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WorkspaceGlobalSchemaClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceGlobalSchemaClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates new or updates existing specified Schema of the workspace in an API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *WorkspaceGlobalSchemaClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, parameters GlobalSchemaContract, options *WorkspaceGlobalSchemaClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, schemaID, parameters, options)
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
func (client *WorkspaceGlobalSchemaClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, parameters GlobalSchemaContract, options *WorkspaceGlobalSchemaClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/schemas/{schemaId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes specific Schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - schemaID - Schema id identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - WorkspaceGlobalSchemaClientDeleteOptions contains the optional parameters for the WorkspaceGlobalSchemaClient.Delete
//     method.
func (client *WorkspaceGlobalSchemaClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, ifMatch string, options *WorkspaceGlobalSchemaClientDeleteOptions) (WorkspaceGlobalSchemaClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, schemaID, ifMatch, options)
	if err != nil {
		return WorkspaceGlobalSchemaClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceGlobalSchemaClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceGlobalSchemaClientDeleteResponse{}, err
	}
	return WorkspaceGlobalSchemaClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceGlobalSchemaClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, ifMatch string, options *WorkspaceGlobalSchemaClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/schemas/{schemaId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the Schema specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - schemaID - Schema id identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceGlobalSchemaClientGetOptions contains the optional parameters for the WorkspaceGlobalSchemaClient.Get
//     method.
func (client *WorkspaceGlobalSchemaClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, options *WorkspaceGlobalSchemaClientGetOptions) (WorkspaceGlobalSchemaClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, schemaID, options)
	if err != nil {
		return WorkspaceGlobalSchemaClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceGlobalSchemaClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceGlobalSchemaClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceGlobalSchemaClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, options *WorkspaceGlobalSchemaClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/schemas/{schemaId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceGlobalSchemaClient) getHandleResponse(resp *http.Response) (WorkspaceGlobalSchemaClientGetResponse, error) {
	result := WorkspaceGlobalSchemaClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalSchemaContract); err != nil {
		return WorkspaceGlobalSchemaClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Schema specified by its identifier.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - schemaID - Schema id identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceGlobalSchemaClientGetEntityTagOptions contains the optional parameters for the WorkspaceGlobalSchemaClient.GetEntityTag
//     method.
func (client *WorkspaceGlobalSchemaClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, options *WorkspaceGlobalSchemaClientGetEntityTagOptions) (WorkspaceGlobalSchemaClientGetEntityTagResponse, error) {
	var err error
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, schemaID, options)
	if err != nil {
		return WorkspaceGlobalSchemaClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceGlobalSchemaClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceGlobalSchemaClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *WorkspaceGlobalSchemaClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, schemaID string, options *WorkspaceGlobalSchemaClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/schemas/{schemaId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if schemaID == "" {
		return nil, errors.New("parameter schemaID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaId}", url.PathEscape(schemaID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *WorkspaceGlobalSchemaClient) getEntityTagHandleResponse(resp *http.Response) (WorkspaceGlobalSchemaClientGetEntityTagResponse, error) {
	result := WorkspaceGlobalSchemaClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists a collection of schemas registered with workspace in a service instance.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceGlobalSchemaClientListByServiceOptions contains the optional parameters for the WorkspaceGlobalSchemaClient.NewListByServicePager
//     method.
func (client *WorkspaceGlobalSchemaClient) NewListByServicePager(resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceGlobalSchemaClientListByServiceOptions) *runtime.Pager[WorkspaceGlobalSchemaClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceGlobalSchemaClientListByServiceResponse]{
		More: func(page WorkspaceGlobalSchemaClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceGlobalSchemaClientListByServiceResponse) (WorkspaceGlobalSchemaClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceGlobalSchemaClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceGlobalSchemaClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceGlobalSchemaClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *WorkspaceGlobalSchemaClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceGlobalSchemaClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/schemas"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *WorkspaceGlobalSchemaClient) listByServiceHandleResponse(resp *http.Response) (WorkspaceGlobalSchemaClientListByServiceResponse, error) {
	result := WorkspaceGlobalSchemaClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalSchemaCollection); err != nil {
		return WorkspaceGlobalSchemaClientListByServiceResponse{}, err
	}
	return result, nil
}
