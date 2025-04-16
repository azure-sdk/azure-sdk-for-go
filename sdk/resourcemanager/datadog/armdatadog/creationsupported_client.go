// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

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

// CreationSupportedClient contains the methods for the CreationSupported group.
// Don't use this type directly, use NewCreationSupportedClient() instead.
type CreationSupportedClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCreationSupportedClient creates a new instance of CreationSupportedClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCreationSupportedClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CreationSupportedClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CreationSupportedClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Informs if the current subscription is being already monitored for selected Datadog organization.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
//   - datadogOrganizationID - Datadog Organization Id
//   - options - CreationSupportedClientGetOptions contains the optional parameters for the CreationSupportedClient.Get method.
func (client *CreationSupportedClient) Get(ctx context.Context, datadogOrganizationID string, options *CreationSupportedClientGetOptions) (CreationSupportedClientGetResponse, error) {
	var err error
	const operationName = "CreationSupportedClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, datadogOrganizationID, options)
	if err != nil {
		return CreationSupportedClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CreationSupportedClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CreationSupportedClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CreationSupportedClient) getCreateRequest(ctx context.Context, datadogOrganizationID string, _ *CreationSupportedClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/subscriptionStatuses/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	reqQP.Set("datadogOrganizationId", datadogOrganizationID)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CreationSupportedClient) getHandleResponse(resp *http.Response) (CreationSupportedClientGetResponse, error) {
	result := CreationSupportedClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateResourceSupportedResponse); err != nil {
		return CreationSupportedClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Informs if the current subscription is being already monitored for selected Datadog organization.
//
// Generated from API version 2023-10-20
//   - datadogOrganizationID - Datadog Organization Id
//   - options - CreationSupportedClientListOptions contains the optional parameters for the CreationSupportedClient.NewListPager
//     method.
func (client *CreationSupportedClient) NewListPager(datadogOrganizationID string, options *CreationSupportedClientListOptions) *runtime.Pager[CreationSupportedClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CreationSupportedClientListResponse]{
		More: func(page CreationSupportedClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CreationSupportedClientListResponse) (CreationSupportedClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CreationSupportedClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, datadogOrganizationID, options)
			}, nil)
			if err != nil {
				return CreationSupportedClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CreationSupportedClient) listCreateRequest(ctx context.Context, datadogOrganizationID string, _ *CreationSupportedClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/subscriptionStatuses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	reqQP.Set("datadogOrganizationId", datadogOrganizationID)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CreationSupportedClient) listHandleResponse(resp *http.Response) (CreationSupportedClientListResponse, error) {
	result := CreationSupportedClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateResourceSupportedResponseList); err != nil {
		return CreationSupportedClientListResponse{}, err
	}
	return result, nil
}
