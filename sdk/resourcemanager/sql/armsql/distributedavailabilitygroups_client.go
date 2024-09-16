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

// DistributedAvailabilityGroupsClient contains the methods for the DistributedAvailabilityGroups group.
// Don't use this type directly, use NewDistributedAvailabilityGroupsClient() instead.
type DistributedAvailabilityGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDistributedAvailabilityGroupsClient creates a new instance of DistributedAvailabilityGroupsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDistributedAvailabilityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DistributedAvailabilityGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DistributedAvailabilityGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a distributed availability group between Sql On-Prem and Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - distributedAvailabilityGroupName - The distributed availability group name.
//   - parameters - The distributed availability group info.
//   - options - DistributedAvailabilityGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the DistributedAvailabilityGroupsClient.BeginCreateOrUpdate
//     method.
func (client *DistributedAvailabilityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DistributedAvailabilityGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DistributedAvailabilityGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DistributedAvailabilityGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a distributed availability group between Sql On-Prem and Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
func (client *DistributedAvailabilityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DistributedAvailabilityGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DistributedAvailabilityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/distributedAvailabilityGroups/{distributedAvailabilityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if distributedAvailabilityGroupName == "" {
		return nil, errors.New("parameter distributedAvailabilityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{distributedAvailabilityGroupName}", url.PathEscape(distributedAvailabilityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Drops a distributed availability group between Sql On-Prem and Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - distributedAvailabilityGroupName - The distributed availability group name.
//   - options - DistributedAvailabilityGroupsClientBeginDeleteOptions contains the optional parameters for the DistributedAvailabilityGroupsClient.BeginDelete
//     method.
func (client *DistributedAvailabilityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, options *DistributedAvailabilityGroupsClientBeginDeleteOptions) (*runtime.Poller[DistributedAvailabilityGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DistributedAvailabilityGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DistributedAvailabilityGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Drops a distributed availability group between Sql On-Prem and Sql Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
func (client *DistributedAvailabilityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, options *DistributedAvailabilityGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DistributedAvailabilityGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, options)
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
func (client *DistributedAvailabilityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, options *DistributedAvailabilityGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/distributedAvailabilityGroups/{distributedAvailabilityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if distributedAvailabilityGroupName == "" {
		return nil, errors.New("parameter distributedAvailabilityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{distributedAvailabilityGroupName}", url.PathEscape(distributedAvailabilityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a distributed availability group info.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - distributedAvailabilityGroupName - The distributed availability group name.
//   - options - DistributedAvailabilityGroupsClientGetOptions contains the optional parameters for the DistributedAvailabilityGroupsClient.Get
//     method.
func (client *DistributedAvailabilityGroupsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, options *DistributedAvailabilityGroupsClientGetOptions) (DistributedAvailabilityGroupsClientGetResponse, error) {
	var err error
	const operationName = "DistributedAvailabilityGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, options)
	if err != nil {
		return DistributedAvailabilityGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DistributedAvailabilityGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DistributedAvailabilityGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DistributedAvailabilityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, options *DistributedAvailabilityGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/distributedAvailabilityGroups/{distributedAvailabilityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if distributedAvailabilityGroupName == "" {
		return nil, errors.New("parameter distributedAvailabilityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{distributedAvailabilityGroupName}", url.PathEscape(distributedAvailabilityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DistributedAvailabilityGroupsClient) getHandleResponse(resp *http.Response) (DistributedAvailabilityGroupsClientGetResponse, error) {
	result := DistributedAvailabilityGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DistributedAvailabilityGroup); err != nil {
		return DistributedAvailabilityGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - Gets a list of a distributed availability groups in instance.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - DistributedAvailabilityGroupsClientListByInstanceOptions contains the optional parameters for the DistributedAvailabilityGroupsClient.NewListByInstancePager
//     method.
func (client *DistributedAvailabilityGroupsClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *DistributedAvailabilityGroupsClientListByInstanceOptions) *runtime.Pager[DistributedAvailabilityGroupsClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DistributedAvailabilityGroupsClientListByInstanceResponse]{
		More: func(page DistributedAvailabilityGroupsClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DistributedAvailabilityGroupsClientListByInstanceResponse) (DistributedAvailabilityGroupsClientListByInstanceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DistributedAvailabilityGroupsClient.NewListByInstancePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			}, nil)
			if err != nil {
				return DistributedAvailabilityGroupsClientListByInstanceResponse{}, err
			}
			return client.listByInstanceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *DistributedAvailabilityGroupsClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *DistributedAvailabilityGroupsClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/distributedAvailabilityGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *DistributedAvailabilityGroupsClient) listByInstanceHandleResponse(resp *http.Response) (DistributedAvailabilityGroupsClientListByInstanceResponse, error) {
	result := DistributedAvailabilityGroupsClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DistributedAvailabilityGroupsListResult); err != nil {
		return DistributedAvailabilityGroupsClientListByInstanceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a distributed availability group replication mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - distributedAvailabilityGroupName - The distributed availability group name.
//   - parameters - The distributed availability group info.
//   - options - DistributedAvailabilityGroupsClientBeginUpdateOptions contains the optional parameters for the DistributedAvailabilityGroupsClient.BeginUpdate
//     method.
func (client *DistributedAvailabilityGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginUpdateOptions) (*runtime.Poller[DistributedAvailabilityGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DistributedAvailabilityGroupsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DistributedAvailabilityGroupsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a distributed availability group replication mode.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
func (client *DistributedAvailabilityGroupsClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DistributedAvailabilityGroupsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, distributedAvailabilityGroupName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *DistributedAvailabilityGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, distributedAvailabilityGroupName string, parameters DistributedAvailabilityGroup, options *DistributedAvailabilityGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/distributedAvailabilityGroups/{distributedAvailabilityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if distributedAvailabilityGroupName == "" {
		return nil, errors.New("parameter distributedAvailabilityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{distributedAvailabilityGroupName}", url.PathEscape(distributedAvailabilityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
