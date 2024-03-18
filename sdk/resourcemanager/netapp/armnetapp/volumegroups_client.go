//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// VolumeGroupsClient contains the methods for the VolumeGroups group.
// Don't use this type directly, use NewVolumeGroupsClient() instead.
type VolumeGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVolumeGroupsClient creates a new instance of VolumeGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVolumeGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumeGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VolumeGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a volume group along with specified volumes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - volumeGroupName - The name of the volumeGroup
//   - body - Volume Group object supplied in the body of the operation.
//   - options - VolumeGroupsClientBeginCreateOptions contains the optional parameters for the VolumeGroupsClient.BeginCreate
//     method.
func (client *VolumeGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsClientBeginCreateOptions) (*runtime.Poller[VolumeGroupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, volumeGroupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeGroupsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeGroupsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a volume group along with specified volumes
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *VolumeGroupsClient) create(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeGroupsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *VolumeGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, body VolumeGroupDetails, options *VolumeGroupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the specified volume group only if there are no volumes under volume group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - volumeGroupName - The name of the volumeGroup
//   - options - VolumeGroupsClientBeginDeleteOptions contains the optional parameters for the VolumeGroupsClient.BeginDelete
//     method.
func (client *VolumeGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsClientBeginDeleteOptions) (*runtime.Poller[VolumeGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, volumeGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the specified volume group only if there are no volumes under volume group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *VolumeGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, options)
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
func (client *VolumeGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get details of the specified volume group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - volumeGroupName - The name of the volumeGroup
//   - options - VolumeGroupsClientGetOptions contains the optional parameters for the VolumeGroupsClient.Get method.
func (client *VolumeGroupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsClientGetOptions) (VolumeGroupsClientGetResponse, error) {
	var err error
	const operationName = "VolumeGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, volumeGroupName, options)
	if err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VolumeGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VolumeGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, volumeGroupName string, options *VolumeGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups/{volumeGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if volumeGroupName == "" {
		return nil, errors.New("parameter volumeGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeGroupName}", url.PathEscape(volumeGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumeGroupsClient) getHandleResponse(resp *http.Response) (VolumeGroupsClientGetResponse, error) {
	result := VolumeGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroupDetails); err != nil {
		return VolumeGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNetAppAccountPager - List all volume groups for given account
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - options - VolumeGroupsClientListByNetAppAccountOptions contains the optional parameters for the VolumeGroupsClient.NewListByNetAppAccountPager
//     method.
func (client *VolumeGroupsClient) NewListByNetAppAccountPager(resourceGroupName string, accountName string, options *VolumeGroupsClientListByNetAppAccountOptions) *runtime.Pager[VolumeGroupsClientListByNetAppAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumeGroupsClientListByNetAppAccountResponse]{
		More: func(page VolumeGroupsClientListByNetAppAccountResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *VolumeGroupsClientListByNetAppAccountResponse) (VolumeGroupsClientListByNetAppAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VolumeGroupsClient.NewListByNetAppAccountPager")
			req, err := client.listByNetAppAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return VolumeGroupsClientListByNetAppAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VolumeGroupsClientListByNetAppAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumeGroupsClientListByNetAppAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNetAppAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNetAppAccountCreateRequest creates the ListByNetAppAccount request.
func (client *VolumeGroupsClient) listByNetAppAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *VolumeGroupsClientListByNetAppAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/volumeGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNetAppAccountHandleResponse handles the ListByNetAppAccount response.
func (client *VolumeGroupsClient) listByNetAppAccountHandleResponse(resp *http.Response) (VolumeGroupsClientListByNetAppAccountResponse, error) {
	result := VolumeGroupsClientListByNetAppAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeGroupList); err != nil {
		return VolumeGroupsClientListByNetAppAccountResponse{}, err
	}
	return result, nil
}
