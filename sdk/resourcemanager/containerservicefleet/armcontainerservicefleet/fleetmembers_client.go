//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

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

// FleetMembersClient contains the methods for the FleetMembers group.
// Don't use this type directly, use NewFleetMembersClient() instead.
type FleetMembersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFleetMembersClient creates a new instance of FleetMembersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFleetMembersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FleetMembersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FleetMembersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - fleetMemberName - The name of the Fleet member resource.
//   - resource - Resource create parameters.
//   - options - FleetMembersClientBeginCreateOptions contains the optional parameters for the FleetMembersClient.BeginCreate
//     method.
func (client *FleetMembersClient) BeginCreate(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, resource FleetMember, options *FleetMembersClientBeginCreateOptions) (*runtime.Poller[FleetMembersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, fleetName, fleetMemberName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetMembersClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetMembersClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *FleetMembersClient) create(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, resource FleetMember, options *FleetMembersClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetMembersClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, resource, options)
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
func (client *FleetMembersClient) createCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, resource FleetMember, options *FleetMembersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - fleetMemberName - The name of the Fleet member resource.
//   - options - FleetMembersClientBeginDeleteOptions contains the optional parameters for the FleetMembersClient.BeginDelete
//     method.
func (client *FleetMembersClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*runtime.Poller[FleetMembersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, fleetMemberName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetMembersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetMembersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *FleetMembersClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetMembersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, options)
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
func (client *FleetMembersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// Get - Get a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - fleetMemberName - The name of the Fleet member resource.
//   - options - FleetMembersClientGetOptions contains the optional parameters for the FleetMembersClient.Get method.
func (client *FleetMembersClient) Get(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientGetOptions) (FleetMembersClientGetResponse, error) {
	var err error
	const operationName = "FleetMembersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, options)
	if err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FleetMembersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FleetMembersClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, options *FleetMembersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FleetMembersClient) getHandleResponse(resp *http.Response) (FleetMembersClientGetResponse, error) {
	result := FleetMembersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetMember); err != nil {
		return FleetMembersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFleetPager - List FleetMember resources by Fleet
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - options - FleetMembersClientListByFleetOptions contains the optional parameters for the FleetMembersClient.NewListByFleetPager
//     method.
func (client *FleetMembersClient) NewListByFleetPager(resourceGroupName string, fleetName string, options *FleetMembersClientListByFleetOptions) *runtime.Pager[FleetMembersClientListByFleetResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetMembersClientListByFleetResponse]{
		More: func(page FleetMembersClientListByFleetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetMembersClientListByFleetResponse) (FleetMembersClientListByFleetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FleetMembersClient.NewListByFleetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFleetCreateRequest(ctx, resourceGroupName, fleetName, options)
			}, nil)
			if err != nil {
				return FleetMembersClientListByFleetResponse{}, err
			}
			return client.listByFleetHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFleetCreateRequest creates the ListByFleet request.
func (client *FleetMembersClient) listByFleetCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetMembersClientListByFleetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFleetHandleResponse handles the ListByFleet response.
func (client *FleetMembersClient) listByFleetHandleResponse(resp *http.Response) (FleetMembersClientListByFleetResponse, error) {
	result := FleetMembersClientListByFleetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetMemberListResult); err != nil {
		return FleetMembersClientListByFleetResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - fleetMemberName - The name of the Fleet member resource.
//   - properties - The resource properties to be updated.
//   - options - FleetMembersClientBeginUpdateOptions contains the optional parameters for the FleetMembersClient.BeginUpdate
//     method.
func (client *FleetMembersClient) BeginUpdate(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, properties FleetMemberUpdate, options *FleetMembersClientBeginUpdateOptions) (*runtime.Poller[FleetMembersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, fleetName, fleetMemberName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetMembersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetMembersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a FleetMember
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *FleetMembersClient) update(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, properties FleetMemberUpdate, options *FleetMembersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetMembersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, fleetName, fleetMemberName, properties, options)
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
func (client *FleetMembersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, fleetMemberName string, properties FleetMemberUpdate, options *FleetMembersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/members/{fleetMemberName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if fleetMemberName == "" {
		return nil, errors.New("parameter fleetMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetMemberName}", url.PathEscape(fleetMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
