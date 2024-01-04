//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

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

// HierarchySettingsClient contains the methods for the HierarchySettings group.
// Don't use this type directly, use NewHierarchySettingsClient() instead.
type HierarchySettingsClient struct {
	internal *arm.Client
}

// NewHierarchySettingsClient creates a new instance of HierarchySettingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHierarchySettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HierarchySettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HierarchySettingsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the hierarchy settings defined at the Management Group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - createTenantSettingsRequest - Tenant level settings request parameter.
//   - options - HierarchySettingsClientCreateOrUpdateOptions contains the optional parameters for the HierarchySettingsClient.CreateOrUpdate
//     method.
func (client *HierarchySettingsClient) CreateOrUpdate(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsClientCreateOrUpdateOptions) (HierarchySettingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "HierarchySettingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, groupID, createTenantSettingsRequest, options)
	if err != nil {
		return HierarchySettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HierarchySettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HierarchySettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HierarchySettingsClient) createOrUpdateCreateRequest(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createTenantSettingsRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HierarchySettingsClient) createOrUpdateHandleResponse(resp *http.Response) (HierarchySettingsClientCreateOrUpdateResponse, error) {
	result := HierarchySettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the hierarchy settings defined at the Management Group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - options - HierarchySettingsClientDeleteOptions contains the optional parameters for the HierarchySettingsClient.Delete
//     method.
func (client *HierarchySettingsClient) Delete(ctx context.Context, groupID string, options *HierarchySettingsClientDeleteOptions) (HierarchySettingsClientDeleteResponse, error) {
	var err error
	const operationName = "HierarchySettingsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HierarchySettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HierarchySettingsClientDeleteResponse{}, err
	}
	return HierarchySettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HierarchySettingsClient) deleteCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management
// Group of the hierarchy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - options - HierarchySettingsClientGetOptions contains the optional parameters for the HierarchySettingsClient.Get method.
func (client *HierarchySettingsClient) Get(ctx context.Context, groupID string, options *HierarchySettingsClientGetOptions) (HierarchySettingsClientGetResponse, error) {
	var err error
	const operationName = "HierarchySettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HierarchySettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HierarchySettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HierarchySettingsClient) getCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HierarchySettingsClient) getHandleResponse(resp *http.Response) (HierarchySettingsClientGetResponse, error) {
	result := HierarchySettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all the hierarchy settings defined at the Management Group level. Settings can only be set on the root Management
// Group of the hierarchy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - options - HierarchySettingsClientListOptions contains the optional parameters for the HierarchySettingsClient.List method.
func (client *HierarchySettingsClient) List(ctx context.Context, groupID string, options *HierarchySettingsClientListOptions) (HierarchySettingsClientListResponse, error) {
	var err error
	const operationName = "HierarchySettingsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, groupID, options)
	if err != nil {
		return HierarchySettingsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HierarchySettingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HierarchySettingsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *HierarchySettingsClient) listCreateRequest(ctx context.Context, groupID string, options *HierarchySettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HierarchySettingsClient) listHandleResponse(resp *http.Response) (HierarchySettingsClientListResponse, error) {
	result := HierarchySettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettingsList); err != nil {
		return HierarchySettingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates the hierarchy settings defined at the Management Group level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - groupID - Management Group ID.
//   - createTenantSettingsRequest - Tenant level settings request parameter.
//   - options - HierarchySettingsClientUpdateOptions contains the optional parameters for the HierarchySettingsClient.Update
//     method.
func (client *HierarchySettingsClient) Update(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsClientUpdateOptions) (HierarchySettingsClientUpdateResponse, error) {
	var err error
	const operationName = "HierarchySettingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, groupID, createTenantSettingsRequest, options)
	if err != nil {
		return HierarchySettingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HierarchySettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HierarchySettingsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *HierarchySettingsClient) updateCreateRequest(ctx context.Context, groupID string, createTenantSettingsRequest CreateOrUpdateSettingsRequest, options *HierarchySettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/settings/default"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createTenantSettingsRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *HierarchySettingsClient) updateHandleResponse(resp *http.Response) (HierarchySettingsClientUpdateResponse, error) {
	result := HierarchySettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HierarchySettings); err != nil {
		return HierarchySettingsClientUpdateResponse{}, err
	}
	return result, nil
}
