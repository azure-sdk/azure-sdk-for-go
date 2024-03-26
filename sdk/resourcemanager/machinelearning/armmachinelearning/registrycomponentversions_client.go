//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// RegistryComponentVersionsClient contains the methods for the RegistryComponentVersions group.
// Don't use this type directly, use NewRegistryComponentVersionsClient() instead.
type RegistryComponentVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegistryComponentVersionsClient creates a new instance of RegistryComponentVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistryComponentVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistryComponentVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistryComponentVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - componentName - Container name.
//   - version - Version identifier.
//   - body - Version entity to create or update.
//   - options - RegistryComponentVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistryComponentVersionsClient.BeginCreateOrUpdate
//     method.
func (client *RegistryComponentVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, body ComponentVersion, options *RegistryComponentVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistryComponentVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, registryName, componentName, version, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryComponentVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryComponentVersionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *RegistryComponentVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, body ComponentVersion, options *RegistryComponentVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryComponentVersionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, registryName, componentName, version, body, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegistryComponentVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, body ComponentVersion, options *RegistryComponentVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/components/{componentName}/versions/{version}"
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
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - componentName - Container name.
//   - version - Version identifier.
//   - options - RegistryComponentVersionsClientBeginDeleteOptions contains the optional parameters for the RegistryComponentVersionsClient.BeginDelete
//     method.
func (client *RegistryComponentVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *RegistryComponentVersionsClientBeginDeleteOptions) (*runtime.Poller[RegistryComponentVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, componentName, version, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistryComponentVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistryComponentVersionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *RegistryComponentVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *RegistryComponentVersionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistryComponentVersionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, componentName, version, options)
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
func (client *RegistryComponentVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *RegistryComponentVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/components/{componentName}/versions/{version}"
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
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - componentName - Container name.
//   - version - Version identifier.
//   - options - RegistryComponentVersionsClientGetOptions contains the optional parameters for the RegistryComponentVersionsClient.Get
//     method.
func (client *RegistryComponentVersionsClient) Get(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *RegistryComponentVersionsClientGetOptions) (RegistryComponentVersionsClientGetResponse, error) {
	var err error
	const operationName = "RegistryComponentVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, componentName, version, options)
	if err != nil {
		return RegistryComponentVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistryComponentVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegistryComponentVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegistryComponentVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *RegistryComponentVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/components/{componentName}/versions/{version}"
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
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
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
func (client *RegistryComponentVersionsClient) getHandleResponse(resp *http.Response) (RegistryComponentVersionsClientGetResponse, error) {
	result := RegistryComponentVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentVersion); err != nil {
		return RegistryComponentVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List versions.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - Name of Azure Machine Learning registry. This is case-insensitive
//   - componentName - Container name.
//   - options - RegistryComponentVersionsClientListOptions contains the optional parameters for the RegistryComponentVersionsClient.NewListPager
//     method.
func (client *RegistryComponentVersionsClient) NewListPager(resourceGroupName string, registryName string, componentName string, options *RegistryComponentVersionsClientListOptions) *runtime.Pager[RegistryComponentVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistryComponentVersionsClientListResponse]{
		More: func(page RegistryComponentVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistryComponentVersionsClientListResponse) (RegistryComponentVersionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegistryComponentVersionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, registryName, componentName, options)
			}, nil)
			if err != nil {
				return RegistryComponentVersionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RegistryComponentVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, componentName string, options *RegistryComponentVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/registries/{registryName}/components/{componentName}/versions"
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
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-04-01-preview")
	if options != nil && options.Stage != nil {
		reqQP.Set("stage", *options.Stage)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistryComponentVersionsClient) listHandleResponse(resp *http.Response) (RegistryComponentVersionsClientListResponse, error) {
	result := RegistryComponentVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentVersionResourceArmPaginatedResult); err != nil {
		return RegistryComponentVersionsClientListResponse{}, err
	}
	return result, nil
}
