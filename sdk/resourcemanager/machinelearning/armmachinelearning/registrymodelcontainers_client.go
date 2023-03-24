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

// RegistryModelContainersClient contains the methods for the RegistryModelContainers group.
// Don't use this type directly, use NewRegistryModelContainersClient() instead.
type RegistryModelContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryModelContainersClient creates a new instance of RegistryModelContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryModelContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryModelContainersClient, error) {
	cl, err := arm.NewClient(moduleName+".RegistryModelContainersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryModelContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update model container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry.
//   - modelName - Container name.
//   - body - Container entity to create or update.
//   - options - RegistryModelContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryModelContainersClient.BeginCreateOrUpdate
//     method.
func (client *RegistryModelContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, modelName string, body ModelContainer, options *RegistryModelContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryModelContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, modelName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryModelContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryModelContainersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update model container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *RegistryModelContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, modelName string, body ModelContainer, options *RegistryModelContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, modelName, body, options)
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
func (client *RegistryModelContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, body ModelContainer, options *RegistryModelContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}"
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
	if modelName == "" {
		return nil, errors.New("parameter modelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", url.PathEscape(modelName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry.
//   - modelName - Container name.
//   - options - RegistryModelContainersClientBeginDeleteOptions contains the optional parameters for the RegistryModelContainersClient.BeginDelete
//     method.
func (client *RegistryModelContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelContainersClientBeginDeleteOptions) (*runtime.Poller[RegistryModelContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, modelName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryModelContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryModelContainersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *RegistryModelContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelContainersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, modelName, options)
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
func (client *RegistryModelContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}"
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
	if modelName == "" {
		return nil, errors.New("parameter modelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", url.PathEscape(modelName))
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

// Get - Get container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry.
//   - modelName - Container name. This is case-sensitive.
//   - options - RegistryModelContainersClientGetOptions contains the optional parameters for the RegistryModelContainersClient.Get
//     method.
func (client *RegistryModelContainersClient) Get(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelContainersClientGetOptions) (RegistryModelContainersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, modelName, options)
	if err != nil {
		return RegistryModelContainersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryModelContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistryModelContainersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistryModelContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}"
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
	if modelName == "" {
		return nil, errors.New("parameter modelName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modelName}", url.PathEscape(modelName))
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
func (client *RegistryModelContainersClient) getHandleResponse(resp *http.Response) (RegistryModelContainersClientGetResponse, error) {
	result := RegistryModelContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelContainer); err != nil {
		return RegistryModelContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List model containers.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry.
//   - options - RegistryModelContainersClientListOptions contains the optional parameters for the RegistryModelContainersClient.NewListPager
//     method.
func (client *RegistryModelContainersClient) NewListPager(resourceGroupName string, registryName string, options *RegistryModelContainersClientListOptions) *runtime.Pager[RegistryModelContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryModelContainersClientListResponse]{
		More: func(page RegistryModelContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryModelContainersClientListResponse) (RegistryModelContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistryModelContainersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegistryModelContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistryModelContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistryModelContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RegistryModelContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models"
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
	reqQP.Set("api-version", "2023-04-01")
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
func (client *RegistryModelContainersClient) listHandleResponse(resp *http.Response) (RegistryModelContainersClientListResponse, error) {
	result := RegistryModelContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelContainerResourceArmPaginatedResult); err != nil {
		return RegistryModelContainersClientListResponse{}, err
	}
	return result, nil
}
