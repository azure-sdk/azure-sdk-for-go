//go:build go1.18
// +build go1.18

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

// Client contains the methods for the MonitorClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal *arm.Client
}

// NewClient creates a new instance of Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		internal: cl,
	}
	return client, nil
}

// BeginCreateNotificationsAtTenantActionGroupResourceLevel - Send test notifications to a set of provided receivers
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - notificationRequest - The notification request body which includes the contact details
//   - options - ClientBeginCreateNotificationsAtTenantActionGroupResourceLevelOptions contains the optional parameters for the
//     Client.BeginCreateNotificationsAtTenantActionGroupResourceLevel method.
func (client *Client) BeginCreateNotificationsAtTenantActionGroupResourceLevel(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, notificationRequest TenantNotificationRequestBody, options *ClientBeginCreateNotificationsAtTenantActionGroupResourceLevelOptions) (*runtime.Poller[ClientCreateNotificationsAtTenantActionGroupResourceLevelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createNotificationsAtTenantActionGroupResourceLevel(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, notificationRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientCreateNotificationsAtTenantActionGroupResourceLevelResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClientCreateNotificationsAtTenantActionGroupResourceLevelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateNotificationsAtTenantActionGroupResourceLevel - Send test notifications to a set of provided receivers
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *Client) createNotificationsAtTenantActionGroupResourceLevel(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, notificationRequest TenantNotificationRequestBody, options *ClientBeginCreateNotificationsAtTenantActionGroupResourceLevelOptions) (*http.Response, error) {
	var err error
	const operationName = "Client.BeginCreateNotificationsAtTenantActionGroupResourceLevel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createNotificationsAtTenantActionGroupResourceLevelCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, notificationRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createNotificationsAtTenantActionGroupResourceLevelCreateRequest creates the CreateNotificationsAtTenantActionGroupResourceLevel request.
func (client *Client) createNotificationsAtTenantActionGroupResourceLevelCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, notificationRequest TenantNotificationRequestBody, options *ClientBeginCreateNotificationsAtTenantActionGroupResourceLevelOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}/createNotifications"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	if err := runtime.MarshalAsJSON(req, notificationRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// GetTestNotificationsAtTenantActionGroupResourceLevel - Get the test notifications by the notification id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - managementGroupID - The management group id.
//   - tenantActionGroupName - The name of the action group.
//   - xmsClientTenantID - The tenant ID of the client making the request.
//   - notificationID - The notification id
//   - options - ClientGetTestNotificationsAtTenantActionGroupResourceLevelOptions contains the optional parameters for the Client.GetTestNotificationsAtTenantActionGroupResourceLevel
//     method.
func (client *Client) GetTestNotificationsAtTenantActionGroupResourceLevel(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, notificationID string, options *ClientGetTestNotificationsAtTenantActionGroupResourceLevelOptions) (ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse, error) {
	var err error
	const operationName = "Client.GetTestNotificationsAtTenantActionGroupResourceLevel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getTestNotificationsAtTenantActionGroupResourceLevelCreateRequest(ctx, managementGroupID, tenantActionGroupName, xmsClientTenantID, notificationID, options)
	if err != nil {
		return ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{}, err
	}
	resp, err := client.getTestNotificationsAtTenantActionGroupResourceLevelHandleResponse(httpResp)
	return resp, err
}

// getTestNotificationsAtTenantActionGroupResourceLevelCreateRequest creates the GetTestNotificationsAtTenantActionGroupResourceLevel request.
func (client *Client) getTestNotificationsAtTenantActionGroupResourceLevelCreateRequest(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, notificationID string, options *ClientGetTestNotificationsAtTenantActionGroupResourceLevelOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Insights/tenantActionGroups/{tenantActionGroupName}/notificationStatus/{notificationId}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if tenantActionGroupName == "" {
		return nil, errors.New("parameter tenantActionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tenantActionGroupName}", url.PathEscape(tenantActionGroupName))
	if notificationID == "" {
		return nil, errors.New("parameter notificationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationId}", url.PathEscape(notificationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-client-tenant-id"] = []string{xmsClientTenantID}
	return req, nil
}

// getTestNotificationsAtTenantActionGroupResourceLevelHandleResponse handles the GetTestNotificationsAtTenantActionGroupResourceLevel response.
func (client *Client) getTestNotificationsAtTenantActionGroupResourceLevelHandleResponse(resp *http.Response) (ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse, error) {
	result := ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestNotificationDetailsResponseAutoGenerated); err != nil {
		return ClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{}, err
	}
	return result, nil
}