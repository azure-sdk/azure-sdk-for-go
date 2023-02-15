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

// DaprResiliencyCircuitBreakerPoliciesClient contains the methods for the DaprResiliencyCircuitBreakerPolicies group.
// Don't use this type directly, use NewDaprResiliencyCircuitBreakerPoliciesClient() instead.
type DaprResiliencyCircuitBreakerPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDaprResiliencyCircuitBreakerPoliciesClient creates a new instance of DaprResiliencyCircuitBreakerPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDaprResiliencyCircuitBreakerPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprResiliencyCircuitBreakerPoliciesClient, error) {
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
	client := &DaprResiliencyCircuitBreakerPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Dapr Resiliency Circuit Breaker Policy in a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - policyName - Name of the Dapr Resiliency Circuit Breaker Policy.
//   - daprResiliencyCircuitBreakerPoliciesEnvelope - Configuration details of the Dapr Resiliency Circuit Breaker Policy.
//   - options - DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions contains the optional parameters for the DaprResiliencyCircuitBreakerPoliciesClient.CreateOrUpdate
//     method.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, policyName string, daprResiliencyCircuitBreakerPoliciesEnvelope DaprResiliencyCircuitBreakerPolicy, options *DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions) (DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, policyName, daprResiliencyCircuitBreakerPoliciesEnvelope, options)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, policyName string, daprResiliencyCircuitBreakerPoliciesEnvelope DaprResiliencyCircuitBreakerPolicy, options *DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
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
func (client *DaprResiliencyCircuitBreakerPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse, error) {
	result := DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPolicy); err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Dapr Resiliency Circuit Breaker Policy from a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - policyName - Name of the Dapr Resiliency Circuit Breaker Policy.
//   - options - DaprResiliencyCircuitBreakerPoliciesClientDeleteOptions contains the optional parameters for the DaprResiliencyCircuitBreakerPoliciesClient.Delete
//     method.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, policyName string, options *DaprResiliencyCircuitBreakerPoliciesClientDeleteOptions) (DaprResiliencyCircuitBreakerPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, policyName, options)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DaprResiliencyCircuitBreakerPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, policyName string, options *DaprResiliencyCircuitBreakerPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
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
//   - environmentName - Name of the Managed Environment.
//   - policyName - Name of the Dapr resiliency circuit breaker policy.
//   - options - DaprResiliencyCircuitBreakerPoliciesClientGetOptions contains the optional parameters for the DaprResiliencyCircuitBreakerPoliciesClient.Get
//     method.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, policyName string, options *DaprResiliencyCircuitBreakerPoliciesClientGetOptions) (DaprResiliencyCircuitBreakerPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, policyName, options)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, policyName string, options *DaprResiliencyCircuitBreakerPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprResiliencyCircuitBreakerPolicies/{policyName}"
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
func (client *DaprResiliencyCircuitBreakerPoliciesClient) getHandleResponse(resp *http.Response) (DaprResiliencyCircuitBreakerPoliciesClientGetResponse, error) {
	result := DaprResiliencyCircuitBreakerPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPolicy); err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Dapr Resiliency Circuit Breaker Policies for a managed environment.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - DaprResiliencyCircuitBreakerPoliciesClientListOptions contains the optional parameters for the DaprResiliencyCircuitBreakerPoliciesClient.NewListPager
//     method.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) NewListPager(resourceGroupName string, environmentName string, options *DaprResiliencyCircuitBreakerPoliciesClientListOptions) *runtime.Pager[DaprResiliencyCircuitBreakerPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DaprResiliencyCircuitBreakerPoliciesClientListResponse]{
		More: func(page DaprResiliencyCircuitBreakerPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprResiliencyCircuitBreakerPoliciesClientListResponse) (DaprResiliencyCircuitBreakerPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DaprResiliencyCircuitBreakerPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DaprResiliencyCircuitBreakerPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *DaprResiliencyCircuitBreakerPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprResiliencyCircuitBreakerPolicies"
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
func (client *DaprResiliencyCircuitBreakerPoliciesClient) listHandleResponse(resp *http.Response) (DaprResiliencyCircuitBreakerPoliciesClientListResponse, error) {
	result := DaprResiliencyCircuitBreakerPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprResiliencyCircuitBreakerPoliciesCollection); err != nil {
		return DaprResiliencyCircuitBreakerPoliciesClientListResponse{}, err
	}
	return result, nil
}
