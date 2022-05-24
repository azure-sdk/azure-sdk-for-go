//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// PolicyClient contains the methods for the Policy group.
// Don't use this type directly, use NewPolicyClient() instead.
type PolicyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPolicyClient creates a new instance of PolicyClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyClient, error) {
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
	client := &PolicyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the global policy configuration of the Api Management service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// policyID - The identifier of the Policy.
// parameters - The policy contents to apply.
// options - PolicyClientCreateOrUpdateOptions contains the optional parameters for the PolicyClient.CreateOrUpdate method.
func (client *PolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, parameters PolicyContract, options *PolicyClientCreateOrUpdateOptions) (PolicyClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, policyID, parameters, options)
	if err != nil {
		return PolicyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PolicyClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, parameters PolicyContract, options *PolicyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PolicyClient) createOrUpdateHandleResponse(resp *http.Response) (PolicyClientCreateOrUpdateResponse, error) {
	result := PolicyClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return PolicyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the global policy configuration of the Api Management Service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// policyID - The identifier of the Policy.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - PolicyClientDeleteOptions contains the optional parameters for the PolicyClient.Delete method.
func (client *PolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, ifMatch string, options *PolicyClientDeleteOptions) (PolicyClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, policyID, ifMatch, options)
	if err != nil {
		return PolicyClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PolicyClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PolicyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, ifMatch string, options *PolicyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Global policy definition of the Api Management service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// policyID - The identifier of the Policy.
// options - PolicyClientGetOptions contains the optional parameters for the PolicyClient.Get method.
func (client *PolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, options *PolicyClientGetOptions) (PolicyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, policyID, options)
	if err != nil {
		return PolicyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, options *PolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("format", string(*options.Format))
	}
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PolicyClient) getHandleResponse(resp *http.Response) (PolicyClientGetResponse, error) {
	result := PolicyClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return PolicyClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Global policy definition in the Api Management service.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// policyID - The identifier of the Policy.
// options - PolicyClientGetEntityTagOptions contains the optional parameters for the PolicyClient.GetEntityTag method.
func (client *PolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, options *PolicyClientGetEntityTagOptions) (PolicyClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, policyID, options)
	if err != nil {
		return PolicyClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *PolicyClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, policyID PolicyIDName, options *PolicyClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *PolicyClient) getEntityTagHandleResponse(resp *http.Response) (PolicyClientGetEntityTagResponse, error) {
	result := PolicyClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// ListByService - Lists all the Global Policy definitions of the Api Management service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - PolicyClientListByServiceOptions contains the optional parameters for the PolicyClient.ListByService method.
func (client *PolicyClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *PolicyClientListByServiceOptions) (PolicyClientListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return PolicyClientListByServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyClientListByServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PolicyClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PolicyClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PolicyClient) listByServiceHandleResponse(resp *http.Response) (PolicyClientListByServiceResponse, error) {
	result := PolicyClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyCollection); err != nil {
		return PolicyClientListByServiceResponse{}, err
	}
	return result, nil
}
