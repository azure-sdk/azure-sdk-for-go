//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// EntitiesClient contains the methods for the Entities group.
// Don't use this type directly, use NewEntitiesClient() instead.
type EntitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEntitiesClient creates a new instance of EntitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEntitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EntitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EntitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Expand - Expands an entity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityID - entity ID
//   - parameters - The parameters required to execute an expand operation on the given entity.
//   - options - EntitiesClientExpandOptions contains the optional parameters for the EntitiesClient.Expand method.
func (client *EntitiesClient) Expand(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters EntityExpandParameters, options *EntitiesClientExpandOptions) (EntitiesClientExpandResponse, error) {
	var err error
	const operationName = "EntitiesClient.Expand"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.expandCreateRequest(ctx, resourceGroupName, workspaceName, entityID, parameters, options)
	if err != nil {
		return EntitiesClientExpandResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EntitiesClientExpandResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EntitiesClientExpandResponse{}, err
	}
	resp, err := client.expandHandleResponse(httpResp)
	return resp, err
}

// expandCreateRequest creates the Expand request.
func (client *EntitiesClient) expandCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters EntityExpandParameters, options *EntitiesClientExpandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/expand"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// expandHandleResponse handles the Expand response.
func (client *EntitiesClient) expandHandleResponse(resp *http.Response) (EntitiesClientExpandResponse, error) {
	result := EntitiesClientExpandResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityExpandResponse); err != nil {
		return EntitiesClientExpandResponse{}, err
	}
	return result, nil
}

// Get - Gets an entity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityID - entity ID
//   - options - EntitiesClientGetOptions contains the optional parameters for the EntitiesClient.Get method.
func (client *EntitiesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, options *EntitiesClientGetOptions) (EntitiesClientGetResponse, error) {
	var err error
	const operationName = "EntitiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, entityID, options)
	if err != nil {
		return EntitiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EntitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EntitiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EntitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, options *EntitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EntitiesClient) getHandleResponse(resp *http.Response) (EntitiesClientGetResponse, error) {
	result := EntitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EntitiesClientGetResponse{}, err
	}
	return result, nil
}

// GetInsights - Execute Insights for an entity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityID - entity ID
//   - parameters - The parameters required to execute insights on the given entity.
//   - options - EntitiesClientGetInsightsOptions contains the optional parameters for the EntitiesClient.GetInsights method.
func (client *EntitiesClient) GetInsights(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters EntityGetInsightsParameters, options *EntitiesClientGetInsightsOptions) (EntitiesClientGetInsightsResponse, error) {
	var err error
	const operationName = "EntitiesClient.GetInsights"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getInsightsCreateRequest(ctx, resourceGroupName, workspaceName, entityID, parameters, options)
	if err != nil {
		return EntitiesClientGetInsightsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EntitiesClientGetInsightsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EntitiesClientGetInsightsResponse{}, err
	}
	resp, err := client.getInsightsHandleResponse(httpResp)
	return resp, err
}

// getInsightsCreateRequest creates the GetInsights request.
func (client *EntitiesClient) getInsightsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, parameters EntityGetInsightsParameters, options *EntitiesClientGetInsightsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/getInsights"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// getInsightsHandleResponse handles the GetInsights response.
func (client *EntitiesClient) getInsightsHandleResponse(resp *http.Response) (EntitiesClientGetInsightsResponse, error) {
	result := EntitiesClientGetInsightsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityGetInsightsResponse); err != nil {
		return EntitiesClientGetInsightsResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all entities.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - EntitiesClientListOptions contains the optional parameters for the EntitiesClient.NewListPager method.
func (client *EntitiesClient) NewListPager(resourceGroupName string, workspaceName string, options *EntitiesClientListOptions) *runtime.Pager[EntitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntitiesClientListResponse]{
		More: func(page EntitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EntitiesClientListResponse) (EntitiesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EntitiesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return EntitiesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EntitiesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *EntitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntitiesClient) listHandleResponse(resp *http.Response) (EntitiesClientListResponse, error) {
	result := EntitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityList); err != nil {
		return EntitiesClientListResponse{}, err
	}
	return result, nil
}

// NewQueriesPager - Get Insights and Activities for an entity.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityID - entity ID
//   - kind - The Kind parameter for queries
//   - options - EntitiesClientQueriesOptions contains the optional parameters for the EntitiesClient.NewQueriesPager method.
func (client *EntitiesClient) NewQueriesPager(resourceGroupName string, workspaceName string, entityID string, kind EntityItemQueryKind, options *EntitiesClientQueriesOptions) *runtime.Pager[EntitiesClientQueriesResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntitiesClientQueriesResponse]{
		More: func(page EntitiesClientQueriesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *EntitiesClientQueriesResponse) (EntitiesClientQueriesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EntitiesClient.NewQueriesPager")
			req, err := client.queriesCreateRequest(ctx, resourceGroupName, workspaceName, entityID, kind, options)
			if err != nil {
				return EntitiesClientQueriesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EntitiesClientQueriesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EntitiesClientQueriesResponse{}, runtime.NewResponseError(resp)
			}
			return client.queriesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// queriesCreateRequest creates the Queries request.
func (client *EntitiesClient) queriesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, kind EntityItemQueryKind, options *EntitiesClientQueriesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/queries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	reqQP.Set("kind", string(kind))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// queriesHandleResponse handles the Queries response.
func (client *EntitiesClient) queriesHandleResponse(resp *http.Response) (EntitiesClientQueriesResponse, error) {
	result := EntitiesClientQueriesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetQueriesResponse); err != nil {
		return EntitiesClientQueriesResponse{}, err
	}
	return result, nil
}

// RunPlaybook - Triggers playbook on a specific entity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityIdentifier - Entity identifier.
//   - options - EntitiesClientRunPlaybookOptions contains the optional parameters for the EntitiesClient.RunPlaybook method.
func (client *EntitiesClient) RunPlaybook(ctx context.Context, resourceGroupName string, workspaceName string, entityIdentifier string, options *EntitiesClientRunPlaybookOptions) (EntitiesClientRunPlaybookResponse, error) {
	var err error
	const operationName = "EntitiesClient.RunPlaybook"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.runPlaybookCreateRequest(ctx, resourceGroupName, workspaceName, entityIdentifier, options)
	if err != nil {
		return EntitiesClientRunPlaybookResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EntitiesClientRunPlaybookResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EntitiesClientRunPlaybookResponse{}, err
	}
	return EntitiesClientRunPlaybookResponse{}, nil
}

// runPlaybookCreateRequest creates the RunPlaybook request.
func (client *EntitiesClient) runPlaybookCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityIdentifier string, options *EntitiesClientRunPlaybookOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityIdentifier}/runPlaybook"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityIdentifier == "" {
		return nil, errors.New("parameter entityIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityIdentifier}", url.PathEscape(entityIdentifier))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.RequestBody != nil {
		if err := runtime.MarshalAsJSON(req, *options.RequestBody); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
