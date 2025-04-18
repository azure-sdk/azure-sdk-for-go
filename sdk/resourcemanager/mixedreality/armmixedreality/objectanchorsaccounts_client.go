// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

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

// ObjectAnchorsAccountsClient contains the methods for the ObjectAnchorsAccounts group.
// Don't use this type directly, use NewObjectAnchorsAccountsClient() instead.
type ObjectAnchorsAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewObjectAnchorsAccountsClient creates a new instance of ObjectAnchorsAccountsClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewObjectAnchorsAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ObjectAnchorsAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ObjectAnchorsAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// Creating or Updating an object anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - objectAnchorsAccount - Object Anchors Account parameter.
//   - options - ObjectAnchorsAccountsClientCreateOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Create
//     method.
func (client *ObjectAnchorsAccountsClient) Create(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount, options *ObjectAnchorsAccountsClientCreateOptions) (ObjectAnchorsAccountsClientCreateResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, objectAnchorsAccount, options)
	if err != nil {
		return ObjectAnchorsAccountsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ObjectAnchorsAccountsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount, _ *ObjectAnchorsAccountsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, objectAnchorsAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ObjectAnchorsAccountsClient) createHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientCreateResponse, error) {
	result := ObjectAnchorsAccountsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectAnchorsAccount); err != nil {
		return ObjectAnchorsAccountsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// Delete an Object Anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - ObjectAnchorsAccountsClientDeleteOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Delete
//     method.
func (client *ObjectAnchorsAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, options *ObjectAnchorsAccountsClientDeleteOptions) (ObjectAnchorsAccountsClientDeleteResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ObjectAnchorsAccountsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientDeleteResponse{}, err
	}
	return ObjectAnchorsAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ObjectAnchorsAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *ObjectAnchorsAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// Retrieve an Object Anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - ObjectAnchorsAccountsClientGetOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Get
//     method.
func (client *ObjectAnchorsAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, options *ObjectAnchorsAccountsClientGetOptions) (ObjectAnchorsAccountsClientGetResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ObjectAnchorsAccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ObjectAnchorsAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *ObjectAnchorsAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}"
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
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectAnchorsAccountsClient) getHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientGetResponse, error) {
	result := ObjectAnchorsAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectAnchorsAccount); err != nil {
		return ObjectAnchorsAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// List Resources by Resource Group
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - options - ObjectAnchorsAccountsClientListByResourceGroupOptions contains the optional parameters for the ObjectAnchorsAccountsClient.NewListByResourceGroupPager
//     method.
func (client *ObjectAnchorsAccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *ObjectAnchorsAccountsClientListByResourceGroupOptions) *runtime.Pager[ObjectAnchorsAccountsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectAnchorsAccountsClientListByResourceGroupResponse]{
		More: func(page ObjectAnchorsAccountsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ObjectAnchorsAccountsClientListByResourceGroupResponse) (ObjectAnchorsAccountsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ObjectAnchorsAccountsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ObjectAnchorsAccountsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ObjectAnchorsAccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *ObjectAnchorsAccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts"
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
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ObjectAnchorsAccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientListByResourceGroupResponse, error) {
	result := ObjectAnchorsAccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectAnchorsAccountPage); err != nil {
		return ObjectAnchorsAccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// List Object Anchors Accounts by Subscription
//
// Generated from API version 2021-03-01-preview
//   - options - ObjectAnchorsAccountsClientListBySubscriptionOptions contains the optional parameters for the ObjectAnchorsAccountsClient.NewListBySubscriptionPager
//     method.
func (client *ObjectAnchorsAccountsClient) NewListBySubscriptionPager(options *ObjectAnchorsAccountsClientListBySubscriptionOptions) *runtime.Pager[ObjectAnchorsAccountsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectAnchorsAccountsClientListBySubscriptionResponse]{
		More: func(page ObjectAnchorsAccountsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ObjectAnchorsAccountsClientListBySubscriptionResponse) (ObjectAnchorsAccountsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ObjectAnchorsAccountsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ObjectAnchorsAccountsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ObjectAnchorsAccountsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *ObjectAnchorsAccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/objectAnchorsAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ObjectAnchorsAccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientListBySubscriptionResponse, error) {
	result := ObjectAnchorsAccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectAnchorsAccountPage); err != nil {
		return ObjectAnchorsAccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// List Both of the 2 Keys of an object anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - ObjectAnchorsAccountsClientListKeysOptions contains the optional parameters for the ObjectAnchorsAccountsClient.ListKeys
//     method.
func (client *ObjectAnchorsAccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string, options *ObjectAnchorsAccountsClientListKeysOptions) (ObjectAnchorsAccountsClientListKeysResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return ObjectAnchorsAccountsClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ObjectAnchorsAccountsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, _ *ObjectAnchorsAccountsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}/listKeys"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ObjectAnchorsAccountsClient) listKeysHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientListKeysResponse, error) {
	result := ObjectAnchorsAccountsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return ObjectAnchorsAccountsClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// Regenerate specified Key of an object anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - regenerate - Required information for key regeneration.
//   - options - ObjectAnchorsAccountsClientRegenerateKeysOptions contains the optional parameters for the ObjectAnchorsAccountsClient.RegenerateKeys
//     method.
func (client *ObjectAnchorsAccountsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest, options *ObjectAnchorsAccountsClientRegenerateKeysOptions) (ObjectAnchorsAccountsClientRegenerateKeysResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.RegenerateKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, accountName, regenerate, options)
	if err != nil {
		return ObjectAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	resp, err := client.regenerateKeysHandleResponse(httpResp)
	return resp, err
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *ObjectAnchorsAccountsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest, _ *ObjectAnchorsAccountsClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}/regenerateKeys"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, regenerate); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *ObjectAnchorsAccountsClient) regenerateKeysHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientRegenerateKeysResponse, error) {
	result := ObjectAnchorsAccountsClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return ObjectAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	return result, nil
}

// Update - > [!NOTE]
// Mixed Reality retirement
// The Mixed Reality service is now deprecated and will be retired.
// Updating an Object Anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - objectAnchorsAccount - Object Anchors Account parameter.
//   - options - ObjectAnchorsAccountsClientUpdateOptions contains the optional parameters for the ObjectAnchorsAccountsClient.Update
//     method.
func (client *ObjectAnchorsAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount, options *ObjectAnchorsAccountsClientUpdateOptions) (ObjectAnchorsAccountsClientUpdateResponse, error) {
	var err error
	const operationName = "ObjectAnchorsAccountsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, objectAnchorsAccount, options)
	if err != nil {
		return ObjectAnchorsAccountsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ObjectAnchorsAccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ObjectAnchorsAccountsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ObjectAnchorsAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectAnchorsAccount ObjectAnchorsAccount, _ *ObjectAnchorsAccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/objectAnchorsAccounts/{accountName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, objectAnchorsAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ObjectAnchorsAccountsClient) updateHandleResponse(resp *http.Response) (ObjectAnchorsAccountsClientUpdateResponse, error) {
	result := ObjectAnchorsAccountsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectAnchorsAccount); err != nil {
		return ObjectAnchorsAccountsClientUpdateResponse{}, err
	}
	return result, nil
}
