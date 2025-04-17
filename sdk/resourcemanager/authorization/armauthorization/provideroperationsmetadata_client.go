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

// ProviderOperationsMetadataClient contains the methods for the ProviderOperationsMetadata group.
// Don't use this type directly, use NewProviderOperationsMetadataClient() instead.
type ProviderOperationsMetadataClient struct {
	internal *arm.Client
}

// NewProviderOperationsMetadataClient creates a new instance of ProviderOperationsMetadataClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProviderOperationsMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderOperationsMetadataClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProviderOperationsMetadataClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets provider operations metadata for the specified resource provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01
//   - resourceProviderNamespace - The namespace of the resource provider.
//   - options - ProviderOperationsMetadataClientGetOptions contains the optional parameters for the ProviderOperationsMetadataClient.Get
//     method.
func (client *ProviderOperationsMetadataClient) Get(ctx context.Context, resourceProviderNamespace string, options *ProviderOperationsMetadataClientGetOptions) (ProviderOperationsMetadataClientGetResponse, error) {
	var err error
	const operationName = "ProviderOperationsMetadataClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProviderOperationsMetadataClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderOperationsMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderOperationsMetadataClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProviderOperationsMetadataClient) getCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProviderOperationsMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProviderOperationsMetadataClient) getHandleResponse(resp *http.Response) (ProviderOperationsMetadataClientGetResponse, error) {
	result := ProviderOperationsMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderOperationsMetadata); err != nil {
		return ProviderOperationsMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets provider operations metadata for all resource providers.
//
// Generated from API version 2022-04-01
//   - options - ProviderOperationsMetadataClientListOptions contains the optional parameters for the ProviderOperationsMetadataClient.NewListPager
//     method.
func (client *ProviderOperationsMetadataClient) NewListPager(options *ProviderOperationsMetadataClientListOptions) *runtime.Pager[ProviderOperationsMetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderOperationsMetadataClientListResponse]{
		More: func(page ProviderOperationsMetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderOperationsMetadataClientListResponse) (ProviderOperationsMetadataClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProviderOperationsMetadataClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ProviderOperationsMetadataClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ProviderOperationsMetadataClient) listCreateRequest(ctx context.Context, options *ProviderOperationsMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/providerOperations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProviderOperationsMetadataClient) listHandleResponse(resp *http.Response) (ProviderOperationsMetadataClientListResponse, error) {
	result := ProviderOperationsMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderOperationsMetadataListResult); err != nil {
		return ProviderOperationsMetadataClientListResponse{}, err
	}
	return result, nil
}
