//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNSiteLinkConnectionsClient contains the methods for the VPNSiteLinkConnections group.
// Don't use this type directly, use NewVPNSiteLinkConnectionsClient() instead.
type VPNSiteLinkConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNSiteLinkConnectionsClient creates a new instance of VPNSiteLinkConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNSiteLinkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNSiteLinkConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VPNSiteLinkConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieves the details of a vpn site link connection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The resource group name of the VpnGateway.
// gatewayName - The name of the gateway.
// connectionName - The name of the vpn connection.
// linkConnectionName - The name of the vpn connection.
// options - VPNSiteLinkConnectionsClientGetOptions contains the optional parameters for the VPNSiteLinkConnectionsClient.Get
// method.
func (client *VPNSiteLinkConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsClientGetOptions) (VPNSiteLinkConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
	if err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNSiteLinkConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNSiteLinkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNSiteLinkConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNSiteLinkConnectionsClient) getHandleResponse(resp *http.Response) (VPNSiteLinkConnectionsClientGetResponse, error) {
	result := VPNSiteLinkConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNSiteLinkConnection); err != nil {
		return VPNSiteLinkConnectionsClientGetResponse{}, err
	}
	return result, nil
}
