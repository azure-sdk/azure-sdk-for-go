//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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

// ConnectedRegistriesClient contains the methods for the ConnectedRegistries group.
// Don't use this type directly, use NewConnectedRegistriesClient() instead.
type ConnectedRegistriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedRegistriesClient creates a new instance of ConnectedRegistriesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedRegistriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedRegistriesClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectedRegistriesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedRegistriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - connectedRegistryCreateParameters - The parameters for creating a connectedRegistry.
//   - options - ConnectedRegistriesClientBeginCreateOptions contains the optional parameters for the ConnectedRegistriesClient.BeginCreate
//     method.
func (client *ConnectedRegistriesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesClientBeginCreateOptions) (*runtime.Poller[ConnectedRegistriesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedRegistriesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConnectedRegistriesClient) create(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedRegistriesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
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

// createCreateRequest creates the Create request.
func (client *ConnectedRegistriesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectedRegistryCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeactivate - Deactivates the connected registry instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - options - ConnectedRegistriesClientBeginDeactivateOptions contains the optional parameters for the ConnectedRegistriesClient.BeginDeactivate
//     method.
func (client *ConnectedRegistriesClient) BeginDeactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeactivateOptions) (*runtime.Poller[ConnectedRegistriesClientDeactivateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deactivate(ctx, resourceGroupName, registryName, connectedRegistryName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedRegistriesClientDeactivateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientDeactivateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Deactivate - Deactivates the connected registry instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConnectedRegistriesClient) deactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeactivateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedRegistriesClient.BeginDeactivate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
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

// deactivateCreateRequest creates the Deactivate request.
func (client *ConnectedRegistriesClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeactivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}/deactivate"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes a connected registry from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - options - ConnectedRegistriesClientBeginDeleteOptions contains the optional parameters for the ConnectedRegistriesClient.BeginDelete
//     method.
func (client *ConnectedRegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeleteOptions) (*runtime.Poller[ConnectedRegistriesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, connectedRegistryName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedRegistriesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a connected registry from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConnectedRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedRegistriesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the connected registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - options - ConnectedRegistriesClientGetOptions contains the optional parameters for the ConnectedRegistriesClient.Get method.
func (client *ConnectedRegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientGetOptions) (ConnectedRegistriesClientGetResponse, error) {
	var err error
	const operationName = "ConnectedRegistriesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistriesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedRegistriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedRegistriesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectedRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedRegistriesClient) getHandleResponse(resp *http.Response) (ConnectedRegistriesClientGetResponse, error) {
	result := ConnectedRegistriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedRegistry); err != nil {
		return ConnectedRegistriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all connected registries for the specified container registry.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - ConnectedRegistriesClientListOptions contains the optional parameters for the ConnectedRegistriesClient.NewListPager
//     method.
func (client *ConnectedRegistriesClient) NewListPager(resourceGroupName string, registryName string, options *ConnectedRegistriesClientListOptions) *runtime.Pager[ConnectedRegistriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedRegistriesClientListResponse]{
		More: func(page ConnectedRegistriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedRegistriesClientListResponse) (ConnectedRegistriesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedRegistriesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedRegistriesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConnectedRegistriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedRegistriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ConnectedRegistriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedRegistriesClient) listHandleResponse(resp *http.Response) (ConnectedRegistriesClientListResponse, error) {
	result := ConnectedRegistriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedRegistryListResult); err != nil {
		return ConnectedRegistriesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a connected registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - connectedRegistryUpdateParameters - The parameters for updating a connectedRegistry.
//   - options - ConnectedRegistriesClientBeginUpdateOptions contains the optional parameters for the ConnectedRegistriesClient.BeginUpdate
//     method.
func (client *ConnectedRegistriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesClientBeginUpdateOptions) (*runtime.Poller[ConnectedRegistriesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedRegistriesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a connected registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ConnectedRegistriesClient) update(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedRegistriesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ConnectedRegistriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectedRegistryUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
