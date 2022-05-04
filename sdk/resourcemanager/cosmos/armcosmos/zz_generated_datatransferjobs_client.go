//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// DataTransferJobsClient contains the methods for the DataTransferJobs group.
// Don't use this type directly, use NewDataTransferJobsClient() instead.
type DataTransferJobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataTransferJobsClient creates a new instance of DataTransferJobsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataTransferJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataTransferJobsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DataTransferJobsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Cancel - Cancels a Data Transfer Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// jobName - Name of the Data Transfer Job
// options - DataTransferJobsClientCancelOptions contains the optional parameters for the DataTransferJobsClient.Cancel method.
func (client *DataTransferJobsClient) Cancel(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientCancelOptions) (DataTransferJobsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, accountName, jobName, options)
	if err != nil {
		return DataTransferJobsClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataTransferJobsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataTransferJobsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *DataTransferJobsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs/{jobName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *DataTransferJobsClient) cancelHandleResponse(resp *http.Response) (DataTransferJobsClientCancelResponse, error) {
	result := DataTransferJobsClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobGetResults); err != nil {
		return DataTransferJobsClientCancelResponse{}, err
	}
	return result, nil
}

// Create - Creates a Data Transfer Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// jobName - Name of the Data Transfer Job
// options - DataTransferJobsClientCreateOptions contains the optional parameters for the DataTransferJobsClient.Create method.
func (client *DataTransferJobsClient) Create(ctx context.Context, resourceGroupName string, accountName string, jobName string, jobCreateParameters CreateJobRequest, options *DataTransferJobsClientCreateOptions) (DataTransferJobsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, jobName, jobCreateParameters, options)
	if err != nil {
		return DataTransferJobsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataTransferJobsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataTransferJobsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DataTransferJobsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, jobName string, jobCreateParameters CreateJobRequest, options *DataTransferJobsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, jobCreateParameters)
}

// createHandleResponse handles the Create response.
func (client *DataTransferJobsClient) createHandleResponse(resp *http.Response) (DataTransferJobsClientCreateResponse, error) {
	result := DataTransferJobsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobGetResults); err != nil {
		return DataTransferJobsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get a Data Transfer Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// jobName - Name of the Data Transfer Job
// options - DataTransferJobsClientGetOptions contains the optional parameters for the DataTransferJobsClient.Get method.
func (client *DataTransferJobsClient) Get(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientGetOptions) (DataTransferJobsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, jobName, options)
	if err != nil {
		return DataTransferJobsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataTransferJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataTransferJobsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataTransferJobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataTransferJobsClient) getHandleResponse(resp *http.Response) (DataTransferJobsClientGetResponse, error) {
	result := DataTransferJobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobGetResults); err != nil {
		return DataTransferJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabaseAccountPager - Get a list of Data Transfer jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// options - DataTransferJobsClientListByDatabaseAccountOptions contains the optional parameters for the DataTransferJobsClient.ListByDatabaseAccount
// method.
func (client *DataTransferJobsClient) NewListByDatabaseAccountPager(resourceGroupName string, accountName string, options *DataTransferJobsClientListByDatabaseAccountOptions) *runtime.Pager[DataTransferJobsClientListByDatabaseAccountResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataTransferJobsClientListByDatabaseAccountResponse]{
		More: func(page DataTransferJobsClientListByDatabaseAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataTransferJobsClientListByDatabaseAccountResponse) (DataTransferJobsClientListByDatabaseAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataTransferJobsClientListByDatabaseAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataTransferJobsClientListByDatabaseAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataTransferJobsClientListByDatabaseAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseAccountHandleResponse(resp)
		},
	})
}

// listByDatabaseAccountCreateRequest creates the ListByDatabaseAccount request.
func (client *DataTransferJobsClient) listByDatabaseAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *DataTransferJobsClientListByDatabaseAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseAccountHandleResponse handles the ListByDatabaseAccount response.
func (client *DataTransferJobsClient) listByDatabaseAccountHandleResponse(resp *http.Response) (DataTransferJobsClientListByDatabaseAccountResponse, error) {
	result := DataTransferJobsClientListByDatabaseAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobFeedResults); err != nil {
		return DataTransferJobsClientListByDatabaseAccountResponse{}, err
	}
	return result, nil
}

// Pause - Pause a Data Transfer Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// jobName - Name of the Data Transfer Job
// options - DataTransferJobsClientPauseOptions contains the optional parameters for the DataTransferJobsClient.Pause method.
func (client *DataTransferJobsClient) Pause(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientPauseOptions) (DataTransferJobsClientPauseResponse, error) {
	req, err := client.pauseCreateRequest(ctx, resourceGroupName, accountName, jobName, options)
	if err != nil {
		return DataTransferJobsClientPauseResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataTransferJobsClientPauseResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataTransferJobsClientPauseResponse{}, runtime.NewResponseError(resp)
	}
	return client.pauseHandleResponse(resp)
}

// pauseCreateRequest creates the Pause request.
func (client *DataTransferJobsClient) pauseCreateRequest(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientPauseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs/{jobName}/pause"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// pauseHandleResponse handles the Pause response.
func (client *DataTransferJobsClient) pauseHandleResponse(resp *http.Response) (DataTransferJobsClientPauseResponse, error) {
	result := DataTransferJobsClientPauseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobGetResults); err != nil {
		return DataTransferJobsClientPauseResponse{}, err
	}
	return result, nil
}

// Resume - Resumes a Data Transfer Job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - Cosmos DB database account name.
// jobName - Name of the Data Transfer Job
// options - DataTransferJobsClientResumeOptions contains the optional parameters for the DataTransferJobsClient.Resume method.
func (client *DataTransferJobsClient) Resume(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientResumeOptions) (DataTransferJobsClientResumeResponse, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, accountName, jobName, options)
	if err != nil {
		return DataTransferJobsClientResumeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataTransferJobsClientResumeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataTransferJobsClientResumeResponse{}, runtime.NewResponseError(resp)
	}
	return client.resumeHandleResponse(resp)
}

// resumeCreateRequest creates the Resume request.
func (client *DataTransferJobsClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, accountName string, jobName string, options *DataTransferJobsClientResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/dataTransferJobs/{jobName}/resume"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resumeHandleResponse handles the Resume response.
func (client *DataTransferJobsClient) resumeHandleResponse(resp *http.Response) (DataTransferJobsClientResumeResponse, error) {
	result := DataTransferJobsClientResumeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataTransferJobGetResults); err != nil {
		return DataTransferJobsClientResumeResponse{}, err
	}
	return result, nil
}
