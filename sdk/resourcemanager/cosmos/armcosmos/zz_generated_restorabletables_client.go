//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// RestorableTablesClient contains the methods for the RestorableTables group.
// Don't use this type directly, use NewRestorableTablesClient() instead.
type RestorableTablesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableTablesClient creates a new instance of RestorableTablesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableTablesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableTablesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Show the event feed of all mutations done on all the Azure Cosmos DB Tables. This helps in scenario where
// table was accidentally deleted. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns an *azcore.ResponseError type.
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableTablesClientListOptions contains the optional parameters for the RestorableTablesClient.List method.
func (client *RestorableTablesClient) NewListPager(location string, instanceID string, options *RestorableTablesClientListOptions) *runtime.Pager[RestorableTablesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RestorableTablesClientListResponse]{
		More: func(page RestorableTablesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableTablesClientListResponse) (RestorableTablesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, instanceID, options)
			if err != nil {
				return RestorableTablesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorableTablesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableTablesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorableTablesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableTablesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableTables"
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
	reqQP.Set("api-version", "2022-04-15-preview")
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableTablesClient) listHandleResponse(resp *http.Response) (RestorableTablesClientListResponse, error) {
	result := RestorableTablesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableTablesListResult); err != nil {
		return RestorableTablesClientListResponse{}, err
	}
	return result, nil
}
