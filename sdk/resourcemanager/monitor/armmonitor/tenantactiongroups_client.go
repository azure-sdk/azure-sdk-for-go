//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// TenantActionGroupsClient contains the methods for the TenantActionGroups group.
// Don't use this type directly, use NewTenantActionGroupsClient() instead.
type TenantActionGroupsClient struct {
	internal *arm.Client
}

// NewTenantActionGroupsClient creates a new instance of TenantActionGroupsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantActionGroupsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantActionGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantActionGroupsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a new tenant action group or update an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - actionGroup - The tenant action group to create or use for the update.
//   - options - TenantActionGroupsClientCreateOrUpdateOptions contains the optional parameters for the TenantActionGroupsClient.CreateOrUpdate
//     method.
func (client *TenantActionGroupsClient) CreateOrUpdate(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, actionGroup TenantActionGroupResource, options *TenantActionGroupsClientCreateOrUpdateOptions) (TenantActionGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TenantActionGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, actionGroup, options)
	if err != nil {
		return TenantActionGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantActionGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TenantActionGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TenantActionGroupsClient) createOrUpdateCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, actionGroup TenantActionGroupResource, options *TenantActionGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	if err := runtime.MarshalAsJSON(req, actionGroup); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TenantActionGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (TenantActionGroupsClientCreateOrUpdateResponse, error) {
	result := TenantActionGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantActionGroupResource); err != nil {
		return TenantActionGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a tenant action group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - options - TenantActionGroupsClientDeleteOptions contains the optional parameters for the TenantActionGroupsClient.Delete
//     method.
func (client *TenantActionGroupsClient) Delete(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *TenantActionGroupsClientDeleteOptions) (TenantActionGroupsClientDeleteResponse, error) {
	var err error
	const operationName = "TenantActionGroupsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, options)
	if err != nil {
		return TenantActionGroupsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantActionGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TenantActionGroupsClientDeleteResponse{}, err
	}
	return TenantActionGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TenantActionGroupsClient) deleteCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *TenantActionGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	return req, nil
}

// Get - Get a tenant action group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - options - TenantActionGroupsClientGetOptions contains the optional parameters for the TenantActionGroupsClient.Get method.
func (client *TenantActionGroupsClient) Get(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *TenantActionGroupsClientGetOptions) (TenantActionGroupsClientGetResponse, error) {
	var err error
	const operationName = "TenantActionGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, options)
	if err != nil {
		return TenantActionGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantActionGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TenantActionGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TenantActionGroupsClient) getCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *TenantActionGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantActionGroupsClient) getHandleResponse(resp *http.Response) (TenantActionGroupsClientGetResponse, error) {
	result := TenantActionGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantActionGroupResource); err != nil {
		return TenantActionGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagementGroupIDPager - Get a list of all tenant action groups in a management group.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - options - TenantActionGroupsClientListByManagementGroupIDOptions contains the optional parameters for the TenantActionGroupsClient.NewListByManagementGroupIDPager
//     method.
func (client *TenantActionGroupsClient) NewListByManagementGroupIDPager(managementGroupID string, xmsClientTenantID string, options *TenantActionGroupsClientListByManagementGroupIDOptions) *runtime.Pager[TenantActionGroupsClientListByManagementGroupIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantActionGroupsClientListByManagementGroupIDResponse]{
		More: func(page TenantActionGroupsClientListByManagementGroupIDResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TenantActionGroupsClientListByManagementGroupIDResponse) (TenantActionGroupsClientListByManagementGroupIDResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TenantActionGroupsClient.NewListByManagementGroupIDPager")
			req, err := client.listByManagementGroupIDCreateRequest(ctx, managementGroupID, xmsClientTenantID, options)
			if err != nil {
				return TenantActionGroupsClientListByManagementGroupIDResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TenantActionGroupsClientListByManagementGroupIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantActionGroupsClientListByManagementGroupIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagementGroupIDHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagementGroupIDCreateRequest creates the ListByManagementGroupID request.
func (client *TenantActionGroupsClient) listByManagementGroupIDCreateRequest(ctx context.Context, managementGroupID string, xmsClientTenantID string, options *TenantActionGroupsClientListByManagementGroupIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	return req, nil
}

// listByManagementGroupIDHandleResponse handles the ListByManagementGroupID response.
func (client *TenantActionGroupsClient) listByManagementGroupIDHandleResponse(resp *http.Response) (TenantActionGroupsClientListByManagementGroupIDResponse, error) {
	result := TenantActionGroupsClientListByManagementGroupIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantActionGroupList); err != nil {
		return TenantActionGroupsClientListByManagementGroupIDResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing tenant action group's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - tenantActionGroupPatch - Parameters supplied to the operation.
//   - options - TenantActionGroupsClientUpdateOptions contains the optional parameters for the TenantActionGroupsClient.Update
//     method.
func (client *TenantActionGroupsClient) Update(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, tenantActionGroupPatch ActionGroupPatchBodyAutoGenerated, options *TenantActionGroupsClientUpdateOptions) (TenantActionGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "TenantActionGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, tenantActionGroupPatch, options)
	if err != nil {
		return TenantActionGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantActionGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TenantActionGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TenantActionGroupsClient) updateCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, tenantActionGroupPatch ActionGroupPatchBodyAutoGenerated, options *TenantActionGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	if err := runtime.MarshalAsJSON(req, tenantActionGroupPatch); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TenantActionGroupsClient) updateHandleResponse(resp *http.Response) (TenantActionGroupsClientUpdateResponse, error) {
	result := TenantActionGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantActionGroupResource); err != nil {
		return TenantActionGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
