//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

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

// TargetsClient contains the methods for the Targets group.
// Don't use this type directly, use NewTargetsClient() instead.
type TargetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTargetsClient creates a new instance of TargetsClient with the specified values.
//   - subscriptionID - GUID that represents an Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTargetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TargetsClient, error) {
	cl, err := arm.NewClient(moduleName+".TargetsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TargetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Target resource that extends a tracked regional resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-27-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - target - Target resource to be created or updated.
//   - options - TargetsClientCreateOrUpdateOptions contains the optional parameters for the TargetsClient.CreateOrUpdate method.
func (client *TargetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, target Target, options *TargetsClientCreateOrUpdateOptions) (TargetsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, target, options)
	if err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TargetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, target Target, options *TargetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, target); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TargetsClient) createOrUpdateHandleResponse(resp *http.Response) (TargetsClientCreateOrUpdateResponse, error) {
	result := TargetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Target); err != nil {
		return TargetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Target resource that extends a tracked regional resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-27-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - options - TargetsClientDeleteOptions contains the optional parameters for the TargetsClient.Delete method.
func (client *TargetsClient) Delete(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *TargetsClientDeleteOptions) (TargetsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, options)
	if err != nil {
		return TargetsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientDeleteResponse{}, err
	}
	return TargetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TargetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *TargetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Target resource that extends a tracked regional resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-27-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - options - TargetsClientGetOptions contains the optional parameters for the TargetsClient.Get method.
func (client *TargetsClient) Get(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *TargetsClientGetOptions) (TargetsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, options)
	if err != nil {
		return TargetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TargetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TargetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TargetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *TargetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-27-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TargetsClient) getHandleResponse(resp *http.Response) (TargetsClientGetResponse, error) {
	result := TargetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Target); err != nil {
		return TargetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Target resources that extend a tracked regional resource.
//
// Generated from API version 2023-10-27-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - options - TargetsClientListOptions contains the optional parameters for the TargetsClient.NewListPager method.
func (client *TargetsClient) NewListPager(resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, options *TargetsClientListOptions) *runtime.Pager[TargetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TargetsClientListResponse]{
		More: func(page TargetsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TargetsClientListResponse) (TargetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TargetsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TargetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TargetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TargetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, options *TargetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-27-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TargetsClient) listHandleResponse(resp *http.Response) (TargetsClientListResponse, error) {
	result := TargetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TargetListResult); err != nil {
		return TargetsClientListResponse{}, err
	}
	return result, nil
}
