//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ClientsClient contains the methods for the Clients group.
// Don't use this type directly, use NewClientsClient() instead.
type ClientsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClientsClient creates a new instance of ClientsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientsClient, error) {
	cl, err := arm.NewClient(moduleName+".ClientsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClientsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a client with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientName - The client name.
//   - clientInfo - Client information.
//   - options - ClientsClientBeginCreateOrUpdateOptions contains the optional parameters for the ClientsClient.BeginCreateOrUpdate
//     method.
func (client *ClientsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, clientInfo Client, options *ClientsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClientsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, namespaceName, clientName, clientInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClientsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a client with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *ClientsClient) createOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, clientInfo Client, options *ClientsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, clientName, clientInfo, options)
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
func (client *ClientsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, clientInfo Client, options *ClientsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clients/{clientName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientName == "" {
		return nil, errors.New("parameter clientName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientName}", url.PathEscape(clientName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clientInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an existing client.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientName - Name of the client.
//   - options - ClientsClientBeginDeleteOptions contains the optional parameters for the ClientsClient.BeginDelete method.
func (client *ClientsClient) BeginDelete(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, options *ClientsClientBeginDeleteOptions) (*runtime.Poller[ClientsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, namespaceName, clientName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClientsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete an existing client.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *ClientsClient) deleteOperation(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, options *ClientsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, clientName, options)
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
func (client *ClientsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, options *ClientsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clients/{clientName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientName == "" {
		return nil, errors.New("parameter clientName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientName}", url.PathEscape(clientName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of a client.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - clientName - Name of the client.
//   - options - ClientsClientGetOptions contains the optional parameters for the ClientsClient.Get method.
func (client *ClientsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, options *ClientsClientGetOptions) (ClientsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, clientName, options)
	if err != nil {
		return ClientsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClientsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, clientName string, options *ClientsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clients/{clientName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if clientName == "" {
		return nil, errors.New("parameter clientName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clientName}", url.PathEscape(clientName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClientsClient) getHandleResponse(resp *http.Response) (ClientsClientGetResponse, error) {
	result := ClientsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Client); err != nil {
		return ClientsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Get all the permission bindings under a namespace.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - options - ClientsClientListByNamespaceOptions contains the optional parameters for the ClientsClient.NewListByNamespacePager
//     method.
func (client *ClientsClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *ClientsClientListByNamespaceOptions) *runtime.Pager[ClientsClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientsClientListByNamespaceResponse]{
		More: func(page ClientsClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientsClientListByNamespaceResponse) (ClientsClientListByNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientsClientListByNamespaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientsClientListByNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientsClientListByNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNamespaceHandleResponse(resp)
		},
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *ClientsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *ClientsClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/clients"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
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

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *ClientsClient) listByNamespaceHandleResponse(resp *http.Response) (ClientsClientListByNamespaceResponse, error) {
	result := ClientsClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientsListResult); err != nil {
		return ClientsClientListByNamespaceResponse{}, err
	}
	return result, nil
}
