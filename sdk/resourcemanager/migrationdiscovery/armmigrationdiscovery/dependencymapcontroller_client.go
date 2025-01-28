//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationdiscovery

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

// DependencyMapControllerClient contains the methods for the DependencyMapController group.
// Don't use this type directly, use NewDependencyMapControllerClient() instead.
type DependencyMapControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDependencyMapControllerClient creates a new instance of DependencyMapControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDependencyMapControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DependencyMapControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DependencyMapControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginClientGroupMembers - API to list client group members for the selected client group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - mapRequest - The content of the action request
//   - options - DependencyMapControllerClientBeginClientGroupMembersOptions contains the optional parameters for the DependencyMapControllerClient.BeginClientGroupMembers
//     method.
func (client *DependencyMapControllerClient) BeginClientGroupMembers(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsClientGroupMembersRequest, options *DependencyMapControllerClientBeginClientGroupMembersOptions) (*runtime.Poller[DependencyMapControllerClientClientGroupMembersResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.clientGroupMembers(ctx, resourceGroupName, siteName, mapRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DependencyMapControllerClientClientGroupMembersResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DependencyMapControllerClientClientGroupMembersResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ClientGroupMembers - API to list client group members for the selected client group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *DependencyMapControllerClient) clientGroupMembers(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsClientGroupMembersRequest, options *DependencyMapControllerClientBeginClientGroupMembersOptions) (*http.Response, error) {
	var err error
	const operationName = "DependencyMapControllerClient.BeginClientGroupMembers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.clientGroupMembersCreateRequest(ctx, resourceGroupName, siteName, mapRequest, options)
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

// clientGroupMembersCreateRequest creates the ClientGroupMembers request.
func (client *DependencyMapControllerClient) clientGroupMembersCreateRequest(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsClientGroupMembersRequest, options *DependencyMapControllerClientBeginClientGroupMembersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/clientGroupMembers"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, mapRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginExportDependencies - API to generate report containing agentless dependencies.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - requestBody - The content of the action request
//   - options - DependencyMapControllerClientBeginExportDependenciesOptions contains the optional parameters for the DependencyMapControllerClient.BeginExportDependencies
//     method.
func (client *DependencyMapControllerClient) BeginExportDependencies(ctx context.Context, resourceGroupName string, siteName string, requestBody DependencyMapServiceMapextensionsExportDependenciesRequest, options *DependencyMapControllerClientBeginExportDependenciesOptions) (*runtime.Poller[DependencyMapControllerClientExportDependenciesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportDependencies(ctx, resourceGroupName, siteName, requestBody, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DependencyMapControllerClientExportDependenciesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DependencyMapControllerClientExportDependenciesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportDependencies - API to generate report containing agentless dependencies.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *DependencyMapControllerClient) exportDependencies(ctx context.Context, resourceGroupName string, siteName string, requestBody DependencyMapServiceMapextensionsExportDependenciesRequest, options *DependencyMapControllerClientBeginExportDependenciesOptions) (*http.Response, error) {
	var err error
	const operationName = "DependencyMapControllerClient.BeginExportDependencies"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportDependenciesCreateRequest(ctx, resourceGroupName, siteName, requestBody, options)
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

// exportDependenciesCreateRequest creates the ExportDependencies request.
func (client *DependencyMapControllerClient) exportDependenciesCreateRequest(ctx context.Context, resourceGroupName string, siteName string, requestBody DependencyMapServiceMapextensionsExportDependenciesRequest, options *DependencyMapControllerClientBeginExportDependenciesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/exportDependencies"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginGenerateCoarseMap - API to generate coarse map for the list of machines.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - mapRequest - The content of the action request
//   - options - DependencyMapControllerClientBeginGenerateCoarseMapOptions contains the optional parameters for the DependencyMapControllerClient.BeginGenerateCoarseMap
//     method.
func (client *DependencyMapControllerClient) BeginGenerateCoarseMap(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsScopeMapRequest, options *DependencyMapControllerClientBeginGenerateCoarseMapOptions) (*runtime.Poller[DependencyMapControllerClientGenerateCoarseMapResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateCoarseMap(ctx, resourceGroupName, siteName, mapRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DependencyMapControllerClientGenerateCoarseMapResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DependencyMapControllerClientGenerateCoarseMapResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GenerateCoarseMap - API to generate coarse map for the list of machines.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *DependencyMapControllerClient) generateCoarseMap(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsScopeMapRequest, options *DependencyMapControllerClientBeginGenerateCoarseMapOptions) (*http.Response, error) {
	var err error
	const operationName = "DependencyMapControllerClient.BeginGenerateCoarseMap"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateCoarseMapCreateRequest(ctx, resourceGroupName, siteName, mapRequest, options)
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

// generateCoarseMapCreateRequest creates the GenerateCoarseMap request.
func (client *DependencyMapControllerClient) generateCoarseMapCreateRequest(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsScopeMapRequest, options *DependencyMapControllerClientBeginGenerateCoarseMapOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/generateCoarseMap"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, mapRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginGenerateDetailedMap - API to generate detailed map for a selected machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - mapRequest - The content of the action request
//   - options - DependencyMapControllerClientBeginGenerateDetailedMapOptions contains the optional parameters for the DependencyMapControllerClient.BeginGenerateDetailedMap
//     method.
func (client *DependencyMapControllerClient) BeginGenerateDetailedMap(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *DependencyMapControllerClientBeginGenerateDetailedMapOptions) (*runtime.Poller[DependencyMapControllerClientGenerateDetailedMapResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateDetailedMap(ctx, resourceGroupName, siteName, mapRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DependencyMapControllerClientGenerateDetailedMapResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DependencyMapControllerClientGenerateDetailedMapResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GenerateDetailedMap - API to generate detailed map for a selected machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *DependencyMapControllerClient) generateDetailedMap(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *DependencyMapControllerClientBeginGenerateDetailedMapOptions) (*http.Response, error) {
	var err error
	const operationName = "DependencyMapControllerClient.BeginGenerateDetailedMap"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateDetailedMapCreateRequest(ctx, resourceGroupName, siteName, mapRequest, options)
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

// generateDetailedMapCreateRequest creates the GenerateDetailedMap request.
func (client *DependencyMapControllerClient) generateDetailedMapCreateRequest(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsSingleMachineDetailedMapRequest, options *DependencyMapControllerClientBeginGenerateDetailedMapOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/generateDetailedMap"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, mapRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginServerGroupMembers - API to list server group members for the selected server group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - mapRequest - The content of the action request
//   - options - DependencyMapControllerClientBeginServerGroupMembersOptions contains the optional parameters for the DependencyMapControllerClient.BeginServerGroupMembers
//     method.
func (client *DependencyMapControllerClient) BeginServerGroupMembers(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsServerGroupMembersRequest, options *DependencyMapControllerClientBeginServerGroupMembersOptions) (*runtime.Poller[DependencyMapControllerClientServerGroupMembersResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.serverGroupMembers(ctx, resourceGroupName, siteName, mapRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DependencyMapControllerClientServerGroupMembersResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DependencyMapControllerClientServerGroupMembersResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ServerGroupMembers - API to list server group members for the selected server group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *DependencyMapControllerClient) serverGroupMembers(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsServerGroupMembersRequest, options *DependencyMapControllerClientBeginServerGroupMembersOptions) (*http.Response, error) {
	var err error
	const operationName = "DependencyMapControllerClient.BeginServerGroupMembers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.serverGroupMembersCreateRequest(ctx, resourceGroupName, siteName, mapRequest, options)
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

// serverGroupMembersCreateRequest creates the ServerGroupMembers request.
func (client *DependencyMapControllerClient) serverGroupMembersCreateRequest(ctx context.Context, resourceGroupName string, siteName string, mapRequest DependencyMapServiceMapextensionsServerGroupMembersRequest, options *DependencyMapControllerClientBeginServerGroupMembersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/serverGroupMembers"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, mapRequest); err != nil {
		return nil, err
	}
	return req, nil
}
