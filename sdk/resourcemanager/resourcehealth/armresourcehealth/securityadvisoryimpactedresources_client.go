//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

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

// SecurityAdvisoryImpactedResourcesClient contains the methods for the SecurityAdvisoryImpactedResources group.
// Don't use this type directly, use NewSecurityAdvisoryImpactedResourcesClient() instead.
type SecurityAdvisoryImpactedResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityAdvisoryImpactedResourcesClient creates a new instance of SecurityAdvisoryImpactedResourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityAdvisoryImpactedResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityAdvisoryImpactedResourcesClient, error) {
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
	client := &SecurityAdvisoryImpactedResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListBySubscriptionIDAndEventIDPager - Lists impacted resources in the subscription by an event (Security Advisory).
// Generated from API version 2022-10-01-preview
// eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
// options - SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions contains the optional parameters
// for the SecurityAdvisoryImpactedResourcesClient.ListBySubscriptionIDAndEventID method.
func (client *SecurityAdvisoryImpactedResourcesClient) NewListBySubscriptionIDAndEventIDPager(eventTrackingID string, options *SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) *runtime.Pager[SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse]{
		More: func(page SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse) (SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionIDAndEventIDCreateRequest(ctx, eventTrackingID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionIDAndEventIDHandleResponse(resp)
		},
	})
}

// listBySubscriptionIDAndEventIDCreateRequest creates the ListBySubscriptionIDAndEventID request.
func (client *SecurityAdvisoryImpactedResourcesClient) listBySubscriptionIDAndEventIDCreateRequest(ctx context.Context, eventTrackingID string, options *SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/listSecurityAdvisoryImpactedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionIDAndEventIDHandleResponse handles the ListBySubscriptionIDAndEventID response.
func (client *SecurityAdvisoryImpactedResourcesClient) listBySubscriptionIDAndEventIDHandleResponse(resp *http.Response) (SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse, error) {
	result := SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResourceListResult); err != nil {
		return SecurityAdvisoryImpactedResourcesClientListBySubscriptionIDAndEventIDResponse{}, err
	}
	return result, nil
}

// NewListByTenantIDAndEventIDPager - Lists impacted resources in the tenant by an event (Security Advisory).
// Generated from API version 2022-10-01-preview
// eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
// options - SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDOptions contains the optional parameters for the
// SecurityAdvisoryImpactedResourcesClient.ListByTenantIDAndEventID method.
func (client *SecurityAdvisoryImpactedResourcesClient) NewListByTenantIDAndEventIDPager(eventTrackingID string, options *SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDOptions) *runtime.Pager[SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse]{
		More: func(page SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse) (SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTenantIDAndEventIDCreateRequest(ctx, eventTrackingID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTenantIDAndEventIDHandleResponse(resp)
		},
	})
}

// listByTenantIDAndEventIDCreateRequest creates the ListByTenantIDAndEventID request.
func (client *SecurityAdvisoryImpactedResourcesClient) listByTenantIDAndEventIDCreateRequest(ctx context.Context, eventTrackingID string, options *SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/listSecurityAdvisoryImpactedResources"
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTenantIDAndEventIDHandleResponse handles the ListByTenantIDAndEventID response.
func (client *SecurityAdvisoryImpactedResourcesClient) listByTenantIDAndEventIDHandleResponse(resp *http.Response) (SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse, error) {
	result := SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventImpactedResourceListResult); err != nil {
		return SecurityAdvisoryImpactedResourcesClientListByTenantIDAndEventIDResponse{}, err
	}
	return result, nil
}
