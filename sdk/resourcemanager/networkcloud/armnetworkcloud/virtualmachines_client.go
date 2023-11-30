//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// VirtualMachinesClient contains the methods for the VirtualMachines group.
// Don't use this type directly, use NewVirtualMachinesClient() instead.
type VirtualMachinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachinesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new virtual machine or update the properties of the existing virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - virtualMachineParameters - The request body.
//   - options - VirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualMachinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineParameters VirtualMachine, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachinesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualMachineName, virtualMachineParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new virtual machine or update the properties of the existing virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineParameters VirtualMachine, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualMachineName, virtualMachineParameters, options)
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
func (client *VirtualMachinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineParameters VirtualMachine, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, virtualMachineParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientBeginDeleteOptions contains the optional parameters for the VirtualMachinesClient.BeginDelete
//     method.
func (client *VirtualMachinesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
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
func (client *VirtualMachinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientGetOptions contains the optional parameters for the VirtualMachinesClient.Get method.
func (client *VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (VirtualMachinesClientGetResponse, error) {
	var err error
	const operationName = "VirtualMachinesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachinesClient) getHandleResponse(resp *http.Response) (VirtualMachinesClientGetResponse, error) {
	result := VirtualMachinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachine); err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of virtual machines in the provided resource group.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - VirtualMachinesClientListByResourceGroupOptions contains the optional parameters for the VirtualMachinesClient.NewListByResourceGroupPager
//     method.
func (client *VirtualMachinesClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualMachinesClientListByResourceGroupOptions) *runtime.Pager[VirtualMachinesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachinesClientListByResourceGroupResponse]{
		More: func(page VirtualMachinesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachinesClientListByResourceGroupResponse) (VirtualMachinesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualMachinesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VirtualMachinesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualMachinesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualMachinesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualMachinesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualMachinesClientListByResourceGroupResponse, error) {
	result := VirtualMachinesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineList); err != nil {
		return VirtualMachinesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of virtual machines in the provided subscription.
//
// Generated from API version 2023-10-01-preview
//   - options - VirtualMachinesClientListBySubscriptionOptions contains the optional parameters for the VirtualMachinesClient.NewListBySubscriptionPager
//     method.
func (client *VirtualMachinesClient) NewListBySubscriptionPager(options *VirtualMachinesClientListBySubscriptionOptions) *runtime.Pager[VirtualMachinesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachinesClientListBySubscriptionResponse]{
		More: func(page VirtualMachinesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachinesClientListBySubscriptionResponse) (VirtualMachinesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualMachinesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VirtualMachinesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VirtualMachinesClient) listBySubscriptionCreateRequest(ctx context.Context, options *VirtualMachinesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VirtualMachinesClient) listBySubscriptionHandleResponse(resp *http.Response) (VirtualMachinesClientListBySubscriptionResponse, error) {
	result := VirtualMachinesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineList); err != nil {
		return VirtualMachinesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginPowerOff - Power off the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientBeginPowerOffOptions contains the optional parameters for the VirtualMachinesClient.BeginPowerOff
//     method.
func (client *VirtualMachinesClient) BeginPowerOff(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginPowerOffOptions) (*runtime.Poller[VirtualMachinesClientPowerOffResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.powerOff(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientPowerOffResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientPowerOffResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// PowerOff - Power off the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) powerOff(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginPowerOffOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginPowerOff"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.powerOffCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
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

// powerOffCreateRequest creates the PowerOff request.
func (client *VirtualMachinesClient) powerOffCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginPowerOffOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/powerOff"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.VirtualMachinePowerOffParameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.VirtualMachinePowerOffParameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginReimage - Reimage the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientBeginReimageOptions contains the optional parameters for the VirtualMachinesClient.BeginReimage
//     method.
func (client *VirtualMachinesClient) BeginReimage(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*runtime.Poller[VirtualMachinesClientReimageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reimage(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientReimageResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientReimageResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reimage - Reimage the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) reimage(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginReimage"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reimageCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
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

// reimageCreateRequest creates the Reimage request.
func (client *VirtualMachinesClient) reimageCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginReimageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/reimage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRestart - Restart the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientBeginRestartOptions contains the optional parameters for the VirtualMachinesClient.BeginRestart
//     method.
func (client *VirtualMachinesClient) BeginRestart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*runtime.Poller[VirtualMachinesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - Restart the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) restart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
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
func (client *VirtualMachinesClient) restartCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Start the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - VirtualMachinesClientBeginStartOptions contains the optional parameters for the VirtualMachinesClient.BeginStart
//     method.
func (client *VirtualMachinesClient) BeginStart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*runtime.Poller[VirtualMachinesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Start the provided virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) start(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
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

// startCreateRequest creates the Start request.
func (client *VirtualMachinesClient) startCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Patch the properties of the provided virtual machine, or update the tags associated with the virtual machine.
// Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - virtualMachineUpdateParameters - The request body.
//   - options - VirtualMachinesClientBeginUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginUpdate
//     method.
func (client *VirtualMachinesClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineUpdateParameters VirtualMachinePatchParameters, options *VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachinesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualMachineName, virtualMachineUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualMachinesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualMachinesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch the properties of the provided virtual machine, or update the tags associated with the virtual machine.
// Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *VirtualMachinesClient) update(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineUpdateParameters VirtualMachinePatchParameters, options *VirtualMachinesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualMachinesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualMachineName, virtualMachineUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *VirtualMachinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineUpdateParameters VirtualMachinePatchParameters, options *VirtualMachinesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, virtualMachineUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
