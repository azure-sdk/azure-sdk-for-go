//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VipSwapClient contains the methods for the VipSwap group.
// Don't use this type directly, use NewVipSwapClient() instead.
type VipSwapClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVipSwapClient creates a new instance of VipSwapClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVipSwapClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VipSwapClient, error) {
	cl, err := arm.NewClient(moduleName+".VipSwapClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VipSwapClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Performs vip swap operation on swappable cloud services.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - groupName - The name of the resource group.
//   - resourceName - The name of the cloud service.
//   - parameters - SwapResource object where slot type should be the target slot after vip swap for the specified cloud service.
//   - options - VipSwapClientBeginCreateOptions contains the optional parameters for the VipSwapClient.BeginCreate method.
func (client *VipSwapClient) BeginCreate(ctx context.Context, groupName string, resourceName string, parameters SwapResource, options *VipSwapClientBeginCreateOptions) (*runtime.Poller[VipSwapClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, groupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[VipSwapClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VipSwapClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Performs vip swap operation on swappable cloud services.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *VipSwapClient) create(ctx context.Context, groupName string, resourceName string, parameters SwapResource, options *VipSwapClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, groupName, resourceName, parameters, options)
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

// createCreateRequest creates the Create request.
func (client *VipSwapClient) createCreateRequest(ctx context.Context, groupName string, resourceName string, parameters SwapResource, options *VipSwapClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Compute/cloudServices/{resourceName}/providers/Microsoft.Network/cloudServiceSlots/{singletonResource}"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	urlPath = strings.ReplaceAll(urlPath, "{singletonResource}", url.PathEscape("swap"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets the SwapResource which identifies the slot type for the specified cloud service. The slot type on a cloud service
// can either be Staging or Production
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - groupName - The name of the resource group.
//   - resourceName - The name of the cloud service.
//   - options - VipSwapClientGetOptions contains the optional parameters for the VipSwapClient.Get method.
func (client *VipSwapClient) Get(ctx context.Context, groupName string, resourceName string, options *VipSwapClientGetOptions) (VipSwapClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, groupName, resourceName, options)
	if err != nil {
		return VipSwapClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VipSwapClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VipSwapClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VipSwapClient) getCreateRequest(ctx context.Context, groupName string, resourceName string, options *VipSwapClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Compute/cloudServices/{resourceName}/providers/Microsoft.Network/cloudServiceSlots/{singletonResource}"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	urlPath = strings.ReplaceAll(urlPath, "{singletonResource}", url.PathEscape("swap"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VipSwapClient) getHandleResponse(resp *http.Response) (VipSwapClientGetResponse, error) {
	result := VipSwapClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SwapResource); err != nil {
		return VipSwapClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of SwapResource which identifies the slot type for the specified cloud service. The slot type on a
// cloud service can either be Staging or Production
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - groupName - The name of the resource group.
//   - resourceName - The name of the cloud service.
//   - options - VipSwapClientListOptions contains the optional parameters for the VipSwapClient.List method.
func (client *VipSwapClient) List(ctx context.Context, groupName string, resourceName string, options *VipSwapClientListOptions) (VipSwapClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, groupName, resourceName, options)
	if err != nil {
		return VipSwapClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VipSwapClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VipSwapClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *VipSwapClient) listCreateRequest(ctx context.Context, groupName string, resourceName string, options *VipSwapClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Compute/cloudServices/{resourceName}/providers/Microsoft.Network/cloudServiceSlots"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VipSwapClient) listHandleResponse(resp *http.Response) (VipSwapClientListResponse, error) {
	result := VipSwapClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SwapResourceListResult); err != nil {
		return VipSwapClientListResponse{}, err
	}
	return result, nil
}
