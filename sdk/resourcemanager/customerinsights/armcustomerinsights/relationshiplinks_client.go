// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// RelationshipLinksClient contains the methods for the RelationshipLinks group.
// Don't use this type directly, use NewRelationshipLinksClient() instead.
type RelationshipLinksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRelationshipLinksClient creates a new instance of RelationshipLinksClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRelationshipLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RelationshipLinksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RelationshipLinksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a relationship link or updates an existing relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipLinkName - The name of the relationship link.
//   - parameters - Parameters supplied to the CreateOrUpdate relationship link operation.
//   - options - RelationshipLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the RelationshipLinksClient.BeginCreateOrUpdate
//     method.
func (client *RelationshipLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, options *RelationshipLinksClientBeginCreateOrUpdateOptions) (*runtime.Poller[RelationshipLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, relationshipLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RelationshipLinksClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RelationshipLinksClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a relationship link or updates an existing relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
func (client *RelationshipLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, options *RelationshipLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RelationshipLinksClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RelationshipLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, _ *RelationshipLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipLinkName - The name of the relationship.
//   - options - RelationshipLinksClientBeginDeleteOptions contains the optional parameters for the RelationshipLinksClient.BeginDelete
//     method.
func (client *RelationshipLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientBeginDeleteOptions) (*runtime.Poller[RelationshipLinksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, relationshipLinkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RelationshipLinksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RelationshipLinksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
func (client *RelationshipLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RelationshipLinksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *RelationshipLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, _ *RelationshipLinksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about the specified relationship Link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipLinkName - The name of the relationship link.
//   - options - RelationshipLinksClientGetOptions contains the optional parameters for the RelationshipLinksClient.Get method.
func (client *RelationshipLinksClient) Get(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientGetOptions) (RelationshipLinksClientGetResponse, error) {
	var err error
	const operationName = "RelationshipLinksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, options)
	if err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RelationshipLinksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RelationshipLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, _ *RelationshipLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RelationshipLinksClient) getHandleResponse(resp *http.Response) (RelationshipLinksClientGetResponse, error) {
	result := RelationshipLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipLinkResourceFormat); err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all relationship links in the hub.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - options - RelationshipLinksClientListByHubOptions contains the optional parameters for the RelationshipLinksClient.NewListByHubPager
//     method.
func (client *RelationshipLinksClient) NewListByHubPager(resourceGroupName string, hubName string, options *RelationshipLinksClientListByHubOptions) *runtime.Pager[RelationshipLinksClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[RelationshipLinksClientListByHubResponse]{
		More: func(page RelationshipLinksClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RelationshipLinksClientListByHubResponse) (RelationshipLinksClientListByHubResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RelationshipLinksClient.NewListByHubPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			}, nil)
			if err != nil {
				return RelationshipLinksClientListByHubResponse{}, err
			}
			return client.listByHubHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *RelationshipLinksClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, _ *RelationshipLinksClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *RelationshipLinksClient) listByHubHandleResponse(resp *http.Response) (RelationshipLinksClientListByHubResponse, error) {
	result := RelationshipLinksClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipLinkListResult); err != nil {
		return RelationshipLinksClientListByHubResponse{}, err
	}
	return result, nil
}
