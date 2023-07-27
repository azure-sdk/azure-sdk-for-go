//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// ApplicationClient contains the methods for the Application group.
// Don't use this type directly, use NewApplicationClient() instead.
type ApplicationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationClient creates a new instance of ApplicationClient with the specified values.
//   - subscriptionID - The customer subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates an application resource with the specified name, description and properties. If an application resource
// with the same name exists, then it is updated with the specified description and
// properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - applicationResourceName - The identity of the application.
//   - applicationResourceDescription - Description for creating a Application resource.
//   - options - ApplicationClientCreateOptions contains the optional parameters for the ApplicationClient.Create method.
func (client *ApplicationClient) Create(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription, options *ApplicationClientCreateOptions) (ApplicationClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, applicationResourceName, applicationResourceDescription, options)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ApplicationClient) createCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, applicationResourceDescription ApplicationResourceDescription, options *ApplicationClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, applicationResourceDescription); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ApplicationClient) createHandleResponse(resp *http.Response) (ApplicationClientCreateResponse, error) {
	result := ApplicationClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescription); err != nil {
		return ApplicationClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the application resource identified by the name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - applicationResourceName - The identity of the application.
//   - options - ApplicationClientDeleteOptions contains the optional parameters for the ApplicationClient.Delete method.
func (client *ApplicationClient) Delete(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientDeleteOptions) (ApplicationClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationClientDeleteResponse{}, err
	}
	return ApplicationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the information about the application resource with the given name. The information include the description
// and other properties of the application.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - applicationResourceName - The identity of the application.
//   - options - ApplicationClientGetOptions contains the optional parameters for the ApplicationClient.Get method.
func (client *ApplicationClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientGetOptions) (ApplicationClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ApplicationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationClient) getHandleResponse(resp *http.Response) (ApplicationClientGetResponse, error) {
	result := ApplicationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescription); err != nil {
		return ApplicationClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets the information about all application resources in a given resource group. The information
// include the description and other properties of the Application.
//
// Generated from API version 2018-09-01-preview
//   - resourceGroupName - Azure resource group name
//   - options - ApplicationClientListByResourceGroupOptions contains the optional parameters for the ApplicationClient.NewListByResourceGroupPager
//     method.
func (client *ApplicationClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationClientListByResourceGroupOptions) *runtime.Pager[ApplicationClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationClientListByResourceGroupResponse]{
		More: func(page ApplicationClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationClientListByResourceGroupResponse) (ApplicationClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications"
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
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationClientListByResourceGroupResponse, error) {
	result := ApplicationClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescriptionList); err != nil {
		return ApplicationClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the information about all application resources in a given resource group. The information
// include the description and other properties of the application.
//
// Generated from API version 2018-09-01-preview
//   - options - ApplicationClientListBySubscriptionOptions contains the optional parameters for the ApplicationClient.NewListBySubscriptionPager
//     method.
func (client *ApplicationClient) NewListBySubscriptionPager(options *ApplicationClientListBySubscriptionOptions) *runtime.Pager[ApplicationClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationClientListBySubscriptionResponse]{
		More: func(page ApplicationClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationClientListBySubscriptionResponse) (ApplicationClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/applications"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationClientListBySubscriptionResponse, error) {
	result := ApplicationClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationResourceDescriptionList); err != nil {
		return ApplicationClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
