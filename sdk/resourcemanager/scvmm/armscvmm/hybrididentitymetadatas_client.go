//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armscvmm

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

// HybridIdentityMetadatasClient contains the methods for the HybridIdentityMetadatas group.
// Don't use this type directly, use NewHybridIdentityMetadatasClient() instead.
type HybridIdentityMetadatasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHybridIdentityMetadatasClient creates a new instance of HybridIdentityMetadatasClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHybridIdentityMetadatasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridIdentityMetadatasClient, error) {
	cl, err := arm.NewClient(moduleName+".HybridIdentityMetadatasClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HybridIdentityMetadatasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create Or Update HybridIdentityMetadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-21-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualMachineName - Name of the vm.
//   - metadataName - Name of the hybridIdentityMetadata.
//   - options - HybridIdentityMetadatasClientCreateOptions contains the optional parameters for the HybridIdentityMetadatasClient.Create
//     method.
func (client *HybridIdentityMetadatasClient) Create(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientCreateOptions) (HybridIdentityMetadatasClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadatasClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadatasClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridIdentityMetadatasClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *HybridIdentityMetadatasClient) createCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-21-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *HybridIdentityMetadatasClient) createHandleResponse(resp *http.Response) (HybridIdentityMetadatasClientCreateResponse, error) {
	result := HybridIdentityMetadatasClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadatasClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Implements HybridIdentityMetadata DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-21-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualMachineName - Name of the vm.
//   - metadataName - Name of the HybridIdentityMetadata.
//   - options - HybridIdentityMetadatasClientDeleteOptions contains the optional parameters for the HybridIdentityMetadatasClient.Delete
//     method.
func (client *HybridIdentityMetadatasClient) Delete(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientDeleteOptions) (HybridIdentityMetadatasClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadatasClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadatasClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HybridIdentityMetadatasClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HybridIdentityMetadatasClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridIdentityMetadatasClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-21-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements HybridIdentityMetadata GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-21-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualMachineName - Name of the vm.
//   - metadataName - Name of the HybridIdentityMetadata.
//   - options - HybridIdentityMetadatasClientGetOptions contains the optional parameters for the HybridIdentityMetadatasClient.Get
//     method.
func (client *HybridIdentityMetadatasClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientGetOptions) (HybridIdentityMetadatasClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadatasClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadatasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridIdentityMetadatasClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridIdentityMetadatasClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, metadataName string, options *HybridIdentityMetadatasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachines/{virtualMachineName}/hybridIdentityMetadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-21-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridIdentityMetadatasClient) getHandleResponse(resp *http.Response) (HybridIdentityMetadatasClientGetResponse, error) {
	result := HybridIdentityMetadatasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadatasClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVMPager - Returns the list of HybridIdentityMetadata of the given vm.
//
// Generated from API version 2022-05-21-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualMachineName - Name of the vm.
//   - options - HybridIdentityMetadatasClientListByVMOptions contains the optional parameters for the HybridIdentityMetadatasClient.NewListByVMPager
//     method.
func (client *HybridIdentityMetadatasClient) NewListByVMPager(resourceGroupName string, virtualMachineName string, options *HybridIdentityMetadatasClientListByVMOptions) *runtime.Pager[HybridIdentityMetadatasClientListByVMResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridIdentityMetadatasClientListByVMResponse]{
		More: func(page HybridIdentityMetadatasClientListByVMResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridIdentityMetadatasClientListByVMResponse) (HybridIdentityMetadatasClientListByVMResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVMCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridIdentityMetadatasClientListByVMResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HybridIdentityMetadatasClientListByVMResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridIdentityMetadatasClientListByVMResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVMHandleResponse(resp)
		},
	})
}

// listByVMCreateRequest creates the ListByVM request.
func (client *HybridIdentityMetadatasClient) listByVMCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *HybridIdentityMetadatasClientListByVMOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachines/{virtualMachineName}/hybridIdentityMetadata"
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
	reqQP.Set("api-version", "2022-05-21-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVMHandleResponse handles the ListByVM response.
func (client *HybridIdentityMetadatasClient) listByVMHandleResponse(resp *http.Response) (HybridIdentityMetadatasClientListByVMResponse, error) {
	result := HybridIdentityMetadatasClientListByVMResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadataList); err != nil {
		return HybridIdentityMetadatasClientListByVMResponse{}, err
	}
	return result, nil
}
