//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// ImportCollectorsClient contains the methods for the ImportCollectors group.
// Don't use this type directly, use NewImportCollectorsClient() instead.
type ImportCollectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImportCollectorsClient creates a new instance of ImportCollectorsClient with the specified values.
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImportCollectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImportCollectorsClient, error) {
	cl, err := arm.NewClient(moduleName+".ImportCollectorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImportCollectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create or Update Import collector
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - importCollectorName - Unique name of a Import collector within a project.
//   - options - ImportCollectorsClientCreateOptions contains the optional parameters for the ImportCollectorsClient.Create method.
func (client *ImportCollectorsClient) Create(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientCreateOptions) (ImportCollectorsClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, importCollectorName, options)
	if err != nil {
		return ImportCollectorsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportCollectorsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ImportCollectorsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ImportCollectorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/importcollectors/{importCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if importCollectorName == "" {
		return nil, errors.New("parameter importCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importCollectorName}", url.PathEscape(importCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CollectorBody != nil {
		if err := runtime.MarshalAsJSON(req, *options.CollectorBody); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ImportCollectorsClient) createHandleResponse(resp *http.Response) (ImportCollectorsClientCreateResponse, error) {
	result := ImportCollectorsClientCreateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportCollector); err != nil {
		return ImportCollectorsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Import collector from the project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - importCollectorName - Unique name of a Import collector within a project.
//   - options - ImportCollectorsClientDeleteOptions contains the optional parameters for the ImportCollectorsClient.Delete method.
func (client *ImportCollectorsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientDeleteOptions) (ImportCollectorsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, importCollectorName, options)
	if err != nil {
		return ImportCollectorsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportCollectorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ImportCollectorsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *ImportCollectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/importcollectors/{importCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if importCollectorName == "" {
		return nil, errors.New("parameter importCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importCollectorName}", url.PathEscape(importCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ImportCollectorsClient) deleteHandleResponse(resp *http.Response) (ImportCollectorsClientDeleteResponse, error) {
	result := ImportCollectorsClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// Get - Get a Import collector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - importCollectorName - Unique name of a Import collector within a project.
//   - options - ImportCollectorsClientGetOptions contains the optional parameters for the ImportCollectorsClient.Get method.
func (client *ImportCollectorsClient) Get(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientGetOptions) (ImportCollectorsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, importCollectorName, options)
	if err != nil {
		return ImportCollectorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportCollectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportCollectorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImportCollectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, importCollectorName string, options *ImportCollectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/importcollectors/{importCollectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if importCollectorName == "" {
		return nil, errors.New("parameter importCollectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importCollectorName}", url.PathEscape(importCollectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImportCollectorsClient) getHandleResponse(resp *http.Response) (ImportCollectorsClientGetResponse, error) {
	result := ImportCollectorsClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportCollector); err != nil {
		return ImportCollectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - Get a list of Import collector.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ImportCollectorsClientListByProjectOptions contains the optional parameters for the ImportCollectorsClient.NewListByProjectPager
//     method.
func (client *ImportCollectorsClient) NewListByProjectPager(resourceGroupName string, projectName string, options *ImportCollectorsClientListByProjectOptions) *runtime.Pager[ImportCollectorsClientListByProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImportCollectorsClientListByProjectResponse]{
		More: func(page ImportCollectorsClientListByProjectResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ImportCollectorsClientListByProjectResponse) (ImportCollectorsClientListByProjectResponse, error) {
			req, err := client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			if err != nil {
				return ImportCollectorsClientListByProjectResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ImportCollectorsClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImportCollectorsClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *ImportCollectorsClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ImportCollectorsClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/importcollectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *ImportCollectorsClient) listByProjectHandleResponse(resp *http.Response) (ImportCollectorsClientListByProjectResponse, error) {
	result := ImportCollectorsClientListByProjectResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportCollectorList); err != nil {
		return ImportCollectorsClientListByProjectResponse{}, err
	}
	return result, nil
}
