//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// SubvolumesClient contains the methods for the Subvolumes group.
// Don't use this type directly, use NewSubvolumesClient() instead.
type SubvolumesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubvolumesClient creates a new instance of SubvolumesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubvolumesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubvolumesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubvolumesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a subvolume in the path or clones the subvolume mentioned in the parentPath
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - subvolumeName - The name of the subvolume.
//   - body - Subvolume object supplied in the body of the operation.
//   - options - SubvolumesClientBeginCreateOptions contains the optional parameters for the SubvolumesClient.BeginCreate method.
func (client *SubvolumesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*runtime.Poller[SubvolumesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubvolumesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubvolumesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a subvolume in the path or clones the subvolume mentioned in the parentPath
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
func (client *SubvolumesClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SubvolumesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *SubvolumesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumeInfo, options *SubvolumesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - subvolumeName - The name of the subvolume.
//   - options - SubvolumesClientBeginDeleteOptions contains the optional parameters for the SubvolumesClient.BeginDelete method.
func (client *SubvolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*runtime.Poller[SubvolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubvolumesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubvolumesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
func (client *SubvolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SubvolumesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
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
func (client *SubvolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the path associated with the subvolumeName provided
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - subvolumeName - The name of the subvolume.
//   - options - SubvolumesClientGetOptions contains the optional parameters for the SubvolumesClient.Get method.
func (client *SubvolumesClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientGetOptions) (SubvolumesClientGetResponse, error) {
	var err error
	const operationName = "SubvolumesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
	if err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubvolumesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SubvolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubvolumesClient) getHandleResponse(resp *http.Response) (SubvolumesClientGetResponse, error) {
	result := SubvolumesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubvolumeInfo); err != nil {
		return SubvolumesClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetMetadata - Get details of the specified subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - subvolumeName - The name of the subvolume.
//   - options - SubvolumesClientBeginGetMetadataOptions contains the optional parameters for the SubvolumesClient.BeginGetMetadata
//     method.
func (client *SubvolumesClient) BeginGetMetadata(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*runtime.Poller[SubvolumesClientGetMetadataResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getMetadata(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubvolumesClientGetMetadataResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubvolumesClientGetMetadataResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GetMetadata - Get details of the specified subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
func (client *SubvolumesClient) getMetadata(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*http.Response, error) {
	var err error
	const operationName = "SubvolumesClient.BeginGetMetadata"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMetadataCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, options)
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

// getMetadataCreateRequest creates the GetMetadata request.
func (client *SubvolumesClient) getMetadataCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, options *SubvolumesClientBeginGetMetadataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}/getMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByVolumePager - Returns a list of the subvolumes in the volume
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - SubvolumesClientListByVolumeOptions contains the optional parameters for the SubvolumesClient.NewListByVolumePager
//     method.
func (client *SubvolumesClient) NewListByVolumePager(resourceGroupName string, accountName string, poolName string, volumeName string, options *SubvolumesClientListByVolumeOptions) *runtime.Pager[SubvolumesClientListByVolumeResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubvolumesClientListByVolumeResponse]{
		More: func(page SubvolumesClientListByVolumeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubvolumesClientListByVolumeResponse) (SubvolumesClientListByVolumeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SubvolumesClient.NewListByVolumePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVolumeCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			}, nil)
			if err != nil {
				return SubvolumesClientListByVolumeResponse{}, err
			}
			return client.listByVolumeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVolumeCreateRequest creates the ListByVolume request.
func (client *SubvolumesClient) listByVolumeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *SubvolumesClientListByVolumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVolumeHandleResponse handles the ListByVolume response.
func (client *SubvolumesClient) listByVolumeHandleResponse(resp *http.Response) (SubvolumesClientListByVolumeResponse, error) {
	result := SubvolumesClientListByVolumeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubvolumesList); err != nil {
		return SubvolumesClientListByVolumeResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - subvolumeName - The name of the subvolume.
//   - body - Subvolume object supplied in the body of the operation.
//   - options - SubvolumesClientBeginUpdateOptions contains the optional parameters for the SubvolumesClient.BeginUpdate method.
func (client *SubvolumesClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*runtime.Poller[SubvolumesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SubvolumesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SubvolumesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch a subvolume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
func (client *SubvolumesClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SubvolumesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, subvolumeName, body, options)
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
func (client *SubvolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, subvolumeName string, body SubvolumePatchRequest, options *SubvolumesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/subvolumes/{subvolumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	if subvolumeName == "" {
		return nil, errors.New("parameter subvolumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subvolumeName}", url.PathEscape(subvolumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
