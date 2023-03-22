//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// PacketCapturesClient contains the methods for the PacketCaptures group.
// Don't use this type directly, use NewPacketCapturesClient() instead.
type PacketCapturesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPacketCapturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PacketCapturesClient, error) {
	cl, err := arm.NewClient(moduleName+".PacketCapturesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PacketCapturesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create and start a packet capture on the specified VM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the network watcher.
//   - packetCaptureName - The name of the packet capture session.
//   - parameters - Parameters that define the create packet capture operation.
//   - options - PacketCapturesClientBeginCreateOptions contains the optional parameters for the PacketCapturesClient.BeginCreate
//     method.
func (client *PacketCapturesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (*runtime.Poller[PacketCapturesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create and start a packet capture on the specified VM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *PacketCapturesClient) create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PacketCapturesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the network watcher.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientBeginDeleteOptions contains the optional parameters for the PacketCapturesClient.BeginDelete
//     method.
func (client *PacketCapturesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*runtime.Poller[PacketCapturesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *PacketCapturesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
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
func (client *PacketCapturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a packet capture session by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the network watcher.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientGetOptions contains the optional parameters for the PacketCapturesClient.Get method.
func (client *PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (PacketCapturesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCapturesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCapturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *PacketCapturesClient) getHandleResponse(resp *http.Response) (PacketCapturesClientGetResponse, error) {
	result := PacketCapturesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureResult); err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetStatus - Query the status of a running packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the Network Watcher resource.
//   - packetCaptureName - The name given to the packet capture session.
//   - options - PacketCapturesClientBeginGetStatusOptions contains the optional parameters for the PacketCapturesClient.BeginGetStatus
//     method.
func (client *PacketCapturesClient) BeginGetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (*runtime.Poller[PacketCapturesClientGetStatusResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getStatus(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientGetStatusResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientGetStatusResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GetStatus - Query the status of a running packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *PacketCapturesClient) getStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (*http.Response, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
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

// getStatusCreateRequest creates the GetStatus request.
func (client *PacketCapturesClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Lists all packet capture sessions within the specified resource group.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the Network Watcher resource.
//   - options - PacketCapturesClientListOptions contains the optional parameters for the PacketCapturesClient.NewListPager method.
func (client *PacketCapturesClient) NewListPager(resourceGroupName string, networkWatcherName string, options *PacketCapturesClientListOptions) *runtime.Pager[PacketCapturesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PacketCapturesClientListResponse]{
		More: func(page PacketCapturesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PacketCapturesClientListResponse) (PacketCapturesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
			if err != nil {
				return PacketCapturesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PacketCapturesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PacketCapturesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PacketCapturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *PacketCapturesClient) listHandleResponse(resp *http.Response) (PacketCapturesClientListResponse, error) {
	result := PacketCapturesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureListResult); err != nil {
		return PacketCapturesClientListResponse{}, err
	}
	return result, nil
}

// BeginStop - Stops a specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkWatcherName - The name of the network watcher.
//   - packetCaptureName - The name of the packet capture session.
//   - options - PacketCapturesClientBeginStopOptions contains the optional parameters for the PacketCapturesClient.BeginStop
//     method.
func (client *PacketCapturesClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*runtime.Poller[PacketCapturesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PacketCapturesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PacketCapturesClientStopResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Stop - Stops a specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *PacketCapturesClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
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

// stopCreateRequest creates the Stop request.
func (client *PacketCapturesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
