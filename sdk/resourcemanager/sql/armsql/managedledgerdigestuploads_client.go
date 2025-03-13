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

// ManagedLedgerDigestUploadsClient contains the methods for the ManagedLedgerDigestUploads group.
// Don't use this type directly, use NewManagedLedgerDigestUploadsClient() instead.
type ManagedLedgerDigestUploadsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedLedgerDigestUploadsClient creates a new instance of ManagedLedgerDigestUploadsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedLedgerDigestUploadsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedLedgerDigestUploadsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedLedgerDigestUploadsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Enables upload ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - ledgerDigestUploads - The name of the Ledger Digest Upload Configurations.
//   - parameters - The Ledger Digest Storage Endpoint.
//   - options - ManagedLedgerDigestUploadsClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedLedgerDigestUploadsClient.BeginCreateOrUpdate
//     method.
func (client *ManagedLedgerDigestUploadsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, parameters ManagedLedgerDigestUploads, options *ManagedLedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedLedgerDigestUploadsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedLedgerDigestUploadsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedLedgerDigestUploadsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Enables upload ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *ManagedLedgerDigestUploadsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, parameters ManagedLedgerDigestUploads, options *ManagedLedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedLedgerDigestUploadsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads, parameters, options)
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
func (client *ManagedLedgerDigestUploadsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, parameters ManagedLedgerDigestUploads, _ *ManagedLedgerDigestUploadsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2022-08-01-preview")
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
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedLedgerDigestUploadsClientBeginDisableOptions contains the optional parameters for the ManagedLedgerDigestUploadsClient.BeginDisable
//     method.
func (client *ManagedLedgerDigestUploadsClient) BeginDisable(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, options *ManagedLedgerDigestUploadsClientBeginDisableOptions) (*runtime.Poller[ManagedLedgerDigestUploadsClientDisableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disable(ctx, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedLedgerDigestUploadsClientDisableResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedLedgerDigestUploadsClientDisableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Disable - Disables uploading ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *ManagedLedgerDigestUploadsClient) disable(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, options *ManagedLedgerDigestUploadsClientBeginDisableOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedLedgerDigestUploadsClient.BeginDisable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disableCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads, options)
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
func (client *ManagedLedgerDigestUploadsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, _ *ManagedLedgerDigestUploadsClientBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}/disable"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the current ledger digest upload configuration for a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedLedgerDigestUploadsClientGetOptions contains the optional parameters for the ManagedLedgerDigestUploadsClient.Get
//     method.
func (client *ManagedLedgerDigestUploadsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, options *ManagedLedgerDigestUploadsClientGetOptions) (ManagedLedgerDigestUploadsClientGetResponse, error) {
	var err error
	const operationName = "ManagedLedgerDigestUploadsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, ledgerDigestUploads, options)
	if err != nil {
		return ManagedLedgerDigestUploadsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedLedgerDigestUploadsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedLedgerDigestUploadsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedLedgerDigestUploadsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, ledgerDigestUploads ManagedLedgerDigestUploadsName, _ *ManagedLedgerDigestUploadsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/ledgerDigestUploads/{ledgerDigestUploads}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedLedgerDigestUploadsClient) getHandleResponse(resp *http.Response) (ManagedLedgerDigestUploadsClientGetResponse, error) {
	result := ManagedLedgerDigestUploadsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedLedgerDigestUploads); err != nil {
		return ManagedLedgerDigestUploadsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets all ledger digest upload settings on a database.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedLedgerDigestUploadsClientListByDatabaseOptions contains the optional parameters for the ManagedLedgerDigestUploadsClient.NewListByDatabasePager
//     method.
func (client *ManagedLedgerDigestUploadsClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedLedgerDigestUploadsClientListByDatabaseOptions) *runtime.Pager[ManagedLedgerDigestUploadsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedLedgerDigestUploadsClientListByDatabaseResponse]{
		More: func(page ManagedLedgerDigestUploadsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedLedgerDigestUploadsClientListByDatabaseResponse) (ManagedLedgerDigestUploadsClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedLedgerDigestUploadsClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			}, nil)
			if err != nil {
				return ManagedLedgerDigestUploadsClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedLedgerDigestUploadsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, _ *ManagedLedgerDigestUploadsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/ledgerDigestUploads"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedLedgerDigestUploadsClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedLedgerDigestUploadsClientListByDatabaseResponse, error) {
	result := ManagedLedgerDigestUploadsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedLedgerDigestUploadsListResult); err != nil {
		return ManagedLedgerDigestUploadsClientListByDatabaseResponse{}, err
	}
	return result, nil
}
