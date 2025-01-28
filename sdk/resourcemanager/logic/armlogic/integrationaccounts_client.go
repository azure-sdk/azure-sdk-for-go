//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationAccountsClient contains the methods for the IntegrationAccounts group.
// Don't use this type directly, use NewIntegrationAccountsClient() instead.
type IntegrationAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationAccountsClient creates a new instance of IntegrationAccountsClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - integrationAccount - The integration account.
//   - options - IntegrationAccountsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountsClient.CreateOrUpdate
//     method.
func (client *IntegrationAccountsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsClientCreateOrUpdateOptions) (IntegrationAccountsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, integrationAccount, options)
	if err != nil {
		return IntegrationAccountsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, integrationAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountsClientDeleteOptions contains the optional parameters for the IntegrationAccountsClient.Delete
//     method.
func (client *IntegrationAccountsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsClientDeleteOptions) (IntegrationAccountsClientDeleteResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
	if err != nil {
		return IntegrationAccountsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientDeleteResponse{}, err
	}
	return IntegrationAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - options - IntegrationAccountsClientGetOptions contains the optional parameters for the IntegrationAccountsClient.Get method.
func (client *IntegrationAccountsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsClientGetOptions) (IntegrationAccountsClientGetResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
	if err != nil {
		return IntegrationAccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountsClient) getHandleResponse(resp *http.Response) (IntegrationAccountsClientGetResponse, error) {
	result := IntegrationAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of integration accounts by resource group.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - options - IntegrationAccountsClientListByResourceGroupOptions contains the optional parameters for the IntegrationAccountsClient.NewListByResourceGroupPager
//     method.
func (client *IntegrationAccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *IntegrationAccountsClientListByResourceGroupOptions) *runtime.Pager[IntegrationAccountsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountsClientListByResourceGroupResponse]{
		More: func(page IntegrationAccountsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountsClientListByResourceGroupResponse) (IntegrationAccountsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return IntegrationAccountsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IntegrationAccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IntegrationAccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IntegrationAccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (IntegrationAccountsClientListByResourceGroupResponse, error) {
	result := IntegrationAccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountListResult); err != nil {
		return IntegrationAccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of integration accounts by subscription.
//
// Generated from API version 2019-05-01
//   - options - IntegrationAccountsClientListBySubscriptionOptions contains the optional parameters for the IntegrationAccountsClient.NewListBySubscriptionPager
//     method.
func (client *IntegrationAccountsClient) NewListBySubscriptionPager(options *IntegrationAccountsClientListBySubscriptionOptions) *runtime.Pager[IntegrationAccountsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountsClientListBySubscriptionResponse]{
		More: func(page IntegrationAccountsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountsClientListBySubscriptionResponse) (IntegrationAccountsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return IntegrationAccountsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IntegrationAccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *IntegrationAccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IntegrationAccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (IntegrationAccountsClientListBySubscriptionResponse, error) {
	result := IntegrationAccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountListResult); err != nil {
		return IntegrationAccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListCallbackURL - Gets the integration account callback URL.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - parameters - The callback URL parameters.
//   - options - IntegrationAccountsClientListCallbackURLOptions contains the optional parameters for the IntegrationAccountsClient.ListCallbackURL
//     method.
func (client *IntegrationAccountsClient) ListCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters, options *IntegrationAccountsClientListCallbackURLOptions) (IntegrationAccountsClientListCallbackURLResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.ListCallbackURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, parameters, options)
	if err != nil {
		return IntegrationAccountsClientListCallbackURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientListCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientListCallbackURLResponse{}, err
	}
	resp, err := client.listCallbackURLHandleResponse(httpResp)
	return resp, err
}

// listCallbackURLCreateRequest creates the ListCallbackURL request.
func (client *IntegrationAccountsClient) listCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, parameters GetCallbackURLParameters, options *IntegrationAccountsClientListCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// listCallbackURLHandleResponse handles the ListCallbackURL response.
func (client *IntegrationAccountsClient) listCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountsClientListCallbackURLResponse, error) {
	result := IntegrationAccountsClientListCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CallbackURL); err != nil {
		return IntegrationAccountsClientListCallbackURLResponse{}, err
	}
	return result, nil
}

// NewListKeyVaultKeysPager - Gets the integration account's Key Vault keys.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - listKeyVaultKeys - The key vault parameters.
//   - options - IntegrationAccountsClientListKeyVaultKeysOptions contains the optional parameters for the IntegrationAccountsClient.NewListKeyVaultKeysPager
//     method.
func (client *IntegrationAccountsClient) NewListKeyVaultKeysPager(resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition, options *IntegrationAccountsClientListKeyVaultKeysOptions) *runtime.Pager[IntegrationAccountsClientListKeyVaultKeysResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountsClientListKeyVaultKeysResponse]{
		More: func(page IntegrationAccountsClientListKeyVaultKeysResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountsClientListKeyVaultKeysResponse) (IntegrationAccountsClientListKeyVaultKeysResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationAccountsClient.NewListKeyVaultKeysPager")
			req, err := client.listKeyVaultKeysCreateRequest(ctx, resourceGroupName, integrationAccountName, listKeyVaultKeys, options)
			if err != nil {
				return IntegrationAccountsClientListKeyVaultKeysResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IntegrationAccountsClientListKeyVaultKeysResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationAccountsClientListKeyVaultKeysResponse{}, runtime.NewResponseError(resp)
			}
			return client.listKeyVaultKeysHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listKeyVaultKeysCreateRequest creates the ListKeyVaultKeys request.
func (client *IntegrationAccountsClient) listKeyVaultKeysCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, listKeyVaultKeys ListKeyVaultKeysDefinition, options *IntegrationAccountsClientListKeyVaultKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/listKeyVaultKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, listKeyVaultKeys); err != nil {
		return nil, err
	}
	return req, nil
}

// listKeyVaultKeysHandleResponse handles the ListKeyVaultKeys response.
func (client *IntegrationAccountsClient) listKeyVaultKeysHandleResponse(resp *http.Response) (IntegrationAccountsClientListKeyVaultKeysResponse, error) {
	result := IntegrationAccountsClientListKeyVaultKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyVaultKeyCollection); err != nil {
		return IntegrationAccountsClientListKeyVaultKeysResponse{}, err
	}
	return result, nil
}

// LogTrackingEvents - Logs the integration account's tracking events.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - logTrackingEvents - The callback URL parameters.
//   - options - IntegrationAccountsClientLogTrackingEventsOptions contains the optional parameters for the IntegrationAccountsClient.LogTrackingEvents
//     method.
func (client *IntegrationAccountsClient) LogTrackingEvents(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition, options *IntegrationAccountsClientLogTrackingEventsOptions) (IntegrationAccountsClientLogTrackingEventsResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.LogTrackingEvents"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.logTrackingEventsCreateRequest(ctx, resourceGroupName, integrationAccountName, logTrackingEvents, options)
	if err != nil {
		return IntegrationAccountsClientLogTrackingEventsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientLogTrackingEventsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientLogTrackingEventsResponse{}, err
	}
	return IntegrationAccountsClientLogTrackingEventsResponse{}, nil
}

// logTrackingEventsCreateRequest creates the LogTrackingEvents request.
func (client *IntegrationAccountsClient) logTrackingEventsCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, logTrackingEvents TrackingEventsDefinition, options *IntegrationAccountsClientLogTrackingEventsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/logTrackingEvents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, logTrackingEvents); err != nil {
		return nil, err
	}
	return req, nil
}

// RegenerateAccessKey - Regenerates the integration account access key.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - regenerateAccessKey - The access key type.
//   - options - IntegrationAccountsClientRegenerateAccessKeyOptions contains the optional parameters for the IntegrationAccountsClient.RegenerateAccessKey
//     method.
func (client *IntegrationAccountsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter, options *IntegrationAccountsClientRegenerateAccessKeyOptions) (IntegrationAccountsClientRegenerateAccessKeyResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.RegenerateAccessKey"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateAccessKeyCreateRequest(ctx, resourceGroupName, integrationAccountName, regenerateAccessKey, options)
	if err != nil {
		return IntegrationAccountsClientRegenerateAccessKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientRegenerateAccessKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientRegenerateAccessKeyResponse{}, err
	}
	resp, err := client.regenerateAccessKeyHandleResponse(httpResp)
	return resp, err
}

// regenerateAccessKeyCreateRequest creates the RegenerateAccessKey request.
func (client *IntegrationAccountsClient) regenerateAccessKeyCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, regenerateAccessKey RegenerateActionParameter, options *IntegrationAccountsClientRegenerateAccessKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/regenerateAccessKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, regenerateAccessKey); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateAccessKeyHandleResponse handles the RegenerateAccessKey response.
func (client *IntegrationAccountsClient) regenerateAccessKeyHandleResponse(resp *http.Response) (IntegrationAccountsClientRegenerateAccessKeyResponse, error) {
	result := IntegrationAccountsClientRegenerateAccessKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsClientRegenerateAccessKeyResponse{}, err
	}
	return result, nil
}

// Update - Updates an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - integrationAccountName - The integration account name.
//   - integrationAccount - The integration account.
//   - options - IntegrationAccountsClientUpdateOptions contains the optional parameters for the IntegrationAccountsClient.Update
//     method.
func (client *IntegrationAccountsClient) Update(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsClientUpdateOptions) (IntegrationAccountsClientUpdateResponse, error) {
	var err error
	const operationName = "IntegrationAccountsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, integrationAccountName, integrationAccount, options)
	if err != nil {
		return IntegrationAccountsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationAccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationAccountsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *IntegrationAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, integrationAccount IntegrationAccount, options *IntegrationAccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, integrationAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *IntegrationAccountsClient) updateHandleResponse(resp *http.Response) (IntegrationAccountsClientUpdateResponse, error) {
	result := IntegrationAccountsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccount); err != nil {
		return IntegrationAccountsClientUpdateResponse{}, err
	}
	return result, nil
}
