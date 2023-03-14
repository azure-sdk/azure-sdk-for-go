//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers

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

// VirtualNetworkSubnetUsageClient contains the methods for the VirtualNetworkSubnetUsage group.
// Don't use this type directly, use NewVirtualNetworkSubnetUsageClient() instead.
type VirtualNetworkSubnetUsageClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkSubnetUsageClient creates a new instance of VirtualNetworkSubnetUsageClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkSubnetUsageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkSubnetUsageClient, error) {
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
	client := &VirtualNetworkSubnetUsageClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Execute - Get virtual network subnet usage for a given vNet resource id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - locationName - The name of the location.
//   - parameters - The required parameters for creating or updating a server.
//   - options - VirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the VirtualNetworkSubnetUsageClient.Execute
//     method.
func (client *VirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageClientExecuteOptions) (VirtualNetworkSubnetUsageClientExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *VirtualNetworkSubnetUsageClient) executeCreateRequest(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *VirtualNetworkSubnetUsageClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/checkVirtualNetworkSubnetUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// executeHandleResponse handles the Execute response.
func (client *VirtualNetworkSubnetUsageClient) executeHandleResponse(resp *http.Response) (VirtualNetworkSubnetUsageClientExecuteResponse, error) {
	result := VirtualNetworkSubnetUsageClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkSubnetUsageResult); err != nil {
		return VirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	return result, nil
}
