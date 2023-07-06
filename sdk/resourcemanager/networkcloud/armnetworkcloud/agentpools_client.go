//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud

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

// AgentPoolsClient contains the methods for the AgentPools group.
// Don't use this type directly, use NewAgentPoolsClient() instead.
type AgentPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAgentPoolsClient creates a new instance of AgentPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentPoolsClient, error) {
	cl, err := arm.NewClient(moduleName+".AgentPoolsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AgentPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new Kubernetes cluster agent pool or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - agentPoolName - The name of the Kubernetes cluster agent pool.
//   - agentPoolParameters - The request body.
//   - options - AgentPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginCreateOrUpdate
//     method.
func (client *AgentPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolParameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AgentPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, agentPoolParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new Kubernetes cluster agent pool or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *AgentPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolParameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, agentPoolParameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AgentPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolParameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/agentPools/{agentPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, agentPoolParameters)
}

// BeginDelete - Delete the provided Kubernetes cluster agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - agentPoolName - The name of the Kubernetes cluster agent pool.
//   - options - AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
func (client *AgentPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*runtime.Poller[AgentPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided Kubernetes cluster agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *AgentPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, options)
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
func (client *AgentPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/agentPools/{agentPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided Kubernetes cluster agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - agentPoolName - The name of the Kubernetes cluster agent pool.
//   - options - AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
func (client *AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *AgentPoolsClientGetOptions) (AgentPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, options *AgentPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/agentPools/{agentPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentPoolsClient) getHandleResponse(resp *http.Response) (AgentPoolsClientGetResponse, error) {
	result := AgentPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPool); err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByKubernetesClusterPager - Get a list of agent pools for the provided Kubernetes cluster.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - options - AgentPoolsClientListByKubernetesClusterOptions contains the optional parameters for the AgentPoolsClient.NewListByKubernetesClusterPager
//     method.
func (client *AgentPoolsClient) NewListByKubernetesClusterPager(resourceGroupName string, kubernetesClusterName string, options *AgentPoolsClientListByKubernetesClusterOptions) *runtime.Pager[AgentPoolsClientListByKubernetesClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgentPoolsClientListByKubernetesClusterResponse]{
		More: func(page AgentPoolsClientListByKubernetesClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgentPoolsClientListByKubernetesClusterResponse) (AgentPoolsClientListByKubernetesClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByKubernetesClusterCreateRequest(ctx, resourceGroupName, kubernetesClusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AgentPoolsClientListByKubernetesClusterResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AgentPoolsClientListByKubernetesClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AgentPoolsClientListByKubernetesClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByKubernetesClusterHandleResponse(resp)
		},
	})
}

// listByKubernetesClusterCreateRequest creates the ListByKubernetesCluster request.
func (client *AgentPoolsClient) listByKubernetesClusterCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, options *AgentPoolsClientListByKubernetesClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/agentPools"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByKubernetesClusterHandleResponse handles the ListByKubernetesCluster response.
func (client *AgentPoolsClient) listByKubernetesClusterHandleResponse(resp *http.Response) (AgentPoolsClientListByKubernetesClusterResponse, error) {
	result := AgentPoolsClientListByKubernetesClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolList); err != nil {
		return AgentPoolsClientListByKubernetesClusterResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch the properties of the provided Kubernetes cluster agent pool, or update the tags associated with the
// Kubernetes cluster agent pool. Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - kubernetesClusterName - The name of the Kubernetes cluster.
//   - agentPoolName - The name of the Kubernetes cluster agent pool.
//   - agentPoolUpdateParameters - The request body.
//   - options - AgentPoolsClientBeginUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginUpdate method.
func (client *AgentPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolUpdateParameters AgentPoolPatchParameters, options *AgentPoolsClientBeginUpdateOptions) (*runtime.Poller[AgentPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, agentPoolUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Patch the properties of the provided Kubernetes cluster agent pool, or update the tags associated with the Kubernetes
// cluster agent pool. Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *AgentPoolsClient) update(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolUpdateParameters AgentPoolPatchParameters, options *AgentPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, kubernetesClusterName, agentPoolName, agentPoolUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AgentPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, kubernetesClusterName string, agentPoolName string, agentPoolUpdateParameters AgentPoolPatchParameters, options *AgentPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/kubernetesClusters/{kubernetesClusterName}/agentPools/{agentPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if kubernetesClusterName == "" {
		return nil, errors.New("parameter kubernetesClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kubernetesClusterName}", url.PathEscape(kubernetesClusterName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, agentPoolUpdateParameters)
}
