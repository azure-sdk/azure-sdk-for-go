//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ReplicationEligibilityResultsClient contains the methods for the ReplicationEligibilityResults group.
// Don't use this type directly, use NewReplicationEligibilityResultsClient() instead.
type ReplicationEligibilityResultsClient struct {
	host              string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationEligibilityResultsClient creates a new instance of ReplicationEligibilityResultsClient with the specified values.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationEligibilityResultsClient(resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationEligibilityResultsClient, error) {
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
	client := &ReplicationEligibilityResultsClient{
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Validates whether a given VM can be protected or not in which case returns list of errors.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// virtualMachineName - Virtual Machine name.
// options - ReplicationEligibilityResultsClientGetOptions contains the optional parameters for the ReplicationEligibilityResultsClient.Get
// method.
func (client *ReplicationEligibilityResultsClient) Get(ctx context.Context, virtualMachineName string, options *ReplicationEligibilityResultsClientGetOptions) (ReplicationEligibilityResultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, virtualMachineName, options)
	if err != nil {
		return ReplicationEligibilityResultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationEligibilityResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationEligibilityResultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationEligibilityResultsClient) getCreateRequest(ctx context.Context, virtualMachineName string, options *ReplicationEligibilityResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default"
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationEligibilityResultsClient) getHandleResponse(resp *http.Response) (ReplicationEligibilityResultsClientGetResponse, error) {
	result := ReplicationEligibilityResultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationEligibilityResults); err != nil {
		return ReplicationEligibilityResultsClientGetResponse{}, err
	}
	return result, nil
}

// List - Validates whether a given VM can be protected or not in which case returns list of errors.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// virtualMachineName - Virtual Machine name.
// options - ReplicationEligibilityResultsClientListOptions contains the optional parameters for the ReplicationEligibilityResultsClient.List
// method.
func (client *ReplicationEligibilityResultsClient) List(ctx context.Context, virtualMachineName string, options *ReplicationEligibilityResultsClientListOptions) (ReplicationEligibilityResultsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, virtualMachineName, options)
	if err != nil {
		return ReplicationEligibilityResultsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationEligibilityResultsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationEligibilityResultsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ReplicationEligibilityResultsClient) listCreateRequest(ctx context.Context, virtualMachineName string, options *ReplicationEligibilityResultsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults"
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationEligibilityResultsClient) listHandleResponse(resp *http.Response) (ReplicationEligibilityResultsClientListResponse, error) {
	result := ReplicationEligibilityResultsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationEligibilityResultsCollection); err != nil {
		return ReplicationEligibilityResultsClientListResponse{}, err
	}
	return result, nil
}
