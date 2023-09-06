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

// PredefinedAcceleratorsClient contains the methods for the PredefinedAccelerators group.
// Don't use this type directly, use NewPredefinedAcceleratorsClient() instead.
type PredefinedAcceleratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPredefinedAcceleratorsClient creates a new instance of PredefinedAcceleratorsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPredefinedAcceleratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PredefinedAcceleratorsClient, error) {
	cl, err := arm.NewClient(moduleName+".PredefinedAcceleratorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PredefinedAcceleratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDisable - Disable predefined accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - predefinedAcceleratorName - The name of the predefined accelerator.
//   - options - PredefinedAcceleratorsClientBeginDisableOptions contains the optional parameters for the PredefinedAcceleratorsClient.BeginDisable
//     method.
func (client *PredefinedAcceleratorsClient) BeginDisable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginDisableOptions) (*runtime.Poller[PredefinedAcceleratorsClientDisableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disable(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PredefinedAcceleratorsClientDisableResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PredefinedAcceleratorsClientDisableResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Disable - Disable predefined accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *PredefinedAcceleratorsClient) disable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginDisableOptions) (*http.Response, error) {
	var err error
	req, err := client.disableCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName, options)
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

// disableCreateRequest creates the Disable request.
func (client *PredefinedAcceleratorsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}/disable"
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
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	if predefinedAcceleratorName == "" {
		return nil, errors.New("parameter predefinedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predefinedAcceleratorName}", url.PathEscape(predefinedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginEnable - Enable predefined accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - predefinedAcceleratorName - The name of the predefined accelerator.
//   - options - PredefinedAcceleratorsClientBeginEnableOptions contains the optional parameters for the PredefinedAcceleratorsClient.BeginEnable
//     method.
func (client *PredefinedAcceleratorsClient) BeginEnable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginEnableOptions) (*runtime.Poller[PredefinedAcceleratorsClientEnableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.enable(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PredefinedAcceleratorsClientEnableResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PredefinedAcceleratorsClientEnableResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Enable - Enable predefined accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
func (client *PredefinedAcceleratorsClient) enable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginEnableOptions) (*http.Response, error) {
	var err error
	req, err := client.enableCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName, options)
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

// enableCreateRequest creates the Enable request.
func (client *PredefinedAcceleratorsClient) enableCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientBeginEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}/enable"
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
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	if predefinedAcceleratorName == "" {
		return nil, errors.New("parameter predefinedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predefinedAcceleratorName}", url.PathEscape(predefinedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the predefined accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - predefinedAcceleratorName - The name of the predefined accelerator.
//   - options - PredefinedAcceleratorsClientGetOptions contains the optional parameters for the PredefinedAcceleratorsClient.Get
//     method.
func (client *PredefinedAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientGetOptions) (PredefinedAcceleratorsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName, options)
	if err != nil {
		return PredefinedAcceleratorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PredefinedAcceleratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PredefinedAcceleratorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PredefinedAcceleratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string, options *PredefinedAcceleratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}"
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
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	if predefinedAcceleratorName == "" {
		return nil, errors.New("parameter predefinedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predefinedAcceleratorName}", url.PathEscape(predefinedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PredefinedAcceleratorsClient) getHandleResponse(resp *http.Response) (PredefinedAcceleratorsClientGetResponse, error) {
	result := PredefinedAcceleratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredefinedAcceleratorResource); err != nil {
		return PredefinedAcceleratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handle requests to list all predefined accelerators.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - options - PredefinedAcceleratorsClientListOptions contains the optional parameters for the PredefinedAcceleratorsClient.NewListPager
//     method.
func (client *PredefinedAcceleratorsClient) NewListPager(resourceGroupName string, serviceName string, applicationAcceleratorName string, options *PredefinedAcceleratorsClientListOptions) *runtime.Pager[PredefinedAcceleratorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PredefinedAcceleratorsClientListResponse]{
		More: func(page PredefinedAcceleratorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PredefinedAcceleratorsClientListResponse) (PredefinedAcceleratorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PredefinedAcceleratorsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PredefinedAcceleratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PredefinedAcceleratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PredefinedAcceleratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *PredefinedAcceleratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators"
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
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PredefinedAcceleratorsClient) listHandleResponse(resp *http.Response) (PredefinedAcceleratorsClientListResponse, error) {
	result := PredefinedAcceleratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredefinedAcceleratorResourceCollection); err != nil {
		return PredefinedAcceleratorsClientListResponse{}, err
	}
	return result, nil
}
