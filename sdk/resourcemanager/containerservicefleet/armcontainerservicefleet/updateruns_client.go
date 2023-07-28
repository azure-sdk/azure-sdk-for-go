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

// UpdateRunsClient contains the methods for the UpdateRuns group.
// Don't use this type directly, use NewUpdateRunsClient() instead.
type UpdateRunsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUpdateRunsClient creates a new instance of UpdateRunsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUpdateRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UpdateRunsClient, error) {
	cl, err := arm.NewClient(moduleName+".UpdateRunsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UpdateRunsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a UpdateRun
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateRunName - The name of the UpdateRun resource.
//   - resource - Resource create parameters.
//   - options - UpdateRunsClientBeginCreateOrUpdateOptions contains the optional parameters for the UpdateRunsClient.BeginCreateOrUpdate
//     method.
func (client *UpdateRunsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, resource UpdateRun, options *UpdateRunsClientBeginCreateOrUpdateOptions) (*runtime.Poller[UpdateRunsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, updateRunName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UpdateRunsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[UpdateRunsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a UpdateRun
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *UpdateRunsClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, resource UpdateRun, options *UpdateRunsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, updateRunName, resource, options)
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
func (client *UpdateRunsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, resource UpdateRun, options *UpdateRunsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}"
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
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a UpdateRun
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateRunName - The name of the UpdateRun resource.
//   - options - UpdateRunsClientBeginDeleteOptions contains the optional parameters for the UpdateRunsClient.BeginDelete method.
func (client *UpdateRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*runtime.Poller[UpdateRunsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, updateRunName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UpdateRunsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[UpdateRunsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a UpdateRun
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *UpdateRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, updateRunName, options)
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
func (client *UpdateRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}"
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
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a UpdateRun
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateRunName - The name of the UpdateRun resource.
//   - options - UpdateRunsClientGetOptions contains the optional parameters for the UpdateRunsClient.Get method.
func (client *UpdateRunsClient) Get(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientGetOptions) (UpdateRunsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, updateRunName, options)
	if err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UpdateRunsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UpdateRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}"
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
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UpdateRunsClient) getHandleResponse(resp *http.Response) (UpdateRunsClientGetResponse, error) {
	result := UpdateRunsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateRun); err != nil {
		return UpdateRunsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFleetPager - List UpdateRun resources by Fleet
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - options - UpdateRunsClientListByFleetOptions contains the optional parameters for the UpdateRunsClient.NewListByFleetPager
//     method.
func (client *UpdateRunsClient) NewListByFleetPager(resourceGroupName string, fleetName string, options *UpdateRunsClientListByFleetOptions) *runtime.Pager[UpdateRunsClientListByFleetResponse] {
	return runtime.NewPager(runtime.PagingHandler[UpdateRunsClientListByFleetResponse]{
		More: func(page UpdateRunsClientListByFleetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UpdateRunsClientListByFleetResponse) (UpdateRunsClientListByFleetResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFleetCreateRequest(ctx, resourceGroupName, fleetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UpdateRunsClientListByFleetResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UpdateRunsClientListByFleetResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UpdateRunsClientListByFleetResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFleetHandleResponse(resp)
		},
	})
}

// listByFleetCreateRequest creates the ListByFleet request.
func (client *UpdateRunsClient) listByFleetCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *UpdateRunsClientListByFleetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns"
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
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFleetHandleResponse handles the ListByFleet response.
func (client *UpdateRunsClient) listByFleetHandleResponse(resp *http.Response) (UpdateRunsClientListByFleetResponse, error) {
	result := UpdateRunsClientListByFleetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateRunListResult); err != nil {
		return UpdateRunsClientListByFleetResponse{}, err
	}
	return result, nil
}

// BeginStart - Starts an UpdateRun.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateRunName - The name of the UpdateRun resource.
//   - options - UpdateRunsClientBeginStartOptions contains the optional parameters for the UpdateRunsClient.BeginStart method.
func (client *UpdateRunsClient) BeginStart(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStartOptions) (*runtime.Poller[UpdateRunsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, fleetName, updateRunName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UpdateRunsClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[UpdateRunsClientStartResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Start - Starts an UpdateRun.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *UpdateRunsClient) start(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStartOptions) (*http.Response, error) {
	var err error
	req, err := client.startCreateRequest(ctx, resourceGroupName, fleetName, updateRunName, options)
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

// startCreateRequest creates the Start request.
func (client *UpdateRunsClient) startCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}/start"
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
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Stops an UpdateRun.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateRunName - The name of the UpdateRun resource.
//   - options - UpdateRunsClientBeginStopOptions contains the optional parameters for the UpdateRunsClient.BeginStop method.
func (client *UpdateRunsClient) BeginStop(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStopOptions) (*runtime.Poller[UpdateRunsClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, fleetName, updateRunName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UpdateRunsClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[UpdateRunsClientStopResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Stop - Stops an UpdateRun.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *UpdateRunsClient) stop(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStopOptions) (*http.Response, error) {
	var err error
	req, err := client.stopCreateRequest(ctx, resourceGroupName, fleetName, updateRunName, options)
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

// stopCreateRequest creates the Stop request.
func (client *UpdateRunsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateRunName string, options *UpdateRunsClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateRuns/{updateRunName}/stop"
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
	if updateRunName == "" {
		return nil, errors.New("parameter updateRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateRunName}", url.PathEscape(updateRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
