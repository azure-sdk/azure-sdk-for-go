//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsphere

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

// DevicesClient contains the methods for the Devices group.
// Don't use this type directly, use NewDevicesClient() instead.
type DevicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDevicesClient creates a new instance of DevicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDevicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DevicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Device. Use '.unassigned' or '.default' for the device group and product names to claim
// a device to the catalog only.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - deviceName - Device name
//   - resource - Resource create parameters.
//   - options - DevicesClientBeginCreateOrUpdateOptions contains the optional parameters for the DevicesClient.BeginCreateOrUpdate
//     method.
func (client *DevicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, resource Device, options *DevicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a Device. Use '.unassigned' or '.default' for the device group and product names to claim a device
// to the catalog only.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DevicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, resource Device, options *DevicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DevicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, resource, options)
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
func (client *DevicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, resource Device, options *DevicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - deviceName - Device name
//   - options - DevicesClientBeginDeleteOptions contains the optional parameters for the DevicesClient.BeginDelete method.
func (client *DevicesClient) BeginDelete(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*runtime.Poller[DevicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DevicesClient) deleteOperation(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DevicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, options)
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
func (client *DevicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginGenerateCapabilityImage - Generates the capability image for the device. Use '.unassigned' or '.default' for the device
// group and product names to generate the image for a device that does not belong to a specific device group
// and product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - deviceName - Device name
//   - generateDeviceCapabilityRequest - Generate capability image request body.
//   - options - DevicesClientBeginGenerateCapabilityImageOptions contains the optional parameters for the DevicesClient.BeginGenerateCapabilityImage
//     method.
func (client *DevicesClient) BeginGenerateCapabilityImage(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, generateDeviceCapabilityRequest GenerateCapabilityImageRequest, options *DevicesClientBeginGenerateCapabilityImageOptions) (*runtime.Poller[DevicesClientGenerateCapabilityImageResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.generateCapabilityImage(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, generateDeviceCapabilityRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientGenerateCapabilityImageResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevicesClientGenerateCapabilityImageResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GenerateCapabilityImage - Generates the capability image for the device. Use '.unassigned' or '.default' for the device
// group and product names to generate the image for a device that does not belong to a specific device group
// and product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DevicesClient) generateCapabilityImage(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, generateDeviceCapabilityRequest GenerateCapabilityImageRequest, options *DevicesClientBeginGenerateCapabilityImageOptions) (*http.Response, error) {
	var err error
	const operationName = "DevicesClient.BeginGenerateCapabilityImage"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateCapabilityImageCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, generateDeviceCapabilityRequest, options)
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

// generateCapabilityImageCreateRequest creates the GenerateCapabilityImage request.
func (client *DevicesClient) generateCapabilityImageCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, generateDeviceCapabilityRequest GenerateCapabilityImageRequest, options *DevicesClientBeginGenerateCapabilityImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices/{deviceName}/generateCapabilityImage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, generateDeviceCapabilityRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a Device. Use '.unassigned' or '.default' for the device group and product names when a device does not belong
// to a device group and product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - deviceName - Device name
//   - options - DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
func (client *DevicesClient) Get(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, options *DevicesClientGetOptions) (DevicesClientGetResponse, error) {
	var err error
	const operationName = "DevicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, options)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DevicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DevicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, options *DevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevicesClient) getHandleResponse(resp *http.Response) (DevicesClientGetResponse, error) {
	result := DevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Device); err != nil {
		return DevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDeviceGroupPager - List Device resources by DeviceGroup. '.default' and '.unassigned' are system defined values
// and cannot be used for product or device group name.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - options - DevicesClientListByDeviceGroupOptions contains the optional parameters for the DevicesClient.NewListByDeviceGroupPager
//     method.
func (client *DevicesClient) NewListByDeviceGroupPager(resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DevicesClientListByDeviceGroupOptions) *runtime.Pager[DevicesClientListByDeviceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevicesClientListByDeviceGroupResponse]{
		More: func(page DevicesClientListByDeviceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevicesClientListByDeviceGroupResponse) (DevicesClientListByDeviceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DevicesClient.NewListByDeviceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDeviceGroupCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, options)
			}, nil)
			if err != nil {
				return DevicesClientListByDeviceGroupResponse{}, err
			}
			return client.listByDeviceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDeviceGroupCreateRequest creates the ListByDeviceGroup request.
func (client *DevicesClient) listByDeviceGroupCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DevicesClientListByDeviceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDeviceGroupHandleResponse handles the ListByDeviceGroup response.
func (client *DevicesClient) listByDeviceGroupHandleResponse(resp *http.Response) (DevicesClientListByDeviceGroupResponse, error) {
	result := DevicesClientListByDeviceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceListResult); err != nil {
		return DevicesClientListByDeviceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Device. Use '.unassigned' or '.default' for the device group and product names to move a device
// to the catalog level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - deviceName - Device name
//   - properties - The resource properties to be updated.
//   - options - DevicesClientBeginUpdateOptions contains the optional parameters for the DevicesClient.BeginUpdate method.
func (client *DevicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, properties DeviceUpdate, options *DevicesClientBeginUpdateOptions) (*runtime.Poller[DevicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DevicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Device. Use '.unassigned' or '.default' for the device group and product names to move a device to the
// catalog level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DevicesClient) update(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, properties DeviceUpdate, options *DevicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DevicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, deviceName, properties, options)
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
func (client *DevicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, deviceName string, properties DeviceUpdate, options *DevicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/devices/{deviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
