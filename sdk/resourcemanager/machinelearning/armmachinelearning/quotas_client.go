//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

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

// QuotasClient contains the methods for the Quotas group.
// Don't use this type directly, use NewQuotasClient() instead.
type QuotasClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewQuotasClient creates a new instance of QuotasClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotasClient, error) {
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
	client := &QuotasClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Gets the currently assigned Workspace Quotas based on VMFamily.
//
// Generated from API version 2023-04-01
//   - location - The location for which resource usage is queried.
//   - options - QuotasClientListOptions contains the optional parameters for the QuotasClient.NewListPager method.
func (client *QuotasClient) NewListPager(location string, options *QuotasClientListOptions) *runtime.Pager[QuotasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QuotasClientListResponse]{
		More: func(page QuotasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotasClientListResponse) (QuotasClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QuotasClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return QuotasClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QuotasClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *QuotasClient) listCreateRequest(ctx context.Context, location string, options *QuotasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/quotas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QuotasClient) listHandleResponse(resp *http.Response) (QuotasClientListResponse, error) {
	result := QuotasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListWorkspaceQuotas); err != nil {
		return QuotasClientListResponse{}, err
	}
	return result, nil
}

// Update - Update quota for each VM family in workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - location - The location for update quota is queried.
//   - parameters - Quota update parameters.
//   - options - QuotasClientUpdateOptions contains the optional parameters for the QuotasClient.Update method.
func (client *QuotasClient) Update(ctx context.Context, location string, parameters QuotaUpdateParameters, options *QuotasClientUpdateOptions) (QuotasClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotasClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *QuotasClient) updateCreateRequest(ctx context.Context, location string, parameters QuotaUpdateParameters, options *QuotasClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/updateQuotas"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *QuotasClient) updateHandleResponse(resp *http.Response) (QuotasClientUpdateResponse, error) {
	result := QuotasClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateWorkspaceQuotasResult); err != nil {
		return QuotasClientUpdateResponse{}, err
	}
	return result, nil
}
