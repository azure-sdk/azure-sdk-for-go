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

// AssessedWebApplicationOperationsClient contains the methods for the AssessedWebApplicationOperations group.
// Don't use this type directly, use NewAssessedWebApplicationOperationsClient() instead.
type AssessedWebApplicationOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessedWebApplicationOperationsClient creates a new instance of AssessedWebApplicationOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessedWebApplicationOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessedWebApplicationOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessedWebApplicationOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessedWebApplication
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - assessmentName - AKS Assessment Name.
//   - assessedWorkload - Assessed Web Application Name.
//   - options - AssessedWebApplicationOperationsClientGetOptions contains the optional parameters for the AssessedWebApplicationOperationsClient.Get
//     method.
func (client *AssessedWebApplicationOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, assessedWorkload string, options *AssessedWebApplicationOperationsClientGetOptions) (AssessedWebApplicationOperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessedWebApplicationOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, assessmentName, assessedWorkload, options)
	if err != nil {
		return AssessedWebApplicationOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessedWebApplicationOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessedWebApplicationOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessedWebApplicationOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, assessedWorkload string, _ *AssessedWebApplicationOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/aksAssessments/{assessmentName}/assessedWebApps/{assessedWorkload}"
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
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assessedWorkload == "" {
		return nil, errors.New("parameter assessedWorkload cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessedWorkload}", url.PathEscape(assessedWorkload))
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
func (client *AssessedWebApplicationOperationsClient) getHandleResponse(resp *http.Response) (AssessedWebApplicationOperationsClientGetResponse, error) {
	result := AssessedWebApplicationOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedWebApplication); err != nil {
		return AssessedWebApplicationOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAksAssessmentPager - List AssessedWebApplication resources by AKSAssessment
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - assessmentName - AKS Assessment Name.
//   - options - AssessedWebApplicationOperationsClientListByAksAssessmentOptions contains the optional parameters for the AssessedWebApplicationOperationsClient.NewListByAksAssessmentPager
//     method.
func (client *AssessedWebApplicationOperationsClient) NewListByAksAssessmentPager(resourceGroupName string, projectName string, assessmentName string, options *AssessedWebApplicationOperationsClientListByAksAssessmentOptions) *runtime.Pager[AssessedWebApplicationOperationsClientListByAksAssessmentResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssessedWebApplicationOperationsClientListByAksAssessmentResponse]{
		More: func(page AssessedWebApplicationOperationsClientListByAksAssessmentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessedWebApplicationOperationsClientListByAksAssessmentResponse) (AssessedWebApplicationOperationsClientListByAksAssessmentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessedWebApplicationOperationsClient.NewListByAksAssessmentPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAksAssessmentCreateRequest(ctx, resourceGroupName, projectName, assessmentName, options)
			}, nil)
			if err != nil {
				return AssessedWebApplicationOperationsClientListByAksAssessmentResponse{}, err
			}
			return client.listByAksAssessmentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAksAssessmentCreateRequest creates the ListByAksAssessment request.
func (client *AssessedWebApplicationOperationsClient) listByAksAssessmentCreateRequest(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, options *AssessedWebApplicationOperationsClientListByAksAssessmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/aksAssessments/{assessmentName}/assessedWebApps"
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-01-01-preview")
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

// listByAksAssessmentHandleResponse handles the ListByAksAssessment response.
func (client *AssessedWebApplicationOperationsClient) listByAksAssessmentHandleResponse(resp *http.Response) (AssessedWebApplicationOperationsClientListByAksAssessmentResponse, error) {
	result := AssessedWebApplicationOperationsClientListByAksAssessmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessedWebApplicationListResult); err != nil {
		return AssessedWebApplicationOperationsClientListByAksAssessmentResponse{}, err
	}
	return result, nil
}
