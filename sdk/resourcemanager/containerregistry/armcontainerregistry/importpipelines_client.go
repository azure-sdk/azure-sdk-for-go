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

// ImportPipelinesClient contains the methods for the ImportPipelines group.
// Don't use this type directly, use NewImportPipelinesClient() instead.
type ImportPipelinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImportPipelinesClient creates a new instance of ImportPipelinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImportPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImportPipelinesClient, error) {
	cl, err := arm.NewClient(moduleName+".ImportPipelinesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImportPipelinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates an import pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - importPipelineName - The name of the import pipeline.
//   - importPipelineCreateParameters - The parameters for creating an import pipeline.
//   - options - ImportPipelinesClientBeginCreateOptions contains the optional parameters for the ImportPipelinesClient.BeginCreate
//     method.
func (client *ImportPipelinesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*runtime.Poller[ImportPipelinesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, importPipelineName, importPipelineCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ImportPipelinesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ImportPipelinesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates an import pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *ImportPipelinesClient) create(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ImportPipelinesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, importPipelineCreateParameters, options)
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
func (client *ImportPipelinesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, importPipelineCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an import pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - importPipelineName - The name of the import pipeline.
//   - options - ImportPipelinesClientBeginDeleteOptions contains the optional parameters for the ImportPipelinesClient.BeginDelete
//     method.
func (client *ImportPipelinesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*runtime.Poller[ImportPipelinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, importPipelineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ImportPipelinesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ImportPipelinesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an import pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *ImportPipelinesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ImportPipelinesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, options)
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
func (client *ImportPipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the import pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - importPipelineName - The name of the import pipeline.
//   - options - ImportPipelinesClientGetOptions contains the optional parameters for the ImportPipelinesClient.Get method.
func (client *ImportPipelinesClient) Get(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientGetOptions) (ImportPipelinesClientGetResponse, error) {
	var err error
	const operationName = "ImportPipelinesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, options)
	if err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportPipelinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImportPipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImportPipelinesClient) getHandleResponse(resp *http.Response) (ImportPipelinesClientGetResponse, error) {
	result := ImportPipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportPipeline); err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all import pipelines for the specified container registry.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - ImportPipelinesClientListOptions contains the optional parameters for the ImportPipelinesClient.NewListPager
//     method.
func (client *ImportPipelinesClient) NewListPager(resourceGroupName string, registryName string, options *ImportPipelinesClientListOptions) *runtime.Pager[ImportPipelinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImportPipelinesClientListResponse]{
		More: func(page ImportPipelinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImportPipelinesClientListResponse) (ImportPipelinesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImportPipelinesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImportPipelinesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ImportPipelinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImportPipelinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ImportPipelinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ImportPipelinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines"
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
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ImportPipelinesClient) listHandleResponse(resp *http.Response) (ImportPipelinesClientListResponse, error) {
	result := ImportPipelinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportPipelineListResult); err != nil {
		return ImportPipelinesClientListResponse{}, err
	}
	return result, nil
}
