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

// TopicSpacesClient contains the methods for the TopicSpaces group.
// Don't use this type directly, use NewTopicSpacesClient() instead.
type TopicSpacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTopicSpacesClient creates a new instance of TopicSpacesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopicSpacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TopicSpacesClient, error) {
	cl, err := arm.NewClient(moduleName+".TopicSpacesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopicSpacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a topic space with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicSpaceName - The topic space name.
//   - topicSpaceInfo - Topic space information.
//   - options - TopicSpacesClientBeginCreateOrUpdateOptions contains the optional parameters for the TopicSpacesClient.BeginCreateOrUpdate
//     method.
func (client *TopicSpacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, topicSpaceInfo TopicSpace, options *TopicSpacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TopicSpacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, namespaceName, topicSpaceName, topicSpaceInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicSpacesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TopicSpacesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a topic space with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *TopicSpacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, topicSpaceInfo TopicSpace, options *TopicSpacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, topicSpaceName, topicSpaceInfo, options)
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
func (client *TopicSpacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, topicSpaceInfo TopicSpace, options *TopicSpacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topicSpaces/{topicSpaceName}"
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
	if topicSpaceName == "" {
		return nil, errors.New("parameter topicSpaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicSpaceName}", url.PathEscape(topicSpaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, topicSpaceInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an existing topic space.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicSpaceName - Name of the Topic space.
//   - options - TopicSpacesClientBeginDeleteOptions contains the optional parameters for the TopicSpacesClient.BeginDelete method.
func (client *TopicSpacesClient) BeginDelete(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, options *TopicSpacesClientBeginDeleteOptions) (*runtime.Poller[TopicSpacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, namespaceName, topicSpaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicSpacesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TopicSpacesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete an existing topic space.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *TopicSpacesClient) deleteOperation(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, options *TopicSpacesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, topicSpaceName, options)
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
func (client *TopicSpacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, options *TopicSpacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topicSpaces/{topicSpaceName}"
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
	if topicSpaceName == "" {
		return nil, errors.New("parameter topicSpaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicSpaceName}", url.PathEscape(topicSpaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of a topic space.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - topicSpaceName - Name of the Topic space.
//   - options - TopicSpacesClientGetOptions contains the optional parameters for the TopicSpacesClient.Get method.
func (client *TopicSpacesClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, options *TopicSpacesClientGetOptions) (TopicSpacesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, topicSpaceName, options)
	if err != nil {
		return TopicSpacesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicSpacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicSpacesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopicSpacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, topicSpaceName string, options *TopicSpacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topicSpaces/{topicSpaceName}"
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
	if topicSpaceName == "" {
		return nil, errors.New("parameter topicSpaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicSpaceName}", url.PathEscape(topicSpaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicSpacesClient) getHandleResponse(resp *http.Response) (TopicSpacesClientGetResponse, error) {
	result := TopicSpacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicSpace); err != nil {
		return TopicSpacesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Get all the topic spaces under a namespace.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - namespaceName - Name of the namespace.
//   - options - TopicSpacesClientListByNamespaceOptions contains the optional parameters for the TopicSpacesClient.NewListByNamespacePager
//     method.
func (client *TopicSpacesClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *TopicSpacesClientListByNamespaceOptions) *runtime.Pager[TopicSpacesClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicSpacesClientListByNamespaceResponse]{
		More: func(page TopicSpacesClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopicSpacesClientListByNamespaceResponse) (TopicSpacesClientListByNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TopicSpacesClientListByNamespaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TopicSpacesClientListByNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicSpacesClientListByNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNamespaceHandleResponse(resp)
		},
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *TopicSpacesClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *TopicSpacesClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/namespaces/{namespaceName}/topicSpaces"
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
	reqQP.Set("api-version", "2023-06-01-preview")
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
func (client *TopicSpacesClient) listByNamespaceHandleResponse(resp *http.Response) (TopicSpacesClientListByNamespaceResponse, error) {
	result := TopicSpacesClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicSpacesListResult); err != nil {
		return TopicSpacesClientListByNamespaceResponse{}, err
	}
	return result, nil
}
