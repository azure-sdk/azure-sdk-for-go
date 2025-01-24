//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationassessment

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AssessedSQLInstanceV2OperationsClient contains the methods for the AssessedSQLInstanceV2Operations group.
// Don't use this type directly, use NewAssessedSQLInstanceV2OperationsClient() instead.
type AssessedSQLInstanceV2OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessedSQLInstanceV2OperationsClient creates a new instance of AssessedSQLInstanceV2OperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessedSQLInstanceV2OperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessedSQLInstanceV2OperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessedSQLInstanceV2OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessedSqlInstanceV2
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-09-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - SQL Assessment arm name.
//   - assessedSQLInstanceName - Sql assessment Assessed Instance ARM name.
//   - options - AssessedSQLInstanceV2OperationsClientGetOptions contains the optional parameters for the AssessedSQLInstanceV2OperationsClient.Get
//     method.
func (client *AssessedSQLInstanceV2OperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedSQLInstanceName string, options *AssessedSQLInstanceV2OperationsClientGetOptions) (AssessedSQLInstanceV2OperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessedSQLInstanceV2OperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedSQLInstanceName, options)
	if err != nil {
		return AssessedSQLInstanceV2OperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessedSQLInstanceV2OperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessedSQLInstanceV2OperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessedSQLInstanceV2OperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedSQLInstanceName string, options *AssessedSQLInstanceV2OperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/sqlAssessments/{assessmentName}/assessedSqlInstances/{assessedSqlInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assessedSQLInstanceName == "" {
		return nil, errors.New("parameter assessedSQLInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedSqlInstanceName}", url.PathEscape(assessedSQLInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-09-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessedSQLInstanceV2OperationsClient) getHandleResponse(resp *http.Response) (AssessedSQLInstanceV2OperationsClientGetResponse, error) {
	result := AssessedSQLInstanceV2OperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedSQLInstanceV2); err != nil {
		return AssessedSQLInstanceV2OperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySQLAssessmentV2Pager - List AssessedSqlInstanceV2 resources by SqlAssessmentV2
//
// Generated from API version 2023-09-09-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - SQL Assessment arm name.
//   - options - AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Options contains the optional parameters for the AssessedSQLInstanceV2OperationsClient.NewListBySQLAssessmentV2Pager
//     method.
func (client *AssessedSQLInstanceV2OperationsClient) NewListBySQLAssessmentV2Pager(resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Options) *runtime.Pager[AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response] {
	return runtime.NewPager(runtime.PagingHandler[AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response]{
		More: func(page AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response) (AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessedSQLInstanceV2OperationsClient.NewListBySQLAssessmentV2Pager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySQLAssessmentV2CreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
			}, nil)
			if err != nil {
				return AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response{}, err
			}
			return client.listBySQLAssessmentV2HandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySQLAssessmentV2CreateRequest creates the ListBySQLAssessmentV2 request.
func (client *AssessedSQLInstanceV2OperationsClient) listBySQLAssessmentV2CreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Options) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/sqlAssessments/{assessmentName}/assessedSqlInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-09-09-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySQLAssessmentV2HandleResponse handles the ListBySQLAssessmentV2 response.
func (client *AssessedSQLInstanceV2OperationsClient) listBySQLAssessmentV2HandleResponse(resp *http.Response) (AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response, error) {
	result := AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedSQLInstanceV2ListResult); err != nil {
		return AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Response{}, err
	}
	return result, nil
}
