//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// DataExportsClient contains the methods for the DataExports group.
// Don't use this type directly, use NewDataExportsClient() instead.
type DataExportsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataExportsClient creates a new instance of DataExportsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataExportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataExportsClient, error) {
	cl, err := arm.NewClient(moduleName+".DataExportsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataExportsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a data export.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataExportName - The data export rule name.
//   - parameters - The parameters required to create or update a data export.
//   - options - DataExportsClientCreateOrUpdateOptions contains the optional parameters for the DataExportsClient.CreateOrUpdate
//     method.
func (client *DataExportsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, parameters DataExport, options *DataExportsClientCreateOrUpdateOptions) (DataExportsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataExportName, parameters, options)
	if err != nil {
		return DataExportsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataExportsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataExportsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataExportsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, parameters DataExport, options *DataExportsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataExports/{dataExportName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataExportName == "" {
		return nil, errors.New("parameter dataExportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataExportName}", url.PathEscape(dataExportName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataExportsClient) createOrUpdateHandleResponse(resp *http.Response) (DataExportsClientCreateOrUpdateResponse, error) {
	result := DataExportsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataExport); err != nil {
		return DataExportsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified data export in a given workspace..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataExportName - The data export rule name.
//   - options - DataExportsClientDeleteOptions contains the optional parameters for the DataExportsClient.Delete method.
func (client *DataExportsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, options *DataExportsClientDeleteOptions) (DataExportsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataExportName, options)
	if err != nil {
		return DataExportsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataExportsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return DataExportsClientDeleteResponse{}, err
	}
	return DataExportsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataExportsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, options *DataExportsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataExports/{dataExportName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataExportName == "" {
		return nil, errors.New("parameter dataExportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataExportName}", url.PathEscape(dataExportName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a data export instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataExportName - The data export rule name.
//   - options - DataExportsClientGetOptions contains the optional parameters for the DataExportsClient.Get method.
func (client *DataExportsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, options *DataExportsClientGetOptions) (DataExportsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataExportName, options)
	if err != nil {
		return DataExportsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataExportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataExportsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataExportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataExportName string, options *DataExportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataExports/{dataExportName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataExportName == "" {
		return nil, errors.New("parameter dataExportName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataExportName}", url.PathEscape(dataExportName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataExportsClient) getHandleResponse(resp *http.Response) (DataExportsClientGetResponse, error) {
	result := DataExportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataExport); err != nil {
		return DataExportsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists the data export instances within a workspace.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - DataExportsClientListByWorkspaceOptions contains the optional parameters for the DataExportsClient.NewListByWorkspacePager
//     method.
func (client *DataExportsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *DataExportsClientListByWorkspaceOptions) *runtime.Pager[DataExportsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataExportsClientListByWorkspaceResponse]{
		More: func(page DataExportsClientListByWorkspaceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DataExportsClientListByWorkspaceResponse) (DataExportsClientListByWorkspaceResponse, error) {
			req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return DataExportsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataExportsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataExportsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *DataExportsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DataExportsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataExports"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *DataExportsClient) listByWorkspaceHandleResponse(resp *http.Response) (DataExportsClientListByWorkspaceResponse, error) {
	result := DataExportsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataExportListResult); err != nil {
		return DataExportsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
