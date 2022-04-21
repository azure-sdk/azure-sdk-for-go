//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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
	"strconv"
	"strings"
)

// OperationResultsClient contains the methods for the OperationResults group.
// Don't use this type directly, use NewOperationResultsClient() instead.
type OperationResultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationResultsClient creates a new instance of OperationResultsClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationResultsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OperationResultsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get asset track operation result.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// trackName - The Asset Track name.
// operationID - Operation Id.
// options - OperationResultsClientGetOptions contains the optional parameters for the OperationResultsClient.Get method.
func (client *OperationResultsClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, operationID string, options *OperationResultsClientGetOptions) (OperationResultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, assetName, trackName, operationID, options)
	if err != nil {
		return OperationResultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return OperationResultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationResultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, trackName string, operationID string, options *OperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/tracks/{trackName}/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if trackName == "" {
		return nil, errors.New("parameter trackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trackName}", url.PathEscape(trackName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationResultsClient) getHandleResponse(resp *http.Response) (OperationResultsClientGetResponse, error) {
	result := OperationResultsClientGetResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return OperationResultsClientGetResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetTrack); err != nil {
		return OperationResultsClientGetResponse{}, err
	}
	return result, nil
}
