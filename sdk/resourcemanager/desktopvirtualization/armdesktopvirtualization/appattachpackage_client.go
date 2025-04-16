// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// AppAttachPackageClient contains the methods for the AppAttachPackage group.
// Don't use this type directly, use NewAppAttachPackageClient() instead.
type AppAttachPackageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAppAttachPackageClient creates a new instance of AppAttachPackageClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAppAttachPackageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AppAttachPackageClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AppAttachPackageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an App Attach package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appAttachPackageName - The name of the App Attach package arm object
//   - appAttachPackage - Object containing App Attach Package definitions.
//   - options - AppAttachPackageClientCreateOrUpdateOptions contains the optional parameters for the AppAttachPackageClient.CreateOrUpdate
//     method.
func (client *AppAttachPackageClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, appAttachPackageName string, appAttachPackage AppAttachPackage, options *AppAttachPackageClientCreateOrUpdateOptions) (AppAttachPackageClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AppAttachPackageClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, appAttachPackageName, appAttachPackage, options)
	if err != nil {
		return AppAttachPackageClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppAttachPackageClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AppAttachPackageClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AppAttachPackageClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, appAttachPackageName string, appAttachPackage AppAttachPackage, _ *AppAttachPackageClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/appAttachPackages/{appAttachPackageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appAttachPackageName == "" {
		return nil, errors.New("parameter appAttachPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appAttachPackageName}", url.PathEscape(appAttachPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, appAttachPackage); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AppAttachPackageClient) createOrUpdateHandleResponse(resp *http.Response) (AppAttachPackageClientCreateOrUpdateResponse, error) {
	result := AppAttachPackageClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackage); err != nil {
		return AppAttachPackageClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove an App Attach Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appAttachPackageName - The name of the App Attach package arm object
//   - options - AppAttachPackageClientDeleteOptions contains the optional parameters for the AppAttachPackageClient.Delete method.
func (client *AppAttachPackageClient) Delete(ctx context.Context, resourceGroupName string, appAttachPackageName string, options *AppAttachPackageClientDeleteOptions) (AppAttachPackageClientDeleteResponse, error) {
	var err error
	const operationName = "AppAttachPackageClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, appAttachPackageName, options)
	if err != nil {
		return AppAttachPackageClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppAttachPackageClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AppAttachPackageClientDeleteResponse{}, err
	}
	return AppAttachPackageClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AppAttachPackageClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, appAttachPackageName string, options *AppAttachPackageClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/appAttachPackages/{appAttachPackageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appAttachPackageName == "" {
		return nil, errors.New("parameter appAttachPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appAttachPackageName}", url.PathEscape(appAttachPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an app attach package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appAttachPackageName - The name of the App Attach package arm object
//   - options - AppAttachPackageClientGetOptions contains the optional parameters for the AppAttachPackageClient.Get method.
func (client *AppAttachPackageClient) Get(ctx context.Context, resourceGroupName string, appAttachPackageName string, options *AppAttachPackageClientGetOptions) (AppAttachPackageClientGetResponse, error) {
	var err error
	const operationName = "AppAttachPackageClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, appAttachPackageName, options)
	if err != nil {
		return AppAttachPackageClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppAttachPackageClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppAttachPackageClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AppAttachPackageClient) getCreateRequest(ctx context.Context, resourceGroupName string, appAttachPackageName string, _ *AppAttachPackageClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/appAttachPackages/{appAttachPackageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appAttachPackageName == "" {
		return nil, errors.New("parameter appAttachPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appAttachPackageName}", url.PathEscape(appAttachPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AppAttachPackageClient) getHandleResponse(resp *http.Response) (AppAttachPackageClientGetResponse, error) {
	result := AppAttachPackageClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackage); err != nil {
		return AppAttachPackageClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List App Attach packages in resource group.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - AppAttachPackageClientListByResourceGroupOptions contains the optional parameters for the AppAttachPackageClient.NewListByResourceGroupPager
//     method.
func (client *AppAttachPackageClient) NewListByResourceGroupPager(resourceGroupName string, options *AppAttachPackageClientListByResourceGroupOptions) *runtime.Pager[AppAttachPackageClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AppAttachPackageClientListByResourceGroupResponse]{
		More: func(page AppAttachPackageClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppAttachPackageClientListByResourceGroupResponse) (AppAttachPackageClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AppAttachPackageClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AppAttachPackageClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AppAttachPackageClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AppAttachPackageClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/appAttachPackages"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AppAttachPackageClient) listByResourceGroupHandleResponse(resp *http.Response) (AppAttachPackageClientListByResourceGroupResponse, error) {
	result := AppAttachPackageClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackageList); err != nil {
		return AppAttachPackageClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List App Attach packages in subscription.
//
// Generated from API version 2024-11-01-preview
//   - options - AppAttachPackageClientListBySubscriptionOptions contains the optional parameters for the AppAttachPackageClient.NewListBySubscriptionPager
//     method.
func (client *AppAttachPackageClient) NewListBySubscriptionPager(options *AppAttachPackageClientListBySubscriptionOptions) *runtime.Pager[AppAttachPackageClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AppAttachPackageClientListBySubscriptionResponse]{
		More: func(page AppAttachPackageClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppAttachPackageClientListBySubscriptionResponse) (AppAttachPackageClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AppAttachPackageClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AppAttachPackageClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AppAttachPackageClient) listBySubscriptionCreateRequest(ctx context.Context, options *AppAttachPackageClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/appAttachPackages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AppAttachPackageClient) listBySubscriptionHandleResponse(resp *http.Response) (AppAttachPackageClientListBySubscriptionResponse, error) {
	result := AppAttachPackageClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackageList); err != nil {
		return AppAttachPackageClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update an App Attach Package
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appAttachPackageName - The name of the App Attach package arm object
//   - options - AppAttachPackageClientUpdateOptions contains the optional parameters for the AppAttachPackageClient.Update method.
func (client *AppAttachPackageClient) Update(ctx context.Context, resourceGroupName string, appAttachPackageName string, options *AppAttachPackageClientUpdateOptions) (AppAttachPackageClientUpdateResponse, error) {
	var err error
	const operationName = "AppAttachPackageClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, appAttachPackageName, options)
	if err != nil {
		return AppAttachPackageClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppAttachPackageClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppAttachPackageClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AppAttachPackageClient) updateCreateRequest(ctx context.Context, resourceGroupName string, appAttachPackageName string, options *AppAttachPackageClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/appAttachPackages/{appAttachPackageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appAttachPackageName == "" {
		return nil, errors.New("parameter appAttachPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appAttachPackageName}", url.PathEscape(appAttachPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.AppAttachPackagePatch != nil {
		if err := runtime.MarshalAsJSON(req, *options.AppAttachPackagePatch); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AppAttachPackageClient) updateHandleResponse(resp *http.Response) (AppAttachPackageClientUpdateResponse, error) {
	result := AppAttachPackageClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackage); err != nil {
		return AppAttachPackageClientUpdateResponse{}, err
	}
	return result, nil
}
