// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

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

// ExadbVMClustersClient contains the methods for the ExadbVMClusters group.
// Don't use this type directly, use NewExadbVMClustersClient() instead.
type ExadbVMClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExadbVMClustersClient creates a new instance of ExadbVMClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExadbVMClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExadbVMClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExadbVMClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - exadbVMClusterName - The name of the ExadbVmCluster
//   - resource - Resource create parameters.
//   - options - ExadbVMClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ExadbVMClustersClient.BeginCreateOrUpdate
//     method.
func (client *ExadbVMClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, exadbVMClusterName string, resource ExadbVMCluster, options *ExadbVMClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExadbVMClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, exadbVMClusterName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExadbVMClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExadbVMClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ExadbVMClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, exadbVMClusterName string, resource ExadbVMCluster, options *ExadbVMClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExadbVMClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, exadbVMClusterName, resource, options)
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
func (client *ExadbVMClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, exadbVMClusterName string, resource ExadbVMCluster, _ *ExadbVMClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters/{exadbVmClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if exadbVMClusterName == "" {
		return nil, errors.New("parameter exadbVMClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exadbVmClusterName}", url.PathEscape(exadbVMClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - exadbVMClusterName - The name of the ExadbVmCluster
//   - options - ExadbVMClustersClientBeginDeleteOptions contains the optional parameters for the ExadbVMClustersClient.BeginDelete
//     method.
func (client *ExadbVMClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, exadbVMClusterName string, options *ExadbVMClustersClientBeginDeleteOptions) (*runtime.Poller[ExadbVMClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, exadbVMClusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExadbVMClustersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExadbVMClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ExadbVMClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, exadbVMClusterName string, options *ExadbVMClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ExadbVMClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, exadbVMClusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExadbVMClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, exadbVMClusterName string, _ *ExadbVMClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters/{exadbVmClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if exadbVMClusterName == "" {
		return nil, errors.New("parameter exadbVMClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exadbVmClusterName}", url.PathEscape(exadbVMClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - exadbVMClusterName - The name of the ExadbVmCluster
//   - options - ExadbVMClustersClientGetOptions contains the optional parameters for the ExadbVMClustersClient.Get method.
func (client *ExadbVMClustersClient) Get(ctx context.Context, resourceGroupName string, exadbVMClusterName string, options *ExadbVMClustersClientGetOptions) (ExadbVMClustersClientGetResponse, error) {
	var err error
	const operationName = "ExadbVMClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, exadbVMClusterName, options)
	if err != nil {
		return ExadbVMClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExadbVMClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExadbVMClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExadbVMClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, exadbVMClusterName string, _ *ExadbVMClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters/{exadbVmClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if exadbVMClusterName == "" {
		return nil, errors.New("parameter exadbVMClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exadbVmClusterName}", url.PathEscape(exadbVMClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExadbVMClustersClient) getHandleResponse(resp *http.Response) (ExadbVMClustersClientGetResponse, error) {
	result := ExadbVMClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExadbVMCluster); err != nil {
		return ExadbVMClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List ExadbVmCluster resources by resource group
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ExadbVMClustersClientListByResourceGroupOptions contains the optional parameters for the ExadbVMClustersClient.NewListByResourceGroupPager
//     method.
func (client *ExadbVMClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *ExadbVMClustersClientListByResourceGroupOptions) *runtime.Pager[ExadbVMClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExadbVMClustersClientListByResourceGroupResponse]{
		More: func(page ExadbVMClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExadbVMClustersClientListByResourceGroupResponse) (ExadbVMClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExadbVMClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ExadbVMClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExadbVMClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *ExadbVMClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExadbVMClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (ExadbVMClustersClientListByResourceGroupResponse, error) {
	result := ExadbVMClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExadbVMClusterListResult); err != nil {
		return ExadbVMClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List ExadbVmCluster resources by subscription ID
//
// Generated from API version 2025-03-01
//   - options - ExadbVMClustersClientListBySubscriptionOptions contains the optional parameters for the ExadbVMClustersClient.NewListBySubscriptionPager
//     method.
func (client *ExadbVMClustersClient) NewListBySubscriptionPager(options *ExadbVMClustersClientListBySubscriptionOptions) *runtime.Pager[ExadbVMClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExadbVMClustersClientListBySubscriptionResponse]{
		More: func(page ExadbVMClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExadbVMClustersClientListBySubscriptionResponse) (ExadbVMClustersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExadbVMClustersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExadbVMClustersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ExadbVMClustersClient) listBySubscriptionCreateRequest(ctx context.Context, _ *ExadbVMClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/exadbVmClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ExadbVMClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (ExadbVMClustersClientListBySubscriptionResponse, error) {
	result := ExadbVMClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExadbVMClusterListResult); err != nil {
		return ExadbVMClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRemoveVMs - Remove VMs from the VM Cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - exadbVMClusterName - The name of the ExadbVmCluster
//   - body - The content of the action request
//   - options - ExadbVMClustersClientBeginRemoveVMsOptions contains the optional parameters for the ExadbVMClustersClient.BeginRemoveVMs
//     method.
func (client *ExadbVMClustersClient) BeginRemoveVMs(ctx context.Context, resourceGroupName string, exadbVMClusterName string, body RemoveVirtualMachineFromExadbVMClusterDetails, options *ExadbVMClustersClientBeginRemoveVMsOptions) (*runtime.Poller[ExadbVMClustersClientRemoveVMsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.removeVMs(ctx, resourceGroupName, exadbVMClusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExadbVMClustersClientRemoveVMsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExadbVMClustersClientRemoveVMsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RemoveVMs - Remove VMs from the VM Cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ExadbVMClustersClient) removeVMs(ctx context.Context, resourceGroupName string, exadbVMClusterName string, body RemoveVirtualMachineFromExadbVMClusterDetails, options *ExadbVMClustersClientBeginRemoveVMsOptions) (*http.Response, error) {
	var err error
	const operationName = "ExadbVMClustersClient.BeginRemoveVMs"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.removeVMsCreateRequest(ctx, resourceGroupName, exadbVMClusterName, body, options)
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

// removeVMsCreateRequest creates the RemoveVMs request.
func (client *ExadbVMClustersClient) removeVMsCreateRequest(ctx context.Context, resourceGroupName string, exadbVMClusterName string, body RemoveVirtualMachineFromExadbVMClusterDetails, _ *ExadbVMClustersClientBeginRemoveVMsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters/{exadbVmClusterName}/removeVms"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if exadbVMClusterName == "" {
		return nil, errors.New("parameter exadbVMClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exadbVmClusterName}", url.PathEscape(exadbVMClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Update a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - exadbVMClusterName - The name of the ExadbVmCluster
//   - properties - The resource properties to be updated.
//   - options - ExadbVMClustersClientBeginUpdateOptions contains the optional parameters for the ExadbVMClustersClient.BeginUpdate
//     method.
func (client *ExadbVMClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, exadbVMClusterName string, properties ExadbVMClusterUpdate, options *ExadbVMClustersClientBeginUpdateOptions) (*runtime.Poller[ExadbVMClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, exadbVMClusterName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExadbVMClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ExadbVMClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a ExadbVmCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
func (client *ExadbVMClustersClient) update(ctx context.Context, resourceGroupName string, exadbVMClusterName string, properties ExadbVMClusterUpdate, options *ExadbVMClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExadbVMClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, exadbVMClusterName, properties, options)
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
func (client *ExadbVMClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, exadbVMClusterName string, properties ExadbVMClusterUpdate, _ *ExadbVMClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Oracle.Database/exadbVmClusters/{exadbVmClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if exadbVMClusterName == "" {
		return nil, errors.New("parameter exadbVMClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exadbVmClusterName}", url.PathEscape(exadbVMClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
