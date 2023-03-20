//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SKUsClient contains the methods for the SKUs group.
// Don't use this type directly, use NewSKUsClient() instead.
type SKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSKUsClient creates a new instance of SKUsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List all the available Skus in the region and information related to them
//
// Generated from API version 2023-01-01
//   - options - SKUsClientListOptions contains the optional parameters for the SKUsClient.NewListPager method.
func (client *SKUsClient) NewListPager(options *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SKUsClientListResponse]{
		More: func(page SKUsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *SKUsClientListResponse) (SKUsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SKUsClient) listCreateRequest(ctx context.Context, options *SKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ElasticSan/skus"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SKUsClient) listHandleResponse(resp *http.Response) (SKUsClientListResponse, error) {
	result := SKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUInformationList); err != nil {
		return SKUsClientListResponse{}, err
	}
	return result, nil
}
