//go:build go1.18
// +build go1.18

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

// NotificationRecipientEmailClient contains the methods for the NotificationRecipientEmail group.
// Don't use this type directly, use NewNotificationRecipientEmailClient() instead.
type NotificationRecipientEmailClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotificationRecipientEmailClient creates a new instance of NotificationRecipientEmailClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotificationRecipientEmailClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotificationRecipientEmailClient, error) {
	cl, err := arm.NewClient(moduleName+".NotificationRecipientEmailClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotificationRecipientEmailClient{
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
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - NotificationRecipientEmailClientCheckEntityExistsOptions contains the optional parameters for the NotificationRecipientEmailClient.CheckEntityExists
//     method.
func (client *NotificationRecipientEmailClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientCheckEntityExistsOptions) (NotificationRecipientEmailClientCheckEntityExistsResponse, error) {
	var err error
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailClientCheckEntityExistsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientEmailClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRecipientEmailClientCheckEntityExistsResponse{}, err
	}
	return NotificationRecipientEmailClientCheckEntityExistsResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *NotificationRecipientEmailClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
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
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - NotificationRecipientEmailClientCreateOrUpdateOptions contains the optional parameters for the NotificationRecipientEmailClient.CreateOrUpdate
//     method.
func (client *NotificationRecipientEmailClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientCreateOrUpdateOptions) (NotificationRecipientEmailClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationRecipientEmailClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
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
func (client *NotificationRecipientEmailClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationRecipientEmailClientCreateOrUpdateResponse, error) {
	result := NotificationRecipientEmailClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailContract); err != nil {
		return NotificationRecipientEmailClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the email from the list of Notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - email - Email identifier.
//   - options - NotificationRecipientEmailClientDeleteOptions contains the optional parameters for the NotificationRecipientEmailClient.Delete
//     method.
func (client *NotificationRecipientEmailClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientDeleteOptions) (NotificationRecipientEmailClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, notificationName, email, options)
	if err != nil {
		return NotificationRecipientEmailClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientEmailClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRecipientEmailClientDeleteResponse{}, err
	}
	return NotificationRecipientEmailClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationRecipientEmailClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, email string, options *NotificationRecipientEmailClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails/{email}"
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
	if email == "" {
		return nil, errors.New("parameter email cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{email}", url.PathEscape(email))
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
//   - notificationName - Notification Name Identifier.
//   - options - NotificationRecipientEmailClientListByNotificationOptions contains the optional parameters for the NotificationRecipientEmailClient.ListByNotification
//     method.
func (client *NotificationRecipientEmailClient) ListByNotification(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientEmailClientListByNotificationOptions) (NotificationRecipientEmailClientListByNotificationResponse, error) {
	var err error
	req, err := client.listByNotificationCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	resp, err := client.listByNotificationHandleResponse(httpResp)
	return resp, err
}

// listByNotificationCreateRequest creates the ListByNotification request.
func (client *NotificationRecipientEmailClient) listByNotificationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationRecipientEmailClientListByNotificationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails"
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
func (client *NotificationRecipientEmailClient) listByNotificationHandleResponse(resp *http.Response) (NotificationRecipientEmailClientListByNotificationResponse, error) {
	result := NotificationRecipientEmailClientListByNotificationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipientEmailCollection); err != nil {
		return NotificationRecipientEmailClientListByNotificationResponse{}, err
	}
	return result, nil
}
