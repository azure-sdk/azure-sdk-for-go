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

// ManagedEnvironmentPrivateLinkResourcesClient contains the methods for the ManagedEnvironmentPrivateLinkResources group.
// Don't use this type directly, use NewManagedEnvironmentPrivateLinkResourcesClient() instead.
type ManagedEnvironmentPrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedEnvironmentPrivateLinkResourcesClient creates a new instance of ManagedEnvironmentPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedEnvironmentPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedEnvironmentPrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedEnvironmentPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List private link resources for a given managed environment.
//
// Generated from API version 2024-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - ManagedEnvironmentPrivateLinkResourcesClientListOptions contains the optional parameters for the ManagedEnvironmentPrivateLinkResourcesClient.NewListPager
//     method.
func (client *ManagedEnvironmentPrivateLinkResourcesClient) NewListPager(resourceGroupName string, environmentName string, options *ManagedEnvironmentPrivateLinkResourcesClientListOptions) *runtime.Pager[ManagedEnvironmentPrivateLinkResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedEnvironmentPrivateLinkResourcesClientListResponse]{
		More: func(page ManagedEnvironmentPrivateLinkResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedEnvironmentPrivateLinkResourcesClientListResponse) (ManagedEnvironmentPrivateLinkResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedEnvironmentPrivateLinkResourcesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			}, nil)
			if err != nil {
				return ManagedEnvironmentPrivateLinkResourcesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedEnvironmentPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, _ *ManagedEnvironmentPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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
func (client *ManagedEnvironmentPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (ManagedEnvironmentPrivateLinkResourcesClientListResponse, error) {
	result := ManagedEnvironmentPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return ManagedEnvironmentPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
