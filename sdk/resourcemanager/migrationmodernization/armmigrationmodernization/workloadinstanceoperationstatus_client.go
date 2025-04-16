// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationmodernization

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

// WorkloadInstanceOperationStatusClient contains the methods for the WorkloadInstanceOperationStatus group.
// Don't use this type directly, use NewWorkloadInstanceOperationStatusClient() instead.
type WorkloadInstanceOperationStatusClient struct {
	internal *arm.Client
}

// NewWorkloadInstanceOperationStatusClient creates a new instance of WorkloadInstanceOperationStatusClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkloadInstanceOperationStatusClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkloadInstanceOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkloadInstanceOperationStatusClient{
		internal: cl,
	}
	return client, nil
}

// Get - Tracks the results of an asynchronous operation on the workload instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - subscriptionID - Azure Subscription Id in which project was created.
//   - resourceGroupName - Name of the Azure Resource Group that project is part of.
//   - modernizeProjectName - ModernizeProject name.
//   - workloadInstanceName - Workload instance name.
//   - operationID - Operation Id.
//   - options - WorkloadInstanceOperationStatusClientGetOptions contains the optional parameters for the WorkloadInstanceOperationStatusClient.Get
//     method.
func (client *WorkloadInstanceOperationStatusClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, operationID string, options *WorkloadInstanceOperationStatusClientGetOptions) (WorkloadInstanceOperationStatusClientGetResponse, error) {
	var err error
	const operationName = "WorkloadInstanceOperationStatusClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, modernizeProjectName, workloadInstanceName, operationID, options)
	if err != nil {
		return WorkloadInstanceOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkloadInstanceOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkloadInstanceOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkloadInstanceOperationStatusClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, operationID string, _ *WorkloadInstanceOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/modernizeProjects/{modernizeProjectName}/workloadInstances/{workloadInstanceName}/operations/{operationId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if modernizeProjectName == "" {
		return nil, errors.New("parameter modernizeProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{modernizeProjectName}", url.PathEscape(modernizeProjectName))
	if workloadInstanceName == "" {
		return nil, errors.New("parameter workloadInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadInstanceName}", url.PathEscape(workloadInstanceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkloadInstanceOperationStatusClient) getHandleResponse(resp *http.Response) (WorkloadInstanceOperationStatusClientGetResponse, error) {
	result := WorkloadInstanceOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return WorkloadInstanceOperationStatusClientGetResponse{}, err
	}
	return result, nil
}
