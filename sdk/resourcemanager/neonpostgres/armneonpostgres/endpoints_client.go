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

// EndpointsClient contains the methods for the Endpoints group.
// Don't use this type directly, use NewEndpointsClient() instead.
type EndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEndpointsClient creates a new instance of EndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - endpointName - The name of the Endpoint
//   - resource - Resource create parameters.
//   - options - EndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the EndpointsClient.BeginCreateOrUpdate
//     method.
func (client *EndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, resource Endpoint, options *EndpointsClientBeginCreateOrUpdateOptions) (*runtime.Poller[EndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *EndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, resource Endpoint, options *EndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, resource, options)
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
func (client *EndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, resource Endpoint, _ *EndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/endpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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

// Delete - Delete a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - endpointName - The name of the Endpoint
//   - options - EndpointsClientDeleteOptions contains the optional parameters for the EndpointsClient.Delete method.
func (client *EndpointsClient) Delete(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, options *EndpointsClientDeleteOptions) (EndpointsClientDeleteResponse, error) {
	var err error
	const operationName = "EndpointsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, options)
	if err != nil {
		return EndpointsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EndpointsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EndpointsClientDeleteResponse{}, err
	}
	return EndpointsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, _ *EndpointsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/endpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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

// Get - Get a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - endpointName - The name of the Endpoint
//   - options - EndpointsClientGetOptions contains the optional parameters for the EndpointsClient.Get method.
func (client *EndpointsClient) Get(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, options *EndpointsClientGetOptions) (EndpointsClientGetResponse, error) {
	var err error
	const operationName = "EndpointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, options)
	if err != nil {
		return EndpointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EndpointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, _ *EndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/endpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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
func (client *EndpointsClient) getHandleResponse(resp *http.Response) (EndpointsClientGetResponse, error) {
	result := EndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Endpoint); err != nil {
		return EndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Endpoint resources by Branch
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - options - EndpointsClientListOptions contains the optional parameters for the EndpointsClient.NewListPager method.
func (client *EndpointsClient) NewListPager(resourceGroupName string, organizationName string, projectName string, branchName string, options *EndpointsClientListOptions) *runtime.Pager[EndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EndpointsClientListResponse]{
		More: func(page EndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointsClientListResponse) (EndpointsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EndpointsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, options)
			}, nil)
			if err != nil {
				return EndpointsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, _ *EndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/endpoints"
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

// listHandleResponse handles the List response.
func (client *EndpointsClient) listHandleResponse(resp *http.Response) (EndpointsClientListResponse, error) {
	result := EndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointListResult); err != nil {
		return EndpointsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Neon Organizations resource
//   - projectName - The name of the Project
//   - branchName - The name of the Branch
//   - endpointName - The name of the Endpoint
//   - properties - The resource properties to be updated.
//   - options - EndpointsClientBeginUpdateOptions contains the optional parameters for the EndpointsClient.BeginUpdate method.
func (client *EndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, properties Endpoint, options *EndpointsClientBeginUpdateOptions) (*runtime.Poller[EndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Endpoint
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01-preview
func (client *EndpointsClient) update(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, properties Endpoint, options *EndpointsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "EndpointsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, projectName, branchName, endpointName, properties, options)
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
func (client *EndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, projectName string, branchName string, endpointName string, properties Endpoint, _ *EndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Neon.Postgres/organizations/{organizationName}/projects/{projectName}/branches/{branchName}/endpoints/{endpointName}"
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
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
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
