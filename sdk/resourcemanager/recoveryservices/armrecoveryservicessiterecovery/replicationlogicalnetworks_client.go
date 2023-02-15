//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery

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

// ReplicationLogicalNetworksClient contains the methods for the ReplicationLogicalNetworks group.
// Don't use this type directly, use NewReplicationLogicalNetworksClient() instead.
type ReplicationLogicalNetworksClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationLogicalNetworksClient creates a new instance of ReplicationLogicalNetworksClient with the specified values.
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationLogicalNetworksClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationLogicalNetworksClient, error) {
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
	client := &ReplicationLogicalNetworksClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Gets the details of a logical network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - fabricName - Server Id.
//   - logicalNetworkName - Logical network name.
//   - options - ReplicationLogicalNetworksClientGetOptions contains the optional parameters for the ReplicationLogicalNetworksClient.Get
//     method.
func (client *ReplicationLogicalNetworksClient) Get(ctx context.Context, fabricName string, logicalNetworkName string, options *ReplicationLogicalNetworksClientGetOptions) (ReplicationLogicalNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, logicalNetworkName, options)
	if err != nil {
		return ReplicationLogicalNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationLogicalNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationLogicalNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationLogicalNetworksClient) getCreateRequest(ctx context.Context, fabricName string, logicalNetworkName string, options *ReplicationLogicalNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks/{logicalNetworkName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if logicalNetworkName == "" {
		return nil, errors.New("parameter logicalNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicalNetworkName}", url.PathEscape(logicalNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationLogicalNetworksClient) getHandleResponse(resp *http.Response) (ReplicationLogicalNetworksClientGetResponse, error) {
	result := ReplicationLogicalNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalNetwork); err != nil {
		return ReplicationLogicalNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists all the logical networks of the Azure Site Recovery fabric.
//
// Generated from API version 2023-01-01
//   - fabricName - Server Id.
//   - options - ReplicationLogicalNetworksClientListByReplicationFabricsOptions contains the optional parameters for the ReplicationLogicalNetworksClient.NewListByReplicationFabricsPager
//     method.
func (client *ReplicationLogicalNetworksClient) NewListByReplicationFabricsPager(fabricName string, options *ReplicationLogicalNetworksClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationLogicalNetworksClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationLogicalNetworksClientListByReplicationFabricsResponse]{
		More: func(page ReplicationLogicalNetworksClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationLogicalNetworksClientListByReplicationFabricsResponse) (ReplicationLogicalNetworksClientListByReplicationFabricsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationFabricsCreateRequest(ctx, fabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationLogicalNetworksClientListByReplicationFabricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationLogicalNetworksClientListByReplicationFabricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationLogicalNetworksClientListByReplicationFabricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationLogicalNetworksClient) listByReplicationFabricsCreateRequest(ctx context.Context, fabricName string, options *ReplicationLogicalNetworksClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationLogicalNetworksClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationLogicalNetworksClientListByReplicationFabricsResponse, error) {
	result := ReplicationLogicalNetworksClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalNetworkCollection); err != nil {
		return ReplicationLogicalNetworksClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}
