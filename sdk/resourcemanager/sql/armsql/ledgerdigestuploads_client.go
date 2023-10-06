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

// LedgerDigestUploadsClient contains the methods for the LedgerDigestUploads group.
// Don't use this type directly, use NewLedgerDigestUploadsClient() instead.
type LedgerDigestUploadsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLedgerDigestUploadsClient creates a new instance of LedgerDigestUploadsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLedgerDigestUploadsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LedgerDigestUploadsClient, error) {
	cl, err := arm.NewClient(moduleName+".LedgerDigestUploadsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LedgerDigestUploadsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Enables upload ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - LedgerDigestUploadsClientBeginCreateOrUpdateOptions contains the optional parameters for the LedgerDigestUploadsClient.BeginCreateOrUpdate
//     method.
func (client *LedgerDigestUploadsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, parameters LedgerDigestUploads, options *LedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LedgerDigestUploadsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, databaseName, ledgerDigestUploads, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LedgerDigestUploadsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LedgerDigestUploadsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Enables upload ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *LedgerDigestUploadsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, parameters LedgerDigestUploads, options *LedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, ledgerDigestUploads, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LedgerDigestUploadsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, parameters LedgerDigestUploads, options *LedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}"
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
	if ledgerDigestUploads == "" {
		return nil, errors.New("parameter ledgerDigestUploads cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerDigestUploads}", url.PathEscape(string(ledgerDigestUploads)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDisable - Disables uploading ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - LedgerDigestUploadsClientBeginDisableOptions contains the optional parameters for the LedgerDigestUploadsClient.BeginDisable
//     method.
func (client *LedgerDigestUploadsClient) BeginDisable(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, options *LedgerDigestUploadsClientBeginDisableOptions) (*runtime.Poller[LedgerDigestUploadsClientDisableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disable(ctx, resourceGroupName, serverName, databaseName, ledgerDigestUploads, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[LedgerDigestUploadsClientDisableResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LedgerDigestUploadsClientDisableResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Disable - Disables uploading ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *LedgerDigestUploadsClient) disable(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, options *LedgerDigestUploadsClientBeginDisableOptions) (*http.Response, error) {
	var err error
	req, err := client.disableCreateRequest(ctx, resourceGroupName, serverName, databaseName, ledgerDigestUploads, options)
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

// disableCreateRequest creates the Disable request.
func (client *LedgerDigestUploadsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, options *LedgerDigestUploadsClientBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}/disable"
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
	if ledgerDigestUploads == "" {
		return nil, errors.New("parameter ledgerDigestUploads cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerDigestUploads}", url.PathEscape(string(ledgerDigestUploads)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the current ledger digest upload configuration for a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - LedgerDigestUploadsClientGetOptions contains the optional parameters for the LedgerDigestUploadsClient.Get method.
func (client *LedgerDigestUploadsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, options *LedgerDigestUploadsClientGetOptions) (LedgerDigestUploadsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, ledgerDigestUploads, options)
	if err != nil {
		return LedgerDigestUploadsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LedgerDigestUploadsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LedgerDigestUploadsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LedgerDigestUploadsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, ledgerDigestUploads LedgerDigestUploadsName, options *LedgerDigestUploadsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}"
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
	if ledgerDigestUploads == "" {
		return nil, errors.New("parameter ledgerDigestUploads cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerDigestUploads}", url.PathEscape(string(ledgerDigestUploads)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LedgerDigestUploadsClient) getHandleResponse(resp *http.Response) (LedgerDigestUploadsClientGetResponse, error) {
	result := LedgerDigestUploadsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LedgerDigestUploads); err != nil {
		return LedgerDigestUploadsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets all ledger digest upload settings on a database.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - LedgerDigestUploadsClientListByDatabaseOptions contains the optional parameters for the LedgerDigestUploadsClient.NewListByDatabasePager
//     method.
func (client *LedgerDigestUploadsClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *LedgerDigestUploadsClientListByDatabaseOptions) *runtime.Pager[LedgerDigestUploadsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[LedgerDigestUploadsClientListByDatabaseResponse]{
		More: func(page LedgerDigestUploadsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LedgerDigestUploadsClientListByDatabaseResponse) (LedgerDigestUploadsClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LedgerDigestUploadsClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LedgerDigestUploadsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LedgerDigestUploadsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *LedgerDigestUploadsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *LedgerDigestUploadsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/ledgerDigestUploads"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *LedgerDigestUploadsClient) listByDatabaseHandleResponse(resp *http.Response) (LedgerDigestUploadsClientListByDatabaseResponse, error) {
	result := LedgerDigestUploadsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LedgerDigestUploadsListResult); err != nil {
		return LedgerDigestUploadsClientListByDatabaseResponse{}, err
	}
	return result, nil
}
