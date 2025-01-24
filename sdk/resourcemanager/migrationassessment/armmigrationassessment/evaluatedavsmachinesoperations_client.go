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

// EvaluatedAvsMachinesOperationsClient contains the methods for the EvaluatedAvsMachinesOperations group.
// Don't use this type directly, use NewEvaluatedAvsMachinesOperationsClient() instead.
type EvaluatedAvsMachinesOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEvaluatedAvsMachinesOperationsClient creates a new instance of EvaluatedAvsMachinesOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEvaluatedAvsMachinesOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EvaluatedAvsMachinesOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EvaluatedAvsMachinesOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a EvaluatedAvsMachine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-09-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - evaluatedAvsMachineName - Business case Evaluated AVS machine ARM name
//   - options - EvaluatedAvsMachinesOperationsClientGetOptions contains the optional parameters for the EvaluatedAvsMachinesOperationsClient.Get
//     method.
func (client *EvaluatedAvsMachinesOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, evaluatedAvsMachineName string, options *EvaluatedAvsMachinesOperationsClientGetOptions) (EvaluatedAvsMachinesOperationsClientGetResponse, error) {
	var err error
	const operationName = "EvaluatedAvsMachinesOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, evaluatedAvsMachineName, options)
	if err != nil {
		return EvaluatedAvsMachinesOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvaluatedAvsMachinesOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EvaluatedAvsMachinesOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EvaluatedAvsMachinesOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, evaluatedAvsMachineName string, options *EvaluatedAvsMachinesOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/evaluatedAvsMachines/{evaluatedAvsMachineName}"
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
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	if evaluatedAvsMachineName == "" {
		return nil, errors.New("parameter evaluatedAvsMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evaluatedAvsMachineName}", url.PathEscape(evaluatedAvsMachineName))
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
func (client *EvaluatedAvsMachinesOperationsClient) getHandleResponse(resp *http.Response) (EvaluatedAvsMachinesOperationsClientGetResponse, error) {
	result := EvaluatedAvsMachinesOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatedAvsMachine); err != nil {
		return EvaluatedAvsMachinesOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBusinessCasePager - List EvaluatedAvsMachine resources by BusinessCase
//
// Generated from API version 2023-09-09-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - options - EvaluatedAvsMachinesOperationsClientListByBusinessCaseOptions contains the optional parameters for the EvaluatedAvsMachinesOperationsClient.NewListByBusinessCasePager
//     method.
func (client *EvaluatedAvsMachinesOperationsClient) NewListByBusinessCasePager(resourceGroupName string, projectName string, businessCaseName string, options *EvaluatedAvsMachinesOperationsClientListByBusinessCaseOptions) *runtime.Pager[EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse]{
		More: func(page EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse) (EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EvaluatedAvsMachinesOperationsClient.NewListByBusinessCasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBusinessCaseCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, options)
			}, nil)
			if err != nil {
				return EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse{}, err
			}
			return client.listByBusinessCaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBusinessCaseCreateRequest creates the ListByBusinessCase request.
func (client *EvaluatedAvsMachinesOperationsClient) listByBusinessCaseCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, options *EvaluatedAvsMachinesOperationsClientListByBusinessCaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/evaluatedAvsMachines"
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
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
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

// listByBusinessCaseHandleResponse handles the ListByBusinessCase response.
func (client *EvaluatedAvsMachinesOperationsClient) listByBusinessCaseHandleResponse(resp *http.Response) (EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse, error) {
	result := EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatedAvsMachineListResult); err != nil {
		return EvaluatedAvsMachinesOperationsClientListByBusinessCaseResponse{}, err
	}
	return result, nil
}
