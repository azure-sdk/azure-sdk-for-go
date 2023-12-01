//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// NetworkFunctionDefinitionVersionsClient contains the methods for the NetworkFunctionDefinitionVersions group.
// Don't use this type directly, use NewNetworkFunctionDefinitionVersionsClient() instead.
type NetworkFunctionDefinitionVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkFunctionDefinitionVersionsClient creates a new instance of NetworkFunctionDefinitionVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkFunctionDefinitionVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkFunctionDefinitionVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkFunctionDefinitionVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network function definition version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - networkFunctionDefinitionVersionName - The name of the network function definition version. The name should conform to
//     the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to the create or update network function definition version operation.
//   - options - NetworkFunctionDefinitionVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkFunctionDefinitionVersionsClient.BeginCreateOrUpdate
//     method.
func (client *NetworkFunctionDefinitionVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersion, options *NetworkFunctionDefinitionVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkFunctionDefinitionVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionDefinitionVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionDefinitionVersionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a network function definition version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *NetworkFunctionDefinitionVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersion, options *NetworkFunctionDefinitionVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionVersionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, parameters, options)
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
func (client *NetworkFunctionDefinitionVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersion, options *NetworkFunctionDefinitionVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions/{networkFunctionDefinitionVersionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if networkFunctionDefinitionVersionName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionVersionName}", url.PathEscape(networkFunctionDefinitionVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified network function definition version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - networkFunctionDefinitionVersionName - The name of the network function definition version. The name should conform to
//     the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - options - NetworkFunctionDefinitionVersionsClientBeginDeleteOptions contains the optional parameters for the NetworkFunctionDefinitionVersionsClient.BeginDelete
//     method.
func (client *NetworkFunctionDefinitionVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, options *NetworkFunctionDefinitionVersionsClientBeginDeleteOptions) (*runtime.Poller[NetworkFunctionDefinitionVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionDefinitionVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionDefinitionVersionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified network function definition version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *NetworkFunctionDefinitionVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, options *NetworkFunctionDefinitionVersionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionVersionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkFunctionDefinitionVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, options *NetworkFunctionDefinitionVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions/{networkFunctionDefinitionVersionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if networkFunctionDefinitionVersionName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionVersionName}", url.PathEscape(networkFunctionDefinitionVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a network function definition version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - networkFunctionDefinitionVersionName - The name of the network function definition version. The name should conform to
//     the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - options - NetworkFunctionDefinitionVersionsClientGetOptions contains the optional parameters for the NetworkFunctionDefinitionVersionsClient.Get
//     method.
func (client *NetworkFunctionDefinitionVersionsClient) Get(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, options *NetworkFunctionDefinitionVersionsClientGetOptions) (NetworkFunctionDefinitionVersionsClientGetResponse, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, options)
	if err != nil {
		return NetworkFunctionDefinitionVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionDefinitionVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionDefinitionVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkFunctionDefinitionVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, options *NetworkFunctionDefinitionVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions/{networkFunctionDefinitionVersionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if networkFunctionDefinitionVersionName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionVersionName}", url.PathEscape(networkFunctionDefinitionVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkFunctionDefinitionVersionsClient) getHandleResponse(resp *http.Response) (NetworkFunctionDefinitionVersionsClientGetResponse, error) {
	result := NetworkFunctionDefinitionVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionVersion); err != nil {
		return NetworkFunctionDefinitionVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNetworkFunctionDefinitionGroupPager - Gets information about a list of network function definition versions under
// a network function definition group.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - options - NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupOptions contains the optional parameters
//     for the NetworkFunctionDefinitionVersionsClient.NewListByNetworkFunctionDefinitionGroupPager method.
func (client *NetworkFunctionDefinitionVersionsClient) NewListByNetworkFunctionDefinitionGroupPager(resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupOptions) *runtime.Pager[NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse]{
		More: func(page NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse) (NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkFunctionDefinitionVersionsClient.NewListByNetworkFunctionDefinitionGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNetworkFunctionDefinitionGroupCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, options)
			}, nil)
			if err != nil {
				return NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse{}, err
			}
			return client.listByNetworkFunctionDefinitionGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNetworkFunctionDefinitionGroupCreateRequest creates the ListByNetworkFunctionDefinitionGroup request.
func (client *NetworkFunctionDefinitionVersionsClient) listByNetworkFunctionDefinitionGroupCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNetworkFunctionDefinitionGroupHandleResponse handles the ListByNetworkFunctionDefinitionGroup response.
func (client *NetworkFunctionDefinitionVersionsClient) listByNetworkFunctionDefinitionGroupHandleResponse(resp *http.Response) (NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse, error) {
	result := NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionVersionListResult); err != nil {
		return NetworkFunctionDefinitionVersionsClientListByNetworkFunctionDefinitionGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a network function definition version resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - networkFunctionDefinitionVersionName - The name of the network function definition version. The name should conform to
//     the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to the create or update network function definition version operation.
//   - options - NetworkFunctionDefinitionVersionsClientUpdateOptions contains the optional parameters for the NetworkFunctionDefinitionVersionsClient.Update
//     method.
func (client *NetworkFunctionDefinitionVersionsClient) Update(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters TagsObject, options *NetworkFunctionDefinitionVersionsClientUpdateOptions) (NetworkFunctionDefinitionVersionsClientUpdateResponse, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionVersionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, parameters, options)
	if err != nil {
		return NetworkFunctionDefinitionVersionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionDefinitionVersionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionDefinitionVersionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *NetworkFunctionDefinitionVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters TagsObject, options *NetworkFunctionDefinitionVersionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions/{networkFunctionDefinitionVersionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if networkFunctionDefinitionVersionName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionVersionName}", url.PathEscape(networkFunctionDefinitionVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *NetworkFunctionDefinitionVersionsClient) updateHandleResponse(resp *http.Response) (NetworkFunctionDefinitionVersionsClientUpdateResponse, error) {
	result := NetworkFunctionDefinitionVersionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionVersion); err != nil {
		return NetworkFunctionDefinitionVersionsClientUpdateResponse{}, err
	}
	return result, nil
}

// BeginUpdateState - Update network function definition version state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - networkFunctionDefinitionVersionName - The name of the network function definition version. The name should conform to
//     the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to update the state of network function definition version.
//   - options - NetworkFunctionDefinitionVersionsClientBeginUpdateStateOptions contains the optional parameters for the NetworkFunctionDefinitionVersionsClient.BeginUpdateState
//     method.
func (client *NetworkFunctionDefinitionVersionsClient) BeginUpdateState(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersionUpdateState, options *NetworkFunctionDefinitionVersionsClientBeginUpdateStateOptions) (*runtime.Poller[NetworkFunctionDefinitionVersionsClientUpdateStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateState(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionDefinitionVersionsClientUpdateStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionDefinitionVersionsClientUpdateStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateState - Update network function definition version state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *NetworkFunctionDefinitionVersionsClient) updateState(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersionUpdateState, options *NetworkFunctionDefinitionVersionsClientBeginUpdateStateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionVersionsClient.BeginUpdateState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateStateCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, networkFunctionDefinitionVersionName, parameters, options)
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

// updateStateCreateRequest creates the UpdateState request.
func (client *NetworkFunctionDefinitionVersionsClient) updateStateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, networkFunctionDefinitionVersionName string, parameters NetworkFunctionDefinitionVersionUpdateState, options *NetworkFunctionDefinitionVersionsClientBeginUpdateStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}/networkFunctionDefinitionVersions/{networkFunctionDefinitionVersionName}/updateState"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if networkFunctionDefinitionVersionName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionVersionName}", url.PathEscape(networkFunctionDefinitionVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
