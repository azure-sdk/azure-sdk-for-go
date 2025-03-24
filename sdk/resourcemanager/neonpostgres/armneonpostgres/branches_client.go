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

// BranchesClient contains the methods for the Branches group.
// Don't use this type directly, use NewBranchesClient() instead.
type BranchesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBranchesClient creates a new instance of BranchesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBranchesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BranchesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BranchesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - resource - Resource create parameters.
//   - options - BranchesClientBeginCreateOrUpdateOptions contains the optional parameters for the BranchesClient.BeginCreateOrUpdate
//     method.
func (client *BranchesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, resource Branch, options *BranchesClientBeginCreateOrUpdateOptions) (*runtime.Poller[BranchesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, organizationName, projectName, branchName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BranchesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BranchesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *BranchesClient) createOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, resource Branch, options *BranchesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BranchesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, resource, options)
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
func (client *BranchesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, resource Branch, _ *BranchesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - options - BranchesClientDeleteOptions contains the optional parameters for the BranchesClient.Delete method.
func (client *BranchesClient) Delete(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, options *BranchesClientDeleteOptions) (BranchesClientDeleteResponse, error) {
	var err error
	const operationName = "BranchesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, options)
	if err != nil {
		return BranchesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BranchesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BranchesClientDeleteResponse{}, err
	}
	return BranchesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BranchesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, _ *BranchesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - options - BranchesClientGetOptions contains the optional parameters for the BranchesClient.Get method.
func (client *BranchesClient) Get(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, options *BranchesClientGetOptions) (BranchesClientGetResponse, error) {
	var err error
	const operationName = "BranchesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, options)
	if err != nil {
		return BranchesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BranchesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BranchesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BranchesClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, _ *BranchesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}"
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
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BranchesClient) getHandleResponse(resp *http.Response) (BranchesClientGetResponse, error) {
	result := BranchesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Branch); err != nil {
		return BranchesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Branch resources by Project
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - options - BranchesClientListOptions contains the optional parameters for the BranchesClient.NewListPager method.
func (client *BranchesClient) NewListPager(resourceGroupName string, organizationName string, projectName string, options *BranchesClientListOptions) *runtime.Pager[BranchesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BranchesClientListResponse]{
		More: func(page BranchesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BranchesClientListResponse) (BranchesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BranchesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, organizationName, projectName, options)
			}, nil)
			if err != nil {
				return BranchesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BranchesClient) listCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, _ *BranchesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BranchesClient) listHandleResponse(resp *http.Response) (BranchesClientListResponse, error) {
	result := BranchesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BranchListResult); err != nil {
		return BranchesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - properties - The resource properties to be updated.
//   - options - BranchesClientBeginUpdateOptions contains the optional parameters for the BranchesClient.BeginUpdate method.
func (client *BranchesClient) BeginUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, properties Branch, options *BranchesClientBeginUpdateOptions) (*runtime.Poller[BranchesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, organizationName, projectName, branchName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BranchesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BranchesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Branch
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *BranchesClient) update(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, properties Branch, options *BranchesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BranchesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, properties, options)
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
func (client *BranchesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, properties Branch, _ *BranchesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
