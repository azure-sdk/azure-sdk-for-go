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

// JobCredentialsClient contains the methods for the JobCredentials group.
// Don't use this type directly, use NewJobCredentialsClient() instead.
type JobCredentialsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobCredentialsClient creates a new instance of JobCredentialsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobCredentialsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobCredentialsClient, error) {
	cl, err := arm.NewClient(moduleName+".JobCredentialsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobCredentialsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a job credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - credentialName - The name of the credential.
//   - parameters - The requested job credential state.
//   - options - JobCredentialsClientCreateOrUpdateOptions contains the optional parameters for the JobCredentialsClient.CreateOrUpdate
//     method.
func (client *JobCredentialsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, parameters JobCredential, options *JobCredentialsClientCreateOrUpdateOptions) (JobCredentialsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, credentialName, parameters, options)
	if err != nil {
		return JobCredentialsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCredentialsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return JobCredentialsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobCredentialsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, parameters JobCredential, options *JobCredentialsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobCredentialsClient) createOrUpdateHandleResponse(resp *http.Response) (JobCredentialsClientCreateOrUpdateResponse, error) {
	result := JobCredentialsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCredential); err != nil {
		return JobCredentialsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a job credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - credentialName - The name of the credential.
//   - options - JobCredentialsClientDeleteOptions contains the optional parameters for the JobCredentialsClient.Delete method.
func (client *JobCredentialsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, options *JobCredentialsClientDeleteOptions) (JobCredentialsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, credentialName, options)
	if err != nil {
		return JobCredentialsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCredentialsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return JobCredentialsClientDeleteResponse{}, err
	}
	return JobCredentialsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobCredentialsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, options *JobCredentialsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a jobs credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - credentialName - The name of the credential.
//   - options - JobCredentialsClientGetOptions contains the optional parameters for the JobCredentialsClient.Get method.
func (client *JobCredentialsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, options *JobCredentialsClientGetOptions) (JobCredentialsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, credentialName, options)
	if err != nil {
		return JobCredentialsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCredentialsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobCredentialsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobCredentialsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, credentialName string, options *JobCredentialsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials/{credentialName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobCredentialsClient) getHandleResponse(resp *http.Response) (JobCredentialsClientGetResponse, error) {
	result := JobCredentialsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCredential); err != nil {
		return JobCredentialsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAgentPager - Gets a list of jobs credentials.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - options - JobCredentialsClientListByAgentOptions contains the optional parameters for the JobCredentialsClient.NewListByAgentPager
//     method.
func (client *JobCredentialsClient) NewListByAgentPager(resourceGroupName string, serverName string, jobAgentName string, options *JobCredentialsClientListByAgentOptions) *runtime.Pager[JobCredentialsClientListByAgentResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobCredentialsClientListByAgentResponse]{
		More: func(page JobCredentialsClientListByAgentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobCredentialsClientListByAgentResponse) (JobCredentialsClientListByAgentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobCredentialsClientListByAgentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobCredentialsClientListByAgentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobCredentialsClientListByAgentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAgentHandleResponse(resp)
		},
	})
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobCredentialsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobCredentialsClientListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/credentials"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobCredentialsClient) listByAgentHandleResponse(resp *http.Response) (JobCredentialsClientListByAgentResponse, error) {
	result := JobCredentialsClientListByAgentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCredentialListResult); err != nil {
		return JobCredentialsClientListByAgentResponse{}, err
	}
	return result, nil
}
