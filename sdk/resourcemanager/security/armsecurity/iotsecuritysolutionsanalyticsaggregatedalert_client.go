//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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
	"strconv"
	"strings"
)

// IotSecuritySolutionsAnalyticsAggregatedAlertClient contains the methods for the IotSecuritySolutionsAnalyticsAggregatedAlert group.
// Don't use this type directly, use NewIotSecuritySolutionsAnalyticsAggregatedAlertClient() instead.
type IotSecuritySolutionsAnalyticsAggregatedAlertClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIotSecuritySolutionsAnalyticsAggregatedAlertClient creates a new instance of IotSecuritySolutionsAnalyticsAggregatedAlertClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIotSecuritySolutionsAnalyticsAggregatedAlertClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IotSecuritySolutionsAnalyticsAggregatedAlertClient, error) {
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
	client := &IotSecuritySolutionsAnalyticsAggregatedAlertClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Dismiss - Use this method to dismiss an aggregated IoT Security Solution Alert.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - solutionName - The name of the IoT Security solution.
//   - aggregatedAlertName - Identifier of the aggregated alert.
//   - options - IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissOptions contains the optional parameters for the IotSecuritySolutionsAnalyticsAggregatedAlertClient.Dismiss
//     method.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) Dismiss(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissOptions) (IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse, error) {
	req, err := client.dismissCreateRequest(ctx, resourceGroupName, solutionName, aggregatedAlertName, options)
	if err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse{}, runtime.NewResponseError(resp)
	}
	return IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse{}, nil
}

// dismissCreateRequest creates the Dismiss request.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) dismissCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}/dismiss"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	if aggregatedAlertName == "" {
		return nil, errors.New("parameter aggregatedAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{aggregatedAlertName}", url.PathEscape(aggregatedAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Use this method to get a single the aggregated alert of yours IoT Security solution. This aggregation is performed
// by alert name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - solutionName - The name of the IoT Security solution.
//   - aggregatedAlertName - Identifier of the aggregated alert.
//   - options - IotSecuritySolutionsAnalyticsAggregatedAlertClientGetOptions contains the optional parameters for the IotSecuritySolutionsAnalyticsAggregatedAlertClient.Get
//     method.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientGetOptions) (IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, solutionName, aggregatedAlertName, options)
	if err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) getCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	if aggregatedAlertName == "" {
		return nil, errors.New("parameter aggregatedAlertName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{aggregatedAlertName}", url.PathEscape(aggregatedAlertName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) getHandleResponse(resp *http.Response) (IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse, error) {
	result := IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecurityAggregatedAlert); err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Use this method to get the aggregated alert list of yours IoT Security solution.
//
// Generated from API version 2019-08-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - solutionName - The name of the IoT Security solution.
//   - options - IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions contains the optional parameters for the IotSecuritySolutionsAnalyticsAggregatedAlertClient.NewListPager
//     method.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) NewListPager(resourceGroupName string, solutionName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions) *runtime.Pager[IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse]{
		More: func(page IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse) (IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, solutionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) listCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IotSecuritySolutionsAnalyticsAggregatedAlertClient) listHandleResponse(resp *http.Response) (IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse, error) {
	result := IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecurityAggregatedAlertList); err != nil {
		return IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse{}, err
	}
	return result, nil
}
