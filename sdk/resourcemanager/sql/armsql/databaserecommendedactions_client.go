// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// DatabaseRecommendedActionsClient contains the methods for the DatabaseRecommendedActions group.
// Don't use this type directly, use NewDatabaseRecommendedActionsClient() instead.
type DatabaseRecommendedActionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseRecommendedActionsClient creates a new instance of DatabaseRecommendedActionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseRecommendedActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseRecommendedActionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseRecommendedActionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a database recommended action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - advisorName - The name of the Database Advisor.
//   - recommendedActionName - The name of Database Recommended Action.
//   - options - DatabaseRecommendedActionsClientGetOptions contains the optional parameters for the DatabaseRecommendedActionsClient.Get
//     method.
func (client *DatabaseRecommendedActionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, options *DatabaseRecommendedActionsClientGetOptions) (DatabaseRecommendedActionsClientGetResponse, error) {
	var err error
	const operationName = "DatabaseRecommendedActionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName, options)
	if err != nil {
		return DatabaseRecommendedActionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseRecommendedActionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseRecommendedActionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseRecommendedActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, _ *DatabaseRecommendedActionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if recommendedActionName == "" {
		return nil, errors.New("parameter recommendedActionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendedActionName}", url.PathEscape(recommendedActionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseRecommendedActionsClient) getHandleResponse(resp *http.Response) (DatabaseRecommendedActionsClientGetResponse, error) {
	result := DatabaseRecommendedActionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedAction); err != nil {
		return DatabaseRecommendedActionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByDatabaseAdvisor - Gets list of Database Recommended Actions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - advisorName - The name of the Database Advisor.
//   - options - DatabaseRecommendedActionsClientListByDatabaseAdvisorOptions contains the optional parameters for the DatabaseRecommendedActionsClient.ListByDatabaseAdvisor
//     method.
func (client *DatabaseRecommendedActionsClient) ListByDatabaseAdvisor(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, options *DatabaseRecommendedActionsClientListByDatabaseAdvisorOptions) (DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse, error) {
	var err error
	const operationName = "DatabaseRecommendedActionsClient.ListByDatabaseAdvisor"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByDatabaseAdvisorCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, options)
	if err != nil {
		return DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse{}, err
	}
	resp, err := client.listByDatabaseAdvisorHandleResponse(httpResp)
	return resp, err
}

// listByDatabaseAdvisorCreateRequest creates the ListByDatabaseAdvisor request.
func (client *DatabaseRecommendedActionsClient) listByDatabaseAdvisorCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, _ *DatabaseRecommendedActionsClientListByDatabaseAdvisorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseAdvisorHandleResponse handles the ListByDatabaseAdvisor response.
func (client *DatabaseRecommendedActionsClient) listByDatabaseAdvisorHandleResponse(resp *http.Response) (DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse, error) {
	result := DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedActionArray); err != nil {
		return DatabaseRecommendedActionsClientListByDatabaseAdvisorResponse{}, err
	}
	return result, nil
}

// Update - Updates a database recommended action.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - advisorName - The name of the Database Advisor.
//   - recommendedActionName - The name of Database Recommended Action.
//   - parameters - The requested recommended action resource state.
//   - options - DatabaseRecommendedActionsClientUpdateOptions contains the optional parameters for the DatabaseRecommendedActionsClient.Update
//     method.
func (client *DatabaseRecommendedActionsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction, options *DatabaseRecommendedActionsClientUpdateOptions) (DatabaseRecommendedActionsClientUpdateResponse, error) {
	var err error
	const operationName = "DatabaseRecommendedActionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, advisorName, recommendedActionName, parameters, options)
	if err != nil {
		return DatabaseRecommendedActionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseRecommendedActionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseRecommendedActionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DatabaseRecommendedActionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, recommendedActionName string, parameters RecommendedAction, _ *DatabaseRecommendedActionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}/recommendedActions/{recommendedActionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if recommendedActionName == "" {
		return nil, errors.New("parameter recommendedActionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendedActionName}", url.PathEscape(recommendedActionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DatabaseRecommendedActionsClient) updateHandleResponse(resp *http.Response) (DatabaseRecommendedActionsClientUpdateResponse, error) {
	result := DatabaseRecommendedActionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedAction); err != nil {
		return DatabaseRecommendedActionsClientUpdateResponse{}, err
	}
	return result, nil
}
