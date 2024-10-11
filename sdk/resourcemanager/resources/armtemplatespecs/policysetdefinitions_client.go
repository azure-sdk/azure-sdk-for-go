//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

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

// PolicySetDefinitionsClient contains the methods for the PolicySetDefinitions group.
// Don't use this type directly, use NewPolicySetDefinitionsClient() instead.
type PolicySetDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPolicySetDefinitionsClient creates a new instance of PolicySetDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicySetDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicySetDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PolicySetDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - This operation creates or updates a policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - policySetDefinitionName - The name of the policy set definition to create.
//   - parameters - The policy set definition properties.
//   - options - PolicySetDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the PolicySetDefinitionsClient.CreateOrUpdate
//     method.
func (client *PolicySetDefinitionsClient) CreateOrUpdate(ctx context.Context, policySetDefinitionName string, parameters PolicySetDefinition, options *PolicySetDefinitionsClientCreateOrUpdateOptions) (PolicySetDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, policySetDefinitionName, parameters, options)
	if err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PolicySetDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, policySetDefinitionName string, parameters PolicySetDefinition, options *PolicySetDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PolicySetDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (PolicySetDefinitionsClientCreateOrUpdateResponse, error) {
	result := PolicySetDefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinition); err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAtManagementGroup - This operation creates or updates a policy set definition in the given management group
// with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to create.
//   - parameters - The policy set definition properties.
//   - options - PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup
//     method.
func (client *PolicySetDefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, parameters PolicySetDefinition, options *PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, parameters, options)
	if err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	resp, err := client.createOrUpdateAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateAtManagementGroupCreateRequest creates the CreateOrUpdateAtManagementGroup request.
func (client *PolicySetDefinitionsClient) createOrUpdateAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, parameters PolicySetDefinition, options *PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateAtManagementGroupHandleResponse handles the CreateOrUpdateAtManagementGroup response.
func (client *PolicySetDefinitionsClient) createOrUpdateAtManagementGroupHandleResponse(resp *http.Response) (PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	result := PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinition); err != nil {
		return PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	return result, nil
}

// Delete - This operation deletes the policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - policySetDefinitionName - The name of the policy set definition to delete.
//   - options - PolicySetDefinitionsClientDeleteOptions contains the optional parameters for the PolicySetDefinitionsClient.Delete
//     method.
func (client *PolicySetDefinitionsClient) Delete(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientDeleteOptions) (PolicySetDefinitionsClientDeleteResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return PolicySetDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientDeleteResponse{}, err
	}
	return PolicySetDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PolicySetDefinitionsClient) deleteCreateRequest(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtManagementGroup - This operation deletes the policy set definition in the given management group with the given
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to delete.
//   - options - PolicySetDefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.DeleteAtManagementGroup
//     method.
func (client *PolicySetDefinitionsClient) DeleteAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *PolicySetDefinitionsClientDeleteAtManagementGroupOptions) (PolicySetDefinitionsClientDeleteAtManagementGroupResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.DeleteAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, options)
	if err != nil {
		return PolicySetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	return PolicySetDefinitionsClientDeleteAtManagementGroupResponse{}, nil
}

