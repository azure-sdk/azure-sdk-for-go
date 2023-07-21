//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuresiterecovery

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

// EmailConfigurationClient contains the methods for the EmailConfiguration group.
// Don't use this type directly, use NewEmailConfigurationClient() instead.
type EmailConfigurationClient struct {
	internal *arm.Client
}

// NewEmailConfigurationClient creates a new instance of EmailConfigurationClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEmailConfigurationClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EmailConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName+".EmailConfigurationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EmailConfigurationClient{
		internal: cl,
	}
	return client, nil
}

// Create - Creates an alert configuration setting for the given vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - emailConfigurationName - Email configuration name.
//   - body - EmailConfiguration model.
//   - options - EmailConfigurationClientCreateOptions contains the optional parameters for the EmailConfigurationClient.Create
//     method.
func (client *EmailConfigurationClient) Create(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, emailConfigurationName string, body EmailConfigurationModel, options *EmailConfigurationClientCreateOptions) (EmailConfigurationClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, emailConfigurationName, body, options)
	if err != nil {
		return EmailConfigurationClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailConfigurationClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EmailConfigurationClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *EmailConfigurationClient) createCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, emailConfigurationName string, body EmailConfigurationModel, options *EmailConfigurationClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/alertSettings/{emailConfigurationName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if emailConfigurationName == "" {
		return nil, errors.New("parameter emailConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailConfigurationName}", url.PathEscape(emailConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *EmailConfigurationClient) createHandleResponse(resp *http.Response) (EmailConfigurationClientCreateResponse, error) {
	result := EmailConfigurationClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailConfigurationModel); err != nil {
		return EmailConfigurationClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the alert configuration setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - emailConfigurationName - Email configuration name.
//   - options - EmailConfigurationClientGetOptions contains the optional parameters for the EmailConfigurationClient.Get method.
func (client *EmailConfigurationClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, emailConfigurationName string, options *EmailConfigurationClientGetOptions) (EmailConfigurationClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, emailConfigurationName, options)
	if err != nil {
		return EmailConfigurationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmailConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EmailConfigurationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EmailConfigurationClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, emailConfigurationName string, options *EmailConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/alertSettings/{emailConfigurationName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if emailConfigurationName == "" {
		return nil, errors.New("parameter emailConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailConfigurationName}", url.PathEscape(emailConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmailConfigurationClient) getHandleResponse(resp *http.Response) (EmailConfigurationClientGetResponse, error) {
	result := EmailConfigurationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailConfigurationModel); err != nil {
		return EmailConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of alert configuration settings for the given vault.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - options - EmailConfigurationClientListOptions contains the optional parameters for the EmailConfigurationClient.NewListPager
//     method.
func (client *EmailConfigurationClient) NewListPager(subscriptionID string, resourceGroupName string, vaultName string, options *EmailConfigurationClientListOptions) *runtime.Pager[EmailConfigurationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EmailConfigurationClientListResponse]{
		More: func(page EmailConfigurationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EmailConfigurationClientListResponse) (EmailConfigurationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EmailConfigurationClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EmailConfigurationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EmailConfigurationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EmailConfigurationClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, options *EmailConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/alertSettings"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EmailConfigurationClient) listHandleResponse(resp *http.Response) (EmailConfigurationClientListResponse, error) {
	result := EmailConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailConfigurationModelCollection); err != nil {
		return EmailConfigurationClientListResponse{}, err
	}
	return result, nil
}
