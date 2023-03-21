//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EnvironmentTypesClient contains the methods for the EnvironmentTypes group.
// Don't use this type directly, use NewEnvironmentTypesClient() instead.
type EnvironmentTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEnvironmentTypesClient creates a new instance of EnvironmentTypesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnvironmentTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentTypesClient, error) {
	cl, err := arm.NewClient(moduleName+".EnvironmentTypesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnvironmentTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an environment type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - environmentTypeName - The name of the environment type.
//   - body - Represents an Environment Type.
//   - options - EnvironmentTypesClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentTypesClient.CreateOrUpdate
//     method.
func (client *EnvironmentTypesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, body EnvironmentType, options *EnvironmentTypesClientCreateOrUpdateOptions) (EnvironmentTypesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, devCenterName, environmentTypeName, body, options)
	if err != nil {
		return EnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentTypesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EnvironmentTypesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, body EnvironmentType, options *EnvironmentTypesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EnvironmentTypesClient) createOrUpdateHandleResponse(resp *http.Response) (EnvironmentTypesClientCreateOrUpdateResponse, error) {
	result := EnvironmentTypesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentType); err != nil {
		return EnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an environment type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - environmentTypeName - The name of the environment type.
//   - options - EnvironmentTypesClientDeleteOptions contains the optional parameters for the EnvironmentTypesClient.Delete method.
func (client *EnvironmentTypesClient) Delete(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, options *EnvironmentTypesClientDeleteOptions) (EnvironmentTypesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, devCenterName, environmentTypeName, options)
	if err != nil {
		return EnvironmentTypesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentTypesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return EnvironmentTypesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return EnvironmentTypesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnvironmentTypesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, options *EnvironmentTypesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an environment type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - environmentTypeName - The name of the environment type.
//   - options - EnvironmentTypesClientGetOptions contains the optional parameters for the EnvironmentTypesClient.Get method.
func (client *EnvironmentTypesClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, options *EnvironmentTypesClientGetOptions) (EnvironmentTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, environmentTypeName, options)
	if err != nil {
		return EnvironmentTypesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnvironmentTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, options *EnvironmentTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnvironmentTypesClient) getHandleResponse(resp *http.Response) (EnvironmentTypesClientGetResponse, error) {
	result := EnvironmentTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentType); err != nil {
		return EnvironmentTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDevCenterPager - Lists environment types for the devcenter.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - options - EnvironmentTypesClientListByDevCenterOptions contains the optional parameters for the EnvironmentTypesClient.NewListByDevCenterPager
//     method.
func (client *EnvironmentTypesClient) NewListByDevCenterPager(resourceGroupName string, devCenterName string, options *EnvironmentTypesClientListByDevCenterOptions) *runtime.Pager[EnvironmentTypesClientListByDevCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentTypesClientListByDevCenterResponse]{
		More: func(page EnvironmentTypesClientListByDevCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentTypesClientListByDevCenterResponse) (EnvironmentTypesClientListByDevCenterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnvironmentTypesClientListByDevCenterResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EnvironmentTypesClientListByDevCenterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnvironmentTypesClientListByDevCenterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDevCenterHandleResponse(resp)
		},
	})
}

// listByDevCenterCreateRequest creates the ListByDevCenter request.
func (client *EnvironmentTypesClient) listByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *EnvironmentTypesClientListByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/environmentTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDevCenterHandleResponse handles the ListByDevCenter response.
func (client *EnvironmentTypesClient) listByDevCenterHandleResponse(resp *http.Response) (EnvironmentTypesClientListByDevCenterResponse, error) {
	result := EnvironmentTypesClientListByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentTypeListResult); err != nil {
		return EnvironmentTypesClientListByDevCenterResponse{}, err
	}
	return result, nil
}

// Update - Partially updates an environment type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - environmentTypeName - The name of the environment type.
//   - body - Updatable environment type properties.
//   - options - EnvironmentTypesClientUpdateOptions contains the optional parameters for the EnvironmentTypesClient.Update method.
func (client *EnvironmentTypesClient) Update(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, body EnvironmentTypeUpdate, options *EnvironmentTypesClientUpdateOptions) (EnvironmentTypesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, devCenterName, environmentTypeName, body, options)
	if err != nil {
		return EnvironmentTypesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentTypesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnvironmentTypesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EnvironmentTypesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, environmentTypeName string, body EnvironmentTypeUpdate, options *EnvironmentTypesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *EnvironmentTypesClient) updateHandleResponse(resp *http.Response) (EnvironmentTypesClientUpdateResponse, error) {
	result := EnvironmentTypesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentType); err != nil {
		return EnvironmentTypesClientUpdateResponse{}, err
	}
	return result, nil
}
