//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PolicySetsClient contains the methods for the PolicySets group.
// Don't use this type directly, use NewPolicySetsClient() instead.
type PolicySetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPolicySetsClient creates a new instance of PolicySetsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicySetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicySetsClient, error) {
	cl, err := arm.NewClient(moduleName+".PolicySetsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PolicySetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// EvaluatePolicies - Evaluates lab policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the policy set.
//   - evaluatePoliciesRequest - Request body for evaluating a policy set.
//   - options - PolicySetsClientEvaluatePoliciesOptions contains the optional parameters for the PolicySetsClient.EvaluatePolicies
//     method.
func (client *PolicySetsClient) EvaluatePolicies(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest, options *PolicySetsClientEvaluatePoliciesOptions) (PolicySetsClientEvaluatePoliciesResponse, error) {
	req, err := client.evaluatePoliciesCreateRequest(ctx, resourceGroupName, labName, name, evaluatePoliciesRequest, options)
	if err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicySetsClientEvaluatePoliciesResponse{}, runtime.NewResponseError(resp)
	}
	return client.evaluatePoliciesHandleResponse(resp)
}

// evaluatePoliciesCreateRequest creates the EvaluatePolicies request.
func (client *PolicySetsClient) evaluatePoliciesCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest EvaluatePoliciesRequest, options *PolicySetsClientEvaluatePoliciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{name}/evaluatePolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, evaluatePoliciesRequest)
}

// evaluatePoliciesHandleResponse handles the EvaluatePolicies response.
func (client *PolicySetsClient) evaluatePoliciesHandleResponse(resp *http.Response) (PolicySetsClientEvaluatePoliciesResponse, error) {
	result := PolicySetsClientEvaluatePoliciesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatePoliciesResponse); err != nil {
		return PolicySetsClientEvaluatePoliciesResponse{}, err
	}
	return result, nil
}

// NewListPager - List policy sets in a given lab.
//
// Generated from API version 2021-09-01
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - options - PolicySetsClientListOptions contains the optional parameters for the PolicySetsClient.NewListPager method.
func (client *PolicySetsClient) NewListPager(resourceGroupName string, labName string, options *PolicySetsClientListOptions) *runtime.Pager[PolicySetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicySetsClientListResponse]{
		More: func(page PolicySetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicySetsClientListResponse) (PolicySetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PolicySetsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PolicySetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PolicySetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PolicySetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *PolicySetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PolicySetsClient) listHandleResponse(resp *http.Response) (PolicySetsClientListResponse, error) {
	result := PolicySetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetList); err != nil {
		return PolicySetsClientListResponse{}, err
	}
	return result, nil
}
