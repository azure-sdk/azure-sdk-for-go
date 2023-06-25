//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights

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

// WorkspacePurgeClient contains the methods for the WorkspacePurge group.
// Don't use this type directly, use NewWorkspacePurgeClient() instead.
type WorkspacePurgeClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspacePurgeClient creates a new instance of WorkspacePurgeClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspacePurgeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspacePurgeClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspacePurgeClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspacePurgeClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetPurgeStatus - Gets status of an ongoing purge operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - purgeID - In a purge status request, this is the Id of the operation the status of which is returned.
//   - options - WorkspacePurgeClientGetPurgeStatusOptions contains the optional parameters for the WorkspacePurgeClient.GetPurgeStatus
//     method.
func (client *WorkspacePurgeClient) GetPurgeStatus(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string, options *WorkspacePurgeClientGetPurgeStatusOptions) (WorkspacePurgeClientGetPurgeStatusResponse, error) {
	req, err := client.getPurgeStatusCreateRequest(ctx, resourceGroupName, workspaceName, purgeID, options)
	if err != nil {
		return WorkspacePurgeClientGetPurgeStatusResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspacePurgeClientGetPurgeStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacePurgeClientGetPurgeStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPurgeStatusHandleResponse(resp)
}

// getPurgeStatusCreateRequest creates the GetPurgeStatus request.
func (client *WorkspacePurgeClient) getPurgeStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string, options *WorkspacePurgeClientGetPurgeStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if purgeID == "" {
		return nil, errors.New("parameter purgeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{purgeId}", url.PathEscape(purgeID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPurgeStatusHandleResponse handles the GetPurgeStatus response.
func (client *WorkspacePurgeClient) getPurgeStatusHandleResponse(resp *http.Response) (WorkspacePurgeClientGetPurgeStatusResponse, error) {
	result := WorkspacePurgeClientGetPurgeStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspacePurgeStatusResponse); err != nil {
		return WorkspacePurgeClientGetPurgeStatusResponse{}, err
	}
	return result, nil
}

// Purge - Purges data in an Log Analytics workspace by a set of user-defined filters.
// In order to manage system resources, purge requests are throttled at 50 requests per hour. You should batch the execution
// of purge requests by sending a single command whose predicate includes all
// user identities that require purging. Use the in operator to specify multiple identities. You should run the query prior
// to using for a purge request to verify that the results are expected. Log
// Analytics only supports purge operations required for compliance with GDPR. The Log Analytics product team reserves the
// right to reject requests for purge operations that are not for the purpose of
// GDPR compliance. In the event of a dispute, please create a support ticket
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - body - Describes the body of a request to purge data in a single table of an Log Analytics Workspace
//   - options - WorkspacePurgeClientPurgeOptions contains the optional parameters for the WorkspacePurgeClient.Purge method.
func (client *WorkspacePurgeClient) Purge(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody, options *WorkspacePurgeClientPurgeOptions) (WorkspacePurgeClientPurgeResponse, error) {
	req, err := client.purgeCreateRequest(ctx, resourceGroupName, workspaceName, body, options)
	if err != nil {
		return WorkspacePurgeClientPurgeResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspacePurgeClientPurgeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return WorkspacePurgeClientPurgeResponse{}, runtime.NewResponseError(resp)
	}
	return client.purgeHandleResponse(resp)
}

// purgeCreateRequest creates the Purge request.
func (client *WorkspacePurgeClient) purgeCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody, options *WorkspacePurgeClientPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/purge"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// purgeHandleResponse handles the Purge response.
func (client *WorkspacePurgeClient) purgeHandleResponse(resp *http.Response) (WorkspacePurgeClientPurgeResponse, error) {
	result := WorkspacePurgeClientPurgeResponse{}
	if val := resp.Header.Get("x-ms-status-location"); val != "" {
		result.XMSStatusLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspacePurgeResponse); err != nil {
		return WorkspacePurgeClientPurgeResponse{}, err
	}
	return result, nil
}
