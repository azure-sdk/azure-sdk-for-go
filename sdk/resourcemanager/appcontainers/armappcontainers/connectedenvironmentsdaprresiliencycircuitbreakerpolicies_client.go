//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient contains the methods for the ConnectedEnvironmentsDaprResiliencyCircuitBreakerPolicies group.
// Don't use this type directly, use NewConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient() instead.
type ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient creates a new instance of ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient, error) {
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
	client := &ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Dapr Resiliency Circuit Breaker Policy in a Connected Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - policyName - Name of the Dapr Resiliency Circuit Breaker Policy.
//   - daprResiliencyCircuitBreakerPoliciesEnvelope - Configuration details of the Dapr Resiliency Circuit Breaker Policy.
//   - options - ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions contains the optional parameters
//     for the ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient.CreateOrUpdate method.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, daprResiliencyCircuitBreakerPoliciesEnvelope DaprResiliencyCircuitBreakerPolicy, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, policyName, daprResiliencyCircuitBreakerPoliciesEnvelope, options)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, daprResiliencyCircuitBreakerPoliciesEnvelope DaprResiliencyCircuitBreakerPolicy, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, daprResiliencyCircuitBreakerPoliciesEnvelope)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse, error) {
	result := ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPolicy); err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Dapr Resiliency Circuit Breaker Policy from a Connected Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - policyName - Name of the Dapr Resiliency Circuit Breaker Policy.
//   - options - ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteOptions contains the optional parameters
//     for the ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient.Delete method.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) Delete(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteOptions) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, policyName, options)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Dapr Resiliency Circuit Breaker policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - policyName - Name of the Dapr resiliency circuit breaker policy.
//   - options - ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetOptions contains the optional parameters for
//     the ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient.Get method.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) Get(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetOptions) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, policyName, options)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, policyName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) getHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse, error) {
	result := ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPolicy); err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Dapr Resiliency Circuit Breaker Policies for a connected environment.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the Connected Environment.
//   - options - ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListOptions contains the optional parameters for
//     the ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient.NewListPager method.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) NewListPager(resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListOptions) *runtime.Pager[ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse]{
		More: func(page ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprResiliencyCircuitBreakerPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClient) listHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse, error) {
	result := ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPoliciesCollection); err != nil {
		return ConnectedEnvironmentsDaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
	}
	return result, nil
}
