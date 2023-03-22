//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// CustomizedAcceleratorsClient contains the methods for the CustomizedAccelerators group.
// Don't use this type directly, use NewCustomizedAcceleratorsClient() instead.
type CustomizedAcceleratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomizedAcceleratorsClient creates a new instance of CustomizedAcceleratorsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomizedAcceleratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomizedAcceleratorsClient, error) {
	cl, err := arm.NewClient(moduleName+".CustomizedAcceleratorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomizedAcceleratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the customized accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - customizedAcceleratorName - The name of the customized accelerator.
//   - customizedAcceleratorResource - The customized accelerator for the create or update operation
//   - options - CustomizedAcceleratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomizedAcceleratorsClient.BeginCreateOrUpdate
//     method.
func (client *CustomizedAcceleratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource CustomizedAcceleratorResource, options *CustomizedAcceleratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[CustomizedAcceleratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, customizedAcceleratorResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CustomizedAcceleratorsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CustomizedAcceleratorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update the customized accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *CustomizedAcceleratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource CustomizedAcceleratorResource, options *CustomizedAcceleratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, customizedAcceleratorResource, options)
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
func (client *CustomizedAcceleratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, customizedAcceleratorResource CustomizedAcceleratorResource, options *CustomizedAcceleratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}"
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
	if customizedAcceleratorName == "" {
		return nil, errors.New("parameter customizedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customizedAcceleratorName}", url.PathEscape(customizedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, customizedAcceleratorResource)
}

// BeginDelete - Delete the customized accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - customizedAcceleratorName - The name of the customized accelerator.
//   - options - CustomizedAcceleratorsClientBeginDeleteOptions contains the optional parameters for the CustomizedAcceleratorsClient.BeginDelete
//     method.
func (client *CustomizedAcceleratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *CustomizedAcceleratorsClientBeginDeleteOptions) (*runtime.Poller[CustomizedAcceleratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CustomizedAcceleratorsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CustomizedAcceleratorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the customized accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *CustomizedAcceleratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *CustomizedAcceleratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, options)
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
func (client *CustomizedAcceleratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *CustomizedAcceleratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}"
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
	if customizedAcceleratorName == "" {
		return nil, errors.New("parameter customizedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customizedAcceleratorName}", url.PathEscape(customizedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the customized accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - customizedAcceleratorName - The name of the customized accelerator.
//   - options - CustomizedAcceleratorsClientGetOptions contains the optional parameters for the CustomizedAcceleratorsClient.Get
//     method.
func (client *CustomizedAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *CustomizedAcceleratorsClientGetOptions) (CustomizedAcceleratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, options)
	if err != nil {
		return CustomizedAcceleratorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomizedAcceleratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomizedAcceleratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomizedAcceleratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, options *CustomizedAcceleratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}"
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
	if customizedAcceleratorName == "" {
		return nil, errors.New("parameter customizedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customizedAcceleratorName}", url.PathEscape(customizedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomizedAcceleratorsClient) getHandleResponse(resp *http.Response) (CustomizedAcceleratorsClientGetResponse, error) {
	result := CustomizedAcceleratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizedAcceleratorResource); err != nil {
		return CustomizedAcceleratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handle requests to list all customized accelerators.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - options - CustomizedAcceleratorsClientListOptions contains the optional parameters for the CustomizedAcceleratorsClient.NewListPager
//     method.
func (client *CustomizedAcceleratorsClient) NewListPager(resourceGroupName string, serviceName string, applicationAcceleratorName string, options *CustomizedAcceleratorsClientListOptions) *runtime.Pager[CustomizedAcceleratorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomizedAcceleratorsClientListResponse]{
		More: func(page CustomizedAcceleratorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomizedAcceleratorsClientListResponse) (CustomizedAcceleratorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomizedAcceleratorsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CustomizedAcceleratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomizedAcceleratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CustomizedAcceleratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *CustomizedAcceleratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators"
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CustomizedAcceleratorsClient) listHandleResponse(resp *http.Response) (CustomizedAcceleratorsClientListResponse, error) {
	result := CustomizedAcceleratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizedAcceleratorResourceCollection); err != nil {
		return CustomizedAcceleratorsClientListResponse{}, err
	}
	return result, nil
}

// Validate - Check the customized accelerator are valid.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationAcceleratorName - The name of the application accelerator.
//   - customizedAcceleratorName - The name of the customized accelerator.
//   - properties - Customized accelerator properties to be validated
//   - options - CustomizedAcceleratorsClientValidateOptions contains the optional parameters for the CustomizedAcceleratorsClient.Validate
//     method.
func (client *CustomizedAcceleratorsClient) Validate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, properties CustomizedAcceleratorProperties, options *CustomizedAcceleratorsClientValidateOptions) (CustomizedAcceleratorsClientValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, customizedAcceleratorName, properties, options)
	if err != nil {
		return CustomizedAcceleratorsClientValidateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomizedAcceleratorsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return CustomizedAcceleratorsClientValidateResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *CustomizedAcceleratorsClient) validateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, customizedAcceleratorName string, properties CustomizedAcceleratorProperties, options *CustomizedAcceleratorsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/customizedAccelerators/{customizedAcceleratorName}/validate"
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
	if customizedAcceleratorName == "" {
		return nil, errors.New("parameter customizedAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customizedAcceleratorName}", url.PathEscape(customizedAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// validateHandleResponse handles the Validate response.
func (client *CustomizedAcceleratorsClient) validateHandleResponse(resp *http.Response) (CustomizedAcceleratorsClientValidateResponse, error) {
	result := CustomizedAcceleratorsClientValidateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizedAcceleratorValidateResult); err != nil {
		return CustomizedAcceleratorsClientValidateResponse{}, err
	}
	return result, nil
}
