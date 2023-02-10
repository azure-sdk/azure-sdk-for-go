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

// PrometheusRuleGroupsClient contains the methods for the PrometheusRuleGroups group.
// Don't use this type directly, use NewPrometheusRuleGroupsClient() instead.
type PrometheusRuleGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrometheusRuleGroupsClient creates a new instance of PrometheusRuleGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrometheusRuleGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrometheusRuleGroupsClient, error) {
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
	client := &PrometheusRuleGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Prometheus rule group definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-22-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleGroupName - The name of the rule group.
//   - parameters - The parameters of the rule group to create or update.
//   - options - PrometheusRuleGroupsClientCreateOrUpdateOptions contains the optional parameters for the PrometheusRuleGroupsClient.CreateOrUpdate
//     method.
func (client *PrometheusRuleGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleGroupName string, parameters PrometheusRuleGroupResource, options *PrometheusRuleGroupsClientCreateOrUpdateOptions) (PrometheusRuleGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleGroupName, parameters, options)
	if err != nil {
		return PrometheusRuleGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrometheusRuleGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrometheusRuleGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrometheusRuleGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleGroupName string, parameters PrometheusRuleGroupResource, options *PrometheusRuleGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/prometheusRuleGroups/{ruleGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PrometheusRuleGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (PrometheusRuleGroupsClientCreateOrUpdateResponse, error) {
	result := PrometheusRuleGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrometheusRuleGroupResource); err != nil {
		return PrometheusRuleGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Prometheus rule group definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-22-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleGroupName - The name of the rule group.
//   - options - PrometheusRuleGroupsClientDeleteOptions contains the optional parameters for the PrometheusRuleGroupsClient.Delete
//     method.
func (client *PrometheusRuleGroupsClient) Delete(ctx context.Context, resourceGroupName string, ruleGroupName string, options *PrometheusRuleGroupsClientDeleteOptions) (PrometheusRuleGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleGroupName, options)
	if err != nil {
		return PrometheusRuleGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrometheusRuleGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PrometheusRuleGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PrometheusRuleGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrometheusRuleGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleGroupName string, options *PrometheusRuleGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/prometheusRuleGroups/{ruleGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve a Prometheus rule group definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-22-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleGroupName - The name of the rule group.
//   - options - PrometheusRuleGroupsClientGetOptions contains the optional parameters for the PrometheusRuleGroupsClient.Get
//     method.
func (client *PrometheusRuleGroupsClient) Get(ctx context.Context, resourceGroupName string, ruleGroupName string, options *PrometheusRuleGroupsClientGetOptions) (PrometheusRuleGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleGroupName, options)
	if err != nil {
		return PrometheusRuleGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrometheusRuleGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrometheusRuleGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrometheusRuleGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleGroupName string, options *PrometheusRuleGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/prometheusRuleGroups/{ruleGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrometheusRuleGroupsClient) getHandleResponse(resp *http.Response) (PrometheusRuleGroupsClientGetResponse, error) {
	result := PrometheusRuleGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrometheusRuleGroupResource); err != nil {
		return PrometheusRuleGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve Prometheus rule group definitions in a resource group.
//
// Generated from API version 2021-07-22-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - PrometheusRuleGroupsClientListByResourceGroupOptions contains the optional parameters for the PrometheusRuleGroupsClient.NewListByResourceGroupPager
//     method.
func (client *PrometheusRuleGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *PrometheusRuleGroupsClientListByResourceGroupOptions) *runtime.Pager[PrometheusRuleGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrometheusRuleGroupsClientListByResourceGroupResponse]{
		More: func(page PrometheusRuleGroupsClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrometheusRuleGroupsClientListByResourceGroupResponse) (PrometheusRuleGroupsClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return PrometheusRuleGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrometheusRuleGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrometheusRuleGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrometheusRuleGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrometheusRuleGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/prometheusRuleGroups"
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
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrometheusRuleGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (PrometheusRuleGroupsClientListByResourceGroupResponse, error) {
	result := PrometheusRuleGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrometheusRuleGroupResourceCollection); err != nil {
		return PrometheusRuleGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve Prometheus rule group definitions in a subscription.
//
// Generated from API version 2021-07-22-preview
//   - options - PrometheusRuleGroupsClientListBySubscriptionOptions contains the optional parameters for the PrometheusRuleGroupsClient.NewListBySubscriptionPager
//     method.
func (client *PrometheusRuleGroupsClient) NewListBySubscriptionPager(options *PrometheusRuleGroupsClientListBySubscriptionOptions) *runtime.Pager[PrometheusRuleGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrometheusRuleGroupsClientListBySubscriptionResponse]{
		More: func(page PrometheusRuleGroupsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrometheusRuleGroupsClientListBySubscriptionResponse) (PrometheusRuleGroupsClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return PrometheusRuleGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrometheusRuleGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrometheusRuleGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PrometheusRuleGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PrometheusRuleGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/prometheusRuleGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PrometheusRuleGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (PrometheusRuleGroupsClientListBySubscriptionResponse, error) {
	result := PrometheusRuleGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrometheusRuleGroupResourceCollection); err != nil {
		return PrometheusRuleGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an Prometheus rule group definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-22-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - ruleGroupName - The name of the rule group.
//   - parameters - The parameters of the rule group to update.
//   - options - PrometheusRuleGroupsClientUpdateOptions contains the optional parameters for the PrometheusRuleGroupsClient.Update
//     method.
func (client *PrometheusRuleGroupsClient) Update(ctx context.Context, resourceGroupName string, ruleGroupName string, parameters PrometheusRuleGroupResourcePatch, options *PrometheusRuleGroupsClientUpdateOptions) (PrometheusRuleGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleGroupName, parameters, options)
	if err != nil {
		return PrometheusRuleGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrometheusRuleGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrometheusRuleGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PrometheusRuleGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleGroupName string, parameters PrometheusRuleGroupResourcePatch, options *PrometheusRuleGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/prometheusRuleGroups/{ruleGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleGroupName == "" {
		return nil, errors.New("parameter ruleGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleGroupName}", url.PathEscape(ruleGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-22-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PrometheusRuleGroupsClient) updateHandleResponse(resp *http.Response) (PrometheusRuleGroupsClientUpdateResponse, error) {
	result := PrometheusRuleGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrometheusRuleGroupResource); err != nil {
		return PrometheusRuleGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
