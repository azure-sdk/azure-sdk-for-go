//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// ApplicationLiveViewsClient contains the methods for the ApplicationLiveViews group.
// Don't use this type directly, use NewApplicationLiveViewsClient() instead.
type ApplicationLiveViewsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationLiveViewsClient creates a new instance of ApplicationLiveViewsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationLiveViewsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationLiveViewsClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationLiveViewsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationLiveViewsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default Application Live View or update the existing Application Live View.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationLiveViewName - The name of Application Live View.
//   - applicationLiveViewResource - Parameters for the update operation
//   - options - ApplicationLiveViewsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationLiveViewsClient.BeginCreateOrUpdate
//     method.
func (client *ApplicationLiveViewsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource ApplicationLiveViewResource, options *ApplicationLiveViewsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationLiveViewsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, applicationLiveViewName, applicationLiveViewResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ApplicationLiveViewsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationLiveViewsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create the default Application Live View or update the existing Application Live View.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *ApplicationLiveViewsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource ApplicationLiveViewResource, options *ApplicationLiveViewsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, applicationLiveViewName, applicationLiveViewResource, options)
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
func (client *ApplicationLiveViewsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource ApplicationLiveViewResource, options *ApplicationLiveViewsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationLiveViewName == "" {
		return nil, errors.New("parameter applicationLiveViewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationLiveViewName}", url.PathEscape(applicationLiveViewName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, applicationLiveViewResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Disable the default Application Live View.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationLiveViewName - The name of Application Live View.
//   - options - ApplicationLiveViewsClientBeginDeleteOptions contains the optional parameters for the ApplicationLiveViewsClient.BeginDelete
//     method.
func (client *ApplicationLiveViewsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *ApplicationLiveViewsClientBeginDeleteOptions) (*runtime.Poller[ApplicationLiveViewsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, applicationLiveViewName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ApplicationLiveViewsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationLiveViewsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Disable the default Application Live View.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *ApplicationLiveViewsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *ApplicationLiveViewsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, applicationLiveViewName, options)
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
func (client *ApplicationLiveViewsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *ApplicationLiveViewsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationLiveViewName == "" {
		return nil, errors.New("parameter applicationLiveViewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationLiveViewName}", url.PathEscape(applicationLiveViewName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Application Live and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - applicationLiveViewName - The name of Application Live View.
//   - options - ApplicationLiveViewsClientGetOptions contains the optional parameters for the ApplicationLiveViewsClient.Get
//     method.
func (client *ApplicationLiveViewsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *ApplicationLiveViewsClientGetOptions) (ApplicationLiveViewsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, applicationLiveViewName, options)
	if err != nil {
		return ApplicationLiveViewsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationLiveViewsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationLiveViewsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationLiveViewsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, options *ApplicationLiveViewsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationLiveViewName == "" {
		return nil, errors.New("parameter applicationLiveViewName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationLiveViewName}", url.PathEscape(applicationLiveViewName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationLiveViewsClient) getHandleResponse(resp *http.Response) (ApplicationLiveViewsClientGetResponse, error) {
	result := ApplicationLiveViewsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationLiveViewResource); err != nil {
		return ApplicationLiveViewsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - ApplicationLiveViewsClientListOptions contains the optional parameters for the ApplicationLiveViewsClient.NewListPager
//     method.
func (client *ApplicationLiveViewsClient) NewListPager(resourceGroupName string, serviceName string, options *ApplicationLiveViewsClientListOptions) *runtime.Pager[ApplicationLiveViewsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationLiveViewsClientListResponse]{
		More: func(page ApplicationLiveViewsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationLiveViewsClientListResponse) (ApplicationLiveViewsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationLiveViewsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationLiveViewsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationLiveViewsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationLiveViewsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApplicationLiveViewsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationLiveViewsClient) listHandleResponse(resp *http.Response) (ApplicationLiveViewsClientListResponse, error) {
	result := ApplicationLiveViewsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationLiveViewResourceCollection); err != nil {
		return ApplicationLiveViewsClientListResponse{}, err
	}
	return result, nil
}
