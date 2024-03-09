//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// DatabaseEncryptionProtectorsClient contains the methods for the DatabaseEncryptionProtectors group.
// Don't use this type directly, use NewDatabaseEncryptionProtectorsClient() instead.
type DatabaseEncryptionProtectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseEncryptionProtectorsClient creates a new instance of DatabaseEncryptionProtectorsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseEncryptionProtectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseEncryptionProtectorsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseEncryptionProtectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginRevalidate - Revalidates an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - encryptionProtectorName - The name of the encryption protector to be updated.
//   - options - DatabaseEncryptionProtectorsClientBeginRevalidateOptions contains the optional parameters for the DatabaseEncryptionProtectorsClient.BeginRevalidate
//     method.
func (client *DatabaseEncryptionProtectorsClient) BeginRevalidate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevalidateOptions) (*runtime.Poller[DatabaseEncryptionProtectorsClientRevalidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revalidate(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseEncryptionProtectorsClientRevalidateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseEncryptionProtectorsClientRevalidateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Revalidate - Revalidates an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
func (client *DatabaseEncryptionProtectorsClient) revalidate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevalidateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseEncryptionProtectorsClient.BeginRevalidate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRevert - Reverts an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - encryptionProtectorName - The name of the encryption protector to be updated.
//   - options - DatabaseEncryptionProtectorsClientBeginRevertOptions contains the optional parameters for the DatabaseEncryptionProtectorsClient.BeginRevert
//     method.
func (client *DatabaseEncryptionProtectorsClient) BeginRevert(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevertOptions) (*runtime.Poller[DatabaseEncryptionProtectorsClientRevertResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revert(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseEncryptionProtectorsClientRevertResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseEncryptionProtectorsClientRevertResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Revert - Reverts an existing encryption protector for a particular database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
func (client *DatabaseEncryptionProtectorsClient) revert(ctx context.Context, resourceGroupName string, serverName string, databaseName string, encryptionProtectorName EncryptionProtectorName, options *DatabaseEncryptionProtectorsClientBeginRevertOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseEncryptionProtectorsClient.BeginRevert"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revertCreateRequest(ctx, resourceGroupName, serverName, databaseName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
