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

// AssessedWebAppV2OperationsClient contains the methods for the AssessedWebAppV2Operations group.
// Don't use this type directly, use NewAssessedWebAppV2OperationsClient() instead.
type AssessedWebAppV2OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessedWebAppV2OperationsClient creates a new instance of AssessedWebAppV2OperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessedWebAppV2OperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessedWebAppV2OperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessedWebAppV2OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessedWebAppV2
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Web app Assessment arm name.
//   - assessedWebAppName - Assessed web app ARM name.
//   - options - AssessedWebAppV2OperationsClientGetOptions contains the optional parameters for the AssessedWebAppV2OperationsClient.Get
//     method.
func (client *AssessedWebAppV2OperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedWebAppName string, options *AssessedWebAppV2OperationsClientGetOptions) (AssessedWebAppV2OperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessedWebAppV2OperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, assessedWebAppName, options)
	if err != nil {
		return AssessedWebAppV2OperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessedWebAppV2OperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessedWebAppV2OperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessedWebAppV2OperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedWebAppName string, options *AssessedWebAppV2OperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/webAppAssessments/{assessmentName}/assessedWebApps/{assessedWebAppName}"
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
	if assessedWebAppName == "" {
		return nil, errors.New("parameter assessedWebAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedWebAppName}", url.PathEscape(assessedWebAppName))
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
func (client *AssessedWebAppV2OperationsClient) getHandleResponse(resp *http.Response) (AssessedWebAppV2OperationsClientGetResponse, error) {
	result := AssessedWebAppV2OperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedWebAppV2); err != nil {
		return AssessedWebAppV2OperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWebAppAssessmentV2Pager - List AssessedWebAppV2 resources by WebAppAssessmentV2
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - groupName - Group ARM name
//   - assessmentName - Web app Assessment arm name.
//   - options - AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Options contains the optional parameters for the AssessedWebAppV2OperationsClient.NewListByWebAppAssessmentV2Pager
//     method.
func (client *AssessedWebAppV2OperationsClient) NewListByWebAppAssessmentV2Pager(resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Options) *runtime.Pager[AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response] {
	return runtime.NewPager(runtime.PagingHandler[AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response]{
		More: func(page AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response) (AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessedWebAppV2OperationsClient.NewListByWebAppAssessmentV2Pager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppAssessmentV2CreateRequest(ctx, resourceGroupName, projectName, groupName, assessmentName, options)
			}, nil)
			if err != nil {
				return AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response{}, err
			}
			return client.listByWebAppAssessmentV2HandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppAssessmentV2CreateRequest creates the ListByWebAppAssessmentV2 request.
func (client *AssessedWebAppV2OperationsClient) listByWebAppAssessmentV2CreateRequest(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Options) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/webAppAssessments/{assessmentName}/assessedWebApps"
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

// listByWebAppAssessmentV2HandleResponse handles the ListByWebAppAssessmentV2 response.
func (client *AssessedWebAppV2OperationsClient) listByWebAppAssessmentV2HandleResponse(resp *http.Response) (AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response, error) {
	result := AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedWebAppV2ListResult); err != nil {
		return AssessedWebAppV2OperationsClientListByWebAppAssessmentV2Response{}, err
	}
	return result, nil
}
