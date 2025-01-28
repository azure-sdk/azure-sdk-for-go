//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// DevOpsConfigurationsClient contains the methods for the DevOpsConfigurations group.
// Don't use this type directly, use NewDevOpsConfigurationsClient() instead.
type DevOpsConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDevOpsConfigurationsClient creates a new instance of DevOpsConfigurationsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDevOpsConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevOpsConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DevOpsConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DevOps Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - devOpsConfiguration - The DevOps configuration resource payload.
//   - options - DevOpsConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the DevOpsConfigurationsClient.BeginCreateOrUpdate
//     method.
func (client *DevOpsConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevOpsConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, securityConnectorName, devOpsConfiguration, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevOpsConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevOpsConfigurationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a DevOps Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
func (client *DevOpsConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DevOpsConfigurationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityConnectorName, devOpsConfiguration, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DevOpsConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, devOpsConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a DevOps Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - DevOpsConfigurationsClientBeginDeleteOptions contains the optional parameters for the DevOpsConfigurationsClient.BeginDelete
//     method.
func (client *DevOpsConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientBeginDeleteOptions) (*runtime.Poller[DevOpsConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, securityConnectorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevOpsConfigurationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevOpsConfigurationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a DevOps Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
func (client *DevOpsConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DevOpsConfigurationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
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
func (client *DevOpsConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a DevOps Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - DevOpsConfigurationsClientGetOptions contains the optional parameters for the DevOpsConfigurationsClient.Get
//     method.
func (client *DevOpsConfigurationsClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientGetOptions) (DevOpsConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "DevOpsConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
	if err != nil {
		return DevOpsConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevOpsConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DevOpsConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DevOpsConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevOpsConfigurationsClient) getHandleResponse(resp *http.Response) (DevOpsConfigurationsClientGetResponse, error) {
	result := DevOpsConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevOpsConfiguration); err != nil {
		return DevOpsConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List DevOps Configurations.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - DevOpsConfigurationsClientListOptions contains the optional parameters for the DevOpsConfigurationsClient.NewListPager
//     method.
func (client *DevOpsConfigurationsClient) NewListPager(resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientListOptions) *runtime.Pager[DevOpsConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevOpsConfigurationsClientListResponse]{
		More: func(page DevOpsConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevOpsConfigurationsClientListResponse) (DevOpsConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DevOpsConfigurationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
			}, nil)
			if err != nil {
				return DevOpsConfigurationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DevOpsConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *DevOpsConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DevOpsConfigurationsClient) listHandleResponse(resp *http.Response) (DevOpsConfigurationsClientListResponse, error) {
	result := DevOpsConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevOpsConfigurationListResponse); err != nil {
		return DevOpsConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a DevOps Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - devOpsConfiguration - The DevOps configuration resource payload.
//   - options - DevOpsConfigurationsClientBeginUpdateOptions contains the optional parameters for the DevOpsConfigurationsClient.BeginUpdate
//     method.
func (client *DevOpsConfigurationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginUpdateOptions) (*runtime.Poller[DevOpsConfigurationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, securityConnectorName, devOpsConfiguration, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevOpsConfigurationsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevOpsConfigurationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a DevOps Configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
func (client *DevOpsConfigurationsClient) update(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DevOpsConfigurationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, securityConnectorName, devOpsConfiguration, options)
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
func (client *DevOpsConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, devOpsConfiguration DevOpsConfiguration, options *DevOpsConfigurationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, devOpsConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}
