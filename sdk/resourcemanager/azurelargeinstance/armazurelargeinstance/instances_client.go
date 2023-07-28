//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurelargeinstance

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

// InstancesClient contains the methods for the AzureLargeInstanceInstances group.
// Don't use this type directly, use NewInstancesClient() instead.
type InstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInstancesClient creates a new instance of InstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InstancesClient, error) {
	cl, err := arm.NewClient(moduleName+".InstancesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListBySubscriptionPager - Gets a list of Azure Large Instances in the specified subscription. The operations returns
// various properties of each Azure Large instance.
//
// Generated from API version 2023-07-20
//   - options - InstancesClientListBySubscriptionOptions contains the optional parameters for the InstancesClient.NewListBySubscriptionPager
//     method.
func (client *InstancesClient) NewListBySubscriptionPager(options *InstancesClientListBySubscriptionOptions) *runtime.Pager[InstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[InstancesClientListBySubscriptionResponse]{
		More: func(page InstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InstancesClientListBySubscriptionResponse) (InstancesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InstancesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InstancesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InstancesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *InstancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *InstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureLargeInstance/azureLargeInstance"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *InstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (InstancesClientListBySubscriptionResponse, error) {
	result := InstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return InstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches the Tags field of a Azure Large instance for the specified subscription, resource group, and instance
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureLargeInstanceName - Name of the AzureLargeInstance.
//   - tagsParameter - Request body that only contains the new Tags field
//   - options - InstancesClientUpdateOptions contains the optional parameters for the InstancesClient.Update method.
func (client *InstancesClient) Update(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, tagsParameter Tags, options *InstancesClientUpdateOptions) (InstancesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureLargeInstanceName, tagsParameter, options)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *InstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureLargeInstanceName string, tagsParameter Tags, options *InstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureLargeInstance/azureLargeInstance/{azureLargeInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureLargeInstanceName == "" {
		return nil, errors.New("parameter azureLargeInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureLargeInstanceName}", url.PathEscape(azureLargeInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, tagsParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *InstancesClient) updateHandleResponse(resp *http.Response) (InstancesClientUpdateResponse, error) {
	result := InstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureLargeInstance); err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	return result, nil
}
