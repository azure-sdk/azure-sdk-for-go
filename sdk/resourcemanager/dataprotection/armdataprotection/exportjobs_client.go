//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection

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

// ExportJobsClient contains the methods for the ExportJobs group.
// Don't use this type directly, use NewExportJobsClient() instead.
type ExportJobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExportJobsClient creates a new instance of ExportJobsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExportJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExportJobsClient, error) {
	cl, err := arm.NewClient(moduleName+".ExportJobsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExportJobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginTrigger - Triggers export of jobs and returns an OperationID to track.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - ExportJobsClientBeginTriggerOptions contains the optional parameters for the ExportJobsClient.BeginTrigger method.
func (client *ExportJobsClient) BeginTrigger(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsClientBeginTriggerOptions) (*runtime.Poller[ExportJobsClientTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.trigger(ctx, resourceGroupName, vaultName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExportJobsClientTriggerResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExportJobsClientTriggerResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Trigger - Triggers export of jobs and returns an OperationID to track.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *ExportJobsClient) trigger(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsClientBeginTriggerOptions) (*http.Response, error) {
	req, err := client.triggerCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *ExportJobsClient) triggerCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *ExportJobsClientBeginTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/exportBackupJobs"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
