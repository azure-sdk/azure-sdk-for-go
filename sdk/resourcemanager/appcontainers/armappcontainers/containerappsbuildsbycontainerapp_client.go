//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsBuildsByContainerAppClient contains the methods for the ContainerAppsBuildsByContainerApp group.
// Don't use this type directly, use NewContainerAppsBuildsByContainerAppClient() instead.
type ContainerAppsBuildsByContainerAppClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsBuildsByContainerAppClient creates a new instance of ContainerAppsBuildsByContainerAppClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsBuildsByContainerAppClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsBuildsByContainerAppClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsBuildsByContainerAppClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List Container Apps Build resources by Container App
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Build is associated.
//   - options - ContainerAppsBuildsByContainerAppClientListOptions contains the optional parameters for the ContainerAppsBuildsByContainerAppClient.NewListPager
//     method.
func (client *ContainerAppsBuildsByContainerAppClient) NewListPager(resourceGroupName string, containerAppName string, options *ContainerAppsBuildsByContainerAppClientListOptions) *runtime.Pager[ContainerAppsBuildsByContainerAppClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsBuildsByContainerAppClientListResponse]{
		More: func(page ContainerAppsBuildsByContainerAppClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsBuildsByContainerAppClientListResponse) (ContainerAppsBuildsByContainerAppClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerAppsBuildsByContainerAppClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, containerAppName, options)
			}, nil)
			if err != nil {
				return ContainerAppsBuildsByContainerAppClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ContainerAppsBuildsByContainerAppClient) listCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsBuildsByContainerAppClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/builds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContainerAppsBuildsByContainerAppClient) listHandleResponse(resp *http.Response) (ContainerAppsBuildsByContainerAppClientListResponse, error) {
	result := ContainerAppsBuildsByContainerAppClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppsBuildCollection); err != nil {
		return ContainerAppsBuildsByContainerAppClientListResponse{}, err
	}
	return result, nil
}
