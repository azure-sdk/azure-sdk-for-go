// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple8000series

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginClone - Clones the backup element as a new volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - backupName - The backup name.
//   - backupElementName - The backup element name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The clone request object.
//   - options - BackupsClientBeginCloneOptions contains the optional parameters for the BackupsClient.BeginClone method.
func (client *BackupsClient) BeginClone(ctx context.Context, deviceName string, backupName string, backupElementName string, resourceGroupName string, managerName string, parameters CloneRequest, options *BackupsClientBeginCloneOptions) (*runtime.Poller[BackupsClientCloneResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.clone(ctx, deviceName, backupName, backupElementName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientCloneResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientCloneResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Clone - Clones the backup element as a new volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *BackupsClient) clone(ctx context.Context, deviceName string, backupName string, backupElementName string, resourceGroupName string, managerName string, parameters CloneRequest, options *BackupsClientBeginCloneOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginClone"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cloneCreateRequest(ctx, deviceName, backupName, backupElementName, resourceGroupName, managerName, parameters, options)
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

// cloneCreateRequest creates the Clone request.
func (client *BackupsClient) cloneCreateRequest(ctx context.Context, deviceName string, backupName string, backupElementName string, resourceGroupName string, managerName string, parameters CloneRequest, _ *BackupsClientBeginCloneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/elements/{backupElementName}/clone"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", backupName)
	urlPath = strings.ReplaceAll(urlPath, "{backupElementName}", backupElementName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the backup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - backupName - The backup name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
func (client *BackupsClient) BeginDelete(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginDeleteOptions) (*runtime.Poller[BackupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, backupName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the backup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *BackupsClient) deleteOperation(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, backupName, resourceGroupName, managerName, options)
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
func (client *BackupsClient) deleteCreateRequest(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, _ *BackupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", backupName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// NewListByDevicePager - Retrieves all the backups in a device.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientListByDeviceOptions contains the optional parameters for the BackupsClient.NewListByDevicePager
//     method.
func (client *BackupsClient) NewListByDevicePager(deviceName string, resourceGroupName string, managerName string, options *BackupsClientListByDeviceOptions) *runtime.Pager[BackupsClientListByDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupsClientListByDeviceResponse]{
		More: func(page BackupsClientListByDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupsClientListByDeviceResponse) (BackupsClientListByDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupsClient.NewListByDevicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
			}, nil)
			if err != nil {
				return BackupsClientListByDeviceResponse{}, err
			}
			return client.listByDeviceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *BackupsClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *BackupsClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *BackupsClient) listByDeviceHandleResponse(resp *http.Response) (BackupsClientListByDeviceResponse, error) {
	result := BackupsClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupList); err != nil {
		return BackupsClientListByDeviceResponse{}, err
	}
	return result, nil
}

// BeginRestore - Restores the backup on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - backupName - The backupSet name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientBeginRestoreOptions contains the optional parameters for the BackupsClient.BeginRestore method.
func (client *BackupsClient) BeginRestore(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginRestoreOptions) (*runtime.Poller[BackupsClientRestoreResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restore(ctx, deviceName, backupName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientRestoreResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientRestoreResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restore - Restores the backup on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *BackupsClient) restore(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginRestoreOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginRestore"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restoreCreateRequest(ctx, deviceName, backupName, resourceGroupName, managerName, options)
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

// restoreCreateRequest creates the Restore request.
func (client *BackupsClient) restoreCreateRequest(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, _ *BackupsClientBeginRestoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/restore"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", backupName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
