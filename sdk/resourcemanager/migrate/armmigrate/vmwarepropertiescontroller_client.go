//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// VmwarePropertiesControllerClient contains the methods for the VmwarePropertiesController group.
// Don't use this type directly, use NewVmwarePropertiesControllerClient() instead.
type VmwarePropertiesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVmwarePropertiesControllerClient creates a new instance of VmwarePropertiesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVmwarePropertiesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VmwarePropertiesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VmwarePropertiesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginUpdateDependencyMapStatus - Method to enable disable dependency map status for machines in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - updateMachineDepMapStatus - The content of the action request
//   - options - VmwarePropertiesControllerClientBeginUpdateDependencyMapStatusOptions contains the optional parameters for the
//     VmwarePropertiesControllerClient.BeginUpdateDependencyMapStatus method.
func (client *VmwarePropertiesControllerClient) BeginUpdateDependencyMapStatus(ctx context.Context, resourceGroupName string, siteName string, updateMachineDepMapStatus UpdateMachineDepMapStatus, options *VmwarePropertiesControllerClientBeginUpdateDependencyMapStatusOptions) (*runtime.Poller[VmwarePropertiesControllerClientUpdateDependencyMapStatusResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateDependencyMapStatus(ctx, resourceGroupName, siteName, updateMachineDepMapStatus, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmwarePropertiesControllerClientUpdateDependencyMapStatusResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmwarePropertiesControllerClientUpdateDependencyMapStatusResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateDependencyMapStatus - Method to enable disable dependency map status for machines in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *VmwarePropertiesControllerClient) updateDependencyMapStatus(ctx context.Context, resourceGroupName string, siteName string, updateMachineDepMapStatus UpdateMachineDepMapStatus, options *VmwarePropertiesControllerClientBeginUpdateDependencyMapStatusOptions) (*http.Response, error) {
	var err error
	const operationName = "VmwarePropertiesControllerClient.BeginUpdateDependencyMapStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateDependencyMapStatusCreateRequest(ctx, resourceGroupName, siteName, updateMachineDepMapStatus, options)
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

// updateDependencyMapStatusCreateRequest creates the UpdateDependencyMapStatus request.
func (client *VmwarePropertiesControllerClient) updateDependencyMapStatusCreateRequest(ctx context.Context, resourceGroupName string, siteName string, updateMachineDepMapStatus UpdateMachineDepMapStatus, options *VmwarePropertiesControllerClientBeginUpdateDependencyMapStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/updateDependencyMapStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateMachineDepMapStatus); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateProperties - Method to update properties for machines in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - metaData - The content of the action request
//   - options - VmwarePropertiesControllerClientBeginUpdatePropertiesOptions contains the optional parameters for the VmwarePropertiesControllerClient.BeginUpdateProperties
//     method.
func (client *VmwarePropertiesControllerClient) BeginUpdateProperties(ctx context.Context, resourceGroupName string, siteName string, metaData MachineMetadataCollection, options *VmwarePropertiesControllerClientBeginUpdatePropertiesOptions) (*runtime.Poller[VmwarePropertiesControllerClientUpdatePropertiesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateProperties(ctx, resourceGroupName, siteName, metaData, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmwarePropertiesControllerClientUpdatePropertiesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmwarePropertiesControllerClientUpdatePropertiesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateProperties - Method to update properties for machines in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *VmwarePropertiesControllerClient) updateProperties(ctx context.Context, resourceGroupName string, siteName string, metaData MachineMetadataCollection, options *VmwarePropertiesControllerClientBeginUpdatePropertiesOptions) (*http.Response, error) {
	var err error
	const operationName = "VmwarePropertiesControllerClient.BeginUpdateProperties"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updatePropertiesCreateRequest(ctx, resourceGroupName, siteName, metaData, options)
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

// updatePropertiesCreateRequest creates the UpdateProperties request.
func (client *VmwarePropertiesControllerClient) updatePropertiesCreateRequest(ctx context.Context, resourceGroupName string, siteName string, metaData MachineMetadataCollection, options *VmwarePropertiesControllerClientBeginUpdatePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/updateProperties"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, metaData); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateRunAsAccount - Method to associate Run as account to machine in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - updateMachineRunAsAccount - The content of the action request
//   - options - VmwarePropertiesControllerClientBeginUpdateRunAsAccountOptions contains the optional parameters for the VmwarePropertiesControllerClient.BeginUpdateRunAsAccount
//     method.
func (client *VmwarePropertiesControllerClient) BeginUpdateRunAsAccount(ctx context.Context, resourceGroupName string, siteName string, updateMachineRunAsAccount UpdateMachineRunAsAccount, options *VmwarePropertiesControllerClientBeginUpdateRunAsAccountOptions) (*runtime.Poller[VmwarePropertiesControllerClientUpdateRunAsAccountResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateRunAsAccount(ctx, resourceGroupName, siteName, updateMachineRunAsAccount, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmwarePropertiesControllerClientUpdateRunAsAccountResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmwarePropertiesControllerClientUpdateRunAsAccountResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateRunAsAccount - Method to associate Run as account to machine in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *VmwarePropertiesControllerClient) updateRunAsAccount(ctx context.Context, resourceGroupName string, siteName string, updateMachineRunAsAccount UpdateMachineRunAsAccount, options *VmwarePropertiesControllerClientBeginUpdateRunAsAccountOptions) (*http.Response, error) {
	var err error
	const operationName = "VmwarePropertiesControllerClient.BeginUpdateRunAsAccount"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateRunAsAccountCreateRequest(ctx, resourceGroupName, siteName, updateMachineRunAsAccount, options)
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

// updateRunAsAccountCreateRequest creates the UpdateRunAsAccount request.
func (client *VmwarePropertiesControllerClient) updateRunAsAccountCreateRequest(ctx context.Context, resourceGroupName string, siteName string, updateMachineRunAsAccount UpdateMachineRunAsAccount, options *VmwarePropertiesControllerClientBeginUpdateRunAsAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/updateRunAsAccount"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateMachineRunAsAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateTags - Method to associate Run as account to machine in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - updateMachineTags - The content of the action request
//   - options - VmwarePropertiesControllerClientBeginUpdateTagsOptions contains the optional parameters for the VmwarePropertiesControllerClient.BeginUpdateTags
//     method.
func (client *VmwarePropertiesControllerClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, siteName string, updateMachineTags UpdateMachineTags, options *VmwarePropertiesControllerClientBeginUpdateTagsOptions) (*runtime.Poller[VmwarePropertiesControllerClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, siteName, updateMachineTags, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VmwarePropertiesControllerClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VmwarePropertiesControllerClientUpdateTagsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateTags - Method to associate Run as account to machine in a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *VmwarePropertiesControllerClient) updateTags(ctx context.Context, resourceGroupName string, siteName string, updateMachineTags UpdateMachineTags, options *VmwarePropertiesControllerClientBeginUpdateTagsOptions) (*http.Response, error) {
	var err error
	const operationName = "VmwarePropertiesControllerClient.BeginUpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, siteName, updateMachineTags, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VmwarePropertiesControllerClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, siteName string, updateMachineTags UpdateMachineTags, options *VmwarePropertiesControllerClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/updateTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateMachineTags); err != nil {
		return nil, err
	}
	return req, nil
}
