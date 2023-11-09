//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ExportJobsOperationResultClient contains the methods for the ExportJobsOperationResult group.
// Don't use this type directly, use NewExportJobsOperationResultClient() instead.
type ExportJobsOperationResultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExportJobsOperationResultClient creates a new instance of ExportJobsOperationResultClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExportJobsOperationResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExportJobsOperationResultClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExportJobsOperationResultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the operation result of operation triggered by Export Jobs API. If the operation is successful, then it also
// contains URL of a Blob and a SAS key to access the same. The blob contains exported
// jobs in JSON serialized format.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - operationID - OperationID which represents the export job.
//   - options - ExportJobsOperationResultClientGetOptions contains the optional parameters for the ExportJobsOperationResultClient.Get
//     method.
func (client *ExportJobsOperationResultClient) Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *ExportJobsOperationResultClientGetOptions) (ExportJobsOperationResultClientGetResponse, error) {
	var err error
	const operationName = "ExportJobsOperationResultClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExportJobsOperationResultClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *ExportJobsOperationResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupJobs/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExportJobsOperationResultClient) getHandleResponse(resp *http.Response) (ExportJobsOperationResultClientGetResponse, error) {
	result := ExportJobsOperationResultClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportJobsResult); err != nil {
		return ExportJobsOperationResultClientGetResponse{}, err
	}
	return result, nil
}
