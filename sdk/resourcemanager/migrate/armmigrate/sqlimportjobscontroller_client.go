// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// SQLImportJobsControllerClient contains the methods for the SQLImportJobsController group.
// Don't use this type directly, use NewSQLImportJobsControllerClient() instead.
type SQLImportJobsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLImportJobsControllerClient creates a new instance of SQLImportJobsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLImportJobsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLImportJobsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLImportJobsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetImportjob - Gets the SQL inventory import job with the given job name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - jobName - Job Arm Name.
//   - options - SQLImportJobsControllerClientGetImportjobOptions contains the optional parameters for the SQLImportJobsControllerClient.GetImportjob
//     method.
func (client *SQLImportJobsControllerClient) GetImportjob(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, jobName string, options *SQLImportJobsControllerClientGetImportjobOptions) (SQLImportJobsControllerClientGetImportjobResponse, error) {
	var err error
	const operationName = "SQLImportJobsControllerClient.GetImportjob"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getImportjobCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, jobName, options)
	if err != nil {
		return SQLImportJobsControllerClientGetImportjobResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLImportJobsControllerClientGetImportjobResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLImportJobsControllerClientGetImportjobResponse{}, err
	}
	resp, err := client.getImportjobHandleResponse(httpResp)
	return resp, err
}

// getImportjobCreateRequest creates the GetImportjob request.
func (client *SQLImportJobsControllerClient) getImportjobCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, jobName string, _ *SQLImportJobsControllerClientGetImportjobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/importJobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getImportjobHandleResponse handles the GetImportjob response.
func (client *SQLImportJobsControllerClient) getImportjobHandleResponse(resp *http.Response) (SQLImportJobsControllerClientGetImportjobResponse, error) {
	result := SQLImportJobsControllerClientGetImportjobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportSQLInventoryJob); err != nil {
		return SQLImportJobsControllerClientGetImportjobResponse{}, err
	}
	return result, nil
}

// NewListImportjobsPager - Method to get all import SQL inventory job for the given site.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLImportJobsControllerClientListImportjobsOptions contains the optional parameters for the SQLImportJobsControllerClient.NewListImportjobsPager
//     method.
func (client *SQLImportJobsControllerClient) NewListImportjobsPager(resourceGroupName string, siteName string, sqlSiteName string, options *SQLImportJobsControllerClientListImportjobsOptions) *runtime.Pager[SQLImportJobsControllerClientListImportjobsResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLImportJobsControllerClientListImportjobsResponse]{
		More: func(page SQLImportJobsControllerClientListImportjobsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLImportJobsControllerClientListImportjobsResponse) (SQLImportJobsControllerClientListImportjobsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLImportJobsControllerClient.NewListImportjobsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listImportjobsCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
			}, nil)
			if err != nil {
				return SQLImportJobsControllerClientListImportjobsResponse{}, err
			}
			return client.listImportjobsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listImportjobsCreateRequest creates the ListImportjobs request.
func (client *SQLImportJobsControllerClient) listImportjobsCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, _ *SQLImportJobsControllerClientListImportjobsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/importJobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listImportjobsHandleResponse handles the ListImportjobs response.
func (client *SQLImportJobsControllerClient) listImportjobsHandleResponse(resp *http.Response) (SQLImportJobsControllerClientListImportjobsResponse, error) {
	result := SQLImportJobsControllerClientListImportjobsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedImportSQLInventoryJob); err != nil {
		return SQLImportJobsControllerClientListImportjobsResponse{}, err
	}
	return result, nil
}
