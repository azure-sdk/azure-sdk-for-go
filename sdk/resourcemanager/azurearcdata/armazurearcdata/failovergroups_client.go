//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurearcdata

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

// FailoverGroupsClient contains the methods for the FailoverGroups group.
// Don't use this type directly, use NewFailoverGroupsClient() instead.
type FailoverGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFailoverGroupsClient creates a new instance of FailoverGroupsClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFailoverGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FailoverGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".FailoverGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FailoverGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces a failover group resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - failoverGroupName - The name of the Failover Group
//   - failoverGroupResource - desc
//   - options - FailoverGroupsClientBeginCreateOptions contains the optional parameters for the FailoverGroupsClient.BeginCreate
//     method.
func (client *FailoverGroupsClient) BeginCreate(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, failoverGroupResource FailoverGroupResource, options *FailoverGroupsClientBeginCreateOptions) (*runtime.Poller[FailoverGroupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sqlManagedInstanceName, failoverGroupName, failoverGroupResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FailoverGroupsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FailoverGroupsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates or replaces a failover group resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *FailoverGroupsClient) create(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, failoverGroupResource FailoverGroupResource, options *FailoverGroupsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, failoverGroupName, failoverGroupResource, options)
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

// createCreateRequest creates the Create request.
func (client *FailoverGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, failoverGroupResource FailoverGroupResource, options *FailoverGroupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}/failoverGroups/{failoverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, failoverGroupResource)
}

// BeginDelete - Deletes a failover group resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - failoverGroupName - The name of the Failover Group
//   - options - FailoverGroupsClientBeginDeleteOptions contains the optional parameters for the FailoverGroupsClient.BeginDelete
//     method.
func (client *FailoverGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*runtime.Poller[FailoverGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlManagedInstanceName, failoverGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[FailoverGroupsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[FailoverGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a failover group resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *FailoverGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, failoverGroupName, options)
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
func (client *FailoverGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *FailoverGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}/failoverGroups/{failoverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves a failover group resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - failoverGroupName - The name of the Failover Group
//   - options - FailoverGroupsClientGetOptions contains the optional parameters for the FailoverGroupsClient.Get method.
func (client *FailoverGroupsClient) Get(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *FailoverGroupsClientGetOptions) (FailoverGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, failoverGroupName, options)
	if err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FailoverGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FailoverGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *FailoverGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}/failoverGroups/{failoverGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	if failoverGroupName == "" {
		return nil, errors.New("parameter failoverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{failoverGroupName}", url.PathEscape(failoverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FailoverGroupsClient) getHandleResponse(resp *http.Response) (FailoverGroupsClientGetResponse, error) {
	result := FailoverGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FailoverGroupResource); err != nil {
		return FailoverGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the failover groups associated with the given sql managed instance.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - options - FailoverGroupsClientListOptions contains the optional parameters for the FailoverGroupsClient.NewListPager method.
func (client *FailoverGroupsClient) NewListPager(resourceGroupName string, sqlManagedInstanceName string, options *FailoverGroupsClientListOptions) *runtime.Pager[FailoverGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FailoverGroupsClientListResponse]{
		More: func(page FailoverGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FailoverGroupsClientListResponse) (FailoverGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FailoverGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FailoverGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FailoverGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FailoverGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *FailoverGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}/failoverGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FailoverGroupsClient) listHandleResponse(resp *http.Response) (FailoverGroupsClientListResponse, error) {
	result := FailoverGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FailoverGroupListResult); err != nil {
		return FailoverGroupsClientListResponse{}, err
	}
	return result, nil
}
