//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

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

// CloudHsmClusterPrivateEndpointConnectionsClient contains the methods for the CloudHsmClusterPrivateEndpointConnections group.
// Don't use this type directly, use NewCloudHsmClusterPrivateEndpointConnectionsClient() instead.
type CloudHsmClusterPrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCloudHsmClusterPrivateEndpointConnectionsClient creates a new instance of CloudHsmClusterPrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCloudHsmClusterPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudHsmClusterPrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".CloudHsmClusterPrivateEndpointConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CloudHsmClusterPrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or updates the private endpoint connection for the Cloud Hsm Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudHsmClusterName - The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must
//     be between 3 and 24 characters in length.
//   - peConnectionName - Name of the private endpoint connection associated with the Cloud HSM Cluster.
//   - properties - Parameters of the PrivateEndpointConnection
//   - options - CloudHsmClusterPrivateEndpointConnectionsClientCreateOptions contains the optional parameters for the CloudHsmClusterPrivateEndpointConnectionsClient.Create
//     method.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) Create(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, properties PrivateEndpointConnection, options *CloudHsmClusterPrivateEndpointConnectionsClientCreateOptions) (CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, cloudHsmClusterName, peConnectionName, properties, options)
	if err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, properties PrivateEndpointConnection, options *CloudHsmClusterPrivateEndpointConnectionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/{cloudHsmClusterName}/privateEndpointConnections/{peConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudHsmClusterName == "" {
		return nil, errors.New("parameter cloudHsmClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudHsmClusterName}", url.PathEscape(cloudHsmClusterName))
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) createHandleResponse(resp *http.Response) (CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse, error) {
	result := CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the private endpoint connection for the Cloud Hsm Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudHsmClusterName - The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must
//     be between 3 and 24 characters in length.
//   - peConnectionName - Name of the private endpoint connection associated with the Cloud HSM Cluster.
//   - options - CloudHsmClusterPrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the CloudHsmClusterPrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *CloudHsmClusterPrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cloudHsmClusterName, peConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[CloudHsmClusterPrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the private endpoint connection for the Cloud Hsm Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-31-preview
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *CloudHsmClusterPrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cloudHsmClusterName, peConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *CloudHsmClusterPrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/{cloudHsmClusterName}/privateEndpointConnections/{peConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudHsmClusterName == "" {
		return nil, errors.New("parameter cloudHsmClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudHsmClusterName}", url.PathEscape(cloudHsmClusterName))
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the private endpoint connection for the Cloud Hsm Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudHsmClusterName - The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must
//     be between 3 and 24 characters in length.
//   - peConnectionName - Name of the private endpoint connection associated with the Cloud HSM Cluster.
//   - options - CloudHsmClusterPrivateEndpointConnectionsClientGetOptions contains the optional parameters for the CloudHsmClusterPrivateEndpointConnectionsClient.Get
//     method.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *CloudHsmClusterPrivateEndpointConnectionsClientGetOptions) (CloudHsmClusterPrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, cloudHsmClusterName, peConnectionName, options)
	if err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CloudHsmClusterPrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, peConnectionName string, options *CloudHsmClusterPrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/{cloudHsmClusterName}/privateEndpointConnections/{peConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudHsmClusterName == "" {
		return nil, errors.New("parameter cloudHsmClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudHsmClusterName}", url.PathEscape(cloudHsmClusterName))
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudHsmClusterPrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (CloudHsmClusterPrivateEndpointConnectionsClientGetResponse, error) {
	result := CloudHsmClusterPrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return CloudHsmClusterPrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}
