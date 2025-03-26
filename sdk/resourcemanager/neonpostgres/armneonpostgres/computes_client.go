// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armneonpostgres

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

// ComputesClient contains the methods for the Computes group.
// Don't use this type directly, use NewComputesClient() instead.
type ComputesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewComputesClient creates a new instance of ComputesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewComputesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComputesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ComputesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - computeName - The name of the Compute
//   - resource - Resource create parameters.
//   - options - ComputesClientBeginCreateOrUpdateOptions contains the optional parameters for the ComputesClient.BeginCreateOrUpdate
//     method.
func (client *ComputesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, resource Compute, options *ComputesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ComputesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ComputesClient) createOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, resource Compute, options *ComputesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, resource, options)
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
func (client *ComputesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, resource Compute, _ *ComputesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if branchName == "" {
		return nil, errors.New("parameter branchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{branchName}", url.PathEscape(branchName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - computeName - The name of the Compute
//   - options - ComputesClientDeleteOptions contains the optional parameters for the ComputesClient.Delete method.
func (client *ComputesClient) Delete(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, options *ComputesClientDeleteOptions) (ComputesClientDeleteResponse, error) {
	var err error
	const operationName = "ComputesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, options)
	if err != nil {
		return ComputesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ComputesClientDeleteResponse{}, err
	}
	return ComputesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ComputesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, _ *ComputesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if branchName == "" {
		return nil, errors.New("parameter branchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{branchName}", url.PathEscape(branchName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - computeName - The name of the Compute
//   - options - ComputesClientGetOptions contains the optional parameters for the ComputesClient.Get method.
func (client *ComputesClient) Get(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, options *ComputesClientGetOptions) (ComputesClientGetResponse, error) {
	var err error
	const operationName = "ComputesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, options)
	if err != nil {
		return ComputesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ComputesClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, _ *ComputesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if branchName == "" {
		return nil, errors.New("parameter branchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{branchName}", url.PathEscape(branchName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComputesClient) getHandleResponse(resp *http.Response) (ComputesClientGetResponse, error) {
	result := ComputesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Compute); err != nil {
		return ComputesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Compute resources by Branch
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - options - ComputesClientListOptions contains the optional parameters for the ComputesClient.NewListPager method.
func (client *ComputesClient) NewListPager(resourceGroupName string, organizationName string, projectName string, branchName string, options *ComputesClientListOptions) *runtime.Pager[ComputesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComputesClientListResponse]{
		More: func(page ComputesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComputesClientListResponse) (ComputesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ComputesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, options)
			}, nil)
			if err != nil {
				return ComputesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ComputesClient) listCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, _ *ComputesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/computes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if branchName == "" {
		return nil, errors.New("parameter branchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{branchName}", url.PathEscape(branchName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComputesClient) listHandleResponse(resp *http.Response) (ComputesClientListResponse, error) {
	result := ComputesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputeListResult); err != nil {
		return ComputesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - computeName - The name of the Compute
//   - properties - The resource properties to be updated.
//   - options - ComputesClientBeginUpdateOptions contains the optional parameters for the ComputesClient.BeginUpdate method.
func (client *ComputesClient) BeginUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, properties Compute, options *ComputesClientBeginUpdateOptions) (*runtime.Poller[ComputesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ComputesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ComputesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Compute
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ComputesClient) update(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, properties Compute, options *ComputesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ComputesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, computeName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ComputesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, computeName string, properties Compute, _ *ComputesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/computes/{computeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if branchName == "" {
		return nil, errors.New("parameter branchName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{branchName}", url.PathEscape(branchName))
	if computeName == "" {
		return nil, errors.New("parameter computeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computeName}", url.PathEscape(computeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
