// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

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

// GenerateDetailedCostReportOperationStatusClient contains the methods for the GenerateDetailedCostReportOperationStatus
// group.
// Don't use this type directly, use NewGenerateDetailedCostReportOperationStatusClient() instead.
type GenerateDetailedCostReportOperationStatusClient struct {
	internal *arm.Client
}

// NewGenerateDetailedCostReportOperationStatusClient creates a new instance of GenerateDetailedCostReportOperationStatusClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGenerateDetailedCostReportOperationStatusClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateDetailedCostReportOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GenerateDetailedCostReportOperationStatusClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get the status of the specified operation. This link is provided in the GenerateDetailedCostReport creation request
// response header.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - operationID - The target operation Id.
//   - scope - The ARM Resource ID for subscription, resource group, billing account, or other billing scopes. For details, see
//     https://aka.ms/costmgmt/scopes.
//   - options - GenerateDetailedCostReportOperationStatusClientGetOptions contains the optional parameters for the GenerateDetailedCostReportOperationStatusClient.Get
//     method.
func (client *GenerateDetailedCostReportOperationStatusClient) Get(ctx context.Context, operationID string, scope string, options *GenerateDetailedCostReportOperationStatusClientGetOptions) (GenerateDetailedCostReportOperationStatusClientGetResponse, error) {
	var err error
	const operationName = "GenerateDetailedCostReportOperationStatusClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, operationID, scope, options)
	if err != nil {
		return GenerateDetailedCostReportOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GenerateDetailedCostReportOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GenerateDetailedCostReportOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GenerateDetailedCostReportOperationStatusClient) getCreateRequest(ctx context.Context, operationID string, scope string, _ *GenerateDetailedCostReportOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/operationStatus/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GenerateDetailedCostReportOperationStatusClient) getHandleResponse(resp *http.Response) (GenerateDetailedCostReportOperationStatusClientGetResponse, error) {
	result := GenerateDetailedCostReportOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GenerateDetailedCostReportOperationStatuses); err != nil {
		return GenerateDetailedCostReportOperationStatusClientGetResponse{}, err
	}
	return result, nil
}
