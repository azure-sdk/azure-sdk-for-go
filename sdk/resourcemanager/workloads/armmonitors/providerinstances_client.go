//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitors

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

// ProviderInstancesClient contains the methods for the ProviderInstances group.
// Don't use this type directly, use NewProviderInstancesClient() instead.
type ProviderInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProviderInstancesClient creates a new instance of ProviderInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProviderInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProviderInstancesClient, error) {
	cl, err := arm.NewClient(moduleName+".ProviderInstancesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProviderInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - providerInstanceName - Name of the provider instance.
//   - options - ProviderInstancesClientBeginCreateOptions contains the optional parameters for the ProviderInstancesClient.BeginCreate
//     method.
func (client *ProviderInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, providerInstanceParameter ProviderInstance, options *ProviderInstancesClientBeginCreateOptions) (*runtime.Poller[ProviderInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, monitorName, providerInstanceName, providerInstanceParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProviderInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ProviderInstancesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *ProviderInstancesClient) create(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, providerInstanceParameter ProviderInstance, options *ProviderInstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, monitorName, providerInstanceName, providerInstanceParameter, options)
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
func (client *ProviderInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, providerInstanceParameter ProviderInstance, options *ProviderInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/providerInstances/{providerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if providerInstanceName == "" {
		return nil, errors.New("parameter providerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerInstanceName}", url.PathEscape(providerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, providerInstanceParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a provider instance for the specified subscription, resource group, SAP monitor name, and resource
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - providerInstanceName - Name of the provider instance.
//   - options - ProviderInstancesClientBeginDeleteOptions contains the optional parameters for the ProviderInstancesClient.BeginDelete
//     method.
func (client *ProviderInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, options *ProviderInstancesClientBeginDeleteOptions) (*runtime.Poller[ProviderInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, providerInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProviderInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ProviderInstancesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a provider instance for the specified subscription, resource group, SAP monitor name, and resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *ProviderInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, options *ProviderInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, providerInstanceName, options)
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
func (client *ProviderInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, options *ProviderInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/providerInstances/{providerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if providerInstanceName == "" {
		return nil, errors.New("parameter providerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerInstanceName}", url.PathEscape(providerInstanceName))
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

// Get - Gets properties of a provider instance for the specified subscription, resource group, SAP monitor name, and resource
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - providerInstanceName - Name of the provider instance.
//   - options - ProviderInstancesClientGetOptions contains the optional parameters for the ProviderInstancesClient.Get method.
func (client *ProviderInstancesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, options *ProviderInstancesClientGetOptions) (ProviderInstancesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, providerInstanceName, options)
	if err != nil {
		return ProviderInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProviderInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProviderInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProviderInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, providerInstanceName string, options *ProviderInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/providerInstances/{providerInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if providerInstanceName == "" {
		return nil, errors.New("parameter providerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerInstanceName}", url.PathEscape(providerInstanceName))
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
func (client *ProviderInstancesClient) getHandleResponse(resp *http.Response) (ProviderInstancesClientGetResponse, error) {
	result := ProviderInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderInstance); err != nil {
		return ProviderInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of provider instances in the specified SAP monitor. The operations returns various properties
// of each provider instances.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - options - ProviderInstancesClientListOptions contains the optional parameters for the ProviderInstancesClient.NewListPager
//     method.
func (client *ProviderInstancesClient) NewListPager(resourceGroupName string, monitorName string, options *ProviderInstancesClientListOptions) *runtime.Pager[ProviderInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProviderInstancesClientListResponse]{
		More: func(page ProviderInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProviderInstancesClientListResponse) (ProviderInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProviderInstancesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProviderInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProviderInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProviderInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *ProviderInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/providerInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
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

// listHandleResponse handles the List response.
func (client *ProviderInstancesClient) listHandleResponse(resp *http.Response) (ProviderInstancesClientListResponse, error) {
	result := ProviderInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProviderInstanceListResult); err != nil {
		return ProviderInstancesClientListResponse{}, err
	}
	return result, nil
}