// deleteAtManagementGroupCreateRequest creates the DeleteAtManagementGroup request.
func (client *PolicySetDefinitionsClient) deleteAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *PolicySetDefinitionsClientDeleteAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - This operation retrieves the policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - PolicySetDefinitionsClientGetOptions contains the optional parameters for the PolicySetDefinitionsClient.Get
//     method.
func (client *PolicySetDefinitionsClient) Get(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientGetOptions) (PolicySetDefinitionsClientGetResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return PolicySetDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PolicySetDefinitionsClient) getCreateRequest(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PolicySetDefinitionsClient) getHandleResponse(resp *http.Response) (PolicySetDefinitionsClientGetResponse, error) {
	result := PolicySetDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinition); err != nil {
		return PolicySetDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAtManagementGroup - This operation retrieves the policy set definition in the given management group with the given
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - PolicySetDefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.GetAtManagementGroup
//     method.
func (client *PolicySetDefinitionsClient) GetAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *PolicySetDefinitionsClientGetAtManagementGroupOptions) (PolicySetDefinitionsClientGetAtManagementGroupResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.GetAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, options)
	if err != nil {
		return PolicySetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	resp, err := client.getAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// getAtManagementGroupCreateRequest creates the GetAtManagementGroup request.
func (client *PolicySetDefinitionsClient) getAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *PolicySetDefinitionsClientGetAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtManagementGroupHandleResponse handles the GetAtManagementGroup response.
func (client *PolicySetDefinitionsClient) getAtManagementGroupHandleResponse(resp *http.Response) (PolicySetDefinitionsClientGetAtManagementGroupResponse, error) {
	result := PolicySetDefinitionsClientGetAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinition); err != nil {
		return PolicySetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - This operation retrieves the built-in policy set definition with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - PolicySetDefinitionsClientGetBuiltInOptions contains the optional parameters for the PolicySetDefinitionsClient.GetBuiltIn
//     method.
func (client *PolicySetDefinitionsClient) GetBuiltIn(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientGetBuiltInOptions) (PolicySetDefinitionsClientGetBuiltInResponse, error) {
	var err error
	const operationName = "PolicySetDefinitionsClient.GetBuiltIn"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBuiltInCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return PolicySetDefinitionsClientGetBuiltInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicySetDefinitionsClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicySetDefinitionsClientGetBuiltInResponse{}, err
	}
	resp, err := client.getBuiltInHandleResponse(httpResp)
	return resp, err
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *PolicySetDefinitionsClient) getBuiltInCreateRequest(ctx context.Context, policySetDefinitionName string, options *PolicySetDefinitionsClientGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}"
	if policySetDefinitionName == "" {
		return nil, errors.New("parameter policySetDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policySetDefinitionName}", url.PathEscape(policySetDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *PolicySetDefinitionsClient) getBuiltInHandleResponse(resp *http.Response) (PolicySetDefinitionsClientGetBuiltInResponse, error) {
	result := PolicySetDefinitionsClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinition); err != nil {
		return PolicySetDefinitionsClientGetBuiltInResponse{}, err
	}
	return result, nil
}

// NewListPager - This operation retrieves a list of all the policy set definitions in a given subscription that match the
// optional given $filter. Valid values for $filter are: 'atExactScope()', 'policyType -eq
// {value}' or 'category eq '{value}”. If $filter is not provided, the unfiltered list includes all policy set definitions
// associated with the subscription, including those that apply directly or from
// management groups that contain the given subscription. If $filter=atExactScope() is provided, the returned list only includes
// all policy set definitions that at the given subscription. If
// $filter='policyType -eq {value}' is provided, the returned list only includes all policy set definitions whose type match
// the {value}. Possible policyType values are NotSpecified, BuiltIn and Custom.
// If $filter='category -eq {value}' is provided, the returned list only includes all policy set definitions whose category
// match the {value}.
//
// Generated from API version 2025-01-01
//   - options - PolicySetDefinitionsClientListOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListPager
//     method.
func (client *PolicySetDefinitionsClient) NewListPager(options *PolicySetDefinitionsClientListOptions) *runtime.Pager[PolicySetDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicySetDefinitionsClientListResponse]{
		More: func(page PolicySetDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicySetDefinitionsClientListResponse) (PolicySetDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PolicySetDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PolicySetDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PolicySetDefinitionsClient) listCreateRequest(ctx context.Context, options *PolicySetDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-01-01")
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
func (client *PolicySetDefinitionsClient) listHandleResponse(resp *http.Response) (PolicySetDefinitionsClientListResponse, error) {
	result := PolicySetDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinitionListResult); err != nil {
		return PolicySetDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBuiltInPager - This operation retrieves a list of all the built-in policy set definitions that match the optional
// given $filter. If $filter='category -eq {value}' is provided, the returned list only includes all
// built-in policy set definitions whose category match the {value}.
//
// Generated from API version 2025-01-01
//   - options - PolicySetDefinitionsClientListBuiltInOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListBuiltInPager
//     method.
func (client *PolicySetDefinitionsClient) NewListBuiltInPager(options *PolicySetDefinitionsClientListBuiltInOptions) *runtime.Pager[PolicySetDefinitionsClientListBuiltInResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicySetDefinitionsClientListBuiltInResponse]{
		More: func(page PolicySetDefinitionsClientListBuiltInResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicySetDefinitionsClientListBuiltInResponse) (PolicySetDefinitionsClientListBuiltInResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PolicySetDefinitionsClient.NewListBuiltInPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBuiltInCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PolicySetDefinitionsClientListBuiltInResponse{}, err
			}
			return client.listBuiltInHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBuiltInCreateRequest creates the ListBuiltIn request.
func (client *PolicySetDefinitionsClient) listBuiltInCreateRequest(ctx context.Context, options *PolicySetDefinitionsClientListBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policySetDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBuiltInHandleResponse handles the ListBuiltIn response.
func (client *PolicySetDefinitionsClient) listBuiltInHandleResponse(resp *http.Response) (PolicySetDefinitionsClientListBuiltInResponse, error) {
	result := PolicySetDefinitionsClientListBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinitionListResult); err != nil {
		return PolicySetDefinitionsClientListBuiltInResponse{}, err
	}
	return result, nil
}

// NewListByManagementGroupPager - This operation retrieves a list of all the policy set definitions in a given management
// group that match the optional given $filter. Valid values for $filter are: 'atExactScope()', 'policyType -eq
// {value}' or 'category eq '{value}”. If $filter is not provided, the unfiltered list includes all policy set definitions
// associated with the management group, including those that apply directly or
// from management groups that contain the given management group. If $filter=atExactScope() is provided, the returned list
// only includes all policy set definitions that at the given management group. If
// $filter='policyType -eq {value}' is provided, the returned list only includes all policy set definitions whose type match
// the {value}. Possible policyType values are NotSpecified, BuiltIn and Custom.
// If $filter='category -eq {value}' is provided, the returned list only includes all policy set definitions whose category
// match the {value}.
//
// Generated from API version 2025-01-01
//   - managementGroupID - The ID of the management group.
//   - options - PolicySetDefinitionsClientListByManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListByManagementGroupPager
//     method.
func (client *PolicySetDefinitionsClient) NewListByManagementGroupPager(managementGroupID string, options *PolicySetDefinitionsClientListByManagementGroupOptions) *runtime.Pager[PolicySetDefinitionsClientListByManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicySetDefinitionsClientListByManagementGroupResponse]{
		More: func(page PolicySetDefinitionsClientListByManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicySetDefinitionsClientListByManagementGroupResponse) (PolicySetDefinitionsClientListByManagementGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PolicySetDefinitionsClient.NewListByManagementGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagementGroupCreateRequest(ctx, managementGroupID, options)
			}, nil)
			if err != nil {
				return PolicySetDefinitionsClientListByManagementGroupResponse{}, err
			}
			return client.listByManagementGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagementGroupCreateRequest creates the ListByManagementGroup request.
func (client *PolicySetDefinitionsClient) listByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *PolicySetDefinitionsClientListByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagementGroupHandleResponse handles the ListByManagementGroup response.
func (client *PolicySetDefinitionsClient) listByManagementGroupHandleResponse(resp *http.Response) (PolicySetDefinitionsClientListByManagementGroupResponse, error) {
	result := PolicySetDefinitionsClientListByManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicySetDefinitionListResult); err != nil {
		return PolicySetDefinitionsClientListByManagementGroupResponse{}, err
	}
	return result, nil
}
