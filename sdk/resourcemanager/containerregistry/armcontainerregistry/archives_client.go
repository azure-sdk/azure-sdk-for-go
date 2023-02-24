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

// ArchivesClient contains the methods for the Archives group.
// Don't use this type directly, use NewArchivesClient() instead.
type ArchivesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewArchivesClient creates a new instance of ArchivesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArchivesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArchivesClient, error) {
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
	client := &ArchivesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package repository.
//   - archiveName - The name of the archive resource.
//   - archiveCreateParameters - The parameters for creating a archive.
//   - options - ArchivesClientBeginCreateOptions contains the optional parameters for the ArchivesClient.BeginCreate method.
func (client *ArchivesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveCreateParameters Archive, options *ArchivesClientBeginCreateOptions) (*runtime.Poller[ArchivesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, packageType, archiveName, archiveCreateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ArchivesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ArchivesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *ArchivesClient) create(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveCreateParameters Archive, options *ArchivesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, archiveCreateParameters, options)
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
func (client *ArchivesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveCreateParameters Archive, options *ArchivesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, archiveCreateParameters)
}

// BeginDelete - Deletes a archive from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package repository.
//   - archiveName - The name of the archive resource.
//   - options - ArchivesClientBeginDeleteOptions contains the optional parameters for the ArchivesClient.BeginDelete method.
func (client *ArchivesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchivesClientBeginDeleteOptions) (*runtime.Poller[ArchivesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, packageType, archiveName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ArchivesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ArchivesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a archive from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *ArchivesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchivesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, options)
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
func (client *ArchivesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchivesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the archive.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package repository.
//   - archiveName - The name of the archive resource.
//   - options - ArchivesClientGetOptions contains the optional parameters for the ArchivesClient.Get method.
func (client *ArchivesClient) Get(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchivesClientGetOptions) (ArchivesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, options)
	if err != nil {
		return ArchivesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ArchivesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArchivesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArchivesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, options *ArchivesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *ArchivesClient) getHandleResponse(resp *http.Response) (ArchivesClientGetResponse, error) {
	result := ArchivesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Archive); err != nil {
		return ArchivesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all archives for the specified container registry and repository type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package repository.
//   - options - ArchivesClientListOptions contains the optional parameters for the ArchivesClient.NewListPager method.
func (client *ArchivesClient) NewListPager(resourceGroupName string, registryName string, packageType string, options *ArchivesClientListOptions) *runtime.Pager[ArchivesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArchivesClientListResponse]{
		More: func(page ArchivesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArchivesClientListResponse) (ArchivesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, packageType, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArchivesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ArchivesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArchivesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ArchivesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, options *ArchivesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *ArchivesClient) listHandleResponse(resp *http.Response) (ArchivesClientListResponse, error) {
	result := ArchivesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArchiveListResult); err != nil {
		return ArchivesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - packageType - The type of the package repository.
//   - archiveName - The name of the archive resource.
//   - archiveUpdateParameters - The parameters for updating a archive.
//   - options - ArchivesClientBeginUpdateOptions contains the optional parameters for the ArchivesClient.BeginUpdate method.
func (client *ArchivesClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveUpdateParameters ArchiveUpdateParameters, options *ArchivesClientBeginUpdateOptions) (*runtime.Poller[ArchivesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, packageType, archiveName, archiveUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ArchivesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ArchivesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a archive for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *ArchivesClient) update(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveUpdateParameters ArchiveUpdateParameters, options *ArchivesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, packageType, archiveName, archiveUpdateParameters, options)
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
func (client *ArchivesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, packageType string, archiveName string, archiveUpdateParameters ArchiveUpdateParameters, options *ArchivesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/packages/{packageType}/archives/{archiveName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, archiveUpdateParameters)
}
