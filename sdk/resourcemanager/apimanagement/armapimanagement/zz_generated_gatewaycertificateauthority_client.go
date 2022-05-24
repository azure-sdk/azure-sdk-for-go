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
	"strconv"
	"strings"
)

// GatewayCertificateAuthorityClient contains the methods for the GatewayCertificateAuthority group.
// Don't use this type directly, use NewGatewayCertificateAuthorityClient() instead.
type GatewayCertificateAuthorityClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGatewayCertificateAuthorityClient creates a new instance of GatewayCertificateAuthorityClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGatewayCertificateAuthorityClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GatewayCertificateAuthorityClient, error) {
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
	client := &GatewayCertificateAuthorityClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Assign Certificate entity to Gateway entity as Certificate Authority.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
// 'managed'
// certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
// options - GatewayCertificateAuthorityClientCreateOrUpdateOptions contains the optional parameters for the GatewayCertificateAuthorityClient.CreateOrUpdate
// method.
func (client *GatewayCertificateAuthorityClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, parameters GatewayCertificateAuthorityContract, options *GatewayCertificateAuthorityClientCreateOrUpdateOptions) (GatewayCertificateAuthorityClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, parameters, options)
	if err != nil {
		return GatewayCertificateAuthorityClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GatewayCertificateAuthorityClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewayCertificateAuthorityClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, parameters GatewayCertificateAuthorityContract, options *GatewayCertificateAuthorityClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) createOrUpdateHandleResponse(resp *http.Response) (GatewayCertificateAuthorityClientCreateOrUpdateResponse, error) {
	result := GatewayCertificateAuthorityClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityContract); err != nil {
		return GatewayCertificateAuthorityClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove relationship between Certificate Authority and Gateway entity.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
// 'managed'
// certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - GatewayCertificateAuthorityClientDeleteOptions contains the optional parameters for the GatewayCertificateAuthorityClient.Delete
// method.
func (client *GatewayCertificateAuthorityClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, ifMatch string, options *GatewayCertificateAuthorityClientDeleteOptions) (GatewayCertificateAuthorityClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, ifMatch, options)
	if err != nil {
		return GatewayCertificateAuthorityClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GatewayCertificateAuthorityClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return GatewayCertificateAuthorityClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayCertificateAuthorityClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, ifMatch string, options *GatewayCertificateAuthorityClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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

// Get - Get assigned Gateway Certificate Authority details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
// 'managed'
// certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
// options - GatewayCertificateAuthorityClientGetOptions contains the optional parameters for the GatewayCertificateAuthorityClient.Get
// method.
func (client *GatewayCertificateAuthorityClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityClientGetOptions) (GatewayCertificateAuthorityClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, options)
	if err != nil {
		return GatewayCertificateAuthorityClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayCertificateAuthorityClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayCertificateAuthorityClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) getHandleResponse(resp *http.Response) (GatewayCertificateAuthorityClientGetResponse, error) {
	result := GatewayCertificateAuthorityClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityContract); err != nil {
		return GatewayCertificateAuthorityClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Checks if Certificate entity is assigned to Gateway entity as Certificate Authority.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
// 'managed'
// certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
// options - GatewayCertificateAuthorityClientGetEntityTagOptions contains the optional parameters for the GatewayCertificateAuthorityClient.GetEntityTag
// method.
func (client *GatewayCertificateAuthorityClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityClientGetEntityTagOptions) (GatewayCertificateAuthorityClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, certificateID, options)
	if err != nil {
		return GatewayCertificateAuthorityClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayCertificateAuthorityClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayCertificateAuthorityClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GatewayCertificateAuthorityClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, certificateID string, options *GatewayCertificateAuthorityClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
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
func (client *GatewayCertificateAuthorityClient) getEntityTagHandleResponse(resp *http.Response) (GatewayCertificateAuthorityClientGetEntityTagResponse, error) {
	result := GatewayCertificateAuthorityClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists the collection of Certificate Authorities for the specified Gateway entity.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
// 'managed'
// options - GatewayCertificateAuthorityClientListByServiceOptions contains the optional parameters for the GatewayCertificateAuthorityClient.ListByService
// method.
func (client *GatewayCertificateAuthorityClient) NewListByServicePager(resourceGroupName string, serviceName string, gatewayID string, options *GatewayCertificateAuthorityClientListByServiceOptions) *runtime.Pager[GatewayCertificateAuthorityClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewayCertificateAuthorityClientListByServiceResponse]{
		More: func(page GatewayCertificateAuthorityClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewayCertificateAuthorityClientListByServiceResponse) (GatewayCertificateAuthorityClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GatewayCertificateAuthorityClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GatewayCertificateAuthorityClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GatewayCertificateAuthorityClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *GatewayCertificateAuthorityClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, options *GatewayCertificateAuthorityClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/certificateAuthorities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
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
func (client *GatewayCertificateAuthorityClient) listByServiceHandleResponse(resp *http.Response) (GatewayCertificateAuthorityClientListByServiceResponse, error) {
	result := GatewayCertificateAuthorityClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayCertificateAuthorityCollection); err != nil {
		return GatewayCertificateAuthorityClientListByServiceResponse{}, err
	}
	return result, nil
}
