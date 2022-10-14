//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcomplianceautomation

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// ReportsClient contains the methods for the Reports group.
// Don't use this type directly, use NewReportsClient() instead.
type ReportsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReportsClient creates a new instance of ReportsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReportsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReportsClient, error) {
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
	client := &ReportsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Get the AppComplianceAutomation report list.
// Generated from API version 2022-11-16-preview
// options - ReportsClientListOptions contains the optional parameters for the ReportsClient.List method.
func (client *ReportsClient) NewListPager(options *ReportsClientListOptions) *runtime.Pager[ReportsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReportsClientListResponse]{
		More: func(page ReportsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReportsClientListResponse) (ReportsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReportsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReportsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReportsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReportsClient) listCreateRequest(ctx context.Context, options *ReportsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppComplianceAutomation/reports"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-16-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.OfferGUID != nil {
		reqQP.Set("offerGuid", *options.OfferGUID)
	}
	if options != nil && options.ReportCreatorTenantID != nil {
		reqQP.Set("reportCreatorTenantId", *options.ReportCreatorTenantID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReportsClient) listHandleResponse(resp *http.Response) (ReportsClientListResponse, error) {
	result := ReportsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReportResourceList); err != nil {
		return ReportsClientListResponse{}, err
	}
	return result, nil
}
