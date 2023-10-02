//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

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

// StorageTargetClient contains the methods for the StorageTarget group.
// Don't use this type directly, use NewStorageTargetClient() instead.
type StorageTargetClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageTargetClient creates a new instance of StorageTargetClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageTargetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageTargetClient, error) {
	cl, err := arm.NewClient(moduleName+".StorageTargetClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageTargetClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginFlush - Tells the cache to write all dirty data to the Storage Target's backend storage. Client requests to this storage
// target's namespace will return errors until the flush operation completes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
//   - storageTargetName - Name of Storage Target.
//   - options - StorageTargetClientBeginFlushOptions contains the optional parameters for the StorageTargetClient.BeginFlush
//     method.
func (client *StorageTargetClient) BeginFlush(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginFlushOptions) (*runtime.Poller[StorageTargetClientFlushResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.flush(ctx, resourceGroupName, cacheName, storageTargetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageTargetClientFlushResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageTargetClientFlushResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Flush - Tells the cache to write all dirty data to the Storage Target's backend storage. Client requests to this storage
// target's namespace will return errors until the flush operation completes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *StorageTargetClient) flush(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginFlushOptions) (*http.Response, error) {
	var err error
	req, err := client.flushCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
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

// flushCreateRequest creates the Flush request.
func (client *StorageTargetClient) flushCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginFlushOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/flush"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginInvalidate - Invalidate all cached data for a storage target. Cached files are discarded and fetched from the back
// end on the next request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
//   - storageTargetName - Name of Storage Target.
//   - options - StorageTargetClientBeginInvalidateOptions contains the optional parameters for the StorageTargetClient.BeginInvalidate
//     method.
func (client *StorageTargetClient) BeginInvalidate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginInvalidateOptions) (*runtime.Poller[StorageTargetClientInvalidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.invalidate(ctx, resourceGroupName, cacheName, storageTargetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageTargetClientInvalidateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageTargetClientInvalidateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Invalidate - Invalidate all cached data for a storage target. Cached files are discarded and fetched from the back end
// on the next request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *StorageTargetClient) invalidate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginInvalidateOptions) (*http.Response, error) {
	var err error
	req, err := client.invalidateCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
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

// invalidateCreateRequest creates the Invalidate request.
func (client *StorageTargetClient) invalidateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginInvalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/invalidate"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginResume - Resumes client access to a previously suspended storage target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
//   - storageTargetName - Name of Storage Target.
//   - options - StorageTargetClientBeginResumeOptions contains the optional parameters for the StorageTargetClient.BeginResume
//     method.
func (client *StorageTargetClient) BeginResume(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginResumeOptions) (*runtime.Poller[StorageTargetClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceGroupName, cacheName, storageTargetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageTargetClientResumeResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageTargetClientResumeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Resume - Resumes client access to a previously suspended storage target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *StorageTargetClient) resume(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginResumeOptions) (*http.Response, error) {
	var err error
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
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

// resumeCreateRequest creates the Resume request.
func (client *StorageTargetClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/resume"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginSuspend - Suspends client access to a storage target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
//   - storageTargetName - Name of Storage Target.
//   - options - StorageTargetClientBeginSuspendOptions contains the optional parameters for the StorageTargetClient.BeginSuspend
//     method.
func (client *StorageTargetClient) BeginSuspend(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginSuspendOptions) (*runtime.Poller[StorageTargetClientSuspendResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.suspend(ctx, resourceGroupName, cacheName, storageTargetName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageTargetClientSuspendResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageTargetClientSuspendResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Suspend - Suspends client access to a storage target.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
func (client *StorageTargetClient) suspend(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginSuspendOptions) (*http.Response, error) {
	var err error
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
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

// suspendCreateRequest creates the Suspend request.
func (client *StorageTargetClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetClientBeginSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/suspend"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
