//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

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

// GroupsClient contains the methods for the SQLVirtualMachineGroups group.
// Don't use this type directly, use NewGroupsClient() instead.
type GroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGroupsClient creates a new instance of GroupsClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - parameters - The SQL virtual machine group.
//   - options - GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate
//     method.
func (client *GroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *GroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
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
func (client *GroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters Group, options *GroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - options - GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
func (client *GroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (*runtime.Poller[GroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *GroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
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
func (client *GroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a SQL virtual machine group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - options - GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
func (client *GroupsClient) Get(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientGetOptions) (GroupsClientGetResponse, error) {
	var err error
	const operationName = "GroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *GroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GroupsClient) getHandleResponse(resp *http.Response) (GroupsClientGetResponse, error) {
	result := GroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Group); err != nil {
		return GroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all SQL virtual machine groups in a subscription.
//
// Generated from API version 2023-10-01
//   - options - GroupsClientListOptions contains the optional parameters for the GroupsClient.NewListPager method.
func (client *GroupsClient) NewListPager(options *GroupsClientListOptions) *runtime.Pager[GroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupsClientListResponse]{
		More: func(page GroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupsClientListResponse) (GroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return GroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GroupsClient) listCreateRequest(ctx context.Context, options *GroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GroupsClient) listHandleResponse(resp *http.Response) (GroupsClientListResponse, error) {
	result := GroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all SQL virtual machine groups in a resource group.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - options - GroupsClientListByResourceGroupOptions contains the optional parameters for the GroupsClient.NewListByResourceGroupPager
//     method.
func (client *GroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *GroupsClientListByResourceGroupOptions) *runtime.Pager[GroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupsClientListByResourceGroupResponse]{
		More: func(page GroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupsClientListByResourceGroupResponse) (GroupsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GroupsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return GroupsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (GroupsClientListByResourceGroupResponse, error) {
	result := GroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates SQL virtual machine group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - parameters - The SQL virtual machine group.
//   - options - GroupsClientBeginUpdateOptions contains the optional parameters for the GroupsClient.BeginUpdate method.
func (client *GroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (*runtime.Poller[GroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates SQL virtual machine group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *GroupsClient) update(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *GroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, parameters GroupUpdate, options *GroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
