//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

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

// RestorableMongodbResourcesClient contains the methods for the RestorableMongodbResources group.
// Don't use this type directly, use NewRestorableMongodbResourcesClient() instead.
type RestorableMongodbResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableMongodbResourcesClient creates a new instance of RestorableMongodbResourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableMongodbResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableMongodbResourcesClient, error) {
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
	client := &RestorableMongodbResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Return a list of database and collection combo that exist on the account at the given timestamp and location.
// This helps in scenarios to validate what resources exist at given timestamp and location.
// This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-15-preview
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableMongodbResourcesClientListOptions contains the optional parameters for the RestorableMongodbResourcesClient.List
// method.
func (client *RestorableMongodbResourcesClient) NewListPager(location string, instanceID string, options *RestorableMongodbResourcesClientListOptions) *runtime.Pager[RestorableMongodbResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableMongodbResourcesClientListResponse]{
		More: func(page RestorableMongodbResourcesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableMongodbResourcesClientListResponse) (RestorableMongodbResourcesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableMongodbResourcesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorableMongodbResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableMongodbResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbResourcesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbResources"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-15-preview")
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
func (client *RestorableMongodbResourcesClient) listHandleResponse(resp *http.Response) (RestorableMongodbResourcesClientListResponse, error) {
	result := RestorableMongodbResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbResourcesListResult); err != nil {
		return RestorableMongodbResourcesClientListResponse{}, err
	}
	return result, nil
}
