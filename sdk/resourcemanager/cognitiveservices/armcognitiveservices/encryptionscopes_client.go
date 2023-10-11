//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// EncryptionScopesClient contains the methods for the EncryptionScopes group.
// Don't use this type directly, use NewEncryptionScopesClient() instead.
type EncryptionScopesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEncryptionScopesClient creates a new instance of EncryptionScopesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEncryptionScopesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EncryptionScopesClient, error) {
	cl, err := arm.NewClient(moduleName+".EncryptionScopesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EncryptionScopesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Update the state of specified encryptionScope associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - encryptionScopeName - The name of the encryptionScope associated with the Cognitive Services Account
//   - encryptionScope - The encryptionScope properties.
//   - options - EncryptionScopesClientCreateOrUpdateOptions contains the optional parameters for the EncryptionScopesClient.CreateOrUpdate
//     method.
func (client *EncryptionScopesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientCreateOrUpdateOptions) (EncryptionScopesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, encryptionScope, options)
	if err != nil {
		return EncryptionScopesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EncryptionScopesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EncryptionScopesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EncryptionScopesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, encryptionScope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EncryptionScopesClient) createOrUpdateHandleResponse(resp *http.Response) (EncryptionScopesClientCreateOrUpdateResponse, error) {
	result := EncryptionScopesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScope); err != nil {
		return EncryptionScopesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified encryptionScope associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - encryptionScopeName - The name of the encryptionScope associated with the Cognitive Services Account
//   - options - EncryptionScopesClientBeginDeleteOptions contains the optional parameters for the EncryptionScopesClient.BeginDelete
//     method.
func (client *EncryptionScopesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientBeginDeleteOptions) (*runtime.Poller[EncryptionScopesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, encryptionScopeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[EncryptionScopesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[EncryptionScopesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified encryptionScope associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *EncryptionScopesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, options)
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
func (client *EncryptionScopesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified EncryptionScope associated with the Cognitive Services account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - encryptionScopeName - The name of the encryptionScope associated with the Cognitive Services Account
//   - options - EncryptionScopesClientGetOptions contains the optional parameters for the EncryptionScopesClient.Get method.
func (client *EncryptionScopesClient) Get(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientGetOptions) (EncryptionScopesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, options)
	if err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EncryptionScopesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EncryptionScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if encryptionScopeName == "" {
		return nil, errors.New("parameter encryptionScopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EncryptionScopesClient) getHandleResponse(resp *http.Response) (EncryptionScopesClientGetResponse, error) {
	result := EncryptionScopesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScope); err != nil {
		return EncryptionScopesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the content filters associated with the Azure OpenAI account.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - options - EncryptionScopesClientListOptions contains the optional parameters for the EncryptionScopesClient.NewListPager
//     method.
func (client *EncryptionScopesClient) NewListPager(resourceGroupName string, accountName string, options *EncryptionScopesClientListOptions) *runtime.Pager[EncryptionScopesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EncryptionScopesClientListResponse]{
		More: func(page EncryptionScopesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EncryptionScopesClientListResponse) (EncryptionScopesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EncryptionScopesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EncryptionScopesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EncryptionScopesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EncryptionScopesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *EncryptionScopesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/encryptionScopes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EncryptionScopesClient) listHandleResponse(resp *http.Response) (EncryptionScopesClientListResponse, error) {
	result := EncryptionScopesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionScopeListResult); err != nil {
		return EncryptionScopesClientListResponse{}, err
	}
	return result, nil
}
