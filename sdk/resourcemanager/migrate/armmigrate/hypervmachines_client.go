//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmigrate

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

// HyperVMachinesClient contains the methods for the HyperVMachines group.
// Don't use this type directly, use NewHyperVMachinesClient() instead.
type HyperVMachinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHyperVMachinesClient creates a new instance of HyperVMachinesClient with the specified values.
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHyperVMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HyperVMachinesClient, error) {
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
	client := &HyperVMachinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewGetAllMachinesInSitePager - Method to get machine.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - options - HyperVMachinesClientGetAllMachinesInSiteOptions contains the optional parameters for the HyperVMachinesClient.NewGetAllMachinesInSitePager
//     method.
func (client *HyperVMachinesClient) NewGetAllMachinesInSitePager(resourceGroupName string, siteName string, options *HyperVMachinesClientGetAllMachinesInSiteOptions) *runtime.Pager[HyperVMachinesClientGetAllMachinesInSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[HyperVMachinesClientGetAllMachinesInSiteResponse]{
		More: func(page HyperVMachinesClientGetAllMachinesInSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HyperVMachinesClientGetAllMachinesInSiteResponse) (HyperVMachinesClientGetAllMachinesInSiteResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getAllMachinesInSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HyperVMachinesClientGetAllMachinesInSiteResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HyperVMachinesClientGetAllMachinesInSiteResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HyperVMachinesClientGetAllMachinesInSiteResponse{}, runtime.NewResponseError(resp)
			}
			return client.getAllMachinesInSiteHandleResponse(resp)
		},
	})
}

// getAllMachinesInSiteCreateRequest creates the GetAllMachinesInSite request.
func (client *HyperVMachinesClient) getAllMachinesInSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *HyperVMachinesClientGetAllMachinesInSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/machines"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-07")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllMachinesInSiteHandleResponse handles the GetAllMachinesInSite response.
func (client *HyperVMachinesClient) getAllMachinesInSiteHandleResponse(resp *http.Response) (HyperVMachinesClientGetAllMachinesInSiteResponse, error) {
	result := HyperVMachinesClientGetAllMachinesInSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVMachineCollection); err != nil {
		return HyperVMachinesClientGetAllMachinesInSiteResponse{}, err
	}
	return result, nil
}

// GetMachine - Method to get machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-07-07
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name.
//   - machineName - Machine ARM name.
//   - options - HyperVMachinesClientGetMachineOptions contains the optional parameters for the HyperVMachinesClient.GetMachine
//     method.
func (client *HyperVMachinesClient) GetMachine(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *HyperVMachinesClientGetMachineOptions) (HyperVMachinesClientGetMachineResponse, error) {
	req, err := client.getMachineCreateRequest(ctx, resourceGroupName, siteName, machineName, options)
	if err != nil {
		return HyperVMachinesClientGetMachineResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HyperVMachinesClientGetMachineResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HyperVMachinesClientGetMachineResponse{}, runtime.NewResponseError(resp)
	}
	return client.getMachineHandleResponse(resp)
}

// getMachineCreateRequest creates the GetMachine request.
func (client *HyperVMachinesClient) getMachineCreateRequest(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *HyperVMachinesClientGetMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/HyperVSites/{siteName}/machines/{machineName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-07-07")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getMachineHandleResponse handles the GetMachine response.
func (client *HyperVMachinesClient) getMachineHandleResponse(resp *http.Response) (HyperVMachinesClientGetMachineResponse, error) {
	result := HyperVMachinesClientGetMachineResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HyperVMachine); err != nil {
		return HyperVMachinesClientGetMachineResponse{}, err
	}
	return result, nil
}
