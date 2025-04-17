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

// SynchronizationSettingsClient contains the methods for the SynchronizationSettings group.
// Don't use this type directly, use NewSynchronizationSettingsClient() instead.
type SynchronizationSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSynchronizationSettingsClient creates a new instance of SynchronizationSettingsClient with the specified values.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSynchronizationSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SynchronizationSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SynchronizationSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a synchronizationSetting
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share to add the synchronization setting to.
//   - synchronizationSettingName - The name of the synchronizationSetting.
//   - synchronizationSetting - The new synchronization setting information.
//   - options - SynchronizationSettingsClientCreateOptions contains the optional parameters for the SynchronizationSettingsClient.Create
//     method.
func (client *SynchronizationSettingsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting SynchronizationSettingClassification, options *SynchronizationSettingsClientCreateOptions) (SynchronizationSettingsClientCreateResponse, error) {
	var err error
	const operationName = "SynchronizationSettingsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, synchronizationSetting, options)
	if err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SynchronizationSettingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, synchronizationSetting SynchronizationSettingClassification, _ *SynchronizationSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, synchronizationSetting); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SynchronizationSettingsClient) createHandleResponse(resp *http.Response) (SynchronizationSettingsClientCreateResponse, error) {
	result := SynchronizationSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SynchronizationSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - synchronizationSettingName - The name of the synchronizationSetting .
//   - options - SynchronizationSettingsClientBeginDeleteOptions contains the optional parameters for the SynchronizationSettingsClient.BeginDelete
//     method.
func (client *SynchronizationSettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientBeginDeleteOptions) (*runtime.Poller[SynchronizationSettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SynchronizationSettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SynchronizationSettingsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
func (client *SynchronizationSettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SynchronizationSettingsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
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
func (client *SynchronizationSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, _ *SynchronizationSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
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

// Get - Get a synchronizationSetting in a share
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - synchronizationSettingName - The name of the synchronizationSetting.
//   - options - SynchronizationSettingsClientGetOptions contains the optional parameters for the SynchronizationSettingsClient.Get
//     method.
func (client *SynchronizationSettingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, options *SynchronizationSettingsClientGetOptions) (SynchronizationSettingsClientGetResponse, error) {
	var err error
	const operationName = "SynchronizationSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareName, synchronizationSettingName, options)
	if err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SynchronizationSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SynchronizationSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, synchronizationSettingName string, _ *SynchronizationSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}"
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
	if synchronizationSettingName == "" {
		return nil, errors.New("parameter synchronizationSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{synchronizationSettingName}", url.PathEscape(synchronizationSettingName))
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
func (client *SynchronizationSettingsClient) getHandleResponse(resp *http.Response) (SynchronizationSettingsClientGetResponse, error) {
	result := SynchronizationSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SynchronizationSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySharePager - List synchronizationSettings in a share
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareName - The name of the share.
//   - options - SynchronizationSettingsClientListByShareOptions contains the optional parameters for the SynchronizationSettingsClient.NewListBySharePager
//     method.
func (client *SynchronizationSettingsClient) NewListBySharePager(resourceGroupName string, accountName string, shareName string, options *SynchronizationSettingsClientListByShareOptions) *runtime.Pager[SynchronizationSettingsClientListByShareResponse] {
	return runtime.NewPager(runtime.PagingHandler[SynchronizationSettingsClientListByShareResponse]{
		More: func(page SynchronizationSettingsClientListByShareResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SynchronizationSettingsClientListByShareResponse) (SynchronizationSettingsClientListByShareResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SynchronizationSettingsClient.NewListBySharePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByShareCreateRequest(ctx, resourceGroupName, accountName, shareName, options)
			}, nil)
			if err != nil {
				return SynchronizationSettingsClientListByShareResponse{}, err
			}
			return client.listByShareHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByShareCreateRequest creates the ListByShare request.
func (client *SynchronizationSettingsClient) listByShareCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareName string, options *SynchronizationSettingsClientListByShareOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings"
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
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByShareHandleResponse handles the ListByShare response.
func (client *SynchronizationSettingsClient) listByShareHandleResponse(resp *http.Response) (SynchronizationSettingsClientListByShareResponse, error) {
	result := SynchronizationSettingsClientListByShareResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SynchronizationSettingList); err != nil {
		return SynchronizationSettingsClientListByShareResponse{}, err
	}
	return result, nil
}
