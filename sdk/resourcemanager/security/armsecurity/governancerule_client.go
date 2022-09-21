//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// GovernanceRuleClient contains the methods for the GovernanceRule group.
// Don't use this type directly, use NewGovernanceRuleClient() instead.
type GovernanceRuleClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGovernanceRuleClient creates a new instance of GovernanceRuleClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGovernanceRuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GovernanceRuleClient, error) {
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
	client := &GovernanceRuleClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Get a list of all relevant governanceRules over a subscription level scope
// Generated from API version 2022-01-01-preview
// options - GovernanceRuleClientListOptions contains the optional parameters for the GovernanceRuleClient.List method.
func (client *GovernanceRuleClient) NewListPager(options *GovernanceRuleClientListOptions) *runtime.Pager[GovernanceRuleClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GovernanceRuleClientListResponse]{
		More: func(page GovernanceRuleClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GovernanceRuleClientListResponse) (GovernanceRuleClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GovernanceRuleClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GovernanceRuleClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GovernanceRuleClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GovernanceRuleClient) listCreateRequest(ctx context.Context, options *GovernanceRuleClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/governanceRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GovernanceRuleClient) listHandleResponse(resp *http.Response) (GovernanceRuleClientListResponse, error) {
	result := GovernanceRuleClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceRuleList); err != nil {
		return GovernanceRuleClientListResponse{}, err
	}
	return result, nil
}
