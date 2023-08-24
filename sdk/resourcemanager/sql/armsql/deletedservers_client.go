//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// DeletedServersClient contains the methods for the DeletedServers group.
// Don't use this type directly, use NewDeletedServersClient() instead.
type DeletedServersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeletedServersClient creates a new instance of DeletedServersClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeletedServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeletedServersClient, error) {
	cl, err := arm.NewClient(moduleName+".DeletedServersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeletedServersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - locationName - The name of the region where the resource is located.
//   - deletedServerName - The name of the deleted server.
//   - options - DeletedServersClientGetOptions contains the optional parameters for the DeletedServersClient.Get method.
func (client *DeletedServersClient) Get(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientGetOptions) (DeletedServersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, locationName, deletedServerName, options)
	if err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeletedServersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeletedServersClient) getCreateRequest(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if deletedServerName == "" {
		return nil, errors.New("parameter deletedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServerName}", url.PathEscape(deletedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeletedServersClient) getHandleResponse(resp *http.Response) (DeletedServersClientGetResponse, error) {
	result := DeletedServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServer); err != nil {
		return DeletedServersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all deleted servers in a subscription.
//
// Generated from API version 2023-02-01-preview
//   - options - DeletedServersClientListOptions contains the optional parameters for the DeletedServersClient.NewListPager method.
func (client *DeletedServersClient) NewListPager(options *DeletedServersClientListOptions) *runtime.Pager[DeletedServersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedServersClientListResponse]{
		More: func(page DeletedServersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedServersClientListResponse) (DeletedServersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeletedServersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeletedServersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeletedServersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DeletedServersClient) listCreateRequest(ctx context.Context, options *DeletedServersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/deletedServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedServersClient) listHandleResponse(resp *http.Response) (DeletedServersClientListResponse, error) {
	result := DeletedServersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServerListResult); err != nil {
		return DeletedServersClientListResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Gets a list of deleted servers for a location.
//
// Generated from API version 2023-02-01-preview
//   - locationName - The name of the region where the resource is located.
//   - options - DeletedServersClientListByLocationOptions contains the optional parameters for the DeletedServersClient.NewListByLocationPager
//     method.
func (client *DeletedServersClient) NewListByLocationPager(locationName string, options *DeletedServersClientListByLocationOptions) *runtime.Pager[DeletedServersClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedServersClientListByLocationResponse]{
		More: func(page DeletedServersClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedServersClientListByLocationResponse) (DeletedServersClientListByLocationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLocationCreateRequest(ctx, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeletedServersClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeletedServersClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeletedServersClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DeletedServersClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *DeletedServersClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *DeletedServersClient) listByLocationHandleResponse(resp *http.Response) (DeletedServersClientListByLocationResponse, error) {
	result := DeletedServersClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServerListResult); err != nil {
		return DeletedServersClientListByLocationResponse{}, err
	}
	return result, nil
}

// BeginRecover - Recovers a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - locationName - The name of the region where the resource is located.
//   - deletedServerName - The name of the deleted server.
//   - options - DeletedServersClientBeginRecoverOptions contains the optional parameters for the DeletedServersClient.BeginRecover
//     method.
func (client *DeletedServersClient) BeginRecover(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (*runtime.Poller[DeletedServersClientRecoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.recoverOperation(ctx, locationName, deletedServerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DeletedServersClientRecoverResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DeletedServersClientRecoverResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Recover - Recovers a deleted server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
func (client *DeletedServersClient) recoverOperation(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (*http.Response, error) {
	var err error
	req, err := client.recoverCreateRequest(ctx, locationName, deletedServerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// recoverCreateRequest creates the Recover request.
func (client *DeletedServersClient) recoverCreateRequest(ctx context.Context, locationName string, deletedServerName string, options *DeletedServersClientBeginRecoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}/recover"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if deletedServerName == "" {
		return nil, errors.New("parameter deletedServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServerName}", url.PathEscape(deletedServerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
