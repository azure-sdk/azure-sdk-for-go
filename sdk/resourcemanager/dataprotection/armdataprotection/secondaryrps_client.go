//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

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

// SecondaryRPsClient contains the methods for the SecondaryRPs group.
// Don't use this type directly, use NewSecondaryRPsClient() instead.
type SecondaryRPsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecondaryRPsClient creates a new instance of SecondaryRPsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecondaryRPsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecondaryRPsClient, error) {
	cl, err := arm.NewClient(moduleName+".SecondaryRPsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecondaryRPsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewFetchPager - Returns a list of Secondary Recovery Points for a DataSource in a vault, that can be used for Cross Region
// Restore.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - Request body for operation
//   - options - SecondaryRPsClientFetchOptions contains the optional parameters for the SecondaryRPsClient.NewFetchPager method.
func (client *SecondaryRPsClient) NewFetchPager(resourceGroupName string, location string, parameters FetchSecondaryRPsRequestParameters, options *SecondaryRPsClientFetchOptions) *runtime.Pager[SecondaryRPsClientFetchResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecondaryRPsClientFetchResponse]{
		More: func(page SecondaryRPsClientFetchResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecondaryRPsClientFetchResponse) (SecondaryRPsClientFetchResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.fetchCreateRequest(ctx, resourceGroupName, location, parameters, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SecondaryRPsClientFetchResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SecondaryRPsClientFetchResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SecondaryRPsClientFetchResponse{}, runtime.NewResponseError(resp)
			}
			return client.fetchHandleResponse(resp)
		},
	})
}

// fetchCreateRequest creates the Fetch request.
func (client *SecondaryRPsClient) fetchCreateRequest(ctx context.Context, resourceGroupName string, location string, parameters FetchSecondaryRPsRequestParameters, options *SecondaryRPsClientFetchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/locations/{location}/fetchSecondaryRecoveryPoints"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// fetchHandleResponse handles the Fetch response.
func (client *SecondaryRPsClient) fetchHandleResponse(resp *http.Response) (SecondaryRPsClientFetchResponse, error) {
	result := SecondaryRPsClientFetchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupRecoveryPointResourceList); err != nil {
		return SecondaryRPsClientFetchResponse{}, err
	}
	return result, nil
}
