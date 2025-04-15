// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

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

// CapabilitiesClient contains the methods for the Capabilities group.
// Don't use this type directly, use NewCapabilitiesClient() instead.
type CapabilitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCapabilitiesClient creates a new instance of CapabilitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCapabilitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapabilitiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CapabilitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Capability resource that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parentProviderNamespace - The parent resource provider namespace.
//   - parentResourceType - The parent resource type.
//   - parentResourceName - The parent resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - resource - Capability resource to be created or updated.
//   - options - CapabilitiesClientCreateOrUpdateOptions contains the optional parameters for the CapabilitiesClient.CreateOrUpdate
//     method.
func (client *CapabilitiesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, resource Capability, options *CapabilitiesClientCreateOrUpdateOptions) (CapabilitiesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "CapabilitiesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, resource, options)
	if err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CapabilitiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, resource Capability, _ *CapabilitiesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
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
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CapabilitiesClient) createOrUpdateHandleResponse(resp *http.Response) (CapabilitiesClientCreateOrUpdateResponse, error) {
	result := CapabilitiesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capability); err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Capability that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parentProviderNamespace - The parent resource provider namespace.
//   - parentResourceType - The parent resource type.
//   - parentResourceName - The parent resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - options - CapabilitiesClientDeleteOptions contains the optional parameters for the CapabilitiesClient.Delete method.
func (client *CapabilitiesClient) Delete(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientDeleteOptions) (CapabilitiesClientDeleteResponse, error) {
	var err error
	const operationName = "CapabilitiesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, options)
	if err != nil {
		return CapabilitiesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CapabilitiesClientDeleteResponse{}, err
	}
	return CapabilitiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapabilitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, _ *CapabilitiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
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
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Capability resource that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parentProviderNamespace - The parent resource provider namespace.
//   - parentResourceType - The parent resource type.
//   - parentResourceName - The parent resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - options - CapabilitiesClientGetOptions contains the optional parameters for the CapabilitiesClient.Get method.
func (client *CapabilitiesClient) Get(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientGetOptions) (CapabilitiesClientGetResponse, error) {
	var err error
	const operationName = "CapabilitiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, options)
	if err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CapabilitiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CapabilitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, _ *CapabilitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
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
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapabilitiesClient) getHandleResponse(resp *http.Response) (CapabilitiesClientGetResponse, error) {
	result := CapabilitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capability); err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Capability resources that extend a Target resource.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parentProviderNamespace - The parent resource provider namespace.
//   - parentResourceType - The parent resource type.
//   - parentResourceName - The parent resource name.
//   - targetName - String that represents a Target resource name.
//   - options - CapabilitiesClientListOptions contains the optional parameters for the CapabilitiesClient.NewListPager method.
func (client *CapabilitiesClient) NewListPager(resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *CapabilitiesClientListOptions) *runtime.Pager[CapabilitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapabilitiesClientListResponse]{
		More: func(page CapabilitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapabilitiesClientListResponse) (CapabilitiesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CapabilitiesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, options)
			}, nil)
			if err != nil {
				return CapabilitiesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CapabilitiesClient) listCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *CapabilitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities"
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
	reqQP.Set("api-version", "2025-01-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapabilitiesClient) listHandleResponse(resp *http.Response) (CapabilitiesClientListResponse, error) {
	result := CapabilitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityListResult); err != nil {
		return CapabilitiesClientListResponse{}, err
	}
	return result, nil
}
