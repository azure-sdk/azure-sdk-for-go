//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// ConnectorApplicationClient contains the methods for the SecurityConnectorApplication group.
// Don't use this type directly, use NewConnectorApplicationClient() instead.
type ConnectorApplicationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectorApplicationClient creates a new instance of ConnectorApplicationClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectorApplicationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectorApplicationClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectorApplicationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectorApplicationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or update a security Application on the given security connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - applicationID - The security Application key - unique key for the standard application
//   - application - Application over a subscription scope
//   - options - ConnectorApplicationClientCreateOrUpdateOptions contains the optional parameters for the ConnectorApplicationClient.CreateOrUpdate
//     method.
func (client *ConnectorApplicationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, application Application, options *ConnectorApplicationClientCreateOrUpdateOptions) (ConnectorApplicationClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityConnectorName, applicationID, application, options)
	if err != nil {
		return ConnectorApplicationClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorApplicationClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorApplicationClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectorApplicationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, application Application, options *ConnectorApplicationClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/applications/{applicationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if applicationID == "" {
		return nil, errors.New("parameter applicationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationId}", url.PathEscape(applicationID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, application); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectorApplicationClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectorApplicationClientCreateOrUpdateResponse, error) {
	result := ConnectorApplicationClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ConnectorApplicationClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an Application over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - applicationID - The security Application key - unique key for the standard application
//   - options - ConnectorApplicationClientDeleteOptions contains the optional parameters for the ConnectorApplicationClient.Delete
//     method.
func (client *ConnectorApplicationClient) Delete(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, options *ConnectorApplicationClientDeleteOptions) (ConnectorApplicationClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, securityConnectorName, applicationID, options)
	if err != nil {
		return ConnectorApplicationClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorApplicationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorApplicationClientDeleteResponse{}, err
	}
	return ConnectorApplicationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectorApplicationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, options *ConnectorApplicationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/applications/{applicationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if applicationID == "" {
		return nil, errors.New("parameter applicationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationId}", url.PathEscape(applicationID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific application for the requested scope by applicationId
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - applicationID - The security Application key - unique key for the standard application
//   - options - ConnectorApplicationClientGetOptions contains the optional parameters for the ConnectorApplicationClient.Get
//     method.
func (client *ConnectorApplicationClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, options *ConnectorApplicationClientGetOptions) (ConnectorApplicationClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, applicationID, options)
	if err != nil {
		return ConnectorApplicationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorApplicationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorApplicationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectorApplicationClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, applicationID string, options *ConnectorApplicationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/applications/{applicationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if applicationID == "" {
		return nil, errors.New("parameter applicationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationId}", url.PathEscape(applicationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectorApplicationClient) getHandleResponse(resp *http.Response) (ConnectorApplicationClientGetResponse, error) {
	result := ConnectorApplicationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Application); err != nil {
		return ConnectorApplicationClientGetResponse{}, err
	}
	return result, nil
}
