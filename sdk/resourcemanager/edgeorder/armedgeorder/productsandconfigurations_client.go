//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

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

// ProductsAndConfigurationsClient contains the methods for the ProductsAndConfigurations group.
// Don't use this type directly, use NewProductsAndConfigurationsClient() instead.
type ProductsAndConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProductsAndConfigurationsClient creates a new instance of ProductsAndConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProductsAndConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductsAndConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProductsAndConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListConfigurationsPager - List configurations for the given product family, product line and product for the given subscription.
//
// Generated from API version 2024-02-01
//   - configurationsRequest - Filters for showing the configurations.
//   - options - ProductsAndConfigurationsClientListConfigurationsOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListConfigurationsPager
//     method.
func (client *ProductsAndConfigurationsClient) NewListConfigurationsPager(configurationsRequest ConfigurationsRequest, options *ProductsAndConfigurationsClientListConfigurationsOptions) *runtime.Pager[ProductsAndConfigurationsClientListConfigurationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsAndConfigurationsClientListConfigurationsResponse]{
		More: func(page ProductsAndConfigurationsClientListConfigurationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsAndConfigurationsClientListConfigurationsResponse) (ProductsAndConfigurationsClientListConfigurationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsAndConfigurationsClient.NewListConfigurationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listConfigurationsCreateRequest(ctx, configurationsRequest, options)
			}, nil)
			if err != nil {
				return ProductsAndConfigurationsClientListConfigurationsResponse{}, err
			}
			return client.listConfigurationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listConfigurationsCreateRequest creates the ListConfigurations request.
func (client *ProductsAndConfigurationsClient) listConfigurationsCreateRequest(ctx context.Context, configurationsRequest ConfigurationsRequest, options *ProductsAndConfigurationsClientListConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/listConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configurationsRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// listConfigurationsHandleResponse handles the ListConfigurations response.
func (client *ProductsAndConfigurationsClient) listConfigurationsHandleResponse(resp *http.Response) (ProductsAndConfigurationsClientListConfigurationsResponse, error) {
	result := ProductsAndConfigurationsClientListConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Configurations); err != nil {
		return ProductsAndConfigurationsClientListConfigurationsResponse{}, err
	}
	return result, nil
}

// NewListProductFamiliesPager - List product families for the given subscription.
//
// Generated from API version 2024-02-01
//   - productFamiliesRequest - Filters for showing the product families.
//   - options - ProductsAndConfigurationsClientListProductFamiliesOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListProductFamiliesPager
//     method.
func (client *ProductsAndConfigurationsClient) NewListProductFamiliesPager(productFamiliesRequest ProductFamiliesRequest, options *ProductsAndConfigurationsClientListProductFamiliesOptions) *runtime.Pager[ProductsAndConfigurationsClientListProductFamiliesResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsAndConfigurationsClientListProductFamiliesResponse]{
		More: func(page ProductsAndConfigurationsClientListProductFamiliesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsAndConfigurationsClientListProductFamiliesResponse) (ProductsAndConfigurationsClientListProductFamiliesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsAndConfigurationsClient.NewListProductFamiliesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listProductFamiliesCreateRequest(ctx, productFamiliesRequest, options)
			}, nil)
			if err != nil {
				return ProductsAndConfigurationsClientListProductFamiliesResponse{}, err
			}
			return client.listProductFamiliesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listProductFamiliesCreateRequest creates the ListProductFamilies request.
func (client *ProductsAndConfigurationsClient) listProductFamiliesCreateRequest(ctx context.Context, productFamiliesRequest ProductFamiliesRequest, options *ProductsAndConfigurationsClientListProductFamiliesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/listProductFamilies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, productFamiliesRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// listProductFamiliesHandleResponse handles the ListProductFamilies response.
func (client *ProductsAndConfigurationsClient) listProductFamiliesHandleResponse(resp *http.Response) (ProductsAndConfigurationsClientListProductFamiliesResponse, error) {
	result := ProductsAndConfigurationsClientListProductFamiliesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductFamilies); err != nil {
		return ProductsAndConfigurationsClientListProductFamiliesResponse{}, err
	}
	return result, nil
}

// NewListProductFamiliesMetadataPager - List product families metadata for the given subscription.
//
// Generated from API version 2024-02-01
//   - options - ProductsAndConfigurationsClientListProductFamiliesMetadataOptions contains the optional parameters for the ProductsAndConfigurationsClient.NewListProductFamiliesMetadataPager
//     method.
func (client *ProductsAndConfigurationsClient) NewListProductFamiliesMetadataPager(options *ProductsAndConfigurationsClientListProductFamiliesMetadataOptions) *runtime.Pager[ProductsAndConfigurationsClientListProductFamiliesMetadataResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductsAndConfigurationsClientListProductFamiliesMetadataResponse]{
		More: func(page ProductsAndConfigurationsClientListProductFamiliesMetadataResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductsAndConfigurationsClientListProductFamiliesMetadataResponse) (ProductsAndConfigurationsClientListProductFamiliesMetadataResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProductsAndConfigurationsClient.NewListProductFamiliesMetadataPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listProductFamiliesMetadataCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ProductsAndConfigurationsClientListProductFamiliesMetadataResponse{}, err
			}
			return client.listProductFamiliesMetadataHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listProductFamiliesMetadataCreateRequest creates the ListProductFamiliesMetadata request.
func (client *ProductsAndConfigurationsClient) listProductFamiliesMetadataCreateRequest(ctx context.Context, options *ProductsAndConfigurationsClientListProductFamiliesMetadataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/productFamiliesMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listProductFamiliesMetadataHandleResponse handles the ListProductFamiliesMetadata response.
func (client *ProductsAndConfigurationsClient) listProductFamiliesMetadataHandleResponse(resp *http.Response) (ProductsAndConfigurationsClientListProductFamiliesMetadataResponse, error) {
	result := ProductsAndConfigurationsClientListProductFamiliesMetadataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductFamiliesMetadata); err != nil {
		return ProductsAndConfigurationsClientListProductFamiliesMetadataResponse{}, err
	}
	return result, nil
}
