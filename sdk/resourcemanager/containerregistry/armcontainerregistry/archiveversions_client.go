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

// ArchiveVersionsClient contains the methods for the ArchiveVersions group.
// Don't use this type directly, use NewArchiveVersionsClient() instead.
type ArchiveVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArchiveVersionsClient creates a new instance of ArchiveVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArchiveVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArchiveVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ArchiveVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArchiveVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package resource.
//   - archiveName - The name of the archive resource.
//   - archiveVersionName - The name of the archive version resource.
//   - options - ArchiveVersionsClientBeginCreateOptions contains the optional parameters for the ArchiveVersionsClient.BeginCreate
//     method.
func (client *ArchiveVersionsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginCreateOptions) (*runtime.Poller[ArchiveVersionsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, packageType, archiveName, archiveVersionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArchiveVersionsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ArchiveVersionsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *ArchiveVersionsClient) create(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ArchiveVersionsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, archiveVersionName, options)
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
func (client *ArchiveVersionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}/versions/{archiveVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if packageType == "" {
		return nil, errors.New("parameter packageType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageType}", url.PathEscape(packageType))
	if archiveName == "" {
		return nil, errors.New("parameter archiveName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveName}", url.PathEscape(archiveName))
	if archiveVersionName == "" {
		return nil, errors.New("parameter archiveVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveVersionName}", url.PathEscape(archiveVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes a archive version from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package resource.
//   - archiveName - The name of the archive resource.
//   - archiveVersionName - The name of the archive version resource.
//   - options - ArchiveVersionsClientBeginDeleteOptions contains the optional parameters for the ArchiveVersionsClient.BeginDelete
//     method.
func (client *ArchiveVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginDeleteOptions) (*runtime.Poller[ArchiveVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, packageType, archiveName, archiveVersionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArchiveVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ArchiveVersionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a archive version from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *ArchiveVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ArchiveVersionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, archiveVersionName, options)
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
func (client *ArchiveVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}/versions/{archiveVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if packageType == "" {
		return nil, errors.New("parameter packageType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageType}", url.PathEscape(packageType))
	if archiveName == "" {
		return nil, errors.New("parameter archiveName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveName}", url.PathEscape(archiveName))
	if archiveVersionName == "" {
		return nil, errors.New("parameter archiveVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveVersionName}", url.PathEscape(archiveVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the archive version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package resource.
//   - archiveName - The name of the archive resource.
//   - archiveVersionName - The name of the archive version resource.
//   - options - ArchiveVersionsClientGetOptions contains the optional parameters for the ArchiveVersionsClient.Get method.
func (client *ArchiveVersionsClient) Get(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientGetOptions) (ArchiveVersionsClientGetResponse, error) {
	var err error
	const operationName = "ArchiveVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, archiveVersionName, options)
	if err != nil {
		return ArchiveVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArchiveVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ArchiveVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ArchiveVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveVersionName string, options *ArchiveVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}/versions/{archiveVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if packageType == "" {
		return nil, errors.New("parameter packageType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageType}", url.PathEscape(packageType))
	if archiveName == "" {
		return nil, errors.New("parameter archiveName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveName}", url.PathEscape(archiveName))
	if archiveVersionName == "" {
		return nil, errors.New("parameter archiveVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveVersionName}", url.PathEscape(archiveVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArchiveVersionsClient) getHandleResponse(resp *http.Response) (ArchiveVersionsClientGetResponse, error) {
	result := ArchiveVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArchiveVersion); err != nil {
		return ArchiveVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all archive versions for the specified container registry, repository type and archive name.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package resource.
//   - archiveName - The name of the archive resource.
//   - options - ArchiveVersionsClientListOptions contains the optional parameters for the ArchiveVersionsClient.NewListPager
//     method.
func (client *ArchiveVersionsClient) NewListPager(resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchiveVersionsClientListOptions) *runtime.Pager[ArchiveVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArchiveVersionsClientListResponse]{
		More: func(page ArchiveVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArchiveVersionsClientListResponse) (ArchiveVersionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ArchiveVersionsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArchiveVersionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ArchiveVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArchiveVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ArchiveVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchiveVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if packageType == "" {
		return nil, errors.New("parameter packageType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageType}", url.PathEscape(packageType))
	if archiveName == "" {
		return nil, errors.New("parameter archiveName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{archiveName}", url.PathEscape(archiveName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArchiveVersionsClient) listHandleResponse(resp *http.Response) (ArchiveVersionsClientListResponse, error) {
	result := ArchiveVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArchiveVersionListResult); err != nil {
		return ArchiveVersionsClientListResponse{}, err
	}
	return result, nil
}
