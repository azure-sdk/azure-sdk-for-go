//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

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

// ResourceGroupsClient contains the methods for the ResourceGroups group.
// Don't use this type directly, use NewResourceGroupsClient() instead.
type ResourceGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceGroupsClient creates a new instance of ResourceGroupsClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckExistence - Checks whether a resource group exists.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group to check. The name is case insensitive.
//   - options - ResourceGroupsClientCheckExistenceOptions contains the optional parameters for the ResourceGroupsClient.CheckExistence
//     method.
func (client *ResourceGroupsClient) CheckExistence(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientCheckExistenceOptions) (ResourceGroupsClientCheckExistenceResponse, error) {
	var err error
	const operationName = "ResourceGroupsClient.CheckExistence"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkExistenceCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ResourceGroupsClientCheckExistenceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGroupsClientCheckExistenceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGroupsClientCheckExistenceResponse{}, err
	}
	return ResourceGroupsClientCheckExistenceResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// checkExistenceCreateRequest creates the CheckExistence request.
func (client *ResourceGroupsClient) checkExistenceCreateRequest(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientCheckExistenceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - Creates or updates a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses,
//     hyphen, period (except at end), and Unicode characters that match the allowed characters.
//   - parameters - Parameters supplied to the create or update a resource group.
//   - options - ResourceGroupsClientCreateOrUpdateOptions contains the optional parameters for the ResourceGroupsClient.CreateOrUpdate
//     method.
func (client *ResourceGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters ResourceGroup, options *ResourceGroupsClientCreateOrUpdateOptions) (ResourceGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ResourceGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, parameters, options)
	if err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, parameters ResourceGroup, options *ResourceGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ResourceGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ResourceGroupsClientCreateOrUpdateResponse, error) {
	result := ResourceGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroup); err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes
// all of its template deployments and currently stored operations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group to delete. The name is case insensitive.
//   - options - ResourceGroupsClientBeginDeleteOptions contains the optional parameters for the ResourceGroupsClient.BeginDelete
//     method.
func (client *ResourceGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientBeginDeleteOptions) (*runtime.Poller[ResourceGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ResourceGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ResourceGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes all
// of its template deployments and currently stored operations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ResourceGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ResourceGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *ResourceGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ForceDeletionTypes != nil {
		reqQP.Set("forceDeletionTypes", *options.ForceDeletionTypes)
	}
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExportTemplate - Captures the specified resource group as a template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - Parameters for exporting the template.
//   - options - ResourceGroupsClientBeginExportTemplateOptions contains the optional parameters for the ResourceGroupsClient.BeginExportTemplate
//     method.
func (client *ResourceGroupsClient) BeginExportTemplate(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest, options *ResourceGroupsClientBeginExportTemplateOptions) (*runtime.Poller[ResourceGroupsClientExportTemplateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportTemplate(ctx, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ResourceGroupsClientExportTemplateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ResourceGroupsClientExportTemplateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportTemplate - Captures the specified resource group as a template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ResourceGroupsClient) exportTemplate(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest, options *ResourceGroupsClientBeginExportTemplateOptions) (*http.Response, error) {
	var err error
	const operationName = "ResourceGroupsClient.BeginExportTemplate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportTemplateCreateRequest(ctx, resourceGroupName, parameters, options)
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

// exportTemplateCreateRequest creates the ExportTemplate request.
func (client *ResourceGroupsClient) exportTemplateCreateRequest(ctx context.Context, resourceGroupName string, parameters ExportTemplateRequest, options *ResourceGroupsClientBeginExportTemplateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/exportTemplate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group to get. The name is case insensitive.
//   - options - ResourceGroupsClientGetOptions contains the optional parameters for the ResourceGroupsClient.Get method.
func (client *ResourceGroupsClient) Get(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientGetOptions) (ResourceGroupsClientGetResponse, error) {
	var err error
	const operationName = "ResourceGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ResourceGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, options *ResourceGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
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
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceGroupsClient) getHandleResponse(resp *http.Response) (ResourceGroupsClientGetResponse, error) {
	result := ResourceGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroup); err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the resource groups for a subscription.
//
// Generated from API version 2023-07-01
//   - options - ResourceGroupsClientListOptions contains the optional parameters for the ResourceGroupsClient.NewListPager method.
func (client *ResourceGroupsClient) NewListPager(options *ResourceGroupsClientListOptions) *runtime.Pager[ResourceGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceGroupsClientListResponse]{
		More: func(page ResourceGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceGroupsClientListResponse) (ResourceGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ResourceGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ResourceGroupsClient) listCreateRequest(ctx context.Context, options *ResourceGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceGroupsClient) listHandleResponse(resp *http.Response) (ResourceGroupsClientListResponse, error) {
	result := ResourceGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroupListResult); err != nil {
		return ResourceGroupsClientListResponse{}, err
	}
	return result, nil
}

// Update - Resource groups can be updated through a simple PATCH operation to a group address. The format of the request
// is the same as that for creating a resource group. If a field is unspecified, the current
// value is retained.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group to update. The name is case insensitive.
//   - parameters - Parameters supplied to update a resource group.
//   - options - ResourceGroupsClientUpdateOptions contains the optional parameters for the ResourceGroupsClient.Update method.
func (client *ResourceGroupsClient) Update(ctx context.Context, resourceGroupName string, parameters ResourceGroupPatchable, options *ResourceGroupsClientUpdateOptions) (ResourceGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "ResourceGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, parameters, options)
	if err != nil {
		return ResourceGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ResourceGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, parameters ResourceGroupPatchable, options *ResourceGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ResourceGroupsClient) updateHandleResponse(resp *http.Response) (ResourceGroupsClientUpdateResponse, error) {
	result := ResourceGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroup); err != nil {
		return ResourceGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
