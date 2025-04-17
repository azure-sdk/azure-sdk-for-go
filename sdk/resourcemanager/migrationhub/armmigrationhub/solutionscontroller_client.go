// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationhub

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

// SolutionsControllerClient contains the methods for the SolutionsController group.
// Don't use this type directly, use NewSolutionsControllerClient() instead.
type SolutionsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSolutionsControllerClient creates a new instance of SolutionsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSolutionsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SolutionsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SolutionsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CleanupData - Cleanup the solution data in the migrate project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - options - SolutionsControllerClientCleanupDataOptions contains the optional parameters for the SolutionsControllerClient.CleanupData
//     method.
func (client *SolutionsControllerClient) CleanupData(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, options *SolutionsControllerClientCleanupDataOptions) (SolutionsControllerClientCleanupDataResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.CleanupData"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cleanupDataCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, options)
	if err != nil {
		return SolutionsControllerClientCleanupDataResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientCleanupDataResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientCleanupDataResponse{}, err
	}
	resp, err := client.cleanupDataHandleResponse(httpResp)
	return resp, err
}

// cleanupDataCreateRequest creates the CleanupData request.
func (client *SolutionsControllerClient) cleanupDataCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, _ *SolutionsControllerClientCleanupDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}/cleanupData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cleanupDataHandleResponse handles the CleanupData response.
func (client *SolutionsControllerClient) cleanupDataHandleResponse(resp *http.Response) (SolutionsControllerClientCleanupDataResponse, error) {
	result := SolutionsControllerClientCleanupDataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return SolutionsControllerClientCleanupDataResponse{}, err
	}
	return result, nil
}

// Create - Creates a solution in the migrate project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - solutionInput - The input for the solution.
//   - options - SolutionsControllerClientCreateOptions contains the optional parameters for the SolutionsControllerClient.Create
//     method.
func (client *SolutionsControllerClient) Create(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput Solution, options *SolutionsControllerClientCreateOptions) (SolutionsControllerClientCreateResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, solutionInput, options)
	if err != nil {
		return SolutionsControllerClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SolutionsControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput Solution, _ *SolutionsControllerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, solutionInput); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SolutionsControllerClient) createHandleResponse(resp *http.Response) (SolutionsControllerClientCreateResponse, error) {
	result := SolutionsControllerClientCreateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return SolutionsControllerClientCreateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsControllerClientCreateResponse{}, err
	}
	return result, nil
}

// DeleteSolution - Delete the solution. Deleting non-existent project is a no-operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - options - SolutionsControllerClientDeleteSolutionOptions contains the optional parameters for the SolutionsControllerClient.DeleteSolution
//     method.
func (client *SolutionsControllerClient) DeleteSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, options *SolutionsControllerClientDeleteSolutionOptions) (SolutionsControllerClientDeleteSolutionResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.DeleteSolution"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteSolutionCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, options)
	if err != nil {
		return SolutionsControllerClientDeleteSolutionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientDeleteSolutionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientDeleteSolutionResponse{}, err
	}
	resp, err := client.deleteSolutionHandleResponse(httpResp)
	return resp, err
}

// deleteSolutionCreateRequest creates the DeleteSolution request.
func (client *SolutionsControllerClient) deleteSolutionCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, _ *SolutionsControllerClientDeleteSolutionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteSolutionHandleResponse handles the DeleteSolution response.
func (client *SolutionsControllerClient) deleteSolutionHandleResponse(resp *http.Response) (SolutionsControllerClientDeleteSolutionResponse, error) {
	result := SolutionsControllerClientDeleteSolutionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return SolutionsControllerClientDeleteSolutionResponse{}, err
	}
	return result, nil
}

// GetConfig - Gets the config for the solution in the migrate project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - options - SolutionsControllerClientGetConfigOptions contains the optional parameters for the SolutionsControllerClient.GetConfig
//     method.
func (client *SolutionsControllerClient) GetConfig(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, options *SolutionsControllerClientGetConfigOptions) (SolutionsControllerClientGetConfigResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.GetConfig"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getConfigCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, options)
	if err != nil {
		return SolutionsControllerClientGetConfigResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientGetConfigResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientGetConfigResponse{}, err
	}
	resp, err := client.getConfigHandleResponse(httpResp)
	return resp, err
}

