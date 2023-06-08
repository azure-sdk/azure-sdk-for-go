//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// WorkspaceNotificationRecipientEmailClient contains the methods for the WorkspaceNotificationRecipientEmail group.
// Don't use this type directly, use NewWorkspaceNotificationRecipientEmailClient() instead.
type WorkspaceNotificationRecipientEmailClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceNotificationRecipientEmailClient creates a new instance of WorkspaceNotificationRecipientEmailClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceNotificationRecipientEmailClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceNotificationRecipientEmailClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceNotificationRecipientEmailClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceNotificationRecipientEmailClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEntityExists - Determine if Notification Recipient Email subscribed to the notification.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - WorkspaceNotificationRecipientEmailClientCheckEntityExistsOptions contains the optional parameters for the WorkspaceNotificationRecipientEmailClient.CheckEntityExists
//     method.
func (client *WorkspaceNotificationRecipientEmailClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientCheckEntityExistsOptions) (WorkspaceNotificationRecipientEmailClientCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, email, options)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientCheckEntityExistsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent, http.StatusNotFound) {
		return WorkspaceNotificationRecipientEmailClientCheckEntityExistsResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceNotificationRecipientEmailClientCheckEntityExistsResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *WorkspaceNotificationRecipientEmailClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - Adds the Email address to the list of Recipients for the Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - WorkspaceNotificationRecipientEmailClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceNotificationRecipientEmailClient.CreateOrUpdate
//     method.
func (client *WorkspaceNotificationRecipientEmailClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientCreateOrUpdateOptions) (WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, email, options)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceNotificationRecipientEmailClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceNotificationRecipientEmailClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse, error) {
	result := WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailContract); err != nil {
		return WorkspaceNotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the email from the list of Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - WorkspaceNotificationRecipientEmailClientDeleteOptions contains the optional parameters for the WorkspaceNotificationRecipientEmailClient.Delete
//     method.
func (client *WorkspaceNotificationRecipientEmailClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientDeleteOptions) (WorkspaceNotificationRecipientEmailClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, email, options)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WorkspaceNotificationRecipientEmailClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceNotificationRecipientEmailClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceNotificationRecipientEmailClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, email string, options *WorkspaceNotificationRecipientEmailClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ListByNotification - Gets the list of the Notification Recipient Emails subscribed to a notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - notificationName - Notification Name Identifier.
//   - options - WorkspaceNotificationRecipientEmailClientListByNotificationOptions contains the optional parameters for the WorkspaceNotificationRecipientEmailClient.ListByNotification
//     method.
func (client *WorkspaceNotificationRecipientEmailClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, options *WorkspaceNotificationRecipientEmailClientListByNotificationOptions) (WorkspaceNotificationRecipientEmailClientListByNotificationResponse, error) {
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, notificationName, options)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceNotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceNotificationRecipientEmailClientListByNotificationResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByNotificationHandleResponse(resp)
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *WorkspaceNotificationRecipientEmailClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, notificationName NotificationName, options *WorkspaceNotificationRecipientEmailClientListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/notifications/{notificationName}/recipientEmails"
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNotificationHandleResponse handles the ListByNotification response.
func (client *WorkspaceNotificationRecipientEmailClient) listByNotificationHandleResponse(resp *http.Response) (WorkspaceNotificationRecipientEmailClientListByNotificationResponse, error) {
	result := WorkspaceNotificationRecipientEmailClientListByNotificationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailCollection); err != nil {
		return WorkspaceNotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	return result, nil
}
