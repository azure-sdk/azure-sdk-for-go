//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// ResourceHealthMetadataClient contains the methods for the ResourceHealthMetadata group.
// Don't use this type directly, use NewResourceHealthMetadataClient() instead.
type ResourceHealthMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceHealthMetadataClient creates a new instance of ResourceHealthMetadataClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceHealthMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceHealthMetadataClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceHealthMetadataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetBySite - Description for Gets the category of ResourceHealthMetadata to use for the given site
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of web app
// options - ResourceHealthMetadataClientGetBySiteOptions contains the optional parameters for the ResourceHealthMetadataClient.GetBySite
// method.
func (client *ResourceHealthMetadataClient) GetBySite(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataClientGetBySiteOptions) (ResourceHealthMetadataClientGetBySiteResponse, error) {
	req, err := client.getBySiteCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return ResourceHealthMetadataClientGetBySiteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceHealthMetadataClientGetBySiteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceHealthMetadataClientGetBySiteResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBySiteHandleResponse(resp)
}

// getBySiteCreateRequest creates the GetBySite request.
func (client *ResourceHealthMetadataClient) getBySiteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataClientGetBySiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBySiteHandleResponse handles the GetBySite response.
func (client *ResourceHealthMetadataClient) getBySiteHandleResponse(resp *http.Response) (ResourceHealthMetadataClientGetBySiteResponse, error) {
	result := ResourceHealthMetadataClientGetBySiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadata); err != nil {
		return ResourceHealthMetadataClientGetBySiteResponse{}, err
	}
	return result, nil
}

// GetBySiteSlot - Description for Gets the category of ResourceHealthMetadata to use for the given site
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of web app
// slot - Name of web app slot. If not specified then will default to production slot.
// options - ResourceHealthMetadataClientGetBySiteSlotOptions contains the optional parameters for the ResourceHealthMetadataClient.GetBySiteSlot
// method.
func (client *ResourceHealthMetadataClient) GetBySiteSlot(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataClientGetBySiteSlotOptions) (ResourceHealthMetadataClientGetBySiteSlotResponse, error) {
	req, err := client.getBySiteSlotCreateRequest(ctx, resourceGroupName, name, slot, options)
	if err != nil {
		return ResourceHealthMetadataClientGetBySiteSlotResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceHealthMetadataClientGetBySiteSlotResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceHealthMetadataClientGetBySiteSlotResponse{}, runtime.NewResponseError(resp)
	}
	return client.getBySiteSlotHandleResponse(resp)
}

