//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// VirtualMachineRunCommandsClient contains the methods for the VirtualMachineRunCommands group.
// Don't use this type directly, use NewVirtualMachineRunCommandsClient() instead.
type VirtualMachineRunCommandsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineRunCommandsClient creates a new instance of VirtualMachineRunCommandsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineRunCommandsClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualMachineRunCommandsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineRunCommandsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the run command should be created or updated.
//   - runCommandName - The name of the virtual machine run command.
//   - runCommand - Parameters supplied to the Create Virtual Machine RunCommand operation.
//   - options - VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualMachineRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineRunCommandsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineRunCommandsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineRunCommandsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *VirtualMachineRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, runCommand)
}

// BeginDelete - The operation to delete the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the run command should be deleted.
//   - runCommandName - The name of the virtual machine run command.
//   - options - VirtualMachineRunCommandsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginDelete
//     method.
func (client *VirtualMachineRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineRunCommandsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmName, runCommandName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineRunCommandsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineRunCommandsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *VirtualMachineRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmName, runCommandName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// Get - Gets specific run command for a subscription in a location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - location - The location upon which run commands is queried.
//   - commandID - The command id.
//   - options - VirtualMachineRunCommandsClientGetOptions contains the optional parameters for the VirtualMachineRunCommandsClient.Get
//     method.
func (client *VirtualMachineRunCommandsClient) Get(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsClientGetOptions) (VirtualMachineRunCommandsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, commandID, options)
	if err != nil {
		return VirtualMachineRunCommandsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineRunCommandsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineRunCommandsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineRunCommandsClient) getCreateRequest(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if commandID == "" {
		return nil, errors.New("parameter commandID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commandId}", url.PathEscape(commandID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineRunCommandsClient) getHandleResponse(resp *http.Response) (VirtualMachineRunCommandsClientGetResponse, error) {
	result := VirtualMachineRunCommandsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunCommandDocument); err != nil {
		return VirtualMachineRunCommandsClientGetResponse{}, err
	}
	return result, nil
}

// GetByVirtualMachine - The operation to get the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine containing the run command.
//   - runCommandName - The name of the virtual machine run command.
//   - options - VirtualMachineRunCommandsClientGetByVirtualMachineOptions contains the optional parameters for the VirtualMachineRunCommandsClient.GetByVirtualMachine
//     method.
func (client *VirtualMachineRunCommandsClient) GetByVirtualMachine(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsClientGetByVirtualMachineOptions) (VirtualMachineRunCommandsClientGetByVirtualMachineResponse, error) {
	req, err := client.getByVirtualMachineCreateRequest(ctx, resourceGroupName, vmName, runCommandName, options)
	if err != nil {
		return VirtualMachineRunCommandsClientGetByVirtualMachineResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineRunCommandsClientGetByVirtualMachineResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineRunCommandsClientGetByVirtualMachineResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByVirtualMachineHandleResponse(resp)
}

// getByVirtualMachineCreateRequest creates the GetByVirtualMachine request.
func (client *VirtualMachineRunCommandsClient) getByVirtualMachineCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsClientGetByVirtualMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getByVirtualMachineHandleResponse handles the GetByVirtualMachine response.
func (client *VirtualMachineRunCommandsClient) getByVirtualMachineHandleResponse(resp *http.Response) (VirtualMachineRunCommandsClientGetByVirtualMachineResponse, error) {
	result := VirtualMachineRunCommandsClientGetByVirtualMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommand); err != nil {
		return VirtualMachineRunCommandsClientGetByVirtualMachineResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all available run commands for a subscription in a location.
//
// Generated from API version 2023-07-01
//   - location - The location upon which run commands is queried.
//   - options - VirtualMachineRunCommandsClientListOptions contains the optional parameters for the VirtualMachineRunCommandsClient.NewListPager
//     method.
func (client *VirtualMachineRunCommandsClient) NewListPager(location string, options *VirtualMachineRunCommandsClientListOptions) *runtime.Pager[VirtualMachineRunCommandsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineRunCommandsClientListResponse]{
		More: func(page VirtualMachineRunCommandsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineRunCommandsClientListResponse) (VirtualMachineRunCommandsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineRunCommandsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualMachineRunCommandsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineRunCommandsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineRunCommandsClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineRunCommandsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineRunCommandsClient) listHandleResponse(resp *http.Response) (VirtualMachineRunCommandsClientListResponse, error) {
	result := VirtualMachineRunCommandsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunCommandListResult); err != nil {
		return VirtualMachineRunCommandsClientListResponse{}, err
	}
	return result, nil
}

// NewListByVirtualMachinePager - The operation to get all run commands of a Virtual Machine.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine containing the run command.
//   - options - VirtualMachineRunCommandsClientListByVirtualMachineOptions contains the optional parameters for the VirtualMachineRunCommandsClient.NewListByVirtualMachinePager
//     method.
func (client *VirtualMachineRunCommandsClient) NewListByVirtualMachinePager(resourceGroupName string, vmName string, options *VirtualMachineRunCommandsClientListByVirtualMachineOptions) *runtime.Pager[VirtualMachineRunCommandsClientListByVirtualMachineResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineRunCommandsClientListByVirtualMachineResponse]{
		More: func(page VirtualMachineRunCommandsClientListByVirtualMachineResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineRunCommandsClientListByVirtualMachineResponse) (VirtualMachineRunCommandsClientListByVirtualMachineResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVirtualMachineCreateRequest(ctx, resourceGroupName, vmName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineRunCommandsClientListByVirtualMachineResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualMachineRunCommandsClientListByVirtualMachineResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineRunCommandsClientListByVirtualMachineResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVirtualMachineHandleResponse(resp)
		},
	})
}

// listByVirtualMachineCreateRequest creates the ListByVirtualMachine request.
func (client *VirtualMachineRunCommandsClient) listByVirtualMachineCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineRunCommandsClientListByVirtualMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listByVirtualMachineHandleResponse handles the ListByVirtualMachine response.
func (client *VirtualMachineRunCommandsClient) listByVirtualMachineHandleResponse(resp *http.Response) (VirtualMachineRunCommandsClientListByVirtualMachineResponse, error) {
	result := VirtualMachineRunCommandsClientListByVirtualMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommandsListResult); err != nil {
		return VirtualMachineRunCommandsClientListByVirtualMachineResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the run command should be updated.
//   - runCommandName - The name of the virtual machine run command.
//   - runCommand - Parameters supplied to the Update Virtual Machine RunCommand operation.
//   - options - VirtualMachineRunCommandsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginUpdate
//     method.
func (client *VirtualMachineRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineRunCommandsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineRunCommandsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineRunCommandsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *VirtualMachineRunCommandsClient) update(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, runCommand)
}
