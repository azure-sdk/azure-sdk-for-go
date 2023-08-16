//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

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

// SubscriptionRequestsClient contains the methods for the SubscriptionRequests group.
// Don't use this type directly, use NewSubscriptionRequestsClient() instead.
type SubscriptionRequestsClient struct {
	internal *arm.Client
}

// NewSubscriptionRequestsClient creates a new instance of SubscriptionRequestsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubscriptionRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionRequestsClient, error) {
	cl, err := arm.NewClient(moduleName+".SubscriptionRequestsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionRequestsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get API to check the status of a subscriptionIds request by requestId. Use the polling API - OperationsStatus URI
// specified in Azure-AsyncOperation header field, with retry-after duration in seconds
// to check the intermediate status. This API provides the finals status with the request details and status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - requestID - Request Id.
//   - options - SubscriptionRequestsClientGetOptions contains the optional parameters for the SubscriptionRequestsClient.Get
//     method.
func (client *SubscriptionRequestsClient) Get(ctx context.Context, mgID string, groupQuotaName string, requestID string, options *SubscriptionRequestsClientGetOptions) (SubscriptionRequestsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, mgID, groupQuotaName, requestID, options)
	if err != nil {
		return SubscriptionRequestsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubscriptionRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubscriptionRequestsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SubscriptionRequestsClient) getCreateRequest(ctx context.Context, mgID string, groupQuotaName string, requestID string, options *SubscriptionRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptionRequests/{requestId}"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	if requestID == "" {
		return nil, errors.New("parameter requestID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{requestId}", url.PathEscape(requestID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionRequestsClient) getHandleResponse(resp *http.Response) (SubscriptionRequestsClientGetResponse, error) {
	result := SubscriptionRequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotaSubscriptionID); err != nil {
		return SubscriptionRequestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List API to check the status of a subscriptionId requests by requestId. Request history is maintained for
// 1 year.
//
// Generated from API version 2023-06-01-preview
//   - mgID - Management Group Id.
//   - groupQuotaName - The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
//   - options - SubscriptionRequestsClientListOptions contains the optional parameters for the SubscriptionRequestsClient.NewListPager
//     method.
func (client *SubscriptionRequestsClient) NewListPager(mgID string, groupQuotaName string, options *SubscriptionRequestsClientListOptions) *runtime.Pager[SubscriptionRequestsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubscriptionRequestsClientListResponse]{
		More: func(page SubscriptionRequestsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionRequestsClientListResponse) (SubscriptionRequestsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, mgID, groupQuotaName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionRequestsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SubscriptionRequestsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionRequestsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SubscriptionRequestsClient) listCreateRequest(ctx context.Context, mgID string, groupQuotaName string, options *SubscriptionRequestsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{mgId}/providers/Microsoft.Quota/groupQuotas/{groupQuotaName}/subscriptionRequests"
	if mgID == "" {
		return nil, errors.New("parameter mgID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mgId}", url.PathEscape(mgID))
	if groupQuotaName == "" {
		return nil, errors.New("parameter groupQuotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupQuotaName}", url.PathEscape(groupQuotaName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionRequestsClient) listHandleResponse(resp *http.Response) (SubscriptionRequestsClientListResponse, error) {
	result := SubscriptionRequestsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupQuotaSubscriptionIDList); err != nil {
		return SubscriptionRequestsClientListResponse{}, err
	}
	return result, nil
}
