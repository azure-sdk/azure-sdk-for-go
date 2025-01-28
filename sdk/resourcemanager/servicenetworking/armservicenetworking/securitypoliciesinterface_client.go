// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armservicenetworking

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

// SecurityPoliciesInterfaceClient contains the methods for the SecurityPoliciesInterface group.
// Don't use this type directly, use NewSecurityPoliciesInterfaceClient() instead.
type SecurityPoliciesInterfaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecurityPoliciesInterfaceClient creates a new instance of SecurityPoliciesInterfaceClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecurityPoliciesInterfaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecurityPoliciesInterfaceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecurityPoliciesInterfaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - securityPolicyName - SecurityPolicy
//   - resource - Resource create parameters.
//   - options - SecurityPoliciesInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the SecurityPoliciesInterfaceClient.BeginCreateOrUpdate
//     method.
func (client *SecurityPoliciesInterfaceClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, resource SecurityPolicy, options *SecurityPoliciesInterfaceClientBeginCreateOrUpdateOptions) (*runtime.Poller[SecurityPoliciesInterfaceClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, trafficControllerName, securityPolicyName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityPoliciesInterfaceClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecurityPoliciesInterfaceClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *SecurityPoliciesInterfaceClient) createOrUpdate(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, resource SecurityPolicy, options *SecurityPoliciesInterfaceClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityPoliciesInterfaceClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, trafficControllerName, securityPolicyName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityPoliciesInterfaceClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, resource SecurityPolicy, _ *SecurityPoliciesInterfaceClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/securityPolicies/{securityPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if securityPolicyName == "" {
		return nil, errors.New("parameter securityPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPolicyName}", url.PathEscape(securityPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - securityPolicyName - SecurityPolicy
//   - options - SecurityPoliciesInterfaceClientBeginDeleteOptions contains the optional parameters for the SecurityPoliciesInterfaceClient.BeginDelete
//     method.
func (client *SecurityPoliciesInterfaceClient) BeginDelete(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, options *SecurityPoliciesInterfaceClientBeginDeleteOptions) (*runtime.Poller[SecurityPoliciesInterfaceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, trafficControllerName, securityPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SecurityPoliciesInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SecurityPoliciesInterfaceClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
func (client *SecurityPoliciesInterfaceClient) deleteOperation(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, options *SecurityPoliciesInterfaceClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SecurityPoliciesInterfaceClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, trafficControllerName, securityPolicyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityPoliciesInterfaceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, _ *SecurityPoliciesInterfaceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/securityPolicies/{securityPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if securityPolicyName == "" {
		return nil, errors.New("parameter securityPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPolicyName}", url.PathEscape(securityPolicyName))
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

// Get - Get a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - securityPolicyName - SecurityPolicy
//   - options - SecurityPoliciesInterfaceClientGetOptions contains the optional parameters for the SecurityPoliciesInterfaceClient.Get
//     method.
func (client *SecurityPoliciesInterfaceClient) Get(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, options *SecurityPoliciesInterfaceClientGetOptions) (SecurityPoliciesInterfaceClientGetResponse, error) {
	var err error
	const operationName = "SecurityPoliciesInterfaceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, trafficControllerName, securityPolicyName, options)
	if err != nil {
		return SecurityPoliciesInterfaceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityPoliciesInterfaceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityPoliciesInterfaceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SecurityPoliciesInterfaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, _ *SecurityPoliciesInterfaceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/securityPolicies/{securityPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if securityPolicyName == "" {
		return nil, errors.New("parameter securityPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPolicyName}", url.PathEscape(securityPolicyName))
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
func (client *SecurityPoliciesInterfaceClient) getHandleResponse(resp *http.Response) (SecurityPoliciesInterfaceClientGetResponse, error) {
	result := SecurityPoliciesInterfaceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPolicy); err != nil {
		return SecurityPoliciesInterfaceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTrafficControllerPager - List SecurityPolicy resources by TrafficController
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - options - SecurityPoliciesInterfaceClientListByTrafficControllerOptions contains the optional parameters for the SecurityPoliciesInterfaceClient.NewListByTrafficControllerPager
//     method.
func (client *SecurityPoliciesInterfaceClient) NewListByTrafficControllerPager(resourceGroupName string, trafficControllerName string, options *SecurityPoliciesInterfaceClientListByTrafficControllerOptions) *runtime.Pager[SecurityPoliciesInterfaceClientListByTrafficControllerResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecurityPoliciesInterfaceClientListByTrafficControllerResponse]{
		More: func(page SecurityPoliciesInterfaceClientListByTrafficControllerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecurityPoliciesInterfaceClientListByTrafficControllerResponse) (SecurityPoliciesInterfaceClientListByTrafficControllerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecurityPoliciesInterfaceClient.NewListByTrafficControllerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTrafficControllerCreateRequest(ctx, resourceGroupName, trafficControllerName, options)
			}, nil)
			if err != nil {
				return SecurityPoliciesInterfaceClientListByTrafficControllerResponse{}, err
			}
			return client.listByTrafficControllerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTrafficControllerCreateRequest creates the ListByTrafficController request.
func (client *SecurityPoliciesInterfaceClient) listByTrafficControllerCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, _ *SecurityPoliciesInterfaceClientListByTrafficControllerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/securityPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
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

// listByTrafficControllerHandleResponse handles the ListByTrafficController response.
func (client *SecurityPoliciesInterfaceClient) listByTrafficControllerHandleResponse(resp *http.Response) (SecurityPoliciesInterfaceClientListByTrafficControllerResponse, error) {
	result := SecurityPoliciesInterfaceClientListByTrafficControllerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPolicyListResult); err != nil {
		return SecurityPoliciesInterfaceClientListByTrafficControllerResponse{}, err
	}
	return result, nil
}

// Update - Update a SecurityPolicy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trafficControllerName - traffic controller name for path
//   - securityPolicyName - SecurityPolicy
//   - properties - The resource properties to be updated.
//   - options - SecurityPoliciesInterfaceClientUpdateOptions contains the optional parameters for the SecurityPoliciesInterfaceClient.Update
//     method.
func (client *SecurityPoliciesInterfaceClient) Update(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, properties SecurityPolicyUpdate, options *SecurityPoliciesInterfaceClientUpdateOptions) (SecurityPoliciesInterfaceClientUpdateResponse, error) {
	var err error
	const operationName = "SecurityPoliciesInterfaceClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, trafficControllerName, securityPolicyName, properties, options)
	if err != nil {
		return SecurityPoliciesInterfaceClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SecurityPoliciesInterfaceClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SecurityPoliciesInterfaceClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SecurityPoliciesInterfaceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, trafficControllerName string, securityPolicyName string, properties SecurityPolicyUpdate, _ *SecurityPoliciesInterfaceClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceNetworking/trafficControllers/{trafficControllerName}/securityPolicies/{securityPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trafficControllerName == "" {
		return nil, errors.New("parameter trafficControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trafficControllerName}", url.PathEscape(trafficControllerName))
	if securityPolicyName == "" {
		return nil, errors.New("parameter securityPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityPolicyName}", url.PathEscape(securityPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SecurityPoliciesInterfaceClient) updateHandleResponse(resp *http.Response) (SecurityPoliciesInterfaceClientUpdateResponse, error) {
	result := SecurityPoliciesInterfaceClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityPolicy); err != nil {
		return SecurityPoliciesInterfaceClientUpdateResponse{}, err
	}
	return result, nil
}
