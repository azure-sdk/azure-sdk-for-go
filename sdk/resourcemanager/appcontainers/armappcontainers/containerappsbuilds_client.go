//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ContainerAppsBuildsClient contains the methods for the ContainerAppsBuilds group.
// Don't use this type directly, use NewContainerAppsBuildsClient() instead.
type ContainerAppsBuildsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsBuildsClient creates a new instance of ContainerAppsBuildsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsBuildsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsBuildsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsBuildsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete a Container Apps Build resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Build is associated.
//   - buildName - The name of a build.
//   - options - ContainerAppsBuildsClientBeginDeleteOptions contains the optional parameters for the ContainerAppsBuildsClient.BeginDelete
//     method.
func (client *ContainerAppsBuildsClient) BeginDelete(ctx context.Context, resourceGroupName string, containerAppName string, buildName string, options *ContainerAppsBuildsClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsBuildsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerAppName, buildName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerAppsBuildsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerAppsBuildsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Container Apps Build resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *ContainerAppsBuildsClient) deleteOperation(ctx context.Context, resourceGroupName string, containerAppName string, buildName string, options *ContainerAppsBuildsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerAppsBuildsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, buildName, options)
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
func (client *ContainerAppsBuildsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, buildName string, options *ContainerAppsBuildsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/builds/{buildName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Container Apps Build resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App the Build is associated.
//   - buildName - The name of a build.
//   - options - ContainerAppsBuildsClientGetOptions contains the optional parameters for the ContainerAppsBuildsClient.Get method.
func (client *ContainerAppsBuildsClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, buildName string, options *ContainerAppsBuildsClientGetOptions) (ContainerAppsBuildsClientGetResponse, error) {
	var err error
	const operationName = "ContainerAppsBuildsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, buildName, options)
	if err != nil {
		return ContainerAppsBuildsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsBuildsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsBuildsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsBuildsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, buildName string, options *ContainerAppsBuildsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/builds/{buildName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if buildName == "" {
		return nil, errors.New("parameter buildName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildName}", url.PathEscape(buildName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsBuildsClient) getHandleResponse(resp *http.Response) (ContainerAppsBuildsClientGetResponse, error) {
	result := ContainerAppsBuildsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppsBuildResource); err != nil {
		return ContainerAppsBuildsClientGetResponse{}, err
	}
	return result, nil
}
