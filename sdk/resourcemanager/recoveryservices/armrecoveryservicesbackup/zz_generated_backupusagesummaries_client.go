//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// BackupUsageSummariesClient contains the methods for the BackupUsageSummaries group.
// Don't use this type directly, use NewBackupUsageSummariesClient() instead.
type BackupUsageSummariesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupUsageSummariesClient creates a new instance of BackupUsageSummariesClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupUsageSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupUsageSummariesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &BackupUsageSummariesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Fetches the backup management usage summaries of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - BackupUsageSummariesClientListOptions contains the optional parameters for the BackupUsageSummariesClient.List
// method.
func (client *BackupUsageSummariesClient) NewListPager(vaultName string, resourceGroupName string, options *BackupUsageSummariesClientListOptions) *runtime.Pager[BackupUsageSummariesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[BackupUsageSummariesClientListResponse]{
		More: func(page BackupUsageSummariesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BackupUsageSummariesClientListResponse) (BackupUsageSummariesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
			if err != nil {
				return BackupUsageSummariesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BackupUsageSummariesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupUsageSummariesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupUsageSummariesClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupUsageSummariesClientListOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupUsageSummaries"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupUsageSummariesClient) listHandleResponse(resp *http.Response) (BackupUsageSummariesClientListResponse, error) {
	result := BackupUsageSummariesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupManagementUsageList); err != nil {
		return BackupUsageSummariesClientListResponse{}, err
	}
	return result, nil
}
