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

// EntityQueryTemplatesClient contains the methods for the EntityQueryTemplates group.
// Don't use this type directly, use NewEntityQueryTemplatesClient() instead.
type EntityQueryTemplatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEntityQueryTemplatesClient creates a new instance of EntityQueryTemplatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEntityQueryTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EntityQueryTemplatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EntityQueryTemplatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an entity query.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityQueryTemplateID - entity query template ID
//   - options - EntityQueryTemplatesClientGetOptions contains the optional parameters for the EntityQueryTemplatesClient.Get
//     method.
func (client *EntityQueryTemplatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string, options *EntityQueryTemplatesClientGetOptions) (EntityQueryTemplatesClientGetResponse, error) {
	var err error
	const operationName = "EntityQueryTemplatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, entityQueryTemplateID, options)
	if err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EntityQueryTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryTemplateID string, _ *EntityQueryTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates/{entityQueryTemplateId}"
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
	if entityQueryTemplateID == "" {
		return nil, errors.New("parameter entityQueryTemplateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityQueryTemplateId}", url.PathEscape(entityQueryTemplateID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EntityQueryTemplatesClient) getHandleResponse(resp *http.Response) (EntityQueryTemplatesClientGetResponse, error) {
	result := EntityQueryTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EntityQueryTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all entity query templates.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - EntityQueryTemplatesClientListOptions contains the optional parameters for the EntityQueryTemplatesClient.NewListPager
//     method.
func (client *EntityQueryTemplatesClient) NewListPager(resourceGroupName string, workspaceName string, options *EntityQueryTemplatesClientListOptions) *runtime.Pager[EntityQueryTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntityQueryTemplatesClientListResponse]{
		More: func(page EntityQueryTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EntityQueryTemplatesClientListResponse) (EntityQueryTemplatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EntityQueryTemplatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return EntityQueryTemplatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EntityQueryTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *EntityQueryTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueryTemplates"
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
	reqQP.Set("api-version", "2025-04-01-preview")
	if options != nil && options.Kind != nil {
		reqQP.Set("kind", string(*options.Kind))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntityQueryTemplatesClient) listHandleResponse(resp *http.Response) (EntityQueryTemplatesClientListResponse, error) {
	result := EntityQueryTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityQueryTemplateList); err != nil {
		return EntityQueryTemplatesClientListResponse{}, err
	}
	return result, nil
}
