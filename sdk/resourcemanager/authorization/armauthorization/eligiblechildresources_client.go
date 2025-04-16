// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// EligibleChildResourcesClient contains the methods for the EligibleChildResources group.
// Don't use this type directly, use NewEligibleChildResourcesClient() instead.
type EligibleChildResourcesClient struct {
	internal *arm.Client
}

// NewEligibleChildResourcesClient creates a new instance of EligibleChildResourcesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEligibleChildResourcesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EligibleChildResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EligibleChildResourcesClient{
		internal: cl,
	}
	return client, nil
}

// NewGetPager - Get the child resources of a resource on which user has eligible access
//
// Generated from API version 2024-09-01-preview
//   - scope - The scope of the role management policy.
//   - options - EligibleChildResourcesClientGetOptions contains the optional parameters for the EligibleChildResourcesClient.NewGetPager
//     method.
func (client *EligibleChildResourcesClient) NewGetPager(scope string, options *EligibleChildResourcesClientGetOptions) *runtime.Pager[EligibleChildResourcesClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[EligibleChildResourcesClientGetResponse]{
		More: func(page EligibleChildResourcesClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EligibleChildResourcesClientGetResponse) (EligibleChildResourcesClientGetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EligibleChildResourcesClient.NewGetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.getCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return EligibleChildResourcesClientGetResponse{}, err
			}
			return client.getHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// getCreateRequest creates the Get request.
func (client *EligibleChildResourcesClient) getCreateRequest(ctx context.Context, scope string, options *EligibleChildResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/eligibleChildResources"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EligibleChildResourcesClient) getHandleResponse(resp *http.Response) (EligibleChildResourcesClientGetResponse, error) {
	result := EligibleChildResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EligibleChildResourcesListResult); err != nil {
		return EligibleChildResourcesClientGetResponse{}, err
	}
	return result, nil
}
