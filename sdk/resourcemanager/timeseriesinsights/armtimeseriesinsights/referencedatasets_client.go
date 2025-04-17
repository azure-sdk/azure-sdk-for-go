// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtimeseriesinsights

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

// ReferenceDataSetsClient contains the methods for the ReferenceDataSets group.
// Don't use this type directly, use NewReferenceDataSetsClient() instead.
type ReferenceDataSetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReferenceDataSetsClient creates a new instance of ReferenceDataSetsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReferenceDataSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReferenceDataSetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReferenceDataSetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a reference data set in the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-31-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - referenceDataSetName - Name of the reference data set.
//   - parameters - Parameters for creating a reference data set.
//   - options - ReferenceDataSetsClientCreateOrUpdateOptions contains the optional parameters for the ReferenceDataSetsClient.CreateOrUpdate
//     method.
func (client *ReferenceDataSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, parameters ReferenceDataSetCreateOrUpdateParameters, options *ReferenceDataSetsClientCreateOrUpdateOptions) (ReferenceDataSetsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ReferenceDataSetsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, referenceDataSetName, parameters, options)
	if err != nil {
		return ReferenceDataSetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReferenceDataSetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ReferenceDataSetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ReferenceDataSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, parameters ReferenceDataSetCreateOrUpdateParameters, _ *ReferenceDataSetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if referenceDataSetName == "" {
		return nil, errors.New("parameter referenceDataSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{referenceDataSetName}", url.PathEscape(referenceDataSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ReferenceDataSetsClient) createOrUpdateHandleResponse(resp *http.Response) (ReferenceDataSetsClientCreateOrUpdateResponse, error) {
	result := ReferenceDataSetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReferenceDataSetResource); err != nil {
		return ReferenceDataSetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the reference data set with the specified name in the specified subscription, resource group, and environment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-31-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - referenceDataSetName - The name of the Time Series Insights reference data set associated with the specified environment.
//   - options - ReferenceDataSetsClientDeleteOptions contains the optional parameters for the ReferenceDataSetsClient.Delete
//     method.
func (client *ReferenceDataSetsClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, options *ReferenceDataSetsClientDeleteOptions) (ReferenceDataSetsClientDeleteResponse, error) {
	var err error
	const operationName = "ReferenceDataSetsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, referenceDataSetName, options)
	if err != nil {
		return ReferenceDataSetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReferenceDataSetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ReferenceDataSetsClientDeleteResponse{}, err
	}
	return ReferenceDataSetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReferenceDataSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, _ *ReferenceDataSetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if referenceDataSetName == "" {
		return nil, errors.New("parameter referenceDataSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{referenceDataSetName}", url.PathEscape(referenceDataSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the reference data set with the specified name in the specified environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-31-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - referenceDataSetName - The name of the Time Series Insights reference data set associated with the specified environment.
//   - options - ReferenceDataSetsClientGetOptions contains the optional parameters for the ReferenceDataSetsClient.Get method.
func (client *ReferenceDataSetsClient) Get(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, options *ReferenceDataSetsClientGetOptions) (ReferenceDataSetsClientGetResponse, error) {
	var err error
	const operationName = "ReferenceDataSetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, referenceDataSetName, options)
	if err != nil {
		return ReferenceDataSetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReferenceDataSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReferenceDataSetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReferenceDataSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, _ *ReferenceDataSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if referenceDataSetName == "" {
		return nil, errors.New("parameter referenceDataSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{referenceDataSetName}", url.PathEscape(referenceDataSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReferenceDataSetsClient) getHandleResponse(resp *http.Response) (ReferenceDataSetsClientGetResponse, error) {
	result := ReferenceDataSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReferenceDataSetResource); err != nil {
		return ReferenceDataSetsClientGetResponse{}, err
	}
	return result, nil
}

// ListByEnvironment - Lists all the available reference data sets associated with the subscription and within the specified
// resource group and environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-31-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - options - ReferenceDataSetsClientListByEnvironmentOptions contains the optional parameters for the ReferenceDataSetsClient.ListByEnvironment
//     method.
func (client *ReferenceDataSetsClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string, options *ReferenceDataSetsClientListByEnvironmentOptions) (ReferenceDataSetsClientListByEnvironmentResponse, error) {
	var err error
	const operationName = "ReferenceDataSetsClient.ListByEnvironment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByEnvironmentCreateRequest(ctx, resourceGroupName, environmentName, options)
	if err != nil {
		return ReferenceDataSetsClientListByEnvironmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReferenceDataSetsClientListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReferenceDataSetsClientListByEnvironmentResponse{}, err
	}
	resp, err := client.listByEnvironmentHandleResponse(httpResp)
	return resp, err
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *ReferenceDataSetsClient) listByEnvironmentCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, _ *ReferenceDataSetsClientListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *ReferenceDataSetsClient) listByEnvironmentHandleResponse(resp *http.Response) (ReferenceDataSetsClientListByEnvironmentResponse, error) {
	result := ReferenceDataSetsClientListByEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReferenceDataSetListResponse); err != nil {
		return ReferenceDataSetsClientListByEnvironmentResponse{}, err
	}
	return result, nil
}

// Update - Updates the reference data set with the specified name in the specified subscription, resource group, and environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-31-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - environmentName - The name of the Time Series Insights environment associated with the specified resource group.
//   - referenceDataSetName - The name of the Time Series Insights reference data set associated with the specified environment.
//   - referenceDataSetUpdateParameters - Request object that contains the updated information for the reference data set.
//   - options - ReferenceDataSetsClientUpdateOptions contains the optional parameters for the ReferenceDataSetsClient.Update
//     method.
func (client *ReferenceDataSetsClient) Update(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, referenceDataSetUpdateParameters ReferenceDataSetUpdateParameters, options *ReferenceDataSetsClientUpdateOptions) (ReferenceDataSetsClientUpdateResponse, error) {
	var err error
	const operationName = "ReferenceDataSetsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, environmentName, referenceDataSetName, referenceDataSetUpdateParameters, options)
	if err != nil {
		return ReferenceDataSetsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReferenceDataSetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReferenceDataSetsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ReferenceDataSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, referenceDataSetName string, referenceDataSetUpdateParameters ReferenceDataSetUpdateParameters, _ *ReferenceDataSetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if referenceDataSetName == "" {
		return nil, errors.New("parameter referenceDataSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{referenceDataSetName}", url.PathEscape(referenceDataSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, referenceDataSetUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ReferenceDataSetsClient) updateHandleResponse(resp *http.Response) (ReferenceDataSetsClientUpdateResponse, error) {
	result := ReferenceDataSetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReferenceDataSetResource); err != nil {
		return ReferenceDataSetsClientUpdateResponse{}, err
	}
	return result, nil
}
