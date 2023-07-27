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

// ContainerAppsSourceControlsClient contains the methods for the ContainerAppsSourceControls group.
// Don't use this type directly, use NewContainerAppsSourceControlsClient() instead.
type ContainerAppsSourceControlsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerAppsSourceControlsClient creates a new instance of ContainerAppsSourceControlsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerAppsSourceControlsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsSourceControlsClient, error) {
	cl, err := arm.NewClient(moduleName+".ContainerAppsSourceControlsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerAppsSourceControlsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the SourceControl for a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - sourceControlName - Name of the Container App SourceControl.
//   - sourceControlEnvelope - Properties used to create a Container App SourceControl
//   - options - ContainerAppsSourceControlsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerAppsSourceControlsClient.BeginCreateOrUpdate
//     method.
func (client *ContainerAppsSourceControlsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, sourceControlEnvelope SourceControl, options *ContainerAppsSourceControlsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerAppsSourceControlsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, containerAppName, sourceControlName, sourceControlEnvelope, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ContainerAppsSourceControlsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsSourceControlsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update the SourceControl for a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *ContainerAppsSourceControlsClient) createOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, sourceControlEnvelope SourceControl, options *ContainerAppsSourceControlsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerAppName, sourceControlName, sourceControlEnvelope, options)
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
func (client *ContainerAppsSourceControlsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, sourceControlEnvelope SourceControl, options *ContainerAppsSourceControlsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/sourcecontrols/{sourceControlName}"
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
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sourceControlEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Container App SourceControl.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - sourceControlName - Name of the Container App SourceControl.
//   - options - ContainerAppsSourceControlsClientBeginDeleteOptions contains the optional parameters for the ContainerAppsSourceControlsClient.BeginDelete
//     method.
func (client *ContainerAppsSourceControlsClient) BeginDelete(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, options *ContainerAppsSourceControlsClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsSourceControlsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerAppName, sourceControlName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ContainerAppsSourceControlsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsSourceControlsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a Container App SourceControl.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *ContainerAppsSourceControlsClient) deleteOperation(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, options *ContainerAppsSourceControlsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, sourceControlName, options)
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
func (client *ContainerAppsSourceControlsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, options *ContainerAppsSourceControlsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/sourcecontrols/{sourceControlName}"
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
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a SourceControl of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - sourceControlName - Name of the Container App SourceControl.
//   - options - ContainerAppsSourceControlsClientGetOptions contains the optional parameters for the ContainerAppsSourceControlsClient.Get
//     method.
func (client *ContainerAppsSourceControlsClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, options *ContainerAppsSourceControlsClientGetOptions) (ContainerAppsSourceControlsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, sourceControlName, options)
	if err != nil {
		return ContainerAppsSourceControlsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerAppsSourceControlsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerAppsSourceControlsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsSourceControlsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, sourceControlName string, options *ContainerAppsSourceControlsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/sourcecontrols/{sourceControlName}"
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
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsSourceControlsClient) getHandleResponse(resp *http.Response) (ContainerAppsSourceControlsClientGetResponse, error) {
	result := ContainerAppsSourceControlsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControl); err != nil {
		return ContainerAppsSourceControlsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByContainerAppPager - Get the Container App SourceControls in a given resource group.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - options - ContainerAppsSourceControlsClientListByContainerAppOptions contains the optional parameters for the ContainerAppsSourceControlsClient.NewListByContainerAppPager
//     method.
func (client *ContainerAppsSourceControlsClient) NewListByContainerAppPager(resourceGroupName string, containerAppName string, options *ContainerAppsSourceControlsClientListByContainerAppOptions) *runtime.Pager[ContainerAppsSourceControlsClientListByContainerAppResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsSourceControlsClientListByContainerAppResponse]{
		More: func(page ContainerAppsSourceControlsClientListByContainerAppResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsSourceControlsClientListByContainerAppResponse) (ContainerAppsSourceControlsClientListByContainerAppResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByContainerAppCreateRequest(ctx, resourceGroupName, containerAppName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsSourceControlsClientListByContainerAppResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ContainerAppsSourceControlsClientListByContainerAppResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsSourceControlsClientListByContainerAppResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByContainerAppHandleResponse(resp)
		},
	})
}

// listByContainerAppCreateRequest creates the ListByContainerApp request.
func (client *ContainerAppsSourceControlsClient) listByContainerAppCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, options *ContainerAppsSourceControlsClientListByContainerAppOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/sourcecontrols"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByContainerAppHandleResponse handles the ListByContainerApp response.
func (client *ContainerAppsSourceControlsClient) listByContainerAppHandleResponse(resp *http.Response) (ContainerAppsSourceControlsClientListByContainerAppResponse, error) {
	result := ContainerAppsSourceControlsClientListByContainerAppResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlCollection); err != nil {
		return ContainerAppsSourceControlsClientListByContainerAppResponse{}, err
	}
	return result, nil
}
