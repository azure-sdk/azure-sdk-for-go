// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// VmwareSoftwareInventoriesControllerClient contains the methods for the VmwareSoftwareInventoriesController group.
// Don't use this type directly, use NewVmwareSoftwareInventoriesControllerClient() instead.
type VmwareSoftwareInventoriesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVmwareSoftwareInventoriesControllerClient creates a new instance of VmwareSoftwareInventoriesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVmwareSoftwareInventoriesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VmwareSoftwareInventoriesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VmwareSoftwareInventoriesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetMachineSoftwareInventory - Method to get a machines software inventory like applications and roles.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - machineName - Machine name
//   - defaultParam - Default value.
//   - options - VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions contains the optional parameters
//     for the VmwareSoftwareInventoriesControllerClient.GetMachineSoftwareInventory method.
func (client *VmwareSoftwareInventoriesControllerClient) GetMachineSoftwareInventory(ctx context.Context, resourceGroupName string, siteName string, machineName string, defaultParam Default, options *VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions) (VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse, error) {
	var err error
	const operationName = "VmwareSoftwareInventoriesControllerClient.GetMachineSoftwareInventory"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMachineSoftwareInventoryCreateRequest(ctx, resourceGroupName, siteName, machineName, defaultParam, options)
	if err != nil {
		return VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse{}, err
	}
	resp, err := client.getMachineSoftwareInventoryHandleResponse(httpResp)
	return resp, err
}

// getMachineSoftwareInventoryCreateRequest creates the GetMachineSoftwareInventory request.
func (client *VmwareSoftwareInventoriesControllerClient) getMachineSoftwareInventoryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, defaultParam Default, _ *VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/machines/{machineName}/softwareInventories/{default}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMachineSoftwareInventoryHandleResponse handles the GetMachineSoftwareInventory response.
func (client *VmwareSoftwareInventoriesControllerClient) getMachineSoftwareInventoryHandleResponse(resp *http.Response) (VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse, error) {
	result := VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareMachineSoftwareInventory); err != nil {
		return VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse{}, err
	}
	return result, nil
}

// NewListByMachineResourcePager - List VmwareMachineSoftwareInventory resources by MachineResource
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - machineName - Machine name
//   - options - VmwareSoftwareInventoriesControllerClientListByMachineResourceOptions contains the optional parameters for the
//     VmwareSoftwareInventoriesControllerClient.NewListByMachineResourcePager method.
func (client *VmwareSoftwareInventoriesControllerClient) NewListByMachineResourcePager(resourceGroupName string, siteName string, machineName string, options *VmwareSoftwareInventoriesControllerClientListByMachineResourceOptions) *runtime.Pager[VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse]{
		More: func(page VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse) (VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VmwareSoftwareInventoriesControllerClient.NewListByMachineResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMachineResourceCreateRequest(ctx, resourceGroupName, siteName, machineName, options)
			}, nil)
			if err != nil {
				return VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse{}, err
			}
			return client.listByMachineResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMachineResourceCreateRequest creates the ListByMachineResource request.
func (client *VmwareSoftwareInventoriesControllerClient) listByMachineResourceCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, _ *VmwareSoftwareInventoriesControllerClientListByMachineResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/machines/{machineName}/softwareinventories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMachineResourceHandleResponse handles the ListByMachineResource response.
func (client *VmwareSoftwareInventoriesControllerClient) listByMachineResourceHandleResponse(resp *http.Response) (VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse, error) {
	result := VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareMachineSoftwareInventoryListResult); err != nil {
		return VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse{}, err
	}
	return result, nil
}
