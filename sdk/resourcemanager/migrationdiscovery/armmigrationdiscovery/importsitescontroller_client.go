//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationdiscovery

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

// ImportSitesControllerClient contains the methods for the ImportSitesController group.
// Don't use this type directly, use NewImportSitesControllerClient() instead.
type ImportSitesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImportSitesControllerClient creates a new instance of ImportSitesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImportSitesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImportSitesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImportSitesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a ImportSite
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - body - Resource create parameters.
//   - options - ImportSitesControllerClientCreateOptions contains the optional parameters for the ImportSitesControllerClient.Create
//     method.
func (client *ImportSitesControllerClient) Create(ctx context.Context, resourceGroupName string, siteName string, body ImportSite, options *ImportSitesControllerClientCreateOptions) (ImportSitesControllerClientCreateResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return ImportSitesControllerClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ImportSitesControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, body ImportSite, options *ImportSitesControllerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ImportSitesControllerClient) createHandleResponse(resp *http.Response) (ImportSitesControllerClientCreateResponse, error) {
	result := ImportSitesControllerClientCreateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return ImportSitesControllerClientCreateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSite); err != nil {
		return ImportSitesControllerClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a ImportSite
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - ImportSitesControllerClientDeleteOptions contains the optional parameters for the ImportSitesControllerClient.Delete
//     method.
func (client *ImportSitesControllerClient) Delete(ctx context.Context, resourceGroupName string, siteName string, options *ImportSitesControllerClientDeleteOptions) (ImportSitesControllerClientDeleteResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, options)
	if err != nil {
		return ImportSitesControllerClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientDeleteResponse{}, err
	}
	return ImportSitesControllerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ImportSitesControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *ImportSitesControllerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteImportedMachines - Deletes the imported machines for site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - body - The content of the action request
//   - options - ImportSitesControllerClientDeleteImportedMachinesOptions contains the optional parameters for the ImportSitesControllerClient.DeleteImportedMachines
//     method.
func (client *ImportSitesControllerClient) DeleteImportedMachines(ctx context.Context, resourceGroupName string, siteName string, body any, options *ImportSitesControllerClientDeleteImportedMachinesOptions) (ImportSitesControllerClientDeleteImportedMachinesResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.DeleteImportedMachines"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteImportedMachinesCreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return ImportSitesControllerClientDeleteImportedMachinesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientDeleteImportedMachinesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientDeleteImportedMachinesResponse{}, err
	}
	resp, err := client.deleteImportedMachinesHandleResponse(httpResp)
	return resp, err
}

// deleteImportedMachinesCreateRequest creates the DeleteImportedMachines request.
func (client *ImportSitesControllerClient) deleteImportedMachinesCreateRequest(ctx context.Context, resourceGroupName string, siteName string, body any, options *ImportSitesControllerClientDeleteImportedMachinesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}/deleteImportedMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// deleteImportedMachinesHandleResponse handles the DeleteImportedMachines response.
func (client *ImportSitesControllerClient) deleteImportedMachinesHandleResponse(resp *http.Response) (ImportSitesControllerClientDeleteImportedMachinesResponse, error) {
	result := ImportSitesControllerClientDeleteImportedMachinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasURIResponse); err != nil {
		return ImportSitesControllerClientDeleteImportedMachinesResponse{}, err
	}
	return result, nil
}

// ExportURI - Method to export a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - body - The content of the action request
//   - options - ImportSitesControllerClientExportURIOptions contains the optional parameters for the ImportSitesControllerClient.ExportURI
//     method.
func (client *ImportSitesControllerClient) ExportURI(ctx context.Context, resourceGroupName string, siteName string, body SasURIResponse, options *ImportSitesControllerClientExportURIOptions) (ImportSitesControllerClientExportURIResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.ExportURI"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportURICreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return ImportSitesControllerClientExportURIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientExportURIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientExportURIResponse{}, err
	}
	resp, err := client.exportURIHandleResponse(httpResp)
	return resp, err
}

// exportURICreateRequest creates the ExportURI request.
func (client *ImportSitesControllerClient) exportURICreateRequest(ctx context.Context, resourceGroupName string, siteName string, body SasURIResponse, options *ImportSitesControllerClientExportURIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}/exportUri"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// exportURIHandleResponse handles the ExportURI response.
func (client *ImportSitesControllerClient) exportURIHandleResponse(resp *http.Response) (ImportSitesControllerClientExportURIResponse, error) {
	result := ImportSitesControllerClientExportURIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasURIResponse); err != nil {
		return ImportSitesControllerClientExportURIResponse{}, err
	}
	return result, nil
}

