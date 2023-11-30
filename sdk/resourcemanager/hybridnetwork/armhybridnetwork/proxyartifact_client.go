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

// ProxyArtifactClient contains the methods for the ProxyArtifact group.
// Don't use this type directly, use NewProxyArtifactClient() instead.
type ProxyArtifactClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProxyArtifactClient creates a new instance of ProxyArtifactClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProxyArtifactClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProxyArtifactClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProxyArtifactClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewGetPager - Get a Artifact overview information.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - artifactName - The name of the artifact.
//   - options - ProxyArtifactClientGetOptions contains the optional parameters for the ProxyArtifactClient.NewGetPager method.
func (client *ProxyArtifactClient) NewGetPager(resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, options *ProxyArtifactClientGetOptions) *runtime.Pager[ProxyArtifactClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProxyArtifactClientGetResponse]{
		More: func(page ProxyArtifactClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProxyArtifactClientGetResponse) (ProxyArtifactClientGetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProxyArtifactClient.NewGetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, artifactName, options)
			}, nil)
			if err != nil {
				return ProxyArtifactClientGetResponse{}, err
			}
			return client.getHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getCreateRequest creates the Get request.
func (client *ProxyArtifactClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, options *ProxyArtifactClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}/artifactVersions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("artifactName", artifactName)
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProxyArtifactClient) getHandleResponse(resp *http.Response) (ProxyArtifactClientGetResponse, error) {
	result := ProxyArtifactClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProxyArtifactVersionsOverviewListResult); err != nil {
		return ProxyArtifactClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the available artifacts in the parent Artifact Store.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - options - ProxyArtifactClientListOptions contains the optional parameters for the ProxyArtifactClient.NewListPager method.
func (client *ProxyArtifactClient) NewListPager(resourceGroupName string, publisherName string, artifactStoreName string, options *ProxyArtifactClientListOptions) *runtime.Pager[ProxyArtifactClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProxyArtifactClientListResponse]{
		More: func(page ProxyArtifactClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProxyArtifactClientListResponse) (ProxyArtifactClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProxyArtifactClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, options)
			}, nil)
			if err != nil {
				return ProxyArtifactClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ProxyArtifactClient) listCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, options *ProxyArtifactClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}/artifacts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
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

// listHandleResponse handles the List response.
func (client *ProxyArtifactClient) listHandleResponse(resp *http.Response) (ProxyArtifactClientListResponse, error) {
	result := ProxyArtifactClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProxyArtifactOverviewListResult); err != nil {
		return ProxyArtifactClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdateState - Change artifact state defined in artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - artifactStoreName - The name of the artifact store.
//   - artifactName - The name of the artifact.
//   - artifactVersionName - The name of the artifact version.
//   - parameters - Parameters supplied to update the state of artifact manifest.
//   - options - ProxyArtifactClientBeginUpdateStateOptions contains the optional parameters for the ProxyArtifactClient.BeginUpdateState
//     method.
func (client *ProxyArtifactClient) BeginUpdateState(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, artifactVersionName string, parameters ArtifactChangeState, options *ProxyArtifactClientBeginUpdateStateOptions) (*runtime.Poller[ProxyArtifactClientUpdateStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateState(ctx, resourceGroupName, publisherName, artifactStoreName, artifactName, artifactVersionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProxyArtifactClientUpdateStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProxyArtifactClientUpdateStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateState - Change artifact state defined in artifact store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *ProxyArtifactClient) updateState(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, artifactVersionName string, parameters ArtifactChangeState, options *ProxyArtifactClientBeginUpdateStateOptions) (*http.Response, error) {
	var err error
	const operationName = "ProxyArtifactClient.BeginUpdateState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateStateCreateRequest(ctx, resourceGroupName, publisherName, artifactStoreName, artifactName, artifactVersionName, parameters, options)
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
func (client *ProxyArtifactClient) updateStateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, artifactVersionName string, parameters ArtifactChangeState, options *ProxyArtifactClientBeginUpdateStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/artifactStores/{artifactStoreName}/artifactVersions/{artifactVersionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if artifactStoreName == "" {
		return nil, errors.New("parameter artifactStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactStoreName}", url.PathEscape(artifactStoreName))
	if artifactVersionName == "" {
		return nil, errors.New("parameter artifactVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactVersionName}", url.PathEscape(artifactVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("artifactName", artifactName)
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
