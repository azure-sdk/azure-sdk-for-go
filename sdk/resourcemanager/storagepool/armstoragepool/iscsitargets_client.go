// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

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

// IscsiTargetsClient contains the methods for the IscsiTargets group.
// Don't use this type directly, use NewIscsiTargetsClient() instead.
type IscsiTargetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIscsiTargetsClient creates a new instance of IscsiTargetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIscsiTargetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IscsiTargetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IscsiTargetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - iscsiTargetName - The name of the iSCSI Target.
//   - iscsiTargetCreatePayload - Request payload for iSCSI Target create operation.
//   - options - IscsiTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginCreateOrUpdate
//     method.
func (client *IscsiTargetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, options *IscsiTargetsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IscsiTargetsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetCreatePayload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IscsiTargetsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IscsiTargetsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *IscsiTargetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, options *IscsiTargetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IscsiTargetsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetCreatePayload, options)
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
func (client *IscsiTargetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetCreatePayload IscsiTargetCreate, _ *IscsiTargetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, iscsiTargetCreatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - iscsiTargetName - The name of the iSCSI Target.
//   - options - IscsiTargetsClientBeginDeleteOptions contains the optional parameters for the IscsiTargetsClient.BeginDelete
//     method.
func (client *IscsiTargetsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientBeginDeleteOptions) (*runtime.Poller[IscsiTargetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IscsiTargetsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IscsiTargetsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *IscsiTargetsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "IscsiTargetsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
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
func (client *IscsiTargetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, _ *IscsiTargetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - iscsiTargetName - The name of the iSCSI Target.
//   - options - IscsiTargetsClientGetOptions contains the optional parameters for the IscsiTargetsClient.Get method.
func (client *IscsiTargetsClient) Get(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, options *IscsiTargetsClientGetOptions) (IscsiTargetsClientGetResponse, error) {
	var err error
	const operationName = "IscsiTargetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, options)
	if err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IscsiTargetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IscsiTargetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, _ *IscsiTargetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IscsiTargetsClient) getHandleResponse(resp *http.Response) (IscsiTargetsClientGetResponse, error) {
	result := IscsiTargetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IscsiTarget); err != nil {
		return IscsiTargetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDiskPoolPager - Get iSCSI Targets in a Disk pool.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - IscsiTargetsClientListByDiskPoolOptions contains the optional parameters for the IscsiTargetsClient.NewListByDiskPoolPager
//     method.
func (client *IscsiTargetsClient) NewListByDiskPoolPager(resourceGroupName string, diskPoolName string, options *IscsiTargetsClientListByDiskPoolOptions) *runtime.Pager[IscsiTargetsClientListByDiskPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[IscsiTargetsClientListByDiskPoolResponse]{
		More: func(page IscsiTargetsClientListByDiskPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IscsiTargetsClientListByDiskPoolResponse) (IscsiTargetsClientListByDiskPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IscsiTargetsClient.NewListByDiskPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDiskPoolCreateRequest(ctx, resourceGroupName, diskPoolName, options)
			}, nil)
			if err != nil {
				return IscsiTargetsClientListByDiskPoolResponse{}, err
			}
			return client.listByDiskPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDiskPoolCreateRequest creates the ListByDiskPool request.
func (client *IscsiTargetsClient) listByDiskPoolCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, _ *IscsiTargetsClientListByDiskPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDiskPoolHandleResponse handles the ListByDiskPool response.
func (client *IscsiTargetsClient) listByDiskPoolHandleResponse(resp *http.Response) (IscsiTargetsClientListByDiskPoolResponse, error) {
	result := IscsiTargetsClientListByDiskPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IscsiTargetList); err != nil {
		return IscsiTargetsClientListByDiskPoolResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - iscsiTargetName - The name of the iSCSI Target.
//   - iscsiTargetUpdatePayload - Request payload for iSCSI Target update operation.
//   - options - IscsiTargetsClientBeginUpdateOptions contains the optional parameters for the IscsiTargetsClient.BeginUpdate
//     method.
func (client *IscsiTargetsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, options *IscsiTargetsClientBeginUpdateOptions) (*runtime.Poller[IscsiTargetsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetUpdatePayload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IscsiTargetsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IscsiTargetsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an iSCSI Target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *IscsiTargetsClient) update(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, options *IscsiTargetsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IscsiTargetsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetUpdatePayload, options)
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
func (client *IscsiTargetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetUpdatePayload IscsiTargetUpdate, _ *IscsiTargetsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	if iscsiTargetName == "" {
		return nil, errors.New("parameter iscsiTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iscsiTargetName}", url.PathEscape(iscsiTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, iscsiTargetUpdatePayload); err != nil {
		return nil, err
	}
	return req, nil
}
