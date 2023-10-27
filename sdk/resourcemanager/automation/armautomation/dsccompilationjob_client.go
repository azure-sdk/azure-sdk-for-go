//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// DscCompilationJobClient contains the methods for the DscCompilationJob group.
// Don't use this type directly, use NewDscCompilationJobClient() instead.
type DscCompilationJobClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDscCompilationJobClient creates a new instance of DscCompilationJobClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDscCompilationJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DscCompilationJobClient, error) {
	cl, err := arm.NewClient(moduleName+".DscCompilationJobClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DscCompilationJobClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - compilationJobName - The DSC configuration Id.
//   - parameters - The parameters supplied to the create compilation job operation.
//   - options - DscCompilationJobClientBeginCreateOptions contains the optional parameters for the DscCompilationJobClient.BeginCreate
//     method.
func (client *DscCompilationJobClient) BeginCreate(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*runtime.Poller[DscCompilationJobClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DscCompilationJobClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DscCompilationJobClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates the Dsc compilation job of the configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
func (client *DscCompilationJobClient) create(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *DscCompilationJobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters DscCompilationJobCreateParameters, options *DscCompilationJobClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if compilationJobName == "" {
		return nil, errors.New("parameter compilationJobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{compilationJobName}", url.PathEscape(compilationJobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Retrieve the Dsc configuration compilation job identified by job id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - compilationJobName - The DSC configuration Id.
//   - options - DscCompilationJobClientGetOptions contains the optional parameters for the DscCompilationJobClient.Get method.
func (client *DscCompilationJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobClientGetOptions) (DscCompilationJobClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, compilationJobName, options)
	if err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscCompilationJobClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DscCompilationJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *DscCompilationJobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if compilationJobName == "" {
		return nil, errors.New("parameter compilationJobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{compilationJobName}", url.PathEscape(compilationJobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscCompilationJobClient) getHandleResponse(resp *http.Response) (DscCompilationJobClientGetResponse, error) {
	result := DscCompilationJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscCompilationJob); err != nil {
		return DscCompilationJobClientGetResponse{}, err
	}
	return result, nil
}

// GetStream - Retrieve the job stream identified by job stream id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - jobID - The job id.
//   - jobStreamID - The job stream id.
//   - options - DscCompilationJobClientGetStreamOptions contains the optional parameters for the DscCompilationJobClient.GetStream
//     method.
func (client *DscCompilationJobClient) GetStream(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobClientGetStreamOptions) (DscCompilationJobClientGetStreamResponse, error) {
	var err error
	req, err := client.getStreamCreateRequest(ctx, resourceGroupName, automationAccountName, jobID, jobStreamID, options)
	if err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	resp, err := client.getStreamHandleResponse(httpResp)
	return resp, err
}

// getStreamCreateRequest creates the GetStream request.
func (client *DscCompilationJobClient) getStreamCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *DscCompilationJobClientGetStreamOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams/{jobStreamId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	urlPath = strings.ReplaceAll(urlPath, "{jobId}", url.PathEscape(jobID))
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStreamHandleResponse handles the GetStream response.
func (client *DscCompilationJobClient) getStreamHandleResponse(resp *http.Response) (DscCompilationJobClientGetStreamResponse, error) {
	result := DscCompilationJobClientGetStreamResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return DscCompilationJobClientGetStreamResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of dsc compilation jobs.
//
// Generated from API version 2023-05-15-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - DscCompilationJobClientListByAutomationAccountOptions contains the optional parameters for the DscCompilationJobClient.NewListByAutomationAccountPager
//     method.
func (client *DscCompilationJobClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *DscCompilationJobClientListByAutomationAccountOptions) *runtime.Pager[DscCompilationJobClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[DscCompilationJobClientListByAutomationAccountResponse]{
		More: func(page DscCompilationJobClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DscCompilationJobClientListByAutomationAccountResponse) (DscCompilationJobClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DscCompilationJobClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DscCompilationJobClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DscCompilationJobClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscCompilationJobClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscCompilationJobClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscCompilationJobClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscCompilationJobClientListByAutomationAccountResponse, error) {
	result := DscCompilationJobClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscCompilationJobListResult); err != nil {
		return DscCompilationJobClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