// Get - Get a ImportSite
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - ImportSitesControllerClientGetOptions contains the optional parameters for the ImportSitesControllerClient.Get
//     method.
func (client *ImportSitesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, options *ImportSitesControllerClientGetOptions) (ImportSitesControllerClientGetResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, options)
	if err != nil {
		return ImportSitesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImportSitesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *ImportSitesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
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
func (client *ImportSitesControllerClient) getHandleResponse(resp *http.Response) (ImportSitesControllerClientGetResponse, error) {
	result := ImportSitesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSite); err != nil {
		return ImportSitesControllerClientGetResponse{}, err
	}
	return result, nil
}

// ImportURI - Method to import a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - body - The content of the action request
//   - options - ImportSitesControllerClientImportURIOptions contains the optional parameters for the ImportSitesControllerClient.ImportURI
//     method.
func (client *ImportSitesControllerClient) ImportURI(ctx context.Context, resourceGroupName string, siteName string, body SasURIResponse, options *ImportSitesControllerClientImportURIOptions) (ImportSitesControllerClientImportURIResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.ImportURI"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.importURICreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return ImportSitesControllerClientImportURIResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientImportURIResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientImportURIResponse{}, err
	}
	resp, err := client.importURIHandleResponse(httpResp)
	return resp, err
}

// importURICreateRequest creates the ImportURI request.
func (client *ImportSitesControllerClient) importURICreateRequest(ctx context.Context, resourceGroupName string, siteName string, body SasURIResponse, options *ImportSitesControllerClientImportURIOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}/importUri"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// importURIHandleResponse handles the ImportURI response.
func (client *ImportSitesControllerClient) importURIHandleResponse(resp *http.Response) (ImportSitesControllerClientImportURIResponse, error) {
	result := ImportSitesControllerClientImportURIResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasURIResponse); err != nil {
		return ImportSitesControllerClientImportURIResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all import sites.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ImportSitesControllerClientListByResourceGroupOptions contains the optional parameters for the ImportSitesControllerClient.NewListByResourceGroupPager
//     method.
func (client *ImportSitesControllerClient) NewListByResourceGroupPager(resourceGroupName string, options *ImportSitesControllerClientListByResourceGroupOptions) *runtime.Pager[ImportSitesControllerClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImportSitesControllerClientListByResourceGroupResponse]{
		More: func(page ImportSitesControllerClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImportSitesControllerClientListByResourceGroupResponse) (ImportSitesControllerClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImportSitesControllerClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ImportSitesControllerClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ImportSitesControllerClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ImportSitesControllerClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ImportSitesControllerClient) listByResourceGroupHandleResponse(resp *http.Response) (ImportSitesControllerClientListByResourceGroupResponse, error) {
	result := ImportSitesControllerClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSiteListResult); err != nil {
		return ImportSitesControllerClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List ImportSite resources by subscription ID
//
// Generated from API version 2023-10-01-preview
//   - options - ImportSitesControllerClientListBySubscriptionOptions contains the optional parameters for the ImportSitesControllerClient.NewListBySubscriptionPager
//     method.
func (client *ImportSitesControllerClient) NewListBySubscriptionPager(options *ImportSitesControllerClientListBySubscriptionOptions) *runtime.Pager[ImportSitesControllerClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImportSitesControllerClientListBySubscriptionResponse]{
		More: func(page ImportSitesControllerClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImportSitesControllerClientListBySubscriptionResponse) (ImportSitesControllerClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImportSitesControllerClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ImportSitesControllerClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ImportSitesControllerClient) listBySubscriptionCreateRequest(ctx context.Context, options *ImportSitesControllerClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OffAzure/importSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ImportSitesControllerClient) listBySubscriptionHandleResponse(resp *http.Response) (ImportSitesControllerClientListBySubscriptionResponse, error) {
	result := ImportSitesControllerClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSiteListResult); err != nil {
		return ImportSitesControllerClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a ImportSite
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - body - The resource properties to be updated.
//   - options - ImportSitesControllerClientUpdateOptions contains the optional parameters for the ImportSitesControllerClient.Update
//     method.
func (client *ImportSitesControllerClient) Update(ctx context.Context, resourceGroupName string, siteName string, body ImportSiteUpdate, options *ImportSitesControllerClientUpdateOptions) (ImportSitesControllerClientUpdateResponse, error) {
	var err error
	const operationName = "ImportSitesControllerClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, siteName, body, options)
	if err != nil {
		return ImportSitesControllerClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImportSitesControllerClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImportSitesControllerClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ImportSitesControllerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, siteName string, body ImportSiteUpdate, options *ImportSitesControllerClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/importSites/{siteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ImportSitesControllerClient) updateHandleResponse(resp *http.Response) (ImportSitesControllerClientUpdateResponse, error) {
	result := ImportSitesControllerClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSite); err != nil {
		return ImportSitesControllerClientUpdateResponse{}, err
	}
	return result, nil
}
