//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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
	"strconv"
	"strings"
)

// AuthorizationProviderClient contains the methods for the AuthorizationProvider group.
// Don't use this type directly, use NewAuthorizationProviderClient() instead.
type AuthorizationProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAuthorizationProviderClient creates a new instance of AuthorizationProviderClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAuthorizationProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationProviderClient, error) {
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
	client := &AuthorizationProviderClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates authorization provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// authorizationProviderID - Identifier of the authorization provider.
// parameters - Create parameters.
// options - AuthorizationProviderClientCreateOrUpdateOptions contains the optional parameters for the AuthorizationProviderClient.CreateOrUpdate
// method.
func (client *AuthorizationProviderClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, parameters AuthorizationProviderContract, options *AuthorizationProviderClientCreateOrUpdateOptions) (AuthorizationProviderClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, authorizationProviderID, parameters, options)
	if err != nil {
		return AuthorizationProviderClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationProviderClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationProviderClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AuthorizationProviderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, parameters AuthorizationProviderContract, options *AuthorizationProviderClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authorizationProviderID == "" {
		return nil, errors.New("parameter authorizationProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationProviderId}", url.PathEscape(authorizationProviderID))
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
func (client *AuthorizationProviderClient) createOrUpdateHandleResponse(resp *http.Response) (AuthorizationProviderClientCreateOrUpdateResponse, error) {
	result := AuthorizationProviderClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationProviderContract); err != nil {
		return AuthorizationProviderClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific authorization provider from the API Management service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// authorizationProviderID - Identifier of the authorization provider.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - AuthorizationProviderClientDeleteOptions contains the optional parameters for the AuthorizationProviderClient.Delete
// method.
func (client *AuthorizationProviderClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, ifMatch string, options *AuthorizationProviderClientDeleteOptions) (AuthorizationProviderClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, authorizationProviderID, ifMatch, options)
	if err != nil {
		return AuthorizationProviderClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationProviderClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AuthorizationProviderClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AuthorizationProviderClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AuthorizationProviderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, ifMatch string, options *AuthorizationProviderClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authorizationProviderID == "" {
		return nil, errors.New("parameter authorizationProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationProviderId}", url.PathEscape(authorizationProviderID))
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

// Get - Gets the details of the authorization provider specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// authorizationProviderID - Identifier of the authorization provider.
// options - AuthorizationProviderClientGetOptions contains the optional parameters for the AuthorizationProviderClient.Get
// method.
func (client *AuthorizationProviderClient) Get(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, options *AuthorizationProviderClientGetOptions) (AuthorizationProviderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, authorizationProviderID, options)
	if err != nil {
		return AuthorizationProviderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationProviderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationProviderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AuthorizationProviderClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, options *AuthorizationProviderClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if authorizationProviderID == "" {
		return nil, errors.New("parameter authorizationProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationProviderId}", url.PathEscape(authorizationProviderID))
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

// getHandleResponse handles the Get response.
func (client *AuthorizationProviderClient) getHandleResponse(resp *http.Response) (AuthorizationProviderClientGetResponse, error) {
	result := AuthorizationProviderClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationProviderContract); err != nil {
		return AuthorizationProviderClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists a collection of authorization providers defined within a service instance.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - AuthorizationProviderClientListByServiceOptions contains the optional parameters for the AuthorizationProviderClient.ListByService
// method.
func (client *AuthorizationProviderClient) NewListByServicePager(resourceGroupName string, serviceName string, options *AuthorizationProviderClientListByServiceOptions) *runtime.Pager[AuthorizationProviderClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthorizationProviderClientListByServiceResponse]{
		More: func(page AuthorizationProviderClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationProviderClientListByServiceResponse) (AuthorizationProviderClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AuthorizationProviderClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AuthorizationProviderClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AuthorizationProviderClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *AuthorizationProviderClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *AuthorizationProviderClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *AuthorizationProviderClient) listByServiceHandleResponse(resp *http.Response) (AuthorizationProviderClientListByServiceResponse, error) {
	result := AuthorizationProviderClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationProviderCollection); err != nil {
		return AuthorizationProviderClientListByServiceResponse{}, err
	}
	return result, nil
}
