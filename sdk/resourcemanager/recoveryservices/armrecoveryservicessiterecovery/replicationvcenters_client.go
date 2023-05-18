//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery

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

// ReplicationvCentersClient contains the methods for the ReplicationvCenters group.
// Don't use this type directly, use NewReplicationvCentersClient() instead.
type ReplicationvCentersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationvCentersClient creates a new instance of ReplicationvCentersClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationvCentersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationvCentersClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationvCentersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationvCentersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to create a vCenter object..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - vcenterName - vcenter name.
//   - addVCenterRequest - The input to the add vCenter operation.
//   - options - ReplicationvCentersClientBeginCreateOptions contains the optional parameters for the ReplicationvCentersClient.BeginCreate
//     method.
func (client *ReplicationvCentersClient) BeginCreate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, addVCenterRequest AddVCenterRequest, options *ReplicationvCentersClientBeginCreateOptions) (*runtime.Poller[ReplicationvCentersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceName, resourceGroupName, fabricName, vcenterName, addVCenterRequest, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationvCentersClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationvCentersClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - The operation to create a vCenter object..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationvCentersClient) create(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, addVCenterRequest AddVCenterRequest, options *ReplicationvCentersClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceName, resourceGroupName, fabricName, vcenterName, addVCenterRequest, options)
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

// createCreateRequest creates the Create request.
func (client *ReplicationvCentersClient) createCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, addVCenterRequest AddVCenterRequest, options *ReplicationvCentersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vcenterName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, addVCenterRequest)
}

// BeginDelete - The operation to remove(unregister) a registered vCenter server from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - vcenterName - vcenter name.
//   - options - ReplicationvCentersClientBeginDeleteOptions contains the optional parameters for the ReplicationvCentersClient.BeginDelete
//     method.
func (client *ReplicationvCentersClient) BeginDelete(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, options *ReplicationvCentersClientBeginDeleteOptions) (*runtime.Poller[ReplicationvCentersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceName, resourceGroupName, fabricName, vcenterName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationvCentersClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationvCentersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to remove(unregister) a registered vCenter server from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationvCentersClient) deleteOperation(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, options *ReplicationvCentersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceName, resourceGroupName, fabricName, vcenterName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationvCentersClient) deleteCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, options *ReplicationvCentersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vcenterName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the details of a registered vCenter server(Add vCenter server).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - vcenterName - vcenter name.
//   - options - ReplicationvCentersClientGetOptions contains the optional parameters for the ReplicationvCentersClient.Get method.
func (client *ReplicationvCentersClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, options *ReplicationvCentersClientGetOptions) (ReplicationvCentersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, vcenterName, options)
	if err != nil {
		return ReplicationvCentersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationvCentersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationvCentersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationvCentersClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, options *ReplicationvCentersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vcenterName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationvCentersClient) getHandleResponse(resp *http.Response) (ReplicationvCentersClientGetResponse, error) {
	result := ReplicationvCentersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VCenter); err != nil {
		return ReplicationvCentersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the vCenter servers registered in the vault.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationvCentersClientListOptions contains the optional parameters for the ReplicationvCentersClient.NewListPager
//     method.
func (client *ReplicationvCentersClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationvCentersClientListOptions) *runtime.Pager[ReplicationvCentersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationvCentersClientListResponse]{
		More: func(page ReplicationvCentersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationvCentersClientListResponse) (ReplicationvCentersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationvCentersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationvCentersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationvCentersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationvCentersClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationvCentersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationvCentersClient) listHandleResponse(resp *http.Response) (ReplicationvCentersClientListResponse, error) {
	result := ReplicationvCentersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VCenterCollection); err != nil {
		return ReplicationvCentersClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists the vCenter servers registered in a fabric.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - options - ReplicationvCentersClientListByReplicationFabricsOptions contains the optional parameters for the ReplicationvCentersClient.NewListByReplicationFabricsPager
//     method.
func (client *ReplicationvCentersClient) NewListByReplicationFabricsPager(resourceName string, resourceGroupName string, fabricName string, options *ReplicationvCentersClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationvCentersClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationvCentersClientListByReplicationFabricsResponse]{
		More: func(page ReplicationvCentersClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationvCentersClientListByReplicationFabricsResponse) (ReplicationvCentersClientListByReplicationFabricsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationFabricsCreateRequest(ctx, resourceName, resourceGroupName, fabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationvCentersClientListByReplicationFabricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationvCentersClientListByReplicationFabricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationvCentersClientListByReplicationFabricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationvCentersClient) listByReplicationFabricsCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, options *ReplicationvCentersClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationvCentersClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationvCentersClientListByReplicationFabricsResponse, error) {
	result := ReplicationvCentersClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VCenterCollection); err != nil {
		return ReplicationvCentersClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update a registered vCenter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - vcenterName - vcenter name.
//   - updateVCenterRequest - The input to the update vCenter operation.
//   - options - ReplicationvCentersClientBeginUpdateOptions contains the optional parameters for the ReplicationvCentersClient.BeginUpdate
//     method.
func (client *ReplicationvCentersClient) BeginUpdate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, updateVCenterRequest UpdateVCenterRequest, options *ReplicationvCentersClientBeginUpdateOptions) (*runtime.Poller[ReplicationvCentersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceName, resourceGroupName, fabricName, vcenterName, updateVCenterRequest, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationvCentersClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationvCentersClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update a registered vCenter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationvCentersClient) update(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, updateVCenterRequest UpdateVCenterRequest, options *ReplicationvCentersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceName, resourceGroupName, fabricName, vcenterName, updateVCenterRequest, options)
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
func (client *ReplicationvCentersClient) updateCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, vcenterName string, updateVCenterRequest UpdateVCenterRequest, options *ReplicationvCentersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vcenterName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if vcenterName == "" {
		return nil, errors.New("parameter vcenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vcenterName}", url.PathEscape(vcenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateVCenterRequest)
}
