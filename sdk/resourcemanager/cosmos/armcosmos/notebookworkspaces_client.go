// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// NotebookWorkspacesClient contains the methods for the NotebookWorkspaces group.
// Don't use this type directly, use NewNotebookWorkspacesClient() instead.
type NotebookWorkspacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotebookWorkspacesClient creates a new instance of NotebookWorkspacesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotebookWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotebookWorkspacesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotebookWorkspacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - notebookCreateUpdateParameters - The notebook workspace to create for the current database account.
//   - options - NotebookWorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the NotebookWorkspacesClient.BeginCreateOrUpdate
//     method.
func (client *NotebookWorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[NotebookWorkspacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NotebookWorkspacesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NotebookWorkspacesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
func (client *NotebookWorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, options *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, notebookCreateUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotebookWorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, notebookCreateUpdateParameters NotebookWorkspaceCreateUpdateParameters, _ *NotebookWorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, notebookCreateUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginDeleteOptions contains the optional parameters for the NotebookWorkspacesClient.BeginDelete
//     method.
func (client *NotebookWorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginDeleteOptions) (*runtime.Poller[NotebookWorkspacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NotebookWorkspacesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NotebookWorkspacesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
func (client *NotebookWorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotebookWorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, _ *NotebookWorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the notebook workspace for a Cosmos DB account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientGetOptions contains the optional parameters for the NotebookWorkspacesClient.Get method.
func (client *NotebookWorkspacesClient) Get(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientGetOptions) (NotebookWorkspacesClientGetResponse, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotebookWorkspacesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NotebookWorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, _ *NotebookWorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NotebookWorkspacesClient) getHandleResponse(resp *http.Response) (NotebookWorkspacesClientGetResponse, error) {
	result := NotebookWorkspacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspace); err != nil {
		return NotebookWorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabaseAccountPager - Gets the notebook workspace resources of an existing Cosmos DB account.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - NotebookWorkspacesClientListByDatabaseAccountOptions contains the optional parameters for the NotebookWorkspacesClient.NewListByDatabaseAccountPager
//     method.
func (client *NotebookWorkspacesClient) NewListByDatabaseAccountPager(resourceGroupName string, accountName string, options *NotebookWorkspacesClientListByDatabaseAccountOptions) *runtime.Pager[NotebookWorkspacesClientListByDatabaseAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotebookWorkspacesClientListByDatabaseAccountResponse]{
		More: func(page NotebookWorkspacesClientListByDatabaseAccountResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *NotebookWorkspacesClientListByDatabaseAccountResponse) (NotebookWorkspacesClientListByDatabaseAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NotebookWorkspacesClient.NewListByDatabaseAccountPager")
			req, err := client.listByDatabaseAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NotebookWorkspacesClientListByDatabaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseAccountCreateRequest creates the ListByDatabaseAccount request.
func (client *NotebookWorkspacesClient) listByDatabaseAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *NotebookWorkspacesClientListByDatabaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseAccountHandleResponse handles the ListByDatabaseAccount response.
func (client *NotebookWorkspacesClient) listByDatabaseAccountHandleResponse(resp *http.Response) (NotebookWorkspacesClientListByDatabaseAccountResponse, error) {
	result := NotebookWorkspacesClientListByDatabaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceListResult); err != nil {
		return NotebookWorkspacesClientListByDatabaseAccountResponse{}, err
	}
	return result, nil
}

// ListConnectionInfo - Retrieves the connection info for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientListConnectionInfoOptions contains the optional parameters for the NotebookWorkspacesClient.ListConnectionInfo
//     method.
func (client *NotebookWorkspacesClient) ListConnectionInfo(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientListConnectionInfoOptions) (NotebookWorkspacesClientListConnectionInfoResponse, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.ListConnectionInfo"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listConnectionInfoCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
	if err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	resp, err := client.listConnectionInfoHandleResponse(httpResp)
	return resp, err
}

// listConnectionInfoCreateRequest creates the ListConnectionInfo request.
func (client *NotebookWorkspacesClient) listConnectionInfoCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, _ *NotebookWorkspacesClientListConnectionInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/listConnectionInfo"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConnectionInfoHandleResponse handles the ListConnectionInfo response.
func (client *NotebookWorkspacesClient) listConnectionInfoHandleResponse(resp *http.Response) (NotebookWorkspacesClientListConnectionInfoResponse, error) {
	result := NotebookWorkspacesClientListConnectionInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotebookWorkspaceConnectionInfoResult); err != nil {
		return NotebookWorkspacesClientListConnectionInfoResponse{}, err
	}
	return result, nil
}

// BeginRegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginRegenerateAuthTokenOptions contains the optional parameters for the NotebookWorkspacesClient.BeginRegenerateAuthToken
//     method.
func (client *NotebookWorkspacesClient) BeginRegenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*runtime.Poller[NotebookWorkspacesClientRegenerateAuthTokenResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.regenerateAuthToken(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NotebookWorkspacesClientRegenerateAuthTokenResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NotebookWorkspacesClientRegenerateAuthTokenResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RegenerateAuthToken - Regenerates the auth token for the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
func (client *NotebookWorkspacesClient) regenerateAuthToken(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*http.Response, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.BeginRegenerateAuthToken"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateAuthTokenCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
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

// regenerateAuthTokenCreateRequest creates the RegenerateAuthToken request.
func (client *NotebookWorkspacesClient) regenerateAuthTokenCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, _ *NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/regenerateAuthToken"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Starts the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - notebookWorkspaceName - The name of the notebook workspace resource.
//   - options - NotebookWorkspacesClientBeginStartOptions contains the optional parameters for the NotebookWorkspacesClient.BeginStart
//     method.
func (client *NotebookWorkspacesClient) BeginStart(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginStartOptions) (*runtime.Poller[NotebookWorkspacesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NotebookWorkspacesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NotebookWorkspacesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Starts the notebook workspace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-15
func (client *NotebookWorkspacesClient) start(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, options *NotebookWorkspacesClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "NotebookWorkspacesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, accountName, notebookWorkspaceName, options)
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

// startCreateRequest creates the Start request.
func (client *NotebookWorkspacesClient) startCreateRequest(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName NotebookWorkspaceName, _ *NotebookWorkspacesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/notebookWorkspaces/{notebookWorkspaceName}/start"
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
	if notebookWorkspaceName == "" {
		return nil, errors.New("parameter notebookWorkspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookWorkspaceName}", url.PathEscape(string(notebookWorkspaceName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
