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

// ExportPipelinesClient contains the methods for the ExportPipelines group.
// Don't use this type directly, use NewExportPipelinesClient() instead.
type ExportPipelinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExportPipelinesClient creates a new instance of ExportPipelinesClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExportPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExportPipelinesClient, error) {
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
	client := &ExportPipelinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates an export pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - exportPipelineCreateParameters - The parameters for creating an export pipeline.
//   - options - ExportPipelinesClientBeginCreateOptions contains the optional parameters for the ExportPipelinesClient.BeginCreate
//     method.
func (client *ExportPipelinesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*runtime.Poller[ExportPipelinesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, exportPipelineName, exportPipelineCreateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ExportPipelinesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ExportPipelinesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates an export pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ExportPipelinesClient) create(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, exportPipelineCreateParameters, options)
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
func (client *ExportPipelinesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
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
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, exportPipelineCreateParameters)
}

// BeginDelete - Deletes an export pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - options - ExportPipelinesClientBeginDeleteOptions contains the optional parameters for the ExportPipelinesClient.BeginDelete
//     method.
func (client *ExportPipelinesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*runtime.Poller[ExportPipelinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, exportPipelineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ExportPipelinesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ExportPipelinesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an export pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
func (client *ExportPipelinesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, options)
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
func (client *ExportPipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
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
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
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

// Get - Gets the properties of the export pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - options - ExportPipelinesClientGetOptions contains the optional parameters for the ExportPipelinesClient.Get method.
func (client *ExportPipelinesClient) Get(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientGetOptions) (ExportPipelinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, options)
	if err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExportPipelinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExportPipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
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
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
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
func (client *ExportPipelinesClient) getHandleResponse(resp *http.Response) (ExportPipelinesClientGetResponse, error) {
	result := ExportPipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportPipeline); err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all export pipelines for the specified container registry.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - ExportPipelinesClientListOptions contains the optional parameters for the ExportPipelinesClient.NewListPager
//     method.
func (client *ExportPipelinesClient) NewListPager(resourceGroupName string, registryName string, options *ExportPipelinesClientListOptions) *runtime.Pager[ExportPipelinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExportPipelinesClientListResponse]{
		More: func(page ExportPipelinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExportPipelinesClientListResponse) (ExportPipelinesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExportPipelinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExportPipelinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExportPipelinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExportPipelinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ExportPipelinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExportPipelinesClient) listHandleResponse(resp *http.Response) (ExportPipelinesClientListResponse, error) {
	result := ExportPipelinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportPipelineListResult); err != nil {
		return ExportPipelinesClientListResponse{}, err
	}
	return result, nil
}
