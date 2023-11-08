//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// LogAnalyticsClient contains the methods for the LogAnalytics group.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLogAnalyticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogAnalyticsClient, error) {
	cl, err := arm.NewClient(moduleName+".LogAnalyticsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LogAnalyticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetLogAnalyticsLocations - Get all available location names for AFD log analytics report.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetLogAnalyticsLocationsOptions contains the optional parameters for the LogAnalyticsClient.GetLogAnalyticsLocations
//     method.
func (client *LogAnalyticsClient) GetLogAnalyticsLocations(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsClientGetLogAnalyticsLocationsOptions) (LogAnalyticsClientGetLogAnalyticsLocationsResponse, error) {
	var err error
	req, err := client.getLogAnalyticsLocationsCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsLocationsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsLocationsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetLogAnalyticsLocationsResponse{}, err
	}
	resp, err := client.getLogAnalyticsLocationsHandleResponse(httpResp)
	return resp, err
}

// getLogAnalyticsLocationsCreateRequest creates the GetLogAnalyticsLocations request.
func (client *LogAnalyticsClient) getLogAnalyticsLocationsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsClientGetLogAnalyticsLocationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogAnalyticsLocationsHandleResponse handles the GetLogAnalyticsLocations response.
func (client *LogAnalyticsClient) getLogAnalyticsLocationsHandleResponse(resp *http.Response) (LogAnalyticsClientGetLogAnalyticsLocationsResponse, error) {
	result := LogAnalyticsClientGetLogAnalyticsLocationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContinentsResponse); err != nil {
		return LogAnalyticsClientGetLogAnalyticsLocationsResponse{}, err
	}
	return result, nil
}

// GetLogAnalyticsMetrics - Get log report for AFD profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetLogAnalyticsMetricsOptions contains the optional parameters for the LogAnalyticsClient.GetLogAnalyticsMetrics
//     method.
func (client *LogAnalyticsClient) GetLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []LogMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity LogMetricsGranularity, customDomains []string, protocols []string, options *LogAnalyticsClientGetLogAnalyticsMetricsOptions) (LogAnalyticsClientGetLogAnalyticsMetricsResponse, error) {
	var err error
	req, err := client.getLogAnalyticsMetricsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, granularity, customDomains, protocols, options)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsMetricsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetLogAnalyticsMetricsResponse{}, err
	}
	resp, err := client.getLogAnalyticsMetricsHandleResponse(httpResp)
	return resp, err
}

// getLogAnalyticsMetricsCreateRequest creates the GetLogAnalyticsMetrics request.
func (client *LogAnalyticsClient) getLogAnalyticsMetricsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []LogMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity LogMetricsGranularity, customDomains []string, protocols []string, options *LogAnalyticsClientGetLogAnalyticsMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", string(qv))
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("granularity", string(granularity))
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	if options != nil && options.Continents != nil {
		for _, qv := range options.Continents {
			reqQP.Add("continents", qv)
		}
	}
	if options != nil && options.CountryOrRegions != nil {
		for _, qv := range options.CountryOrRegions {
			reqQP.Add("countryOrRegions", qv)
		}
	}
	for _, qv := range customDomains {
		reqQP.Add("customDomains", qv)
	}
	for _, qv := range protocols {
		reqQP.Add("protocols", qv)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogAnalyticsMetricsHandleResponse handles the GetLogAnalyticsMetrics response.
func (client *LogAnalyticsClient) getLogAnalyticsMetricsHandleResponse(resp *http.Response) (LogAnalyticsClientGetLogAnalyticsMetricsResponse, error) {
	result := LogAnalyticsClientGetLogAnalyticsMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricsResponse); err != nil {
		return LogAnalyticsClientGetLogAnalyticsMetricsResponse{}, err
	}
	return result, nil
}

