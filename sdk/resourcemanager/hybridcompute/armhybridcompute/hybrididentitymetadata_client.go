//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

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

// HybridIdentityMetadataClient contains the methods for the HybridIdentityMetadata group.
// Don't use this type directly, use NewHybridIdentityMetadataClient() instead.
type HybridIdentityMetadataClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHybridIdentityMetadataClient creates a new instance of HybridIdentityMetadataClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHybridIdentityMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HybridIdentityMetadataClient, error) {
	cl, err := arm.NewClient(moduleName+".HybridIdentityMetadataClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HybridIdentityMetadataClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Implements HybridIdentityMetadata GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - metadataName - Name of the HybridIdentityMetadata.
//   - options - HybridIdentityMetadataClientGetOptions contains the optional parameters for the HybridIdentityMetadataClient.Get
//     method.
func (client *HybridIdentityMetadataClient) Get(ctx context.Context, resourceGroupName string, machineName string, metadataName string, options *HybridIdentityMetadataClientGetOptions) (HybridIdentityMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, metadataName, options)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridIdentityMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridIdentityMetadataClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, metadataName string, options *HybridIdentityMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/hybridIdentityMetadata/{metadataName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridIdentityMetadataClient) getHandleResponse(resp *http.Response) (HybridIdentityMetadataClientGetResponse, error) {
	result := HybridIdentityMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadata); err != nil {
		return HybridIdentityMetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMachinesPager - Returns the list of HybridIdentityMetadata of the given machine.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - HybridIdentityMetadataClientListByMachinesOptions contains the optional parameters for the HybridIdentityMetadataClient.NewListByMachinesPager
//     method.
func (client *HybridIdentityMetadataClient) NewListByMachinesPager(resourceGroupName string, machineName string, options *HybridIdentityMetadataClientListByMachinesOptions) *runtime.Pager[HybridIdentityMetadataClientListByMachinesResponse] {
	return runtime.NewPager(runtime.PagingHandler[HybridIdentityMetadataClientListByMachinesResponse]{
		More: func(page HybridIdentityMetadataClientListByMachinesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HybridIdentityMetadataClientListByMachinesResponse) (HybridIdentityMetadataClientListByMachinesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByMachinesCreateRequest(ctx, resourceGroupName, machineName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HybridIdentityMetadataClientListByMachinesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HybridIdentityMetadataClientListByMachinesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HybridIdentityMetadataClientListByMachinesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMachinesHandleResponse(resp)
		},
	})
}

// listByMachinesCreateRequest creates the ListByMachines request.
func (client *HybridIdentityMetadataClient) listByMachinesCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *HybridIdentityMetadataClientListByMachinesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/hybridIdentityMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMachinesHandleResponse handles the ListByMachines response.
func (client *HybridIdentityMetadataClient) listByMachinesHandleResponse(resp *http.Response) (HybridIdentityMetadataClientListByMachinesResponse, error) {
	result := HybridIdentityMetadataClientListByMachinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridIdentityMetadataList); err != nil {
		return HybridIdentityMetadataClientListByMachinesResponse{}, err
	}
	return result, nil
}
