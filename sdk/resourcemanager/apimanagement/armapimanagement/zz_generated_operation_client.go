//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// OperationClient contains the methods for the Operation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationClient, error) {
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
	client := &OperationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByTagsPager - Lists a collection of operations associated with tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// options - OperationClientListByTagsOptions contains the optional parameters for the OperationClient.ListByTags method.
func (client *OperationClient) NewListByTagsPager(resourceGroupName string, serviceName string, apiID string, options *OperationClientListByTagsOptions) *runtime.Pager[OperationClientListByTagsResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationClientListByTagsResponse]{
		More: func(page OperationClientListByTagsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationClientListByTagsResponse) (OperationClientListByTagsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTagsCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationClientListByTagsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OperationClientListByTagsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationClientListByTagsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTagsHandleResponse(resp)
		},
	})
}

// listByTagsCreateRequest creates the ListByTags request.
func (client *OperationClient) listByTagsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *OperationClientListByTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.IncludeNotTaggedOperations != nil {
		reqQP.Set("includeNotTaggedOperations", strconv.FormatBool(*options.IncludeNotTaggedOperations))
	}
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTagsHandleResponse handles the ListByTags response.
func (client *OperationClient) listByTagsHandleResponse(resp *http.Response) (OperationClientListByTagsResponse, error) {
	result := OperationClientListByTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagResourceCollection); err != nil {
		return OperationClientListByTagsResponse{}, err
	}
	return result, nil
}
