//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package sapmonitors

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

// AlertTemplatesClient contains the methods for the AlertTemplates group.
// Don't use this type directly, use NewAlertTemplatesClient() instead.
type AlertTemplatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAlertTemplatesClient creates a new instance of AlertTemplatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertTemplatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertTemplatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets properties of an alert template for the specified subscription, resource group, SAP monitor name, and resource
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - alertTemplateName - SAP monitor alert template resource name.
//   - options - AlertTemplatesClientGetOptions contains the optional parameters for the AlertTemplatesClient.Get method.
func (client *AlertTemplatesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, alertTemplateName string, options *AlertTemplatesClientGetOptions) (AlertTemplatesClientGetResponse, error) {
	var err error
	const operationName = "AlertTemplatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, alertTemplateName, options)
	if err != nil {
		return AlertTemplatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertTemplatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AlertTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, alertTemplateName string, options *AlertTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/alertTemplates/{alertTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if alertTemplateName == "" {
		return nil, errors.New("parameter alertTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertTemplateName}", url.PathEscape(alertTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertTemplatesClient) getHandleResponse(resp *http.Response) (AlertTemplatesClientGetResponse, error) {
	result := AlertTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertTemplate); err != nil {
		return AlertTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of alert templates in the specified SAP monitor. The operations returns various properties of
// each alert template.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - options - AlertTemplatesClientListOptions contains the optional parameters for the AlertTemplatesClient.NewListPager method.
func (client *AlertTemplatesClient) NewListPager(resourceGroupName string, monitorName string, options *AlertTemplatesClientListOptions) *runtime.Pager[AlertTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertTemplatesClientListResponse]{
		More: func(page AlertTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertTemplatesClientListResponse) (AlertTemplatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AlertTemplatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			}, nil)
			if err != nil {
				return AlertTemplatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AlertTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *AlertTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/alertTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	if options != nil && options.ProviderType != nil {
		reqQP.Set("providerType", *options.ProviderType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertTemplatesClient) listHandleResponse(resp *http.Response) (AlertTemplatesClientListResponse, error) {
	result := AlertTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertTemplateListResult); err != nil {
		return AlertTemplatesClientListResponse{}, err
	}
	return result, nil
}
