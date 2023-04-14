//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
	parentType     PrivateEndpointConnectionsParentType
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\'.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, parentType PrivateEndpointConnectionsParentType, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateEndpointConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		parentType:     parentType,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete a specific private endpoint connection under a topic, domain, or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, parentName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateEndpointConnectionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a specific private endpoint connection under a topic, domain, or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, parentName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.parentType == "" {
		return nil, errors.New("parameter client.parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(client.parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific private endpoint connection under a topic, domain, or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
//     method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.parentType == "" {
		return nil, errors.New("parameter client.parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(client.parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourcePager - Get all private endpoint connections under a topic, domain, or partner namespace.
//
// Generated from API version 2022-06-15
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
//   - options - PrivateEndpointConnectionsClientListByResourceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByResourcePager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByResourcePager(resourceGroupName string, parentName string, options *PrivateEndpointConnectionsClientListByResourceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByResourceResponse]{
		More: func(page PrivateEndpointConnectionsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByResourceResponse) (PrivateEndpointConnectionsClientListByResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceCreateRequest(ctx, resourceGroupName, parentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateEndpointConnectionsClientListByResourceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateEndpointConnectionsClientListByResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionsClientListByResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceHandleResponse(resp)
		},
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *PrivateEndpointConnectionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, parentName string, options *PrivateEndpointConnectionsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.parentType == "" {
		return nil, errors.New("parameter client.parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(client.parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *PrivateEndpointConnectionsClient) listByResourceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByResourceResponse, error) {
	result := PrivateEndpointConnectionsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByResourceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a specific private endpoint connection under a topic, domain or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - privateEndpointConnection - The private endpoint connection object to update.
//   - options - PrivateEndpointConnectionsClientBeginUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginUpdate
//     method.
func (client *PrivateEndpointConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, parentName, privateEndpointConnectionName, privateEndpointConnection, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateEndpointConnectionsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a specific private endpoint connection under a topic, domain or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-15
func (client *PrivateEndpointConnectionsClient) update(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, parentName, privateEndpointConnectionName, privateEndpointConnection, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *PrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.parentType == "" {
		return nil, errors.New("parameter client.parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(client.parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, privateEndpointConnection)
}
