//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// GenerateDetailedCostReportOperationResultsClient contains the methods for the GenerateDetailedCostReportOperationResults group.
// Don't use this type directly, use NewGenerateDetailedCostReportOperationResultsClient() instead.
type GenerateDetailedCostReportOperationResultsClient struct {
	internal *arm.Client
}

// NewGenerateDetailedCostReportOperationResultsClient creates a new instance of GenerateDetailedCostReportOperationResultsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGenerateDetailedCostReportOperationResultsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GenerateDetailedCostReportOperationResultsClient, error) {
	cl, err := arm.NewClient(moduleName+".GenerateDetailedCostReportOperationResultsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GenerateDetailedCostReportOperationResultsClient{
		internal: cl,
	}
	return client, nil
}

// BeginGet - Gets the result of the specified operation. The link with this operationId is provided as a response header
// of the initial request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - operationID - The target operation Id.
//   - scope - The ARM Resource ID for subscription, resource group, billing account, or other billing scopes. For details, see
//     https://aka.ms/costmgmt/scopes.
//   - options - GenerateDetailedCostReportOperationResultsClientBeginGetOptions contains the optional parameters for the GenerateDetailedCostReportOperationResultsClient.BeginGet
//     method.
func (client *GenerateDetailedCostReportOperationResultsClient) BeginGet(ctx context.Context, operationID string, scope string, options *GenerateDetailedCostReportOperationResultsClientBeginGetOptions) (*runtime.Poller[GenerateDetailedCostReportOperationResultsClientGetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.get(ctx, operationID, scope, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GenerateDetailedCostReportOperationResultsClientGetResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[GenerateDetailedCostReportOperationResultsClientGetResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Get - Gets the result of the specified operation. The link with this operationId is provided as a response header of the
// initial request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *GenerateDetailedCostReportOperationResultsClient) get(ctx context.Context, operationID string, scope string, options *GenerateDetailedCostReportOperationResultsClientBeginGetOptions) (*http.Response, error) {
	req, err := client.getCreateRequest(ctx, operationID, scope, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// getCreateRequest creates the Get request.
func (client *GenerateDetailedCostReportOperationResultsClient) getCreateRequest(ctx context.Context, operationID string, scope string, options *GenerateDetailedCostReportOperationResultsClientBeginGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/operationResults/{operationId}"
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
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
