//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconnectedcache

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

// EnterpriseCustomerOperationsClient contains the methods for the EnterpriseCustomerOperations group.
// Don't use this type directly, use NewEnterpriseCustomerOperationsClient() instead.
type EnterpriseCustomerOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEnterpriseCustomerOperationsClient creates a new instance of EnterpriseCustomerOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnterpriseCustomerOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnterpriseCustomerOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnterpriseCustomerOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a cacheNodes with the specified create parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - customerResourceName - Name of the Customer resource
//   - resource - Resource create parameters.
//   - options - EnterpriseCustomerOperationsClientBeginCreateOrUpdateOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.BeginCreateOrUpdate
//     method.
func (client *EnterpriseCustomerOperationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, customerResourceName string, resource EnterprisePreviewResource, options *EnterpriseCustomerOperationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[EnterpriseCustomerOperationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, customerResourceName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[EnterpriseCustomerOperationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[EnterpriseCustomerOperationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a cacheNodes with the specified create parameters
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *EnterpriseCustomerOperationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, customerResourceName string, resource EnterprisePreviewResource, options *EnterpriseCustomerOperationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "EnterpriseCustomerOperationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, customerResourceName, resource, options)
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
func (client *EnterpriseCustomerOperationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, customerResourceName string, resource EnterprisePreviewResource, options *EnterpriseCustomerOperationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedCache/enterpriseCustomers/{customerResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customerResourceName == "" {
		return nil, errors.New("parameter customerResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerResourceName}", url.PathEscape(customerResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes an existing customer Enterprise resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - customerResourceName - Name of the Customer resource
//   - options - EnterpriseCustomerOperationsClientDeleteOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.Delete
//     method.
func (client *EnterpriseCustomerOperationsClient) Delete(ctx context.Context, resourceGroupName string, customerResourceName string, options *EnterpriseCustomerOperationsClientDeleteOptions) (EnterpriseCustomerOperationsClientDeleteResponse, error) {
	var err error
	const operationName = "EnterpriseCustomerOperationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customerResourceName, options)
	if err != nil {
		return EnterpriseCustomerOperationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnterpriseCustomerOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EnterpriseCustomerOperationsClientDeleteResponse{}, err
	}
	return EnterpriseCustomerOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnterpriseCustomerOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customerResourceName string, options *EnterpriseCustomerOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedCache/enterpriseCustomers/{customerResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customerResourceName == "" {
		return nil, errors.New("parameter customerResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerResourceName}", url.PathEscape(customerResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the properties of a Enterprise customer
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - customerResourceName - Name of the Customer resource
//   - options - EnterpriseCustomerOperationsClientGetOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.Get
//     method.
func (client *EnterpriseCustomerOperationsClient) Get(ctx context.Context, resourceGroupName string, customerResourceName string, options *EnterpriseCustomerOperationsClientGetOptions) (EnterpriseCustomerOperationsClientGetResponse, error) {
	var err error
	const operationName = "EnterpriseCustomerOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, customerResourceName, options)
	if err != nil {
		return EnterpriseCustomerOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnterpriseCustomerOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnterpriseCustomerOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EnterpriseCustomerOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, customerResourceName string, options *EnterpriseCustomerOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedCache/enterpriseCustomers/{customerResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customerResourceName == "" {
		return nil, errors.New("parameter customerResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerResourceName}", url.PathEscape(customerResourceName))
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
func (client *EnterpriseCustomerOperationsClient) getHandleResponse(resp *http.Response) (EnterpriseCustomerOperationsClientGetResponse, error) {
	result := EnterpriseCustomerOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePreviewResource); err != nil {
		return EnterpriseCustomerOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieves the properties of all ConnectedCache enterpriseCustomers
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - EnterpriseCustomerOperationsClientListByResourceGroupOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.NewListByResourceGroupPager
//     method.
func (client *EnterpriseCustomerOperationsClient) NewListByResourceGroupPager(resourceGroupName string, options *EnterpriseCustomerOperationsClientListByResourceGroupOptions) *runtime.Pager[EnterpriseCustomerOperationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnterpriseCustomerOperationsClientListByResourceGroupResponse]{
		More: func(page EnterpriseCustomerOperationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnterpriseCustomerOperationsClientListByResourceGroupResponse) (EnterpriseCustomerOperationsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EnterpriseCustomerOperationsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return EnterpriseCustomerOperationsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *EnterpriseCustomerOperationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *EnterpriseCustomerOperationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedCache/enterpriseCustomers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *EnterpriseCustomerOperationsClient) listByResourceGroupHandleResponse(resp *http.Response) (EnterpriseCustomerOperationsClientListByResourceGroupResponse, error) {
	result := EnterpriseCustomerOperationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePreviewResourceListResult); err != nil {
		return EnterpriseCustomerOperationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieves the properties of all ConnectedCaches
//
// Generated from API version 2023-05-01-preview
//   - options - EnterpriseCustomerOperationsClientListBySubscriptionOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.NewListBySubscriptionPager
//     method.
func (client *EnterpriseCustomerOperationsClient) NewListBySubscriptionPager(options *EnterpriseCustomerOperationsClientListBySubscriptionOptions) *runtime.Pager[EnterpriseCustomerOperationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnterpriseCustomerOperationsClientListBySubscriptionResponse]{
		More: func(page EnterpriseCustomerOperationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnterpriseCustomerOperationsClientListBySubscriptionResponse) (EnterpriseCustomerOperationsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EnterpriseCustomerOperationsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return EnterpriseCustomerOperationsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *EnterpriseCustomerOperationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *EnterpriseCustomerOperationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedCache/enterpriseCustomers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *EnterpriseCustomerOperationsClient) listBySubscriptionHandleResponse(resp *http.Response) (EnterpriseCustomerOperationsClientListBySubscriptionResponse, error) {
	result := EnterpriseCustomerOperationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePreviewResourceListResult); err != nil {
		return EnterpriseCustomerOperationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - updates an existing enterpriseCustomers
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - customerResourceName - Name of the Customer resource
//   - properties - The resource properties to be updated.
//   - options - EnterpriseCustomerOperationsClientUpdateOptions contains the optional parameters for the EnterpriseCustomerOperationsClient.Update
//     method.
func (client *EnterpriseCustomerOperationsClient) Update(ctx context.Context, resourceGroupName string, customerResourceName string, properties PatchResource, options *EnterpriseCustomerOperationsClientUpdateOptions) (EnterpriseCustomerOperationsClientUpdateResponse, error) {
	var err error
	const operationName = "EnterpriseCustomerOperationsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, customerResourceName, properties, options)
	if err != nil {
		return EnterpriseCustomerOperationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnterpriseCustomerOperationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnterpriseCustomerOperationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *EnterpriseCustomerOperationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, customerResourceName string, properties PatchResource, options *EnterpriseCustomerOperationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedCache/enterpriseCustomers/{customerResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customerResourceName == "" {
		return nil, errors.New("parameter customerResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerResourceName}", url.PathEscape(customerResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *EnterpriseCustomerOperationsClient) updateHandleResponse(resp *http.Response) (EnterpriseCustomerOperationsClientUpdateResponse, error) {
	result := EnterpriseCustomerOperationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnterprisePreviewResource); err != nil {
		return EnterpriseCustomerOperationsClientUpdateResponse{}, err
	}
	return result, nil
}
