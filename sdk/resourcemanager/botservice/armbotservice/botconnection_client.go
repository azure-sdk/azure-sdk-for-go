//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbotservice

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BotConnectionClient contains the methods for the BotConnection group.
// Don't use this type directly, use NewBotConnectionClient() instead.
type BotConnectionClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBotConnectionClient creates a new instance of BotConnectionClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBotConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BotConnectionClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &BotConnectionClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Register a new Auth Connection for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - connectionName - The name of the Bot Service Connection Setting resource.
//   - parameters - The parameters to provide for creating the Connection Setting.
//   - options - BotConnectionClientCreateOptions contains the optional parameters for the BotConnectionClient.Create method.
func (client *BotConnectionClient) Create(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionClientCreateOptions) (BotConnectionClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, connectionName, parameters, options)
	if err != nil {
		return BotConnectionClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotConnectionClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *BotConnectionClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *BotConnectionClient) createHandleResponse(resp *http.Response) (BotConnectionClientCreateResponse, error) {
	result := BotConnectionClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Connection Setting registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - connectionName - The name of the Bot Service Connection Setting resource.
//   - options - BotConnectionClientDeleteOptions contains the optional parameters for the BotConnectionClient.Delete method.
func (client *BotConnectionClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientDeleteOptions) (BotConnectionClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BotConnectionClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BotConnectionClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BotConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Connection Setting registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - connectionName - The name of the Bot Service Connection Setting resource.
//   - options - BotConnectionClientGetOptions contains the optional parameters for the BotConnectionClient.Get method.
func (client *BotConnectionClient) Get(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientGetOptions) (BotConnectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BotConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BotConnectionClient) getHandleResponse(resp *http.Response) (BotConnectionClientGetResponse, error) {
	result := BotConnectionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBotServicePager - Returns all the Connection Settings registered to a particular BotService resource
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - options - BotConnectionClientListByBotServiceOptions contains the optional parameters for the BotConnectionClient.NewListByBotServicePager
//     method.
func (client *BotConnectionClient) NewListByBotServicePager(resourceGroupName string, resourceName string, options *BotConnectionClientListByBotServiceOptions) *runtime.Pager[BotConnectionClientListByBotServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[BotConnectionClientListByBotServiceResponse]{
		More: func(page BotConnectionClientListByBotServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BotConnectionClientListByBotServiceResponse) (BotConnectionClientListByBotServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByBotServiceCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BotConnectionClientListByBotServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BotConnectionClientListByBotServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BotConnectionClientListByBotServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByBotServiceHandleResponse(resp)
		},
	})
}

// listByBotServiceCreateRequest creates the ListByBotService request.
func (client *BotConnectionClient) listByBotServiceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *BotConnectionClientListByBotServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBotServiceHandleResponse handles the ListByBotService response.
func (client *BotConnectionClient) listByBotServiceHandleResponse(resp *http.Response) (BotConnectionClientListByBotServiceResponse, error) {
	result := BotConnectionClientListByBotServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSettingResponseList); err != nil {
		return BotConnectionClientListByBotServiceResponse{}, err
	}
	return result, nil
}

// ListServiceProviders - Lists the available Service Providers for creating Connection Settings
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - options - BotConnectionClientListServiceProvidersOptions contains the optional parameters for the BotConnectionClient.ListServiceProviders
//     method.
func (client *BotConnectionClient) ListServiceProviders(ctx context.Context, options *BotConnectionClientListServiceProvidersOptions) (BotConnectionClientListServiceProvidersResponse, error) {
	req, err := client.listServiceProvidersCreateRequest(ctx, options)
	if err != nil {
		return BotConnectionClientListServiceProvidersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientListServiceProvidersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionClientListServiceProvidersResponse{}, runtime.NewResponseError(resp)
	}
	return client.listServiceProvidersHandleResponse(resp)
}

// listServiceProvidersCreateRequest creates the ListServiceProviders request.
func (client *BotConnectionClient) listServiceProvidersCreateRequest(ctx context.Context, options *BotConnectionClientListServiceProvidersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/listAuthServiceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listServiceProvidersHandleResponse handles the ListServiceProviders response.
func (client *BotConnectionClient) listServiceProvidersHandleResponse(resp *http.Response) (BotConnectionClientListServiceProvidersResponse, error) {
	result := BotConnectionClientListServiceProvidersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceProviderResponseList); err != nil {
		return BotConnectionClientListServiceProvidersResponse{}, err
	}
	return result, nil
}

// ListWithSecrets - Get a Connection Setting registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - connectionName - The name of the Bot Service Connection Setting resource.
//   - options - BotConnectionClientListWithSecretsOptions contains the optional parameters for the BotConnectionClient.ListWithSecrets
//     method.
func (client *BotConnectionClient) ListWithSecrets(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientListWithSecretsOptions) (BotConnectionClientListWithSecretsResponse, error) {
	req, err := client.listWithSecretsCreateRequest(ctx, resourceGroupName, resourceName, connectionName, options)
	if err != nil {
		return BotConnectionClientListWithSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientListWithSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BotConnectionClientListWithSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listWithSecretsHandleResponse(resp)
}

// listWithSecretsCreateRequest creates the ListWithSecrets request.
func (client *BotConnectionClient) listWithSecretsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, options *BotConnectionClientListWithSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}/listWithSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithSecretsHandleResponse handles the ListWithSecrets response.
func (client *BotConnectionClient) listWithSecretsHandleResponse(resp *http.Response) (BotConnectionClientListWithSecretsResponse, error) {
	result := BotConnectionClientListWithSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionClientListWithSecretsResponse{}, err
	}
	return result, nil
}

// Update - Updates a Connection Setting registration for a Bot Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-15
//   - resourceGroupName - The name of the Bot resource group in the user subscription.
//   - resourceName - The name of the Bot resource.
//   - connectionName - The name of the Bot Service Connection Setting resource.
//   - parameters - The parameters to provide for updating the Connection Setting.
//   - options - BotConnectionClientUpdateOptions contains the optional parameters for the BotConnectionClient.Update method.
func (client *BotConnectionClient) Update(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionClientUpdateOptions) (BotConnectionClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, connectionName, parameters, options)
	if err != nil {
		return BotConnectionClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BotConnectionClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BotConnectionClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *BotConnectionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, connectionName string, parameters ConnectionSetting, options *BotConnectionClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *BotConnectionClient) updateHandleResponse(resp *http.Response) (BotConnectionClientUpdateResponse, error) {
	result := BotConnectionClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionSetting); err != nil {
		return BotConnectionClientUpdateResponse{}, err
	}
	return result, nil
}
