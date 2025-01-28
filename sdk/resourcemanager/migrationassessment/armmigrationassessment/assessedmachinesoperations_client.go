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

// AssessedMachinesOperationsClient contains the methods for the AssessedMachinesOperations group.
// Don't use this type directly, use NewAssessedMachinesOperationsClient() instead.
type AssessedMachinesOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessedMachinesOperationsClient creates a new instance of AssessedMachinesOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessedMachinesOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessedMachinesOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessedMachinesOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessedMachine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Machine Assessment ARM name
//   - assessedMachineName - Machine assessment Assessed Machine ARM name
//   - options - AssessedMachinesOperationsClientGetOptions contains the optional parameters for the AssessedMachinesOperationsClient.Get
//     method.
func (client *AssessedMachinesOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string, options *AssessedMachinesOperationsClientGetOptions) (AssessedMachinesOperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessedMachinesOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedMachineName, options)
	if err != nil {
		return AssessedMachinesOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessedMachinesOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessedMachinesOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessedMachinesOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string, options *AssessedMachinesOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}"
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
	if assessedMachineName == "" {
		return nil, errors.New("parameter assessedMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedMachineName}", url.PathEscape(assessedMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssessedMachinesOperationsClient) getHandleResponse(resp *http.Response) (AssessedMachinesOperationsClientGetResponse, error) {
	result := AssessedMachinesOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedMachine); err != nil {
		return AssessedMachinesOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAssessmentPager - List AssessedMachine resources by Assessment
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Machine Assessment ARM name
//   - options - AssessedMachinesOperationsClientListByAssessmentOptions contains the optional parameters for the AssessedMachinesOperationsClient.NewListByAssessmentPager
//     method.
func (client *AssessedMachinesOperationsClient) NewListByAssessmentPager(resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedMachinesOperationsClientListByAssessmentOptions) *runtime.Pager[AssessedMachinesOperationsClientListByAssessmentResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssessedMachinesOperationsClientListByAssessmentResponse]{
		More: func(page AssessedMachinesOperationsClientListByAssessmentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessedMachinesOperationsClientListByAssessmentResponse) (AssessedMachinesOperationsClientListByAssessmentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessedMachinesOperationsClient.NewListByAssessmentPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
			}, nil)
			if err != nil {
				return AssessedMachinesOperationsClientListByAssessmentResponse{}, err
			}
			return client.listByAssessmentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentCreateRequest creates the ListByAssessment request.
func (client *AssessedMachinesOperationsClient) listByAssessmentCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedMachinesOperationsClientListByAssessmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines"
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
	reqQP.Set("api-version", "2024-01-01-preview")
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

// listByAssessmentHandleResponse handles the ListByAssessment response.
func (client *AssessedMachinesOperationsClient) listByAssessmentHandleResponse(resp *http.Response) (AssessedMachinesOperationsClientListByAssessmentResponse, error) {
	result := AssessedMachinesOperationsClientListByAssessmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedMachineListResult); err != nil {
		return AssessedMachinesOperationsClientListByAssessmentResponse{}, err
	}
	return result, nil
}
