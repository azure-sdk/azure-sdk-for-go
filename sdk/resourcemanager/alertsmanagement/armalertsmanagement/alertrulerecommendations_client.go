// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

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

// AlertRuleRecommendationsClient contains the methods for the AlertRuleRecommendations group.
// Don't use this type directly, use NewAlertRuleRecommendationsClient() instead.
type AlertRuleRecommendationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAlertRuleRecommendationsClient creates a new instance of AlertRuleRecommendationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertRuleRecommendationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRuleRecommendationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertRuleRecommendationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByResourcePager - Retrieve alert rule recommendations for a resource.
//
// Generated from API version 2023-08-01-preview
//   - resourceURI - The identifier of the resource.
//   - options - AlertRuleRecommendationsClientListByResourceOptions contains the optional parameters for the AlertRuleRecommendationsClient.NewListByResourcePager
//     method.
func (client *AlertRuleRecommendationsClient) NewListByResourcePager(resourceURI string, options *AlertRuleRecommendationsClientListByResourceOptions) *runtime.Pager[AlertRuleRecommendationsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertRuleRecommendationsClientListByResourceResponse]{
		More: func(page AlertRuleRecommendationsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRuleRecommendationsClientListByResourceResponse) (AlertRuleRecommendationsClientListByResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AlertRuleRecommendationsClient.NewListByResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return AlertRuleRecommendationsClientListByResourceResponse{}, err
			}
			return client.listByResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *AlertRuleRecommendationsClient) listByResourceCreateRequest(ctx context.Context, resourceURI string, _ *AlertRuleRecommendationsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.AlertsManagement/alertRuleRecommendations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *AlertRuleRecommendationsClient) listByResourceHandleResponse(resp *http.Response) (AlertRuleRecommendationsClientListByResourceResponse, error) {
	result := AlertRuleRecommendationsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleRecommendationsListResponse); err != nil {
		return AlertRuleRecommendationsClientListByResourceResponse{}, err
	}
	return result, nil
}

// NewListByTargetTypePager - Retrieve alert rule recommendations for a target type.
//
// Generated from API version 2023-08-01-preview
//   - targetType - The recommendations target type.
//   - options - AlertRuleRecommendationsClientListByTargetTypeOptions contains the optional parameters for the AlertRuleRecommendationsClient.NewListByTargetTypePager
//     method.
func (client *AlertRuleRecommendationsClient) NewListByTargetTypePager(targetType string, options *AlertRuleRecommendationsClientListByTargetTypeOptions) *runtime.Pager[AlertRuleRecommendationsClientListByTargetTypeResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertRuleRecommendationsClientListByTargetTypeResponse]{
		More: func(page AlertRuleRecommendationsClientListByTargetTypeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRuleRecommendationsClientListByTargetTypeResponse) (AlertRuleRecommendationsClientListByTargetTypeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AlertRuleRecommendationsClient.NewListByTargetTypePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTargetTypeCreateRequest(ctx, targetType, options)
			}, nil)
			if err != nil {
				return AlertRuleRecommendationsClientListByTargetTypeResponse{}, err
			}
			return client.listByTargetTypeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTargetTypeCreateRequest creates the ListByTargetType request.
func (client *AlertRuleRecommendationsClient) listByTargetTypeCreateRequest(ctx context.Context, targetType string, _ *AlertRuleRecommendationsClientListByTargetTypeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alertRuleRecommendations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	reqQP.Set("targetType", targetType)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTargetTypeHandleResponse handles the ListByTargetType response.
func (client *AlertRuleRecommendationsClient) listByTargetTypeHandleResponse(resp *http.Response) (AlertRuleRecommendationsClientListByTargetTypeResponse, error) {
	result := AlertRuleRecommendationsClientListByTargetTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleRecommendationsListResponse); err != nil {
		return AlertRuleRecommendationsClientListByTargetTypeResponse{}, err
	}
	return result, nil
}
