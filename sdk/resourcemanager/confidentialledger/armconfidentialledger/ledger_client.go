// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger

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

// LedgerClient contains the methods for the Ledger group.
// Don't use this type directly, use NewLedgerClient() instead.
type LedgerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLedgerClient creates a new instance of LedgerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLedgerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LedgerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LedgerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginBackup - Backs up a Confidential Ledger Resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - confidentialLedger - Confidential Ledger Backup Request Body
//   - options - LedgerClientBeginBackupOptions contains the optional parameters for the LedgerClient.BeginBackup method.
func (client *LedgerClient) BeginBackup(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Backup, options *LedgerClientBeginBackupOptions) (*runtime.Poller[LedgerClientBackupResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.backup(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LedgerClientBackupResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LedgerClientBackupResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Backup - Backs up a Confidential Ledger Resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
func (client *LedgerClient) backup(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Backup, options *LedgerClientBeginBackupOptions) (*http.Response, error) {
	var err error
	const operationName = "LedgerClient.BeginBackup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.backupCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
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

// backupCreateRequest creates the Backup request.
func (client *LedgerClient) backupCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Backup, _ *LedgerClientBeginBackupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}/backup"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, confidentialLedger); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreate - Creates a Confidential Ledger with the specified ledger parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - confidentialLedger - Confidential Ledger Create Request Body
//   - options - LedgerClientBeginCreateOptions contains the optional parameters for the LedgerClient.BeginCreate method.
func (client *LedgerClient) BeginCreate(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginCreateOptions) (*runtime.Poller[LedgerClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LedgerClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LedgerClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a Confidential Ledger with the specified ledger parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
func (client *LedgerClient) create(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "LedgerClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *LedgerClient) createCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, _ *LedgerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, confidentialLedger); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an existing Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - options - LedgerClientBeginDeleteOptions contains the optional parameters for the LedgerClient.BeginDelete method.
func (client *LedgerClient) BeginDelete(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientBeginDeleteOptions) (*runtime.Poller[LedgerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, ledgerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LedgerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LedgerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an existing Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
func (client *LedgerClient) deleteOperation(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LedgerClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ledgerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LedgerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, _ *LedgerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the properties of a Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - options - LedgerClientGetOptions contains the optional parameters for the LedgerClient.Get method.
func (client *LedgerClient) Get(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientGetOptions) (LedgerClientGetResponse, error) {
	var err error
	const operationName = "LedgerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ledgerName, options)
	if err != nil {
		return LedgerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LedgerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LedgerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LedgerClient) getCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, _ *LedgerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LedgerClient) getHandleResponse(resp *http.Response) (LedgerClientGetResponse, error) {
	result := LedgerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfidentialLedger); err != nil {
		return LedgerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieves the properties of all Confidential Ledgers.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - LedgerClientListByResourceGroupOptions contains the optional parameters for the LedgerClient.NewListByResourceGroupPager
//     method.
func (client *LedgerClient) NewListByResourceGroupPager(resourceGroupName string, options *LedgerClientListByResourceGroupOptions) *runtime.Pager[LedgerClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[LedgerClientListByResourceGroupResponse]{
		More: func(page LedgerClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LedgerClientListByResourceGroupResponse) (LedgerClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LedgerClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return LedgerClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LedgerClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LedgerClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LedgerClient) listByResourceGroupHandleResponse(resp *http.Response) (LedgerClientListByResourceGroupResponse, error) {
	result := LedgerClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return LedgerClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieves the properties of all Confidential Ledgers.
//
// Generated from API version 2024-09-19-preview
//   - options - LedgerClientListBySubscriptionOptions contains the optional parameters for the LedgerClient.NewListBySubscriptionPager
//     method.
func (client *LedgerClient) NewListBySubscriptionPager(options *LedgerClientListBySubscriptionOptions) *runtime.Pager[LedgerClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[LedgerClientListBySubscriptionResponse]{
		More: func(page LedgerClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LedgerClientListBySubscriptionResponse) (LedgerClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LedgerClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return LedgerClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LedgerClient) listBySubscriptionCreateRequest(ctx context.Context, options *LedgerClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConfidentialLedger/ledgers/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LedgerClient) listBySubscriptionHandleResponse(resp *http.Response) (LedgerClientListBySubscriptionResponse, error) {
	result := LedgerClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return LedgerClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRestore - Restores a Confidential Ledger Resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - confidentialLedger - Confidential Ledger Restore Request Body
//   - options - LedgerClientBeginRestoreOptions contains the optional parameters for the LedgerClient.BeginRestore method.
func (client *LedgerClient) BeginRestore(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Restore, options *LedgerClientBeginRestoreOptions) (*runtime.Poller[LedgerClientRestoreResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restore(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LedgerClientRestoreResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LedgerClientRestoreResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restore - Restores a Confidential Ledger Resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
func (client *LedgerClient) restore(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Restore, options *LedgerClientBeginRestoreOptions) (*http.Response, error) {
	var err error
	const operationName = "LedgerClient.BeginRestore"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restoreCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
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

// restoreCreateRequest creates the Restore request.
func (client *LedgerClient) restoreCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger Restore, _ *LedgerClientBeginRestoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}/restore"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, confidentialLedger); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Updates properties of Confidential Ledger
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ledgerName - Name of the Confidential Ledger
//   - confidentialLedger - Confidential Ledger request body for Updating Ledger
//   - options - LedgerClientBeginUpdateOptions contains the optional parameters for the LedgerClient.BeginUpdate method.
func (client *LedgerClient) BeginUpdate(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginUpdateOptions) (*runtime.Poller[LedgerClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LedgerClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LedgerClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates properties of Confidential Ledger
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-19-preview
func (client *LedgerClient) update(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "LedgerClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *LedgerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, _ *LedgerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-19-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, confidentialLedger); err != nil {
		return nil, err
	}
	return req, nil
}
