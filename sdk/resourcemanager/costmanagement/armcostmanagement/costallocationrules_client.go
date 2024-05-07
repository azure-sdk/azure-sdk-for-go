//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

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

// CostAllocationRulesClient contains the methods for the CostAllocationRules group.
// Don't use this type directly, use NewCostAllocationRulesClient() instead.
type CostAllocationRulesClient struct {
	internal *arm.Client
}

// NewCostAllocationRulesClient creates a new instance of CostAllocationRulesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCostAllocationRulesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CostAllocationRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CostAllocationRulesClient{
		internal: cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks availability and correctness of a name for a cost allocation rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - billingAccountID - BillingAccount ID
//   - costAllocationRuleCheckNameAvailabilityRequest - Cost allocation rule to be created or updated
//   - options - CostAllocationRulesClientCheckNameAvailabilityOptions contains the optional parameters for the CostAllocationRulesClient.CheckNameAvailability
//     method.
func (client *CostAllocationRulesClient) CheckNameAvailability(ctx context.Context, billingAccountID string, costAllocationRuleCheckNameAvailabilityRequest CostAllocationRuleCheckNameAvailabilityRequest, options *CostAllocationRulesClientCheckNameAvailabilityOptions) (CostAllocationRulesClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "CostAllocationRulesClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, billingAccountID, costAllocationRuleCheckNameAvailabilityRequest, options)
	if err != nil {
		return CostAllocationRulesClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CostAllocationRulesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CostAllocationRulesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *CostAllocationRulesClient) checkNameAvailabilityCreateRequest(ctx context.Context, billingAccountID string, costAllocationRuleCheckNameAvailabilityRequest CostAllocationRuleCheckNameAvailabilityRequest, options *CostAllocationRulesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/checkNameAvailability"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, costAllocationRuleCheckNameAvailabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *CostAllocationRulesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (CostAllocationRulesClientCheckNameAvailabilityResponse, error) {
	result := CostAllocationRulesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CostAllocationRuleCheckNameAvailabilityResponse); err != nil {
		return CostAllocationRulesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Create/Update a rule to allocate cost between different resources within a billing account or enterprise
// enrollment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - billingAccountID - BillingAccount ID
//   - ruleName - Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_'
//     and '-'. The max length is 260 characters.
//   - costAllocationRule - Cost allocation rule to be created or updated
//   - options - CostAllocationRulesClientCreateOrUpdateOptions contains the optional parameters for the CostAllocationRulesClient.CreateOrUpdate
//     method.
func (client *CostAllocationRulesClient) CreateOrUpdate(ctx context.Context, billingAccountID string, ruleName string, costAllocationRule CostAllocationRuleDefinition, options *CostAllocationRulesClientCreateOrUpdateOptions) (CostAllocationRulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "CostAllocationRulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, billingAccountID, ruleName, costAllocationRule, options)
	if err != nil {
		return CostAllocationRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CostAllocationRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CostAllocationRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CostAllocationRulesClient) createOrUpdateCreateRequest(ctx context.Context, billingAccountID string, ruleName string, costAllocationRule CostAllocationRuleDefinition, options *CostAllocationRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/{ruleName}"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, costAllocationRule); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CostAllocationRulesClient) createOrUpdateHandleResponse(resp *http.Response) (CostAllocationRulesClientCreateOrUpdateResponse, error) {
	result := CostAllocationRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CostAllocationRuleDefinition); err != nil {
		return CostAllocationRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete cost allocation rule for billing account or enterprise enrollment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - billingAccountID - BillingAccount ID
//   - ruleName - Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_'
//     and '-'. The max length is 260 characters.
//   - options - CostAllocationRulesClientDeleteOptions contains the optional parameters for the CostAllocationRulesClient.Delete
//     method.
func (client *CostAllocationRulesClient) Delete(ctx context.Context, billingAccountID string, ruleName string, options *CostAllocationRulesClientDeleteOptions) (CostAllocationRulesClientDeleteResponse, error) {
	var err error
	const operationName = "CostAllocationRulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, billingAccountID, ruleName, options)
	if err != nil {
		return CostAllocationRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CostAllocationRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CostAllocationRulesClientDeleteResponse{}, err
	}
	return CostAllocationRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CostAllocationRulesClient) deleteCreateRequest(ctx context.Context, billingAccountID string, ruleName string, options *CostAllocationRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/{ruleName}"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a cost allocation rule by rule name and billing account or enterprise enrollment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - billingAccountID - BillingAccount ID
//   - ruleName - Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_'
//     and '-'. The max length is 260 characters.
//   - options - CostAllocationRulesClientGetOptions contains the optional parameters for the CostAllocationRulesClient.Get method.
func (client *CostAllocationRulesClient) Get(ctx context.Context, billingAccountID string, ruleName string, options *CostAllocationRulesClientGetOptions) (CostAllocationRulesClientGetResponse, error) {
	var err error
	const operationName = "CostAllocationRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, billingAccountID, ruleName, options)
	if err != nil {
		return CostAllocationRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CostAllocationRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CostAllocationRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CostAllocationRulesClient) getCreateRequest(ctx context.Context, billingAccountID string, ruleName string, options *CostAllocationRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/{ruleName}"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CostAllocationRulesClient) getHandleResponse(resp *http.Response) (CostAllocationRulesClientGetResponse, error) {
	result := CostAllocationRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CostAllocationRuleDefinition); err != nil {
		return CostAllocationRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the list of all cost allocation rules for a billing account or enterprise enrollment.
//
// Generated from API version 2023-11-01
//   - billingAccountID - BillingAccount ID
//   - options - CostAllocationRulesClientListOptions contains the optional parameters for the CostAllocationRulesClient.NewListPager
//     method.
func (client *CostAllocationRulesClient) NewListPager(billingAccountID string, options *CostAllocationRulesClientListOptions) *runtime.Pager[CostAllocationRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CostAllocationRulesClientListResponse]{
		More: func(page CostAllocationRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CostAllocationRulesClientListResponse) (CostAllocationRulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CostAllocationRulesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, billingAccountID, options)
			}, nil)
			if err != nil {
				return CostAllocationRulesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CostAllocationRulesClient) listCreateRequest(ctx context.Context, billingAccountID string, options *CostAllocationRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CostAllocationRulesClient) listHandleResponse(resp *http.Response) (CostAllocationRulesClientListResponse, error) {
	result := CostAllocationRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CostAllocationRuleList); err != nil {
		return CostAllocationRulesClientListResponse{}, err
	}
	return result, nil
}