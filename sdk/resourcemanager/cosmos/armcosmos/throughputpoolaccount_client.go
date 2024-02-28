//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// ThroughputPoolAccountClient contains the methods for the ThroughputPoolAccount group.
// Don't use this type directly, use NewThroughputPoolAccountClient() instead.
type ThroughputPoolAccountClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewThroughputPoolAccountClient creates a new instance of ThroughputPoolAccountClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewThroughputPoolAccountClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThroughputPoolAccountClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ThroughputPoolAccountClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or updates an Azure Cosmos DB ThroughputPool account. The "Update" method is preferred when performing
// updates on an account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - throughputPoolName - Cosmos DB Throughput Pool name.
//   - throughputPoolAccountName - Cosmos DB global database account in a Throughput Pool
//   - body - The parameters to provide for the current ThroughputPoolAccount.
//   - options - ThroughputPoolAccountClientBeginCreateOptions contains the optional parameters for the ThroughputPoolAccountClient.BeginCreate
//     method.
func (client *ThroughputPoolAccountClient) BeginCreate(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, body ThroughputPoolAccountResource, options *ThroughputPoolAccountClientBeginCreateOptions) (*runtime.Poller[ThroughputPoolAccountClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, throughputPoolName, throughputPoolAccountName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ThroughputPoolAccountClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ThroughputPoolAccountClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates or updates an Azure Cosmos DB ThroughputPool account. The "Update" method is preferred when performing
// updates on an account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-15-preview
func (client *ThroughputPoolAccountClient) create(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, body ThroughputPoolAccountResource, options *ThroughputPoolAccountClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ThroughputPoolAccountClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, throughputPoolName, throughputPoolAccountName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ThroughputPoolAccountClient) createCreateRequest(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, body ThroughputPoolAccountResource, options *ThroughputPoolAccountClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/throughputPools/{throughputPoolName}/throughputPoolAccounts/{throughputPoolAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if throughputPoolName == "" {
		return nil, errors.New("parameter throughputPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolName}", url.PathEscape(throughputPoolName))
	if throughputPoolAccountName == "" {
		return nil, errors.New("parameter throughputPoolAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolAccountName}", url.PathEscape(throughputPoolAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Removes an existing Azure Cosmos DB database account from a throughput pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - throughputPoolName - Cosmos DB Throughput Pool name.
//   - throughputPoolAccountName - Cosmos DB global database account in a Throughput Pool
//   - options - ThroughputPoolAccountClientBeginDeleteOptions contains the optional parameters for the ThroughputPoolAccountClient.BeginDelete
//     method.
func (client *ThroughputPoolAccountClient) BeginDelete(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, options *ThroughputPoolAccountClientBeginDeleteOptions) (*runtime.Poller[ThroughputPoolAccountClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, throughputPoolName, throughputPoolAccountName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ThroughputPoolAccountClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ThroughputPoolAccountClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Removes an existing Azure Cosmos DB database account from a throughput pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-15-preview
func (client *ThroughputPoolAccountClient) deleteOperation(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, options *ThroughputPoolAccountClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ThroughputPoolAccountClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, throughputPoolName, throughputPoolAccountName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ThroughputPoolAccountClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, options *ThroughputPoolAccountClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/throughputPools/{throughputPoolName}/throughputPoolAccounts/{throughputPoolAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if throughputPoolName == "" {
		return nil, errors.New("parameter throughputPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolName}", url.PathEscape(throughputPoolName))
	if throughputPoolAccountName == "" {
		return nil, errors.New("parameter throughputPoolAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolAccountName}", url.PathEscape(throughputPoolAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the properties of an existing Azure Cosmos DB Throughput Pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - throughputPoolName - Cosmos DB Throughput Pool name.
//   - throughputPoolAccountName - Cosmos DB global database account in a Throughput Pool
//   - options - ThroughputPoolAccountClientGetOptions contains the optional parameters for the ThroughputPoolAccountClient.Get
//     method.
func (client *ThroughputPoolAccountClient) Get(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, options *ThroughputPoolAccountClientGetOptions) (ThroughputPoolAccountClientGetResponse, error) {
	var err error
	const operationName = "ThroughputPoolAccountClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, throughputPoolName, throughputPoolAccountName, options)
	if err != nil {
		return ThroughputPoolAccountClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ThroughputPoolAccountClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ThroughputPoolAccountClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ThroughputPoolAccountClient) getCreateRequest(ctx context.Context, resourceGroupName string, throughputPoolName string, throughputPoolAccountName string, options *ThroughputPoolAccountClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/throughputPools/{throughputPoolName}/throughputPoolAccounts/{throughputPoolAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if throughputPoolName == "" {
		return nil, errors.New("parameter throughputPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolName}", url.PathEscape(throughputPoolName))
	if throughputPoolAccountName == "" {
		return nil, errors.New("parameter throughputPoolAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolAccountName}", url.PathEscape(throughputPoolAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ThroughputPoolAccountClient) getHandleResponse(resp *http.Response) (ThroughputPoolAccountClientGetResponse, error) {
	result := ThroughputPoolAccountClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputPoolAccountResource); err != nil {
		return ThroughputPoolAccountClientGetResponse{}, err
	}
	return result, nil
}
