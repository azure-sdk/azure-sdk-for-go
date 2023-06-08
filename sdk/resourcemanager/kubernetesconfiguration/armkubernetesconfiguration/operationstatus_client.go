//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkubernetesconfiguration

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

// OperationStatusClient contains the methods for the OperationStatus group.
// Don't use this type directly, use NewOperationStatusClient() instead.
type OperationStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationStatusClient creates a new instance of OperationStatusClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationStatusClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get Async Operation status
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
//   - clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
//   - clusterName - The name of the kubernetes cluster.
//   - extensionName - Name of the Extension.
//   - operationID - operation Id
//   - options - OperationStatusClientGetOptions contains the optional parameters for the OperationStatusClient.Get method.
func (client *OperationStatusClient) Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, operationID string, options *OperationStatusClientGetOptions) (OperationStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName, operationID, options)
	if err != nil {
		return OperationStatusClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, extensionName string, operationID string, options *OperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensions/{extensionName}/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionName == "" {
		return nil, errors.New("parameter extensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionName}", url.PathEscape(extensionName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationStatusClient) getHandleResponse(resp *http.Response) (OperationStatusClientGetResponse, error) {
	result := OperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatusResult); err != nil {
		return OperationStatusClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Async Operations, currently in progress, in a cluster
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterRp - The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
//   - clusterResourceName - The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
//   - clusterName - The name of the kubernetes cluster.
//   - options - OperationStatusClientListOptions contains the optional parameters for the OperationStatusClient.NewListPager
//     method.
func (client *OperationStatusClient) NewListPager(resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *OperationStatusClientListOptions) *runtime.Pager[OperationStatusClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationStatusClientListResponse]{
		More: func(page OperationStatusClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationStatusClientListResponse) (OperationStatusClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationStatusClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OperationStatusClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationStatusClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, options *OperationStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/operations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(clusterRp))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(clusterResourceName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationStatusClient) listHandleResponse(resp *http.Response) (OperationStatusClientListResponse, error) {
	result := OperationStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatusList); err != nil {
		return OperationStatusClientListResponse{}, err
	}
	return result, nil
}
