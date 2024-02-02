//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

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

// SetDefinitionsClient contains the methods for the PolicySetDefinitions group.
// Don't use this type directly, use NewSetDefinitionsClient() instead.
type SetDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSetDefinitionsClient creates a new instance of SetDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSetDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SetDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SetDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - This operation creates or updates a policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - policySetDefinitionName - The name of the policy set definition to create.
//   - parameters - The policy set definition properties.
//   - options - SetDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the SetDefinitionsClient.CreateOrUpdate
//     method.
func (client *SetDefinitionsClient) CreateOrUpdate(ctx context.Context, policySetDefinitionName string, parameters SetDefinition, options *SetDefinitionsClientCreateOrUpdateOptions) (SetDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, policySetDefinitionName, parameters, options)
	if err != nil {
		return SetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SetDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, policySetDefinitionName string, parameters SetDefinition, options *SetDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SetDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (SetDefinitionsClientCreateOrUpdateResponse, error) {
	result := SetDefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinition); err != nil {
		return SetDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAtManagementGroup - This operation creates or updates a policy set definition in the given management group
// with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to create.
//   - parameters - The policy set definition properties.
//   - options - SetDefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the SetDefinitionsClient.CreateOrUpdateAtManagementGroup
//     method.
func (client *SetDefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, parameters SetDefinition, options *SetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.CreateOrUpdateAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, parameters, options)
	if err != nil {
		return SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	resp, err := client.createOrUpdateAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateAtManagementGroupCreateRequest creates the CreateOrUpdateAtManagementGroup request.
func (client *SetDefinitionsClient) createOrUpdateAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, parameters SetDefinition, options *SetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateAtManagementGroupHandleResponse handles the CreateOrUpdateAtManagementGroup response.
func (client *SetDefinitionsClient) createOrUpdateAtManagementGroupHandleResponse(resp *http.Response) (SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	result := SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinition); err != nil {
		return SetDefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	return result, nil
}

// Delete - This operation deletes the policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - policySetDefinitionName - The name of the policy set definition to delete.
//   - options - SetDefinitionsClientDeleteOptions contains the optional parameters for the SetDefinitionsClient.Delete method.
func (client *SetDefinitionsClient) Delete(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientDeleteOptions) (SetDefinitionsClientDeleteResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return SetDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientDeleteResponse{}, err
	}
	return SetDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SetDefinitionsClient) deleteCreateRequest(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientDeleteOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtManagementGroup - This operation deletes the policy set definition in the given management group with the given
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to delete.
//   - options - SetDefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the SetDefinitionsClient.DeleteAtManagementGroup
//     method.
func (client *SetDefinitionsClient) DeleteAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *SetDefinitionsClientDeleteAtManagementGroupOptions) (SetDefinitionsClientDeleteAtManagementGroupResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.DeleteAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, options)
	if err != nil {
		return SetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	return SetDefinitionsClientDeleteAtManagementGroupResponse{}, nil
}

// deleteAtManagementGroupCreateRequest creates the DeleteAtManagementGroup request.
func (client *SetDefinitionsClient) deleteAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *SetDefinitionsClientDeleteAtManagementGroupOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - This operation retrieves the policy set definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - SetDefinitionsClientGetOptions contains the optional parameters for the SetDefinitionsClient.Get method.
func (client *SetDefinitionsClient) Get(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientGetOptions) (SetDefinitionsClientGetResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return SetDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SetDefinitionsClient) getCreateRequest(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientGetOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SetDefinitionsClient) getHandleResponse(resp *http.Response) (SetDefinitionsClientGetResponse, error) {
	result := SetDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinition); err != nil {
		return SetDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAtManagementGroup - This operation retrieves the policy set definition in the given management group with the given
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - managementGroupID - The ID of the management group.
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - SetDefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the SetDefinitionsClient.GetAtManagementGroup
//     method.
func (client *SetDefinitionsClient) GetAtManagementGroup(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *SetDefinitionsClientGetAtManagementGroupOptions) (SetDefinitionsClientGetAtManagementGroupResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.GetAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAtManagementGroupCreateRequest(ctx, managementGroupID, policySetDefinitionName, options)
	if err != nil {
		return SetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	resp, err := client.getAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// getAtManagementGroupCreateRequest creates the GetAtManagementGroup request.
func (client *SetDefinitionsClient) getAtManagementGroupCreateRequest(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *SetDefinitionsClientGetAtManagementGroupOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtManagementGroupHandleResponse handles the GetAtManagementGroup response.
func (client *SetDefinitionsClient) getAtManagementGroupHandleResponse(resp *http.Response) (SetDefinitionsClientGetAtManagementGroupResponse, error) {
	result := SetDefinitionsClientGetAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinition); err != nil {
		return SetDefinitionsClientGetAtManagementGroupResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - This operation retrieves the built-in policy set definition with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - policySetDefinitionName - The name of the policy set definition to get.
//   - options - SetDefinitionsClientGetBuiltInOptions contains the optional parameters for the SetDefinitionsClient.GetBuiltIn
//     method.
func (client *SetDefinitionsClient) GetBuiltIn(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientGetBuiltInOptions) (SetDefinitionsClientGetBuiltInResponse, error) {
	var err error
	const operationName = "SetDefinitionsClient.GetBuiltIn"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBuiltInCreateRequest(ctx, policySetDefinitionName, options)
	if err != nil {
		return SetDefinitionsClientGetBuiltInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SetDefinitionsClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SetDefinitionsClientGetBuiltInResponse{}, err
	}
	resp, err := client.getBuiltInHandleResponse(httpResp)
	return resp, err
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *SetDefinitionsClient) getBuiltInCreateRequest(ctx context.Context, policySetDefinitionName string, options *SetDefinitionsClientGetBuiltInOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *SetDefinitionsClient) getBuiltInHandleResponse(resp *http.Response) (SetDefinitionsClientGetBuiltInResponse, error) {
	result := SetDefinitionsClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinition); err != nil {
		return SetDefinitionsClientGetBuiltInResponse{}, err
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
// Generated from API version 2023-04-01
//   - options - SetDefinitionsClientListOptions contains the optional parameters for the SetDefinitionsClient.NewListPager method.
func (client *SetDefinitionsClient) NewListPager(options *SetDefinitionsClientListOptions) *runtime.Pager[SetDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SetDefinitionsClientListResponse]{
		More: func(page SetDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SetDefinitionsClientListResponse) (SetDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SetDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SetDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SetDefinitionsClient) listCreateRequest(ctx context.Context, options *SetDefinitionsClientListOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
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
func (client *SetDefinitionsClient) listHandleResponse(resp *http.Response) (SetDefinitionsClientListResponse, error) {
	result := SetDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinitionListResult); err != nil {
		return SetDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBuiltInPager - This operation retrieves a list of all the built-in policy set definitions that match the optional
// given $filter. If $filter='category -eq {value}' is provided, the returned list only includes all
// built-in policy set definitions whose category match the {value}.
//
// Generated from API version 2023-04-01
//   - options - SetDefinitionsClientListBuiltInOptions contains the optional parameters for the SetDefinitionsClient.NewListBuiltInPager
//     method.
func (client *SetDefinitionsClient) NewListBuiltInPager(options *SetDefinitionsClientListBuiltInOptions) *runtime.Pager[SetDefinitionsClientListBuiltInResponse] {
	return runtime.NewPager(runtime.PagingHandler[SetDefinitionsClientListBuiltInResponse]{
		More: func(page SetDefinitionsClientListBuiltInResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SetDefinitionsClientListBuiltInResponse) (SetDefinitionsClientListBuiltInResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SetDefinitionsClient.NewListBuiltInPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBuiltInCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SetDefinitionsClientListBuiltInResponse{}, err
			}
			return client.listBuiltInHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBuiltInCreateRequest creates the ListBuiltIn request.
func (client *SetDefinitionsClient) listBuiltInCreateRequest(ctx context.Context, options *SetDefinitionsClientListBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policySetDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
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
func (client *SetDefinitionsClient) listBuiltInHandleResponse(resp *http.Response) (SetDefinitionsClientListBuiltInResponse, error) {
	result := SetDefinitionsClientListBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinitionListResult); err != nil {
		return SetDefinitionsClientListBuiltInResponse{}, err
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
// Generated from API version 2023-04-01
//   - managementGroupID - The ID of the management group.
//   - options - SetDefinitionsClientListByManagementGroupOptions contains the optional parameters for the SetDefinitionsClient.NewListByManagementGroupPager
//     method.
func (client *SetDefinitionsClient) NewListByManagementGroupPager(managementGroupID string, options *SetDefinitionsClientListByManagementGroupOptions) *runtime.Pager[SetDefinitionsClientListByManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SetDefinitionsClientListByManagementGroupResponse]{
		More: func(page SetDefinitionsClientListByManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SetDefinitionsClientListByManagementGroupResponse) (SetDefinitionsClientListByManagementGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SetDefinitionsClient.NewListByManagementGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagementGroupCreateRequest(ctx, managementGroupID, options)
			}, nil)
			if err != nil {
				return SetDefinitionsClientListByManagementGroupResponse{}, err
			}
			return client.listByManagementGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagementGroupCreateRequest creates the ListByManagementGroup request.
func (client *SetDefinitionsClient) listByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *SetDefinitionsClientListByManagementGroupOptions) (*policy.Request, error) {
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
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
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
func (client *SetDefinitionsClient) listByManagementGroupHandleResponse(resp *http.Response) (SetDefinitionsClientListByManagementGroupResponse, error) {
	result := SetDefinitionsClientListByManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SetDefinitionListResult); err != nil {
		return SetDefinitionsClientListByManagementGroupResponse{}, err
	}
	return result, nil
}
