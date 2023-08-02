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
	"strconv"
	"strings"
)

// NotificationClient contains the methods for the Notification group.
// Don't use this type directly, use NewNotificationClient() instead.
type NotificationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotificationClient creates a new instance of NotificationClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotificationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotificationClient, error) {
	cl, err := arm.NewClient(moduleName+".NotificationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotificationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update API Management publisher notification.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - options - NotificationClientCreateOrUpdateOptions contains the optional parameters for the NotificationClient.CreateOrUpdate
//     method.
func (client *NotificationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientCreateOrUpdateOptions) (NotificationClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationClientCreateOrUpdateResponse, error) {
	result := NotificationClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationContract); err != nil {
		return NotificationClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the Notification specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - notificationName - Notification Name Identifier.
//   - options - NotificationClientGetOptions contains the optional parameters for the NotificationClient.Get method.
func (client *NotificationClient) Get(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientGetOptions) (NotificationClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, notificationName, options)
	if err != nil {
		return NotificationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotificationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NotificationClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, notificationName NotificationName, options *NotificationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}"
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

// getHandleResponse handles the Get response.
func (client *NotificationClient) getHandleResponse(resp *http.Response) (NotificationClientGetResponse, error) {
	result := NotificationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationContract); err != nil {
		return NotificationClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists a collection of properties defined within a service instance.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - NotificationClientListByServiceOptions contains the optional parameters for the NotificationClient.NewListByServicePager
//     method.
func (client *NotificationClient) NewListByServicePager(resourceGroupName string, serviceName string, options *NotificationClientListByServiceOptions) *runtime.Pager[NotificationClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotificationClientListByServiceResponse]{
		More: func(page NotificationClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NotificationClientListByServiceResponse) (NotificationClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NotificationClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NotificationClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NotificationClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NotificationClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NotificationClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NotificationClient) listByServiceHandleResponse(resp *http.Response) (NotificationClientListByServiceResponse, error) {
	result := NotificationClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationCollection); err != nil {
		return NotificationClientListByServiceResponse{}, err
	}
	return result, nil
}
