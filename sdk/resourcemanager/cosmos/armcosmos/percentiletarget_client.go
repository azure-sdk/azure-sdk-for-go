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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PercentileTargetClient contains the methods for the PercentileTarget group.
// Don't use this type directly, use NewPercentileTargetClient() instead.
type PercentileTargetClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPercentileTargetClient creates a new instance of PercentileTargetClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPercentileTargetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PercentileTargetClient, error) {
	cl, err := arm.NewClient(moduleName+".PercentileTargetClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PercentileTargetClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given account target region. This url
// is only for PBS and Replication Latency data
//
// Generated from API version 2023-08-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - targetRegion - Target region to which data is written. Cosmos DB region, with spaces between words and each word capitalized.
//   - filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
//     name.value (name of the metric, can have an or of multiple names), startTime, endTime,
//     and timeGrain. The supported operator is eq.
//   - options - PercentileTargetClientListMetricsOptions contains the optional parameters for the PercentileTargetClient.NewListMetricsPager
//     method.
func (client *PercentileTargetClient) NewListMetricsPager(resourceGroupName string, accountName string, targetRegion string, filter string, options *PercentileTargetClientListMetricsOptions) *runtime.Pager[PercentileTargetClientListMetricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[PercentileTargetClientListMetricsResponse]{
		More: func(page PercentileTargetClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PercentileTargetClientListMetricsResponse) (PercentileTargetClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, targetRegion, filter, options)
			if err != nil {
				return PercentileTargetClientListMetricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PercentileTargetClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PercentileTargetClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *PercentileTargetClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, targetRegion string, filter string, options *PercentileTargetClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/targetRegion/{targetRegion}/percentile/metrics"
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
	if targetRegion == "" {
		return nil, errors.New("parameter targetRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetRegion}", url.PathEscape(targetRegion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *PercentileTargetClient) listMetricsHandleResponse(resp *http.Response) (PercentileTargetClientListMetricsResponse, error) {
	result := PercentileTargetClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PercentileMetricListResult); err != nil {
		return PercentileTargetClientListMetricsResponse{}, err
	}
	return result, nil
}