// getConfigCreateRequest creates the GetConfig request.
func (client *SolutionsControllerClient) getConfigCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, _ *SolutionsControllerClientGetConfigOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}/getConfig"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getConfigHandleResponse handles the GetConfig response.
func (client *SolutionsControllerClient) getConfigHandleResponse(resp *http.Response) (SolutionsControllerClientGetConfigResponse, error) {
	result := SolutionsControllerClientGetConfigResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionConfig); err != nil {
		return SolutionsControllerClientGetConfigResponse{}, err
	}
	return result, nil
}

// GetSolution - Gets a solution in the migrate project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - options - SolutionsControllerClientGetSolutionOptions contains the optional parameters for the SolutionsControllerClient.GetSolution
//     method.
func (client *SolutionsControllerClient) GetSolution(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, options *SolutionsControllerClientGetSolutionOptions) (SolutionsControllerClientGetSolutionResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.GetSolution"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSolutionCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, options)
	if err != nil {
		return SolutionsControllerClientGetSolutionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientGetSolutionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientGetSolutionResponse{}, err
	}
	resp, err := client.getSolutionHandleResponse(httpResp)
	return resp, err
}

// getSolutionCreateRequest creates the GetSolution request.
func (client *SolutionsControllerClient) getSolutionCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, _ *SolutionsControllerClientGetSolutionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSolutionHandleResponse handles the GetSolution response.
func (client *SolutionsControllerClient) getSolutionHandleResponse(resp *http.Response) (SolutionsControllerClientGetSolutionResponse, error) {
	result := SolutionsControllerClientGetSolutionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsControllerClientGetSolutionResponse{}, err
	}
	return result, nil
}

// NewListSolutionsPager - Gets the list of solutions in the migrate project.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - options - SolutionsControllerClientListSolutionsOptions contains the optional parameters for the SolutionsControllerClient.NewListSolutionsPager
//     method.
func (client *SolutionsControllerClient) NewListSolutionsPager(resourceGroupName string, migrateProjectName string, options *SolutionsControllerClientListSolutionsOptions) *runtime.Pager[SolutionsControllerClientListSolutionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SolutionsControllerClientListSolutionsResponse]{
		More: func(page SolutionsControllerClientListSolutionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SolutionsControllerClientListSolutionsResponse) (SolutionsControllerClientListSolutionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SolutionsControllerClient.NewListSolutionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSolutionsCreateRequest(ctx, resourceGroupName, migrateProjectName, options)
			}, nil)
			if err != nil {
				return SolutionsControllerClientListSolutionsResponse{}, err
			}
			return client.listSolutionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSolutionsCreateRequest creates the ListSolutions request.
func (client *SolutionsControllerClient) listSolutionsCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, _ *SolutionsControllerClientListSolutionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSolutionsHandleResponse handles the ListSolutions response.
func (client *SolutionsControllerClient) listSolutionsHandleResponse(resp *http.Response) (SolutionsControllerClientListSolutionsResponse, error) {
	result := SolutionsControllerClientListSolutionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SolutionsCollection); err != nil {
		return SolutionsControllerClientListSolutionsResponse{}, err
	}
	return result, nil
}

// Update - Update a solution with specified name. Supports partial updates, for example only tags can be provided.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - migrateProjectName - Name of the Azure Migrate project.
//   - solutionName - Unique name of a migration solution within a migrate project.
//   - solutionInput - The input for the solution.
//   - options - SolutionsControllerClientUpdateOptions contains the optional parameters for the SolutionsControllerClient.Update
//     method.
func (client *SolutionsControllerClient) Update(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput Solution, options *SolutionsControllerClientUpdateOptions) (SolutionsControllerClientUpdateResponse, error) {
	var err error
	const operationName = "SolutionsControllerClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, migrateProjectName, solutionName, solutionInput, options)
	if err != nil {
		return SolutionsControllerClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SolutionsControllerClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SolutionsControllerClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SolutionsControllerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, migrateProjectName string, solutionName string, solutionInput Solution, _ *SolutionsControllerClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if migrateProjectName == "" {
		return nil, errors.New("parameter migrateProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrateProjectName}", url.PathEscape(migrateProjectName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, solutionInput); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SolutionsControllerClient) updateHandleResponse(resp *http.Response) (SolutionsControllerClientUpdateResponse, error) {
	result := SolutionsControllerClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Solution); err != nil {
		return SolutionsControllerClientUpdateResponse{}, err
	}
	return result, nil
}
