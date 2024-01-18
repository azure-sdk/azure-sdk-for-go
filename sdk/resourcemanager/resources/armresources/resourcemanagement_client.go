//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ResourceManagementClient contains the methods for the ResourceManagementClient group.
// Don't use this type directly, use NewResourceManagementClient() instead.
type ResourceManagementClient struct {
	internal *arm.Client
}

// NewResourceManagementClient creates a new instance of ResourceManagementClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceManagementClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceManagementClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceManagementClient{
		internal: cl,
	}
	return client, nil
}

// GetDataBoundaryAtScope - Get data boundary at specified scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - scope - The resource scope.
//   - options - ResourceManagementClientGetDataBoundaryAtScopeOptions contains the optional parameters for the ResourceManagementClient.GetDataBoundaryAtScope
//     method.
func (client *ResourceManagementClient) GetDataBoundaryAtScope(ctx context.Context, scope string, options *ResourceManagementClientGetDataBoundaryAtScopeOptions) (ResourceManagementClientGetDataBoundaryAtScopeResponse, error) {
	var err error
	const operationName = "ResourceManagementClient.GetDataBoundaryAtScope"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDataBoundaryAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return ResourceManagementClientGetDataBoundaryAtScopeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceManagementClientGetDataBoundaryAtScopeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceManagementClientGetDataBoundaryAtScopeResponse{}, err
	}
	resp, err := client.getDataBoundaryAtScopeHandleResponse(httpResp)
	return resp, err
}

// getDataBoundaryAtScopeCreateRequest creates the GetDataBoundaryAtScope request.
func (client *ResourceManagementClient) getDataBoundaryAtScopeCreateRequest(ctx context.Context, scope string, options *ResourceManagementClientGetDataBoundaryAtScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/dataBoundaries/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
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

// getDataBoundaryAtScopeHandleResponse handles the GetDataBoundaryAtScope response.
func (client *ResourceManagementClient) getDataBoundaryAtScopeHandleResponse(resp *http.Response) (ResourceManagementClientGetDataBoundaryAtScopeResponse, error) {
	result := ResourceManagementClientGetDataBoundaryAtScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ResourceManagementClientGetDataBoundaryAtScopeResponse{}, err
	}
	return result, nil
}

// GetTenantDataBoundary - Get data boundary of tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - options - ResourceManagementClientGetTenantDataBoundaryOptions contains the optional parameters for the ResourceManagementClient.GetTenantDataBoundary
//     method.
func (client *ResourceManagementClient) GetTenantDataBoundary(ctx context.Context, options *ResourceManagementClientGetTenantDataBoundaryOptions) (ResourceManagementClientGetTenantDataBoundaryResponse, error) {
	var err error
	const operationName = "ResourceManagementClient.GetTenantDataBoundary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTenantDataBoundaryCreateRequest(ctx, options)
	if err != nil {
		return ResourceManagementClientGetTenantDataBoundaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceManagementClientGetTenantDataBoundaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceManagementClientGetTenantDataBoundaryResponse{}, err
	}
	resp, err := client.getTenantDataBoundaryHandleResponse(httpResp)
	return resp, err
}

// getTenantDataBoundaryCreateRequest creates the GetTenantDataBoundary request.
func (client *ResourceManagementClient) getTenantDataBoundaryCreateRequest(ctx context.Context, options *ResourceManagementClientGetTenantDataBoundaryOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/dataBoundaries/default"
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

// getTenantDataBoundaryHandleResponse handles the GetTenantDataBoundary response.
func (client *ResourceManagementClient) getTenantDataBoundaryHandleResponse(resp *http.Response) (ResourceManagementClientGetTenantDataBoundaryResponse, error) {
	result := ResourceManagementClientGetTenantDataBoundaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ResourceManagementClientGetTenantDataBoundaryResponse{}, err
	}
	return result, nil
}

// PutDataBoundary - Opt-in tenant to data boundary.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - dataBoundaryDefinition - The data boundary definition.
//   - options - ResourceManagementClientPutDataBoundaryOptions contains the optional parameters for the ResourceManagementClient.PutDataBoundary
//     method.
func (client *ResourceManagementClient) PutDataBoundary(ctx context.Context, dataBoundaryDefinition any, options *ResourceManagementClientPutDataBoundaryOptions) (ResourceManagementClientPutDataBoundaryResponse, error) {
	var err error
	const operationName = "ResourceManagementClient.PutDataBoundary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putDataBoundaryCreateRequest(ctx, dataBoundaryDefinition, options)
	if err != nil {
		return ResourceManagementClientPutDataBoundaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceManagementClientPutDataBoundaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ResourceManagementClientPutDataBoundaryResponse{}, err
	}
	resp, err := client.putDataBoundaryHandleResponse(httpResp)
	return resp, err
}

// putDataBoundaryCreateRequest creates the PutDataBoundary request.
func (client *ResourceManagementClient) putDataBoundaryCreateRequest(ctx context.Context, dataBoundaryDefinition any, options *ResourceManagementClientPutDataBoundaryOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/dataBoundaries/default"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dataBoundaryDefinition); err != nil {
		return nil, err
	}
	return req, nil
}

// putDataBoundaryHandleResponse handles the PutDataBoundary response.
func (client *ResourceManagementClient) putDataBoundaryHandleResponse(resp *http.Response) (ResourceManagementClientPutDataBoundaryResponse, error) {
	result := ResourceManagementClientPutDataBoundaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataBoundaryDefinition); err != nil {
		return ResourceManagementClientPutDataBoundaryResponse{}, err
	}
	return result, nil
}
