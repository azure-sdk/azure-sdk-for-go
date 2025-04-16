// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// SecretValueClient contains the methods for the SecretValue group.
// Don't use this type directly, use NewSecretValueClient() instead.
type SecretValueClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecretValueClient creates a new instance of SecretValueClient with the specified values.
//   - subscriptionID - The customer subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecretValueClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecretValueClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecretValueClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates a new value of the specified secret resource. The name of the value is typically the version identifier.
// Once created the value cannot be changed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - secretValueResourceName - The name of the secret resource value which is typically the version identifier for the value.
//   - secretValueResourceDescription - Description for creating a value of a secret resource.
//   - options - SecretValueClientCreateOptions contains the optional parameters for the SecretValueClient.Create method.
func (client *SecretValueClient) Create(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription, options *SecretValueClientCreateOptions) (SecretValueClientCreateResponse, error) {
	var err error
	const operationName = "SecretValueClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, secretResourceName, secretValueResourceName, secretValueResourceDescription, options)
	if err != nil {
		return SecretValueClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretValueClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return SecretValueClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SecretValueClient) createCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, secretValueResourceDescription SecretValueResourceDescription, _ *SecretValueClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{secretValueResourceName}", secretValueResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, secretValueResourceDescription); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SecretValueClient) createHandleResponse(resp *http.Response) (SecretValueClientCreateResponse, error) {
	result := SecretValueClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretValueResourceDescription); err != nil {
		return SecretValueClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the secret value resource identified by the name. The name of the resource is typically the version associated
// with that value. Deletion will fail if the specified value is in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - secretValueResourceName - The name of the secret resource value which is typically the version identifier for the value.
//   - options - SecretValueClientDeleteOptions contains the optional parameters for the SecretValueClient.Delete method.
func (client *SecretValueClient) Delete(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, options *SecretValueClientDeleteOptions) (SecretValueClientDeleteResponse, error) {
	var err error
	const operationName = "SecretValueClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, secretResourceName, secretValueResourceName, options)
	if err != nil {
		return SecretValueClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretValueClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SecretValueClientDeleteResponse{}, err
	}
	return SecretValueClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecretValueClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, _ *SecretValueClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{secretValueResourceName}", secretValueResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the information about the specified named secret value resources. The information does not include the actual
// value of the secret.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - secretValueResourceName - The name of the secret resource value which is typically the version identifier for the value.
//   - options - SecretValueClientGetOptions contains the optional parameters for the SecretValueClient.Get method.
func (client *SecretValueClient) Get(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, options *SecretValueClientGetOptions) (SecretValueClientGetResponse, error) {
	var err error
	const operationName = "SecretValueClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, secretResourceName, secretValueResourceName, options)
	if err != nil {
		return SecretValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecretValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecretValueClient) getCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, _ *SecretValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{secretValueResourceName}", secretValueResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecretValueClient) getHandleResponse(resp *http.Response) (SecretValueClientGetResponse, error) {
	result := SecretValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretValueResourceDescription); err != nil {
		return SecretValueClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets information about all secret value resources of the specified secret resource. The information includes
// the names of the secret value resources, but not the actual values.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - options - SecretValueClientListOptions contains the optional parameters for the SecretValueClient.NewListPager method.
func (client *SecretValueClient) NewListPager(resourceGroupName string, secretResourceName string, options *SecretValueClientListOptions) *runtime.Pager[SecretValueClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecretValueClientListResponse]{
		More: func(page SecretValueClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecretValueClientListResponse) (SecretValueClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecretValueClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, secretResourceName, options)
			}, nil)
			if err != nil {
				return SecretValueClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecretValueClient) listCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, _ *SecretValueClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecretValueClient) listHandleResponse(resp *http.Response) (SecretValueClientListResponse, error) {
	result := SecretValueClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretValueResourceDescriptionList); err != nil {
		return SecretValueClientListResponse{}, err
	}
	return result, nil
}

// ListValue - Lists the decrypted value of the specified named value of the secret resource. This is a privileged operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - secretResourceName - The name of the secret resource.
//   - secretValueResourceName - The name of the secret resource value which is typically the version identifier for the value.
//   - options - SecretValueClientListValueOptions contains the optional parameters for the SecretValueClient.ListValue method.
func (client *SecretValueClient) ListValue(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, options *SecretValueClientListValueOptions) (SecretValueClientListValueResponse, error) {
	var err error
	const operationName = "SecretValueClient.ListValue"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listValueCreateRequest(ctx, resourceGroupName, secretResourceName, secretValueResourceName, options)
	if err != nil {
		return SecretValueClientListValueResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecretValueClientListValueResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecretValueClientListValueResponse{}, err
	}
	resp, err := client.listValueHandleResponse(httpResp)
	return resp, err
}

// listValueCreateRequest creates the ListValue request.
func (client *SecretValueClient) listValueCreateRequest(ctx context.Context, resourceGroupName string, secretResourceName string, secretValueResourceName string, _ *SecretValueClientListValueOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}/list_value"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{secretResourceName}", secretResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{secretValueResourceName}", secretValueResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listValueHandleResponse handles the ListValue response.
func (client *SecretValueClient) listValueHandleResponse(resp *http.Response) (SecretValueClientListValueResponse, error) {
	result := SecretValueClientListValueResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretValue); err != nil {
		return SecretValueClientListValueResponse{}, err
	}
	return result, nil
}
