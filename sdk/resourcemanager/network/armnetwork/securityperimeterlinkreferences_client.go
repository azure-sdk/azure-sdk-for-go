//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SecurityPerimeterLinkReferencesClient contains the methods for the NetworkSecurityPerimeterLinkReferences group.
// Don't use this type directly, use NewSecurityPerimeterLinkReferencesClient() instead.
type SecurityPerimeterLinkReferencesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityPerimeterLinkReferencesClient creates a new instance of SecurityPerimeterLinkReferencesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityPerimeterLinkReferencesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityPerimeterLinkReferencesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityPerimeterLinkReferencesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists the NSP LinkReference resources in the specified network security perimeter.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - networkSecurityPerimeterName - The name of the network security perimeter.
//   - options - SecurityPerimeterLinkReferencesClientListOptions contains the optional parameters for the SecurityPerimeterLinkReferencesClient.NewListPager
//     method.
func (client *SecurityPerimeterLinkReferencesClient) NewListPager(resourceGroupName string, networkSecurityPerimeterName string, options *SecurityPerimeterLinkReferencesClientListOptions) *runtime.Pager[SecurityPerimeterLinkReferencesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityPerimeterLinkReferencesClientListResponse]{
		More: func(page SecurityPerimeterLinkReferencesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityPerimeterLinkReferencesClientListResponse) (SecurityPerimeterLinkReferencesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecurityPerimeterLinkReferencesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, networkSecurityPerimeterName, options)
			}, nil)
			if err != nil {
				return SecurityPerimeterLinkReferencesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecurityPerimeterLinkReferencesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, options *SecurityPerimeterLinkReferencesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/linkReferences"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityPerimeterName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterName}", url.PathEscape(networkSecurityPerimeterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityPerimeterLinkReferencesClient) listHandleResponse(resp *http.Response) (SecurityPerimeterLinkReferencesClientListResponse, error) {
	result := SecurityPerimeterLinkReferencesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NspLinkReferenceListResult); err != nil {
		return SecurityPerimeterLinkReferencesClientListResponse{}, err
	}
	return result, nil
}
