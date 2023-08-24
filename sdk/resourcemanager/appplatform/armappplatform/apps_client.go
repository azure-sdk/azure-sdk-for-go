//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// AppsClient contains the methods for the Apps group.
// Don't use this type directly, use NewAppsClient() instead.
type AppsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAppsClient creates a new instance of AppsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AppsClient, error) {
	cl, err := arm.NewClient(moduleName+".AppsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AppsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new App or update an exiting App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - appResource - Parameters for the create or update operation
//   - options - AppsClientBeginCreateOrUpdateOptions contains the optional parameters for the AppsClient.BeginCreateOrUpdate
//     method.
func (client *AppsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AppsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, appName, appResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AppsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AppsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new App or update an exiting App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *AppsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, appName, appResource, options)
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
func (client *AppsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, appResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Operation to delete an App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - options - AppsClientBeginDeleteOptions contains the optional parameters for the AppsClient.BeginDelete method.
func (client *AppsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientBeginDeleteOptions) (*runtime.Poller[AppsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, appName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AppsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AppsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Operation to delete an App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *AppsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, appName, options)
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
func (client *AppsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an App and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - options - AppsClientGetOptions contains the optional parameters for the AppsClient.Get method.
func (client *AppsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientGetOptions) (AppsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, appName, options)
	if err != nil {
		return AppsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AppsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	if options != nil && options.SyncStatus != nil {
		reqQP.Set("syncStatus", *options.SyncStatus)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AppsClient) getHandleResponse(resp *http.Response) (AppsClientGetResponse, error) {
	result := AppsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResource); err != nil {
		return AppsClientGetResponse{}, err
	}
	return result, nil
}

// GetResourceUploadURL - Get an resource upload URL for an App, which may be artifacts or source archive.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - options - AppsClientGetResourceUploadURLOptions contains the optional parameters for the AppsClient.GetResourceUploadURL
//     method.
func (client *AppsClient) GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientGetResourceUploadURLOptions) (AppsClientGetResourceUploadURLResponse, error) {
	var err error
	req, err := client.getResourceUploadURLCreateRequest(ctx, resourceGroupName, serviceName, appName, options)
	if err != nil {
		return AppsClientGetResourceUploadURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppsClientGetResourceUploadURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppsClientGetResourceUploadURLResponse{}, err
	}
	resp, err := client.getResourceUploadURLHandleResponse(httpResp)
	return resp, err
}

// getResourceUploadURLCreateRequest creates the GetResourceUploadURL request.
func (client *AppsClient) getResourceUploadURLCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, options *AppsClientGetResourceUploadURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/getResourceUploadUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getResourceUploadURLHandleResponse handles the GetResourceUploadURL response.
func (client *AppsClient) getResourceUploadURLHandleResponse(resp *http.Response) (AppsClientGetResourceUploadURLResponse, error) {
	result := AppsClientGetResourceUploadURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceUploadDefinition); err != nil {
		return AppsClientGetResourceUploadURLResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - AppsClientListOptions contains the optional parameters for the AppsClient.NewListPager method.
func (client *AppsClient) NewListPager(resourceGroupName string, serviceName string, options *AppsClientListOptions) *runtime.Pager[AppsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AppsClientListResponse]{
		More: func(page AppsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppsClientListResponse) (AppsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AppsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AppsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *AppsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AppsClient) listHandleResponse(resp *http.Response) (AppsClientListResponse, error) {
	result := AppsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResourceCollection); err != nil {
		return AppsClientListResponse{}, err
	}
	return result, nil
}

// BeginSetActiveDeployments - Set existing Deployment under the app as active
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - activeDeploymentCollection - A list of Deployment name to be active.
//   - options - AppsClientBeginSetActiveDeploymentsOptions contains the optional parameters for the AppsClient.BeginSetActiveDeployments
//     method.
func (client *AppsClient) BeginSetActiveDeployments(ctx context.Context, resourceGroupName string, serviceName string, appName string, activeDeploymentCollection ActiveDeploymentCollection, options *AppsClientBeginSetActiveDeploymentsOptions) (*runtime.Poller[AppsClientSetActiveDeploymentsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.setActiveDeployments(ctx, resourceGroupName, serviceName, appName, activeDeploymentCollection, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AppsClientSetActiveDeploymentsResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AppsClientSetActiveDeploymentsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// SetActiveDeployments - Set existing Deployment under the app as active
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *AppsClient) setActiveDeployments(ctx context.Context, resourceGroupName string, serviceName string, appName string, activeDeploymentCollection ActiveDeploymentCollection, options *AppsClientBeginSetActiveDeploymentsOptions) (*http.Response, error) {
	var err error
	req, err := client.setActiveDeploymentsCreateRequest(ctx, resourceGroupName, serviceName, appName, activeDeploymentCollection, options)
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

// setActiveDeploymentsCreateRequest creates the SetActiveDeployments request.
func (client *AppsClient) setActiveDeploymentsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, activeDeploymentCollection ActiveDeploymentCollection, options *AppsClientBeginSetActiveDeploymentsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/setActiveDeployments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, activeDeploymentCollection); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Operation to update an exiting App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - appResource - Parameters for the update operation
//   - options - AppsClientBeginUpdateOptions contains the optional parameters for the AppsClient.BeginUpdate method.
func (client *AppsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginUpdateOptions) (*runtime.Poller[AppsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serviceName, appName, appResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AppsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AppsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Operation to update an exiting App.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *AppsClient) update(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, appName, appResource, options)
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
func (client *AppsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource AppResource, options *AppsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, appResource); err != nil {
		return nil, err
	}
	return req, nil
}

// ValidateDomain - Check the resource name is valid as well as not in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - appName - The name of the App resource.
//   - validatePayload - Custom domain payload to be validated
//   - options - AppsClientValidateDomainOptions contains the optional parameters for the AppsClient.ValidateDomain method.
func (client *AppsClient) ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload CustomDomainValidatePayload, options *AppsClientValidateDomainOptions) (AppsClientValidateDomainResponse, error) {
	var err error
	req, err := client.validateDomainCreateRequest(ctx, resourceGroupName, serviceName, appName, validatePayload, options)
	if err != nil {
		return AppsClientValidateDomainResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppsClientValidateDomainResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppsClientValidateDomainResponse{}, err
	}
	resp, err := client.validateDomainHandleResponse(httpResp)
	return resp, err
}

// validateDomainCreateRequest creates the ValidateDomain request.
func (client *AppsClient) validateDomainCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload CustomDomainValidatePayload, options *AppsClientValidateDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/validateDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, validatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// validateDomainHandleResponse handles the ValidateDomain response.
func (client *AppsClient) validateDomainHandleResponse(resp *http.Response) (AppsClientValidateDomainResponse, error) {
	result := AppsClientValidateDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainValidateResult); err != nil {
		return AppsClientValidateDomainResponse{}, err
	}
	return result, nil
}
