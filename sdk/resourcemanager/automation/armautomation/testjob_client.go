//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// TestJobClient contains the methods for the TestJob group.
// Don't use this type directly, use NewTestJobClient() instead.
type TestJobClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTestJobClient creates a new instance of TestJobClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTestJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TestJobClient, error) {
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
	client := &TestJobClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create a test job of the runbook.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-30-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// runbookName - The parameters supplied to the create test job operation.
// parameters - The parameters supplied to the create test job operation.
// options - TestJobClientCreateOptions contains the optional parameters for the TestJobClient.Create method.
func (client *TestJobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters TestJobCreateParameters, options *TestJobClientCreateOptions) (TestJobClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, parameters, options)
	if err != nil {
		return TestJobClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return TestJobClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TestJobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, parameters TestJobCreateParameters, options *TestJobClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *TestJobClient) createHandleResponse(resp *http.Response) (TestJobClientCreateResponse, error) {
	result := TestJobClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestJob); err != nil {
		return TestJobClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Retrieve the test job for the specified runbook.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-30-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// runbookName - The runbook name.
// options - TestJobClientGetOptions contains the optional parameters for the TestJobClient.Get method.
func (client *TestJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientGetOptions) (TestJobClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return TestJobClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestJobClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TestJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestJobClient) getHandleResponse(resp *http.Response) (TestJobClientGetResponse, error) {
	result := TestJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestJob); err != nil {
		return TestJobClientGetResponse{}, err
	}
	return result, nil
}

// Resume - Resume the test job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-30-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// runbookName - The runbook name.
// options - TestJobClientResumeOptions contains the optional parameters for the TestJobClient.Resume method.
func (client *TestJobClient) Resume(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientResumeOptions) (TestJobClientResumeResponse, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return TestJobClientResumeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobClientResumeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestJobClientResumeResponse{}, runtime.NewResponseError(resp)
	}
	return TestJobClientResumeResponse{}, nil
}

// resumeCreateRequest creates the Resume request.
func (client *TestJobClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/resume"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - Stop the test job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-30-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// runbookName - The runbook name.
// options - TestJobClientStopOptions contains the optional parameters for the TestJobClient.Stop method.
func (client *TestJobClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientStopOptions) (TestJobClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return TestJobClientStopResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestJobClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return TestJobClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *TestJobClient) stopCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Suspend - Suspend the test job.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-30-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// runbookName - The runbook name.
// options - TestJobClientSuspendOptions contains the optional parameters for the TestJobClient.Suspend method.
func (client *TestJobClient) Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientSuspendOptions) (TestJobClientSuspendResponse, error) {
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
	if err != nil {
		return TestJobClientSuspendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobClientSuspendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestJobClientSuspendResponse{}, runtime.NewResponseError(resp)
	}
	return TestJobClientSuspendResponse{}, nil
}

// suspendCreateRequest creates the Suspend request.
func (client *TestJobClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobClientSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/suspend"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
