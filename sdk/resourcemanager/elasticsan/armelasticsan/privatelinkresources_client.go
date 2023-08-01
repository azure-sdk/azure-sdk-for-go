//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateLinkResourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ListByElasticSan - Gets the private link resources that need to be created for a elastic San.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - elasticSanName - The name of the ElasticSan.
//   - options - PrivateLinkResourcesClientListByElasticSanOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByElasticSan
//     method.
func (client *PrivateLinkResourcesClient) ListByElasticSan(ctx context.Context, resourceGroupName string, elasticSanName string, options *PrivateLinkResourcesClientListByElasticSanOptions) (PrivateLinkResourcesClientListByElasticSanResponse, error) {
	var err error
	req, err := client.listByElasticSanCreateRequest(ctx, resourceGroupName, elasticSanName, options)
	if err != nil {
		return PrivateLinkResourcesClientListByElasticSanResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesClientListByElasticSanResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourcesClientListByElasticSanResponse{}, err
	}
	resp, err := client.listByElasticSanHandleResponse(httpResp)
	return resp, err
}

// listByElasticSanCreateRequest creates the ListByElasticSan request.
func (client *PrivateLinkResourcesClient) listByElasticSanCreateRequest(ctx context.Context, resourceGroupName string, elasticSanName string, options *PrivateLinkResourcesClientListByElasticSanOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if elasticSanName == "" {
		return nil, errors.New("parameter elasticSanName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticSanName}", url.PathEscape(elasticSanName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByElasticSanHandleResponse handles the ListByElasticSan response.
func (client *PrivateLinkResourcesClient) listByElasticSanHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByElasticSanResponse, error) {
	result := PrivateLinkResourcesClientListByElasticSanResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByElasticSanResponse{}, err
	}
	return result, nil
}
