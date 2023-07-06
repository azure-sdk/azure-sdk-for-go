//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud

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

// RacksClient contains the methods for the Racks group.
// Don't use this type directly, use NewRacksClient() instead.
type RacksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRacksClient creates a new instance of RacksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRacksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RacksClient, error) {
	cl, err := arm.NewClient(moduleName+".RacksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RacksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new rack or update properties of the existing one. All customer initiated requests will
// be rejected as the life cycle of this resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - rackName - The name of the rack.
//   - rackParameters - The request body.
//   - options - RacksClientBeginCreateOrUpdateOptions contains the optional parameters for the RacksClient.BeginCreateOrUpdate
//     method.
func (client *RacksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, rackName string, rackParameters Rack, options *RacksClientBeginCreateOrUpdateOptions) (*runtime.Poller[RacksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, rackName, rackParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RacksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RacksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new rack or update properties of the existing one. All customer initiated requests will be rejected
// as the life cycle of this resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *RacksClient) createOrUpdate(ctx context.Context, resourceGroupName string, rackName string, rackParameters Rack, options *RacksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, rackName, rackParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RacksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, rackName string, rackParameters Rack, options *RacksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/racks/{rackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rackName == "" {
		return nil, errors.New("parameter rackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rackName}", url.PathEscape(rackName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, rackParameters)
}

// BeginDelete - Delete the provided rack. All customer initiated requests will be rejected as the life cycle of this resource
// is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - rackName - The name of the rack.
//   - options - RacksClientBeginDeleteOptions contains the optional parameters for the RacksClient.BeginDelete method.
func (client *RacksClient) BeginDelete(ctx context.Context, resourceGroupName string, rackName string, options *RacksClientBeginDeleteOptions) (*runtime.Poller[RacksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, rackName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RacksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RacksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided rack. All customer initiated requests will be rejected as the life cycle of this resource
// is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *RacksClient) deleteOperation(ctx context.Context, resourceGroupName string, rackName string, options *RacksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, rackName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RacksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, rackName string, options *RacksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/racks/{rackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rackName == "" {
		return nil, errors.New("parameter rackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rackName}", url.PathEscape(rackName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided rack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - rackName - The name of the rack.
//   - options - RacksClientGetOptions contains the optional parameters for the RacksClient.Get method.
func (client *RacksClient) Get(ctx context.Context, resourceGroupName string, rackName string, options *RacksClientGetOptions) (RacksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, rackName, options)
	if err != nil {
		return RacksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RacksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RacksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RacksClient) getCreateRequest(ctx context.Context, resourceGroupName string, rackName string, options *RacksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/racks/{rackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rackName == "" {
		return nil, errors.New("parameter rackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rackName}", url.PathEscape(rackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RacksClient) getHandleResponse(resp *http.Response) (RacksClientGetResponse, error) {
	result := RacksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Rack); err != nil {
		return RacksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of racks in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - RacksClientListByResourceGroupOptions contains the optional parameters for the RacksClient.NewListByResourceGroupPager
//     method.
func (client *RacksClient) NewListByResourceGroupPager(resourceGroupName string, options *RacksClientListByResourceGroupOptions) *runtime.Pager[RacksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[RacksClientListByResourceGroupResponse]{
		More: func(page RacksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RacksClientListByResourceGroupResponse) (RacksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RacksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RacksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RacksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RacksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RacksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/racks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RacksClient) listByResourceGroupHandleResponse(resp *http.Response) (RacksClientListByResourceGroupResponse, error) {
	result := RacksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RackList); err != nil {
		return RacksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of racks in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - RacksClientListBySubscriptionOptions contains the optional parameters for the RacksClient.NewListBySubscriptionPager
//     method.
func (client *RacksClient) NewListBySubscriptionPager(options *RacksClientListBySubscriptionOptions) *runtime.Pager[RacksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[RacksClientListBySubscriptionResponse]{
		More: func(page RacksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RacksClientListBySubscriptionResponse) (RacksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RacksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RacksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RacksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *RacksClient) listBySubscriptionCreateRequest(ctx context.Context, options *RacksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/racks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *RacksClient) listBySubscriptionHandleResponse(resp *http.Response) (RacksClientListBySubscriptionResponse, error) {
	result := RacksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RackList); err != nil {
		return RacksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch properties of the provided rack, or update the tags associated with the rack. Properties and tag updates
// can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - rackName - The name of the rack.
//   - rackUpdateParameters - The request body.
//   - options - RacksClientBeginUpdateOptions contains the optional parameters for the RacksClient.BeginUpdate method.
func (client *RacksClient) BeginUpdate(ctx context.Context, resourceGroupName string, rackName string, rackUpdateParameters RackPatchParameters, options *RacksClientBeginUpdateOptions) (*runtime.Poller[RacksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, rackName, rackUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RacksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RacksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Patch properties of the provided rack, or update the tags associated with the rack. Properties and tag updates
// can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *RacksClient) update(ctx context.Context, resourceGroupName string, rackName string, rackUpdateParameters RackPatchParameters, options *RacksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, rackName, rackUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RacksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, rackName string, rackUpdateParameters RackPatchParameters, options *RacksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/racks/{rackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if rackName == "" {
		return nil, errors.New("parameter rackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rackName}", url.PathEscape(rackName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, rackUpdateParameters)
}
