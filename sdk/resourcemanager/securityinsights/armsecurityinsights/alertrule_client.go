//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// AlertRuleClient contains the methods for the AlertRule group.
// Don't use this type directly, use NewAlertRuleClient() instead.
type AlertRuleClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAlertRuleClient creates a new instance of AlertRuleClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertRuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRuleClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertRuleClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginTriggerRuleRun - triggers analytics rule run
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - ruleID - Alert rule ID
//   - analyticsRuleRunTriggerParameter - The Analytics Rule Run Trigger parameter
//   - options - AlertRuleClientBeginTriggerRuleRunOptions contains the optional parameters for the AlertRuleClient.BeginTriggerRuleRun
//     method.
func (client *AlertRuleClient) BeginTriggerRuleRun(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, analyticsRuleRunTriggerParameter AnalyticsRuleRunTrigger, options *AlertRuleClientBeginTriggerRuleRunOptions) (*runtime.Poller[AlertRuleClientTriggerRuleRunResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.triggerRuleRun(ctx, resourceGroupName, workspaceName, ruleID, analyticsRuleRunTriggerParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AlertRuleClientTriggerRuleRunResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AlertRuleClientTriggerRuleRunResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// TriggerRuleRun - triggers analytics rule run
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
func (client *AlertRuleClient) triggerRuleRun(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, analyticsRuleRunTriggerParameter AnalyticsRuleRunTrigger, options *AlertRuleClientBeginTriggerRuleRunOptions) (*http.Response, error) {
	var err error
	const operationName = "AlertRuleClient.BeginTriggerRuleRun"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.triggerRuleRunCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, analyticsRuleRunTriggerParameter, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// triggerRuleRunCreateRequest creates the TriggerRuleRun request.
func (client *AlertRuleClient) triggerRuleRunCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, analyticsRuleRunTriggerParameter AnalyticsRuleRunTrigger, options *AlertRuleClientBeginTriggerRuleRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/triggerRuleRun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, analyticsRuleRunTriggerParameter); err != nil {
		return nil, err
	}
	return req, nil
}
