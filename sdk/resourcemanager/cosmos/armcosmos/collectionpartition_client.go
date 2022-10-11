//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

import (
	"context"
	"errors"
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

// CollectionPartitionClient contains the methods for the CollectionPartition group.
// Don't use this type directly, use NewCollectionPartitionClient() instead.
type CollectionPartitionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCollectionPartitionClient creates a new instance of CollectionPartitionClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCollectionPartitionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CollectionPartitionClient, error) {
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
	client := &CollectionPartitionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given collection, split by partition.
// Generated from API version 2022-05-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// databaseRid - Cosmos DB database rid.
// collectionRid - Cosmos DB collection rid.
// filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
// name.value (name of the metric, can have an or of multiple names), startTime, endTime,
// and timeGrain. The supported operator is eq.
// options - CollectionPartitionClientListMetricsOptions contains the optional parameters for the CollectionPartitionClient.ListMetrics
// method.
func (client *CollectionPartitionClient) NewListMetricsPager(resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionClientListMetricsOptions) *runtime.Pager[CollectionPartitionClientListMetricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionPartitionClientListMetricsResponse]{
		More: func(page CollectionPartitionClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionPartitionClientListMetricsResponse) (CollectionPartitionClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter, options)
			if err != nil {
				return CollectionPartitionClientListMetricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CollectionPartitionClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionPartitionClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionPartitionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionPartitionClient) listMetricsHandleResponse(resp *http.Response) (CollectionPartitionClientListMetricsResponse, error) {
	result := CollectionPartitionClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartitionMetricListResult); err != nil {
		return CollectionPartitionClientListMetricsResponse{}, err
	}
	return result, nil
}

// NewListUsagesPager - Retrieves the usages (most recent storage data) for the given collection, split by partition.
// Generated from API version 2022-05-15-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// databaseRid - Cosmos DB database rid.
// collectionRid - Cosmos DB collection rid.
// options - CollectionPartitionClientListUsagesOptions contains the optional parameters for the CollectionPartitionClient.ListUsages
// method.
func (client *CollectionPartitionClient) NewListUsagesPager(resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionClientListUsagesOptions) *runtime.Pager[CollectionPartitionClientListUsagesResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionPartitionClientListUsagesResponse]{
		More: func(page CollectionPartitionClientListUsagesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionPartitionClientListUsagesResponse) (CollectionPartitionClientListUsagesResponse, error) {
			req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
			if err != nil {
				return CollectionPartitionClientListUsagesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CollectionPartitionClientListUsagesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionPartitionClientListUsagesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listUsagesHandleResponse(resp)
		},
	})
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *CollectionPartitionClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionClientListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *CollectionPartitionClient) listUsagesHandleResponse(resp *http.Response) (CollectionPartitionClientListUsagesResponse, error) {
	result := CollectionPartitionClientListUsagesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartitionUsagesResult); err != nil {
		return CollectionPartitionClientListUsagesResponse{}, err
	}
	return result, nil
}
