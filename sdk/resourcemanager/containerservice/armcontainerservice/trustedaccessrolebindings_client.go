//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TrustedAccessRoleBindingsClient contains the methods for the TrustedAccessRoleBindings group.
// Don't use this type directly, use NewTrustedAccessRoleBindingsClient() instead.
type TrustedAccessRoleBindingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTrustedAccessRoleBindingsClient creates a new instance of TrustedAccessRoleBindingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrustedAccessRoleBindingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrustedAccessRoleBindingsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &TrustedAccessRoleBindingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a trusted access role binding
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - trustedAccessRoleBinding - A trusted access role binding
//   - options - TrustedAccessRoleBindingsClientCreateOrUpdateOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.CreateOrUpdate
//     method.
func (client *TrustedAccessRoleBindingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding, options *TrustedAccessRoleBindingsClientCreateOrUpdateOptions) (TrustedAccessRoleBindingsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, trustedAccessRoleBinding, options)
	if err != nil {
		return TrustedAccessRoleBindingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrustedAccessRoleBindingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrustedAccessRoleBindingsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TrustedAccessRoleBindingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, trustedAccessRoleBinding TrustedAccessRoleBinding, options *TrustedAccessRoleBindingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, trustedAccessRoleBinding)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TrustedAccessRoleBindingsClient) createOrUpdateHandleResponse(resp *http.Response) (TrustedAccessRoleBindingsClientCreateOrUpdateResponse, error) {
	result := TrustedAccessRoleBindingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedAccessRoleBinding); err != nil {
		return TrustedAccessRoleBindingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a trusted access role binding.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - options - TrustedAccessRoleBindingsClientDeleteOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.Delete
//     method.
func (client *TrustedAccessRoleBindingsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientDeleteOptions) (TrustedAccessRoleBindingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, options)
	if err != nil {
		return TrustedAccessRoleBindingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrustedAccessRoleBindingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TrustedAccessRoleBindingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return TrustedAccessRoleBindingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TrustedAccessRoleBindingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a trusted access role binding.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - trustedAccessRoleBindingName - The name of trusted access role binding.
//   - options - TrustedAccessRoleBindingsClientGetOptions contains the optional parameters for the TrustedAccessRoleBindingsClient.Get
//     method.
func (client *TrustedAccessRoleBindingsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientGetOptions) (TrustedAccessRoleBindingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, trustedAccessRoleBindingName, options)
	if err != nil {
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TrustedAccessRoleBindingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TrustedAccessRoleBindingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TrustedAccessRoleBindingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, trustedAccessRoleBindingName string, options *TrustedAccessRoleBindingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/trustedAccessRoleBindings/{trustedAccessRoleBindingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
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
// Generated from API version 2022-10-02-preview
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
			resp, err := client.pl.Do(req)
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-02-preview")
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
