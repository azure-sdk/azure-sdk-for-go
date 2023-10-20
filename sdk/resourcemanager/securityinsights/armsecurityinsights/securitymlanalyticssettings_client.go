//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// SecurityMLAnalyticsSettingsClient contains the methods for the SecurityMLAnalyticsSettings group.
// Don't use this type directly, use NewSecurityMLAnalyticsSettingsClient() instead.
type SecurityMLAnalyticsSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityMLAnalyticsSettingsClient creates a new instance of SecurityMLAnalyticsSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityMLAnalyticsSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityMLAnalyticsSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".SecurityMLAnalyticsSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityMLAnalyticsSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the Security ML Analytics Settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - settingsResourceName - Security ML Analytics Settings resource name
//   - securityMLAnalyticsSetting - The security ML Analytics setting
//   - options - SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.CreateOrUpdate
//     method.
func (client *SecurityMLAnalyticsSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, securityMLAnalyticsSetting SecurityMLAnalyticsSettingClassification, options *SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions) (SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, settingsResourceName, securityMLAnalyticsSetting, options)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityMLAnalyticsSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, securityMLAnalyticsSetting SecurityMLAnalyticsSettingClassification, options *SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, securityMLAnalyticsSetting); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SecurityMLAnalyticsSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse, error) {
	result := SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the Security ML Analytics Settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - settingsResourceName - Security ML Analytics Settings resource name
//   - options - SecurityMLAnalyticsSettingsClientDeleteOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.Delete
//     method.
func (client *SecurityMLAnalyticsSettingsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *SecurityMLAnalyticsSettingsClientDeleteOptions) (SecurityMLAnalyticsSettingsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, settingsResourceName, options)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SecurityMLAnalyticsSettingsClientDeleteResponse{}, err
	}
	return SecurityMLAnalyticsSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityMLAnalyticsSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *SecurityMLAnalyticsSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
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

// Get - Gets the Security ML Analytics Settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - settingsResourceName - Security ML Analytics Settings resource name
//   - options - SecurityMLAnalyticsSettingsClientGetOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.Get
//     method.
func (client *SecurityMLAnalyticsSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *SecurityMLAnalyticsSettingsClientGetOptions) (SecurityMLAnalyticsSettingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, settingsResourceName, options)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityMLAnalyticsSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityMLAnalyticsSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecurityMLAnalyticsSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *SecurityMLAnalyticsSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings/{settingsResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if settingsResourceName == "" {
		return nil, errors.New("parameter settingsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsResourceName}", url.PathEscape(settingsResourceName))
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
func (client *SecurityMLAnalyticsSettingsClient) getHandleResponse(resp *http.Response) (SecurityMLAnalyticsSettingsClientGetResponse, error) {
	result := SecurityMLAnalyticsSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return SecurityMLAnalyticsSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Security ML Analytics Settings.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - SecurityMLAnalyticsSettingsClientListOptions contains the optional parameters for the SecurityMLAnalyticsSettingsClient.NewListPager
//     method.
func (client *SecurityMLAnalyticsSettingsClient) NewListPager(resourceGroupName string, workspaceName string, options *SecurityMLAnalyticsSettingsClientListOptions) *runtime.Pager[SecurityMLAnalyticsSettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityMLAnalyticsSettingsClientListResponse]{
		More: func(page SecurityMLAnalyticsSettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityMLAnalyticsSettingsClientListResponse) (SecurityMLAnalyticsSettingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityMLAnalyticsSettingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecurityMLAnalyticsSettingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityMLAnalyticsSettingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SecurityMLAnalyticsSettingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SecurityMLAnalyticsSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/securityMLAnalyticsSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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
func (client *SecurityMLAnalyticsSettingsClient) listHandleResponse(resp *http.Response) (SecurityMLAnalyticsSettingsClientListResponse, error) {
	result := SecurityMLAnalyticsSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityMLAnalyticsSettingsList); err != nil {
		return SecurityMLAnalyticsSettingsClientListResponse{}, err
	}
	return result, nil
}
