//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ConnectedEnvironmentsCertificatesClient contains the methods for the ConnectedEnvironmentsCertificates group.
// Don't use this type directly, use NewConnectedEnvironmentsCertificatesClient() instead.
type ConnectedEnvironmentsCertificatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedEnvironmentsCertificatesClient creates a new instance of ConnectedEnvironmentsCertificatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedEnvironmentsCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedEnvironmentsCertificatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedEnvironmentsCertificatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update a Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - certificateName - Name of the Certificate.
//   - options - ConnectedEnvironmentsCertificatesClientCreateOrUpdateOptions contains the optional parameters for the ConnectedEnvironmentsCertificatesClient.CreateOrUpdate
//     method.
func (client *ConnectedEnvironmentsCertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientCreateOrUpdateOptions) (ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsCertificatesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, certificateName, options)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectedEnvironmentsCertificatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CertificateEnvelope != nil {
		if err := runtime.MarshalAsJSON(req, *options.CertificateEnvelope); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectedEnvironmentsCertificatesClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse, error) {
	result := ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return ConnectedEnvironmentsCertificatesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - certificateName - Name of the Certificate.
//   - options - ConnectedEnvironmentsCertificatesClientDeleteOptions contains the optional parameters for the ConnectedEnvironmentsCertificatesClient.Delete
//     method.
func (client *ConnectedEnvironmentsCertificatesClient) Delete(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientDeleteOptions) (ConnectedEnvironmentsCertificatesClientDeleteResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsCertificatesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, certificateName, options)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsCertificatesClientDeleteResponse{}, err
	}
	return ConnectedEnvironmentsCertificatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedEnvironmentsCertificatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified Certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - certificateName - Name of the Certificate.
//   - options - ConnectedEnvironmentsCertificatesClientGetOptions contains the optional parameters for the ConnectedEnvironmentsCertificatesClient.Get
//     method.
func (client *ConnectedEnvironmentsCertificatesClient) Get(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientGetOptions) (ConnectedEnvironmentsCertificatesClientGetResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsCertificatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, certificateName, options)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsCertificatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectedEnvironmentsCertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, options *ConnectedEnvironmentsCertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedEnvironmentsCertificatesClient) getHandleResponse(resp *http.Response) (ConnectedEnvironmentsCertificatesClientGetResponse, error) {
	result := ConnectedEnvironmentsCertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return ConnectedEnvironmentsCertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Certificates in a given connected environment.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - options - ConnectedEnvironmentsCertificatesClientListOptions contains the optional parameters for the ConnectedEnvironmentsCertificatesClient.NewListPager
//     method.
func (client *ConnectedEnvironmentsCertificatesClient) NewListPager(resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsCertificatesClientListOptions) *runtime.Pager[ConnectedEnvironmentsCertificatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedEnvironmentsCertificatesClientListResponse]{
		More: func(page ConnectedEnvironmentsCertificatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedEnvironmentsCertificatesClientListResponse) (ConnectedEnvironmentsCertificatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedEnvironmentsCertificatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
			}, nil)
			if err != nil {
				return ConnectedEnvironmentsCertificatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedEnvironmentsCertificatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsCertificatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedEnvironmentsCertificatesClient) listHandleResponse(resp *http.Response) (ConnectedEnvironmentsCertificatesClientListResponse, error) {
	result := ConnectedEnvironmentsCertificatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateCollection); err != nil {
		return ConnectedEnvironmentsCertificatesClientListResponse{}, err
	}
	return result, nil
}

// Update - Patches a certificate. Currently only patching of tags is supported
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - certificateName - Name of the Certificate.
//   - certificateEnvelope - Properties of a certificate that need to be updated
//   - options - ConnectedEnvironmentsCertificatesClientUpdateOptions contains the optional parameters for the ConnectedEnvironmentsCertificatesClient.Update
//     method.
func (client *ConnectedEnvironmentsCertificatesClient) Update(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, certificateEnvelope CertificatePatch, options *ConnectedEnvironmentsCertificatesClientUpdateOptions) (ConnectedEnvironmentsCertificatesClientUpdateResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsCertificatesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, certificateName, certificateEnvelope, options)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsCertificatesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsCertificatesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ConnectedEnvironmentsCertificatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, certificateName string, certificateEnvelope CertificatePatch, options *ConnectedEnvironmentsCertificatesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, certificateEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConnectedEnvironmentsCertificatesClient) updateHandleResponse(resp *http.Response) (ConnectedEnvironmentsCertificatesClientUpdateResponse, error) {
	result := ConnectedEnvironmentsCertificatesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return ConnectedEnvironmentsCertificatesClientUpdateResponse{}, err
	}
	return result, nil
}
