//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleAssignmentSchedulesClient contains the methods for the RoleAssignmentSchedules group.
// Don't use this type directly, use NewRoleAssignmentSchedulesClient() instead.
type RoleAssignmentSchedulesClient struct {
	internal *arm.Client
}

// NewRoleAssignmentSchedulesClient creates a new instance of RoleAssignmentSchedulesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleAssignmentSchedulesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleAssignmentSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleAssignmentSchedulesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get the specified role assignment schedule for a resource scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01-preview
//   - scope - The scope of the role assignment schedule.
//   - roleAssignmentScheduleName - The name (guid) of the role assignment schedule to get.
//   - options - RoleAssignmentSchedulesClientGetOptions contains the optional parameters for the RoleAssignmentSchedulesClient.Get
//     method.
func (client *RoleAssignmentSchedulesClient) Get(ctx context.Context, scope string, roleAssignmentScheduleName string, options *RoleAssignmentSchedulesClientGetOptions) (RoleAssignmentSchedulesClientGetResponse, error) {
	var err error
	const operationName = "RoleAssignmentSchedulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentScheduleName, options)
	if err != nil {
		return RoleAssignmentSchedulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RoleAssignmentSchedulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentSchedulesClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentScheduleName string, options *RoleAssignmentSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentSchedules/{roleAssignmentScheduleName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentScheduleName == "" {
		return nil, errors.New("parameter roleAssignmentScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentScheduleName}", url.PathEscape(roleAssignmentScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentSchedulesClient) getHandleResponse(resp *http.Response) (RoleAssignmentSchedulesClientGetResponse, error) {
	result := RoleAssignmentSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentSchedule); err != nil {
		return RoleAssignmentSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role assignment schedules for a resource scope.
//
// Generated from API version 2024-09-01-preview
//   - scope - The scope of the role assignments schedules.
//   - options - RoleAssignmentSchedulesClientListForScopeOptions contains the optional parameters for the RoleAssignmentSchedulesClient.NewListForScopePager
//     method.
func (client *RoleAssignmentSchedulesClient) NewListForScopePager(scope string, options *RoleAssignmentSchedulesClientListForScopeOptions) *runtime.Pager[RoleAssignmentSchedulesClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleAssignmentSchedulesClientListForScopeResponse]{
		More: func(page RoleAssignmentSchedulesClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleAssignmentSchedulesClientListForScopeResponse) (RoleAssignmentSchedulesClientListForScopeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RoleAssignmentSchedulesClient.NewListForScopePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForScopeCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return RoleAssignmentSchedulesClientListForScopeResponse{}, err
			}
			return client.listForScopeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentSchedulesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentSchedulesClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignmentSchedules"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentSchedulesClient) listForScopeHandleResponse(resp *http.Response) (RoleAssignmentSchedulesClientListForScopeResponse, error) {
	result := RoleAssignmentSchedulesClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentScheduleListResult); err != nil {
		return RoleAssignmentSchedulesClientListForScopeResponse{}, err
	}
	return result, nil
}
