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

// VirtualMachineTemplatesClient contains the methods for the VirtualMachineTemplates group.
// Don't use this type directly, use NewVirtualMachineTemplatesClient() instead.
type VirtualMachineTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineTemplatesClient creates a new instance of VirtualMachineTemplatesClient with the specified values.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineTemplatesClient, error) {
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
	client := &VirtualMachineTemplatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create Or Update virtual machine template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualMachineTemplateName - Name of the virtual machine template resource.
//   - body - Request payload.
//   - options - VirtualMachineTemplatesClientBeginCreateOptions contains the optional parameters for the VirtualMachineTemplatesClient.BeginCreate
//     method.
func (client *VirtualMachineTemplatesClient) BeginCreate(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOptions) (*runtime.Poller[VirtualMachineTemplatesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachineTemplatesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineTemplatesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create Or Update virtual machine template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
func (client *VirtualMachineTemplatesClient) create(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
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

// createCreateRequest creates the Create request.
func (client *VirtualMachineTemplatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements virtual machine template DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualMachineTemplateName - Name of the virtual machine template resource.
//   - options - VirtualMachineTemplatesClientBeginDeleteOptions contains the optional parameters for the VirtualMachineTemplatesClient.BeginDelete
//     method.
func (client *VirtualMachineTemplatesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineTemplatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualMachineTemplateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineTemplatesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineTemplatesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Implements virtual machine template DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
func (client *VirtualMachineTemplatesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, options)
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
func (client *VirtualMachineTemplatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements virtual machine template GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualMachineTemplateName - Name of the virtual machine template resource.
//   - options - VirtualMachineTemplatesClientGetOptions contains the optional parameters for the VirtualMachineTemplatesClient.Get
//     method.
func (client *VirtualMachineTemplatesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (VirtualMachineTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, options)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
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
func (client *VirtualMachineTemplatesClient) getHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientGetResponse, error) {
	result := VirtualMachineTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplate); err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List of virtualMachineTemplates in a subscription.
//
// Generated from API version 2022-07-15-preview
//   - options - VirtualMachineTemplatesClientListOptions contains the optional parameters for the VirtualMachineTemplatesClient.NewListPager
//     method.
func (client *VirtualMachineTemplatesClient) NewListPager(options *VirtualMachineTemplatesClientListOptions) *runtime.Pager[VirtualMachineTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineTemplatesClientListResponse]{
		More: func(page VirtualMachineTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineTemplatesClientListResponse) (VirtualMachineTemplatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineTemplatesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineTemplatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineTemplatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineTemplatesClient) listCreateRequest(ctx context.Context, options *VirtualMachineTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates"
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

// listHandleResponse handles the List response.
func (client *VirtualMachineTemplatesClient) listHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientListResponse, error) {
	result := VirtualMachineTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplatesList); err != nil {
		return VirtualMachineTemplatesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of virtualMachineTemplates in a resource group.
//
// Generated from API version 2022-07-15-preview
//   - resourceGroupName - The Resource Group Name.
//   - options - VirtualMachineTemplatesClientListByResourceGroupOptions contains the optional parameters for the VirtualMachineTemplatesClient.NewListByResourceGroupPager
//     method.
func (client *VirtualMachineTemplatesClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualMachineTemplatesClientListByResourceGroupOptions) *runtime.Pager[VirtualMachineTemplatesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineTemplatesClientListByResourceGroupResponse]{
		More: func(page VirtualMachineTemplatesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineTemplatesClientListByResourceGroupResponse) (VirtualMachineTemplatesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualMachineTemplatesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualMachineTemplatesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates"
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualMachineTemplatesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientListByResourceGroupResponse, error) {
	result := VirtualMachineTemplatesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplatesList); err != nil {
		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the virtual machine template resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-15-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualMachineTemplateName - Name of the virtual machine template resource.
//   - body - Resource properties to update.
//   - options - VirtualMachineTemplatesClientUpdateOptions contains the optional parameters for the VirtualMachineTemplatesClient.Update
//     method.
func (client *VirtualMachineTemplatesClient) Update(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body ResourcePatch, options *VirtualMachineTemplatesClientUpdateOptions) (VirtualMachineTemplatesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
	if err != nil {
		return VirtualMachineTemplatesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineTemplatesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineTemplatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body ResourcePatch, options *VirtualMachineTemplatesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *VirtualMachineTemplatesClient) updateHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientUpdateResponse, error) {
	result := VirtualMachineTemplatesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplate); err != nil {
		return VirtualMachineTemplatesClientUpdateResponse{}, err
	}
	return result, nil
}
