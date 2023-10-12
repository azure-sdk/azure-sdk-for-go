//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// CustomizationTasksClient contains the methods for the CustomizationTasks group.
// Don't use this type directly, use NewCustomizationTasksClient() instead.
type CustomizationTasksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomizationTasksClient creates a new instance of CustomizationTasksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomizationTasksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomizationTasksClient, error) {
	cl, err := arm.NewClient(moduleName+".CustomizationTasksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomizationTasksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a Task from the catalog
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - taskName - The name of the Task.
//   - options - CustomizationTasksClientGetOptions contains the optional parameters for the CustomizationTasksClient.Get method.
func (client *CustomizationTasksClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, taskName string, options *CustomizationTasksClientGetOptions) (CustomizationTasksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, taskName, options)
	if err != nil {
		return CustomizationTasksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomizationTasksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomizationTasksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CustomizationTasksClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, taskName string, options *CustomizationTasksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/tasks/{taskName}"
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
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomizationTasksClient) getHandleResponse(resp *http.Response) (CustomizationTasksClientGetResponse, error) {
	result := CustomizationTasksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizationTask); err != nil {
		return CustomizationTasksClientGetResponse{}, err
	}
	return result, nil
}

// GetErrorDetails - Gets Customization Task error details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - taskName - The name of the Task.
//   - options - CustomizationTasksClientGetErrorDetailsOptions contains the optional parameters for the CustomizationTasksClient.GetErrorDetails
//     method.
func (client *CustomizationTasksClient) GetErrorDetails(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, taskName string, options *CustomizationTasksClientGetErrorDetailsOptions) (CustomizationTasksClientGetErrorDetailsResponse, error) {
	var err error
	req, err := client.getErrorDetailsCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, taskName, options)
	if err != nil {
		return CustomizationTasksClientGetErrorDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomizationTasksClientGetErrorDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomizationTasksClientGetErrorDetailsResponse{}, err
	}
	resp, err := client.getErrorDetailsHandleResponse(httpResp)
	return resp, err
}

// getErrorDetailsCreateRequest creates the GetErrorDetails request.
func (client *CustomizationTasksClient) getErrorDetailsCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, taskName string, options *CustomizationTasksClientGetErrorDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/tasks/{taskName}/getErrorDetails"
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
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getErrorDetailsHandleResponse handles the GetErrorDetails response.
func (client *CustomizationTasksClient) getErrorDetailsHandleResponse(resp *http.Response) (CustomizationTasksClientGetErrorDetailsResponse, error) {
	result := CustomizationTasksClientGetErrorDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatalogResourceValidationErrorDetails); err != nil {
		return CustomizationTasksClientGetErrorDetailsResponse{}, err
	}
	return result, nil
}

// NewListByCatalogPager - List Tasks in the catalog.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CustomizationTasksClientListByCatalogOptions contains the optional parameters for the CustomizationTasksClient.NewListByCatalogPager
//     method.
func (client *CustomizationTasksClient) NewListByCatalogPager(resourceGroupName string, devCenterName string, catalogName string, options *CustomizationTasksClientListByCatalogOptions) *runtime.Pager[CustomizationTasksClientListByCatalogResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomizationTasksClientListByCatalogResponse]{
		More: func(page CustomizationTasksClientListByCatalogResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomizationTasksClientListByCatalogResponse) (CustomizationTasksClientListByCatalogResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCatalogCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomizationTasksClientListByCatalogResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CustomizationTasksClientListByCatalogResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomizationTasksClientListByCatalogResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCatalogHandleResponse(resp)
		},
	})
}

// listByCatalogCreateRequest creates the ListByCatalog request.
func (client *CustomizationTasksClient) listByCatalogCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CustomizationTasksClientListByCatalogOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/tasks"
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
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCatalogHandleResponse handles the ListByCatalog response.
func (client *CustomizationTasksClient) listByCatalogHandleResponse(resp *http.Response) (CustomizationTasksClientListByCatalogResponse, error) {
	result := CustomizationTasksClientListByCatalogResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomizationTaskListResult); err != nil {
		return CustomizationTasksClientListByCatalogResponse{}, err
	}
	return result, nil
}
