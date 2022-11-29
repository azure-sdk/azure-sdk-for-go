//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// LogAnalyticsClient contains the methods for the LogAnalytics group.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLogAnalyticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogAnalyticsClient, error) {
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
	client := &LogAnalyticsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window
// to show throttling activities.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// location - The location upon which virtual-machine-sizes is queried.
// parameters - Parameters supplied to the LogAnalytics getRequestRateByInterval Api.
// options - LogAnalyticsClientBeginExportRequestRateByIntervalOptions contains the optional parameters for the LogAnalyticsClient.BeginExportRequestRateByInterval
// method.
func (client *LogAnalyticsClient) BeginExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (*runtime.Poller[LogAnalyticsClientExportRequestRateByIntervalResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportRequestRateByInterval(ctx, location, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LogAnalyticsClientExportRequestRateByIntervalResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LogAnalyticsClientExportRequestRateByIntervalResponse](options.ResumeToken, client.pl, nil)
	}
}

// ExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to
// show throttling activities.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
func (client *LogAnalyticsClient) exportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (*http.Response, error) {
	req, err := client.exportRequestRateByIntervalCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// exportRequestRateByIntervalCreateRequest creates the ExportRequestRateByInterval request.
func (client *LogAnalyticsClient) exportRequestRateByIntervalCreateRequest(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getRequestRateByInterval"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time
// window.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// location - The location upon which virtual-machine-sizes is queried.
// parameters - Parameters supplied to the LogAnalytics getThrottledRequests Api.
// options - LogAnalyticsClientBeginExportThrottledRequestsOptions contains the optional parameters for the LogAnalyticsClient.BeginExportThrottledRequests
// method.
func (client *LogAnalyticsClient) BeginExportThrottledRequests(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsClientBeginExportThrottledRequestsOptions) (*runtime.Poller[LogAnalyticsClientExportThrottledRequestsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportThrottledRequests(ctx, location, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LogAnalyticsClientExportThrottledRequestsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LogAnalyticsClientExportThrottledRequestsResponse](options.ResumeToken, client.pl, nil)
	}
}

// ExportThrottledRequests - Export logs that show total throttled Api requests for this subscription in the given time window.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
func (client *LogAnalyticsClient) exportThrottledRequests(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsClientBeginExportThrottledRequestsOptions) (*http.Response, error) {
	req, err := client.exportThrottledRequestsCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// exportThrottledRequestsCreateRequest creates the ExportThrottledRequests request.
func (client *LogAnalyticsClient) exportThrottledRequestsCreateRequest(ctx context.Context, location string, parameters ThrottledRequestsInput, options *LogAnalyticsClientBeginExportThrottledRequestsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/logAnalytics/apiAccess/getThrottledRequests"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
