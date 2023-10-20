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

// TriggeredAnalyticsRuleRunClient contains the methods for the TriggeredAnalyticsRuleRun group.
// Don't use this type directly, use NewTriggeredAnalyticsRuleRunClient() instead.
type TriggeredAnalyticsRuleRunClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTriggeredAnalyticsRuleRunClient creates a new instance of TriggeredAnalyticsRuleRunClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTriggeredAnalyticsRuleRunClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TriggeredAnalyticsRuleRunClient, error) {
	cl, err := arm.NewClient(moduleName+".TriggeredAnalyticsRuleRunClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TriggeredAnalyticsRuleRunClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the triggered analytics rule run.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - ruleRunID - the triggered rule id
//   - options - TriggeredAnalyticsRuleRunClientGetOptions contains the optional parameters for the TriggeredAnalyticsRuleRunClient.Get
//     method.
func (client *TriggeredAnalyticsRuleRunClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleRunID string, options *TriggeredAnalyticsRuleRunClientGetOptions) (TriggeredAnalyticsRuleRunClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, ruleRunID, options)
	if err != nil {
		return TriggeredAnalyticsRuleRunClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TriggeredAnalyticsRuleRunClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TriggeredAnalyticsRuleRunClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TriggeredAnalyticsRuleRunClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleRunID string, options *TriggeredAnalyticsRuleRunClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/triggeredAnalyticsRuleRuns/{ruleRunId}"
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
	if ruleRunID == "" {
		return nil, errors.New("parameter ruleRunID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleRunId}", url.PathEscape(ruleRunID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TriggeredAnalyticsRuleRunClient) getHandleResponse(resp *http.Response) (TriggeredAnalyticsRuleRunClientGetResponse, error) {
	result := TriggeredAnalyticsRuleRunClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggeredAnalyticsRuleRun); err != nil {
		return TriggeredAnalyticsRuleRunClientGetResponse{}, err
	}
	return result, nil
}
