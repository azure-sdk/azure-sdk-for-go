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

// RestorableMongodbDatabasesClient contains the methods for the RestorableMongodbDatabases group.
// Don't use this type directly, use NewRestorableMongodbDatabasesClient() instead.
type RestorableMongodbDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableMongodbDatabasesClient creates a new instance of RestorableMongodbDatabasesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableMongodbDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableMongodbDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableMongodbDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Show the event feed of all mutations done on all the Azure Cosmos DB MongoDB databases under the restorable
// account. This helps in scenario where database was accidentally deleted to get the deletion
// time. This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
//
// Generated from API version 2024-08-15
//   - location - Cosmos DB region, with spaces between words and each word capitalized.
//   - instanceID - The instanceId GUID of a restorable database account.
//   - options - RestorableMongodbDatabasesClientListOptions contains the optional parameters for the RestorableMongodbDatabasesClient.NewListPager
//     method.
func (client *RestorableMongodbDatabasesClient) NewListPager(location string, instanceID string, options *RestorableMongodbDatabasesClientListOptions) *runtime.Pager[RestorableMongodbDatabasesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableMongodbDatabasesClientListResponse]{
		More: func(page RestorableMongodbDatabasesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableMongodbDatabasesClientListResponse) (RestorableMongodbDatabasesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RestorableMongodbDatabasesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableMongodbDatabasesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableMongodbDatabasesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableMongodbDatabasesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbDatabasesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbDatabasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbDatabases"
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
	reqQP.Set("api-version", "2024-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbDatabasesClient) listHandleResponse(resp *http.Response) (RestorableMongodbDatabasesClientListResponse, error) {
	result := RestorableMongodbDatabasesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbDatabasesListResult); err != nil {
		return RestorableMongodbDatabasesClientListResponse{}, err
	}
	return result, nil
}
