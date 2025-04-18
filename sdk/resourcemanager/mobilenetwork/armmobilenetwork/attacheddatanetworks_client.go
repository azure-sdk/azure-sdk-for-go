// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// AttachedDataNetworksClient contains the methods for the AttachedDataNetworks group.
// Don't use this type directly, use NewAttachedDataNetworksClient() instead.
type AttachedDataNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAttachedDataNetworksClient creates a new instance of AttachedDataNetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAttachedDataNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AttachedDataNetworksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AttachedDataNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an attached data network. Must be created in the same location as its parent packet
// core data plane.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCoreDataPlaneName - The name of the packet core data plane.
//   - attachedDataNetworkName - The name of the attached data network.
//   - parameters - Parameters supplied to the create or update attached data network operation.
//   - options - AttachedDataNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the AttachedDataNetworksClient.BeginCreateOrUpdate
//     method.
func (client *AttachedDataNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, options *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[AttachedDataNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AttachedDataNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AttachedDataNetworksClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an attached data network. Must be created in the same location as its parent packet
// core data plane.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *AttachedDataNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, options *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AttachedDataNetworksClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AttachedDataNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters AttachedDataNetwork, _ *AttachedDataNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCoreDataPlaneName - The name of the packet core data plane.
//   - attachedDataNetworkName - The name of the attached data network.
//   - options - AttachedDataNetworksClientBeginDeleteOptions contains the optional parameters for the AttachedDataNetworksClient.BeginDelete
//     method.
func (client *AttachedDataNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientBeginDeleteOptions) (*runtime.Poller[AttachedDataNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AttachedDataNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AttachedDataNetworksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *AttachedDataNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AttachedDataNetworksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
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
func (client *AttachedDataNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, _ *AttachedDataNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified attached data network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCoreDataPlaneName - The name of the packet core data plane.
//   - attachedDataNetworkName - The name of the attached data network.
//   - options - AttachedDataNetworksClientGetOptions contains the optional parameters for the AttachedDataNetworksClient.Get
//     method.
func (client *AttachedDataNetworksClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, options *AttachedDataNetworksClientGetOptions) (AttachedDataNetworksClientGetResponse, error) {
	var err error
	const operationName = "AttachedDataNetworksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, options)
	if err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AttachedDataNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AttachedDataNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, _ *AttachedDataNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AttachedDataNetworksClient) getHandleResponse(resp *http.Response) (AttachedDataNetworksClientGetResponse, error) {
	result := AttachedDataNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetwork); err != nil {
		return AttachedDataNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPacketCoreDataPlanePager - Gets all the attached data networks associated with a packet core data plane.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCoreDataPlaneName - The name of the packet core data plane.
//   - options - AttachedDataNetworksClientListByPacketCoreDataPlaneOptions contains the optional parameters for the AttachedDataNetworksClient.NewListByPacketCoreDataPlanePager
//     method.
func (client *AttachedDataNetworksClient) NewListByPacketCoreDataPlanePager(resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, options *AttachedDataNetworksClientListByPacketCoreDataPlaneOptions) *runtime.Pager[AttachedDataNetworksClientListByPacketCoreDataPlaneResponse] {
	return runtime.NewPager(runtime.PagingHandler[AttachedDataNetworksClientListByPacketCoreDataPlaneResponse]{
		More: func(page AttachedDataNetworksClientListByPacketCoreDataPlaneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AttachedDataNetworksClientListByPacketCoreDataPlaneResponse) (AttachedDataNetworksClientListByPacketCoreDataPlaneResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AttachedDataNetworksClient.NewListByPacketCoreDataPlanePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPacketCoreDataPlaneCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, options)
			}, nil)
			if err != nil {
				return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, err
			}
			return client.listByPacketCoreDataPlaneHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPacketCoreDataPlaneCreateRequest creates the ListByPacketCoreDataPlane request.
func (client *AttachedDataNetworksClient) listByPacketCoreDataPlaneCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, _ *AttachedDataNetworksClientListByPacketCoreDataPlaneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPacketCoreDataPlaneHandleResponse handles the ListByPacketCoreDataPlane response.
func (client *AttachedDataNetworksClient) listByPacketCoreDataPlaneHandleResponse(resp *http.Response) (AttachedDataNetworksClientListByPacketCoreDataPlaneResponse, error) {
	result := AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetworkListResult); err != nil {
		return AttachedDataNetworksClientListByPacketCoreDataPlaneResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an attached data network tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - packetCoreDataPlaneName - The name of the packet core data plane.
//   - attachedDataNetworkName - The name of the attached data network.
//   - parameters - Parameters supplied to update attached data network tags.
//   - options - AttachedDataNetworksClientUpdateTagsOptions contains the optional parameters for the AttachedDataNetworksClient.UpdateTags
//     method.
func (client *AttachedDataNetworksClient) UpdateTags(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters TagsObject, options *AttachedDataNetworksClientUpdateTagsOptions) (AttachedDataNetworksClientUpdateTagsResponse, error) {
	var err error
	const operationName = "AttachedDataNetworksClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName, parameters, options)
	if err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *AttachedDataNetworksClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string, attachedDataNetworkName string, parameters TagsObject, _ *AttachedDataNetworksClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/packetCoreDataPlanes/{packetCoreDataPlaneName}/attachedDataNetworks/{attachedDataNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if packetCoreDataPlaneName == "" {
		return nil, errors.New("parameter packetCoreDataPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreDataPlaneName}", url.PathEscape(packetCoreDataPlaneName))
	if attachedDataNetworkName == "" {
		return nil, errors.New("parameter attachedDataNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{attachedDataNetworkName}", url.PathEscape(attachedDataNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *AttachedDataNetworksClient) updateTagsHandleResponse(resp *http.Response) (AttachedDataNetworksClientUpdateTagsResponse, error) {
	result := AttachedDataNetworksClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AttachedDataNetwork); err != nil {
		return AttachedDataNetworksClientUpdateTagsResponse{}, err
	}
	return result, nil
}
