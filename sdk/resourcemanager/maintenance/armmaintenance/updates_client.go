// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// UpdatesClient contains the methods for the Updates group.
// Don't use this type directly, use NewUpdatesClient() instead.
type UpdatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUpdatesClient creates a new instance of UpdatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUpdatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UpdatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UpdatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get updates to resources.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - UpdatesClientListOptions contains the optional parameters for the UpdatesClient.NewListPager method.
func (client *UpdatesClient) NewListPager(resourceGroupName string, providerName string, resourceType string, resourceName string, options *UpdatesClientListOptions) *runtime.Pager[UpdatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UpdatesClientListResponse]{
		More: func(page UpdatesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UpdatesClientListResponse) (UpdatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UpdatesClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, options)
			if err != nil {
				return UpdatesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UpdatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UpdatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *UpdatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, _ *UpdatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UpdatesClient) listHandleResponse(resp *http.Response) (UpdatesClientListResponse, error) {
	result := UpdatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUpdatesResult); err != nil {
		return UpdatesClientListResponse{}, err
	}
	return result, nil
}

// NewListParentPager - Get updates to resources.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - UpdatesClientListParentOptions contains the optional parameters for the UpdatesClient.NewListParentPager method.
func (client *UpdatesClient) NewListParentPager(resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *UpdatesClientListParentOptions) *runtime.Pager[UpdatesClientListParentResponse] {
	return runtime.NewPager(runtime.PagingHandler[UpdatesClientListParentResponse]{
		More: func(page UpdatesClientListParentResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *UpdatesClientListParentResponse) (UpdatesClientListParentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UpdatesClient.NewListParentPager")
			req, err := client.listParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, options)
			if err != nil {
				return UpdatesClientListParentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UpdatesClientListParentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UpdatesClientListParentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listParentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listParentCreateRequest creates the ListParent request.
func (client *UpdatesClient) listParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, _ *UpdatesClientListParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listParentHandleResponse handles the ListParent response.
func (client *UpdatesClient) listParentHandleResponse(resp *http.Response) (UpdatesClientListParentResponse, error) {
	result := UpdatesClientListParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUpdatesResult); err != nil {
		return UpdatesClientListParentResponse{}, err
	}
	return result, nil
}
