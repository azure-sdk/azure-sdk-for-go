//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// RestorableTableResourcesClient contains the methods for the RestorableTableResources group.
// Don't use this type directly, use NewRestorableTableResourcesClient() instead.
type RestorableTableResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableTableResourcesClient creates a new instance of RestorableTableResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableTableResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableTableResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableTableResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Return a list of tables that exist on the account at the given timestamp and location. This helps in scenarios
// to validate what resources exist at given timestamp and location. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission.
//
// Generated from API version 2024-05-15-preview
//   - location - Cosmos DB region, with spaces between words and each word capitalized.
//   - instanceID - The instanceId GUID of a restorable database account.
//   - options - RestorableTableResourcesClientListOptions contains the optional parameters for the RestorableTableResourcesClient.NewListPager
//     method.
func (client *RestorableTableResourcesClient) NewListPager(location string, instanceID string, options *RestorableTableResourcesClientListOptions) *runtime.Pager[RestorableTableResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableTableResourcesClientListResponse]{
		More: func(page RestorableTableResourcesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableTableResourcesClientListResponse) (RestorableTableResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorableTableResourcesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableTableResourcesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableTableResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableTableResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RestorableTableResourcesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableTableResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableTableResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	if options != nil && options.RestoreLocation != nil {
		reqQP.Set("restoreLocation", *options.RestoreLocation)
	}
	if options != nil && options.RestoreTimestampInUTC != nil {
		reqQP.Set("restoreTimestampInUtc", *options.RestoreTimestampInUTC)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableTableResourcesClient) listHandleResponse(resp *http.Response) (RestorableTableResourcesClientListResponse, error) {
	result := RestorableTableResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableTableResourcesListResult); err != nil {
		return RestorableTableResourcesClientListResponse{}, err
	}
	return result, nil
}
