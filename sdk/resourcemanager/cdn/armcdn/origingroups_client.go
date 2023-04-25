//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn

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

// OriginGroupsClient contains the methods for the OriginGroups group.
// Don't use this type directly, use NewOriginGroupsClient() instead.
type OriginGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOriginGroupsClient creates a new instance of OriginGroupsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOriginGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OriginGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".OriginGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OriginGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new origin group within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originGroupName - Name of the origin group which is unique within the endpoint.
//   - originGroup - Origin group properties
//   - options - OriginGroupsClientBeginCreateOptions contains the optional parameters for the OriginGroupsClient.BeginCreate
//     method.
func (client *OriginGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsClientBeginCreateOptions) (*runtime.Poller[OriginGroupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroup, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginGroupsClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginGroupsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a new origin group within the specified endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
func (client *OriginGroupsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroup, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OriginGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroup OriginGroup, options *OriginGroupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, originGroup)
}

// BeginDelete - Deletes an existing origin group within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originGroupName - Name of the origin group which is unique within the endpoint.
//   - options - OriginGroupsClientBeginDeleteOptions contains the optional parameters for the OriginGroupsClient.BeginDelete
//     method.
func (client *OriginGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsClientBeginDeleteOptions) (*runtime.Poller[OriginGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginGroupsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing origin group within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
func (client *OriginGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OriginGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing origin group within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originGroupName - Name of the origin group which is unique within the endpoint.
//   - options - OriginGroupsClientGetOptions contains the optional parameters for the OriginGroupsClient.Get method.
func (client *OriginGroupsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsClientGetOptions) (OriginGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, options)
	if err != nil {
		return OriginGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OriginGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OriginGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OriginGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, options *OriginGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OriginGroupsClient) getHandleResponse(resp *http.Response) (OriginGroupsClientGetResponse, error) {
	result := OriginGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginGroup); err != nil {
		return OriginGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointPager - Lists all of the existing origin groups within an endpoint.
//
// Generated from API version 2023-04-11-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - options - OriginGroupsClientListByEndpointOptions contains the optional parameters for the OriginGroupsClient.NewListByEndpointPager
//     method.
func (client *OriginGroupsClient) NewListByEndpointPager(resourceGroupName string, profileName string, endpointName string, options *OriginGroupsClientListByEndpointOptions) *runtime.Pager[OriginGroupsClientListByEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[OriginGroupsClientListByEndpointResponse]{
		More: func(page OriginGroupsClientListByEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OriginGroupsClientListByEndpointResponse) (OriginGroupsClientListByEndpointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OriginGroupsClientListByEndpointResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OriginGroupsClientListByEndpointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OriginGroupsClientListByEndpointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEndpointHandleResponse(resp)
		},
	})
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *OriginGroupsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *OriginGroupsClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *OriginGroupsClient) listByEndpointHandleResponse(resp *http.Response) (OriginGroupsClientListByEndpointResponse, error) {
	result := OriginGroupsClientListByEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OriginGroupListResult); err != nil {
		return OriginGroupsClientListByEndpointResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing origin group within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - originGroupName - Name of the origin group which is unique within the endpoint.
//   - originGroupUpdateProperties - Origin group properties
//   - options - OriginGroupsClientBeginUpdateOptions contains the optional parameters for the OriginGroupsClient.BeginUpdate
//     method.
func (client *OriginGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsClientBeginUpdateOptions) (*runtime.Poller[OriginGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroupUpdateProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OriginGroupsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OriginGroupsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing origin group within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-11-preview
func (client *OriginGroupsClient) update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, endpointName, originGroupName, originGroupUpdateProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OriginGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originGroupName string, originGroupUpdateProperties OriginGroupUpdateParameters, options *OriginGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if originGroupName == "" {
		return nil, errors.New("parameter originGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{originGroupName}", url.PathEscape(originGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, originGroupUpdateProperties)
}
