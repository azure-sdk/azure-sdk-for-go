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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationEventsClient contains the methods for the ReplicationEvents group.
// Don't use this type directly, use NewReplicationEventsClient() instead.
type ReplicationEventsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationEventsClient creates a new instance of ReplicationEventsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationEventsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationEventsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationEventsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - The operation to get the details of an Azure Site recovery event.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - eventName - The name of the Azure Site Recovery event.
//   - options - ReplicationEventsClientGetOptions contains the optional parameters for the ReplicationEventsClient.Get method.
func (client *ReplicationEventsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, eventName string, options *ReplicationEventsClientGetOptions) (ReplicationEventsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, eventName, options)
	if err != nil {
		return ReplicationEventsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationEventsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationEventsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationEventsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, eventName string, options *ReplicationEventsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventName == "" {
		return nil, errors.New("parameter eventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventName}", url.PathEscape(eventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationEventsClient) getHandleResponse(resp *http.Response) (ReplicationEventsClientGetResponse, error) {
	result := ReplicationEventsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Event); err != nil {
		return ReplicationEventsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of Azure Site Recovery events for the vault.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationEventsClientListOptions contains the optional parameters for the ReplicationEventsClient.NewListPager
//     method.
func (client *ReplicationEventsClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationEventsClientListOptions) *runtime.Pager[ReplicationEventsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationEventsClientListResponse]{
		More: func(page ReplicationEventsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationEventsClientListResponse) (ReplicationEventsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationEventsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationEventsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationEventsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationEventsClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationEventsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationEventsClient) listHandleResponse(resp *http.Response) (ReplicationEventsClientListResponse, error) {
	result := ReplicationEventsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventCollection); err != nil {
		return ReplicationEventsClientListResponse{}, err
	}
	return result, nil
}
