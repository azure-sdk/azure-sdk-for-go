// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetworkfabric

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

// NetworkToNetworkInterconnectsClient contains the methods for the NetworkToNetworkInterconnects group.
// Don't use this type directly, use NewNetworkToNetworkInterconnectsClient() instead.
type NetworkToNetworkInterconnectsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkToNetworkInterconnectsClient creates a new instance of NetworkToNetworkInterconnectsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkToNetworkInterconnectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkToNetworkInterconnectsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkToNetworkInterconnectsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Configuration used to setup CE-PE connectivity PUT Method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - body - Request payload.
//   - options - NetworkToNetworkInterconnectsClientBeginCreateOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.BeginCreate
//     method.
func (client *NetworkToNetworkInterconnectsClient) BeginCreate(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, options *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkToNetworkInterconnectsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Configuration used to setup CE-PE connectivity PUT Method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkToNetworkInterconnectsClient) create(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, options *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
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

// createCreateRequest creates the Create request.
func (client *NetworkToNetworkInterconnectsClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnect, _ *NetworkToNetworkInterconnectsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Implements NetworkToNetworkInterconnects DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - options - NetworkToNetworkInterconnectsClientBeginDeleteOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.BeginDelete
//     method.
func (client *NetworkToNetworkInterconnectsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkToNetworkInterconnectsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Implements NetworkToNetworkInterconnects DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkToNetworkInterconnectsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkToNetworkInterconnectsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, _ *NetworkToNetworkInterconnectsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements NetworkToNetworkInterconnects GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - options - NetworkToNetworkInterconnectsClientGetOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.Get
//     method.
func (client *NetworkToNetworkInterconnectsClient) Get(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *NetworkToNetworkInterconnectsClientGetOptions) (NetworkToNetworkInterconnectsClientGetResponse, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, options)
	if err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkToNetworkInterconnectsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, _ *NetworkToNetworkInterconnectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkToNetworkInterconnectsClient) getHandleResponse(resp *http.Response) (NetworkToNetworkInterconnectsClientGetResponse, error) {
	result := NetworkToNetworkInterconnectsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkToNetworkInterconnect); err != nil {
		return NetworkToNetworkInterconnectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNetworkFabricPager - Implements Network To Network Interconnects list by Network Fabric GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - options - NetworkToNetworkInterconnectsClientListByNetworkFabricOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.NewListByNetworkFabricPager
//     method.
func (client *NetworkToNetworkInterconnectsClient) NewListByNetworkFabricPager(resourceGroupName string, networkFabricName string, options *NetworkToNetworkInterconnectsClientListByNetworkFabricOptions) *runtime.Pager[NetworkToNetworkInterconnectsClientListByNetworkFabricResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkToNetworkInterconnectsClientListByNetworkFabricResponse]{
		More: func(page NetworkToNetworkInterconnectsClientListByNetworkFabricResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkToNetworkInterconnectsClientListByNetworkFabricResponse) (NetworkToNetworkInterconnectsClientListByNetworkFabricResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkToNetworkInterconnectsClient.NewListByNetworkFabricPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNetworkFabricCreateRequest(ctx, resourceGroupName, networkFabricName, options)
			}, nil)
			if err != nil {
				return NetworkToNetworkInterconnectsClientListByNetworkFabricResponse{}, err
			}
			return client.listByNetworkFabricHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNetworkFabricCreateRequest creates the ListByNetworkFabric request.
func (client *NetworkToNetworkInterconnectsClient) listByNetworkFabricCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, _ *NetworkToNetworkInterconnectsClientListByNetworkFabricOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNetworkFabricHandleResponse handles the ListByNetworkFabric response.
func (client *NetworkToNetworkInterconnectsClient) listByNetworkFabricHandleResponse(resp *http.Response) (NetworkToNetworkInterconnectsClientListByNetworkFabricResponse, error) {
	result := NetworkToNetworkInterconnectsClientListByNetworkFabricResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkToNetworkInterconnectsList); err != nil {
		return NetworkToNetworkInterconnectsClientListByNetworkFabricResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update certain properties of the Network To NetworkInterconnects resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - body - Network to Network Interconnect properties to update.
//   - options - NetworkToNetworkInterconnectsClientBeginUpdateOptions contains the optional parameters for the NetworkToNetworkInterconnectsClient.BeginUpdate
//     method.
func (client *NetworkToNetworkInterconnectsClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnectPatch, options *NetworkToNetworkInterconnectsClientBeginUpdateOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkToNetworkInterconnectsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update certain properties of the Network To NetworkInterconnects resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkToNetworkInterconnectsClient) update(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnectPatch, options *NetworkToNetworkInterconnectsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *NetworkToNetworkInterconnectsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body NetworkToNetworkInterconnectPatch, _ *NetworkToNetworkInterconnectsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateAdministrativeState - Updates the Admin State.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - body - Request payload.
//   - options - NetworkToNetworkInterconnectsClientBeginUpdateAdministrativeStateOptions contains the optional parameters for
//     the NetworkToNetworkInterconnectsClient.BeginUpdateAdministrativeState method.
func (client *NetworkToNetworkInterconnectsClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, options *NetworkToNetworkInterconnectsClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateAdministrativeState - Updates the Admin State.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkToNetworkInterconnectsClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, options *NetworkToNetworkInterconnectsClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.BeginUpdateAdministrativeState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
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

// updateAdministrativeStateCreateRequest creates the UpdateAdministrativeState request.
func (client *NetworkToNetworkInterconnectsClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, _ *NetworkToNetworkInterconnectsClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}/updateAdministrativeState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateNpbStaticRouteBfdAdministrativeState - Updates the NPB Static Route BFD Administrative State.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkFabricName - Name of the Network Fabric.
//   - networkToNetworkInterconnectName - Name of the Network to Network Interconnect.
//   - body - Request payload.
//   - options - NetworkToNetworkInterconnectsClientBeginUpdateNpbStaticRouteBfdAdministrativeStateOptions contains the optional
//     parameters for the NetworkToNetworkInterconnectsClient.BeginUpdateNpbStaticRouteBfdAdministrativeState method.
func (client *NetworkToNetworkInterconnectsClient) BeginUpdateNpbStaticRouteBfdAdministrativeState(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, options *NetworkToNetworkInterconnectsClientBeginUpdateNpbStaticRouteBfdAdministrativeStateOptions) (*runtime.Poller[NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateNpbStaticRouteBfdAdministrativeState(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateNpbStaticRouteBfdAdministrativeState - Updates the NPB Static Route BFD Administrative State.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkToNetworkInterconnectsClient) updateNpbStaticRouteBfdAdministrativeState(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, options *NetworkToNetworkInterconnectsClientBeginUpdateNpbStaticRouteBfdAdministrativeStateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkToNetworkInterconnectsClient.BeginUpdateNpbStaticRouteBfdAdministrativeState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateNpbStaticRouteBfdAdministrativeStateCreateRequest(ctx, resourceGroupName, networkFabricName, networkToNetworkInterconnectName, body, options)
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

// updateNpbStaticRouteBfdAdministrativeStateCreateRequest creates the UpdateNpbStaticRouteBfdAdministrativeState request.
func (client *NetworkToNetworkInterconnectsClient) updateNpbStaticRouteBfdAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body UpdateAdministrativeState, _ *NetworkToNetworkInterconnectsClientBeginUpdateNpbStaticRouteBfdAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkFabrics/{networkFabricName}/networkToNetworkInterconnects/{networkToNetworkInterconnectName}/updateNpbStaticRouteBfdAdministrativeState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkFabricName == "" {
		return nil, errors.New("parameter networkFabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFabricName}", url.PathEscape(networkFabricName))
	if networkToNetworkInterconnectName == "" {
		return nil, errors.New("parameter networkToNetworkInterconnectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkToNetworkInterconnectName}", url.PathEscape(networkToNetworkInterconnectName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
