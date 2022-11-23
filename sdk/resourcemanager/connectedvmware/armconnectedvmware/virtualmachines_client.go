//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

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

// VirtualMachinesClient contains the methods for the VirtualMachines group.
// Don't use this type directly, use NewVirtualMachinesClient() instead.
type VirtualMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachinesClient creates a new instance of VirtualMachinesClient with the specified values.
// subscriptionID - The Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachinesClient, error) {
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
	client := &VirtualMachinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginAssessPatches - The operation to assess patches on a vSphere VMware machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The name of the resource group.
// virtualMachineName - The name of the vSphere VMware machine.
// options - VirtualMachinesClientBeginAssessPatchesOptions contains the optional parameters for the VirtualMachinesClient.BeginAssessPatches
// method.
func (client *VirtualMachinesClient) BeginAssessPatches(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginAssessPatchesOptions) (*runtime.Poller[VirtualMachinesClientAssessPatchesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.assessPatches(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientAssessPatchesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientAssessPatchesResponse](options.ResumeToken, client.pl, nil)
	}
}

// AssessPatches - The operation to assess patches on a vSphere VMware machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) assessPatches(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginAssessPatchesOptions) (*http.Response, error) {
	req, err := client.assessPatchesCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// assessPatchesCreateRequest creates the AssessPatches request.
func (client *VirtualMachinesClient) assessPatchesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginAssessPatchesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/assessPatches"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Create Or Update virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginCreateOrUpdate
// method.
func (client *VirtualMachinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachinesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create Or Update virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Implements virtual machine DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginDeleteOptions contains the optional parameters for the VirtualMachinesClient.BeginDelete
// method.
func (client *VirtualMachinesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachinesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Implements virtual machine DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	if options != nil && options.Retain != nil {
		reqQP.Set("retain", strconv.FormatBool(*options.Retain))
	}
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements virtual machine GET method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientGetOptions contains the optional parameters for the VirtualMachinesClient.Get method.
func (client *VirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (VirtualMachinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
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

// BeginInstallPatches - The operation to install patches on a vSphere VMware machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The name of the resource group.
// virtualMachineName - The name of the vSphere VMware machine.
// installPatchesInput - Input for InstallPatches as directly received by the API
// options - VirtualMachinesClientBeginInstallPatchesOptions contains the optional parameters for the VirtualMachinesClient.BeginInstallPatches
// method.
func (client *VirtualMachinesClient) BeginInstallPatches(ctx context.Context, resourceGroupName string, virtualMachineName string, installPatchesInput VirtualMachineInstallPatchesParameters, options *VirtualMachinesClientBeginInstallPatchesOptions) (*runtime.Poller[VirtualMachinesClientInstallPatchesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.installPatches(ctx, resourceGroupName, virtualMachineName, installPatchesInput, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientInstallPatchesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientInstallPatchesResponse](options.ResumeToken, client.pl, nil)
	}
}

// InstallPatches - The operation to install patches on a vSphere VMware machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) installPatches(ctx context.Context, resourceGroupName string, virtualMachineName string, installPatchesInput VirtualMachineInstallPatchesParameters, options *VirtualMachinesClientBeginInstallPatchesOptions) (*http.Response, error) {
	req, err := client.installPatchesCreateRequest(ctx, resourceGroupName, virtualMachineName, installPatchesInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// installPatchesCreateRequest creates the InstallPatches request.
func (client *VirtualMachinesClient) installPatchesCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, installPatchesInput VirtualMachineInstallPatchesParameters, options *VirtualMachinesClientBeginInstallPatchesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/installPatches"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, installPatchesInput)
}

// NewListPager - List of virtualMachines in a resource group.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// options - VirtualMachinesClientListOptions contains the optional parameters for the VirtualMachinesClient.List method.
func (client *VirtualMachinesClient) NewListPager(resourceGroupName string, options *VirtualMachinesClientListOptions) *runtime.Pager[VirtualMachinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachinesClientListResponse]{
		More: func(page VirtualMachinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachinesClientListResponse) (VirtualMachinesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualMachinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachinesClient) listHandleResponse(resp *http.Response) (VirtualMachinesClientListResponse, error) {
	result := VirtualMachinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachinesList); err != nil {
		return VirtualMachinesClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - List of virtualMachines in a subscription.
// Generated from API version 2022-07-15-preview
// options - VirtualMachinesClientListAllOptions contains the optional parameters for the VirtualMachinesClient.ListAll method.
func (client *VirtualMachinesClient) NewListAllPager(options *VirtualMachinesClientListAllOptions) *runtime.Pager[VirtualMachinesClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachinesClientListAllResponse]{
		More: func(page VirtualMachinesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachinesClientListAllResponse) (VirtualMachinesClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachinesClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachinesClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachinesClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *VirtualMachinesClient) listAllCreateRequest(ctx context.Context, options *VirtualMachinesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *VirtualMachinesClient) listAllHandleResponse(resp *http.Response) (VirtualMachinesClientListAllResponse, error) {
	result := VirtualMachinesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachinesList); err != nil {
		return VirtualMachinesClientListAllResponse{}, err
	}
	return result, nil
}

// BeginRestart - Restart virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginRestartOptions contains the optional parameters for the VirtualMachinesClient.BeginRestart
// method.
func (client *VirtualMachinesClient) BeginRestart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*runtime.Poller[VirtualMachinesClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientRestartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientRestartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Restart - Restart virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) restart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *VirtualMachinesClient) restartCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/restart"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Start virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginStartOptions contains the optional parameters for the VirtualMachinesClient.BeginStart
// method.
func (client *VirtualMachinesClient) BeginStart(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*runtime.Poller[VirtualMachinesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientStartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Start - Start virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) start(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *VirtualMachinesClient) startCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/start"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Stop virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginStopOptions contains the optional parameters for the VirtualMachinesClient.BeginStop
// method.
func (client *VirtualMachinesClient) BeginStop(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*runtime.Poller[VirtualMachinesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachinesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientStopResponse](options.ResumeToken, client.pl, nil)
	}
}

// Stop - Stop virtual machine.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) stop(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *VirtualMachinesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}/stop"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginUpdate - API to update certain properties of the virtual machine resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
// resourceGroupName - The Resource Group Name.
// virtualMachineName - Name of the virtual machine resource.
// options - VirtualMachinesClientBeginUpdateOptions contains the optional parameters for the VirtualMachinesClient.BeginUpdate
// method.
func (client *VirtualMachinesClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachinesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualMachineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachinesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachinesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - API to update certain properties of the virtual machine resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-15-preview
func (client *VirtualMachinesClient) update(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *VirtualMachinesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/{virtualMachineName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}
