//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armalertsmanagement

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

// AlertProcessingRulesClient contains the methods for the AlertProcessingRules group.
// Don't use this type directly, use NewAlertProcessingRulesClient() instead.
type AlertProcessingRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAlertProcessingRulesClient creates a new instance of AlertProcessingRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertProcessingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertProcessingRulesClient, error) {
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
	client := &AlertProcessingRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an alert processing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-08
//   - resourceGroupName - Resource group name where the resource is created.
//   - alertProcessingRuleName - The name of the alert processing rule that needs to be created/updated.
//   - alertProcessingRule - Alert processing rule to be created/updated.
//   - options - AlertProcessingRulesClientCreateOrUpdateOptions contains the optional parameters for the AlertProcessingRulesClient.CreateOrUpdate
//     method.
func (client *AlertProcessingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRule AlertProcessingRule, options *AlertProcessingRulesClientCreateOrUpdateOptions) (AlertProcessingRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, alertProcessingRuleName, alertProcessingRule, options)
	if err != nil {
		return AlertProcessingRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertProcessingRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AlertProcessingRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AlertProcessingRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRule AlertProcessingRule, options *AlertProcessingRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{alertProcessingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if alertProcessingRuleName == "" {
		return nil, errors.New("parameter alertProcessingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertProcessingRuleName}", url.PathEscape(alertProcessingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, alertProcessingRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AlertProcessingRulesClient) createOrUpdateHandleResponse(resp *http.Response) (AlertProcessingRulesClientCreateOrUpdateResponse, error) {
	result := AlertProcessingRulesClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertProcessingRule); err != nil {
		return AlertProcessingRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an alert processing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-08
//   - resourceGroupName - Resource group name where the resource is created.
//   - alertProcessingRuleName - The name of the alert processing rule that needs to be deleted.
//   - options - AlertProcessingRulesClientDeleteOptions contains the optional parameters for the AlertProcessingRulesClient.Delete
//     method.
func (client *AlertProcessingRulesClient) Delete(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *AlertProcessingRulesClientDeleteOptions) (AlertProcessingRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, alertProcessingRuleName, options)
	if err != nil {
		return AlertProcessingRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertProcessingRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AlertProcessingRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *AlertProcessingRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *AlertProcessingRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{alertProcessingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if alertProcessingRuleName == "" {
		return nil, errors.New("parameter alertProcessingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertProcessingRuleName}", url.PathEscape(alertProcessingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *AlertProcessingRulesClient) deleteHandleResponse(resp *http.Response) (AlertProcessingRulesClientDeleteResponse, error) {
	result := AlertProcessingRulesClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// GetByName - Get an alert processing rule by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-08
//   - resourceGroupName - Resource group name where the resource is created.
//   - alertProcessingRuleName - The name of the alert processing rule that needs to be fetched.
//   - options - AlertProcessingRulesClientGetByNameOptions contains the optional parameters for the AlertProcessingRulesClient.GetByName
//     method.
func (client *AlertProcessingRulesClient) GetByName(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *AlertProcessingRulesClientGetByNameOptions) (AlertProcessingRulesClientGetByNameResponse, error) {
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, alertProcessingRuleName, options)
	if err != nil {
		return AlertProcessingRulesClientGetByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertProcessingRulesClientGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertProcessingRulesClientGetByNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByNameHandleResponse(resp)
}

// getByNameCreateRequest creates the GetByName request.
func (client *AlertProcessingRulesClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *AlertProcessingRulesClientGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{alertProcessingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if alertProcessingRuleName == "" {
		return nil, errors.New("parameter alertProcessingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertProcessingRuleName}", url.PathEscape(alertProcessingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *AlertProcessingRulesClient) getByNameHandleResponse(resp *http.Response) (AlertProcessingRulesClientGetByNameResponse, error) {
	result := AlertProcessingRulesClientGetByNameResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertProcessingRule); err != nil {
		return AlertProcessingRulesClientGetByNameResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all alert processing rules in a resource group.
//
// Generated from API version 2021-08-08
//   - resourceGroupName - Resource group name where the resource is created.
//   - options - AlertProcessingRulesClientListByResourceGroupOptions contains the optional parameters for the AlertProcessingRulesClient.NewListByResourceGroupPager
//     method.
func (client *AlertProcessingRulesClient) NewListByResourceGroupPager(resourceGroupName string, options *AlertProcessingRulesClientListByResourceGroupOptions) *runtime.Pager[AlertProcessingRulesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertProcessingRulesClientListByResourceGroupResponse]{
		More: func(page AlertProcessingRulesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertProcessingRulesClientListByResourceGroupResponse) (AlertProcessingRulesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertProcessingRulesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertProcessingRulesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertProcessingRulesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AlertProcessingRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AlertProcessingRulesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AlertProcessingRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (AlertProcessingRulesClientListByResourceGroupResponse, error) {
	result := AlertProcessingRulesClientListByResourceGroupResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertProcessingRulesList); err != nil {
		return AlertProcessingRulesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all alert processing rules in a subscription.
//
// Generated from API version 2021-08-08
//   - options - AlertProcessingRulesClientListBySubscriptionOptions contains the optional parameters for the AlertProcessingRulesClient.NewListBySubscriptionPager
//     method.
func (client *AlertProcessingRulesClient) NewListBySubscriptionPager(options *AlertProcessingRulesClientListBySubscriptionOptions) *runtime.Pager[AlertProcessingRulesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertProcessingRulesClientListBySubscriptionResponse]{
		More: func(page AlertProcessingRulesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertProcessingRulesClientListBySubscriptionResponse) (AlertProcessingRulesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertProcessingRulesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AlertProcessingRulesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertProcessingRulesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AlertProcessingRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AlertProcessingRulesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/actionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AlertProcessingRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (AlertProcessingRulesClientListBySubscriptionResponse, error) {
	result := AlertProcessingRulesClientListBySubscriptionResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertProcessingRulesList); err != nil {
		return AlertProcessingRulesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Enable, disable, or update tags for an alert processing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-08
//   - resourceGroupName - Resource group name where the resource is created.
//   - alertProcessingRuleName - The name that needs to be updated.
//   - alertProcessingRulePatch - Parameters supplied to the operation.
//   - options - AlertProcessingRulesClientUpdateOptions contains the optional parameters for the AlertProcessingRulesClient.Update
//     method.
func (client *AlertProcessingRulesClient) Update(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRulePatch PatchObject, options *AlertProcessingRulesClientUpdateOptions) (AlertProcessingRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, alertProcessingRuleName, alertProcessingRulePatch, options)
	if err != nil {
		return AlertProcessingRulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AlertProcessingRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AlertProcessingRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AlertProcessingRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRulePatch PatchObject, options *AlertProcessingRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{alertProcessingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if alertProcessingRuleName == "" {
		return nil, errors.New("parameter alertProcessingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertProcessingRuleName}", url.PathEscape(alertProcessingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, alertProcessingRulePatch)
}

// updateHandleResponse handles the Update response.
func (client *AlertProcessingRulesClient) updateHandleResponse(resp *http.Response) (AlertProcessingRulesClientUpdateResponse, error) {
	result := AlertProcessingRulesClientUpdateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertProcessingRule); err != nil {
		return AlertProcessingRulesClientUpdateResponse{}, err
	}
	return result, nil
}
