// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API, command line tools, or the portal.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListSupportedPager - Gets a list of all supported private link resource types for the given service.
//
// Generated from API version 2025-05-01
//   - resourceGroupName - The name of the resource group within the current subscription. You can obtain this value from the
//     Azure Resource Manager API or the portal.
//   - searchServiceName - The name of the Azure AI Search service associated with the specified resource group.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - PrivateLinkResourcesClientListSupportedOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListSupportedPager
//     method.
func (client *PrivateLinkResourcesClient) NewListSupportedPager(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *PrivateLinkResourcesClientListSupportedOptions) *runtime.Pager[PrivateLinkResourcesClientListSupportedResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListSupportedResponse]{
		More: func(page PrivateLinkResourcesClientListSupportedResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListSupportedResponse) (PrivateLinkResourcesClientListSupportedResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkResourcesClient.NewListSupportedPager")
			req, err := client.listSupportedCreateRequest(ctx, resourceGroupName, searchServiceName, searchManagementRequestOptions, options)
			if err != nil {
				return PrivateLinkResourcesClientListSupportedResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkResourcesClientListSupportedResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientListSupportedResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSupportedHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSupportedCreateRequest creates the ListSupported request.
func (client *PrivateLinkResourcesClient) listSupportedCreateRequest(ctx context.Context, resourceGroupName string, searchServiceName string, searchManagementRequestOptions *SearchManagementRequestOptions, _ *PrivateLinkResourcesClientListSupportedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if searchServiceName == "" {
		return nil, errors.New("parameter searchServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{searchServiceName}", url.PathEscape(searchServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// listSupportedHandleResponse handles the ListSupported response.
func (client *PrivateLinkResourcesClient) listSupportedHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListSupportedResponse, error) {
	result := PrivateLinkResourcesClientListSupportedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourcesResult); err != nil {
		return PrivateLinkResourcesClientListSupportedResponse{}, err
	}
	return result, nil
}
