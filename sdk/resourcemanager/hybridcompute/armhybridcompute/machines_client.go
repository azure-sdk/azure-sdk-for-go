//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcompute

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

// MachinesClient contains the methods for the Machines group.
// Don't use this type directly, use NewMachinesClient() instead.
type MachinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachinesClient creates a new instance of MachinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachinesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAssessPatches - The operation to assess patches on a hybrid machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the hybrid machine.
//   - options - MachinesClientBeginAssessPatchesOptions contains the optional parameters for the MachinesClient.BeginAssessPatches
//     method.
func (client *MachinesClient) BeginAssessPatches(ctx context.Context, resourceGroupName string, name string, options *MachinesClientBeginAssessPatchesOptions) (*runtime.Poller[MachinesClientAssessPatchesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.assessPatches(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachinesClientAssessPatchesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachinesClientAssessPatchesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// AssessPatches - The operation to assess patches on a hybrid machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
func (client *MachinesClient) assessPatches(ctx context.Context, resourceGroupName string, name string, options *MachinesClientBeginAssessPatchesOptions) (*http.Response, error) {
	var err error
	const operationName = "MachinesClient.BeginAssessPatches"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.assessPatchesCreateRequest(ctx, resourceGroupName, name, options)
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

// assessPatchesCreateRequest creates the AssessPatches request.
func (client *MachinesClient) assessPatchesCreateRequest(ctx context.Context, resourceGroupName string, name string, options *MachinesClientBeginAssessPatchesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/assessPatches"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// CreateOrUpdate - The operation to create or update a hybrid machine. Please note some properties can be set only during
// machine creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - parameters - Parameters supplied to the Create hybrid machine operation.
//   - options - MachinesClientCreateOrUpdateOptions contains the optional parameters for the MachinesClient.CreateOrUpdate method.
func (client *MachinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, machineName string, parameters Machine, options *MachinesClientCreateOrUpdateOptions) (MachinesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "MachinesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, machineName, parameters, options)
	if err != nil {
		return MachinesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MachinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, parameters Machine, options *MachinesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MachinesClient) createOrUpdateHandleResponse(resp *http.Response) (MachinesClientCreateOrUpdateResponse, error) {
	result := MachinesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Machine); err != nil {
		return MachinesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation to delete a hybrid machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - MachinesClientDeleteOptions contains the optional parameters for the MachinesClient.Delete method.
func (client *MachinesClient) Delete(ctx context.Context, resourceGroupName string, machineName string, options *MachinesClientDeleteOptions) (MachinesClientDeleteResponse, error) {
	var err error
	const operationName = "MachinesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, machineName, options)
	if err != nil {
		return MachinesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MachinesClientDeleteResponse{}, err
	}
	return MachinesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MachinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachinesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about the model view or the instance view of a hybrid machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - MachinesClientGetOptions contains the optional parameters for the MachinesClient.Get method.
func (client *MachinesClient) Get(ctx context.Context, resourceGroupName string, machineName string, options *MachinesClientGetOptions) (MachinesClientGetResponse, error) {
	var err error
	const operationName = "MachinesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, options)
	if err != nil {
		return MachinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachinesClient) getHandleResponse(resp *http.Response) (MachinesClientGetResponse, error) {
	result := MachinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Machine); err != nil {
		return MachinesClientGetResponse{}, err
	}
	return result, nil
}

// BeginInstallPatches - The operation to install patches on a hybrid machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group.
//   - name - The name of the hybrid machine.
//   - installPatchesInput - Input for InstallPatches as directly received by the API
//   - options - MachinesClientBeginInstallPatchesOptions contains the optional parameters for the MachinesClient.BeginInstallPatches
//     method.
func (client *MachinesClient) BeginInstallPatches(ctx context.Context, resourceGroupName string, name string, installPatchesInput MachineInstallPatchesParameters, options *MachinesClientBeginInstallPatchesOptions) (*runtime.Poller[MachinesClientInstallPatchesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.installPatches(ctx, resourceGroupName, name, installPatchesInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachinesClientInstallPatchesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MachinesClientInstallPatchesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// InstallPatches - The operation to install patches on a hybrid machine identity in Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
func (client *MachinesClient) installPatches(ctx context.Context, resourceGroupName string, name string, installPatchesInput MachineInstallPatchesParameters, options *MachinesClientBeginInstallPatchesOptions) (*http.Response, error) {
	var err error
	const operationName = "MachinesClient.BeginInstallPatches"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.installPatchesCreateRequest(ctx, resourceGroupName, name, installPatchesInput, options)
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

// installPatchesCreateRequest creates the InstallPatches request.
func (client *MachinesClient) installPatchesCreateRequest(ctx context.Context, resourceGroupName string, name string, installPatchesInput MachineInstallPatchesParameters, options *MachinesClientBeginInstallPatchesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/installPatches"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, installPatchesInput); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListByResourceGroupPager - Lists all the hybrid machines in the specified resource group. Use the nextLink property
// in the response to get the next page of hybrid machines.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MachinesClientListByResourceGroupOptions contains the optional parameters for the MachinesClient.NewListByResourceGroupPager
//     method.
func (client *MachinesClient) NewListByResourceGroupPager(resourceGroupName string, options *MachinesClientListByResourceGroupOptions) *runtime.Pager[MachinesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachinesClientListByResourceGroupResponse]{
		More: func(page MachinesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachinesClientListByResourceGroupResponse) (MachinesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachinesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MachinesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MachinesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MachinesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines"
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
	reqQP.Set("api-version", "2023-10-03-preview")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MachinesClient) listByResourceGroupHandleResponse(resp *http.Response) (MachinesClientListByResourceGroupResponse, error) {
	result := MachinesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineListResult); err != nil {
		return MachinesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the hybrid machines in the specified subscription. Use the nextLink property in
// the response to get the next page of hybrid machines.
//
// Generated from API version 2023-10-03-preview
//   - options - MachinesClientListBySubscriptionOptions contains the optional parameters for the MachinesClient.NewListBySubscriptionPager
//     method.
func (client *MachinesClient) NewListBySubscriptionPager(options *MachinesClientListBySubscriptionOptions) *runtime.Pager[MachinesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachinesClientListBySubscriptionResponse]{
		More: func(page MachinesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachinesClientListBySubscriptionResponse) (MachinesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachinesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MachinesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MachinesClient) listBySubscriptionCreateRequest(ctx context.Context, options *MachinesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/machines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MachinesClient) listBySubscriptionHandleResponse(resp *http.Response) (MachinesClientListBySubscriptionResponse, error) {
	result := MachinesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineListResult); err != nil {
		return MachinesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - The operation to update a hybrid machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-03-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - parameters - Parameters supplied to the Update hybrid machine operation.
//   - options - MachinesClientUpdateOptions contains the optional parameters for the MachinesClient.Update method.
func (client *MachinesClient) Update(ctx context.Context, resourceGroupName string, machineName string, parameters MachineUpdate, options *MachinesClientUpdateOptions) (MachinesClientUpdateResponse, error) {
	var err error
	const operationName = "MachinesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, machineName, parameters, options)
	if err != nil {
		return MachinesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MachinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, parameters MachineUpdate, options *MachinesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-03-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MachinesClient) updateHandleResponse(resp *http.Response) (MachinesClientUpdateResponse, error) {
	result := MachinesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Machine); err != nil {
		return MachinesClientUpdateResponse{}, err
	}
	return result, nil
}
