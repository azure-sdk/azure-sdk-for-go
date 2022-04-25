//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// CollectionRegionClient contains the methods for the CollectionRegion group.
// Don't use this type directly, use NewCollectionRegionClient() instead.
type CollectionRegionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCollectionRegionClient creates a new instance of CollectionRegionClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCollectionRegionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CollectionRegionClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CollectionRegionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given database account, collection and
// region.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// region - Cosmos DB region, with spaces between words and each word capitalized.
// databaseRid - Cosmos DB database rid.
// collectionRid - Cosmos DB collection rid.
// filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
// name.value (name of the metric, can have an or of multiple names), startTime, endTime,
// and timeGrain. The supported operator is eq.
// options - CollectionRegionClientListMetricsOptions contains the optional parameters for the CollectionRegionClient.ListMetrics
// method.
func (client *CollectionRegionClient) NewListMetricsPager(resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionRegionClientListMetricsOptions) *runtime.Pager[CollectionRegionClientListMetricsResponse] {
	return runtime.NewPager(runtime.PageProcessor[CollectionRegionClientListMetricsResponse]{
		More: func(page CollectionRegionClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionRegionClientListMetricsResponse) (CollectionRegionClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, region, databaseRid, collectionRid, filter, options)
			if err != nil {
				return CollectionRegionClientListMetricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CollectionRegionClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionRegionClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionRegionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionRegionClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/metrics"
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
	if region == "" {
		return nil, errors.New("parameter region cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{region}", url.PathEscape(region))
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
	reqQP.Set("api-version", "2022-04-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionRegionClient) listMetricsHandleResponse(resp *http.Response) (CollectionRegionClientListMetricsResponse, error) {
	result := CollectionRegionClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricListResult); err != nil {
		return CollectionRegionClientListMetricsResponse{}, err
	}
	return result, nil
}