// GetLogAnalyticsRankings - Get log analytics ranking report for AFD profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetLogAnalyticsRankingsOptions contains the optional parameters for the LogAnalyticsClient.GetLogAnalyticsRankings
//     method.
func (client *LogAnalyticsClient) GetLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, rankings []LogRanking, metrics []LogRankingMetric, maxRanking int32, dateTimeBegin time.Time, dateTimeEnd time.Time, options *LogAnalyticsClientGetLogAnalyticsRankingsOptions) (LogAnalyticsClientGetLogAnalyticsRankingsResponse, error) {
	var err error
	req, err := client.getLogAnalyticsRankingsCreateRequest(ctx, resourceGroupName, profileName, rankings, metrics, maxRanking, dateTimeBegin, dateTimeEnd, options)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsRankingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsRankingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetLogAnalyticsRankingsResponse{}, err
	}
	resp, err := client.getLogAnalyticsRankingsHandleResponse(httpResp)
	return resp, err
}

// getLogAnalyticsRankingsCreateRequest creates the GetLogAnalyticsRankings request.
func (client *LogAnalyticsClient) getLogAnalyticsRankingsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, rankings []LogRanking, metrics []LogRankingMetric, maxRanking int32, dateTimeBegin time.Time, dateTimeEnd time.Time, options *LogAnalyticsClientGetLogAnalyticsRankingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsRankings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	for _, qv := range rankings {
		reqQP.Add("rankings", string(qv))
	}
	for _, qv := range metrics {
		reqQP.Add("metrics", string(qv))
	}
	reqQP.Set("maxRanking", strconv.FormatInt(int64(maxRanking), 10))
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	if options != nil && options.CustomDomains != nil {
		for _, qv := range options.CustomDomains {
			reqQP.Add("customDomains", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogAnalyticsRankingsHandleResponse handles the GetLogAnalyticsRankings response.
func (client *LogAnalyticsClient) getLogAnalyticsRankingsHandleResponse(resp *http.Response) (LogAnalyticsClientGetLogAnalyticsRankingsResponse, error) {
	result := LogAnalyticsClientGetLogAnalyticsRankingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RankingsResponse); err != nil {
		return LogAnalyticsClientGetLogAnalyticsRankingsResponse{}, err
	}
	return result, nil
}

// GetLogAnalyticsResources - Get all endpoints and custom domains available for AFD log report
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetLogAnalyticsResourcesOptions contains the optional parameters for the LogAnalyticsClient.GetLogAnalyticsResources
//     method.
func (client *LogAnalyticsClient) GetLogAnalyticsResources(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsClientGetLogAnalyticsResourcesOptions) (LogAnalyticsClientGetLogAnalyticsResourcesResponse, error) {
	var err error
	req, err := client.getLogAnalyticsResourcesCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsResourcesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetLogAnalyticsResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetLogAnalyticsResourcesResponse{}, err
	}
	resp, err := client.getLogAnalyticsResourcesHandleResponse(httpResp)
	return resp, err
}

// getLogAnalyticsResourcesCreateRequest creates the GetLogAnalyticsResources request.
func (client *LogAnalyticsClient) getLogAnalyticsResourcesCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *LogAnalyticsClientGetLogAnalyticsResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getLogAnalyticsResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogAnalyticsResourcesHandleResponse handles the GetLogAnalyticsResources response.
func (client *LogAnalyticsClient) getLogAnalyticsResourcesHandleResponse(resp *http.Response) (LogAnalyticsClientGetLogAnalyticsResourcesResponse, error) {
	result := LogAnalyticsClientGetLogAnalyticsResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcesResponse); err != nil {
		return LogAnalyticsClientGetLogAnalyticsResourcesResponse{}, err
	}
	return result, nil
}

// GetWafLogAnalyticsMetrics - Get Waf related log analytics report for AFD profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetWafLogAnalyticsMetricsOptions contains the optional parameters for the LogAnalyticsClient.GetWafLogAnalyticsMetrics
//     method.
func (client *LogAnalyticsClient) GetWafLogAnalyticsMetrics(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity WafGranularity, options *LogAnalyticsClientGetWafLogAnalyticsMetricsOptions) (LogAnalyticsClientGetWafLogAnalyticsMetricsResponse, error) {
	var err error
	req, err := client.getWafLogAnalyticsMetricsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, granularity, options)
	if err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{}, err
	}
	resp, err := client.getWafLogAnalyticsMetricsHandleResponse(httpResp)
	return resp, err
}

