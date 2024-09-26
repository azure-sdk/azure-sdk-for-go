//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// DaprComponentResiliencyPoliciesClient contains the methods for the DaprComponentResiliencyPolicies group.
// Don't use this type directly, use NewDaprComponentResiliencyPoliciesClient() instead.
type DaprComponentResiliencyPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDaprComponentResiliencyPoliciesClient creates a new instance of DaprComponentResiliencyPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDaprComponentResiliencyPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprComponentResiliencyPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DaprComponentResiliencyPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a resiliency policy for a Dapr component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - name - Name of the Dapr Component Resiliency Policy.
//   - daprComponentResiliencyPolicyEnvelope - Configuration details of the Dapr Component Resiliency Policy.
//   - options - DaprComponentResiliencyPoliciesClientCreateOrUpdateOptions contains the optional parameters for the DaprComponentResiliencyPoliciesClient.CreateOrUpdate
//     method.
func (client *DaprComponentResiliencyPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, daprComponentResiliencyPolicyEnvelope DaprComponentResiliencyPolicy, options *DaprComponentResiliencyPoliciesClientCreateOrUpdateOptions) (DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DaprComponentResiliencyPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, componentName, name, daprComponentResiliencyPolicyEnvelope, options)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprComponentResiliencyPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, daprComponentResiliencyPolicyEnvelope DaprComponentResiliencyPolicy, options *DaprComponentResiliencyPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, daprComponentResiliencyPolicyEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprComponentResiliencyPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse, error) {
	result := DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponentResiliencyPolicy); err != nil {
		return DaprComponentResiliencyPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a resiliency policy for a Dapr component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - name - Name of the Dapr Component Resiliency Policy.
//   - options - DaprComponentResiliencyPoliciesClientDeleteOptions contains the optional parameters for the DaprComponentResiliencyPoliciesClient.Delete
//     method.
func (client *DaprComponentResiliencyPoliciesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *DaprComponentResiliencyPoliciesClientDeleteOptions) (DaprComponentResiliencyPoliciesClientDeleteResponse, error) {
	var err error
	const operationName = "DaprComponentResiliencyPoliciesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, componentName, name, options)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DaprComponentResiliencyPoliciesClientDeleteResponse{}, err
	}
	return DaprComponentResiliencyPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprComponentResiliencyPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *DaprComponentResiliencyPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Dapr component resiliency policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - name - Name of the Dapr Component Resiliency Policy.
//   - options - DaprComponentResiliencyPoliciesClientGetOptions contains the optional parameters for the DaprComponentResiliencyPoliciesClient.Get
//     method.
func (client *DaprComponentResiliencyPoliciesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *DaprComponentResiliencyPoliciesClientGetOptions) (DaprComponentResiliencyPoliciesClientGetResponse, error) {
	var err error
	const operationName = "DaprComponentResiliencyPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, componentName, name, options)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentResiliencyPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DaprComponentResiliencyPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DaprComponentResiliencyPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, name string, options *DaprComponentResiliencyPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprComponentResiliencyPoliciesClient) getHandleResponse(resp *http.Response) (DaprComponentResiliencyPoliciesClientGetResponse, error) {
	result := DaprComponentResiliencyPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponentResiliencyPolicy); err != nil {
		return DaprComponentResiliencyPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the resiliency policies for a Dapr component.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - options - DaprComponentResiliencyPoliciesClientListOptions contains the optional parameters for the DaprComponentResiliencyPoliciesClient.NewListPager
//     method.
func (client *DaprComponentResiliencyPoliciesClient) NewListPager(resourceGroupName string, environmentName string, componentName string, options *DaprComponentResiliencyPoliciesClientListOptions) *runtime.Pager[DaprComponentResiliencyPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DaprComponentResiliencyPoliciesClientListResponse]{
		More: func(page DaprComponentResiliencyPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprComponentResiliencyPoliciesClientListResponse) (DaprComponentResiliencyPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DaprComponentResiliencyPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, componentName, options)
			}, nil)
			if err != nil {
				return DaprComponentResiliencyPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DaprComponentResiliencyPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentResiliencyPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}/resiliencyPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprComponentResiliencyPoliciesClient) listHandleResponse(resp *http.Response) (DaprComponentResiliencyPoliciesClientListResponse, error) {
	result := DaprComponentResiliencyPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponentResiliencyPoliciesCollection); err != nil {
		return DaprComponentResiliencyPoliciesClientListResponse{}, err
	}
	return result, nil
}
