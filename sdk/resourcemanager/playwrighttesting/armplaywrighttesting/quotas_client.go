// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armplaywrighttesting

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

// QuotasClient contains the methods for the Quotas group.
// Don't use this type directly, use NewQuotasClient() instead.
type QuotasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQuotasClient creates a new instance of QuotasClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QuotasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get subscription quota by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - location - The location of quota in ARM Normalized format like eastus, southeastasia etc.
//   - quotaName - The quota name.
//   - options - QuotasClientGetOptions contains the optional parameters for the QuotasClient.Get method.
func (client *QuotasClient) Get(ctx context.Context, location string, quotaName QuotaNames, options *QuotasClientGetOptions) (QuotasClientGetResponse, error) {
	var err error
	const operationName = "QuotasClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, quotaName, options)
	if err != nil {
		return QuotasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QuotasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QuotasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *QuotasClient) getCreateRequest(ctx context.Context, location string, quotaName QuotaNames, _ *QuotasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzurePlaywrightService/locations/{location}/quotas/{quotaName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if quotaName == "" {
		return nil, errors.New("parameter quotaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaName}", url.PathEscape(string(quotaName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QuotasClient) getHandleResponse(resp *http.Response) (QuotasClientGetResponse, error) {
	result := QuotasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Quota); err != nil {
		return QuotasClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List quotas for a given subscription Id.
//
// Generated from API version 2024-02-01-preview
//   - location - The location of quota in ARM Normalized format like eastus, southeastasia etc.
//   - options - QuotasClientListBySubscriptionOptions contains the optional parameters for the QuotasClient.NewListBySubscriptionPager
//     method.
func (client *QuotasClient) NewListBySubscriptionPager(location string, options *QuotasClientListBySubscriptionOptions) *runtime.Pager[QuotasClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[QuotasClientListBySubscriptionResponse]{
		More: func(page QuotasClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotasClientListBySubscriptionResponse) (QuotasClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QuotasClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return QuotasClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *QuotasClient) listBySubscriptionCreateRequest(ctx context.Context, location string, _ *QuotasClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzurePlaywrightService/locations/{location}/quotas"
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
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *QuotasClient) listBySubscriptionHandleResponse(resp *http.Response) (QuotasClientListBySubscriptionResponse, error) {
	result := QuotasClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaListResult); err != nil {
		return QuotasClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
