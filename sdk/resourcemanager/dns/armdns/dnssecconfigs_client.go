//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdns

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

// DnssecConfigsClient contains the methods for the DnssecConfigs group.
// Don't use this type directly, use NewDnssecConfigsClient() instead.
type DnssecConfigsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDnssecConfigsClient creates a new instance of DnssecConfigsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDnssecConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DnssecConfigsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DnssecConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the DNSSEC configuration on a DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - DnssecConfigsClientBeginCreateOrUpdateOptions contains the optional parameters for the DnssecConfigsClient.BeginCreateOrUpdate
//     method.
func (client *DnssecConfigsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DnssecConfigsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, zoneName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DnssecConfigsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DnssecConfigsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the DNSSEC configuration on a DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *DnssecConfigsClient) createOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DnssecConfigsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, zoneName, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DnssecConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/dnssecConfigs/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	return req, nil
}

// BeginDelete - Deletes the DNSSEC configuration on a DNS zone. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - DnssecConfigsClientBeginDeleteOptions contains the optional parameters for the DnssecConfigsClient.BeginDelete
//     method.
func (client *DnssecConfigsClient) BeginDelete(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginDeleteOptions) (*runtime.Poller[DnssecConfigsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, zoneName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DnssecConfigsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DnssecConfigsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the DNSSEC configuration on a DNS zone. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *DnssecConfigsClient) deleteOperation(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DnssecConfigsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, zoneName, options)
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
func (client *DnssecConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/dnssecConfigs/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// Get - Gets the DNSSEC configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - DnssecConfigsClientGetOptions contains the optional parameters for the DnssecConfigsClient.Get method.
func (client *DnssecConfigsClient) Get(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientGetOptions) (DnssecConfigsClientGetResponse, error) {
	var err error
	const operationName = "DnssecConfigsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, zoneName, options)
	if err != nil {
		return DnssecConfigsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DnssecConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DnssecConfigsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DnssecConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/dnssecConfigs/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DnssecConfigsClient) getHandleResponse(resp *http.Response) (DnssecConfigsClientGetResponse, error) {
	result := DnssecConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DnssecConfig); err != nil {
		return DnssecConfigsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDNSZonePager - Lists the DNSSEC configurations in a DNS zone.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - DnssecConfigsClientListByDNSZoneOptions contains the optional parameters for the DnssecConfigsClient.NewListByDNSZonePager
//     method.
func (client *DnssecConfigsClient) NewListByDNSZonePager(resourceGroupName string, zoneName string, options *DnssecConfigsClientListByDNSZoneOptions) *runtime.Pager[DnssecConfigsClientListByDNSZoneResponse] {
	return runtime.NewPager(runtime.PagingHandler[DnssecConfigsClientListByDNSZoneResponse]{
		More: func(page DnssecConfigsClientListByDNSZoneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DnssecConfigsClientListByDNSZoneResponse) (DnssecConfigsClientListByDNSZoneResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DnssecConfigsClient.NewListByDNSZonePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDNSZoneCreateRequest(ctx, resourceGroupName, zoneName, options)
			}, nil)
			if err != nil {
				return DnssecConfigsClientListByDNSZoneResponse{}, err
			}
			return client.listByDNSZoneHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDNSZoneCreateRequest creates the ListByDNSZone request.
func (client *DnssecConfigsClient) listByDNSZoneCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *DnssecConfigsClientListByDNSZoneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/dnssecConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDNSZoneHandleResponse handles the ListByDNSZone response.
func (client *DnssecConfigsClient) listByDNSZoneHandleResponse(resp *http.Response) (DnssecConfigsClientListByDNSZoneResponse, error) {
	result := DnssecConfigsClientListByDNSZoneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DnssecConfigListResult); err != nil {
		return DnssecConfigsClientListByDNSZoneResponse{}, err
	}
	return result, nil
}
