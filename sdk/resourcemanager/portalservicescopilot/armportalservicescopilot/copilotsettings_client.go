// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armportalservicescopilot

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// CopilotSettingsClient contains the methods for the CopilotSettings group.
// Don't use this type directly, use NewCopilotSettingsClient() instead.
type CopilotSettingsClient struct {
	internal *arm.Client
}

// NewCopilotSettingsClient creates a new instance of CopilotSettingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCopilotSettingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CopilotSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CopilotSettingsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a CopilotSettingsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resource - Resource create parameters.
//   - options - CopilotSettingsClientCreateOrUpdateOptions contains the optional parameters for the CopilotSettingsClient.CreateOrUpdate
//     method.
func (client *CopilotSettingsClient) CreateOrUpdate(ctx context.Context, resource CopilotSettingsResource, options *CopilotSettingsClientCreateOrUpdateOptions) (CopilotSettingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "CopilotSettingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resource, options)
	if err != nil {
		return CopilotSettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CopilotSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CopilotSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CopilotSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resource CopilotSettingsResource, _ *CopilotSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PortalServices/copilotSettings/default"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CopilotSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (CopilotSettingsClientCreateOrUpdateResponse, error) {
	result := CopilotSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CopilotSettingsResource); err != nil {
		return CopilotSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a CopilotSettingsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - options - CopilotSettingsClientDeleteOptions contains the optional parameters for the CopilotSettingsClient.Delete method.
func (client *CopilotSettingsClient) Delete(ctx context.Context, options *CopilotSettingsClientDeleteOptions) (CopilotSettingsClientDeleteResponse, error) {
	var err error
	const operationName = "CopilotSettingsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return CopilotSettingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CopilotSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CopilotSettingsClientDeleteResponse{}, err
	}
	return CopilotSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CopilotSettingsClient) deleteCreateRequest(ctx context.Context, _ *CopilotSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PortalServices/copilotSettings/default"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a CopilotSettingsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - options - CopilotSettingsClientGetOptions contains the optional parameters for the CopilotSettingsClient.Get method.
func (client *CopilotSettingsClient) Get(ctx context.Context, options *CopilotSettingsClientGetOptions) (CopilotSettingsClientGetResponse, error) {
	var err error
	const operationName = "CopilotSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return CopilotSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CopilotSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CopilotSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CopilotSettingsClient) getCreateRequest(ctx context.Context, _ *CopilotSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PortalServices/copilotSettings/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CopilotSettingsClient) getHandleResponse(resp *http.Response) (CopilotSettingsClientGetResponse, error) {
	result := CopilotSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CopilotSettingsResource); err != nil {
		return CopilotSettingsClientGetResponse{}, err
	}
	return result, nil
}

// Update - Update a CopilotSettingsResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - properties - The resource properties to be updated.
//   - options - CopilotSettingsClientUpdateOptions contains the optional parameters for the CopilotSettingsClient.Update method.
func (client *CopilotSettingsClient) Update(ctx context.Context, properties CopilotSettingsResourceUpdate, options *CopilotSettingsClientUpdateOptions) (CopilotSettingsClientUpdateResponse, error) {
	var err error
	const operationName = "CopilotSettingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, properties, options)
	if err != nil {
		return CopilotSettingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CopilotSettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CopilotSettingsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CopilotSettingsClient) updateCreateRequest(ctx context.Context, properties CopilotSettingsResourceUpdate, _ *CopilotSettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PortalServices/copilotSettings/default"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CopilotSettingsClient) updateHandleResponse(resp *http.Response) (CopilotSettingsClientUpdateResponse, error) {
	result := CopilotSettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CopilotSettingsResource); err != nil {
		return CopilotSettingsClientUpdateResponse{}, err
	}
	return result, nil
}
