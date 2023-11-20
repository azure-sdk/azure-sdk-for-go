//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnewrelicobservability

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

// OrganizationsClient contains the methods for the Organizations group.
// Don't use this type directly, use NewOrganizationsClient() instead.
type OrganizationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOrganizationsClient creates a new instance of OrganizationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrganizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrganizationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List all the existing organizations
//
// Generated from API version 2023-11-01-preview
//   - userEmail - User Email.
//   - location - Location for NewRelic.
//   - options - OrganizationsClientListOptions contains the optional parameters for the OrganizationsClient.NewListPager method.
func (client *OrganizationsClient) NewListPager(userEmail string, location string, options *OrganizationsClientListOptions) *runtime.Pager[OrganizationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationsClientListResponse]{
		More: func(page OrganizationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationsClientListResponse) (OrganizationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrganizationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, userEmail, location, options)
			}, nil)
			if err != nil {
				return OrganizationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OrganizationsClient) listCreateRequest(ctx context.Context, userEmail string, location string, options *OrganizationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/NewRelic.Observability/organizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	reqQP.Set("userEmail", userEmail)
	reqQP.Set("location", location)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OrganizationsClient) listHandleResponse(resp *http.Response) (OrganizationsClientListResponse, error) {
	result := OrganizationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationsListResponse); err != nil {
		return OrganizationsClientListResponse{}, err
	}
	return result, nil
}
