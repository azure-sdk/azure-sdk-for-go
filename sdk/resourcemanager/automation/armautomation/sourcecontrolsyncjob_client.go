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

// SourceControlSyncJobClient contains the methods for the SourceControlSyncJob group.
// Don't use this type directly, use NewSourceControlSyncJobClient() instead.
type SourceControlSyncJobClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSourceControlSyncJobClient creates a new instance of SourceControlSyncJobClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSourceControlSyncJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SourceControlSyncJobClient, error) {
	cl, err := arm.NewClient(moduleName+".SourceControlSyncJobClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SourceControlSyncJobClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates the sync job for a source control.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - sourceControlName - The source control name.
//   - sourceControlSyncJobID - The source control sync job id.
//   - parameters - The parameters supplied to the create source control sync job operation.
//   - options - SourceControlSyncJobClientCreateOptions contains the optional parameters for the SourceControlSyncJobClient.Create
//     method.
func (client *SourceControlSyncJobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, parameters SourceControlSyncJobCreateParameters, options *SourceControlSyncJobClientCreateOptions) (SourceControlSyncJobClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, parameters, options)
	if err != nil {
		return SourceControlSyncJobClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SourceControlSyncJobClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SourceControlSyncJobClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SourceControlSyncJobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, parameters SourceControlSyncJobCreateParameters, options *SourceControlSyncJobClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlSyncJobId}", url.PathEscape(sourceControlSyncJobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SourceControlSyncJobClient) createHandleResponse(resp *http.Response) (SourceControlSyncJobClientCreateResponse, error) {
	result := SourceControlSyncJobClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlSyncJob); err != nil {
		return SourceControlSyncJobClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Retrieve the source control sync job identified by job id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - sourceControlName - The source control name.
//   - sourceControlSyncJobID - The source control sync job id.
//   - options - SourceControlSyncJobClientGetOptions contains the optional parameters for the SourceControlSyncJobClient.Get
//     method.
func (client *SourceControlSyncJobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, options *SourceControlSyncJobClientGetOptions) (SourceControlSyncJobClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, sourceControlSyncJobID, options)
	if err != nil {
		return SourceControlSyncJobClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SourceControlSyncJobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SourceControlSyncJobClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SourceControlSyncJobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, sourceControlSyncJobID string, options *SourceControlSyncJobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs/{sourceControlSyncJobId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlSyncJobId}", url.PathEscape(sourceControlSyncJobID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SourceControlSyncJobClient) getHandleResponse(resp *http.Response) (SourceControlSyncJobClientGetResponse, error) {
	result := SourceControlSyncJobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlSyncJobByID); err != nil {
		return SourceControlSyncJobClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of source control sync jobs.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - sourceControlName - The source control name.
//   - options - SourceControlSyncJobClientListByAutomationAccountOptions contains the optional parameters for the SourceControlSyncJobClient.NewListByAutomationAccountPager
//     method.
func (client *SourceControlSyncJobClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlSyncJobClientListByAutomationAccountOptions) *runtime.Pager[SourceControlSyncJobClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[SourceControlSyncJobClientListByAutomationAccountResponse]{
		More: func(page SourceControlSyncJobClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlSyncJobClientListByAutomationAccountResponse) (SourceControlSyncJobClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, sourceControlName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SourceControlSyncJobClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SourceControlSyncJobClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SourceControlSyncJobClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *SourceControlSyncJobClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, sourceControlName string, options *SourceControlSyncJobClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/sourceControls/{sourceControlName}/sourceControlSyncJobs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if sourceControlName == "" {
		return nil, errors.New("parameter sourceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceControlName}", url.PathEscape(sourceControlName))
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
	reqQP.Set("api-version", "2022-08-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *SourceControlSyncJobClient) listByAutomationAccountHandleResponse(resp *http.Response) (SourceControlSyncJobClientListByAutomationAccountResponse, error) {
	result := SourceControlSyncJobClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceControlSyncJobListResult); err != nil {
		return SourceControlSyncJobClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
