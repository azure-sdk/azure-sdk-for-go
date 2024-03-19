//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// LogAnalyticsClient contains the methods for the LogAnalytics group.
// Don't use this type directly, use NewLogAnalyticsClient() instead.
type LogAnalyticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLogAnalyticsClient creates a new instance of LogAnalyticsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLogAnalyticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogAnalyticsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LogAnalyticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window
// to show throttling activities.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - location - The location upon which virtual-machine-sizes is queried.
//   - parameters - Parameters supplied to the LogAnalytics getRequestRateByInterval Api.
//   - options - LogAnalyticsClientBeginExportRequestRateByIntervalOptions contains the optional parameters for the LogAnalyticsClient.BeginExportRequestRateByInterval
//     method.
func (client *LogAnalyticsClient) BeginExportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (*runtime.Poller[LogAnalyticsClientExportRequestRateByIntervalResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportRequestRateByInterval(ctx, location, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LogAnalyticsClientExportRequestRateByIntervalResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LogAnalyticsClientExportRequestRateByIntervalResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportRequestRateByInterval - Export logs that show Api requests made by this subscription in the given time window to
// show throttling activities.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *LogAnalyticsClient) exportRequestRateByInterval(ctx context.Context, location string, parameters RequestRateByIntervalInput, options *LogAnalyticsClientBeginExportRequestRateByIntervalOptions) (*http.Response, error) {
	var err error
	const operationName = "LogAnalyticsClient.BeginExportRequestRateByInterval"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportRequestRateByIntervalCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
