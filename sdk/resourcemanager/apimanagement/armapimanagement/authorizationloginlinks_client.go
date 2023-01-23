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
	"strings"
)

// AuthorizationLoginLinksClient contains the methods for the AuthorizationLoginLinks group.
// Don't use this type directly, use NewAuthorizationLoginLinksClient() instead.
type AuthorizationLoginLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAuthorizationLoginLinksClient creates a new instance of AuthorizationLoginLinksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAuthorizationLoginLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationLoginLinksClient, error) {
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
	client := &AuthorizationLoginLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Post - Gets authorization login links.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - authorizationProviderID - Identifier of the authorization provider.
//   - authorizationID - Identifier of the authorization.
//   - parameters - Create parameters.
//   - options - AuthorizationLoginLinksClientPostOptions contains the optional parameters for the AuthorizationLoginLinksClient.Post
//     method.
func (client *AuthorizationLoginLinksClient) Post(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, authorizationID string, parameters AuthorizationLoginRequestContract, options *AuthorizationLoginLinksClientPostOptions) (AuthorizationLoginLinksClientPostResponse, error) {
	req, err := client.postCreateRequest(ctx, resourceGroupName, serviceName, authorizationProviderID, authorizationID, parameters, options)
	if err != nil {
		return AuthorizationLoginLinksClientPostResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationLoginLinksClientPostResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationLoginLinksClientPostResponse{}, runtime.NewResponseError(resp)
	}
	return client.postHandleResponse(resp)
}

// postCreateRequest creates the Post request.
func (client *AuthorizationLoginLinksClient) postCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, authorizationProviderID string, authorizationID string, parameters AuthorizationLoginRequestContract, options *AuthorizationLoginLinksClientPostOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationProviders/{authorizationProviderId}/authorizations/{authorizationId}/getLoginLinks"
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
	if authorizationID == "" {
		return nil, errors.New("parameter authorizationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationId}", url.PathEscape(authorizationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// postHandleResponse handles the Post response.
func (client *AuthorizationLoginLinksClient) postHandleResponse(resp *http.Response) (AuthorizationLoginLinksClientPostResponse, error) {
	result := AuthorizationLoginLinksClientPostResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationLoginResponseContract); err != nil {
		return AuthorizationLoginLinksClientPostResponse{}, err
	}
	return result, nil
}
