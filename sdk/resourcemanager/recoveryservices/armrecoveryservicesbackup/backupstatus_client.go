//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

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

// BackupStatusClient contains the methods for the BackupStatus group.
// Don't use this type directly, use NewBackupStatusClient() instead.
type BackupStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupStatusClient creates a new instance of BackupStatusClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupStatusClient, error) {
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
	client := &BackupStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the container backup status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - azureRegion - Azure region to hit Api
//   - parameters - Container Backup Status Request
//   - options - BackupStatusClientGetOptions contains the optional parameters for the BackupStatusClient.Get method.
func (client *BackupStatusClient) Get(ctx context.Context, azureRegion string, parameters BackupStatusRequest, options *BackupStatusClientGetOptions) (BackupStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, azureRegion, parameters, options)
	if err != nil {
		return BackupStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupStatusClient) getCreateRequest(ctx context.Context, azureRegion string, parameters BackupStatusRequest, options *BackupStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{azureRegion}/backupStatus"
	if azureRegion == "" {
		return nil, errors.New("parameter azureRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureRegion}", url.PathEscape(azureRegion))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// getHandleResponse handles the Get response.
func (client *BackupStatusClient) getHandleResponse(resp *http.Response) (BackupStatusClientGetResponse, error) {
	result := BackupStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupStatusResponse); err != nil {
		return BackupStatusClientGetResponse{}, err
	}
	return result, nil
}
