// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// SharesClient contains the methods for the Shares group.
// Don't use this type directly, use NewSharesClient() instead.
type SharesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharesClient creates a new instance of SharesClient with the specified values.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - share - The share payload
//   - options - SharesClientCreateOptions contains the optional parameters for the SharesClient.Create method.
func (client *SharesClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share, options *SharesClientCreateOptions) (SharesClientCreateResponse, error) {
	var err error
	const operationName = "SharesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareName, share, options)
	if err != nil {
		return SharesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SharesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SharesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, share Share, _ *SharesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, share); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SharesClient) createHandleResponse(resp *http.Response) (SharesClientCreateResponse, error) {
	result := SharesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Share); err != nil {
		return SharesClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - options - SharesClientBeginDeleteOptions contains the optional parameters for the SharesClient.BeginDelete method.
func (client *SharesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesClientBeginDeleteOptions) (*runtime.Poller[SharesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, shareName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SharesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SharesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
func (client *SharesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SharesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
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
func (client *SharesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, _ *SharesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share to retrieve.
//   - options - SharesClientGetOptions contains the optional parameters for the SharesClient.Get method.
func (client *SharesClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesClientGetOptions) (SharesClientGetResponse, error) {
	var err error
	const operationName = "SharesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
	if err != nil {
		return SharesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SharesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, _ *SharesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharesClient) getHandleResponse(resp *http.Response) (SharesClientGetResponse, error) {
	result := SharesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Share); err != nil {
		return SharesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - List shares in an account
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - options - SharesClientListByAccountOptions contains the optional parameters for the SharesClient.NewListByAccountPager
//     method.
func (client *SharesClient) NewListByAccountPager(resourceGroupName string, accountName string, options *SharesClientListByAccountOptions) *runtime.Pager[SharesClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharesClientListByAccountResponse]{
		More: func(page SharesClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharesClientListByAccountResponse) (SharesClientListByAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharesClient.NewListByAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return SharesClientListByAccountResponse{}, err
			}
			return client.listByAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *SharesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SharesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *SharesClient) listByAccountHandleResponse(resp *http.Response) (SharesClientListByAccountResponse, error) {
	result := SharesClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShareList); err != nil {
		return SharesClientListByAccountResponse{}, err
	}
	return result, nil
}

// NewListSynchronizationDetailsPager - List synchronization details
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - shareSynchronization - Share Synchronization payload.
//   - options - SharesClientListSynchronizationDetailsOptions contains the optional parameters for the SharesClient.NewListSynchronizationDetailsPager
//     method.
func (client *SharesClient) NewListSynchronizationDetailsPager(resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, options *SharesClientListSynchronizationDetailsOptions) *runtime.Pager[SharesClientListSynchronizationDetailsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharesClientListSynchronizationDetailsResponse]{
		More: func(page SharesClientListSynchronizationDetailsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharesClientListSynchronizationDetailsResponse) (SharesClientListSynchronizationDetailsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharesClient.NewListSynchronizationDetailsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSynchronizationDetailsCreateRequest(ctx, resourceGroupName, accountName, shareName, shareSynchronization, options)
			}, nil)
			if err != nil {
				return SharesClientListSynchronizationDetailsResponse{}, err
			}
			return client.listSynchronizationDetailsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSynchronizationDetailsCreateRequest creates the ListSynchronizationDetails request.
func (client *SharesClient) listSynchronizationDetailsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, shareSynchronization ShareSynchronization, options *SharesClientListSynchronizationDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizationDetails"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, shareSynchronization); err != nil {
		return nil, err
	}
	return req, nil
}

// listSynchronizationDetailsHandleResponse handles the ListSynchronizationDetails response.
func (client *SharesClient) listSynchronizationDetailsHandleResponse(resp *http.Response) (SharesClientListSynchronizationDetailsResponse, error) {
	result := SharesClientListSynchronizationDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SynchronizationDetailsList); err != nil {
		return SharesClientListSynchronizationDetailsResponse{}, err
	}
	return result, nil
}

// NewListSynchronizationsPager - List synchronizations of a share
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - options - SharesClientListSynchronizationsOptions contains the optional parameters for the SharesClient.NewListSynchronizationsPager
//     method.
func (client *SharesClient) NewListSynchronizationsPager(resourceGroupName string, accountName string, shareName string, options *SharesClientListSynchronizationsOptions) *runtime.Pager[SharesClientListSynchronizationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharesClientListSynchronizationsResponse]{
		More: func(page SharesClientListSynchronizationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharesClientListSynchronizationsResponse) (SharesClientListSynchronizationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharesClient.NewListSynchronizationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSynchronizationsCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
			}, nil)
			if err != nil {
				return SharesClientListSynchronizationsResponse{}, err
			}
			return client.listSynchronizationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSynchronizationsCreateRequest creates the ListSynchronizations request.
func (client *SharesClient) listSynchronizationsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SharesClientListSynchronizationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/listSynchronizations"
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
	if shareName == "" {
		return nil, errors.New("parameter shareName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareName}", url.PathEscape(shareName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSynchronizationsHandleResponse handles the ListSynchronizations response.
func (client *SharesClient) listSynchronizationsHandleResponse(resp *http.Response) (SharesClientListSynchronizationsResponse, error) {
	result := SharesClientListSynchronizationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShareSynchronizationList); err != nil {
		return SharesClientListSynchronizationsResponse{}, err
	}
	return result, nil
}
