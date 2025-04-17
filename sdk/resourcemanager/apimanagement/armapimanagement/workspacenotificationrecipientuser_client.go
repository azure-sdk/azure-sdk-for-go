// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// WorkspaceNotificationRecipientUserClient contains the methods for the WorkspaceNotificationRecipientUser group.
// Don't use this type directly, use NewWorkspaceNotificationRecipientUserClient() instead.
type WorkspaceNotificationRecipientUserClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceNotificationRecipientUserClient creates a new instance of WorkspaceNotificationRecipientUserClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceNotificationRecipientUserClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceNotificationRecipientUserClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceNotificationRecipientUserClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEntityExists - Determine if the Notification Recipient User is subscribed to the notification.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceNotificationRecipientUserClientCheckEntityExistsOptions contains the optional parameters for the WorkspaceNotificationRecipientUserClient.CheckEntityExists
//     method.
func (client *WorkspaceNotificationRecipientUserClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, options *WorkspaceNotificationRecipientUserClientCheckEntityExistsOptions) (WorkspaceNotificationRecipientUserClientCheckEntityExistsResponse, error) {
	var err error
	const operationName = "WorkspaceNotificationRecipientUserClient.CheckEntityExists"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, userID, options)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientCheckEntityExistsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceNotificationRecipientUserClientCheckEntityExistsResponse{}, err
	}
	return WorkspaceNotificationRecipientUserClientCheckEntityExistsResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *WorkspaceNotificationRecipientUserClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, _ *WorkspaceNotificationRecipientUserClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - Adds the API Management User to the list of Recipients for the Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceNotificationRecipientUserClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceNotificationRecipientUserClient.CreateOrUpdate
//     method.
func (client *WorkspaceNotificationRecipientUserClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, options *WorkspaceNotificationRecipientUserClientCreateOrUpdateOptions) (WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceNotificationRecipientUserClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, userID, options)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceNotificationRecipientUserClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, _ *WorkspaceNotificationRecipientUserClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceNotificationRecipientUserClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse, error) {
	result := WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserContract); err != nil {
		return WorkspaceNotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the API Management user from the list of Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceNotificationRecipientUserClientDeleteOptions contains the optional parameters for the WorkspaceNotificationRecipientUserClient.Delete
//     method.
func (client *WorkspaceNotificationRecipientUserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, options *WorkspaceNotificationRecipientUserClientDeleteOptions) (WorkspaceNotificationRecipientUserClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceNotificationRecipientUserClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, userID, options)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceNotificationRecipientUserClientDeleteResponse{}, err
	}
	return WorkspaceNotificationRecipientUserClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceNotificationRecipientUserClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, userID string, _ *WorkspaceNotificationRecipientUserClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ListByNotification - Gets the list of the Notification Recipient User subscribed to the notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - options - WorkspaceNotificationRecipientUserClientListByNotificationOptions contains the optional parameters for the WorkspaceNotificationRecipientUserClient.ListByNotification
//     method.
func (client *WorkspaceNotificationRecipientUserClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, options *WorkspaceNotificationRecipientUserClientListByNotificationOptions) (WorkspaceNotificationRecipientUserClientListByNotificationResponse, error) {
	var err error
	const operationName = "WorkspaceNotificationRecipientUserClient.ListByNotification"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, options)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientListByNotificationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientUserClientListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceNotificationRecipientUserClientListByNotificationResponse{}, err
	}
	resp, err := client.listByNotificationHandleResponse(httpResp)
	return resp, err
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *WorkspaceNotificationRecipientUserClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, _ *WorkspaceNotificationRecipientUserClientListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientUsers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if notificationName == "" {
		return nil, errors.New("parameter notificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationName}", url.PathEscape(string(notificationName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNotificationHandleResponse handles the ListByNotification response.
func (client *WorkspaceNotificationRecipientUserClient) listByNotificationHandleResponse(resp *http.Response) (WorkspaceNotificationRecipientUserClientListByNotificationResponse, error) {
	result := WorkspaceNotificationRecipientUserClientListByNotificationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserCollection); err != nil {
		return WorkspaceNotificationRecipientUserClientListByNotificationResponse{}, err
	}
	return result, nil
}
