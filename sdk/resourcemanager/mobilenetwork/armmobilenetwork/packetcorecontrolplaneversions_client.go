//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

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
	"strings"
)

// PacketCoreControlPlaneVersionsClient contains the methods for the PacketCoreControlPlaneVersions group.
// Don't use this type directly, use NewPacketCoreControlPlaneVersionsClient() instead.
type PacketCoreControlPlaneVersionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPacketCoreControlPlaneVersionsClient creates a new instance of PacketCoreControlPlaneVersionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPacketCoreControlPlaneVersionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PacketCoreControlPlaneVersionsClient, error) {
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
	client := &PacketCoreControlPlaneVersionsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Gets information about the specified packet core control plane version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - versionName - The name of the packet core control plane version.
//   - options - PacketCoreControlPlaneVersionsClientGetOptions contains the optional parameters for the PacketCoreControlPlaneVersionsClient.Get
//     method.
func (client *PacketCoreControlPlaneVersionsClient) Get(ctx context.Context, versionName string, options *PacketCoreControlPlaneVersionsClientGetOptions) (PacketCoreControlPlaneVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, versionName, options)
	if err != nil {
		return PacketCoreControlPlaneVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCoreControlPlaneVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCoreControlPlaneVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCoreControlPlaneVersionsClient) getCreateRequest(ctx context.Context, versionName string, options *PacketCoreControlPlaneVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions/{versionName}"
	if versionName == "" {
		return nil, errors.New("parameter versionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionName}", url.PathEscape(versionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PacketCoreControlPlaneVersionsClient) getHandleResponse(resp *http.Response) (PacketCoreControlPlaneVersionsClientGetResponse, error) {
	result := PacketCoreControlPlaneVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlaneVersion); err != nil {
		return PacketCoreControlPlaneVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all supported packet core control planes versions.
//
// Generated from API version 2022-11-01
//   - options - PacketCoreControlPlaneVersionsClientListOptions contains the optional parameters for the PacketCoreControlPlaneVersionsClient.NewListPager
//     method.
func (client *PacketCoreControlPlaneVersionsClient) NewListPager(options *PacketCoreControlPlaneVersionsClientListOptions) *runtime.Pager[PacketCoreControlPlaneVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PacketCoreControlPlaneVersionsClientListResponse]{
		More: func(page PacketCoreControlPlaneVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PacketCoreControlPlaneVersionsClientListResponse) (PacketCoreControlPlaneVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PacketCoreControlPlaneVersionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PacketCoreControlPlaneVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PacketCoreControlPlaneVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PacketCoreControlPlaneVersionsClient) listCreateRequest(ctx context.Context, options *PacketCoreControlPlaneVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.MobileNetwork/packetCoreControlPlaneVersions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PacketCoreControlPlaneVersionsClient) listHandleResponse(resp *http.Response) (PacketCoreControlPlaneVersionsClientListResponse, error) {
	result := PacketCoreControlPlaneVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCoreControlPlaneVersionListResult); err != nil {
		return PacketCoreControlPlaneVersionsClientListResponse{}, err
	}
	return result, nil
}
