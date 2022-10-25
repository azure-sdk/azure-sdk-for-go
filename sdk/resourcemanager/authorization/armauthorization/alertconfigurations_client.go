//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// AlertConfigurationsClient contains the methods for the AlertConfigurations group.
// Don't use this type directly, use NewAlertConfigurationsClient() instead.
type AlertConfigurationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAlertConfigurationsClient creates a new instance of AlertConfigurationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAlertConfigurationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertConfigurationsClient, error) {
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
	client := &AlertConfigurationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get the specified alert configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-07
// scope - The scope of the alert configuration. The scope can be any REST resource instance. For example, use '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/'
// for a subscription,
// '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource
// group, and
// '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
// for a resource.
// alertID - The name of the alert configuration to get.
// options - AlertConfigurationsClientGetOptions contains the optional parameters for the AlertConfigurationsClient.Get method.
func (client *AlertConfigurationsClient) Get(ctx context.Context, scope string, alertID string, options *AlertConfigurationsClientGetOptions) (AlertConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, alertID, options)
	if err != nil {
		return AlertConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AlertConfigurationsClient) getCreateRequest(ctx context.Context, scope string, alertID string, options *AlertConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementAlerts/alertConfigurations/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertConfigurationsClient) getHandleResponse(resp *http.Response) (AlertConfigurationsClientGetResponse, error) {
	result := AlertConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertConfiguration); err != nil {
		return AlertConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets alert configurations for a resource scope.
// Generated from API version 2022-01-07
// scope - The scope of the alert configuration.
// options - AlertConfigurationsClientListForScopeOptions contains the optional parameters for the AlertConfigurationsClient.ListForScope
// method.
func (client *AlertConfigurationsClient) NewListForScopePager(scope string, options *AlertConfigurationsClientListForScopeOptions) *runtime.Pager[AlertConfigurationsClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertConfigurationsClientListForScopeResponse]{
		More: func(page AlertConfigurationsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertConfigurationsClientListForScopeResponse) (AlertConfigurationsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertConfigurationsClientListForScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertConfigurationsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertConfigurationsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *AlertConfigurationsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *AlertConfigurationsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementAlerts/alertConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *AlertConfigurationsClient) listForScopeHandleResponse(resp *http.Response) (AlertConfigurationsClientListForScopeResponse, error) {
	result := AlertConfigurationsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertConfigurationListResult); err != nil {
		return AlertConfigurationsClientListForScopeResponse{}, err
	}
	return result, nil
}

// Update - Update an alert configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-07
// scope - The scope of the alert configuration.
// alertID - The name of the alert configuration to update.
// parameters - Parameters for the alert configuration.
// options - AlertConfigurationsClientUpdateOptions contains the optional parameters for the AlertConfigurationsClient.Update
// method.
func (client *AlertConfigurationsClient) Update(ctx context.Context, scope string, alertID string, parameters AlertConfiguration, options *AlertConfigurationsClientUpdateOptions) (AlertConfigurationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, scope, alertID, parameters, options)
	if err != nil {
		return AlertConfigurationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertConfigurationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertConfigurationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AlertConfigurationsClient) updateCreateRequest(ctx context.Context, scope string, alertID string, parameters AlertConfiguration, options *AlertConfigurationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementAlerts/alertConfigurations/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AlertConfigurationsClient) updateHandleResponse(resp *http.Response) (AlertConfigurationsClientUpdateResponse, error) {
	result := AlertConfigurationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertConfiguration); err != nil {
		return AlertConfigurationsClientUpdateResponse{}, err
	}
	return result, nil
}
