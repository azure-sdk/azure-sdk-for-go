//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// DatabaseEncryptionProtectorsClient contains the methods for the DatabaseEncryptionProtectors group.
// Don't use this type directly, use NewDatabaseEncryptionProtectorsClient() instead.
type DatabaseEncryptionProtectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDatabaseEncryptionProtectorsClient creates a new instance of DatabaseEncryptionProtectorsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDatabaseEncryptionProtectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseEncryptionProtectorsClient, error) {
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
	client := &DatabaseEncryptionProtectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginRevalidate - Revalidates an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// encryptionProtectorName - The name of the encryption protector to be updated.
// options - DatabaseEncryptionProtectorsClientBeginRevalidateOptions contains the optional parameters for the DatabaseEncryptionProtectorsClient.BeginRevalidate
// method.
func (client *DatabaseEncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevalidateOptions) (*runtime.Poller[DatabaseEncryptionProtectorsClientRevalidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revalidate(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabaseEncryptionProtectorsClientRevalidateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabaseEncryptionProtectorsClientRevalidateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Revalidate - Revalidates an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
func (client *DatabaseEncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *DatabaseEncryptionProtectorsClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/encryptionProtector/{encryptionProtectorName}/revalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginRevert - Reverts an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// encryptionProtectorName - The name of the encryption protector to be updated.
// options - DatabaseEncryptionProtectorsClientBeginRevertOptions contains the optional parameters for the DatabaseEncryptionProtectorsClient.BeginRevert
// method.
func (client *DatabaseEncryptionProtectorsClient) BeginRevert(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevertOptions) (*runtime.Poller[DatabaseEncryptionProtectorsClientRevertResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revert(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabaseEncryptionProtectorsClientRevertResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabaseEncryptionProtectorsClientRevertResponse](options.ResumeToken, client.pl, nil)
	}
}

// Revert - Reverts an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
func (client *DatabaseEncryptionProtectorsClient) revert(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevertOptions) (*http.Response, error) {
	req, err := client.revertCreateRequest(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revertCreateRequest creates the Revert request.
func (client *DatabaseEncryptionProtectorsClient) revertCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevertOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/encryptionProtector/{encryptionProtectorName}/revert"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
