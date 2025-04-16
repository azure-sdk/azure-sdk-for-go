// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// AddonsClient contains the methods for the Addons group.
// Don't use this type directly, use NewAddonsClient() instead.
type AddonsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAddonsClient creates a new instance of AddonsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAddonsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AddonsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AddonsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a addon.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - addonName - The addon name.
//   - resourceGroupName - The resource group name.
//   - addon - The addon properties.
//   - options - AddonsClientBeginCreateOrUpdateOptions contains the optional parameters for the AddonsClient.BeginCreateOrUpdate
//     method.
func (client *AddonsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, options *AddonsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AddonsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, roleName, addonName, resourceGroupName, addon, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AddonsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AddonsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a addon.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *AddonsClient) createOrUpdate(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, options *AddonsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AddonsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, addon, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AddonsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, _ *AddonsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, addon); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the addon on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - addonName - The addon name.
//   - resourceGroupName - The resource group name.
//   - options - AddonsClientBeginDeleteOptions contains the optional parameters for the AddonsClient.BeginDelete method.
func (client *AddonsClient) BeginDelete(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsClientBeginDeleteOptions) (*runtime.Poller[AddonsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, roleName, addonName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AddonsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AddonsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the addon on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *AddonsClient) deleteOperation(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AddonsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AddonsClient) deleteCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, _ *AddonsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a specific addon by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - addonName - The addon name.
//   - resourceGroupName - The resource group name.
//   - options - AddonsClientGetOptions contains the optional parameters for the AddonsClient.Get method.
func (client *AddonsClient) Get(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsClientGetOptions) (AddonsClientGetResponse, error) {
	var err error
	const operationName = "AddonsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, options)
	if err != nil {
		return AddonsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AddonsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AddonsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AddonsClient) getCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, _ *AddonsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AddonsClient) getHandleResponse(resp *http.Response) (AddonsClientGetResponse, error) {
	result := AddonsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AddonsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRolePager - Lists all the addons configured in the role.
//
// Generated from API version 2023-12-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - resourceGroupName - The resource group name.
//   - options - AddonsClientListByRoleOptions contains the optional parameters for the AddonsClient.NewListByRolePager method.
func (client *AddonsClient) NewListByRolePager(deviceName string, roleName string, resourceGroupName string, options *AddonsClientListByRoleOptions) *runtime.Pager[AddonsClientListByRoleResponse] {
	return runtime.NewPager(runtime.PagingHandler[AddonsClientListByRoleResponse]{
		More: func(page AddonsClientListByRoleResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AddonsClientListByRoleResponse) (AddonsClientListByRoleResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AddonsClient.NewListByRolePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByRoleCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AddonsClientListByRoleResponse{}, err
			}
			return client.listByRoleHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByRoleCreateRequest creates the ListByRole request.
func (client *AddonsClient) listByRoleCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, _ *AddonsClientListByRoleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRoleHandleResponse handles the ListByRole response.
func (client *AddonsClient) listByRoleHandleResponse(resp *http.Response) (AddonsClientListByRoleResponse, error) {
	result := AddonsClientListByRoleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddonList); err != nil {
		return AddonsClientListByRoleResponse{}, err
	}
	return result, nil
}
