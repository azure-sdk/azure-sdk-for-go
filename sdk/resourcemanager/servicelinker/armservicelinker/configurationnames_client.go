//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ConfigurationNamesClient contains the methods for the ConfigurationNames group.
// Don't use this type directly, use NewConfigurationNamesClient() instead.
type ConfigurationNamesClient struct {
	internal *arm.Client
}

// NewConfigurationNamesClient creates a new instance of ConfigurationNamesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationNamesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationNamesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationNamesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists the configuration names generated by Service Connector for all target, client types, auth types.
//
// Generated from API version 2022-11-01-preview
//   - options - ConfigurationNamesClientListOptions contains the optional parameters for the ConfigurationNamesClient.NewListPager
//     method.
func (client *ConfigurationNamesClient) NewListPager(options *ConfigurationNamesClientListOptions) *runtime.Pager[ConfigurationNamesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationNamesClientListResponse]{
		More: func(page ConfigurationNamesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConfigurationNamesClientListResponse) (ConfigurationNamesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConfigurationNamesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ConfigurationNamesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationNamesClient) listCreateRequest(ctx context.Context, options *ConfigurationNamesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ServiceLinker/configurationNames"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationNamesClient) listHandleResponse(resp *http.Response) (ConfigurationNamesClientListResponse, error) {
	result := ConfigurationNamesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationNameResult); err != nil {
		return ConfigurationNamesClientListResponse{}, err
	}
	return result, nil
}
