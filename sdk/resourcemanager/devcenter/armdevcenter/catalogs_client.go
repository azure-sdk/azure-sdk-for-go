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

// CatalogsClient contains the methods for the Catalogs group.
// Don't use this type directly, use NewCatalogsClient() instead.
type CatalogsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCatalogsClient creates a new instance of CatalogsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCatalogsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CatalogsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CatalogsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginConnect - Connects a catalog to enable syncing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogsClientBeginConnectOptions contains the optional parameters for the CatalogsClient.BeginConnect method.
func (client *CatalogsClient) BeginConnect(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginConnectOptions) (*runtime.Poller[CatalogsClientConnectResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.connect(ctx, resourceGroupName, devCenterName, catalogName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CatalogsClientConnectResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CatalogsClientConnectResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Connect - Connects a catalog to enable syncing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
func (client *CatalogsClient) connect(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginConnectOptions) (*http.Response, error) {
	var err error
	const operationName = "CatalogsClient.BeginConnect"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.connectCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// connectCreateRequest creates the Connect request.
func (client *CatalogsClient) connectCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginConnectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/connect"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Creates or updates a catalog.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - body - Represents a catalog.
//   - options - CatalogsClientBeginCreateOrUpdateOptions contains the optional parameters for the CatalogsClient.BeginCreateOrUpdate
//     method.
func (client *CatalogsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body Catalog, options *CatalogsClientBeginCreateOrUpdateOptions) (*runtime.Poller[CatalogsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, devCenterName, catalogName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CatalogsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CatalogsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a catalog.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
func (client *CatalogsClient) createOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body Catalog, options *CatalogsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CatalogsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, body, options)
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
func (client *CatalogsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body Catalog, options *CatalogsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a catalog resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogsClientBeginDeleteOptions contains the optional parameters for the CatalogsClient.BeginDelete method.
func (client *CatalogsClient) BeginDelete(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginDeleteOptions) (*runtime.Poller[CatalogsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, devCenterName, catalogName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CatalogsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CatalogsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a catalog resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
func (client *CatalogsClient) deleteOperation(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CatalogsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CatalogsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}"
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

// Get - Gets a catalog
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogsClientGetOptions contains the optional parameters for the CatalogsClient.Get method.
func (client *CatalogsClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientGetOptions) (CatalogsClientGetResponse, error) {
	var err error
	const operationName = "CatalogsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
	if err != nil {
		return CatalogsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CatalogsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CatalogsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CatalogsClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}"
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
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CatalogsClient) getHandleResponse(resp *http.Response) (CatalogsClientGetResponse, error) {
	result := CatalogsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Catalog); err != nil {
		return CatalogsClientGetResponse{}, err
	}
	return result, nil
}

// GetSyncErrorDetails - Gets catalog synchronization error details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogsClientGetSyncErrorDetailsOptions contains the optional parameters for the CatalogsClient.GetSyncErrorDetails
//     method.
func (client *CatalogsClient) GetSyncErrorDetails(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientGetSyncErrorDetailsOptions) (CatalogsClientGetSyncErrorDetailsResponse, error) {
	var err error
	const operationName = "CatalogsClient.GetSyncErrorDetails"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSyncErrorDetailsCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
	if err != nil {
		return CatalogsClientGetSyncErrorDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CatalogsClientGetSyncErrorDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CatalogsClientGetSyncErrorDetailsResponse{}, err
	}
	resp, err := client.getSyncErrorDetailsHandleResponse(httpResp)
	return resp, err
}

// getSyncErrorDetailsCreateRequest creates the GetSyncErrorDetails request.
func (client *CatalogsClient) getSyncErrorDetailsCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientGetSyncErrorDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/getSyncErrorDetails"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSyncErrorDetailsHandleResponse handles the GetSyncErrorDetails response.
func (client *CatalogsClient) getSyncErrorDetailsHandleResponse(resp *http.Response) (CatalogsClientGetSyncErrorDetailsResponse, error) {
	result := CatalogsClientGetSyncErrorDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncErrorDetails); err != nil {
		return CatalogsClientGetSyncErrorDetailsResponse{}, err
	}
	return result, nil
}

// NewListByDevCenterPager - Lists catalogs for a devcenter.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - options - CatalogsClientListByDevCenterOptions contains the optional parameters for the CatalogsClient.NewListByDevCenterPager
//     method.
func (client *CatalogsClient) NewListByDevCenterPager(resourceGroupName string, devCenterName string, options *CatalogsClientListByDevCenterOptions) *runtime.Pager[CatalogsClientListByDevCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[CatalogsClientListByDevCenterResponse]{
		More: func(page CatalogsClientListByDevCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CatalogsClientListByDevCenterResponse) (CatalogsClientListByDevCenterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CatalogsClient.NewListByDevCenterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, options)
			}, nil)
			if err != nil {
				return CatalogsClientListByDevCenterResponse{}, err
			}
			return client.listByDevCenterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDevCenterCreateRequest creates the ListByDevCenter request.
func (client *CatalogsClient) listByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *CatalogsClientListByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDevCenterHandleResponse handles the ListByDevCenter response.
func (client *CatalogsClient) listByDevCenterHandleResponse(resp *http.Response) (CatalogsClientListByDevCenterResponse, error) {
	result := CatalogsClientListByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatalogListResult); err != nil {
		return CatalogsClientListByDevCenterResponse{}, err
	}
	return result, nil
}

// BeginSync - Syncs templates for a template source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogsClientBeginSyncOptions contains the optional parameters for the CatalogsClient.BeginSync method.
func (client *CatalogsClient) BeginSync(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginSyncOptions) (*runtime.Poller[CatalogsClientSyncResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.syncOperation(ctx, resourceGroupName, devCenterName, catalogName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CatalogsClientSyncResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CatalogsClientSyncResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Sync - Syncs templates for a template source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
func (client *CatalogsClient) syncOperation(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginSyncOptions) (*http.Response, error) {
	var err error
	const operationName = "CatalogsClient.BeginSync"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.syncCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// syncCreateRequest creates the Sync request.
func (client *CatalogsClient) syncCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogsClientBeginSyncOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/sync"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Partially updates a catalog.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - body - Updatable catalog properties.
//   - options - CatalogsClientBeginUpdateOptions contains the optional parameters for the CatalogsClient.BeginUpdate method.
func (client *CatalogsClient) BeginUpdate(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body CatalogUpdate, options *CatalogsClientBeginUpdateOptions) (*runtime.Poller[CatalogsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, devCenterName, catalogName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CatalogsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CatalogsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Partially updates a catalog.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01-preview
func (client *CatalogsClient) update(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body CatalogUpdate, options *CatalogsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CatalogsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, body, options)
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
func (client *CatalogsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body CatalogUpdate, options *CatalogsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
