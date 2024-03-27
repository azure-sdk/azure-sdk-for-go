//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// CollectionClient contains the methods for the Collection group.
// Don't use this type directly, use NewCollectionClient() instead.
type CollectionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCollectionClient creates a new instance of CollectionClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCollectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CollectionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CollectionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListMetricDefinitionsPager - Retrieves metric definitions for the given collection.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - collectionRid - Cosmos DB collection rid.
//   - options - CollectionClientListMetricDefinitionsOptions contains the optional parameters for the CollectionClient.NewListMetricDefinitionsPager
//     method.
func (client *CollectionClient) NewListMetricDefinitionsPager(resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionClientListMetricDefinitionsOptions) *runtime.Pager[CollectionClientListMetricDefinitionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionClientListMetricDefinitionsResponse]{
		More: func(page CollectionClientListMetricDefinitionsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionClientListMetricDefinitionsResponse) (CollectionClientListMetricDefinitionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CollectionClient.NewListMetricDefinitionsPager")
			req, err := client.listMetricDefinitionsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
			if err != nil {
				return CollectionClientListMetricDefinitionsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CollectionClientListMetricDefinitionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionClientListMetricDefinitionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricDefinitionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listMetricDefinitionsCreateRequest creates the ListMetricDefinitions request.
func (client *CollectionClient) listMetricDefinitionsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionClientListMetricDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricDefinitionsHandleResponse handles the ListMetricDefinitions response.
func (client *CollectionClient) listMetricDefinitionsHandleResponse(resp *http.Response) (CollectionClientListMetricDefinitionsResponse, error) {
	result := CollectionClientListMetricDefinitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionsListResult); err != nil {
		return CollectionClientListMetricDefinitionsResponse{}, err
	}
	return result, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given database account and collection.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - collectionRid - Cosmos DB collection rid.
//   - filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
//     name.value (name of the metric, can have an or of multiple names), startTime, endTime,
//     and timeGrain. The supported operator is eq.
//   - options - CollectionClientListMetricsOptions contains the optional parameters for the CollectionClient.NewListMetricsPager
//     method.
func (client *CollectionClient) NewListMetricsPager(resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionClientListMetricsOptions) *runtime.Pager[CollectionClientListMetricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionClientListMetricsResponse]{
		More: func(page CollectionClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionClientListMetricsResponse) (CollectionClientListMetricsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CollectionClient.NewListMetricsPager")
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter, options)
			if err != nil {
				return CollectionClientListMetricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CollectionClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("$filter", filter)
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionClient) listMetricsHandleResponse(resp *http.Response) (CollectionClientListMetricsResponse, error) {
	result := CollectionClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return CollectionClientListMetricsResponse{}, err
	}
	return result, nil
}

// NewListUsagesPager - Retrieves the usages (most recent storage data) for the given collection.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - databaseRid - Cosmos DB database rid.
//   - collectionRid - Cosmos DB collection rid.
//   - options - CollectionClientListUsagesOptions contains the optional parameters for the CollectionClient.NewListUsagesPager
//     method.
func (client *CollectionClient) NewListUsagesPager(resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionClientListUsagesOptions) *runtime.Pager[CollectionClientListUsagesResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionClientListUsagesResponse]{
		More: func(page CollectionClientListUsagesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionClientListUsagesResponse) (CollectionClientListUsagesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CollectionClient.NewListUsagesPager")
			req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
			if err != nil {
				return CollectionClientListUsagesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CollectionClientListUsagesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionClientListUsagesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listUsagesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *CollectionClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionClientListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *CollectionClient) listUsagesHandleResponse(resp *http.Response) (CollectionClientListUsagesResponse, error) {
	result := CollectionClientListUsagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesResult); err != nil {
		return CollectionClientListUsagesResponse{}, err
	}
	return result, nil
}
