// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

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

// AzureBareMetalStorageInstancesClient contains the methods for the AzureBareMetalStorageInstances group.
// Don't use this type directly, use NewAzureBareMetalStorageInstancesClient() instead.
type AzureBareMetalStorageInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureBareMetalStorageInstancesClient creates a new instance of AzureBareMetalStorageInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureBareMetalStorageInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureBareMetalStorageInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureBareMetalStorageInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates an Azure Bare Metal Storage Instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalStorageInstanceName - Name of the Azure Bare Metal Storage Instance, also known as the ResourceName.
//   - requestBodyParameters - request body for put call
//   - options - AzureBareMetalStorageInstancesClientCreateOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Create
//     method.
func (client *AzureBareMetalStorageInstancesClient) Create(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, requestBodyParameters AzureBareMetalStorageInstance, options *AzureBareMetalStorageInstancesClientCreateOptions) (AzureBareMetalStorageInstancesClientCreateResponse, error) {
	var err error
	const operationName = "AzureBareMetalStorageInstancesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, azureBareMetalStorageInstanceName, requestBodyParameters, options)
	if err != nil {
		return AzureBareMetalStorageInstancesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalStorageInstancesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalStorageInstancesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *AzureBareMetalStorageInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, requestBodyParameters AzureBareMetalStorageInstance, _ *AzureBareMetalStorageInstancesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances/{azureBareMetalStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalStorageInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalStorageInstanceName}", url.PathEscape(azureBareMetalStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBodyParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *AzureBareMetalStorageInstancesClient) createHandleResponse(resp *http.Response) (AzureBareMetalStorageInstancesClientCreateResponse, error) {
	result := AzureBareMetalStorageInstancesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalStorageInstance); err != nil {
		return AzureBareMetalStorageInstancesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Azure Bare Metal Storage Instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalStorageInstanceName - Name of the Azure Bare Metal Storage Instance, also known as the ResourceName.
//   - options - AzureBareMetalStorageInstancesClientDeleteOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Delete
//     method.
func (client *AzureBareMetalStorageInstancesClient) Delete(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, options *AzureBareMetalStorageInstancesClientDeleteOptions) (AzureBareMetalStorageInstancesClientDeleteResponse, error) {
	var err error
	const operationName = "AzureBareMetalStorageInstancesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, azureBareMetalStorageInstanceName, options)
	if err != nil {
		return AzureBareMetalStorageInstancesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalStorageInstancesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalStorageInstancesClientDeleteResponse{}, err
	}
	return AzureBareMetalStorageInstancesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AzureBareMetalStorageInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, _ *AzureBareMetalStorageInstancesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances/{azureBareMetalStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalStorageInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalStorageInstanceName}", url.PathEscape(azureBareMetalStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Azure Bare Metal Storage instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalStorageInstanceName - Name of the Azure Bare Metal Storage Instance, also known as the ResourceName.
//   - options - AzureBareMetalStorageInstancesClientGetOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Get
//     method.
func (client *AzureBareMetalStorageInstancesClient) Get(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, options *AzureBareMetalStorageInstancesClientGetOptions) (AzureBareMetalStorageInstancesClientGetResponse, error) {
	var err error
	const operationName = "AzureBareMetalStorageInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureBareMetalStorageInstanceName, options)
	if err != nil {
		return AzureBareMetalStorageInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalStorageInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalStorageInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureBareMetalStorageInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, _ *AzureBareMetalStorageInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances/{azureBareMetalStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalStorageInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalStorageInstanceName}", url.PathEscape(azureBareMetalStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureBareMetalStorageInstancesClient) getHandleResponse(resp *http.Response) (AzureBareMetalStorageInstancesClientGetResponse, error) {
	result := AzureBareMetalStorageInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalStorageInstance); err != nil {
		return AzureBareMetalStorageInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of AzureBareMetalStorage instances in the specified subscription and resource
// group. The operations returns various properties of each Azure Bare Metal Instance.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AzureBareMetalStorageInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.NewListByResourceGroupPager
//     method.
func (client *AzureBareMetalStorageInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureBareMetalStorageInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureBareMetalStorageInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalStorageInstancesClientListByResourceGroupResponse]{
		More: func(page AzureBareMetalStorageInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalStorageInstancesClientListByResourceGroupResponse) (AzureBareMetalStorageInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureBareMetalStorageInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AzureBareMetalStorageInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureBareMetalStorageInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *AzureBareMetalStorageInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureBareMetalStorageInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureBareMetalStorageInstancesClientListByResourceGroupResponse, error) {
	result := AzureBareMetalStorageInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalStorageInstancesListResult); err != nil {
		return AzureBareMetalStorageInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of AzureBareMetalStorage instances in the specified subscription. The operations
// returns various properties of each Azure Bare Metal Instance.
//
// Generated from API version 2024-08-01-preview
//   - options - AzureBareMetalStorageInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.NewListBySubscriptionPager
//     method.
func (client *AzureBareMetalStorageInstancesClient) NewListBySubscriptionPager(options *AzureBareMetalStorageInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureBareMetalStorageInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalStorageInstancesClientListBySubscriptionResponse]{
		More: func(page AzureBareMetalStorageInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalStorageInstancesClientListBySubscriptionResponse) (AzureBareMetalStorageInstancesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureBareMetalStorageInstancesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AzureBareMetalStorageInstancesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureBareMetalStorageInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, _ *AzureBareMetalStorageInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureBareMetalStorageInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureBareMetalStorageInstancesClientListBySubscriptionResponse, error) {
	result := AzureBareMetalStorageInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalStorageInstancesListResult); err != nil {
		return AzureBareMetalStorageInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches the Tags field of a Azure BareMetalStorage instance for the specified subscription, resource group, and
// instance name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureBareMetalStorageInstanceName - Name of the Azure Bare Metal Storage Instance, also known as the ResourceName.
//   - azureBareMetalStorageInstanceBodyParameter - Request body that only contains the Tags and Identity Field
//   - options - AzureBareMetalStorageInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalStorageInstancesClient.Update
//     method.
func (client *AzureBareMetalStorageInstancesClient) Update(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, azureBareMetalStorageInstanceBodyParameter AzureBareMetalStorageInstanceBody, options *AzureBareMetalStorageInstancesClientUpdateOptions) (AzureBareMetalStorageInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "AzureBareMetalStorageInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureBareMetalStorageInstanceName, azureBareMetalStorageInstanceBodyParameter, options)
	if err != nil {
		return AzureBareMetalStorageInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureBareMetalStorageInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureBareMetalStorageInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AzureBareMetalStorageInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalStorageInstanceName string, azureBareMetalStorageInstanceBodyParameter AzureBareMetalStorageInstanceBody, _ *AzureBareMetalStorageInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalStorageInstances/{azureBareMetalStorageInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalStorageInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalStorageInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalStorageInstanceName}", url.PathEscape(azureBareMetalStorageInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, azureBareMetalStorageInstanceBodyParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AzureBareMetalStorageInstancesClient) updateHandleResponse(resp *http.Response) (AzureBareMetalStorageInstancesClientUpdateResponse, error) {
	result := AzureBareMetalStorageInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalStorageInstance); err != nil {
		return AzureBareMetalStorageInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
