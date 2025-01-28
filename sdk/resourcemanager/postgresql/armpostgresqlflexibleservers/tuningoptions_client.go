//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// TuningOptionsClient contains the methods for the TuningOptions group.
// Don't use this type directly, use NewTuningOptionsClient() instead.
type TuningOptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTuningOptionsClient creates a new instance of TuningOptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTuningOptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TuningOptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TuningOptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the tuning option on a server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - tuningOption - The name of the tuning option.
//   - options - TuningOptionsClientGetOptions contains the optional parameters for the TuningOptionsClient.Get method.
func (client *TuningOptionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, tuningOption TuningOptionEnum, options *TuningOptionsClientGetOptions) (TuningOptionsClientGetResponse, error) {
	var err error
	const operationName = "TuningOptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, tuningOption, options)
	if err != nil {
		return TuningOptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TuningOptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TuningOptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TuningOptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, tuningOption TuningOptionEnum, options *TuningOptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/tuningOptions/{tuningOption}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if tuningOption == "" {
		return nil, errors.New("parameter tuningOption cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tuningOption}", url.PathEscape(string(tuningOption)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TuningOptionsClient) getHandleResponse(resp *http.Response) (TuningOptionsClientGetResponse, error) {
	result := TuningOptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TuningOptionsResource); err != nil {
		return TuningOptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Retrieve the list of available tuning options.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - TuningOptionsClientListByServerOptions contains the optional parameters for the TuningOptionsClient.NewListByServerPager
//     method.
func (client *TuningOptionsClient) NewListByServerPager(resourceGroupName string, serverName string, options *TuningOptionsClientListByServerOptions) *runtime.Pager[TuningOptionsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[TuningOptionsClientListByServerResponse]{
		More: func(page TuningOptionsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TuningOptionsClientListByServerResponse) (TuningOptionsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TuningOptionsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return TuningOptionsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *TuningOptionsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *TuningOptionsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/tuningOptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *TuningOptionsClient) listByServerHandleResponse(resp *http.Response) (TuningOptionsClientListByServerResponse, error) {
	result := TuningOptionsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TuningOptionsListResult); err != nil {
		return TuningOptionsClientListByServerResponse{}, err
	}
	return result, nil
}

// NewListRecommendationsPager - Retrieve the list of available tuning index recommendations.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - tuningOption - The name of the tuning option.
//   - options - TuningOptionsClientListRecommendationsOptions contains the optional parameters for the TuningOptionsClient.NewListRecommendationsPager
//     method.
func (client *TuningOptionsClient) NewListRecommendationsPager(resourceGroupName string, serverName string, tuningOption TuningOptionEnum, options *TuningOptionsClientListRecommendationsOptions) *runtime.Pager[TuningOptionsClientListRecommendationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[TuningOptionsClientListRecommendationsResponse]{
		More: func(page TuningOptionsClientListRecommendationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TuningOptionsClientListRecommendationsResponse) (TuningOptionsClientListRecommendationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TuningOptionsClient.NewListRecommendationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listRecommendationsCreateRequest(ctx, resourceGroupName, serverName, tuningOption, options)
			}, nil)
			if err != nil {
				return TuningOptionsClientListRecommendationsResponse{}, err
			}
			return client.listRecommendationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listRecommendationsCreateRequest creates the ListRecommendations request.
func (client *TuningOptionsClient) listRecommendationsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, tuningOption TuningOptionEnum, options *TuningOptionsClientListRecommendationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/tuningOptions/{tuningOption}/recommendations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if tuningOption == "" {
		return nil, errors.New("parameter tuningOption cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tuningOption}", url.PathEscape(string(tuningOption)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.RecommendationType != nil {
		reqQP.Set("$recommendationType", string(*options.RecommendationType))
	}
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRecommendationsHandleResponse handles the ListRecommendations response.
func (client *TuningOptionsClient) listRecommendationsHandleResponse(resp *http.Response) (TuningOptionsClientListRecommendationsResponse, error) {
	result := TuningOptionsClientListRecommendationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IndexRecommendationListResult); err != nil {
		return TuningOptionsClientListRecommendationsResponse{}, err
	}
	return result, nil
}
