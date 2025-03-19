// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// PreconfiguredEndpointsClient contains the methods for the PreconfiguredEndpoints group.
// Don't use this type directly, use NewPreconfiguredEndpointsClient() instead.
type PreconfiguredEndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPreconfiguredEndpointsClient creates a new instance of PreconfiguredEndpointsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPreconfiguredEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PreconfiguredEndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PreconfiguredEndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets a list of Preconfigured Endpoints
//
// Generated from API version 2019-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - The Profile identifier associated with the Tenant and Partner
//   - options - PreconfiguredEndpointsClientListOptions contains the optional parameters for the PreconfiguredEndpointsClient.NewListPager
//     method.
func (client *PreconfiguredEndpointsClient) NewListPager(resourceGroupName string, profileName string, options *PreconfiguredEndpointsClientListOptions) *runtime.Pager[PreconfiguredEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PreconfiguredEndpointsClientListResponse]{
		More: func(page PreconfiguredEndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PreconfiguredEndpointsClientListResponse) (PreconfiguredEndpointsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PreconfiguredEndpointsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, profileName, options)
			}, nil)
			if err != nil {
				return PreconfiguredEndpointsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PreconfiguredEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, profileName string, _ *PreconfiguredEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/PreconfiguredEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PreconfiguredEndpointsClient) listHandleResponse(resp *http.Response) (PreconfiguredEndpointsClientListResponse, error) {
	result := PreconfiguredEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PreconfiguredEndpointList); err != nil {
		return PreconfiguredEndpointsClientListResponse{}, err
	}
	return result, nil
}
