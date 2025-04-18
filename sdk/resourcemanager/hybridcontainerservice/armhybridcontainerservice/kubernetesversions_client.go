// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcontainerservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// KubernetesVersionsClient contains the methods for the KubernetesVersions group.
// Don't use this type directly, use NewKubernetesVersionsClient() instead.
type KubernetesVersionsClient struct {
	internal *arm.Client
}

// NewKubernetesVersionsClient creates a new instance of KubernetesVersionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKubernetesVersionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*KubernetesVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KubernetesVersionsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists the supported kubernetes versions for the specified custom location
//
// Generated from API version 2024-01-01
//   - customLocationResourceURI - The fully qualified Azure Resource Manager identifier of the custom location resource.
//   - options - KubernetesVersionsClientListOptions contains the optional parameters for the KubernetesVersionsClient.NewListPager
//     method.
func (client *KubernetesVersionsClient) NewListPager(customLocationResourceURI string, options *KubernetesVersionsClientListOptions) *runtime.Pager[KubernetesVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[KubernetesVersionsClientListResponse]{
		More: func(page KubernetesVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KubernetesVersionsClientListResponse) (KubernetesVersionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "KubernetesVersionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, customLocationResourceURI, options)
			}, nil)
			if err != nil {
				return KubernetesVersionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *KubernetesVersionsClient) listCreateRequest(ctx context.Context, customLocationResourceURI string, _ *KubernetesVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{customLocationResourceUri}/providers/Microsoft.HybridContainerService/kubernetesVersions"
	urlPath = strings.ReplaceAll(urlPath, "{customLocationResourceUri}", customLocationResourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *KubernetesVersionsClient) listHandleResponse(resp *http.Response) (KubernetesVersionsClientListResponse, error) {
	result := KubernetesVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KubernetesVersionProfileList); err != nil {
		return KubernetesVersionsClientListResponse{}, err
	}
	return result, nil
}
