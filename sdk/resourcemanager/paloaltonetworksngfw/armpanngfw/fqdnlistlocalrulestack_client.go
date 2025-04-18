// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpanngfw

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

// FqdnListLocalRulestackClient contains the methods for the FqdnListLocalRulestack group.
// Don't use this type directly, use NewFqdnListLocalRulestackClient() instead.
type FqdnListLocalRulestackClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFqdnListLocalRulestackClient creates a new instance of FqdnListLocalRulestackClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFqdnListLocalRulestackClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FqdnListLocalRulestackClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FqdnListLocalRulestackClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a FqdnListLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - fqdn list name
//   - resource - Resource create parameters.
//   - options - FqdnListLocalRulestackClientBeginCreateOrUpdateOptions contains the optional parameters for the FqdnListLocalRulestackClient.BeginCreateOrUpdate
//     method.
func (client *FqdnListLocalRulestackClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource FqdnListLocalRulestackResource, options *FqdnListLocalRulestackClientBeginCreateOrUpdateOptions) (*runtime.Poller[FqdnListLocalRulestackClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, localRulestackName, name, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FqdnListLocalRulestackClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FqdnListLocalRulestackClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a FqdnListLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-06-preview
func (client *FqdnListLocalRulestackClient) createOrUpdate(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource FqdnListLocalRulestackResource, options *FqdnListLocalRulestackClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FqdnListLocalRulestackClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, localRulestackName, name, resource, options)
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
func (client *FqdnListLocalRulestackClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, resource FqdnListLocalRulestackResource, _ *FqdnListLocalRulestackClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/fqdnlists/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a FqdnListLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - fqdn list name
//   - options - FqdnListLocalRulestackClientBeginDeleteOptions contains the optional parameters for the FqdnListLocalRulestackClient.BeginDelete
//     method.
func (client *FqdnListLocalRulestackClient) BeginDelete(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *FqdnListLocalRulestackClientBeginDeleteOptions) (*runtime.Poller[FqdnListLocalRulestackClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, localRulestackName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FqdnListLocalRulestackClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FqdnListLocalRulestackClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a FqdnListLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-06-preview
func (client *FqdnListLocalRulestackClient) deleteOperation(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *FqdnListLocalRulestackClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FqdnListLocalRulestackClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, localRulestackName, name, options)
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
func (client *FqdnListLocalRulestackClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, _ *FqdnListLocalRulestackClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/fqdnlists/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a FqdnListLocalRulestackResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - name - fqdn list name
//   - options - FqdnListLocalRulestackClientGetOptions contains the optional parameters for the FqdnListLocalRulestackClient.Get
//     method.
func (client *FqdnListLocalRulestackClient) Get(ctx context.Context, resourceGroupName string, localRulestackName string, name string, options *FqdnListLocalRulestackClientGetOptions) (FqdnListLocalRulestackClientGetResponse, error) {
	var err error
	const operationName = "FqdnListLocalRulestackClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, localRulestackName, name, options)
	if err != nil {
		return FqdnListLocalRulestackClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FqdnListLocalRulestackClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FqdnListLocalRulestackClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FqdnListLocalRulestackClient) getCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, name string, _ *FqdnListLocalRulestackClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/fqdnlists/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FqdnListLocalRulestackClient) getHandleResponse(resp *http.Response) (FqdnListLocalRulestackClientGetResponse, error) {
	result := FqdnListLocalRulestackClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FqdnListLocalRulestackResource); err != nil {
		return FqdnListLocalRulestackClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocalRulestacksPager - List FqdnListLocalRulestackResource resources by LocalRulestacks
//
// Generated from API version 2025-02-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - localRulestackName - LocalRulestack resource name
//   - options - FqdnListLocalRulestackClientListByLocalRulestacksOptions contains the optional parameters for the FqdnListLocalRulestackClient.NewListByLocalRulestacksPager
//     method.
func (client *FqdnListLocalRulestackClient) NewListByLocalRulestacksPager(resourceGroupName string, localRulestackName string, options *FqdnListLocalRulestackClientListByLocalRulestacksOptions) *runtime.Pager[FqdnListLocalRulestackClientListByLocalRulestacksResponse] {
	return runtime.NewPager(runtime.PagingHandler[FqdnListLocalRulestackClientListByLocalRulestacksResponse]{
		More: func(page FqdnListLocalRulestackClientListByLocalRulestacksResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FqdnListLocalRulestackClientListByLocalRulestacksResponse) (FqdnListLocalRulestackClientListByLocalRulestacksResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FqdnListLocalRulestackClient.NewListByLocalRulestacksPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocalRulestacksCreateRequest(ctx, resourceGroupName, localRulestackName, options)
			}, nil)
			if err != nil {
				return FqdnListLocalRulestackClientListByLocalRulestacksResponse{}, err
			}
			return client.listByLocalRulestacksHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocalRulestacksCreateRequest creates the ListByLocalRulestacks request.
func (client *FqdnListLocalRulestackClient) listByLocalRulestacksCreateRequest(ctx context.Context, resourceGroupName string, localRulestackName string, _ *FqdnListLocalRulestackClientListByLocalRulestacksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/PaloAltoNetworks.Cloudngfw/localRulestacks/{localRulestackName}/fqdnlists"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localRulestackName == "" {
		return nil, errors.New("parameter localRulestackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localRulestackName}", url.PathEscape(localRulestackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocalRulestacksHandleResponse handles the ListByLocalRulestacks response.
func (client *FqdnListLocalRulestackClient) listByLocalRulestacksHandleResponse(resp *http.Response) (FqdnListLocalRulestackClientListByLocalRulestacksResponse, error) {
	result := FqdnListLocalRulestackClientListByLocalRulestacksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FqdnListLocalRulestackResourceListResult); err != nil {
		return FqdnListLocalRulestackClientListByLocalRulestacksResponse{}, err
	}
	return result, nil
}