// getBySiteSlotCreateRequest creates the GetBySiteSlot request.
func (client *ResourceHealthMetadataClient) getBySiteSlotCreateRequest(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataClientGetBySiteSlotOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if slot == "" {
		return nil, errors.New("parameter slot cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{slot}", url.PathEscape(slot))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBySiteSlotHandleResponse handles the GetBySiteSlot response.
func (client *ResourceHealthMetadataClient) getBySiteSlotHandleResponse(resp *http.Response) (ResourceHealthMetadataClientGetBySiteSlotResponse, error) {
	result := ResourceHealthMetadataClientGetBySiteSlotResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadata); err != nil {
		return ResourceHealthMetadataClientGetBySiteSlotResponse{}, err
	}
	return result, nil
}

// NewListPager - Description for List all ResourceHealthMetadata for all sites in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ResourceHealthMetadataClientListOptions contains the optional parameters for the ResourceHealthMetadataClient.List
// method.
func (client *ResourceHealthMetadataClient) NewListPager(options *ResourceHealthMetadataClientListOptions) *runtime.Pager[ResourceHealthMetadataClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourceHealthMetadataClientListResponse]{
		More: func(page ResourceHealthMetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceHealthMetadataClientListResponse) (ResourceHealthMetadataClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceHealthMetadataClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceHealthMetadataClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceHealthMetadataClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ResourceHealthMetadataClient) listCreateRequest(ctx context.Context, options *ResourceHealthMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceHealthMetadataClient) listHandleResponse(resp *http.Response) (ResourceHealthMetadataClientListResponse, error) {
	result := ResourceHealthMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Description for List all ResourceHealthMetadata for all sites in the resource group in the
// subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group to which the resource belongs.
// options - ResourceHealthMetadataClientListByResourceGroupOptions contains the optional parameters for the ResourceHealthMetadataClient.ListByResourceGroup
// method.
func (client *ResourceHealthMetadataClient) NewListByResourceGroupPager(resourceGroupName string, options *ResourceHealthMetadataClientListByResourceGroupOptions) *runtime.Pager[ResourceHealthMetadataClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourceHealthMetadataClientListByResourceGroupResponse]{
		More: func(page ResourceHealthMetadataClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceHealthMetadataClientListByResourceGroupResponse) (ResourceHealthMetadataClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceHealthMetadataClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceHealthMetadataClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceHealthMetadataClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ResourceHealthMetadataClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ResourceHealthMetadataClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ResourceHealthMetadataClient) listByResourceGroupHandleResponse(resp *http.Response) (ResourceHealthMetadataClientListByResourceGroupResponse, error) {
	result := ResourceHealthMetadataClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySitePager - Description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of web app.
// options - ResourceHealthMetadataClientListBySiteOptions contains the optional parameters for the ResourceHealthMetadataClient.ListBySite
// method.
func (client *ResourceHealthMetadataClient) NewListBySitePager(resourceGroupName string, name string, options *ResourceHealthMetadataClientListBySiteOptions) *runtime.Pager[ResourceHealthMetadataClientListBySiteResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourceHealthMetadataClientListBySiteResponse]{
		More: func(page ResourceHealthMetadataClientListBySiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceHealthMetadataClientListBySiteResponse) (ResourceHealthMetadataClientListBySiteResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySiteCreateRequest(ctx, resourceGroupName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceHealthMetadataClientListBySiteResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceHealthMetadataClientListBySiteResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceHealthMetadataClientListBySiteResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySiteHandleResponse(resp)
		},
	})
}

// listBySiteCreateRequest creates the ListBySite request.
func (client *ResourceHealthMetadataClient) listBySiteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ResourceHealthMetadataClientListBySiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySiteHandleResponse handles the ListBySite response.
func (client *ResourceHealthMetadataClient) listBySiteHandleResponse(resp *http.Response) (ResourceHealthMetadataClientListBySiteResponse, error) {
	result := ResourceHealthMetadataClientListBySiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataClientListBySiteResponse{}, err
	}
	return result, nil
}

// NewListBySiteSlotPager - Description for Gets the category of ResourceHealthMetadata to use for the given site as a collection
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of web app.
// slot - Name of web app slot. If not specified then will default to production slot.
// options - ResourceHealthMetadataClientListBySiteSlotOptions contains the optional parameters for the ResourceHealthMetadataClient.ListBySiteSlot
// method.
func (client *ResourceHealthMetadataClient) NewListBySiteSlotPager(resourceGroupName string, name string, slot string, options *ResourceHealthMetadataClientListBySiteSlotOptions) *runtime.Pager[ResourceHealthMetadataClientListBySiteSlotResponse] {
	return runtime.NewPager(runtime.PageProcessor[ResourceHealthMetadataClientListBySiteSlotResponse]{
		More: func(page ResourceHealthMetadataClientListBySiteSlotResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceHealthMetadataClientListBySiteSlotResponse) (ResourceHealthMetadataClientListBySiteSlotResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySiteSlotCreateRequest(ctx, resourceGroupName, name, slot, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceHealthMetadataClientListBySiteSlotResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceHealthMetadataClientListBySiteSlotResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceHealthMetadataClientListBySiteSlotResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySiteSlotHandleResponse(resp)
		},
	})
}

// listBySiteSlotCreateRequest creates the ListBySiteSlot request.
func (client *ResourceHealthMetadataClient) listBySiteSlotCreateRequest(ctx context.Context, resourceGroupName string, name string, slot string, options *ResourceHealthMetadataClientListBySiteSlotOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if slot == "" {
		return nil, errors.New("parameter slot cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{slot}", url.PathEscape(slot))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySiteSlotHandleResponse handles the ListBySiteSlot response.
func (client *ResourceHealthMetadataClient) listBySiteSlotHandleResponse(resp *http.Response) (ResourceHealthMetadataClientListBySiteSlotResponse, error) {
	result := ResourceHealthMetadataClientListBySiteSlotResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceHealthMetadataCollection); err != nil {
		return ResourceHealthMetadataClientListBySiteSlotResponse{}, err
	}
	return result, nil
}
