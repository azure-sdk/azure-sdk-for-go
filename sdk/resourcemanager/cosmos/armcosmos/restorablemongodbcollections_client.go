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

// RestorableMongodbCollectionsClient contains the methods for the RestorableMongodbCollections group.
// Don't use this type directly, use NewRestorableMongodbCollectionsClient() instead.
type RestorableMongodbCollectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableMongodbCollectionsClient creates a new instance of RestorableMongodbCollectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableMongodbCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableMongodbCollectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableMongodbCollectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Show the event feed of all mutations done on all the Azure Cosmos DB MongoDB collections under a specific
// database. This helps in scenario where container was accidentally deleted. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
//
// Generated from API version 2023-11-15-preview
//   - location - Cosmos DB region, with spaces between words and each word capitalized.
//   - instanceID - The instanceId GUID of a restorable database account.
//   - options - RestorableMongodbCollectionsClientListOptions contains the optional parameters for the RestorableMongodbCollectionsClient.NewListPager
//     method.
func (client *RestorableMongodbCollectionsClient) NewListPager(location string, instanceID string, options *RestorableMongodbCollectionsClientListOptions) *runtime.Pager[RestorableMongodbCollectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableMongodbCollectionsClientListResponse]{
		More: func(page RestorableMongodbCollectionsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableMongodbCollectionsClientListResponse) (RestorableMongodbCollectionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorableMongodbCollectionsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableMongodbCollectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableMongodbCollectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableMongodbCollectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbCollectionsClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbCollectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbCollections"
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
	reqQP.Set("api-version", "2023-11-15-preview")
	if options != nil && options.RestorableMongodbDatabaseRid != nil {
		reqQP.Set("restorableMongodbDatabaseRid", *options.RestorableMongodbDatabaseRid)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbCollectionsClient) listHandleResponse(resp *http.Response) (RestorableMongodbCollectionsClientListResponse, error) {
	result := RestorableMongodbCollectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbCollectionsListResult); err != nil {
		return RestorableMongodbCollectionsClientListResponse{}, err
	}
	return result, nil
}
