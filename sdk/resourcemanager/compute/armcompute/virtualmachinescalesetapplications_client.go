//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// VirtualMachineScaleSetApplicationsClient contains the methods for the VirtualMachineScaleSetApplications group.
// Don't use this type directly, use NewVirtualMachineScaleSetApplicationsClient() instead.
type VirtualMachineScaleSetApplicationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineScaleSetApplicationsClient creates a new instance of VirtualMachineScaleSetApplicationsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineScaleSetApplicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineScaleSetApplicationsClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualMachineScaleSetApplicationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineScaleSetApplicationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Removes an application from a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - vmScaleSetName - The name of the VM scale set containing the extension.
//   - applicationName - The name of the application to query.
//   - options - VirtualMachineScaleSetApplicationsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineScaleSetApplicationsClient.BeginDelete
//     method.
func (client *VirtualMachineScaleSetApplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *VirtualMachineScaleSetApplicationsClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineScaleSetApplicationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, applicationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachineScaleSetApplicationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineScaleSetApplicationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Removes an application from a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *VirtualMachineScaleSetApplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *VirtualMachineScaleSetApplicationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineScaleSetApplicationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, applicationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineScaleSetApplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *VirtualMachineScaleSetApplicationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an application in a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - vmScaleSetName - The name of the VM scale set containing the extension.
//   - applicationName - The name of the application to query.
//   - options - VirtualMachineScaleSetApplicationsClientGetOptions contains the optional parameters for the VirtualMachineScaleSetApplicationsClient.Get
//     method.
func (client *VirtualMachineScaleSetApplicationsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *VirtualMachineScaleSetApplicationsClientGetOptions) (VirtualMachineScaleSetApplicationsClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachineScaleSetApplicationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, applicationName, options)
	if err != nil {
		return VirtualMachineScaleSetApplicationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetApplicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineScaleSetApplicationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetApplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *VirtualMachineScaleSetApplicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetApplicationsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetApplicationsClientGetResponse, error) {
	result := VirtualMachineScaleSetApplicationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VMApplicationProxyResource); err != nil {
		return VirtualMachineScaleSetApplicationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all applications in a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - vmScaleSetName - The name of the VM scale set containing the extension.
//   - options - VirtualMachineScaleSetApplicationsClientListOptions contains the optional parameters for the VirtualMachineScaleSetApplicationsClient.List
//     method.
func (client *VirtualMachineScaleSetApplicationsClient) List(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetApplicationsClientListOptions) (VirtualMachineScaleSetApplicationsClientListResponse, error) {
	var err error
	const operationName = "VirtualMachineScaleSetApplicationsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return VirtualMachineScaleSetApplicationsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineScaleSetApplicationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachineScaleSetApplicationsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetApplicationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetApplicationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/applications"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetApplicationsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetApplicationsClientListResponse, error) {
	result := VirtualMachineScaleSetApplicationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineApplicationsProxyResourceListResult); err != nil {
		return VirtualMachineScaleSetApplicationsClientListResponse{}, err
	}
	return result, nil
}

// BeginPut - Adds an application to a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group.
//   - vmScaleSetName - The name of the VM scale set containing the extension.
//   - applicationName - The name of the application to query.
//   - application - The definition of the VMApplication to add the virtual machine
//   - options - VirtualMachineScaleSetApplicationsClientBeginPutOptions contains the optional parameters for the VirtualMachineScaleSetApplicationsClient.BeginPut
//     method.
func (client *VirtualMachineScaleSetApplicationsClient) BeginPut(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, application VMApplicationProxyResource, options *VirtualMachineScaleSetApplicationsClientBeginPutOptions) (*runtime.Poller[VirtualMachineScaleSetApplicationsClientPutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.put(ctx, resourceGroupName, vmScaleSetName, applicationName, application, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[VirtualMachineScaleSetApplicationsClientPutResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineScaleSetApplicationsClientPutResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Put - Adds an application to a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *VirtualMachineScaleSetApplicationsClient) put(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, application VMApplicationProxyResource, options *VirtualMachineScaleSetApplicationsClientBeginPutOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachineScaleSetApplicationsClient.BeginPut"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, resourceGroupName, vmScaleSetName, applicationName, application, options)
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

// putCreateRequest creates the Put request.
func (client *VirtualMachineScaleSetApplicationsClient) putCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, application VMApplicationProxyResource, options *VirtualMachineScaleSetApplicationsClientBeginPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/applications/{applicationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, application); err != nil {
		return nil, err
	}
	return req, nil
}
