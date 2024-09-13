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

// ProjectsClient contains the methods for the Projects group.
// Don't use this type directly, use NewProjectsClient() instead.
type ProjectsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProjectsClient creates a new instance of ProjectsClient with the specified values.
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProjectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProjectsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AssessmentOptions - Get all available options for the properties of an assessment on a project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - assessmentOptionsName - Name of the assessment options. The only name accepted in default.
//   - options - ProjectsClientAssessmentOptionsOptions contains the optional parameters for the ProjectsClient.AssessmentOptions
//     method.
func (client *ProjectsClient) AssessmentOptions(ctx context.Context, resourceGroupName string, projectName string, assessmentOptionsName string, options *ProjectsClientAssessmentOptionsOptions) (ProjectsClientAssessmentOptionsResponse, error) {
	var err error
	const operationName = "ProjectsClient.AssessmentOptions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.assessmentOptionsCreateRequest(ctx, resourceGroupName, projectName, assessmentOptionsName, options)
	if err != nil {
		return ProjectsClientAssessmentOptionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientAssessmentOptionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientAssessmentOptionsResponse{}, err
	}
	resp, err := client.assessmentOptionsHandleResponse(httpResp)
	return resp, err
}

// assessmentOptionsCreateRequest creates the AssessmentOptions request.
func (client *ProjectsClient) assessmentOptionsCreateRequest(ctx context.Context, resourceGroupName string, projectName string, assessmentOptionsName string, options *ProjectsClientAssessmentOptionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions/{assessmentOptionsName}"
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
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// assessmentOptionsHandleResponse handles the AssessmentOptions response.
func (client *ProjectsClient) assessmentOptionsHandleResponse(resp *http.Response) (ProjectsClientAssessmentOptionsResponse, error) {
	result := ProjectsClientAssessmentOptionsResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentOptions); err != nil {
		return ProjectsClientAssessmentOptionsResponse{}, err
	}
	return result, nil
}

// NewAssessmentOptionsListPager - Gets list of all available options for the properties of an assessment on a project.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ProjectsClientAssessmentOptionsListOptions contains the optional parameters for the ProjectsClient.NewAssessmentOptionsListPager
//     method.
func (client *ProjectsClient) NewAssessmentOptionsListPager(resourceGroupName string, projectName string, options *ProjectsClientAssessmentOptionsListOptions) *runtime.Pager[ProjectsClientAssessmentOptionsListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientAssessmentOptionsListResponse]{
		More: func(page ProjectsClientAssessmentOptionsListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientAssessmentOptionsListResponse) (ProjectsClientAssessmentOptionsListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProjectsClient.NewAssessmentOptionsListPager")
			req, err := client.assessmentOptionsListCreateRequest(ctx, resourceGroupName, projectName, options)
			if err != nil {
				return ProjectsClientAssessmentOptionsListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProjectsClientAssessmentOptionsListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectsClientAssessmentOptionsListResponse{}, runtime.NewResponseError(resp)
			}
			return client.assessmentOptionsListHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// assessmentOptionsListCreateRequest creates the AssessmentOptionsList request.
func (client *ProjectsClient) assessmentOptionsListCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientAssessmentOptionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions"
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
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// assessmentOptionsListHandleResponse handles the AssessmentOptionsList response.
func (client *ProjectsClient) assessmentOptionsListHandleResponse(resp *http.Response) (ProjectsClientAssessmentOptionsListResponse, error) {
	result := ProjectsClientAssessmentOptionsListResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssessmentOptionsResultList); err != nil {
		return ProjectsClientAssessmentOptionsListResponse{}, err
	}
	return result, nil
}

// Create - Create a project with specified name. If a project already exists, update it.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ProjectsClientCreateOptions contains the optional parameters for the ProjectsClient.Create method.
func (client *ProjectsClient) Create(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientCreateOptions) (ProjectsClientCreateResponse, error) {
	var err error
	const operationName = "ProjectsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return ProjectsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ProjectsClient) createCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Project != nil {
		if err := runtime.MarshalAsJSON(req, *options.Project); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ProjectsClient) createHandleResponse(resp *http.Response) (ProjectsClientCreateResponse, error) {
	result := ProjectsClientCreateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the project. Deleting non-existent project is a no-operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ProjectsClientDeleteOptions contains the optional parameters for the ProjectsClient.Delete method.
func (client *ProjectsClient) Delete(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientDeleteOptions) (ProjectsClientDeleteResponse, error) {
	var err error
	const operationName = "ProjectsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return ProjectsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *ProjectsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ProjectsClient) deleteHandleResponse(resp *http.Response) (ProjectsClientDeleteResponse, error) {
	result := ProjectsClientDeleteResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	return result, nil
}

// Get - Get the project with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
func (client *ProjectsClient) Get(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientGetOptions) (ProjectsClientGetResponse, error) {
	var err error
	const operationName = "ProjectsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProjectsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}"
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
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectsClient) getHandleResponse(resp *http.Response) (ProjectsClientGetResponse, error) {
	result := ProjectsClientGetResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the projects in the resource group.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - options - ProjectsClientListOptions contains the optional parameters for the ProjectsClient.NewListPager method.
func (client *ProjectsClient) NewListPager(resourceGroupName string, options *ProjectsClientListOptions) *runtime.Pager[ProjectsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListResponse]{
		More: func(page ProjectsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListResponse) (ProjectsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProjectsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ProjectsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ProjectsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ProjectsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProjectsClient) listHandleResponse(resp *http.Response) (ProjectsClientListResponse, error) {
	result := ProjectsClientListResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectResultList); err != nil {
		return ProjectsClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get all the projects in the subscription.
//
// Generated from API version 2019-10-01
//   - options - ProjectsClientListBySubscriptionOptions contains the optional parameters for the ProjectsClient.NewListBySubscriptionPager
//     method.
func (client *ProjectsClient) NewListBySubscriptionPager(options *ProjectsClientListBySubscriptionOptions) *runtime.Pager[ProjectsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListBySubscriptionResponse]{
		More: func(page ProjectsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListBySubscriptionResponse) (ProjectsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ProjectsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ProjectsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProjectsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProjectsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/assessmentProjects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProjectsClient) listBySubscriptionHandleResponse(resp *http.Response) (ProjectsClientListBySubscriptionResponse, error) {
	result := ProjectsClientListBySubscriptionResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectResultList); err != nil {
		return ProjectsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a project with specified name. Supports partial updates, for example only tags can be provided.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - projectName - Name of the Azure Migrate project.
//   - options - ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
func (client *ProjectsClient) Update(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientUpdateOptions) (ProjectsClientUpdateResponse, error) {
	var err error
	const operationName = "ProjectsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, options)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProjectsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ProjectsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Project != nil {
		if err := runtime.MarshalAsJSON(req, *options.Project); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ProjectsClient) updateHandleResponse(resp *http.Response) (ProjectsClientUpdateResponse, error) {
	result := ProjectsClientUpdateResponse{}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	return result, nil
}