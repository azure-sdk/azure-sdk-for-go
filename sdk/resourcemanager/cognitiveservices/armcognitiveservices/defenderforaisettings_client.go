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

// DefenderForAISettingsClient contains the methods for the DefenderForAISettings group.
// Don't use this type directly, use NewDefenderForAISettingsClient() instead.
type DefenderForAISettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDefenderForAISettingsClient creates a new instance of DefenderForAISettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDefenderForAISettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefenderForAISettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DefenderForAISettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates the specified Defender for AI setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - defenderForAISettingName - The name of the defender for AI setting.
//   - defenderForAISettings - Properties describing the Defender for AI setting.
//   - options - DefenderForAISettingsClientCreateOrUpdateOptions contains the optional parameters for the DefenderForAISettingsClient.CreateOrUpdate
//     method.
func (client *DefenderForAISettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings DefenderForAISetting, options *DefenderForAISettingsClientCreateOrUpdateOptions) (DefenderForAISettingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DefenderForAISettingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, defenderForAISettingName, defenderForAISettings, options)
	if err != nil {
		return DefenderForAISettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefenderForAISettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DefenderForAISettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DefenderForAISettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings DefenderForAISetting, options *DefenderForAISettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/defenderForAISettings/{defenderForAISettingName}"
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
	if defenderForAISettingName == "" {
		return nil, errors.New("parameter defenderForAISettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{defenderForAISettingName}", url.PathEscape(defenderForAISettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, defenderForAISettings); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DefenderForAISettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DefenderForAISettingsClientCreateOrUpdateResponse, error) {
	result := DefenderForAISettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderForAISetting); err != nil {
		return DefenderForAISettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the specified Defender for AI setting by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - defenderForAISettingName - The name of the defender for AI setting.
//   - options - DefenderForAISettingsClientGetOptions contains the optional parameters for the DefenderForAISettingsClient.Get
//     method.
func (client *DefenderForAISettingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, options *DefenderForAISettingsClientGetOptions) (DefenderForAISettingsClientGetResponse, error) {
	var err error
	const operationName = "DefenderForAISettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, defenderForAISettingName, options)
	if err != nil {
		return DefenderForAISettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefenderForAISettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefenderForAISettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DefenderForAISettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, options *DefenderForAISettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/defenderForAISettings/{defenderForAISettingName}"
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
	if defenderForAISettingName == "" {
		return nil, errors.New("parameter defenderForAISettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{defenderForAISettingName}", url.PathEscape(defenderForAISettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefenderForAISettingsClient) getHandleResponse(resp *http.Response) (DefenderForAISettingsClientGetResponse, error) {
	result := DefenderForAISettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderForAISetting); err != nil {
		return DefenderForAISettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Defender for AI settings.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - options - DefenderForAISettingsClientListOptions contains the optional parameters for the DefenderForAISettingsClient.NewListPager
//     method.
func (client *DefenderForAISettingsClient) NewListPager(resourceGroupName string, accountName string, options *DefenderForAISettingsClientListOptions) *runtime.Pager[DefenderForAISettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefenderForAISettingsClientListResponse]{
		More: func(page DefenderForAISettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefenderForAISettingsClientListResponse) (DefenderForAISettingsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DefenderForAISettingsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return DefenderForAISettingsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DefenderForAISettingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *DefenderForAISettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/defenderForAISettings"
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
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefenderForAISettingsClient) listHandleResponse(resp *http.Response) (DefenderForAISettingsClientListResponse, error) {
	result := DefenderForAISettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderForAISettingResult); err != nil {
		return DefenderForAISettingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates the specified Defender for AI setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of Cognitive Services account.
//   - defenderForAISettingName - The name of the defender for AI setting.
//   - defenderForAISettings - Properties describing the Defender for AI setting.
//   - options - DefenderForAISettingsClientUpdateOptions contains the optional parameters for the DefenderForAISettingsClient.Update
//     method.
func (client *DefenderForAISettingsClient) Update(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings DefenderForAISetting, options *DefenderForAISettingsClientUpdateOptions) (DefenderForAISettingsClientUpdateResponse, error) {
	var err error
	const operationName = "DefenderForAISettingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, defenderForAISettingName, defenderForAISettings, options)
	if err != nil {
		return DefenderForAISettingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefenderForAISettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefenderForAISettingsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DefenderForAISettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, defenderForAISettingName string, defenderForAISettings DefenderForAISetting, options *DefenderForAISettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/defenderForAISettings/{defenderForAISettingName}"
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
	if defenderForAISettingName == "" {
		return nil, errors.New("parameter defenderForAISettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{defenderForAISettingName}", url.PathEscape(defenderForAISettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, defenderForAISettings); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DefenderForAISettingsClient) updateHandleResponse(resp *http.Response) (DefenderForAISettingsClientUpdateResponse, error) {
	result := DefenderForAISettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefenderForAISetting); err != nil {
		return DefenderForAISettingsClientUpdateResponse{}, err
	}
	return result, nil
}
