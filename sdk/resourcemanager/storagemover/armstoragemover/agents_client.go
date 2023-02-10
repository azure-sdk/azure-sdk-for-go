//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AgentsClient contains the methods for the Agents group.
// Don't use this type directly, use NewAgentsClient() instead.
type AgentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAgentsClient creates a new instance of AgentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AgentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - agentName - The name of the Agent resource.
//   - options - AgentsClientCreateOrUpdateOptions contains the optional parameters for the AgentsClient.CreateOrUpdate method.
func (client *AgentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent Agent, options *AgentsClientCreateOrUpdateOptions) (AgentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, agent, options)
	if err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AgentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent Agent, options *AgentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, agent)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AgentsClient) createOrUpdateHandleResponse(resp *http.Response) (AgentsClientCreateOrUpdateResponse, error) {
	result := AgentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes an Agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - agentName - The name of the Agent resource.
//   - options - AgentsClientBeginDeleteOptions contains the optional parameters for the AgentsClient.BeginDelete method.
func (client *AgentsClient) BeginDelete(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (*runtime.Poller[AgentsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageMoverName, agentName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[AgentsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AgentsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an Agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *AgentsClient) deleteOperation(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AgentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - agentName - The name of the Agent resource.
//   - options - AgentsClientGetOptions contains the optional parameters for the AgentsClient.Get method.
func (client *AgentsClient) Get(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientGetOptions) (AgentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, options)
	if err != nil {
		return AgentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, options *AgentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentsClient) getHandleResponse(resp *http.Response) (AgentsClientGetResponse, error) {
	result := AgentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all Agents in a Storage Mover.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - options - AgentsClientListOptions contains the optional parameters for the AgentsClient.NewListPager method.
func (client *AgentsClient) NewListPager(resourceGroupName string, storageMoverName string, options *AgentsClientListOptions) *runtime.Pager[AgentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgentsClientListResponse]{
		More: func(page AgentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgentsClientListResponse) (AgentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, storageMoverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AgentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AgentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AgentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AgentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, options *AgentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentsClient) listHandleResponse(resp *http.Response) (AgentsClientListResponse, error) {
	result := AgentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentList); err != nil {
		return AgentsClientListResponse{}, err
	}
	return result, nil
}

// Update - Creates or updates an Agent resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageMoverName - The name of the Storage Mover resource.
//   - agentName - The name of the Agent resource.
//   - options - AgentsClientUpdateOptions contains the optional parameters for the AgentsClient.Update method.
func (client *AgentsClient) Update(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent AgentUpdateParameters, options *AgentsClientUpdateOptions) (AgentsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageMoverName, agentName, agent, options)
	if err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AgentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageMoverName string, agentName string, agent AgentUpdateParameters, options *AgentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageMover/storageMovers/{storageMoverName}/agents/{agentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageMoverName == "" {
		return nil, errors.New("parameter storageMoverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageMoverName}", url.PathEscape(storageMoverName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, agent)
}

// updateHandleResponse handles the Update response.
func (client *AgentsClient) updateHandleResponse(resp *http.Response) (AgentsClientUpdateResponse, error) {
	result := AgentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Agent); err != nil {
		return AgentsClientUpdateResponse{}, err
	}
	return result, nil
}
