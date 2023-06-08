//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

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

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExtensionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ExtensionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExtensionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - extensionName - The name of the machine extension.
//   - extension - Details of the Machine Extension to be created.
//   - options - ExtensionsClientBeginCreateOptions contains the optional parameters for the ExtensionsClient.BeginCreate method.
func (client *ExtensionsClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*runtime.Poller[ExtensionsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExtensionsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExtensionsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ExtensionsClient) create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extension)
}

// BeginDelete - Delete particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - extensionName - The name of the machine extension.
//   - options - ExtensionsClientBeginDeleteOptions contains the optional parameters for the ExtensionsClient.BeginDelete method.
func (client *ExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*runtime.Poller[ExtensionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExtensionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExtensionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - extensionName - The name of the machine extension.
//   - options - ExtensionsClientGetOptions contains the optional parameters for the ExtensionsClient.Get method.
func (client *ExtensionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientGetOptions) (ExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, options)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, options *ExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *http.Response) (ExtensionsClientGetResponse, error) {
	result := ExtensionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Extension); err != nil {
		return ExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByArcSettingPager - List all Extensions under ArcSetting resource.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - options - ExtensionsClientListByArcSettingOptions contains the optional parameters for the ExtensionsClient.NewListByArcSettingPager
//     method.
func (client *ExtensionsClient) NewListByArcSettingPager(resourceGroupName string, clusterName string, arcSettingName string, options *ExtensionsClientListByArcSettingOptions) *runtime.Pager[ExtensionsClientListByArcSettingResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExtensionsClientListByArcSettingResponse]{
		More: func(page ExtensionsClientListByArcSettingResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExtensionsClientListByArcSettingResponse) (ExtensionsClientListByArcSettingResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByArcSettingCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExtensionsClientListByArcSettingResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExtensionsClientListByArcSettingResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExtensionsClientListByArcSettingResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByArcSettingHandleResponse(resp)
		},
	})
}

// listByArcSettingCreateRequest creates the ListByArcSetting request.
func (client *ExtensionsClient) listByArcSettingCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ExtensionsClientListByArcSettingOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByArcSettingHandleResponse handles the ListByArcSetting response.
func (client *ExtensionsClient) listByArcSettingHandleResponse(resp *http.Response) (ExtensionsClientListByArcSettingResponse, error) {
	result := ExtensionsClientListByArcSettingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionList); err != nil {
		return ExtensionsClientListByArcSettingResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - extensionName - The name of the machine extension.
//   - extension - Details of the Machine Extension to be created.
//   - options - ExtensionsClientBeginUpdateOptions contains the optional parameters for the ExtensionsClient.BeginUpdate method.
func (client *ExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (*runtime.Poller[ExtensionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExtensionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaOriginalURI,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExtensionsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update Extension for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ExtensionsClient) update(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extension, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extension Extension, options *ExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extension)
}

// BeginUpgrade - Upgrade a particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - extensionName - The name of the machine extension.
//   - extensionUpgradeParameters - Parameters supplied to the Upgrade Extensions operation.
//   - options - ExtensionsClientBeginUpgradeOptions contains the optional parameters for the ExtensionsClient.BeginUpgrade method.
func (client *ExtensionsClient) BeginUpgrade(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extensionUpgradeParameters ExtensionUpgradeParameters, options *ExtensionsClientBeginUpgradeOptions) (*runtime.Poller[ExtensionsClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extensionUpgradeParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExtensionsClientUpgradeResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExtensionsClientUpgradeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Upgrade - Upgrade a particular Arc Extension of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ExtensionsClient) upgrade(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extensionUpgradeParameters ExtensionUpgradeParameters, options *ExtensionsClientBeginUpgradeOptions) (*http.Response, error) {
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, extensionName, extensionUpgradeParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *ExtensionsClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, extensionName string, extensionUpgradeParameters ExtensionUpgradeParameters, options *ExtensionsClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/extensions/{extensionName}/upgrade"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extensionUpgradeParameters)
}
