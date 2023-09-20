//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageactions

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

// StorageTaskAssignmentClient contains the methods for the StorageTaskAssignment group.
// Don't use this type directly, use NewStorageTaskAssignmentClient() instead.
type StorageTaskAssignmentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageTaskAssignmentClient creates a new instance of StorageTaskAssignmentClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageTaskAssignmentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageTaskAssignmentClient, error) {
	cl, err := arm.NewClient(moduleName+".StorageTaskAssignmentClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageTaskAssignmentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all the storage tasks available under the given resource group.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageTaskName - The name of the storage task within the specified resource group. Storage task names must be between
//     3 and 18 characters in length and use numbers and lower-case letters only.
//   - options - StorageTaskAssignmentClientListOptions contains the optional parameters for the StorageTaskAssignmentClient.NewListPager
//     method.
func (client *StorageTaskAssignmentClient) NewListPager(resourceGroupName string, storageTaskName string, options *StorageTaskAssignmentClientListOptions) *runtime.Pager[StorageTaskAssignmentClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageTaskAssignmentClientListResponse]{
		More: func(page StorageTaskAssignmentClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageTaskAssignmentClientListResponse) (StorageTaskAssignmentClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, storageTaskName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageTaskAssignmentClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StorageTaskAssignmentClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageTaskAssignmentClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *StorageTaskAssignmentClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageTaskName string, options *StorageTaskAssignmentClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageActions/storageTasks/{storageTaskName}/storageTaskAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageTaskName == "" {
		return nil, errors.New("parameter storageTaskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTaskName}", url.PathEscape(storageTaskName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", *options.Maxpagesize)
	}
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StorageTaskAssignmentClient) listHandleResponse(resp *http.Response) (StorageTaskAssignmentClientListResponse, error) {
	result := StorageTaskAssignmentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageTaskAssignmentsListResult); err != nil {
		return StorageTaskAssignmentClientListResponse{}, err
	}
	return result, nil
}
