//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// HybridIdentityMetadataClient contains the methods for the HybridIdentityMetadata group.
// Don't use this type directly, use NewHybridIdentityMetadataClient() instead.
type HybridIdentityMetadataClient struct {
	internal *arm.Client
}

// NewHybridIdentityMetadataClient creates a new instance of HybridIdentityMetadataClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHybridIdentityMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridIdentityMetadataClient, error) {
	cl, err := arm.NewClient(moduleName+".HybridIdentityMetadataClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HybridIdentityMetadataClient{
		internal: cl,
	}
	return client, nil
}

// Get - Implements HybridIdentityMetadata GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
//   - options - HybridIdentityMetadataClientGetOptions contains the optional parameters for the HybridIdentityMetadataClient.Get
//     method.
func (client *HybridIdentityMetadataClient) Get(ctx context.Context, resourceURI string, options *HybridIdentityMetadataClientGetOptions) (HybridIdentityMetadataClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HybridIdentityMetadataClient) getCreateRequest(ctx context.Context, resourceURI string, options *HybridIdentityMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default/hybridIdentityMetadata/default"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridIdentityMetadataClient) getHandleResponse(resp *http.Response) (HybridIdentityMetadataClientGetResponse, error) {
	result := HybridIdentityMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns the list of HybridIdentityMetadata of the given vm.
//
// Generated from API version 2023-07-01-preview
//   - resourceURI - The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
//   - options - HybridIdentityMetadataClientListOptions contains the optional parameters for the HybridIdentityMetadataClient.NewListPager
//     method.
func (client *HybridIdentityMetadataClient) NewListPager(resourceURI string, options *HybridIdentityMetadataClientListOptions) *runtime.Pager[HybridIdentityMetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridIdentityMetadataClientListResponse]{
		More: func(page HybridIdentityMetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridIdentityMetadataClientListResponse) (HybridIdentityMetadataClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridIdentityMetadataClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HybridIdentityMetadataClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridIdentityMetadataClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HybridIdentityMetadataClient) listCreateRequest(ctx context.Context, resourceURI string, options *HybridIdentityMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AzureStackHCI/virtualMachineInstances/default/hybridIdentityMetadata"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HybridIdentityMetadataClient) listHandleResponse(resp *http.Response) (HybridIdentityMetadataClientListResponse, error) {
	result := HybridIdentityMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadataList); err != nil {
		return HybridIdentityMetadataClientListResponse{}, err
	}
	return result, nil
}
