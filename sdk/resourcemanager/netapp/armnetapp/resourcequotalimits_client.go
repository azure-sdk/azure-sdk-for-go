//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// ResourceQuotaLimitsClient contains the methods for the NetAppResourceQuotaLimits group.
// Don't use this type directly, use NewResourceQuotaLimitsClient() instead.
type ResourceQuotaLimitsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceQuotaLimitsClient creates a new instance of ResourceQuotaLimitsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceQuotaLimitsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceQuotaLimitsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceQuotaLimitsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the default and current subscription quota limit
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - location - The name of the Azure region.
//   - quotaLimitName - The name of the Quota Limit
//   - options - ResourceQuotaLimitsClientGetOptions contains the optional parameters for the ResourceQuotaLimitsClient.Get method.
func (client *ResourceQuotaLimitsClient) Get(ctx context.Context, location string, quotaLimitName string, options *ResourceQuotaLimitsClientGetOptions) (ResourceQuotaLimitsClientGetResponse, error) {
	var err error
	const operationName = "ResourceQuotaLimitsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, quotaLimitName, options)
	if err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ResourceQuotaLimitsClient) getCreateRequest(ctx context.Context, location string, quotaLimitName string, options *ResourceQuotaLimitsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/quotaLimits/{quotaLimitName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if quotaLimitName == "" {
		return nil, errors.New("parameter quotaLimitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaLimitName}", url.PathEscape(quotaLimitName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceQuotaLimitsClient) getHandleResponse(resp *http.Response) (ResourceQuotaLimitsClientGetResponse, error) {
	result := ResourceQuotaLimitsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotaItem); err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the default and current limits for quotas
//
// Generated from API version 2024-07-01
//   - location - The name of the Azure region.
//   - options - ResourceQuotaLimitsClientListOptions contains the optional parameters for the ResourceQuotaLimitsClient.NewListPager
//     method.
func (client *ResourceQuotaLimitsClient) NewListPager(location string, options *ResourceQuotaLimitsClientListOptions) *runtime.Pager[ResourceQuotaLimitsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceQuotaLimitsClientListResponse]{
		More: func(page ResourceQuotaLimitsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ResourceQuotaLimitsClientListResponse) (ResourceQuotaLimitsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceQuotaLimitsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, location, options)
			if err != nil {
				return ResourceQuotaLimitsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ResourceQuotaLimitsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceQuotaLimitsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ResourceQuotaLimitsClient) listCreateRequest(ctx context.Context, location string, options *ResourceQuotaLimitsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/quotaLimits"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceQuotaLimitsClient) listHandleResponse(resp *http.Response) (ResourceQuotaLimitsClientListResponse, error) {
	result := ResourceQuotaLimitsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotaItemList); err != nil {
		return ResourceQuotaLimitsClientListResponse{}, err
	}
	return result, nil
}
