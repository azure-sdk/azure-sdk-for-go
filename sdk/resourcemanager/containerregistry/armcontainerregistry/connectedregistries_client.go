//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectedRegistriesClient contains the methods for the ConnectedRegistries group.
// Don't use this type directly, use NewConnectedRegistriesClient() instead.
type ConnectedRegistriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectedRegistriesClient creates a new instance of ConnectedRegistriesClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedRegistriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedRegistriesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedRegistriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
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
		return runtime.NewPoller[ConnectedRegistriesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ConnectedRegistriesClient) create(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ConnectedRegistriesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, connectedRegistryCreateParameters)
}

// BeginDeactivate - Deactivates the connected registry instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
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
		return runtime.NewPoller[ConnectedRegistriesClientDeactivateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientDeactivateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Deactivate - Deactivates the connected registry instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ConnectedRegistriesClient) deactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeactivateOptions) (*http.Response, error) {
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deactivateCreateRequest creates the Deactivate request.
func (client *ConnectedRegistriesClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeactivateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}/deactivate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes a connected registry from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
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
		return runtime.NewPoller[ConnectedRegistriesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a connected registry from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ConnectedRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the connected registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - connectedRegistryName - The name of the connected registry.
//   - options - ConnectedRegistriesClientGetOptions contains the optional parameters for the ConnectedRegistriesClient.Get method.
func (client *ConnectedRegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientGetOptions) (ConnectedRegistriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedRegistriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedRegistriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
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
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - ConnectedRegistriesClientListOptions contains the optional parameters for the ConnectedRegistriesClient.NewListPager
//     method.
func (client *ConnectedRegistriesClient) NewListPager(resourceGroupName string, registryName string, options *ConnectedRegistriesClientListOptions) *runtime.Pager[ConnectedRegistriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedRegistriesClientListResponse]{
		More: func(page ConnectedRegistriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedRegistriesClientListResponse) (ConnectedRegistriesClientListResponse, error) {
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
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedRegistriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedRegistriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ConnectedRegistriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
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
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
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
		return runtime.NewPoller[ConnectedRegistriesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectedRegistriesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a connected registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ConnectedRegistriesClient) update(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConnectedRegistriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, connectedRegistryUpdateParameters)
}
