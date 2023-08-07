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

// TestJobStreamsClient contains the methods for the TestJobStreams group.
// Don't use this type directly, use NewTestJobStreamsClient() instead.
type TestJobStreamsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTestJobStreamsClient creates a new instance of TestJobStreamsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTestJobStreamsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TestJobStreamsClient, error) {
	cl, err := arm.NewClient(moduleName+".TestJobStreamsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TestJobStreamsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve a test job stream of the test job identified by runbook name and stream id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - jobStreamID - The job stream id.
//   - options - TestJobStreamsClientGetOptions contains the optional parameters for the TestJobStreamsClient.Get method.
func (client *TestJobStreamsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string, options *TestJobStreamsClientGetOptions) (TestJobStreamsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, jobStreamID, options)
	if err != nil {
		return TestJobStreamsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TestJobStreamsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TestJobStreamsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TestJobStreamsClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string, options *TestJobStreamsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}"
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
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
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
func (client *TestJobStreamsClient) getHandleResponse(resp *http.Response) (TestJobStreamsClientGetResponse, error) {
	result := TestJobStreamsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return TestJobStreamsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTestJobPager - Retrieve a list of test job streams identified by runbook name.
//
// Generated from API version 2022-08-08
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - runbookName - The runbook name.
//   - options - TestJobStreamsClientListByTestJobOptions contains the optional parameters for the TestJobStreamsClient.NewListByTestJobPager
//     method.
func (client *TestJobStreamsClient) NewListByTestJobPager(resourceGroupName string, automationAccountName string, runbookName string, options *TestJobStreamsClientListByTestJobOptions) *runtime.Pager[TestJobStreamsClientListByTestJobResponse] {
	return runtime.NewPager(runtime.PagingHandler[TestJobStreamsClientListByTestJobResponse]{
		More: func(page TestJobStreamsClientListByTestJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TestJobStreamsClientListByTestJobResponse) (TestJobStreamsClientListByTestJobResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByTestJobCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TestJobStreamsClientListByTestJobResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TestJobStreamsClientListByTestJobResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TestJobStreamsClientListByTestJobResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTestJobHandleResponse(resp)
		},
	})
}

// listByTestJobCreateRequest creates the ListByTestJob request.
func (client *TestJobStreamsClient) listByTestJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobStreamsClientListByTestJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams"
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

// listByTestJobHandleResponse handles the ListByTestJob response.
func (client *TestJobStreamsClient) listByTestJobHandleResponse(resp *http.Response) (TestJobStreamsClientListByTestJobResponse, error) {
	result := TestJobStreamsClientListByTestJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStreamListResult); err != nil {
		return TestJobStreamsClientListByTestJobResponse{}, err
	}
	return result, nil
}
