//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

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

// GroupsClient contains the methods for the ManagedNetworkGroups group.
// Don't use this type directly, use NewGroupsClient() instead.
type GroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGroupsClient creates a new instance of GroupsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
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

// BeginCreateOrUpdate - The Put ManagedNetworkGroups operation creates or updates a Managed Network Group resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkGroupName - The name of the Managed Network Group.
//   - managedNetworkGroup - Parameters supplied to the create/update a Managed Network Group resource
//   - options - GroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the GroupsClient.BeginCreateOrUpdate
//     method.
func (client *GroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, managedNetworkGroup, options)
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

// CreateOrUpdate - The Put ManagedNetworkGroups operation creates or updates a Managed Network Group resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *GroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, managedNetworkGroup, options)
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
func (client *GroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, managedNetworkGroup Group, options *GroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, managedNetworkGroup); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The Delete ManagedNetworkGroups operation deletes a Managed Network Group specified by the resource group,
// Managed Network name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkGroupName - The name of the Managed Network Group.
//   - options - GroupsClientBeginDeleteOptions contains the optional parameters for the GroupsClient.BeginDelete method.
func (client *GroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*runtime.Poller[GroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[GroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[GroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The Delete ManagedNetworkGroups operation deletes a Managed Network Group specified by the resource group, Managed
// Network name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *GroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "GroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
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
func (client *GroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The Get ManagedNetworkGroups operation gets a Managed Network Group specified by the resource group, Managed Network
// name, and group name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetworkGroupName - The name of the Managed Network Group.
//   - options - GroupsClientGetOptions contains the optional parameters for the GroupsClient.Get method.
func (client *GroupsClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientGetOptions) (GroupsClientGetResponse, error) {
	var err error
	const operationName = "GroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName, options)
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
func (client *GroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string, options *GroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	if managedNetworkGroupName == "" {
		return nil, errors.New("parameter managedNetworkGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkGroupName}", url.PathEscape(managedNetworkGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
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

// NewListByManagedNetworkPager - The ListByManagedNetwork ManagedNetworkGroup operation retrieves all the Managed Network
// Groups in a specified Managed Networks in a paginated format.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - options - GroupsClientListByManagedNetworkOptions contains the optional parameters for the GroupsClient.NewListByManagedNetworkPager
//     method.
func (client *GroupsClient) NewListByManagedNetworkPager(resourceGroupName string, managedNetworkName string, options *GroupsClientListByManagedNetworkOptions) *runtime.Pager[GroupsClientListByManagedNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[GroupsClientListByManagedNetworkResponse]{
		More: func(page GroupsClientListByManagedNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GroupsClientListByManagedNetworkResponse) (GroupsClientListByManagedNetworkResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GroupsClient.NewListByManagedNetworkPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedNetworkCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
			}, nil)
			if err != nil {
				return GroupsClientListByManagedNetworkResponse{}, err
			}
			return client.listByManagedNetworkHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedNetworkCreateRequest creates the ListByManagedNetwork request.
func (client *GroupsClient) listByManagedNetworkCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *GroupsClientListByManagedNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedNetworkHandleResponse handles the ListByManagedNetwork response.
func (client *GroupsClient) listByManagedNetworkHandleResponse(resp *http.Response) (GroupsClientListByManagedNetworkResponse, error) {
	result := GroupsClientListByManagedNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GroupListResult); err != nil {
		return GroupsClientListByManagedNetworkResponse{}, err
	}
	return result, nil
}
