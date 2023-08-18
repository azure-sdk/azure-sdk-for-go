//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// PrivateEndpointConnectionClient contains the methods for the PrivateEndpointConnection group.
// Don't use this type directly, use NewPrivateEndpointConnectionClient() instead.
type PrivateEndpointConnectionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionClient creates a new instance of PrivateEndpointConnectionClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateEndpointConnectionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new Private Endpoint Connection or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - privateEndpointConnectionName - Name of the private endpoint connection.
//   - options - PrivateEndpointConnectionClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginCreateOrUpdate
//     method.
func (client *PrivateEndpointConnectionClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, privateEndpointConnectionRequest PrivateEndpointConnectionRequest, options *PrivateEndpointConnectionClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, privateEndpointConnectionName, privateEndpointConnectionRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PrivateEndpointConnectionClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new Private Endpoint Connection or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *PrivateEndpointConnectionClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, privateEndpointConnectionRequest PrivateEndpointConnectionRequest, options *PrivateEndpointConnectionClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, privateEndpointConnectionName, privateEndpointConnectionRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, privateEndpointConnectionRequest PrivateEndpointConnectionRequest, options *PrivateEndpointConnectionClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, privateEndpointConnectionRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified Private Endpoint Connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - privateEndpointConnectionName - Name of the private endpoint connection.
//   - options - PrivateEndpointConnectionClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginDelete
//     method.
func (client *PrivateEndpointConnectionClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PrivateEndpointConnectionClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified Private Endpoint Connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *PrivateEndpointConnectionClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetByName - Gets the details of the Private Endpoint Connection specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - privateEndpointConnectionName - Name of the private endpoint connection.
//   - options - PrivateEndpointConnectionClientGetByNameOptions contains the optional parameters for the PrivateEndpointConnectionClient.GetByName
//     method.
func (client *PrivateEndpointConnectionClient) GetByName(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientGetByNameOptions) (PrivateEndpointConnectionClientGetByNameResponse, error) {
	var err error
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, serviceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionClientGetByNameResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionClientGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionClientGetByNameResponse{}, err
	}
	resp, err := client.getByNameHandleResponse(httpResp)
	return resp, err
}

// getByNameCreateRequest creates the GetByName request.
func (client *PrivateEndpointConnectionClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionClientGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *PrivateEndpointConnectionClient) getByNameHandleResponse(resp *http.Response) (PrivateEndpointConnectionClientGetByNameResponse, error) {
	result := PrivateEndpointConnectionClientGetByNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionClientGetByNameResponse{}, err
	}
	return result, nil
}

// GetPrivateLinkResource - Gets the private link resources
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - privateLinkSubResourceName - Name of the private link resource.
//   - options - PrivateEndpointConnectionClientGetPrivateLinkResourceOptions contains the optional parameters for the PrivateEndpointConnectionClient.GetPrivateLinkResource
//     method.
func (client *PrivateEndpointConnectionClient) GetPrivateLinkResource(ctx context.Context, resourceGroupName string, serviceName string, privateLinkSubResourceName string, options *PrivateEndpointConnectionClientGetPrivateLinkResourceOptions) (PrivateEndpointConnectionClientGetPrivateLinkResourceResponse, error) {
	var err error
	req, err := client.getPrivateLinkResourceCreateRequest(ctx, resourceGroupName, serviceName, privateLinkSubResourceName, options)
	if err != nil {
		return PrivateEndpointConnectionClientGetPrivateLinkResourceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionClientGetPrivateLinkResourceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionClientGetPrivateLinkResourceResponse{}, err
	}
	resp, err := client.getPrivateLinkResourceHandleResponse(httpResp)
	return resp, err
}

// getPrivateLinkResourceCreateRequest creates the GetPrivateLinkResource request.
func (client *PrivateEndpointConnectionClient) getPrivateLinkResourceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, privateLinkSubResourceName string, options *PrivateEndpointConnectionClientGetPrivateLinkResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateLinkResources/{privateLinkSubResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if privateLinkSubResourceName == "" {
		return nil, errors.New("parameter privateLinkSubResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkSubResourceName}", url.PathEscape(privateLinkSubResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPrivateLinkResourceHandleResponse handles the GetPrivateLinkResource response.
func (client *PrivateEndpointConnectionClient) getPrivateLinkResourceHandleResponse(resp *http.Response) (PrivateEndpointConnectionClientGetPrivateLinkResourceResponse, error) {
	result := PrivateEndpointConnectionClientGetPrivateLinkResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateEndpointConnectionClientGetPrivateLinkResourceResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists all private endpoint connections of the API Management service instance.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PrivateEndpointConnectionClientListByServiceOptions contains the optional parameters for the PrivateEndpointConnectionClient.NewListByServicePager
//     method.
func (client *PrivateEndpointConnectionClient) NewListByServicePager(resourceGroupName string, serviceName string, options *PrivateEndpointConnectionClientListByServiceOptions) *runtime.Pager[PrivateEndpointConnectionClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionClientListByServiceResponse]{
		More: func(page PrivateEndpointConnectionClientListByServiceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionClientListByServiceResponse) (PrivateEndpointConnectionClientListByServiceResponse, error) {
			req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			if err != nil {
				return PrivateEndpointConnectionClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateEndpointConnectionClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PrivateEndpointConnectionClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PrivateEndpointConnectionClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateEndpointConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PrivateEndpointConnectionClient) listByServiceHandleResponse(resp *http.Response) (PrivateEndpointConnectionClientListByServiceResponse, error) {
	result := PrivateEndpointConnectionClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListPrivateLinkResources - Gets the private link resources
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PrivateEndpointConnectionClientListPrivateLinkResourcesOptions contains the optional parameters for the PrivateEndpointConnectionClient.ListPrivateLinkResources
//     method.
func (client *PrivateEndpointConnectionClient) ListPrivateLinkResources(ctx context.Context, resourceGroupName string, serviceName string, options *PrivateEndpointConnectionClientListPrivateLinkResourcesOptions) (PrivateEndpointConnectionClientListPrivateLinkResourcesResponse, error) {
	var err error
	req, err := client.listPrivateLinkResourcesCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return PrivateEndpointConnectionClientListPrivateLinkResourcesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionClientListPrivateLinkResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionClientListPrivateLinkResourcesResponse{}, err
	}
	resp, err := client.listPrivateLinkResourcesHandleResponse(httpResp)
	return resp, err
}

// listPrivateLinkResourcesCreateRequest creates the ListPrivateLinkResources request.
func (client *PrivateEndpointConnectionClient) listPrivateLinkResourcesCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PrivateEndpointConnectionClientListPrivateLinkResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPrivateLinkResourcesHandleResponse handles the ListPrivateLinkResources response.
func (client *PrivateEndpointConnectionClient) listPrivateLinkResourcesHandleResponse(resp *http.Response) (PrivateEndpointConnectionClientListPrivateLinkResourcesResponse, error) {
	result := PrivateEndpointConnectionClientListPrivateLinkResourcesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateEndpointConnectionClientListPrivateLinkResourcesResponse{}, err
	}
	return result, nil
}
