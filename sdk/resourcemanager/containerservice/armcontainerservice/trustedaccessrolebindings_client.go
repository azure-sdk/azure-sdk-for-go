//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// TrustedAccessRoleBindingsClient contains the methods for the TrustedAccessRoleBindings group.
// Don't use this type directly, use NewTrustedAccessRoleBindingsClient() instead.
type TrustedAccessRoleBindingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTrustedAccessRoleBindingsClient creates a new instance of TrustedAccessRoleBindingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrustedAccessRoleBindingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrustedAccessRoleBindingsClient, error) {
	cl, err := arm.NewClient(moduleName+".TrustedAccessRoleBindingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TrustedAccessRoleBindingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a trusted access role binding
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - trustedAccessRoleBinding - A trusted access role binding
//   - options - TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.BeginCreateOrUpdate
//     method.
func (client *TrustedAccessRoleBindingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding, options *TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions) (*runtime.Poller[TrustedAccessRoleBindingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, trustedAccessRoleBinding, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TrustedAccessRoleBindingsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TrustedAccessRoleBindingsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a trusted access role binding
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *TrustedAccessRoleBindingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding, options *TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, trustedAccessRoleBinding, options)
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
func (client *TrustedAccessRoleBindingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding, options *TrustedAccessRoleBindingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if trustedAccessRoleBindingName == "" {
		return nil, errors.New("parameter trustedAccessRoleBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedAccessRoleBindingName}", url.PathEscape(trustedAccessRoleBindingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trustedAccessRoleBinding); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a trusted access role binding.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - options - TrustedAccessRoleBindingsClientBeginDeleteOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.BeginDelete
//     method.
func (client *TrustedAccessRoleBindingsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientBeginDeleteOptions) (*runtime.Poller[TrustedAccessRoleBindingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TrustedAccessRoleBindingsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TrustedAccessRoleBindingsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a trusted access role binding.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *TrustedAccessRoleBindingsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, options)
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
func (client *TrustedAccessRoleBindingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if trustedAccessRoleBindingName == "" {
		return nil, errors.New("parameter trustedAccessRoleBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedAccessRoleBindingName}", url.PathEscape(trustedAccessRoleBindingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a trusted access role binding.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - options - TrustedAccessRoleBindingsClientGetOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.Get
//     method.
func (client *TrustedAccessRoleBindingsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientGetOptions) (TrustedAccessRoleBindingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, options)
	if err != nil {
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TrustedAccessRoleBindingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if trustedAccessRoleBindingName == "" {
		return nil, errors.New("parameter trustedAccessRoleBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trustedAccessRoleBindingName}", url.PathEscape(trustedAccessRoleBindingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TrustedAccessRoleBindingsClient) getHandleResponse(resp *http.Response) (TrustedAccessRoleBindingsClientGetResponse, error) {
	result := TrustedAccessRoleBindingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedAccessRoleBinding); err != nil {
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List trusted access role bindings.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - TrustedAccessRoleBindingsClientListOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.NewListPager
//     method.
func (client *TrustedAccessRoleBindingsClient) NewListPager(resourceGroupName string, resourceName string, options *TrustedAccessRoleBindingsClientListOptions) *runtime.Pager[TrustedAccessRoleBindingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrustedAccessRoleBindingsClientListResponse]{
		More: func(page TrustedAccessRoleBindingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrustedAccessRoleBindingsClientListResponse) (TrustedAccessRoleBindingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TrustedAccessRoleBindingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TrustedAccessRoleBindingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TrustedAccessRoleBindingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TrustedAccessRoleBindingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *TrustedAccessRoleBindingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TrustedAccessRoleBindingsClient) listHandleResponse(resp *http.Response) (TrustedAccessRoleBindingsClientListResponse, error) {
	result := TrustedAccessRoleBindingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedAccessRoleBindingListResult); err != nil {
		return TrustedAccessRoleBindingsClientListResponse{}, err
	}
	return result, nil
}
