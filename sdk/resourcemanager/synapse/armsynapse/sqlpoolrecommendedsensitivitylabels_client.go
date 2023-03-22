//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// SQLPoolRecommendedSensitivityLabelsClient contains the methods for the SQLPoolRecommendedSensitivityLabels group.
// Don't use this type directly, use NewSQLPoolRecommendedSensitivityLabelsClient() instead.
type SQLPoolRecommendedSensitivityLabelsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLPoolRecommendedSensitivityLabelsClient creates a new instance of SQLPoolRecommendedSensitivityLabelsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLPoolRecommendedSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolRecommendedSensitivityLabelsClient, error) {
	cl, err := arm.NewClient(moduleName+".SQLPoolRecommendedSensitivityLabelsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLPoolRecommendedSensitivityLabelsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Update - Update recommended sensitivity labels states of a given SQL Pool using an operations batch.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sqlPoolName - SQL pool name
//   - options - SQLPoolRecommendedSensitivityLabelsClientUpdateOptions contains the optional parameters for the SQLPoolRecommendedSensitivityLabelsClient.Update
//     method.
func (client *SQLPoolRecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters RecommendedSensitivityLabelUpdateList, options *SQLPoolRecommendedSensitivityLabelsClientUpdateOptions) (SQLPoolRecommendedSensitivityLabelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, parameters, options)
	if err != nil {
		return SQLPoolRecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLPoolRecommendedSensitivityLabelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolRecommendedSensitivityLabelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return SQLPoolRecommendedSensitivityLabelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *SQLPoolRecommendedSensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, parameters RecommendedSensitivityLabelUpdateList, options *SQLPoolRecommendedSensitivityLabelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/recommendedSensitivityLabels"
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
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}
