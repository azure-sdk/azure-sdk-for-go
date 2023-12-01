//go:build go1.18
// +build go1.18

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
	"strconv"
	"strings"
)

// AssessedSQLMachinesOperationsClient contains the methods for the AssessedSQLMachinesOperations group.
// Don't use this type directly, use NewAssessedSQLMachinesOperationsClient() instead.
type AssessedSQLMachinesOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessedSQLMachinesOperationsClient creates a new instance of AssessedSQLMachinesOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessedSQLMachinesOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessedSQLMachinesOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessedSQLMachinesOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessedSqlMachine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - SQL Assessment arm name.
//   - assessedSQLMachineName - Sql assessment Assessed Machine ARM name.
//   - options - AssessedSQLMachinesOperationsClientGetOptions contains the optional parameters for the AssessedSQLMachinesOperationsClient.Get
//     method.
func (client *AssessedSQLMachinesOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedSQLMachineName string, options *AssessedSQLMachinesOperationsClientGetOptions) (AssessedSQLMachinesOperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessedSQLMachinesOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedSQLMachineName, options)
	if err != nil {
		return AssessedSQLMachinesOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessedSQLMachinesOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessedSQLMachinesOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessedSQLMachinesOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedSQLMachineName string, options *AssessedSQLMachinesOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/sqlAssessments/{assessmentName}/assessedSqlMachines/{assessedSqlMachineName}"
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
	if assessedSQLMachineName == "" {
		return nil, errors.New("parameter assessedSQLMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedSqlMachineName}", url.PathEscape(assessedSQLMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessedSQLMachinesOperationsClient) getHandleResponse(resp *http.Response) (AssessedSQLMachinesOperationsClientGetResponse, error) {
	result := AssessedSQLMachinesOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedSQLMachine); err != nil {
		return AssessedSQLMachinesOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySQLAssessmentV2Pager - List AssessedSqlMachine resources by SqlAssessmentV2
//
// Generated from API version 2023-03-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - SQL Assessment arm name.
//   - options - AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Options contains the optional parameters for the AssessedSQLMachinesOperationsClient.NewListBySQLAssessmentV2Pager
//     method.
func (client *AssessedSQLMachinesOperationsClient) NewListBySQLAssessmentV2Pager(resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Options) *runtime.Pager[AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response] {
	return runtime.NewPager(runtime.PagingHandler[AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response]{
		More: func(page AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response) (AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessedSQLMachinesOperationsClient.NewListBySQLAssessmentV2Pager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySQLAssessmentV2CreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
			}, nil)
			if err != nil {
				return AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response{}, err
			}
			return client.listBySQLAssessmentV2HandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySQLAssessmentV2CreateRequest creates the ListBySQLAssessmentV2 request.
func (client *AssessedSQLMachinesOperationsClient) listBySQLAssessmentV2CreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Options) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/sqlAssessments/{assessmentName}/assessedSqlMachines"
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
	reqQP.Set("api-version", "2023-03-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySQLAssessmentV2HandleResponse handles the ListBySQLAssessmentV2 response.
func (client *AssessedSQLMachinesOperationsClient) listBySQLAssessmentV2HandleResponse(resp *http.Response) (AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response, error) {
	result := AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedSQLMachineListResult); err != nil {
		return AssessedSQLMachinesOperationsClientListBySQLAssessmentV2Response{}, err
	}
	return result, nil
}
