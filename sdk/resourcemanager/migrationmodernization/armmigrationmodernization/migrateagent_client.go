// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationmodernization

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

// MigrateAgentClient contains the methods for the MigrateAgent group.
// Don't use this type directly, use NewMigrateAgentClient() instead.
type MigrateAgentClient struct {
	internal *arm.Client
}

// NewMigrateAgentClient creates a new instance of MigrateAgentClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMigrateAgentClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MigrateAgentClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MigrateAgentClient{
		internal: cl,
	}
	return client, nil
}

// Create - Creates the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject name.
//   - agentName - MigrateAgent name.
//   - body - MigrateAgent model.
//   - options - MigrateAgentClientCreateOptions contains the optional parameters for the MigrateAgentClient.Create method.
func (client *MigrateAgentClient) Create(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, body MigrateAgentModel, options *MigrateAgentClientCreateOptions) (MigrateAgentClientCreateResponse, error) {
	var err error
	const operationName = "MigrateAgentClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, body, options)
	if err != nil {
		return MigrateAgentClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MigrateAgentClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MigrateAgentClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *MigrateAgentClient) createCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, body MigrateAgentModel, _ *MigrateAgentClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/migrateAgents/{agentName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *MigrateAgentClient) createHandleResponse(resp *http.Response) (MigrateAgentClientCreateResponse, error) {
	result := MigrateAgentClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrateAgentModel); err != nil {
		return MigrateAgentClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject Name.
//   - agentName - MigrateAgent Name.
//   - options - MigrateAgentClientBeginDeleteOptions contains the optional parameters for the MigrateAgentClient.BeginDelete
//     method.
func (client *MigrateAgentClient) BeginDelete(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, options *MigrateAgentClientBeginDeleteOptions) (*runtime.Poller[MigrateAgentClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MigrateAgentClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MigrateAgentClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
func (client *MigrateAgentClient) deleteOperation(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, options *MigrateAgentClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MigrateAgentClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, options)
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
func (client *MigrateAgentClient) deleteCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, _ *MigrateAgentClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/migrateAgents/{agentName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject name.
//   - agentName - MigrateAgent name.
//   - options - MigrateAgentClientGetOptions contains the optional parameters for the MigrateAgentClient.Get method.
func (client *MigrateAgentClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, options *MigrateAgentClientGetOptions) (MigrateAgentClientGetResponse, error) {
	var err error
	const operationName = "MigrateAgentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, options)
	if err != nil {
		return MigrateAgentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MigrateAgentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MigrateAgentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MigrateAgentClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, _ *MigrateAgentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/migrateAgents/{agentName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrateAgentClient) getHandleResponse(resp *http.Response) (MigrateAgentClientGetResponse, error) {
	result := MigrateAgentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrateAgentModel); err != nil {
		return MigrateAgentClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of modernizeProject agents in the given modernizeProject.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject name.
//   - options - MigrateAgentClientListOptions contains the optional parameters for the MigrateAgentClient.List method.
func (client *MigrateAgentClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, options *MigrateAgentClientListOptions) (MigrateAgentClientListResponse, error) {
	var err error
	const operationName = "MigrateAgentClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, options)
	if err != nil {
		return MigrateAgentClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MigrateAgentClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MigrateAgentClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *MigrateAgentClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, _ *MigrateAgentClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/migrateAgents"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MigrateAgentClient) listHandleResponse(resp *http.Response) (MigrateAgentClientListResponse, error) {
	result := MigrateAgentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrateAgentModelCollection); err != nil {
		return MigrateAgentClientListResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refreshes the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject name.
//   - agentName - MigrateAgent name.
//   - options - MigrateAgentClientBeginRefreshOptions contains the optional parameters for the MigrateAgentClient.BeginRefresh
//     method.
func (client *MigrateAgentClient) BeginRefresh(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, options *MigrateAgentClientBeginRefreshOptions) (*runtime.Poller[MigrateAgentClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MigrateAgentClientRefreshResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MigrateAgentClientRefreshResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Refresh - Refreshes the modernizeProject agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
func (client *MigrateAgentClient) refresh(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, options *MigrateAgentClientBeginRefreshOptions) (*http.Response, error) {
	var err error
	const operationName = "MigrateAgentClient.BeginRefresh"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, agentName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *MigrateAgentClient) refreshCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, agentName string, _ *MigrateAgentClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/migrateAgents/{agentName}/refresh"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	if agentName == "" {
		return nil, errors.New("parameter agentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentName}", url.PathEscape(agentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
