//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ApplicationGroupsClient contains the methods for the ApplicationGroups group.
// Don't use this type directly, use NewApplicationGroupsClient() instead.
type ApplicationGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationGroupsClient creates a new instance of ApplicationGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - applicationGroup - Object containing ApplicationGroup definitions.
//   - options - ApplicationGroupsClientCreateOrUpdateOptions contains the optional parameters for the ApplicationGroupsClient.CreateOrUpdate
//     method.
func (client *ApplicationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsClientCreateOrUpdateOptions) (ApplicationGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationGroupName, applicationGroup, options)
	if err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup ApplicationGroup, options *ApplicationGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, applicationGroup)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationGroupsClientCreateOrUpdateResponse, error) {
	result := ApplicationGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - options - ApplicationGroupsClientDeleteOptions contains the optional parameters for the ApplicationGroupsClient.Delete
//     method.
func (client *ApplicationGroupsClient) Delete(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientDeleteOptions) (ApplicationGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ApplicationGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ApplicationGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an application group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - options - ApplicationGroupsClientGetOptions contains the optional parameters for the ApplicationGroupsClient.Get method.
func (client *ApplicationGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientGetOptions) (ApplicationGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationGroupsClient) getHandleResponse(resp *http.Response) (ApplicationGroupsClientGetResponse, error) {
	result := ApplicationGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List applicationGroups.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ApplicationGroupsClientListByResourceGroupOptions contains the optional parameters for the ApplicationGroupsClient.NewListByResourceGroupPager
//     method.
func (client *ApplicationGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationGroupsClientListByResourceGroupOptions) *runtime.Pager[ApplicationGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationGroupsClientListByResourceGroupResponse]{
		More: func(page ApplicationGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationGroupsClientListByResourceGroupResponse) (ApplicationGroupsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups"
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
	reqQP.Set("api-version", "2023-01-30-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationGroupsClientListByResourceGroupResponse, error) {
	result := ApplicationGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List applicationGroups in subscription.
//
// Generated from API version 2023-01-30-preview
//   - options - ApplicationGroupsClientListBySubscriptionOptions contains the optional parameters for the ApplicationGroupsClient.NewListBySubscriptionPager
//     method.
func (client *ApplicationGroupsClient) NewListBySubscriptionPager(options *ApplicationGroupsClientListBySubscriptionOptions) *runtime.Pager[ApplicationGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationGroupsClientListBySubscriptionResponse]{
		More: func(page ApplicationGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationGroupsClientListBySubscriptionResponse) (ApplicationGroupsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/applicationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationGroupsClientListBySubscriptionResponse, error) {
	result := ApplicationGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroupList); err != nil {
		return ApplicationGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an applicationGroup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - options - ApplicationGroupsClientUpdateOptions contains the optional parameters for the ApplicationGroupsClient.Update
//     method.
func (client *ApplicationGroupsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientUpdateOptions) (ApplicationGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
	if err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ApplicationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *ApplicationGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ApplicationGroup != nil {
		return req, runtime.MarshalAsJSON(req, *options.ApplicationGroup)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ApplicationGroupsClient) updateHandleResponse(resp *http.Response) (ApplicationGroupsClientUpdateResponse, error) {
	result := ApplicationGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGroup); err != nil {
		return ApplicationGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
