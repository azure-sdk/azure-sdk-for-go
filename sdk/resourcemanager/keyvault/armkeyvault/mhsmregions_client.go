//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkeyvault

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

// MHSMRegionsClient contains the methods for the MHSMRegions group.
// Don't use this type directly, use NewMHSMRegionsClient() instead.
type MHSMRegionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMHSMRegionsClient creates a new instance of MHSMRegionsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMHSMRegionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MHSMRegionsClient, error) {
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
	client := &MHSMRegionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByResourcePager - The List operation gets information about the regions associated with the managed HSM Pool.
//
// Generated from API version 2023-02-01
//   - resourceGroupName - Name of the resource group that contains the managed HSM pool.
//   - name - Name of the managed HSM Pool
//   - options - MHSMRegionsClientListByResourceOptions contains the optional parameters for the MHSMRegionsClient.NewListByResourcePager
//     method.
func (client *MHSMRegionsClient) NewListByResourcePager(resourceGroupName string, name string, options *MHSMRegionsClientListByResourceOptions) *runtime.Pager[MHSMRegionsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[MHSMRegionsClientListByResourceResponse]{
		More: func(page MHSMRegionsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MHSMRegionsClientListByResourceResponse) (MHSMRegionsClientListByResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceCreateRequest(ctx, resourceGroupName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MHSMRegionsClientListByResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MHSMRegionsClientListByResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MHSMRegionsClientListByResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceHandleResponse(resp)
		},
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *MHSMRegionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, name string, options *MHSMRegionsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/regions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *MHSMRegionsClient) listByResourceHandleResponse(resp *http.Response) (MHSMRegionsClientListByResourceResponse, error) {
	result := MHSMRegionsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MHSMRegionsListResult); err != nil {
		return MHSMRegionsClientListByResourceResponse{}, err
	}
	return result, nil
}
