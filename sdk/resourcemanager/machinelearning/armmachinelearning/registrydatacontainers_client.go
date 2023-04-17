//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

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

// RegistryDataContainersClient contains the methods for the RegistryDataContainers group.
// Don't use this type directly, use NewRegistryDataContainersClient() instead.
type RegistryDataContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryDataContainersClient creates a new instance of RegistryDataContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryDataContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryDataContainersClient, error) {
	cl, err := arm.NewClient(moduleName+".RegistryDataContainersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryDataContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - name - Container name.
//   - body - Container entity to create or update.
//   - options - RegistryDataContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryDataContainersClient.BeginCreateOrUpdate
//     method.
func (client *RegistryDataContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, name string, body DataContainer, options *RegistryDataContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryDataContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, name, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryDataContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryDataContainersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *RegistryDataContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, name string, body DataContainer, options *RegistryDataContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, name, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistryDataContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, name string, body DataContainer, options *RegistryDataContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/data/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - name - Container name.
//   - options - RegistryDataContainersClientBeginDeleteOptions contains the optional parameters for the RegistryDataContainersClient.BeginDelete
//     method.
func (client *RegistryDataContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, name string, options *RegistryDataContainersClientBeginDeleteOptions) (*runtime.Poller[RegistryDataContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryDataContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryDataContainersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *RegistryDataContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, name string, options *RegistryDataContainersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistryDataContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, name string, options *RegistryDataContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/data/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - name - Container name.
//   - options - RegistryDataContainersClientGetOptions contains the optional parameters for the RegistryDataContainersClient.Get
//     method.
func (client *RegistryDataContainersClient) Get(ctx context.Context, resourceGroupName string, registryName string, name string, options *RegistryDataContainersClientGetOptions) (RegistryDataContainersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, name, options)
	if err != nil {
		return RegistryDataContainersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryDataContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistryDataContainersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistryDataContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, name string, options *RegistryDataContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/data/{name}"
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
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistryDataContainersClient) getHandleResponse(resp *http.Response) (RegistryDataContainersClientGetResponse, error) {
	result := RegistryDataContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataContainer); err != nil {
		return RegistryDataContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Data containers.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - options - RegistryDataContainersClientListOptions contains the optional parameters for the RegistryDataContainersClient.NewListPager
//     method.
func (client *RegistryDataContainersClient) NewListPager(resourceGroupName string, registryName string, options *RegistryDataContainersClientListOptions) *runtime.Pager[RegistryDataContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryDataContainersClientListResponse]{
		More: func(page RegistryDataContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryDataContainersClientListResponse) (RegistryDataContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistryDataContainersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegistryDataContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistryDataContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistryDataContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistryDataContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/data"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryDataContainersClient) listHandleResponse(resp *http.Response) (RegistryDataContainersClientListResponse, error) {
	result := RegistryDataContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataContainerResourceArmPaginatedResult); err != nil {
		return RegistryDataContainersClientListResponse{}, err
	}
	return result, nil
}
