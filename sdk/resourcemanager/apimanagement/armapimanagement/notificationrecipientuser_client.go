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

// NotificationRecipientUserClient contains the methods for the NotificationRecipientUser group.
// Don't use this type directly, use NewNotificationRecipientUserClient() instead.
type NotificationRecipientUserClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotificationRecipientUserClient creates a new instance of NotificationRecipientUserClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotificationRecipientUserClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotificationRecipientUserClient, error) {
	cl, err := arm.NewClient(moduleName+".NotificationRecipientUserClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotificationRecipientUserClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEntityExists - Determine if the Notification Recipient User is subscribed to the notification.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - NotificationRecipientUserClientCheckEntityExistsOptions contains the optional parameters for the NotificationRecipientUserClient.CheckEntityExists
//     method.
func (client *NotificationRecipientUserClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientCheckEntityExistsOptions) (NotificationRecipientUserClientCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserClientCheckEntityExistsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientUserClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent, http.StatusNotFound) {
		return NotificationRecipientUserClientCheckEntityExistsResponse{}, runtime.NewResponseError(resp)
	}
	return NotificationRecipientUserClientCheckEntityExistsResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *NotificationRecipientUserClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - Adds the API Management User to the list of Recipients for the Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - NotificationRecipientUserClientCreateOrUpdateOptions contains the optional parameters for the NotificationRecipientUserClient.CreateOrUpdate
//     method.
func (client *NotificationRecipientUserClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientCreateOrUpdateOptions) (NotificationRecipientUserClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return NotificationRecipientUserClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationRecipientUserClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationRecipientUserClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationRecipientUserClientCreateOrUpdateResponse, error) {
	result := NotificationRecipientUserClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserContract); err != nil {
		return NotificationRecipientUserClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the API Management user from the list of Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - userID - User identifier. Must be unique in the current API Management service instance.
//   - options - NotificationRecipientUserClientDeleteOptions contains the optional parameters for the NotificationRecipientUserClient.Delete
//     method.
func (client *NotificationRecipientUserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientDeleteOptions) (NotificationRecipientUserClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, notificationName, userID, options)
	if err != nil {
		return NotificationRecipientUserClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientUserClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return NotificationRecipientUserClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return NotificationRecipientUserClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationRecipientUserClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, userID string, options *NotificationRecipientUserClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers/{userId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ListByNotification - Gets the list of the Notification Recipient User subscribed to the notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - options - NotificationRecipientUserClientListByNotificationOptions contains the optional parameters for the NotificationRecipientUserClient.ListByNotification
//     method.
func (client *NotificationRecipientUserClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientUserClientListByNotificationOptions) (NotificationRecipientUserClientListByNotificationResponse, error) {
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationRecipientUserClientListByNotificationResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientUserClientListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationRecipientUserClientListByNotificationResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByNotificationHandleResponse(resp)
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *NotificationRecipientUserClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientUserClientListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
func (client *NotificationRecipientUserClient) listByNotificationHandleResponse(resp *http.Response) (NotificationRecipientUserClientListByNotificationResponse, error) {
	result := NotificationRecipientUserClientListByNotificationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientUserCollection); err != nil {
		return NotificationRecipientUserClientListByNotificationResponse{}, err
	}
	return result, nil
}
