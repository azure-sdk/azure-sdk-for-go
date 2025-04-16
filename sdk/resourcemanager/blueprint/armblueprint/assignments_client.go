// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// AssignmentsClient contains the methods for the Assignments group.
// Don't use this type directly, use NewAssignmentsClient() instead.
type AssignmentsClient struct {
	internal *arm.Client
}

// NewAssignmentsClient creates a new instance of AssignmentsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssignmentsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a blueprint assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01-preview
//   - resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: '/subscriptions/{subscriptionId}').
//   - assignmentName - Name of the blueprint assignment.
//   - assignment - Blueprint assignment object to save.
//   - options - AssignmentsClientCreateOrUpdateOptions contains the optional parameters for the AssignmentsClient.CreateOrUpdate
//     method.
func (client *AssignmentsClient) CreateOrUpdate(ctx context.Context, resourceScope string, assignmentName string, assignment Assignment, options *AssignmentsClientCreateOrUpdateOptions) (AssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AssignmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceScope, assignmentName, assignment, options)
	if err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceScope string, assignmentName string, assignment Assignment, _ *AssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, assignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (AssignmentsClientCreateOrUpdateResponse, error) {
	result := AssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a blueprint assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01-preview
//   - resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: '/subscriptions/{subscriptionId}').
//   - assignmentName - Name of the blueprint assignment.
//   - options - AssignmentsClientDeleteOptions contains the optional parameters for the AssignmentsClient.Delete method.
func (client *AssignmentsClient) Delete(ctx context.Context, resourceScope string, assignmentName string, options *AssignmentsClientDeleteOptions) (AssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "AssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceScope, assignmentName, options)
	if err != nil {
		return AssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *AssignmentsClient) deleteCreateRequest(ctx context.Context, resourceScope string, assignmentName string, options *AssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	if options != nil && options.DeleteBehavior != nil {
		reqQP.Set("deleteBehavior", string(*options.DeleteBehavior))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *AssignmentsClient) deleteHandleResponse(resp *http.Response) (AssignmentsClientDeleteResponse, error) {
	result := AssignmentsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get a blueprint assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01-preview
//   - resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: '/subscriptions/{subscriptionId}').
//   - assignmentName - Name of the blueprint assignment.
//   - options - AssignmentsClientGetOptions contains the optional parameters for the AssignmentsClient.Get method.
func (client *AssignmentsClient) Get(ctx context.Context, resourceScope string, assignmentName string, options *AssignmentsClientGetOptions) (AssignmentsClientGetResponse, error) {
	var err error
	const operationName = "AssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceScope, assignmentName, options)
	if err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AssignmentsClient) getCreateRequest(ctx context.Context, resourceScope string, assignmentName string, _ *AssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssignmentsClient) getHandleResponse(resp *http.Response) (AssignmentsClientGetResponse, error) {
	result := AssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List blueprint assignments within a subscription or a management group.
//
// Generated from API version 2018-11-01-preview
//   - resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: '/subscriptions/{subscriptionId}').
//   - options - AssignmentsClientListOptions contains the optional parameters for the AssignmentsClient.NewListPager method.
func (client *AssignmentsClient) NewListPager(resourceScope string, options *AssignmentsClientListOptions) *runtime.Pager[AssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssignmentsClientListResponse]{
		More: func(page AssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssignmentsClientListResponse) (AssignmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AssignmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceScope, options)
			}, nil)
			if err != nil {
				return AssignmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AssignmentsClient) listCreateRequest(ctx context.Context, resourceScope string, _ *AssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssignmentsClient) listHandleResponse(resp *http.Response) (AssignmentsClientListResponse, error) {
	result := AssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return AssignmentsClientListResponse{}, err
	}
	return result, nil
}

// WhoIsBlueprint - Get Blueprints service SPN objectId
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-11-01-preview
//   - resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format: '/subscriptions/{subscriptionId}').
//   - assignmentName - Name of the blueprint assignment.
//   - options - AssignmentsClientWhoIsBlueprintOptions contains the optional parameters for the AssignmentsClient.WhoIsBlueprint
//     method.
func (client *AssignmentsClient) WhoIsBlueprint(ctx context.Context, resourceScope string, assignmentName string, options *AssignmentsClientWhoIsBlueprintOptions) (AssignmentsClientWhoIsBlueprintResponse, error) {
	var err error
	const operationName = "AssignmentsClient.WhoIsBlueprint"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.whoIsBlueprintCreateRequest(ctx, resourceScope, assignmentName, options)
	if err != nil {
		return AssignmentsClientWhoIsBlueprintResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsClientWhoIsBlueprintResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AssignmentsClientWhoIsBlueprintResponse{}, err
	}
	resp, err := client.whoIsBlueprintHandleResponse(httpResp)
	return resp, err
}

// whoIsBlueprintCreateRequest creates the WhoIsBlueprint request.
func (client *AssignmentsClient) whoIsBlueprintCreateRequest(ctx context.Context, resourceScope string, assignmentName string, _ *AssignmentsClientWhoIsBlueprintOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/whoIsBlueprint"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if assignmentName == "" {
		return nil, errors.New("parameter assignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentName}", url.PathEscape(assignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// whoIsBlueprintHandleResponse handles the WhoIsBlueprint response.
func (client *AssignmentsClient) whoIsBlueprintHandleResponse(resp *http.Response) (AssignmentsClientWhoIsBlueprintResponse, error) {
	result := AssignmentsClientWhoIsBlueprintResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WhoIsBlueprintContract); err != nil {
		return AssignmentsClientWhoIsBlueprintResponse{}, err
	}
	return result, nil
}
