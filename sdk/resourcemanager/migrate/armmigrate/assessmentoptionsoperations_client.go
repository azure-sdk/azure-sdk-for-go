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
	"strings"
)

// AssessmentOptionsOperationsClient contains the methods for the AssessmentOptionsOperations group.
// Don't use this type directly, use NewAssessmentOptionsOperationsClient() instead.
type AssessmentOptionsOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssessmentOptionsOperationsClient creates a new instance of AssessmentOptionsOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssessmentOptionsOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssessmentOptionsOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssessmentOptionsOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AssessmentOptions
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - assessmentOptionsName - assessment options ARM name. Accepted value is 'default'
//   - options - AssessmentOptionsOperationsClientGetOptions contains the optional parameters for the AssessmentOptionsOperationsClient.Get
//     method.
func (client *AssessmentOptionsOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, assessmentOptionsName string, options *AssessmentOptionsOperationsClientGetOptions) (AssessmentOptionsOperationsClientGetResponse, error) {
	var err error
	const operationName = "AssessmentOptionsOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, assessmentOptionsName, options)
	if err != nil {
		return AssessmentOptionsOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssessmentOptionsOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssessmentOptionsOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssessmentOptionsOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, assessmentOptionsName string, options *AssessmentOptionsOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions/{assessmentOptionsName}"
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
	if assessmentOptionsName == "" {
		return nil, errors.New("parameter assessmentOptionsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentOptionsName}", url.PathEscape(assessmentOptionsName))
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
func (client *AssessmentOptionsOperationsClient) getHandleResponse(resp *http.Response) (AssessmentOptionsOperationsClientGetResponse, error) {
	result := AssessmentOptionsOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentOptions); err != nil {
		return AssessmentOptionsOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAssessmentProjectPager - List AssessmentOptions resources by AssessmentProject
//
// Generated from API version 2023-03-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - options - AssessmentOptionsOperationsClientListByAssessmentProjectOptions contains the optional parameters for the AssessmentOptionsOperationsClient.NewListByAssessmentProjectPager
//     method.
func (client *AssessmentOptionsOperationsClient) NewListByAssessmentProjectPager(resourceGroupName string, projectName string, options *AssessmentOptionsOperationsClientListByAssessmentProjectOptions) *runtime.Pager[AssessmentOptionsOperationsClientListByAssessmentProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssessmentOptionsOperationsClientListByAssessmentProjectResponse]{
		More: func(page AssessmentOptionsOperationsClientListByAssessmentProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssessmentOptionsOperationsClientListByAssessmentProjectResponse) (AssessmentOptionsOperationsClientListByAssessmentProjectResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssessmentOptionsOperationsClient.NewListByAssessmentProjectPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAssessmentProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			}, nil)
			if err != nil {
				return AssessmentOptionsOperationsClientListByAssessmentProjectResponse{}, err
			}
			return client.listByAssessmentProjectHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAssessmentProjectCreateRequest creates the ListByAssessmentProject request.
func (client *AssessmentOptionsOperationsClient) listByAssessmentProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *AssessmentOptionsOperationsClientListByAssessmentProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions"
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

// listByAssessmentProjectHandleResponse handles the ListByAssessmentProject response.
func (client *AssessmentOptionsOperationsClient) listByAssessmentProjectHandleResponse(resp *http.Response) (AssessmentOptionsOperationsClientListByAssessmentProjectResponse, error) {
	result := AssessmentOptionsOperationsClientListByAssessmentProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentOptionsListResult); err != nil {
		return AssessmentOptionsOperationsClientListByAssessmentProjectResponse{}, err
	}
	return result, nil
}
