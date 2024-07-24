//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

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

// ConnectedClusterClient contains the methods for the ConnectedCluster group.
// Don't use this type directly, use NewConnectedClusterClient() instead.
type ConnectedClusterClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedClusterClient creates a new instance of ConnectedClusterClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedClusterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedClusterClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedClusterClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - API to register a new Kubernetes cluster and create a tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kubernetes cluster on which get is called.
//   - connectedCluster - Parameters supplied to Create a Connected Cluster.
//   - options - ConnectedClusterClientBeginCreateOptions contains the optional parameters for the ConnectedClusterClient.BeginCreate
//     method.
func (client *ConnectedClusterClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*runtime.Poller[ConnectedClusterClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterName, connectedCluster, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedClusterClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConnectedClusterClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - API to register a new Kubernetes cluster and create a tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
func (client *ConnectedClusterClient) create(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedClusterClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, connectedCluster, options)
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

// createCreateRequest creates the Create request.
func (client *ConnectedClusterClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectedCluster); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a connected cluster, removing the tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kubernetes cluster on which get is called.
//   - options - ConnectedClusterClientBeginDeleteOptions contains the optional parameters for the ConnectedClusterClient.BeginDelete
//     method.
func (client *ConnectedClusterClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*runtime.Poller[ConnectedClusterClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectedClusterClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConnectedClusterClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a connected cluster, removing the tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
func (client *ConnectedClusterClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectedClusterClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
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
func (client *ConnectedClusterClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the properties of the specified connected cluster, including name, identity, properties, and additional cluster
// details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kubernetes cluster on which get is called.
//   - options - ConnectedClusterClientGetOptions contains the optional parameters for the ConnectedClusterClient.Get method.
func (client *ConnectedClusterClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientGetOptions) (ConnectedClusterClientGetResponse, error) {
	var err error
	const operationName = "ConnectedClusterClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedClusterClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectedClusterClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedClusterClient) getHandleResponse(resp *http.Response) (ConnectedClusterClientGetResponse, error) {
	result := ConnectedClusterClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedCluster); err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - API to enumerate registered connected K8s clusters under a Resource Group
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ConnectedClusterClientListByResourceGroupOptions contains the optional parameters for the ConnectedClusterClient.NewListByResourceGroupPager
//     method.
func (client *ConnectedClusterClient) NewListByResourceGroupPager(resourceGroupName string, options *ConnectedClusterClientListByResourceGroupOptions) *runtime.Pager[ConnectedClusterClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedClusterClientListByResourceGroupResponse]{
		More: func(page ConnectedClusterClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedClusterClientListByResourceGroupResponse) (ConnectedClusterClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedClusterClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ConnectedClusterClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConnectedClusterClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConnectedClusterClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters"
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
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConnectedClusterClient) listByResourceGroupHandleResponse(resp *http.Response) (ConnectedClusterClientListByResourceGroupResponse, error) {
	result := ConnectedClusterClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedClusterList); err != nil {
		return ConnectedClusterClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - API to enumerate registered connected K8s clusters under a Subscription
//
// Generated from API version 2024-07-15-preview
//   - options - ConnectedClusterClientListBySubscriptionOptions contains the optional parameters for the ConnectedClusterClient.NewListBySubscriptionPager
//     method.
func (client *ConnectedClusterClient) NewListBySubscriptionPager(options *ConnectedClusterClientListBySubscriptionOptions) *runtime.Pager[ConnectedClusterClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedClusterClientListBySubscriptionResponse]{
		More: func(page ConnectedClusterClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedClusterClientListBySubscriptionResponse) (ConnectedClusterClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedClusterClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ConnectedClusterClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConnectedClusterClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConnectedClusterClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Kubernetes/connectedClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConnectedClusterClient) listBySubscriptionHandleResponse(resp *http.Response) (ConnectedClusterClientListBySubscriptionResponse, error) {
	result := ConnectedClusterClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedClusterList); err != nil {
		return ConnectedClusterClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListClusterUserCredential - Gets cluster user credentials of the connected cluster with a specified resource group and
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kubernetes cluster on which get is called.
//   - properties - ListClusterUserCredential properties
//   - options - ConnectedClusterClientListClusterUserCredentialOptions contains the optional parameters for the ConnectedClusterClient.ListClusterUserCredential
//     method.
func (client *ConnectedClusterClient) ListClusterUserCredential(ctx context.Context, resourceGroupName string, clusterName string, properties ListClusterUserCredentialProperties, options *ConnectedClusterClientListClusterUserCredentialOptions) (ConnectedClusterClientListClusterUserCredentialResponse, error) {
	var err error
	const operationName = "ConnectedClusterClient.ListClusterUserCredential"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listClusterUserCredentialCreateRequest(ctx, resourceGroupName, clusterName, properties, options)
	if err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	resp, err := client.listClusterUserCredentialHandleResponse(httpResp)
	return resp, err
}

// listClusterUserCredentialCreateRequest creates the ListClusterUserCredential request.
func (client *ConnectedClusterClient) listClusterUserCredentialCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, properties ListClusterUserCredentialProperties, options *ConnectedClusterClientListClusterUserCredentialOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}/listClusterUserCredential"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// listClusterUserCredentialHandleResponse handles the ListClusterUserCredential response.
func (client *ConnectedClusterClient) listClusterUserCredentialHandleResponse(resp *http.Response) (ConnectedClusterClientListClusterUserCredentialResponse, error) {
	result := ConnectedClusterClientListClusterUserCredentialResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResults); err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the connected cluster resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kubernetes cluster on which get is called.
//   - connectedClusterPatch - Parameters supplied to update Connected Cluster.
//   - options - ConnectedClusterClientUpdateOptions contains the optional parameters for the ConnectedClusterClient.Update method.
func (client *ConnectedClusterClient) Update(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch ConnectedClusterPatch, options *ConnectedClusterClientUpdateOptions) (ConnectedClusterClientUpdateResponse, error) {
	var err error
	const operationName = "ConnectedClusterClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, connectedClusterPatch, options)
	if err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedClusterClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ConnectedClusterClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch ConnectedClusterPatch, options *ConnectedClusterClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectedClusterPatch); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConnectedClusterClient) updateHandleResponse(resp *http.Response) (ConnectedClusterClientUpdateResponse, error) {
	result := ConnectedClusterClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedCluster); err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	return result, nil
}
