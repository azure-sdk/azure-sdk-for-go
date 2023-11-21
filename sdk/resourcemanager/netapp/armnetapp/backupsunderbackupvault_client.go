//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// BackupsUnderBackupVaultClient contains the methods for the BackupsUnderBackupVault group.
// Don't use this type directly, use NewBackupsUnderBackupVaultClient() instead.
type BackupsUnderBackupVaultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsUnderBackupVaultClient creates a new instance of BackupsUnderBackupVaultClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsUnderBackupVaultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsUnderBackupVaultClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsUnderBackupVaultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginRestoreFiles - Restore the specified files from the specified backup to the active filesystem
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - backupName - The name of the backup
//   - body - Restore payload supplied in the body of the operation.
//   - options - BackupsUnderBackupVaultClientBeginRestoreFilesOptions contains the optional parameters for the BackupsUnderBackupVaultClient.BeginRestoreFiles
//     method.
func (client *BackupsUnderBackupVaultClient) BeginRestoreFiles(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupRestoreFiles, options *BackupsUnderBackupVaultClientBeginRestoreFilesOptions) (*runtime.Poller[BackupsUnderBackupVaultClientRestoreFilesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restoreFiles(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsUnderBackupVaultClientRestoreFilesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsUnderBackupVaultClientRestoreFilesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RestoreFiles - Restore the specified files from the specified backup to the active filesystem
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *BackupsUnderBackupVaultClient) restoreFiles(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupRestoreFiles, options *BackupsUnderBackupVaultClientBeginRestoreFilesOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsUnderBackupVaultClient.BeginRestoreFiles"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restoreFilesCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
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

// restoreFilesCreateRequest creates the RestoreFiles request.
func (client *BackupsUnderBackupVaultClient) restoreFilesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupRestoreFiles, options *BackupsUnderBackupVaultClientBeginRestoreFilesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups/{backupName}/restoreFiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
