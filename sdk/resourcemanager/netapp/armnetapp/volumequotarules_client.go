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

// VolumeQuotaRulesClient contains the methods for the VolumeQuotaRules group.
// Don't use this type directly, use NewVolumeQuotaRulesClient() instead.
type VolumeQuotaRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVolumeQuotaRulesClient creates a new instance of VolumeQuotaRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVolumeQuotaRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumeQuotaRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VolumeQuotaRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create the specified quota rule within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - volumeQuotaRuleName - The name of volume quota rule
//   - body - Quota rule object supplied in the body of the operation.
//   - options - VolumeQuotaRulesClientBeginCreateOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginCreate
//     method.
func (client *VolumeQuotaRulesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*runtime.Poller[VolumeQuotaRulesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeQuotaRulesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeQuotaRulesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create the specified quota rule within the given volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *VolumeQuotaRulesClient) create(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeQuotaRulesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
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

// createCreateRequest creates the Create request.
func (client *VolumeQuotaRulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRule, options *VolumeQuotaRulesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
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
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete quota rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - volumeQuotaRuleName - The name of volume quota rule
//   - options - VolumeQuotaRulesClientBeginDeleteOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginDelete
//     method.
func (client *VolumeQuotaRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*runtime.Poller[VolumeQuotaRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeQuotaRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeQuotaRulesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete quota rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *VolumeQuotaRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeQuotaRulesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
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
func (client *VolumeQuotaRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
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
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get details of the specified quota rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - volumeQuotaRuleName - The name of volume quota rule
//   - options - VolumeQuotaRulesClientGetOptions contains the optional parameters for the VolumeQuotaRulesClient.Get method.
func (client *VolumeQuotaRulesClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientGetOptions) (VolumeQuotaRulesClientGetResponse, error) {
	var err error
	const operationName = "VolumeQuotaRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, options)
	if err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VolumeQuotaRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, options *VolumeQuotaRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
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
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumeQuotaRulesClient) getHandleResponse(resp *http.Response) (VolumeQuotaRulesClientGetResponse, error) {
	result := VolumeQuotaRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeQuotaRule); err != nil {
		return VolumeQuotaRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVolumePager - List all quota rules associated with the volume
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - VolumeQuotaRulesClientListByVolumeOptions contains the optional parameters for the VolumeQuotaRulesClient.NewListByVolumePager
//     method.
func (client *VolumeQuotaRulesClient) NewListByVolumePager(resourceGroupName string, accountName string, poolName string, volumeName string, options *VolumeQuotaRulesClientListByVolumeOptions) *runtime.Pager[VolumeQuotaRulesClientListByVolumeResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumeQuotaRulesClientListByVolumeResponse]{
		More: func(page VolumeQuotaRulesClientListByVolumeResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *VolumeQuotaRulesClientListByVolumeResponse) (VolumeQuotaRulesClientListByVolumeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VolumeQuotaRulesClient.NewListByVolumePager")
			req, err := client.listByVolumeCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
			if err != nil {
				return VolumeQuotaRulesClientListByVolumeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VolumeQuotaRulesClientListByVolumeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumeQuotaRulesClientListByVolumeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVolumeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVolumeCreateRequest creates the ListByVolume request.
func (client *VolumeQuotaRulesClient) listByVolumeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *VolumeQuotaRulesClientListByVolumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVolumeHandleResponse handles the ListByVolume response.
func (client *VolumeQuotaRulesClient) listByVolumeHandleResponse(resp *http.Response) (VolumeQuotaRulesClientListByVolumeResponse, error) {
	result := VolumeQuotaRulesClientListByVolumeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeQuotaRulesList); err != nil {
		return VolumeQuotaRulesClientListByVolumeResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a quota rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - volumeQuotaRuleName - The name of volume quota rule
//   - body - Quota rule object supplied in the body of the operation.
//   - options - VolumeQuotaRulesClientBeginUpdateOptions contains the optional parameters for the VolumeQuotaRulesClient.BeginUpdate
//     method.
func (client *VolumeQuotaRulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*runtime.Poller[VolumeQuotaRulesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumeQuotaRulesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VolumeQuotaRulesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch a quota rule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *VolumeQuotaRulesClient) update(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VolumeQuotaRulesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, volumeQuotaRuleName, body, options)
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
func (client *VolumeQuotaRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, volumeQuotaRuleName string, body VolumeQuotaRulePatch, options *VolumeQuotaRulesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/volumeQuotaRules/{volumeQuotaRuleName}"
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
	if volumeQuotaRuleName == "" {
		return nil, errors.New("parameter volumeQuotaRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeQuotaRuleName}", url.PathEscape(volumeQuotaRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