// getWafLogAnalyticsMetricsCreateRequest creates the GetWafLogAnalyticsMetrics request.
func (client *LogAnalyticsClient) getWafLogAnalyticsMetricsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, granularity WafGranularity, options *LogAnalyticsClientGetWafLogAnalyticsMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getWafLogAnalyticsMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", string(qv))
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("granularity", string(granularity))
	if options != nil && options.Actions != nil {
		for _, qv := range options.Actions {
			reqQP.Add("actions", string(qv))
		}
	}
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	if options != nil && options.RuleTypes != nil {
		for _, qv := range options.RuleTypes {
			reqQP.Add("ruleTypes", string(qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWafLogAnalyticsMetricsHandleResponse handles the GetWafLogAnalyticsMetrics response.
func (client *LogAnalyticsClient) getWafLogAnalyticsMetricsHandleResponse(resp *http.Response) (LogAnalyticsClientGetWafLogAnalyticsMetricsResponse, error) {
	result := LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WafMetricsResponse); err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsMetricsResponse{}, err
	}
	return result, nil
}

// GetWafLogAnalyticsRankings - Get WAF log analytics charts for AFD profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
//     group. which is unique within the resource group.
//   - options - LogAnalyticsClientGetWafLogAnalyticsRankingsOptions contains the optional parameters for the LogAnalyticsClient.GetWafLogAnalyticsRankings
//     method.
func (client *LogAnalyticsClient) GetWafLogAnalyticsRankings(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, maxRanking int32, rankings []WafRankingType, options *LogAnalyticsClientGetWafLogAnalyticsRankingsOptions) (LogAnalyticsClientGetWafLogAnalyticsRankingsResponse, error) {
	var err error
	req, err := client.getWafLogAnalyticsRankingsCreateRequest(ctx, resourceGroupName, profileName, metrics, dateTimeBegin, dateTimeEnd, maxRanking, rankings, options)
	if err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{}, err
	}
	resp, err := client.getWafLogAnalyticsRankingsHandleResponse(httpResp)
	return resp, err
}

// getWafLogAnalyticsRankingsCreateRequest creates the GetWafLogAnalyticsRankings request.
func (client *LogAnalyticsClient) getWafLogAnalyticsRankingsCreateRequest(ctx context.Context, resourceGroupName string, profileName string, metrics []WafMetric, dateTimeBegin time.Time, dateTimeEnd time.Time, maxRanking int32, rankings []WafRankingType, options *LogAnalyticsClientGetWafLogAnalyticsRankingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/getWafLogAnalyticsRankings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	for _, qv := range metrics {
		reqQP.Add("metrics", string(qv))
	}
	reqQP.Set("dateTimeBegin", dateTimeBegin.Format(time.RFC3339Nano))
	reqQP.Set("dateTimeEnd", dateTimeEnd.Format(time.RFC3339Nano))
	reqQP.Set("maxRanking", strconv.FormatInt(int64(maxRanking), 10))
	for _, qv := range rankings {
		reqQP.Add("rankings", string(qv))
	}
	if options != nil && options.Actions != nil {
		for _, qv := range options.Actions {
			reqQP.Add("actions", string(qv))
		}
	}
	if options != nil && options.RuleTypes != nil {
		for _, qv := range options.RuleTypes {
			reqQP.Add("ruleTypes", string(qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWafLogAnalyticsRankingsHandleResponse handles the GetWafLogAnalyticsRankings response.
func (client *LogAnalyticsClient) getWafLogAnalyticsRankingsHandleResponse(resp *http.Response) (LogAnalyticsClientGetWafLogAnalyticsRankingsResponse, error) {
	result := LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WafRankingsResponse); err != nil {
		return LogAnalyticsClientGetWafLogAnalyticsRankingsResponse{}, err
	}
	return result, nil
}
