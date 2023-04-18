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
	"strconv"
	"strings"
)

// RegistryModelVersionsClient contains the methods for the RegistryModelVersions group.
// Don't use this type directly, use NewRegistryModelVersionsClient() instead.
type RegistryModelVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryModelVersionsClient creates a new instance of RegistryModelVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryModelVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryModelVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".RegistryModelVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryModelVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrGetStartPendingUpload - Generate a storage location and credential for the client to upload a model asset to.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - modelName - Model name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - body - Pending upload request object
//   - options - RegistryModelVersionsClientCreateOrGetStartPendingUploadOptions contains the optional parameters for the RegistryModelVersionsClient.CreateOrGetStartPendingUpload
//     method.
func (client *RegistryModelVersionsClient) CreateOrGetStartPendingUpload(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body PendingUploadRequestDto, options *RegistryModelVersionsClientCreateOrGetStartPendingUploadOptions) (RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse, error) {
	req, err := client.createOrGetStartPendingUploadCreateRequest(ctx, resourceGroupName, registryName, modelName, version, body, options)
	if err != nil {
		return RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrGetStartPendingUploadHandleResponse(resp)
}

// createOrGetStartPendingUploadCreateRequest creates the CreateOrGetStartPendingUpload request.
func (client *RegistryModelVersionsClient) createOrGetStartPendingUploadCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body PendingUploadRequestDto, options *RegistryModelVersionsClientCreateOrGetStartPendingUploadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}/versions/{version}/startPendingUpload"
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
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrGetStartPendingUploadHandleResponse handles the CreateOrGetStartPendingUpload response.
func (client *RegistryModelVersionsClient) createOrGetStartPendingUploadHandleResponse(resp *http.Response) (RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse, error) {
	result := RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PendingUploadResponseDto); err != nil {
		return RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - modelName - Container name.
//   - version - Version identifier.
//   - body - Version entity to create or update.
//   - options - RegistryModelVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryModelVersionsClient.BeginCreateOrUpdate
//     method.
func (client *RegistryModelVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body ModelVersion, options *RegistryModelVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryModelVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, modelName, version, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryModelVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryModelVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *RegistryModelVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body ModelVersion, options *RegistryModelVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, modelName, version, body, options)
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
func (client *RegistryModelVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body ModelVersion, options *RegistryModelVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}/versions/{version}"
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
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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

// BeginDelete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - modelName - Container name.
//   - version - Version identifier.
//   - options - RegistryModelVersionsClientBeginDeleteOptions contains the optional parameters for the RegistryModelVersionsClient.BeginDelete
//     method.
func (client *RegistryModelVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *RegistryModelVersionsClientBeginDeleteOptions) (*runtime.Poller[RegistryModelVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, modelName, version, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryModelVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RegistryModelVersionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *RegistryModelVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *RegistryModelVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, modelName, version, options)
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
func (client *RegistryModelVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *RegistryModelVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}/versions/{version}"
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
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - modelName - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - RegistryModelVersionsClientGetOptions contains the optional parameters for the RegistryModelVersionsClient.Get
//     method.
func (client *RegistryModelVersionsClient) Get(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *RegistryModelVersionsClientGetOptions) (RegistryModelVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, modelName, version, options)
	if err != nil {
		return RegistryModelVersionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryModelVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegistryModelVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegistryModelVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *RegistryModelVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}/versions/{version}"
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
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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
func (client *RegistryModelVersionsClient) getHandleResponse(resp *http.Response) (RegistryModelVersionsClientGetResponse, error) {
	result := RegistryModelVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelVersion); err != nil {
		return RegistryModelVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List versions.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - modelName - Container name. This is case-sensitive.
//   - options - RegistryModelVersionsClientListOptions contains the optional parameters for the RegistryModelVersionsClient.NewListPager
//     method.
func (client *RegistryModelVersionsClient) NewListPager(resourceGroupName string, registryName string, modelName string, options *RegistryModelVersionsClientListOptions) *runtime.Pager[RegistryModelVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryModelVersionsClientListResponse]{
		More: func(page RegistryModelVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryModelVersionsClientListResponse) (RegistryModelVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, modelName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegistryModelVersionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RegistryModelVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegistryModelVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegistryModelVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, modelName string, options *RegistryModelVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/models/{modelName}/versions"
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Version != nil {
		reqQP.Set("version", *options.Version)
	}
	if options != nil && options.Description != nil {
		reqQP.Set("description", *options.Description)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", *options.Tags)
	}
	if options != nil && options.Properties != nil {
		reqQP.Set("properties", *options.Properties)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryModelVersionsClient) listHandleResponse(resp *http.Response) (RegistryModelVersionsClientListResponse, error) {
	result := RegistryModelVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModelVersionResourceArmPaginatedResult); err != nil {
		return RegistryModelVersionsClientListResponse{}, err
	}
	return result, nil
}
