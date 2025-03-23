// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// MetricAlertsClient contains the methods for the MetricAlerts group.
// Don't use this type directly, use NewMetricAlertsClient() instead.
type MetricAlertsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetricAlertsClient creates a new instance of MetricAlertsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetricAlertsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricAlertsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetricAlertsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group.
//   - ruleName - The name of the alert rule.
//   - parameters - The metric alert rule resource.
//   - options - MetricAlertsClientCreateOrUpdateOptions contains the optional parameters for the MetricAlertsClient.CreateOrUpdate
//     method.
func (client *MetricAlertsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource, options *MetricAlertsClientCreateOrUpdateOptions) (MetricAlertsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MetricAlertsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResource, _ *MetricAlertsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MetricAlertsClient) createOrUpdateHandleResponse(resp *http.Response) (MetricAlertsClientCreateOrUpdateResponse, error) {
	result := MetricAlertsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group.
//   - ruleName - The name of the alert rule.
//   - options - MetricAlertsClientDeleteOptions contains the optional parameters for the MetricAlertsClient.Delete method.
func (client *MetricAlertsClient) Delete(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientDeleteOptions) (MetricAlertsClientDeleteResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientDeleteResponse{}, err
	}
	return MetricAlertsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MetricAlertsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, _ *MetricAlertsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve an alert rule definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group.
//   - ruleName - The name of the alert rule.
//   - options - MetricAlertsClientGetOptions contains the optional parameters for the MetricAlertsClient.Get method.
func (client *MetricAlertsClient) Get(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsClientGetOptions) (MetricAlertsClientGetResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MetricAlertsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, _ *MetricAlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MetricAlertsClient) getHandleResponse(resp *http.Response) (MetricAlertsClientGetResponse, error) {
	result := MetricAlertsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieve alert rule definitions in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group.
//   - options - MetricAlertsClientListByResourceGroupOptions contains the optional parameters for the MetricAlertsClient.ListByResourceGroup
//     method.
func (client *MetricAlertsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *MetricAlertsClientListByResourceGroupOptions) (MetricAlertsClientListByResourceGroupResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.ListByResourceGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return MetricAlertsClientListByResourceGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientListByResourceGroupResponse{}, err
	}
	resp, err := client.listByResourceGroupHandleResponse(httpResp)
	return resp, err
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MetricAlertsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *MetricAlertsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts"
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
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MetricAlertsClient) listByResourceGroupHandleResponse(resp *http.Response) (MetricAlertsClientListByResourceGroupResponse, error) {
	result := MetricAlertsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResourceCollection); err != nil {
		return MetricAlertsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieve alert rule definitions in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - options - MetricAlertsClientListBySubscriptionOptions contains the optional parameters for the MetricAlertsClient.ListBySubscription
//     method.
func (client *MetricAlertsClient) ListBySubscription(ctx context.Context, options *MetricAlertsClientListBySubscriptionOptions) (MetricAlertsClientListBySubscriptionResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.ListBySubscription"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return MetricAlertsClientListBySubscriptionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientListBySubscriptionResponse{}, err
	}
	resp, err := client.listBySubscriptionHandleResponse(httpResp)
	return resp, err
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MetricAlertsClient) listBySubscriptionCreateRequest(ctx context.Context, _ *MetricAlertsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.insights/metricAlerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MetricAlertsClient) listBySubscriptionHandleResponse(resp *http.Response) (MetricAlertsClientListBySubscriptionResponse, error) {
	result := MetricAlertsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResourceCollection); err != nil {
		return MetricAlertsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an alert rule. Only tags, enabled, and actions fields can be updated.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group.
//   - ruleName - The name of the alert rule.
//   - parameters - The metric alert rule resource for patch operations.
//   - options - MetricAlertsClientUpdateOptions contains the optional parameters for the MetricAlertsClient.Update method.
func (client *MetricAlertsClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch, options *MetricAlertsClientUpdateOptions) (MetricAlertsClientUpdateResponse, error) {
	var err error
	const operationName = "MetricAlertsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ruleName, parameters, options)
	if err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricAlertsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MetricAlertsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, parameters MetricAlertResourcePatch, _ *MetricAlertsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/metricAlerts/{ruleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MetricAlertsClient) updateHandleResponse(resp *http.Response) (MetricAlertsClientUpdateResponse, error) {
	result := MetricAlertsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertResource); err != nil {
		return MetricAlertsClientUpdateResponse{}, err
	}
	return result, nil
}
