//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// EurekaServersClient contains the methods for the EurekaServers group.
// Don't use this type directly, use NewEurekaServersClient() instead.
type EurekaServersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEurekaServersClient creates a new instance of EurekaServersClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEurekaServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EurekaServersClient, error) {
	cl, err := arm.NewClient(moduleName+".EurekaServersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EurekaServersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - EurekaServersClientGetOptions contains the optional parameters for the EurekaServersClient.Get method.
func (client *EurekaServersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *EurekaServersClientGetOptions) (EurekaServersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return EurekaServersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EurekaServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EurekaServersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EurekaServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EurekaServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/eurekaServers/default"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EurekaServersClient) getHandleResponse(resp *http.Response) (EurekaServersClientGetResponse, error) {
	result := EurekaServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EurekaServerResource); err != nil {
		return EurekaServersClientGetResponse{}, err
	}
	return result, nil
}

// List - List the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - EurekaServersClientListOptions contains the optional parameters for the EurekaServersClient.List method.
func (client *EurekaServersClient) List(ctx context.Context, resourceGroupName string, serviceName string, options *EurekaServersClientListOptions) (EurekaServersClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return EurekaServersClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EurekaServersClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EurekaServersClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *EurekaServersClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EurekaServersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/eurekaServers"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EurekaServersClient) listHandleResponse(resp *http.Response) (EurekaServersClientListResponse, error) {
	result := EurekaServersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EurekaServerResourceCollection); err != nil {
		return EurekaServersClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdatePatch - Update the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - eurekaServerResource - Parameters for the update operation
//   - options - EurekaServersClientBeginUpdatePatchOptions contains the optional parameters for the EurekaServersClient.BeginUpdatePatch
//     method.
func (client *EurekaServersClient) BeginUpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePatchOptions) (*runtime.Poller[EurekaServersClientUpdatePatchResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePatch(ctx, resourceGroupName, serviceName, eurekaServerResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EurekaServersClientUpdatePatchResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EurekaServersClientUpdatePatchResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdatePatch - Update the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *EurekaServersClient) updatePatch(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePatchOptions) (*http.Response, error) {
	req, err := client.updatePatchCreateRequest(ctx, resourceGroupName, serviceName, eurekaServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePatchCreateRequest creates the UpdatePatch request.
func (client *EurekaServersClient) updatePatchCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/eurekaServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, eurekaServerResource)
}

// BeginUpdatePut - Update the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - eurekaServerResource - Parameters for the update operation
//   - options - EurekaServersClientBeginUpdatePutOptions contains the optional parameters for the EurekaServersClient.BeginUpdatePut
//     method.
func (client *EurekaServersClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePutOptions) (*runtime.Poller[EurekaServersClientUpdatePutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePut(ctx, resourceGroupName, serviceName, eurekaServerResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EurekaServersClientUpdatePutResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EurekaServersClientUpdatePutResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdatePut - Update the eureka server settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *EurekaServersClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePutOptions) (*http.Response, error) {
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, eurekaServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePutCreateRequest creates the UpdatePut request.
func (client *EurekaServersClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource EurekaServerResource, options *EurekaServersClientBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/eurekaServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, eurekaServerResource)
}
