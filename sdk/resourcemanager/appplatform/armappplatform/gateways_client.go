//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// GatewaysClient contains the methods for the Gateways group.
// Don't use this type directly, use NewGatewaysClient() instead.
type GatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGatewaysClient creates a new instance of GatewaysClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GatewaysClient, error) {
	cl, err := arm.NewClient(moduleName+".GatewaysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - gatewayResource - The gateway for the create or update operation
//   - options - GatewaysClientBeginCreateOrUpdateOptions contains the optional parameters for the GatewaysClient.BeginCreateOrUpdate
//     method.
func (client *GatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (*runtime.Poller[GatewaysClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, gatewayName, gatewayResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GatewaysClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GatewaysClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *GatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, gatewayResource, options)
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
func (client *GatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource GatewayResource, options *GatewaysClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, gatewayResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Disable the default Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - options - GatewaysClientBeginDeleteOptions contains the optional parameters for the GatewaysClient.BeginDelete method.
func (client *GatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*runtime.Poller[GatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GatewaysClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GatewaysClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Disable the default Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *GatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
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
func (client *GatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Spring Cloud Gateway and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - options - GatewaysClientGetOptions contains the optional parameters for the GatewaysClient.Get method.
func (client *GatewaysClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientGetOptions) (GatewaysClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GatewaysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewaysClient) getHandleResponse(resp *http.Response) (GatewaysClientGetResponse, error) {
	result := GatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResource); err != nil {
		return GatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - GatewaysClientListOptions contains the optional parameters for the GatewaysClient.NewListPager method.
func (client *GatewaysClient) NewListPager(resourceGroupName string, serviceName string, options *GatewaysClientListOptions) *runtime.Pager[GatewaysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewaysClientListResponse]{
		More: func(page GatewaysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewaysClientListResponse) (GatewaysClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GatewaysClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GatewaysClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GatewaysClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *GatewaysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GatewaysClient) listHandleResponse(resp *http.Response) (GatewaysClientListResponse, error) {
	result := GatewaysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResourceCollection); err != nil {
		return GatewaysClientListResponse{}, err
	}
	return result, nil
}

// ListEnvSecrets - List sensitive environment variables of Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - options - GatewaysClientListEnvSecretsOptions contains the optional parameters for the GatewaysClient.ListEnvSecrets method.
func (client *GatewaysClient) ListEnvSecrets(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientListEnvSecretsOptions) (GatewaysClientListEnvSecretsResponse, error) {
	var err error
	req, err := client.listEnvSecretsCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
	if err != nil {
		return GatewaysClientListEnvSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewaysClientListEnvSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GatewaysClientListEnvSecretsResponse{}, err
	}
	resp, err := client.listEnvSecretsHandleResponse(httpResp)
	return resp, err
}

// listEnvSecretsCreateRequest creates the ListEnvSecrets request.
func (client *GatewaysClient) listEnvSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientListEnvSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/listEnvSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEnvSecretsHandleResponse handles the ListEnvSecrets response.
func (client *GatewaysClient) listEnvSecretsHandleResponse(resp *http.Response) (GatewaysClientListEnvSecretsResponse, error) {
	result := GatewaysClientListEnvSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return GatewaysClientListEnvSecretsResponse{}, err
	}
	return result, nil
}

// BeginRestart - Restart the Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - options - GatewaysClientBeginRestartOptions contains the optional parameters for the GatewaysClient.BeginRestart method.
func (client *GatewaysClient) BeginRestart(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginRestartOptions) (*runtime.Poller[GatewaysClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, serviceName, gatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GatewaysClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GatewaysClientRestartResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Restart - Restart the Spring Cloud Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *GatewaysClient) restart(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginRestartOptions) (*http.Response, error) {
	var err error
	req, err := client.restartCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, options)
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

// restartCreateRequest creates the Restart request.
func (client *GatewaysClient) restartCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, options *GatewaysClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdateCapacity - Operation to update an exiting Spring Cloud Gateway capacity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - gatewayCapacityResource - The gateway capacity for the update operation
//   - options - GatewaysClientBeginUpdateCapacityOptions contains the optional parameters for the GatewaysClient.BeginUpdateCapacity
//     method.
func (client *GatewaysClient) BeginUpdateCapacity(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayCapacityResource SKUObject, options *GatewaysClientBeginUpdateCapacityOptions) (*runtime.Poller[GatewaysClientUpdateCapacityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateCapacity(ctx, resourceGroupName, serviceName, gatewayName, gatewayCapacityResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GatewaysClientUpdateCapacityResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GatewaysClientUpdateCapacityResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateCapacity - Operation to update an exiting Spring Cloud Gateway capacity.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *GatewaysClient) updateCapacity(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayCapacityResource SKUObject, options *GatewaysClientBeginUpdateCapacityOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCapacityCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, gatewayCapacityResource, options)
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

// updateCapacityCreateRequest creates the UpdateCapacity request.
func (client *GatewaysClient) updateCapacityCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayCapacityResource SKUObject, options *GatewaysClientBeginUpdateCapacityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, gatewayCapacityResource); err != nil {
		return nil, err
	}
	return req, nil
}

// ValidateDomain - Check the domains are valid as well as not in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - gatewayName - The name of Spring Cloud Gateway.
//   - validatePayload - Custom domain payload to be validated
//   - options - GatewaysClientValidateDomainOptions contains the optional parameters for the GatewaysClient.ValidateDomain method.
func (client *GatewaysClient) ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, validatePayload CustomDomainValidatePayload, options *GatewaysClientValidateDomainOptions) (GatewaysClientValidateDomainResponse, error) {
	var err error
	req, err := client.validateDomainCreateRequest(ctx, resourceGroupName, serviceName, gatewayName, validatePayload, options)
	if err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GatewaysClientValidateDomainResponse{}, err
	}
	resp, err := client.validateDomainHandleResponse(httpResp)
	return resp, err
}

// validateDomainCreateRequest creates the ValidateDomain request.
func (client *GatewaysClient) validateDomainCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, validatePayload CustomDomainValidatePayload, options *GatewaysClientValidateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/validateDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, validatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// validateDomainHandleResponse handles the ValidateDomain response.
func (client *GatewaysClient) validateDomainHandleResponse(resp *http.Response) (GatewaysClientValidateDomainResponse, error) {
	result := GatewaysClientValidateDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainValidateResult); err != nil {
		return GatewaysClientValidateDomainResponse{}, err
	}
	return result, nil
}
