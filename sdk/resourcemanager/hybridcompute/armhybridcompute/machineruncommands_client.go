//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// MachineRunCommandsClient contains the methods for the MachineRunCommands group.
// Don't use this type directly, use NewMachineRunCommandsClient() instead.
type MachineRunCommandsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachineRunCommandsClient creates a new instance of MachineRunCommandsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachineRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachineRunCommandsClient, error) {
	cl, err := arm.NewClient(moduleName+".MachineRunCommandsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachineRunCommandsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - runCommandProperties - Parameters supplied to the Create Run Command.
//   - options - MachineRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the MachineRunCommandsClient.BeginCreateOrUpdate
//     method.
func (client *MachineRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, options *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MachineRunCommandsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineRunCommandsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[MachineRunCommandsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
func (client *MachineRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, options *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
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
func (client *MachineRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommand, options *MachineRunCommandsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, runCommandProperties)
}

// BeginDelete - The operation to delete a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - options - MachineRunCommandsClientBeginDeleteOptions contains the optional parameters for the MachineRunCommandsClient.BeginDelete
//     method.
func (client *MachineRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientBeginDeleteOptions) (*runtime.Poller[MachineRunCommandsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, machineName, runCommandName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MachineRunCommandsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[MachineRunCommandsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
func (client *MachineRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, machineName, runCommandName, options)
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
func (client *MachineRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation to get a run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - options - MachineRunCommandsClientGetOptions contains the optional parameters for the MachineRunCommandsClient.Get method.
func (client *MachineRunCommandsClient) Get(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientGetOptions) (MachineRunCommandsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, runCommandName, options)
	if err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MachineRunCommandsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MachineRunCommandsClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, options *MachineRunCommandsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachineRunCommandsClient) getHandleResponse(resp *http.Response) (MachineRunCommandsClientGetResponse, error) {
	result := MachineRunCommandsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineRunCommand); err != nil {
		return MachineRunCommandsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation to get all the run commands of a non-Azure machine.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - MachineRunCommandsClientListOptions contains the optional parameters for the MachineRunCommandsClient.NewListPager
//     method.
func (client *MachineRunCommandsClient) NewListPager(resourceGroupName string, machineName string, options *MachineRunCommandsClientListOptions) *runtime.Pager[MachineRunCommandsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachineRunCommandsClientListResponse]{
		More: func(page MachineRunCommandsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachineRunCommandsClientListResponse) (MachineRunCommandsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, machineName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MachineRunCommandsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MachineRunCommandsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MachineRunCommandsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MachineRunCommandsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *MachineRunCommandsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MachineRunCommandsClient) listHandleResponse(resp *http.Response) (MachineRunCommandsClientListResponse, error) {
	result := MachineRunCommandsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachineRunCommandsListResult); err != nil {
		return MachineRunCommandsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - runCommandName - The name of the run command.
//   - runCommandProperties - Parameters supplied to the Create Run Command.
//   - options - MachineRunCommandsClientBeginUpdateOptions contains the optional parameters for the MachineRunCommandsClient.BeginUpdate
//     method.
func (client *MachineRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommandUpdate, options *MachineRunCommandsClientBeginUpdateOptions) (*runtime.Poller[MachineRunCommandsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MachineRunCommandsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[MachineRunCommandsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update the run command.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
func (client *MachineRunCommandsClient) update(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommandUpdate, options *MachineRunCommandsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, machineName, runCommandName, runCommandProperties, options)
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
func (client *MachineRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, machineName string, runCommandName string, runCommandProperties MachineRunCommandUpdate, options *MachineRunCommandsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, runCommandProperties)
}
