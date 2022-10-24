//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// AccessReviewHistoryDefinitionsClient contains the methods for the AccessReviewHistoryDefinitions group.
// Don't use this type directly, use NewAccessReviewHistoryDefinitionsClient() instead.
type AccessReviewHistoryDefinitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccessReviewHistoryDefinitionsClient creates a new instance of AccessReviewHistoryDefinitionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccessReviewHistoryDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewHistoryDefinitionsClient, error) {
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
	client := &AccessReviewHistoryDefinitionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetByID - Get access review history definition by definition Id
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// historyDefinitionID - The id of the access review history definition.
// options - AccessReviewHistoryDefinitionsClientGetByIDOptions contains the optional parameters for the AccessReviewHistoryDefinitionsClient.GetByID
// method.
func (client *AccessReviewHistoryDefinitionsClient) GetByID(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionsClientGetByIDOptions) (AccessReviewHistoryDefinitionsClientGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, historyDefinitionID, options)
	if err != nil {
		return AccessReviewHistoryDefinitionsClientGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewHistoryDefinitionsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewHistoryDefinitionsClientGetByIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewHistoryDefinitionsClient) getByIDCreateRequest(ctx context.Context, historyDefinitionID string, options *AccessReviewHistoryDefinitionsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions/{historyDefinitionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if historyDefinitionID == "" {
		return nil, errors.New("parameter historyDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{historyDefinitionId}", url.PathEscape(historyDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewHistoryDefinitionsClient) getByIDHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionsClientGetByIDResponse, error) {
	result := AccessReviewHistoryDefinitionsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryDefinition); err != nil {
		return AccessReviewHistoryDefinitionsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the accessReviewHistoryDefinitions available from this provider, definition instances are only available
// for 30 days after creation.
// Generated from API version 2021-12-01-preview
// options - AccessReviewHistoryDefinitionsClientListOptions contains the optional parameters for the AccessReviewHistoryDefinitionsClient.List
// method.
func (client *AccessReviewHistoryDefinitionsClient) NewListPager(options *AccessReviewHistoryDefinitionsClientListOptions) *runtime.Pager[AccessReviewHistoryDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewHistoryDefinitionsClientListResponse]{
		More: func(page AccessReviewHistoryDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewHistoryDefinitionsClientListResponse) (AccessReviewHistoryDefinitionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessReviewHistoryDefinitionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessReviewHistoryDefinitionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessReviewHistoryDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewHistoryDefinitionsClient) listCreateRequest(ctx context.Context, options *AccessReviewHistoryDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewHistoryDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewHistoryDefinitionsClient) listHandleResponse(resp *http.Response) (AccessReviewHistoryDefinitionsClientListResponse, error) {
	result := AccessReviewHistoryDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewHistoryDefinitionListResult); err != nil {
		return AccessReviewHistoryDefinitionsClientListResponse{}, err
	}
	return result, nil
}
